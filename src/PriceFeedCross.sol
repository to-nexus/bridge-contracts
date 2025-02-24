// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {TokenStorage} from "./abstract/TokenStorage.sol";
import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {IPriceFeed} from "./interface/IPriceFeed.sol";
import {PriceFeedLib} from "./lib/PriceFeedLib.sol";

contract PriceFeedCross is TokenStorage, ValidatorManager, IPriceFeed {
    error PriceFeedCanNotZeroAddress(string name);
    error PriceFeedCanNotZeroValue(string name);
    error PriceFeedNotExistToken(address token);
    error PriceFeedInvalidLength();
    error PriceFeedInvalidPriceAt(uint at, uint blocktime);
    error PriceFeedNotValidPrice(address token, uint deadline);

    event PriceFeedPriceUpdated(address indexed token, uint price, uint timestamp);
    event PriceFeedDeadlineChanged(uint deadline);

    address public constant coin = address(1);
    uint public constant dollarDecimals = 6;
    uint public deadline;
    mapping(address => PriceData) private prices;

    uint[46] private __gap;

    /// @notice Initializes the contract with default values.
    function initialize(uint8 _threshold) public initializer {
        __Validator_init(_msgSender(), _threshold);
        deadline = 2 hours;
        _addToken(coin);

        // The initial price for coin is set.
        // lastUpdated is initialized to the maximum possible value, ensuring the initial price is considered valid before any updates.
        // (1 xcross == 0.001 dollar)
        prices[coin] = PriceData({token: coin, price: (10 ** dollarDecimals) / 1000, lastUpdated: type(uint).max});
    }

    /// @notice Retrieves the price data for multiple tokens.
    /// @param token An array of token addresses.
    /// @return data An array of PriceData structs containing price information for each token.
    function getPrices(address[] memory token) external view returns (PriceData[] memory data) {
        data = new PriceData[](token.length);
        for (uint i = 0; i < token.length;) {
            data[i] = prices[token[i]];
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Retrieves the valid price for a given token.
    /// @param token The address of the token.
    /// @return price The price of the token.
    function getValidPrice(address token) external view returns (uint price) {
        require(contains(token), PriceFeedNotExistToken(token));
        PriceData memory data = prices[token];
        require(
            (data.lastUpdated == type(uint).max || block.timestamp <= data.lastUpdated + deadline),
            PriceFeedNotValidPrice(token, data.lastUpdated + deadline)
        );
        price = data.price;
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
        prices[token] = PriceData({token: token, price: price, lastUpdated: priceAt});
        emit PriceFeedPriceUpdated(token, price, priceAt);
    }
}
