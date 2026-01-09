// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {BSCBridgeV2} from "../src/BSCBridgeV2.sol";
import {BaseBridge} from "../src/BaseBridge.sol";

import {BridgeExecutor, IBridgeExecutor} from "../src/BridgeExecutor.sol";
import {CrossBridge} from "../src/CrossBridge.sol";
import {PriceFeed} from "../src/PriceFeed.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Upgrades} from "lib/openzeppelin-foundry-upgrades/src/Upgrades.sol";

import {Const} from "../src/lib/Const.sol";
import {Script, console} from "forge-std/Script.sol";

/**
 * @title ImplementationScript
 * @notice Implementation 컨트랙트 배포 및 업그레이드 스크립트
 * @dev UUPS 프록시 패턴을 사용하는 컨트랙트들의 Implementation 배포 및 업그레이드를 수행합니다.
 *
 * 제공 기능:
 * - deployPriceFeedImpl(): PriceFeed Implementation 배포
 * - deployBaseBridgeImpl(): BaseBridge Implementation 배포
 * - deployCrossBridgeImpl(): CrossBridge Implementation 배포
 * - upgradeBridgeImplWithoutCall(): 프록시의 Implementation 업그레이드
 *
 * 사용법:
 *   # Implementation 배포
 *   forge script script/Implementation.s.sol:ImplementationScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "deployBaseBridgeImpl()" \
 *     --broadcast
 *
 *   # 프록시 업그레이드
 *   forge script script/Implementation.s.sol:ImplementationScript \
 *     --rpc-url <RPC_URL> \
 *     --sig "upgradeBridgeImplWithoutCall(address,address)" \
 *     <PROXY> <NEW_IMPL> \
 *     --broadcast
 */
contract ImplementationScript is Script {
    function setUp() public {}

    /**
     * @notice PriceFeed Implementation 배포
     * @dev 프록시 없이 Implementation만 배포합니다.
     *      초기 배포 또는 업그레이드용 새 Implementation 배포 시 사용합니다.
     */
    function deployPriceFeedImpl() public {
        vm.broadcast();
        PriceFeed impl = new PriceFeed();
        console.log("PriceFeed Implementation deployed to:", address(impl));
    }

    /**
     * @notice BaseBridge Implementation 배포
     * @dev 프록시 없이 Implementation만 배포합니다.
     *      초기 배포 또는 업그레이드용 새 Implementation 배포 시 사용합니다.
     */
    function deployBaseBridgeImpl() public {
        vm.broadcast();
        BaseBridge impl = new BaseBridge();
        console.log("BaseBridge Implementation deployed to:", address(impl));
    }

    /**
     * @notice CrossBridge Implementation 배포
     * @dev 프록시 없이 Implementation만 배포합니다.
     *      초기 배포 또는 업그레이드용 새 Implementation 배포 시 사용합니다.
     */
    function deployCrossBridgeImpl() public {
        vm.broadcast();
        CrossBridge impl = new CrossBridge();
        console.log("CrossBridge Implementation deployed to:", address(impl));
    }

    /**
     * @notice BSCBridgeV2 Implementation 배포
     * @dev 프록시 없이 Implementation만 배포합니다.
     *      초기 배포 또는 업그레이드용 새 Implementation 배포 시 사용합니다.
     */
    function deployBSCBridgeV2Impl() public {
        vm.broadcast();
        BSCBridgeV2 impl = new BSCBridgeV2();
        console.log("BSCBridgeV2 Implementation deployed to:", address(impl));
    }

    /**
     * @notice 프록시의 Implementation 업그레이드 (초기화 호출 없음)
     * @dev UUPS 프록시의 Implementation을 새 버전으로 업그레이드합니다.
     *      reinitialize가 필요 없는 경우 사용합니다.
     * @param proxy 프록시 컨트랙트 주소
     * @param impl 새 Implementation 주소
     */
    function upgradeBridgeImplWithoutCall(address proxy, address impl) public {
        console.log("Legacy Implementation:", Upgrades.getImplementationAddress(address(proxy)));
        console.log("New Implementation:", impl);

        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(impl, "");

        console.log("Upgraded Implementation:", Upgrades.getImplementationAddress(address(proxy)));
    }

    /**
     * @notice 프록시의 Implementation 업그레이드 및 setMaxExtraDataLength 호출
     * @dev UUPS 프록시의 Implementation을 새 버전으로 업그레이드하고, setMaxExtraDataLength 등 필요한 설정을 수행합니다.
     * @param proxy 프록시 컨트랙트 주소
     * @param impl 새 Implementation 주소
     * @param editor setMaxExtraDataLength를 호출할 권한을 가진 주소
     * @param extraDataLength 설정할 최대 extra data 길이
     * @param executor (옵션) BridgeExecutor 컨트랙트 주소. 없으면 address(0)
     * @param whitelist (옵션) BridgeExecutor에 추가할 화이트리스트 타겟 주소 배열
     */
    function upgradeBridgeImplWithSetMaxExtraDataLength(
        address proxy,
        address impl,
        address editor,
        uint extraDataLength,
        address executor,
        address[] calldata whitelist
    ) public {
        console.log("Legacy Implementation:", Upgrades.getImplementationAddress(address(proxy)));
        console.log("New Implementation:", impl);

        vm.broadcast();
        UUPSUpgradeable(payable(proxy)).upgradeToAndCall(impl, "");

        vm.prank(editor);
        BaseBridge(payable(proxy)).setMaxExtraDataLength(extraDataLength);

        if (executor != address(0)) {
            vm.prank(msg.sender);
            BaseBridge(payable(proxy)).setBridgeExecutor(IBridgeExecutor(executor));
            console.log("Executor:", executor);

            if (whitelist.length > 0) {
                for (uint i = 0; i < whitelist.length; i++) {
                    vm.prank(msg.sender);
                    BridgeExecutor(payable(executor)).addWhitelistTarget(whitelist[i]);
                    console.log("Whitelisted target:", whitelist[i]);
                }
            }
        }

        console.log("Upgraded Implementation:", Upgrades.getImplementationAddress(address(proxy)));
    }
}

/*
 * ===================================
 * 사용 예제
 * ===================================
 *
 * # --------------------------------------------------
 * # Implementation 배포
 * # --------------------------------------------------
 *
 * # PriceFeed Implementation 배포
 * # forge script script/Implementation.s.sol:ImplementationScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "deployPriceFeedImpl()" \
 * #   --broadcast
 *
 * # BaseBridge Implementation 배포
 * # forge script script/Implementation.s.sol:ImplementationScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "deployBaseBridgeImpl()" \
 * #   --broadcast
 *
 * # CrossBridge Implementation 배포
 * # forge script script/Implementation.s.sol:ImplementationScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "deployCrossBridgeImpl()" \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 프록시 업그레이드
 * # --------------------------------------------------
 *
 * # 프록시의 Implementation 업그레이드 (reinitialize 불필요 시)
 * # forge script script/Implementation.s.sol:ImplementationScript \
 * #   --rpc-url $RPC_URL \
 * #   --private-key $PRIVATE_KEY \
 * #   --sig "upgradeBridgeImplWithoutCall(address,address)" \
 * #   $PROXY_ADDRESS $NEW_IMPL_ADDRESS \
 * #   --broadcast
 *
 * # --------------------------------------------------
 * # 현재 Implementation 확인
 * # --------------------------------------------------
 *
 * # ERC1967 프록시의 Implementation 주소 조회
 * # cast implementation $PROXY_ADDRESS --rpc-url $RPC_URL
 *
 * # --------------------------------------------------
 * # 참고사항
 * # --------------------------------------------------
 *
 * # 업그레이드 시 ADMIN_ROLE 권한이 필요합니다.
 * # reinitialize가 필요한 경우 별도의 함수를 사용하거나,
 * # upgradeToAndCall의 두 번째 파라미터에 초기화 calldata를 전달해야 합니다.
 */
