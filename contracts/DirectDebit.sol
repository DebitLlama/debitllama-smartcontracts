// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./Errors.sol";

// This contract  implements direct debit using crypto notes
// The user can create a Note off-chain that is stored encrypted inside the smart contract
// Then create an account which contains commitment from the note, this account can be topped up as needed.
// The Note can be used off-chain to compute proofs that allow a payee to withdraw funds from this account.

// The data for the accounts is stored in this struct
struct AccountData {
    bool active;
    address creator;
    IERC20 token;
    uint256 balance;
}

/*
  The the debit history is used for tracking payment intents and nullifying them!
  it tracks withdrawalCount and the lastDate the proof was used for withdrawing value
*/
struct PaymentIntentHistory {
    bool isNullified;
    uint256 withdrawalCount;
    uint256 lastDate;
}

// The interface of the Verifier contract generated from the circuit
interface IVerifier {
    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[6] memory _input
    ) external returns (bool);
}

/*
 */
contract DirectDebit is ReentrancyGuard, DirectDebitErrors, DirectDebitEvents {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    IVerifier public immutable verifier; // The verifier address

    address payable public _owner; // The owner of the contract that can update the feeless token's address

    /**
     1% fee sent to the deployer of this contract on all transactions. 
     We divide the payout with the ownerFeeDivider
    */
    uint256 public immutable ownerFeeDivider = 100;

    /**
    The commitments for each address are stored in a mapping for easy access
     */
    mapping(address => mapping(uint256 => bytes32)) public commitments;

    /**
    The counter is used to access the account commitment index
     */
    mapping(address => uint256) public accountCounter;

    /* PaymentIntent Hashes are keys to the history of payments
     Each payment intent has a unique hash thanks to the nonce added to the nullifier secret nullifier
     The mapping tracks how many times a payment intent was used and will nullify future payments
     */
    mapping(bytes32 => PaymentIntentHistory) public paymentIntents;

    /*
       accounts are keys to Account data
    */
    mapping(bytes32 => AccountData) public accounts;

    /*
       The Secret notes are encrypted and stored on-chain
       This feature relies on metamask. The note was encrypted with the public key aquired with eth_getEncryptionPublicKey
       Decryption requires the use of eth_decrypt on the client
       These notes are used to compute the ZKP (Payment Intent) off-chain that can be used to debit an account
    */
    mapping(bytes32 => string) public encryptedNotes;

    /**
        @dev : the constructor
        @param _verifier is the address of SNARK verifier contract        
    */

    constructor(IVerifier _verifier) {
        verifier = _verifier;
        _owner = payable(msg.sender);
    }

    /**
     @dev calculate the fee used for the payment
     @param amount is the debited amount
   */
    function calculateFee(
        uint256 amount
    ) public pure returns (uint256 ownerFee, uint256 payment) {
        ownerFee = amount.div(ownerFeeDivider);
        payment = amount.sub(ownerFee);
    }

    /**
      @dev : Create a new Account by depositing ETH
      @param _commitment is the poseidon hash created for the note on client side
      @param balance that is the value of the note. 
      @param encryptedNote is the crypto note that is encrypted client side and stored in the contract. 
      Storing the note allows for meta transactions where the user needs to decrypt instead of sign a message, and compute an off-chain zkp.
      The zkp created is a "payment intent" and no gas fees are involved creating those!
    */

    function depositEth(
        bytes32 _commitment,
        uint256 balance,
        string calldata encryptedNote
    ) external payable nonReentrant {
        if (accounts[_commitment].active) revert AccountAlreadyActive();
        // When cancelling a commitment and withdrawing, accounts.active will be set to false, but the creator address won't be zero!
        // This is an edge case I check for here, to not deposit into inactivated but previously existing accounts!
        if (accounts[_commitment].creator != address(0))
            revert AccountAlreadyExists();
        if (balance == 0) revert ZeroTopup();
        if (balance != msg.value) revert NotEnoughValue();

        // A convenience mapping to fetch the commitments by address to access accounts later!
        commitments[msg.sender][accountCounter[msg.sender]] = _commitment;
        accountCounter[msg.sender] += 1;

        // Record the Account creation
        accounts[_commitment].active = true;
        accounts[_commitment].creator = msg.sender;
        accounts[_commitment].balance = balance;
        // Save the encrypted note string to storage!
        encryptedNotes[_commitment] = encryptedNote;

        emit NewEthAccount(_commitment, msg.sender, balance);
    }

    /**
   @dev : depositToken is for creating a bunnyNote by depositing tokens, the wallet calling this function must approve ERC20 spend first
   @param _commitment is the poseidon hash of the note
   @param balance is the amount of token transferred to the contract that represents the note's value. balance does not contain the fee
   @param token is the ERC20 token that is used for this deposits  
   @param encryptedNote is the crypto note created for this account
    */
    function depositToken(
        bytes32 _commitment,
        uint256 balance,
        address token,
        string calldata encryptedNote
    ) external nonReentrant {
        if (accounts[_commitment].active) revert AccountAlreadyActive();
        if (balance == 0) revert ZeroTopup();
        // When cancelling a commitment and withdrawing, active will be set back to false, but the creator address won't be zero!
        if (accounts[_commitment].creator != address(0))
            revert AccountAlreadyExists();

        // A convenience mapping to fetch the commitments by address to access accounts later!
        commitments[msg.sender][accountCounter[msg.sender]] = _commitment;
        accountCounter[msg.sender] += 1;

        accounts[_commitment].active = true;
        accounts[_commitment].creator = msg.sender;
        accounts[_commitment].balance = balance;
        accounts[_commitment].token = IERC20(token);
        encryptedNotes[_commitment] = encryptedNote;
        //If the user doesn't have enough token balance this will throw
        IERC20(token).safeTransferFrom(msg.sender, address(this), balance);

        emit NewTokenAccount(_commitment, msg.sender, balance, token);
    }

    /**
      Top up your balance with ETH or other gas tokens
      @param _commitment is the identifier of the account
      @param balance is the top up balance to add to the account

      It is allowed for a user to top up another user's account.
    */

    function topUpETH(
        bytes32 _commitment,
        uint256 balance
    ) external payable nonReentrant {
        if (!accounts[_commitment].active) revert InactiveAccount();
        if (balance == 0) revert ZeroTopup();
        if (balance != msg.value) revert NotEnoughValue();
        if (accounts[_commitment].token != IERC20(address(0)))
            revert NotEthAccount();
        // Adds the balance to the account!
        accounts[_commitment].balance += balance;
        // Emit a top up event
        emit TopUpETH(_commitment, balance);
    }

    /**
      Top up your account balance with tokens
      @param _commitment is the identifier of the account
      @param balance is the amount of top up balance
     */

    function topUpTokens(
        bytes32 _commitment,
        uint256 balance
    ) external nonReentrant {
        if (!accounts[_commitment].active) revert InactiveAccount();
        if (balance == 0) revert ZeroTopup();
        if (accounts[_commitment].token == IERC20(address(0)))
            revert NotTokenAccount();

        accounts[_commitment].balance += balance;
        IERC20(accounts[_commitment].token).safeTransferFrom(
            msg.sender,
            address(this),
            balance
        );

        emit TopUpToken(
            _commitment,
            balance,
            address(accounts[_commitment].token)
        );
    }

    /**
      A function that allows direct debit with a reusable proof
      N times to M address with L max amount that can be withdrawn
      The proof and public inputs are the PaymentIntent
      @param proof contains the zkSnark
      @param hashes are [0] = paymentIntent [1] = commitment
      @param payee is the account recieving the payment
      @param debit[4] are [0] = max debit amount, [1] = debitTimes, [2] = debitInterval, [3] = payment amount. 
      Last param is not used in the circuit but it must be smaller than the max debit amount
      By using a separate max debit amount and a payment amount we can create dynamic subscriptions, where the final price varies 
      but can't be bigger than the allowed amount!
    */
    function directdebit(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) external nonReentrant {
        _verifyPaymentIntent(proof, hashes, payee, debit);
        _processPaymentIntent(hashes, payee, debit);
    }

    /**
      Cancels the payment intent! The caller must be the creator of the account and must have the zksnark (paymentIntent)
      The zksnark is needed so the Payment Intent can be cancelled before it's used and for that we need proof that it exists!
      @param proof is the snark
      @param hashes are [0] = paymentIntent [1] = commitment
      @param payee is the account recieving the payment
      @param debit [4] are [0] = max debit amount, [1] = debitTimes, [2] = debitInterval, [3] = payment amount. 
      In case of cancellation the 4th value in the debit array can be arbitrary. It is kept here to keep the verifier function's interface
     */
    function cancelPaymentIntent(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) external {
        if (!_verifyProof(proof, hashes, payee, debit)) revert InvalidProof();
        if (msg.sender != accounts[hashes[1]].creator && msg.sender != payee)
            revert OnlyRelatedPartiesCanCancel();
        paymentIntents[hashes[0]].isNullified = true;
        emit PaymentIntentCancelled(hashes[1], hashes[0], payee);
    }

    // The account creator can withdraw the value deposited and close the account
    // This will set the active false but the creator address remains, hence the edge case we handled on account creation
    function withdraw(bytes32 commitment) external nonReentrant {
        if (!accounts[commitment].active) revert InactiveAccount();
        if (msg.sender != accounts[commitment].creator)
            revert OnlyAccountOwner();

        accounts[commitment].active = false;
        uint256 balance = accounts[commitment].balance;

        if (address(accounts[commitment].token) != address(0)) {
            _processTokenWithdraw(commitment, msg.sender, balance);
        } else {
            _processEthWithdraw(commitment, msg.sender, balance);
        }
        emit AccountClosed(commitment);
    }

    /**
       The direct debit transaction is verified using these conditions!
    */

    function _verifyPaymentIntent(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) internal {
        // The payment intent was cancelled by the account!
        if (paymentIntents[hashes[0]].isNullified)
            revert PaymentIntentNullified();
        // The payment intent was withdrawn already N times.
        if (paymentIntents[hashes[0]].withdrawalCount == debit[1])
            revert PaymentIntentExpired();
        // The account we are charging is inactive!
        if (!accounts[hashes[1]].active) revert InactiveAccount();

        // The authorized amount must be bigger or equal than the amount withdrawn!
        if (debit[0] < debit[3]) revert PaymentNotAuthorized();

        // The account has insufficient balance to continue
        if (debit[3] > accounts[hashes[1]].balance)
            revert NotEnoughAccountBalance();

        // Enforce the debit interval!
        // The debitInterval is always in days!
        if (
            paymentIntents[hashes[0]].lastDate.add(debit[2] * 1 days) >
            block.timestamp
        ) revert EarlyPaymentNotAllowed();
        // verify the ZKP
        if (!_verifyProof(proof, hashes, payee, debit)) revert InvalidProof();
    }

    function _verifyProof(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) internal returns (bool) {
        return
            verifier.verifyProof(
                [proof[0], proof[1]],
                [[proof[2], proof[3]], [proof[4], proof[5]]],
                [proof[6], proof[7]],
                [
                    uint256(hashes[0]),
                    uint256(hashes[1]),
                    uint256(uint160(payee)),
                    uint256(debit[0]),
                    uint256(debit[1]),
                    uint256(debit[2])
                ]
            );
    }

    /**
      This function will process the payment intent and trigger the withdrawals
      It sends rewards to the relayers or cashback to the account user!
    */

    function _processPaymentIntent(
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) internal {
        // Calculate the fee
        (uint256 ownerFee, uint256 payment) = calculateFee(debit[3]);

        // Add a debit time to the nullifier
        paymentIntents[hashes[0]].withdrawalCount += 1;
        paymentIntents[hashes[0]].lastDate = block.timestamp;

        // Process the withdraw
        if (address(accounts[hashes[1]].token) != address(0)) {
            _processTokenWithdraw(hashes[1], payee, payment);
            _processTokenWithdraw(hashes[1], _owner, ownerFee);
        } else {
            // Transfer the eth
            _processEthWithdraw(hashes[1], payee, payment);
            // Send the fee to the owner
            _processEthWithdraw(hashes[1], _owner, ownerFee);
        }

        emit AccountDebited(hashes[1], payee, payment);
    }

    /*
       Process the token withdrawal and decrease the account balance
    */

    function _processTokenWithdraw(
        bytes32 commitment,
        address payee,
        uint256 payment
    ) internal {
        accounts[commitment].token.safeTransfer(payable(payee), payment);
        accounts[commitment].balance -= payment;
    }

    /**
      Process the eth withdraw and decrease the account balance
    */

    function _processEthWithdraw(
        bytes32 commitment,
        address payee,
        uint256 payment
    ) internal {
        Address.sendValue(payable(payee), payment);
        accounts[commitment].balance -= payment;
    }
}
