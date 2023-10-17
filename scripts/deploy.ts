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

// DONAU TESTNET ADDRESSES: (redeployed with finalized circuit)
// Verifier contract is deployed to  0xde8E09BC372EE26Bb25A1e7325FceF5af523281C
// Virtual Accounts contract is deployed to :  0xb66adC34cb968bB67cC7656fA3Aba875a322FF35
// Connected Wallets contract is deployed to:  0x59b3202ed8C7e2459FA9c6e9592e9Dae671644A5