// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Const} from "../lib/Const.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IBridgeRegistry {
    enum PendingType {
        TokenPaused,
        VerificationAmountThresholdExceeded,
        PeriodTotalValueThresholdExceeded,
        TokenCurrentVolumeOverflow,
        TransferFailed,
        MintFailed
    }

    struct FinalizeArguments {
        uint fromChainID;
        uint index;
        IERC20 toToken;
        address to;
        uint value;
        bytes extraData;
    }

    struct Chain {
        uint remoteChainID;
        uint initiateIndex;
        uint finalizeIndex;
        bool paused;
    }

    struct TokenPair {
        address localToken; // local token address
        address remoteToken; // remote token address
        bool isOrigin; // whether the token is origin token
        bool paused; // whether the token is paused
        uint deposited; // deposited amount of the token
        uint pendingAmount; // pending amount of the token
    }

    struct PendingData {
        FinalizeArguments args;
        Const.FinalizeStatus status;
        uint delayExpiration;
        bytes reason;
    }

    function allChainIDs() external view returns (uint[] memory);
    function allTokenPairs(uint remoteChainID) external view returns (TokenPair[] memory);
    function allPendingIndex(uint remoteChainID) external view returns (uint[] memory);
    function getTokenPair(uint remoteChainID, address token) external view returns (TokenPair memory);
    function getNextInitiateIndex(uint remoteChainID) external view returns (uint);
    function getNextFinalizeIndex(uint remoteChainID) external view returns (uint);
    function isPending(uint remoteChainID, uint index) external view returns (bool);
    function getPendingArguments(uint remoteChainID, uint index) external view returns (PendingData memory);
}
