//@ts-ignore
import { utils } from "ffjavascript";
import crypto from "crypto";
//@ts-ignore
import { groth16 } from "snarkjs";
import bigInt from "big-integer";
//@ts-ignore
import { poseidon } from "circomlibjs";

/** Generates the ZKP for the payment intent!
 */
export async function createPaymentIntent(
  { paymentIntentSecret, snarkArtifacts }: {
    paymentIntentSecret: PaymentIntentSecret;
    snarkArtifacts?: SnarkArtifacts;
  },
): Promise<FullProof> {
  const nonce = rbigint();
  const secrets = decodeAccountSecrets(paymentIntentSecret.note);
  const paymentIntent = generatePaymentIntentNullifier(
    secrets.nullifier,
    nonce,
  );
  console.log("Generate proof start");
  const input = {
    // public inputs
    paymentIntent,
    commitmentHash: secrets.commitment,
    payee: paymentIntentSecret.payee,
    maxDebitAmount: paymentIntentSecret.maxDebitAmount,
    debitTimes: paymentIntentSecret.debitTimes,
    debitInterval: paymentIntentSecret.debitInterval,
    // Private Inputs

    secret: secrets.secret,
    nonce,
    nullifier: secrets.nullifier,
  };
  console.time("Proof Time");

  if (!snarkArtifacts) {
    snarkArtifacts = {
      wasmFilePath: `circuits/directDebit/directDebit_js/directDebit.wasm`,
      zkeyFilePath: `circuits/directDebit/directDebit_0001.zkey`,
    };
  }

  const { proof, publicSignals } = await groth16.fullProve(
    input,
    snarkArtifacts.wasmFilePath,
    snarkArtifacts.zkeyFilePath,
  );
  console.timeEnd("Proof Time");

  return { proof, publicSignals };
}

/**
 * Verifies a SnarkJS proof.
 * @param verificationKey The zero-knowledge verification key.
 * @param fullProof The SnarkJS full proof.
 * @returns True if the proof is valid, false otherwise.
 */

export function verifySixPublicSignals(
  verificationKey: any,
  { proof, publicSignals }: FullProof,
): Promise<boolean> {
  return groth16.verify(
    verificationKey,
    [
      publicSignals[0],
      publicSignals[1],
      publicSignals[2],
      publicSignals[3],
      publicSignals[4],
      publicSignals[5],
    ],
    proof,
  );
}

/**
 * @returns AccountSecrets
 */
export function createAccountSecrets(
  { nullifier, secret }: { nullifier: bigint; secret: bigint },
): AccountSecrets {
  return {
    nullifier,
    secret,
    preimage: Buffer.concat([
      utils.leInt2Buff(nullifier, 31),
      utils.leInt2Buff(secret, 31),
    ]),
    commitment: generateCommitmentHash(nullifier, secret),
  };
}

export function newAccountSecrets(): string {
  const secrets = createAccountSecrets({
    nullifier: rbigint(),
    secret: rbigint(),
  });
  return toNoteHex(secrets.preimage, 62);
}

export function decodeAccountSecrets(note: string): AccountSecrets {
  const buf = Buffer.from(note);
  const nullifier = utils.leBuff2int(buf.slice(0, 31));
  const secret = utils.leBuff2int(buf.slice(31, 62));
  const secrets = createAccountSecrets({ nullifier, secret });
  return secrets;
}

export function generateCommitmentHash(
  nullifier: BigNumberish,
  secret: BigNumberish,
): bigint {
  return poseidon([BigInt(nullifier), BigInt(secret)]);
}

export function generatePaymentIntentNullifier(
  nullifier: BigNumberish,
  nonce: BigNumberish,
): bigint {
  return poseidon([BigInt(nullifier), BigInt(nonce)]);
}

export function rbigint(): bigint {
  return utils.leBuff2int(crypto.randomBytes(31));
}

/**
 * Makes a proof compatible with the Verifier.sol method inputs.
 * @param proof The proof generated with SnarkJS.
 * @returns The Solidity compatible proof.
 */
export default function packToSolidityProof(proof: Proof): SolidityProof {
  return [
    proof.pi_a[0],
    proof.pi_a[1],
    proof.pi_b[0][1],
    proof.pi_b[0][0],
    proof.pi_b[1][1],
    proof.pi_b[1][0],
    proof.pi_c[0],
    proof.pi_c[1],
  ];
}

export type AccountSecrets = {
  nullifier: bigint;
  secret: bigint;
  preimage: Buffer;
  commitment: bigint;
};

export type PaymentIntentSecret = {
  note: string;
  payee: string;
  maxDebitAmount: string;
  debitTimes: number;
  debitInterval: number;
};

export type PaymentIntentPublicSignals = {
  commitment: bigint;
  paymentIntent: bigint;
  payee: string;
  maxDebitAmount: string;
  debitTimes: number;
  debitInterval: number;
};

export type Proof = {
  pi_a: BigNumberish[];
  pi_b: BigNumberish[][];
  pi_c: BigNumberish[];
  protocol: string;
  curve: string;
};

export type SolidityProof = [
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish,
];
export type BigNumberish = string | bigint;

/** BigNumber to hex string of specified length */
export function toNoteHex(number: Buffer | any, length = 32) {
  const str = number instanceof Buffer
    ? number.toString("hex")
    : bigInt(number).toString(16);
  return "0x" + str.padStart(length * 2, "0");
}

export type SnarkArtifacts = {
  wasmFilePath: string;
  zkeyFilePath: string;
};

export type FullProof = {
  proof: Proof;
  publicSignals: Array<any>;
};
