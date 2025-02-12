// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";

import {BridgeStandard} from "./abstract/BridgeStandard.sol";
import {ICrossMintableERC20, ICrossMintableERC20Code} from "./interface/ICrossMintableERC20.sol";

contract BridgeCross is BridgeStandard {
    using Address for address payable;
    using SafeERC20 for IERC20;

    error BridgeCrossInvalidValueUnit(uint value);
    error BridgeCrossInsufficientValue(uint expected, uint actual);
    error BridgeCrossBurnFailed(address token, address from, uint value);

    address public xcross;
    ICrossMintableERC20Code private _crossMintableERC20Code;

    function initialize(address crossMintableERC20Code, address rewardWallet_, address BridgeFeeManager)
        external
        initializer
    {
        __BridgeStandard_init(rewardWallet_, BridgeFeeManager);
        xcross = address(1);
        _crossMintableERC20Code = ICrossMintableERC20Code(crossMintableERC20Code);
    }

    function addTokenDeploy(IERC20 pair, string memory symbol, uint8 decimals)
        public
        onlyOwner
        returns (address tokenAddress)
    {
        string memory name = string(abi.encodePacked("Cross Bridge ", symbol));
        bytes32 salt = keccak256(abi.encodePacked(pair));
        bytes memory bytecode = abi.encodePacked(_crossMintableERC20Code.code(), abi.encode(name, symbol, decimals));
        tokenAddress = Create2.deploy(0, salt, bytecode);
        addToken(IERC20(tokenAddress), pair);
    }

    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal override {
        if (address(token) == xcross) {
            require((msg.value / _exrate) * _exrate == msg.value, BridgeCrossInvalidValueUnit(msg.value));
            require(msg.value == value + fee, BridgeCrossInsufficientValue(value + fee, msg.value));
            if (fee > 0) rewardWallet().sendValue(fee);
        } else {
            if (fee > 0) token.safeTransferFrom(from, rewardWallet(), fee);
            require(
                ICrossMintableERC20(address(token)).burn(from, value),
                BridgeCrossBurnFailed(address(token), from, value)
            );
        }
    }

    function _finalizeBridge(IERC20 token, address to, uint value)
        internal
        override
        returns (bool ok, bytes memory reason)
    {
        if (value > 0) {
            if (address(token) == xcross) {
                payable(to).sendValue(value * _exrate);
            } else {
                try ICrossMintableERC20(address(token)).mint(to, value) returns (bool success) {
                    if (success) {
                        ok = true;
                        reason = "";
                    } else {
                        ok = false;
                        reason = "BridgeCross: mint failed";
                    }
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
}
