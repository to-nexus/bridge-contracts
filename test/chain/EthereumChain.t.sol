// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier, IBridgeVerifier} from "../../src/BridgeVerifier.sol";
import {EthereumBridge} from "../../src/EthereumBridge.sol";

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {TestToken} from "../token/TestToken.sol";
import {CrossChainTest} from "./CrossChain.t.sol";

import {IPriceFeed, PriceFeed} from "../../src/PriceFeed.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract EthereumChainTest is CrossChainTest {
    bool internal bridgeRevertEthereum = false;
    bool internal finalizeRevertEthereum = false;

    PriceFeed internal priceFeedEthereum;
    uint internal nextIndexEthereum;
    EthereumBridge internal bridgeEthereum;
    BridgeVerifier internal bridgeVerifierEthereum;

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
            bridgeEthereum.initializeEthereumBridge(
                OWNER, REWARD, threshold, CROSS_CHAIN_ID, address(cross), CROSS_FOUNDATION_INITIAL_SUPPLY
            );

            bridgeEthereum.grantRole(ADMIN_ROLE, OWNER); // for test
            bridgeEthereum.grantRole(EDITOR_ROLE, OWNER); // for test
            bridgeEthereum.grantRole(OPERATOR_ROLE, OWNER); // for test
            bridgeEthereum.grantRole(PRICER_ROLE, OWNER); // for test
            // bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, address(cross), address(NATIVE_TOKEN)); // already registered
            bytes32[] memory roles = new bytes32[](5);
            for (uint i = 0; i < 5; i++) {
                roles[i] = VALIDATOR_ROLE;
            }
            bridgeEthereum.grantRoleBatch(roles, VALIDATORS);
        }

        {
            // fee table

            address priceFeedEthereumImpl = address(new PriceFeed());
            ERC1967Proxy priceFeedEthereumProxy = new ERC1967Proxy(priceFeedEthereumImpl, bytes(""));
            priceFeedEthereum = PriceFeed(address(priceFeedEthereumProxy));
            priceFeedEthereum.initialize(OWNER, DOLLAR_DECIMALS);
            priceFeedEthereum.grantRole(PRICER_ROLE, OWNER); // for test

            {
                // token price update
                address[] memory tokens = new address[](3);
                uint[] memory prices = new uint[](3);
                uint[] memory pricesAt = new uint[](3);

                tokens[0] = address(NATIVE_TOKEN);
                prices[0] = 1;
                pricesAt[0] = 0;

                tokens[1] = address(testTokenEthereum);
                prices[1] = 1;
                pricesAt[1] = 0;

                tokens[2] = address(cross);
                prices[2] = 1;
                pricesAt[2] = 0;

                priceFeedEthereum.updatePrice(tokens, prices, pricesAt);
            }

            bridgeVerifierEthereum = new BridgeVerifier(
                OWNER, address(bridgeEthereum), address(priceFeedEthereum), 200000, 0, 0, 0, 0, 0, 2 hours
            );
            bridgeVerifierEthereum.grantRole(PRICER_ROLE, OWNER);
            bridgeVerifierEthereum.grantRole(ADMIN_ROLE, OWNER);
            bridgeVerifierEthereum.grantRole(EDITOR_ROLE, OWNER);

            bridgeEthereum.setBridgeVerifier(bridgeVerifierEthereum);

            bytes32[] memory roles = new bytes32[](5);
            for (uint i = 0; i < 5; i++) {
                roles[i] = VALIDATOR_ROLE;
            }
            priceFeedEthereum.grantRoleBatch(roles, VALIDATORS);
        }

        // add token to bridge (ethereum chain)
        {
            bridgeEthereum.registerToken(CROSS_CHAIN_ID, true, true, address(NATIVE_TOKEN), address(weth));
            bridgeEthereum.registerToken(
                CROSS_CHAIN_ID, true, true, address(testTokenEthereum), address(testTokenCross)
            );

            // Mint tokens excluding the pre-allocated supply amount (CROSS_FOUNDATION_INITIAL_SUPPLY)
            TestToken(address(cross)).mint(OWNER, INITIAL_SUPPLY - CROSS_FOUNDATION_INITIAL_SUPPLY);
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
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, CROSS_CHAIN_ID, index, token, to, value, NULLDATA));
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

    function ethereumCalcFee(IERC20 token, uint totalValue) public returns (uint value, uint gas, uint ex) {
        vm.selectFork(ethereumForkID);
        if (address(bridgeVerifierEthereum) == address(0)) return (totalValue, 0, 0);

        bool ok;
        (ok, value, gas, ex) = estimateMaxValue(bridgeVerifierEthereum, ETHEREUM_CHAIN_ID, token, totalValue);
        assertTrue(ok);
    }

    function ethereumGetTokenFee(IERC20 token) public returns (uint minimum, uint gasFee, uint exFeeRate) {
        vm.selectFork(ethereumForkID);
        if (address(bridgeVerifierEthereum) == address(0)) return (0, 0, 0);
        (minimum, gasFee, exFeeRate) = bridgeVerifierEthereum.getTokenConfig(ETHEREUM_CHAIN_ID, token);
    }
}
