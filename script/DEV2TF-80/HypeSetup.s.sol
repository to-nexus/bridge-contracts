// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BaseBridge} from "../../src/BaseBridge.sol";
import {BridgeExecutor} from "../../src/BridgeExecutor.sol";
import {BridgeVerifier} from "../../src/BridgeVerifier.sol";
import {PriceFeed} from "../../src/PriceFeed.sol";
import {IBridgeExecutor} from "../../src/interface/IBridgeExecutor.sol";
import {IBridgeRegistry} from "../../src/interface/IBridgeRegistry.sol";
import {IBridgeVerifier} from "../../src/interface/IBridgeVerifier.sol";
import {Const} from "../../src/lib/Const.sol";
import {HyperMintableERC20} from "../../src/token/HyperMintableERC20.sol";
import {HyperMintableERC20Code} from "../../src/token/HyperMintableERC20Code.sol";
import {ICrossMintableERC20Code} from "../../src/token/ICrossMintableERC20Code.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title HypeSetup — [Mainnet HyperEVM 999] DEV2TF-382 · 383 · 384 · 435(STEP 1·2) 를 tx 파일 하나로
 * @notice `gorge script` 로 미서명 tx 묶음을 만든다 → `gorgef sign` 1회 → `gorge send` 1회.
 *         09-hype-setup.sh 가 환경을 채워 부른다. 직접 부르지 않는다.
 *
 *   infra()                   9·10·11 — PriceFeed · 브릿지 프록시 · 역할 · BridgeVerifier · 팩토리 · 비콘 · executor
 *   token(bridge,factory)     435     — 토큰 생성(ONEchain / ONE) → 가격 → «링크 가능 상태» 확인 → 999 쪽 페어 등록
 *   all()                     위 둘을 한 파일로
 *   verify(bridge,token)      전송 «후» 온체인 재검증 (조회만 — tx 없음)
 *
 * ★ 운영 주소는 이 파일에 없다 (퍼블릭 레포). 전부 환경변수로 받는다 — 정본은 레포 밖 chains.env:
 *     SENDER                    운영 키 A — 모든 tx 의 발신자 · 각 컨트랙트의 DEFAULT_ADMIN
 *     HYPE_BIMPL                0번(DEV2TF-657)에서 선배포 · 보안팀 검수한 BaseBridge 구현
 *     HYPE_BIMPL_CODEHASH       그 검수본의 온체인 keccak (검수 보고서 값)
 *     HYPE_DEV                  수수료 수취 dev (09-hype-setup.sh 가 ONE·BSC 브릿지 dev() 에서 읽어 넣는다)
 *     HYPE_BROLE_*              브릿지 역할 목표 — DEFAULT_ADMIN · ADMIN · VALIDATOR · EDITOR · OPERATOR · VERIFIER · INITIATOR(빈 값)
 *     HYPE_VROLE_*              verifier 역할 목표 — DEFAULT_ADMIN · ADMIN · EDITOR · PRICER
 *     HYPE_PROLE_*              PriceFeed 역할 목표 — ADMIN · PRICER
 *     HYPE_EROLE_*              executor 역할 목표 — DEFAULT_ADMIN · ADMIN
 *   쉼표로 구분한 주소 목록이다. 시뮬레이션 끝의 require 가 «온체인 == 이 목표» 를 집합으로 비교한다.
 *
 * ★ 하나라도 어긋나면 forge 가 실패하고 tx 파일이 «만들어지지 않는다».
 * ★ 주소는 SENDER 의 nonce 로 정해진다. tx 파일을 만든 뒤 전송까지 SENDER 로 다른 tx 를 보내지 않는다.
 * ★ big block 은 이 묶음 «밖» 이다 (L1 액션). 전송 전 ON, 전송 후 OFF.
 * ★ 12번(ONE 허브 등록)은 체인이 달라 여기 없다 → 12-hub999.sh
 */
contract HypeSetup is Script {
    // ── 확정 구성 (DEV2TF-381) — 운영 주소가 아닌 값만 둔다 ──────────────────────────────
    uint internal constant CHAIN_ID = 999;
    uint internal constant HUB = 612055;
    address internal constant NATIVE = 0x0000000000000000000000000000000000000001;
    uint8 internal constant THRESHOLD = 2;
    uint8 internal constant DOLLAR_DECIMALS = 6; // ONE PriceFeed 실측
    uint internal constant FINALIZE_GAS = 300000;
    uint internal constant DEFAULT_TOKEN_PRICE = 10000; // ★ defaultExFeeRate 보다 «앞» 인자
    uint internal constant DEFAULT_EX_FEE_RATE = 0;
    uint internal constant MIN_TOKEN_VALUE = 5000000;
    uint internal constant TIME_WINDOW = 21600;
    uint internal constant MAX_EXTRA = 1024;
    uint internal constant MAX_RETURN_DATA = 1024;
    uint8 internal constant TOKEN_DECIMALS = 18;
    // ★ 생성 후 바꿀 수 없다 — 토큰 주소(CREATE2 initcode)에도 들어간다 (2026-10-26 확정)
    string internal constant TOKEN_NAME = "ONEchain";
    string internal constant TOKEN_SYMBOL = "ONE";
    // ── BSC 메인넷(56) 과 같은 값으로 맞춘다 — 빌드 때 BSC 실측값(HYPE_REF_*)과 같은지 강제한다 ──
    uint internal constant GAS_PRICE_HUB = 0; // 수수료① networkFee — 999→612055 (BSC→612055 와 같음)
    uint internal constant TOKEN_EX_FEE = 0; // 수수료② exFee — 래핑 ONE 토큰별 exFeeRate (BSC ONE 토큰과 같음)
    uint internal constant TOKEN_PRICE = 1000000; // PriceFeed — 래핑 ONE 토큰 가격 1.00 USD (BSC ONE 토큰과 같음)
    uint internal constant HUB_NATIVE_PRICE = 1000000; // PriceFeed — 612055 네이티브 가격 1.00 USD (BSC 와 같음)
    uint internal constant VERIFICATION_DELAY = 86400; // 브릿지 _verificationDelay (slot 1) — initialize 기본값
    uint internal constant POST_CALL_GAS_RESERVE = 0; // 브릿지 _postCallGasReserve (slot 54) — 0 = 기본 150000
    uint internal constant MIN_VALUE_AT_PRICE = 5 ether; // getTokenConfig(612055, ONE).minimumValue — BSC 실측 5e18

    bytes32 internal constant IMPL_SLOT = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
    bytes32 internal constant BEACON_SLOT = 0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50;
    bytes32 internal constant HC_DEPLOYER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;
    bytes4 internal constant INVALID_INITIALIZATION = 0xf92ee8a9;

    struct Cfg {
        address owner;
        address payable dev;
        address bimpl;
        bytes32 bimplCodehash;
        address[] brAdmin;
        address[] brValidator;
        address[] brEditor;
        address[] brOperator;
        address[] brVerifier;
        address[] vAdmin;
        address[] vEditor;
        address[] vPricer;
        address[] pAdmin;
        address[] pPricer;
    }

    struct Infra {
        address priceFeedImpl;
        address priceFeed;
        address bridge;
        address verifier;
        address tokenImpl;
        address factoryImpl;
        address factory;
        address beacon;
        address executor;
    }

    // =====================================================================
    // 진입점
    // =====================================================================

    function infra() external returns (Infra memory d) {
        Cfg memory c = _cfg();
        _requireNonce(c);
        d = _infra(c);
        _verify(c, d.bridge, address(0), "", ""); // 브로드캐스트 밖 — 전송 후 검증과 «같은 코드» 로 확인
    }

    function token(address bridge, address factory) external returns (address tok) {
        Cfg memory c = _cfg();
        _requireNonce(c);
        tok = _token(c, bridge, factory, TOKEN_NAME, TOKEN_SYMBOL);
        _checkToken(c, bridge, factory, tok, TOKEN_NAME, TOKEN_SYMBOL);
        console.log("TOKEN999                =", tok);
    }

    function all() external returns (Infra memory d, address tok) {
        Cfg memory c = _cfg();
        _requireNonce(c);
        d = _infra(c);
        tok = _token(c, d.bridge, d.factory, TOKEN_NAME, TOKEN_SYMBOL);
        _verify(c, d.bridge, tok, TOKEN_NAME, TOKEN_SYMBOL);
    }

    /**
     * @notice 전송 «후» 온체인 재검증 — 읽기 전용. tx 를 만들지 않으므로 gorge 가 아니라 forge script 로 돌린다.
     * @dev 브릿지 주소 하나에서 나머지를 따라가 찾는다. 역할 · 설정 · 연결 + 런타임 바이트코드 대조.
     *      바이트코드는 같은 빌드로 시뮬레이션 안에서 새로 만든 사본과 비교한다. UUPS __self immutable
     *      (PriceFeed · HyperMintableERC20Code) 은 사본 주소를 온체인 주소로 치환한 뒤 비교한다.
     *      ★ executor 는 AccessControl(열거 불가)이라 «선언한 주소의 hasRole» 만 본다 — 추가 보유자는 못 찾는다.
     * @param tok 0 이면 infra 만 보낸 경우다 — 이때는 페어가 «하나도 없어야» 한다
     */
    function verify(address bridge, address tok) external {
        _verify(_cfg(), bridge, tok, TOKEN_NAME, TOKEN_SYMBOL);
    }

    /// @dev forge 가 시뮬레이션한 nonce == 사전 조회한 체인 nonce. gorge 는 nonce 만 다시 매기고
    ///      CREATE 주소가 박힌 calldata 는 고치지 않으므로, 둘이 다르면 주소 참조가 어긋난다.
    function _requireNonce(Cfg memory c) internal view {
        require(vm.getNonce(c.owner) == vm.envUint("HYPE_EXPECT_NONCE"), "simulated nonce != HYPE_EXPECT_NONCE");
    }

    // =====================================================================
    // 환경 → 목표 구성
    // =====================================================================

    function _cfg() internal view returns (Cfg memory c) {
        require(block.chainid == CHAIN_ID, "chainid != 999");
        c.owner = vm.envAddress("SENDER");
        c.dev = payable(vm.envAddress("HYPE_DEV"));
        c.bimpl = vm.envAddress("HYPE_BIMPL");
        c.bimplCodehash = vm.envBytes32("HYPE_BIMPL_CODEHASH");
        c.brAdmin = vm.envAddress("HYPE_BROLE_ADMIN", ",");
        c.brValidator = vm.envAddress("HYPE_BROLE_VALIDATOR", ",");
        c.brEditor = vm.envAddress("HYPE_BROLE_EDITOR", ",");
        c.brOperator = vm.envAddress("HYPE_BROLE_OPERATOR", ",");
        c.brVerifier = vm.envAddress("HYPE_BROLE_VERIFIER", ",");
        c.vAdmin = vm.envAddress("HYPE_VROLE_ADMIN", ",");
        c.vEditor = vm.envAddress("HYPE_VROLE_EDITOR", ",");
        c.vPricer = vm.envAddress("HYPE_VROLE_PRICER", ",");
        c.pAdmin = vm.envAddress("HYPE_PROLE_ADMIN", ",");
        c.pPricer = vm.envAddress("HYPE_PROLE_PRICER", ",");

        // ── 목표 자체의 규칙 (DEV2TF-381 · DEV2TF-80 §4-2) ──
        //    DEFAULT_ADMIN 은 initialize / 생성자가 owner 에게만 준다 — 목표도 owner 하나여야 한다
        _eqSet(
            vm.envAddress("HYPE_BROLE_DEFAULT_ADMIN", ","), _one(c.owner), "env HYPE_BROLE_DEFAULT_ADMIN != [SENDER]"
        );
        _eqSet(
            vm.envAddress("HYPE_VROLE_DEFAULT_ADMIN", ","), _one(c.owner), "env HYPE_VROLE_DEFAULT_ADMIN != [SENDER]"
        );
        _eqSet(
            vm.envAddress("HYPE_EROLE_DEFAULT_ADMIN", ","), _one(c.owner), "env HYPE_EROLE_DEFAULT_ADMIN != [SENDER]"
        );
        //    executor 생성자는 owner 에게 ADMIN 까지 준다 — 그 외 ADMIN 은 이 묶음이 만들지 않는다
        _eqSet(vm.envAddress("HYPE_EROLE_ADMIN", ","), _one(c.owner), "env HYPE_EROLE_ADMIN != [SENDER]");
        //    이 묶음의 setter(setBridgeVerifier 등)는 owner 의 ADMIN 으로 보낸다
        require(_has(c.brAdmin, c.owner), "env HYPE_BROLE_ADMIN must include SENDER");
        //    registerToken 은 EDITOR, 가격 투입은 PriceFeed PRICER 로 보낸다
        require(_has(c.brEditor, c.owner), "env HYPE_BROLE_EDITOR must include SENDER");
        require(_has(c.pPricer, c.owner), "env HYPE_PROLE_PRICER must include SENDER");
        //    INITIATOR 는 두지 않는다 (DEV2TF-381 결정 5) — 변수는 «있되 비어» 있어야 한다
        require(bytes(vm.envString("HYPE_BROLE_INITIATOR")).length == 0, "env HYPE_BROLE_INITIATOR must be empty");
        //    서명 정족수 — 운영 키 금지 · threshold 이상
        require(!_has(c.brValidator, c.owner), "SENDER must not be VALIDATOR");
        require(c.brValidator.length >= THRESHOLD, "VALIDATOR count < threshold");
        //    목표 목록에 중복이 있으면 집합 비교가 «예상 밖 보유자» 를 숨길 수 있다
        _noDup(c.brAdmin, "duplicate in HYPE_BROLE_ADMIN");
        _noDup(c.brValidator, "duplicate in HYPE_BROLE_VALIDATOR");
        _noDup(c.brEditor, "duplicate in HYPE_BROLE_EDITOR");
        _noDup(c.brOperator, "duplicate in HYPE_BROLE_OPERATOR");
        _noDup(c.brVerifier, "duplicate in HYPE_BROLE_VERIFIER");
        _noDup(c.vAdmin, "duplicate in HYPE_VROLE_ADMIN");
        _noDup(c.vEditor, "duplicate in HYPE_VROLE_EDITOR");
        _noDup(c.vPricer, "duplicate in HYPE_VROLE_PRICER");
        _noDup(c.pAdmin, "duplicate in HYPE_PROLE_ADMIN");
        _noDup(c.pPricer, "duplicate in HYPE_PROLE_PRICER");
        require(c.dev != address(0), "HYPE_DEV zero");
        require(c.bimpl.codehash == c.bimplCodehash, "BIMPL codehash != audited (DEV2TF-657)");

        // ── BSC 메인넷 실측값(09-hype-setup.sh 가 빌드 직전에 읽어 넘긴다)과 이 파일의 상수가 같은가 ──
        //    다르면 «BSC 에 맞춘다» 는 결정과 어긋난 것이다 — 상수를 고칠지 사람이 정해야 한다
        _ref("HYPE_REF_FINALIZE_GAS", FINALIZE_GAS);
        _ref("HYPE_REF_DEFAULT_EX_FEE_RATE", DEFAULT_EX_FEE_RATE);
        _ref("HYPE_REF_DEFAULT_TOKEN_PRICE", DEFAULT_TOKEN_PRICE);
        _ref("HYPE_REF_MIN_TOKEN_VALUE", MIN_TOKEN_VALUE);
        _ref("HYPE_REF_VERIFICATION_AMOUNT_THRESHOLD", 0);
        _ref("HYPE_REF_PERIOD_TOTAL_VALUE_THRESHOLD", 0);
        _ref("HYPE_REF_TIME_WINDOW", TIME_WINDOW);
        _ref("HYPE_REF_GAS_PRICE_HUB", GAS_PRICE_HUB);
        _ref("HYPE_REF_TOKEN_EX_FEE", TOKEN_EX_FEE);
        _ref("HYPE_REF_DOLLAR_DECIMALS", DOLLAR_DECIMALS);
        _ref("HYPE_REF_TOKEN_PRICE", TOKEN_PRICE);
        _ref("HYPE_REF_HUB_NATIVE_PRICE", HUB_NATIVE_PRICE);
        _ref("HYPE_REF_MIN_VALUE_AT_PRICE", MIN_VALUE_AT_PRICE);
        _ref("HYPE_REF_VERIFICATION_DELAY", VERIFICATION_DELAY);
        _ref("HYPE_REF_MAX_EXTRA", MAX_EXTRA);
        _ref("HYPE_REF_POST_CALL_GAS_RESERVE", POST_CALL_GAS_RESERVE);
        _ref("HYPE_REF_THRESHOLD", THRESHOLD);
        _ref("HYPE_REF_VALUE_LIMIT_WHITELIST_LEN", 0);
    }

    function _ref(string memory key, uint mine) internal view {
        require(vm.envUint(key) == mine, string.concat("BSC mainnet ", key, " != HypeSetup constant"));
    }

    // =====================================================================
    // 9 · 10 · 11
    // =====================================================================

    function _infra(Cfg memory c) internal returns (Infra memory d) {
        vm.startBroadcast(c.owner);

        // ── 9 STEP 2 — PriceFeed (구현 → 프록시 → ADMIN · PRICER) ──────────────
        //    initialize 는 DEFAULT_ADMIN 만 준다. ADMIN(업그레이드 _authorizeUpgrade) · PRICER 는 따로.
        d.priceFeedImpl = address(new PriceFeed());
        d.priceFeed =
            address(new ERC1967Proxy(d.priceFeedImpl, abi.encodeCall(PriceFeed.initialize, (c.owner, DOLLAR_DECIMALS))));
        _grant(d.priceFeed, Const.ADMIN_ROLE, c.pAdmin, Const.PRICER_ROLE, c.pPricer);
        //    612055 네이티브 가격 — BSC PriceFeed 와 같은 1.00 USD. owner 는 바로 위에서 PRICER 를 받았다.
        //    pricesAt 0 = 블록 시각 (PriceFeed._updateNativeTokenPrice)
        PriceFeed(d.priceFeed).updateNativeTokenPrice(_u(HUB), _u(HUB_NATIVE_PRICE), _u(0));

        // ── 9 STEP 4 — 브릿지 프록시 (구현 = 0번 BIMPL · owner · dev · threshold 2) ─
        d.bridge =
            address(new ERC1967Proxy(c.bimpl, abi.encodeCall(BaseBridge.initialize, (c.owner, c.dev, THRESHOLD))));

        // ── 9 STEP 4-2 — 브릿지 역할 (HYPE_BROLE_*) ─────────────────────────
        //    DEFAULT_ADMIN 은 initialize 가 이미 줬다 (재부여는 무시되므로 넣지 않는다).
        {
            bytes32[] memory r = new bytes32[](
                c.brAdmin.length + c.brValidator.length + c.brEditor.length + c.brOperator.length + c.brVerifier.length
            );
            address[] memory a = new address[](r.length);
            uint k;
            k = _fill(r, a, k, Const.ADMIN_ROLE, c.brAdmin);
            k = _fill(r, a, k, Const.VALIDATOR_ROLE, c.brValidator);
            k = _fill(r, a, k, Const.EDITOR_ROLE, c.brEditor);
            k = _fill(r, a, k, Const.OPERATOR_ROLE, c.brOperator);
            k = _fill(r, a, k, Const.VERIFIER_ROLE, c.brVerifier);
            BaseBridge(payable(d.bridge)).grantRoleBatch(r, a);
        }

        // ── 9 STEP 5 — BridgeVerifier · 역할 · 연결 ─────────────────────────
        //    인자 순서: owner, bridge, priceFeed, finalizeBridgeGas, defaultTokenPrice, defaultExFeeRate,
        //               minimumTokenValue, verificationAmountThreshold, periodTotalValueThreshold, timeWindow
        d.verifier = address(_newVerifier(c.owner, d.bridge, d.priceFeed));
        {
            //    DEFAULT_ADMIN · BRIDGE(=브릿지 단독) 는 생성자가 줬다.
            bytes32[] memory r = new bytes32[](c.vAdmin.length + c.vEditor.length + c.vPricer.length);
            address[] memory a = new address[](r.length);
            uint k;
            k = _fill(r, a, k, Const.ADMIN_ROLE, c.vAdmin);
            k = _fill(r, a, k, Const.EDITOR_ROLE, c.vEditor);
            k = _fill(r, a, k, Const.PRICER_ROLE, c.vPricer);
            BridgeVerifier(d.verifier).grantRoleBatch(r, a);
        }
        BaseBridge(payable(d.bridge)).setBridgeVerifier(IBridgeVerifier(d.verifier));

        // ── 10 STEP 1·2 — 토큰 구현 · 팩토리(비콘을 스스로 생성 · 소유) · 브릿지에 등록 ─
        d.tokenImpl = address(new HyperMintableERC20());
        d.factoryImpl = address(new HyperMintableERC20Code());
        d.factory = address(
            new ERC1967Proxy(
                d.factoryImpl,
                // initialize(initialOwner, initialBridge, tokenImplementation) — 인자 3개 (DEV2TF-383 정정)
                abi.encodeWithSignature("initialize(address,address,address)", c.owner, d.bridge, d.tokenImpl)
            )
        );
        BaseBridge(payable(d.bridge)).setCrossMintableERC20Code(ICrossMintableERC20Code(d.factory));

        // ── 11 STEP 1·2 — executor (owner, executor=브릿지) · 브릿지에 연결 ─────
        //    STEP 4 gasPrice 999→612055 = 0 은 미설정 기본값과 같아 보내지 않는다 (verify 에서 0 확인).
        //    maxExtraDataLength 1024 는 새 프록시 initialize 의 기본값이다 (verify 에서 확인).
        d.executor = address(new BridgeExecutor(c.owner, d.bridge));
        BaseBridge(payable(d.bridge)).setBridgeExecutor(IBridgeExecutor(d.executor));

        vm.stopBroadcast();
    }

    function _newVerifier(address owner, address bridge, address priceFeed) internal returns (BridgeVerifier) {
        return new BridgeVerifier(
            owner,
            bridge,
            priceFeed,
            FINALIZE_GAS,
            DEFAULT_TOKEN_PRICE,
            DEFAULT_EX_FEE_RATE,
            MIN_TOKEN_VALUE,
            0, // verificationAmountThreshold
            0, // periodTotalValueThreshold
            TIME_WINDOW
        );
    }

    // =====================================================================
    // 435 STEP 1·2 — 토큰 생성 → 링크 가능 상태 확인 → 999 쪽 페어 등록
    // =====================================================================

    function _token(Cfg memory c, address bridge, address factory, string memory name_, string memory symbol_)
        internal
        returns (address tok)
    {
        HyperMintableERC20Code f = HyperMintableERC20Code(factory);
        require(address(BaseBridge(payable(bridge)).crossMintableERC20Code()) == factory, "bridge factory mismatch");
        require(f.tokenForPair(HUB, NATIVE) == address(0), "pair token already created");
        require(bytes(name_).length != 0 && bytes(symbol_).length != 0, "empty name/symbol");
        address predicted = f.computeTokenAddressWithName(HUB, NATIVE, name_, symbol_, TOKEN_DECIMALS, bridge);

        // ── STEP 1 — 팩토리 경로 (ADMIN). minter = 브릿지. 브릿지 createToken 은 페어까지 자동 등록하므로 쓰지 않는다.
        vm.startBroadcast(c.owner);
        tok = f.createHyperMintableERC20(HUB, NATIVE, name_, symbol_, TOKEN_DECIMALS, bridge);
        vm.stopBroadcast();

        // ── 생성 확인 + ★ 링크 가능 상태 (DEV2TF-435 §3-4) — 페어 등록 «전» ──
        HyperMintableERC20 t = HyperMintableERC20(tok);
        require(tok == predicted, "token != computeTokenAddressWithName");
        require(keccak256(bytes(t.name())) == keccak256(bytes(name_)), "name");
        require(keccak256(bytes(t.symbol())) == keccak256(bytes(symbol_)), "symbol");
        require(t.totalSupply() == 0, "totalSupply != 0");
        _checkTokenLinkReady(c, bridge, factory, tok);

        // ── 토큰 가격 — BSC PriceFeed 의 ONE 토큰과 같은 1.00 USD. 없으면 최소금액이 defaultTokenPrice 기준 500 ONE 이 된다.
        //    페어 등록(개방) «전» 에 넣는다.
        address pf = address(BridgeVerifier(address(BaseBridge(payable(bridge)).bridgeVerifier())).priceFeed());
        vm.startBroadcast(c.owner);
        PriceFeed(pf).updatePrice(_a(tok), _u(TOKEN_PRICE), _u(0));
        vm.stopBroadcast();

        // ── STEP 2 — 999 쪽 페어 (EDITOR). isOrigin 은 «두 번째» · 999 는 false (ONE 쪽은 12번에서 true)
        vm.startBroadcast(c.owner);
        BaseBridge(payable(bridge)).registerToken(HUB, false, tok, NATIVE);
        vm.stopBroadcast();

        _checkPair(bridge, tok);
    }

    function _checkToken(
        Cfg memory c,
        address bridge,
        address factory,
        address tok,
        string memory name_,
        string memory symbol_
    ) internal {
        require(bytes(name_).length != 0 && bytes(symbol_).length != 0, "verify: name/symbol required with token");
        require(keccak256(bytes(HyperMintableERC20(tok).name())) == keccak256(bytes(name_)), "token name");
        require(keccak256(bytes(HyperMintableERC20(tok).symbol())) == keccak256(bytes(symbol_)), "token symbol");
        _checkTokenLinkReady(c, bridge, factory, tok);
        _checkPair(bridge, tok);
        // ── BSC 와 같은 가격 · 수수료 — 토큰 단위 ──
        BridgeVerifier v = BridgeVerifier(address(BaseBridge(payable(bridge)).bridgeVerifier()));
        (bool ex, uint price,) = PriceFeed(address(v.priceFeed())).getTokenPriceInDollars(tok);
        require(ex && price == TOKEN_PRICE, "PF token price != BSC (1.00 USD)");
        require(uint(vm.load(address(v), keccak256(abi.encode(tok, uint(9))))) == TOKEN_EX_FEE, "V exFeeRate[token]");
        (uint minV, uint netFee, uint exFee) = v.getTokenConfig(HUB, IERC20(tok));
        require(minV == MIN_VALUE_AT_PRICE && netFee == 0 && exFee == 0, "V getTokenConfig(612055, token) != BSC");
        // 이 팩토리 비콘을 가리키는 정품 BeaconProxy 여야 한다 (_beacon immutable 까지 같아야 일치)
        _sameCode(tok, address(new BeaconProxy(HyperMintableERC20Code(factory).beacon(), "")), "token proxy code");
    }

    /// @dev 나중에 HyperCore 링크를 «할 수 있는» 구조인지 (DEV2TF-1189 §0). 여기서 틀리면 링크할 방법이 없다.
    function _checkTokenLinkReady(Cfg memory c, address bridge, address factory, address tok) internal view {
        HyperMintableERC20Code f = HyperMintableERC20Code(factory);
        HyperMintableERC20 t = HyperMintableERC20(tok);
        require(f.tokenForPair(HUB, NATIVE) == tok, "tokenForPair");
        require(f.isHyperMintableERC20(tok), "isHyperMintableERC20");
        require(address(uint160(uint(vm.load(tok, BEACON_SLOT)))) == f.beacon(), "token beacon slot");
        require(t.decimals() == TOKEN_DECIMALS, "decimals");
        require(t.hasRole(Const.MINTER_ROLE, bridge), "MINTER != bridge");
        require(t.defaultAdmin() == factory, "token defaultAdmin != factory");
        require(t.factoryLinker() == factory, "factoryLinker != factory");
        require(t.hasRole(Const.LINKER_ROLE, factory), "factory LINKER_ROLE");
        require(t.isLinkAuthority(factory), "factory not link authority");
        require(!t.isLinkAuthority(c.owner), "SENDER must not be token link authority");
        require(!t.isCoreTokenIndexSet(), "coreTokenIndex already set");
        require(vm.load(tok, HC_DEPLOYER_SLOT) == bytes32(0), "finalizer slot not empty");
        require(t.HYPERCORE_DEPLOYER_SLOT() == HC_DEPLOYER_SLOT, "HYPERCORE_DEPLOYER_SLOT");
        // CBU-10 — 링크 전 coreTokenIndex() 는 0x935dcbfa 로 revert 해야 한다
        (bool ok, bytes memory err) = tok.staticcall(abi.encodeCall(HyperMintableERC20.coreTokenIndex, ()));
        require(
            !ok && err.length >= 4 && bytes4(err) == HyperMintableERC20.HyperMintableERC20CoreTokenIndexNotSet.selector,
            "CBU-10: coreTokenIndex must revert 0x935dcbfa before link"
        );
    }

    function _checkPair(address bridge, address tok) internal view {
        IBridgeRegistry.TokenPair memory p = BaseBridge(payable(bridge)).getTokenPair(HUB, tok);
        require(p.localToken == tok && p.remoteToken == NATIVE, "pair tokens");
        require(!p.isOrigin, "999 pair must be isOrigin=false");
        require(!p.paused, "pair paused");
        require(p.pendingAmount == 0, "pair pendingAmount");
        uint[] memory ids = BaseBridge(payable(bridge)).allChainIDs();
        require(ids.length == 1 && ids[0] == HUB, "allChainIDs != [612055]");
    }

    // =====================================================================
    // 검증 (시뮬레이션 끝 · 전송 후 verify 공용)
    // =====================================================================

    function _verify(Cfg memory c, address bridge, address tok, string memory name_, string memory symbol_) internal {
        BaseBridge br = BaseBridge(payable(bridge));
        Infra memory d;
        d.bridge = bridge;
        d.verifier = address(br.bridgeVerifier());
        d.executor = address(br.bridgeExecutor());
        d.factory = address(br.crossMintableERC20Code());
        d.priceFeed = address(BridgeVerifier(d.verifier).priceFeed());
        d.priceFeedImpl = _impl(d.priceFeed);
        d.factoryImpl = _impl(d.factory);
        d.beacon = HyperMintableERC20Code(d.factory).beacon();
        d.tokenImpl = UpgradeableBeacon(d.beacon).implementation();

        _checkInfra(c, d);

        // ── 런타임 바이트코드 — 이 빌드(FOUNDRY_PROFILE=deploy)와 같은가 ──
        _sameCode(d.priceFeedImpl, address(new PriceFeed()), "PriceFeed impl code");
        _sameCode(d.verifier, address(_newVerifier(c.owner, bridge, d.priceFeed)), "verifier code");
        _sameCode(d.tokenImpl, address(new HyperMintableERC20()), "token impl code");
        _sameCode(d.factoryImpl, address(new HyperMintableERC20Code()), "factory impl code");
        _sameCode(d.executor, address(new BridgeExecutor(c.owner, bridge)), "executor code");
        address anyImpl = address(new PriceFeed());
        _sameCode(d.priceFeed, address(new ERC1967Proxy(anyImpl, "")), "PriceFeed proxy code");
        _sameCode(bridge, address(new ERC1967Proxy(anyImpl, "")), "bridge proxy code");
        _sameCode(d.factory, address(new ERC1967Proxy(anyImpl, "")), "factory proxy code");
        _sameCode(d.beacon, address(new UpgradeableBeacon(anyImpl, c.owner)), "beacon code");

        if (tok != address(0)) {
            _checkToken(c, bridge, d.factory, tok, name_, symbol_);
        } else {
            // infra 만 보낸 상태 — 페어 · 체인 등록이 «하나도» 없어야 한다 (11번 최종 점검)
            require(br.allChainIDs().length == 0, "infra-only: allChainIDs must be empty");
            require(
                HyperMintableERC20Code(d.factory).tokenForPair(HUB, NATIVE) == address(0), "infra-only: token exists"
            );
        }
        _report(d);
        if (tok != address(0)) console.log("TOKEN999                =", tok);
        console.log("VERIFY OK");
    }

    function _checkInfra(Cfg memory c, Infra memory d) internal {
        // ── PriceFeed ──
        PriceFeed pf = PriceFeed(d.priceFeed);
        require(pf.dollarDecimals() == DOLLAR_DECIMALS, "PF dollarDecimals");
        (bool nex, uint nprice,) = pf.getNativeTokenPrice(HUB);
        require(nex && nprice == HUB_NATIVE_PRICE, "PF native price[612055] != BSC (1.00 USD)");
        _eqSet(pf.getRoleMembers(bytes32(0)), _one(c.owner), "PF DEFAULT_ADMIN");
        _eqSet(pf.getRoleMembers(Const.ADMIN_ROLE), c.pAdmin, "PF ADMIN");
        _eqSet(pf.getRoleMembers(Const.PRICER_ROLE), c.pPricer, "PF PRICER");
        _sealed(d.priceFeed, abi.encodeCall(PriceFeed.initialize, (c.owner, DOLLAR_DECIMALS)), "PF proxy init sealed");
        _sealed(d.priceFeedImpl, abi.encodeCall(PriceFeed.initialize, (c.owner, DOLLAR_DECIMALS)), "PF impl sealed");

        // ── 브릿지 ──
        BaseBridge br = BaseBridge(payable(d.bridge));
        require(_impl(d.bridge) == c.bimpl, "bridge impl != BIMPL");
        require(c.bimpl.codehash == c.bimplCodehash, "BIMPL codehash != audited (DEV2TF-657)");
        require(br.dev() == c.dev, "bridge dev");
        require(br.threshold() == THRESHOLD, "bridge threshold");
        require(br.maxExtraDataLength() == MAX_EXTRA, "bridge maxExtraDataLength");
        require(uint(vm.load(d.bridge, bytes32(uint(1)))) == VERIFICATION_DELAY, "bridge verificationDelay (slot 1)");
        require(
            uint(vm.load(d.bridge, bytes32(uint(54)))) == POST_CALL_GAS_RESERVE, "bridge postCallGasReserve (slot 54)"
        );
        require(!br.paused(), "bridge paused");
        require(address(br.bridgeVerifier()) == d.verifier, "bridge verifier");
        require(address(br.bridgeExecutor()) == d.executor, "bridge executor");
        require(address(uint160(uint(vm.load(d.bridge, bytes32(0))))) == d.factory, "bridge slot0 != factory");
        _eqSet(br.getRoleMembers(bytes32(0)), _one(c.owner), "BR DEFAULT_ADMIN");
        _eqSet(br.getRoleMembers(Const.ADMIN_ROLE), c.brAdmin, "BR ADMIN");
        _eqSet(br.getRoleMembers(Const.VALIDATOR_ROLE), c.brValidator, "BR VALIDATOR");
        _eqSet(br.getRoleMembers(Const.EDITOR_ROLE), c.brEditor, "BR EDITOR");
        _eqSet(br.getRoleMembers(Const.OPERATOR_ROLE), c.brOperator, "BR OPERATOR");
        _eqSet(br.getRoleMembers(Const.VERIFIER_ROLE), c.brVerifier, "BR VERIFIER");
        require(br.getRoleMembers(Const.INITIATOR_ROLE).length == 0, "BR INITIATOR must be empty");
        require(!br.hasRole(Const.VALIDATOR_ROLE, c.owner), "SENDER must not be VALIDATOR");
        _sealed(d.bridge, abi.encodeCall(BaseBridge.initialize, (c.owner, c.dev, THRESHOLD)), "bridge init sealed");

        // ── BridgeVerifier (slot 1~8 = DEV2TF-381 확정값) ──
        BridgeVerifier v = BridgeVerifier(d.verifier);
        require(address(v.priceFeed()) == d.priceFeed, "V priceFeed");
        require(uint(vm.load(d.verifier, bytes32(uint(2)))) == FINALIZE_GAS, "V finalizeBridgeGas");
        require(uint(vm.load(d.verifier, bytes32(uint(3)))) == DEFAULT_EX_FEE_RATE, "V defaultExFeeRate");
        require(uint(vm.load(d.verifier, bytes32(uint(4)))) == DEFAULT_TOKEN_PRICE, "V defaultTokenPrice");
        require(uint(vm.load(d.verifier, bytes32(uint(5)))) == MIN_TOKEN_VALUE, "V minimumTokenValue");
        require(uint(vm.load(d.verifier, bytes32(uint(6)))) == 0, "V verificationAmountThreshold");
        require(uint(vm.load(d.verifier, bytes32(uint(7)))) == 0, "V periodTotalValueThreshold");
        require(uint(vm.load(d.verifier, bytes32(uint(8)))) == TIME_WINDOW, "V timeWindow");
        require(v.getGasPrice(HUB) == GAS_PRICE_HUB, "V gasPrice[612055] != BSC");
        require(v.getValueLimitWhitelistLength() == 0, "V valueLimitWhitelist not empty");
        _eqSet(v.getRoleMembers(bytes32(0)), _one(c.owner), "V DEFAULT_ADMIN");
        _eqSet(v.getRoleMembers(Const.ADMIN_ROLE), c.vAdmin, "V ADMIN");
        _eqSet(v.getRoleMembers(Const.EDITOR_ROLE), c.vEditor, "V EDITOR");
        _eqSet(v.getRoleMembers(Const.PRICER_ROLE), c.vPricer, "V PRICER");
        _eqSet(v.getRoleMembers(Const.BRIDGE_ROLE), _one(d.bridge), "V BRIDGE (bridge only)");

        // ── 팩토리 · 비콘 ──
        HyperMintableERC20Code f = HyperMintableERC20Code(d.factory);
        require(UpgradeableBeacon(d.beacon).implementation() == d.tokenImpl, "beacon impl != tokenImpl");
        require(UpgradeableBeacon(d.beacon).owner() == d.factory, "beacon owner != factory (must not be EOA)");
        require(f.defaultAdmin() == c.owner, "factory defaultAdmin");
        require(f.hasRole(Const.ADMIN_ROLE, c.owner), "factory ADMIN");
        require(f.hasRole(Const.BRIDGE_ROLE, d.bridge), "factory BRIDGE_ROLE != bridge (initialBridge)");
        _sealed(
            d.factory,
            abi.encodeWithSignature("initialize(address,address,address)", c.owner, d.bridge, d.tokenImpl),
            "factory proxy init sealed"
        );
        _sealed(
            d.factoryImpl,
            abi.encodeWithSignature("initialize(address,address,address)", c.owner, c.owner, c.owner),
            "factory impl sealed"
        );

        // ── executor (AccessControl — 열거 불가, 선언 주소만) ──
        BridgeExecutor e = BridgeExecutor(payable(d.executor));
        require(e.hasRole(bytes32(0), c.owner), "E DEFAULT_ADMIN");
        require(e.hasRole(Const.ADMIN_ROLE, c.owner), "E ADMIN");
        require(e.hasRole(Const.EXECUTOR_ROLE, d.bridge), "E EXECUTOR != bridge");
        require(!e.hasRole(Const.EXECUTOR_ROLE, c.owner), "SENDER must not be EXECUTOR");
        require(e.maxReturnDataSize() == MAX_RETURN_DATA, "E maxReturnDataSize");
    }

    // =====================================================================
    // 보조
    // =====================================================================

    function _grant(address target, bytes32 r1, address[] memory a1, bytes32 r2, address[] memory a2) internal {
        bytes32[] memory r = new bytes32[](a1.length + a2.length);
        address[] memory a = new address[](r.length);
        uint k = _fill(r, a, 0, r1, a1);
        _fill(r, a, k, r2, a2);
        PriceFeed(target).grantRoleBatch(r, a);
    }

    function _fill(bytes32[] memory r, address[] memory a, uint k, bytes32 role, address[] memory who)
        internal
        pure
        returns (uint)
    {
        for (uint i = 0; i < who.length; ++i) {
            require(who[i] != address(0), "zero address in role list");
            (r[k], a[k]) = (role, who[i]);
            ++k;
        }
        return k;
    }

    function _impl(address proxy) internal view returns (address) {
        return address(uint160(uint(vm.load(proxy, IMPL_SLOT))));
    }

    /// @dev 재초기화가 InvalidInitialization 으로 막히는지 — 브로드캐스트 밖에서만 부른다
    function _sealed(address target, bytes memory data, string memory what) internal {
        (bool ok, bytes memory ret) = target.call(data);
        require(!ok && ret.length >= 4 && bytes4(ret) == INVALID_INITIALIZATION, what);
    }

    /// @dev onchain 런타임 == fresh(이 빌드로 방금 만든 사본) 런타임. fresh 코드 안의 «자기 주소»
    ///      (UUPS __self immutable)만 onchain 주소로 바꿔 넣고 전체 keccak 을 비교한다.
    function _sameCode(address onchain, address fresh, string memory what) internal view {
        bytes memory a = onchain.code;
        bytes memory b = fresh.code;
        require(a.length > 0 && a.length == b.length, what);
        bytes20 f = bytes20(fresh);
        bytes20 o = bytes20(onchain);
        for (uint i = 0; i + 20 <= b.length; ++i) {
            bytes20 w;
            assembly {
                // 상위 20바이트만 남긴다 (bytes20 은 좌측 정렬 — 하위 12바이트를 지워 둬야 비교가 정확하다)
                w := and(mload(add(add(b, 32), i)), not(sub(shl(96, 1), 1)))
            }
            if (w == f) {
                assembly {
                    let p := add(add(b, 32), i)
                    // 앞 20바이트 = onchain 주소, 뒤 12바이트 = 원래 값 그대로
                    mstore(p, or(o, and(mload(p), sub(shl(96, 1), 1))))
                }
                i += 19;
            }
        }
        require(keccak256(a) == keccak256(b), what);
    }

    function _eqSet(address[] memory got, address[] memory want, string memory what) internal pure {
        require(got.length == want.length, what);
        for (uint i = 0; i < want.length; ++i) {
            require(_has(got, want[i]), what);
        }
    }

    function _has(address[] memory xs, address x) internal pure returns (bool) {
        for (uint i = 0; i < xs.length; ++i) {
            if (xs[i] == x) return true;
        }
        return false;
    }

    function _noDup(address[] memory xs, string memory what) internal pure {
        for (uint i = 0; i < xs.length; ++i) {
            for (uint j = i + 1; j < xs.length; ++j) {
                require(xs[i] != xs[j], what);
            }
        }
    }

    function _u(uint x) internal pure returns (uint[] memory r) {
        r = new uint[](1);
        r[0] = x;
    }

    function _a(address x) internal pure returns (address[] memory r) {
        return _one(x);
    }

    function _one(address x) internal pure returns (address[] memory r) {
        r = new address[](1);
        r[0] = x;
    }

    function _report(Infra memory d) internal pure {
        console.log("");
        console.log("== 9/10/11 addresses (chain 999 - compare with receipts after send) ==");
        console.log("PriceFeed impl          =", d.priceFeedImpl);
        console.log("PriceFeed (PF999)       =", d.priceFeed);
        console.log("Bridge proxy (P999)     =", d.bridge);
        console.log("BridgeVerifier (V999)   =", d.verifier);
        console.log("Token impl (TOKIMPL)    =", d.tokenImpl);
        console.log("Factory impl (FACIMPL)  =", d.factoryImpl);
        console.log("Factory proxy (FACTORY) =", d.factory);
        console.log("Beacon                  =", d.beacon);
        console.log("BridgeExecutor (E999)   =", d.executor);
        console.log("");
        console.log("== chains.env ==");
        console.log(string.concat("HYPE_BRIDGE=", vm.toString(d.bridge)));
        console.log(string.concat("HYPE_VROLE_BRIDGE=", vm.toString(d.bridge)));
        console.log(string.concat("HYPE_EROLE_EXECUTOR=", vm.toString(d.bridge)));
        console.log(string.concat("HYPE_TOKEN_FACTORY=", vm.toString(d.factory)));
    }
}
