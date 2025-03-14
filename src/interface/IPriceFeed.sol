// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IPriceOracle} from "./IPriceOracle.sol";

interface IPriceFeed is IPriceOracle {
    struct PriceData {
        address token;
        uint price;
        uint lastUpdated;
    }

    function allPrices() external view returns (bool[] memory exist, uint[] memory prices, uint updatedAt_);
    function getTokenPriceInDollars(address token) external view returns (bool exist, uint price, uint updatedAt_);
    function getNativeTokenPrice(uint chainID) external view returns (bool exist, uint price, uint updatedAt_);
    function getPrices(address[] memory tokens)
        external
        view
        returns (bool[] memory exist, uint[] memory prices, uint updatedAt_);
    function getPriceData(address token) external view returns (bool exist, PriceData memory data);
    function getPriceDataList(address[] memory tokens)
        external
        view
        returns (bool[] memory exist, PriceData[] memory list);
    function updatePrice(address[] memory tokens, uint[] memory _prices, uint[] memory pricesAt) external;
}
