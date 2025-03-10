// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBridgeFeeManager {
    struct TokenFee {
        IERC20 token;
        uint minimumValue;
        uint exFeeRate;
    }

    function denominator() external pure returns (uint);
    function getGasPrice(uint remoteChainID) external view returns (uint);
    function getTokenFee(uint remoteChainID, IERC20 token)
        external
        view
        returns (uint minimumValue, uint gasFee, uint exFeeRate);
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint gasFee, uint exFee);
}
