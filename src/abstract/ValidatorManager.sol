// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import {Const} from "../lib/Const.sol";
import {RoleManager} from "./RoleManager.sol";
/**
 * @title ValidatorManager
 * @notice Abstract contract for managing bridge validators and signature verification
 * @dev This contract extends Role to provide validator-specific functionality
 * - Signature verification logic using EIP-712
 * - Threshold-based multi-signature validation
 * - Duplicate signature detection
 * - Validator role management using bytes32 role identifiers
 */

abstract contract ValidatorManager is RoleManager, EIP712Upgradeable {
    using ECDSA for bytes32;

    error ValidatorThresholdCanNotZero();
    error ValidatorInsufficientSignature(uint length);
    error ValidatorInvalidSignatures();
    error ValidatorNotAuthorized(address account);

    /**
     * @notice Emitted when the signature threshold is changed
     * @param threshold New threshold value
     */
    event ThresholdChanged(uint8 threshold);

    /// @dev Minimum number of validator signatures required
    /// @custom:storage-location erc7201:cross.bridge.ValidatorManager
    struct ValidatorManagerStorage {
        uint8 _threshold;
    }

    // keccak256(abi.encode(uint256(keccak256("erc7201:cross.bridge.ValidatorManager")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ValidatorManagerStorageLocation =
        0x303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200;

    function _getValidatorManagerStorage() private pure returns (ValidatorManagerStorage storage $) {
        assembly {
            $.slot := ValidatorManagerStorageLocation
        }
    }

    /**
     * @notice Initializes the ValidatorManager
     * @dev Sets up EIP712 domain for signature verification and initializes threshold
     * - Configures EIP-712 domain with name "Validator" and version "1.0.0"
     * - Sets initial threshold value
     * @param threshold_ Required number of validator signatures
     */
    function __Validator_init(uint8 threshold_) internal onlyInitializing {
        require(threshold_ != 0, ValidatorThresholdCanNotZero());
        __EIP712_init("Validator", "1.0.0");

        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._threshold = threshold_;
    }

    /**
     * @notice Returns the current threshold value
     * @return Current threshold value
     */
    function threshold() external view returns (uint8) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._threshold;
    }

    /**
     * @notice Changes the threshold value
     * @dev Updates minimum required signatures and emits event
     * @param threshold_ New threshold value
     */
    function changeThreshold(uint8 threshold_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(threshold_ != 0, ValidatorThresholdCanNotZero());
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._threshold = threshold_;
        emit ThresholdChanged(threshold_);
    }

    /**
     * @notice Validates signatures from validators
     * @dev Verifies that there are enough valid signatures from unique validators
     * - Checks signature array lengths match
     * - Verifies minimum threshold of signatures
     * - Recovers signer addresses from signatures
     * - Validates signers have Validator role using Const.VALIDATOR_ROLE identifier
     * - Detects and ignores duplicate signatures by comparing validator addresses
     * - Optimizes duplicate checking by requiring signatures to be sorted in ascending order by validator address
     * - Ensures final valid signature count meets threshold
     * @param messageHash Message hash to verify
     * @param v Array of signature v values
     * @param r Array of signature r values
     * @param s Array of signature s values
     * @notice The signatures must be sorted in ascending order by validator address for proper duplicate detection
     */
    function _validateSignature(bytes32 messageHash, uint8[] memory v, bytes32[] memory r, bytes32[] memory s)
        internal
        view
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        uint sigsLength = v.length;
        require(sigsLength == r.length && sigsLength == s.length, ValidatorInvalidSignatures());
        require(sigsLength >= $._threshold, ValidatorInsufficientSignature(sigsLength));

        uint valid = 0;
        address before = address(0);
        for (uint i = 0; i < sigsLength; ++i) {
            // Recover signer address using EIP-712 typed data hash
            address validator = _hashTypedDataV4(messageHash).recover(v[i], r[i], s[i]);

            // Optimize duplicate detection by comparing with the previous validator address
            // Note: Signatures MUST be sorted in ascending order by validator address for this to work correctly
            if (before < validator && hasRole(Const.VALIDATOR_ROLE, validator)) {
                unchecked {
                    ++valid;
                }
                before = validator;
            }
        }

        // Ensure enough unique valid signatures
        require(valid >= $._threshold, ValidatorInsufficientSignature(valid));
    }
}
