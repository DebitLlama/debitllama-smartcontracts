// SPDX-License-Identifire: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// THIS CONTRACT IS USED ONLY FOR TESTING PURPOSES!!

contract MOCKERC20 is ERC20("USDTMOCK", "USDTM") {
    function mint(address account, uint256 amount) public {
        _mint(account, amount);
    }
}
