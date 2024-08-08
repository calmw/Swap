// contracts/GLDToken.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ToxA is RoleAccess, ERC20 {
    constructor(uint256 initialSupply) ERC20("INTOVERSE TOKEN", "TOXA") {
        _addAdmin(msg.sender);
        _mint(msg.sender, initialSupply);
    }

    function mint(address account, uint256 amount) public onlyAdmin {
        _mint(account, amount);
    }
}
