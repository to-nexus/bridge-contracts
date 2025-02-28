// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {IPriceFeed} from "../interface/IPriceFeed.sol";

library PriceFeedLib {
    using Math for uint;

    error PriceFeedLibCanNotZeroValue(string name);
    error PriceFeedLibOverflow();

    /**
     * @notice Calculates an amount of token A to an equivalent amount of token B based on their prices.
     * @param feed The IPriceFeed contract instance.
     * @param tokenA The address of token A.
     * @param tokenB The address of token B.
     * @param amountA The amount of token A to calculate.
     * @return ok True if both token prices are valid, false otherwise.
     * @return amountB The equivalent amount of token B.
     * @return updatedAt The updated time.
     */
    function calculateAmountB(IPriceFeed feed, address tokenA, address tokenB, uint amountA)
        external
        view
        returns (bool ok, uint amountB, uint updatedAt)
    {
        if (tokenA == tokenB) return (true, amountA, block.timestamp);

        address[] memory tokens = new address[](2);
        tokens[0] = tokenA;
        tokens[1] = tokenB;

        bool[] memory exist;
        uint[] memory prices;
        (exist, prices, updatedAt) = feed.getPrices(tokens);
        ok = exist[0] && exist[1];
        if (!ok) return (false, 0, 0);

        (uint8 decimalA, uint8 decimalB) = (_decimals(feed, tokenA), _decimals(feed, tokenB));
        amountB = calculateAmountBWithPrice(amountA, prices[0], prices[1], decimalA, decimalB);
    }

    /// @notice Calculates an amount of token A to an equivalent amount of token B using provided price data.
    /// @param amountA The amount of token A to calculate.
    /// @param priceA The price of token A.
    /// @param priceB The price of token B.
    /// @param decimalA The number of decimals for token A.
    /// @param decimalB The number of decimals for token B.
    /// @return amountB The equivalent amount of token B.
    function calculateAmountBWithPrice(uint amountA, uint priceA, uint priceB, uint8 decimalA, uint8 decimalB)
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
        decimals = token == feed.nativeToken() ? uint8(18) : IERC20Metadata(token).decimals();
    }
}
