// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ITokenStorage} from "../interface/ITokenStorage.sol";

interface IBridgeTokenInfo is ITokenStorage {
    struct TokenInfo {
        address token;
        uint minimumValue;
        uint gasFee; // fixed
        uint serviceFee; // rate
    }

    function addTokenInfo(address token, uint minimumValue, uint gasFee, uint serviceFee) external;
    function removeTokenInfo(address token) external;
    function addTokenInfoMany(TokenInfo[] memory tokenInfoList) external;
    function updateTokenInfoMany(TokenInfo[] memory tokenInfoList) external;
    function removeTokenInfoMany(address[] memory token) external;

    function denominator() external view returns (uint);
    function calculate(address token, uint value) external view returns (uint minimum, uint gas, uint service);
    function getTokenInfo(address token) external view returns (TokenInfo memory);
    function getTokenInfoList(address[] memory tokens) external view returns (TokenInfo[] memory);
    function allTokenInfo() external view returns (TokenInfo[] memory);
    function tokenInfoByIndex(uint index) external view returns (TokenInfo memory);
}
