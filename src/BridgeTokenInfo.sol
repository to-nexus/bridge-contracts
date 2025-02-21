// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IPriceFeed, PriceFeedLib} from "./PriceFeedCross.sol";
import {TokenStorage} from "./abstract/TokenStorage.sol";
import {IBridgeTokenInfo} from "./interface/IBridgeTokenInfo.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {IERC20, IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract BridgeTokenInfo is Ownable, IBridgeTokenInfo {
    using EnumerableSet for EnumerableSet.AddressSet;
    using PriceFeedLib for IPriceFeed;
    using Math for uint;

    error BridgeTokenInfoNotFoundETHPrice(address eth);
    error BridgeTokenInfoCanNotZeroAddress(string name);
    error BridgeTokenInfoCanNotZeroValue(string name);

    event BridgeTokenInfoGasPriceUpdated(uint gasPrice);
    event BridgeTokenInfoExchangeFeeUpdated(uint exFee);
    event BridgeTokenInfoPriceFeedUpdated(IPriceFeed priceFeed);

    struct Token {
        IERC20 token;
        uint minimumAmount;
        uint exchangeFee;
    }

    IPriceFeed public priceFeed;
    address public wcoin; // crosschain wrapped eth
    uint private constant DENOMINATOR = 1000;
    uint private _exFee = 10;
    uint private _gasPrice = 5 gwei;
    uint private _finalizeGas = 100000;

    // mapping(IERC20 => uint) public exchangeFee;
    mapping(IERC20 => Token) public _tokenInfo;

    constructor(address _wcoin) Ownable(_msgSender()) {
        require(_wcoin != address(0), BridgeTokenInfoCanNotZeroAddress("_wcoin"));
        wcoin = _wcoin;
    }

    function denominator() external pure returns (uint) {
        return DENOMINATOR;
    }

    function getTokenInfo(IERC20 token) public view returns (uint minimum, uint gas, uint ex, bool isValid) {
        Token memory info = _tokenInfo[token];
        (gas, isValid) = _convertGasToToken(token);
        (minimum, ex) = info.token == token ? (info.minimumAmount, info.exchangeFee) : (0, _exFee);
    }

    function calculate(IERC20 token, uint value)
        external
        view
        returns (uint minimum, uint gas, uint ex, bool isValid)
    {
        (minimum, gas, ex, isValid) = getTokenInfo(token);
        ex = value.mulDiv(ex, DENOMINATOR);
        isValid = isValid && value >= minimum;
    }

    function calculateMax(IERC20 token, uint totalValue)
        external
        view
        returns (uint value, uint gas, uint ex, bool isValid)
    {
        uint minimum;
        (minimum, gas, ex, isValid) = getTokenInfo(token);
        if (!isValid || totalValue <= gas) return (0, 0, 0, false);
        uint v = totalValue - gas;
        value = v.mulDiv(DENOMINATOR, (DENOMINATOR + ex));
        ex = value.mulDiv(ex, DENOMINATOR);
    }

    function _convertGasToToken(IERC20 token) private view returns (uint gas, bool isValid) {
        if (address(priceFeed) != address(0)) {
            (gas, isValid) = priceFeed.convertAtoB(wcoin, address(token), _gasPrice * _finalizeGas);
        } else {
            isValid = true;
        }
    }

    function updateGasPrice(uint gasPrice) external onlyOwner {
        require(gasPrice != 0, BridgeTokenInfoCanNotZeroValue("gasPrice"));
        _gasPrice = gasPrice;
        emit BridgeTokenInfoGasPriceUpdated(_gasPrice);
    }

    function updateFinalizeGas(uint finalizeGas) external onlyOwner {
        require(finalizeGas != 0, BridgeTokenInfoCanNotZeroValue("finalizeGas"));
        _finalizeGas = finalizeGas;
        emit BridgeTokenInfoGasPriceUpdated(_finalizeGas);
    }

    function changeEx(uint ex) external onlyOwner {
        _exFee = ex;
        emit BridgeTokenInfoExchangeFeeUpdated(ex);
    }

    function addTokenInfo(IERC20 token, uint minimum, uint ex) external onlyOwner {
        require(address(token) != address(0), BridgeTokenInfoCanNotZeroAddress("token"));
        _tokenInfo[token] = Token(token, minimum, ex);
    }

    function removeTokenInfo(IERC20 token) external onlyOwner {
        require(address(token) != address(0), BridgeTokenInfoCanNotZeroAddress("token"));
        if (_tokenInfo[token].token == token) delete(_tokenInfo[token]);
    }

    function setPriceFeed(IPriceFeed _priceFeed) external onlyOwner {
        require(address(_priceFeed) != address(0), BridgeTokenInfoCanNotZeroAddress("_priceFeed")); // allow zero address
        (bool isValid) = _priceFeed.contains(wcoin);
        require(isValid, BridgeTokenInfoNotFoundETHPrice(wcoin));
        priceFeed = _priceFeed;
        emit BridgeTokenInfoPriceFeedUpdated(priceFeed);
    }

    function removePriceFeed() external onlyOwner {
        priceFeed = IPriceFeed(address(0));
        emit BridgeTokenInfoPriceFeedUpdated(priceFeed);
    }
}
