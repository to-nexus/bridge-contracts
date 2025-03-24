// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {RoleManager} from "./abstract/RoleManager.sol";
import {IPriceFeed} from "./interface/IPriceFeed.sol";
import {IPriceOracle} from "./interface/IPriceOracle.sol";
import {CalcAmountLib} from "./lib/CalcAmountLib.sol";
import {Const} from "./lib/Const.sol";

/**
 * @title PriceFeed
 * @notice Oracle contract that manages price data for tokens across different chains
 * @dev Provides token price data management with the following features:
 * - Price storage and retrieval for ERC20 and native tokens
 * - Price updates with timestamp tracking
 */
contract PriceFeed is UUPSUpgradeable, RoleManager, IPriceFeed {
    using EnumerableSet for EnumerableSet.AddressSet;

    error PriceFeedCanNotZeroValue(string name);
    error PriceFeedInvalidLength();
    error PriceFeedInvalidPriceAt(uint at, uint blocktime);
    error PriceFeedNoSource(address token);

    event PriceUpdated(address indexed token, uint price, uint priceAt);
    event NativeTokenPriceUpdated(uint indexed chainID, uint price, uint priceAt);

    /// @dev Decimal places used for dollar price representation
    uint8 public dollarDecimals;
    /// @dev Timestamp when prices were last updated
    uint public updatedAt;
    /// @dev Set of all tokens tracked by this price feed
    EnumerableSet.AddressSet private _tokens;

    /// @dev Mapping from token address to its price data
    mapping(address => PriceData) private _priceData;
    /// @dev Mapping from chain ID to native token price for that chain
    mapping(uint => NativeTokenPriceData) private _nativeTokenPrice;

    uint[44] private __gap;

    /**
     * @notice Initializes the price feed contract
     * @param owner Admin address with price update privileges
     * @param _dollarDecimals Precision for price representation
     */
    function initialize(address owner, uint8 _dollarDecimals) public initializer {
        __UUPSUpgradeable_init();
        __RoleManager_init(owner);
        _grantRole(Const.PRICER_ROLE, owner);

        dollarDecimals = _dollarDecimals;
        updatedAt = block.timestamp;
    }

    /**
     * @notice Returns prices for all tracked tokens
     * @return exist Array of price existence flags
     * @return prices Array of token prices
     * @return updatedAt_ Last price update timestamp
     */
    function allPrices() external view returns (bool[] memory exist, uint[] memory prices, uint updatedAt_) {
        return getPrices(_tokens.values());
    }

    /**
     * @notice Gets the exchange rate between two tokens
     * @dev Calculates relative value using dollar prices of both tokens
     * @param tokenA Source token address
     * @param tokenB Target token address
     * @return price How much of tokenB equals 1 tokenA
     * @return lastUpdate Last price update time
     */
    function getPrice(address tokenA, address tokenB) external view returns (uint price, uint lastUpdate) {
        if (tokenA == tokenB) return (1, block.number);

        address[] memory tokens = new address[](2);
        tokens[0] = tokenA;
        tokens[1] = tokenB;

        bool[] memory exist;
        uint[] memory prices;
        (exist, prices, lastUpdate) = getPrices(tokens);
        require(exist[0], PriceFeedNoSource(tokenA));
        require(exist[1], PriceFeedNoSource(tokenB));

        (uint8 decimalA, uint8 decimalB) = (CalcAmountLib.decimals(tokenA), CalcAmountLib.decimals(tokenB));
        price = CalcAmountLib.calculateAmountBWithPrice(1, prices[0], prices[1], decimalA, decimalB);
        return (price, updatedAt);
    }

    /**
     * @notice Gets the dollar price for a specific token
     * @param token Token address
     * @return exist Whether price data exists
     * @return price Dollar price of the token
     * @return updatedAt_ Last price update timestamp
     */
    function getTokenPriceInDollars(address token) external view returns (bool exist, uint price, uint updatedAt_) {
        if (_priceData[token].token != token) return (false, 0, 0);
        return (true, _priceData[token].price, updatedAt);
    }

    /**
     * @notice Gets the native token price for a specific chain
     * @param chainID Chain identifier
     * @return exist Whether price data exists
     * @return price Dollar price of the native token
     * @return updatedAt_ Last price update timestamp
     */
    function getNativeTokenPrice(uint chainID) external view returns (bool exist, uint price, uint updatedAt_) {
        NativeTokenPriceData memory priceData = _nativeTokenPrice[chainID];
        if (priceData.chainID != chainID) return (false, 0, 0);
        return (true, priceData.price, updatedAt);
    }

    /**
     * @notice Gets prices for a list of tokens in dollars
     * @param tokens Array of token addresses
     * @return exist Array of existence flags
     * @return prices Array of token prices
     * @return updatedAt_ Last price update timestamp
     */
    function getPrices(address[] memory tokens)
        public
        view
        returns (bool[] memory exist, uint[] memory prices, uint updatedAt_)
    {
        exist = new bool[](tokens.length);
        prices = new uint[](tokens.length);
        for (uint i = 0; i < tokens.length;) {
            if (_priceData[tokens[i]].token == tokens[i]) {
                exist[i] = true;
                prices[i] = _priceData[tokens[i]].price;
            }
            unchecked {
                ++i;
            }
        }
        updatedAt_ = updatedAt;
    }

    /**
     * @notice Gets detailed price data for a specific token
     * @dev Returns the full PriceData struct and existence flag
     * @param token Address of the token to query
     * @return exist Boolean indicating if price data exists
     * @return data PriceData struct containing price and last update timestamp
     */
    function getPriceData(address token) public view returns (bool exist, PriceData memory data) {
        if (_priceData[token].token == token) {
            exist = true;
            data = _priceData[token];
        }
    }

    /**
     * @notice Gets detailed price data for multiple tokens
     * @dev Returns arrays of existence flags and PriceData structs
     * @param tokens Array of token addresses to query
     * @return exist Array of booleans indicating if price data exists for each token
     * @return list Array of PriceData structs for each token
     */
    function getPriceDataList(address[] memory tokens)
        external
        view
        returns (bool[] memory exist, PriceData[] memory list)
    {
        exist = new bool[](tokens.length);
        list = new PriceData[](tokens.length);
        for (uint i = 0; i < tokens.length;) {
            if (_priceData[tokens[i]].token == tokens[i]) {
                exist[i] = true;
                list[i] = _priceData[tokens[i]];
            }
            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice Updates prices for multiple tokens
     * @dev Only callable by validators, updates stored price data
     * @param tokens Array of token addresses to update
     * @param prices Array of new prices for each token
     * @param pricesAt Array of timestamps for when each price was recorded
     */
    function updatePrice(address[] memory tokens, uint[] memory prices, uint[] memory pricesAt)
        external
        onlyRole(Const.PRICER_ROLE)
    {
        uint tokensLen = tokens.length;
        require(tokensLen == prices.length && tokensLen == pricesAt.length, PriceFeedInvalidLength());
        for (uint i = 0; i < tokensLen;) {
            _updatePrice(tokens[i], prices[i], pricesAt[i]);
            unchecked {
                ++i;
            }
        }
        updatedAt = block.timestamp;
    }

    /**
     * @notice Updates native token prices for multiple chains
     * @dev Only callable by validators, updates stored native token prices
     * @param chainIDs Array of chain IDs to update
     * @param prices Array of new prices for each chain's native token
     * @param pricesAt Array of timestamps for when each price was recorded
     */
    function updateNativeTokenPrice(uint[] memory chainIDs, uint[] memory prices, uint[] memory pricesAt)
        external
        onlyRole(Const.PRICER_ROLE)
    {
        uint chainsLen = chainIDs.length;
        require(chainsLen == prices.length && chainsLen == pricesAt.length, PriceFeedInvalidLength());
        for (uint i = 0; i < chainsLen;) {
            _updateNativeTokenPrice(chainIDs[i], prices[i], pricesAt[i]);
            unchecked {
                ++i;
            }
        }
        updatedAt = block.timestamp;
    }

    /**
     * @notice Internal function to update a token's price
     * @dev Validates price and timestamp, updates storage, and emits event
     * @param token Address of the token to update
     * @param price New price for the token
     * @param priceAt Timestamp when the price was recorded
     */
    function _updatePrice(address token, uint price, uint priceAt) private {
        require(price != 0, PriceFeedCanNotZeroValue("price"));
        require(priceAt <= block.timestamp, PriceFeedInvalidPriceAt(priceAt, block.timestamp));
        _tokens.add(token);
        _priceData[token] = PriceData({token: token, price: price, lastUpdated: priceAt});
        emit PriceUpdated(token, price, priceAt);
    }

    /**
     * @notice Internal function to update a chain's native token price
     * @dev Validates price and timestamp, updates storage
     * @param chainID ID of the chain to update
     * @param price New price for the chain's native token
     * @param priceAt Timestamp when the price was recorded
     */
    function _updateNativeTokenPrice(uint chainID, uint price, uint priceAt) private {
        require(price != 0, PriceFeedCanNotZeroValue("price"));
        require(priceAt <= block.timestamp, PriceFeedInvalidPriceAt(priceAt, block.timestamp));
        _nativeTokenPrice[chainID] = NativeTokenPriceData({chainID: chainID, price: price, lastUpdated: priceAt});
        emit NativeTokenPriceUpdated(chainID, price, priceAt);
    }

    /**
     * @notice Authorizes an upgrade to a new implementation.
     * @param _newImplementation The address of the new implementation.
     */
    function _authorizeUpgrade(address _newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
