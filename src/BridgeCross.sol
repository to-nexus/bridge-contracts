// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

import {BridgeStandard} from "./abstract/BridgeStandard.sol";
import {ICrossMintableERC20, ICrossMintableERC20Code} from "./interface/ICrossMintableERC20.sol";

/**
 * @title BridgeCross
 * @notice This contract implements the BridgeStandard for a specific cross-chain bridge.
 * It handles the initiation and finalization of cross-chain transfers, including minting and burning of wrapped tokens.
 * It uses Create2 for deploying new wrapped tokens.
 */
contract BridgeCross is BridgeStandard {
    using Address for address payable;
    using SafeERC20 for IERC20;

    error BridgeCrossInvalidValueUnit(uint value);
    error BridgeCrossInvalidValue(uint expected, uint actual);
    error BridgeCrossBurnFailed(address token, address from, uint value);

    address public weth;
    ICrossMintableERC20Code private _crossMintableERC20Code; // Bytecode for deploying new cross-mintable ERC20 tokens

    uint[48] private __gap;

    receive() external payable {
        assert(msg.value != 0); // Ensure the receive function only accepts non-zero value
    }

    /**
     * @notice Initializes the BridgeCross contract.
     * @param crossMintableERC20Code The address of the contract containing the bytecode for cross-mintable ERC20 tokens.
     * @param rewardWallet_ The address of the reward wallet.
     * @param _bridgeTokenInfo The address of the BridgeTokenInfo contract.
     */
    function initialize(address crossMintableERC20Code, address rewardWallet_, address _bridgeTokenInfo)
        external
        initializer
    {
        __BridgeStandard_init(rewardWallet_, _bridgeTokenInfo);
        _crossMintableERC20Code = ICrossMintableERC20Code(crossMintableERC20Code);
        weth = addTokenDeploy(IERC20(address(1)), "ETH", 18);
    }

    /**
     * @notice Deploys a new wrapped token using Create2.
     * @param pair The address of the corresponding token on the destination chain.
     * @param symbol The symbol of the new wrapped token.
     * @param decimals The number of decimals for the new wrapped token.
     * @return tokenAddress The address of the newly deployed wrapped token.
     */
    function addTokenDeploy(IERC20 pair, string memory symbol, uint8 decimals)
        public
        onlyOwner
        returns (address tokenAddress)
    {
        string memory name = string(abi.encodePacked("Cross Bridge ", symbol));
        bytes32 salt = keccak256(abi.encodePacked(pair)); // Create a deterministic salt based on the paired token
        bytes memory bytecode = abi.encodePacked(_crossMintableERC20Code.code(), abi.encode(name, symbol, decimals)); // Combine creation code and constructor arguments
        tokenAddress = Create2.deploy(0, salt, bytecode); // Deploy the wrapped token using Create2
        addToken(IERC20(tokenAddress), pair); // Register the new token with the bridge
    }

    /**
     * @notice Handles the initiation of a bridge transaction on the source chain.
     * @param token The address of the token being bridged.
     * @param from The address of the user initiating the bridge.
     * @param value The amount of tokens being bridged.
     * @param fee The total fees (gas + ex) for the bridge transaction.
     */
    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal override {
        if (address(token) == coin) {
            // Handling native token transfers (e.g., CROSS, ETH, BNB)
            require(value % EX_RATE == 0, BridgeCrossInvalidValueUnit(msg.value));
            require(msg.value == value + fee, BridgeCrossInvalidValue(value + fee, msg.value));
            if (fee != 0) rewardWallet().sendValue(fee); // Send fees to the reward wallet
        } else {
            // Handling ERC20 token transfers
            if (fee != 0) token.safeTransferFrom(from, rewardWallet(), fee); // Transfer fees to the reward wallet
            require(
                ICrossMintableERC20(address(token)).burn(from, value), // Burn the wrapped tokens on the source chain
                BridgeCrossBurnFailed(address(token), from, value)
            );
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
            payable(to).sendValue(value * EX_RATE); // Send native tokens to the recipient
            ok = true;
            reason = "";
        } else {
            // Handling ERC20 token transfers
            try ICrossMintableERC20(address(token)).mint(to, value) returns (bool success) {
                // Mint wrapped tokens on the destination chain
                if (success) {
                    ok = true;
                    reason = "";
                } else {
                    ok = false;
                    reason = "BridgeCross: mint failed";
                }
                // Catch potential errors during minting and provide revert reasons
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
