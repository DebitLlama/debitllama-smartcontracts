import { ethers } from "hardhat";

export enum VirtualAccountsContractAddress {
  BTT_TESTNET = "0xF75515Df5AC843a8B261E232bB890dc2F75A4066",
  BTT_MAINNET = "0xc4Cf42D5a6F4F061cf5F98d0338FC5913b6fF581",
}

export enum ConnectedWalletsContractAddress {
  BTT_TESTNET = "0x9c85da9E45126Fd45BC62656026A2E7226bba239",
  BTT_MAINNET = "0xF9962f3C23De4e864E56ef29125D460c785905c6",
}

//ONLY OWNER
async function main() {
  const directDebitFactory = await ethers.getContractFactory("VirtualAccounts");
  //Virtual accounts on btt testnet
  const contract = await directDebitFactory.attach(
    VirtualAccountsContractAddress.BTT_MAINNET,
  );
  await contract.updateFee("20").then((res) => {
    console.log("finished with res", res);
  }).catch((err) => {
    console.log("error occured");
    console.error(err);
  });
  //Increase the fee to 5% flat rate with a  20 divider value
  //The fee divider is used in arithmetic with the paid amount to calculate the fee.
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
