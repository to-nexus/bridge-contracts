// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

// Uniswap V3 official interfaces

import {IPeripheryImmutableState} from "@uniswap/v3-periphery/contracts/interfaces/IPeripheryImmutableState.sol";
import {IQuoterV2} from "@uniswap/v3-periphery/contracts/interfaces/IQuoterV2.sol";
import {ISwapRouter} from "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

/// @title IWETH9
/// @notice Interface for WETH9 (defined locally due to Solidity version incompatibility)
interface IWETH9 is IERC20 {
    function deposit() external payable;
    function withdraw(uint) external;
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
        uint initiateIndex,
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
