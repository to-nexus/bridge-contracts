// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {IBaseBridge} from "./interface/IBaseBridge.sol";
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

    error SBRInvalidAddress();
    error SBRDeadlineExpired();
    error SBRInsufficientOutput();
    error SBRInvalidValue();
    error SBRRefundFailed();
    error SBRInvalidAmountIn();

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
    function swapExactInputSingleAndBridge(ExactInputSingleAndBridgeParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        // Transfer tokenIn from user
        IERC20(params.tokenIn).safeTransferFrom(msg.sender, address(this), params.amountIn);

        // Approve swapRouter
        IERC20(params.tokenIn).forceApprove(address(swapRouter), params.amountIn);

        // Execute swap
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

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, amountOut, params.bridgeParams);

        emit SwapAndBridge(
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
    function swapExactInputAndBridge(ExactInputAndBridgeParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        // Extract tokenIn and tokenOut from path
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, false);

        // Transfer tokenIn from user
        IERC20(tokenIn).safeTransferFrom(msg.sender, address(this), params.amountIn);

        // Approve swapRouter
        IERC20(tokenIn).forceApprove(address(swapRouter), params.amountIn);

        // Execute swap
        amountOut = swapRouter.exactInput(
            ISwapRouter.ExactInputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum
            })
        );

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, amountOut, params.bridgeParams);

        emit SwapAndBridge(
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
    function swapExactOutputSingleAndBridge(ExactOutputSingleAndBridgeParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        // Transfer max tokenIn from user
        IERC20(params.tokenIn).safeTransferFrom(msg.sender, address(this), params.amountInMaximum);

        // Approve swapRouter
        IERC20(params.tokenIn).forceApprove(address(swapRouter), params.amountInMaximum);

        // Execute swap
        amountIn = swapRouter.exactOutputSingle(
            ISwapRouter.ExactOutputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountOut: params.amountOut,
                amountInMaximum: params.amountInMaximum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        // Clear approval
        IERC20(params.tokenIn).forceApprove(address(swapRouter), 0);

        // Refund excess tokenIn
        if (amountIn < params.amountInMaximum) {
            IERC20(params.tokenIn).safeTransfer(msg.sender, params.amountInMaximum - amountIn);
        }

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, params.amountOut, params.bridgeParams);

        emit SwapAndBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            amountIn,
            params.amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapExactOutputAndBridge(ExactOutputAndBridgeParams calldata params, uint deadline)
        external
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        // For exactOutput, path is reversed: tokenOut is first, tokenIn is last
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, true);

        // Transfer max tokenIn from user
        IERC20(tokenIn).safeTransferFrom(msg.sender, address(this), params.amountInMaximum);

        // Approve swapRouter
        IERC20(tokenIn).forceApprove(address(swapRouter), params.amountInMaximum);

        // Execute swap
        amountIn = swapRouter.exactOutput(
            ISwapRouter.ExactOutputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountOut: params.amountOut,
                amountInMaximum: params.amountInMaximum
            })
        );

        // Clear approval
        IERC20(tokenIn).forceApprove(address(swapRouter), 0);

        // Refund excess tokenIn
        if (amountIn < params.amountInMaximum) {
            IERC20(tokenIn).safeTransfer(msg.sender, params.amountInMaximum - amountIn);
        }

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, params.amountOut, params.bridgeParams);

        emit SwapAndBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            amountIn,
            params.amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    // ============ Native Token (ETH) Swap + Bridge Functions ============

    /// @inheritdoc ISwapBridgeRouter
    function swapExactInputSingleETHAndBridge(ExactInputSingleAndBridgeParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        require(msg.value == params.amountIn, SBRInvalidValue());
        require(params.tokenIn == address(WETH9), SBRInvalidAddress());

        // Wrap ETH to WETH
        WETH9.deposit{value: msg.value}();

        // Approve swapRouter
        WETH9.approve(address(swapRouter), params.amountIn);

        // Execute swap
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

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, amountOut, params.bridgeParams);

        emit SwapAndBridge(
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
    function swapExactInputETHAndBridge(ExactInputAndBridgeParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountOut)
    {
        require(msg.value == params.amountIn, SBRInvalidValue());

        // Extract tokenIn and tokenOut from path
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, false);
        require(tokenIn == address(WETH9), SBRInvalidAddress());

        // Wrap ETH to WETH
        WETH9.deposit{value: msg.value}();

        // Approve swapRouter
        WETH9.approve(address(swapRouter), params.amountIn);

        // Execute swap
        amountOut = swapRouter.exactInput(
            ISwapRouter.ExactInputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountIn: params.amountIn,
                amountOutMinimum: params.amountOutMinimum
            })
        );

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, amountOut, params.bridgeParams);

        emit SwapAndBridge(
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
    function swapExactOutputSingleETHAndBridge(ExactOutputSingleAndBridgeParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        require(msg.value >= params.amountInMaximum, SBRInvalidValue());
        require(params.tokenIn == address(WETH9), SBRInvalidAddress());

        // Wrap ETH to WETH
        WETH9.deposit{value: params.amountInMaximum}();

        // Approve swapRouter
        WETH9.approve(address(swapRouter), params.amountInMaximum);

        // Execute swap
        amountIn = swapRouter.exactOutputSingle(
            ISwapRouter.ExactOutputSingleParams({
                tokenIn: params.tokenIn,
                tokenOut: params.tokenOut,
                fee: params.fee,
                recipient: address(this),
                deadline: deadline,
                amountOut: params.amountOut,
                amountInMaximum: params.amountInMaximum,
                sqrtPriceLimitX96: params.sqrtPriceLimitX96
            })
        );

        // Clear approval
        WETH9.approve(address(swapRouter), 0);

        // Refund excess WETH as ETH
        if (amountIn < params.amountInMaximum) {
            uint refundAmount = params.amountInMaximum - amountIn;
            WETH9.withdraw(refundAmount);
            (bool success,) = msg.sender.call{value: refundAmount}("");
            require(success, SBRRefundFailed());
        }

        // Refund any extra ETH sent beyond amountInMaximum
        if (msg.value > params.amountInMaximum) {
            (bool success,) = msg.sender.call{value: msg.value - params.amountInMaximum}("");
            require(success, SBRRefundFailed());
        }

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(params.tokenOut, params.amountOut, params.bridgeParams);

        emit SwapAndBridge(
            msg.sender,
            params.tokenIn,
            params.tokenOut,
            amountIn,
            params.amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
            networkFee,
            exFee
        );
    }

    /// @inheritdoc ISwapBridgeRouter
    function swapExactOutputETHAndBridge(ExactOutputAndBridgeParams calldata params, uint deadline)
        external
        payable
        override
        nonReentrant
        checkDeadline(deadline)
        returns (uint amountIn)
    {
        require(msg.value >= params.amountInMaximum, SBRInvalidValue());

        // For exactOutput, path is reversed: tokenOut is first, tokenIn is last
        (address tokenIn, address tokenOut) = _decodePathTokens(params.path, true);
        require(tokenIn == address(WETH9), SBRInvalidAddress());

        // Wrap ETH to WETH
        WETH9.deposit{value: params.amountInMaximum}();

        // Approve swapRouter
        WETH9.approve(address(swapRouter), params.amountInMaximum);

        // Execute swap
        amountIn = swapRouter.exactOutput(
            ISwapRouter.ExactOutputParams({
                path: params.path,
                recipient: address(this),
                deadline: deadline,
                amountOut: params.amountOut,
                amountInMaximum: params.amountInMaximum
            })
        );

        // Clear approval
        WETH9.approve(address(swapRouter), 0);

        // Refund excess WETH as ETH
        if (amountIn < params.amountInMaximum) {
            uint refundAmount = params.amountInMaximum - amountIn;
            WETH9.withdraw(refundAmount);
            (bool success,) = msg.sender.call{value: refundAmount}("");
            require(success, SBRRefundFailed());
        }

        // Refund any extra ETH sent beyond amountInMaximum
        if (msg.value > params.amountInMaximum) {
            (bool success,) = msg.sender.call{value: msg.value - params.amountInMaximum}("");
            require(success, SBRRefundFailed());
        }

        // Bridge the swapped tokens
        (uint initiateIndex, uint bridgeValue, uint networkFee, uint exFee) =
            _bridgeToken(tokenOut, params.amountOut, params.bridgeParams);

        emit SwapAndBridge(
            msg.sender,
            tokenIn,
            tokenOut,
            amountIn,
            params.amountOut,
            params.bridgeParams.toChainID,
            initiateIndex,
            bridgeValue,
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
        returns (bool ok, uint bridgeValue, uint networkFee, uint exFee)
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
        returns (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee)
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
            return (false, 0, 0, 0, 0);
        }

        // Calculate bridge fees
        (ok, bridgeValue, networkFee, exFee) = _getExpectedBridgeAmount(toChainID, IERC20(tokenOut), swapAmountOut);
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeOutMultihop(uint toChainID, bytes memory path, uint amountIn)
        external
        override
        returns (bool ok, uint swapAmountOut, uint bridgeValue, uint networkFee, uint exFee)
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
            return (false, 0, 0, 0, 0);
        }

        // Calculate bridge fees
        (ok, bridgeValue, networkFee, exFee) = _getExpectedBridgeAmount(toChainID, IERC20(tokenOut), swapAmountOut);
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeIn(uint toChainID, address tokenIn, address tokenOut, uint24 fee, uint bridgeValue)
        external
        override
        returns (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee)
    {
        // Calculate required swap output for desired bridge value
        (ok, swapAmountOut, networkFee, exFee) = _getRequiredSwapOutput(toChainID, IERC20(tokenOut), bridgeValue);
        if (!ok) return (false, 0, 0, 0, 0);

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
            return (false, 0, 0, 0, 0);
        }

        ok = true;
    }

    /// @inheritdoc ISwapBridgeRouter
    function getAmountSwapBridgeInMultihop(uint toChainID, bytes memory path, uint bridgeValue)
        external
        override
        returns (bool ok, uint amountIn, uint swapAmountOut, uint networkFee, uint exFee)
    {
        // For exactOutput path, tokenOut is first 20 bytes
        require(path.length >= 43, SBRInvalidAddress());
        address tokenOut = _extractTokenFromPath(path, true);

        // Calculate required swap output for desired bridge value
        (ok, swapAmountOut, networkFee, exFee) = _getRequiredSwapOutput(toChainID, IERC20(tokenOut), bridgeValue);
        if (!ok) return (false, 0, 0, 0, 0);

        // Get required input from Uniswap V3 QuoterV2
        try quoter.quoteExactOutput(path, swapAmountOut) returns (
            uint _amountIn, uint160[] memory, uint32[] memory, uint
        ) {
            amountIn = _amountIn;
        } catch {
            return (false, 0, 0, 0, 0);
        }

        ok = true;
    }

    // ============ Internal Functions ============

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
        require(path.length >= 43, SBRInvalidAddress()); // minimum: 20 + 3 + 20

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
        require(path.length >= 43, SBRInvalidAddress());

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
     * @return ok Whether the calculation succeeded
     * @return bridgeValue The amount that will be bridged (after fees)
     * @return networkFee The network fee
     * @return exFee The exchange fee
     */
    function _getExpectedBridgeAmount(uint toChainID, IERC20 token, uint totalAmount)
        internal
        view
        returns (bool ok, uint bridgeValue, uint networkFee, uint exFee)
    {
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();

        uint minimumValue;
        uint exFeeRate;
        (minimumValue, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);

        if (totalAmount <= networkFee) return (false, 0, 0, 0);

        uint denominator = bridgeVerifier.denominator();
        uint remaining = totalAmount - networkFee;

        // bridgeValue = remaining * denominator / (denominator + exFeeRate)
        bridgeValue = remaining.mulDiv(denominator, denominator + exFeeRate);
        exFee = bridgeValue.mulDiv(exFeeRate, denominator);

        ok = bridgeValue >= minimumValue;
    }

    /**
     * @notice Internal function to calculate required swap output for desired bridge value
     * @param toChainID Target chain ID
     * @param token Token to bridge
     * @param bridgeValue Desired bridge value (after fees)
     * @return ok Whether the calculation succeeded
     * @return totalAmount Required total amount (swap output) to achieve desired bridge value
     * @return networkFee The network fee
     * @return exFee The exchange fee
     */
    function _getRequiredSwapOutput(uint toChainID, IERC20 token, uint bridgeValue)
        internal
        view
        returns (bool ok, uint totalAmount, uint networkFee, uint exFee)
    {
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();

        uint minimumValue;
        uint exFeeRate;
        (minimumValue, networkFee, exFeeRate) = bridgeVerifier.getTokenConfig(toChainID, token);

        if (bridgeValue < minimumValue) return (false, 0, 0, 0);

        uint denominator = bridgeVerifier.denominator();

        // exFee = bridgeValue * exFeeRate / denominator
        exFee = bridgeValue.mulDiv(exFeeRate, denominator);

        // totalAmount = bridgeValue + exFee + networkFee
        totalAmount = bridgeValue + exFee + networkFee;
        ok = true;
    }

    /**
     * @notice Internal function to bridge tokens after swap
     * @param token Token to bridge
     * @param totalAmount Total amount received from swap
     * @param bridgeParams Bridge parameters
     * @return initiateIndex_ The bridge initiate index for tracking
     * @return bridgeValue_ The amount that will be bridged (after fees)
     * @return networkFee_ The network fee
     * @return exFee_ The exchange fee
     */
    function _bridgeToken(address token, uint totalAmount, SwapAndBridgeParams calldata bridgeParams)
        internal
        returns (uint initiateIndex_, uint bridgeValue_, uint networkFee_, uint exFee_)
    {
        // Calculate fees and bridge value (gas optimized: internal call instead of external)
        bool ok;
        (ok, bridgeValue_, networkFee_, exFee_) =
            _getExpectedBridgeAmount(bridgeParams.toChainID, IERC20(token), totalAmount);
        require(ok, SBRInsufficientOutput());

        // Get the initiate index before bridge call
        initiateIndex_ = bridge.getNextInitiateIndex(bridgeParams.toChainID);

        // Approve bridge to spend tokens
        IERC20(token).forceApprove(address(bridge), totalAmount);

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
    }

    /**
     * @notice Receive function to accept ETH from WETH withdraw
     */
    receive() external payable {
        // Only accept ETH from WETH contract
        require(msg.sender == address(WETH9), SBRInvalidAddress());
    }
}
