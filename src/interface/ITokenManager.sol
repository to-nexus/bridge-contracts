// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ITokenStorage} from "./ITokenStorage.sol";

interface ITokenManager is ITokenStorage {
    function isValidToken(address token) external view returns (bool);
    function getPairToken(address token) external view returns (address);
    function allPairs() external view returns (address[] memory pairs);
}
