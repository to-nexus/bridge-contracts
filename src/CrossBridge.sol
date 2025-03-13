// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title CrossBridge
 * @notice Cross-chain bridge with predeploy functionality
 * @dev Extends BaseBridge with deterministic deployment features
 * - Uses predefined implementation address for easier verification
 * - Implements proxy security checks
 */
contract CrossBridge is BaseBridge {
    /// @dev Predefined address for the predeployed implementation
    address private constant PREDEPLOYED_IMPLEMENTATION_ADDRESS = address(0xb81d6e000000000000000000000000000000C0de);

    /// @dev Storage gap for future upgrades
    uint[50] private __gap;

    /**
     * @notice Verifies proxy delegation
     * @dev Ensures contract is called via proper proxy
     * - Prevents direct calls to implementation
     * - Verifies implementation address matches predeploy address
     */
    function _checkProxy() internal view override {
        if (
            address(this) == PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through delegatecall
                || ERC1967Utils.getImplementation() != PREDEPLOYED_IMPLEMENTATION_ADDRESS // Must be called through an active proxy
        ) revert UUPSUnauthorizedCallContext();
    }
}
