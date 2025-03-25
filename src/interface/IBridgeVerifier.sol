// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {Const} from "../lib/Const.sol";
import {IPriceFeed} from "./IPriceFeed.sol";

interface IBridgeVerifier {
    function priceFeed() external view returns (IPriceFeed);
    function validateBridgeTokenValue(IERC20 token, uint value) external returns (Const.FinalizeStatus status);
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint networkFee, uint exFee);
    function getTokenPrice(IERC20 token) external view returns (bool exist, uint price);
    function getTokenPriceWithValue(IERC20 token, uint value) external view returns (bool exist, uint price);
    function getGasPrice(uint remoteChainID) external view returns (uint);
    function getTokenConfig(uint remoteChainID, IERC20 token)
        external
        view
        returns (uint minimumValue, uint networkFee, uint exFeeRate);
    function getMinimumTokenValue() external view returns (uint);
    function getVerificationAmountThreshold() external view returns (uint);
    function getTimeWindow() external view returns (uint);
    function getPeriodTotalValueThreshold() external view returns (uint);
    function getTokenCurrentScore(IERC20 token) external view returns (uint);
    function getTokenMovementHistory(IERC20 token) external view returns (bytes32[] memory history);
    function denominator() external pure returns (uint);
}
