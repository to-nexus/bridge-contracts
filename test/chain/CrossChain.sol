// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeManager} from "../../src/BridgeManager.sol";
import {CrossBridge} from "../../src/CrossBridge.sol";

import {IPriceFeed, PriceFeed} from "../../src/PriceFeed.sol";
import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";

import {CrossMintableERC20} from "../../src/token/CrossMintableERC20.sol";
import {CrossMintableERC20Code} from "../../src/token/CrossMintableERC20Code.sol";
import {ICrossMintableERC20Code} from "../../src/token/ICrossMintableERC20Code.sol";
import {SettingTest} from "./Setting.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {console} from "forge-std/Test.sol";

contract CrossChainTest is SettingTest {
    bool internal bridgeRevertCross = false;
    bool internal finalizeRevertCross = false;

    uint internal nextIndexCross;
    CrossBridge internal bridgeCross;
    BridgeManager internal bridgeManagerCross;
    PriceFeed internal priceFeed;
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
            bridgeCross.registerToken(ETHEREUM_CHAIN_ID, false, address(NATIVE_TOKEN), address(cross), int(EX_RATE));
            bridgeCross.grantRoleBatch(VALIDATOR_ROLE, VALIDATORS);

            vm.label(address(bridgeCross), "CrossBridge");
        }

        // add token to bridge (cross chain)
        {
            // test token
            address ttAddress = bridgeCross.createToken(ETHEREUM_CHAIN_ID, address(testTokenEthereum), 0, "TT", 18);
            testTokenCross = IERC20(ttAddress);

            // weth
            address wethAddress = bridgeCross.createToken(ETHEREUM_CHAIN_ID, address(NATIVE_TOKEN), 0, "ETH", 18);
            weth = IERC20(wethAddress);
        }

        // fee table
        {
            address priceFeedImpl = address(new PriceFeed());
            ERC1967Proxy priceFeedProxy = new ERC1967Proxy(priceFeedImpl, bytes(""));
            priceFeed = PriceFeed(address(priceFeedProxy));
            priceFeed.initialize(CrossOWNER, uint8(6), 100000);
            priceFeed.grantRoleBatch(UPDATOR_ROLE, VALIDATORS);

            bridgeManagerCross =
                new BridgeManager(CrossOWNER, address(bridgeCross), 200000, 10000, 10, 10_000, 0, 0, 2 hours);
            bridgeManagerCross.grantRole(UPDATOR_ROLE, CrossOWNER);
            bridgeManagerCross.setPriceFeed(IPriceFeed(address(priceFeed)));
        }

        bridgeCross.setBridgeManager(bridgeManagerCross);

        vm.deal(address(bridgeCross), INITIAL_SUPPLY * EX_RATE);
        vm.stopPrank();

        address[] memory tokens = new address[](2);
        uint[] memory prices = new uint[](2);
        uint[] memory pricesAt = new uint[](2);

        tokens[0] = address(weth);
        prices[0] = 10 * (10 ** 6);
        pricesAt[0] = type(uint).max;

        tokens[1] = address(testTokenCross);
        prices[1] = 1 * (10 ** 6);
        pricesAt[1] = type(uint).max;

        vm.prank(VALIDATOR1);
        priceFeed.updatePrice(tokens, prices, pricesAt);
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
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, to, value, NULLDATA));
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
        if (address(bridgeManagerCross) == address(0)) return (totalValue, 0, 0);

        bool ok;
        (ok, value, gas, ex) = estimateMaxValue(bridgeManagerCross, ETHEREUM_CHAIN_ID, token, totalValue);
        assertTrue(ok);
    }

    function crossGetTokenFee(IERC20 token) public returns (uint minimum, uint gasFee, uint exFeeRate) {
        vm.selectFork(crossForkID);
        if (address(bridgeManagerCross) == address(0)) return (0, 0, 0);
        (minimum, gasFee, exFeeRate) = bridgeManagerCross.getTokenConfig(ETHEREUM_CHAIN_ID, token);
    }
}
