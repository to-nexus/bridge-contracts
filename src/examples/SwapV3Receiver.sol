// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {IBridgeReceiver} from "../interface/IBridgeReceiver.sol";

/**
 * @title SwapV3Receiver
 * @notice IBridgeReceiver implementation that swaps tokens using Uniswap V3 after bridge
 * @dev Automatically swaps all bridged tokens to a target token using Uniswap V3 SwapRouter
 *
 * Use case:
 * 1. User bridges USDT from BSC to Cross Chain
 * 2. This contract receives USDT
 * 3. Automatically swaps 100% USDT to target token (e.g., WETH, CROSS)
 * 4. Sends swapped tokens to user
 * 5. If swap fails, sends original USDT to user
 */
contract SwapV3Receiver is IBridgeReceiver, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    error SwapV3InvalidBridge();
    error SwapV3InvalidRouter();
    error SwapV3AlreadyProcessed();
    error SwapV3InsufficientOutput();
    error SwapV3InvalidBalanceState();

    event BridgeReceived(uint indexed fromChainID, uint indexed index, IERC20 indexed token, address to, uint amount);
    event Swapped(address indexed to, IERC20 tokenIn, IERC20 tokenOut, uint amountIn, uint amountOut);
    event SwapFailed(uint indexed fromChainID, uint indexed index, string reason);
    event TokenTransferred(address indexed to, IERC20 indexed token, uint amount);
    event SwapRouterUpdated(address indexed oldRouter, address indexed newRouter);

    /// @notice Allowed bridge contract
    address public immutable bridge;

    /// @notice Uniswap V3 SwapRouter
    ISwapRouter public swapRouter;

    /// @notice Track processed bridges to prevent replay
    mapping(bytes32 => bool) public processedBridges;

    /**
     * @notice Constructor
     * @param _bridge Bridge contract address
     * @param _swapRouter Uniswap V3 SwapRouter address
     */
    constructor(address _bridge, address _swapRouter) Ownable(msg.sender) {
        require(_bridge != address(0), SwapV3InvalidBridge());
        require(_swapRouter != address(0), SwapV3InvalidRouter());
        bridge = _bridge;
        swapRouter = ISwapRouter(_swapRouter);
    }

    /**
     * @notice Called by bridge after finalization
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param token Token received
     * @param amount Amount received
     * @param extraData Encoded: (address to, address tokenOut, uint24 fee, uint amountOutMinimum, uint deadline)
     * @return selector IBridgeReceiver.onBridgeReceived.selector
     */
    function onBridgeReceived(uint fromChainID, uint index, IERC20 token, uint amount, bytes calldata extraData)
        external
        override
        nonReentrant
        returns (bytes4)
    {
        // 1. Only bridge contract can call
        require(msg.sender == bridge, SwapV3InvalidBridge());

        // 2. Prevent replay attacks
        bytes32 bridgeId = keccak256(abi.encodePacked(fromChainID, index));
        require(!processedBridges[bridgeId], SwapV3AlreadyProcessed());
        processedBridges[bridgeId] = true;

        emit BridgeReceived(fromChainID, index, token, msg.sender, amount);

        // 3. Decode extraData
        (address to, address tokenOut, uint24 fee, uint amountOutMinimum, uint deadline) =
            abi.decode(extraData, (address, address, uint24, uint, uint));

        // Validate addresses
        require(to != address(0), "Invalid recipient");
        require(tokenOut != address(0), "Invalid tokenOut");

        // Validate deadline
        require(deadline >= block.timestamp, "Deadline expired");

        // 4. Try swap, if fails transfer original token to user
        try this._processSwap(token, amount, to, tokenOut, fee, amountOutMinimum, deadline) {
            // Success - swap completed
        } catch Error(string memory reason) {
            emit SwapFailed(fromChainID, index, reason);
            // Transfer original token to user
            token.safeTransfer(to, amount);
            emit TokenTransferred(to, token, amount);
        } catch (bytes memory) {
            emit SwapFailed(fromChainID, index, "Unknown error");
            // Transfer original token to user
            token.safeTransfer(to, amount);
            emit TokenTransferred(to, token, amount);
        }

        return IBridgeReceiver.onBridgeReceived.selector;
    }

    /**
     * @notice Internal function to process swap
     * @dev Separated for better error handling with try-catch
     * - Uses balance comparison to ensure only this transaction's tokens are refunded
     * - Prevents mixing with accidentally sent tokens
     * - Refunds exact remainder if swap consumes less than amountIn
     * @param tokenIn Input token
     * @param amountIn Input amount
     * @param to Recipient address
     * @param tokenOut Output token
     * @param fee Pool fee (500, 3000, 10000)
     * @param amountOutMinimum Minimum output amount (slippage protection)
     * @param deadline Swap deadline timestamp
     */
    function _processSwap(
        IERC20 tokenIn,
        uint amountIn,
        address to,
        address tokenOut,
        uint24 fee,
        uint amountOutMinimum,
        uint deadline
    ) external {
        require(msg.sender == address(this), "Only self");

        // 1. Record balance before swap
        uint balanceBefore = tokenIn.balanceOf(address(this));

        // 2. Approve SwapRouter
        tokenIn.forceApprove(address(swapRouter), amountIn);

        // 3. Setup swap parameters
        ISwapRouter.ExactInputSingleParams memory params = ISwapRouter.ExactInputSingleParams({
            tokenIn: address(tokenIn),
            tokenOut: tokenOut,
            fee: fee,
            recipient: to,
            deadline: deadline,
            amountIn: amountIn,
            amountOutMinimum: amountOutMinimum,
            sqrtPriceLimitX96: 0 // No price limit
        });

        // 4. Execute swap
        uint amountOut = swapRouter.exactInputSingle(params);

        // 5. Verify minimum output
        require(amountOut >= amountOutMinimum, SwapV3InsufficientOutput());

        emit Swapped(to, tokenIn, IERC20(tokenOut), amountIn, amountOut);

        // 6. Calculate and return only the tokens from this transaction
        uint balanceAfter = tokenIn.balanceOf(address(this));

        // Safety check: balance should not increase after swap
        if (balanceAfter > balanceBefore) revert SwapV3InvalidBalanceState();

        uint consumed = balanceBefore - balanceAfter;

        // If swap consumed less than amountIn, refund the difference
        if (consumed < amountIn) {
            uint refund = amountIn - consumed;
            tokenIn.safeTransfer(to, refund);
        }
    }

    /**
     * @notice Update swap router
     * @param _swapRouter New SwapRouter address
     */
    function setSwapRouter(address _swapRouter) external onlyOwner {
        require(_swapRouter != address(0), SwapV3InvalidRouter());
        address oldRouter = address(swapRouter);
        swapRouter = ISwapRouter(_swapRouter);
        emit SwapRouterUpdated(oldRouter, _swapRouter);
    }

    /**
     * @notice Emergency withdraw (owner only)
     * @param token Token to withdraw
     * @param amount Amount to withdraw
     */
    function emergencyWithdraw(IERC20 token, uint amount) external onlyOwner {
        token.safeTransfer(msg.sender, amount);
    }
}

/**
 * @notice Uniswap V3 SwapRouter interface
 * @dev Minimal interface for exactInputSingle function
 */
interface ISwapRouter {
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

    /// @notice Swaps amountIn of one token for as much as possible of another token
    /// @param params The parameters necessary for the swap, encoded as ExactInputSingleParams in calldata
    /// @return amountOut The amount of the received token
    function exactInputSingle(ExactInputSingleParams calldata params) external payable returns (uint amountOut);
}
