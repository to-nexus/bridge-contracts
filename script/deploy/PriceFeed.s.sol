// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {PriceFeed} from "../../src/PriceFeed.sol";
import {Const} from "../../src/lib/Const.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

contract PriceFeedScript is Script {
    // Common env variables
    string OWNER = "OWNER";
    string IMPLEMENTATION = "IMPLEMENTATION";
    string PRICE_FEED = "PRICE_FEED";
    string PRICER = "PRICER";
    // PriceFeed env variables
    string PRICE_FEED_DOLLAR_DECIMALS = "DOLLAR_DECIMALS";
    string PRICE_FEED_ROLES = "PRICE_FEED_ROLES";
    string PRICE_FEED_ROLE_MEMBERS = "PRICE_FEED_ROLE_MEMBERS";
    string PRICE_FEED_TOKEN_ADDRESSES = "PRICE_FEED_TOKEN_ADDRESSES";
    string PRICE_FEED_TOKEN_PRICES = "PRICE_FEED_TOKEN_PRICES";
    string PRICE_FEED_TOKEN_PRICE_AT = "PRICE_FEED_TOKEN_PRICE_AT";
    string PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS = "PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS";
    string PRICE_FEED_NATIVE_TOKEN_PRICES = "PRICE_FEED_NATIVE_TOKEN_PRICES";
    string PRICE_FEED_NATIVE_TOKEN_PRICE_AT = "PRICE_FEED_NATIVE_TOKEN_PRICE_AT";

    address owner;
    address impl;
    uint8 dollarDecimals;
    bytes32[] roles;
    address[] addresses;

    address[] tokenAddresses;
    uint[] tokenPrices;
    uint[] tokenPriceAt;
    uint[] nativeTokenChainIDs;
    uint[] nativeTokenPrices;
    uint[] nativeTokenPriceAt;

    bytes32[] emptyBytes32Array;
    address[] emptyAddressArray;
    uint[] emptyUintArray;

    function setUp() public {
        // load env variables
        owner = vm.envAddress(OWNER);
        impl = vm.envAddress(IMPLEMENTATION);
        dollarDecimals = uint8(vm.envUint(PRICE_FEED_DOLLAR_DECIMALS));
        roles = vm.envOr(PRICE_FEED_ROLES, ",", emptyBytes32Array);
        addresses = vm.envOr(PRICE_FEED_ROLE_MEMBERS, ",", emptyAddressArray);

        tokenAddresses = vm.envOr(PRICE_FEED_TOKEN_ADDRESSES, ",", emptyAddressArray);
        tokenPrices = vm.envOr(PRICE_FEED_TOKEN_PRICES, ",", emptyUintArray);
        tokenPriceAt = vm.envOr(PRICE_FEED_TOKEN_PRICE_AT, ",", emptyUintArray);
        nativeTokenChainIDs = vm.envOr(PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS, ",", emptyUintArray);
        nativeTokenPrices = vm.envOr(PRICE_FEED_NATIVE_TOKEN_PRICES, ",", emptyUintArray);
        nativeTokenPriceAt = vm.envOr(PRICE_FEED_NATIVE_TOKEN_PRICE_AT, ",", emptyUintArray);

        // print env variables
        console.log("impl", impl);
        console.log("owner", owner);

        require(roles.length == addresses.length, "roles and addresses length mismatch");
        console.log("roles");
        for (uint i = 0; i < roles.length; i++) {
            console.logString(
                string(abi.encodePacked("role: ", vm.toString(roles[i]), ", account: ", vm.toString(addresses[i])))
            );
        }

        if (tokenAddresses.length != tokenPrices.length || tokenAddresses.length != tokenPriceAt.length) {
            console.log("tokenAddresses.length", tokenAddresses.length);
            console.log("tokenPrices.length", tokenPrices.length);
            console.log("tokenPriceAt.length", tokenPriceAt.length);
            revert("Invalid input");
        }
        for (uint i = 0; i < tokenAddresses.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "TokenPrice[",
                        vm.toString(i),
                        "]: {address: ",
                        vm.toString(tokenAddresses[i]),
                        ", price: ",
                        vm.toString(tokenPrices[i]),
                        ", priceAt: ",
                        vm.toString(tokenPriceAt[i]),
                        "}"
                    )
                )
            );
        }

        require(
            nativeTokenChainIDs.length == nativeTokenPrices.length
                && nativeTokenChainIDs.length == nativeTokenPriceAt.length,
            "Invalid input"
        );
        for (uint i = 0; i < nativeTokenChainIDs.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "NativeTokenPrice[",
                        vm.toString(i),
                        "]: {chainID: ",
                        vm.toString(nativeTokenChainIDs[i]),
                        ", price: ",
                        vm.toString(nativeTokenPrices[i]),
                        ", priceAt: ",
                        vm.toString(nativeTokenPriceAt[i]),
                        "}"
                    )
                )
            );
        }
    }

    function setupPriceFeed() public {
        address pricer = vm.envAddress(PRICER);
        PriceFeed priceFeed = PriceFeed(deployProxy(owner, impl, owner, dollarDecimals));
        setRoles(owner, priceFeed, roles, addresses);

        if (tokenAddresses.length > 0) updatePrice(pricer, priceFeed, tokenAddresses, tokenPrices, tokenPriceAt);

        if (nativeTokenChainIDs.length > 0) {
            updateNativeTokenPrice(pricer, priceFeed, nativeTokenChainIDs, nativeTokenPrices, nativeTokenPriceAt);
        }
    }

    function deployProxy(address signer, address impl_, address owner_, uint8 dollarDecimals_)
        public
        returns (address)
    {
        vm.broadcast(signer);
        address proxy =
            address(new ERC1967Proxy(address(impl_), abi.encodeCall(PriceFeed.initialize, (owner_, dollarDecimals_))));
        console.log("PriceFeed deployed at", proxy);
        return proxy;
    }

    function setRoles(address signer, PriceFeed priceFeed, bytes32[] memory roles_, address[] memory addresses_)
        public
    {
        require(roles_.length == addresses_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.grantRoleBatch(roles_, addresses_);
    }

    function updatePrice(
        address signer,
        PriceFeed priceFeed,
        address[] memory tokens_,
        uint[] memory prices_,
        uint[] memory priceAt_
    ) public {
        require(tokens_.length == prices_.length && tokens_.length == priceAt_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.updatePrice(tokens_, prices_, priceAt_);
    }

    function updateNativeTokenPrice(
        address signer,
        PriceFeed priceFeed,
        uint[] memory chainIDs,
        uint[] memory prices_,
        uint[] memory priceAt_
    ) public {
        require(chainIDs.length == prices_.length && chainIDs.length == priceAt_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.updateNativeTokenPrice(chainIDs, prices_, priceAt_);
    }
}
