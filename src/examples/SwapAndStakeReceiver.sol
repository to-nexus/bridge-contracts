// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {IBridgeReceiver} from "../interface/IBridgeReceiver.sol";

/**
 * @title SwapAndStakeReceiver
 * @notice Example implementation of IBridgeReceiver that swaps and stakes tokens after bridge
 * @dev Demonstrates post-finalize action pattern
 *
 * Use case:
 * 1. User bridges USDT from BSC to Cross Chain
 * 2. This contract receives USDT
 * 3. Automatically swaps USDT to CROSS token
 * 4. Stakes CROSS tokens for the user
 */
contract SwapAndStakeReceiver is IBridgeReceiver, Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    error SwapAndStakeInvalidBridge();
    error SwapAndStakeInsufficientOutput();
    error SwapAndStakeAlreadyProcessed();

    event BridgeReceived(uint indexed fromChainID, uint indexed index, IERC20 indexed token, address user, uint amount);
    event Swapped(address indexed user, IERC20 tokenIn, IERC20 tokenOut, uint amountIn, uint amountOut);
    event Staked(address indexed user, uint amount, uint lockPeriod);
    event ReceiverFailed(uint indexed fromChainID, uint indexed index, string reason);

    /// @notice Allowed bridge contract
    address public immutable bridge;

    /// @notice Mock swap router (in production, use real DEX)
    ISwapRouter public swapRouter;

    /// @notice Mock staking contract (in production, use real staking)
    IStakingContract public stakingContract;

    /// @notice Track processed bridges to prevent replay
    mapping(bytes32 => bool) public processedBridges;

    /**
     * @notice Constructor
     * @param _bridge Bridge contract address
     * @param _swapRouter Swap router address
     * @param _stakingContract Staking contract address
     */
    constructor(address _bridge, address _swapRouter, address _stakingContract) Ownable(msg.sender) {
        require(_bridge != address(0), SwapAndStakeInvalidBridge());
        bridge = _bridge;
        swapRouter = ISwapRouter(_swapRouter);
        stakingContract = IStakingContract(_stakingContract);
    }

    /**
     * @notice Called by bridge after finalization
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param token Token received
     * @param amount Amount received
     * @param extraData Encoded: (address user, address tokenOut, uint minAmountOut, uint lockPeriod)
     */
    function onBridgeReceived(uint fromChainID, uint index, IERC20 token, uint amount, bytes calldata extraData)
        external
        override
        nonReentrant
        returns (bytes4)
    {
        // 1. Only bridge contract can call
        require(msg.sender == bridge, SwapAndStakeInvalidBridge());

        // 2. Prevent replay attacks
        bytes32 bridgeId = keccak256(abi.encodePacked(fromChainID, index));
        require(!processedBridges[bridgeId], SwapAndStakeAlreadyProcessed());
        processedBridges[bridgeId] = true;

        emit BridgeReceived(fromChainID, index, token, msg.sender, amount);

        // 3. Decode extraData
        (address user, address tokenOut, uint minAmountOut, uint lockPeriod) =
            abi.decode(extraData, (address, address, uint, uint));

        try this._processSwapAndStake(token, amount, user, tokenOut, minAmountOut, lockPeriod) {
            // Success
        } catch Error(string memory reason) {
            emit ReceiverFailed(fromChainID, index, reason);
            // Transfer tokens directly to user if processing fails
            token.safeTransfer(user, amount);
        } catch (bytes memory) {
            emit ReceiverFailed(fromChainID, index, "Unknown error");
            // Transfer tokens directly to user if processing fails
            token.safeTransfer(user, amount);
        }

        return IBridgeReceiver.onBridgeReceived.selector;
    }

    /**
     * @notice Internal function to process swap and stake
     * @dev Separated for better error handling
     */
    function _processSwapAndStake(
        IERC20 tokenIn,
        uint amountIn,
        address user,
        address tokenOut,
        uint minAmountOut,
        uint lockPeriod
    ) external {
        require(msg.sender == address(this), "Only self");

        // 1. Approve router
        tokenIn.forceApprove(address(swapRouter), amountIn);

        // 2. Swap
        uint amountOut = swapRouter.swap(address(tokenIn), tokenOut, amountIn, minAmountOut);
        require(amountOut >= minAmountOut, SwapAndStakeInsufficientOutput());

        emit Swapped(user, tokenIn, IERC20(tokenOut), amountIn, amountOut);

        // 3. Approve staking contract
        IERC20(tokenOut).forceApprove(address(stakingContract), amountOut);

        // 4. Stake for user
        stakingContract.stakeFor(user, tokenOut, amountOut, lockPeriod);

        emit Staked(user, amountOut, lockPeriod);
    }

    /**
     * @notice Update swap router
     */
    function setSwapRouter(address _swapRouter) external onlyOwner {
        swapRouter = ISwapRouter(_swapRouter);
    }

    /**
     * @notice Update staking contract
     */
    function setStakingContract(address _stakingContract) external onlyOwner {
        stakingContract = IStakingContract(_stakingContract);
    }

    /**
     * @notice Emergency withdraw (owner only)
     */
    function emergencyWithdraw(IERC20 token, uint amount) external onlyOwner {
        token.safeTransfer(msg.sender, amount);
    }
}

/**
 * @notice Mock swap router interface
 */
interface ISwapRouter {
    function swap(address tokenIn, address tokenOut, uint amountIn, uint minAmountOut)
        external
        returns (uint amountOut);
}

/**
 * @notice Mock staking contract interface
 */
interface IStakingContract {
    function stakeFor(address user, address token, uint amount, uint lockPeriod) external;
}

