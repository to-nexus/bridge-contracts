// SPDX-License-Identifier: MIT
pragma solidity >=0.6.2;

/// @notice Chainlink price oracle
interface ISwapBridgeRouter {
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
        address[] calldata path,
        uint deadline
    ) external;

    function swapExactETHForTokensBridge(
        uint toChainID,
        address to,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] calldata path,
        uint deadline
    ) external payable;

    function swapTokensForExactETHBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint amountInMax,
        uint maxNetworkFee,
        uint maxExFee,
        address[] calldata path,
        uint deadline
    ) external;

    function swapExactTokensForETHBridge(
        uint toChainID,
        address to,
        uint amountIn,
        uint amountOutMin,
        uint maxNetworkFee,
        uint maxExFee,
        address[] calldata path,
        uint deadline
    ) external payable;

    function swapETHForExactTokensBridge(
        uint toChainID,
        address to,
        uint amountOut,
        uint maxNetworkFee,
        uint maxExFee,
        address[] calldata path,
        uint deadline
    ) external payable;

    function getSwapBridgesOut(uint toChainID, uint amountIn, address[] calldata path)
        external
        view
        returns (uint[] memory amounts, uint bridgeValue, uint networkFee, uint exFee);

    function getSwapBridgesIn(uint toChainID, uint amountOut, address[] calldata path)
        external
        view
        returns (uint[] memory amounts, uint networkFee, uint exFee);
}
