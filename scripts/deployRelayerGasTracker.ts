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
    "0x8c142b126fad0E0553aA1d4c84Ae33eA5FcBF0C5",
  );

  const res = await contract.setRelayer(
    "0xaaCb9bf503Dfb3A8a77BB5c459f45f495B7ad392",
  );
  await res.wait().then(async (receipt) => {
    console.log("Finished with status : ", receipt.status);
    const relayer = await contract.relayer();
    console.log(relayer);
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
// Relayer Gas tracker is deployed to :  0x8c142b126fad0E0553aA1d4c84Ae33eA5FcBF0C5 on BTT Mainnet
