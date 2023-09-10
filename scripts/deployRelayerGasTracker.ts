import { ethers } from "hardhat";

async function deploy() {
  const factory = await ethers.getContractFactory("RelayerGasTracker");
  const contract = await factory.deploy();
  await contract.deployed();

  console.log("Relayer Gas tracker is deployed to : ", contract.address);
}

async function setRelayer() {
  const contract = await ethers.getContractAt(
    "RelayerGasTracker",
    "0x3e4E07926c1c4AC9f29539E385fBbF700b49F221",
  );
  const res = await contract.setRelayer(
    "0x71A713135d57911631Bb54259026Eaa030F7B881",
  );
  await res.wait().then((receipt) => {
    console.log("Finished with status : ", receipt.status);
  });
}

async function main() {
  // await deploy();
  await setRelayer();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// Relayer Gas tracker is deployed to :  0x3e4E07926c1c4AC9f29539E385fBbF700b49F221 on Donau Testnet
