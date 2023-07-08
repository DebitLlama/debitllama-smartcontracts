import { expect } from "chai";
import { ethers } from "hardhat";
import packToSolidityProof, {
  createPaymentIntent,
  decodeAccountSecrets,
  newAccountSecrets,
  toNoteHex,
} from "../lib/directDebit";
import { ZEROADDRESS } from "./prooftest";
import { setupTests } from "./setup";
import { parseEther } from "ethers/lib/utils";
//TESTS ARE USING UNENCRYPTED CRYPTO NOTES ON CHAIN BECAUSE METAMASK IS NOT AVAILABLE
// THE CLIENT DEPENDS ON eth_getEncryptionPublicKey and eth_decrypt!!

describe("It should deploy a contract and test it", function () {
  it("Should deposit ETH, create payment intent and withdraw", async function () {
    const { owner, alice, bob, relayer, directDebit, MOCKERC20 } =
      await setupTests();

    // Gonna deposit ETH and create an account
    const ethAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(ethAccountNote);
    const ethAccountCommitment = toNoteHex(accountSecrets.commitment);
    // The note here is not encrypted! That will be done on the client!
    await directDebit.connect(alice).depositEth(
      ethAccountCommitment,
      parseEther("10"),
      ethAccountNote,
      { value: parseEther("10") },
    );

    // I expect there is an eth account open for the commitment and an encrypted note for the commitment

    const accountCounter = await directDebit.accountCounter(alice.address);
    expect(accountCounter).to.equal(1);
    const commitment = await directDebit.commitments(
      alice.address,
      accountCounter.sub(1),
    );
    expect(commitment).to.equal(ethAccountCommitment);

    let ethAccountData = await directDebit.accounts(commitment);
    const savedNote = await directDebit.encryptedNotes(commitment);

    expect(savedNote).to.equal(ethAccountNote);
    expect(ethAccountData.active).to.equal(true);
    expect(ethAccountData.creator).to.equal(alice.address);
    expect(ethAccountData.token).to.equal(ZEROADDRESS);
    expect(ethAccountData.balance).to.equal(parseEther("10"));

    // Now I top it up with 1 more ETH
    await directDebit.connect(alice).topUpETH(
      commitment,
      parseEther("10"),
      { value: parseEther("10") },
    );
    ethAccountData = await directDebit.accounts(commitment);
    expect(ethAccountData.balance).to.equal(parseEther("20"));

    // Now I create a payment intent for 1 ETH and the relayer will debit that

    const paymentIntent = await createPaymentIntent({
      paymentIntentSecret: {
        note: ethAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("1").toString(),
        debitTimes: 2,
        debitInterval: 0,
      },
    });
    const bobBalanceBefore = await bob.getBalance();

    // Now the relayer will send this tx
    await directDebit.connect(relayer).directdebit(
      packToSolidityProof(paymentIntent.proof),
      [
        toNoteHex(paymentIntent.publicSignals[0]),
        toNoteHex(paymentIntent.publicSignals[1]),
      ],
      bob.address,
      [parseEther("1"), 2, 0, parseEther("1")],
    );

    const bobBalanceAfter = await bob.getBalance();

    const fees = await directDebit.calculateFee(parseEther("1"));

    //so there are 0.5% fees that I get by dividing the amount with 200
    expect(fees[0]).to.equal(parseEther("0.005"));
    expect(fees[1]).to.equal(parseEther("0.005"));
    expect(fees[2]).to.equal(parseEther("0.99"));

    expect(bobBalanceAfter).to.equal(
      bobBalanceBefore.add(fees[2]),
    );
    ethAccountData = await directDebit.accounts(commitment);
    expect(ethAccountData.balance).to.equal(parseEther("19"));

    const paymentIntentHistory = await directDebit.paymentIntents(
      toNoteHex(paymentIntent.publicSignals[0]),
    );

    expect(paymentIntentHistory.isNullified).to.equal(false);
    expect(paymentIntentHistory.withdrawalCount).to.equal(1);

    // Now I withdraw the contents of the account and close it!

    let aliceBalanceBefore = await alice.getBalance();

    const withdrawTx = await directDebit.connect(alice).withdraw(
      ethAccountCommitment,
    );
    let aliceBalanceAfter = await alice.getBalance();

    const withdrawReceipt = await withdrawTx.wait();

    // I check if the correct amount was withdrawn by calculating the gas used and checking it.

    const expectedTxGas = parseEther("19").sub(
      aliceBalanceAfter.sub(aliceBalanceBefore),
    );
    const actualGasUsed = withdrawReceipt.gasUsed.mul(
      withdrawReceipt.effectiveGasPrice,
    );

    expect(actualGasUsed).to.equal(expectedTxGas);

    ethAccountData = await directDebit.accounts(commitment);
    expect(ethAccountData.active).to.equal(false);

    let errorOccured = false;
    let errorMessage = "";
    try {
      await directDebit.connect(relayer).directdebit(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("1"), 2, 0, parseEther("1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("InactiveAccount")).to.be.true;
  });

  it("Should deposit an ERC20, top it up, direct debit and withdraw", async function () {
    const { owner, alice, bob, relayer, directDebit, MOCKERC20 } =
      await setupTests();

    await MOCKERC20.mint(alice.address, parseEther("1000"));
    await MOCKERC20.connect(alice).approve(
      directDebit.address,
      parseEther("1000"),
    );

    expect(await MOCKERC20.balanceOf(alice.address)).to.equal(
      parseEther("1000"),
    );

    // Gonna deposit tokens and create a token account
    const tokenAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(tokenAccountNote);
    const tokenAccountCommitment = toNoteHex(accountSecrets.commitment);
    // The note here is not encrypted. That will be done on the clinent

    await directDebit.connect(alice).depositToken(
      tokenAccountCommitment,
      parseEther("100"),
      MOCKERC20.address,
      tokenAccountNote,
    );
    let aliceTokenBalance = await MOCKERC20.balanceOf(alice.address);
    expect(aliceTokenBalance).to.equal(parseEther("900"));

    const accountCounter = await directDebit.accountCounter(alice.address);
    expect(accountCounter).to.equal(1);
    const commitment = await directDebit.commitments(
      alice.address,
      accountCounter.sub(1),
    );
    expect(commitment).to.equal(tokenAccountCommitment);

    let tokenAccountData = await directDebit.accounts(commitment);
    const savedNote = await directDebit.encryptedNotes(commitment);
    expect(savedNote).to.equal(tokenAccountNote);
    expect(tokenAccountData.active).to.equal(true);
    expect(tokenAccountData.creator).to.equal(alice.address);
    expect(tokenAccountData.token).to.equal(MOCKERC20.address);
    expect(tokenAccountData.balance).to.equal(parseEther("100"));

    // Now top it up

    await directDebit.connect(alice).topUpTokens(
      tokenAccountCommitment,
      parseEther("10"),
    );

    tokenAccountData = await directDebit.accounts(commitment);
    expect(tokenAccountData.balance).to.equal(parseEther("110"));

    // and create a payment intent
    const paymentIntent = await createPaymentIntent({
      paymentIntentSecret: {
        note: tokenAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("10").toString(),
        debitTimes: 1,
        debitInterval: 0,
      },
    });

    expect(await MOCKERC20.balanceOf(bob.address)).to.equal(parseEther("0"));

    await directDebit.connect(alice).directdebit(
      packToSolidityProof(paymentIntent.proof),
      [
        toNoteHex(paymentIntent.publicSignals[0]),
        toNoteHex(paymentIntent.publicSignals[1]),
      ],
      bob.address,
      [parseEther("10"), 1, 0, parseEther("5")], // Max 10 allowed, we debit 5
    );

    const fees = await directDebit.calculateFee(parseEther("5"));

    expect(fees[0]).to.equal(parseEther("0.025"));
    expect(fees[1]).to.equal(parseEther("0.025"));
    expect(fees[2]).to.equal(parseEther("4.95"));

    expect(await MOCKERC20.balanceOf(bob.address)).to.equal(parseEther("4.95"));

    const paymentIntentHistory = await directDebit.paymentIntents(
      toNoteHex(paymentIntent.publicSignals[0]),
    );
    expect(paymentIntentHistory.isNullified).to.equal(false);

    expect(paymentIntentHistory.withdrawalCount).to.equal(1);

    // Now If I try to withdraw again it should error because the debit times was 1...

    let errorOccured = false;
    let errorMessage = "";
    try {
      await directDebit.connect(alice).directdebit(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("10"), 1, 0, parseEther("5")], // Max 10 allowed, we debit 5
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("PaymentIntentExpired")).to.be.true;
  });

  it("Deposit ETH and Top up ETH, Closed Account Errors", async function () {
    const { owner, alice, bob, relayer, directDebit, MOCKERC20 } =
      await setupTests();

    let errorOccured = false;
    let errorMessage = "";

    // Tests errors of depositETH

    // try to deposit with zero

    const ethAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(ethAccountNote);
    const ethAccountCommitment = toNoteHex(accountSecrets.commitment);
    // The note here is not encrypted! That will be done on the client!

    try {
      await directDebit.connect(alice).depositEth(
        ethAccountCommitment,
        parseEther("0"),
        ethAccountNote,
        { value: parseEther("0") },
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("ZeroTopup"));

    errorOccured = false;
    errorMessage = "";
    // try to deposit with not enough value
    try {
      await directDebit.connect(alice).depositEth(
        ethAccountCommitment,
        parseEther("10"),
        ethAccountNote,
        { value: parseEther("4") },
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("NotEnoughValue"));

    // I create an account to do further tests
    await directDebit.connect(alice).depositEth(
      ethAccountCommitment,
      parseEther("10"),
      ethAccountNote,
      { value: parseEther("10") },
    );

    //Try to top it up as a token account
    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).topUpTokens(
        ethAccountCommitment,
        parseEther("10"),
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("NotTokenAccount")).to.be.true;

    // Try to create the account again
    errorOccured = false;
    errorMessage = "";

    try {
      await directDebit.connect(alice).depositEth(
        ethAccountCommitment,
        parseEther("10"),
        ethAccountNote,
        { value: parseEther("10") },
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("AccountAlreadyActive()"));

    errorOccured = false;
    errorMessage = "";
    // Now I withdraw and inactivate the account
    try {
      await directDebit.withdraw(ethAccountCommitment);
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("OnlyAccountOwner()")).to.be.true;
    // withdraw success
    await directDebit.connect(alice).withdraw(ethAccountCommitment);

    // Try to create an account that already existed before but now inactive

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).depositEth(
        ethAccountCommitment,
        parseEther("10"),
        ethAccountNote,
        { value: parseEther("10") },
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("AccountAlreadyExists()"));

    // Now I try to withdraw from an inactive account!
    errorOccured = false;
    errorMessage = "";
    // Now I withdraw and inactivate the account
    try {
      await directDebit.withdraw(ethAccountCommitment);
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("InactiveAccount()")).to.be.true;
  });

  it("Deposit Token and top up Tokens errors", async function () {
    const { owner, alice, bob, relayer, directDebit, MOCKERC20 } =
      await setupTests();

    let errorOccured = false;
    let errorMessage = "";

    const tokenAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(tokenAccountNote);
    const tokenAccountCommitment = toNoteHex(accountSecrets.commitment);

    await MOCKERC20.mint(alice.address, parseEther("100"));
    await MOCKERC20.connect(alice).approve(
      directDebit.address,
      parseEther("10"),
    );

    // Test zero topup
    try {
      await directDebit.connect(alice).depositToken(
        tokenAccountCommitment,
        parseEther("0"),
        MOCKERC20.address,
        tokenAccountNote,
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("ZeroTopup")).to.be.true;

    // Now I create it to test already active errors
    await directDebit.connect(alice).depositToken(
      tokenAccountCommitment,
      parseEther("1"),
      MOCKERC20.address,
      tokenAccountNote,
    );

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).depositToken(
        tokenAccountCommitment,
        parseEther("1"),
        MOCKERC20.address,
        tokenAccountNote,
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("AccountAlreadyActive")).to.be.true;

    //Try to top up eth to token account
    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).topUpETH(
        tokenAccountCommitment,
        parseEther("1"),
        { value: parseEther("1") },
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("NotEthAccount")).to.be.true;

    // Now I'm gonna close this account and test the remaining errors

    await directDebit.connect(alice).withdraw(tokenAccountCommitment);
    // Try to recreate account

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).depositToken(
        tokenAccountCommitment,
        parseEther("1"),
        MOCKERC20.address,
        tokenAccountNote,
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("AccountAlreadyExists")).to.be.true;

    // Try to top up closed account
    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.connect(alice).topUpTokens(
        tokenAccountCommitment,
        parseEther("1"),
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("InactiveAccount")).to.be.true;
  });

  it("Direct debit errors tests, cancel payment intent", async function () {
    const { owner, alice, bob, relayer, directDebit, MOCKERC20 } =
      await setupTests();

    // I test direct debit with an ETH Account
    const ethAccountNote = newAccountSecrets();
    const accountSecrets = decodeAccountSecrets(ethAccountNote);
    const ethAccountCommitment = toNoteHex(accountSecrets.commitment);

    await directDebit.connect(alice).depositEth(
      ethAccountCommitment,
      parseEther("10"),
      ethAccountNote,
      { value: parseEther("10") },
    );

    const paymentIntent = await createPaymentIntent({
      paymentIntentSecret: {
        note: ethAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("1").toString(),
        debitTimes: 2,
        debitInterval: 0, // the interval will remain zero as it's hard to test passing time with unit tests
      },
    });

    // The relayer will direct debit that eth account, I need to test the verifyPaymentIntent function

    let errorOccured = false;
    let errorMessage = "";
    try {
      // InvalidProof

      await directDebit.directdebit(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("100"), 1, 0, parseEther("10")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("InvalidProof")).to.be.true;
    // PaymentNotAuthorized

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("10"), 2, 0, parseEther("100")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("PaymentNotAuthorized")).to.be.true;

    // Cancel payment intent OnlyRelatedPartiesCanCancel
    errorOccured = false;
    errorMessage = "";

    try {
      await directDebit.cancelPaymentIntent(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("1"), 2, 0, parseEther("1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("OnlyRelatedPartiesCanCancel"));

    await directDebit.connect(alice).cancelPaymentIntent(
      packToSolidityProof(paymentIntent.proof),
      [
        toNoteHex(paymentIntent.publicSignals[0]),
        toNoteHex(paymentIntent.publicSignals[1]),
      ],
      bob.address,
      [parseEther("1"), 2, 0, parseEther("1")],
    );

    // PaymentIntentNullified

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent.proof),
        [
          toNoteHex(paymentIntent.publicSignals[0]),
          toNoteHex(paymentIntent.publicSignals[1]),
        ],
        bob.address,
        [parseEther("1"), 2, 0, parseEther("1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }

    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("PaymentIntentNullified")).to.be.true;

    // PaymentIntentExpired
    const paymentIntent2 = await createPaymentIntent({
      paymentIntentSecret: {
        note: ethAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("10").toString(),
        debitTimes: 1,
        debitInterval: 0,
      },
    });

    await directDebit.directdebit(
      packToSolidityProof(paymentIntent2.proof),
      [
        toNoteHex(paymentIntent2.publicSignals[0]),
        toNoteHex(paymentIntent2.publicSignals[1]),
      ],
      bob.address,
      [parseEther("10"), 1, 0, parseEther("1")],
    );

    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent2.proof),
        [
          toNoteHex(paymentIntent2.publicSignals[0]),
          toNoteHex(paymentIntent2.publicSignals[1]),
        ],
        bob.address,
        [parseEther("10"), 1, 0, parseEther("1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("PaymentIntentExpired")).to.be.true;

    // NotEnoughAccountBalance
    const paymentIntent3 = await createPaymentIntent({
      paymentIntentSecret: {
        note: ethAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("100").toString(),
        debitTimes: 1,
        debitInterval: 0,
      },
    });

    errorMessage = "";
    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent3.proof),
        [
          toNoteHex(paymentIntent3.publicSignals[0]),
          toNoteHex(paymentIntent3.publicSignals[1]),
        ],
        bob.address,
        [parseEther("100"), 1, 0, parseEther("100")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("NotEnoughAccountBalance")).to.be.true;

    // EarlyPaymentNotAllowed

    const paymentIntent4 = await createPaymentIntent({
      paymentIntentSecret: {
        note: ethAccountNote,
        payee: bob.address,
        maxDebitAmount: parseEther("0.1").toString(),
        debitTimes: 10,
        debitInterval: 10000,
      },
    });

    await directDebit.directdebit(
      packToSolidityProof(paymentIntent4.proof),
      [
        toNoteHex(paymentIntent4.publicSignals[0]),
        toNoteHex(paymentIntent4.publicSignals[1]),
      ],
      bob.address,
      [parseEther("0.1"), 10, 10000, parseEther("0.1")],
    );
    errorOccured = false;
    errorMessage = "";

    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent4.proof),
        [
          toNoteHex(paymentIntent4.publicSignals[0]),
          toNoteHex(paymentIntent4.publicSignals[1]),
        ],
        bob.address,
        [parseEther("0.1"), 10, 10000, parseEther("0.1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("EarlyPaymentNotAllowed")).to.be.true;

    // InactiveAccount

    await directDebit.connect(alice).withdraw(ethAccountCommitment);
    errorOccured = false;
    errorMessage = "";
    try {
      await directDebit.directdebit(
        packToSolidityProof(paymentIntent4.proof),
        [
          toNoteHex(paymentIntent4.publicSignals[0]),
          toNoteHex(paymentIntent4.publicSignals[1]),
        ],
        bob.address,
        [parseEther("0.1"), 10, 10000, parseEther("0.1")],
      );
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;
    expect(errorMessage.includes("InactiveAccount")).to.be.true;
  });
});
