// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IBridgeExecutor
 * @notice Interface for executing bridge operations with extra data
 * @dev This interface is called by BaseBridge when extradata is provided during finalization
 */
interface IBridgeExecutor {
    /**
     * @notice Executes extra call with bridge finalization data
     * @param fromChainID Source chain ID
     * @param index Bridge operation index
     * @param toToken Token being bridged
     * @param to Original recipient address
     * @param value Amount of tokens
     * @param extraData Target contract address (20 bytes) + calldata
     * @return success Whether the execution was successful
     * @return remaining Amount of tokens remaining (not consumed by target)
     */
    function executeExtraCall(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable returns (bool success, uint remaining);

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
     * @notice Get the current post-call gas reserve value
     * @return The gas reserve amount
     */
    function postCallGasReserve() external view returns (uint);

    /**
     * @notice Set the post-call gas reserve value
     * @param value New gas reserve amount
     */
    function setPostCallGasReserve(uint value) external;

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
