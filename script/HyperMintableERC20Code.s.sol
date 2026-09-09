// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {Const} from "../src/lib/Const.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {HyperMintableERC20} from "../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {IHyperMintableERC20Code} from "../src/token/IHyperMintableERC20Code.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title HyperMintableERC20CodeScript
 * @notice HyperMintableERC20Code(+토큰 beacon) 배포 · 업그레이드 · Bridge 연결 · HyperCore 링크
 *         슬롯 조작 스크립트
 * @dev `HyperMintableERC20Code`는 `CrossMintableERC20V2Code`를 상속한다:
 *      - 토큰(`HyperMintableERC20`)은 `UpgradeableBeacon` + `BeaconProxy`이며, 그
 *        beacon은 외부(멀티시그) 주소가 아니라 팩토리 자신이 소유한다 — 팩토리의
 *        `initialize` 내부에서 팩토리 자신이 beacon을 만들고 스스로 소유한다(`beacon()`으로
 *        조회). beacon `owner()`가 팩토리 컨트랙트 자신이므로 beacon에 직접 `upgradeTo`를
 *        보내는 경로는 권한이 없어 실패한다 — beacon 업그레이드는 반드시 팩토리의
 *        `upgradeBeacon(address)`(ADMIN_ROLE 게이트)를 거쳐야 한다.
 *      - 팩토리(`HyperMintableERC20Code`)는 UUPS(`ERC1967Proxy`) 단일 인스턴스이며, 별도의
 *        `tokenAdmin` 파라미터를 받지 않는다 — 새로 생성되는 모든 토큰의
 *        `defaultAdmin()`은 팩토리 자신(`address(this)`)이 된다. `initialize`는 상속받은
 *        `CrossMintableERC20V2Code.initialize(initialOwner, initialBridge, tokenImplementation)`
 *        (3-인자) 시그니처를 그대로 쓴다.
 *      두 impl 모두 생성자에서 `_disableInitializers()`를 호출하므로 `ERC1967Proxy` /
 *      `BeaconProxy`의 초기화 데이터를 통해서만 초기화된다(원자적 초기화).
 *      CREATE2 토큰 주소의 initcode에는 **beacon 주소만** 들어가고 그 순간의 로직 impl 주소는
 *      들어가지 않으므로, 토큰 로직을 업그레이드해도 향후 예측 주소는 바뀌지 않는다.
 *      하단 주석에 HyperCore 링크 런북 전체를 담아 둔다 — 실제 운영 절차는 그 표를 따를 것.
 *
 * 사용법:
 *   forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "deployAll(address,address)" \
 *     <BRIDGE> <FACTORY_ADMIN> \
 *     --broadcast
 */
contract HyperMintableERC20CodeScript is Script {
    /// @dev keccak256("HyperCore deployer") — the script's OWN literal copy of the slot
    /// constant, never read from the target token's getter: trusting the target to name
    /// its own verification slot defeats the point of an independent check.
    bytes32 internal constant HYPERCORE_DEPLOYER_SLOT =
        0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

    function setUp() public {}

    // =====================================================================
    // 배포 — 개별 단계
    // =====================================================================

    /**
     * @notice `HyperMintableERC20` 로직(impl) 배포
     * @dev 생성자에서 `_disableInitializers()`를 호출하므로 이 주소를 직접 `initialize`할
     *      수 없다 — 반드시 beacon을 거쳐 `BeaconProxy`로만 초기화된다. 이 impl 주소가 바로
     *      팩토리 `initialize`의 `tokenImplementation` 인자가 된다 — 팩토리가 스스로 beacon을
     *      만들어 소유하므로 이 스크립트가 별도로 beacon을 배포하지 않는다.
     * @return implementation 배포된 impl 주소
     */
    function deployHyperMintableERC20Implementation() public returns (address implementation) {
        vm.startBroadcast();
        HyperMintableERC20 impl = new HyperMintableERC20();
        implementation = address(impl);
        vm.stopBroadcast();

        console.log("HyperMintableERC20 implementation deployed to:", implementation);
    }

    /**
     * @notice `HyperMintableERC20Code` 로직(impl) 배포
     * @dev 생성자에서 `_disableInitializers()`를 호출하므로 이 주소를 직접 `initialize`할
     *      수 없다 — 반드시 `ERC1967Proxy`를 거쳐서만 초기화된다.
     * @return implementation 배포된 impl 주소
     */
    function deployHyperMintableERC20CodeImplementation() public returns (address implementation) {
        vm.startBroadcast();
        HyperMintableERC20Code impl = new HyperMintableERC20Code();
        implementation = address(impl);
        vm.stopBroadcast();

        console.log("HyperMintableERC20Code implementation deployed to:", implementation);
    }

    /**
     * @notice 팩토리 `ERC1967Proxy` 배포 + 원자 초기화(자체 beacon 생성 포함) + Bridge 연결
     * @dev `initialize`는 이 함수 내부에서 자신의 beacon을 새로 만들어 스스로 소유한다 —
     *      사전 배포된 beacon 주소를 인자로 받지 않는다. 새로 생성되는 모든 토큰의
     *      `defaultAdmin()`은 팩토리 자신이 된다(별도의 `tokenAdmin` 파라미터가 없으므로). 이 함수는 내부적으로
     *      `bridge.setCrossMintableERC20Code`를 호출하므로 브로드캐스트 계정은 bridge의
     *      `ADMIN_ROLE`을 보유해야 한다.
     * @param bridge Bridge 컨트랙트 주소 (프록시 주소)
     * @param codeImplementation `deployHyperMintableERC20CodeImplementation`의 반환값
     * @param factoryAdmin 팩토리 `ADMIN_ROLE`을 받을 주소 (beacon 업그레이드 승인 주체이기도
     *        하다 — beacon owner가 곧 팩토리 자신이고, 팩토리는 `ADMIN_ROLE`로만 그 beacon을
     *        움직이므로 **반드시 멀티시그·타임락 주소를 사용할 것**)
     * @param tokenImplementation `deployHyperMintableERC20Implementation`의 반환값 — 팩토리가
     *        스스로 만드는 beacon의 초기 구현이 된다
     * @return code 배포된 팩토리 프록시 주소
     */
    function deployHyperMintableERC20CodeProxy(
        BaseBridge bridge,
        address codeImplementation,
        address factoryAdmin,
        address tokenImplementation
    ) public returns (address code) {
        vm.startBroadcast();
        // `initialize` is inherited by `HyperMintableERC20Code` from `CrossMintableERC20V2Code`
        // unchanged (not redeclared) — referenced via the declaring contract here since
        // `abi.encodeCall`'s magic member lookup only resolves functions declared directly on
        // the named type, not ones merely inherited. The selector only depends on the function
        // signature, not which contract name encodes it, so this dispatches correctly against
        // `codeImplementation`'s actual (inherited) `initialize`.
        ERC1967Proxy proxy = new ERC1967Proxy(
            codeImplementation,
            abi.encodeCall(CrossMintableERC20V2Code.initialize, (factoryAdmin, address(bridge), tokenImplementation))
        );
        code = address(proxy);
        console.log("HyperMintableERC20Code proxy deployed to:", code);
        console.log("  implementation:", codeImplementation);
        console.log("  factoryAdmin (ADMIN_ROLE):", factoryAdmin);
        console.log("  tokenImplementation:", tokenImplementation);
        console.log("  beacon (self-owned by this factory):", HyperMintableERC20Code(code).beacon());

        bridge.setCrossMintableERC20Code(ICrossMintableERC20Code(code));
        console.log("HyperMintableERC20Code set to bridge");
        vm.stopBroadcast();
    }

    /**
     * @notice impl(토큰) + impl(팩토리) + 팩토리 프록시(자체 beacon 생성 포함)를 한 번에
     *         배포하고 Bridge에 연결
     * @dev 위 세 함수를 순서대로 묶은 편의 함수. HyperEVM은 big block이 필요할 수 있다.
     * @param bridge Bridge 컨트랙트 주소 (ADMIN_ROLE 계정으로 브로드캐스트)
     * @param factoryAdmin 팩토리 `ADMIN_ROLE`을 받을 주소. **반드시 멀티시그/타임락**
     *        (beacon owner가 곧 팩토리이므로 이 role이 토큰 로직 전체를 지배한다)
     * @return tokenImplementation 토큰 impl 주소
     * @return codeImplementation 팩토리 impl 주소
     * @return code 팩토리 프록시 주소 (배포 직후 `code.beacon()`으로 자체 생성된 beacon 주소 확인 가능)
     */
    function deployAll(BaseBridge bridge, address factoryAdmin)
        public
        returns (address tokenImplementation, address codeImplementation, address code)
    {
        tokenImplementation = deployHyperMintableERC20Implementation();
        codeImplementation = deployHyperMintableERC20CodeImplementation();
        code = deployHyperMintableERC20CodeProxy(bridge, codeImplementation, factoryAdmin, tokenImplementation);
    }

    // =====================================================================
    // 업그레이드
    // =====================================================================

    /**
     * @notice 토큰 로직을 업그레이드한다 — 이 팩토리가 만든 beacon을 가리키는 모든 토큰에
     *         즉시 반영된다
     * @dev beacon owner가 팩토리 자신이므로, beacon에 직접 `upgradeTo`를 보내는 것이
     *      아니라 팩토리의 `upgradeBeacon(address)`을 거친다 — 팩토리의 `ADMIN_ROLE` 계정으로
     *      브로드캐스트해야 한다. 단계적 롤아웃이 불가하므로 업그레이드 전 전수 테스트가
     *      필수다. CREATE2 initcode에는 beacon 주소만 들어가므로 이 업그레이드는 향후
     *      `computeTokenAddress`/`computeTokenAddressWithName` 예측값에 영향을 주지 않는다.
     * @param code 팩토리 프록시 주소
     * @param newImplementation 새 `HyperMintableERC20` impl 주소
     */
    function upgradeBeacon(address code, address newImplementation) public {
        vm.startBroadcast();
        HyperMintableERC20Code(code).upgradeBeacon(newImplementation);
        vm.stopBroadcast();

        console.log("beacon upgraded via factory:");
        console.log("  factory:", code);
        console.log("  beacon:", HyperMintableERC20Code(code).beacon());
        console.log("  newImplementation:", newImplementation);
    }

    /**
     * @notice 팩토리 로직을 업그레이드한다
     * @dev 팩토리의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 한다(`_authorizeUpgrade`).
     * @param code 팩토리 프록시 주소
     * @param newImplementation 새 `HyperMintableERC20Code` impl 주소
     */
    function upgradeHyperMintableERC20Code(address code, address newImplementation) public {
        vm.startBroadcast();
        HyperMintableERC20Code(code).upgradeToAndCall(newImplementation, bytes(""));
        vm.stopBroadcast();

        console.log("HyperMintableERC20Code upgraded:");
        console.log("  proxy:", code);
        console.log("  newImplementation:", newImplementation);
    }

    // =====================================================================
    // 토큰 생성 — 파생 경로 (심볼에서 name/symbol 파생, 브릿지 경유)
    // =====================================================================

    /**
     * @notice 배포 전 토큰 CREATE2 주소를 미리 계산한다 (읽기 전용, 트랜잭션 없음)
     * @dev `minter`에는 **반드시 실제로 `createToken`을 호출할 브리지 주소**를 넣을 것.
     *      틀린 주소를 넣어도 revert하지 않고 다른 주소를 조용히 반환한다(의도적 설계 —
     *      `IHyperMintableERC20Code.computeTokenAddress` 참조). 예측이 필요 없는 기본 경로
     *      (먼저 배포 → 실제 주소를 Core에 등록)에서는 이 함수를 건너뛰어도 된다.
     * @param code HyperMintableERC20Code 주소
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
            IHyperMintableERC20Code(code).computeTokenAddress(remoteChainID, remoteToken, symbol, decimals, minter);
        console.log("predicted HyperMintableERC20 address:", predicted);
        console.log("  minter used for prediction (must equal the actual createToken caller):", minter);
    }

    /**
     * @notice `createToken`을 실행하고, 필요하면 사전 예측과 실제 주소가 일치하는지 검증한다
     * @dev `expected != 0`이면 사전 예측을 강제하며 불일치 시 revert한다(단계 1→2 연결 확인용).
     *      `expected == 0`이면 예측 없이 그냥 배포하고 실제 주소를 로그로 남긴다(기본 경로).
     *      브로드캐스트 계정은 bridge의 EDITOR_ROLE을 보유해야 한다.
     * @param bridge Bridge 컨트랙트 주소
     * @param code HyperMintableERC20Code 주소
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
        vm.startBroadcast();
        // HyperEVM: 토큰 CREATE2 배포는 big block 가스 한도(30M)가 필요할 수 있다 — small
        // block(2M)에서는 실패할 수 있으므로 big block으로 전송할 것.
        tokenAddress = bridge.createToken(remoteChainID, remoteToken, symbol, decimals);
        vm.stopBroadcast();

        console.log("HyperMintableERC20 created at:", tokenAddress);
        console.log("  isHyperMintableERC20:", IHyperMintableERC20Code(code).isHyperMintableERC20(tokenAddress));

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
     * @param code HyperMintableERC20Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals 토큰 decimals
     * @param minter `createHyperMintableERC20` 호출 시 넘길 minter 인자 (호출자가 아니라
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
        predicted = IHyperMintableERC20Code(code).computeTokenAddressWithName(
            remoteChainID, remoteToken, name_, symbol_, decimals, minter
        );
        console.log("predicted HyperMintableERC20 address (explicit name/symbol):", predicted);
    }

    /**
     * @notice 임의 name/symbol + 명시 minter로 토큰을 생성한다 — 브릿지의 `createToken`을
     *         거치지 않으므로 필요하면 별도로 `bridge.registerToken`을 호출해야 한다
     * @dev 팩토리의 `ADMIN_ROLE` 계정으로 브로드캐스트해야 한다. `expected != 0`이면 사전
     *      예측과 일치하는지 검증한다.
     * @param code HyperMintableERC20Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param name_ ERC20 name
     * @param symbol_ ERC20 symbol
     * @param decimals 토큰 decimals
     * @param minter `MINTER_ROLE`을 받을 주소
     * @param expected 사전 예측 주소. 0이면 검증을 건너뜀
     * @return tokenAddress 실제 생성된 토큰 주소
     */
    function createHyperMintableERC20AndVerify(
        address code,
        uint remoteChainID,
        address remoteToken,
        string memory name_,
        string memory symbol_,
        uint8 decimals,
        address minter,
        address expected
    ) public returns (address tokenAddress) {
        vm.startBroadcast();
        // HyperEVM: 토큰 CREATE2 배포는 big block 가스 한도(30M)가 필요할 수 있다.
        tokenAddress = IHyperMintableERC20Code(code).createHyperMintableERC20(
            remoteChainID, remoteToken, name_, symbol_, decimals, minter
        );
        vm.stopBroadcast();

        console.log("HyperMintableERC20 created at (explicit name/symbol):", tokenAddress);
        console.log("  name:", name_);
        console.log("  symbol:", symbol_);
        console.log("  minter (MINTER_ROLE):", minter);
        console.log(
            "  NOTE: not routed through bridge.createToken() -> call bridge.registerToken(...) separately if this token must be bridgeable"
        );

        if (expected != address(0)) {
            require(tokenAddress == expected, "createHyperMintableERC20AndVerify: address mismatch vs prediction");
            console.log("  matches pre-computed address:", expected);
        }
    }

    // =====================================================================
    // HyperCore 링크 슬롯 설정 및 검증
    // =====================================================================

    /**
     * @notice 팩토리 경유로 토큰의 HyperCore 링크 슬롯에 finalizer를 기록한다 (런북 단계 5)
     * @dev 직후 반드시 `verifyHyperCoreDeployer`로 원시 슬롯을 확인할 것. 이 시점부터 Core
     *      finalize(단계 7)가 끝날 때까지 LINKER_ROLE 변경·재설정 금지 — 경쟁 창(race window) 방어.
     *      Core 인덱스가 이미 확정된 토큰에서는 온체인에서 `HyperCoreLinkAlreadyFinalized`로
     *      거부된다 — 재실행하기 전에 `verifyHyperCoreDeployer`/`coreTokenIndex`로 상태를
     *      먼저 확인할 것.
     * @param code HyperMintableERC20Code 주소 (ADMIN_ROLE 계정으로 브로드캐스트)
     * @param token 대상 HyperMintableERC20 주소 (이 팩토리가 생성한 토큰이어야 함)
     * @param finalizer HyperCore에서 실제 서명할 주소 (EOA / API wallet)
     */
    function setHyperCoreDeployer(address code, address token, address finalizer) public {
        vm.startBroadcast();
        IHyperMintableERC20Code(code).setHyperCoreDeployer(token, finalizer);
        vm.stopBroadcast();

        console.log("setHyperCoreDeployer sent:");
        console.log("  token:", token);
        console.log("  finalizer:", finalizer);
    }

    /**
     * @notice 팩토리 경유로 토큰의 HyperCore 스팟 자산 인덱스를 기록한다 (런북 단계 8) — 성공은
     *         토큰당 정확히 1회뿐이다
     * @dev **Core finalize(단계 7) 성공 후에만** 호출할 것. 온체인에서 Core read precompile로
     *      `tokenInfo(index).evmContract == token`과 decimals 정합성을 검증하므로,
     *      스팟 페어 인덱스(예: 2902)와 토큰 인덱스(예: 2895)를 혼동해 넣으면 `CoreLinkNotFinalized`로
     *      즉시 거부된다 — 실제 운영에서 이런 오입력이 발생한 적이 있다. 인덱스를 바꾸기
     *      전에는 직전 `coreSystemAddress()`의 EVM ERC-20 잔고와 Core spot 잔고가 둘 다 0인지
     *      오프체인에서 먼저 확인한다 — 잘못 나간 자금은 회수 불가하다. 잠금은
     *      오타·절차오류 방지용이며 beacon 업그레이드로는 무력화될 수 있다 — 업그레이드
     *      권한은 반드시 멀티시그/타임락으로 분리해 둘 것.
     * @param code HyperMintableERC20Code 주소 (ADMIN_ROLE 계정으로 브로드캐스트)
     * @param token 대상 HyperMintableERC20 주소
     * @param index HyperCore 스팟 토큰 인덱스
     */
    function setCoreTokenIndex(address code, address token, uint64 index) public {
        vm.startBroadcast();
        IHyperMintableERC20Code(code).setCoreTokenIndex(token, index);
        vm.stopBroadcast();

        console.log("setCoreTokenIndex sent:");
        console.log("  token:", token);
        console.log("  index:", index);
    }

    /**
     * @notice `keccak256("HyperCore deployer")` 슬롯의 원시 값을 직접 읽어 finalizer를 검증한다
     * @dev getter(`hyperCoreDeployer()`)도, 대상 토큰의 슬롯 상수 getter도 신뢰하지 않는다 —
     *      대상이 자기가 검사받을 슬롯을 스스로 고르게 하면 독립 검증이 무너진다. 이 스크립트
     *      자신의 리터럴 상수(`HYPERCORE_DEPLOYER_SLOT`)만 사용한다. Core도 finalize 시점에
     *      원시 슬롯을 읽으므로, 우리도 같은 방식으로 검증해야 실제 판정과 일치한다. 읽기
     *      전용, 트랜잭션 없음. 각 단계가 독립 트랜잭션이므로 재개 지점 판별에도 쓸 수 있다.
     * @param token 대상 HyperMintableERC20 주소
     * @param expected 기대하는 finalizer 주소 (0이면 슬롯이 비어 있어야 함)
     */
    function verifyHyperCoreDeployer(address token, address expected) public view {
        require(token.code.length > 0, "verifyHyperCoreDeployer: token has no code (EOA or wrong address)");

        bytes32 raw = vm.load(token, HYPERCORE_DEPLOYER_SLOT);
        bytes32 expectedWord = bytes32(uint(uint160(expected)));

        console.log("raw HyperCore deployer slot:", vm.toString(raw));
        console.log("  decoded finalizer (log only, not used for comparison):", address(uint160(uint(raw))));
        console.log("  expected finalizer:", expected);

        if (raw != expectedWord) {
            console.log("  expected word:", vm.toString(expectedWord));
            console.log("  actual word:  ", vm.toString(raw));
            revert("verifyHyperCoreDeployer: slot does not match expected finalizer");
        }
        console.log("  matches expected finalizer:", expected);
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
 * #   - deployAll/deployHyperMintableERC20CodeProxy는 bridge의 ADMIN_ROLE 계정으로
 * #     브로드캐스트해야 함 (setCrossMintableERC20Code 호출)
 * #   - createTokenAndVerify는 bridge의 EDITOR_ROLE 계정으로 브로드캐스트해야 함
 * #     (createToken → 내부 registerToken 호출). ADMIN_ROLE과 EDITOR_ROLE은 서로 다른
 * #     권한이므로, 두 role을 다른 계정이 보유한다면 단계마다 브로드캐스트 키를 바꿔야 함
 *
 * # factoryAdmin: 팩토리 ADMIN_ROLE을 받을 주소
 * #   - setHyperCoreDeployer / setCoreTokenIndex 위임 호출, createHyperMintableERC20(명시
 * #     name/symbol 경로), 팩토리 업그레이드 승인, **beacon 업그레이드 승인**(beacon
 * #     owner가 팩토리 자신이므로 이 role이 곧 beacon을 지배한다) 주체. **반드시 멀티시그·
 * #     타임락 주소를 사용할 것**
 *
 * # (별도의 tokenAdmin 파라미터는 없다 — 새로 생성되는 모든 토큰의 defaultAdmin()은 항상
 * #   팩토리 자신이다. 별도의 beaconOwner 파라미터도 없다 — beacon은 팩토리
 * #   initialize 내부에서 스스로 만들어 스스로 소유한다.)
 *
 * # --------------------------------------------------
 * # 배포 (impl 2종 + 팩토리 프록시, 원자 초기화 시 자체 beacon 생성, Bridge 연결까지 한 번에)
 * # --------------------------------------------------
 *
 * # PRIVATE_KEY는 bridge의 ADMIN_ROLE 계정이어야 함 (setCrossMintableERC20Code 호출) —
 * # 아래 "토큰 생성" 단계의 EDITOR_ROLE 계정과는 다른 권한이다.
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "deployAll(address,address)" \
 * #   $BRIDGE $FACTORY_ADMIN \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # (선택) 사전 주소 계산 — 파생 경로
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --sig "computeTokenAddress(address,uint256,address,string,uint8,address)" \
 * #   $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $BRIDGE
 *
 * # --------------------------------------------------
 * # 토큰 생성 — 파생 경로 (HyperEVM은 big block 필요)
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "createTokenAndVerify(address,address,uint256,address,string,uint8,address)" \
 * #   $BRIDGE $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $EXPECTED_OR_ZERO \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 토큰 생성 — 명시 name/symbol 경로 (ADMIN_ROLE, 브릿지 미경유)
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "createHyperMintableERC20AndVerify(address,uint256,address,string,string,uint8,address,address)" \
 * #   $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $NAME $SYMBOL $DECIMALS $MINTER $EXPECTED_OR_ZERO \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # HyperCore 링크 슬롯 설정 및 검증
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "setHyperCoreDeployer(address,address,address)" \
 * #   $CODE $TOKEN $FINALIZER \
 * #   --broadcast
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --sig "verifyHyperCoreDeployer(address,address)" \
 * #   $TOKEN $FINALIZER
 *
 * # --------------------------------------------------
 * # (Core finalize 성공 후) Core 토큰 인덱스 설정 — 토큰당 1회만 성공
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "setCoreTokenIndex(address,address,uint64)" \
 * #   $CODE $TOKEN $CORE_TOKEN_INDEX \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 업그레이드
 * # --------------------------------------------------
 *
 * # 토큰 로직 업그레이드 (factoryAdmin 계정 — beacon owner가 팩토리 자신이므로 팩토리의
 * # ADMIN_ROLE을 거친다. 이 beacon을 쓰는 모든 토큰에 즉시 반영됨)
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "upgradeBeacon(address,address)" \
 * #   $CODE $NEW_TOKEN_IMPLEMENTATION \
 * #   --broadcast
 *
 * # 팩토리 로직 업그레이드 (factoryAdmin 계정, ADMIN_ROLE)
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "upgradeHyperMintableERC20Code(address,address)" \
 * #   $CODE $NEW_CODE_IMPLEMENTATION \
 * #   --broadcast
 *
 * ===================================================================================
 * 링크 런북 (기본 경로 = 먼저 배포하고 실제 주소를 Core에 등록)
 * ===================================================================================
 *
 * | # | 단계                              | 확인 사항 / 중단 조건                                            |
 * |---|-----------------------------------|-------------------------------------------------------------------|
 * | 0 | 입력 고정                          | remoteChainID, remoteToken, symbol, decimals + 팩토리 주소 + 브리지  |
 * |   |                                    | 주소 기록 (tokenAdmin 개념은 없다 — 항상 팩토리 자신)        |
 * | 1 | (선택) computeTokenAddress         | 사전 조율이 필요할 때만. minter 인자에 반드시 address(bridge)          |
 * | 2 | createToken (EDITOR_ROLE)          | HyperEVM big block 사용. 반환된 실제 주소 기록. 1을 했다면            |
 * |   |                                    | createTokenAndVerify로 일치 확인                                    |
 * | 3 | Core requestEvmContract             | 단계 2의 실제 주소 사용. evmExtraWeiDecimals = EVM decimals −        |
 * |   |                                    | Core weiDecimals, 범위 [-2,18]                                     |
 * | 4 | finalizer 확정                     | 슬롯에 넣을 주소가 Core에서 실제 서명하는 주소(API wallet 포함)와        |
 * |   |                                    | 문자 단위로 동일한지 확인. 온체인 강제 불가                  |
 * | 5 | setHyperCoreDeployer                | 직후 단계 6 검증. 이 시점부터 단계 7 완료까지 팩토리 LINKER_ROLE       |
 * |   |                                    | 변경·재설정 금지. 모니터링 대상 이벤트(아래 표) 감시                  |
 * | 6 | verifyHyperCoreDeployer             | vm.load로 원시 슬롯을 읽어 일치 확인. 불일치면 중단                    |
 * | 7 | Core finalizeEvmContract            | finalizer가 전송. 단계 5 직후 지체 없이. Core 상태와 tx hash 교차확인   |
 * |   | {customStorageSlot}                 |                                                                     |
 * | 8 | setCoreTokenIndex                   | 단계 7 성공 후에만. 온체인에서 tokenInfo(index).evmContract ==        |
 * |   |                                    | token && decimals 정합성을 자동 검증 — 스팟 페어 인덱스와    |
 * |   |                                    | 토큰 인덱스를 혼동해 넣으면 즉시 거부된다. 성공은 토큰당 1회뿐.    |
 * |   |                                    | 직전 coreSystemAddress()의 ① ERC-20 EVM 잔고와 ② Core spot 잔고가    |
 * |   |                                    | 둘 다 0인지 확인, 아니면 중단                                  |
 * | 9 | 시스템 주소 프로비저닝                | 아래 운영 장부 필드를 채운다                                         |
 *
 * ===================================================================================
 * 모니터링 대상 이벤트 (단계 5~7 구간, 경쟁 창 방어)
 * ===================================================================================
 *
 * LINKER_ROLE 권한 주체는 정확히 둘이다 — 현재 토큰 defaultAdmin()과, LINKER_ROLE을 보유한
 * factoryLinker(생성 팩토리). 새로 생성된 토큰은 이 둘이 같은 주소(팩토리 자신)로 수렴한다.
 * hasRole(LINKER_ROLE, account)는 role 보유를, isLinkAuthority(account)는 실효 권한(슬롯을
 * 실제로 쓸 수 있는지)을 뜻한다 — 둘은 같은 말이 아니다: 현재 defaultAdmin()은 LINKER_ROLE을
 * 보유하지 않아도 권한이 있고, factoryLinker가 아닌 제3자는 LINKER_ROLE을 받아도 권한이 생기지
 * 않는다. 슬롯을 실제로 쓸 수 있는 것은 항상 현재 defaultAdmin()과 factoryLinker(LINKER_ROLE
 * 보유 시)뿐이다.
 *
 * | 이벤트                                                  | 무엇을 뜻하는가                          |
 * |------------------------------------------------------------|---------------------------------------------|
 * | HyperCoreDeployerSet                                        | finalizer 슬롯이 바뀌었다                    |
 * | CoreTokenIndexSet                                           | Core 인덱스(=시스템 주소)가 확정됐다(1회뿐)   |
 * | RoleGranted / RoleRevoked (role = LINKER_ROLE)              | factoryLinker의 링크 권한이 켜지거나 꺼졌다.  |
 * |                                                              | factoryLinker 외 주소에 대한 grant는 실효     |
 * |                                                              | 권한이 아니지만 의도를 드러내므로 함께 본다    |
 * | RoleGranted / RoleRevoked (role = DEFAULT_ADMIN_ROLE)       | admin 교체가 실제로 완료됐다 — 이것이 실효    |
 * |                                                              | 권한이 옮겨간 시점이다                       |
 * | DefaultAdminTransferScheduled / DefaultAdminTransferCanceled | admin 교체 의도/취소 (예고일 뿐, 완료 아님)   |
 * | Upgraded (beacon / 팩토리 프록시)                            | 로직 impl이 바뀌었다 — 링크 잠금을      |
 * |                                                              | 무력화할 수 있는 유일한 경로이므로 특히 주시   |
 *
 * 추가로 단계 5~7 구간 동안 defaultAdmin()과 isLinkAuthority()를 주기적으로 재조회해 대조할
 * 것 — 이벤트를 놓쳐도 실효 권한의 현재 값을 직접 확인할 수 있어야 한다.
 *
 * 운영 장부 필수 필드: token address · factory address · beacon address(=factory.beacon()) ·
 * core token index · system address · evmExtraWeiDecimals · provisioned amount · Core spot
 * balance of system address · EVM balance of system address · bridge minted (registry 회계) ·
 * finalizer · current defaultAdmin() · factoryLinker() (immutable) · hasRole(LINKER_ROLE,
 * factoryLinker) (팩토리 권한 on/off) · extraneous LINKER_ROLE holders (실효 권한 없음 — 이상
 * 징후로만 기록). CoreTokenIndexSet / HyperCoreDeployerSet / RoleGranted / RoleRevoked /
 * Transfer(to=system address) / Upgraded 이벤트를 인덱싱해 장부와 정기 대조.
 *
 * ===================================================================================
 * 잠금의 한계 — 반드시 숙지할 것
 * ===================================================================================
 *
 * `setCoreTokenIndex`의 1회 고정과 `setHyperCoreDeployer`의 확정 후 잠금은
 * **오타·절차오류를 막는 장치**이지, 권한 탈취에 대한 방어가 아니다. 이 토큰/팩토리는
 * 업그레이더블이므로, 팩토리 `ADMIN_ROLE`(팩토리 자신이 소유한 beacon과 팩토리 자체 로직
 * 둘 다 이 한 role로 움직인다)이 로직을 교체하면 잠금 자체를 무력화하는 새 구현을 올릴 수
 * 있다. 권한 탈취 방어는 오직 **팩토리 `ADMIN_ROLE`을 멀티시그/타임락으로 두는 것**으로만
 * 확보된다 — 이 스크립트의 `factoryAdmin` 파라미터에 EOA를 넣지 말 것.
 *
 * ===================================================================================
 * 팩토리 교체 시 기존 토큰 관리
 * ===================================================================================
 *
 * 기존 토큰의 실효 링크 권한(isLinkAuthority)은 생성 당시 팩토리(factoryLinker, LINKER_ROLE을
 * 보유하는 동안)와 현재 토큰 defaultAdmin()뿐이다 — defaultAdmin은 LINKER_ROLE을 보유하지
 * 않아도 권한이 있다.
 *   ① 기존 팩토리를 폐기하지 않고 유지해 계속 사용 (권장), 또는
 *   ② 토큰 defaultAdmin이 토큰 함수를 직접 호출 (항상 가능한 폴백).
 * isHyperMintableERC20/isCrossMintableERC20 allowlist 때문에 **새 팩토리로의 관리 이관은
 * 불가(의도된 비목표)**다. 새 팩토리를 배포해도 예전 팩토리가 만든 토큰은 예전 팩토리로만
 * setHyperCoreDeployer / setCoreTokenIndex 위임을 받을 수 있다. 다만 토큰 로직 자체는 팩토리와
 * 무관하게 (예전 팩토리 자신을 통한) beacon 업그레이드로 여전히 고칠 수 있다 — beacon은 팩토리
 * 교체와 독립적으로 유지된다.
 *
 * ===================================================================================
 * 참고사항
 * ===================================================================================
 *
 * # HyperMintableERC20Code가 설정되면 Bridge에서 createToken() 호출 시 HyperMintableERC20
 * # 토큰이 CREATE2로 생성되고 자동 등록됨(파생 name/symbol 경로만). 기존에 이미 Code가
 * # 설정되어 있으면 업데이트됨.
 *
 * # 명시 name/symbol 경로(createHyperMintableERC20)는 브릿지의 createToken을 거치지
 * # 않으므로 자동 registerToken이 일어나지 않는다 — 그 토큰을 브릿지로 이동시키려면
 * # bridge.registerToken(remoteChainID, false, tokenAddress, remoteToken)을 EDITOR_ROLE
 * # 계정으로 별도 호출해야 한다.
 *
 * # 팩토리 initialize의 tokenImplementation은 팩토리 스토리지가 아니라 팩토리가 스스로 만든
 * # beacon의 초기 구현으로만 쓰인다 — beacon 자체는 `factory.beacon()`으로 조회 가능하고
 * # 팩토리 재배포 없이 `upgradeBeacon`으로 계속 바꿀 수 있다. 이미 배포된 토큰의 admin은
 * # 토큰 자체의 beginDefaultAdminTransfer / acceptDefaultAdminTransfer로 이전 가능하고, 토큰
 * # 로직 자체는 beacon 업그레이드로 팩토리 재배포 없이 고칠 수 있다.
 */
