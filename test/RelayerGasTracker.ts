import { expect } from "chai";
import { parseEther } from "ethers/lib/utils";
import { setupRelayerGasTracker } from "../test/setup";

describe("Should test relayer gas tracker", function () {
  it("Should deploy the gas tracker and top up the relayer", async function () {
    const { gasTracker, owner, alice, relayer } =
      await setupRelayerGasTracker();

    const savedRelayer = await gasTracker.relayer();
    expect(savedRelayer).to.equal(relayer.address);

    let errorOccured = false;
    let errorMessage = "";
    try {
      await gasTracker.connect(alice).setRelayer(alice.address);
    } catch (err: any) {
      errorOccured = true;
      errorMessage = err.message;
    }
    expect(errorOccured).to.be.true;

    const relayerBalance = await relayer.getBalance();
    await gasTracker.topUpRelayer({ value: parseEther("10") });

    const relayerBalanceAgain = await relayer.getBalance();

    expect(relayerBalance.add(parseEther("10"))).to.equal(relayerBalanceAgain);
    // Relayer balance added!
  });
});
