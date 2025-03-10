// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IPriceFeed} from "./PriceFeed.sol";
import {CalcGasFeeLib} from "./lib/CalcGasFeeLib.sol";

import {IBridgeFeeManager} from "./interface/IBridgeFeeManager.sol";

/**
 * @title BridgeFeeManager
 * @notice Contract for calculating and managing bridge operation fees
 * @dev Provides functionality to:
 * - Calculate gas fees for bridge transactions based on destination chain
 * - Determine exchange fees based on token and amount
 * - Set and update fee parameters for different tokens and chains
 * - Integrate with PriceFeed for price-aware fee calculations
 */
contract BridgeFeeManager is Ownable, IBridgeFeeManager {
    using EnumerableSet for EnumerableSet.AddressSet;
    using CalcGasFeeLib for IPriceFeed;
    using Math for uint;

    error BridgeFeeManagerChainAleadyExist(uint chainID);
    error BridgeFeeManagerCanNotZeroValue(string name);
    error BridgeFeeManagerInvalidLength();

    event BridgeFeeManagerFinalizeBridgeGasSet(uint finalizeBridgeGas);
    event BridgeFeeManagerGasPriceUpdated(uint indexed remoteChainID, uint gasPrice);
    event BridgeFeeManagerExchangeFeeUpdated(uint exFeeRate);
    event BridgeFeeManagerPriceFeedUpdated(IPriceFeed indexed priceFeed);

    /// @dev Base denominator for fee calculations (1000 = 100%)
    uint private constant DENOMINATOR = 1000;

    /// @dev Price feed contract for token price information
    IPriceFeed private _priceFeed;
    /// @dev Global exchange fee rate applied when no token-specific rate is set
    uint private _exFeeRate;
    /// @dev Gas amount required for finalize bridge operations
    uint private _finalizeBridgeGas;

    /// @dev Mapping from token to its fee configuration
    mapping(IERC20 => TokenFee) private _tokensFee;
    /// @dev Mapping from chain ID to gas price on that chain
    mapping(uint => uint) private _gasPrice;

    /**
     * @notice Initializes the BridgeFeeManager contract
     * @dev Sets initial fee parameters
     * @param exFeeRate Global exchange fee rate (with DENOMINATOR as base)
     * @param finalizeBridgeGas Gas amount required for finalize bridge operations
     */
    constructor(uint exFeeRate, uint finalizeBridgeGas) Ownable(_msgSender()) {
        require(finalizeBridgeGas != 0, BridgeFeeManagerCanNotZeroValue("finalizeBridgeGas"));
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

    /**
     * @notice Calculates total fees for a bridge operation
     * @dev Determines minimum value, gas fee, and exchange fee
     * @param remoteChainID Target chain identifier
     * @param token Token being bridged
     * @param value Amount being bridged
     * @return minimumValue Minimum amount required for bridging
     * @return gasFee Gas fee for target chain transaction
     * @return exFee Exchange fee based on token amount
     */
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint gasFee, uint exFee)
    {
        uint exFeeRate;
        (minimumValue, gasFee, exFeeRate) = getTokenFee(remoteChainID, token);

        exFee = value * exFeeRate / DENOMINATOR;
    }

    /**
     * @notice Gets gas information for a chain
     * @dev Returns stored gas configuration
     * @param remoteChainID Chain ID to query
     * @return Gas price for the specified chain
     */
    function getGasPrice(uint remoteChainID) external view returns (uint) {
        return _gasPrice[remoteChainID];
    }

    /**
     * @notice Calculates fees for a token
     * @dev Determines minimum amount and fees
     * - Gets token-specific or default fee rates
     * - Calculates gas fee using price feed
     * - Handles fee exemptions
     * @param remoteChainID Chain ID for fee calculation
     * @param token Token to calculate fees for
     * @return minimumValue Minimum required amount
     * @return gasFee Gas fee in token amount
     * @return exFeeRate Exchange fee rate
     */
    function getTokenFee(uint remoteChainID, IERC20 token)
        public
        view
        returns (uint minimumValue, uint gasFee, uint exFeeRate)
    {
        (minimumValue, exFeeRate) = _getTokenFee(token);

        (gasFee,) = _calculateGasFee(remoteChainID, token);

        // If exFeeRate is set to max value, set it to 0 (fee exemption)
        if (exFeeRate == type(uint).max) exFeeRate = 0;
    }

    /**
     * @notice Gets token-specific fee configuration
     * @dev Internal helper for fee calculation
     * - Checks for token-specific settings
     * - Falls back to default rate
     * @param token Token to get fees for
     * @return minimumValue Minimum required amount
     * @return exFeeRate Exchange fee rate
     */
    function _getTokenFee(IERC20 token) private view returns (uint minimumValue, uint exFeeRate) {
        TokenFee memory tokenFee = _tokensFee[token];
        if (tokenFee.token == token) return (tokenFee.minimumValue, tokenFee.exFeeRate);
        return (0, _exFeeRate);
    }

    /**
     * @notice Calculates gas fee in token amount
     * @dev Uses price feed for conversion
     * - Validates price feed availability
     * - Converts gas cost to token amount
     * @param remoteChainID Chain ID for gas price
     * @param token Token to calculate fee in
     * @return gasFee Calculated gas fee
     * @return updatedAt Timestamp of price data
     */
    function _calculateGasFee(uint remoteChainID, IERC20 token) private view returns (uint gasFee, uint updatedAt) {
        uint gasPrice = _gasPrice[remoteChainID];
        if (address(_priceFeed) == address(0) || gasPrice == 0) return (0, 0);

        (, gasFee, updatedAt) =
            _priceFeed.calculateTokenAmountForGasFee(remoteChainID, address(token), _finalizeBridgeGas * gasPrice);
    }

    /**
     * @notice Sets gas price for a chain
     * @dev Initializes gas price for a new chain
     * - Prevents overwriting existing prices
     * - Validates non-zero price
     * @param remoteChainID Chain ID to set price for
     * @param gasPrice Gas price value for the chain
     */
    function setGasPrice(uint remoteChainID, uint gasPrice) external onlyOwner {
        require(_gasPrice[remoteChainID] == 0, BridgeFeeManagerChainAleadyExist(remoteChainID));
        require(gasPrice != 0, BridgeFeeManagerCanNotZeroValue("gasPrice"));
        _gasPrice[remoteChainID] = gasPrice;
    }

    /**
     * @notice Sets token fee parameters
     * @dev Configures token-specific fees
     * - Validates token address
     * - Updates minimum amount and fee rate
     * @param token Token to configure
     * @param minimumValue Minimum bridgeable amount
     * @param exFeeRate Exchange fee rate
     */
    function setTokenFee(IERC20 token, uint minimumValue, uint exFeeRate) public onlyOwner {
        require(address(token) != address(0), BridgeFeeManagerCanNotZeroValue("token"));
        _tokensFee[token] = TokenFee({token: token, minimumValue: minimumValue, exFeeRate: exFeeRate});
    }

    /**
     * @notice Sets fee parameters for multiple tokens
     * @dev Batch configuration of token fees
     * - Validates input array lengths match
     * - Updates each token's parameters
     * @param tokenList Array of token addresses
     * @param minimumValueList Array of minimum amounts
     * @param exFeeRateList Array of exchange fee rates
     */
    function setTokenFeeBatch(IERC20[] memory tokenList, uint[] memory minimumValueList, uint[] memory exFeeRateList)
        external
        onlyOwner
    {
        require(
            tokenList.length == minimumValueList.length && tokenList.length == exFeeRateList.length,
            BridgeFeeManagerInvalidLength()
        );
        for (uint i = 0; i < tokenList.length; ++i) {
            setTokenFee(tokenList[i], minimumValueList[i], exFeeRateList[i]);
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
        require(address(token) != address(0), BridgeFeeManagerCanNotZeroValue("token"));
        if (_tokensFee[token].token == token) delete(_tokensFee[token]);
    }

    /**
     * @notice Updates gas price for an existing chain
     * @dev Modifies gas price and emits event
     * @param remoteChainID Chain ID to update
     * @param gasPrice New gas price value
     */
    function updateGasPrice(uint remoteChainID, uint gasPrice) external onlyOwner {
        // require(gasPrice != 0, BridgeFeeManagerCanNotZeroValue("gasPrice"));
        _gasPrice[remoteChainID] = gasPrice;
        emit BridgeFeeManagerGasPriceUpdated(remoteChainID, gasPrice);
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
        require(finalizeBridgeGas != 0, BridgeFeeManagerCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        emit BridgeFeeManagerFinalizeBridgeGasSet(_finalizeBridgeGas);
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
        emit BridgeFeeManagerExchangeFeeUpdated(_exFeeRate);
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
        require(address(priceFeed) != address(0), BridgeFeeManagerCanNotZeroValue("priceFeed")); // allow zero address
        _priceFeed = priceFeed;
        emit BridgeFeeManagerPriceFeedUpdated(_priceFeed);
    }

    /**
     * @notice Removes price feed integration
     * @dev Disables price-based fee calculation
     * - Sets price feed to zero address
     * - Emits update event
     */
    function removePriceFeed() external onlyOwner {
        _priceFeed = IPriceFeed(address(0));
        emit BridgeFeeManagerPriceFeedUpdated(_priceFeed);
    }
}
