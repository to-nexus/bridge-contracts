// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {Const} from "../lib/Const.sol";

abstract contract RoleManager is AccessControlUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error RoleManagerAlreadyHasRole(address account, bytes32 role);
    error RoleManagerDoesNotHaveRole(address account, bytes32 role);
    error RoleManagerMissmatchLength();

    /// @custom:storage-location erc7201:cross.bridge.RoleManager
    struct RoleManagerStorage {
        mapping(bytes32 role => EnumerableSet.AddressSet) _roles;
    }

    // keccak256(abi.encode(uint256(keccak256("erc7201:cross.bridge.RoleManager")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant RoleManagerStorageLocation =
        0x66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00;

    function _getRoleManagerStorage() private pure returns (RoleManagerStorage storage $) {
        assembly {
            $.slot := RoleManagerStorageLocation
        }
    }

    function __RoleManager_init(address owner) internal onlyInitializing {
        __RoleManager_init_unchained(owner);
    }

    function __RoleManager_init_unchained(address owner) internal onlyInitializing {
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, owner);
    }

    /**
     * @notice Returns all members of a specific role
     * @param role The role to query
     * @return members Array of addresses that have the specified role
     */
    function getRoleMembers(bytes32 role) external view returns (address[] memory) {
        RoleManagerStorage storage $ = _getRoleManagerStorage();
        return $._roles[role].values();
    }

    /**
     * @notice Grants multiple roles to a list of accounts
     * @dev Validates input array lengths match
     * - Grants each role to the corresponding account
     * @param roles Array of roles to grant
     * @param accounts Array of accounts to grant roles to
     */
    function grantRoleBatch(bytes32[] memory roles, address[] memory accounts) external {
        require(roles.length == accounts.length, RoleManagerMissmatchLength());
        for (uint i = 0; i < accounts.length; ++i) {
            grantRole(roles[i], accounts[i]);
        }
    }

    /**
     * @notice Revokes multiple roles from a list of accounts
     * @dev Validates input array lengths match
     * - Revokes each role from the corresponding account
     * @param roles Array of roles to revoke
     * @param accounts Array of accounts to revoke roles from
     */
    function revokeRoleBatch(bytes32[] memory roles, address[] memory accounts) external {
        require(roles.length == accounts.length, RoleManagerMissmatchLength());
        for (uint i = 0; i < accounts.length; ++i) {
            revokeRole(roles[i], accounts[i]);
        }
    }

    function _grantRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._grantRole(role, account);
        if (ok) {
            RoleManagerStorage storage $ = _getRoleManagerStorage();
            require($._roles[role].add(account), RoleManagerAlreadyHasRole(account, role));
        }
    }

    function _revokeRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._revokeRole(role, account);
        if (ok) {
            RoleManagerStorage storage $ = _getRoleManagerStorage();
            require($._roles[role].remove(account), RoleManagerDoesNotHaveRole(account, role));
        }
    }
}
