pragma solidity 0.8.28;

import {IValidatorManager} from "../interface/IValidatorManager.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract ValidatorManager is OwnableUpgradeable, EIP712Upgradeable, IValidatorManager {
    using ECDSA for bytes32;
    using EnumerableSet for EnumerableSet.AddressSet;

    error ValidatorManagerNotValidator(address account);
    error ValidatorManagerAlreadyExistValidator(address account);
    error ValidatorManagerNotExistValidator(address account);

    event ValidatorSet(address validators, bool status);
    event ThresholdChanged(uint8 threshold);

    EnumerableSet.AddressSet private _validators;
    uint8 private _threshold;

    uint[48] private __gap;

    modifier onlyValidator() {
        require(isValidator(_msgSender()), ValidatorManagerNotValidator(_msgSender()));
        _;
    }

    function __Validator_init(address initialOwner) internal onlyInitializing {
        __Ownable_init(initialOwner);
        __EIP712_init("Validator", "1.0.0");
        _threshold = 3;
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
        _setValidator(validator, true);
    }

    function setValidators(address[] memory validators) external {
        for (uint i = 0; i < validators.length; ++i) {
            _setValidator(validators[i], true);
        }
    }

    function removeValidator(address validator) external {
        _setValidator(validator, false);
    }

    function removeValidators(address[] memory validators) external {
        for (uint i = 0; i < validators.length; ++i) {
            _setValidator(validators[i], false);
        }
    }

    function _setValidator(address account, bool set) internal onlyOwner {
        if (set) require(_validators.add(account), ValidatorManagerAlreadyExistValidator(account));
        else require(_validators.remove(account), ValidatorManagerNotExistValidator(account));
        emit ValidatorSet(account, set);
    }

    function _validate(bytes32 h, bytes[] memory sigs) internal view returns (bool) {
        uint sigsLength = sigs.length;
        if (sigsLength < _threshold) return false;

        uint valid = 0;
        address[] memory _signed = new address[](sigsLength);
        for (uint i = 0; i < sigs.length; ++i) {
            bytes memory sig = sigs[i];
            if (sig.length < 65) return false;

            address validator = _hashTypedDataV4(h).recover(sig);

            if (!isValidator(validator)) return false; // 유효하지 않은 서명이 하나라도 있다면 false

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

        return valid >= _threshold;
    }
}
