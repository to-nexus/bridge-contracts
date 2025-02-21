// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {TokenStorage} from "./abstract/TokenStorage.sol";
import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {IPriceFeed} from "./interface/IPriceFeed.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

library PriceFeedLib {
    using Math for uint;

    error PriceFeedLibCanNotZeroValue(string name);
    error PriceFeedLibOverflow();

    /// @notice Retrieves price data for all tracked tokens.
    /// @param feed The PriceFeed contract instance.
    /// @return prices An array of PriceData structs containing price information for each token.
    function allPrices(IPriceFeed feed) external view returns (IPriceFeed.PriceData[] memory prices) {
        return feed.getPrices(feed.allTokens());
    }

    /// @notice Converts an amount of token A to an equivalent amount of token B based on their prices.
    /// @param feed The PriceFeed contract instance.
    /// @param tokenA The address of token A.
    /// @param tokenB The address of token B.
    /// @param amountA The amount of token A to convert.
    /// @return amountB The equivalent amount of token B.
    /// @return isValid True if both token prices are valid, false otherwise.
    function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint amountA)
        external
        view
        returns (uint amountB, bool isValid)
    {
        if (tokenA == tokenB) return (amountA, true);

        (uint priceA, bool isValidA) = feed.getValidPrice(tokenA);
        (uint priceB, bool isValidB) = feed.getValidPrice(tokenB);
        isValid = isValidA && isValidB;
        (uint8 decimalA, uint8 decimalB) = (_decimals(feed, tokenA), _decimals(feed, tokenB));
        amountB = convertAtoBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
    }

    /// @notice Converts an amount of token A to an equivalent amount of token B using provided price data.
    /// @param amountA The amount of token A to convert.
    /// @param priceA The price of token A.
    /// @param priceB The price of token B.
    /// @param decimalA The number of decimals for token A.
    /// @param decimalB The number of decimals for token B.
    /// @return amountB The equivalent amount of token B.
    function convertAtoBWithPrice(uint amountA, uint priceA, uint priceB, uint8 decimalA, uint8 decimalB)
        public
        pure
        returns (uint amountB)
    {
        require(priceA != 0, PriceFeedLibCanNotZeroValue("priceA"));
        bool ok;
        (ok, amountB) = decimalB >= decimalA
            ? amountA.tryMul(priceB.mulDiv(10 ** (decimalB - decimalA), priceA))
            : amountA.tryMul(priceB / (10 ** (decimalA - decimalB)) / priceA);
        require(ok, PriceFeedLibOverflow());
    }

    /// @notice Retrieves the number of decimals for a given token.
    /// @param token The address of the token.
    /// @return decimals The number of decimals.
    function _decimals(IPriceFeed feed, address token) private view returns (uint8 decimals) {
        decimals = token == feed.coin() ? uint8(18) : IERC20Metadata(token).decimals();
    }
}

contract PriceFeedCross is TokenStorage, ValidatorManager, IPriceFeed {
    error PriceFeedCanNotZeroAddress(string name);
    error PriceFeedCanNotZeroValue(string name);
    error PriceFeedNotExistToken(address token);
    error PriceFeedInvalidLength();
    error PriceFeedInvalidPriceAt(uint at, uint blocktime);

    event PriceFeedPriceUpdated(address indexed token, uint price, uint timestamp);
    event PriceFeedDeadlineChanged(uint deadline);

    address public coin;
    uint public deadline;
    uint public dollarDecimals;
    mapping(address => PriceData) private prices;

    uint[46] private __gap;

    /// @notice Initializes the contract with default values.
    function initialize() public initializer {
        __Validator_init(_msgSender());
        deadline = 2 hours;
        dollarDecimals = 6;
        coin = address(1);
        _addToken(coin);

        // The initial price for coin is set.
        // lastUpdated is initialized to the maximum possible value, ensuring the initial price is considered valid before any updates.
        // (1 xcross == 0.001 dollar)
        prices[coin] = PriceData(coin, (10 ** dollarDecimals) / 1000, type(uint).max);
    }

    /// @notice Retrieves the valid price for a given token.
    /// @param token The address of the token.
    /// @return price The price of the token.
    /// @return isValid True if the price is valid, false otherwise.
    function getValidPrice(address token) external view returns (uint price, bool isValid) {
        require(contains(token), PriceFeedNotExistToken(token));
        PriceData memory data = prices[token];
        price = data.price;
        if (data.lastUpdated == type(uint).max || block.timestamp <= data.lastUpdated + deadline) isValid = true;
    }

    /// @notice Retrieves the price data for multiple tokens.
    /// @param token An array of token addresses.
    /// @return data An array of PriceData structs containing price information for each token.
    function getPrices(address[] memory token) external view returns (PriceData[] memory data) {
        data = new PriceData[](token.length);
        for (uint i = 0; i < token.length;) {
            data[i] = getPrice(token[i]);
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Retrieves the price data for a single token.
    /// @param token The address of the token.
    /// @return data A PriceData struct containing the price information for the token.
    function getPrice(address token) public view returns (PriceData memory data) {
        require(contains(token), PriceFeedNotExistToken(token));
        return prices[token];
    }

    /// @notice Updates the price for a given token. Only callable by a validator.
    /// @param token The address of the token.
    /// @param price The new price of the token.
    /// @param priceAt The timestamp of the price update. Must be less than or equal to the current block timestamp.
    function updatePrice(address token, uint price, uint priceAt) public onlyValidator {
        require(contains(token), PriceFeedNotExistToken(token));
        _updatePrice(token, price, priceAt);
    }

    /// @notice Batch updates the price for multiple tokens.  Only callable by a validator.
    /// @param tokens An array of token addresses.
    /// @param _prices An array of new prices for the tokens.
    /// @param pricesAt An array of timestamps for the price updates. Each timestamp must be less than or equal to the current block timestamp.
    function updatePriceBatch(address[] memory tokens, uint[] memory _prices, uint[] memory pricesAt) external {
        require(tokens.length == _prices.length, PriceFeedInvalidLength());
        for (uint i = 0; i < tokens.length;) {
            updatePrice(tokens[i], _prices[i], pricesAt[i]);
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Adds a new token to the price feed. Only callable by the contract owner.
    /// @param token The address of the new token.
    /// @param price The initial price of the token.
    /// @param priceAt The timestamp of the initial price.  Must be greater than or equal to the current block timestamp.  Consider removing this requirement.
    function addToken(address token, uint price, uint priceAt) external onlyOwner {
        require(token != address(0), PriceFeedCanNotZeroAddress("token"));
        _addToken(token);
        _updatePrice(token, price, priceAt);
    }

    /// @notice Removes a token from the price feed. Only callable by the contract owner.
    /// @param token The address of the token to remove.
    function removeToken(address token) external onlyOwner {
        _removeToken(token);
        delete prices[token];
    }

    /// @notice Changes the price validity deadline. Only callable by the contract owner.
    /// @param _deadline The new price validity deadline.
    function changeDeadline(uint _deadline) external onlyOwner {
        deadline = _deadline;
        emit PriceFeedDeadlineChanged(_deadline);
    }

    function _updatePrice(address token, uint price, uint priceAt) private {
        require(contains(token), PriceFeedNotExistToken(token));
        require(price != 0, PriceFeedCanNotZeroValue("price"));
        require(
            priceAt == type(uint).max || priceAt <= block.timestamp, PriceFeedInvalidPriceAt(priceAt, block.timestamp)
        );
        prices[token] = PriceData(token, price, priceAt);
        emit PriceFeedPriceUpdated(token, price, priceAt);
    }
}
