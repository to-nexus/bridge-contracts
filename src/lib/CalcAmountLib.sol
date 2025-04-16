// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {IPriceFeed} from "../interface/IPriceFeed.sol";
import {Const} from "./Const.sol";

library CalcAmountLib {
    using Math for uint;

    error CalcAmountLibCanNotZeroValue(string name);
    error CalcAmountLibOverflow();

    uint8 constant NATIVE_TOKEN_DECIMALS = 18;

    /**
     * @notice Calculates the equivalent amount of token based on the native token amount.
     * @param feed The IPriceFeed contract instance.
     * @param toChainID The chain ID of the destination chain.
     * @param token The address of token.
     * @param nativeTokenAmount The amount of native token to calculate.
     * @return ok True if both token prices are valid, false otherwise.
     * @return tokenAmount The equivalent amount of token.
     * @return updatedAt The updated time.
     */
    function calculateTokenAmountForNetworkFee(IPriceFeed feed, uint toChainID, address token, uint nativeTokenAmount)
        internal
        view
        returns (bool ok, uint tokenAmount, uint updatedAt)
    {
        uint nativeTokenPrice;
        (ok, nativeTokenPrice, updatedAt) = feed.getNativeTokenPrice(toChainID);
        if (!ok) return (false, 0, 0);

        uint price;
        (ok, price, updatedAt) = feed.getTokenPriceInDollars(token);
        if (!ok) return (false, 0, 0);

        tokenAmount = calculateAmountBWithPrice(
            nativeTokenAmount, nativeTokenPrice, price, NATIVE_TOKEN_DECIMALS, decimals(token)
        );
    }

    /// @notice Calculates an amount of token A to an equivalent amount of token B using provided price data.
    /// @param amountA The amount of token A to calculate.
    /// @param priceA The price of token A.
    /// @param priceB The price of token B.
    /// @param decimalA The number of decimals for token A.
    /// @param decimalB The number of decimals for token B.
    /// @return amountB The equivalent amount of token B.
    function calculateAmountBWithPrice(uint amountA, uint priceA, uint priceB, uint8 decimalA, uint8 decimalB)
        internal
        pure
        returns (uint amountB)
    {
        require(priceA != 0, CalcAmountLibCanNotZeroValue("priceA"));
        require(priceB != 0, CalcAmountLibCanNotZeroValue("priceB"));

        if (decimalA == decimalB) return amountA.mulDiv(priceA, priceB);
        amountB = decimalB > decimalA
            ? amountA.mulDiv(priceA * 10 ** (decimalB - decimalA), priceB)
            : amountA.mulDiv(priceA, priceB * 10 ** (decimalA - decimalB));
    }

    /// @notice Retrieves the number of decimals for a given token.
    /// @param token The address of the token.
    /// @return _decimals The number of decimals.
    function decimals(address token) internal view returns (uint8 _decimals) {
        _decimals = token == Const.NATIVE_TOKEN ? uint8(NATIVE_TOKEN_DECIMALS) : IERC20Metadata(token).decimals();
    }
}
