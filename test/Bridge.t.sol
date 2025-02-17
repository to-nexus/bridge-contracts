// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeCross} from "../src/BridgeCross.sol";
import {BridgeEthereum} from "../src/BridgeEthereum.sol";

import {BridgeFeeManager, IBridgeFeeManager} from "../src/BridgeFeeManager.sol";
import {CrossMintableERC20, CrossMintableERC20Code} from "../src/CrossMintableERC20.sol";
import {BridgeStandard} from "../src/abstract/BridgeStandard.sol";
import {TestToken} from "./token/TestToken.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Create2} from "@openzeppelin/contracts/utils/Create2.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {Test, console} from "forge-std/Test.sol";

contract BridgeTest is Test {
    event BridgeInitiated(uint indexed index, address indexed from, address indexed to, uint value);
    event BridgeFinalized(uint indexed index, address indexed to, uint value, bytes32 indexed otherChainHash);

    bool public bridgeRevert = false;
    bool public finalizeRevert = false;
    uint public exrate = 100;

    bytes32 public constant FINALIZE_TYPEHASH =
        keccak256("Finalize(uint256 index,address token,address to,uint256 value,bytes[] extraData)");

    uint internal constant OWNER_PK = uint(bytes32("owner"));
    uint internal constant USER_PK = uint(bytes32("user"));
    uint internal constant REWARD_PK = uint(bytes32("reward"));
    uint internal constant VALIDATOR_PK1 = uint(bytes32("validator1"));
    uint internal constant VALIDATOR_PK2 = uint(bytes32("validator2"));
    uint internal constant VALIDATOR_PK3 = uint(bytes32("validator3"));
    uint internal constant VALIDATOR_PK4 = uint(bytes32("validator4"));
    uint internal constant VALIDATOR_PK5 = uint(bytes32("validator5"));

    address public OWNER;
    address public USER;
    address public REWARD;
    address public VALIDATOR1;
    address public VALIDATOR2;
    address public VALIDATOR3;
    address public VALIDATOR4;
    address public VALIDATOR5;

    address[5] public VALIDATORS;
    uint[5] public VALIDATOR_PKs = [VALIDATOR_PK1, VALIDATOR_PK2, VALIDATOR_PK3, VALIDATOR_PK4, VALIDATOR_PK5];

    TestToken public cross;
    IERC20 public xcross = IERC20(address(1));
    IERC20 public testTokenAtEthereum;
    IERC20 public testTokenAtCross;

    bytes[] NULLDATA;

    struct BridgeInfo {
        BridgeStandard b;
        uint nextIndex;
        address otherBridge;
    }

    mapping(address => BridgeInfo) public bridges;

    BridgeInfo public bridgeEthereum;
    BridgeInfo public bridgeCross;
    BridgeFeeManager public bridgeFeeManager;
    CrossMintableERC20Code public crossMintableERC20Code;

    function setUp() public {
        OWNER = vm.addr(OWNER_PK);
        USER = vm.addr(USER_PK);
        REWARD = vm.addr(REWARD_PK);

        VALIDATOR1 = vm.addr(VALIDATOR_PK1);
        VALIDATOR2 = vm.addr(VALIDATOR_PK2);
        VALIDATOR3 = vm.addr(VALIDATOR_PK3);
        VALIDATOR4 = vm.addr(VALIDATOR_PK4);
        VALIDATOR5 = vm.addr(VALIDATOR_PK5);

        VALIDATORS = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4, VALIDATOR5];

        vm.label(OWNER, "owner");
        vm.label(USER, "user");
        vm.label(REWARD, "reward");
        vm.label(VALIDATOR1, "validator1");
        vm.label(VALIDATOR2, "validator2");
        vm.label(VALIDATOR3, "validator3");
        vm.label(VALIDATOR4, "validator4");
        vm.label(VALIDATOR5, "validator5");

        vm.startPrank(OWNER);

        address[] memory validators = new address[](5);
        for (uint i = 0; i < 5; i++) {
            validators[i] = VALIDATORS[i];
        }

        // token setup
        {
            cross = new TestToken("Cross", "CROSS", 18);
            TestToken tt = new TestToken("Test Token", "TT", 18);
            testTokenAtEthereum = IERC20(address(tt));
            crossMintableERC20Code = new CrossMintableERC20Code();
        }

        // bridge(cross) fee table setup
        {
            bridgeFeeManager = new BridgeFeeManager();
            IBridgeFeeManager.FeeInfo[] memory FeeInfoList = new IBridgeFeeManager.FeeInfo[](2);
            FeeInfoList[0] = IBridgeFeeManager.FeeInfo(address(cross), 100 * 1e18, 10);
            FeeInfoList[1] = IBridgeFeeManager.FeeInfo(address(testTokenAtEthereum), 100 * 1e18, 10);

            bridgeFeeManager.addFeeInfoMany(FeeInfoList);
        }

        // bridge setup
        {
            BridgeEthereum bridgeEthereumImpl = new BridgeEthereum();
            BridgeCross bridgeCrossImpl = new BridgeCross();
            ERC1967Proxy bridgeEthereumProxy = new ERC1967Proxy(address(bridgeEthereumImpl), bytes(""));
            ERC1967Proxy bridgeCrossProxy = new ERC1967Proxy(address(bridgeCrossImpl), bytes(""));

            bridgeEthereum = BridgeInfo(BridgeEthereum(payable(address(bridgeEthereumProxy))), 1, address(0));
            bridgeCross = BridgeInfo(BridgeCross(payable(address(bridgeCrossProxy))), 1, address(0));

            // initialize
            BridgeEthereum(address(bridgeEthereum.b)).initialize(IERC20(address(cross)), REWARD, address(0));
            BridgeCross(payable(address(bridgeCross.b))).initialize(
                address(crossMintableERC20Code), REWARD, address(bridgeFeeManager)
            );

            // set validator
            bridgeEthereum.b.setValidators(validators);
            bridgeCross.b.setValidators(validators);

            bridges[address(bridgeEthereum.b)] = bridgeEthereum;
            bridges[address(bridgeCross.b)] = bridgeCross;

            bridgeEthereum.otherBridge = address(bridgeCross.b);
            bridgeCross.otherBridge = address(bridgeEthereum.b);
        }

        // add token to bridge
        {
            string memory name = string(abi.encodePacked("Cross Bridge ", "TT"));
            bytes32 salt = keccak256(abi.encodePacked(testTokenAtEthereum));
            bytes memory bytecode =
                abi.encodePacked(type(CrossMintableERC20).creationCode, abi.encode(name, "TT", uint8(18)));

            BridgeCross(payable(address(bridgeCross.b))).addTokenDeploy(testTokenAtEthereum, "TT", 18);
            address ttAddress = Create2.computeAddress(salt, keccak256(bytecode), address(bridgeCross.b));
            testTokenAtCross = IERC20(ttAddress);
            BridgeCross(payable(address(bridgeCross.b))).addToken(xcross, IERC20(address(cross)));

            BridgeEthereum(address(bridgeEthereum.b)).addToken(IERC20(address(cross)), xcross);
            BridgeEthereum(address(bridgeEthereum.b)).addToken(testTokenAtEthereum, testTokenAtCross);

            uint initialSupply = 1000000000 * 1e18;
            cross.mint(OWNER, initialSupply);
            TestToken(address(testTokenAtEthereum)).mint(OWNER, initialSupply);
            vm.deal(address(bridgeCross.b), initialSupply);
        }

        vm.stopPrank();
    }

    // ----- Functions -----
    function bridge(BridgeInfo memory info, address token, uint value, uint gas, uint service)
        public
        returns (uint index, bool ok)
    {
        // bridge
        index = info.nextIndex;
        vm.prank(USER);

        if (bridgeRevert) {
            bridgeRevert = false;
            vm.expectRevert();
        }

        if (token == address(xcross)) {
            assertTrue(USER.balance >= value + gas + service);
            ok = info.b.bridge{value: value + gas + service}(IERC20(token), value, gas, service, NULLDATA);
        } else {
            ok = info.b.bridge(IERC20(token), value, gas, service, NULLDATA);
        }
    }

    function finalize(address toBridge, uint index, address token, uint value, uint sigCount)
        public
        returns (bool ok)
    {
        if (sigCount > 5) sigCount = 5;

        // create finalize validator signature
        bytes32 h = keccak256(abi.encode(FINALIZE_TYPEHASH, index, token, USER, value, NULLDATA));
        bytes32 hash = MessageHashUtils.toTypedDataHash(BridgeStandard(toBridge).domainSeparator(), h);

        bytes[] memory sigs = new bytes[](sigCount);
        for (uint i = 0; i < sigCount; i++) {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR_PKs[i], hash);
            sigs[i] = abi.encodePacked(r, s, v);
        }

        // finalize
        vm.prank(VALIDATOR1);
        if (finalizeRevert) {
            finalizeRevert = false;
            vm.expectRevert();
        }
        ok = bridges[toBridge].b.finalize(index, IERC20(token), USER, value, NULLDATA, sigs);
    }

    function calcFee(IERC20 token, uint value) public view returns (uint bridgeValue, uint gas, uint service) {
        BridgeFeeManager.FeeInfo memory feeInfo = bridgeFeeManager.getTokenFee(address(token));
        require(feeInfo.gasFee < value, "value too low");

        gas = feeInfo.gasFee;
        uint denom = bridgeFeeManager.denominator();
        uint serviceRate = feeInfo.serviceFee;

        uint v = value - gas;
        uint t = denom + serviceRate;

        bridgeValue = (v * denom / t);
        service = (v * serviceRate / t);
    }

    function deposit(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum.b));
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross.b).balance;

        bool ok;
        (index, ok) = bridge(bridgeEthereum, address(cross), amount, 0, 0);
        if (ok) {
            bridgeEthereum.nextIndex++;
            finalize(bridgeEthereum.otherBridge, index, address(xcross), amount, validatorNum);
        }

        if (!isRevert) {
            assertEq(userTokenBalance - amount, cross.balanceOf(USER));
            assertEq(bridgeTokenBalance + amount, cross.balanceOf(address(bridgeEthereum.b)));
            assertEq(userCoinBalance + (amount * exrate), USER.balance);
            assertEq(bridgeCoinBalance - (amount * exrate), address(bridgeCross.b).balance);
        }
    }

    function withdraw(bool isRevert, uint amount, uint validatorNum) public returns (uint index) {
        require(amount / exrate > 0, "amount too low");
        (uint value, uint gas, uint service) = calcFee(xcross, amount);
        require(value + gas + service <= amount, "invalid fee");

        // initiate
        uint userTokenBalance = cross.balanceOf(USER);
        uint bridgeTokenBalance = cross.balanceOf(address(bridgeEthereum.b));
        uint userCoinBalance = USER.balance;
        uint bridgeCoinBalance = address(bridgeCross.b).balance;
        uint rewardWalletBalance = REWARD.balance;

        value = (value / exrate) * exrate;
        require(value > 0, "value too low");
        service = (service / exrate) * exrate;

        bool ok;
        (index, ok) = bridge(bridgeCross, address(xcross), value, gas, service);
        if (ok) {
            bridgeCross.nextIndex++;
            finalize(bridgeCross.otherBridge, index, address(cross), value, validatorNum);
        }

        if (!isRevert) {
            assertEq(userTokenBalance + (value / exrate), cross.balanceOf(USER));
            assertEq(bridgeTokenBalance - (value / exrate), cross.balanceOf(address(bridgeEthereum.b)));
            assertEq(userCoinBalance - value + gas + service, USER.balance);
            assertEq(bridgeCoinBalance + value, address(bridgeCross.b).balance);
            assertEq(rewardWalletBalance + gas + service, REWARD.balance);
        }
    }

    function depositToken(bool isRevert, uint amount, uint validatorNum) public {
        // initiate
        uint userEthereumBalance = testTokenAtEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenAtEthereum.balanceOf(address(bridgeEthereum.b));
        uint userCrossBalance = testTokenAtCross.balanceOf(USER);

        (uint index, bool ok) = bridge(bridgeEthereum, address(testTokenAtEthereum), amount, 0, 0);
        if (ok) {
            bridgeEthereum.nextIndex++;
            finalize(bridgeEthereum.otherBridge, index, address(testTokenAtCross), amount, validatorNum);
        }

        if (!isRevert) {
            assertEq(userEthereumBalance - amount, testTokenAtEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance + amount, testTokenAtEthereum.balanceOf(address(bridgeEthereum.b)));
            assertEq(userCrossBalance + amount, testTokenAtCross.balanceOf(USER));
        }
    }

    function withdrawToken(bool isRevert, uint amount, uint validatorNum) public {
        (uint value, uint gas, uint service) = calcFee(IERC20(address(testTokenAtCross)), amount);
        require(value + gas + service <= amount, "invalid fee");
        uint total = value + gas + service;
        assertTrue(total <= amount);

        // initiate
        uint userEthereumBalance = testTokenAtEthereum.balanceOf(USER);
        uint bridgeEthereumBalance = testTokenAtEthereum.balanceOf(address(bridgeEthereum.b));
        uint userCrossBalance = testTokenAtCross.balanceOf(USER);
        uint rewardWalletBalance = testTokenAtCross.balanceOf(REWARD);

        (uint index, bool ok) = bridge(bridgeCross, address(testTokenAtCross), value, gas, service);
        if (ok) {
            bridgeCross.nextIndex++;
            finalize(bridgeCross.otherBridge, index, address(testTokenAtEthereum), value, validatorNum);
        }

        if (!isRevert) {
            assertEq(userEthereumBalance + value, testTokenAtEthereum.balanceOf(USER));
            assertEq(bridgeEthereumBalance - value, testTokenAtEthereum.balanceOf(address(bridgeEthereum.b)));
            assertEq(userCrossBalance - total, testTokenAtCross.balanceOf(USER));
            assertEq(rewardWalletBalance + gas + service, testTokenAtCross.balanceOf(REWARD));
        }
    }
}
