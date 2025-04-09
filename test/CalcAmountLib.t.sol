// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IPriceFeed} from "../src/interface/IPriceFeed.sol";
import {CalcAmountLib} from "../src/lib/CalcAmountLib.sol";
import {Const} from "../src/lib/Const.sol";

import {TestToken} from "./token/TestToken.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {Test} from "forge-std/Test.sol";

// Wrapper contract for testing CalcAmountLib's calculateAmountBWithPrice function
contract CalcAmountLibWrapper {
    function calculateAmountBWithPrice(uint amountA, uint priceA, uint priceB, uint8 decimalA, uint8 decimalB)
        external
        pure
        returns (uint)
    {
        return CalcAmountLib.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
    }

    function decimals(address token) external view returns (uint8) {
        return CalcAmountLib.decimals(token);
    }
}

contract CalcAmountLibTest is Test {
    CalcAmountLibWrapper public wrapper;
    TestToken public token6; // Token with 6 decimals (e.g., USDC)
    TestToken public token8; // Token with 8 decimals (e.g., WBTC)
    TestToken public token18; // Token with 18 decimals (e.g., most ERC20 tokens)

    // Test setup
    function setUp() public {
        wrapper = new CalcAmountLibWrapper();
        token6 = new TestToken("token6", "token6", 6);
        token8 = new TestToken("token8", "token8", 8);
        token18 = new TestToken("token18", "token18", 18);
    }

    // Basic case test: Same decimals (18 -> 18)
    function test_same_decimals() public view {
        uint amountA = 1 ether; // 1 ETH (10^18 wei)
        uint priceA = 1500 * 10 ** 8; // ETH price: $1,500 (8 decimal precision)
        uint priceB = 1 * 10 ** 8; // Token price: $1 (8 decimal precision)
        uint8 decimalA = 18;
        uint8 decimalB = 18;

        uint actualAmount = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
        uint expectedAmount = 1500 * 10 ** 18; // $1,500 worth of token B

        assertEq(actualAmount, expectedAmount);
    }

    // Different decimals case: Higher to lower decimals (18 -> 6)
    function test_higher_to_lower_decimals() public view {
        uint amountA = 1 ether; // 1 ETH (10^18 wei)
        uint priceA = 1500 * 10 ** 8; // ETH price: $1,500
        uint priceB = 1 * 10 ** 8; // USDC price: $1
        uint8 decimalA = 18;
        uint8 decimalB = 6;

        uint actualAmount = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
        uint expectedAmount = 1500 * 10 ** 6; // $1,500 worth of token B (6 decimals)

        assertEq(actualAmount, expectedAmount);
    }

    // Different decimals case: Lower to higher decimals (6 -> 18)
    function test_lower_to_higher_decimals() public view {
        uint amountA = 1000 * 10 ** 6; // 1000 USDC
        uint priceA = 1 * 10 ** 8; // USDC price: $1
        uint priceB = 1500 * 10 ** 8; // ETH price: $1,500
        uint8 decimalA = 6;
        uint8 decimalB = 18;

        uint actualAmount = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
        uint expectedAmount = 666666666666666666; // Approximately 0.6667 ETH

        assertApproxEqRel(actualAmount, expectedAmount, 1e9);
    }

    // Extreme price difference test
    function test_extreme_price_difference() public view {
        uint amountA = 1 ether; // 1 ETH
        uint priceA = 1500 * 10 ** 8; // ETH price: $1,500
        uint priceB = 50000 * 10 ** 8; // BTC price: $50,000
        uint8 decimalA = 18;
        uint8 decimalB = 8;

        uint actualAmount = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
        uint expectedAmount = 3000000; // 0.03 BTC (8 decimals)

        assertApproxEqRel(actualAmount, expectedAmount, 1e9);
    }

    // Zero amount test
    function test_zero_amount() public view {
        uint amountA = 0;
        uint priceA = 1500 * 10 ** 8;
        uint priceB = 1 * 10 ** 8;
        uint8 decimalA = 18;
        uint8 decimalB = 6;

        uint expectedAmount = 0;
        uint actualAmount = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);

        assertEq(actualAmount, expectedAmount);
    }

    // Test revert when priceA is zero
    function test_revert_when_price_a_is_zero() public {
        uint amountA = 1 ether;
        uint priceA = 0; // Zero price
        uint priceB = 1 * 10 ** 8;
        uint8 decimalA = 18;
        uint8 decimalB = 6;

        vm.expectRevert(); // Expect revert
        wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
    }

    // Overflow test
    function test_overflow() public view {
        uint amountA = type(uint).max; // Maximum uint value
        uint priceA = 1 * 10 ** 8;
        uint priceB = 2 * 10 ** 8; // Double price
        uint8 decimalA = 18;
        uint8 decimalB = 18;

        uint amountB = wrapper.calculateAmountBWithPrice(amountA, priceA, priceB, decimalA, decimalB);
        assertEq(amountB, type(uint).max / 2);
    }
}
