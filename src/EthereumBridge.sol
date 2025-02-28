// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {StandardBridge} from "./abstract/StandardBridge.sol";

/**
 * @title EthereumBridge
 * @notice This contract acts as a bridge contract on the Ethereum network, facilitating cross-chain token transfers.
 */
contract EthereumBridge is StandardBridge {
    uint[50] private __gap;

    function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode)
        external
        initializer
    {
        __StandardBridge_init(_threshold, _rewardWallet, _crossMintableERC20FactoryCode);
    }
}
