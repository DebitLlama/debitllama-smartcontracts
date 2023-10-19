# Direct Debit Using Note Accounts

                                                                                                                                        
This repository contaisn the implementation of a direct debit system that allows merchants to pull payments from accounts using zksnarks created by customers off-chain. The core of the direct debit implementation can be found in DirectDebit.sol, while the child contracts implement different account types.

## How it works: Circuits

There is a DirectDebit() template in the directDebit.circom file.
This contains **public inputs paymentIntent, commitmentHash, payee, maxDebitAmount, debitTimes, debitInterval** and private inputs secret, nonce and nullifier.

When creating an account we calculate the hash of the secret with the nullifier.
The commitment is a poseidon hash of (nullifier,secret)

When creating a payment intent, the payment intent variable itself is a nullifier, we make the account secrets reusable by hashing the secret nullifier with a nonce

PaymentIntent = hash(nullifier,nonce)

The nonce is random.

**To verify other public inputs** like payee, maxDebitAmount,debitTimes and debitInterval **we use Hidden Signals** in the circuit. These signals are used to make sure that these parameters of the payment intent cannot be altered.

## How it works: Smart Contract

The **DirectDebit.sol** smart contract contains the implementation for the basic direct debit payments.
This is an abstract contract that needs to be implemented to create different debit accounts.
Currently Virtual Accounts and Connected Wallets are implemented.

**Virtual Accounts** are smart contract accounts that store the value inside the contract. This is currently the only way to implement direct debit for ETH, but supports ERC-20 tokens also.

The **Connected Wallets** contract supports external connected wallets that allow spending ERC-20 tokens on behalf of the account owner. This account type supports only ERC-20 tokens, but it works with external cold wallets.


## How it's used?

Users create Accounts using the computed commitment.

Payment Intents are created off-chain, passed to the merchant or relayer and then used to pull payments from the account at a specified time.

### Deployment:
The verifier contract must be deployed first, then the constructor of the Direct Debit contract uses the verifier contract address and sets the address of the owner of the contract.

### API

**Direct Debit**

The underying direct debit contract exposes the following external functions:

`calculateFee` Calculate the fee of the transaction

`directdebit` Debit the account using the zksnark proof, and it's public inputs. 
        You can find more information about it in the source code!

`cancelPaymentIntent` The account owner or the payee can cancel the payment intent so it's not usable anymore for future payments.

`updateFee` The owner of the contract can update how the fee is calculated

`togglePause` The owner of the contract can halt all pull payments from the contract as a safety mechanism, suspicious transactions can be automaticly detected.

`getAccount` A virtual view function to get the account. This is used instead of directly accessing the mapping. The child contracts override it use it to get account balance. When using connectedWallets the balance is calculated using the erc20 allowance and balance and not the stored value.

**Virtual Accounts**

`depositEth` is used to create an account. The user computed the commitment off-chain and deposits value. The secret and the nullifier must be encrypted first.

The *crypto note* (an encoded format of the secret) is essentially saved on the chain so the users don't have to download it. This was insipired by tornado cash note accounts which kind of do the same thing but with a work around for privacy. Here privacy is not a concern.

`depositToken` works the same way as depositEth but will need to pass the address of the ERC20 token also

Accounts only work with a single kind of currency at once, so different currencies need different accounts.

`topUpETH` will let the user top up their account, identified by the commitment 

`topUpTokens` will do the same for ERC20 tokens.

`withdraw` will let the account creator withdraw the deposited value and this will also close the account. 

**Connected Wallets**

`connectWallet` is used to connect an external wallet account to the smart contract using a commitment, the token to use for payments and the encrypted note.

ConnectedWallets only use ERC-20 tokens.

`disconnectWallet` The owner of the wallet can disable the account and disconnect his wallet. Nullifying all future payments

`connectedWalletAlready` Is a mapping (bytes32 => bool) to check if a wallet connected already to disallow wallets connecting twice to the same smart contarct with the same tokens.
Each smart contract manages 1 connected wallet per token!

`getHashByAddresses` is a pure function to calculate a keccak256 hash from 2 addresses, a wallet address and a token contract address. This is used to access connectedWalletAlready mapping!

## Client Side 

To create accounts and payment intents, we use cryptography that is not included by default in the browser. 
The `/lib` directory contains the code needed to do this.
`directDebit.ts` and `directDebit.js` is identical, but the js version was needed for more simple bundling configurations.
To Bundle the code Rollup is used, it's building a minified iife that can be imported by the browser.
You can run `npm run build-dep` to build the dependency or find the file commited in the repository at `directdebit_bundle.js`

### API
I'm only going to mention some of the more important functions here

`newAccountSecrets` is used to compute a new account secret that can be saved as an encrypted note inside the smart contract. Always encrypt this before saving it else anyone can access it! We use double layer encryption in DebitLlama for this, that is both symmetric and asymmetric encryption before saving it.

`decodeAccountSecrets` will take the unencrypted note and extract the secret.

`createPaymentIntent` will compute a zksnark using the note and the payment intent parameters

**Account Note Encryption**

`encryptData` will use a public key and encrypt the note using eth-sig-util, the encrypted data will need to be packed with `packEncryptedMEssage` before it can be saved on-chain. The encryption uses a public key used by ethereum wallets.

`decryptData` will decrypt the encrypted data, but before it can do it we need to unpack it using `unpackEncryptedMessage`. The decryption uses an etherem wallet private key.

The client must make sure to implement a symmetric encryption if it has no access to the wallet private key like this is done with DebitLlama where the ethereum key used to encrypt is supplied by the server, but the final decryption happens in the browser using a user supplied password. This way DebitLlama has zero access to the underlying secret at any time and creating payment intents and spending from a wallet is abstracted to just supplying a password by the unserlying service!


## How to Run Tests

`npx hardhat test`

## Deployment

Configure hardhat.config.ts with the network and run the deploy script

`npx hardhat run scripts/deploy.ts --network <network>`


# Other contracts
The repository contains other contracts used by DebitLlama like the *RelayerGasTracker*.
This is a simple convenience contract that allows depositing gas to a relayer, instead of sending a transaction directly we use a smart contract for this to store the top up history. `TopUpEvent` events are emitted that can be fetched by a relayer to process top up transactions to it using contract calls.

# Trusted Setup
The project uses the Polygon Hermez ptau file

The phase-2 ceremony was done using snarkyceremonies.com with anonymous contributors!

snarkjs zkey verify ./circuits/directDebit/directDebit.r1cs ./circuits/powersOfTau28_hez_final_15.ptau  ./circuits/directDebit/directDebit_0010.zkey 

### Adding a random beacon

For the beacon, I chose to use bitcoin block hash 812650
`00000000000000000002863dc27bf05659898e74dcb5e20167a16d71d1024612`

`snarkjs zkey beacon directDebit_0010.zkey directDebit_final.zkey 00000000000000000002863dc27bf05659898e74dcb5e20167a16d71d1024612 10 -n="Final Beacon Phase 2 DirectDebit Mainnet is Ready :)"`

Verify the final zkey 
`snarkjs zkey verify directDebit.r1cs ../powersOfTau28_hez_final_15.ptau directDebit_final.zkey`

Export the verification key:

`snarkjs zkey export verificationkey directDebit_final.zkey verification_key.json`

Export the verifier smart contract

`snarkjs zkey export solidityverifier directDebit_final.zkey PaymentIntentVerifier.sol`

And finally the verifier was copied to the contracts library

`mv PaymentIntentVerifier.sol ../../contracts/PaymentIntentVerifier.sol`


# Latest Deployments

### Doanu Testnet (latest)
`Verifier contract is deployed to  0xA0c953Db12f02e0E8f41EFd5Ea857259a694069d`
`Virtual Accounts contract is deployed to :  0x2137F4096365bCA1457E945838e0d7EC1925A973`
`Connected Wallets contract is deployed to:  0xc65DDA2E81dB71C998D08A525D70dFA844BF5D3e`

### BTT MAINNET ADDRESSES

`Verifier contract is deployed to  0x5e93788886D8712C0cDe623fB22dCf979ed07724`
`Virtual Accounts contract is deployed to :  0xc4Cf42D5a6F4F061cf5F98d0338FC5913b6fF581`
`Connected Wallets contract is deployed to:  0xF9962f3C23De4e864E56ef29125D460c785905c6`