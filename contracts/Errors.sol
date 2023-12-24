// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface DirectDebitErrors {
    /*
   Account already active.This error is thrown if we try to create an account with a commitment that exists already!
   This error occurs when creating accounts!
*/
    error AccountAlreadyActive();
    /*
  If the account needs to be active but it is not InactiveAccount() error is thrown.
  This error can occur when topping up an account or when trying to debit from a disabled or non-existent account
*/
    error InactiveAccount();

    /*
 This error is thrown when an account was already created and has a creator address.
 It could occur if an account was deactivated but existed before and a user tries to create it again, which should fail
*/
    error AccountAlreadyExists();

    /*
 This error is thrown if a user tries to deposit zero amount when creating or when topping up an account
*/

    error ZeroTopup();

    /*
This error is thrown when an ETH top up msg.value is not balance
*/
    error NotEnoughValue();

    /*
 This error is thrown if we try to TopUp an account with tokens that supports ETH
*/
    error NotTokenAccount();

    /**
  This error is thrown when we try to top up a token account with ETH
   */
    error NotEthAccount();

    /*
  This is thrown when a payment intent was cancelled by it's owner.
*/
    error PaymentIntentNullified();

    /*
 The payment intent has expired if all the allowed payments were made with it!
*/
    error PaymentIntentExpired();

    /**
 directdebit throws this if the requested amount is higher than the account balance 
*/
    error NotEnoughAccountBalance();

    /**
  DirectDebit will throw this if the payment is larger than the max allowed amount!
 */

    error PaymentNotAuthorized();

    /**
   DirectDebit will throw this if the debit interval does not allow withdrawal,yet.
   */
    error EarlyPaymentNotAllowed();

    /**
    Error thrown if the zkSnark verificaiton failed
     */
    error InvalidProof();

    /**
      Access control, thrown if now account owner is calling the function
     */
    error OnlyAccountOwner();

    /**
      Only Related parties can cancel a payment intent. This means the creator or the payee
     */
    error OnlyRelatedPartiesCanCancel();

    /**
    Thrown if the commitment on a payment intent  history does not match the account commitment!
     */
    error CommitmentMismatch();

    /**
    Only owner can call this function!
     */
    error OnlyOwner();

    /**
    Thrown when an unsupported funciton is called in a child contract
     */
    error FunctionNotSupported();

    /**
     This error occurs when a Wallet Tries to connect with a zero address token!
     */
    error ZeroAddressConnected();

    /**
    This error occurs when a wallet tries to connect twice with the the same tokens. That doesn't work.
     */
    error WalletAlreadyConnected();

    /**
    This error occurs on direct debit if the relayer was not approved
     */
    error OnlyApprovedRelayer();
}

/**

This interface contains the custom events that are emitted from the DirectDebit smart contract

 */

interface DirectDebitEvents {
    /**
      Emitted when gas tokens are deposited and a new account is created
     */
    event NewEthAccount(
        bytes32 indexed commitment,
        address depositFor,
        uint256 balance
    );
    /**
   Emitted when gas tokens are added to an existing account!
   */

    event TopUpETH(bytes32 indexed commitment, uint256 balance);
    /**
    Emitted when a new Token account is created
   */
    event NewTokenAccount(
        bytes32 indexed commitment,
        address depositFor,
        uint256 amount,
        address token
    );
    /**
      Emitted when a token account is topped up
    */
    event TopUpToken(bytes32 indexed commitment, uint256 amount, address token);

    /**
    Emitted when a token account is debited
     */
    event AccountDebited(
        bytes32 indexed commitment,
        address payee,
        uint256 payment
    );

    /**
    Emitted when a payment intent is cancelled
    */
    event PaymentIntentCancelled(
        bytes32 indexed commitment,
        bytes32 indexed paymentIntent,
        address payee
    );

    /**
    Emitted when an account was closed
    */
    event AccountClosed(bytes32 indexed commitment);

    /**
    Emitted when a new wallet was connected!
     */
    event NewWalletConnected(
        bytes32 indexed commitment,
        address indexed creator,
        address indexed token
    );
}
