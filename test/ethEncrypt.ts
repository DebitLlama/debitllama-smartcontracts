import { expect } from "chai";
import {
  decryptData,
  encryptData,
  packEncryptedMessage,
  unpackEncryptedMessage,
  //@ts-ignore
} from "../lib/ethencrypt.js";
import { ethers } from "ethers";
import { getEncryptionPublicKey } from "@metamask/eth-sig-util";

describe("ethEncrypt tests", function () {
  it("Should encrypt and decrypt a message", async function () {
    const message = "I am satoshi buterin";
    const privateKey = ethers.Wallet.createRandom().privateKey;
    const encryptionPublicKey = getEncryptionPublicKey(privateKey.substring(2));
    const encryptedData = encryptData(encryptionPublicKey, message);

    const packedData = packEncryptedMessage(encryptedData);
    const unpackedData = unpackEncryptedMessage(packedData);
    const decipheredtext = decryptData(privateKey, unpackedData);
    expect(message).to.equal(decipheredtext);
  });
});
