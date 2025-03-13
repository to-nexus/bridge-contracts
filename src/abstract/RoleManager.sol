// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract RoleManager is AccessControlUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error RoleManagerAlreadyHasRole(address account, bytes32 role);
    error RoleManagerDoesNotHaveRole(address account, bytes32 role);

    bytes32 internal constant ADMIN_ROLE = ("ADMIN");
    bytes32 internal constant OPERATOR_ROLE = ("OPERATOR");
    bytes32 internal constant VALIDATOR_ROLE = ("VALIDATOR");
    bytes32 internal constant UPDATOR_ROLE = ("UPDATOR");
    bytes32 internal constant BRIDGE_ROLE = ("BRIDGE");

    mapping(bytes32 role => EnumerableSet.AddressSet) private _roles;

    function __RoleManager_init(address owner) internal onlyInitializing {
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, owner);
    }

    function getRoleMembers(bytes32 role) external view returns (address[] memory) {
        return _roles[role].values();
    }

    function grantRoleBatch(bytes32 role, address[] memory accounts) external onlyRole(getRoleAdmin(role)) {
        for (uint i = 0; i < accounts.length; i++) {
            _grantRole(role, accounts[i]);
        }
    }

    function revokeRoleBatch(bytes32 role, address[] memory accounts) external onlyRole(getRoleAdmin(role)) {
        for (uint i = 0; i < accounts.length; i++) {
            _revokeRole(role, accounts[i]);
        }
    }

    function _grantRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._grantRole(role, account);
        if (ok) require(_roles[role].add(account), RoleManagerAlreadyHasRole(account, role));
    }

    function _revokeRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._revokeRole(role, account);
        if (ok) require(_roles[role].remove(account), RoleManagerDoesNotHaveRole(account, role));
    }
}
