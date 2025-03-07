pragma solidity 0.8.11;

// https://github.com/alpaca-finance/bsc-alpaca-contract v1.2.30
// github.com/alpaca-finance/bsc-alpaca-contract/contracts/8.11/protocol/interfaces/IPriceOracle.sol

interface IPriceOracle {
    /// @dev Return the wad price of token0/token1, multiplied by 1e18
    /// NOTE: (if you have 1 token0 how much you can sell it for token1)
    function getPrice(address token0, address token1) external view returns (uint price, uint lastUpdate);
}
