import { ethers } from "hardhat";

export async function setupTests() {
  const [owner, alice, bob, relayer] = await ethers.getSigners();
  const verifierFactory = await ethers.getContractFactory(
    "contracts/PaymentIntentVerifier.sol:Verifier",
  );
  const Verifier = await verifierFactory.deploy();
  await Verifier.deployed();
  const DirectDebitFactory = await ethers.getContractFactory("DirectDebit");
  const directDebit = await DirectDebitFactory.deploy(Verifier.address);
  await directDebit.deployed();

  const MOCKERC20Factory = await ethers.getContractFactory("MOCKERC20");
  const MOCKERC20 = await MOCKERC20Factory.deploy();
  await MOCKERC20.deployed();

  return {
    owner,
    alice,
    bob,
    relayer,
    directDebit,
    MOCKERC20,
  };
}

export async function setupRelayerGasTracker() {
  const [owner, alice, relayer] = await ethers.getSigners();
  const factory = await ethers.getContractFactory("RelayerGasTracker");
  const gasTracker = await factory.deploy();
  await gasTracker.deployed();
  await gasTracker.setRelayer(relayer.address);
  return { gasTracker, owner, alice, relayer };
}
