// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface IValidatorManager {
    function isValidator(address validator) external view returns (bool);
    function threshold() external view returns (uint8);
    function allValidators() external view returns (address[] memory);
    function validatorLength() external view returns (uint);
    function validatorByIndex(uint index) external view returns (address);
}
