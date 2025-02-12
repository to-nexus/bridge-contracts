// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossMintableERC20} from "../CrossMintableERC20.sol";
import {ITokenManager} from "../interface/ITokenManager.sol";
import {TokenStorage} from "./TokenStorage.sol";

abstract contract TokenManager is TokenStorage, ITokenManager {
    error TokenManagerInvalidToken(address token);

    event PairUpdated(address indexed token, address indexed pair, bool registered);
    event TokenPaused(address indexed token);
    event TokenUnpaused(address indexed token);

    mapping(address => address) private _pair;
    mapping(address => bool) private _paused;

    uint[48] private __gap;

    modifier checkValidToken(address token) {
        require(isValidToken(token), TokenManagerInvalidToken(token));
        _;
    }

    function isValidToken(address token) public view override returns (bool) {
        return (contains(token) && !_paused[token]);
    }

    function getPairToken(address token) public view returns (address) {
        return _pair[token];
    }

    function allPairs() public view returns (address[] memory) {
        address[] memory tokens = allTokens();
        address[] memory pairs = new address[](tokens.length);
        for (uint i = 0; i < tokens.length; i++) {
            pairs[i] = _pair[tokens[i]];
        }
        return pairs;
    }

    function _addTokenPair(address token, address pair) internal {
        _addToken(token);
        _pair[token] = pair;
        emit PairUpdated(token, pair, true);
    }

    function _removeTokenPair(address token) internal {
        _removeToken(token);
        address pair = _pair[token];
        delete (_pair[token]);
        emit PairUpdated(token, pair, false);
    }

    function _pauseToken(address token) internal {
        require(isValidToken(token), TokenStorageNotExistToken(token));
        _paused[token] = true;
        emit TokenPaused(token);
    }

    function _unpauseToken(address token) internal {
        require(contains(token) && _paused[token], TokenStorageNotExistToken(token));
        _paused[token] = false;
        emit TokenUnpaused(token);
    }
}
