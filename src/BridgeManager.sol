// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IPriceFeed} from "./PriceFeed.sol";
import {CalcGasFeeLib} from "./lib/CalcGasFeeLib.sol";

import {IBridgeManager} from "./interface/IBridgeManager.sol";

/**
 * @title BridgeManager
 * @notice Contract for calculating and managing bridge operation fees
 * @dev Provides fee calculations, token monitoring, and safety threshold management
 * Key features:
 * - Gas and exchange fee calculations based on token type and chain
 * - Token volume monitoring for security threshold enforcement
 * - Time-window based transfer volume restrictions
 */
contract BridgeManager is AccessControl, IBridgeManager {
    using Math for uint;
    using EnumerableSet for EnumerableSet.AddressSet;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using CalcGasFeeLib for IPriceFeed;

    error BridgeManagerChainAleadyExist(uint chainID);
    error BridgeManagerCanNotZeroValue(string name);
    error BridgeManagerInvalidLength();

    event BridgeManagerFinalizeBridgeGasSet(uint finalizeBridgeGas);
    event BridgeManagerGasPriceUpdated(uint indexed remoteChainID, uint gasPrice);
    event BridgeManagerExchangeFeeUpdated(address indexed token, uint exFeeRate);
    event BridgeManagerPriceFeedUpdated(IPriceFeed indexed priceFeed);
    event VerificationAmountThresholdSet(uint verificationAmountThreshold);
    event DefaultTokenPriceSet(uint defaultTokenPrice);
    event TimeWindowSet(uint timeWindow);
    event PeriodTotalValueThresholdSet(uint periodTotalValueThreshold);
    event MinimumTokenValueSet(uint minimumTokenValue);

    /// @dev Base denominator for fee calculations (1000 = 100%)
    uint private constant DENOMINATOR = 1000;
    bytes32 private constant ADMIN_ROLE = bytes32("ADMIN");
    bytes32 private constant UPDATOR_ROLE = bytes32("UPDATOR");
    bytes32 private constant BRIDGE_ROLE = bytes32("BRIDGE");
    uint private constant TOKEN_SCORE_MASK = type(uint).max >> 64;

    /// @dev Price feed contract for token price information
    IPriceFeed private _priceFeed;
    /// @dev Gas amount required for finalize bridge operations
    uint private _finalizeBridgeGas;
    /// @dev Global exchange fee rate (default if no token-specific rate)
    uint private _defaultExFeeRate;
    /// @dev Default token price for monitoring when no price feed is available
    uint private _defaultTokenPrice;
    /// @dev Minimum token value for transfer (in dollars)
    uint private _minimumTokenValue;
    /// @dev Maximum dollar value for a single transaction before verification
    uint private _verificationAmountThreshold;
    /// @dev Time window for monitoring token volume (in seconds)
    uint private _timeWindow;
    /// @dev Maximum dollar value of tokens that can be processed within the time window
    uint private _periodTotalValueThreshold;

    /// @dev Token-specific exchange fee rates
    mapping(IERC20 => uint) private _exFeeRate;
    /// @dev Chain-specific gas prices
    mapping(uint => uint) private _gasPrice;

    /// @dev Current token volume within the time window
    mapping(IERC20 => uint) private _tokenCurrentVolume;
    /// @dev History of token movements for volume tracking
    mapping(IERC20 => DoubleEndedQueue.Bytes32Deque) private _tokenMovementHistory;

    /**
     * @notice Initializes the BridgeManager contract
     * @param initialOwner Admin address
     * @param _bridge Bridge contract address
     * @param finalizeBridgeGas Gas amount for finalization
     * @param defaultTokenPrice Default price for tokens without price feed
     * @param defaultExFeeRate Default exchange fee rate
     * @param minimumTokenValue Minimum token value in USD
     * @param verificationAmountThreshold Threshold for single transaction verification
     * @param periodTotalValueThreshold Threshold for period volume
     * @param timeWindow Monitoring time window in seconds
     */
    constructor(
        address initialOwner,
        address _bridge,
        uint finalizeBridgeGas,
        uint defaultTokenPrice,
        uint defaultExFeeRate,
        uint minimumTokenValue,
        uint verificationAmountThreshold,
        uint periodTotalValueThreshold,
        uint timeWindow
    ) {
        require(finalizeBridgeGas != 0, BridgeManagerCanNotZeroValue("finalizeBridgeGas"));
        require(initialOwner != address(0), BridgeManagerCanNotZeroValue("initialOwner"));
        require(_bridge != address(0), BridgeManagerCanNotZeroValue("_bridge"));
        _grantRole(DEFAULT_ADMIN_ROLE, initialOwner);
        _grantRole(ADMIN_ROLE, initialOwner);
        _grantRole(BRIDGE_ROLE, _bridge);

        _finalizeBridgeGas = finalizeBridgeGas;
        _defaultTokenPrice = defaultTokenPrice;
        _defaultExFeeRate = defaultExFeeRate;
        _minimumTokenValue = minimumTokenValue;
        _verificationAmountThreshold = verificationAmountThreshold;
        _periodTotalValueThreshold = periodTotalValueThreshold;
        _timeWindow = timeWindow;
    }

    /**
     * @notice Verifies token transfer value against security thresholds
     * @dev Implements a time-window based monitoring system with dual thresholds:
     * 1. Single transfer threshold (_verificationAmountThreshold)
     * 2. Time window total volume threshold (_periodTotalValueThreshold)
     *
     * The system tracks token movements in a sliding time window, removing
     * expired records and adding new ones. Dollar value of tokens is calculated
     * using price feed or default price.
     *
     * @param token The token to verify
     * @param value The amount of tokens
     * @return ok Boolean indicating verification passed
     * @return reason Error message if verification failed
     */
    function validateBridgeTokenValue(IERC20 token, uint value)
        external
        onlyRole(BRIDGE_ROLE)
        returns (bool ok, bytes memory reason)
    {
        // Verify token price and threshold
        (, uint score) = getTokenPriceWithValue(token, value);
        if (_verificationAmountThreshold != 0 && _verificationAmountThreshold < score) {
            return (false, "verification amount threshold exceeded");
        }

        // Manage history and verify thresholds
        DoubleEndedQueue.Bytes32Deque storage deque = _tokenMovementHistory[token];
        uint currentTime = block.timestamp;
        uint windowStartTime = currentTime - _timeWindow;

        while (deque.length() > 0) {
            bytes32 frontRecord = deque.front();
            uint recordTime = uint(frontRecord >> 192);

            if (recordTime < windowStartTime) {
                uint oldScore = uint(frontRecord) & TOKEN_SCORE_MASK;
                uint currentScore = _tokenCurrentVolume[token];
                _tokenCurrentVolume[token] = currentScore > oldScore ? currentScore - oldScore : 0;
                deque.popFront();
            } else {
                break;
            }
        }

        bytes32 packedMovement = bytes32((currentTime << 192) | score);
        deque.pushBack(packedMovement);

        _tokenCurrentVolume[token] += score;

        if (_periodTotalValueThreshold != 0 && _tokenCurrentVolume[token] > _periodTotalValueThreshold) {
            return (false, "token volume threshold exceeded");
        }
        return (true, "");
    }

    /**
     * @notice Calculates total fees for a bridge operation
     * @param remoteChainID Target chain ID
     * @param token Token being bridged
     * @param value Amount being bridged
     * @return minimumValue Minimum amount required
     * @return gasFee Gas fee for target chain
     * @return exFee Exchange fee
     */
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint gasFee, uint exFee)
    {
        uint exFeeRate;
        (minimumValue, gasFee, exFeeRate) = getTokenConfig(remoteChainID, token);

        exFee = value * exFeeRate / DENOMINATOR;
    }

    /**
     * @notice Gets the token price for standardized unit amount
     * @dev Returns price for a single token unit based on decimals
     * - Returns both existence flag and price
     * - Uses price feed if available
     * @param token Token to get price for
     * @return exist Whether price feed returned a valid price
     * @return price Dollar price for a standard token unit
     */
    function getTokenPrice(IERC20 token) external view returns (bool exist, uint price) {
        return getTokenPriceWithValue(token, (10 ** _priceFeed.decimals(address(token))));
    }

    /**
     * @notice Gets the price of a token with a specific value
     * @dev Used for calculating dollar-denominated values for monitoring
     * - Returns price information with existence flag
     * - Falls back to default price if price feed unavailable
     * - Scales price based on token decimals
     * @param token Token to get price for
     * @param value Amount of tokens to calculate price for
     * @return exist Whether price feed returned a valid price
     * @return price Calculated dollar price for the token amount
     */
    function getTokenPriceWithValue(IERC20 token, uint value) public view returns (bool exist, uint price) {
        if (address(_priceFeed) == address(0)) exist = false;
        else (exist, price,) = _priceFeed.getTokenPriceInDollars(address(token));
        if (!exist) price = _defaultTokenPrice;
        price = price.mulDiv(value, (10 ** _priceFeed.decimals(address(token))));
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
    function getTokenConfig(uint remoteChainID, IERC20 token)
        public
        view
        returns (uint minimumValue, uint gasFee, uint exFeeRate)
    {
        // Calculate minimum token amount based on dollar value
        bool exist;
        uint tokenPrice;

        // Get token price per unit
        if (address(_priceFeed) == address(0)) {
            exist = false;
            tokenPrice = _defaultTokenPrice;
        } else {
            (exist, tokenPrice,) = _priceFeed.getTokenPriceInDollars(address(token));
            if (!exist) tokenPrice = _defaultTokenPrice;
        }

        // Calculate how many tokens are required to meet the minimum dollar value
        // If token price is 0, set a default minimum value to prevent division by zero
        if (tokenPrice == 0) {
            minimumValue = 1e18; // Default to 1 token with 18 decimals if price is unknown
        } else {
            uint tokenDecimals = 10 ** _priceFeed.decimals(address(token));
            minimumValue = _minimumTokenValue * tokenDecimals / tokenPrice;
        }

        // Get exchange fee rate
        exFeeRate = _exFeeRate[token];
        // If exFeeRate is set to max value, set it to 0 (fee exemption)
        if (exFeeRate == type(uint).max) exFeeRate = 0;
        else if (exFeeRate == 0) exFeeRate = _defaultExFeeRate;

        // Calculate gas fee
        (gasFee,) = _calculateGasFee(remoteChainID, token);
    }

    /**
     * @notice Gets the minimum token value in USD
     * @return Minimum token value in USD
     */
    function getMinimumTokenValue() external view returns (uint) {
        return _minimumTokenValue;
    }

    /**
     * @notice Returns the verification amount threshold
     * @return Single transaction threshold in dollar units
     */
    function getVerificationAmountThreshold() external view returns (uint) {
        return _verificationAmountThreshold;
    }

    /**
     * @notice Returns the current time window for monitoring
     * @return Time window in seconds
     */
    function getTimeWindow() external view returns (uint) {
        return _timeWindow;
    }

    /**
     * @notice Returns the period total value threshold
     * @return Total value threshold in dollar units
     */
    function getPeriodTotalValueThreshold() external view returns (uint) {
        return _periodTotalValueThreshold;
    }

    /**
     * @notice Returns current token volume score
     * @param token Token address
     * @return Current volume score in dollar units
     */
    function getTokenCurrentScore(IERC20 token) external view returns (uint) {
        return _tokenCurrentVolume[token];
    }

    /**
     * @notice Returns token movement history
     * @param token Token address
     * @return history Array of packed movement records
     */
    function getTokenMovementHistory(IERC20 token) external view returns (bytes32[] memory history) {
        uint length = _tokenMovementHistory[token].length();
        history = new bytes32[](length);
        for (uint i = 0; i < length; ++i) {
            history[i] = _tokenMovementHistory[token].at(i);
        }
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
     * @notice Updates gas price for an existing chain
     * @dev Modifies gas price and emits event
     * @param remoteChainID Chain ID to update
     * @param gasPrice New gas price value
     */
    function updateGasPrice(uint remoteChainID, uint gasPrice) external onlyRole(UPDATOR_ROLE) {
        _gasPrice[remoteChainID] = gasPrice;
        emit BridgeManagerGasPriceUpdated(remoteChainID, gasPrice);
    }

    /**
     * @notice Sets gas prices for multiple chains at once
     * @dev Batch configuration of chain gas prices
     * - Validates input array lengths match
     * - Updates each chain's gas price
     * @param remoteChainIDs Array of chain IDs
     * @param gasPrices Array of gas prices
     */
    function updateGasPriceBatch(uint[] memory remoteChainIDs, uint[] memory gasPrices)
        external
        onlyRole(UPDATOR_ROLE)
    {
        require(remoteChainIDs.length == gasPrices.length, BridgeManagerInvalidLength());
        for (uint i = 0; i < remoteChainIDs.length; ++i) {
            _gasPrice[remoteChainIDs[i]] = gasPrices[i];
            emit BridgeManagerGasPriceUpdated(remoteChainIDs[i], gasPrices[i]);
        }
    }

    /**
     * @notice Sets token fee parameters
     * @dev Configures token-specific fees
     * - Validates token address
     * - Updates minimum amount and fee rate
     * @param token Token to configure
     * @param exFeeRate Exchange fee rate
     */
    function setExFeeRate(IERC20 token, uint exFeeRate) public onlyRole(UPDATOR_ROLE) {
        require(address(token) != address(0), BridgeManagerCanNotZeroValue("token"));
        _exFeeRate[token] = exFeeRate;
        emit BridgeManagerExchangeFeeUpdated(address(token), exFeeRate);
    }

    /**
     * @notice Sets fee parameters for multiple tokens
     * @dev Batch configuration of token fees
     * - Validates input array lengths match
     * - Updates each token's parameters
     * @param tokenList Array of token addresses
     * @param exFeeRateList Array of exchange fee rates
     */
    function setExFeeRateBatch(IERC20[] memory tokenList, uint[] memory exFeeRateList)
        external
        onlyRole(UPDATOR_ROLE)
    {
        require(tokenList.length == exFeeRateList.length, BridgeManagerInvalidLength());
        for (uint i = 0; i < tokenList.length; ++i) {
            setExFeeRate(tokenList[i], exFeeRateList[i]);
        }
    }

    /**
     * @notice Sets gas amount for bridge finalization
     * @dev Updates gas cost estimation
     * - Validates non-zero value
     * - Updates stored gas amount
     * - Emits update event
     * @param finalizeBridgeGas New gas amount
     */
    function setFinalizeBridgeGas(uint finalizeBridgeGas) external onlyRole(ADMIN_ROLE) {
        require(finalizeBridgeGas != 0, BridgeManagerCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        emit BridgeManagerFinalizeBridgeGasSet(_finalizeBridgeGas);
    }

    /**
     * @notice Sets the default token price when price feed is unavailable
     * @dev Updates the fallback price used when price feed doesn't return a valid price
     * - Updates stored default price
     * - Emits update event
     * @param defaultTokenPrice New default token price value
     */
    function setDefaultTokenPrice(uint defaultTokenPrice) external onlyRole(ADMIN_ROLE) {
        _defaultTokenPrice = defaultTokenPrice;
        emit DefaultTokenPriceSet(defaultTokenPrice);
    }

    /**
     * @notice Sets global exchange fee rate
     * @dev Updates default fee rate
     * - Updates stored rate
     * - Emits update event
     * @param exFeeRate New exchange fee rate
     */
    function setDefaultExFeeRate(uint exFeeRate) external onlyRole(ADMIN_ROLE) {
        _defaultExFeeRate = exFeeRate;
        emit BridgeManagerExchangeFeeUpdated(address(0), _defaultExFeeRate);
    }

    /**
     * @notice Sets the minimum token value in USD
     * @param minimumTokenValue New minimum token value
     */
    function setMinimumTokenValue(uint minimumTokenValue) external onlyRole(ADMIN_ROLE) {
        _minimumTokenValue = minimumTokenValue;
        emit MinimumTokenValueSet(minimumTokenValue);
    }

    /**
     * @notice Sets verification amount threshold for a token pair
     * @dev Updates maximum bridgeable amount
     * - Validates token exists
     * - Updates verification amount threshold value
     * @param verificationAmountThreshold New verification amount threshold value
     */
    function setVerificationAmountThreshold(uint verificationAmountThreshold) external onlyRole(ADMIN_ROLE) {
        _verificationAmountThreshold = verificationAmountThreshold;
        emit VerificationAmountThresholdSet(verificationAmountThreshold);
    }

    /**
     * @notice Sets the time window for monitoring token transfers
     * @dev Updates the duration over which token volumes are tracked
     * - The time window is used for rolling volume monitoring
     * - Measured in seconds
     * - Emits update event
     * @param timeWindow New time window duration in seconds
     */
    function setTimeWindow(uint timeWindow) external onlyRole(ADMIN_ROLE) {
        _timeWindow = timeWindow;
        emit TimeWindowSet(timeWindow);
    }

    /**
     * @notice Sets the maximum value threshold for transfers within a time window
     * @dev Updates the dollar value threshold for all transfers in the monitoring period
     * - Restricts total volume of transfers within the time window
     * - Value is measured in dollar units
     * - Emits update event
     * @param periodTotalValueThreshold New period total value threshold
     */
    function setPeriodTotalValueThreshold(uint periodTotalValueThreshold) external onlyRole(ADMIN_ROLE) {
        _periodTotalValueThreshold = periodTotalValueThreshold;
        emit PeriodTotalValueThresholdSet(periodTotalValueThreshold);
    }

    /**
     * @notice Sets price feed contract
     * @dev Updates price oracle reference
     * - Validates non-zero address
     * - Updates stored reference
     * - Emits update event
     * @param priceFeed New price feed contract
     */
    function setPriceFeed(IPriceFeed priceFeed) external onlyRole(ADMIN_ROLE) {
        require(address(priceFeed) != address(0), BridgeManagerCanNotZeroValue("priceFeed")); // allow zero address
        _priceFeed = priceFeed;
        emit BridgeManagerPriceFeedUpdated(_priceFeed);
    }

    /**
     * @notice Removes price feed integration
     * @dev Disables price-based fee calculation
     * - Sets price feed to zero address
     * - Emits update event
     */
    function removePriceFeed() external onlyRole(ADMIN_ROLE) {
        _priceFeed = IPriceFeed(address(0));
        emit BridgeManagerPriceFeedUpdated(_priceFeed);
    }
}
