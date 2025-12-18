// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {CrossMintableERC20V2} from "../src/token/CrossMintableERC20V2.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossMintableERC20V2Script
 * @notice CrossMintableERC20V2 토큰 배포 스크립트
 * @dev Bridge에서 mint/burn 가능한 ERC20 토큰을 독립적으로 배포합니다.
 *      CrossMintableERC20V2Code를 통한 자동 생성과 달리, 수동으로 토큰을 배포할 때 사용합니다.
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
     * @notice CrossMintableERC20V2 토큰 배포
     * @param initialOwner 토큰의 Admin 권한을 가질 주소
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
        vm.broadcast();
        CrossMintableERC20V2 token = new CrossMintableERC20V2(initialOwner, initialMinter, name, symbol, decimals);
        console.log("CrossMintableERC20V2 deployed to:", address(token));
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
