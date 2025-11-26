// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IBridgeReceiver
 * @notice Interface for contracts that want to receive bridge finalization callbacks
 * @dev Contracts implementing this interface can execute custom logic after receiving bridged tokens
 */
interface IBridgeReceiver {
    /**
     * @notice Called after bridge finalization is complete
     * @dev This function is called by the bridge contract after tokens are transferred/minted
     * @param fromChainID Source chain identifier
     * @param index Bridge operation index
     * @param token Token that was received
     * @param amount Amount of tokens received
     * @param extraData Additional data from the bridge initiator
     * @return selector Must return IBridgeReceiver.onBridgeReceived.selector for validation
     */
    function onBridgeReceived(uint fromChainID, uint index, IERC20 token, uint amount, bytes calldata extraData)
        external
        returns (bytes4);
}
