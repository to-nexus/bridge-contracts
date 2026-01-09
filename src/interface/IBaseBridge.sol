// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";

import {IBridgeRegistry} from "./IBridgeRegistry.sol";
import {IBridgeVerifier} from "./IBridgeVerifier.sol";

interface IBaseBridge is IBridgeRegistry {
    struct BridgeTokenArguments {
        uint toChainID;
        IERC20 fromToken;
        address from;
        address to;
        uint value;
        uint networkFee;
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
        uint toChainID,
        IERC20 fromToken,
        address to,
        uint value,
        uint networkFee,
        uint exFee,
        bytes calldata extraData
    ) external payable returns (bool);
    function permitBridgeTokenBatch(BridgeTokenArguments[] calldata args, PermitArguments[] calldata permitArgs)
        external
        payable;
    function finalizeBridgeBatch(
        FinalizeArguments[] calldata args,
        uint8[][] memory v,
        bytes32[][] memory r,
        bytes32[][] memory s
    ) external payable returns (bool);
    function bridgeVerifier() external view returns (IBridgeVerifier);
    function releasePending(uint remoteChainID, uint index) external;
    function domainSeparator() external view returns (bytes32);
    function initializedAt() external view returns (uint);
}
