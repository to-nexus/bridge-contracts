// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BaseBridge} from "../../src/BaseBridge.sol";
import {BridgeManager, IBridgeManager} from "../../src/BridgeManager.sol";

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {TestToken} from "../token/TestToken.sol";
import {CrossChainTest} from "./CrossChain.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract EthereumChainTest is CrossChainTest {
    bool internal bridgeRevertEthereum = false;
    bool internal finalizeRevertEthereum = false;

    uint internal nextIndexEthereum;
    BaseBridge internal bridgeEthereum;
    IBridgeManager internal bridgeManagerEthereum;

    function setUp() public virtual override {
        super.setUp();
        nextIndexEthereum = 1;
        vm.selectFork(ethereumForkID);
        vm.startPrank(OWNER);

        // bridge setup
        {
            BaseBridge bridgeEthereumImpl = new BaseBridge();
            ERC1967Proxy bridgeEthereumProxy = new ERC1967Proxy(address(bridgeEthereumImpl), bytes(""));
            bridgeEthereum = BaseBridge(payable(address(bridgeEthereumProxy)));
            bridgeEthereum.initialize(OWNER, threshold, REWARD);

            bridgeEthereum.grantRole(ADMIN_ROLE, OWNER); // for test
            bridgeEthereum.grantRole(OPERATOR_ROLE, OWNER); // for test
            bridgeEthereum.grantRole(UPDATOR_ROLE, OWNER); // for test
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address(cross), address(NATIVE_TOKEN), -(int(EX_RATE)));
            bridgeEthereum.grantRoleBatch(VALIDATOR_ROLE, VALIDATORS);
        }

        // add token to bridge (ethereum chain)
        {
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address((NATIVE_TOKEN)), address(weth), 0);
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address(testTokenEthereum), address(testTokenCross), 0);

            TestToken(address(cross)).mint(OWNER, INITIAL_SUPPLY);
            TestToken(address(testTokenEthereum)).mint(OWNER, INITIAL_SUPPLY);
        }
        vm.stopPrank();
    }

    function ethereumIncrementIndex() public {
        nextIndexEthereum++;
    }

    // ----- Functions -----
    function ethereumBridge(address token, address from, address to, uint value, uint gas, uint service)
        public
        returns (uint index, bool ok)
    {
        vm.selectFork(ethereumForkID);
        // bridge
        index = nextIndexEthereum;
        vm.prank(from);

        if (bridgeRevertEthereum) {
            bridgeRevertEthereum = false;
            vm.expectRevert();
        }

        if (token == address(NATIVE_TOKEN)) {
            assertTrue(from.balance >= value + gas + service);
            ok = bridgeEthereum.bridgeToken{value: value + gas + service}(
                CROSS_CHAIN_ID, IERC20(token), to, value, gas, service, NULLDATA
            );
        } else {
            ok = bridgeEthereum.bridgeToken(CROSS_CHAIN_ID, IERC20(token), to, value, gas, service, NULLDATA);
        }
    }

    function ethereumFinalize(uint index, address token, address to, uint value, uint sigCount)
        public
        returns (bool ok)
    {
        vm.selectFork(ethereumForkID);
        if (sigCount > threshold) sigCount = threshold;

        // create finalize validator signature
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, to, value, NULLDATA));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeEthereum.domainSeparator(), h);

        uint8[] memory v = new uint8[](sigCount);
        bytes32[] memory r = new bytes32[](sigCount);
        bytes32[] memory s = new bytes32[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], hash);
        }

        // finalize
        vm.prank(VALIDATOR1);
        if (finalizeRevertEthereum) {
            finalizeRevertEthereum = false;
            vm.expectRevert();
        }
        ok = bridgeEthereum.finalizeBridge(
            IBridgeRegistry.FinalizeArguments({
                fromChainID: CROSS_CHAIN_ID,
                index: index,
                toToken: IERC20(token),
                to: to,
                value: value,
                extraData: NULLDATA
            }),
            v,
            r,
            s
        );
    }

    function ethereumCalcFee(IERC20, uint) public pure returns (uint value, uint gas, uint ex) {
        return (value, 0, 0); // no fee
            // vm.selectFork(ethereumForkID);
            // if (address(bridgeManagerEthereum) == address(0)) return (value, 0, 0);
            // bool ok;
            // (value, gas, ex, ok) = bridgeManagerEthereum.calculateMax(token, totalValue);
            // assertTrue(ok);
    }
}
