// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {CrossMintableERC20} from "./CrossMintableERC20.sol";
import {TokenStorage} from "./abstract/TokenStorage.sol";
import {IBridgeFeeManager} from "./interface/IBridgeFeeManager.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract BridgeFeeManager is Ownable, TokenStorage, IBridgeFeeManager {
    using EnumerableSet for EnumerableSet.AddressSet;

    error BridgeFeeManagerInvalidToken(address token);

    event FeeInfoAdded(address indexed token, uint gasFee, uint serviceFee);
    event FeeInfoUpdated(address indexed token, uint gasFee, uint serviceFee);
    event FeeInfoRemoved(address indexed token, uint gasFee, uint serviceFee);

    mapping(address => FeeInfo) private _feeInfo;
    uint private _denominator = 1000;

    constructor() Ownable(_msgSender()) {}

    // address(1) == native coin
    function addFeeInfo(address token, uint gasFee, uint serviceFee) public {
        _addToken(token);
        _feeInfo[token] = FeeInfo(token, gasFee, serviceFee);
        emit FeeInfoAdded(token, gasFee, serviceFee);
    }

    function addFeeInfoMany(FeeInfo[] memory feeInfoList) external {
        for (uint i = 0; i < feeInfoList.length; i++) {
            addFeeInfo(feeInfoList[i].token, feeInfoList[i].gasFee, feeInfoList[i].serviceFee);
        }
    }

    function updateFeeInfo(address token, uint gasFee, uint serviceFee) public {
        require(contains(token), BridgeFeeManagerInvalidToken(token));
        _feeInfo[token].gasFee = gasFee;
        _feeInfo[token].serviceFee = serviceFee;
        emit FeeInfoUpdated(token, gasFee, serviceFee);
    }

    function updateFeeInfoMany(FeeInfo[] memory feeInfoList) external {
        for (uint i = 0; i < feeInfoList.length; i++) {
            updateFeeInfo(feeInfoList[i].token, feeInfoList[i].gasFee, feeInfoList[i].serviceFee);
        }
    }

    function removeFeeInfo(address token) public {
        _removeToken(token);
        delete(_feeInfo[token]);
        emit FeeInfoRemoved(token, _feeInfo[token].gasFee, _feeInfo[token].serviceFee);
    }

    function removeFeeInfoMany(address[] memory token) external {
        for (uint i = 0; i < token.length; i++) {
            removeFeeInfo(token[i]);
        }
    }

    function denominator() external view returns (uint) {
        return _denominator;
    }

    function calculateFee(address token, uint value) external view returns (uint gas, uint service) {
        FeeInfo memory info = getTokenFee(token);
        return (info.gasFee, value * info.serviceFee / _denominator);
    }

    function getTokenFee(address token) public view returns (FeeInfo memory) {
        if (contains(token)) return _feeInfo[token];
        return FeeInfo(token, 0, 0);
    }

    function getTokenFeeList(address[] memory tokens) public view returns (FeeInfo[] memory) {
        uint _length = tokensLength();
        FeeInfo[] memory info = new FeeInfo[](_length);

        for (uint i = 0; i < _length; i++) {
            info[i] = getTokenFee(tokens[i]);
        }

        return info;
    }

    function feeInfoByIndex(uint index) external view returns (FeeInfo memory) {
        return _feeInfo[tokenByIndex(index)];
    }

    function allFeeInfo() external view returns (FeeInfo[] memory) {
        uint _length = tokensLength();
        FeeInfo[] memory info = new FeeInfo[](_length);

        for (uint i = 0; i < _length; i++) {
            info[i] = _feeInfo[tokenByIndex(i)];
        }

        return info;
    }
}
