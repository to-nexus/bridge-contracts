// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {CrossMintableERC20V2} from "../src/token/CrossMintableERC20V2.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossMintableERC20V2Script
 * @notice CrossMintableERC20V2 토큰을 `CrossMintableERC20V2Code` 팩토리 없이 독립적으로 배포
 * @dev 개정 3(BeaconProxy 전환)부터 토큰은 상수생성자 대신 `initialize`를 쓰고, 생성자에서
 *      `_disableInitializers()`를 호출하므로 `new CrossMintableERC20V2(...)`로 곧장 초기화할
 *      수 없다 — 반드시 `BeaconProxy`를 거쳐야 한다. 팩토리를 쓰지 않는 이 독립 배포 경로는
 *      매 호출마다 전용 1회용 `UpgradeableBeacon`을 만들어 `initialOwner`에게 소유시킨다
 *      (팩토리가 만드는 토큰들과 달리 다른 토큰과 beacon을 공유하지 않음 — 한 토큰만 관리하는
 *      한도 내에서 계속 업그레이드 가능하게 하려는 목적).
 *
 * 사용법:
 *   forge script script/CrossMintableERC20V2.s.sol:CrossMintableERC20V2Script \
 *     --rpc-url <RPC_URL> \
 *     --sig "deployCrossMintableERC20V2(address,address,string,string,uint8)" \
 *     <OWNER> <MINTER> "<NAME>" "<SYMBOL>" <DECIMALS> \
 *     --broadcast
 */
contract CrossMintableERC20V2Script is Script {
    function setUp() public {}

    /**
     * @notice CrossMintableERC20V2 토큰을 전용 beacon과 함께 배포
     * @param initialOwner 토큰의 `defaultAdmin()`이 될 주소 (이 배포의 전용 beacon `owner()`도 됨)
     * @param initialMinter mint 권한을 가질 주소 (보통 Bridge 컨트랙트)
     * @param name 토큰 이름 (예: "Wrapped BTC")
     * @param symbol 토큰 심볼 (예: "WBTC")
     * @param decimals 소수점 자릿수 (보통 18)
     */
    function deployCrossMintableERC20V2(
        address initialOwner,
        address initialMinter,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) public {
        vm.startBroadcast();
        CrossMintableERC20V2 implementation = new CrossMintableERC20V2();
        UpgradeableBeacon beacon = new UpgradeableBeacon(address(implementation), initialOwner);
        BeaconProxy proxy = new BeaconProxy(
            address(beacon),
            abi.encodeCall(CrossMintableERC20V2.initialize, (initialOwner, initialMinter, name, symbol, decimals))
        );
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2 implementation deployed to:", address(implementation));
        console.log("CrossMintableERC20V2 beacon (owner, must be multisig/timelock) deployed to:", address(beacon));
        console.log("CrossMintableERC20V2 deployed to:", address(proxy));
    }
}

/*
 * ===================================
 * 사용 예제
 * ===================================
 *
 * # --------------------------------------------------
 * # 파라미터 설명
 * # --------------------------------------------------
 *
 * # initialOwner: 토큰의 Admin 권한을 가질 주소
 * #   - DEFAULT_ADMIN_ROLE 부여됨
 * #   - 역할 관리, 토큰 설정 변경 가능
 *
 * # initialMinter: mint/burn 권한을 가질 주소
 * #   - MINTER_ROLE 부여됨
 * #   - 보통 Bridge 컨트랙트 주소를 지정
 *
 * # name: 토큰 이름 (예: "Wrapped Bitcoin", "Cross USDT")
 * # symbol: 토큰 심볼 (예: "WBTC", "xUSDT")
 * # decimals: 소수점 자릿수 (원본 토큰과 동일하게 설정)
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # 예시: Wrapped BTC 토큰 배포
 * # forge script script/CrossMintableERC20V2.s.sol:CrossMintableERC20V2Script \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "deployCrossMintableERC20V2(address,address,string,string,uint8)" \
 * #   0x...owner 0x...bridge "Wrapped BTC" "WBTC" 8 \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 배포 후 확인
 * # --------------------------------------------------
 *
 * # 토큰 정보 확인
 * # cast call $TOKEN "name()" --rpc-url $RPC_URL
 * # cast call $TOKEN "symbol()" --rpc-url $RPC_URL
 * # cast call $TOKEN "decimals()" --rpc-url $RPC_URL
 *
 * # Minter 역할 확인
 * # cast call $TOKEN "hasRole(bytes32,address)" \
 * #   $(cast keccak "MINTER_ROLE") $BRIDGE --rpc-url $RPC_URL
 */
