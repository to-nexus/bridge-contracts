// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface IValidatorManager {
    function isValidator(address validator) external view returns (bool);
    function threshold() external view returns (uint8);
}
