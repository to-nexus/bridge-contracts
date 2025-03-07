pragma solidity 0.8.28;

import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import {RoleManager} from "./RoleManager.sol";

/**
 * @title ValidatorManager
 * @notice Abstract contract for managing bridge validators and signature verification
 * @dev This contract extends RoleManager to provide validator-specific functionality
 * - Signature verification logic using EIP-712
 * - Threshold-based multi-signature validation
 * - Duplicate signature detection
 * - Validator role management using bytes32 role identifiers
 */
abstract contract ValidatorManager is RoleManager, EIP712Upgradeable {
    using ECDSA for bytes32;

    error ValidatorInsufficientSignature(uint length);
    error ValidatorInvalidSignatures(uint v, uint r, uint s);

    /**
     * @notice Emitted when the signature threshold is changed
     * @param threshold New threshold value
     */
    event ThresholdChanged(uint8 threshold);

    /// @dev Constant role identifier for validator role
    bytes32 private constant VALIDATOR_ROLE = bytes32("VALIDATOR");

    /// @dev Minimum number of validator signatures required
    uint8 private _threshold;

    /// @dev Storage gap for future upgrades
    uint[49] private __gap;

    /**
     * @notice Modifier to restrict function calls to validators only
     * @dev Checks if caller has Validator role using the VALIDATOR_ROLE identifier
     */
    modifier onlyValidator() {
        require(hasRole(VALIDATOR_ROLE, _msgSender()), RoleNotAuthorized(VALIDATOR_ROLE, _msgSender()));
        _;
    }

    /**
     * @notice Initializes the ValidatorManager
     * @dev Sets up EIP712 domain for signature verification and initializes threshold
     * - Configures EIP-712 domain with name "Validator" and version "1.0.0"
     * - Sets initial threshold value
     * @param threshold_ Required number of validator signatures
     */
    function __Validator_init(uint8 threshold_) internal onlyInitializing {
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
    function changeThreshold(uint8 threshold_) external onlyOwner {
        _threshold = threshold_;
        emit ThresholdChanged(threshold_);
    }

    /**
     * @notice Validates signatures from validators
     * @dev Verifies that there are enough valid signatures from unique validators
     * - Checks signature array lengths match
     * - Verifies minimum threshold of signatures
     * - Recovers signer addresses from signatures
     * - Validates signers have Validator role using VALIDATOR_ROLE identifier
     * - Detects and ignores duplicate signatures
     * - Ensures final valid signature count meets threshold
     * @param h Message hash to verify
     * @param v Array of signature v values
     * @param r Array of signature r values
     * @param s Array of signature s values
     */
    function _validateSignature(bytes32 h, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) internal view {
        uint sigsLength = v.length;
        require(
            sigsLength == r.length && sigsLength == s.length, ValidatorInvalidSignatures(sigsLength, r.length, s.length)
        );
        require(sigsLength >= _threshold, ValidatorInsufficientSignature(sigsLength));

        uint valid = 0;
        address[] memory _signed = new address[](sigsLength);
        for (uint i = 0; i < sigsLength; ++i) {
            // Recover signer address using EIP-712 typed data hash
            address validator = _hashTypedDataV4(h).recover(v[i], r[i], s[i]);

            // Verify signer has Validator role
            require(hasRole(VALIDATOR_ROLE, validator), RoleNotAuthorized(VALIDATOR_ROLE, validator));

            // Check for duplicate signatures
            bool dup = false;
            for (uint j = 0; j < valid; j++) {
                if (validator == _signed[j]) {
                    dup = true;
                    break;
                }
            }

            // Count unique valid signatures
            if (!dup) {
                _signed[valid] = validator;
                unchecked {
                    ++valid;
                }
            }
        }

        // Ensure enough unique valid signatures
        require(valid >= _threshold, ValidatorInsufficientSignature(valid));
    }
}
