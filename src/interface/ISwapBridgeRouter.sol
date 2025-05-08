// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;

interface ISwapBridgeRouter {
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
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee);

    function getSwapBridgeInCross(address token, uint amountOut)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee);

    function getSwapBridgeOut(uint toChainID, uint amountIn, address[] memory path)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee);

    function getSwapBridgeIn(uint toChainID, uint amountOut, address[] memory path)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee);
}
