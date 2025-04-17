// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.6;
pragma abicoder v2;

import "forge-std/Script.sol";
import "forge-std/console.sol";

import "../test/token/TestERC20.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/PancakeV3Factory.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/PancakeV3PoolDeployer.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/IPancakeV3Factory.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/IPancakeV3Pool.sol";
import "pancake-v3-contracts/projects/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "pancake-v3-contracts/projects/v3-core/contracts/interfaces/callback/IPancakeV3MintCallback.sol";

/**
 * @title PancakeSwapDeployScript
 * @notice PancakeSwap V3 풀 배포 스크립트
 * @dev 실행 방법: forge script script/PancakeSwapDeploy:PancakeSwapDeployScript --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast -vvv
 * @dev 컴파일러 버전 명시: forge script script/PancakeSwapDeploy:PancakeSwapDeployScript --use solc:0.7.6 --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast -vvv
 */
contract PancakeSwapDeployScript is Script {
    // Fee 상수
    uint24 public constant FEE_LOWEST = 100;  // 0.01%
    uint24 public constant FEE_LOW = 500;     // 0.05%
    uint24 public constant FEE_MEDIUM = 2500; // 0.25%
    uint24 public constant FEE_HIGH = 10000;  // 1%

    // 틱 간격 상수
    int24 public constant TICK_SPACING_LOWEST = 1;
    int24 public constant TICK_SPACING_LOW = 10;
    int24 public constant TICK_SPACING_MEDIUM = 50;
    int24 public constant TICK_SPACING_HIGH = 200;

    // 유동성 Tick
    int24 public constant MIN_TICK = -887272;
    int24 public constant MAX_TICK = 887272;

    // 테스트 토큰 초기 발행량
    uint256 public constant INITIAL_MINT_AMOUNT = 200_000_000 ether;

    // 초기 유동성 공급량
    uint256 public constant INITIAL_LIQUIDITY_AMOUNT = 100_000 ether;

    // 배포된 컨트랙트 주소
    address public pancakeV3PoolDeployer;
    address public pancakeV3Factory;
    address public pancakeV3Pool;

    address public token0;
    address public token1;

    // 풀 생성에 사용할 fee tier
    uint24 private _fee = FEE_LOW;

    // 가격 비율
    // 1:0.1 비율 (약 -23028 틱): 25054144837500004539826032225
    // 1:1   비율 (0 틱):         79228162514264337593543950336
    // 1:10  비율 (약 23050 틱):  250830151128512055300121390458
    uint160 public sqrtPriceX96 = 79228162514264337593543950336;

    /// @dev 초기 유동성 공급 여부
    bool public constant INIT_ADD_LIQUIDITY = false;


    function run() public {

        vm.startBroadcast();

        address deployer = msg.sender;
        console.log("Deploying with address:", deployer);

        // 1. PancakeV3PoolDeployer 배포
        pancakeV3PoolDeployer = deployPoolDeployer();
        console.log("PancakeV3PoolDeployer deployed at:", pancakeV3PoolDeployer);

        // 2. PancakeV3Factory 배포
        pancakeV3Factory = deployFactory(pancakeV3PoolDeployer);
        console.log("PancakeV3Factory deployed at:", pancakeV3Factory);

        // 3. 팩토리 주소를 풀 배포자에 설정
        PancakeV3PoolDeployer(pancakeV3PoolDeployer).setFactoryAddress(pancakeV3Factory);
        console.log("Factory address set in PancakeV3PoolDeployer");

        // 4. 테스트 토큰 배포
        (address tokenA, address tokenB) = deployTestTokens();
        console.log("Test tokens deployed:");
        console.log("Token A:", tokenA);
        console.log("Token B:", tokenB);

        // 토큰을 주소 순으로 정렬 (Uniswap V3 요구사항)
        if (tokenA < tokenB) {
            token0 = tokenA;
            token1 = tokenB;
        } else {
            token0 = tokenB;
            token1 = tokenA;
        }
        console.log("Token0 (address-sorted):", token0);
        console.log("Token1 (address-sorted):", token1);

        // 5. 풀 생성
        pancakeV3Pool = createPool(token0, token1, _fee);
        console.log("PancakeV3Pool created at:", pancakeV3Pool);

        // 6. 풀 초기화 (초기 가격 설정)
        initializePool(pancakeV3Pool);
        console.log("Pool initialized");

        if (INIT_ADD_LIQUIDITY) {
            addLiquidity(pancakeV3Pool);
            console.log("Init add liquidity");
        }

        vm.stopBroadcast();

        // 배포 정보 요약
        console.log("\nDeployment Summary:");
        console.log("-----------------");
        console.log("PancakeV3PoolDeployer: ", pancakeV3PoolDeployer);
        console.log("PancakeV3Factory:      ", pancakeV3Factory);
        console.log("Token0:                ", token0);
        console.log("Token1:                ", token1);
        console.log("PancakeV3Pool:         ", pancakeV3Pool);
        console.log("\nRun command: forge script test/PancakeSwapDeploy.s.sol:PancakeSwapDeployScript --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast -vv");
    }

    /**
     * @notice PancakeV3PoolDeployer 컨트랙트 배포
     */
    function deployPoolDeployer() internal returns (address) {
        PancakeV3PoolDeployer deployer = new PancakeV3PoolDeployer();
        return address(deployer);
    }

    /**
     * @notice PancakeV3Factory 컨트랙트 배포
     * @param _poolDeployer PancakeV3PoolDeployer 주소
     */
    function deployFactory(address _poolDeployer) internal returns (address) {
        PancakeV3Factory factory = new PancakeV3Factory(_poolDeployer);
        return address(factory);
    }

    /**
     * @notice 테스트용 ERC20 토큰 배포
     * @return 배포된 토큰 주소들
     */
    function deployTestTokens() internal returns (address, address) {
        TestERC20 tokenA = new TestERC20(INITIAL_MINT_AMOUNT);
        TestERC20 tokenB = new TestERC20(INITIAL_MINT_AMOUNT);
        return (address(tokenA), address(tokenB));
    }

    /**
     * @notice PancakeSwap V3 풀 생성
     * @param _token0 토큰0 주소
     * @param _token1 토큰1 주소
     * @param _feeTier 수수료 티어
     * @return 생성된 풀 주소
     */
    function createPool(address _token0, address _token1, uint24 _feeTier) internal returns (address) {
        address pool = IPancakeV3Factory(pancakeV3Factory).createPool(
            _token0,
            _token1,
            _feeTier
        );
        return pool;
    }

    /**
     * @notice 풀 초기화 (초기 가격 설정)
     * @param _pool 초기화할 풀 주소
     */
    function initializePool(address _pool) internal {
        IPancakeV3Pool(_pool).initialize(sqrtPriceX96);
        
        // 현재 풀 상태 확인
        (uint160 currentPrice, int24 currentTick, , , , , ) = IPancakeV3Pool(_pool).slot0();
        console.log("Pool initialized with sqrt price:", uint256(currentPrice));
        console.log("Current tick:", int256(currentTick));
    }

    /**
     * @notice 풀에 유동성 추가
     * @param _pool 대상 풀 주소
     */
    function addLiquidity(address _pool) public {

        // TODO: 다른형태에 유동성 공급이 필요하면 구현
        // 현재 풀 상태 확인
        IPancakeV3Pool(_pool).slot0();

        // 전 범위로 유동성 공급
        int24 tickLower = (MIN_TICK / TICK_SPACING_LOW) * TICK_SPACING_LOW;
        int24 tickUpper = (MAX_TICK / TICK_SPACING_LOW) * TICK_SPACING_LOW;

        _addLiquidity(_pool, tickLower, tickUpper, INITIAL_LIQUIDITY_AMOUNT,INITIAL_LIQUIDITY_AMOUNT);
    }
    /**
     * @notice 풀에 유동성 추가
     * @param _pool 대상 풀 주소
     * @param _tickLower 하한 틱
     * @param _tickUpper 상한 틱
     */
    function _addLiquidity(
        address _pool,
        int24 _tickLower,
        int24 _tickUpper,
        uint256, //_amount0
        uint256  //_amount1
    ) private {
        uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(_tickLower);
        uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(_tickUpper);

        uint128 liquidity = LiquidityAmounts.getLiquidityForAmounts(
            sqrtPriceX96,       // 현재 가격
            sqrtRatioAX96,     // 하한 가격
            sqrtRatioBX96,     // 상한 가격
            INITIAL_LIQUIDITY_AMOUNT,      // 목표 token0 수량
            INITIAL_LIQUIDITY_AMOUNT       // 목표 token1 수량
        );
        console.log("Liquidity:", uint256(liquidity));

        SimpleRouter router = new SimpleRouter(token0, token1);
        TestERC20(token0).transfer(address(router), INITIAL_LIQUIDITY_AMOUNT);
        TestERC20(token1).transfer(address(router), INITIAL_LIQUIDITY_AMOUNT);

        // 유동성 공급 (Owner)
        (uint256 amount0Liquidity, uint256 amount1Liquidity) = router.mint(_pool, msg.sender, _tickLower, _tickUpper, liquidity);

        console.log("Token0 Liquidity:", amount0Liquidity);
        console.log("Token1 Liquidity:", amount1Liquidity);
    }
}

contract SimpleRouter is IPancakeV3MintCallback {

    address private _token0;
    address private _token1;

    constructor(address token0, address token1) {
        _token0 = token0;
        _token1 = token1;
    }

    function mint(
        address pool,
        address recipient,
        int24 _tickLower,
        int24 _tickUpper,
        uint128 _liquidity
    ) external returns (uint256 amount0, uint256 amount1){

        // 유동성 공급 (Owner)
        (uint256 amount0Liquidity, uint256 amount1Liquidity) = IPancakeV3Pool(pool).mint(
            recipient,
            _tickLower,
            _tickUpper,
            _liquidity,
            abi.encode(address(_token0), address(_token1))
        );

        return (amount0Liquidity, amount1Liquidity);
    }

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
}