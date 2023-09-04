import { ethers } from "hardhat";

async function main() {
  const factory = await ethers.getContractFactory("RelayerGasTracker");
  const contract = await factory.deploy();
  await contract.deployed();

  console.log("Relayer Gas tracker is deployed to : ", contract.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// Relayer Gas tracker is deployed to :  0xB66D7b1e294a7f19A236DdAb6371D5f0b5acD722 on Donau Testnet