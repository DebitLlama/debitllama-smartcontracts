import { ethers } from "hardhat";
import { getEncryptionPublicKey } from "@metamask/eth-sig-util";

async function main() {
  const wallet = ethers.Wallet.createRandom();

  const encryptionPublicKey = getEncryptionPublicKey(
    wallet.privateKey.substring(2),
  );
  console.log("Generated private key : ", wallet.privateKey);
  console.log("Generatex public key: ", await wallet.getAddress());
  console.log("Generated Encryption public key: ", encryptionPublicKey);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
