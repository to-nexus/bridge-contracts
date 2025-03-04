pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IValidatorManager} from "../interface/IValidatorManager.sol";

abstract contract ValidatorManager is OwnableUpgradeable, EIP712Upgradeable, IValidatorManager {
    using ECDSA for bytes32;
    using EnumerableSet for EnumerableSet.AddressSet;

    error errValidatorManagerNotValidator(address account);
    error errValidatorManagerAlreadyExistValidator(address account);
    error errValidatorManagerNotExistValidator(address account);
    error errValidatorManagerInsufficientSignature(uint length);
    error errValidatorManagerInvalidSignatures(uint v, uint r, uint s);

    event ValidatorUpdated(address indexed validator, bool indexed status);
    event ThresholdChanged(uint8 threshold);

    uint8 private _threshold;
    EnumerableSet.AddressSet private _validators;

    uint[47] private __gap;

    modifier onlyValidator() {
        require(isValidator(_msgSender()), errValidatorManagerNotValidator(_msgSender()));
        _;
    }

    function __Validator_init(uint8 threshold_) internal onlyInitializing {
        __EIP712_init("Validator", "1.0.0");
        _threshold = threshold_;
    }

    function isValidator(address validator) public view returns (bool) {
        return _validators.contains(validator);
    }

    function threshold() external view returns (uint8) {
        return _threshold;
    }

    function allValidators() external view returns (address[] memory) {
        return _validators.values();
    }

    function validatorLength() external view returns (uint) {
        return _validators.length();
    }

    function validatorByIndex(uint index) external view returns (address) {
        return _validators.at(index);
    }

    // ----- Set Functions -----
    function changeThreshold(uint8 threshold_) external onlyOwner {
        _threshold = threshold_;
        emit ThresholdChanged(threshold_);
    }

    function setValidator(address validator) external {
        _updateValidator(validator, true);
    }

    function setValidators(address[] memory validators) public {
        for (uint i = 0; i < validators.length; ++i) {
            _updateValidator(validators[i], true);
        }
    }

    function removeValidator(address validator) external {
        _updateValidator(validator, false);
    }

    function removeValidators(address[] memory validators) public {
        for (uint i = 0; i < validators.length; ++i) {
            _updateValidator(validators[i], false);
        }
    }

    function resetValidators(address[] memory validators) external {
        removeValidators(_validators.values());
        setValidators(validators);
    }

    function _updateValidator(address account, bool set) internal onlyOwner {
        if (set) require(_validators.add(account), errValidatorManagerAlreadyExistValidator(account));
        else require(_validators.remove(account), errValidatorManagerNotExistValidator(account));
        emit ValidatorUpdated(account, set);
    }

    function _validateSignature(bytes32 h, uint8[] memory v, bytes32[] memory r, bytes32[] memory s) internal view {
        uint sigsLength = v.length;
        require(
            sigsLength == r.length && sigsLength == s.length,
            errValidatorManagerInvalidSignatures(sigsLength, r.length, s.length)
        );
        require(sigsLength >= _threshold, errValidatorManagerInsufficientSignature(sigsLength));

        uint valid = 0;
        address[] memory _signed = new address[](sigsLength);
        for (uint i = 0; i < sigsLength; ++i) {
            address validator = _hashTypedDataV4(h).recover(v[i], r[i], s[i]);

            // If there is even one invalid signature, false
            require(isValidator(validator), errValidatorManagerNotValidator(validator));

            bool dup = false;
            for (uint j = 0; j < _signed.length; j++) {
                if (validator == _signed[j]) {
                    dup = true;
                    break;
                }
            }

            if (!dup) {
                unchecked {
                    ++valid;
                }
            }
        }

        // Check one more time for duplicate verification
        require(valid >= _threshold, errValidatorManagerInsufficientSignature(valid));
    }
}
