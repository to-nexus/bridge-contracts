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
    error ValidatorInvalidSignatures(uint vLength, uint rLength, uint sLength);
    error ValidatorNotAuthorized(address account);

    /**
     * @notice Emitted when the signature threshold is changed
     * @param threshold New threshold value
     */
    event ThresholdChanged(uint8 threshold);

    /// @dev Minimum number of validator signatures required
    uint8 private _threshold;

    /// @dev Storage gap for future upgrades
    uint[49] private __gap;

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
        _threshold = threshold_;
    }

    /**
     * @notice Returns the current threshold value
     * @return Current threshold value
     */
    function threshold() external view returns (uint8) {
        return _threshold;
    }

    /**
     * @notice Changes the threshold value
     * @dev Updates minimum required signatures and emits event
     * @param threshold_ New threshold value
     */
    function changeThreshold(uint8 threshold_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(threshold_ != 0, ValidatorThresholdCanNotZero());
        _threshold = threshold_;
        emit ThresholdChanged(threshold_);
    }

    /**
     * @notice Validates signatures from validators
     * @dev Verifies that there are enough valid signatures from unique validators
     * - Checks signature array lengths match
     * - Verifies minimum threshold of signatures
     * - Recovers signer addresses from signatures
     * - Validates signers have Validator role using Const.VALIDATOR_ROLE identifier
     * - Detects and ignores duplicate signatures
     * - Ensures final valid signature count meets threshold
     * @param messageHash Message hash to verify
     * @param v Array of signature v values
     * @param r Array of signature r values
     * @param s Array of signature s values
     */
    function _validateSignature(bytes32 messageHash, uint8[] memory v, bytes32[] memory r, bytes32[] memory s)
        internal
        view
    {
        uint sigsLength = v.length;
        require(
            sigsLength == r.length && sigsLength == s.length, ValidatorInvalidSignatures(sigsLength, r.length, s.length)
        );
        require(sigsLength >= _threshold, ValidatorInsufficientSignature(sigsLength));

        uint valid = 0;
        address[] memory signed = new address[](sigsLength);
        for (uint i = 0; i < sigsLength; ++i) {
            // Recover signer address using EIP-712 typed data hash
            address validator = _hashTypedDataV4(messageHash).recover(v[i], r[i], s[i]);

            // Check for duplicate signatures
            bool dup = false;
            for (uint j = 0; j < valid; ++j) {
                if (validator == signed[j]) {
                    dup = true;
                    break;
                }
            }

            // Count unique valid signatures
            if (!dup && hasRole(Const.VALIDATOR_ROLE, validator)) {
                signed[valid] = validator;
                unchecked {
                    ++valid;
                }
            }
        }

        // Ensure enough unique valid signatures
        require(valid >= _threshold, ValidatorInsufficientSignature(valid));
    }
}
