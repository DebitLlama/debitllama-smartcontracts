// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Address.sol";

// A helper smart contract to send the relayer gas
// This is used to fund transaction relaying.
// The relayer tracks the submitted transactions externally but uses this convenience contract to trigger on deposit events
// The relayer is operated by the DebitLlama service. It is a centralized relayer.

contract RelayerGasTracker is Ownable {
    using SafeMath for uint256;

    // The address of the relayer
    address payable public relayer;

    // The total amount sent by a payee, the relayer has it's own database of transaction history to compare this to
    mapping(address => uint256) public total;

    // The event that is emitted when a relayer is topped up!
    event TopUpEvent(address indexed from, uint256 indexed amount);

    // The owner can update the relayer address any time
    function setRelayer(address to) external onlyOwner {
        relayer = payable(to);
    }

    // The relayer is topped up by the payee
    // This will trigger a TopUpEvent to track the payments!
    function topUpRelayer() external payable {
        total[msg.sender] += msg.value;
        Address.sendValue(relayer, msg.value);
        emit TopUpEvent(msg.sender, msg.value);
    }
}
