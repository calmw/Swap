// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IWithdrawLimit {
    function addStartWithdrawAmount(
        address sender,
        uint256 amount,
        uint256 types
    ) external;
}
