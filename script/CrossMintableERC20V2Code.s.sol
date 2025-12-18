// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2Code, ICrossMintableERC20Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossMintableERC20V2CodeScript
 * @notice CrossMintableERC20V2Code 배포 및 Bridge 연결 스크립트
 * @dev Bridge에서 새로운 CrossMintableERC20V2 토큰을 자동 생성할 수 있게 하는 Code 컨트랙트를 배포합니다.
 *      배포 후 Bridge에 setCrossMintableERC20Code()로 자동 연결합니다.
 *
 * 사용법:
 *   forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "deployCrossMintableERC20V2Code(address,address)" \
 *     <BRIDGE> <ADMIN> \
 *     --broadcast
 */
contract CrossMintableERC20V2CodeScript is Script {
    function setUp() public {}

    /**
     * @notice CrossMintableERC20V2Code 배포 및 Bridge에 설정
     * @param bridge Bridge 컨트랙트 주소 (프록시 주소)
     * @param admin Code 컨트랙트의 Admin 권한을 가질 주소
     */
    function deployCrossMintableERC20V2Code(BaseBridge bridge, address admin) public {
        vm.startBroadcast();
        CrossMintableERC20V2Code code = new CrossMintableERC20V2Code(admin, address(bridge));
        console.log("CrossMintableERC20V2Code deployed to:", address(code));

        bridge.setCrossMintableERC20Code(ICrossMintableERC20Code(address(code)));
        console.log("CrossMintableERC20V2Code set to bridge");
        vm.stopBroadcast();
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
 * # bridge: BaseBridge 프록시 주소
 * #   - 토큰 자동 생성 기능을 사용할 Bridge
 * #   - ADMIN_ROLE 권한이 있는 계정으로 실행해야 함
 *
 * # admin: Code 컨트랙트의 Admin 권한을 가질 주소
 * #   - 생성되는 토큰들의 기본 Admin 설정에 사용됨
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # CrossMintableERC20V2Code 배포 및 Bridge 연결
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "deployCrossMintableERC20V2Code(address,address)" \
 * #   $BRIDGE $ADMIN \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 배포 후 확인
 * # --------------------------------------------------
 *
 * # Bridge에 Code가 설정되었는지 확인
 * # cast call $BRIDGE "crossMintableERC20Code()" --rpc-url $RPC_URL
 *
 * # --------------------------------------------------
 * # 참고사항
 * # --------------------------------------------------
 *
 * # CrossMintableERC20V2Code가 설정되면 Bridge에서 addTokenPair() 호출 시
 * # 원격 체인의 토큰에 대응하는 CrossMintableERC20V2 토큰이 자동으로 생성됩니다.
 * # 기존에 이미 Code가 설정되어 있으면 업데이트됩니다.
 */
