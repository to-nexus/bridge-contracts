// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface ICrossMintableERC20Code {
    function createCrossMintableERC20(uint remoteChainID, address remoteToken, string memory symbol, uint8 decimals)
        external
        returns (address tokenAddress);
}
