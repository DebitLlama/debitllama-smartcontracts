# Direct Debit Using Note Accounts

### How it works: Circuits

There is a DirectDebit() template in the directDebit.circom file.
This contains public inputs paymentIntent, commitmentHash, payee, maxDebitAmount, debitTimes, debitInterval and private inputs secret, nonce and nullifier.

When creating an account we calculate the hash of the secret with the nullifier.
Commitment is hash(nullifier,secret)

When creating a payment intent, the intent itself is the nullifier, we make it reusable by hashing it with a nonce
PaymentIntent = hash(nullifier,nonce)
The nonce is random.

To verify other public inputs like payee, maxDebitAmount,debitTimes and debitInterval we use Hidden Signals. These signals are used to make sure that these parameters of the payment intent cannot be altered.

### How it works: Smart Contract

The DirectDebit.sol smart contract contains the implementation for the payments.
users create Accounts with AccountData and have PaymentIntentHistory on processed payments. 

Payment Intents are created off-chain, passed to the merchant or relayer and then used to charge the account.

Constructor takes the verifier contract address and the owner of the contract.

A function is used to calculate the fees that is sent to the relayer, the owner and the final payment. If the creator of the account is the relayer then it's a 0.5% Cashback basicly and not a fee.

`depositEth` is used to create an account. The user computed the commitment off-chain and deposits value. The secret and the nullifier must be encrypted with with the wallet using the RPC call `eth_getEncryptionPublicKey` to get the key.  The crypto note is essentially saved on the chain so the users don't have to download a file. This was insipired by tornado cash note accounts which kind of do the same thing but with a work around for privacy. Here privacy is not a concern.

`depositToken` works the same way as depositEth but will need to pass the address of the ERC20 token also

Accounts only work with a single kind of currency at once, so different currencies need different accounts.

`topUpETH` will let the user top up their account, identified by the commitment and `topUpTokens` will do the same for ERC20 tokens.

`directdebit` The function is used to process payment intents that were created off chain and belong to an account. this will verify and process the payment

`cancelPaymentIntent` will let the account owner to cancel a payment intent so it can't be used anymore. The account owner must posess the payment intent to do this. This will not close the account, only cancel all future direct debits with the Payment Intent

`withdraw` will let the account creator withdraw the deposited value and this will also close the account. 