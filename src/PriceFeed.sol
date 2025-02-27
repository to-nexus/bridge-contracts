// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ValidatorManager} from "./abstract/ValidatorManager.sol";
import {IPriceFeed} from "./interface/IPriceFeed.sol";
import {PriceFeedLib} from "./lib/PriceFeedLib.sol";

contract PriceFeed is ValidatorManager, IPriceFeed {
    using EnumerableSet for EnumerableSet.AddressSet;

    error PriceFeedCanNotZeroAddress(string name);
    error PriceFeedCanNotZeroValue(string name);
    error PriceFeedNotExistToken(address token);
    error PriceFeedInvalidLength();
    error PriceFeedInvalidPriceAt(uint at, uint blocktime);

    event PriceFeedPriceUpdated(address indexed token, uint price, uint timestamp);

    address internal constant NATIVE_TOKEN = address(1);

    uint8 public dollarDecimals;
    uint public updatedAt;
    EnumerableSet.AddressSet private _tokens;

    mapping(address => PriceData) private _priceData;

    uint[46] private __gap;

    function initialize(uint8 _dollarDecimals) public initializer {
        __Ownable_init(_msgSender());
        __Validator_init(0);

        dollarDecimals = _dollarDecimals;
        updatedAt = block.timestamp;

        // The initial price for NATIVE_TOKEN is set.
        // lastUpdated is initialized to the maximum possible value, ensuring the initial price is considered valid before any updates.
        // (1 xcross == 0.001 dollar)
        _priceData[NATIVE_TOKEN] =
            PriceData({token: NATIVE_TOKEN, price: (10 ** dollarDecimals) / 1000, lastUpdated: block.timestamp});
    }

    function allPrices() external view returns (bool[] memory exist, uint[] memory prices, uint updatedAt_) {
        return getPrices(_tokens.values());
    }

    function getPrice(address token) public view returns (bool exist, uint price, uint updatedAt_) {
        if (!(_priceData[token].token == token)) {
            exist = true;
            price = _priceData[token].price;
        }
        updatedAt_ = updatedAt;
    }

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

    function getPriceData(address token) public view returns (bool exist, PriceData memory data) {
        if (_priceData[token].token == token) {
            exist = true;
            data = _priceData[token];
        }
    }

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

    function updatePrice(address[] memory tokens, uint[] memory prices, uint[] memory pricesAt)
        external
        onlyValidator
    {
        require(tokens.length == prices.length, PriceFeedInvalidLength());
        for (uint i = 0; i < tokens.length;) {
            _updatePrice(tokens[i], prices[i], pricesAt[i]);
            unchecked {
                ++i;
            }
        }
        updatedAt = block.timestamp;
    }

    function nativeCoin() public pure returns (address) {
        return NATIVE_TOKEN;
    }

    function _updatePrice(address token, uint price, uint priceAt) private {
        require(price != 0, PriceFeedCanNotZeroValue("price"));
        require(
            priceAt == type(uint).max || priceAt <= block.timestamp, PriceFeedInvalidPriceAt(priceAt, block.timestamp)
        );
        _tokens.add(token);
        _priceData[token] = PriceData({token: token, price: price, lastUpdated: priceAt});
        emit PriceFeedPriceUpdated(token, price, priceAt);
    }
}
