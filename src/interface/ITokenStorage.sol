// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface ITokenStorage {
    function contains(address token) external view returns (bool);
    function tokensLength() external view returns (uint);
    function tokenByIndex(uint i) external view returns (address);
    function allTokens() external view returns (address[] memory tokens);
}
