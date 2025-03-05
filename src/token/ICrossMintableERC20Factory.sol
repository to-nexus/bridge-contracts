// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ICrossMintableERC20} from "./ICrossMintableERC20.sol";

interface ICrossMintableERC20Factory {
    function createCrossMintableERC20(bytes32 salt, string memory name, string memory symbol, uint8 decimals)
        external
        returns (address tokenAddress);

    function allTokens() external view returns (address[] memory);
    function tokenByIndex(uint index) external view returns (address);
    function tokensLength() external view returns (uint);
}
