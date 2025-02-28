// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {StandardBridge} from "./abstract/StandardBridge.sol";

/**
 * @title BranchBridge
 * @notice This contract acts as a bridge contract on the remote network, facilitating cross-chain token transfers.
 * It inherits from StandardBridge and implements the necessary functions for initiating and finalizing bridges.
 */
contract BranchBridge is StandardBridge {
    uint[50] private __gap;

    function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode)
        external
        initializer
    {
        __StandardBridge_init(_threshold, _rewardWallet, _crossMintableERC20FactoryCode);
    }
}
