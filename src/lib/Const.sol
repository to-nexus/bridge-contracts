// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * @title Const
 * @notice Constants and enums used throughout the contracts
 * @dev All constants and enums are centralized here and imported by other contracts
 */
library Const {
    /**
     * @notice Enum defining bridge pending status types
     * @dev Values indicating reasons why a bridge operation cannot be executed immediately
     */
    enum FinalizeStatus {
        None,
        Success,
        TokenPaused,
        VerificationAmountThresholdExceeded,
        PeriodTotalValueThresholdExceeded,
        TransferFailed,
        MintFailed,
        TokenScoreOverflow,
        TokenCurrentVolumeOverflow,
        CrossSupplyLimitExceeded
    }

    /**
     * @notice Special address value representing native token
     * @dev address(1) is a conventional address used to represent native tokens (ETH, MATIC, etc.)
     */
    address internal constant NATIVE_TOKEN = address(1);

    /// @dev The interval for the period of time during which the total value of transfers is calculated
    uint internal constant PERIOD_INTERVAL = 1 hours;

    /**
     * @notice Constants for contract role definitions
     * @dev Role identifiers used for access control
     */
    bytes32 internal constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 internal constant VALIDATOR_ROLE = keccak256("VALIDATOR_ROLE");
    bytes32 internal constant EDITOR_ROLE = keccak256("EDITOR_ROLE");
    bytes32 internal constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");
    bytes32 internal constant PRICER_ROLE = keccak256("PRICER_ROLE");
    bytes32 internal constant VERIFIER_ROLE = keccak256("VERIFIER_ROLE");
    bytes32 internal constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");
    bytes32 internal constant MINTER_ROLE = keccak256("MINTER_ROLE");
}
