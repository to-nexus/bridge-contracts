// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "./IIndexer.sol";
import "./ITokenManager.sol";
import "./IValidatorManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBridgeStandard is ITokenManager, IIndexer, IValidatorManager {
    struct FinalizeArguments {
        uint index;
        IERC20 token;
        address to;
        uint value;
        bytes[] extraData;
    }

    function bridge(IERC20 token, uint value, uint gas, uint service, bytes[] calldata extraData)
        external
        payable
        returns (bool);
    function bridgeTo(IERC20 token, address to, uint value, uint gas, uint service, bytes[] calldata extraData)
        external
        payable
        returns (bool);
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
