const { utils } = require("ffjavascript");
const crypto = require("crypto");
const { groth16 } = require("snarkjs");
const bigInt = require("big-integer");
const { poseidon } = require("circomlibjs");
const Buffer = require('buffer/').Buffer;

const { encryptData, decryptData, packEncryptedMessage, unpackEncryptedMessage } = require("./ethencrypt");


/**
 * THIS IS THE JS VERSION OF THE LIB TO BE BUILT INTO A BROWSER IMPORTABLE DEPENDENCY!
 * 
 */

/** Generates the ZKP for the payment intent!
 */
async function createPaymentIntent(
    { paymentIntentSecret, snarkArtifacts }
) {
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

function verifySixPublicSignals(
    verificationKey,
    { proof, publicSignals },
) {
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
function createAccountSecrets(
    { nullifier, secret },
) {
    return {
        nullifier,
        secret,
        preimage: Buffer.concat([
            Buffer.from(utils.leInt2Buff(nullifier, 31).buffer),
            Buffer.from(utils.leInt2Buff(secret, 31).buffer),
        ]),
        commitment: generateCommitmentHash(nullifier, secret),
    };
}

function newAccountSecrets() {
    const secrets = createAccountSecrets({
        nullifier: rbigint(),
        secret: rbigint(),
    });
    return toNoteHex(secrets.preimage, 62);
}

function decodeAccountSecrets(note) {
    const buf = Buffer.from(note);
    const nullifier = utils.leBuff2int(buf.slice(0, 31));
    const secret = utils.leBuff2int(buf.slice(31, 62));
    const secrets = createAccountSecrets({ nullifier, secret });
    return secrets;
}

function generateCommitmentHash(
    nullifier,
    secret,
) {
    return poseidon([BigInt(nullifier), BigInt(secret)]);
}

function generatePaymentIntentNullifier(
    nullifier,
    nonce,
) {
    return poseidon([BigInt(nullifier), BigInt(nonce)]);
}

function rbigint() {
    return utils.leBuff2int(getRandomBytes(31));
}
function getRandomBytes(n) {
    let array = new Uint8Array(n);
    if (typeof globalThis.crypto !== "undefined") { // Supported
        globalThis.crypto.getRandomValues(array);
    } else { // NodeJS
        crypto.randomFillSync(array);
    }
    return array;
}
/**
 * Makes a proof compatible with the Verifier.sol method inputs.
 * @param proof The proof generated with SnarkJS.
 * @returns The Solidity compatible proof.
 */
function packToSolidityProof(proof) {
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



/** BigNumber to hex string of specified length */
function toNoteHex(number, length = 32) {
    const str = number instanceof Buffer
        ? number.toString("hex")
        : bigInt(number).toString(16);
    return "0x" + str.padStart(length * 2, "0");
}

module.exports = {
    createPaymentIntent,
    verifySixPublicSignals,
    createAccountSecrets,
    newAccountSecrets,
    decodeAccountSecrets,
    generateCommitmentHash,
    generatePaymentIntentNullifier,
    rbigint,
    getRandomBytes,
    packToSolidityProof,
    toNoteHex, Buffer,
    utils,
    encryptData, 
    decryptData, 
    packEncryptedMessage, 
    unpackEncryptedMessage
}