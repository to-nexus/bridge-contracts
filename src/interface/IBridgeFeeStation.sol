// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBridgeFeeStation {
    struct TokenFee {
        IERC20 token;
        uint minimumAmount;
        uint exFeeRate;
    }

    struct GasInfo {
        uint chainID;
        uint gasPrice;
        IERC20 gasToken;
    }

    function denominator() external pure returns (uint);
    function getGasInfo(uint remoteChainID) external view returns (GasInfo memory);
    function getTokenFee(uint remoteChainID, IERC20 token)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFeeRate);
    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee);
}
