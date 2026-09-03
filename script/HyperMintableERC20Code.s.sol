// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../src/BaseBridge.sol";
import {Const} from "../src/lib/Const.sol";
import {HyperMintableERC20Code} from "../src/token/HyperMintableERC20Code.sol";
import {ICrossMintableERC20Code} from "../src/token/ICrossMintableERC20Code.sol";
import {IHyperMintableERC20Code} from "../src/token/IHyperMintableERC20Code.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title HyperMintableERC20CodeScript
 * @notice HyperMintableERC20Code 배포 · Bridge 연결 · HyperCore 링크 슬롯 조작 스크립트
 * @dev CrossMintableERC20V2Code와 달리 생성자가 3인자(`initialOwner, initialTokenAdmin,
 *      initialBridge`)이고, 토큰 CREATE2 주소를 배포 전에 예측할 수 있다. 하단 주석에
 *      링크 런북(§15) 전체를 담아 둔다 — 실제 운영 절차는 그 표를 따를 것.
 *
 * 사용법:
 *   forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "deployHyperMintableERC20Code(address,address,address)" \
 *     <BRIDGE> <FACTORY_ADMIN> <TOKEN_ADMIN> \
 *     --broadcast
 */
contract HyperMintableERC20CodeScript is Script {
    /// @dev keccak256("HyperCore deployer") — the script's OWN literal copy of the slot
    /// constant, never read from the target token's getter (A2): trusting the target to name
    /// its own verification slot defeats the point of an independent check.
    bytes32 internal constant HYPERCORE_DEPLOYER_SLOT =
        0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

    function setUp() public {}

    /**
     * @notice HyperMintableERC20Code 배포 및 Bridge에 설정
     * @dev `tokenAdmin`은 팩토리에서 immutable이라 재배포 없이는 바꿀 수 없다. CREATE2
     *      결정성을 위한 의도적 트레이드오프이므로(A3, AR-4) **반드시 멀티시그·타임락
     *      주소**를 넣을 것 — EOA를 넣으면 키 유출 시 팩토리 재배포 외에 회복 수단이 없다.
     *      이 함수는 내부적으로 `bridge.setCrossMintableERC20Code`를 호출하므로 브로드캐스트
     *      계정은 bridge의 `ADMIN_ROLE`을 보유해야 한다. 이후 `createTokenAndVerify` 단계는
     *      이와 별개로 bridge의 `EDITOR_ROLE`을 요구한다 — 두 role을 다른 계정이 보유한다면
     *      단계별로 브로드캐스트 키를 바꿔야 한다.
     * @param bridge Bridge 컨트랙트 주소 (프록시 주소). ADMIN_ROLE 계정으로 브로드캐스트해야
     *        `setCrossMintableERC20Code` 호출이 성공한다.
     * @param factoryAdmin 팩토리 `ADMIN_ROLE`을 받을 주소 (링크 슬롯 위임 setter 호출 주체)
     * @param tokenAdmin 새로 생성되는 모든 토큰의 초기 owner(defaultAdmin) — LINKER_ROLE 보유
     *        여부와 무관하게 isLinkAuthority로 항상 링크 권한을 가진다. **멀티시그/타임락 권장**
     * @return code 배포된 HyperMintableERC20Code 주소
     */
    function deployHyperMintableERC20Code(BaseBridge bridge, address factoryAdmin, address tokenAdmin)
        public
        returns (address code)
    {
        vm.startBroadcast();
        HyperMintableERC20Code deployed = new HyperMintableERC20Code(factoryAdmin, tokenAdmin, address(bridge));
        code = address(deployed);
        console.log("HyperMintableERC20Code deployed to:", code);
        console.log("  tokenAdmin (immutable):", tokenAdmin);
        console.log("  factoryAdmin (ADMIN_ROLE):", factoryAdmin);

        bridge.setCrossMintableERC20Code(ICrossMintableERC20Code(code));
        console.log("HyperMintableERC20Code set to bridge");
        vm.stopBroadcast();
    }

    /**
     * @notice 배포 전 토큰 CREATE2 주소를 미리 계산한다 (읽기 전용, 트랜잭션 없음)
     * @dev `minter`에는 **반드시 실제로 `createToken`을 호출할 브리지 주소**를 넣을 것.
     *      틀린 주소를 넣어도 revert하지 않고 다른 주소를 조용히 반환한다(의도적 설계 —
     *      `IHyperMintableERC20Code.computeTokenAddress` 참조). 예측이 필요 없는 기본 경로
     *      (먼저 배포 → 실제 주소를 Core에 등록)에서는 이 함수를 건너뛰어도 된다.
     * @param code HyperMintableERC20Code 주소
     * @param remoteChainID 원격 체인 ID
     * @param remoteToken 원격 토큰 주소
     * @param symbol 토큰 심볼
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
        // block(2M)에서는 실패할 수 있으므로 big block으로 전송할 것 (§13, A4).
        tokenAddress = bridge.createToken(remoteChainID, remoteToken, symbol, decimals);
        vm.stopBroadcast();

        console.log("HyperMintableERC20 created at:", tokenAddress);
        console.log("  isHyperMintableERC20:", IHyperMintableERC20Code(code).isHyperMintableERC20(tokenAddress));

        if (expected != address(0)) {
            require(tokenAddress == expected, "createTokenAndVerify: address mismatch vs prediction");
            console.log("  matches pre-computed address:", expected);
        }
    }

    /**
     * @notice 팩토리 경유로 토큰의 HyperCore 링크 슬롯에 finalizer를 기록한다 (런북 단계 5)
     * @dev 직후 반드시 `verifyHyperCoreDeployer`로 원시 슬롯을 확인할 것. 이 시점부터 Core
     *      finalize(단계 7)가 끝날 때까지 LINKER_ROLE 변경·재설정 금지(§12 경쟁 창 방어).
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
     * @notice 팩토리 경유로 토큰의 HyperCore 스팟 자산 인덱스를 기록한다 (런북 단계 8)
     * @dev **Core finalize(단계 7) 성공 후에만** 호출할 것. 인덱스를 바꾸기 전에는 직전
     *      `coreSystemAddress()`의 EVM ERC-20 잔고와 Core spot 잔고가 둘 다 0인지 오프체인에서
     *      먼저 확인한다 — 잘못 나간 자금은 회수 불가(AR-3).
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
     * @dev getter(`hyperCoreDeployer()`)도, 대상 토큰의 슬롯 상수 getter도 신뢰하지 않는다(A2) —
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
 * #   - deployHyperMintableERC20Code는 bridge의 ADMIN_ROLE 계정으로 브로드캐스트해야 함
 * #     (setCrossMintableERC20Code 호출)
 * #   - createTokenAndVerify는 bridge의 EDITOR_ROLE 계정으로 브로드캐스트해야 함
 * #     (createToken → 내부 registerToken 호출). ADMIN_ROLE과 EDITOR_ROLE은 서로 다른
 * #     권한이므로, 두 role을 다른 계정이 보유한다면 단계마다 브로드캐스트 키를 바꿔야 함
 *
 * # factoryAdmin: 팩토리 ADMIN_ROLE을 받을 주소
 * #   - setHyperCoreDeployer / setCoreTokenIndex 위임 호출 주체
 *
 * # tokenAdmin: 새로 생성되는 모든 토큰의 초기 owner(defaultAdmin)
 * #   - LINKER_ROLE 보유 여부와 무관하게 isLinkAuthority로 항상 링크 권한을 가짐
 * #   - 팩토리에서 immutable — 반드시 멀티시그/타임락 주소를 사용할 것
 *
 * # --------------------------------------------------
 * # 배포
 * # --------------------------------------------------
 *
 * # PRIVATE_KEY는 bridge의 ADMIN_ROLE 계정이어야 함 (setCrossMintableERC20Code 호출) —
 * # 아래 "토큰 생성" 단계의 EDITOR_ROLE 계정과는 다른 권한이다.
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "deployHyperMintableERC20Code(address,address,address)" \
 * #   $BRIDGE $FACTORY_ADMIN $TOKEN_ADMIN \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # (선택) 사전 주소 계산
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL \
 * #   --sig "computeTokenAddress(address,uint256,address,string,uint8,address)" \
 * #   $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $BRIDGE
 *
 * # --------------------------------------------------
 * # 토큰 생성 (HyperEVM은 big block 필요 — §13, A4)
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "createTokenAndVerify(address,address,uint256,address,string,uint8,address)" \
 * #   $BRIDGE $CODE $REMOTE_CHAIN_ID $REMOTE_TOKEN $SYMBOL $DECIMALS $EXPECTED_OR_ZERO \
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
 * # (Core finalize 성공 후) Core 토큰 인덱스 설정
 * # --------------------------------------------------
 *
 * # forge script script/HyperMintableERC20Code.s.sol:HyperMintableERC20CodeScript \
 * #   --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
 * #   --sig "setCoreTokenIndex(address,address,uint64)" \
 * #   $CODE $TOKEN $CORE_TOKEN_INDEX \
 * #   --broadcast
 *
 * ===================================================================================
 * 링크 런북 (기본 경로 = 먼저 배포하고 실제 주소를 Core에 등록)
 * ===================================================================================
 *
 * | # | 단계                              | 확인 사항 / 중단 조건                                            |
 * |---|-----------------------------------|-------------------------------------------------------------------|
 * | 0 | 입력 고정                          | remoteChainID, remoteToken, symbol, decimals + 팩토리 주소 + 브리지  |
 * |   |                                    | 주소 + tokenAdmin() 기록                                            |
 * | 1 | (선택) computeTokenAddress         | 사전 조율이 필요할 때만. minter 인자에 반드시 address(bridge)          |
 * | 2 | createToken (EDITOR_ROLE)          | HyperEVM big block 사용. 반환된 실제 주소 기록. 1을 했다면            |
 * |   |                                    | createTokenAndVerify로 일치 확인                                    |
 * | 3 | Core requestEvmContract             | 단계 2의 실제 주소 사용. evmExtraWeiDecimals = EVM decimals −        |
 * |   |                                    | Core weiDecimals, 범위 [-2,18]                                     |
 * | 4 | finalizer 확정                     | 슬롯에 넣을 주소가 Core에서 실제 서명하는 주소(API wallet 포함)와        |
 * |   |                                    | 문자 단위로 동일한지 확인. 온체인 강제 불가(A1, AR-6)                  |
 * | 5 | setHyperCoreDeployer                | 직후 단계 6 검증. 이 시점부터 단계 7 완료까지 팩토리 LINKER_ROLE       |
 * |   |                                    | 변경·재설정 금지. 모니터링 대상 이벤트(아래 표) 감시                  |
 * | 6 | verifyHyperCoreDeployer             | vm.load로 원시 슬롯을 읽어 일치 확인. 불일치면 중단                    |
 * | 7 | Core finalizeEvmContract            | finalizer가 전송. 단계 5 직후 지체 없이. Core 상태와 tx hash 교차확인   |
 * |   | {customStorageSlot}                 |                                                                     |
 * | 8 | setCoreTokenIndex                   | 단계 7 성공 후에만. Core tokenInfo 교차확인. 직전 coreSystemAddress()  |
 * |   |                                    | 의 ① ERC-20 EVM 잔고와 ② Core spot 잔고가 둘 다 0인지 확인, 아니면    |
 * |   |                                    | 중단(AR-3)                                                          |
 * | 9 | 시스템 주소 프로비저닝                | 아래 운영 장부 필드를 채운다                                         |
 *
 * ===================================================================================
 * 모니터링 대상 이벤트 (단계 5~7 구간, §12/AR-2 경쟁 창 방어)
 * ===================================================================================
 *
 * LINKER_ROLE 권한 주체는 정확히 둘이다 — 현재 토큰 defaultAdmin()과, LINKER_ROLE을 보유한
 * factoryLinker(생성 팩토리, immutable). hasRole(LINKER_ROLE, account)는 role 보유를,
 * isLinkAuthority(account)는 실효 권한(슬롯을 실제로 쓸 수 있는지)을 뜻한다 — 둘은 같은 말이
 * 아니다: 현재 defaultAdmin()은 LINKER_ROLE을 보유하지 않아도 권한이 있고, factoryLinker가
 * 아닌 제3자는 LINKER_ROLE을 받아도 권한이 생기지 않는다. 슬롯을 실제로 쓸 수 있는 것은 항상
 * 현재 defaultAdmin()과 factoryLinker(LINKER_ROLE 보유 시)뿐이다.
 *
 * | 이벤트                                                  | 무엇을 뜻하는가                          |
 * |------------------------------------------------------------|---------------------------------------------|
 * | HyperCoreDeployerSet                                        | finalizer 슬롯이 바뀌었다                    |
 * | CoreTokenIndexSet                                           | Core 인덱스(=시스템 주소)가 바뀌었다          |
 * | RoleGranted / RoleRevoked (role = LINKER_ROLE)              | factoryLinker의 링크 권한이 켜지거나 꺼졌다.  |
 * |                                                              | factoryLinker 외 주소에 대한 grant는 실효     |
 * |                                                              | 권한이 아니지만 의도를 드러내므로 함께 본다    |
 * | RoleGranted / RoleRevoked (role = DEFAULT_ADMIN_ROLE)       | admin 교체가 실제로 완료됐다 — 이것이 실효    |
 * |                                                              | 권한이 옮겨간 시점이다                       |
 * | DefaultAdminTransferScheduled / DefaultAdminTransferCanceled | admin 교체 의도/취소 (예고일 뿐, 완료 아님)   |
 *
 * 추가로 단계 5~7 구간 동안 defaultAdmin()과 isLinkAuthority()를 주기적으로 재조회해 대조할
 * 것 — 이벤트를 놓쳐도 실효 권한의 현재 값을 직접 확인할 수 있어야 한다.
 *
 * 운영 장부 필수 필드: token address · factory address · core token index · system address ·
 * evmExtraWeiDecimals · provisioned amount · Core spot balance of system address ·
 * EVM balance of system address · bridge minted (registry 회계) · finalizer ·
 * current defaultAdmin() · factoryLinker() (immutable) · hasRole(LINKER_ROLE, factoryLinker)
 * (팩토리 권한 on/off) · extraneous LINKER_ROLE holders (실효 권한 없음 — 이상 징후로만 기록).
 * CoreTokenIndexSet / HyperCoreDeployerSet / RoleGranted / RoleRevoked / Transfer(to=system
 * address) 이벤트를 인덱싱해 장부와 정기 대조.
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
 * isHyperMintableERC20 allowlist 때문에 **새 팩토리로의 관리 이관은 불가(의도된 비목표)**다.
 * 새 팩토리를 배포해도 예전 팩토리가 만든 토큰은 예전 팩토리로만 setHyperCoreDeployer /
 * setCoreTokenIndex 위임을 받을 수 있다.
 *
 * ===================================================================================
 * 참고사항
 * ===================================================================================
 *
 * # HyperMintableERC20Code가 설정되면 Bridge에서 createToken() 호출 시 HyperMintableERC20
 * # 토큰이 CREATE2로 생성되고 자동 등록됨. 기존에 이미 Code가 설정되어 있으면 업데이트됨.
 *
 * # 팩토리 tokenAdmin은 immutable이므로 새로 생성될 토큰의 초기 owner를 바꾸려면 팩토리를
 * # 재배포해야 한다(AR-4). 이미 배포된 토큰의 admin은 토큰 자체의
 * # beginDefaultAdminTransfer / acceptDefaultAdminTransfer로 이전 가능.
 */
