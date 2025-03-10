// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

interface IRoleManager {
    function hasRole(bytes32 role, address account) external view returns (bool);
    function getRoleMembers(bytes32 role) external view returns (address[] memory);
    function setRole(bytes32 role, address[] memory accounts, bool set) external;
    function resetRole(bytes32 role, address[] memory newAccounts) external;
}
