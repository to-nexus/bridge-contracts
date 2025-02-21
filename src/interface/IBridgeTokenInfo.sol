// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBridgeTokenInfo {
    function denominator() external pure returns (uint);
    function getTokenInfo(IERC20 token) external view returns (uint minimum, uint gas, uint ex, bool isValid);
    function calculate(IERC20 token, uint value)
        external
        view
        returns (uint minimum, uint gas, uint ex, bool isValid);
    function calculateMax(IERC20 token, uint totalValue)
        external
        view
        returns (uint value, uint gas, uint ex, bool isValid);
}
