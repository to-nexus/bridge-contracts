// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title CrossBridge
 * @notice Implementation of BaseBridge for a specific cross-chain bridge deployment
 * @dev This contract extends BaseBridge with predeploy functionality for specific chain deployments
 * - Uses a predefined address for implementation
 * - Implements proxy verification logic
 * - Ensures contract is only called through a valid proxy
 * - Provides security against direct calls to implementation
 */
contract CrossBridge is BaseBridge {
    /// @dev Predefined address for the predeployed implementation
    address private constant PREDEPLOYED_IMPLEMENTATION_ADDRESS = address(0xb81d6e000000000000000000000000000000C0de);

    /// @dev Storage gap for future upgrades
    uint[50] private __gap;

    /**
     * @notice Verifies the contract is called through a valid proxy
     * @dev Override for predeploy implementation check
     * - Ensures the contract is called via delegatecall
     * - Verifies the implementation address matches the predeploy address
     * - Reverts if called directly or through an invalid proxy
     * - This prevents unauthorized access to the implementation contract
     */
    function _checkProxy() internal view override {
        if (
            address(this) == PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through delegatecall
                || ERC1967Utils.getImplementation() != PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through an active proxy
        ) revert UUPSUnauthorizedCallContext();
    }
}
