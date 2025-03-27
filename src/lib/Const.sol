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

    /**
     * @notice CROSS token contract address on Cross chain
     * @dev Used for token pairing between Ethereum and Cross chains
     */
    address internal constant CROSS_TOKEN = address(0x5061C090bf18246890F88AB504Cd562632f83faa);

    /**
     * @notice Initial supply of CROSS tokens for the CROSS Foundation
     * @dev Represents 10 million tokens with 18 decimals (ether denomination)
     * This is the pre-minted amount allocated at contract initialization
     */
    uint internal constant CROSS_INITIAL_SUPPLY = 10_000_000 ether;

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
}
