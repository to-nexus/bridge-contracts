// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeEthereum} from "../../src/BridgeEthereum.sol";
import {BridgeTokenInfo, IBridgeTokenInfo} from "../../src/BridgeTokenInfo.sol";
import {TestToken} from "../token/TestToken.sol";
import {CrossChainTest} from "./CrossChain.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract EthereumChainTest is CrossChainTest {
    bool public bridgeRevertEthereum = false;
    bool public finalizeRevertEthereum = false;

    uint public nextIndexEthereum;
    BridgeEthereum public bridgeEthereum;
    BridgeTokenInfo public bridgeTokenInfoEthereum;

    function setUp() public virtual override {
        super.setUp();
        nextIndexEthereum = 1;
        vm.selectFork(ethereumChainID);
        vm.startPrank(OWNER);
        // bridge setup
        {
            BridgeEthereum bridgeEthereumImpl = new BridgeEthereum();
            ERC1967Proxy bridgeEthereumProxy = new ERC1967Proxy(address(bridgeEthereumImpl), bytes(""));
            bridgeEthereum = BridgeEthereum(payable(address(bridgeEthereumProxy)));
            bridgeEthereum.initialize(IERC20(address(cross)), threshold, REWARD, address(0));
            address[] memory validators = new address[](VALIDATORS.length);
            for (uint i = 0; i < VALIDATORS.length; i++) {
                validators[i] = VALIDATORS[i];
            }
            bridgeEthereum.setValidators(validators);
        }

        // add token to bridge (ethereum chain)
        {
            bridgeEthereum.addToken(coin, weth);
            bridgeEthereum.addToken(testTokenEthereum, IERC20(testTokenCross));

            uint initialSupply = 1000000000 * 1e18;
            TestToken(address(cross)).mint(OWNER, initialSupply);
            TestToken(address(testTokenEthereum)).mint(OWNER, initialSupply);
        }
        vm.stopPrank();
    }

    function ethereumIncrementIndex() public {
        nextIndexEthereum++;
    }

    // ----- Functions -----
    function ethereumBridge(address token, uint value, uint gas, uint service) public returns (uint index, bool ok) {
        vm.selectFork(ethereumChainID);
        // bridge
        index = nextIndexEthereum;
        vm.prank(USER);

        if (bridgeRevertEthereum) {
            bridgeRevertEthereum = false;
            vm.expectRevert();
        }

        if (token == address(coin)) {
            assertTrue(USER.balance >= value + gas + service);
            ok = bridgeEthereum.bridge{value: value + gas + service}(IERC20(token), value, gas, service, NULLDATA);
        } else {
            ok = bridgeEthereum.bridge(IERC20(token), value, gas, service, NULLDATA);
        }
    }

    function ethereumFinalize(uint index, address token, uint value, uint sigCount) public returns (bool ok) {
        vm.selectFork(ethereumChainID);
        if (sigCount > 5) sigCount = 5;

        // create finalize validator signature
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, USER, value, NULLDATA));
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
        ok = bridgeEthereum.finalize(index, IERC20(token), USER, value, NULLDATA, v, r, s);
    }

    function ethereumCalcFee(IERC20 token, uint totalValue) public returns (uint value, uint gas, uint ex) {
        vm.selectFork(ethereumChainID);
        if (address(bridgeTokenInfoEthereum) == address(0)) return (value, 0, 0);
        bool ok;
        (value, gas, ex, ok) = bridgeTokenInfoEthereum.calculateMax(token, totalValue);
        assertTrue(ok);
    }
}
