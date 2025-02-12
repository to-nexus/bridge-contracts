// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {BridgeStandard} from "./abstract/BridgeStandard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableMap} from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

contract BridgeEthereum is BridgeStandard {
    using SafeERC20 for IERC20;

    IERC20 public cross; // @DEV: eth는 weth? eth를 받아서 wrap을 해줄지?

    function initialize(IERC20 _cross, address rewardWallet_, address BridgeFeeManager) public initializer {
        __BridgeStandard_init(rewardWallet_, BridgeFeeManager);
        cross = IERC20(_cross);
    }

    function bridgeCross(address account, uint value, uint deadline, bytes memory permitSig, bytes[] calldata extraData)
        public
        returns (bool)
    {
        return bridgeToCross(account, account, value, deadline, permitSig, extraData);
    }

    function bridgeToCross(
        address from,
        address to,
        uint value,
        uint deadline,
        bytes memory permitSig,
        bytes[] calldata extraData
    ) public returns (bool) {
        return permitBridgeTo(cross, from, to, value, 0, 0, deadline, permitSig, extraData);
    }

    function _initiateBridge(IERC20 token, address from, uint value, uint fee) internal override {
        token.safeTransferFrom(from, address(this), value + fee);
        if (fee > 0) token.safeTransfer(rewardWallet(), fee);
    }

    function _finalizeBridge(IERC20 token, address to, uint value)
        internal
        override
        returns (bool ok, bytes memory reason)
    {
        if (value > 0) {
            if (address(token) == address(cross)) value = value / _exrate;
            try IERC20(token).transfer(to, value) returns (bool success) {
                if (success) {
                    ok = true;
                    reason = "";
                } else {
                    ok = false;
                    reason = "BridgeEthereum: transfer failed";
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
