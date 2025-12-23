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
    /// @notice Deposit ETH and receive WETH
    function deposit() external payable;

    /// @notice Withdraw WETH and receive ETH
    /// @param amount Amount of WETH to withdraw
    function withdraw(uint amount) external;
}

/// @title ISwapBridgeRouter
/// @notice Interface for SwapBridgeRouter that combines Uniswap V3 swaps with bridge operations
/// @dev Supports both exact input and exact output swap modes with single pool and multi-hop paths
interface ISwapBridgeRouter {
    /**
     * @notice Status codes for bridge calculation results
     * @dev Used to provide detailed error information in view/quote functions
     */
    enum QuoteStatus {
        Invalid, // 0 - Default/uninitialized value
        Success, // 1 - Calculation succeeded
        NoPair, // 2 - Token pair not registered for target chain
        InsufficientForFee, // 3 - Amount too low to cover network fee
        InsufficientValue, // 4 - Bridge value below minimum required
        InvalidSwap // 5 - Quoter failed (no pool, insufficient liquidity, etc.)

    }

    /**
     * @notice Bridge parameters for swap and bridge operations
     * @param toChainID Target chain ID for the bridge operation
     * @param recipient Address to receive tokens on the target chain
     * @param extraData Additional data for bridge executor (target contract address + calldata)
     */
    struct BridgeParams {
        uint toChainID;
        address recipient;
        bytes extraData;
    }

    /**
     * @notice Parameters for swapBridgeExactInputSingle (single pool exact input swap + bridge)
     * @param tokenIn Address of the input token to swap
     * @param tokenOut Address of the output token to receive and bridge
     * @param fee Uniswap V3 pool fee tier (e.g., 500, 3000, 10000)
     * @param amountIn Exact amount of tokenIn to swap
     * @param amountOutMinimum Minimum amount of tokenOut to receive from swap (slippage protection)
     * @param sqrtPriceLimitX96 Price limit for the swap (0 for no limit)
     * @param bridgeParams Bridge operation parameters
     */
    struct SwapBridgeExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        uint amountIn;
        uint amountOutMinimum;
        uint160 sqrtPriceLimitX96;
        BridgeParams bridgeParams;
    }

    /**
     * @notice Parameters for swapBridgeExactInput (multi-hop exact input swap + bridge)
     * @param path Encoded swap path (tokenIn + fee + ... + tokenOut)
     * @param amountIn Exact amount of first token to swap
     * @param amountOutMinimum Minimum amount of output token to receive (slippage protection)
     * @param bridgeParams Bridge operation parameters
     */
    struct SwapBridgeExactInputParams {
        bytes path;
        uint amountIn;
        uint amountOutMinimum;
        BridgeParams bridgeParams;
    }

    /**
     * @notice Parameters for swapBridgeExactOutputSingle (single pool exact output swap + bridge)
     * @dev User specifies the exact amount they want to receive on the target chain
     * @param tokenIn Address of the input token to swap
     * @param tokenOut Address of the output token to receive and bridge
     * @param fee Uniswap V3 pool fee tier (e.g., 500, 3000, 10000)
     * @param amountOut Exact amount user wants to receive after bridge (on target chain)
     * @param amountInMaximum Maximum amount of tokenIn willing to spend (slippage protection)
     * @param sqrtPriceLimitX96 Price limit for the swap (0 for no limit)
     * @param bridgeParams Bridge operation parameters
     */
    struct SwapBridgeExactOutputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        uint amountOut;
        uint amountInMaximum;
        uint160 sqrtPriceLimitX96;
        BridgeParams bridgeParams;
    }

    /**
     * @notice Parameters for swapBridgeExactOutput (multi-hop exact output swap + bridge)
     * @dev User specifies the exact amount they want to receive on the target chain
     * @param path Encoded swap path in reverse order (tokenOut + fee + ... + tokenIn)
     * @param amountOut Exact amount user wants to receive after bridge (on target chain)
     * @param amountInMaximum Maximum amount of first token willing to spend (slippage protection)
     * @param bridgeParams Bridge operation parameters
     */
    struct SwapBridgeExactOutputParams {
        bytes path;
        uint amountOut;
        uint amountInMaximum;
        BridgeParams bridgeParams;
    }

    // ============ Events ============

    /**
     * @notice Emitted when a swap and bridge operation is executed
     * @param user Address of the user initiating the swap and bridge
     * @param tokenIn Address of the input token swapped
     * @param tokenOut Address of the output token bridged
     * @param amountIn Amount of tokenIn spent in the swap
     * @param amountOut Amount of tokenOut received from the swap
     * @param toChainID Target chain ID for the bridge
     * @param initiateIndex Index of the bridge initiation
     * @param bridgeValue Amount that will be received on the target chain (after fees)
     * @param networkFee Network fee deducted for the bridge
     * @param exFee Exchange fee deducted for the bridge
     */
    event SwapBridge(
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

    /**
     * @notice Swap exact input single pool and bridge
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountOut The amount of tokenOut received from swap
     */
    function swapBridgeExactInputSingle(SwapBridgeExactInputSingleParams calldata params, uint deadline)
        external
        returns (uint amountOut);

    /**
     * @notice Swap exact input multi-hop and bridge
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountOut The amount of output token received from swap
     */
    function swapBridgeExactInput(SwapBridgeExactInputParams calldata params, uint deadline)
        external
        returns (uint amountOut);

    /**
     * @notice Swap to receive exact bridge output (single pool)
     * @dev Calculates required swap output including bridge fees to achieve exact bridgeValue
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountIn The amount of tokenIn spent
     */
    function swapBridgeExactOutputSingle(SwapBridgeExactOutputSingleParams calldata params, uint deadline)
        external
        returns (uint amountIn);

    /**
     * @notice Swap to receive exact bridge output (multi-hop)
     * @dev Calculates required swap output including bridge fees to achieve exact bridgeValue
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountIn The amount of input token spent
     */
    function swapBridgeExactOutput(SwapBridgeExactOutputParams calldata params, uint deadline)
        external
        returns (uint amountIn);

    // ============ Native Token (ETH) Swap + Bridge Functions ============

    /**
     * @notice Swap exact ETH input single pool and bridge
     * @dev msg.value should match params.amountIn
     * @param params The parameters for the swap and bridge (tokenIn should be WETH)
     * @param deadline Transaction deadline timestamp
     * @return amountOut The amount of tokenOut received from swap
     */
    function swapBridgeExactInputSingleETH(SwapBridgeExactInputSingleParams calldata params, uint deadline)
        external
        payable
        returns (uint amountOut);

    /**
     * @notice Swap exact ETH input multi-hop and bridge
     * @dev msg.value should match params.amountIn
     * @param params The parameters for the swap and bridge (path should start with WETH)
     * @param deadline Transaction deadline timestamp
     * @return amountOut The amount of output token received from swap
     */
    function swapBridgeExactInputETH(SwapBridgeExactInputParams calldata params, uint deadline)
        external
        payable
        returns (uint amountOut);

    /**
     * @notice Swap ETH to receive exact bridge output (single pool)
     * @dev msg.value should be >= params.amountInMaximum, excess is refunded
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountIn The amount of ETH spent (excess refunded)
     */
    function swapBridgeExactOutputSingleETH(SwapBridgeExactOutputSingleParams calldata params, uint deadline)
        external
        payable
        returns (uint amountIn);

    /**
     * @notice Swap ETH to receive exact bridge output (multi-hop)
     * @dev msg.value should be >= params.amountInMaximum, excess is refunded
     * @param params The parameters for the swap and bridge
     * @param deadline Transaction deadline timestamp
     * @return amountIn The amount of ETH spent (excess refunded)
     */
    function swapBridgeExactOutputETH(SwapBridgeExactOutputParams calldata params, uint deadline)
        external
        payable
        returns (uint amountIn);

    // ============ View Functions ============

    /**
     * @notice Calculate expected bridge amount after fees
     * @param toChainID Target chain ID
     * @param token Token to bridge
     * @param totalAmount Total amount before fees
     * @return status Calculation result status (Success, NoPair, InsufficientForFee, InsufficientValue)
     * @return bridgeValue The amount that will be bridged (after fees)
     * @return networkFee The network fee
     * @return exFee The exchange fee
     */
    function getExpectedBridgeAmount(uint toChainID, IERC20 token, uint totalAmount)
        external
        view
        returns (QuoteStatus status, uint bridgeValue, uint networkFee, uint exFee);

    /**
     * @notice Calculate bridge fees for a given bridge value
     * @param toChainID Target chain ID
     * @param token Token to bridge
     * @param value Bridge value (amount user wants to receive, not total amount)
     * @return minimumValue Minimum required value for bridge
     * @return networkFee The network fee
     * @return exFee The exchange fee
     */
    function calculateBridgeFees(uint toChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint networkFee, uint exFee);

    // ============ Quote Functions (Swap + Bridge) ============

    /**
     * @notice Get expected bridge output for a single pool swap
     * @dev Uses Uniswap quoter to simulate swap, then calculates bridge fees
     * @param toChainID Target chain ID for bridge
     * @param tokenIn Token to swap from
     * @param tokenOut Token to swap to (and bridge)
     * @param fee Pool fee tier
     * @param amountIn Amount of tokenIn to swap
     * @return status Calculation result status (Success, NoPair, InsufficientForFee, InsufficientValue)
     * @return swapAmountOut Amount received from swap
     * @return bridgeValue Amount that will be bridged (after fees)
     * @return networkFee Network fee for bridge
     * @return exFee Exchange fee for bridge
     */
    function getAmountSwapBridgeOut(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint amountIn)
        external
        returns (QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee);

    /**
     * @notice Get expected bridge output for a multi-hop swap
     * @dev Uses Uniswap quoter to simulate swap, then calculates bridge fees
     * @param toChainID Target chain ID for bridge
     * @param path The swap path (tokenIn -> fee -> ... -> tokenOut)
     * @param amountIn Amount of first token to swap
     * @return status Calculation result status (Success, NoPair, InsufficientForFee, InsufficientValue)
     * @return swapAmountOut Amount received from swap
     * @return bridgeValue Amount that will be bridged (after fees)
     * @return networkFee Network fee for bridge
     * @return exFee Exchange fee for bridge
     */
    function getAmountSwapBridgeOutMultihop(uint toChainID, bytes memory path, uint amountIn)
        external
        returns (QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee);

    /**
     * @notice Get required swap input for a desired bridge output (single pool)
     * @dev Calculates total swap output needed (bridgeValue + fees), then uses quoter for input
     * @param toChainID Target chain ID for bridge
     * @param tokenIn Token to swap from
     * @param tokenOut Token to swap to (and bridge)
     * @param fee Pool fee tier
     * @param bridgeValue Desired amount to be bridged (after fees)
     * @return status Calculation result status (Success, NoPair, InsufficientForFee, InsufficientValue)
     * @return amountIn Amount of tokenIn required
     * @return swapAmountOut Total amount from swap (before bridge fees)
     * @return networkFee Network fee for bridge
     * @return exFee Exchange fee for bridge
     */
    function getAmountSwapBridgeIn(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint bridgeValue)
        external
        returns (QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee);

    /**
     * @notice Get required swap input for a desired bridge output (multi-hop)
     * @dev Calculates total swap output needed (bridgeValue + fees), then uses quoter for input
     * @param toChainID Target chain ID for bridge
     * @param path The swap path (tokenOut -> fee -> ... -> tokenIn for exactOutput)
     * @param bridgeValue Desired amount to be bridged (after fees)
     * @return status Calculation result status (Success, NoPair, InsufficientForFee, InsufficientValue)
     * @return amountIn Amount of first token required
     * @return swapAmountOut Total amount from swap (before bridge fees)
     * @return networkFee Network fee for bridge
     * @return exFee Exchange fee for bridge
     */
    function getAmountSwapBridgeInMultihop(uint toChainID, bytes memory path, uint bridgeValue)
        external
        returns (QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee);
}
