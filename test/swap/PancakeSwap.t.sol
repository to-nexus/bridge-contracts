// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../token/TestToken.sol";
import "forge-std/Test.sol";
import "pancake-smart-contracts/contracts/interfaces/IPancakeRouter02.sol";

contract PancakeSwapTest is Test {
    string constant bscNodeURL = "https://bsc-mainnet.crosstoken.io/2272489872e4f1475ff25d57ce93b51989f933c7"; // use bsc-mainnet

    IPancakeRouter02 public router;
    TestToken public tokenA;
    TestToken public tokenB;
    uint internal constant OWNER_PK = uint(bytes32("owner"));
    uint internal constant USER_PK = uint(bytes32("user"));
    address internal OWNER;
    address internal USER;

    function setUp() public {
        vm.createSelectFork(bscNodeURL);

        OWNER = vm.addr(OWNER_PK);
        USER = vm.addr(USER_PK);

        router = IPancakeRouter02(address(0x10ED43C718714eb63d5aA57B78B54704E256024E)); // use pancake router on bsc-mainnet

        // 테스트용 토큰 배포
        tokenA = new TestToken("Token A", "TKA", 18);
        tokenB = new TestToken("Token B", "TKB", 18);

        // 테스트용 토큰 발행
        tokenA.mint(OWNER, 1000000000 * 10 ** 18);
        tokenB.mint(OWNER, 1000000000 * 10 ** 18);
    }

    // ERC20-ERC20 스왑 테스트
    function testERC20Swap() public {
        // 1. 스왑풀 생성
        vm.startPrank(OWNER);
        tokenA.approve(address(router), type(uint).max);
        tokenB.approve(address(router), type(uint).max);

        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * 10 ** 18, // 1000 TKA
            100 * 10 ** 18, // 100 TKB
            900 * 10 ** 18, // 최소 900 TKA
            90 * 10 ** 18, // 최소 90 TKB
            OWNER,
            block.timestamp + 300
        );

        tokenA.transfer(USER, 1000 * 10 ** 18);
        vm.stopPrank();

        // 2. 스왑 준비
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // getAmountsOut 테스트
        uint[] memory amountsOut = router.getAmountsOut(10 * 10 ** 18, path);
        assertTrue(amountsOut[0] == 10 * 10 ** 18);
        assertTrue(amountsOut[1] > 0);

        // getAmountsIn 테스트
        uint[] memory amountsIn = router.getAmountsIn(10 * 10 ** 18, path);
        assertTrue(amountsIn[0] > 0);
        assertTrue(amountsIn[1] == 10 * 10 ** 18);

        // 3. 스왑 실행
        vm.startPrank(USER);
        tokenA.approve(address(router), 10 * 10 ** 18);
        uint balanceBefore = tokenB.balanceOf(USER);
        router.swapExactTokensForTokens(
            10 * 10 ** 18, // 10 TKA
            amountsOut[1], // 최소 받을 TKB 양
            path,
            USER,
            block.timestamp + 300
        );
        vm.stopPrank();
        // 4. 검증
        uint balanceAfter = tokenB.balanceOf(USER);
        assertTrue(balanceAfter > balanceBefore);
    }

    // ETH-ERC20 스왑 테스트
    // function testETHSwap() public {
    //     // 5. 스왑풀 생성
    //     tokenA.approve(address(router), type(uint).max);

    //     router.addLiquidityETH{value: 100 * 10 ** 18}(
    //         address(tokenA),
    //         100 * 10 ** 18, // 100 TKA
    //         90 * 10 ** 18, // 최소 90 TKA
    //         90 * 10 ** 18, // 최소 90 ETH
    //         address(this),
    //         block.timestamp + 300
    //     );

    //     // 6. 스왑 준비
    //     address[] memory path = new address[](2);
    //     path[0] = address(wbnb);
    //     path[1] = address(tokenA);

    //     // ETH -> ERC20 스왑 준비
    //     uint[] memory amountsOut = router.getAmountsOut(1 * 10 ** 18, path);
    //     assertTrue(amountsOut[0] == 1 * 10 ** 18);
    //     assertTrue(amountsOut[1] > 0);

    //     // 7. ETH -> ERC20 스왑
    //     uint balanceBefore = tokenA.balanceOf(address(this));
    //     router.swapExactETHForTokens{value: 1 * 10 ** 18}(
    //         amountsOut[1], // 최소 받을 토큰 양
    //         path,
    //         address(this),
    //         block.timestamp + 300
    //     );

    //     // ETH -> ERC20 검증
    //     uint balanceAfter = tokenA.balanceOf(address(this));
    //     assertTrue(balanceAfter > balanceBefore);

    //     // ERC20 -> ETH 스왑 준비
    //     path[0] = address(tokenA);
    //     path[1] = address(wbnb);
    //     amountsOut = router.getAmountsOut(10 * 10 ** 18, path);

    //     // ERC20 -> ETH 스왑
    //     balanceBefore = address(this).balance;
    //     router.swapExactTokensForETH(
    //         10 * 10 ** 18, // 10 TKA
    //         amountsOut[1], // 최소 받을 ETH 양
    //         path,
    //         address(this),
    //         block.timestamp + 300
    //     );

    //     // ERC20 -> ETH 검증
    //     balanceAfter = address(this).balance;
    //     assertTrue(balanceAfter > balanceBefore);
    // }
}
