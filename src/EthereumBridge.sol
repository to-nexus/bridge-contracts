// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";

import {Const} from "./lib/Const.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title EthereumBridge
 * @notice Ethereum-side bridge contract for cross-chain token transfers
 * @dev Extends BaseBridge to handle Ethereum-specific bridge operations
 * - Handles token transfers between Ethereum and Cross chain
 * - Manages CROSS token pair registration and initial supply
 */
contract EthereumBridge is BaseBridge {
    /**
     * @notice Initializes the EthereumBridge contract
     * @dev Sets up the contract with initial configuration
     * - Calls the base initialization in BaseBridge
     * - Registers CROSS token as a token pair with Cross chain
     * - Sets the initial deposit amount for CROSS tokens
     * @param owner_ Address that will receive admin role
     * @param dev_ Address of the developer account for receiving fees
     * @param _threshold Minimum number of validators required for validation
     */
    function initialize(address owner_, address payable dev_, uint8 _threshold) external override initializer {
        __BaseBridge_init(owner_, dev_, _threshold);

        // Register CROSS token as a token pair
        // This pairs the CROSS ERC20 token on this chain with the Native CROSS token on Cross chain (chain ID 612055)
        _registerToken(612055, true, Const.CROSS_TOKEN, Const.NATIVE_TOKEN);

        // Deposit the initial supply of CROSS tokens
        _depositToken(612055, Const.CROSS_TOKEN, Const.CROSS_INITIAL_SUPPLY);
    }
}
