// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @title IWETH9
/// @notice Interface for WETH9
interface IWETH9 is IERC20 {
    function deposit() external payable;
    function withdraw(uint) external;
}

/// @title IUniswapV3Quoter
/// @notice Minimal interface for Uniswap V3 Quoter
/// @dev Used for getting swap quotes without executing the swap
interface IUniswapV3Quoter {
    /// @notice Returns the amount out received for a given exact input swap without executing the swap
    /// @param tokenIn The token being swapped in
    /// @param tokenOut The token being swapped out
    /// @param fee The fee of the token pool to consider for the pair
    /// @param amountIn The amount of tokenIn to swap
    /// @param sqrtPriceLimitX96 The price limit of the pool that cannot be exceeded by the swap
    /// @return amountOut The amount of tokenOut that would be received
    function quoteExactInputSingle(
        address tokenIn,
        address tokenOut,
        uint24 fee,
        uint amountIn,
        uint160 sqrtPriceLimitX96
    ) external returns (uint amountOut);

    /// @notice Returns the amount out received for a given exact input but for a swap of multiple pools
    /// @param path The path of the swap, i.e. each token pair and the pool fee
    /// @param amountIn The amount of the first token to swap
    /// @return amountOut The amount of the last token that would be received
    function quoteExactInput(bytes memory path, uint amountIn) external returns (uint amountOut);

    /// @notice Returns the amount in required for a given exact output swap without executing the swap
    /// @param tokenIn The token being swapped in
    /// @param tokenOut The token being swapped out
    /// @param fee The fee of the token pool to consider for the pair
    /// @param amountOut The desired amount of tokenOut
    /// @param sqrtPriceLimitX96 The price limit of the pool that cannot be exceeded by the swap
    /// @return amountIn The amount of tokenIn that would be required
    function quoteExactOutputSingle(
        address tokenIn,
        address tokenOut,
        uint24 fee,
        uint amountOut,
        uint160 sqrtPriceLimitX96
    ) external returns (uint amountIn);

    /// @notice Returns the amount in required to receive the given exact output amount but for a swap of multiple pools
    /// @param path The path of the swap, i.e. each token pair and the pool fee (reversed)
    /// @param amountOut The amount of the last token to receive
    /// @return amountIn The amount of first token required to be paid
    function quoteExactOutput(bytes memory path, uint amountOut) external returns (uint amountIn);
}

/// @title IUniswapV3SwapRouter
/// @notice Minimal interface for Uniswap V3 SwapRouter
/// @dev Includes WETH9() from PeripheryImmutableState
interface IUniswapV3SwapRouter {
    /// @notice Returns the address of WETH9
    function WETH9() external view returns (address);

    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint deadline;
        uint amountIn;
        uint amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }

    struct ExactInputParams {
        bytes path;
        address recipient;
        uint deadline;
        uint amountIn;
        uint amountOutMinimum;
    }

    struct ExactOutputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint deadline;
        uint amountOut;
        uint amountInMaximum;
        uint160 sqrtPriceLimitX96;
    }

    struct ExactOutputParams {
        bytes path;
        address recipient;
        uint deadline;
        uint amountOut;
        uint amountInMaximum;
    }

    function exactInputSingle(ExactInputSingleParams calldata params) external payable returns (uint amountOut);
    function exactInput(ExactInputParams calldata params) external payable returns (uint amountOut);
    function exactOutputSingle(ExactOutputSingleParams calldata params) external payable returns (uint amountIn);
    function exactOutput(ExactOutputParams calldata params) external payable returns (uint amountIn);
}

/// @title ISwapBridgeRouter
/// @notice Interface for SwapBridgeRouter that combines Uniswap V3 swaps with bridge operations
interface ISwapBridgeRouter {
    /// @notice Parameters for swap and bridge operations
    struct SwapAndBridgeParams {
        uint toChainID;
        address recipient;
        bytes extraData;
    }

    /// @notice Parameters for exact input single swap and bridge
    struct ExactInputSingleAndBridgeParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        uint amountIn;
        uint amountOutMinimum;
        uint160 sqrtPriceLimitX96;
        SwapAndBridgeParams bridgeParams;
    }

    /// @notice Parameters for exact input multi-hop swap and bridge
    struct ExactInputAndBridgeParams {
        bytes path;
        uint amountIn;
        uint amountOutMinimum;
        SwapAndBridgeParams bridgeParams;
    }

    /// @notice Parameters for exact output single swap and bridge
    struct ExactOutputSingleAndBridgeParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        uint amountOut;
        uint amountInMaximum;
        uint160 sqrtPriceLimitX96;
        SwapAndBridgeParams bridgeParams;
    }

    /// @notice Parameters for exact output multi-hop swap and bridge
    struct ExactOutputAndBridgeParams {
        bytes path;
        uint amountOut;
        uint amountInMaximum;
        SwapAndBridgeParams bridgeParams;
    }

    // ============ Events ============

    event SwapAndBridge(
        address indexed user,
        address indexed tokenIn,
        address indexed tokenOut,
        uint amountIn,
        uint amountOut,
        uint toChainID,
        uint bridgeValue,
        uint networkFee,
        uint exFee
    );

    // ============ ERC20 Swap + Bridge Functions ============

    /// @notice Swap exact input single pool and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountOut The amount of tokenOut received from swap
    function swapExactInputSingleAndBridge(ExactInputSingleAndBridgeParams calldata params, uint deadline)
        external
        returns (uint amountOut);

    /// @notice Swap exact input multi-hop and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountOut The amount of output token received from swap
    function swapExactInputAndBridge(ExactInputAndBridgeParams calldata params, uint deadline)
        external
        returns (uint amountOut);

    /// @notice Swap exact output single pool and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountIn The amount of tokenIn spent
    function swapExactOutputSingleAndBridge(ExactOutputSingleAndBridgeParams calldata params, uint deadline)
        external
        returns (uint amountIn);

    /// @notice Swap exact output multi-hop and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountIn The amount of input token spent
    function swapExactOutputAndBridge(ExactOutputAndBridgeParams calldata params, uint deadline)
        external
        returns (uint amountIn);

    // ============ Native Token (ETH) Swap + Bridge Functions ============

    /// @notice Swap exact ETH input single pool and bridge
    /// @param params The parameters for the swap and bridge (tokenIn should be WETH)
    /// @param deadline Transaction deadline
    /// @return amountOut The amount of tokenOut received from swap
    function swapExactInputSingleETHAndBridge(ExactInputSingleAndBridgeParams calldata params, uint deadline)
        external
        payable
        returns (uint amountOut);

    /// @notice Swap exact ETH input multi-hop and bridge
    /// @param params The parameters for the swap and bridge (path should start with WETH)
    /// @param deadline Transaction deadline
    /// @return amountOut The amount of output token received from swap
    function swapExactInputETHAndBridge(ExactInputAndBridgeParams calldata params, uint deadline)
        external
        payable
        returns (uint amountOut);

    /// @notice Swap exact output with ETH input single pool and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountIn The amount of ETH spent (excess refunded)
    function swapExactOutputSingleETHAndBridge(ExactOutputSingleAndBridgeParams calldata params, uint deadline)
        external
        payable
        returns (uint amountIn);

    /// @notice Swap exact output with ETH input multi-hop and bridge
    /// @param params The parameters for the swap and bridge
    /// @param deadline Transaction deadline
    /// @return amountIn The amount of ETH spent (excess refunded)
    function swapExactOutputETHAndBridge(ExactOutputAndBridgeParams calldata params, uint deadline)
        external
        payable
        returns (uint amountIn);

    // ============ View Functions ============

    /// @notice Calculate expected bridge amount after fees
    /// @param toChainID Target chain ID
    /// @param token Token to bridge
    /// @param totalAmount Total amount before fees
    /// @return ok Whether the calculation succeeded
    /// @return bridgeValue The amount that will be bridged (after fees)
    /// @return networkFee The network fee
    /// @return exFee The exchange fee
    function getExpectedBridgeAmount(uint toChainID, IERC20 token, uint totalAmount)
        external
        view
        returns (bool ok, uint bridgeValue, uint networkFee, uint exFee);

    /// @notice Calculate bridge fees for a given amount
    /// @param toChainID Target chain ID
    /// @param token Token to bridge
    /// @param value Bridge value (not total amount)
    /// @return minimumValue Minimum required value
    /// @return networkFee The network fee
    /// @return exFee The exchange fee
    function calculateBridgeFees(uint toChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint networkFee, uint exFee);

    // ============ Quote Functions (Swap + Bridge) ============

    /// @notice Get expected bridge output for a single pool swap
    /// @param toChainID Target chain ID for bridge
    /// @param tokenIn Token to swap from
    /// @param tokenOut Token to swap to (and bridge)
    /// @param fee Pool fee tier
    /// @param amountIn Amount of tokenIn to swap
    /// @return ok Whether the operation would succeed
    /// @return swapAmountOut Amount received from swap
    /// @return bridgeValue Amount that will be bridged (after fees)
    /// @return networkFee Network fee for bridge
    /// @return exFee Exchange fee for bridge
    function getAmountSwapBridgeOut(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint amountIn)
        external
        returns (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee);

    /// @notice Get expected bridge output for a multi-hop swap
    /// @param toChainID Target chain ID for bridge
    /// @param path The swap path (tokenIn -> fee -> ... -> tokenOut)
    /// @param amountIn Amount of first token to swap
    /// @return ok Whether the operation would succeed
    /// @return swapAmountOut Amount received from swap
    /// @return bridgeValue Amount that will be bridged (after fees)
    /// @return networkFee Network fee for bridge
    /// @return exFee Exchange fee for bridge
    function getAmountSwapBridgeOutMultihop(uint toChainID, bytes memory path, uint amountIn)
        external
        returns (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee);

    /// @notice Get required swap input for a desired bridge output (single pool)
    /// @param toChainID Target chain ID for bridge
    /// @param tokenIn Token to swap from
    /// @param tokenOut Token to swap to (and bridge)
    /// @param fee Pool fee tier
    /// @param bridgeValue Desired amount to be bridged (after fees)
    /// @return ok Whether the operation would succeed
    /// @return amountIn Amount of tokenIn required
    /// @return swapAmountOut Total amount from swap (before bridge fees)
    /// @return networkFee Network fee for bridge
    /// @return exFee Exchange fee for bridge
    function getAmountSwapBridgeIn(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint bridgeValue)
        external
        returns (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee);

    /// @notice Get required swap input for a desired bridge output (multi-hop)
    /// @param toChainID Target chain ID for bridge
    /// @param path The swap path (tokenOut -> fee -> ... -> tokenIn for exactOutput)
    /// @param bridgeValue Desired amount to be bridged (after fees)
    /// @return ok Whether the operation would succeed
    /// @return amountIn Amount of first token required
    /// @return swapAmountOut Total amount from swap (before bridge fees)
    /// @return networkFee Network fee for bridge
    /// @return exFee Exchange fee for bridge
    function getAmountSwapBridgeInMultihop(uint toChainID, bytes memory path, uint bridgeValue)
        external
        returns (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee);
}
