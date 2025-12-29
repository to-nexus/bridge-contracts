// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BSCBridge} from "../../src/BSCBridge.sol";
import {BridgeVerifier, IBridgeVerifier} from "../../src/BridgeVerifier.sol";

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {TestToken} from "../token/TestToken.sol";
import {CrossChainTest} from "./CrossChain.t.sol";

import {BSCBridgeV2} from "../../src/BSCBridgeV2.sol";
import {IPriceFeed, PriceFeed} from "../../src/PriceFeed.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract BSCTest is CrossChainTest {
    bool internal bridgeRevertBSC = false;
    bool internal finalizeRevertBSC = false;

    PriceFeed internal priceFeedBSC;
    uint internal nextIndexBSC;
    BSCBridge internal bridgeBSC;
    BridgeVerifier internal bridgeVerifierBSC;

    function setUp() public virtual override {
        super.setUp();
        nextIndexBSC = 1;
        vm.selectFork(bscForkID);
        vm.startPrank(OWNER);

        // bridge setup
        {
            BSCBridge bridgeBSCImpl = new BSCBridge();
            ERC1967Proxy bridgeBSCProxy = new ERC1967Proxy(address(bridgeBSCImpl), bytes(""));
            bridgeBSC = BSCBridge(payable(address(bridgeBSCProxy)));
            bridgeBSC.initializeBSCBridge(
                OWNER, REWARD, threshold, CROSS_CHAIN_ID, address(cross), CROSS_FOUNDATION_INITIAL_SUPPLY
            );

            bridgeBSC.grantRole(ADMIN_ROLE, OWNER); // for test
            bridgeBSC.grantRole(EDITOR_ROLE, OWNER); // for test
            bridgeBSC.grantRole(OPERATOR_ROLE, OWNER); // for test
            bridgeBSC.grantRole(PRICER_ROLE, OWNER); // for test
            // bridgeBSC.registerToken(CROSS_CHAIN_ID, true, address(cross), address(NATIVE_TOKEN)); // already registered
            bytes32[] memory roles = new bytes32[](5);
            for (uint i = 0; i < 5; i++) {
                roles[i] = VALIDATOR_ROLE;
            }
            bridgeBSC.grantRoleBatch(roles, VALIDATORS);
        }

        {
            // fee table

            address priceFeedBSCImpl = address(new PriceFeed());
            ERC1967Proxy priceFeedBSCProxy = new ERC1967Proxy(priceFeedBSCImpl, bytes(""));
            priceFeedBSC = PriceFeed(address(priceFeedBSCProxy));
            priceFeedBSC.initialize(OWNER, DOLLAR_DECIMALS);
            priceFeedBSC.grantRole(PRICER_ROLE, OWNER); // for test

            {
                // token price update
                address[] memory tokens = new address[](3);
                uint[] memory prices = new uint[](3);
                uint[] memory pricesAt = new uint[](3);

                tokens[0] = address(NATIVE_TOKEN);
                prices[0] = 1;
                pricesAt[0] = 0;

                tokens[1] = address(testTokenBSC);
                prices[1] = 1;
                pricesAt[1] = 0;

                tokens[2] = address(cross);
                prices[2] = 1;
                pricesAt[2] = 0;

                priceFeedBSC.updatePrice(tokens, prices, pricesAt);
            }

            bridgeVerifierBSC =
                new BridgeVerifier(OWNER, address(bridgeBSC), address(priceFeedBSC), 200000, 0, 0, 0, 0, 0, 2 hours);
            bridgeVerifierBSC.grantRole(PRICER_ROLE, OWNER);
            bridgeVerifierBSC.grantRole(ADMIN_ROLE, OWNER);
            bridgeVerifierBSC.grantRole(EDITOR_ROLE, OWNER);

            bridgeBSC.setBridgeVerifier(bridgeVerifierBSC);

            bytes32[] memory roles = new bytes32[](5);
            for (uint i = 0; i < 5; i++) {
                roles[i] = VALIDATOR_ROLE;
            }
            priceFeedBSC.grantRoleBatch(roles, VALIDATORS);
        }

        // add token to bridge (bsc)
        {
            bridgeBSC.registerToken(CROSS_CHAIN_ID, true, address(NATIVE_TOKEN), address(weth));
            bridgeBSC.registerToken(CROSS_CHAIN_ID, true, address(testTokenBSC), address(testTokenCross));

            // Mint tokens excluding the pre-allocated supply amount (CROSS_FOUNDATION_INITIAL_SUPPLY)
            TestToken(address(cross)).mint(OWNER, INITIAL_SUPPLY - CROSS_FOUNDATION_INITIAL_SUPPLY);
            TestToken(address(testTokenBSC)).mint(OWNER, INITIAL_SUPPLY);
        }
        vm.stopPrank();
    }

    function bscIncrementIndex() public {
        nextIndexBSC++;
    }

    // ----- Functions -----
    function bscBridge(address token, address from, address to, uint value, uint gas, uint service)
        public
        returns (uint index, bool ok)
    {
        return bscBridge(token, from, to, value, gas, service, NULLDATA);
    }

    function bscBridge(
        address token,
        address from,
        address to,
        uint value,
        uint gas,
        uint service,
        bytes memory extraData
    ) public returns (uint index, bool ok) {
        vm.selectFork(bscForkID);
        // bridge
        index = nextIndexBSC;
        vm.prank(from);

        if (bridgeRevertBSC) {
            bridgeRevertBSC = false;
            vm.expectRevert();
        }

        if (token == address(NATIVE_TOKEN)) {
            assertTrue(from.balance >= value + gas + service);
            ok = bridgeBSC.bridgeToken{value: value + gas + service}(
                CROSS_CHAIN_ID, IERC20(token), to, value, gas, service, extraData
            );
        } else {
            ok = bridgeBSC.bridgeToken(CROSS_CHAIN_ID, IERC20(token), to, value, gas, service, extraData);
        }
    }

    function bscFinalize(uint index, address token, address to, uint value, uint sigCount) public returns (bool ok) {
        return bscFinalize(index, token, to, value, sigCount, NULLDATA);
    }

    function bscFinalize(uint index, address token, address to, uint value, uint sigCount, bytes memory extraData)
        public
        returns (bool ok)
    {
        vm.selectFork(bscForkID);
        if (sigCount > threshold) sigCount = threshold;

        // create finalize validator signature
        bytes32 h =
            keccak256(abi.encode(FINALIZE_TYPEHASH, CROSS_CHAIN_ID, index, token, to, value, keccak256(extraData)));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeBSC.domainSeparator(), h);

        uint8[] memory v = new uint8[](sigCount);
        bytes32[] memory r = new bytes32[](sigCount);
        bytes32[] memory s = new bytes32[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], hash);
        }

        // finalize
        vm.prank(VALIDATOR1);
        if (finalizeRevertBSC) {
            finalizeRevertBSC = false;
            vm.expectRevert();
        }
        ok = bridgeBSC.finalizeBridge(
            IBridgeRegistry.FinalizeArguments({
                fromChainID: CROSS_CHAIN_ID,
                index: index,
                toToken: IERC20(token),
                to: to,
                value: value,
                extraData: extraData
            }),
            v,
            r,
            s
        );
    }

    function bscBurnCrossToDeadWallet(address from, address to, uint value, bool alreadyTransferred)
        public
        returns (uint index, bool ok)
    {
        vm.selectFork(bscForkID);
        // burn cross to dead wallet
        index = nextIndexBSC;
        vm.prank(from);

        if (bridgeRevertBSC) {
            bridgeRevertBSC = false;
            vm.expectRevert();
        }

        ok = BSCBridgeV2(payable(address(bridgeBSC))).burnCrossToDeadWallet(to, value, alreadyTransferred);
    }

    function bscCalcFee(IERC20 token, uint totalValue) public returns (uint value, uint gas, uint ex) {
        vm.selectFork(bscForkID);
        if (address(bridgeVerifierBSC) == address(0)) return (totalValue, 0, 0);

        bool ok;
        (ok, value, gas, ex) = estimateMaxValue(bridgeVerifierBSC, BSC_CHAIN_ID, token, totalValue);
        assertTrue(ok);
    }

    function bscGetTokenFee(IERC20 token) public returns (uint minimum, uint gasFee, uint exFeeRate) {
        vm.selectFork(bscForkID);
        if (address(bridgeVerifierBSC) == address(0)) return (0, 0, 0);
        (minimum, gasFee, exFeeRate) = bridgeVerifierBSC.getTokenConfig(BSC_CHAIN_ID, token);
    }
}
