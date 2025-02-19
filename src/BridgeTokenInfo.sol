// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {TokenStorage} from "./abstract/TokenStorage.sol";
import {IBridgeTokenInfo} from "./interface/IBridgeTokenInfo.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract BridgeTokenInfo is Ownable, TokenStorage, IBridgeTokenInfo {
    using EnumerableSet for EnumerableSet.AddressSet;

    error BridgeTokenInfoInvalidToken(address token);

    event TokenInfoUpdated(address indexed token, uint minimumValue, uint gasFee, uint serviceFee);
    event TokenInfoRemoved(address indexed token);

    mapping(address => TokenInfo) private _tokenInfo;
    uint private _denominator = 1000;

    constructor() Ownable(_msgSender()) {}

    // address(1) == native coin
    function addTokenInfo(address token, uint minimumValue, uint gasFee, uint serviceFee) public {
        _addToken(token);
        _tokenInfo[token] = TokenInfo(token, minimumValue, gasFee, serviceFee);
        emit TokenInfoUpdated(token, minimumValue, gasFee, serviceFee);
    }

    function addTokenInfoMany(TokenInfo[] memory tokenInfoList) external {
        for (uint i = 0; i < tokenInfoList.length; i++) {
            addTokenInfo(
                tokenInfoList[i].token,
                tokenInfoList[i].minimumValue,
                tokenInfoList[i].gasFee,
                tokenInfoList[i].serviceFee
            );
        }
    }

    function updateTokenInfo(address token, uint minimum, uint gasFee, uint serviceFee) public {
        require(contains(token), BridgeTokenInfoInvalidToken(token));
        _tokenInfo[token].minimumValue = minimum;
        _tokenInfo[token].gasFee = gasFee;
        _tokenInfo[token].serviceFee = serviceFee;
        emit TokenInfoUpdated(token, minimum, gasFee, serviceFee);
    }

    function updateTokenInfoMany(TokenInfo[] memory tokenInfoList) external {
        for (uint i = 0; i < tokenInfoList.length; i++) {
            updateTokenInfo(
                tokenInfoList[i].token,
                tokenInfoList[i].minimumValue,
                tokenInfoList[i].gasFee,
                tokenInfoList[i].serviceFee
            );
        }
    }

    function removeTokenInfo(address token) public {
        _removeToken(token);
        delete(_tokenInfo[token]);
        emit TokenInfoRemoved(token);
    }

    function removeTokenInfoMany(address[] memory token) external {
        for (uint i = 0; i < token.length; i++) {
            removeTokenInfo(token[i]);
        }
    }

    function denominator() external view returns (uint) {
        return _denominator;
    }

    function calculate(address token, uint value) external view returns (uint minimum, uint gas, uint service) {
        TokenInfo memory info = getTokenInfo(token);
        return (info.minimumValue, info.gasFee, value * info.serviceFee / _denominator);
    }

    function getTokenInfo(address token) public view returns (TokenInfo memory) {
        if (contains(token)) return _tokenInfo[token];
        return TokenInfo(token, 0, 0, 0);
    }

    function getTokenInfoList(address[] memory tokens) public view returns (TokenInfo[] memory) {
        uint _length = tokensLength();
        TokenInfo[] memory info = new TokenInfo[](_length);

        for (uint i = 0; i < _length; i++) {
            info[i] = getTokenInfo(tokens[i]);
        }

        return info;
    }

    function tokenInfoByIndex(uint index) external view returns (TokenInfo memory) {
        return _tokenInfo[tokenByIndex(index)];
    }

    function allTokenInfo() external view returns (TokenInfo[] memory) {
        uint _length = tokensLength();
        TokenInfo[] memory info = new TokenInfo[](_length);

        for (uint i = 0; i < _length; i++) {
            info[i] = _tokenInfo[tokenByIndex(i)];
        }

        return info;
    }
}
