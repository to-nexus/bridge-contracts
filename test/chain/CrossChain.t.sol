// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier} from "../../src/BridgeVerifier.sol";
import {CrossBridge} from "../../src/CrossBridge.sol";

import {IPriceFeed, PriceFeed} from "../../src/PriceFeed.sol";
import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";

import {CrossMintableERC20} from "../../src/token/CrossMintableERC20.sol";
import {CrossMintableERC20Code} from "../../src/token/CrossMintableERC20Code.sol";
import {ICrossMintableERC20Code} from "../../src/token/ICrossMintableERC20Code.sol";
import {SettingTest} from "./Setting.t.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {console} from "forge-std/Test.sol";

contract CrossChainTest is SettingTest {
    bool internal bridgeRevertCross = false;
    bool internal finalizeRevertCross = false;

    uint internal nextIndexCross;
    CrossBridge internal bridgeCross;
    BridgeVerifier internal bridgeVerifierCross;
    PriceFeed internal priceFeedCross;
    ICrossMintableERC20Code internal crossMintableERC20Code;

    function setUp() public virtual override {
        super.setUp();
        nextIndexCross = 1;

        vm.selectFork(crossForkID);
        vm.startPrank(CrossOWNER);

        // bridge setup
        {
            // bridge
            CrossBridge bridgeCrossImpl = new CrossBridge();
            ERC1967Proxy bridgeCrossProxy = new ERC1967Proxy(address(bridgeCrossImpl), bytes(""));
            bridgeCross = CrossBridge(payable(address(bridgeCrossProxy)));
            bridgeCross.initialize(CrossOWNER, threshold, REWARD);

            crossMintableERC20Code = ICrossMintableERC20Code(address(new CrossMintableERC20Code(address(bridgeCross))));
            bridgeCross.setCrossMintableERC20Code(crossMintableERC20Code);

            bridgeCross.grantRole(ADMIN_ROLE, CrossOWNER); // for test
            bridgeCross.grantRole(OPERATOR_ROLE, CrossOWNER); // for test
            bridgeCross.grantRole(UPDATOR_ROLE, CrossOWNER); // for test
            bridgeCross.registerToken(ETHEREUM_CHAIN_ID, false, address(NATIVE_TOKEN), address(cross));
            bridgeCross.grantRoleBatch(VALIDATOR_ROLE, VALIDATORS);

            vm.label(address(bridgeCross), "CrossBridge");
        }

        // add token to bridge (cross chain)
        {
            // test token
            address ttAddress = bridgeCross.createToken(ETHEREUM_CHAIN_ID, address(testTokenEthereum), "TT", 18);
            testTokenCross = IERC20(ttAddress);

            // weth
            address wethAddress = bridgeCross.createToken(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN), "ETH", 18);
            weth = IERC20(wethAddress);
        }

        // fee table
        {
            address priceFeedCrossImpl = address(new PriceFeed());
            ERC1967Proxy priceFeedCrossProxy = new ERC1967Proxy(priceFeedCrossImpl, bytes(""));
            priceFeedCross = PriceFeed(address(priceFeedCrossProxy));
            priceFeedCross.initialize(CrossOWNER, DOLLAR_DECIMALS);

            {
                // token price update
                address[] memory tokens = new address[](3);
                uint[] memory prices = new uint[](3);
                uint[] memory pricesAt = new uint[](3);

                tokens[0] = address(weth);
                prices[0] = 10 * (10 ** 6);
                pricesAt[0] = 0;

                tokens[1] = address(testTokenCross);
                prices[1] = 1 * (10 ** 6);
                pricesAt[1] = 0;

                tokens[2] = address(NATIVE_TOKEN);
                prices[2] = (10 ** DOLLAR_DECIMALS) / 10;
                pricesAt[2] = 0;

                priceFeedCross.updatePrice(tokens, prices, pricesAt);
            }

            {
                // native token price update
                uint[] memory chainIDs = new uint[](1);
                uint[] memory prices = new uint[](1);
                uint[] memory pricesAt = new uint[](1);

                chainIDs[0] = ETHEREUM_CHAIN_ID;
                prices[0] = 100_000;
                pricesAt[0] = 0;

                priceFeedCross.updateNativeTokenPrice(chainIDs, prices, pricesAt);
            }

            priceFeedCross.grantRoleBatch(UPDATOR_ROLE, VALIDATORS);

            bridgeVerifierCross = new BridgeVerifier(
                CrossOWNER, address(bridgeCross), address(priceFeedCross), 200000, 10000, 10, 10_000, 0, 0, 2 hours
            );
            bridgeVerifierCross.grantRole(UPDATOR_ROLE, CrossOWNER);
            bridgeVerifierCross.updateGasPrice(ETHEREUM_CHAIN_ID, 1 gwei);
        }

        bridgeCross.setBridgeVerifier(bridgeVerifierCross);

        vm.deal(address(bridgeCross), INITIAL_SUPPLY);
        vm.stopPrank();
    }

    function crossIncrementIndex() public {
        nextIndexCross++;
    }

    // ----- Functions -----
    function crossBridge(address token, address from, address to, uint value, uint gas, uint service)
        public
        returns (uint index, bool ok)
    {
        vm.selectFork(crossForkID);

        // bridge
        index = nextIndexCross;
        vm.prank(from);

        if (bridgeRevertCross) {
            bridgeRevertCross = false;
            vm.expectRevert();
        }

        if (token == address(NATIVE_TOKEN)) {
            ok = bridgeCross.bridgeToken{value: value + gas + service}(
                ETHEREUM_CHAIN_ID, IERC20(token), to, value, gas, service, NULLDATA
            );
        } else {
            ok = bridgeCross.bridgeToken(ETHEREUM_CHAIN_ID, IERC20(token), to, value, gas, service, NULLDATA);
        }
    }

    function crossFinalize(uint index, address token, address to, uint value, uint sigCount) public returns (bool ok) {
        vm.selectFork(crossForkID);
        if (sigCount > threshold) sigCount = threshold;

        // create finalize validator signature
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, ETHEREUM_CHAIN_ID, index, token, to, value, NULLDATA));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), h);

        uint8[] memory v = new uint8[](sigCount);
        bytes32[] memory r = new bytes32[](sigCount);
        bytes32[] memory s = new bytes32[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (v[i], r[i], s[i]) = vm.sign(VALIDATOR_PKs[i], hash);
        }

        // finalize
        vm.prank(VALIDATOR1);
        if (finalizeRevertCross) {
            finalizeRevertCross = false;
            vm.expectRevert();
        }
        ok = bridgeCross.finalizeBridge(
            IBridgeRegistry.FinalizeArguments({
                fromChainID: ETHEREUM_CHAIN_ID,
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

    function crossCalcFee(IERC20 token, uint totalValue) public returns (uint value, uint gas, uint ex) {
        vm.selectFork(crossForkID);
        if (address(bridgeVerifierCross) == address(0)) return (totalValue, 0, 0);

        bool ok;
        (ok, value, gas, ex) = estimateMaxValue(bridgeVerifierCross, ETHEREUM_CHAIN_ID, token, totalValue);
        assertTrue(ok);
    }

    function crossGetTokenFee(IERC20 token) public returns (uint minimum, uint gasFee, uint exFeeRate) {
        vm.selectFork(crossForkID);
        if (address(bridgeVerifierCross) == address(0)) return (0, 0, 0);
        (minimum, gasFee, exFeeRate) = bridgeVerifierCross.getTokenConfig(ETHEREUM_CHAIN_ID, token);
    }
}
