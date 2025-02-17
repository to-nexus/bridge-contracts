// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./IIndexer.sol";
import "./ITokenManager.sol";
import "./IValidatorManager.sol";
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

    function bridge(IERC20 token, uint value, uint gas, uint service, bytes[] calldata extraData)
        external
        payable
        returns (bool);
    function bridgeTo(IERC20 token, address to, uint value, uint gas, uint service, bytes[] calldata extraData)
        external
        payable
        returns (bool);
    function permitBridge(
        IERC20 token,
        address account,
        uint value,
        uint gas,
        uint service,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) external payable returns (bool);
    function permitBridgeTo(
        IERC20 token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        PermitArguments memory permitArgs,
        bytes[] calldata extraData
    ) external payable returns (bool);
    function finalize(
        uint index,
        IERC20 token,
        address to,
        uint value,
        bytes[] calldata extraData,
        bytes[] calldata sigs
    ) external payable returns (bool);
    function finalizeBatch(FinalizeArguments[] calldata args, bytes[][] memory sigs) external payable returns (bool);
    function addToken(IERC20 token, IERC20 pair) external;
    function removeToken(IERC20 token) external;
    function pauseToken(IERC20 token) external;
    function unpauseToken(IERC20 token) external;
    function calculateFee(IERC20 token, uint value) external view returns (uint gas, uint service);
}
