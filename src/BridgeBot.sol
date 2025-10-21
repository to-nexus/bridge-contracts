// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import {BaseBridge} from "./BaseBridge.sol";
import {IBridgeVerifier} from "./interface/IBridgeVerifier.sol";

/**
 * @title BridgeBot
 * @notice 주기적으로 토큰을 브릿지하는 봇 컨트랙트
 * @dev 설정된 주기에 따라 토큰을 자동으로 브릿지하는 기능을 제공
 * - 주기적 브릿지 실행
 * - 오너 권한으로 설정 변경 및 출금
 * - 누구나 브릿지 실행 가능
 */
contract BridgeBot is Ownable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    error BridgeBotCanNotZeroAddress();
    error BridgeBotCanNotZeroValue();
    error BridgeBotNotTimeYet(uint nextAvailableTime);
    error BridgeBotInsufficientBalance(uint required, uint available);
    error BridgeBotBridgeFailed();

    /**
     * @notice 브릿지 설정 구조체
     * @param tokenAddress 브릿지할 토큰 주소
     * @param recipient 받는 지갑 주소
     * @param toChainID 대상 체인 ID
     * @param interval 브릿지 실행 간격 (초 단위)
     * @param lastExecuted 마지막 실행 시점
     * @param enabled 활성화 여부
     */
    struct BridgeConfig {
        address tokenAddress;
        address recipient;
        uint toChainID;
        uint interval;
        uint lastExecuted;
        bool enabled;
    }

    /**
     * @notice 브릿지 실행 이벤트
     * @param configId 설정 ID
     * @param tokenAddress 토큰 주소
     * @param amount 브릿지된 수량
     * @param recipient 받는 주소
     * @param toChainID 대상 체인 ID
     * @param executor 실행자
     * @param timestamp 실행 시간
     */
    event BridgeExecuted(
        uint indexed configId,
        address indexed tokenAddress,
        uint amount,
        address indexed recipient,
        uint toChainID,
        address executor,
        uint timestamp
    );

    /**
     * @notice 브릿지 설정 추가 이벤트
     */
    event BridgeConfigAdded(uint indexed configId, BridgeConfig config);

    /**
     * @notice 브릿지 설정 업데이트 이벤트
     */
    event BridgeConfigUpdated(uint indexed configId, BridgeConfig config);

    /**
     * @notice 브릿지 설정 활성화/비활성화 이벤트
     */
    event BridgeConfigToggled(uint indexed configId, bool enabled);

    /**
     * @notice 토큰 출금 이벤트
     */
    event TokenWithdrawn(address indexed token, address indexed to, uint amount);

    /**
     * @notice 네이티브 토큰 출금 이벤트
     */
    event NativeWithdrawn(address indexed to, uint amount);

    /// @dev 네이티브 토큰 주소 상수
    address public constant NATIVE_TOKEN = address(1);

    /// @dev 브릿지 컨트랙트
    BaseBridge public immutable bridge;

    /// @dev 브릿지 설정들
    mapping(uint => BridgeConfig) public bridgeConfigs;

    /// @dev 다음 설정 ID
    uint public nextConfigId;

    /**
     * @notice 컨트랙트 생성자
     * @param _bridge 브릿지 컨트랙트 주소
     * @param _owner 컨트랙트 오너 주소
     */
    constructor(address _bridge, address _owner) Ownable(_owner) {
        require(_bridge != address(0), BridgeBotCanNotZeroAddress());
        bridge = BaseBridge(payable(_bridge));
    }

    /**
     * @notice 브릿지 설정 추가
     * @param tokenAddress 토큰 주소
     * @param recipient 받는 주소
     * @param toChainID 대상 체인 ID
     * @param interval 실행 간격 (초)
     * @return configId 설정 ID
     */
    function addBridgeConfig(address tokenAddress, address recipient, uint toChainID, uint interval)
        external
        onlyOwner
        returns (uint configId)
    {
        require(tokenAddress != address(0), BridgeBotCanNotZeroAddress());
        require(recipient != address(0), BridgeBotCanNotZeroAddress());
        require(interval > 0, BridgeBotCanNotZeroValue());

        configId = nextConfigId++;

        bridgeConfigs[configId] = BridgeConfig({
            tokenAddress: tokenAddress,
            recipient: recipient,
            toChainID: toChainID,
            interval: interval,
            lastExecuted: 0,
            enabled: true
        });

        emit BridgeConfigAdded(configId, bridgeConfigs[configId]);
    }

    /**
     * @notice 브릿지 설정 업데이트
     * @param configId 설정 ID
     * @param tokenAddress 토큰 주소
     * @param recipient 받는 주소
     * @param toChainID 대상 체인 ID
     * @param interval 실행 간격 (초)
     */
    function updateBridgeConfig(uint configId, address tokenAddress, address recipient, uint toChainID, uint interval)
        external
        onlyOwner
    {
        require(bridgeConfigs[configId].tokenAddress != address(0), "Config not exists");
        require(tokenAddress != address(0), BridgeBotCanNotZeroAddress());
        require(recipient != address(0), BridgeBotCanNotZeroAddress());
        require(interval > 0, BridgeBotCanNotZeroValue());

        bridgeConfigs[configId].tokenAddress = tokenAddress;
        bridgeConfigs[configId].recipient = recipient;
        bridgeConfigs[configId].toChainID = toChainID;
        bridgeConfigs[configId].interval = interval;

        emit BridgeConfigUpdated(configId, bridgeConfigs[configId]);
    }

    /**
     * @notice 브릿지 설정 활성화/비활성화
     * @param configId 설정 ID
     * @param enabled 활성화 여부
     */
    function toggleBridgeConfig(uint configId, bool enabled) external onlyOwner {
        require(bridgeConfigs[configId].tokenAddress != address(0), "Config not exists");

        bridgeConfigs[configId].enabled = enabled;
        emit BridgeConfigToggled(configId, enabled);
    }

    /**
     * @notice 브릿지 실행 가능 여부 확인
     * @param configId 설정 ID
     * @return canExecute 실행 가능 여부
     * @return nextAvailableTime 다음 실행 가능 시간
     */
    function canExecuteBridge(uint configId) public view returns (bool canExecute, uint nextAvailableTime) {
        BridgeConfig memory config = bridgeConfigs[configId];

        if (!config.enabled || config.tokenAddress == address(0)) return (false, 0);

        // 첫 번째 실행인 경우 (lastExecuted == 0)
        if (config.lastExecuted == 0) return (true, 0);

        uint currentPeriod = block.timestamp - (block.timestamp % config.interval);
        uint lastExecutedPeriod = config.lastExecuted - (config.lastExecuted % config.interval);

        canExecute = currentPeriod > lastExecutedPeriod;
        nextAvailableTime = canExecute ? 0 : lastExecutedPeriod + config.interval;
    }

    /**
     * @notice 브릿지 실행 (누구나 호출 가능)
     * @param configId 설정 ID
     * @param amount 브릿지할 수량
     */
    function executeBridge(uint configId, uint amount) external nonReentrant {
        _executeBridgeInternal(configId, amount);
    }

    /**
     * @notice 내부 브릿지 실행 함수
     * @param configId 설정 ID
     * @param amount 브릿지할 수량
     */
    function _executeBridgeInternal(uint configId, uint amount) internal {
        require(amount > 0, BridgeBotCanNotZeroValue());

        BridgeConfig storage config = bridgeConfigs[configId];
        require(config.enabled && config.tokenAddress != address(0), "Config not available");

        (bool canExecute, uint nextAvailableTime) = canExecuteBridge(configId);
        require(canExecute, BridgeBotNotTimeYet(nextAvailableTime));

        // 잔액 확인
        uint balance;
        if (config.tokenAddress == NATIVE_TOKEN) balance = address(this).balance;
        else balance = IERC20(config.tokenAddress).balanceOf(address(this));

        // 수수료 계산
        IBridgeVerifier bridgeVerifier = bridge.bridgeVerifier();
        (uint minimumValue, uint gasFee, uint exFee) =
            bridgeVerifier.calculateFee(config.toChainID, IERC20(config.tokenAddress), amount);

        uint totalRequired = amount + gasFee + exFee;
        require(balance >= totalRequired, BridgeBotInsufficientBalance(totalRequired, balance));
        require(amount >= minimumValue, "Amount below minimum");

        // 토큰 승인 (ERC20인 경우)
        if (config.tokenAddress != NATIVE_TOKEN) {
            IERC20 token = IERC20(config.tokenAddress);
            if (token.allowance(address(this), address(bridge)) < totalRequired) {
                token.approve(address(bridge), type(uint).max);
            }
        }

        // 브릿지 실행
        bool success;
        if (config.tokenAddress == NATIVE_TOKEN) {
            success = bridge.bridgeToken{value: totalRequired}(
                config.toChainID, IERC20(config.tokenAddress), config.recipient, amount, gasFee, exFee, ""
            );
        } else {
            success = bridge.bridgeToken(
                config.toChainID, IERC20(config.tokenAddress), config.recipient, amount, gasFee, exFee, ""
            );
        }

        require(success, BridgeBotBridgeFailed());

        // 마지막 실행 시간 업데이트
        config.lastExecuted = block.timestamp;

        emit BridgeExecuted(
            configId, config.tokenAddress, amount, config.recipient, config.toChainID, msg.sender, block.timestamp
        );
    }

    /**
     * @notice 여러 브릿지 설정을 일괄 실행
     * @param configIds 설정 ID 배열
     * @param amounts 각 설정에 대응하는 브릿지 수량 배열
     */
    function executeBridgeBatch(uint[] calldata configIds, uint[] calldata amounts) external nonReentrant {
        require(configIds.length == amounts.length, "Array length mismatch");

        for (uint i = 0; i < configIds.length; i++) {
            (bool canExecute,) = canExecuteBridge(configIds[i]);
            if (canExecute && amounts[i] > 0) _executeBridgeInternal(configIds[i], amounts[i]);
        }
    }

    /**
     * @notice ERC20 토큰 출금 (오너만 가능)
     * @param token 토큰 주소
     * @param to 받을 주소
     * @param amount 출금할 수량
     */
    function withdrawToken(address token, address to, uint amount) external onlyOwner {
        require(token != address(0), BridgeBotCanNotZeroAddress());
        require(to != address(0), BridgeBotCanNotZeroAddress());
        require(amount > 0, BridgeBotCanNotZeroValue());

        IERC20(token).safeTransfer(to, amount);
        emit TokenWithdrawn(token, to, amount);
    }

    /**
     * @notice 네이티브 토큰 출금 (오너만 가능)
     * @param to 받을 주소
     * @param amount 출금할 수량
     */
    function withdrawNative(address payable to, uint amount) external onlyOwner {
        require(to != address(0), BridgeBotCanNotZeroAddress());
        require(amount > 0, BridgeBotCanNotZeroValue());
        require(address(this).balance >= amount, "Insufficient balance");

        (bool success,) = to.call{value: amount}("");
        require(success, "Transfer failed");

        emit NativeWithdrawn(to, amount);
    }

    /**
     * @notice 모든 ERC20 토큰 출금 (오너만 가능)
     * @param token 토큰 주소
     * @param to 받을 주소
     */
    function withdrawAllTokens(address token, address to) external onlyOwner {
        require(token != address(0), BridgeBotCanNotZeroAddress());
        require(to != address(0), BridgeBotCanNotZeroAddress());

        uint balance = IERC20(token).balanceOf(address(this));
        if (balance > 0) {
            IERC20(token).safeTransfer(to, balance);
            emit TokenWithdrawn(token, to, balance);
        }
    }

    /**
     * @notice 모든 네이티브 토큰 출금 (오너만 가능)
     * @param to 받을 주소
     */
    function withdrawAllNative(address payable to) external onlyOwner {
        require(to != address(0), BridgeBotCanNotZeroAddress());

        uint balance = address(this).balance;
        if (balance > 0) {
            (bool success,) = to.call{value: balance}("");
            require(success, "Transfer failed");
            emit NativeWithdrawn(to, balance);
        }
    }

    /**
     * @notice 브릿지 설정 조회
     * @param configId 설정 ID
     * @return config 브릿지 설정
     */
    function getBridgeConfig(uint configId) external view returns (BridgeConfig memory config) {
        return bridgeConfigs[configId];
    }

    /**
     * @notice 실행 가능한 브릿지 설정들 조회
     * @param maxConfigs 최대 조회할 설정 개수
     * @return executableConfigs 실행 가능한 설정 ID 배열
     */
    function getExecutableConfigs(uint maxConfigs) external view returns (uint[] memory executableConfigs) {
        uint[] memory temp = new uint[](maxConfigs);
        uint count = 0;

        for (uint i = 0; i < nextConfigId && count < maxConfigs; i++) {
            (bool canExecute,) = canExecuteBridge(i);
            if (canExecute) {
                temp[count] = i;
                count++;
            }
        }

        executableConfigs = new uint[](count);
        for (uint i = 0; i < count; i++) {
            executableConfigs[i] = temp[i];
        }
    }

    /**
     * @notice 네이티브 토큰 수신
     */
    receive() external payable {}

    /**
     * @notice fallback 함수
     */
    fallback() external payable {}
}
