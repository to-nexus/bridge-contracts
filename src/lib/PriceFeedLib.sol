// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IPriceFeed} from "../interface/IPriceFeed.sol";
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

        uint deadline = feed.deadline();

        address[] memory tokens = new address[](2);
        tokens[0] = tokenA;
        tokens[1] = tokenB;
        IPriceFeed.PriceData[] memory data = feed.getPrices(tokens);
        require(data[0].token == tokenA && data[1].token == tokenB);

        isValid = (data[0].lastUpdated == type(uint).max || data[0].lastUpdated + deadline >= block.timestamp)
            && (data[1].lastUpdated == type(uint).max || data[1].lastUpdated + deadline > block.timestamp);
        (uint8 decimalA, uint8 decimalB) = (_decimals(feed, tokenA), _decimals(feed, tokenB));
        amountB = convertAtoBWithPrice(amountA, data[0].price, data[1].price, decimalA, decimalB);
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
