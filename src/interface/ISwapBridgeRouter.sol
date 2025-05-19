// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;

import {IBaseBridge} from "./IBaseBridge.sol";
import {IPancakeRouter02} from "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

interface ISwapBridgeRouter {
    /// @notice Enum defining swap types with exact inputs or outputs
    enum SwapType {
        EXACT_TOKEN_FOR_TOKEN,
        EXACT_ETH_FOR_TOKEN,
        EXACT_TOKEN_FOR_ETH,
        TOKEN_FOR_EXACT_TOKEN,
        ETH_FOR_EXACT_TOKEN,
        TOKEN_FOR_EXACT_ETH
    }

    /// @notice Struct for fee data to reduce stack variables
    struct FeeData {
        uint bridgeValue;
        uint networkFee;
        uint exFee;
    }

    function swapBridgeExactTokensForCrossTokens(
        address sourceToken,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external;

    function swapBridgeTokensForExactCrossTokens(
        address sourceToken,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external;

    function swapBridgeExactETHForCrossTokens(
        address to,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external payable;

    function swapBridgeETHForExactCrossTokens(
        address to,
        uint amountOut,
        uint networkFeeMax,
        uint exFeeMax,
        uint deadline
    ) external payable;

    function swapBridgeExactTokensForTokens(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external;

    function swapBridgeTokensForExactTokens(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external;

    function swapBridgeExactETHForTokens(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external payable;

    function swapBridgeTokensForExactETH(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external;

    function swapBridgeExactTokensForETH(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external payable;

    function swapBridgeETHForExactTokens(
        uint toChainID,
        address to,
        uint amountOut,
        uint networkFeeMax,
        uint exFeeMax,
        address[] memory path,
        uint deadline
    ) external payable;

    function WETH() external view returns (address);
    function BRIDGE() external view returns (IBaseBridge);
    function SWAP_ROUTER() external view returns (IPancakeRouter02);

    function getSourceTokens() external view returns (address[] memory);
    function getPath(address token) external view returns (address[] memory);
    function getSwapBridgeOutCross(address token, uint amountIn)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee, uint priceImpactBps);
    function getSwapBridgeInCross(address token, uint amountOut)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee, uint priceImpactBps);
    function getSwapBridgeOut(uint toChainID, uint amountIn, address[] memory path)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee, uint priceImpactBps);
    function getSwapBridgeIn(uint toChainID, uint amountOut, address[] memory path)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee, uint priceImpactBps);
    function getReserves(address tokenA, address tokenB) external view returns (uint112 reserveA, uint112 reserveB);
    function quote(uint amountA, uint reserveA, uint reserveB) external view returns (uint amountB);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external view returns (uint amountOut);
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external view returns (uint amountIn);
    function getAmountsOut(uint amountIn, address[] memory path) external view returns (uint[] memory amounts);
    function getAmountsIn(uint amountOut, address[] memory path) external view returns (uint[] memory amounts);
    function getPriceImpact(address[] memory path, uint amountIn) external view returns (uint priceImpactBps);
}
