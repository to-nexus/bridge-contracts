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
    error BridgeFeeStationCanNotZeroValue(string name);
    error BridgeFeeStationInvalidLength();

    event BridgeFeeStationFinalizeBridgeGasSet(uint finalizeBridgeGas);
    event BridgeFeeStationGasPriceUpdated(uint indexed remoteChainID, uint gasPrice);
    event BridgeFeeStationExchangeFeeUpdated(uint exFeeRate);
    event BridgeFeeStationPriceFeedUpdated(IPriceFeed indexed priceFeed);

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

    /**
     * @notice Returns the fee denominator
     * @dev Used for fee calculations (1000 = 100%)
     * @return Denominator constant value
     */
    function denominator() external pure returns (uint) {
        return DENOMINATOR;
    }

    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumAmount, uint gasFee, uint exFee)
    {
        uint exFeeRate;
        (minimumAmount, gasFee, exFeeRate) = getTokenFee(remoteChainID, token);

        exFee = value * exFeeRate / DENOMINATOR;
    }

    /**
     * @notice Gets gas information for a chain
     * @dev Returns stored gas configuration
     * @param remoteChainID Chain ID to query
     * @return GasInfo struct containing gas token and price
     */
    function getGasInfo(uint remoteChainID) external view returns (GasInfo memory) {
        return _gasInfo[remoteChainID];
    }

    /**
     * @notice Calculates fees for a token
     * @dev Determines minimum amount and fees
     * - Gets token-specific or default fee rates
     * - Calculates gas fee using price feed
     * - Handles fee exemptions
     * @param remoteChainID Chain ID for fee calculation
     * @param token Token to calculate fees for
     * @return minimumAmount Minimum required amount
     * @return gasFee Gas fee in token amount
     * @return exFeeRate Exchange fee rate
     */
    function getTokenFee(uint remoteChainID, IERC20 token)
        public
        view
        returns (uint minimumAmount, uint gasFee, uint exFeeRate)
    {
        (minimumAmount, exFeeRate) = _getTokenFee(token);

        (gasFee,) = _calculateGasFee(remoteChainID, token, _finalizeBridgeGas);

        // exFeeRate가 max값이면 0으로 설정 (수수료 면제)
        if (exFeeRate == type(uint).max) exFeeRate = 0;
    }

    /**
     * @notice Gets token-specific fee configuration
     * @dev Internal helper for fee calculation
     * - Checks for token-specific settings
     * - Falls back to default rate
     * @param token Token to get fees for
     * @return minimumAmount Minimum required amount
     * @return exFeeRate Exchange fee rate
     */
    function _getTokenFee(IERC20 token) private view returns (uint minimumAmount, uint exFeeRate) {
        TokenFee memory tokenFee = _tokensFee[token];
        if (tokenFee.token == token) return (tokenFee.minimumAmount, tokenFee.exFeeRate);
        return (0, _exFeeRate);
    }

    /**
     * @notice Calculates gas fee in token amount
     * @dev Uses price feed for conversion
     * - Validates price feed availability
     * - Converts gas cost to token amount
     * @param remoteChainID Chain ID for gas price
     * @param token Token to calculate fee in
     * @param gasAmount Gas amount to calculate for
     * @return gasFee Calculated gas fee
     * @return updatedAt Timestamp of price data
     */
    function _calculateGasFee(uint remoteChainID, IERC20 token, uint gasAmount)
        private
        view
        returns (uint gasFee, uint updatedAt)
    {
        GasInfo memory gasInfo = _gasInfo[remoteChainID];
        if (address(_priceFeed) == address(0) || gasInfo.chainID != remoteChainID) return (0, 0);

        (, gasFee, updatedAt) = _priceFeed.calculateAmountB(address(gasInfo.gasToken), address(token), gasAmount);
    }

    // set functions
    function setGasInfo(uint remoteChainID, IERC20 gasToken, uint gasPrice) external onlyOwner {
        require(_gasInfo[remoteChainID].chainID == 0, BridgeFeeStationChainAleadyExist(remoteChainID));
        require(address(gasToken) != address(0), BridgeFeeStationCanNotZeroValue("gasToken"));
        _gasInfo[remoteChainID] = GasInfo({chainID: remoteChainID, gasToken: gasToken, gasPrice: gasPrice});
    }

    /**
     * @notice Sets token fee parameters
     * @dev Configures token-specific fees
     * - Validates token address
     * - Updates minimum amount and fee rate
     * @param token Token to configure
     * @param minimumAmount Minimum bridgeable amount
     * @param exFeeRate Exchange fee rate
     */
    function setTokenFee(IERC20 token, uint minimumAmount, uint exFeeRate) public onlyOwner {
        require(address(token) != address(0), BridgeFeeStationCanNotZeroValue("token"));
        _tokensFee[token] = TokenFee({token: token, minimumAmount: minimumAmount, exFeeRate: exFeeRate});
    }

    /**
     * @notice Sets fee parameters for multiple tokens
     * @dev Batch configuration of token fees
     * - Validates input array lengths match
     * - Updates each token's parameters
     * @param tokenList Array of token addresses
     * @param minimumAmountList Array of minimum amounts
     * @param exFeeRateList Array of exchange fee rates
     */
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

    /**
     * @notice Removes fee configuration for a token
     * @dev Deletes token-specific fee settings
     * - Validates token address
     * - Removes custom fee configuration
     * @param token Token to remove configuration for
     */
    function removeTokenFee(IERC20 token) external onlyOwner {
        require(address(token) != address(0), BridgeFeeStationCanNotZeroValue("token"));
        if (_tokensFee[token].token == token) delete(_tokensFee[token]);
    }

    function updateGasPrice(uint remoteChainID, uint gasPrice) external onlyOwner {
        // require(gasPrice != 0, BridgeFeeStationCanNotZeroValue("gasPrice"));
        _gasInfo[remoteChainID].gasPrice = gasPrice;
        emit BridgeFeeStationGasPriceUpdated(remoteChainID, gasPrice);
    }

    /**
     * @notice Sets gas amount for bridge finalization
     * @dev Updates gas cost estimation
     * - Validates non-zero value
     * - Updates stored gas amount
     * - Emits update event
     * @param finalizeBridgeGas New gas amount
     */
    function setFinalizeBridgeGas(uint finalizeBridgeGas) external onlyOwner {
        require(finalizeBridgeGas != 0, BridgeFeeStationCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        emit BridgeFeeStationFinalizeBridgeGasSet(_finalizeBridgeGas);
    }

    /**
     * @notice Sets global exchange fee rate
     * @dev Updates default fee rate
     * - Updates stored rate
     * - Emits update event
     * @param exFeeRate New exchange fee rate
     */
    function setExFeeRate(uint exFeeRate) external onlyOwner {
        _exFeeRate = exFeeRate;
        emit BridgeFeeStationExchangeFeeUpdated(_exFeeRate);
    }

    /**
     * @notice Sets price feed contract
     * @dev Updates price oracle reference
     * - Validates non-zero address
     * - Updates stored reference
     * - Emits update event
     * @param priceFeed New price feed contract
     */
    function setPriceFeed(IPriceFeed priceFeed) external onlyOwner {
        require(address(priceFeed) != address(0), BridgeFeeStationCanNotZeroValue("priceFeed")); // allow zero address
        _priceFeed = priceFeed;
        emit BridgeFeeStationPriceFeedUpdated(_priceFeed);
    }

    /**
     * @notice Removes price feed integration
     * @dev Disables price-based fee calculation
     * - Sets price feed to zero address
     * - Emits update event
     */
    function removePriceFeed() external onlyOwner {
        _priceFeed = IPriceFeed(address(0));
        emit BridgeFeeStationPriceFeedUpdated(_priceFeed);
    }
}
