// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IEarlybird {
    function addBscPledge(address from, uint256 amount, uint256 types) external;
}
