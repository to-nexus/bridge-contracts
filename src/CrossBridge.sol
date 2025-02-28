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

    function initialize(
        uint ethereumChainID,
        IERC20 cross,
        uint8 _threshold,
        address _rewardWallet,
        address _crossMintableERC20FactoryCode,
        address _bridgeFeeStation
    ) external initializer {
        __StandardBridge_init(_threshold, _rewardWallet, _crossMintableERC20FactoryCode, _bridgeFeeStation);

        require(address(cross) != address(0), CrossBridgeCanNotZeroAddress("cross"));
        _ethereumChainID = ethereumChainID;

        setChain(ethereumChainID);
        registerToken(_ethereumChainID, false, NATIVE_TOKEN, address(cross));
    }

    function _initiateBridge(uint remoteChainID, IERC20 token, address from, uint value, uint fee) internal override {
        if (remoteChainID == _ethereumChainID && address(token) == NATIVE_TOKEN) {
            require(value % EX_RATE == 0, CrossBridgeInvalidValueUnit(msg.value));
        }
        super._initiateBridge(remoteChainID, token, from, value, fee);
    }

    function _finalizeBridge(uint remoteChainID, IERC20 token, address to, uint value)
        internal
        override
        returns (bool ok, string memory reason)
    {
        if (address(token) == NATIVE_TOKEN) value = value * EX_RATE;
        return super._finalizeBridge(remoteChainID, token, to, value);
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
