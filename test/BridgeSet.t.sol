// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {BridgeVerifier, IBridgeVerifier} from "../src/BridgeVerifier.sol";
import {IPriceFeed} from "../src/PriceFeed.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {BridgeTest} from "./Bridge.t.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BridgeSetTest is BridgeTest {
    function test_set_validator() public {
        vm.startPrank(OWNER);

        vm.selectFork(bscForkID);
        bytes32[] memory roles = new bytes32[](5);
        for (uint i = 0; i < 5; i++) {
            roles[i] = VALIDATOR_ROLE;
        }
        bridgeBSC.revokeRoleBatch(roles, VALIDATORS);
        bool ok = bridgeBSC.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertFalse(ok);

        bridgeBSC.grantRoleBatch(roles, VALIDATORS);
        ok = bridgeBSC.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertTrue(ok);
        vm.stopPrank();

        vm.startPrank(CrossOWNER);
        vm.selectFork(crossForkID);
        bridgeCross.revokeRoleBatch(roles, VALIDATORS);
        ok = bridgeCross.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertFalse(ok);
        bridgeCross.grantRoleBatch(roles, VALIDATORS);
        ok = bridgeCross.hasRole(VALIDATOR_ROLE, VALIDATOR5);
        vm.assertTrue(ok);
        vm.stopPrank();
    }

    function test_set_pause() public {
        uint amount = 1000 ether;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        deposit(false, amount, 5);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setPause(true);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(true);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        vm.prank(USER);
        deposit(true, amount, 5);

        bridgeRevertCross = true;
        vm.prank(USER);
        withdraw(true, amount, 5);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setPause(false);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setPause(false);

        vm.prank(USER);
        deposit(false, amount, 5);
        vm.prank(USER);
        withdraw(false, amount, 5);
    }

    function test_set_threshold() public {
        uint amount = 1000;

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        vm.prank(USER);
        deposit(false, amount, 3);

        vm.selectFork(bscForkID);
        threshold = 4;
        vm.prank(OWNER);
        bridgeBSC.changeThreshold(threshold);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.changeThreshold(threshold);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        finalizeRevertCross = true;
        vm.prank(USER);
        uint index = deposit(true, amount, 3);

        vm.selectFork(crossForkID);
        crossFinalize(index, address(NATIVE_TOKEN), USER, amount, 4);

        vm.selectFork(bscForkID);
        threshold = 1;
        vm.prank(OWNER);
        bridgeBSC.changeThreshold(threshold);
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.changeThreshold(threshold);

        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        vm.prank(USER);
        deposit(false, amount, 1);
    }

    /**
     * @notice Test setting the exchange fee rate
     * @dev Verifies that token-specific exchange fee rates can be set and updated
     */
    function test_set_ex_fee_rate() public {
        uint exFeeRate = 50; // 5%

        // Get bridge manager from the Cross bridge
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        // Set exchange fee rate for a token
        vm.prank(CrossOWNER);
        bridgeVerifierCross.setExFeeRate(IERC20(address(weth)), exFeeRate);

        // Test setting exchange fee rates in batch
        IERC20[] memory tokens = new IERC20[](2);
        tokens[0] = IERC20(address(weth));
        tokens[1] = IERC20(address(testTokenCross));

        uint[] memory rates = new uint[](2);
        rates[0] = 30; // 3%
        rates[1] = 20; // 2%

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setExFeeRateBatch(tokens, rates);
    }

    /**
     * @notice Test setting the finalize bridge gas
     * @dev Verifies that the gas amount required for finalizing bridge operations can be set
     */
    function test_set_finalize_bridge_gas() public {
        uint finalizeBridgeGas = 500000;

        // Set finalize bridge gas on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setFinalizeBridgeGas(finalizeBridgeGas);
    }

    /**
     * @notice Test setting the default token price
     * @dev Verifies that the default token price used when price feed is unavailable can be set
     */
    function test_set_default_token_price() public {
        uint defaultPrice = 1000 * 10 ** 6; // $1000

        // Set default token price on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setDefaultTokenPrice(defaultPrice);
    }

    /**
     * @notice Test setting the default exchange fee rate
     * @dev Verifies that the global default exchange fee rate can be set
     */
    function test_set_default_ex_fee_rate() public {
        uint defaultExFeeRate = 20; // 2%

        // Set default exchange fee rate on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setDefaultExFeeRate(defaultExFeeRate);
    }

    /**
     * @notice Test setting the minimum token value
     * @dev Verifies that the minimum USD value of tokens that can be bridged can be set
     */
    function test_set_minimum_token_value() public {
        uint minimumValue = 10 * 10 ** 6; // $10

        // Set minimum token value on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setMinimumTokenValue(minimumValue);
    }

    /**
     * @notice Test setting the verification amount threshold
     * @dev Verifies that the threshold for requiring verification of large transfers can be set
     */
    function test_set_verification_amount_threshold() public {
        uint threshold = 100000 * 10 ** 6; // $100,000

        // Set verification amount threshold on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setVerificationAmountThreshold(threshold);
    }

    /**
     * @notice Test setting the monitoring time window
     * @dev Verifies that the time window for monitoring token transfer volume can be set
     */
    function test_set_time_window() public {
        uint timeWindow = 4 hours;

        // Set time window on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setTimeWindow(timeWindow);
    }

    /**
     * @notice Test setting the period total value threshold
     * @dev Verifies that the maximum USD value of tokens that can be processed within the time window can be set
     */
    function test_set_period_total_value_threshold() public {
        uint threshold = 1000000 * 10 ** 6; // $1,000,000

        // Set period total value threshold on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setPeriodTotalValueThreshold(threshold);
    }

    /**
     * @notice Test setting the price feed
     * @dev Verifies that the price feed contract reference can be set and removed
     */
    function test_set_price_feed() public {
        address newPriceFeed = address(0xFEED);

        // Set price feed on Cross chain
        vm.selectFork(crossForkID);
        BridgeVerifier bridgeVerifierCross = BridgeVerifier(address(bridgeCross.bridgeVerifier()));

        vm.prank(CrossOWNER);
        bridgeVerifierCross.setPriceFeed(IPriceFeed(newPriceFeed));
    }

    /**
     * @notice Test setting the dev wallet address
     * @dev Verifies that the dev wallet can be updated
     */
    function test_set_dev() public {
        address payable newDev = payable(address(0xbeef));

        // Set dev wallet on BSC chain
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setDev(newDev);

        // Verify dev wallet is set correctly
        address devAddress = bridgeBSC.dev();
        assertEq(devAddress, address(newDev));

        // Set dev wallet on Cross chain
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setDev(newDev);

        // Verify dev wallet is set correctly
        devAddress = bridgeCross.dev();
        assertEq(devAddress, address(newDev));
    }

    /**
     * @notice Test setting chain pause status
     * @dev Verifies that entire chains can be paused and unpaused
     */
    function test_set_chain_pause() public {
        // Pause chain on BSC
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setChainPause(CROSS_CHAIN_ID, true);

        // Try to deposit tokens (should fail due to chain pause)
        uint amount = 1000 ether;
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        vm.prank(USER);
        deposit(true, amount, 5);

        // Unpause chain
        vm.prank(OWNER);
        bridgeBSC.setChainPause(CROSS_CHAIN_ID, false);

        // Now deposit should succeed
        bridgeRevertBSC = false;
        vm.prank(USER);
        deposit(false, amount, 5);

        // Test same operations on Cross chain
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setChainPause(BSC_CHAIN_ID, true);

        // Unpause chain
        vm.prank(CrossOWNER);
        bridgeCross.setChainPause(BSC_CHAIN_ID, false);
    }

    /**
     * @notice Test setting token pause status
     * @dev Verifies that tokens can be paused and unpaused in the bridge
     */
    function test_set_token_pause() public {
        // Pause token on BSC chain
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), true, false);

        // Try to deposit tokens (should fail due to pause)
        uint amount = 1000 ether;
        vm.prank(OWNER);
        cross.transfer(USER, amount);
        vm.prank(USER);
        cross.approve(address(bridgeBSC), amount);

        bridgeRevertBSC = true;
        vm.prank(USER);
        deposit(true, amount, 5);

        // Unpause token
        vm.prank(OWNER);
        bridgeBSC.setTokenPause(CROSS_CHAIN_ID, address(cross), false, false);

        // Now deposit should succeed
        bridgeRevertBSC = false;
        vm.prank(USER);
        deposit(false, amount, 5);

        // Test same operations on Cross chain
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(weth), true, false);

        // Unpause token
        vm.prank(CrossOWNER);
        bridgeCross.setTokenPause(BSC_CHAIN_ID, address(weth), false, false);
    }

    /**
     * @notice Test setting base verification delay
     * @dev Verifies that verification delay can be configured
     */
    function test_set_verification_delay() public {
        uint newDelay = 2 hours;

        // Set verification delay on BSC chain
        vm.selectFork(bscForkID);
        vm.prank(OWNER);
        bridgeBSC.setVerificationDelay(newDelay);

        // Set verification delay on Cross chain
        vm.selectFork(crossForkID);
        vm.prank(CrossOWNER);
        bridgeCross.setVerificationDelay(newDelay);
    }
}
