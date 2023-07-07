pragma circom 2.0.0;
include "../../node_modules/circomlib/circuits/poseidon.circom";

template CommitmentHasher(){
    signal input nullifier;
    signal input secret;
    signal input nonce;
    signal output commitment;
    signal output paymentIntent;

    component commitmentHasher  = Poseidon(2);

    component nullifierHasher = Poseidon(2);

    commitmentHasher.inputs[0] <== nullifier;
    commitmentHasher.inputs[1] <== secret;

    commitment <==commitmentHasher.out;

    nullifierHasher.inputs[0] <== nullifier;
    nullifierHasher.inputs[1] <== nonce;

    paymentIntent <== nullifierHasher.out;
}

template DirectDebit(){
    signal input paymentIntent;
    signal input commitmentHash;

    signal input payee;

    // The max amount that can be debited with the proof
    signal input maxDebitAmount;
    
    // The amount of times this proof can be used ti withdraw max amount
    signal input debitTimes;

    // The time that needs to pass before the proof can be used to debit an account again
    signal input debitInterval;

   // Private inputs!
    signal input secret;
    // A nonce for the nullifier so the note is reusable!
    signal input nonce;
    // The nullifier is used to calculate the secret with the commitment
    // And also the payment intent with the nonce!
    signal input nullifier;
    
    //Hidden signals to verify inputs
    signal payeeSquare;
    signal maxDebitAmountSquare;
    signal debitTimesSquare;
    signal debitIntervalSquare;

    // Hashing the commitment and the nullifier
    component commitmentHasher = CommitmentHasher();
    commitmentHasher.nullifier <== nullifier;
    commitmentHasher.secret <== secret;
    commitmentHasher.nonce <== nonce;

    commitmentHasher.paymentIntent === paymentIntent;
    commitmentHasher.commitment === commitmentHash;

    payeeSquare <== payee * payee;
    maxDebitAmountSquare <== maxDebitAmount * maxDebitAmount;
    debitTimesSquare <== debitTimes * debitTimes;
    debitIntervalSquare <==debitInterval * debitInterval;
}

component main {public [paymentIntent,commitmentHash, payee, maxDebitAmount, debitTimes,debitInterval]} = DirectDebit();