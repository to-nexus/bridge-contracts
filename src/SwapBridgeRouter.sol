// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {IBaseBridge} from "./interface/IBaseBridge.sol";
import {IBridgeRegistry} from "./interface/IBridgeRegistry.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {
    IPeripheryImmutableState,
    IQuoterV2,
    ISwapBridgeRouter,
    ISwapRouter,
    IWETH9
} from "./interface/ISwapBridgeRouter.sol";

/**
 * @title SwapBridgeRouter
 * @notice Router for swapping tokens via Uniswap V3 and bridging them in a single transaction
 * @dev Combines Uniswap V3 swap functionality with bridge operations
 *
 * Key features:
 * - Supports both ERC20 and native token (ETH) swaps
 * - Exact input and exact output swap modes
 * - Single pool and multi-hop swaps
 * - Automatic fee calculation for bridge operations
 */
contract SwapBridgeRouter is ReentrancyGuardTransient, ISwapBridgeRouter {
    using SafeERC20 for IERC20;
    using Math for uint;

    /// @notice Thrown when an invalid address (zero address) is provided
    error SBRInvalidAddress();

    /// @notice Thrown when transaction deadline has expired
    error SBRDeadlineExpired();

    /// @notice Thrown when msg.value doesn't match expected value
    error SBRInvalidValue();

    /// @notice Thrown when ETH refund transfer fails
    error SBRRefundFailed();

    /// @notice Thrown when Uniswap V3 path format is invalid
    error SBRInvalidPath();

    /// @notice Thrown when bridge calculation fails (includes failure reason)
    error SBRQuoteFailed(QuoteStatus status);

    /// @notice Uniswap V3 SwapRouter address
    ISwapRouter public immutable swapRouter;

    /// @notice Uniswap V3 QuoterV2 address
    IQuoterV2 public immutable quoter;

    /// @notice Bridge contract address
    IBaseBridge public immutable bridge;

    /// @notice WETH9 contract address
    IWETH9 public immutable WETH9;

    /// @notice Modifier to check deadline
    modifier checkDeadline(uint deadline) {
        require(block.timestamp <= deadline, SBRDeadlineExpired());
        _;
    }

    /**
     * @notice Constructor
     * @param _swapRouter Uniswap V3 SwapRouter address
     * @param _quoter Uniswap V3 Quoter address
     * @param _bridge Bridge contract address
     */
    constructor(address _swapRouter, address _quoter, address _bridge) {
        require(_swapRouter != address(0) && _quoter != address(0) && _bridge != address(0), SBRInvalidAddress());

        swapRouter = ISwapRouter(_swapRouter);
        quoter = IQuoterV2(_quoter);
        bridge = IBaseBridge(_bridge);
        WETH9 = IWETH9(IPeripheryImmutableState(_swapRouter).WETH9());
    }

    // ============ ERC20 Swap + Bridge Functions ============

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactInputSingle(SwapBridgeExactInputSingleParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        _prepareSwap(params.tokenIn, params.amountIn);

        amountOut = swapRouter.exactInputSingle(
            ISwapRouter.ExactInputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        // Refund unspent tokens (when sqrtPriceLimitX96 causes early termination)
        _refundUnspent(params.tokenIn);

        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, amountOut, params.bridgeParams);

        emit SwapBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            params.amountIn,
            amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactInput(SwapBridgeExactInputParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, false);
        _prepareSwap(tokenIn, params.amountIn);

        amountOut = swapRouter.exactInput(
            ISwapRouter.ExactInputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum
            })
        );

        // Refund unspent tokens (safety measure for edge cases)
        _refundUnspent(tokenIn);

        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, amountOut, params.bridgeParams);

        emit SwapBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            params.amountIn,
            amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactOutputSingle(SwapBridgeExactOutputSingleParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        // Calculate fees for exact bridgeValue (no reverse calculation needed)
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint networkFee, uint exFee) =
            bridgeVerifier.calculateFee(params.bridgeParams.toChainID, IERC20(params.tokenOut), params.amountOut);
        require(params.amountOut >= minimumValue, SBRQuoteFailed(QuoteStatus.InsufficientValue));

        uint totalAmount = params.amountOut + networkFee + exFee;

        _prepareSwap(params.tokenIn, params.amountInMaximum);

        amountIn = swapRouter.exactOutputSingle(
            ISwapRouter.ExactOutputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountOut: totalAmount,
                amountInMaximum: params.amountInMaximum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        _refundExcess(params.tokenIn, params.amountInMaximum, amountIn);

        // Bridge with pre-calculated values (no reverse calculation)
        uint initiateIndex = bridge.getNextInitiateIndex(params.bridgeParams.toChainID);
        IERC20(params.tokenOut).forceApprove(address(bridge), totalAmount);
        bridge.bridgeToken(
            params.bridgeParams.toChainID,
            IERC20(params.tokenOut),
            params.bridgeParams.recipient,
            params.amountOut,
            networkFee,
            exFee,
            params.bridgeParams.extraData
        );

        emit SwapBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            amountIn,
            totalAmount,
            params.bridgeParams.toChainID,
            initiateIndex,
            params.amountOut,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactOutput(SwapBridgeExactOutputParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, true);

        // Calculate fees for exact bridgeValue (no reverse calculation needed)
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint networkFee, uint exFee) =
            bridgeVerifier.calculateFee(params.bridgeParams.toChainID, IERC20(tokenOut), params.amountOut);
        require(params.amountOut >= minimumValue, SBRQuoteFailed(QuoteStatus.InsufficientValue));

        uint totalAmount = params.amountOut + networkFee + exFee;

        _prepareSwap(tokenIn, params.amountInMaximum);

        amountIn = swapRouter.exactOutput(
            ISwapRouter.ExactOutputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountOut: totalAmount,
                amountInMaximum: params.amountInMaximum
            })
        );

        _refundExcess(tokenIn, params.amountInMaximum, amountIn);

        // Bridge with pre-calculated values (no reverse calculation)
        uint initiateIndex = bridge.getNextInitiateIndex(params.bridgeParams.toChainID);
        IERC20(tokenOut).forceApprove(address(bridge), totalAmount);
        bridge.bridgeToken(
            params.bridgeParams.toChainID,
            IERC20(tokenOut),
            params.bridgeParams.recipient,
            params.amountOut,
            networkFee,
            exFee,
            params.bridgeParams.extraData
        );

        emit SwapBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            amountIn,
            totalAmount,
            params.bridgeParams.toChainID,
            initiateIndex,
            params.amountOut,
            networkFee,
            exFee
        );
    }

    // ============ Native Token (ETH) Swap + Bridge Functions ============

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactInputSingleETH(SwapBridgeExactInputSingleParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        require(msg.value == params.amountIn, SBRInvalidValue());
        require(params.tokenIn == address(WETH9), SBRInvalidAddress());

        _prepareSwapETH(params.amountIn);

        amountOut = swapRouter.exactInputSingle(
            ISwapRouter.ExactInputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        // Refund unspent WETH as ETH (when sqrtPriceLimitX96 causes early termination)
        _refundUnspentETH();

        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, amountOut, params.bridgeParams);

        emit SwapBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            params.amountIn,
            amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactInputETH(SwapBridgeExactInputParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        require(msg.value == params.amountIn, SBRInvalidValue());

        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, false);
        require(tokenIn == address(WETH9), SBRInvalidAddress());

        _prepareSwapETH(params.amountIn);

        amountOut = swapRouter.exactInput(
            ISwapRouter.ExactInputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum
            })
        );

        // Refund unspent WETH as ETH (safety measure for edge cases)
        _refundUnspentETH();

        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, amountOut, params.bridgeParams);

        emit SwapBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            params.amountIn,
            amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactOutputSingleETH(SwapBridgeExactOutputSingleParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        require(msg.value >= params.amountInMaximum, SBRInvalidValue());
        require(params.tokenIn == address(WETH9), SBRInvalidAddress());

        // Calculate fees for exact bridgeValue (no reverse calculation needed)
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint networkFee, uint exFee) =
            bridgeVerifier.calculateFee(params.bridgeParams.toChainID, IERC20(params.tokenOut), params.amountOut);
        require(params.amountOut >= minimumValue, SBRQuoteFailed(QuoteStatus.InsufficientValue));

        uint totalAmount = params.amountOut + networkFee + exFee;

        _prepareSwapETH(params.amountInMaximum);

        amountIn = swapRouter.exactOutputSingle(
            ISwapRouter.ExactOutputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountOut: totalAmount,
                amountInMaximum: params.amountInMaximum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        _refundExcessETH(params.amountInMaximum, amountIn);

        // Bridge with pre-calculated values (no reverse calculation)
        uint initiateIndex = bridge.getNextInitiateIndex(params.bridgeParams.toChainID);
        IERC20(params.tokenOut).forceApprove(address(bridge), totalAmount);
        bridge.bridgeToken(
            params.bridgeParams.toChainID,
            IERC20(params.tokenOut),
            params.bridgeParams.recipient,
            params.amountOut,
            networkFee,
            exFee,
            params.bridgeParams.extraData
        );

        emit SwapBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            amountIn,
            totalAmount,
            params.bridgeParams.toChainID,
            initiateIndex,
            params.amountOut,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapBridgeExactOutputETH(SwapBridgeExactOutputParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        require(msg.value >= params.amountInMaximum, SBRInvalidValue());

        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, true);
        require(tokenIn == address(WETH9), SBRInvalidAddress());

        // Calculate fees for exact bridgeValue (no reverse calculation needed)
        IBridgeVerifier bridgeVerifier2 = bridge.bridgeVerifier();
        (uint minimumValue2, uint networkFee, uint exFee) =
            bridgeVerifier2.calculateFee(params.bridgeParams.toChainID, IERC20(tokenOut), params.amountOut);
        require(params.amountOut >= minimumValue2, SBRQuoteFailed(QuoteStatus.InsufficientValue));

        uint totalAmount = params.amountOut + networkFee + exFee;

        _prepareSwapETH(params.amountInMaximum);

        amountIn = swapRouter.exactOutput(
            ISwapRouter.ExactOutputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountOut: totalAmount,
                amountInMaximum: params.amountInMaximum
            })
        );

        _refundExcessETH(params.amountInMaximum, amountIn);

        // Bridge with pre-calculated values (no reverse calculation)
        uint initiateIndex = bridge.getNextInitiateIndex(params.bridgeParams.toChainID);
        IERC20(tokenOut).forceApprove(address(bridge), totalAmount);
        bridge.bridgeToken(
            params.bridgeParams.toChainID,
            IERC20(tokenOut),
            params.bridgeParams.recipient,
            params.amountOut,
            networkFee,
            exFee,
            params.bridgeParams.extraData
        );

        emit SwapBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            amountIn,
            totalAmount,
            params.bridgeParams.toChainID,
            initiateIndex,
            params.amountOut,
            networkFee,
            exFee
        );
    }

    // ============ View Functions ============

    /// @inheritdoc ISwapBridgeRouter
    function getExpectedBridgeAmount(uint toChainID, IERC20 token, uint totalAmount)
        external
        view
        override
        returns (QuoteStatus status, uint bridgeValue, uint networkFee, uint exFee)
    {
        return _getExpectedBridgeAmount(toChainID, token, totalAmount);
    }

    /// @inheritdoc ISwapBridgeRouter
    function calculateBridgeFees(uint toChainID, IERC20 token, uint value)
        external
        view
        override
        returns (uint minimumValue, uint networkFee, uint exFee)
    {
        return bridge.bridgeVerifier().calculateFee(toChainID, token, value);
    }

    // ============ Quote Functions (Swap + Bridge) ============

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeOut(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint amountIn)
        external
        override
        returns (QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee)
    {
        // Get swap quote from Uniswap V3 QuoterV2
        try quoter.quoteExactInputSingle(
            IQuoterV2.QuoteExactInputSingleParams({
                tokenIn: tokenIn,
                tokenOut: tokenOut,
                amountIn: amountIn,
                fee: fee,
                sqrtPriceLimitX96: 0
            })
        ) returns (uint _swapAmountOut, uint160, uint32, uint) {
            swapAmountOut = _swapAmountOut;
        } catch {
            return (QuoteStatus.InvalidSwap, 0, 0, 0, 0);
        }

        // Calculate bridge fees
        (status, bridgeValue, networkFee, exFee) = _getExpectedBridgeAmount(toChainID, IERC20(tokenOut), swapAmountOut);
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeOutMultihop(uint toChainID, bytes memory path, uint amountIn)
        external
        override
        returns (QuoteStatus status, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee)
    {
        // Extract tokenOut from path (last 20 bytes)
        require(path.length >= 43, SBRInvalidAddress());
        address tokenOut = _extractTokenFromPath(path, false);

        // Get swap quote from Uniswap V3 QuoterV2
        try quoter.quoteExactInput(path, amountIn) returns (
            uint _swapAmountOut, uint160[] memory, uint32[] memory, uint
        ) {
            swapAmountOut = _swapAmountOut;
        } catch {
            return (QuoteStatus.InvalidSwap, 0, 0, 0, 0);
        }

        // Calculate bridge fees
        (status, bridgeValue, networkFee, exFee) = _getExpectedBridgeAmount(toChainID, IERC20(tokenOut), swapAmountOut);
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeIn(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint bridgeValue)
        external
        override
        returns (QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee)
    {
        // Validate token pair exists for the target chain
        IBridgeRegistry.TokenPair memory pair = bridge.getTokenPair(toChainID, tokenOut);
        if (pair.localToken != tokenOut) return (QuoteStatus.NoPair, 0, 0, 0, 0);

        // Calculate fees for desired bridge value
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        uint minimumValue;
        (minimumValue, networkFee, exFee) = bridgeVerifier.calculateFee(toChainID, IERC20(tokenOut), bridgeValue);
        if (bridgeValue < minimumValue) return (QuoteStatus.InsufficientValue, 0, 0, networkFee, exFee);

        swapAmountOut = bridgeValue + networkFee + exFee;

        // Get required input from Uniswap V3 QuoterV2
        try quoter.quoteExactOutputSingle(
            IQuoterV2.QuoteExactOutputSingleParams({
                tokenIn: tokenIn,
                tokenOut: tokenOut,
                amount: swapAmountOut,
                fee: fee,
                sqrtPriceLimitX96: 0
            })
        ) returns (uint _amountIn, uint160, uint32, uint) {
            amountIn = _amountIn;
        } catch {
            return (QuoteStatus.InvalidSwap, 0, 0, 0, 0);
        }

        status = QuoteStatus.Success;
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeInMultihop(uint toChainID, bytes memory path, uint bridgeValue)
        external
        override
        returns (QuoteStatus status, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee)
    {
        // For exactOutput path, tokenOut is first 20 bytes
        require(path.length >= 43, SBRInvalidAddress());
        address tokenOut = _extractTokenFromPath(path, true);

        // Validate token pair exists for the target chain
        IBridgeRegistry.TokenPair memory pair = bridge.getTokenPair(toChainID, tokenOut);
        if (pair.localToken != tokenOut) return (QuoteStatus.NoPair, 0, 0, 0, 0);

        // Calculate fees for desired bridge value
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        uint minimumValue;
        (minimumValue, networkFee, exFee) = bridgeVerifier.calculateFee(toChainID, IERC20(tokenOut), bridgeValue);
        if (bridgeValue < minimumValue) return (QuoteStatus.InsufficientValue, 0, 0, networkFee, exFee);

        swapAmountOut = bridgeValue + networkFee + exFee;

        // Get required input from Uniswap V3 QuoterV2
        try quoter.quoteExactOutput(path, swapAmountOut) returns (
            uint _amountIn, uint160[] memory, uint32[] memory, uint
        ) {
            amountIn = _amountIn;
        } catch {
            return (QuoteStatus.InvalidSwap, 0, 0, 0, 0);
        }

        status = QuoteStatus.Success;
    }

    // ============ Internal Helper Functions ============

    /**
     * @notice Prepare ERC20 swap by transferring tokens and setting approval
     * @param tokenIn Input token address
     * @param amountIn Amount to transfer and approve
     */
    function _prepareSwap(address tokenIn, uint amountIn) internal {
        IERC20(tokenIn).safeTransferFrom(msg.sender, address(this), amountIn);
        IERC20(tokenIn).forceApprove(address(swapRouter), amountIn);
    }

    /**
     * @notice Prepare ETH swap by wrapping to WETH and setting approval
     * @param amountIn Amount to wrap and approve
     */
    function _prepareSwapETH(uint amountIn) internal {
        WETH9.deposit{value: amountIn}();
        WETH9.approve(address(swapRouter), amountIn);
    }

    /**
     * @notice Refund excess ERC20 tokens after exactOutput swap
     * @param tokenIn Input token address
     * @param amountInMaximum Maximum amount that was approved
     * @param amountIn Actual amount used
     */
    function _refundExcess(address tokenIn, uint amountInMaximum, uint amountIn) internal {
        IERC20(tokenIn).forceApprove(address(swapRouter), 0);
        if (amountIn < amountInMaximum) IERC20(tokenIn).safeTransfer(msg.sender, amountInMaximum - amountIn);
    }

    /**
     * @notice Refund excess ETH after exactOutput swap
     * @param amountInMaximum Maximum amount that was wrapped
     * @param amountIn Actual amount used
     */
    function _refundExcessETH(uint amountInMaximum, uint amountIn) internal {
        WETH9.approve(address(swapRouter), 0);
        if (amountIn < amountInMaximum) {
            uint refundAmount = amountInMaximum - amountIn;
            WETH9.withdraw(refundAmount);
            (bool success,) = msg.sender.call{value: refundAmount}("");
            require(success, SBRRefundFailed());
        }
        // Refund any extra ETH sent beyond amountInMaximum
        if (msg.value > amountInMaximum) {
            (bool success,) = msg.sender.call{value: msg.value - amountInMaximum}("");
            require(success, SBRRefundFailed());
        }
    }

    /**
     * @notice Refund unspent ERC20 tokens after exactInputSingle swap
     * @dev When sqrtPriceLimitX96 is non-zero, swap may terminate early leaving unspent tokens
     * @param tokenIn Input token address
     */
    function _refundUnspent(address tokenIn) internal {
        // Clear allowance
        IERC20(tokenIn).forceApprove(address(swapRouter), 0);
        // Refund any unspent tokens
        uint remaining = IERC20(tokenIn).balanceOf(address(this));
        if (remaining > 0) IERC20(tokenIn).safeTransfer(msg.sender, remaining);
    }

    /**
     * @notice Refund unspent WETH as ETH after exactInputSingle swap
     * @dev When sqrtPriceLimitX96 is non-zero, swap may terminate early leaving unspent WETH
     */
    function _refundUnspentETH() internal {
        // Clear allowance
        WETH9.approve(address(swapRouter), 0);
        // Refund any unspent WETH as ETH
        uint remaining = WETH9.balanceOf(address(this));
        if (remaining > 0) {
            WETH9.withdraw(remaining);
            (bool success,) = msg.sender.call{value: remaining}("");
            require(success, SBRRefundFailed());
        }
    }

    /**
     * @notice Validate V3 path length format
     * @dev V3 path format: 20 + n*(3+20) bytes where n >= 1
     * @param pathLength Length of the path bytes
     */
    function _validatePathLength(uint pathLength) internal pure {
        require(pathLength >= 43, SBRInvalidAddress()); // minimum: 20 + 3 + 20
        require((pathLength - 20) % 23 == 0, SBRInvalidPath()); // V3 path format
    }

    /**
     * @notice Decode tokenIn and tokenOut from path bytes
     * @param path The swap path
     * @param reversed If true, path is in reverse order (exactOutput format)
     * @return tokenIn The input token address
     * @return tokenOut The output token address
     */
    function _decodePathTokens(bytes calldata path, bool reversed)
        internal
        pure
        returns (address tokenIn, address tokenOut)
    {
        _validatePathLength(path.length);

        if (reversed) {
            // exactOutput path: tokenOut -> fee -> ... -> tokenIn
            tokenOut = address(bytes20(path[0:20]));
            tokenIn = address(bytes20(path[path.length - 20:]));
        } else {
            // exactInput path: tokenIn -> fee -> ... -> tokenOut
            tokenIn = address(bytes20(path[0:20]));
            tokenOut = address(bytes20(path[path.length - 20:]));
        }
    }

    /**
     * @notice Extract token address from memory path bytes
     * @param path The swap path (memory)
     * @param first If true, extract first token (bytes 0-20), else last token
     * @return token The extracted token address
     */
    function _extractTokenFromPath(bytes memory path, bool first) internal pure returns (address token) {
        _validatePathLength(path.length);

        if (first) {
            // Extract first 20 bytes (address at the beginning of path)
            assembly {
                // Load 32 bytes starting at path data (skip length prefix)
                // Then shift right by 96 bits (12 bytes) to get the address in lower 160 bits
                token := shr(96, mload(add(path, 32)))
            }
        } else {
            // Extract last 20 bytes (address at the end of path)
            uint offset = path.length - 20;
            assembly {
                token := shr(96, mload(add(add(path, 32), offset)))
            }
        }
    }

    /**
     * @notice Internal function to calculate expected bridge amount after fees
     * @param toChainID Target chain ID
     * @param token Token to bridge
     * @param totalAmount Total amount before fees
     * @return status Calculation result status
     * @return bridgeValue The amount that will be bridged (after fees)
     * @return networkFee The network fee
     * @return exFee The exchange fee
     */
    function _getExpectedBridgeAmount(uint toChainID, IERC20 token, uint totalAmount)
        internal
        view
        returns (QuoteStatus status, uint bridgeValue, uint networkFee, uint exFee)
    {
        // Validate token pair exists for the target chain
        IBridgeRegistry.TokenPair memory pair = bridge.getTokenPair(toChainID, address(token));
        if (pair.localToken != address(token)) return (QuoteStatus.NoPair, 0, 0, 0);

        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();

        uint minimumValue;
        uint exFeeRate;
        (minimumValue, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);

        if (totalAmount <= networkFee) return (QuoteStatus.InsufficientForFee, 0, networkFee, 0);

        uint denominator = bridgeVerifier.denominator();
        uint remaining = totalAmount - networkFee;

        // bridgeValue = remaining * denominator / (denominator + exFeeRate)
        bridgeValue = remaining.mulDiv(denominator, denominator + exFeeRate);
        exFee = bridgeValue.mulDiv(exFeeRate, denominator);

        status = bridgeValue >= minimumValue ? QuoteStatus.Success : QuoteStatus.InsufficientValue;
    }

    /**
     * /**
     * @notice Internal function to bridge tokens after swap
     * @param token Token to bridge
     * @param totalAmount Total amount received from swap
     * @param bridgeParams Bridge parameters
     * @return initiateIndex_ The bridge initiate index for tracking
     * @return bridgeValue_ The amount that will be bridged (after fees)
     * @return networkFee_ The network fee
     * @return exFee_ The exchange fee
     */
    function _bridgeToken(address token, uint totalAmount, BridgeParams calldata bridgeParams)
        internal
        returns (uint initiateIndex_, uint bridgeValue_, uint networkFee_, uint exFee_)
    {
        // Calculate fees and bridge value (gas optimized: internal call instead of external)
        QuoteStatus status;
        (status, bridgeValue_, networkFee_, exFee_) =
            _getExpectedBridgeAmount(bridgeParams.toChainID, IERC20(token), totalAmount);
        require(status == QuoteStatus.Success, SBRQuoteFailed(status));

        // Rounding dust handling:
        // - Quote math uses floor division, so `bridgeValue + networkFee + exFee` can be < totalAmount.
        // - We cannot "stuff" dust into `exFee` because BaseBridge clamps fees to the minimum calculated fee.
        // Strategy:
        // - Bridge only the spendable amount (`bridgeValue + minimum fees`)
        // - Refund any leftover dust back to the user to avoid locking tokens in this router
        uint spendAmount = bridgeValue_ + networkFee_ + exFee_;

        // Get the initiate index before bridge call
        initiateIndex_ = bridge.getNextInitiateIndex(bridgeParams.toChainID);

        // Approve bridge to spend tokens
        IERC20(token).forceApprove(address(bridge), spendAmount);

        // Execute bridge
        bridge.bridgeToken(
            bridgeParams.toChainID,
            IERC20(token),
            bridgeParams.recipient,
            bridgeValue_,
            networkFee_,
            exFee_,
            bridgeParams.extraData
        );

        // Refund any remaining dust so it doesn't get stuck in the router
        if (totalAmount > spendAmount) IERC20(token).safeTransfer(msg.sender, totalAmount - spendAmount);
    }

    /**
     * @notice Receive function to accept ETH from WETH withdraw
     */
    receive() external payable {
        // Only accept ETH from WETH contract
        require(msg.sender == address(WETH9), SBRInvalidAddress());
    }
}
