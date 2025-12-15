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
     */
    function executeExtraCall(
        uint fromChainID,
        uint index,
        IERC20 toToken,
        address to,
        uint value,
        bytes calldata extraData
    ) external payable returns (bool success);

    function isWhitelistedTarget(address target) external view returns (bool);
}
