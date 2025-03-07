pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IRoleManager} from "../interface/IRoleManager.sol";

/**
 * @title RoleManager
 * @notice Abstract contract for managing bridge roles and signature verification
 * @dev This contract provides the following key features:
 * - Adding/removing/managing roles
 * - Basic role management functionality
 * - Role-based access control for bridge operations using bytes32 role identifiers
 * - Enumerable role membership tracking
 */
abstract contract RoleManager is OwnableUpgradeable, IRoleManager {
    using EnumerableSet for EnumerableSet.AddressSet;

    error RoleCanNotZeroAddress();
    error RoleInvalid();
    error RoleNotAuthorized(bytes32 role, address account);
    error RoleAlreadyGranted(bytes32 role, address account);
    error RoleNotGranted(bytes32 role, address account);

    /**
     * @notice Emitted when a role is added or removed
     * @param account Address that received/lost the role
     * @param role Role identifier that was modified
     * @param status True if role was added, false if removed
     */
    event RoleUpdated(address indexed account, bytes32 indexed role, bool status);

    /// @dev Mapping from role identifier to set of addresses with that role
    mapping(bytes32 => EnumerableSet.AddressSet) private _roles;

    /// @dev Storage gap for future upgrades
    uint[49] private __gap;

    /**
     * @notice Modifier to restrict function calls to specific role
     * @dev Checks if the caller has the required role
     * @param role The role identifier required to call the function
     */
    modifier onlyRole(bytes32 role) {
        require(hasRole(role, _msgSender()), RoleNotAuthorized(role, _msgSender()));
        _;
    }

    /**
     * @notice Checks if an address has a specific role
     * @dev Uses EnumerableSet.contains to efficiently check membership
     * @param role Role identifier to check
     * @param account Address to check
     * @return Boolean indicating if the address has the role
     */
    function hasRole(bytes32 role, address account) public view returns (bool) {
        return _roles[role].contains(account);
    }

    /**
     * @notice Returns all addresses with a specific role
     * @dev Uses EnumerableSet.values to get all members
     * @param role Role identifier to query
     * @return Array of addresses with the role
     */
    function getRoleMembers(bytes32 role) external view returns (address[] memory) {
        return _roles[role].values();
    }

    /**
     * @notice Adds or removes multiple addresses to/from a role
     * @dev Iterates through accounts array and updates each one
     * @param role Role identifier to modify
     * @param accounts Array of addresses to update
     * @param set True to add, false to remove
     */
    function setRole(bytes32 role, address[] memory accounts, bool set) public {
        for (uint i = 0; i < accounts.length; ++i) {
            _updateRole(role, accounts[i], set);
        }
    }

    /**
     * @notice Resets all addresses for a role and sets new ones
     * @dev First removes all existing members, then adds new ones
     * @param role Role identifier to reset
     * @param newAccounts Array of addresses for the new set
     */
    function resetRole(bytes32 role, address[] memory newAccounts) external {
        setRole(role, _roles[role].values(), false);
        setRole(role, newAccounts, true);
    }

    /**
     * @notice Updates role status for an account
     * @dev Internal function to add or remove a role
     * - Validates account is not zero address
     * - Validates role is not zero bytes32
     * - Adds or removes role based on grant parameter
     * - Emits RoleUpdated event
     * @param account Address to update
     * @param role Role identifier to modify
     * @param grant True to add, false to remove
     */
    function _updateRole(bytes32 role, address account, bool grant) internal virtual onlyOwner {
        require(account != address(0), RoleCanNotZeroAddress());
        require(role != bytes32(0), RoleInvalid());

        if (grant) require(_roles[role].add(account), RoleAlreadyGranted(role, account));
        else require(_roles[role].remove(account), RoleNotGranted(role, account));

        emit RoleUpdated(account, role, grant);
    }
}
