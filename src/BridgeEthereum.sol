// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeStandard} from "./abstract/BridgeStandard.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title BridgeEthereum
 * @notice This contract acts as a bridge contract on the Ethereum network, facilitating cross-chain token transfers.
 * It inherits from BridgeStandard and implements the necessary functions for initiating and finalizing bridges.
 */
contract BridgeEthereum is BridgeStandard {
    using Address for address payable;
    using SafeERC20 for IERC20;

    error BridgeEthereumCanNotZeroAddress(string name);
    error BridgeEthereumInvalidValue(uint expected, uint actual);

    IERC20 public cross; // The ERC20 token representing the wrapped native token on the destination chain.

    uint[49] private __gap;

    /**
     * @notice Initializes the BridgeEthereum contract.
     * @param _cross The address of the wrapped native token on the destination chain (e.g., WETH).
     * @param rewardWallet_ The address of the reward wallet to receive fees.
     * @param _bridgeTokenInfo The address of the BridgeTokenInfo contract.
     */
    function initialize(IERC20 _cross, uint8 _threshold, address rewardWallet_, address _bridgeTokenInfo)
        public
        initializer
    {
        require(address(_cross) != address(0), BridgeEthereumCanNotZeroAddress("_cross"));
        __BridgeStandard_init(_threshold, rewardWallet_, _bridgeTokenInfo);
        cross = _cross;
        addToken(_cross, IERC20(coin));
    }

    /**
     * @notice Initiates a cross-chain bridge transaction from Ethereum.
     * @param account The address initiating the bridge (both sender and recipient on Ethereum side if `to` is not specified in `bridgeToCross`).
     * @param value The amount of `cross` tokens to bridge.
     * @param permitArgs Permit arguments for approving token spending. See `PermitArguments` struct.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge initiation was successful.
     */
    function bridgeCross(address account, uint value, PermitArguments memory permitArgs, bytes[] calldata extraData)
        public
        returns (bool)
    {
        return bridgeToCross(account, account, value, permitArgs, extraData);
    }

    /**
     * @notice Initiates a cross-chain bridge transaction from Ethereum to a specific recipient on Ethereum.
     * @param from The address initiating the bridge (sender of tokens on Ethereum side).
     * @param to The address of the recipient on Ethereum side.
     * @param value The amount of `cross` tokens to bridge.
     * @param permitArgs Permit arguments for approving token spending. See `PermitArguments` struct.
     * @param extraData Additional data for the bridge transaction.
     * @return True if the bridge initiation was successful.
     */
    function bridgeToCross(
        address from,
        address to,
        uint value,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) public returns (bool) {
        return permitBridgeTo(cross, from, to, value, 0, 0, permitArgs, extraData);
    }

    /**
     * @notice Handles the initiation of a bridge transaction on the source chain (Ethereum).
     * @param token The address of the token being bridged (should be `cross`).
     * @param from The address of the user initiating the bridge.
     * @param value The amount of tokens being bridged.
     * @param fee The total fees (gas + ex) for the bridge transaction.
     */
    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal override {
        if (address(token) == coin) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(msg.value == value + fee, BridgeEthereumInvalidValue(value + fee, msg.value)); // Verify correct amount received
            if (fee != 0) rewardWallet().sendValue(fee); // Send fees to the reward wallet
        } else {
            // Transfers the tokens and fees from the user to the contract.
            token.safeTransferFrom(from, address(this), value + fee);
            // Transfers the fees to the reward wallet.
            if (fee > 0) token.safeTransfer(rewardWallet(), fee);
        }
    }

    /**
     * @notice Handles the finalization of a bridge transaction on the destination chain.
     * @param token The address of the token being bridged.
     * @param to The address of the recipient on the destination chain.
     * @param value The amount of tokens being bridged.
     * @return ok True if the finalization was successful, false otherwise.
     * @return reason The reason for failure if the finalization was unsuccessful.
     */
    function _finalizeBridge(IERC20 token, address to, uint value)
        internal
        override
        returns (bool ok, bytes memory reason)
    {
        if (value == 0) return (true, "");
        if (address(token) == coin) {
            // Handling native token transfers
            payable(to).sendValue(value); // Send native tokens to the recipient
            ok = true;
            reason = "";
        } else {
            // Adjust the value based on the exchange rate if the token is the CROSS.
            if (address(token) == address(cross)) value = value / EX_RATE;
            try IERC20(token).transfer(to, value) returns (bool success) {
                if (success) {
                    ok = true;
                    reason = "";
                } else {
                    ok = false;
                    reason = "BridgeEthereum: transfer failed";
                }
                // Catch potential errors during token transfer and provide reasons.
            } catch Error(string memory _reason) {
                ok = false;
                reason = bytes(_reason);
            } catch Panic(uint _errorCode) {
                ok = false;
                reason = abi.encodePacked(_errorCode);
            } catch (bytes memory _lowLevelData) {
                ok = false;
                reason = _lowLevelData;
            }
        }
    }
}
