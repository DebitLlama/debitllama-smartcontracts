import { expect } from "chai";
import { setupTests } from "./setup";
import { parseEther } from "ethers/lib/utils";

import {
  createPaymentIntent,
  decodeAccountSecrets,
  newAccountSecrets,
  packToSolidityProof,
  toNoteHex,
} from "../lib/directDebit";

describe("Test Connected Wallets", function () {
  // Most of the direct debit functionality is inherited and tested already with VirtualAccounts
  // So I only test connected wallets here
  it("Should connect a wallet,create a payment intent and withdraw it", async function () {
    const { owner, alice, bob, relayer, connectedWallets, MOCKERC20 } =
      await setupTests();

    await MOCKERC20.mint(alice.address, parseEther("1000"));

    // Alice approves the ConnectedWallet to spend her allowance!
    // This lets the contract directly debit her account!
    await MOCKERC20.connect(alice).approve(
      connectedWallets.address,
      parseEther("1000"),
    );

    //Gonna create a new account
    // Gonna deposit tokens and create a token account
    const tokenAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(tokenAccountNote);
    const tokenAccountCommitment = toNoteHex(accountSecrets.commitment);
    // The note here is not encrypted. That will be done on the client!

    await connectedWallets.connect(alice).connectWallet(
      tokenAccountCommitment,
      MOCKERC20.address,
      tokenAccountNote,
    );

    const accountCounter = await connectedWallets.accountCounter(alice.address);
    expect(accountCounter).to.equal(1);
    const commitment = await connectedWallets.commitments(
      alice.address,
      accountCounter.sub(1),
    );
    expect(commitment).to.equal(tokenAccountCommitment);

    let tokenAccountData = await connectedWallets.accounts(commitment);
    const savedNote = await connectedWallets.encryptedNotes(commitment);
    expect(savedNote).to.equal(tokenAccountNote);
    expect(tokenAccountData.active).to.equal(true);
    expect(tokenAccountData.creator).to.equal(alice.address);
    expect(tokenAccountData.token).to.equal(MOCKERC20.address);
    expect(tokenAccountData.balance).to.equal(parseEther("0")); // There is no balance tracking for connected wallets

    // testing  the getAccount view function!
    const getAccountRes = await connectedWallets.getAccount(commitment);
    expect(tokenAccountData.active).to.equal(getAccountRes.active);
    expect(tokenAccountData.creator).to.equal(getAccountRes.creator);
    expect(tokenAccountData.token).to.equal(getAccountRes.token);
    expect(tokenAccountData.balance).to.not.equal(getAccountRes.balance); // There is no balance tracking for connected wallets
    expect(getAccountRes.balance).to.equal(parseEther("1000"));

    // now create the payment intent
    const paymentIntent = await createPaymentIntent({
      paymentIntentSecret: {
        note: tokenAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("10").toString(),
        debitTimes: 1,
        debitInterval: 0,
      },
    });

    // Now the relayer will process the direct debit
    await connectedWallets.connect(relayer).directdebit(
      packToSolidityProof(paymentIntent.proof),
      [
        toNoteHex(paymentIntent.publicSignals[0]),
        toNoteHex(paymentIntent.publicSignals[1]),
      ],
      bob.address,
      [parseEther("10"), 1, 0, parseEther("5")], // Max 10 allowed, we debit only 5
    );
    const fees = await connectedWallets.calculateFee(parseEther("5"));

    expect(fees[0]).to.equal(parseEther("0.05"));
    expect(fees[1]).to.equal(parseEther("4.95"));

    expect(await MOCKERC20.balanceOf(bob.address)).to.equal(parseEther("4.95"));

    const paymentIntentHistory = await connectedWallets.paymentIntents(
      toNoteHex(paymentIntent.publicSignals[0]),
    );
    expect(paymentIntentHistory.isNullified).to.equal(false);

    expect(paymentIntentHistory.withdrawalCount).to.equal(1);
    // Testing the getAccount function balance calculation again!
    const getAccountResAgain = await connectedWallets.getAccount(commitment);

    expect(getAccountResAgain.balance).to.equal(parseEther("995"));

    // Now I disconnect the wallet!

    // I try to disconnect somebody else's wallet first
    let errorOccured = false;
    let errorMessage = "";
    try {
      await connectedWallets.disconnectWallet(
        toNoteHex(paymentIntent.publicSignals[1]),
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("OnlyAccountOwner")).to.be.true;

    await connectedWallets.connect(alice).disconnectWallet(
      toNoteHex(paymentIntent.publicSignals[1]),
    );
    // Now the account should be closed!

    let acc = await connectedWallets.accounts(commitment);
    expect(acc.active).to.be.false;
  });
});
