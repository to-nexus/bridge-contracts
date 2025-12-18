// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import {BaseBridge} from "../src/BaseBridge.sol";
import {BridgeVerifier} from "../src/BridgeVerifier.sol";
import {CrossMintableERC20V2Code} from "../src/token/CrossMintableERC20V2Code.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title BridgeScript
 * @notice Bridge 전체 설정 스크립트 (BaseBridge + Verifier + CrossMintableERC20V2Code)
 * @dev 환경변수 기반으로 Bridge 시스템 전체를 배포하고 설정합니다.
 *      - BaseBridge 프록시 배포 및 초기화
 *      - CrossMintableERC20V2Code 배포 및 연결
 *      - BridgeVerifier 배포 및 연결
 *      - 역할 일괄 부여
 *
 * 제공 기능:
 * - setupBaseBridge(): 전체 Bridge 시스템 설정
 * - deployBaseBridgeProxy(): BaseBridge 프록시만 배포
 *
 * 사용법:
 *   forge script script/Bridge.s.sol:BridgeScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "setupBaseBridge()" \
 *     --broadcast
 */
contract BridgeScript is Script {
    // ============ 공통 환경변수 키 ============
    string OWNER = "OWNER";
    string PRICE_FEED = "PRICE_FEED";
    string IMPLEMENTATION = "IMPLEMENTATION";

    // ============ Bridge 환경변수 키 ============
    string BRIDGE_DEV = "DEV";
    string BRIDGE_THRESHOLD = "THRESHOLD";
    string BRIDGE_CROSS_CHAIN_ID = "CROSS_CHAIN_ID";
    string BRIDGE_BSC_CHAIN_ID = "BSC_CHAIN_ID";
    string BRIDGE_CROSS = "CROSS";
    string BRIDGE_CROSS_INITIAL_SUPPLY = "CROSS_INITIAL_SUPPLY";
    string BRIDGE_ROLES = "BRIDGE_ROLES";
    string BRIDGE_ROLE_MEMBERS = "BRIDGE_ROLE_MEMBERS";

    // ============ Verifier 환경변수 키 ============
    string VERIFIER_PRICER = "PRICER";
    string VERIFIER_FINALIZE_BRIDGE_GAS = "FINALIZE_BRIDGE_GAS";
    string VERIFIER_DEFAULT_TOKEN_PRICE = "DEFAULT_TOKEN_PRICE";
    string VERIFIER_DEFAULT_EX_FEE_RATE = "DEFAULT_EX_FEE_RATE";
    string VERIFIER_MINIMUM_TOKEN_VALUE = "MINIMUM_TOKEN_VALUE";
    string VERIFIER_VERIFICATION_AMOUNT_THRESHOLD = "VERIFICATION_AMOUNT_THRESHOLD";
    string VERIFIER_PERIOD_TOTAL_VALUE_THRESHOLD = "PERIOD_TOTAL_VALUE_THRESHOLD";
    string VERIFIER_TIME_WINDOW = "TIME_WINDOW";
    string VERIFIER_ROLES = "VERIFIER_ROLES";
    string VERIFIER_ROLE_MEMBERS = "VERIFIER_ROLE_MEMBERS";
    string VERIFIER_GAS_PRICES = "VERIFIER_GAS_PRICES";
    string VERIFIER_GAS_PRICE_CHAINS = "VERIFIER_GAS_PRICE_CHAINS";

    // ============ 환경변수에서 로드된 값들 ============
    address priceFeed;
    address impl;
    address owner;
    address payable dev;
    uint8 threshold;
    uint crossChainID;
    uint bscChainID;
    address cross;
    uint crossInitialSupply;
    uint finalizeBridgeGas;
    uint defaultTokenPrice;
    uint defaultExFeeRate;
    uint minimumTokenValue;
    uint verificationAmountThreshold;
    uint periodTotalValueThreshold;
    uint timeWindow;
    bytes32[] bridgeRoles;
    address[] bridgeRoleMembers;
    bytes32[] verifierRoles;
    address[] verifierRoleMembers;
    uint[] verifierGasPrices;
    uint[] verifierGasPriceChains;

    // 기본값용 빈 배열
    bytes32[] emptyBytes32Array;
    address[] emptyAddressArray;
    uint[] emptyUintArray;

    /**
     * @notice 환경변수 로드 및 유효성 검증
     */
    function setUp() public virtual {
        // 환경변수 로드
        priceFeed = vm.envAddress(PRICE_FEED);
        impl = vm.envAddress(IMPLEMENTATION);
        owner = vm.envAddress(OWNER);
        dev = payable(vm.envAddress(BRIDGE_DEV));
        threshold = uint8(vm.envUint(BRIDGE_THRESHOLD));
        cross = vm.envAddress(BRIDGE_CROSS);
        crossInitialSupply = vm.envUint(BRIDGE_CROSS_INITIAL_SUPPLY) * 1 ether;
        finalizeBridgeGas = vm.envUint(VERIFIER_FINALIZE_BRIDGE_GAS);
        defaultTokenPrice = vm.envUint(VERIFIER_DEFAULT_TOKEN_PRICE);
        defaultExFeeRate = vm.envUint(VERIFIER_DEFAULT_EX_FEE_RATE);
        minimumTokenValue = vm.envUint(VERIFIER_MINIMUM_TOKEN_VALUE);
        verificationAmountThreshold = vm.envUint(VERIFIER_VERIFICATION_AMOUNT_THRESHOLD);
        periodTotalValueThreshold = vm.envUint(VERIFIER_PERIOD_TOTAL_VALUE_THRESHOLD);
        timeWindow = vm.envUint(VERIFIER_TIME_WINDOW);
        bridgeRoles = vm.envOr(BRIDGE_ROLES, ",", emptyBytes32Array);
        bridgeRoleMembers = vm.envOr(BRIDGE_ROLE_MEMBERS, ",", emptyAddressArray);
        verifierRoles = vm.envOr(VERIFIER_ROLES, ",", emptyBytes32Array);
        verifierRoleMembers = vm.envOr(VERIFIER_ROLE_MEMBERS, ",", emptyAddressArray);
        verifierGasPrices = vm.envOr(VERIFIER_GAS_PRICES, ",", emptyUintArray);
        verifierGasPriceChains = vm.envOr(VERIFIER_GAS_PRICE_CHAINS, ",", emptyUintArray);

        // 로드된 값 출력
        console.log("priceFeed:", priceFeed);
        console.log("owner:", owner);
        console.log("dev:", dev);
        console.log("threshold:", threshold);
        console.log("cross:", cross);
        console.log("crossInitialSupply:", crossInitialSupply);
        console.log("finalizeBridgeGas:", finalizeBridgeGas);
        console.log("defaultTokenPrice:", defaultTokenPrice);
        console.log("defaultExFeeRate:", defaultExFeeRate);
        console.log("minimumTokenValue:", minimumTokenValue);
        console.log("verificationAmountThreshold:", verificationAmountThreshold);
        console.log("periodTotalValueThreshold:", periodTotalValueThreshold);

        // Bridge 역할 배열 유효성 검증
        require(bridgeRoles.length == bridgeRoleMembers.length, "bridgeRoles and bridgeRoleMembers length mismatch");
        console.log("bridgeRoles:");
        for (uint i = 0; i < bridgeRoles.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "role: ", vm.toString(bridgeRoles[i]), ", account: ", vm.toString(bridgeRoleMembers[i])
                    )
                )
            );
        }

        // Verifier 역할 배열 유효성 검증
        require(
            verifierRoles.length == verifierRoleMembers.length, "verifierRoles and verifierRoleMembers length mismatch"
        );
        console.log("verifierRoles:");
        for (uint i = 0; i < verifierRoles.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "role: ", vm.toString(verifierRoles[i]), ", account: ", vm.toString(verifierRoleMembers[i])
                    )
                )
            );
        }

        // Gas Price 배열 유효성 검증
        require(
            verifierGasPrices.length == verifierGasPriceChains.length,
            "verifierGasPrices and verifierGasPriceChains length mismatch"
        );
        console.log("verifierGasPrices:");
        for (uint i = 0; i < verifierGasPrices.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "chainId: ",
                        vm.toString(verifierGasPriceChains[i]),
                        ", gasPrice: ",
                        vm.toString(verifierGasPrices[i])
                    )
                )
            );
        }
    }

    /**
     * @notice BaseBridge 전체 설정 (프록시 배포 + Verifier + ERC20Code)
     * @dev 환경변수 기반으로 Bridge 시스템 전체를 배포하고 설정합니다.
     */
    function setupBaseBridge() public {
        crossChainID = vm.envUint(BridgeScript.BRIDGE_CROSS_CHAIN_ID);
        console.log("crossChainID:", crossChainID);

        BaseBridge baseBridge = BaseBridge(payable(deployBaseBridgeProxy()));

        _setupBridge(address(baseBridge));
    }

    /**
     * @notice BaseBridge 프록시 배포
     * @return 배포된 프록시 주소
     */
    function deployBaseBridgeProxy() public returns (address) {
        vm.broadcast();
        address proxy = address(new ERC1967Proxy(impl, abi.encodeCall(BaseBridge.initialize, (owner, dev, threshold))));
        console.log("BaseBridgeProxy deployed to:", proxy);
        return proxy;
    }

    /**
     * @notice Bridge 초기화 후 설정 (내부 함수)
     * @dev CrossMintableERC20V2Code, BridgeVerifier 배포 및 연결
     * @param bridge BaseBridge 프록시 주소
     */
    function _setupBridge(address bridge) internal {
        BaseBridge baseBridge = BaseBridge(payable(bridge));

        vm.startBroadcast();
        // Bridge 역할 부여
        baseBridge.grantRoleBatch(bridgeRoles, bridgeRoleMembers);

        // CrossMintableERC20V2Code 배포 및 연결
        CrossMintableERC20V2Code erc20 = new CrossMintableERC20V2Code(owner, address(baseBridge));
        baseBridge.setCrossMintableERC20Code(erc20);

        // BridgeVerifier 배포 및 연결
        BridgeVerifier verifier = new BridgeVerifier(
            owner,
            address(baseBridge),
            priceFeed,
            finalizeBridgeGas,
            defaultTokenPrice,
            defaultExFeeRate,
            minimumTokenValue,
            verificationAmountThreshold,
            periodTotalValueThreshold,
            timeWindow
        );
        verifier.grantRoleBatch(verifierRoles, verifierRoleMembers);

        baseBridge.setBridgeVerifier(verifier);

        console.log("CrossMintableERC20V2Code deployed to:", address(erc20));
        console.log("BridgeVerifier deployed to:", address(verifier));
        vm.stopBroadcast();

        // Gas Price 설정 (선택적)
        if (verifierGasPriceChains.length > 0) {
            vm.broadcast(vm.envAddress(VERIFIER_PRICER));
            verifier.updateGasPriceBatch(verifierGasPriceChains, verifierGasPrices);
        }
    }
}

/*
 * ===================================
 * 환경변수 예제 (.env)
 * ===================================
 *
 * # --------------------------------------------------
 * # 공통 환경변수
 * # --------------------------------------------------
 *
 * # RPC 노드 URL
 * RPC_URL="{node url}"
 *
 * # Admin 권한을 가질 주소
 * OWNER=0x0000000000000000000000000000000000000000
 *
 * # PriceFeed 프록시 주소 (미리 배포 필요)
 * PRICE_FEED=0x0000000000000000000000000000000000000000
 *
 * # BaseBridge Implementation 주소 (미리 배포 필요)
 * IMPLEMENTATION=0x0000000000000000000000000000000000000000
 *
 * # --------------------------------------------------
 * # Bridge 환경변수
 * # --------------------------------------------------
 *
 * # 수수료 수취 주소 (dev fee)
 * DEV=0x0000000000000000000000000000000000000000
 *
 * # Multi-sig 임계값 (서명 필요 최소 인원)
 * THRESHOLD=0
 *
 * # 연결할 원격 체인 ID (CrossBridge용)
 * CROSS_CHAIN_ID=0
 *
 * # BSC 체인 ID (BSCBridge용)
 * BSC_CHAIN_ID=0
 *
 * # CROSS 토큰 주소 (선택적)
 * CROSS=0x0000000000000000000000000000000000000000
 *
 * # CROSS 토큰 초기 공급량 (단위: ether 없이 숫자만)
 * CROSS_INITIAL_SUPPLY=0
 *
 * # --------------------------------------------------
 * # Verifier 환경변수
 * # --------------------------------------------------
 *
 * # 가격 설정 권한을 가질 주소
 * PRICER=0x0000000000000000000000000000000000000000
 *
 * # finalizeBridge 실행 시 예상 가스비 (wei)
 * FINALIZE_BRIDGE_GAS=0
 *
 * # 토큰 기본 가격 (달러 기준, dollarDecimals 적용)
 * DEFAULT_TOKEN_PRICE=0
 *
 * # 기본 교환 수수료율 (1e6 = 100%)
 * DEFAULT_EX_FEE_RATE=0
 *
 * # 최소 브릿지 가능 토큰 가치 (달러 기준)
 * MINIMUM_TOKEN_VALUE=0
 *
 * # 추가 검증이 필요한 금액 임계값
 * VERIFICATION_AMOUNT_THRESHOLD=0
 *
 * # 기간당 총 가치 임계값 (속도 제한용)
 * PERIOD_TOTAL_VALUE_THRESHOLD=0
 *
 * # 속도 제한 시간 윈도우 (초)
 * TIME_WINDOW=0
 *
 * # --------------------------------------------------
 * # 역할 설정
 * # --------------------------------------------------
 *
 * # 역할 해시값 참고:
 * # ADMIN_ROLE: 0xa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775
 * # PRICER_ROLE: 0xc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a
 * # EDITOR_ROLE: 0x21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c
 * # VERIFIER_ROLE: 0x0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09
 * # OPERATOR_ROLE: 0x97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929
 * # VALIDATOR_ROLE: 0x21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926
 *
 * # Bridge 역할 목록 (쉼표로 구분)
 * BRIDGE_ROLES=0xa498...,0xc682...
 *
 * # Bridge 역할을 받을 주소 목록 (BRIDGE_ROLES와 1:1 매칭)
 * BRIDGE_ROLE_MEMBERS=0x...,0x...
 *
 * # Verifier 역할 목록 (쉼표로 구분)
 * VERIFIER_ROLES=0xa498...,0xc682...
 *
 * # Verifier 역할을 받을 주소 목록 (VERIFIER_ROLES와 1:1 매칭)
 * VERIFIER_ROLE_MEMBERS=0x...,0x...
 *
 * # --------------------------------------------------
 * # Gas Price 설정 (선택적)
 * # --------------------------------------------------
 *
 * # 체인 ID 목록
 * VERIFIER_GAS_PRICE_CHAINS=1,56
 *
 * # 각 체인의 가스 가격 (wei)
 * VERIFIER_GAS_PRICES=20000000000,5000000000
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # BaseBridge 전체 설정 (프록시 + Verifier + ERC20Code)
 * # forge script script/Bridge.s.sol:BridgeScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "setupBaseBridge()" \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 배포 후 확인
 * # --------------------------------------------------
 *
 * # Bridge 설정 확인
 * # cast call $BRIDGE "bridgeVerifier()" --rpc-url $RPC_URL
 * # cast call $BRIDGE "crossMintableERC20Code()" --rpc-url $RPC_URL
 * # cast call $BRIDGE "dev()" --rpc-url $RPC_URL
 * # cast call $BRIDGE "threshold()" --rpc-url $RPC_URL
 *
 * # Verifier 설정 확인
 * # cast call $VERIFIER "finalizeBridgeGas()" --rpc-url $RPC_URL
 * # cast call $VERIFIER "defaultTokenPrice()" --rpc-url $RPC_URL
 * # cast call $VERIFIER "defaultExFeeRate()" --rpc-url $RPC_URL
 */
