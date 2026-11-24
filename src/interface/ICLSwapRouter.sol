// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/// @title ICLSwapRouter
/// @notice Minimal interface of the Slipstream (concentrated liquidity) SwapRouter
/// @dev Mirrors `contracts/periphery/interfaces/ISwapRouter.sol` of the Slipstream periphery.
/// The single-pool structs carry `int24 tickSpacing` where Uniswap V3 carries `uint24 fee`, so the
/// selectors differ (exactInputSingle 0xa026383e vs V3 0x414bf389). Multi-hop paths keep the
/// `token | 3 bytes | token` layout, but the 3 bytes are an int24 tick spacing instead of a fee.
interface ICLSwapRouter {
    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        int24 tickSpacing;
        address recipient;
        uint deadline;
        uint amountIn;
        uint amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }

    function exactInputSingle(ExactInputSingleParams calldata params) external payable returns (uint amountOut);

    struct ExactInputParams {
        bytes path;
        address recipient;
        uint deadline;
        uint amountIn;
        uint amountOutMinimum;
    }

    function exactInput(ExactInputParams calldata params) external payable returns (uint amountOut);

    struct ExactOutputSingleParams {
        address tokenIn;
        address tokenOut;
        int24 tickSpacing;
        address recipient;
        uint deadline;
        uint amountOut;
        uint amountInMaximum;
        uint160 sqrtPriceLimitX96;
    }

    function exactOutputSingle(ExactOutputSingleParams calldata params) external payable returns (uint amountIn);

    struct ExactOutputParams {
        bytes path;
        address recipient;
        uint deadline;
        uint amountOut;
        uint amountInMaximum;
    }

    function exactOutput(ExactOutputParams calldata params) external payable returns (uint amountIn);
}

/// @title ICLQuoterV2
/// @notice Minimal interface of the Slipstream QuoterV2
/// @dev quoteExactInputSingle 0x9e7defe6 (V3 0xc6a5026a). Field order matches V3 with `fee` replaced by
/// `int24 tickSpacing`.
interface ICLQuoterV2 {
    struct QuoteExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint amountIn;
        int24 tickSpacing;
        uint160 sqrtPriceLimitX96;
    }

    struct QuoteExactOutputSingleParams {
        address tokenIn;
        address tokenOut;
        uint amount;
        int24 tickSpacing;
        uint160 sqrtPriceLimitX96;
    }

    function quoteExactInput(bytes memory path, uint amountIn)
        external
        returns (
            uint amountOut,
            uint160[] memory sqrtPriceX96AfterList,
            uint32[] memory initializedTicksCrossedList,
            uint gasEstimate
        );

    function quoteExactInputSingle(QuoteExactInputSingleParams memory params)
        external
        returns (uint amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint gasEstimate);

    function quoteExactOutput(bytes memory path, uint amountOut)
        external
        returns (
            uint amountIn,
            uint160[] memory sqrtPriceX96AfterList,
            uint32[] memory initializedTicksCrossedList,
            uint gasEstimate
        );

    function quoteExactOutputSingle(QuoteExactOutputSingleParams memory params)
        external
        returns (uint amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint gasEstimate);
}
