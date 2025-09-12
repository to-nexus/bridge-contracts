// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {IBridgeRegistry} from "./interface/IBridgeRegistry.sol";

import {Const} from "./lib/Const.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title BSCBridgeV2
 * @notice BSC-side bridge contract for cross-chain token transfers
 * @dev Extends BaseBridge to handle BSC-specific bridge operations
 * - Handles token transfers between BSC and Cross chain
 * - Manages CROSS token pair registration and initial supply
 */
contract BSCBridgeV2 is BaseBridge {
    using SafeERC20 for IERC20;

    error BSCBridgeCanNotZeroAddress();
    error BSCBridgeCanNotZero();
    error BSCBridgeInvalidDeadWallet();

    uint private crossChainID;
    IERC20 private cross;

    mapping(address => bool) private deadWallets;

    /// @dev Storage gap for future upgrades
    uint[47] private __gap;

    /**
     * @notice Reinitializes the BSCBridge contract
     * @param crossChainID_ Cross chain ID (e.g., 612055 for Cross chain, or other chain IDs)
     * @param cross_ Address of the CROSS ERC20 token on this chain
     */
    function reinitializeBSCBridge(uint crossChainID_, IERC20 cross_)
        external
        onlyRole(Const.ADMIN_ROLE)
        reinitializer(2)
    {
        require(crossChainID_ != 0, BSCBridgeCanNotZero());
        require(address(cross_) != address(0), BSCBridgeCanNotZeroAddress());

        crossChainID = crossChainID_;
        cross = cross_;

        deadWallets[address(0x000000000000000000000000000000000000dEaD)] = true;
        deadWallets[address(0xdeaDDeADDEaDdeaDdEAddEADDEAdDeadDEADDEaD)] = true;
        deadWallets[address(0xdEad000000000000000000000000000000000000)] = true;
    }

    /**
     * @notice Burn CROSS tokens by sending them to a dead wallet
     * @dev Emits BridgeInitiated to coordinate a cross-chain burn.
     * Semantics of `alreadyTransferred`:
     * - When `alreadyTransferred == true`:
     *   - Caller MUST have ADMIN role.
     *   - No token transfer occurs on this chain.
     *   - Emitted BridgeInitiated uses `from = address(this)`.
     * - When `alreadyTransferred == false`:
     *   - No special role required.
     *   - Transfers caller's CROSS tokens to `deadWallet` on this chain.
     *   - Emitted BridgeInitiated uses `from = _msgSender()`.
     * @param deadWallet The address of the dead wallet
     * @param amount The amount of CROSS tokens to burn on the cross chain
     * @param alreadyTransferred Whether tokens were already transferred to the dead wallet on this chain
     */
    function burnCrossToDeadWallet(address deadWallet, uint amount, bool alreadyTransferred)
        external
        whenNotPaused
        nonReentrant
        returns (bool)
    {
        require(deadWallets[deadWallet], BSCBridgeInvalidDeadWallet());
        require(amount > 0, BSCBridgeCanNotZero());

        address from;
        if (!alreadyTransferred) {
            cross.safeTransferFrom(_msgSender(), deadWallet, amount);
            from = _msgSender();
        } else {
            _checkRole(Const.ADMIN_ROLE, _msgSender());
            from = address(this);
        }

        uint index = _useInitiateIndex(crossChainID);
        emit BridgeInitiated(
            crossChainID,
            index,
            IERC20(cross),
            IERC20(Const.NATIVE_TOKEN),
            from,
            deadWallet,
            amount,
            0,
            0,
            "",
            block.timestamp
        );

        return true;
    }
}
