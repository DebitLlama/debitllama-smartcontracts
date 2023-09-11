// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "./DirectDebit.sol";

// This contract implements direct debit from a connected wallet
// It supports only ERC-20 tokens
// The wallet must connect and then approve spending tokens for the direct debit to function!

contract ConnectedWallets is DirectDebit {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /**
        @dev : the constructor
        @param _verifier is the address of SNARK verifier contract        
    */

    constructor(IVerifier _verifier) DirectDebit(_verifier) {}

    /**
      @dev  Connect an external wallet to create an account that allows directly debiting it!
      @param _commitment is the poseidon hash of the note
      @param token is the ERC20 token that is used for transactions, the connected wallet must approve allowance!
      @param encryptedNote is the crypto note created for this account
     */
    function connectWallet(
        bytes32 _commitment,
        address token,
        string calldata encryptedNote
    ) external nonReentrant {
        if (accounts[_commitment].active) revert AccountAlreadyActive();
        // When cancelling a commitment accounts.active will be set to false, but the creator address won't be zero!
        // This is an edge case I check for here, to not deposit into inactivated but previously existing accounts!
        if (accounts[_commitment].creator != address(0))
            revert AccountAlreadyExists();

        if (token == (address(0))) revert ZeroAddressConnected();

        // A convenience mapping to fetch the commitments by address to access accounts later!
        commitments[msg.sender][accountCounter[msg.sender]] = _commitment;
        accountCounter[msg.sender] += 1;

        // Record the Account creation
        accounts[_commitment].active = true;
        accounts[_commitment].creator = msg.sender;
        accounts[_commitment].token = IERC20(token);

        // For connected wallets the accounts.balance parameter is not used for balance
        // The erc-20 token allowance is used instead!

        // Save the encrypted note string to storage!
        encryptedNotes[_commitment] = encryptedNote;

        emit NewWalletConnected(_commitment, msg.sender, token);
    }

    /**
      @dev  disconnect the wallet from the account, the payment intents can't be used to debit it from now on!
      @param  commitment is the commitment of the account    
     */

    function disconnectWallet(bytes32 commitment) external nonReentrant {
        if (!accounts[commitment].active) revert InactiveAccount();
        if (msg.sender != accounts[commitment].creator)
            revert OnlyAccountOwner();

        accounts[commitment].active = false;
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
        accounts[commitment].token.safeTransferFrom(
            accounts[commitment].creator,
            payable(payee),
            payment
        );
    }

    /**
      Processing ETH withdrawals is not supported with connected wallets!
    */

    function _processEthWithdraw(
        bytes32 commitment,
        address payee,
        uint256 payment
    ) internal pure override {
        revert FunctionNotSupported();
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

        // The connected wallet has insufficient allowance I throw an error
        if (
            debit[3] >
            IERC20(accounts[hashes[1]].token).allowance(
                accounts[hashes[1]].creator,
                address(this)
            )
        ) revert NotEnoughAccountBalance();

        // Enforce the debit interval!
        // The debitInterval is always in days!
        if (
            paymentIntents[hashes[0]].lastDate.add(debit[2] * 1 days) >
            block.timestamp
        ) revert EarlyPaymentNotAllowed();

        // verify the ZKP
        if (!_verifyProof(proof, hashes, payee, debit)) revert InvalidProof();
    }
}
