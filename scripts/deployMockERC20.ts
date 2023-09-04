import { ethers } from "hardhat";

async function main() {
  const factory = await ethers.getContractFactory("MOCKERC20");
  const contract = await factory.deploy();
  await contract.deployed();

  console.log("USDTM is deployed to : ", contract.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// USDTM is deployed to :  0x4420a4415033bd22393d3A918EF8d2c9c62efD99 on Donau testnet
