// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {BaseBridge} from "../../src/BaseBridge.sol";
import {BridgeVerifier} from "../../src/BridgeVerifier.sol";
import {CrossMintableERC20V2Code} from "../../src/token/CrossMintableERC20V2Code.sol";
import {Script, console} from "forge-std/Script.sol";

contract BridgeScript is Script {
    // Common env variables
    string OWNER = "OWNER";
    string PRICE_FEED = "PRICE_FEED";
    string IMPLEMENTATION = "IMPLEMENTATION";

    // Bridge env variables
    string BRIDGE_DEV = "DEV";
    string BRIDGE_THRESHOLD = "THRESHOLD";
    string BRIDGE_CROSS_CHAIN_ID = "CROSS_CHAIN_ID";
    string BRIDGE_BSC_CHAIN_ID = "BSC_CHAIN_ID";
    string BRIDGE_CROSS = "CROSS";
    string BRIDGE_CROSS_INITIAL_SUPPLY = "CROSS_INITIAL_SUPPLY";
    string BRIDGE_ROLES = "BRIDGE_ROLES";
    string BRIDGE_ROLE_MEMBERS = "BRIDGE_ROLE_MEMBERS";

    // Bridge Verifier env variables
    string VERIFIER_PRICER = "PRICER";
    string VERIFIER_FINALIZE_BRIDGE_GAS = "FINALIZE_BRIDGE_GAS";
    string VERIFIER_DEFAULT_TOKEN_PRICE = "DEFAULT_TOKEN_PRICE";
    string VERIFIER_DEFAULT_EX_FEE_RATE = "DEFAULT_EX_FEE_RATE";
    string VERIFIER_MINIMUM_TOKEN_VALUE = "MINIMUM_TOKEN_VALUE";
    string VERIFIER_VERIFICATION_AMOUNT_THRESHOLD = "VERIFICATION_AMOUNT_THRESHOLD";
    string VERIFIER_PERIOD_TOTAL_VALUE_THRESHOLD = "PERIOD_TOTAL_VALUE_THRESHOLD";
    string VERIFIER_TIME_WINDOW = "TIME_WINDOW";
    string VERIFIER_ROLES = "VERIFIER_ROLES";
    string VERIFIER_ROLE_MEMBERS = "VERIFIER_ROLE_MEMBERS";
    string VERIFIER_GAS_PRICES = "VERIFIER_GAS_PRICES";
    string VERIFIER_GAS_PRICE_CHAINS = "VERIFIER_GAS_PRICE_CHAINS";

    address priceFeed;
    address impl;
    address owner;
    address payable dev;
    uint8 threshold;
    uint crossChainID;
    uint bscChainID;
    address cross;
    uint crossInitialSupply;
    uint finalizeBridgeGas;
    uint defaultTokenPrice;
    uint defaultExFeeRate;
    uint minimumTokenValue;
    uint verificationAmountThreshold;
    uint periodTotalValueThreshold;
    uint timeWindow;
    bytes32[] bridgeRoles;
    address[] bridgeRoleMembers;
    bytes32[] verifierRoles;
    address[] verifierRoleMembers;
    uint[] verifierGasPrices;
    uint[] verifierGasPriceChains;

    bytes32[] emptyBytes32Array;
    address[] emptyAddressArray;
    uint[] emptyUintArray;

    function setUp() public virtual {
        // load env variables
        priceFeed = vm.envAddress(PRICE_FEED);
        impl = vm.envAddress(IMPLEMENTATION);
        owner = vm.envAddress(OWNER);
        dev = payable(vm.envAddress(BRIDGE_DEV));
        threshold = uint8(vm.envUint(BRIDGE_THRESHOLD));
        cross = vm.envAddress(BRIDGE_CROSS);
        crossInitialSupply = vm.envUint(BRIDGE_CROSS_INITIAL_SUPPLY) * 1 ether;
        finalizeBridgeGas = vm.envUint(VERIFIER_FINALIZE_BRIDGE_GAS);
        defaultTokenPrice = vm.envUint(VERIFIER_DEFAULT_TOKEN_PRICE);
        defaultExFeeRate = vm.envUint(VERIFIER_DEFAULT_EX_FEE_RATE);
        minimumTokenValue = vm.envUint(VERIFIER_MINIMUM_TOKEN_VALUE);
        verificationAmountThreshold = vm.envUint(VERIFIER_VERIFICATION_AMOUNT_THRESHOLD);
        periodTotalValueThreshold = vm.envUint(VERIFIER_PERIOD_TOTAL_VALUE_THRESHOLD);
        timeWindow = vm.envUint(VERIFIER_TIME_WINDOW);
        bridgeRoles = vm.envOr(BRIDGE_ROLES, ",", emptyBytes32Array);
        bridgeRoleMembers = vm.envOr(BRIDGE_ROLE_MEMBERS, ",", emptyAddressArray);
        verifierRoles = vm.envOr(VERIFIER_ROLES, ",", emptyBytes32Array);
        verifierRoleMembers = vm.envOr(VERIFIER_ROLE_MEMBERS, ",", emptyAddressArray);
        verifierGasPrices = vm.envOr(VERIFIER_GAS_PRICES, ",", emptyUintArray);
        verifierGasPriceChains = vm.envOr(VERIFIER_GAS_PRICE_CHAINS, ",", emptyUintArray);

        // print env variables
        console.log("priceFeed", priceFeed);
        console.log("owner", owner);
        console.log("dev", dev);
        console.log("threshold", threshold);
        console.log("cross", cross);
        console.log("crossInitialSupply", crossInitialSupply);
        console.log("finalizeBridgeGas", finalizeBridgeGas);
        console.log("defaultTokenPrice", defaultTokenPrice);
        console.log("defaultExFeeRate", defaultExFeeRate);
        console.log("minimumTokenValue", minimumTokenValue);
        console.log("verificationAmountThreshold", verificationAmountThreshold);
        console.log("periodTotalValueThreshold", periodTotalValueThreshold);

        require(bridgeRoles.length == bridgeRoleMembers.length, "bridgeRoles and bridgeRoleMembers length mismatch");
        console.log("bridgeRoles");
        for (uint i = 0; i < bridgeRoles.length; i++) {
            // console.logAddress("bridgeRoleMember", bridgeRoleMembers[i]);
            console.logString(
                string(
                    abi.encodePacked(
                        "role: ", vm.toString(bridgeRoles[i]), ", account: ", vm.toString(bridgeRoleMembers[i])
                    )
                )
            );
        }

        require(
            verifierRoles.length == verifierRoleMembers.length, "verifierRoles and verifierRoleMembers length mismatch"
        );
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

    /**
     * @notice Setup Base Bridge contract and initialize
     * command
     * forge script ./script/deploy/Bridge.sol:BridgeScript \
     * -f $RPC_URL \
     * --sender $SENDER \
     * --sig "setupBaseBridge()"
     */
    function setupBaseBridge() public {
        crossChainID = vm.envUint(BridgeScript.BRIDGE_CROSS_CHAIN_ID);
        console.log("crossChainID", crossChainID);

        BaseBridge baseBridge = BaseBridge(payable(deployBaseBridgeProxy()));

        _setupBridge(address(baseBridge));
    }

    function deployBaseBridgeProxy() public returns (address) {
        vm.broadcast();
        address proxy = address(new ERC1967Proxy(impl, abi.encodeCall(BaseBridge.initialize, (owner, dev, threshold))));
        console.log("BaseBridgeProxy will be deployed to", proxy);
        return proxy;
    }

    /**
     * @notice bridge setup after initialize
     */
    function _setupBridge(address bridge) internal {
        BaseBridge baseBridge = BaseBridge(payable(bridge));

        vm.startBroadcast();
        // set bridge
        baseBridge.grantRoleBatch(bridgeRoles, bridgeRoleMembers);
        CrossMintableERC20V2Code erc20 = new CrossMintableERC20V2Code(owner, address(baseBridge));
        baseBridge.setCrossMintableERC20Code(erc20);

        // set verifier
        BridgeVerifier verifier = new BridgeVerifier(
            owner,
            address(baseBridge),
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

        baseBridge.setBridgeVerifier(verifier);

        console.log("CrossMintableERC20V2Code will be deployed to", address(erc20));
        console.log("BridgeVerifier will be deployed to", address(verifier));
        vm.stopBroadcast();

        if (verifierGasPriceChains.length > 0) {
            vm.broadcast(vm.envAddress(VERIFIER_PRICER));
            verifier.updateGasPriceBatch(verifierGasPriceChains, verifierGasPrices);
        }
    }
}
