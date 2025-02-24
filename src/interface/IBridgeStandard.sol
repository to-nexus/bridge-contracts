// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IBridgeTokenInfo} from "./IBridgeTokenInfo.sol";
import {IIndexer} from "./IIndexer.sol";
import {ITokenManager} from "./ITokenManager.sol";
import {IValidatorManager} from "./IValidatorManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

interface IBridgeStandard is ITokenManager, IIndexer, IValidatorManager {
    struct FinalizeArguments {
        uint index;
        IERC20 token;
        address to;
        uint value;
        bytes[] extraData;
    }

    struct PermitArguments {
        IERC20Permit token;
        address account;
        uint value;
        uint deadline;
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    function bridge(IERC20 token, uint value, uint gas, uint ex, bytes[] calldata extraData)
        external
        payable
        returns (bool);
    function bridgeTo(IERC20 token, address to, uint value, uint gas, uint ex, bytes[] calldata extraData)
        external
        payable
        returns (bool);
    function permitBridge(
        IERC20 token,
        address account,
        uint value,
        uint gas,
        uint ex,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) external payable returns (bool);
    function permitBridgeTo(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint ex,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) external payable returns (bool);
    function finalize(
        uint index,
        IERC20 token,
        address to,
        uint value,
        bytes[] calldata extraData,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) external payable returns (bool);
    function finalizeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool);
    function retryFinalize(uint index) external returns (bool);
    function retryFinalizeBatch(uint[] memory indexes) external returns (bool);
    function initializedAt() external view returns (uint);
    function domainSeparator() external view returns (bytes32);
    function rewardWallet() external view returns (address payable);
    function denominator() external view returns (uint);
    function calculate(IERC20 token, uint value)
        external
        view
        returns (uint minimum, uint gas, uint ex, bool isValid);
    function revertedArguments(uint index) external view returns (FinalizeArguments memory);
    function revertedReason(uint index) external view returns (bytes memory);
    function addToken(IERC20 token, IERC20 pair) external;
    function removeToken(IERC20 token) external;
    function pauseToken(IERC20 token) external;
    function unpauseToken(IERC20 token) external;
    function pause() external;
    function unpause() external;
    function setTokenInfo(IBridgeTokenInfo _bridgeTokenInfo) external;
    function removeTokenInfo() external;
}
