// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRelation {
    // 获取直推下级地址
    function getInvList(
        address addr
    ) external view returns (address[] memory _addrsList);

    // 获取直推人数
    function invListLength(address addr) external view returns (uint256);

    // 获取父级
    function Inviter(address _addr) external view returns (address);
}
