rm -rf directDebit_js
rm -f verification_key.json
rm -f directDebit_0000.zkey
rm -f directDebit_0001.zkey
rm -f directDebit.r1cs
rm -f directDebit.sym


circom directDebit.circom --r1cs --wasm --sym

# get the .zkey
snarkjs groth16 setup directDebit.r1cs ../powersOfTau28_hez_final_15.ptau directDebit_0000.zkey

# Contribute to the phase 2 ceremony, Add your name if you are not me XD
snarkjs zkey contribute directDebit_0000.zkey directDebit_0001.zkey --name="StrawberryChocolateFudge"

#export verification key
snarkjs zkey export verificationkey directDebit_0001.zkey verification_key.json

#generate the verifier.sol
snarkjs zkey export solidityverifier directDebit_0001.zkey PaymentIntentVerifier.sol

# copy the PaymentIntentVerifier.sol to the contracts directory
mv PaymentIntentVerifier.sol ../../contracts/PaymentIntentVerifier.sol

cp directDebit_0001.zkey ../../dist/directDebit_0001.zkey

cp directDebit_js/directDebit.wasm ../../dist/directDebit.wasm