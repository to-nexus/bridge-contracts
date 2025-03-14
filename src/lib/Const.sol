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
        Success,
        Reverted,
        TokenPaused,
        VerificationAmountThresholdExceeded,
        PeriodTotalValueThresholdExceeded,
        TransferFailed,
        MintFailed,
        TokenScoreOverflow,
        TokenCurrentVolumeOverflow
    }

    /**
     * @notice Special address value representing native token
     * @dev address(1) is a conventional address used to represent native tokens (ETH, MATIC, etc.)
     */
    address internal constant NATIVE_TOKEN = address(1);

    uint constant PERIOD_INTERVAL = 1 hours;

    /**
     * @notice Constants for contract role definitions
     * @dev Role identifiers used for access control
     */
    bytes32 internal constant ADMIN_ROLE = ("ADMIN");
    bytes32 internal constant OPERATOR_ROLE = ("OPERATOR");
    bytes32 internal constant VALIDATOR_ROLE = ("VALIDATOR");
    bytes32 internal constant UPDATOR_ROLE = ("UPDATOR");
    bytes32 internal constant VERIFIER_ROLE = ("VERIFIER");
    bytes32 internal constant BRIDGE_ROLE = ("BRIDGE");
}
