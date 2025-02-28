// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {StandardBridge} from "./abstract/StandardBridge.sol";

/**
 * @title EthereumBridge
 * @notice This contract acts as a bridge contract on the Ethereum network, facilitating cross-chain token transfers.
 */
contract EthereumBridge is StandardBridge {
    error EthereumBridgeCanNotZeroAddress(string name);
    error EthereumBridgeInvalidValue(uint expected, uint actual);

    uint private constant EX_RATE = 100; // cross : xcross, 1 : 100

    IERC20 private _cross; // The ERC20 token representing the wrapped native token on the destination chain.

    uint[49] private __gap;

    function initialize(
        uint crossChainID,
        IERC20 cross,
        uint8 _threshold,
        address _rewardWallet,
        address _crossMintableERC20FactoryCode,
        address _bridgeFeeStation
    ) external initializer {
        __StandardBridge_init(_threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation);

        require(address(cross) != address(0), EthereumBridgeCanNotZeroAddress("cross"));
        _cross = cross;

        setChain(crossChainID);
        registerToken(crossChainID, true, address(cross), NATIVE_TOKEN);
    }

    function _finalizeBridge(uint remoteChainID, IERC20 token, address to, uint value)
        internal
        override
        returns (bool ok, string memory reason)
    {
        if (address(token) == address(_cross)) value = value / EX_RATE;
        return super._finalizeBridge(remoteChainID, token, to, value);
    }
}
