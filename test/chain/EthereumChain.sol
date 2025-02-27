// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeFeeStation, IBridgeFeeStation} from "../../src/BridgeFeeStation.sol";
import {EthereumBridge} from "../../src/EthereumBridge.sol";
import {TestToken} from "../token/TestToken.sol";
import {CrossChainTest} from "./CrossChain.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract EthereumChainTest is CrossChainTest {
    bool internal bridgeRevertEthereum = false;
    bool internal finalizeRevertEthereum = false;

    uint internal nextIndexEthereum;
    EthereumBridge internal bridgeEthereum;
    IBridgeFeeStation internal bridgeFeeStationEthereum;

    function setUp() public virtual override {
        super.setUp();
        nextIndexEthereum = 1;
        vm.selectFork(ethereumForkID);
        vm.startPrank(OWNER);

        // bridge setup
        {
            EthereumBridge bridgeEthereumImpl = new EthereumBridge();
            ERC1967Proxy bridgeEthereumProxy = new ERC1967Proxy(address(bridgeEthereumImpl), bytes(""));
            bridgeEthereum = EthereumBridge(payable(address(bridgeEthereumProxy)));
            bridgeEthereum.initialize(CROSS_CHAIN_ID, cross, threshold, REWARD, address(0), address(0));

            bridgeEthereum.setValidators(VALIDATORS);
        }

        // add token to bridge (ethereum chain)
        {
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address(NATIVE_TOKEN), address(weth));
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address(testTokenEthereum), address(testTokenCross));

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
        ok = bridgeEthereum.finalizeBridge(CROSS_CHAIN_ID, index, IERC20(token), to, value, NULLDATA, v, r, s);
    }

    function ethereumCalcFee(IERC20, uint) public pure returns (uint value, uint gas, uint ex) {
        return (value, 0, 0); // no fee
            // vm.selectFork(ethereumForkID);
            // if (address(bridgeFeeStationEthereum) == address(0)) return (value, 0, 0);
            // bool ok;
            // (value, gas, ex, ok) = bridgeFeeStationEthereum.calculateMax(token, totalValue);
            // assertTrue(ok);
    }
}
