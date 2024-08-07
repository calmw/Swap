// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface IBscBalance {
    function addBscBalance(address from, uint256 amount, uint256 types) external ;
}
