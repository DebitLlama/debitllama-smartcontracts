import { ethers } from "hardhat";

export async function setupTests() {
  const [owner, alice, bob, relayer] = await ethers.getSigners();
  const verifierFactory = await ethers.getContractFactory(
    "contracts/PaymentIntentVerifier.sol:Verifier",
  );
  const Verifier = await verifierFactory.deploy();
  await Verifier.deployed();
  const VirtualAccountsFactory = await ethers.getContractFactory(
    "VirtualAccounts",
  );
  const virtualAccounts = await VirtualAccountsFactory.deploy(Verifier.address);
  await virtualAccounts.deployed();

  const ConnectedWalletsFactory = await ethers.getContractFactory(
    "ConnectedWallets",
  );
  const connectedWallets = await ConnectedWalletsFactory.deploy(
    Verifier.address,
  );
  await connectedWallets.deployed();

  const MOCKERC20Factory = await ethers.getContractFactory("MOCKERC20");
  const MOCKERC20 = await MOCKERC20Factory.deploy();
  await MOCKERC20.deployed();

  await virtualAccounts.approveRelayer(relayer.address, true);
  await connectedWallets.approveRelayer(relayer.address, true);
  await virtualAccounts.approveRelayer(owner.address, true);
  await connectedWallets.approveRelayer(owner.address, true);


  return {
    owner,
    alice,
    bob,
    relayer,
    virtualAccounts,
    MOCKERC20,
    connectedWallets,
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
