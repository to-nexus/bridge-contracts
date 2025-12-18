// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {PriceFeed} from "../src/PriceFeed.sol";
import {Const} from "../src/lib/Const.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title PriceFeedScript
 * @notice PriceFeed 배포 및 설정 스크립트
 * @dev PriceFeed 프록시 배포, 역할 부여, 토큰/네이티브 토큰 가격 설정을 수행합니다.
 *
 * 제공 기능:
 * - setupPriceFeed(): 환경변수 기반 전체 설정 (프록시 배포 + 역할 + 가격)
 * - deployProxy(): PriceFeed 프록시 배포
 * - setRoles(): 역할 일괄 부여
 * - updatePrice(): ERC20 토큰 가격 업데이트
 * - updateNativeTokenPrice(): 네이티브 토큰 가격 업데이트
 *
 * 환경 변수:
 * - OWNER (required): Admin 권한을 가질 주소
 * - IMPLEMENTATION (required): PriceFeed Implementation 주소
 * - PRICER (required): 가격 설정 권한을 가질 주소
 * - DOLLAR_DECIMALS (required): 달러 소수점 자릿수
 * - PRICE_FEED_ROLES (optional): 부여할 역할 목록 (쉼표 구분)
 * - PRICE_FEED_ROLE_MEMBERS (optional): 역할을 받을 주소 목록 (쉼표 구분)
 * - PRICE_FEED_TOKEN_ADDRESSES (optional): 토큰 주소 목록
 * - PRICE_FEED_TOKEN_PRICES (optional): 토큰 가격 목록
 * - PRICE_FEED_TOKEN_PRICE_AT (optional): 가격 시점 목록
 * - PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS (optional): 네이티브 토큰 체인 ID 목록
 * - PRICE_FEED_NATIVE_TOKEN_PRICES (optional): 네이티브 토큰 가격 목록
 * - PRICE_FEED_NATIVE_TOKEN_PRICE_AT (optional): 가격 시점 목록
 *
 * 사용법:
 *   forge script script/PriceFeed.s.sol:PriceFeedScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "setupPriceFeed()" \
 *     --broadcast
 */
contract PriceFeedScript is Script {
    // 공통 환경변수 키
    string OWNER = "OWNER";
    string IMPLEMENTATION = "IMPLEMENTATION";
    string PRICE_FEED = "PRICE_FEED";
    string PRICER = "PRICER";

    // PriceFeed 전용 환경변수 키
    string PRICE_FEED_DOLLAR_DECIMALS = "DOLLAR_DECIMALS";
    string PRICE_FEED_ROLES = "PRICE_FEED_ROLES";
    string PRICE_FEED_ROLE_MEMBERS = "PRICE_FEED_ROLE_MEMBERS";
    string PRICE_FEED_TOKEN_ADDRESSES = "PRICE_FEED_TOKEN_ADDRESSES";
    string PRICE_FEED_TOKEN_PRICES = "PRICE_FEED_TOKEN_PRICES";
    string PRICE_FEED_TOKEN_PRICE_AT = "PRICE_FEED_TOKEN_PRICE_AT";
    string PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS = "PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS";
    string PRICE_FEED_NATIVE_TOKEN_PRICES = "PRICE_FEED_NATIVE_TOKEN_PRICES";
    string PRICE_FEED_NATIVE_TOKEN_PRICE_AT = "PRICE_FEED_NATIVE_TOKEN_PRICE_AT";

    // 환경변수에서 로드된 값들
    address owner;
    address impl;
    uint8 dollarDecimals;
    bytes32[] roles;
    address[] addresses;

    address[] tokenAddresses;
    uint[] tokenPrices;
    uint[] tokenPriceAt;
    uint[] nativeTokenChainIDs;
    uint[] nativeTokenPrices;
    uint[] nativeTokenPriceAt;

    // 기본값용 빈 배열
    bytes32[] emptyBytes32Array;
    address[] emptyAddressArray;
    uint[] emptyUintArray;

    /**
     * @notice 환경변수 로드 및 유효성 검증
     */
    function setUp() public {
        // 환경변수 로드
        owner = vm.envAddress(OWNER);
        impl = vm.envAddress(IMPLEMENTATION);
        dollarDecimals = uint8(vm.envUint(PRICE_FEED_DOLLAR_DECIMALS));
        roles = vm.envOr(PRICE_FEED_ROLES, ",", emptyBytes32Array);
        addresses = vm.envOr(PRICE_FEED_ROLE_MEMBERS, ",", emptyAddressArray);

        tokenAddresses = vm.envOr(PRICE_FEED_TOKEN_ADDRESSES, ",", emptyAddressArray);
        tokenPrices = vm.envOr(PRICE_FEED_TOKEN_PRICES, ",", emptyUintArray);
        tokenPriceAt = vm.envOr(PRICE_FEED_TOKEN_PRICE_AT, ",", emptyUintArray);
        nativeTokenChainIDs = vm.envOr(PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS, ",", emptyUintArray);
        nativeTokenPrices = vm.envOr(PRICE_FEED_NATIVE_TOKEN_PRICES, ",", emptyUintArray);
        nativeTokenPriceAt = vm.envOr(PRICE_FEED_NATIVE_TOKEN_PRICE_AT, ",", emptyUintArray);

        // 로드된 값 출력
        console.log("impl:", impl);
        console.log("owner:", owner);

        // 역할 배열 유효성 검증
        require(roles.length == addresses.length, "roles and addresses length mismatch");
        console.log("roles:");
        for (uint i = 0; i < roles.length; i++) {
            console.logString(
                string(abi.encodePacked("role: ", vm.toString(roles[i]), ", account: ", vm.toString(addresses[i])))
            );
        }

        // 토큰 가격 배열 유효성 검증
        if (tokenAddresses.length != tokenPrices.length || tokenAddresses.length != tokenPriceAt.length) {
            console.log("tokenAddresses.length:", tokenAddresses.length);
            console.log("tokenPrices.length:", tokenPrices.length);
            console.log("tokenPriceAt.length:", tokenPriceAt.length);
            revert("Invalid input");
        }
        for (uint i = 0; i < tokenAddresses.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "TokenPrice[",
                        vm.toString(i),
                        "]: {address: ",
                        vm.toString(tokenAddresses[i]),
                        ", price: ",
                        vm.toString(tokenPrices[i]),
                        ", priceAt: ",
                        vm.toString(tokenPriceAt[i]),
                        "}"
                    )
                )
            );
        }

        // 네이티브 토큰 가격 배열 유효성 검증
        require(
            nativeTokenChainIDs.length == nativeTokenPrices.length
                && nativeTokenChainIDs.length == nativeTokenPriceAt.length,
            "Invalid input"
        );
        for (uint i = 0; i < nativeTokenChainIDs.length; i++) {
            console.logString(
                string(
                    abi.encodePacked(
                        "NativeTokenPrice[",
                        vm.toString(i),
                        "]: {chainID: ",
                        vm.toString(nativeTokenChainIDs[i]),
                        ", price: ",
                        vm.toString(nativeTokenPrices[i]),
                        ", priceAt: ",
                        vm.toString(nativeTokenPriceAt[i]),
                        "}"
                    )
                )
            );
        }
    }

    /**
     * @notice 전체 PriceFeed 설정 (프록시 배포 + 역할 + 가격)
     * @dev 환경변수 기반으로 PriceFeed를 완전히 설정합니다.
     */
    function setupPriceFeed() public {
        address pricer = vm.envAddress(PRICER);
        PriceFeed priceFeed = PriceFeed(deployProxy(owner, impl, owner, dollarDecimals));
        setRoles(owner, priceFeed, roles, addresses);

        if (tokenAddresses.length > 0) updatePrice(pricer, priceFeed, tokenAddresses, tokenPrices, tokenPriceAt);

        if (nativeTokenChainIDs.length > 0) {
            updateNativeTokenPrice(pricer, priceFeed, nativeTokenChainIDs, nativeTokenPrices, nativeTokenPriceAt);
        }
    }

    /**
     * @notice PriceFeed 프록시 배포
     * @param signer 트랜잭션 서명자
     * @param impl_ Implementation 주소
     * @param owner_ Admin 권한을 가질 주소
     * @param dollarDecimals_ 달러 소수점 자릿수
     * @return 배포된 프록시 주소
     */
    function deployProxy(address signer, address impl_, address owner_, uint8 dollarDecimals_)
        public
        returns (address)
    {
        vm.broadcast(signer);
        address proxy =
            address(new ERC1967Proxy(address(impl_), abi.encodeCall(PriceFeed.initialize, (owner_, dollarDecimals_))));
        console.log("PriceFeed deployed at:", proxy);
        return proxy;
    }

    /**
     * @notice 역할 일괄 부여
     * @param signer 트랜잭션 서명자 (Admin 권한 필요)
     * @param priceFeed PriceFeed 컨트랙트
     * @param roles_ 부여할 역할 배열
     * @param addresses_ 역할을 받을 주소 배열
     */
    function setRoles(address signer, PriceFeed priceFeed, bytes32[] memory roles_, address[] memory addresses_)
        public
    {
        require(roles_.length == addresses_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.grantRoleBatch(roles_, addresses_);
    }

    /**
     * @notice ERC20 토큰 가격 업데이트
     * @param signer 트랜잭션 서명자 (PRICER_ROLE 필요)
     * @param priceFeed PriceFeed 컨트랙트
     * @param tokens_ 토큰 주소 배열
     * @param prices_ 가격 배열 (dollarDecimals 기준)
     * @param priceAt_ 가격 시점 배열 (timestamp)
     */
    function updatePrice(
        address signer,
        PriceFeed priceFeed,
        address[] memory tokens_,
        uint[] memory prices_,
        uint[] memory priceAt_
    ) public {
        require(tokens_.length == prices_.length && tokens_.length == priceAt_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.updatePrice(tokens_, prices_, priceAt_);
    }

    /**
     * @notice 네이티브 토큰 가격 업데이트
     * @param signer 트랜잭션 서명자 (PRICER_ROLE 필요)
     * @param priceFeed PriceFeed 컨트랙트
     * @param chainIDs 체인 ID 배열
     * @param prices_ 가격 배열 (dollarDecimals 기준)
     * @param priceAt_ 가격 시점 배열 (timestamp)
     */
    function updateNativeTokenPrice(
        address signer,
        PriceFeed priceFeed,
        uint[] memory chainIDs,
        uint[] memory prices_,
        uint[] memory priceAt_
    ) public {
        require(chainIDs.length == prices_.length && chainIDs.length == priceAt_.length, "Invalid input");
        vm.broadcast(signer);
        priceFeed.updateNativeTokenPrice(chainIDs, prices_, priceAt_);
    }
}

/*
 * ===================================
 * 환경변수 예제 (.env)
 * ===================================
 *
 * # --------------------------------------------------
 * # 필수 환경 변수
 * # --------------------------------------------------
 *
 * # RPC 노드 URL
 * RPC_URL="{node url}"
 *
 * # Admin 권한을 가질 주소 (DEFAULT_ADMIN_ROLE 부여됨)
 * OWNER=0x0000000000000000000000000000000000000000
 *
 * # PriceFeed Implementation 주소 (미리 배포된 Implementation)
 * IMPLEMENTATION=0x0000000000000000000000000000000000000000
 *
 * # 가격 설정 권한을 가질 주소 (PRICER_ROLE 부여됨)
 * PRICER=0x0000000000000000000000000000000000000000
 *
 * # 달러 소수점 자릿수 (예: 6이면 1 USD = 1000000)
 * DOLLAR_DECIMALS=6
 *
 * # --------------------------------------------------
 * # 토큰 가격 설정 (선택)
 * # --------------------------------------------------
 *
 * # ERC20 토큰 주소 목록 (쉼표로 구분)
 * PRICE_FEED_TOKEN_ADDRESSES=0x...,0x...
 *
 * # 토큰 가격 목록 (dollarDecimals 기준, 쉼표로 구분)
 * # 예: DOLLAR_DECIMALS=6일 때 1 USD = 1000000
 * PRICE_FEED_TOKEN_PRICES=1000000,2000000
 *
 * # 가격 설정 시점 (timestamp, 0이면 현재 시점)
 * PRICE_FEED_TOKEN_PRICE_AT=0,0
 *
 * # --------------------------------------------------
 * # 네이티브 토큰 가격 설정 (선택)
 * # --------------------------------------------------
 *
 * # 네이티브 토큰의 체인 ID 목록
 * PRICE_FEED_NATIVE_TOKEN_CHAIN_IDS=1,56
 *
 * # 네이티브 토큰 가격 목록 (dollarDecimals 기준)
 * PRICE_FEED_NATIVE_TOKEN_PRICES=2000000000,300000000
 *
 * # 가격 설정 시점 (timestamp, 0이면 현재 시점)
 * PRICE_FEED_NATIVE_TOKEN_PRICE_AT=0,0
 *
 * # --------------------------------------------------
 * # 역할 설정 (선택)
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
 * # 부여할 역할 목록 (쉼표로 구분)
 * PRICE_FEED_ROLES=0xa498...,0xc682...
 *
 * # 역할을 받을 주소 목록 (PRICE_FEED_ROLES와 1:1 매칭)
 * PRICE_FEED_ROLE_MEMBERS=0x...,0x...
 *
 * # --------------------------------------------------
 * # 배포 명령어 예시
 * # --------------------------------------------------
 *
 * # 전체 설정 (프록시 배포 + 역할 + 가격)
 * # forge script script/PriceFeed.s.sol:PriceFeedScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "setupPriceFeed()" \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 배포 후 확인
 * # --------------------------------------------------
 *
 * # 달러 소수점 확인
 * # cast call $PRICE_FEED "dollarDecimals()" --rpc-url $RPC_URL
 *
 * # 토큰 가격 확인
 * # cast call $PRICE_FEED "getPrice(address)" $TOKEN --rpc-url $RPC_URL
 *
 * # 네이티브 토큰 가격 확인
 * # cast call $PRICE_FEED "getNativeTokenPrice(uint256)" $CHAIN_ID --rpc-url $RPC_URL
 */
