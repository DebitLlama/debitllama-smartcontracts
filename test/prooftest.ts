import { expect } from "chai";
import fs from "fs";
import {
  createPaymentIntent,
  decodeAccountSecrets,
  newAccountSecrets,
  toNoteHex,
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
    const res2 = await verifySixPublicSignals(verificationKey, {
      proof,
      publicSignals: [
        publicSignals[0],
        publicSignals[1],
        publicSignals[2],
        publicSignals[3],
        publicSignals[4],
        "11",
      ],
    });
    expect(res2).to.be.false;

    const origin = "16000534653676138996713327308904050705310694920288652176050461465752592825931";
    const want ="0x235ffb4f845bec7dcb1fffbd82391cecbc278bca9187fbcd7cddf854fed5be4b"
    const got = toNoteHex(origin);
    expect(want).to.equal(got);


  });
});

