import { expect } from "chai";
import fs from "fs";
import {
  createPaymentIntent,
  decodeAccountSecrets,
  newAccountSecrets,
  verifySixPublicSignals,
} from "../lib/directDebit";
export const ZEROADDRESS = "0x0000000000000000000000000000000000000000";

describe("Test the circom circuit", function () {
  it("Should create a account secret, then create a payment intent and verify it", async function () {
    const note = newAccountSecrets();
    const secrets = decodeAccountSecrets(note);

    const paymentIntent = await createPaymentIntent({
      paymentIntentSecret: {
        note,
        payee: ZEROADDRESS,
        maxDebitAmount: "10",
        debitTimes: 1,
        debitInterval: 10000,
      },
    });
    const { proof, publicSignals } = paymentIntent;

    const verificationKeyFile = fs.readFileSync(
      "circuits/directDebit/verification_key.json",
      "utf-8",
    );
    const verificationKey = JSON.parse(verificationKeyFile);
    const res = await verifySixPublicSignals(verificationKey, {
      proof,
      publicSignals,
    });
    expect(res).to.be.true;
  });
});
