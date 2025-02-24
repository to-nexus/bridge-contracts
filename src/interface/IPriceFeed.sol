// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ITokenStorage} from "./ITokenStorage.sol";
import {IValidatorManager} from "./IValidatorManager.sol";

interface IPriceFeed is ITokenStorage, IValidatorManager {
    struct PriceData {
        address token;
        uint price;
        uint lastUpdated;
    }

    function coin() external view returns (address);
    function dollarDecimals() external view returns (uint);
    function deadline() external view returns (uint);
    function getPrice(address token) external view returns (PriceData memory data);
    function getValidPrice(address token) external view returns (uint price);
    function getPrices(address[] memory token) external view returns (PriceData[] memory data);
}
