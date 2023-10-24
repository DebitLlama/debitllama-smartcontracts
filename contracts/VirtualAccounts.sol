// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "./DirectDebit.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./Errors.sol";

//   o              o     o                o                                 o               o                                                                      o
//  <|>            <|>  _<|>_             <|>                               <|>             <|>                                                                    <|>
//  < >            < >                    < >                               / \             / \                                                                    < >
//   \o            o/     o    \o__ __o    |       o       o      o__ __o/  \o/           o/   \o           __o__      __o__    o__ __o     o       o   \o__ __o    |
//    v\          /v     <|>    |     |>   o__/_  <|>     <|>    /v     |    |           <|__ __|>         />  \      />  \    /v     v\   <|>     <|>   |     |>   o__/_
//     <\        />      / \   / \   < >   |      < >     < >   />     / \  / \          /       \       o/         o/        />       <\  < >     < >  / \   / \   |
//       \o    o/        \o/   \o/         |       |       |    \      \o/  \o/        o/         \o    <|         <|         \         /   |       |   \o/   \o/   |
//        v\  /v          |     |          o       o       o     o      |    |        /v           v\    \\         \\         o       o    o       o    |     |    o
//         <\/>          / \   / \         <\__    <\__ __/>     <\__  / \  / \      />             <\    _\o__</    _\o__</   <\__ __/>    <\__ __/>   / \   / \   <\__

// This contract implements the direct debit from virtual accounts
// The accounts support ETH and ERC-20 tokens
// The accounts need to be topped up and manually and they can be closed by the wallet that created them!

contract VirtualAccounts is DirectDebit {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /**
        @dev : the constructor
        @param _verifier is the address of SNARK verifier contract        
    */

    constructor(IVerifier _verifier) DirectDebit(_verifier) {}

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
   @dev : depositToken is for creating an account by depositing tokens, the wallet calling this function must approve ERC20 spend first
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

    /*
       Process the token withdrawal and decrease the account balance
    */

    function _processTokenWithdraw(
        bytes32 commitment,
        address payee,
        uint256 payment
    ) internal override {
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
    ) internal override {
        Address.sendValue(payable(payee), payment);
        accounts[commitment].balance -= payment;
    }

    /**
       The direct debit transaction is verified using these conditions!
    */

    function _verifyPaymentIntent(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) internal override {
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

    /**
    A view function to get and display the account's balance
    This is useful when the account balance is calculated from external wallet's balance!
     */
    function getAccount(
        bytes32 commitment
    ) external view override returns (AccountData memory) {
        return accounts[commitment];
    }
}
