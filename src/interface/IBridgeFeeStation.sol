// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBridgeFeeStation {
    function denominator() external pure returns (uint);
    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee);
    function estimateMaxValue(uint remoteChainID, IERC20 token, uint totalValue)
        external
        view
        returns (bool ok, uint value, uint gasFee, uint exFee);
}
