// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {BaseBridge} from "./BaseBridge.sol";

import {Const} from "./lib/Const.sol";
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

    /**
     * @notice Emitted when the cross-chain supply limit is updated
     * @param crossSupplyLimit The new maximum supply limit for CROSS native token transfers
     */
    event CrossSupplyLimitSet(uint crossSupplyLimit);

    /// @dev Maximum issuance limit for CROSS native token on the Cross chain
    /// @notice Defines the maximum amount of CROSS native tokens that can be unlocked from the bridge contract.
    /// @notice This limit is designed to be dynamically adjustable to mitigate security risks from potential
    /// @notice bridge exploits, as most of the token supply is initially locked in the bridge contract.
    uint public crossSupplyLimit;
    /// @dev Initial balance of the contract used for cross-chain supply calculations
    /// @notice Stores the initial balance of the bridge contract to accurately track the amount of tokens
    /// @notice that have been unlocked. Since the Cross chain begins with most of its native token supply
    /// @notice locked in the bridge contract, this initial value is essential for calculating the
    /// @notice circulating supply that has been released through bridging from Ethereum.
    uint private _crossInitializeBalance;

    /// @dev Storage gap for future upgrades
    uint[48] private __gap;

    /**
     * @notice Initializes the CrossBridge contract
     * @dev Sets up the contract with initial configuration
     * - Calls the base initialization in BaseBridge
     * - Records the initial balance for CROSS token supply tracking
     * - Sets the initial CROSS token supply limit to 0
     * @param owner_ Address that will receive admin role
     * @param _threshold Minimum number of validators required for validation
     * @param dev_ Address of the developer account for receiving fees
     */
    function initialize(address owner_, uint8 _threshold, address dev_) external override initializer {
        __BaseBridge_init(owner_, _threshold, dev_);
        _crossInitializeBalance = address(this).balance;
        crossSupplyLimit = 0;
    }

    /**
     * @notice Sets the maximum issuance limit for CROSS native token
     * @dev Only callable by admin role
     * - Updates the issuance limit for CROSS native token on the Cross chain
     * - Emits CrossSupplyLimitSet event upon successful update
     * @param _crossSupplyLimit New maximum issuance limit for CROSS native token
     */
    function setCrossSupplyLimit(uint _crossSupplyLimit) external onlyRole(Const.ADMIN_ROLE) {
        crossSupplyLimit = _crossSupplyLimit;
        emit CrossSupplyLimitSet(crossSupplyLimit);
    }

    /**
     * @notice Verifies if a finalization amount is within allowed limits
     * @dev Extends the base implementation with CROSS token issuance limit checks
     * - First performs standard amount verification via parent contract
     * - For CROSS native token, enforces the issuance limit
     * - Calculates current supply by comparing current balance to initial balance
     * @param fromChainID Source chain ID of the transfer
     * @param token Address of the token being transferred (Const.NATIVE_TOKEN for CROSS native token)
     * @param value Amount of tokens to finalize
     * @param retry Whether this is a retry of a previous finalization attempt
     * @return status Status code indicating the result of the check
     * @return delay Boolean indicating if finalization should be delayed
     */
    function _checkFinalizeAmount(uint fromChainID, address token, uint value, bool retry)
        internal
        override
        returns (Const.FinalizeStatus status, bool delay)
    {
        (status, delay) = super._checkFinalizeAmount(fromChainID, token, value, retry);
        if (status != Const.FinalizeStatus.Success) return (status, delay);

        if (token == Const.NATIVE_TOKEN) {
            // Calculate how much CROSS native token has been bridged out (initial balance minus current balance)
            // If current balance is higher than initial (unlikely), default to 0 supply to prevent underflow
            uint supply =
                _crossInitializeBalance >= address(this).balance ? _crossInitializeBalance - address(this).balance : 0;

            // Check if the new transfer would exceed the configured CROSS token issuance limit
            // If limit is exceeded, return a specific error status and mark for delay
            if (supply + value > crossSupplyLimit) return (Const.FinalizeStatus.CrossSupplyLimitExceeded, true);
        }

        return (status, delay);
    }

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
