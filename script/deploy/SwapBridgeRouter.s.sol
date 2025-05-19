// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {SwapBridgeRouter} from "../../src/SwapBridgeRouter.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

contract SwapBridgeRouterScript is Script {
    function setUp() public {}

    /**
     * @notice Setup SwapBridgeRouter contract and initialize
     * command
     * forge script ./script/deploy/SwapBridgeRouter.s.sol:SwapBridgeRouterScript \
     * -f $RPC_URL \
     * --sender $SENDER \
     * --sig "deploySwapBridgeRouter(address,address,address,address,uint256)" \
     */
    function deploySwapBridgeRouter(
        address owner,
        address bridge,
        address swapRouter,
        address crossToken,
        uint crossChainID
    ) public returns (address) {
        vm.broadcast();
        SwapBridgeRouter swapBridgeRouter = new SwapBridgeRouter(owner, bridge, swapRouter, crossToken, crossChainID);
        console.log("SwapBridgeRouter will be deployed to", address(swapBridgeRouter));
        return address(swapBridgeRouter);
    }
}
