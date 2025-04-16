// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";

import {Const} from "./lib/Const.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title BSCBridge
 * @notice BSC-side bridge contract for cross-chain token transfers
 * @dev Extends BaseBridge to handle BSC-specific bridge operations
 * - Handles token transfers between BSC and Cross chain
 * - Manages CROSS token pair registration and initial supply
 */
contract BSCBridge is BaseBridge {
    error BSCBridgeCanNotZeroAddress();
    error BSCBridgeCanNotZero();

    /// @dev Storage gap for future upgrades
    uint[50] private __gap;

    /**
     * @notice Initializes the BSCBridge contract
     * @dev Sets up the contract with initial configuration
     * - Calls the base initialization in BaseBridge
     * - Registers CROSS token as a token pair with Cross chain
     * - Sets the initial deposit amount for CROSS tokens
     * @param owner Address that will receive admin role
     * @param dev_ Address of the developer account for receiving fees
     * @param threshold_ Minimum number of validators required for validation
     * @param crossChainID Cross chain ID (e.g., 612055 for Cross chain, or other chain IDs)
     * @param cross Address of the CROSS ERC20 token on this chain
     * @param crossInitialSupply Pre-minted supply of CROSS tokens for the CROSS Foundation
     */
    function initializeBSCBridge(
        address owner,
        address payable dev_,
        uint8 threshold_,
        uint crossChainID,
        address cross,
        uint crossInitialSupply
    ) external initializer {
        require(crossChainID != 0, BSCBridgeCanNotZero());
        require(cross != address(0), BSCBridgeCanNotZeroAddress());
        __BaseBridge_init(owner, dev_, threshold_);

        // Register CROSS token as a token pair
        // This pairs the CROSS ERC20 token on this chain with the Native CROSS token on Cross chain
        _registerToken(crossChainID, true, cross, Const.NATIVE_TOKEN);

        // Deposit the initial supply of CROSS tokens
        if (crossInitialSupply > 0) _depositToken(crossChainID, cross, crossInitialSupply);
    }
}
