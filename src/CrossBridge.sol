// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1967Utils} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {StandardBridge} from "./abstract/StandardBridge.sol";
import {ICrossMintableERC20} from "./token/ICrossMintableERC20.sol";

/**
 * @title CrossBridge
 * @notice This contract implements the StandardBridge for a specific cross-chain bridge.
 */
contract CrossBridge is StandardBridge {
    error CrossBridgeInvalidValue(uint expected, uint actual);
    error CrossBridgeInvalidValueUnit(uint value);
    error CrossBridgeCanNotZeroAddress(string name);

    uint private constant EX_RATE = 100; // cross : xcross, 1 : 100
    address private constant PREDEPLOY_ADDRESS = address(0x0956d70000000000000000000000000000000101); // predeployed implementation address

    uint private _ethereumChainID;

    uint[49] private __gap;

    function initialize(uint8 _threshold, address _rewardWallet, address _crossMintableERC20FactoryCode)
        external
        initializer
    {
        __StandardBridge_init(_threshold, _rewardWallet, _crossMintableERC20FactoryCode);
    }

    /**
     * @dev Override for predeploy. When upgrading the implementation, remove it.
     */
    function _checkProxy() internal view override {
        if (
            address(this) == PREDEPLOY_ADDRESS // Must be called through delegatecall
                || ERC1967Utils.getImplementation() != PREDEPLOY_ADDRESS // Must be called through an active proxy
        ) revert UUPSUnauthorizedCallContext();
    }
}
