// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IPriceFeed} from "./PriceFeed.sol";
import {PriceFeedLib} from "./lib/PriceFeedLib.sol";

import {IBridgeFeeStation} from "./interface/IBridgeFeeStation.sol";

contract BridgeFeeStation is Ownable, IBridgeFeeStation {
    using EnumerableSet for EnumerableSet.AddressSet;
    using PriceFeedLib for IPriceFeed;
    using Math for uint;

    error BridgeFeeStationChainAleadyExist(uint chainID);
    error BridgeFeeStationCanNotZeroAddress(string name);
    error BridgeFeeStationCanNotZeroValue(string name);
    error BridgeFeeStationInvalidLength();

    event BridgeFeeStationFinalizeBridgeGasSet(uint finalizeBridgeGas);
    event BridgeFeeStationGasPriceUpdated(uint remoteChainID, uint gasPrice);
    event BridgeFeeStationExchangeFeeUpdated(uint exFee);
    event BridgeFeeStationPriceFeedUpdated(IPriceFeed priceFeed);

    uint private constant DENOMINATOR = 1000;

    IPriceFeed private _priceFeed;
    uint private _exFeeRate;
    uint private _finalizeBridgeGas;

    mapping(IERC20 => TokenFee) private _tokensFee;
    mapping(uint => GasInfo) private _gasInfo;

    constructor(uint exFeeRate, uint finalizeBridgeGas) Ownable(_msgSender()) {
        require(finalizeBridgeGas != 0, BridgeFeeStationCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        _exFeeRate = exFeeRate;
    }

    function denominator() external pure returns (uint) {
        return DENOMINATOR;
    }

    function getGasInfo(uint remoteChainID) external view returns (GasInfo memory) {
        return _gasInfo[remoteChainID];
    }

    // priceFeed 가 없거나, token 시세가 존재하지 않으면 gas 수수료 0
    function getTokenFee(uint remoteChainID, IERC20 token)
        public
        view
        returns (uint minimumAmount, uint gasFee, uint exFeeRate)
    {
        (gasFee,) = _estimateGasToToken(remoteChainID, token);

        TokenFee memory tokenFee = _tokensFee[token];
        (minimumAmount, exFeeRate) =
            tokenFee.token == token ? (tokenFee.minimumAmount, tokenFee.exFeeRate) : (0, _exFeeRate);
        if (exFeeRate == type(uint).max) exFeeRate = 0;
    }

    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee)
    {
        uint exFeeRate;
        (minimumAmount, gasFee, exFeeRate) = getTokenFee(remoteChainID, token);

        exFee = value.mulDiv(exFeeRate, DENOMINATOR);
    }

    function _estimateGasToToken(uint remoteChainID, IERC20 toToken)
        private
        view
        returns (uint gasFee, uint updatedAt)
    {
        GasInfo memory gasInfo = _gasInfo[remoteChainID];
        if (address(_priceFeed) == address(0) || gasInfo.chainID != remoteChainID) return (0, 0);

        bool ok;
        (ok, gasFee, updatedAt) = _priceFeed.calculateAmountB(
            address(gasInfo.gasToken), address(toToken), gasInfo.gasPrice * _finalizeBridgeGas
        );
        if (!ok) return (0, 0);
    }

    function setGasInfo(uint remoteChainID, IERC20 gasToken, uint gasPrice) external onlyOwner {
        require(_gasInfo[remoteChainID].chainID == 0, BridgeFeeStationChainAleadyExist(remoteChainID));
        require(address(gasToken) != address(0), BridgeFeeStationCanNotZeroValue("gasToken"));
        _gasInfo[remoteChainID] = GasInfo({chainID: remoteChainID, gasToken: gasToken, gasPrice: gasPrice});
    }

    function setTokenFee(IERC20 token, uint minimumAmount, uint exFeeRate) public onlyOwner {
        require(address(token) != address(0), BridgeFeeStationCanNotZeroAddress("token"));
        _tokensFee[token] = TokenFee({token: token, minimumAmount: minimumAmount, exFeeRate: exFeeRate});
    }

    function setTokenFeeBatch(IERC20[] memory tokenList, uint[] memory minimumAmountList, uint[] memory exFeeRateList)
        external
        onlyOwner
    {
        require(
            tokenList.length == minimumAmountList.length && tokenList.length == exFeeRateList.length,
            BridgeFeeStationInvalidLength()
        );
        for (uint i = 0; i < tokenList.length; ++i) {
            setTokenFee(tokenList[i], minimumAmountList[i], exFeeRateList[i]);
        }
    }

    function removeTokenFee(IERC20 token) external onlyOwner {
        require(address(token) != address(0), BridgeFeeStationCanNotZeroAddress("token"));
        if (_tokensFee[token].token == token) delete(_tokensFee[token]);
    }

    function updateGasPrice(uint remoteChainID, uint gasPrice) external onlyOwner {
        // require(gasPrice != 0, BridgeFeeStationCanNotZeroValue("gasPrice"));
        _gasInfo[remoteChainID].gasPrice = gasPrice;
        emit BridgeFeeStationGasPriceUpdated(remoteChainID, gasPrice);
    }

    function setFinalizeBridgeGas(uint finalizeBridgeGas) external onlyOwner {
        require(finalizeBridgeGas != 0, BridgeFeeStationCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        emit BridgeFeeStationFinalizeBridgeGasSet(_finalizeBridgeGas);
    }

    function setExFeeRate(uint exFeeRate) external onlyOwner {
        _exFeeRate = exFeeRate;
        emit BridgeFeeStationExchangeFeeUpdated(_exFeeRate);
    }

    function setPriceFeed(IPriceFeed priceFeed) external onlyOwner {
        require(address(priceFeed) != address(0), BridgeFeeStationCanNotZeroAddress("priceFeed")); // allow zero address
        _priceFeed = priceFeed;
        emit BridgeFeeStationPriceFeedUpdated(_priceFeed);
    }

    function removePriceFeed() external onlyOwner {
        _priceFeed = IPriceFeed(address(0));
        emit BridgeFeeStationPriceFeedUpdated(_priceFeed);
    }
}
