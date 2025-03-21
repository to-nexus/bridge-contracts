// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

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

    function getRoleMembers(bytes32 role) external view returns (address[] memory) {
        RoleManagerStorage storage $ = _getRoleManagerStorage();
        return $._roles[role].values();
    }

    function grantRoleBatch(bytes32[] memory roles, address[] memory accounts) external {
        require(roles.length == accounts.length, RoleManagerMissmatchLength());
        for (uint i = 0; i < accounts.length; ++i) {
            grantRole(roles[i], accounts[i]);
        }
    }

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
