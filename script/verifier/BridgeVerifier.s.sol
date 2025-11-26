// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../../src/BaseBridge.sol";
import {BridgeVerifier} from "../../src/BridgeVerifier.sol";
import {PriceFeed} from "../../src/PriceFeed.sol";
import {Const} from "../../src/lib/Const.sol";
import {Script, console} from "forge-std/Script.sol";

contract BridgeVerifierScript is Script {
    // Bridge Verifier env variables
    string OWNER = "OWNER";
    string PRICE_FEED = "PRICE_FEED";
    string BRIDGE = "BRIDGE";
    string PRICER = "PRICER";
    string FINALIZE_BRIDGE_GAS = "FINALIZE_BRIDGE_GAS";
    string DEFAULT_TOKEN_PRICE = "DEFAULT_TOKEN_PRICE";
    string DEFAULT_EX_FEE_RATE = "DEFAULT_EX_FEE_RATE";
    string MINIMUM_TOKEN_VALUE = "MINIMUM_TOKEN_VALUE";
    string VERIFICATION_AMOUNT_THRESHOLD = "VERIFICATION_AMOUNT_THRESHOLD";
    string PERIOD_TOTAL_VALUE_THRESHOLD = "PERIOD_TOTAL_VALUE_THRESHOLD";
    string TIME_WINDOW = "TIME_WINDOW";
    string ROLES = "ROLES";
    string ROLE_MEMBERS = "ROLE_MEMBERS";
    string GAS_PRICES = "GAS_PRICES";
    string GAS_PRICE_CHAINS = "GAS_PRICE_CHAINS";

    address priceFeed;
    address owner;
    BaseBridge bridge;
    uint8 threshold;
    uint bscChainID;
    uint finalizeBridgeGas;
    uint defaultTokenPrice;
    uint defaultExFeeRate;
    uint minimumTokenValue;
    uint verificationAmountThreshold;
    uint periodTotalValueThreshold;
    uint timeWindow;
    bytes32[] verifierRoles;
    address[] verifierRoleMembers;
    uint[] verifierGasPrices;
    uint[] verifierGasPriceChains;

    bytes32[] emptyBytes32Array;
    address[] emptyAddressArray;
    uint[] emptyUintArray;

    function setUp() public virtual {
        // load env variables
        bridge = BaseBridge(vm.envAddress(BRIDGE));
        priceFeed = vm.envAddress(PRICE_FEED);
        owner = vm.envAddress(OWNER);
        finalizeBridgeGas = vm.envUint(FINALIZE_BRIDGE_GAS);
        defaultTokenPrice = vm.envUint(DEFAULT_TOKEN_PRICE);
        defaultExFeeRate = vm.envUint(DEFAULT_EX_FEE_RATE);
        minimumTokenValue = vm.envUint(MINIMUM_TOKEN_VALUE);
        verificationAmountThreshold = vm.envUint(VERIFICATION_AMOUNT_THRESHOLD);
        periodTotalValueThreshold = vm.envUint(PERIOD_TOTAL_VALUE_THRESHOLD);
        timeWindow = vm.envUint(TIME_WINDOW);
        verifierRoles = vm.envOr(ROLES, ",", emptyBytes32Array);
        verifierRoleMembers = vm.envOr(ROLE_MEMBERS, ",", emptyAddressArray);
        verifierGasPrices = vm.envOr(GAS_PRICES, ",", emptyUintArray);
        verifierGasPriceChains = vm.envOr(GAS_PRICE_CHAINS, ",", emptyUintArray);

        // print env variables
        console.log("bridge", address(bridge));
        console.log("priceFeed", priceFeed);
        console.log("owner", owner);
        console.log("finalizeBridgeGas", finalizeBridgeGas);
        console.log("defaultTokenPrice", defaultTokenPrice);
        console.log("defaultExFeeRate", defaultExFeeRate);
        console.log("minimumTokenValue", minimumTokenValue);
        console.log("verificationAmountThreshold", verificationAmountThreshold);
        console.log("periodTotalValueThreshold", periodTotalValueThreshold);

        // require(
        //     verifierRoles.length == verifierRoleMembers.length, "verifierRoles and verifierRoleMembers length mismatch"
        // );
        console.log("verifierRoles");
        for (uint i = 0; i < verifierRoles.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "role: ", vm.toString(verifierRoles[i]), ", account: ", vm.toString(verifierRoleMembers[i])
                    )
                )
            );
        }

        require(
            verifierGasPrices.length == verifierGasPriceChains.length,
            "verifierGasPrices and verifierGasPriceChains length mismatch"
        );
        console.log("verifierGasPrices");
        for (uint i = 0; i < verifierGasPrices.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "chainId: ",
                        vm.toString(verifierGasPriceChains[i]),
                        ", gasPrice: ",
                        vm.toString(verifierGasPrices[i])
                    )
                )
            );
        }
    }

    function deployBridgeVerifier() public {
        vm.startBroadcast();
        // set verifier
        BridgeVerifier verifier = new BridgeVerifier(
            owner,
            address(bridge), // bridge
            priceFeed,
            finalizeBridgeGas,
            defaultTokenPrice,
            defaultExFeeRate,
            minimumTokenValue,
            verificationAmountThreshold,
            periodTotalValueThreshold,
            timeWindow
        );
        verifier.grantRoleBatch(verifierRoles, verifierRoleMembers);

        bridge.setBridgeVerifier(verifier);
        console.log("BridgeVerifier will be deployed to", address(verifier));
        vm.stopBroadcast();

        if (verifierGasPriceChains.length > 0) {
            vm.broadcast(vm.envAddress(PRICER));
            verifier.updateGasPriceBatch(verifierGasPriceChains, verifierGasPrices);
        }
    }
}
