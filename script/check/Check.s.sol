// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BSCBridge} from "../../src/BSCBridge.sol";
import {BridgeVerifier, IBridgeVerifier} from "../../src/BridgeVerifier.sol";
import {CrossBridge} from "../../src/CrossBridge.sol";
import {PriceFeed} from "../../src/PriceFeed.sol";

import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {Const} from "../../src/lib/Const.sol";
import {CrossMintableERC20Code} from "../../src/token/CrossMintableERC20Code.sol";
import {Script, console} from "forge-std/Script.sol";

contract CheckScript is Script {
    function setUp() public {}

    function checkBSCContract(address cross, PriceFeed priceFeed, BSCBridge bscBridge) public view {
        console.log("--- priceFeed args ---", address(priceFeed));
        console.log("role_admin", priceFeed.getRoleMembers(Const.ADMIN_ROLE)[0]);
        console.log("dollar_decimals", priceFeed.dollarDecimals());
        (, uint crossPrice,) = priceFeed.getTokenPriceInDollars(cross);
        console.log("crossPrice", crossPrice);
        (, uint crossNativePrice,) = priceFeed.getNativeTokenPrice(612055);
        console.log("crossNativePrice", crossNativePrice);

        // ----------------------------------------------------------

        console.log("--- bscBridge args ---", address(bscBridge));
        IBridgeVerifier bridgeVerifier = bscBridge.bridgeVerifier();
        CrossMintableERC20Code crossMintableERC20Code =
            CrossMintableERC20Code(address(bscBridge.crossMintableERC20Code()));
        console.log("role_admin", bscBridge.getRoleMembers(Const.ADMIN_ROLE)[0]);
        console.log("bridgeVerifier", address(bridgeVerifier));
        console.log("crossMintableERC20Code", address(crossMintableERC20Code));
        console.log("threshold", bscBridge.threshold());
        IBridgeRegistry.TokenPair memory tokenPair = bscBridge.getTokenPair(612055, cross);
        console.log("cross token pair");
        console.log("localToken", tokenPair.localToken);
        console.log("remoteToken", tokenPair.remoteToken);
        console.log("isOrigin", tokenPair.isOrigin);
        console.log("paused", tokenPair.paused);
        console.log("pendingAmount", tokenPair.pendingAmount);
        console.log("deposited", tokenPair.deposited);
        console.log("minted", tokenPair.minted);

        console.log("--- bridgeVerifier args ---", address(bridgeVerifier));
        console.log("priceFeed", address(bridgeVerifier.priceFeed()));
        console.log("minimumValue", bridgeVerifier.getMinimumTokenValue());
        console.log("verificationAmountThreshold", bridgeVerifier.getVerificationAmountThreshold());
        console.log("timeWindow", bridgeVerifier.getTimeWindow());
        console.log("periodTotalValueThreshold", bridgeVerifier.getPeriodTotalValueThreshold());
        console.log("denominator", bridgeVerifier.denominator());

        console.log("--- crossMintableERC20Code args ---", address(crossMintableERC20Code));
        console.log("bridge", crossMintableERC20Code.bridge());
    }

    function checkCrossContract(PriceFeed priceFeed, CrossBridge crossBridge) public view {
        console.log("--- priceFeed args ---", address(priceFeed));
        console.log("role_admin", priceFeed.getRoleMembers(Const.ADMIN_ROLE)[0]);
        console.log("dollar_decimals", priceFeed.dollarDecimals());
        (, uint crossPrice,) = priceFeed.getTokenPriceInDollars(Const.NATIVE_TOKEN);
        console.log("crossPrice", crossPrice);
        (, uint bscNativePrice,) = priceFeed.getNativeTokenPrice(56);
        console.log("bscNativePrice", bscNativePrice);

        // ----------------------------------------------------------

        console.log("--- crossBridge args ---", address(crossBridge));
        IBridgeVerifier bridgeVerifier = crossBridge.bridgeVerifier();
        CrossMintableERC20Code crossMintableERC20Code =
            CrossMintableERC20Code(address(crossBridge.crossMintableERC20Code()));
        console.log("role_admin", crossBridge.getRoleMembers(Const.ADMIN_ROLE)[0]);
        console.log("bridgeVerifier", address(bridgeVerifier));
        console.log("crossMintableERC20Code", address(crossMintableERC20Code));
        console.log("threshold", crossBridge.threshold());
        IBridgeRegistry.TokenPair memory tokenPair = crossBridge.getTokenPair(56, Const.NATIVE_TOKEN);
        console.log("cross token pair");
        console.log("localToken", tokenPair.localToken);
        console.log("remoteToken", tokenPair.remoteToken);
        console.log("isOrigin", tokenPair.isOrigin);
        console.log("paused", tokenPair.paused);
        console.log("pendingAmount", tokenPair.pendingAmount);
        console.log("deposited", tokenPair.deposited);
        console.log("minted", tokenPair.minted);

        console.log("--- bridgeVerifier args ---", address(bridgeVerifier));
        console.log("priceFeed", address(bridgeVerifier.priceFeed()));
        console.log("minimumValue", bridgeVerifier.getMinimumTokenValue());
        console.log("verificationAmountThreshold", bridgeVerifier.getVerificationAmountThreshold());
        console.log("timeWindow", bridgeVerifier.getTimeWindow());
        console.log("periodTotalValueThreshold", bridgeVerifier.getPeriodTotalValueThreshold());
        console.log("denominator", bridgeVerifier.denominator());

        console.log("--- crossMintableERC20Code args ---", address(crossMintableERC20Code));
        console.log("bridge", crossMintableERC20Code.bridge());
    }
}
