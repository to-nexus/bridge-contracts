// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
// import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {CrossMintableERC20} from "../CrossMintableERC20.sol";
import {ITokenStorage} from "../interface/ITokenStorage.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

abstract contract TokenStorage is ITokenStorage {
    using EnumerableSet for EnumerableSet.AddressSet;

    error TokenStorageNotExistToken(address token);
    error TokenStorageAleadyExistToken(address token);
    error TokenStorageCanNotZeroAddress(string name);

    event TokenAdded(address indexed token);
    event TokenRemoved(address indexed token);

    EnumerableSet.AddressSet private _tokens;

    uint[49] private __gap;

    modifier checkZeroAddress(address token) {
        require(token != address(0), TokenStorageCanNotZeroAddress("token"));
        _;
    }

    function contains(address token) public view returns (bool) {
        return _tokens.contains(address(token));
    }

    function tokensLength() public view returns (uint) {
        return _tokens.length();
    }

    function tokenByIndex(uint i) public view returns (address) {
        return _tokens.at(i);
    }

    function allTokens() public view returns (address[] memory tokens) {
        uint length = _tokens.length();
        tokens = new address[](length);
        return _tokens.values();
    }

    function _addToken(address token) internal checkZeroAddress(token) {
        require(_tokens.add(token), TokenStorageAleadyExistToken(token));
        emit TokenAdded(token);
    }

    function _removeToken(address token) internal checkZeroAddress(token) {
        require(_tokens.remove(token), TokenStorageNotExistToken(token));
        emit TokenRemoved(token);
    }
}
