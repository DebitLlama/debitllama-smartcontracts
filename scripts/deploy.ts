import { ethers } from "hardhat";

async function main() {
  const verifierFactory = await ethers.getContractFactory(
    "contracts/PaymentIntentVerifier.sol:Verifier",
  );
  const Verifier = await verifierFactory.deploy();
  await Verifier.deployed();
  const DirectDebitFactory = await ethers.getContractFactory("DirectDebit");
  const directDebit = await DirectDebitFactory.deploy(Verifier.address);
  await directDebit.deployed();

  console.log("Direct Debit is deployed to : ", directDebit.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// Direct Debit is deployed to :  0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 on Donau testnet
