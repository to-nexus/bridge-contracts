// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IBridgeExecutor
 * @notice Interface for executing bridge operations with extra data
 * @dev This interface is called by BaseBridge when extradata is provided during finalization
 *
 * Flow:
 * 1. BaseBridge approves executor to spend tokens (ERC20) or sends native value
 * 2. Executor pulls tokens via transferFrom (ERC20) or receives native value
 * 3. Executor calls target contract with tokens
 * 4. Executor returns consumed amount; remaining is sent to user by executor
 * 5. If executor reverts, entire transaction rolls back (safe for bridge)
 */
interface IBridgeExecutor {
    /**
     * @notice Executes extra call with bridge finalization data
     * @dev This function WILL revert on any failure. On revert, bridge falls back to normal transfer.
     *      - For ERC20: pulls tokens from msg.sender via transferFrom
     *      - For Native: receives value via msg.value
     *      - Returns consumed amount; sends remaining to user directly
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param toToken Token being bridged
     * @param to Original recipient address
     * @param value Amount of tokens
     * @param extraData Target contract address (20 bytes) + calldata
     * @return consumed Amount of tokens actually consumed by target
     * @return returnData Return data from target call (capped at 1KB)
     */
    function executeExtraCall(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable returns (uint consumed, bytes memory returnData);

    /**
     * @notice Check if a target contract is whitelisted
     * @param target Target contract address to check
     * @return True if target is whitelisted
     */
    function isWhitelistedTarget(address target) external view returns (bool);

    /**
     * @notice Check if method check is enabled for a target
     * @param target Target contract address to check
     * @return True if method check is enabled
     */
    function isMethodCheckEnabled(address target) external view returns (bool);

    /**
     * @notice Check if a method selector is whitelisted for a target
     * @param target Target contract address
     * @param methodID Function selector
     * @return True if method is whitelisted
     */
    function isWhitelistedMethod(address target, bytes4 methodID) external view returns (bool);

    /**
     * @notice Add multiple method selectors to the whitelist for a target
     * @param target Target contract address
     * @param methodIDs Array of function selectors to whitelist
     */
    function addWhitelistMethods(address target, bytes4[] calldata methodIDs) external;

    /**
     * @notice Remove multiple method selectors from the whitelist for a target
     * @param target Target contract address
     * @param methodIDs Array of function selectors to remove
     */
    function removeWhitelistMethods(address target, bytes4[] calldata methodIDs) external;
}
