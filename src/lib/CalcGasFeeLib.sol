// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {IPriceFeed} from "../interface/IPriceFeed.sol";

library CalcGasFeeLib {
    using Math for uint;

    error CalcGasFeeLibCanNotZeroValue(string name);
    error CalcGasFeeLibOverflow();

    /**
     * @notice Calculates an amount of token A to an equivalent amount of token B based on their prices.
     * @param feed The IPriceFeed contract instance.
     * @param toChainID The chain ID of the destination chain.
     * @param token The address of token A.
     * @param nativeTokenAmount The amount of native token to calculate.
     * @return ok True if both token prices are valid, false otherwise.
     * @return amountB The equivalent amount of token B.
     * @return updatedAt The updated time.
     */
    function calculateTokenAmountForGasFee(IPriceFeed feed, uint toChainID, address token, uint nativeTokenAmount)
        external
        view
        returns (bool ok, uint amountB, uint updatedAt)
    {
        uint nativeTokenPrice;
        (ok, nativeTokenPrice, updatedAt) = feed.getNativeTokenPrice(toChainID);
        if (!ok) return (false, 0, 0);

        uint price;
        (ok, price, updatedAt) = feed.getTokenPriceInDollars(token);
        if (!ok) return (false, 0, 0);

        amountB = calculateAmountBWithPrice(nativeTokenAmount, nativeTokenPrice, price, 18, decimals(feed, token));
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
        require(priceA != 0, CalcGasFeeLibCanNotZeroValue("priceA"));
        bool ok;
        (ok, amountB) = decimalB >= decimalA
            ? amountA.tryMul(priceB.mulDiv(10 ** (decimalB - decimalA), priceA))
            : amountA.tryMul(priceB / (10 ** (decimalA - decimalB)) / priceA);
        require(ok, CalcGasFeeLibOverflow());
    }

    /// @notice Retrieves the number of decimals for a given token.
    /// @param token The address of the token.
    /// @return _decimals The number of decimals.
    function decimals(IPriceFeed feed, address token) public view returns (uint8 _decimals) {
        _decimals = token == feed.nativeToken() ? uint8(18) : IERC20Metadata(token).decimals();
    }
}
