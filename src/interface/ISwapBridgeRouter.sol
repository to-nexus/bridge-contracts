// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;

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

    function swapExactTokensForCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external;

    function swapTokensForExactCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external;

    function swapExactETHForCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external payable;

    function swapETHForExactCrossTokensBridge(
        address tokenForSwap,
        address to,
        uint amountOut,
        uint maxNetworkFee,
        uint maxExFee,
        uint deadline
    ) external payable;

    function swapExactTokensForTokensBridge(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external;

    function swapTokensForExactTokensBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external;

    function swapExactETHForTokensBridge(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external payable;

    function swapTokensForExactETHBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external;

    function swapExactTokensForETHBridge(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external payable;

    function swapETHForExactTokensBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint maxNetworkFee,
        uint maxExFee,
        address[] memory path,
        uint deadline
    ) external payable;

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
