// Sources flattened with hardhat v2.15.0 https://hardhat.org

// File @openzeppelin/contracts/utils/Context.sol@v4.9.1

// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

pragma solidity ^0.8.0;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}


// File @openzeppelin/contracts/access/Ownable.sol@v4.9.1

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}


// File @openzeppelin/contracts/security/Pausable.sol@v4.9.1


/**
 * @dev Contract module which allows children to implement an emergency stop
 * mechanism that can be triggered by an authorized account.
 *
 * This module is used through inheritance. It will make available the
 * modifiers `whenNotPaused` and `whenPaused`, which can be applied to
 * the functions of your contract. Note that they will not be pausable by
 * simply including this module, only once the modifiers are put in place.
 */
abstract contract Pausable is Context {
    /**
     * @dev Emitted when the pause is triggered by `account`.
     */
    event Paused(address account);

    /**
     * @dev Emitted when the pause is lifted by `account`.
     */
    event Unpaused(address account);

    bool private _paused;

    /**
     * @dev Initializes the contract in unpaused state.
     */
    constructor() {
        _paused = false;
    }

    /**
     * @dev Modifier to make a function callable only when the contract is not paused.
     *
     * Requirements:
     *
     * - The contract must not be paused.
     */
    modifier whenNotPaused() {
        _requireNotPaused();
        _;
    }

    /**
     * @dev Modifier to make a function callable only when the contract is paused.
     *
     * Requirements:
     *
     * - The contract must be paused.
     */
    modifier whenPaused() {
        _requirePaused();
        _;
    }

    /**
     * @dev Returns true if the contract is paused, and false otherwise.
     */
    function paused() public view virtual returns (bool) {
        return _paused;
    }

    /**
     * @dev Throws if the contract is paused.
     */
    function _requireNotPaused() internal view virtual {
        require(!paused(), "Pausable: paused");
    }

    /**
     * @dev Throws if the contract is not paused.
     */
    function _requirePaused() internal view virtual {
        require(paused(), "Pausable: not paused");
    }

    /**
     * @dev Triggers stopped state.
     *
     * Requirements:
     *
     * - The contract must not be paused.
     */
    function _pause() internal virtual whenNotPaused {
        _paused = true;
        emit Paused(_msgSender());
    }

    /**
     * @dev Returns to normal state.
     *
     * Requirements:
     *
     * - The contract must be paused.
     */
    function _unpause() internal virtual whenPaused {
        _paused = false;
        emit Unpaused(_msgSender());
    }
}


// File @openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol@v4.9.1


/**
 * @dev Interface of the ERC20 Permit extension allowing approvals to be made via signatures, as defined in
 * https://eips.ethereum.org/EIPS/eip-2612[EIP-2612].
 *
 * Adds the {permit} method, which can be used to change an account's ERC20 allowance (see {IERC20-allowance}) by
 * presenting a message signed by the account. By not relying on {IERC20-approve}, the token holder account doesn't
 * need to send a transaction, and thus is not required to hold Ether at all.
 */
interface IERC20Permit {
    /**
     * @dev Sets `value` as the allowance of `spender` over ``owner``'s tokens,
     * given ``owner``'s signed approval.
     *
     * IMPORTANT: The same issues {IERC20-approve} has related to transaction
     * ordering also apply here.
     *
     * Emits an {Approval} event.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `deadline` must be a timestamp in the future.
     * - `v`, `r` and `s` must be a valid `secp256k1` signature from `owner`
     * over the EIP712-formatted function arguments.
     * - the signature must use ``owner``'s current nonce (see {nonces}).
     *
     * For more information on the signature format, see the
     * https://eips.ethereum.org/EIPS/eip-2612#specification[relevant EIP
     * section].
     */
    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;

    /**
     * @dev Returns the current nonce for `owner`. This value must be
     * included whenever a signature is generated for {permit}.
     *
     * Every successful call to {permit} increases ``owner``'s nonce by one. This
     * prevents a signature from being used multiple times.
     */
    function nonces(address owner) external view returns (uint256);

    /**
     * @dev Returns the domain separator used in the encoding of the signature for {permit}, as defined by {EIP712}.
     */
    // solhint-disable-next-line func-name-mixedcase
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}


// File @openzeppelin/contracts/token/ERC20/IERC20.sol@v4.9.1


/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `from` to `to` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}


// File @openzeppelin/contracts/utils/Address.sol@v4.9.1


/**
 * @dev Collection of functions related to the address type
 */
library Address {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     *
     * Furthermore, `isContract` will also return true if the target contract within
     * the same transaction is already scheduled for destruction by `SELFDESTRUCT`,
     * which only has an effect at the end of a transaction.
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://consensys.net/diligence/blog/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.8.0/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(address target, bytes memory data, uint256 value) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verify that a low level call to smart-contract was successful, and revert (either by bubbling
     * the revert reason or using the provided one) in case of unsuccessful call or if target was not a contract.
     *
     * _Available since v4.8._
     */
    function verifyCallResultFromTarget(
        address target,
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        if (success) {
            if (returndata.length == 0) {
                // only check isContract if the call was successful and the return data is empty
                // otherwise we already know that it was a contract
                require(isContract(target), "Address: call to non-contract");
            }
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    /**
     * @dev Tool to verify that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason or using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    function _revert(bytes memory returndata, string memory errorMessage) private pure {
        // Look for revert reason and bubble it up if present
        if (returndata.length > 0) {
            // The easiest way to bubble the revert reason is using memory via assembly
            /// @solidity memory-safe-assembly
            assembly {
                let returndata_size := mload(returndata)
                revert(add(32, returndata), returndata_size)
            }
        } else {
            revert(errorMessage);
        }
    }
}


// File @openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol@v4.9.1



/**
 * @title SafeERC20
 * @dev Wrappers around ERC20 operations that throw on failure (when the token
 * contract returns false). Tokens that return no value (and instead revert or
 * throw on failure) are also supported, non-reverting calls are assumed to be
 * successful.
 * To use this library you can add a `using SafeERC20 for IERC20;` statement to your contract,
 * which allows you to call the safe operations as `token.safeTransfer(...)`, etc.
 */
library SafeERC20 {
    using Address for address;

    /**
     * @dev Transfer `value` amount of `token` from the calling contract to `to`. If `token` returns no value,
     * non-reverting calls are assumed to be successful.
     */
    function safeTransfer(IERC20 token, address to, uint256 value) internal {
        _callOptionalReturn(token, abi.encodeWithSelector(token.transfer.selector, to, value));
    }

    /**
     * @dev Transfer `value` amount of `token` from `from` to `to`, spending the approval given by `from` to the
     * calling contract. If `token` returns no value, non-reverting calls are assumed to be successful.
     */
    function safeTransferFrom(IERC20 token, address from, address to, uint256 value) internal {
        _callOptionalReturn(token, abi.encodeWithSelector(token.transferFrom.selector, from, to, value));
    }

    /**
     * @dev Deprecated. This function has issues similar to the ones found in
     * {IERC20-approve}, and its usage is discouraged.
     *
     * Whenever possible, use {safeIncreaseAllowance} and
     * {safeDecreaseAllowance} instead.
     */
    function safeApprove(IERC20 token, address spender, uint256 value) internal {
        // safeApprove should only be called when setting an initial allowance,
        // or when resetting it to zero. To increase and decrease it, use
        // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
        require(
            (value == 0) || (token.allowance(address(this), spender) == 0),
            "SafeERC20: approve from non-zero to non-zero allowance"
        );
        _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, value));
    }

    /**
     * @dev Increase the calling contract's allowance toward `spender` by `value`. If `token` returns no value,
     * non-reverting calls are assumed to be successful.
     */
    function safeIncreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        uint256 oldAllowance = token.allowance(address(this), spender);
        _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, oldAllowance + value));
    }

    /**
     * @dev Decrease the calling contract's allowance toward `spender` by `value`. If `token` returns no value,
     * non-reverting calls are assumed to be successful.
     */
    function safeDecreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        unchecked {
            uint256 oldAllowance = token.allowance(address(this), spender);
            require(oldAllowance >= value, "SafeERC20: decreased allowance below zero");
            _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, oldAllowance - value));
        }
    }

    /**
     * @dev Set the calling contract's allowance toward `spender` to `value`. If `token` returns no value,
     * non-reverting calls are assumed to be successful. Compatible with tokens that require the approval to be set to
     * 0 before setting it to a non-zero value.
     */
    function forceApprove(IERC20 token, address spender, uint256 value) internal {
        bytes memory approvalCall = abi.encodeWithSelector(token.approve.selector, spender, value);

        if (!_callOptionalReturnBool(token, approvalCall)) {
            _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, 0));
            _callOptionalReturn(token, approvalCall);
        }
    }

    /**
     * @dev Use a ERC-2612 signature to set the `owner` approval toward `spender` on `token`.
     * Revert on invalid signature.
     */
    function safePermit(
        IERC20Permit token,
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal {
        uint256 nonceBefore = token.nonces(owner);
        token.permit(owner, spender, value, deadline, v, r, s);
        uint256 nonceAfter = token.nonces(owner);
        require(nonceAfter == nonceBefore + 1, "SafeERC20: permit did not succeed");
    }

    /**
     * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
     * on the return value: the return value is optional (but if data is returned, it must not be false).
     * @param token The token targeted by the call.
     * @param data The call data (encoded using abi.encode or one of its variants).
     */
    function _callOptionalReturn(IERC20 token, bytes memory data) private {
        // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
        // we're implementing it ourselves. We use {Address-functionCall} to perform this call, which verifies that
        // the target address contains contract code and also asserts for success in the low-level call.

        bytes memory returndata = address(token).functionCall(data, "SafeERC20: low-level call failed");
        require(returndata.length == 0 || abi.decode(returndata, (bool)), "SafeERC20: ERC20 operation did not succeed");
    }

    /**
     * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
     * on the return value: the return value is optional (but if data is returned, it must not be false).
     * @param token The token targeted by the call.
     * @param data The call data (encoded using abi.encode or one of its variants).
     *
     * This is a variant of {_callOptionalReturn} that silents catches all reverts and returns a bool instead.
     */
    function _callOptionalReturnBool(IERC20 token, bytes memory data) private returns (bool) {
        // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
        // we're implementing it ourselves. We cannot use {Address-functionCall} here since this should return false
        // and not revert is the subcall reverts.

        (bool success, bytes memory returndata) = address(token).call(data);
        return
            success && (returndata.length == 0 || abi.decode(returndata, (bool))) && Address.isContract(address(token));
    }
}


// File @openzeppelin/contracts/security/ReentrancyGuard.sol@v4.9.1


/**
 * @dev Contract module that helps prevent reentrant calls to a function.
 *
 * Inheriting from `ReentrancyGuard` will make the {nonReentrant} modifier
 * available, which can be applied to functions to make sure there are no nested
 * (reentrant) calls to them.
 *
 * Note that because there is a single `nonReentrant` guard, functions marked as
 * `nonReentrant` may not call one another. This can be worked around by making
 * those functions `private`, and then adding `external` `nonReentrant` entry
 * points to them.
 *
 * TIP: If you would like to learn more about reentrancy and alternative ways
 * to protect against it, check out our blog post
 * https://blog.openzeppelin.com/reentrancy-after-istanbul/[Reentrancy After Istanbul].
 */
abstract contract ReentrancyGuard {
    // Booleans are more expensive than uint256 or any type that takes up a full
    // word because each write operation emits an extra SLOAD to first read the
    // slot's contents, replace the bits taken up by the boolean, and then write
    // back. This is the compiler's defense against contract upgrades and
    // pointer aliasing, and it cannot be disabled.

    // The values being non-zero value makes deployment a bit more expensive,
    // but in exchange the refund on every call to nonReentrant will be lower in
    // amount. Since refunds are capped to a percentage of the total
    // transaction's gas, it is best to keep them low in cases like this one, to
    // increase the likelihood of the full refund coming into effect.
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;

    uint256 private _status;

    constructor() {
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and making it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        _nonReentrantBefore();
        _;
        _nonReentrantAfter();
    }

    function _nonReentrantBefore() private {
        // On the first call to nonReentrant, _status will be _NOT_ENTERED
        require(_status != _ENTERED, "ReentrancyGuard: reentrant call");

        // Any calls to nonReentrant after this point will fail
        _status = _ENTERED;
    }

    function _nonReentrantAfter() private {
        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Returns true if the reentrancy guard is currently set to "entered", which indicates there is a
     * `nonReentrant` function in the call stack.
     */
    function _reentrancyGuardEntered() internal view returns (bool) {
        return _status == _ENTERED;
    }
}


// File @openzeppelin/contracts/utils/math/SafeMath.sol@v4.9.1


// CAUTION
// This version of SafeMath should only be used with Solidity 0.8 or later,
// because it relies on the compiler's built in overflow checks.

/**
 * @dev Wrappers over Solidity's arithmetic operations.
 *
 * NOTE: `SafeMath` is generally not needed starting with Solidity 0.8, since the compiler
 * now has built in overflow checking.
 */
library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryAdd(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            uint256 c = a + b;
            if (c < a) return (false, 0);
            return (true, c);
        }
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function trySub(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b > a) return (false, 0);
            return (true, a - b);
        }
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryMul(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
            // benefit is lost if 'b' is also tested.
            // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
            if (a == 0) return (true, 0);
            uint256 c = a * b;
            if (c / a != b) return (false, 0);
            return (true, c);
        }
    }

    /**
     * @dev Returns the division of two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryDiv(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b == 0) return (false, 0);
            return (true, a / b);
        }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryMod(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        unchecked {
            if (b == 0) return (false, 0);
            return (true, a % b);
        }
    }

    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     *
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return a - b;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     *
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        return a * b;
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator.
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return a / b;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return a % b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {trySub}.
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        unchecked {
            require(b <= a, errorMessage);
            return a - b;
        }
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        unchecked {
            require(b > 0, errorMessage);
            return a / b;
        }
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting with custom message when dividing by zero.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryMod}.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        unchecked {
            require(b > 0, errorMessage);
            return a % b;
        }
    }
}


// File contracts/Errors.sol


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


// File contracts/DirectDebit.sol




//   o__ __o         o                                        o           o__ __o                     o             o     o     
//  <|     v\      _<|>_                                     <|>         <|     v\                   <|>          _<|>_  <|>    
//  / \     <\                                               < >         / \     <\                  / >                 < >    
//  \o/       \o     o    \o__ __o     o__  __o       __o__   |          \o/       \o     o__  __o   \o__ __o       o     |     
//   |         |>   <|>    |     |>   /v      |>     />  \    o__/_       |         |>   /v      |>   |     v\     <|>    o__/_ 
//  / \       //    / \   / \   < >  />      //    o/         |          / \       //   />      //   / \     <\    / \    |     
//  \o/      /      \o/   \o/        \o    o/     <|          |          \o/      /     \o    o/     \o/      /    \o/    |     
//   |      o        |     |          v\  /v __o   \\         o           |      o       v\  /v __o   |      o      |     o     
//  / \  __/>       / \   / \          <\/> __/>    _\o__</   <\__       / \  __/>        <\/> __/>  / \  __/>     / \    <\__  
                                                                                                                             
                                                                                                                             
                                                                                                                             


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

/**
   The direct debit contract is a parent contract that implements the basic functionality to create debit accounts
   Deploy only the ConnectedWallets or the VirtualAccounts contracts
 */
abstract contract DirectDebit is
    ReentrancyGuard,
    DirectDebitErrors,
    DirectDebitEvents,
    Pausable,
    Ownable
{
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    IVerifier public immutable verifier; // The verifier address

    /**
     1% fee sent to the deployer of this contract on all transactions. 
     We divide the payout with the ownerFeeDivider
    */
    uint256 public ownerFeeDivider = 100;

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
       The Secret notes are encrypted and stored on-chain, Insipred by Tornado Cash note accounts
       The note's are encrypted with symmetric and asymmetric encryption before stored on chain using a password and an identity provider's public key.
       These notes are used to compute the ZKP (Payment Intent) off-chain that can be used to debit an account
    */
    mapping(bytes32 => string) public encryptedNotes;

    /**
        @dev : the constructor
        @param _verifier is the address of SNARK verifier contract        
    */

    constructor(IVerifier _verifier) {
        verifier = _verifier;
    }

    /**
     @dev calculate the fee used for the payment
     @param amount is the debited amount
   */
    function calculateFee(
        uint256 amount
    ) public view returns (uint256 ownerFee, uint256 payment) {
        ownerFee = amount.div(ownerFeeDivider);
        payment = amount.sub(ownerFee);
    }

    /**
       @dev update the fee calculation
       @param newFeeDivider is the variable used for calculating the fee
     */
    function updateFee(uint256 newFeeDivider) external onlyOwner {
        ownerFeeDivider = newFeeDivider;
    }

    /**
      The Owner can pause and unpause the contract to stop pull payments from it, the withdrawals and wallet disconnecting will still work!
      This is a safety mechanism, in the event of a decryption key leak, an attacker will be able to decrypt the asymmetric encrypted crypto notes,
      However the attacker still needs time to brute force passwords and won't be able to drain balances fast!
      The owner of the smart contract can pause all directdebit and let the account creators disconnect/withdraw their funds.
      Incidents/unauthorized access can be detected automaticly as successful directdebit transactions not sent by the relayer.
     */
    function togglePause() external onlyOwner {
        if (paused()) {
            _unpause();
        } else {
            _pause();
        }
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
      This function can be paused by the owner to disable it!
    */
    function directdebit(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) external nonReentrant whenNotPaused {
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

    /**
       The direct debit transaction is verified using these conditions!
       The verification is slightly different for child contracts so then need to implement it.
    */

    function _verifyPaymentIntent(
        uint256[8] calldata proof,
        bytes32[2] calldata hashes,
        address payee,
        uint256[4] calldata debit
    ) internal virtual;

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
            _processTokenWithdraw(hashes[1], payable(owner()), ownerFee);
        } else {
            // Transfer the eth
            _processEthWithdraw(hashes[1], payee, payment);
            // Send the fee to the owner
            _processEthWithdraw(hashes[1], payable(owner()), ownerFee);
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
    ) internal virtual;

    /**
      Process the eth withdraw and decrease the account balance
    */

    function _processEthWithdraw(
        bytes32 commitment,
        address payee,
        uint256 payment
    ) internal virtual;

    /**
    A view function to get and display the account's balance
    This is useful when the account balance is calculated from external wallet's balance!
     */
    function getAccount(
        bytes32 commitment
    ) external view virtual returns (AccountData memory);
}


// File contracts/ConnectedWallets.sol


//       o__ __o         o__ __o        o          o    o          o    o__ __o__/_       o__ __o    ____o__ __o____   o__ __o__/_   o__ __o            o__ __o__/_       o__ __o               o               o              o           o           o            o            o__ __o__/_  ____o__ __o____ 
//      /v     v\       /v     v\      <|\        <|>  <|\        <|>  <|    v           /v     v\    /   \   /   \   <|    v       <|     v\          <|    v           /v     v\             <|>             <|>            <|>         <|>         <|>          <|>          <|    v        /   \   /   \  
//     />       <\     />       <\     / \\o      / \  / \\o      / \  < >              />       <\        \o/        < >           / \     <\         < >              />       <\            / \             / \            / \         / \         / \          / \          < >                 \o/       
//   o/              o/           \o   \o/ v\     \o/  \o/ v\     \o/   |             o/                    |          |            \o/       \o        |             o/           \o        o/   \o           \o/            \o/       o/   \o       \o/          \o/           |                   |        
//  <|              <|             |>   |   <\     |    |   <\     |    o__/_        <|                    < >         o__/_         |         |>       o__/_        <|             |>      <|__ __|>           |              |       <|__ __|>       |            |            o__/_              < >       
//   \\              \\           //   / \    \o  / \  / \    \o  / \   |             \\                    |          |            / \       //        |             \\           //       /       \          < >            < >      /       \      / \          / \           |                   |        
//     \         /     \         /     \o/     v\ \o/  \o/     v\ \o/  <o>              \         /         o         <o>           \o/      /         <o>              \         /       o/         \o         \o    o/\o    o/     o/         \o    \o/          \o/          <o>                  o        
//      o       o       o       o       |       <\ |    |       <\ |    |                o       o         <|          |             |      o           |                o       o       /v           v\         v\  /v  v\  /v     /v           v\    |            |            |                  <|        
//      <\__ __/>       <\__ __/>      / \        < \  / \        < \  / \  _\o__/_      <\__ __/>         / \        / \  _\o__/_  / \  __/>          / \  _\o__/_      <\__ __/>      />             <\         <\/>    <\/>     />             <\  / \ _\o__/_  / \ _\o__/_  / \  _\o__/_        / \       
                                                                                                                                                                                                                                                                                                           
                                                                                                                                                                                                                                                                                                           
                                                                                                                                                                                                                                                                                                           

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
      A mapping to store if a wallet was connected already with the same token to avoid creating 2 connected wallets. 
      This is useful, because 2 connected wallets with the same address and tokens have the same balance also anyways.
      The key is keccak256(creator,token)
     */
    mapping(bytes32 => bool) public connectedWalletAlready;

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

        bytes32 connectedHash = getHashByAddresses(msg.sender, token);

        if (connectedWalletAlready[connectedHash])
            revert WalletAlreadyConnected();

        connectedWalletAlready[connectedHash] = true;

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

        bytes32 connectedHash = getHashByAddresses(
            accounts[commitment].creator,
            address(accounts[commitment].token)
        );
        connectedWalletAlready[connectedHash] = false;
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

    /**
    A view function to get and display the account's balance
    This is useful when the account balance is calculated from external wallet's balance!
     */
    function getAccount(
        bytes32 commitment
    ) external view override returns (AccountData memory) {
        AccountData memory accountdata = accounts[commitment];
        AccountData memory result;
        result.active = accountdata.active;
        result.creator = accountdata.creator;
        result.token = accountdata.token;

        // Deactivated accounts should have zero balance always!
        if (!accountdata.active) {
            result.balance = 0;
            return result;
        }

        uint256 allowance = IERC20(accountdata.token).allowance(
            accountdata.creator,
            address(this)
        );
        uint256 balance = IERC20(accountdata.token).balanceOf(
            accountdata.creator
        );

        if (allowance > balance) {
            result.balance = balance;
        } else {
            result.balance = allowance;
        }
        return result;
    }

    /**
     This is used with the connectedWalletAlready mapping to access it. Pass in addresses and get the hash!
     */
    function getHashByAddresses(
        address _creator,
        address _token
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_creator, _token));
    }
}
