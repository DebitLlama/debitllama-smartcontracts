import { ethers } from "hardhat";

//This will do a full deploy!

async function main() {
  const verifierFactory = await ethers.getContractFactory(
    "contracts/PaymentIntentVerifier.sol:Verifier",
  );
  const Verifier = await verifierFactory.deploy();
  await Verifier.deployed().then(async () => {
    console.log("Verifier contract is deployed to ", Verifier.address);

    const VirtualAccountsFactory = await ethers.getContractFactory(
      "VirtualAccounts",
    );

    const virtualAccounts = await VirtualAccountsFactory.deploy(
      Verifier.address,
    );

    await virtualAccounts.deployed().then(async () => {
      console.log(
        "Virtual Accounts contract is deployed to : ",
        virtualAccounts.address,
      );

      const ConnectedWalletsFactory = await ethers.getContractFactory(
        "ConnectedWallets",
      );
      const connectedWallets = await ConnectedWalletsFactory.deploy(
        "0x22c025aa2009DfAbbc10F5262512A999D2a73E0d",
      );
      await connectedWallets.deployed();

      console.log(
        "Connected Wallets contract is deployed to: ",
        connectedWallets.address,
      );
    });
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// DONAU TESTNET ADDRESSES: (redeployed with new connected wallet and direct debit features)
// Verifier contract is deployed to  0xA0c953Db12f02e0E8f41EFd5Ea857259a694069d
// Virtual Accounts contract is deployed to :  0x2137F4096365bCA1457E945838e0d7EC1925A973
// Connected Wallets contract is deployed to:  0xc65DDA2E81dB71C998D08A525D70dFA844BF5D3e
