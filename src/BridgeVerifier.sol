// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {IPriceFeed} from "./PriceFeed.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";
import {CalcAmountLib} from "./lib/CalcAmountLib.sol";
import {Const} from "./lib/Const.sol";

/**
 * @title BridgeVerifier
 * @notice Contract for calculating and managing bridge operation fees
 * @dev Provides fee calculations, token monitoring, and safety threshold management
 * Key features:
 * - Network and exchange fee calculations based on token type and chain
 * - Token volume monitoring for security threshold enforcement
 * - Time-window based transfer volume restrictions
 */
contract BridgeVerifier is AccessControl, IBridgeVerifier {
    using Math for uint;
    using EnumerableSet for EnumerableSet.AddressSet;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using CalcAmountLib for IPriceFeed;

    error BridgeVerifierCanNotZeroValue(string name);
    error BridgeVerifierInvalidLength();
    error BridgeVerifierAlreadyHasRole(address account, bytes32 role);
    error BridgeVerifierDoesNotHaveRole(address account, bytes32 role);
    error BridgeVerifierMissmatchLength();
    error BridgeVerifierInvalidTimeWindow();

    event FinalizeBridgeGasSet(uint finalizeBridgeGas);
    event GasPriceUpdated(uint indexed remoteChainID, uint gasPrice);
    event ExchangeFeeUpdated(address indexed token, uint exFeeRate);
    event PriceFeedUpdated(IPriceFeed indexed priceFeed);
    event VerificationAmountThresholdSet(uint verificationAmountThreshold);
    event DefaultTokenPriceSet(uint defaultTokenPrice);
    event TimeWindowSet(uint timeWindow);
    event PeriodTotalValueThresholdSet(uint periodTotalValueThreshold);
    event MinimumTokenValueSet(uint minimumTokenValue);

    /// @dev Base denominator for fee calculations (10000 = 100%)
    uint private constant DENOMINATOR = 10000;
    uint private constant TOKEN_SCORE_MASK = type(uint).max >> 64;

    /// @dev Price feed contract for token price information
    IPriceFeed public priceFeed;
    /// @dev Gas amount required for finalize bridge operations
    uint private _finalizeBridgeGas;
    /// @dev Global exchange fee rate (default if no token-specific rate)
    uint private _defaultExFeeRate;
    /// @dev Default token price for monitoring when no price feed is available
    uint private _defaultTokenPrice;
    /// @dev Minimum token value for transfer (in dollars)
    uint private _minimumTokenValue;
    /// @dev Maximum dollar value for a single transaction before verification
    /// @notice Limits the dollar value of a single bridge request. Amounts exceeding this threshold
    /// @notice require additional verification procedures to ensure security for unusual large transfers.
    uint private _verificationAmountThreshold;
    /// @dev Maximum dollar value of tokens that can be processed within the time window
    /// @notice Sets the maximum dollar value that can be transferred within the specified time window
    /// @notice without additional verification. If this threshold is exceeded, additional security
    /// @notice checks are triggered to protect against unusual transfer volume spikes.
    uint private _periodTotalValueThreshold;
    /// @dev Time window for monitoring token volume (in seconds)
    /// @notice Defines the duration for which token volume is monitored. The system tracks all
    /// @notice token movements within this rolling time window to enforce the period total value threshold.
    uint private _timeWindow;

    /// @dev Token-specific exchange fee rates
    mapping(IERC20 => uint) private _exFeeRate;
    /// @dev Chain-specific gas prices
    mapping(uint => uint) private _gasPrice;

    /// @dev Current token volume within the time window
    mapping(IERC20 => uint) private _tokenCurrentVolume;
    /// @dev History of token movements for volume tracking
    mapping(IERC20 => DoubleEndedQueue.Bytes32Deque) private _tokenMovementHistory;

    mapping(bytes32 role => EnumerableSet.AddressSet) _roles;

    /**
     * @notice Initializes the BridgeVerifier contract
     * @param initialOwner Admin address
     * @param bridge Bridge contract address
     * @param _priceFeed Price feed contract address
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
        address bridge,
        address _priceFeed,
        uint finalizeBridgeGas,
        uint defaultTokenPrice,
        uint defaultExFeeRate,
        uint minimumTokenValue,
        uint verificationAmountThreshold,
        uint periodTotalValueThreshold,
        uint timeWindow
    ) {
        require(finalizeBridgeGas != 0, BridgeVerifierCanNotZeroValue("finalizeBridgeGas"));
        require(initialOwner != address(0), BridgeVerifierCanNotZeroValue("initialOwner"));
        require(bridge != address(0), BridgeVerifierCanNotZeroValue("bridge"));
        require(_priceFeed != address(0), BridgeVerifierCanNotZeroValue("_priceFeed"));
        require(timeWindow % Const.PERIOD_INTERVAL == 0, BridgeVerifierInvalidTimeWindow());
        _grantRole(DEFAULT_ADMIN_ROLE, initialOwner);
        _grantRole(Const.BRIDGE_ROLE, bridge);

        priceFeed = IPriceFeed(_priceFeed);
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
     * @return status Success status
     */
    function validateBridgeTokenValue(IERC20 token, uint value)
        external
        onlyRole(Const.BRIDGE_ROLE)
        returns (Const.FinalizeStatus status)
    {
        // Verify token price and threshold
        (, uint score) = getTokenPriceWithValue(token, value);
        if (score > type(uint192).max) return Const.FinalizeStatus.TokenScoreOverflow;

        if (_verificationAmountThreshold != 0 && _verificationAmountThreshold < score) {
            status = Const.FinalizeStatus.VerificationAmountThresholdExceeded;
        } else {
            status = Const.FinalizeStatus.Success;
        }

        if (_timeWindow == 0 || _periodTotalValueThreshold == 0) {
            return status;
        } else {
            // Check for potential overflow before adding new score to the current volume
            if (type(uint192).max - score < _tokenCurrentVolume[token]) {
                return Const.FinalizeStatus.TokenCurrentVolumeOverflow;
            }

            // Update current token volume
            _tokenCurrentVolume[token] += score;
        }

        // Manage history and verify thresholds
        DoubleEndedQueue.Bytes32Deque storage deque = _tokenMovementHistory[token];
        // Calculate current time period and the beginning of the monitoring window
        uint currentTime = (block.timestamp / Const.PERIOD_INTERVAL) * Const.PERIOD_INTERVAL;
        uint windowStartTime = currentTime - _timeWindow;

        // Remove expired records (older than the window start time)
        // and subtract their values from the current token volume tracking
        while (deque.length() > 0) {
            bytes32 frontRecord = deque.front();
            // Extract timestamp from the upper 64 bits of the packed record
            uint recordTime = uint(frontRecord >> 192);

            if (recordTime < windowStartTime) {
                // Extract score value from the lower 192 bits using bitmask
                uint oldScore = uint(frontRecord) & TOKEN_SCORE_MASK;
                uint currentScore = _tokenCurrentVolume[token];
                _tokenCurrentVolume[token] = currentScore > oldScore ? currentScore - oldScore : 0;
                deque.popFront();
            } else {
                break;
            }
        }

        // If there's already a record for the current time period,
        // update it by combining with the new score rather than creating a duplicate entry
        if (deque.length() != 0 && uint(deque.back() >> 192) == currentTime) {
            score = (uint(deque.popBack()) & TOKEN_SCORE_MASK) + score;
        }

        // Create a new movement record by packing time (upper 64 bits) and score (lower 192 bits)
        bytes32 packedMovement = bytes32((currentTime << 192) | score);
        deque.pushBack(packedMovement);

        if (status == Const.FinalizeStatus.Success && _tokenCurrentVolume[token] > _periodTotalValueThreshold) {
            status = Const.FinalizeStatus.PeriodTotalValueThresholdExceeded;
        }

        return status;
    }

    /**
     * @notice Calculates total fees for a bridge operation
     * @param remoteChainID Target chain ID
     * @param token Token being bridged
     * @param value Amount being bridged
     * @return minimumValue Minimum amount required
     * @return networkFee Network fee for target chain
     * @return exFee Exchange fee
     */
    function calculateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimumValue, uint networkFee, uint exFee)
    {
        uint exFeeRate;
        (minimumValue, networkFee, exFeeRate) = getTokenConfig(remoteChainID, token);

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
        return getTokenPriceWithValue(token, (10 ** CalcAmountLib.decimals(address(token))));
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
        (exist, price,) = priceFeed.getTokenPriceInDollars(address(token));
        if (!exist) price = _defaultTokenPrice;
        price = price.mulDiv(value, (10 ** CalcAmountLib.decimals(address(token))));
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
     * - Calculates network fee using price feed
     * - Handles fee exemptions
     * @param remoteChainID Chain ID for fee calculation
     * @param token Token to calculate fees for
     * @return minimumValue Minimum required amount
     * @return networkFee Network fee in token amount
     * @return exFeeRate Exchange fee rate
     */
    function getTokenConfig(uint remoteChainID, IERC20 token)
        public
        view
        returns (uint minimumValue, uint networkFee, uint exFeeRate)
    {
        // Get token price per unit
        (bool exist, uint tokenPrice,) = priceFeed.getTokenPriceInDollars(address(token));
        if (!exist) tokenPrice = _defaultTokenPrice;

        uint tokenDecimals = 10 ** CalcAmountLib.decimals(address(token));
        // Calculate how many tokens are required to meet the minimum dollar value
        // If token price is 0, set a default minimum value to prevent division by zero
        if (tokenPrice == 0) minimumValue = tokenDecimals; // Default to 1 token with decimals if price is unknown

        else minimumValue = _minimumTokenValue * tokenDecimals / tokenPrice;

        // Get exchange fee rate
        exFeeRate = _exFeeRate[token];
        // If exFeeRate is set to max value, set it to 0 (fee exemption)
        if (exFeeRate == type(uint).max) exFeeRate = 0;
        else if (exFeeRate == 0) exFeeRate = _defaultExFeeRate;

        // Calculate network fee
        (networkFee,) = _calculateNetworkFee(remoteChainID, token);
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

    function dollarDecimals() external view returns (uint) {
        return priceFeed.dollarDecimals();
    }

    /**
     * @notice Returns the fee denominator
     * @dev Used for fee calculations (10000 = 100%)
     * @return Denominator constant value
     */
    function denominator() external pure returns (uint) {
        return DENOMINATOR;
    }

    /**
     * @notice Calculates network fee in token amount
     * @dev Uses price feed for conversion
     * - Validates price feed availability
     * - Converts gas cost to token amount
     * @param remoteChainID Chain ID for gas price
     * @param token Token to calculate fee in
     * @return networkFee Calculated network fee
     * @return updatedAt Timestamp of price data
     */
    function _calculateNetworkFee(uint remoteChainID, IERC20 token)
        private
        view
        returns (uint networkFee, uint updatedAt)
    {
        uint gasPrice = _gasPrice[remoteChainID];
        if (gasPrice == 0) return (0, 0); // If gas price is 0, return 0
        (, networkFee, updatedAt) =
            priceFeed.calculateTokenAmountForNetworkFee(remoteChainID, address(token), _finalizeBridgeGas * gasPrice);
    }

    /**
     * @notice Updates gas price for an existing chain
     * @dev Modifies gas price and emits event
     * @param remoteChainID Chain ID to update
     * @param gasPrice New gas price value
     */
    function updateGasPrice(uint remoteChainID, uint gasPrice) external onlyRole(Const.PRICER_ROLE) {
        _gasPrice[remoteChainID] = gasPrice;
        emit GasPriceUpdated(remoteChainID, gasPrice);
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
        onlyRole(Const.PRICER_ROLE)
    {
        require(remoteChainIDs.length == gasPrices.length, BridgeVerifierInvalidLength());
        for (uint i = 0; i < remoteChainIDs.length; ++i) {
            _gasPrice[remoteChainIDs[i]] = gasPrices[i];
            emit GasPriceUpdated(remoteChainIDs[i], gasPrices[i]);
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
    function setExFeeRate(IERC20 token, uint exFeeRate) public onlyRole(Const.EDITOR_ROLE) {
        require(address(token) != address(0), BridgeVerifierCanNotZeroValue("token"));
        _exFeeRate[token] = exFeeRate;
        emit ExchangeFeeUpdated(address(token), exFeeRate);
    }

    /**
     * @notice Sets fee parameters for multiple tokens
     * @dev Batch configuration of token fees
     * - Validates input array lengths match
     * - Updates each token's parameters
     * @param tokenList Array of token addresses
     * @param exFeeRateList Array of exchange fee rates
     */
    function setExFeeRateBatch(IERC20[] memory tokenList, uint[] memory exFeeRateList) external {
        require(tokenList.length == exFeeRateList.length, BridgeVerifierInvalidLength());
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
    function setFinalizeBridgeGas(uint finalizeBridgeGas) external onlyRole(Const.ADMIN_ROLE) {
        require(finalizeBridgeGas != 0, BridgeVerifierCanNotZeroValue("finalizeBridgeGas"));
        _finalizeBridgeGas = finalizeBridgeGas;
        emit FinalizeBridgeGasSet(_finalizeBridgeGas);
    }

    /**
     * @notice Sets the default token price when price feed is unavailable
     * @dev Updates the fallback price used when price feed doesn't return a valid price
     * - Updates stored default price
     * - Emits update event
     * @param defaultTokenPrice New default token price value
     */
    function setDefaultTokenPrice(uint defaultTokenPrice) external onlyRole(Const.EDITOR_ROLE) {
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
    function setDefaultExFeeRate(uint exFeeRate) external onlyRole(Const.EDITOR_ROLE) {
        _defaultExFeeRate = exFeeRate;
        emit ExchangeFeeUpdated(address(0), _defaultExFeeRate);
    }

    /**
     * @notice Sets the minimum token value in USD
     * @param minimumTokenValue New minimum token value
     */
    function setMinimumTokenValue(uint minimumTokenValue) external onlyRole(Const.EDITOR_ROLE) {
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
    function setVerificationAmountThreshold(uint verificationAmountThreshold) external onlyRole(Const.ADMIN_ROLE) {
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
    function setTimeWindow(uint timeWindow) external onlyRole(Const.ADMIN_ROLE) {
        require(timeWindow % Const.PERIOD_INTERVAL == 0, BridgeVerifierInvalidTimeWindow());
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
    function setPeriodTotalValueThreshold(uint periodTotalValueThreshold) external onlyRole(Const.ADMIN_ROLE) {
        _periodTotalValueThreshold = periodTotalValueThreshold;
        emit PeriodTotalValueThresholdSet(periodTotalValueThreshold);
    }

    /**
     * @notice Sets price feed contract
     * @dev Updates price oracle reference
     * - Validates non-zero address
     * - Updates stored reference
     * - Emits update event
     * @param _priceFeed New price feed contract
     */
    function setPriceFeed(IPriceFeed _priceFeed) external onlyRole(Const.ADMIN_ROLE) {
        require(address(_priceFeed) != address(0), BridgeVerifierCanNotZeroValue("_priceFeed")); // allow zero address
        priceFeed = _priceFeed;
        emit PriceFeedUpdated(priceFeed);
    }

    /**
     * @notice Returns all members of a specific role
     * @param role The role to query
     * @return members Array of addresses that have the specified role
     */
    function getRoleMembers(bytes32 role) external view returns (address[] memory) {
        return _roles[role].values();
    }

    /**
     * @notice Grants multiple roles to a list of accounts
     * @dev Validates input array lengths match
     * - Grants each role to the corresponding account
     * @param roles Array of roles to grant
     * @param accounts Array of accounts to grant roles to
     */
    function grantRoleBatch(bytes32[] memory roles, address[] memory accounts) external {
        require(roles.length == accounts.length, BridgeVerifierMissmatchLength());
        for (uint i = 0; i < accounts.length; ++i) {
            grantRole(roles[i], accounts[i]);
        }
    }

    /**
     * @notice Revokes multiple roles from a list of accounts
     * @dev Validates input array lengths match
     * - Revokes each role from the corresponding account
     * @param roles Array of roles to revoke
     * @param accounts Array of accounts to revoke roles from
     */
    function revokeRoleBatch(bytes32[] memory roles, address[] memory accounts) external {
        require(roles.length == accounts.length, BridgeVerifierMissmatchLength());
        for (uint i = 0; i < accounts.length; ++i) {
            revokeRole(roles[i], accounts[i]);
        }
    }

    function _grantRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._grantRole(role, account);
        if (ok) require(_roles[role].add(account), BridgeVerifierAlreadyHasRole(account, role));
    }

    function _revokeRole(bytes32 role, address account) internal override returns (bool ok) {
        ok = super._revokeRole(role, account);
        if (ok) require(_roles[role].remove(account), BridgeVerifierDoesNotHaveRole(account, role));
    }
}
