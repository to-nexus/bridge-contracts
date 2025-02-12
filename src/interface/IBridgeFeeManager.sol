// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ITokenStorage} from "../interface/ITokenStorage.sol";

interface IBridgeFeeManager is ITokenStorage {
    struct FeeInfo {
        address token;
        uint gasFee; // fixed
        uint serviceFee; // rate
    }

    function addFeeInfo(address token, uint gasFee, uint serviceFee) external;
    function removeFeeInfo(address token) external;
    function addFeeInfoMany(FeeInfo[] memory feeInfoList) external;
    function updateFeeInfoMany(FeeInfo[] memory feeInfoList) external;
    function removeFeeInfoMany(address[] memory token) external;

    function denominator() external view returns (uint);
    function calculateFee(address token, uint value) external view returns (uint gas, uint service);
    function getTokenFee(address token) external view returns (FeeInfo memory);
    function getTokenFeeList(address[] memory tokens) external view returns (FeeInfo[] memory);
    function allFeeInfo() external view returns (FeeInfo[] memory);
    function feeInfoByIndex(uint index) external view returns (FeeInfo memory);
}
