// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.6;
pragma abicoder v2;

import "./token/TestERC20.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/PancakeV3Factory.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/PancakeV3PoolDeployer.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/IPancakeV3Factory.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/IPancakeV3Pool.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/IPancakeV3PoolDeployer.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/callback/IPancakeV3MintCallback.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/callback/IPancakeV3SwapCallback.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/libraries/TickMath.sol";
import "pancake-v3-contracts/projects/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import {Test, console} from "forge-std/Test.sol";

contract PancakeSwapTest is Test, IPancakeV3MintCallback, IPancakeV3SwapCallback {

    uint internal constant OWNER_PK = uint(bytes32("owner"));

    // Fee and tick spacing constants
    uint24 public constant FEE_LOWEST = 100;  // 0.01%
    uint24 public constant FEE_LOW = 500;     // 0.05%
    uint24 public constant FEE_MEDIUM = 2500; // 0.25%
    uint24 public constant FEE_HIGH = 10000;  // 1%

    int24 public constant TICK_SPACING_LOWEST = 1;
    int24 public constant TICK_SPACING_LOW = 10;
    int24 public constant TICK_SPACING_MEDIUM = 50;
    int24 public constant TICK_SPACING_HIGH = 200;

    // Liquidity constants
    int24 public constant MIN_TICK = -887272;
    int24 public constant MAX_TICK = 887272;

    // Test amounts
    uint256 public constant INITIAL_MINT_AMOUNT = 200_000_000 ether;
    uint256 public constant INITIAL_TESTER_AMOUNT = LIQUIDITY_AMOUNT;
    uint256 public constant LIQUIDITY_AMOUNT = 100_000_000 ether;
    uint256 public constant SWAP_AMOUNT = 1_000 ether;

    uint internal bscForkID;

    address public pancakeV3Factory;
    address public pancakeV3PoolDeployer;
    IPancakeV3Pool public pancakeV3Pool;

    TestERC20 public token0;
    TestERC20 public token1;

    uint24 private _fee = FEE_LOWEST;

    address internal OWNER;

    // PancakeV3MintCallback implementation
    function pancakeV3MintCallback(
        uint256 amount0Owed,
        uint256 amount1Owed,
        bytes calldata data
    ) external override {
        // Transfer required tokens to the pool
        (address token0Address, address token1Address) = abi.decode(data, (address, address));

        if (amount0Owed > 0) {
            TestERC20(token0Address).transfer(msg.sender, amount0Owed);
        }
        if (amount1Owed > 0) {
            TestERC20(token1Address).transfer(msg.sender, amount1Owed);
        }
    }

    // IPancakeSwapCallback implementation
    function pancakeV3SwapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external override {
        // Check that callback is called from the pool
        require(msg.sender == address(pancakeV3Pool), "Invalid callback sender");

        // Decode token information
        (address token0Address, address token1Address) = abi.decode(data, (address, address));

        // Transfer tokens (pay only positive amounts)
        if (amount0Delta > 0) {
            TestERC20(token0Address).transfer(msg.sender, uint256(amount0Delta));
        }
        if (amount1Delta > 0) {
            TestERC20(token1Address).transfer(msg.sender, uint256(amount1Delta));
        }
    }

    function setUp() public {
        bscForkID = vm.createFork("http://localhost:8545"); // bsc test chain
        vm.selectFork(bscForkID);

        OWNER = vm.addr(OWNER_PK);

        // Deploy test tokens
        vm.startPrank(OWNER);
        token0 = new TestERC20(INITIAL_MINT_AMOUNT);
        token1 = new TestERC20(INITIAL_MINT_AMOUNT);

        // Sort tokens by address
        if (address(token0) > address(token1)) {
            (token0, token1) = (token1, token0);
        }

        // Transfer tokens to test contract
        token0.transfer(address(this), INITIAL_TESTER_AMOUNT);
        token1.transfer(address(this), INITIAL_TESTER_AMOUNT);

        vm.stopPrank();
    }

    function test_deploy_PancakeV3Factory() public {
        console.log("------------------------------------------deploy factory-------------------------------------");
        vm.startPrank(OWNER);

        // Deploy PoolDeployer
        pancakeV3PoolDeployer = address(new PancakeV3PoolDeployer());
        console.log("PancakeV3PoolDeployer deployed at:", pancakeV3PoolDeployer);

        // Factory deployment includes default fee settings
        // 100(0.01%), 500(0.05%), 2500(0.25%), 10000(1%) fee tiers are enabled by default
        pancakeV3Factory = address(new PancakeV3Factory(pancakeV3PoolDeployer));
        console.log("PancakeV3Factory deployed at:", pancakeV3Factory);

        // Set Factory address in PoolDeployer
        PancakeV3PoolDeployer(pancakeV3PoolDeployer).setFactoryAddress(pancakeV3Factory);
        console.log("Factory address set in PoolDeployer", pancakeV3PoolDeployer);

        vm.stopPrank();

        // Validate factory state
        validateFactoryState();
    }

    function test_create_pool() public {
        // 0. deploy factory
        test_deploy_PancakeV3Factory();

        console.log("----------------------------------------create pool-----------------------------------------");

        vm.startPrank(OWNER);

        // 1. Create pool
        pancakeV3Pool = IPancakeV3Pool(IPancakeV3Factory(pancakeV3Factory).createPool(
            address(token0),
            address(token1),
            _fee
        ));
        console.log("PancakeV3Pool created at:", address(pancakeV3Pool));

        vm.stopPrank();

        // Check pool existence
        address retrievedPool = IPancakeV3Factory(pancakeV3Factory).getPool(
            address(token0),
            address(token1),
            _fee
        );

        // 2. Validation
        assertNotEq(address(pancakeV3Pool), address(0), "Pool address should not be zero");
        assertEq(address(pancakeV3Pool), retrievedPool, "Created pool and retrieved pool should match");
    }

    /// @dev Pool 초기화
    function test_initialize_pool() public {
        // 0. create pool(token0 : token1)
        test_create_pool();

        console.log("---------------------------------------initialize(set price ratio)-------------------------------");

        vm.startPrank(OWNER);

        // 1. Initialize pool (set price ratio)
        uint160 sqrtPriceX96 = 25054144837500004539826032225;
        /// @dev (1 : 0.1) 25054144837500004539826032225  (Tick : -23028)
        /// @dev (1 :   1) 79228162514264337593543950336  (Tick : 0)
        /// @dev (1 :  10) 250830151128512055300121390458 (Tick : 23050)
        /// @dev (1 : 100) 792281625142643375935439503360 (Tick : 46054)
        pancakeV3Pool.initialize(sqrtPriceX96);
        console.log("Pool initialized with price");

        vm.stopPrank();

        // 2. Validation
        /// @dev 초기화(initialization) 후 Slot0의 상태값이 적용 되었는지 체크
        (uint160 sqrtPriceX96After, int24 tick,,,,,) = pancakeV3Pool.slot0();
        if (sqrtPriceX96After != sqrtPriceX96) {
            console.log("Expected sqrtPriceX96", uint256(sqrtPriceX96));
            console.log("Actual sqrtPriceX96", uint256(sqrtPriceX96After));
            revert("Pool price should be set correctly after initialization");
        }
        if (tick != TickMath.getTickAtSqrtRatio(sqrtPriceX96)) {
            console.log("Expected tick", tick);
            console.log("Actual tick", TickMath.getTickAtSqrtRatio(sqrtPriceX96));
            revert("Tick is invalid");
        }
        console.log("Current tick:", tick);
        console.log("Current sqrt price:", uint256(sqrtPriceX96After));
    }

    function test_pancake_add_liquidity_token0() public {
        // create pool(token0 : token1) and initialize(set price ratio)
        test_initialize_pool();

        console.log("---------------------------------------add_liquidity(min <-> max)-----------------------------------");

        (uint160 sqrtPriceX96, int24 tick,,,,,) = pancakeV3Pool.slot0();
        int24 tickLower = tick+1;
        int24 tickUpper = tick+2;

        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(tickUpper);

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(
            sqrtPriceX96,         // 현재 가격
            sqrtRatioAX96,       // 하한 가격
            sqrtRatioBX96,       // 상한 가격
            LIQUIDITY_AMOUNT,        // 목표 token0 수량
            type(uint256).max        // token1은 최대치로 설정 (제한 없음)
        );
        console.log("Liquidity:", uint256(liquidity));

        // 유동성 공급 (Owner)
        (uint256 amount0, uint256 amount1) = pancakeV3Pool.mint(
            address(this),
            tickLower,
            tickUpper,
            liquidity,
            abi.encode(address(token0), address(token1))
        );

        console.log("Token0 Liquidity:", amount0);
        console.log("Token1 Liquidity:", amount1);
    }

    function test_pancake_add_liquidity_token1() public {
        // create pool(token0 : token1) and initialize(set price ratio)
        test_initialize_pool();

        console.log("---------------------------------------add_liquidity(min <-> max)-----------------------------------");

        (uint160 sqrtPriceX96, int24 tick,,,,,) = pancakeV3Pool.slot0();
        int24 tickLower = tick-2;
        int24 tickUpper = tick-1;

        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(tickUpper);

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(
            sqrtPriceX96,       // 현재 가격
            sqrtRatioAX96,     // 하한 가격
            sqrtRatioBX96,     // 상한 가격
            type(uint256).max,     // token0은 최대치로 설정 (제한 없음)
            LIQUIDITY_AMOUNT       // 목표 token1 수량
        );
        console.log("Liquidity:", uint256(liquidity));

        // 유동성 공급 (Owner)
        (uint256 amount0, uint256 amount1) = pancakeV3Pool.mint(
            OWNER,
            tickLower,
            tickUpper,
            liquidity,
            abi.encode(address(token0), address(token1))
        );

        console.log("Token0 Liquidity:", amount0);
        console.log("Token1 Liquidity:", amount1);
    }

    function test_pancake_add_liquidity_token0_token1() public {
        // create pool(token0 : token1) and initialize(set price ratio)
        test_initialize_pool();

        console.log("---------------------------------------add_liquidity(min <-> max)-----------------------------------");

        (uint160 sqrtPriceX96, int24 tick,,,,,) = pancakeV3Pool.slot0();

        int24 tickSpacing = pancakeV3Pool.tickSpacing();
        int24 tickLower = (tick / tickSpacing) * tickSpacing;
        int24 tickUpper = tickLower + tickSpacing;

        console.log("TickLower:", tickLower);
        console.log("TickUpper:", tickUpper);

        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(tickUpper);

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(
            sqrtPriceX96,       // 현재 가격
            sqrtRatioAX96,     // 하한 가격
            sqrtRatioBX96,     // 상한 가격
            LIQUIDITY_AMOUNT,      // 목표 token0 수량
            LIQUIDITY_AMOUNT       // 목표 token1 수량
        );
        console.log("Liquidity:", uint256(liquidity));

        // 유동성 공급 (Owner)
        (uint256 amount0, uint256 amount1) = pancakeV3Pool.mint(
            OWNER,
            tickLower,
            tickUpper,
            liquidity,
            abi.encode(address(token0), address(token1))
        );

        console.log("Token0 Liquidity:", amount0);
        console.log("Token1 Liquidity:", amount1);
    }

    /// @dev 유동성 공급
    function test_add_liquidity_full_range() public {
        // create pool(token0 : token1) and initialize(set price ratio)
        test_initialize_pool();

        console.log("---------------------------------------add_liquidity(min <-> max)-----------------------------------");

        // Approve tokens
        token0.approve(address(pancakeV3Pool), INITIAL_TESTER_AMOUNT);
        token1.approve(address(pancakeV3Pool), INITIAL_TESTER_AMOUNT);

        (uint160 sqrtPriceX96,,,,,,) = pancakeV3Pool.slot0();

        // Set tick range for liquidity (full range)
        int24 tickLower = (MIN_TICK / TICK_SPACING_LOW) * TICK_SPACING_LOW;
        int24 tickUpper = (MAX_TICK / TICK_SPACING_LOW) * TICK_SPACING_LOW;

        console.log("TickLower:", tickLower);
        console.log("TickUpper:", tickUpper);

        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(tickUpper);

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(
            sqrtPriceX96,       // 현재 가격
            sqrtRatioAX96,     // 하한 가격
            sqrtRatioBX96,     // 상한 가격
            LIQUIDITY_AMOUNT,      // 목표 token0 수량
            LIQUIDITY_AMOUNT       // 목표 token1 수량
        );
        console.log("Liquidity:", uint256(liquidity));

        // 유동성 공급 (Owner)
        (uint256 amount0, uint256 amount1) = pancakeV3Pool.mint(
            OWNER,
            tickLower,
            tickUpper,
            liquidity,
            abi.encode(address(token0), address(token1))
        );

        console.log("Token0 Liquidity:", amount0);
        console.log("Token1 Liquidity:", amount1);

        // 유동성 포지션 확인
        bytes32 positionKey = keccak256(
                    abi.encodePacked(OWNER, tickLower, tickUpper)
        );

        (uint128 liquidityActual,,,,) = IPancakeV3Pool(pancakeV3Pool).positions(positionKey);
        assertEq(uint256(liquidityActual), liquidity, "Liquidity should be added correctly");
    }

    function test_add_liquidity_expected() public {
        // 팩토리와 풀 배포
        test_initialize_pool();

        // 3. 목표 토큰 수량 설정
//        uint256 targetToken0Amount = 100_000_000 ether; // 1억 token0

        // 4. 토큰 승인
        token0.approve(address(pancakeV3Pool), INITIAL_TESTER_AMOUNT);
        token1.approve(address(pancakeV3Pool), INITIAL_TESTER_AMOUNT);

        // 5. 특정 틱 범위 설정
        (, int24 tick,,,,,) = pancakeV3Pool.slot0();

        int24 tickSpacing = pancakeV3Pool.tickSpacing();
        int24 tickLower = (tick / tickSpacing) * tickSpacing;
        int24 tickUpper = tickLower + tickSpacing;

        // 6. 틱에 해당하는 가격 제곱근 계산
        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(tickUpper);

        // 7. 토큰 초기 잔액 확인
        console.log("Token0 balance before mint:", token0.balanceOf(address(this)));
        console.log("Token1 balance before mint:", token1.balanceOf(address(this)));

        // 8. 필요한 유동성 계산 (이진 탐색)
        uint128 lowerLiquidity = 0;
        uint128 upperLiquidity = type(uint128).max / 2;
        uint128 midLiquidity;
        uint256 calculatedToken0;
        uint256 calculatedToken1;

        for (uint i = 0; i < 50; i++) { // 최대 50회 반복
            midLiquidity = uint128((uint256(lowerLiquidity) + uint256(upperLiquidity)) / 2);

            // 계산된 token0 수량 확인
            calculatedToken0 = LiquidityAmounts.getAmount0ForLiquidity(
                sqrtRatioAX96,
                sqrtRatioBX96,
                midLiquidity
            );

            if (calculatedToken0 > 100_000_000 ether) {
                upperLiquidity = midLiquidity;
            } else if (calculatedToken0 < 100_000_000 ether) {
                lowerLiquidity = midLiquidity;
            } else {
                break; // 정확히 일치하면 종료
            }

            // 충분히 근접하면 종료
            if (upperLiquidity - lowerLiquidity < 100) {
                break;
            }
        }

        // 9. 최종 유동성으로 필요한 토큰 양 계산
        calculatedToken0 = LiquidityAmounts.getAmount0ForLiquidity(
            sqrtRatioAX96,
            sqrtRatioBX96,
            midLiquidity
        );
        calculatedToken1 = LiquidityAmounts.getAmount1ForLiquidity(
            sqrtRatioAX96,
            sqrtRatioBX96,
            midLiquidity
        );

        // 10. 유동성 공급
        (uint256 amount0, uint256 amount1) = pancakeV3Pool.mint(
            address(this),
            tickLower,
            tickUpper,
            midLiquidity,
            abi.encode(address(token0), address(token1))
        );


        // 11. 토큰 잔액 변화 확인
        console.log("Token0 balance after mint:", token0.balanceOf(address(this)));
        console.log("Token1 balance after mint:", token1.balanceOf(address(this)));

        console.log("Token0 Liquidity:", amount0);
        console.log("Token1 Liquidity:", amount1);

        // 12. 포지션 확인
        bytes32 positionKey = keccak256(
            abi.encodePacked(address(this), tickLower, tickUpper)
        );

        (uint128 liquidity,,,, ) = IPancakeV3Pool(pancakeV3Pool).positions(positionKey);
        assertEq(uint256(liquidity), uint256(midLiquidity), "invalid liquidity");

        console.log("Test success");
    }

    function test_swap() public {
        // 유동성 공급
        test_add_liquidity_full_range();

        console.log("-------------------------------------------test_swap-----------------------------------------------");

        // 현재 가격 정보 확인
        (uint160 currentSqrtPriceX96, int24 currentTick, , , , , ) = pancakeV3Pool.slot0();
        console.log("Current tick:", int256(currentTick));
        console.log("Current sqrt price:", uint256(currentSqrtPriceX96));

        // 토큰 잔액 확인
        uint256 token0BalanceBefore = token0.balanceOf(address(this));
        uint256 token1BalanceBefore = token1.balanceOf(address(this));
        console.log("Token0 balance before:", token0BalanceBefore);
        console.log("Token1 balance before:", token1BalanceBefore);

        // 매우 작은 금액으로 스왑
        uint256 microSwapAmount = 0.00001 ether;

        // 현재 틱이 음수면 가격을 올리는 스왑이 더 효과적
        bool zeroForOne = false;
        uint160 sqrtPriceLimitX96 = uint160((uint256(currentSqrtPriceX96) * 10001) / 10000);

        // zeroForOne이 true면 가격이 내려가므로 현재 틱이 양수일 때 더 효과적
        console.log("Swap amount:", microSwapAmount);
        console.log("Price limit:", uint256(sqrtPriceLimitX96));

        console.log("-----------------------------------token1 -> token0 swap------------------------------------------");

        // 스왑 시도
        try pancakeV3Pool.swap(
            address(this),
            zeroForOne,
            int256(zeroForOne ? microSwapAmount : microSwapAmount),
            sqrtPriceLimitX96,
            abi.encode(address(token0), address(token1))
        ) returns (int256 amount0Delta, int256 amount1Delta) {
            console.log("Swap succeeded direction (zeroForOne):", zeroForOne);
            console.log("Amount0 delta:", uint256(amount0Delta > 0 ? amount0Delta : -amount0Delta));
            console.log("Amount1 delta:", uint256(amount1Delta > 0 ? amount1Delta : -amount1Delta));

            // 확인
            assertNotEq(amount0Delta, 0, "Amount0 delta should not be zero");
            assertNotEq(amount1Delta, 0, "Amount1 delta should not be zero");

            // 토큰 잔액 변화 확인
            uint256 token0BalanceAfter = token0.balanceOf(address(this));
            uint256 token1BalanceAfter = token1.balanceOf(address(this));
            console.log("Token0 balance change:", int256(token0BalanceAfter) - int256(token0BalanceBefore));
            console.log("Token1 balance change:", int256(token1BalanceAfter) - int256(token1BalanceBefore));
        } catch Error(string memory reason) {
            console.log("Swap failed with reason:", reason);
            // 현재 상태 확인
            (, int24 tickAfterFail, , , , , ) = pancakeV3Pool.slot0();
            console.log("Tick after failure:", int256(tickAfterFail));
        } catch {
            console.log("Swap failed with unknown error");
        }


        console.log("-----------------------------------token0 -> token1 swap------------------------------------------");

        zeroForOne = true;
        // 가격 제한을 현재 가격보다 약간 낮게 설정
        // 주의: 우리는 -10~10 틱 범위에 유동성을 공급했으므로
        // 가격 제한을 tickLower(현재 -10) 근처로 설정하면 안됨
        sqrtPriceLimitX96 = uint160((uint256(currentSqrtPriceX96) * 9999) / 10000);

        // 스왑 시도
        try pancakeV3Pool.swap(
            address(this),
            zeroForOne,
            int256(zeroForOne ? microSwapAmount : microSwapAmount),
            sqrtPriceLimitX96,
            abi.encode(address(token0), address(token1))
        ) returns (int256 amount0Delta, int256 amount1Delta) {
            console.log("Swap succeeded direction (zeroForOne):", zeroForOne);
            console.log("Amount0 delta:", uint256(amount0Delta > 0 ? amount0Delta : -amount0Delta));
            console.log("Amount1 delta:", uint256(amount1Delta > 0 ? amount1Delta : -amount1Delta));

            // 확인
            assertNotEq(amount0Delta, 0, "Amount0 delta should not be zero");
            assertNotEq(amount1Delta, 0, "Amount1 delta should not be zero");

            // 토큰 잔액 변화 확인
            uint256 token0BalanceAfter = token0.balanceOf(address(this));
            uint256 token1BalanceAfter = token1.balanceOf(address(this));
            console.log("Token0 balance change:", int256(token0BalanceAfter) - int256(token0BalanceBefore));
            console.log("Token1 balance change:", int256(token1BalanceAfter) - int256(token1BalanceBefore));
        } catch Error(string memory reason) {
            console.log("Swap failed with reason:", reason);
            // 현재 상태 확인
            (, int24 tickAfterFail, , , , , ) = IPancakeV3Pool(pancakeV3Pool).slot0();
            console.log("Tick after failure:", int256(tickAfterFail));
        } catch {
            console.log("Swap failed with unknown error");
        }

        // 현재 가격 정보 확인
        (currentSqrtPriceX96, currentTick, , , , , ) = IPancakeV3Pool(pancakeV3Pool).slot0();
        console.log("Current tick:", int256(currentTick));
        console.log("Current sqrt price:", uint256(currentSqrtPriceX96));
    }

    function validateFactoryState() internal view {
        // Validate factory state
        address factoryOwner = IPancakeV3Factory(pancakeV3Factory).owner();
        assertEq(factoryOwner, OWNER, "Factory owner is incorrect");
        
        // Validate fee settings
        int24 tickSpacingLow = IPancakeV3Factory(pancakeV3Factory).feeAmountTickSpacing(FEE_LOW);
        int24 tickSpacingMedium = IPancakeV3Factory(pancakeV3Factory).feeAmountTickSpacing(FEE_MEDIUM);
        int24 tickSpacingHigh = IPancakeV3Factory(pancakeV3Factory).feeAmountTickSpacing(FEE_HIGH);
        int24 tickSpacingLowest = IPancakeV3Factory(pancakeV3Factory).feeAmountTickSpacing(FEE_LOWEST);
        
        assertEq(tickSpacingLowest, TICK_SPACING_LOWEST, "Lowest fee tick spacing is incorrect");
        assertEq(tickSpacingLow, TICK_SPACING_LOW, "Low fee tick spacing is incorrect");
        assertEq(tickSpacingMedium, TICK_SPACING_MEDIUM, "Medium fee tick spacing is incorrect");
        assertEq(tickSpacingHigh, TICK_SPACING_HIGH, "High fee tick spacing is incorrect");
        
        // Validate fee activation status
        (bool whitelistRequestedLow, bool enabledLow) = IPancakeV3Factory(pancakeV3Factory).feeAmountTickSpacingExtraInfo(_fee);
        assertEq(whitelistRequestedLow, false, "Expected whitelistRequestedLow to be false, indicating no whitelist request is pending");
        assertEq(enabledLow, true, "Low fee tier should be enabled");
    }
}