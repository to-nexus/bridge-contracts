// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {IBridgeRegistry} from "../src/interface/IBridgeRegistry.sol";
import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2} from "../src/token/CrossMintableERC20V2.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {ICrossMintableERC20V2Code} from "../src/token/ICrossMintableERC20V2Code.sol";

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title CrossMintableERC20V2CodeScript
 * @notice `CrossMintableERC20V2Code`(+ 토큰 beacon) 배포 · 업그레이드 · Bridge 연결 스크립트
 * @dev 두 컨트랙트 모두 프록시다:
 *      - 토큰(`CrossMintableERC20V2`)은 `UpgradeableBeacon` + `BeaconProxy`. 토큰이 몇 개든
 *        beacon 업그레이드 한 번으로 전부 반영된다.
 *      - 팩토리(`CrossMintableERC20V2Code`)는 UUPS(`ERC1967Proxy`), 단일 인스턴스.
 *      두 impl 모두 생성자에서 `_disableInitializers()`를 호출하므로 `ERC1967Proxy` /
 *      `BeaconProxy`의 초기화 데이터를 통해서만 초기화된다(원자적 초기화).
 *
 *      `HyperMintableERC20Code`와 달리 beacon은 별도로 배포·소유하지 않는다 — 팩토리
 *      `initialize`가 beacon을 직접 만들고 스스로 소유한다. 따라서 beacon 업그레이드도
 *      beacon을 직접 부르지 않고 **팩토리의 `upgradeBeacon`을 거쳐야 한다** — beacon
 *      `owner()`가 팩토리 프록시 주소이지 EOA/멀티시그가 아니기 때문이다. 신뢰가 팩토리
 *      `ADMIN_ROLE` 하나로 모이므로 **그 계정은 반드시 멀티시그/타임락**이어야 한다.
 *
 *      CREATE2 토큰 주소의 initcode에는 **beacon 주소만** 들어가고 그 순간의 로직 impl 주소는
 *      들어가지 않으므로, 토큰 로직을 업그레이드해도 향후 예측 주소는 바뀌지 않는다.
 *
 * 사용법:
 *   forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 *     --rpc-url $RPC_URL \
 *     --sig "deployAll(address,address,address)" \
 *     $BRIDGE $FACTORY_ADMIN $INITIAL_BRIDGE_ROLE \
 *     --broadcast
 *
 * CROSS 테스트넷(비표준 포트 RPC)에서는 `forge script --broadcast`가 조용히 실패할 수 있으므로
 * (시뮬레이션 로그만 출력되고 트랜잭션 미전송) 상태 변경은 `cast send`로 하고, 매 단계 온체인
 * 상태를 직접 조회해 확인할 것 — 아래 각 함수를 단계별로 나눠 둔 이유이기도 하다(재개 지점
 * 판별용). BSC 테스트넷은 `--gas-price`를 명시할 것(실측 0.05 gwei 표시가는 mempool에 걸림).
 */
contract CrossMintableERC20V2CodeScript is Script {
    function setUp() public {}

    // =====================================================================
    // 배포 — 개별 단계 (4단계: 토큰 impl -> 팩토리 impl -> 팩토리 프록시(비콘 자체 생성) -> Bridge 연결)
    // =====================================================================

    /**
     * @notice `CrossMintableERC20V2` 로직(impl) 배포 — 팩토리가 소유할 beacon이 가리킬 구현체
     * @dev 생성자에서 `_disableInitializers()`를 호출하므로 이 주소를 직접 `initialize`할 수
     *      없다 — 반드시 beacon을 거쳐 `BeaconProxy`로만 초기화된다.
     * @return implementation 배포된 impl 주소
     */
    function deployCrossMintableERC20V2Implementation() public returns (address implementation) {
        vm.startBroadcast();
        CrossMintableERC20V2 impl = new CrossMintableERC20V2();
        implementation = address(impl);
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2 implementation deployed to:", implementation);
    }

    /**
     * @notice `CrossMintableERC20V2Code` 로직(impl) 배포
     * @dev 생성자에서 `_disableInitializers()`를 호출하므로 이 주소를 직접 `initialize`할
     *      수 없다 — 반드시 `ERC1967Proxy`를 거쳐서만 초기화된다.
     * @return implementation 배포된 impl 주소
     */
    function deployCrossMintableERC20V2CodeImplementation() public returns (address implementation) {
        vm.startBroadcast();
        CrossMintableERC20V2Code impl = new CrossMintableERC20V2Code();
        implementation = address(impl);
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2Code implementation deployed to:", implementation);
    }

    /**
     * @notice 팩토리 `ERC1967Proxy` 배포 + 원자 초기화 (beacon은 `initialize` 내부에서 팩토리
     *         자신이 직접 만들어 소유한다 — 별도 beacon 배포 단계가 없다)
     * @dev `factoryAdmin`은 **반드시 멀티시그/타임락 주소**를 쓸 것 — 비콘 업그레이드
     *      (전 토큰 로직 교체) · 토큰 역할 관리 패스쓰루 · 팩토리 자신의 업그레이드 승인을
     *      전부 이 한 계정이 쥔다. 브로드캐스트 계정은 `factoryAdmin`일 필요는
     *      없다(`initialize`가 public `grantRole`이 아니라 `_grantRole`을 쓰므로) —
     *      단순히 이 프록시를 배포·초기화할 수 있는 아무 계정이면 된다.
     * @param codeImplementation `deployCrossMintableERC20V2CodeImplementation`의 반환값
     * @param factoryAdmin 팩토리 `ADMIN_ROLE`을 받을 주소(canonical owner, **멀티시그/타임락 필수**)
     * @param initialBridge 팩토리 `BRIDGE_ROLE`을 받을 주소(레거시 `createCrossMintableERC20`
     *        경로를 호출할 Bridge). 0이면 건너뜀 — 나중에 팩토리 `ADMIN_ROLE`로 별도 부여 가능
     * @param tokenImplementation `deployCrossMintableERC20V2Implementation`의 반환값 — 팩토리가
     *        생성할 beacon의 초기 구현체
     * @return code 배포된 팩토리 프록시 주소
     */
    function deployCrossMintableERC20V2CodeProxy(
        address codeImplementation,
        address factoryAdmin,
        address initialBridge,
        address tokenImplementation
    ) public returns (address code) {
        vm.startBroadcast();
        ERC1967Proxy proxy = new ERC1967Proxy(
            codeImplementation,
            abi.encodeCall(
                CrossMintableERC20V2Code.initialize, (factoryAdmin, initialBridge, tokenImplementation)
            )
        );
        code = address(proxy);
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2Code proxy deployed to:", code);
        console.log("  implementation:", codeImplementation);
        console.log("  factoryAdmin (ADMIN_ROLE, must be multisig/timelock):", factoryAdmin);
        console.log("  initialBridge (BRIDGE_ROLE):", initialBridge);
        console.log("  beacon (created by initialize):", ICrossMintableERC20V2Code(code).beacon());
    }

    /**
     * @notice 팩토리를 Bridge에 연결한다 (Bridge의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 함)
     * @dev 이후 Bridge의 `createToken`이 이 팩토리를 거쳐 레거시 파생-name/symbol 경로로
     *      토큰을 만든다. **호출 전에 반드시 프리플라이트
     *      (`preflightCheckDuplicateRemoteToken`)를 대상 체인의 기존 페어 전부에 대해
     *      실행해 둘 것** — 이 함수 자체는 페어 단위가 아니라 팩토리 교체 자체이므로
     *      remoteToken 중복을 검사하지 않는다.
     * @param bridge Bridge 컨트랙트 주소 (프록시 주소)
     * @param code `deployCrossMintableERC20V2CodeProxy`의 반환값
     */
    function connectToBridge(BaseBridge bridge, address code) public {
        vm.startBroadcast();
        bridge.setCrossMintableERC20Code(ICrossMintableERC20Code(code));
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2Code set to bridge:");
        console.log("  bridge:", address(bridge));
        console.log("  code:", code);
    }

    /**
     * @notice impl(토큰) + impl(팩토리) + 팩토리 프록시(비콘 내부 생성) + Bridge 연결을 한 번에
     * @dev 위 네 함수를 순서대로 묶은 편의 함수. **CROSS 테스트넷 RPC에서는 `--broadcast`가
     *      조용히 실패할 수 있으므로** 각 단계를 개별 함수로 나눠 두었다 — 실패가 의심되면
     *      이 편의 함수 대신 단계별로 실행하고 매 단계 온체인 상태를 직접 조회해 확인할 것.
     * @param bridge Bridge 컨트랙트 주소 (연결 단계는 bridge의 `ADMIN_ROLE` 계정으로 브로드캐스트)
     * @param factoryAdmin 팩토리 `ADMIN_ROLE`을 받을 주소. **멀티시그/타임락 필수**
     * @param initialBridge 팩토리 `BRIDGE_ROLE`을 받을 주소 (보통 `bridge`와 동일)
     * @return tokenImplementation 토큰 impl 주소
     * @return codeImplementation 팩토리 impl 주소
     * @return code 팩토리 프록시 주소
     */
    function deployAll(BaseBridge bridge, address factoryAdmin, address initialBridge)
        public
        returns (address tokenImplementation, address codeImplementation, address code)
    {
        tokenImplementation = deployCrossMintableERC20V2Implementation();
        codeImplementation = deployCrossMintableERC20V2CodeImplementation();
        code = deployCrossMintableERC20V2CodeProxy(codeImplementation, factoryAdmin, initialBridge, tokenImplementation);
        connectToBridge(bridge, code);
    }

    // =====================================================================
    // 업그레이드
    // =====================================================================

    /**
     * @notice 토큰 로직을 업그레이드한다 — 이 팩토리가 만든 모든 토큰에 즉시 반영된다
     * @dev beacon `owner()`는 팩토리 프록시 자신이므로 beacon을 직접 부르지 않고 **팩토리의
     *      `upgradeBeacon`을 거친다** — 팩토리의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 한다.
     *      단계적 롤아웃이 불가하므로 업그레이드 전 전수 테스트가 필수다. CREATE2 initcode에는
     *      beacon 주소만 들어가므로 이 업그레이드는 향후 `computeTokenAddress`/
     *      `computeTokenAddressWithName` 예측값에 영향을 주지 않는다.
     * @param code 팩토리 프록시 주소
     * @param newImplementation 새 `CrossMintableERC20V2` impl 주소
     */
    function upgradeTokenBeacon(address code, address newImplementation) public {
        vm.startBroadcast();
        ICrossMintableERC20V2Code(code).upgradeBeacon(newImplementation);
        vm.stopBroadcast();

        console.log("token beacon upgraded via factory:");
        console.log("  code:", code);
        console.log("  beacon:", ICrossMintableERC20V2Code(code).beacon());
        console.log("  newImplementation:", newImplementation);
    }

    /**
     * @notice 팩토리 로직을 업그레이드한다
     * @dev 팩토리의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 한다(`_authorizeUpgrade`).
     * @param code 팩토리 프록시 주소
     * @param newImplementation 새 `CrossMintableERC20V2Code` impl 주소
     */
    function upgradeCrossMintableERC20V2Code(address code, address newImplementation) public {
        vm.startBroadcast();
        CrossMintableERC20V2Code(code).upgradeToAndCall(newImplementation, bytes(""));
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2Code upgraded:");
        console.log("  proxy:", code);
        console.log("  newImplementation:", newImplementation);
    }

    // =====================================================================
    // 소유권 이전 — DEFAULT_ADMIN_ROLE 과 ADMIN_ROLE 은 별개이므로 4단계
    // =====================================================================
    //
    // 4단계는 서명자가 다르다(1=구 owner, 2~4=신 owner). `msg.sender`로 서명자를 판별하지
    // 않는다 — forge 스크립트에서 `msg.sender`는 스크립트 컨트랙트 자신이고 `vm.startBroadcast`는
    // 기록되는 외부 호출의 발신자만 바꾼다. 따라서 검사 대상 주소를 전부 명시 인자로 받고, 각
    // 실행 래퍼는 브로드캐스트 발신자도 명시한다(`vm.startBroadcast(oldOwner)` /
    // `vm.startBroadcast(newOwner)`) — 무인자 `vm.startBroadcast()`는 두 신원을 동시에
    // 만족시킬 수 없다.
    //
    // 검증(view)과 실행(브로드캐스트)을 분리한다: `beginTransferStatus` / `acceptTransferStatus`
    // / `revokeOldAdminStatus`는 순수 view이며 상태만 보고 `Ready`/`AlreadyDone`/`Inconsistent`를
    // 반환한다. 공개 실행 함수는 이 판정을 호출해 `Inconsistent`면 그 `reason`으로 revert,
    // `AlreadyDone`이면 로그만 남기고 return, `Ready`면 브로드캐스트한다.
    //
    // `begin`/`accept`는 멱등 skip(`AlreadyDone`)이 안전하다 — 사후조건이 신원을 특정하기
    // 때문이다. 반대로 **`revokeOldAdmin`은 절대 `AlreadyDone`을 반환하지 않는다** —
    // `ADMIN_ROLE` 부재만으로는 "이미 취소됨"과 "oldOwner 를 잘못 넣음"을 스테이트리스하게
    // 구분할 수 없으므로, 멱등 skip 을 두면 오타 하나가 구 관리자의 특권을 조용히 방치하는
    // 경로가 된다. `revokeOldAdmin`을 재실행하면 (이미 완료됐어도) revert하는 것이
    // 의도된 안전 동작이다 — 온체인 상태를 직접 확인해야 한다.

    /// @dev 4단계 이전의 진행 상태 판정 결과. `Ready`=사전조건 충족, 실행 함수가 브로드캐스트를
    /// 진행해야 함. `AlreadyDone`=사후조건이 이미 충족됨, 실행 함수는 로그만 남기고 스킵.
    /// `Inconsistent`=사전조건도 사후조건도 아님, 실행 함수는 `reason`으로 revert해야 함.
    enum TransferStage {
        Ready,
        AlreadyDone,
        Inconsistent
    }

    /**
     * @notice 1단계의 진행 상태를 판정한다 (순수 view, 브로드캐스트 없음)
     * @dev `Ready`: `defaultAdmin() == oldOwner`이고 아직 `newOwner`로 pending 되지 않음.
     *      `AlreadyDone`: `defaultAdmin() == newOwner`(1~2단계 완료) 이거나
     *      `pendingDefaultAdmin() == newOwner`(1단계만 완료, `defaultAdmin()`은 여전히
     *      `oldOwner`). `Inconsistent`: 위 어느 것도 아님 — `defaultAdmin()`이 `oldOwner`도
     *      `newOwner`도 아닌 제3의 주소인 경우. `defaultAdmin() == oldOwner` 확인이 pending
     *      분기보다 먼저 이뤄져야 한다(그렇지 않으면 제3자가 잘못 시작한 pending 이전을
     *      "이미 완료"로 오판할 수 있다).
     * @param factory 팩토리 프록시 주소
     * @param oldOwner 현재 `defaultAdmin()`이어야 하는 주소
     * @param newOwner `DEFAULT_ADMIN_ROLE`을 이전받을 주소
     */
    function beginTransferStatus(address factory, address oldOwner, address newOwner)
        public
        view
        returns (TransferStage, string memory reason)
    {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);
        address currentAdmin = target.defaultAdmin();

        if (currentAdmin == newOwner) {
            return (TransferStage.AlreadyDone, "defaultAdmin() is already newOwner - transfer already completed");
        }

        if (currentAdmin != oldOwner) {
            return (
                TransferStage.Inconsistent,
                string.concat(
                    "defaultAdmin() is ", vm.toString(currentAdmin), ", expected oldOwner ", vm.toString(oldOwner)
                )
            );
        }

        (address pendingNewAdmin,) = target.pendingDefaultAdmin();
        if (pendingNewAdmin == newOwner) {
            return (TransferStage.AlreadyDone, "pendingDefaultAdmin() is already newOwner - transfer already begun");
        }

        return (TransferStage.Ready, "");
    }

    /**
     * @notice 2·3단계의 진행 상태를 판정한다 (순수 view, 브로드캐스트 없음)
     * @dev `Ready`: `pendingDefaultAdmin() == newOwner`(수락 대기 중) 이거나
     *      `defaultAdmin() == newOwner`(수락은 끝났고 `ADMIN_ROLE` 부여만 남음). `AlreadyDone`:
     *      사후조건 `defaultAdmin()==newOwner && hasRole(ADMIN_ROLE,newOwner)`이 이미 참 —
     *      신원이 `newOwner` 하나로 특정되므로 멱등 skip 이 안전하다. `Inconsistent`:
     *      `defaultAdmin()`도 `pendingDefaultAdmin()`도 `newOwner`가 아님 —
     *      `beginOwnershipTransfer`가 아직 실행되지 않았다는 뜻.
     * @param factory 팩토리 프록시 주소
     * @param newOwner `DEFAULT_ADMIN_ROLE`을 수락하고 `ADMIN_ROLE`을 받을 주소
     */
    function acceptTransferStatus(address factory, address newOwner)
        public
        view
        returns (TransferStage, string memory reason)
    {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);
        bool defaultAdminIsNewOwner = target.defaultAdmin() == newOwner;

        if (defaultAdminIsNewOwner && target.hasRole(Const.ADMIN_ROLE, newOwner)) {
            return (
                TransferStage.AlreadyDone,
                "defaultAdmin() and ADMIN_ROLE both already confirmed for newOwner - transfer already completed"
            );
        }

        if (!defaultAdminIsNewOwner) {
            (address pendingNewAdmin,) = target.pendingDefaultAdmin();
            if (pendingNewAdmin != newOwner) {
                return (
                    TransferStage.Inconsistent,
                    string.concat(
                        "neither defaultAdmin() (",
                        vm.toString(target.defaultAdmin()),
                        ") nor pendingDefaultAdmin() (",
                        vm.toString(pendingNewAdmin),
                        ") is newOwner (",
                        vm.toString(newOwner),
                        ") - run beginOwnershipTransfer first"
                    )
                );
            }
        }

        return (TransferStage.Ready, "");
    }

    /**
     * @notice 4단계의 진행 상태를 판정한다 (순수 view, 브로드캐스트 없음)
     * @dev **절대 `AlreadyDone`을 반환하지 않는다** — `hasRole(ADMIN_ROLE, oldOwner) == false`는
     *      "이미 취소됨"과 "oldOwner 가 애초에 틀렸음"을 스테이트리스하게 구분할 수 없기
     *      때문이다. 세 사전조건을 모두 요구하며, 어느 하나라도 깨지면 진단 가능한
     *      `reason`과 함께 `Inconsistent`를 반환한다:
     *        1. `defaultAdmin() == newOwner` — 이전이 실제로 완료됐는지
     *        2. `hasRole(ADMIN_ROLE, newOwner) == true` — 관리자 0명 방지
     *        3. `hasRole(ADMIN_ROLE, oldOwner) == true` — 신원 확인 (없으면 조용한 통과 차단)
     *      더해 `oldOwner != newOwner`도 요구한다 — 그렇지 않으면 조건 2·3이 같은 주소를
     *      가리키게 되어 자기 자신의 `ADMIN_ROLE`을 회수하는 `Ready` 판정이 나올 수 있다.
     * @param factory 팩토리 프록시 주소
     * @param oldOwner `ADMIN_ROLE`을 회수당할 주소
     * @param newOwner 이미 `ADMIN_ROLE`을 보유하고 있어야 하는 주소
     */
    function revokeOldAdminStatus(address factory, address oldOwner, address newOwner)
        public
        view
        returns (TransferStage, string memory reason)
    {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);

        if (oldOwner == newOwner) {
            return (TransferStage.Inconsistent, "oldOwner == newOwner");
        }

        address currentAdmin = target.defaultAdmin();
        if (currentAdmin != newOwner) {
            return (
                TransferStage.Inconsistent,
                string.concat(
                    "defaultAdmin() is ",
                    vm.toString(currentAdmin),
                    ", expected newOwner ",
                    vm.toString(newOwner),
                    " - ownership transfer (begin/accept) not complete yet"
                )
            );
        }

        if (!target.hasRole(Const.ADMIN_ROLE, newOwner)) {
            return (
                TransferStage.Inconsistent,
                string.concat(
                    "newOwner (",
                    vm.toString(newOwner),
                    ") is missing ADMIN_ROLE - run acceptOwnershipAndGrantAdmin first"
                )
            );
        }

        if (!target.hasRole(Const.ADMIN_ROLE, oldOwner)) {
            return (
                TransferStage.Inconsistent,
                string.concat(
                    "oldOwner (",
                    vm.toString(oldOwner),
                    ") is missing ADMIN_ROLE - already revoked, or oldOwner is the wrong address; verify manually"
                )
            );
        }

        return (TransferStage.Ready, "");
    }

    /**
     * @notice 1단계: 팩토리 `DEFAULT_ADMIN_ROLE`을 `oldOwner`에서 `newOwner`로 이전
     *         시작한다 (`oldOwner`의 키로 브로드캐스트됨 — `vm.startBroadcast(oldOwner)`)
     * @dev 판정은 `beginTransferStatus`에 위임하는 얇은 래퍼다. `Inconsistent`면 그 `reason`으로
     *      revert, `AlreadyDone`이면 로그만 남기고 return, `Ready`면 브로드캐스트한다.
     * @param factory 팩토리 프록시 주소
     * @param oldOwner 현재 `defaultAdmin()`이어야 하는 주소
     * @param newOwner `DEFAULT_ADMIN_ROLE`을 이전받을 주소
     */
    function beginOwnershipTransfer(address factory, address oldOwner, address newOwner) public {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);

        (TransferStage stage, string memory reason) = beginTransferStatus(factory, oldOwner, newOwner);
        if (stage == TransferStage.Inconsistent) {
            revert(reason);
        }
        if (stage == TransferStage.AlreadyDone) {
            console.log("beginOwnershipTransfer: already done, skipping broadcast.");
            console.log("  factory:", factory);
            console.log("  newOwner:", newOwner);
            return;
        }

        vm.startBroadcast(oldOwner);
        target.beginDefaultAdminTransfer(newOwner);
        vm.stopBroadcast();

        (address pendingNewAdmin,) = target.pendingDefaultAdmin();
        require(pendingNewAdmin == newOwner, "beginOwnershipTransfer: postcondition failed, pendingDefaultAdmin() != newOwner");

        console.log("beginOwnershipTransfer: pendingDefaultAdmin() ==", newOwner);
        console.log("  factory:", factory);
        console.log("  oldOwner:", oldOwner);
    }

    /**
     * @notice 2·3단계: `newOwner`가 `DEFAULT_ADMIN_ROLE`을 수락하고 스스로에게
     *         `ADMIN_ROLE`을 부여한다 (`newOwner`의 키로 브로드캐스트됨 —
     *         `vm.startBroadcast(newOwner)`)
     * @dev 판정은 `acceptTransferStatus`에 위임하는 얇은 래퍼다. `Inconsistent`면 그 `reason`으로
     *      revert, `AlreadyDone`이면 로그만 남기고 return, `Ready`면 두 하위 단계(수락 /
     *      `ADMIN_ROLE` 부여)를 필요한 것만 브로드캐스트한다 — 둘 다 신원이 `newOwner`로
     *      특정되므로 독립적으로 멱등 skip 해도 안전하다.
     * @param factory 팩토리 프록시 주소
     * @param newOwner `DEFAULT_ADMIN_ROLE`을 수락하고 `ADMIN_ROLE`을 받을 주소
     */
    function acceptOwnershipAndGrantAdmin(address factory, address newOwner) public {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);

        (TransferStage stage, string memory reason) = acceptTransferStatus(factory, newOwner);
        if (stage == TransferStage.Inconsistent) {
            revert(reason);
        }
        if (stage == TransferStage.AlreadyDone) {
            console.log("acceptOwnershipAndGrantAdmin: already done (defaultAdmin() and ADMIN_ROLE both confirmed), skipping.");
            console.log("  factory:", factory);
            console.log("  newOwner:", newOwner);
            return;
        }

        if (target.defaultAdmin() == newOwner) {
            console.log("acceptOwnershipAndGrantAdmin: DEFAULT_ADMIN_ROLE already accepted by newOwner, skipping accept step.");
        } else {
            vm.startBroadcast(newOwner);
            target.acceptDefaultAdminTransfer();
            vm.stopBroadcast();

            require(
                target.defaultAdmin() == newOwner,
                "acceptOwnershipAndGrantAdmin: defaultAdmin() != newOwner after acceptDefaultAdminTransfer"
            );
        }

        if (target.hasRole(Const.ADMIN_ROLE, newOwner)) {
            console.log("acceptOwnershipAndGrantAdmin: newOwner already has ADMIN_ROLE, skipping grant step.");
        } else {
            vm.startBroadcast(newOwner);
            target.grantRole(Const.ADMIN_ROLE, newOwner);
            vm.stopBroadcast();
        }

        require(target.defaultAdmin() == newOwner, "acceptOwnershipAndGrantAdmin: postcondition failed, defaultAdmin() != newOwner");
        require(
            target.hasRole(Const.ADMIN_ROLE, newOwner),
            "acceptOwnershipAndGrantAdmin: postcondition failed, newOwner missing ADMIN_ROLE"
        );

        console.log("acceptOwnershipAndGrantAdmin: defaultAdmin() and ADMIN_ROLE both confirmed for", newOwner);
        console.log("  factory:", factory);
    }

    /**
     * @notice 4단계: `oldOwner`의 `ADMIN_ROLE`을 회수한다 (`newOwner`의 키로
     *         브로드캐스트됨 — `vm.startBroadcast(newOwner)`. `newOwner`가
     *         `acceptOwnershipAndGrantAdmin` 이후 `DEFAULT_ADMIN_ROLE` 보유자이므로
     *         `ADMIN_ROLE`의 role-admin 자격이 있다)
     * @dev 판정은 `revokeOldAdminStatus`에 위임하는 얇은 래퍼다. **그 함수는 절대
     *      `AlreadyDone`을 반환하지 않으므로** 이 함수도 멱등 skip 이 없다 — `Inconsistent`면
     *      그 `reason`으로 revert한다(이미 완료됐거나 `oldOwner`가 잘못됐다는 뜻이므로 온체인
     *      상태를 직접 확인해야 한다). 재실행 시 매번 revert하는 것이 의도된 안전
     *      동작이다.
     * @param factory 팩토리 프록시 주소
     * @param oldOwner `ADMIN_ROLE`을 회수당할 주소
     * @param newOwner 이미 `ADMIN_ROLE`을 보유하고 있어야 하는 주소
     */
    function revokeOldAdmin(address factory, address oldOwner, address newOwner) public {
        CrossMintableERC20V2Code target = CrossMintableERC20V2Code(factory);

        (TransferStage stage, string memory reason) = revokeOldAdminStatus(factory, oldOwner, newOwner);
        if (stage == TransferStage.Inconsistent) {
            revert(reason);
        }
        // revokeOldAdminStatus never returns AlreadyDone (see its @dev) - Ready is the only
        // remaining case here.

        vm.startBroadcast(newOwner);
        target.revokeRole(Const.ADMIN_ROLE, oldOwner);
        vm.stopBroadcast();

        require(!target.hasRole(Const.ADMIN_ROLE, oldOwner), "revokeOldAdmin: postcondition failed, oldOwner still has ADMIN_ROLE");

        console.log("revokeOldAdmin: ADMIN_ROLE revoked from oldOwner:", oldOwner);
        console.log("  factory:", factory);
        console.log("  newOwner (canonical owner going forward):", newOwner);
    }

    // =====================================================================
    // 중복 remoteToken 프리플라이트 — 온체인 가드가 아니라 배포 절차 가드
    // =====================================================================

    /**
     * @notice 대상 `remoteChainID`에 이미 같은 `remoteToken`으로 등록된 페어가 있으면 revert한다
     * @dev `BridgeRegistry._registerToken`은 localToken 기준으로만 중복을 막으므로,
     *      새 팩토리가 다른 localToken을 만들면 이 가드를 그냥 통과해 버린다 — 즉 리매핑이
     *      아니라 두 번째 페어가 추가 등록되어 기존 잔고와 신규 유동성이 분리된 채 공존하게
     *      된다. `createToken`(레거시, 브릿지 경유)과 `createMintableERC20`+수동
     *      `registerToken`(신규 경로) **양쪽 모두 호출 전에** 이 함수를 먼저 실행할 것.
     *      읽기 전용, 트랜잭션 없음.
     * @param bridge Bridge 컨트랙트 주소
     * @param remoteChainID 확인할 원격 체인 ID
     * @param remoteToken 확인할 원격 토큰 주소
     */
    function preflightCheckDuplicateRemoteToken(BaseBridge bridge, uint remoteChainID, address remoteToken)
        public
        view
    {
        IBridgeRegistry.TokenPair[] memory pairs = bridge.allTokenPairs(remoteChainID);
        for (uint i = 0; i < pairs.length; ++i) {
            if (pairs[i].remoteToken == remoteToken) {
                console.log("preflight FAILED: remoteToken already registered under existing localToken:");
                console.log("  remoteChainID:", remoteChainID);
                console.log("  remoteToken:", remoteToken);
                console.log("  existing localToken:", pairs[i].localToken);
                revert(
                    "preflightCheckDuplicateRemoteToken: remoteToken already registered for this remoteChainID under a different localToken"
                );
            }
        }
        console.log("preflight OK: no existing pair for this remoteChainID/remoteToken:");
        console.log("  remoteChainID:", remoteChainID);
        console.log("  remoteToken:", remoteToken);
    }

    // =====================================================================
    // 토큰 생성 — 파생 경로 (심볼에서 name/symbol 파생, 브릿지 경유, 자동 registerToken)
    // =====================================================================

    /**
     * @notice 배포 전 토큰 CREATE2 주소를 미리 계산한다 (읽기 전용, 트랜잭션 없음)
     * @dev `minter`에는 **반드시 실제로 `createToken`을 호출할 브리지 주소**를 넣을 것.
     *      틀린 주소를 넣어도 revert하지 않고 다른 주소를 조용히 반환한다(의도적 설계).
     * @param code CrossMintableERC20V2Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param symbol 토큰 심볼 (name은 "Cross Bridge <symbol>", symbol은 "<symbol>x"로 파생)
     * @param decimals 토큰 decimals
     * @param minter 실제 `createCrossMintableERC20`을 호출할 주소 (보통 bridge)
     * @return predicted 예측된 토큰 주소
     */
    function computeTokenAddress(
        address code,
        uint remoteChainID,
        address remoteToken,
        string memory symbol,
        uint8 decimals,
        address minter
    ) public view returns (address predicted) {
        predicted =
            ICrossMintableERC20V2Code(code).computeTokenAddress(remoteChainID, remoteToken, symbol, decimals, minter);
        console.log("predicted CrossMintableERC20V2 address:", predicted);
        console.log("  minter used for prediction (must equal the actual createToken caller):", minter);
    }

    /**
     * @notice `preflightCheckDuplicateRemoteToken` 실행 후 `createToken`을 실행하고, 필요하면
     *         사전 예측과 실제 주소가 일치하는지 검증한다
     * @dev `expected != 0`이면 사전 예측을 강제하며 불일치 시 revert한다(단계 1->2 연결 확인용).
     *      `expected == 0`이면 예측 없이 그냥 배포하고 실제 주소를 로그로 남긴다(기본 경로).
     *      브로드캐스트 계정은 bridge의 `EDITOR_ROLE`을 보유해야 한다(`registerToken` 게이트).
     * @param bridge Bridge 컨트랙트 주소
     * @param code CrossMintableERC20V2Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param symbol 토큰 심볼
     * @param decimals 토큰 decimals
     * @param expected 사전 예측 주소. 0이면 검증을 건너뜀
     * @return tokenAddress 실제 생성된 토큰 주소
     */
    function createTokenAndVerify(
        BaseBridge bridge,
        address code,
        uint remoteChainID,
        address remoteToken,
        string memory symbol,
        uint8 decimals,
        address expected
    ) public returns (address tokenAddress) {
        preflightCheckDuplicateRemoteToken(bridge, remoteChainID, remoteToken);

        vm.startBroadcast();
        tokenAddress = bridge.createToken(remoteChainID, remoteToken, symbol, decimals);
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2 created at:", tokenAddress);
        console.log("  isCrossMintableERC20:", ICrossMintableERC20V2Code(code).isCrossMintableERC20(tokenAddress));

        if (expected != address(0)) {
            require(tokenAddress == expected, "createTokenAndVerify: address mismatch vs prediction");
            console.log("  matches pre-computed address:", expected);
        }
    }

    // =====================================================================
    // 토큰 생성 — 명시 경로 (임의 name/symbol + minter 직접 지정, 브릿지 미경유)
    // =====================================================================

    /**
     * @notice 배포 전 토큰 CREATE2 주소를 미리 계산한다 (명시 name/symbol 경로, 읽기 전용)
     * @param code CrossMintableERC20V2Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals 토큰 decimals
     * @param minter `createMintableERC20` 호출 시 넘길 minter 인자 (호출자가 아니라
     *        `MINTER_ROLE`을 받을 주소, 보통 bridge)
     * @return predicted 예측된 토큰 주소
     */
    function computeTokenAddressWithName(
        address code,
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter
    ) public view returns (address predicted) {
        predicted = ICrossMintableERC20V2Code(code).computeTokenAddressWithName(
            remoteChainID, remoteToken, name_, symbol_, decimals, minter
        );
        console.log("predicted CrossMintableERC20V2 address (explicit name/symbol):", predicted);
    }

    /**
     * @notice `preflightCheckDuplicateRemoteToken` 실행 후, 임의 name/symbol + 명시 minter로
     *         토큰을 생성한다 — 브릿지의 `createToken`을 거치지 않으므로 필요하면 별도로
     *         `bridge.registerToken`을 호출해야 한다
     * @dev 팩토리의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 한다. `expected != 0`이면 사전
     *      예측과 일치하는지 검증한다.
     * @param bridge 중복 프리플라이트에 사용할 Bridge 컨트랙트 주소
     * @param code CrossMintableERC20V2Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals 토큰 decimals
     * @param minter `MINTER_ROLE`을 받을 주소
     * @param expected 사전 예측 주소. 0이면 검증을 건너뜀
     * @return tokenAddress 실제 생성된 토큰 주소
     */
    function createMintableERC20AndVerify(
        BaseBridge bridge,
        address code,
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter,
        address expected
    ) public returns (address tokenAddress) {
        preflightCheckDuplicateRemoteToken(bridge, remoteChainID, remoteToken);

        vm.startBroadcast();
        tokenAddress =
            ICrossMintableERC20V2Code(code).createMintableERC20(remoteChainID, remoteToken, name_, symbol_, decimals, minter);
        vm.stopBroadcast();

        console.log("CrossMintableERC20V2 created at (explicit name/symbol):", tokenAddress);
        console.log("  name:", name_);
        console.log("  symbol:", symbol_);
        console.log("  minter (MINTER_ROLE):", minter);
        console.log(
            "  NOTE: not routed through bridge.createToken() -> call bridge.registerToken(...) separately if this token must be bridgeable"
        );

        if (expected != address(0)) {
            require(tokenAddress == expected, "createMintableERC20AndVerify: address mismatch vs prediction");
            console.log("  matches pre-computed address:", expected);
        }
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
 * #   - connectToBridge/deployAll은 bridge의 ADMIN_ROLE 계정으로 브로드캐스트해야 함
 * #     (setCrossMintableERC20Code 호출)
 * #   - createTokenAndVerify는 bridge의 EDITOR_ROLE 계정으로 브로드캐스트해야 함
 * #     (createToken -> 내부 registerToken 호출)
 *
 * # factoryAdmin: 팩토리 ADMIN_ROLE을 받을 주소
 * #   - 비콘 업그레이드(전 토큰 로직 교체) · 토큰 역할 관리 패스쓰루
 * #     (grantTokenRole/revokeTokenRole/beginTokenDefaultAdminTransfer) · 팩토리 자신의
 * #     업그레이드 승인을 전부 이 한 계정이 쥔다. **반드시 멀티시그/타임락**.
 *
 * # --------------------------------------------------
 * # 배포 (impl 2종 + 팩토리 프록시(비콘 내부 생성) + Bridge 연결까지 한 번에)
 * # --------------------------------------------------
 *
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "deployAll(address,address,address)" \
 * #   $BRIDGE $FACTORY_ADMIN $INITIAL_BRIDGE_ROLE \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # (선택) 사전 주소 계산 — 파생 경로
 * # --------------------------------------------------
 *
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --sig "computeTokenAddress(address,uint256,address,string,uint8,address)" \
 * #   $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $BRIDGE
 *
 * # --------------------------------------------------
 * # 중복 remoteToken 프리플라이트 (createToken/createMintableERC20 전에 필수)
 * # --------------------------------------------------
 *
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --sig "preflightCheckDuplicateRemoteToken(address,uint256,address)" \
 * #   $BRIDGE $REMOTE_CHAIN_ID $REMOTE_TOKEN
 *
 * # --------------------------------------------------
 * # 토큰 생성 — 파생 경로
 * # --------------------------------------------------
 *
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "createTokenAndVerify(address,address,uint256,address,string,uint8,address)" \
 * #   $BRIDGE $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $EXPECTED_OR_ZERO \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 토큰 생성 — 명시 name/symbol 경로 (ADMIN_ROLE, 브릿지 미경유)
 * # --------------------------------------------------
 *
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "createMintableERC20AndVerify(address,address,uint256,address,string,string,uint8,address,address)" \
 * #   $BRIDGE $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $NAME $SYMBOL $DECIMALS $MINTER $EXPECTED_OR_ZERO \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 업그레이드
 * # --------------------------------------------------
 *
 * # 토큰 로직 업그레이드 (factoryAdmin 계정 — 팩토리의 upgradeBeacon을 거침, 전 토큰에 즉시 반영)
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "upgradeTokenBeacon(address,address)" \
 * #   $CODE $NEW_TOKEN_IMPLEMENTATION \
 * #   --broadcast
 *
 * # 팩토리 로직 업그레이드 (factoryAdmin 계정, ADMIN_ROLE)
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "upgradeCrossMintableERC20V2Code(address,address)" \
 * #   $CODE $NEW_CODE_IMPLEMENTATION \
 * #   --broadcast
 *
 * ===================================================================================
 * 배포 전 체인별 사전 조사 — 구현 범위 밖, 배포 체크리스트 항목
 * ===================================================================================
 *
 * 각 대상 체인에서 `bridge.crossMintableERC20Code()`(현재 팩토리 종류)와
 * `bridge.allTokenPairs(remoteChainID)`(기존 페어)를 실측해 교체 여부를 정한다.
 * 실측 결과(작성 시점): CROSS 테스트넷 612044 · CROSS 메인넷 612055 · BSC 테스트넷 97 ·
 * BSC 메인넷 56 네 곳 모두 V1 팩토리였다. 미확인 대상: `1000` · `8217`(Kaia) — 배포 전
 * 반드시 재확인할 것(시간이 지나 팩토리가 이미 교체됐을 수 있다).
 *
 * ===================================================================================
 * 참고사항
 * ===================================================================================
 *
 * # CrossMintableERC20V2Code가 Bridge에 설정되면 createToken() 호출 시 CrossMintableERC20V2
 * # 토큰이 CREATE2로 생성되고 자동 등록됨(파생 name/symbol 경로만). 기존에 이미 Code가
 * # 설정되어 있으면 업데이트됨 — 교체 직후 온체인 재조회로 어느 팩토리가 활성인지 확인할 것
 * # (팩토리 교체와 토큰 생성이 경합하면 어느 팩토리가 만들었는지 갈릴 수 있다).
 *
 * # 명시 name/symbol 경로(createMintableERC20)는 브릿지의 createToken을 거치지 않으므로
 * # 자동 registerToken이 일어나지 않는다 — 그 토큰을 브릿지로 이동시키려면
 * # bridge.registerToken(remoteChainID, false, tokenAddress, remoteToken)을 EDITOR_ROLE
 * # 계정으로 별도 호출해야 한다.
 *
 * # 팩토리 소유권 이전 — DEFAULT_ADMIN_ROLE과 ADMIN_ROLE은 별개다. 토큰 관리
 * # 패스쓰루·비콘 업그레이드는 ADMIN_ROLE로 게이트되므로 beginDefaultAdminTransfer만으로는
 * # 새 owner가 아무것도 할 수 없다. 이전은 반드시 4단계이며, 위 세 함수(begin/accept/revoke)가
 * # 각 단계를 구현한다. begin/accept는 이미 완료된 단계면 revert 없이 로그만 남기고 통과하지만,
 * # revokeOldAdmin은 절대 멱등 skip하지 않는다 — 재실행하면 (이미 완료됐어도) revert한다. 이는
 * # 의도된 안전 동작이다: ADMIN_ROLE 부재만으로는 "이미 취소됨"과 "oldOwner를 잘못 넣음"을
 * # 구분할 수 없기 때문이다. 실행 경로는 환경에 따라 셋으로 나뉜다.
 * #
 * # ---- 경로 1: 개발/테스트넷 (EOA로 통제되는 환경 전용) ----
 * #
 * # 래퍼가 vm.startBroadcast(<명시 주소>)를 쓰므로 --private-key가 그 주소를 정확히 유도해야
 * # 한다 — 아니면 조용히 다른 계정으로 서명되는 게 아니라 즉시 실행 에러가 난다.
 * #   beginOwnershipTransfer          -> oldOwner의 키 ($OLD_OWNER_PRIVATE_KEY)
 * #   acceptOwnershipAndGrantAdmin    -> newOwner의 키 ($NEW_OWNER_PRIVATE_KEY)
 * #   revokeOldAdmin                  -> newOwner의 키 ($NEW_OWNER_PRIVATE_KEY)
 * #
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $OLD_OWNER_PRIVATE_KEY \
 * #   --sig "beginOwnershipTransfer(address,address,address)" \
 * #   $CODE $OLD_OWNER $NEW_OWNER \
 * #   --broadcast
 * #
 * # 1단계 직후, 2단계를 제출하기 전에 acceptSchedule을 확인하고 그 시각을 지난 뒤에 제출할 것
 * # — delay가 0이어도 예약 시각(acceptSchedule) 이후여야 하므로, 곧바로 연달아 제출하면 2단계가
 * # revert한다:
 * #
 * # cast call $CODE "pendingDefaultAdmin()(address,uint48)" --rpc-url $RPC_URL
 * #   -> (newOwner, acceptSchedule) 확인, block.timestamp > acceptSchedule 이 된 뒤 2단계 제출
 * #
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $NEW_OWNER_PRIVATE_KEY \
 * #   --sig "acceptOwnershipAndGrantAdmin(address,address)" \
 * #   $CODE $NEW_OWNER \
 * #   --broadcast
 * #
 * # forge script script/CrossMintableERC20V2Code.s.sol:CrossMintableERC20V2CodeScript \
 * #   --rpc-url $RPC_URL --private-key $NEW_OWNER_PRIVATE_KEY \
 * #   --sig "revokeOldAdmin(address,address,address)" \
 * #   $CODE $OLD_OWNER $NEW_OWNER \
 * #   --broadcast
 * #
 * # 마지막 단계(revokeOldAdmin)를 빠뜨리면 구 owner가 토큰 관리·비콘 업그레이드 권한을 계속
 * # 보유한다 — revokeOldAdminStatus의 사전조건(newOwner가 이미 ADMIN_ROLE 보유)이 3단계를
 * # 건너뛴 채 4단계만 실행해 관리자가 아무도 없는 팩토리가 되는 것을 막는다.
 * #
 * # ---- 경로 1 예외: CROSS 체인은 forge script를 쓰지 않는다 ----
 * #
 * # CROSS 테스트넷(비표준 포트 RPC)에서는 forge script --broadcast가 조용히 실패할 수 있다
 * # (시뮬레이션 로그만 찍히고 트랜잭션 미전송). CROSS에서는 위 4단계를 cast send로 직접
 * # 보내고, 매 단계 후 cast call로 온체인 상태를 직접 조회해 확인한다(스크립트 로그를 성공
 * # 근거로 삼지 않는다). cast send는 forge script의 revokeOldAdmin 래퍼(내장 3중 사전조건)를
 * # 거치지 않고 직접 revokeRole을 호출하므로, 아래 $RECORDED_OLD_OWNER 절차와
 * # 4단계 사전조건을 동일하게 수동으로 지켜야 한다:
 * #
 * # $RECORDED_OLD_OWNER 확정 (필수): 이전 시작 **전**에 defaultAdmin()을 조회해 별도
 * # 변수로 기록해 둔다. 4단계 revokeRole에 넣을 주소는 그때그때 다시 손으로 입력하지 않고
 * # 이 기록된 값 하나만 계속 쓴다 — 사전조건 확인·실행·사후검증 전부 동일하다.
 * #
 * # cast call $CODE "defaultAdmin()(address)" --rpc-url $CROSS_RPC_URL
 * #   -> 이전 시작 전에 실행, 반환값을 아래 변수로 기록해 둘 것
 * #
 * # 이전 시작 '전' 에 defaultAdmin() 을 조회해 얻은 주소로 아래 0x... 를 교체한다
 * # export RECORDED_OLD_OWNER=0x0000000000000000000000000000000000000000
 * #
 * # cast send $CODE "beginDefaultAdminTransfer(address)" $NEW_OWNER \
 * #   --rpc-url $CROSS_RPC_URL --private-key $OLD_OWNER_PRIVATE_KEY
 * # cast call  $CODE "pendingDefaultAdmin()(address,uint48)" --rpc-url $CROSS_RPC_URL
 * #   -> (newOwner, acceptSchedule) 확인
 * #
 * # 다음 명령(수락) 제출 전에 반드시 block.timestamp > acceptSchedule 이 될 때까지 기다릴 것
 * # — delay가 0이어도 예약 시각 이후여야 하므로, 곧바로 연달아 보내면 아래 acceptDefaultAdminTransfer
 * # 호출이 revert한다.
 * #
 * # cast send $CODE "acceptDefaultAdminTransfer()" \
 * #   --rpc-url $CROSS_RPC_URL --private-key $NEW_OWNER_PRIVATE_KEY
 * # cast call  $CODE "defaultAdmin()(address)" --rpc-url $CROSS_RPC_URL
 * #
 * # cast send $CODE "grantRole(bytes32,address)" $ADMIN_ROLE $NEW_OWNER \
 * #   --rpc-url $CROSS_RPC_URL --private-key $NEW_OWNER_PRIVATE_KEY
 * # cast call  $CODE "hasRole(bytes32,address)(bool)" $ADMIN_ROLE $NEW_OWNER --rpc-url $CROSS_RPC_URL
 * #
 * # ---- 4단계(revokeRole, CROSS) 제출 전 필수 사전조건 ----
 * #
 * # revokeRole은 역할이 없는 주소에 대해서도 조용히 성공하는 no-op이다 — 잘못된 주소를 넣어도
 * # revert하지 않는다. 4단계 cast send를 보내기 **전에** 반드시 아래 세 가지를 조회해 확인한다
 * # (기록된 $RECORDED_OLD_OWNER 사용):
 * #
 * # cast call $CODE "defaultAdmin()(address)" --rpc-url $CROSS_RPC_URL
 * #   -> $NEW_OWNER 와 같아야 함
 * # cast call $CODE "hasRole(bytes32,address)(bool)" $ADMIN_ROLE $NEW_OWNER --rpc-url $CROSS_RPC_URL
 * #   -> true 여야 함
 * # cast call $CODE "hasRole(bytes32,address)(bool)" $ADMIN_ROLE $RECORDED_OLD_OWNER --rpc-url $CROSS_RPC_URL
 * #   -> true 여야 함 ← 핵심
 * #
 * # 세 번째가 핵심이다. false면 이미 취소됐거나 $RECORDED_OLD_OWNER가 잘못된 주소이므로, 그
 * # 상태에서 4단계를 제출하면 **아무 일도 하지 않으면서(no-op) 성공한 것처럼 보인다.** 반드시
 * # 제출을 중단하고 원인을 확인한다.
 * #
 * # cast send $CODE "revokeRole(bytes32,address)" $ADMIN_ROLE $RECORDED_OLD_OWNER \
 * #   --rpc-url $CROSS_RPC_URL --private-key $NEW_OWNER_PRIVATE_KEY
 * #
 * # 실행 후 사후검증도 calldata에 쓴 값이 아니라 기록된 주소로 한다:
 * #
 * # cast call $CODE "hasRole(bytes32,address)(bool)" $ADMIN_ROLE $RECORDED_OLD_OWNER --rpc-url $CROSS_RPC_URL
 * #   -> false 여야 함 (기록된 $RECORDED_OLD_OWNER 기준)
 * #
 * # ($ADMIN_ROLE 값은 아래 경로 2를 참고.)
 * #
 * # ---- 경로 2: 프로덕션 (멀티시그/타임락) ----
 * #
 * # 팩토리 ADMIN_ROLE 계정은 반드시 멀티시그/타임락이다 — 개인키가 없으므로
 * # 위 forge script/cast send 명령(--private-key)은 프로덕션 실행 경로가 될 수 없다. 대신
 * # 재현 가능한 calldata를 멀티시그·타임락 UI에 그대로 입력한다. target은 네 건 모두 팩토리
 * # 프록시 주소($CODE)다. ADMIN_ROLE = keccak256("ADMIN_ROLE")의 리터럴 값(멀티시그 UI에
 * # bytes32를 직접 입력해야 하므로 미리 계산해 둔다):
 * #
 * #   ADMIN_ROLE = 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775
 * #
 * # $RECORDED_OLD_OWNER 확정 (필수): 이전 시작 **전**의 defaultAdmin() 조회값이나
 * # 거버넌스 기록에서 확정해 별도 변수로 적어 둔다. 4단계 calldata에 넣을 그 값과 **같은
 * # 출처(그때그때 수동 타이핑)에서 다시 가져오지 않는다** — 사전조건 확인·계산·사후검증에
 * # 전부 이 기록된 값을 쓴다.
 * #
 * # cast call $CODE "defaultAdmin()(address)" --rpc-url $RPC_URL
 * #   -> 이전 시작 전에 실행, 반환값을 아래 변수로 기록해 둘 것
 * #
 * # 이전 시작 '전' 에 defaultAdmin() 을 조회해 얻은 주소로 아래 0x... 를 교체한다
 * # export RECORDED_OLD_OWNER=0x0000000000000000000000000000000000000000
 * #
 * # cast calldata "beginDefaultAdminTransfer(address)" $NEW_OWNER
 * #
 * # 1단계(begin) 제출 후, 2단계(accept) calldata를 제출하기 **전에** acceptSchedule을 확인하고
 * # 그 시각을 지난 뒤에 제출할 것 — delay가 0이어도 예약 시각 이후여야 하므로, 곧바로 연달아
 * # 제출하면 2단계가 revert한다:
 * #
 * # cast call $CODE "pendingDefaultAdmin()(address,uint48)" --rpc-url $RPC_URL
 * #   -> (newOwner, acceptSchedule) 확인, block.timestamp > acceptSchedule 이 된 뒤 제출
 * #
 * # cast calldata "acceptDefaultAdminTransfer()"
 * # cast calldata "grantRole(bytes32,address)"  0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $NEW_OWNER
 * #
 * # ---- 4단계(revokeRole) 제출 전 필수 사전조건 ----
 * #
 * # revokeRole은 역할이 없는 주소에 대해서도 조용히 성공하는 no-op이다 — 잘못된 주소를 넣어도
 * # revert하지 않고, 문서화된 사후조건("그 주소가 ADMIN_ROLE을 잃었는지 확인")도 자동으로
 * # 통과해 버린다. 4단계 calldata를 멀티시그/타임락 UI에 제출하기 **전에** 반드시 아래 세
 * # 가지를 조회해 확인한다(기록된 $RECORDED_OLD_OWNER 사용):
 * #
 * # cast call $CODE "defaultAdmin()(address)"                                                                              --rpc-url $RPC_URL
 * #   -> $NEW_OWNER 와 같아야 함
 * # cast call $CODE "hasRole(bytes32,address)(bool)" 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $NEW_OWNER            --rpc-url $RPC_URL
 * #   -> true 여야 함
 * # cast call $CODE "hasRole(bytes32,address)(bool)" 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $RECORDED_OLD_OWNER   --rpc-url $RPC_URL
 * #   -> true 여야 함 ← 핵심
 * #
 * # 세 번째가 핵심이다. false면 이미 취소됐거나 $RECORDED_OLD_OWNER가 잘못된 주소이므로, 그
 * # 상태에서 4단계를 제출하면 **아무 일도 하지 않으면서(no-op) 성공한 것처럼 보인다.** 반드시
 * # 제출을 중단하고 원인을 확인한다.
 * #
 * # cast calldata "revokeRole(bytes32,address)" 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $RECORDED_OLD_OWNER
 * #
 * # 타임락 주의: 타임락은 각 작업을 schedule과 execute로 나눠 자기 지연시간·predecessor 규칙에
 * # 따라 별도 제출해야 한다 — 위 4건을 한 번에 밀어넣을 수 없다. 1단계(begin)는 구 owner
 * # 멀티시그/타임락이, 2~4단계(accept/grant/revoke)는 신 owner 멀티시그/타임락이 제출한다.
 * # 4단계(revoke)의 위 세 가지 사전조건은 **schedule 제출 시점과 execute 제출 시점 양쪽
 * # 모두**에서 반복 확인한다 — 그 사이에 상태가 바뀔 수 있다(예: 다른 경로로 이미
 * # 취소됨). 2단계(accept)의 acceptSchedule 대기도 schedule/execute 각 제출 전에
 * # 마찬가지로 재확인한다.
 * #
 * # 각 트랜잭션 실행 후 온체인 사후조건을 직접 조회해 확인한다(스크립트/멀티시그 UI 로그를
 * # 성공 근거로 삼지 않는다). 4단계 사후조건은 **calldata 인자가 아니라 기록된
 * # $RECORDED_OLD_OWNER**로 확인한다:
 * #
 * # cast call $CODE "defaultAdmin()(address)"                                                                              --rpc-url $RPC_URL
 * # cast call $CODE "pendingDefaultAdmin()(address,uint48)"                                                                --rpc-url $RPC_URL
 * # cast call $CODE "hasRole(bytes32,address)(bool)" 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $NEW_OWNER            --rpc-url $RPC_URL
 * # cast call $CODE "hasRole(bytes32,address)(bool)" 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775 $RECORDED_OLD_OWNER   --rpc-url $RPC_URL
 * #   -> false 여야 함 (기록된 $RECORDED_OLD_OWNER 기준 — calldata에 쓴 값이 아니라)
 */
