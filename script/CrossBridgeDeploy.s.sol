// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {CrossBridge} from "../src/CrossBridge.sol";

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {Upgrades} from "lib/openzeppelin-foundry-upgrades/src/Upgrades.sol";

import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossBridgeDeploy
 * @notice CrossBridge + ForwardLib 배포 스크립트
 * @dev `Implementation.s.sol`을 건드리지 않고 CrossBridge 전용 배포 경로만 제공한다.
 *      프록시는 이미 배포되어 있다고 가정한다 — 이 스크립트는 구현 배포와 업그레이드만
 *      담당하며, 프록시 생성이나 초기화는 하지 않는다(CrossBridge는 initializer가 봉인되어 있다).
 *
 * ============================================================
 * 왜 2단계인가 — ForwardLib은 external 라이브러리다
 * ============================================================
 * `CrossBridge`는 `ForwardLib`의 external 함수를 DELEGATECALL로 호출하므로,
 * 컴파일 산출물에 relocation 3곳이 남는다:
 *
 *   out/CrossBridge.sol/CrossBridge.json
 *     → deployedBytecode.linkReferences["src/lib/ForwardLib.sol"]["ForwardLib"]
 *
 * 따라서 라이브러리를 먼저 배포하고, 그 주소로 링크해서 구현을 배포해야 한다.
 *
 * ============================================================
 * 사용법
 * ============================================================
 *
 * 1) ForwardLib 배포 — 출력된 주소를 기록해 둘 것
 *
 *   forge script script/CrossBridgeDeploy.s.sol:CrossBridgeDeploy \
 *     --sig "deployForwardLib()" \
 *     --rpc-url <RPC_URL> --broadcast
 *
 * 2) 그 주소로 링크해서 CrossBridge 구현 배포
 *
 *   forge script script/CrossBridgeDeploy.s.sol:CrossBridgeDeploy \
 *     --sig "deployCrossBridgeImpl()" \
 *     --libraries src/lib/ForwardLib.sol:ForwardLib:<FORWARD_LIB_ADDR> \
 *     --rpc-url <RPC_URL> --broadcast
 *
 *   `--libraries`를 생략하면 Foundry가 ForwardLib을 자동 배포해 링크한다(주소는
 *   broadcast 산출물에 기록됨). 편하지만 배포마다 라이브러리 주소가 달라지므로,
 *   운영 배포에서는 1)에서 얻은 주소를 명시적으로 링크할 것을 권장한다.
 *
 * 3) 업그레이드 전 링크 무결성 검증 (필수 게이트 — 읽기 전용, 트랜잭션 없음)
 *
 *   forge script script/ForwardLibVerify.s.sol \
 *     --sig 'run(address,address)' <IMPL_ADDR> <FORWARD_LIB_ADDR> \
 *     --rpc-url <RPC_URL>
 *
 *   revert하면 업그레이드를 진행하지 말 것. 자세한 내용은 FORWARD_DEPLOYMENT.md 참조.
 *
 * 4) 프록시 업그레이드 (CrossBridge는 initializer가 봉인되어 있어 초기화 호출 없음)
 *
 *   forge script script/CrossBridgeDeploy.s.sol:CrossBridgeDeploy \
 *     --sig "upgradeToCrossBridge(address,address)" <PROXY> <IMPL_ADDR> \
 *     --rpc-url <RPC_URL> --broadcast
 *
 *   `Implementation.s.sol:upgradeBridgeImplWithoutCall(address,address)`와 동작이
 *   같으므로 그쪽을 써도 무방하다.
 *
 * 5) 업그레이드 후 executor 설정 — `BridgeExecutorDeploy.s.sol` 사용
 *    (target 화이트리스트 + bridgeTokenForwarded selector 단독 등록)
 *
 * ============================================================
 * 가스 주의 (HyperEVM 등 블록 가스 한도가 낮은 체인)
 * ============================================================
 * CrossBridge 구현은 런타임 22,036 B → 코드 예치만 약 4.41M 가스이며, calldata와
 * 실행분을 더하면 배포 트랜잭션이 약 4.8M 가스에 이른다. 블록 가스 한도가 이보다 낮은
 * 체인에서는 배포가 불가능하므로 큰 블록(또는 상응하는 모드)이 필요하다.
 * ForwardLib은 1,381 B(약 0.32M 가스)로 작아 문제되지 않는다.
 */
contract CrossBridgeDeploy is Script {
    function setUp() public {}

    /**
     * @notice ForwardLib 배포
     * @dev 라이브러리는 `new`로 생성할 수 없어 산출물 바이트코드로 배포한다.
     * @return lib 배포된 ForwardLib 주소 — 2)의 `--libraries`에 그대로 사용한다
     */
    function deployForwardLib() public returns (address lib) {
        vm.broadcast();
        lib = deployCode("src/lib/ForwardLib.sol:ForwardLib");

        require(lib.code.length > 0, "ForwardLib deploy failed");

        console.log("ForwardLib deployed to:", lib);
        console.log("  codehash:", vm.toString(lib.codehash));
        console.log("  runtime size:", lib.code.length);
        console.log("");
        console.log("Next: deployCrossBridgeImpl() with");
        console.log("  --libraries src/lib/ForwardLib.sol:ForwardLib:%s", vm.toString(lib));
    }

    /**
     * @notice CrossBridge Implementation 배포
     * @dev ForwardLib이 링크된 상태로 컴파일되어 있어야 한다(위 사용법 2 참조).
     *      프록시는 배포하지 않는다 — 이 스크립트는 기존 프록시 업그레이드 전용이다.
     * @return impl 배포된 구현 주소
     */
    function deployCrossBridgeImpl() public returns (address impl) {
        vm.broadcast();
        CrossBridge implementation = new CrossBridge();
        impl = address(implementation);

        console.log("CrossBridge Implementation deployed to:", impl);
        console.log("  runtime size:", impl.code.length);
        console.log("");
        console.log("Next: run ForwardLibVerify.s.sol against this impl BEFORE upgrading.");
    }

    /**
     * @notice ForwardLib + CrossBridge 구현을 한 번에 배포
     * @dev 편의용. Foundry가 ForwardLib을 자동 배포·링크하므로 라이브러리 주소가
     *      배포마다 달라진다. 운영 배포에서는 2단계 절차를 권장한다.
     * @return lib 배포된 ForwardLib 주소
     * @return impl 배포된 구현 주소
     */
    function deployAll() public returns (address lib, address impl) {
        lib = deployForwardLib();
        impl = deployCrossBridgeImpl();
    }

    /**
     * @notice 프록시를 CrossBridge 구현으로 업그레이드 (초기화 호출 없음)
     * @dev CrossBridge는 `initialize`가 revert 스텁으로 봉인되어 있으므로
     *      업그레이드 시 초기화 calldata를 넘기지 않는다. 스토리지 레이아웃은
     *      CrossBridge와 동일하게 재선언되어 있어 마이그레이션이 필요 없다.
     * @param proxy 프록시 주소
     * @param impl 새 CrossBridge 구현 주소
     */
    function upgradeToCrossBridge(address proxy, address impl) public {
        console.log("Legacy Implementation:", Upgrades.getImplementationAddress(proxy));
        console.log("New Implementation:", impl);

        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(impl, "");

        console.log("Upgraded Implementation:", Upgrades.getImplementationAddress(proxy));
    }
}
