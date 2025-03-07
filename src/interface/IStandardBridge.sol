// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

import {IBridgeFeeStation} from "./IBridgeFeeStation.sol";
import {IBridgeRegistry} from "./IBridgeRegistry.sol";
import {IRoleManager} from "./IRoleManager.sol";

interface IStandardBridge is IBridgeRegistry {
    struct BridgeTokenArguments {
        uint remoteChainID;
        IERC20 token;
        address to;
        uint value;
        uint gasFee;
        uint exFee;
        bytes extraData;
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

    function bridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData
    ) external payable returns (bool);
    function permitBridgeToken(
        uint remoteChainID,
        IERC20 token,
        address to,
        uint value,
        uint gasFee,
        uint exFee,
        bytes calldata extraData,
        PermitArguments calldata permitArgs
    ) external payable returns (bool);
    function permitBridgeTokenBatch(BridgeTokenArguments[] calldata args, PermitArguments[] calldata permitArgs)
        external
        payable;
    function finalizeBridge(FinalizeArguments calldata args, uint8[] memory v, bytes32[] memory r, bytes32[] memory s)
        external
        payable
        returns (bool);
    function finalizeBridgeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool);
    function bridgeFeeStation() external view returns (IBridgeFeeStation);
    function retryFinalizeBridge(uint remoteChainID, uint index) external returns (bool);
    function retryFinalizeBridgeBatch(uint remoteChainID, uint[] memory indexes) external returns (bool);
    function domainSeparator() external view returns (bytes32);
    function initializedAt() external view returns (uint);
    function estimateFee(uint remoteChainID, IERC20 token, uint value)
        external
        view
        returns (uint minimum, uint gasFee, uint exFee);
}
