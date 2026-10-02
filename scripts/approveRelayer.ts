import { ethers } from "hardhat";

// Arbitrum Sepolia

const VERIFIER_ADDRESS = "0x26Cb79a592dC3F2F4e588Ed19A5A5314779fe80E";
const VIRTUAL_ACCOUNTS_ADDRESS = "0x5586938a2fC4489661E868c5800769Fb10847fC5";
const CONNECTED_WALLETS_ADDRESS = "0x3Cad43A3038F0E657753C0129ce7Ea4a5801EC90";

const MOCKADDRESS = "0x8c2d2a0C51f8F9476423476a79A572C46b622D6e";

async function main() {
  const [owner] = await ethers.getSigners();

  console.log("Owner:", owner.address);
  console.log("Approving relayer:", MOCKADDRESS);

  const VirtualAccountsFactory = await ethers.getContractFactory(
    "VirtualAccounts"
  );

  const virtualAccounts = VirtualAccountsFactory.attach(
    VIRTUAL_ACCOUNTS_ADDRESS
  );

  // The owner approves MOCKADDRESS as a relayer
  const tx = await virtualAccounts.approveRelayer(
    MOCKADDRESS,
    true
  );

  console.log("Transaction:", tx.hash);

  await tx.wait();

  console.log(
    `Relayer ${MOCKADDRESS} approved successfully`
  );

  // Verify
  const approved = await virtualAccounts.approvedRelayers(
    MOCKADDRESS
  );

  console.log("approvedRelayers:", approved);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});