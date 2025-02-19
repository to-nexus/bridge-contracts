// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeCross} from "../../src/BridgeCross.sol";
import {BridgeTokenInfo, IBridgeTokenInfo} from "../../src/BridgeTokenInfo.sol";
import {CrossMintableERC20, CrossMintableERC20Code} from "../../src/CrossMintableERC20.sol";
import {ConstTest} from "./Const.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract CrossChainTest is ConstTest {
    bool public bridgeRevertCross = false;
    bool public finalizeRevertCross = false;

    uint public nextIndexCross;
    BridgeCross public bridgeCross;
    BridgeTokenInfo public bridgeTokenInfoCross;
    CrossMintableERC20Code public crossMintableERC20Code;

    function setUp() public virtual override {
        super.setUp();
        nextIndexCross = 1;

        vm.selectFork(crossChainID);
        vm.startPrank(OWNER);
        // cross token
        {
            crossMintableERC20Code = new CrossMintableERC20Code();
        }

        // bridge setup
        {
            // bridge
            BridgeCross bridgeCrossImpl = new BridgeCross();
            ERC1967Proxy bridgeCrossProxy = new ERC1967Proxy(address(bridgeCrossImpl), bytes(""));
            bridgeCross = BridgeCross(payable(address(bridgeCrossProxy)));
            bridgeCross.initialize(address(crossMintableERC20Code), REWARD, address(0));

            address[] memory validators = new address[](VALIDATORS.length);
            for (uint i = 0; i < VALIDATORS.length; i++) {
                validators[i] = VALIDATORS[i];
            }
            bridgeCross.setValidators(validators);
        }

        // add token to bridge (cross chain)
        {
            string memory name = string(abi.encodePacked("Cross Bridge ", "TT"));
            bytes32 salt = keccak256(abi.encodePacked(testTokenEthereum));
            bytes memory bytecode =
                abi.encodePacked(type(CrossMintableERC20).creationCode, abi.encode(name, "TT", uint8(18)));

            bridgeCross.addTokenDeploy(IERC20(testTokenEthereum), "TT", 18);
            address ttAddress = Create2.computeAddress(salt, keccak256(bytecode), address(bridgeCross));
            testTokenCross = IERC20(ttAddress);
            bridgeCross.addToken(xcross, IERC20(address(cross)));
        }

        // fee table
        {
            bridgeTokenInfoCross = new BridgeTokenInfo();
            IBridgeTokenInfo.TokenInfo[] memory TokenInfoList = new IBridgeTokenInfo.TokenInfo[](2);
            TokenInfoList[0] = IBridgeTokenInfo.TokenInfo(address(xcross), 0, 100 * 1e18, 10);
            TokenInfoList[1] = IBridgeTokenInfo.TokenInfo(address(testTokenCross), 0, 100 * 1e18, 10);
            bridgeTokenInfoCross.addTokenInfoMany(TokenInfoList);
        }

        bridgeCross.setTokenInfo(bridgeTokenInfoCross);

        vm.deal(address(bridgeCross), 1000000000 * 1e18);
        vm.stopPrank();
    }

    function crossIncrementIndex() public {
        nextIndexCross++;
    }

    // ----- Functions -----
    function crossBridge(address token, uint value, uint gas, uint service) public returns (uint index, bool ok) {
        vm.selectFork(crossChainID);

        // bridge
        index = nextIndexCross;
        vm.prank(USER);

        if (bridgeRevertCross) {
            bridgeRevertCross = false;
            vm.expectRevert();
        }

        if (token == address(xcross)) {
            assertTrue(USER.balance >= value + gas + service);
            ok = bridgeCross.bridge{value: value + gas + service}(IERC20(token), value, gas, service, NULLDATA);
        } else {
            ok = bridgeCross.bridge(IERC20(token), value, gas, service, NULLDATA);
        }
    }

    function crossFinalize(uint index, address token, uint value, uint sigCount) public returns (bool ok) {
        vm.selectFork(crossChainID);
        if (sigCount > 5) sigCount = 5;

        // create finalize validator signature
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, USER, value, NULLDATA));
        bytes32 hash = MessageHashUtils.toTypedDataHash(bridgeCross.domainSeparator(), h);

        bytes[] memory sigs = new bytes[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR_PKs[i], hash);
            sigs[i] = abi.encodePacked(r, s, v);
        }

        // finalize
        vm.prank(VALIDATOR1);
        if (finalizeRevertCross) {
            finalizeRevertCross = false;
            vm.expectRevert();
        }
        ok = bridgeCross.finalize(index, IERC20(token), USER, value, NULLDATA, sigs);
    }

    function crossCalcFee(IERC20 token, uint value) public returns (uint bridgeValue, uint gas, uint service) {
        vm.selectFork(crossChainID);
        BridgeTokenInfo.TokenInfo memory tokenInfo = bridgeTokenInfoCross.getTokenInfo(address(token));
        assertTrue(tokenInfo.gasFee < value);

        gas = tokenInfo.gasFee;
        uint denom = bridgeTokenInfoCross.denominator();
        uint serviceRate = tokenInfo.serviceFee;

        uint v = value - gas;

        bridgeValue = (v * denom / (denom + serviceRate));
        bridgeValue = (bridgeValue / exrate) * exrate;
        service = (bridgeValue * serviceRate / denom);
    }
}
