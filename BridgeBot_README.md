# BridgeBot - 주기적 토큰 브릿지 자동화 컨트랙트

## 개요

BridgeBot은 설정된 주기에 따라 자동으로 토큰을 브릿지하는 스마트 컨트랙트입니다. 사용자가 설정한 간격에 따라 지정된 수량의 토큰을 대상 체인으로 자동 전송합니다.

## 주요 기능

### 🔄 주기적 브릿지 실행
- 설정된 시간 간격(초 단위)에 따라 자동 브릿지 실행
- 예: 86400초 (24시간) 간격으로 매일 실행
- 스마트한 시간 체크: `현재시간 - (현재시간 % 간격)`으로 중복 실행 방지

### ⚙️ 유연한 설정 관리
- 토큰 주소, 수량, 받는 주소, 대상 체인 ID 설정 가능
- 오너 권한으로 설정 추가/수정/활성화/비활성화
- 여러 브릿지 설정을 동시에 운영 가능

### 🔐 보안 기능
- 오너만 설정 변경 및 출금 가능
- 누구나 브릿지 실행 가능 (실행 권한 분리)
- 재진입 공격 방지

### 💰 자금 관리
- ERC20 토큰 및 네이티브 토큰 지원
- 오너 권한으로 언제든지 출금 가능
- 부분/전체 출금 기능

## 사용법

### 1. 컨트랙트 배포

```solidity
// BridgeBot 배포
BridgeBot bridgeBot = new BridgeBot(
    address(baseBridge),  // 기존 브릿지 컨트랙트 주소
    ownerAddress         // 오너 주소
);
```

### 2. 브릿지 설정 추가

```solidity
// 매일 CROSS 토큰 3500개를 Cross 체인으로 브릿지
uint256 configId = bridgeBot.addBridgeConfig(
    address(crossToken),    // 토큰 주소
    3500 ether,            // 브릿지할 수량
    recipientAddress,       // 받는 주소
    612044,                // Cross 체인 ID
    86400                  // 24시간 간격 (초)
);
```

### 3. 자금 입금

```solidity
// ERC20 토큰 전송
IERC20(crossToken).transfer(address(bridgeBot), 10000 ether);

// 네이티브 토큰 전송
payable(address(bridgeBot)).transfer(1 ether);
```

### 4. 브릿지 실행

```solidity
// 개별 실행 (누구나 호출 가능)
bridgeBot.executeBridge(configId);

// 일괄 실행
uint256[] memory configIds = new uint256[](2);
configIds[0] = 0;
configIds[1] = 1;
bridgeBot.executeBridgeBatch(configIds);

// 실행 가능한 설정 자동 조회 후 실행
uint256[] memory executableConfigs = bridgeBot.getExecutableConfigs(10);
if (executableConfigs.length > 0) {
    bridgeBot.executeBridgeBatch(executableConfigs);
}
```

## 스크립트 사용 예제

### Foundry 스크립트로 배포 및 설정

```bash
# 배포 및 초기 설정
forge script script/BridgeBotExample.s.sol:BridgeBotExample --sig "deployAndSetup()" --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast

# 브릿지 설정 추가
forge script script/BridgeBotExample.s.sol:BridgeBotExample --sig "addBridgeConfigs()" --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast

# 브릿지 실행
forge script script/BridgeBotExample.s.sol:BridgeBotExample --sig "executeBridges()" --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> --broadcast
```

## 주요 함수

### 오너 전용 함수

| 함수명 | 설명 |
|--------|------|
| `addBridgeConfig()` | 새로운 브릿지 설정 추가 |
| `updateBridgeConfig()` | 기존 설정 수정 |
| `toggleBridgeConfig()` | 설정 활성화/비활성화 |
| `withdrawToken()` | ERC20 토큰 출금 |
| `withdrawNative()` | 네이티브 토큰 출금 |
| `withdrawAllTokens()` | 모든 ERC20 토큰 출금 |
| `withdrawAllNative()` | 모든 네이티브 토큰 출금 |

### 공개 함수

| 함수명 | 설명 |
|--------|------|
| `executeBridge()` | 개별 브릿지 실행 |
| `executeBridgeBatch()` | 여러 브릿지 일괄 실행 |
| `canExecuteBridge()` | 실행 가능 여부 확인 |
| `getBridgeConfig()` | 설정 정보 조회 |
| `getExecutableConfigs()` | 실행 가능한 설정 목록 조회 |

## 설정 구조체

```solidity
struct BridgeConfig {
    address tokenAddress;    // 브릿지할 토큰 주소
    uint256 amount;         // 브릿지할 수량
    address recipient;       // 받는 주소
    uint256 toChainID;      // 대상 체인 ID
    uint256 interval;       // 실행 간격 (초)
    uint256 lastExecuted;   // 마지막 실행 시간
    bool enabled;           // 활성화 여부
}
```

## 주기 실행 로직

브릿지 실행 가능 여부는 다음 로직으로 결정됩니다:

```solidity
uint256 currentPeriod = block.timestamp - (block.timestamp % interval);
uint256 lastExecutedPeriod = lastExecuted - (lastExecuted % interval);
bool canExecute = currentPeriod > lastExecutedPeriod;
```

이 로직은 다음과 같이 작동합니다:
- **간격이 86400초(24시간)인 경우**: 오늘 한 번 실행했다면, 내일까지 다시 실행할 수 없음
- **간격이 3600초(1시간)인 경우**: 이번 시간에 한 번 실행했다면, 다음 시간까지 다시 실행할 수 없음

## 이벤트

```solidity
// 브릿지 실행 시 발생
event BridgeExecuted(
    uint256 indexed configId,
    address indexed tokenAddress,
    uint256 amount,
    address indexed recipient,
    uint256 toChainID,
    address executor,
    uint256 timestamp
);

// 설정 관련 이벤트
event BridgeConfigAdded(uint256 indexed configId, BridgeConfig config);
event BridgeConfigUpdated(uint256 indexed configId, BridgeConfig config);
event BridgeConfigToggled(uint256 indexed configId, bool enabled);

// 출금 관련 이벤트
event TokenWithdrawn(address indexed token, address indexed to, uint256 amount);
event NativeWithdrawn(address indexed to, uint256 amount);
```

## 테스트

```bash
# 테스트 실행
forge test --match-contract BridgeBotTest -vv

# 특정 테스트 함수 실행
forge test --match-test testExecuteBridge -vv

# 가스 사용량 포함하여 테스트
forge test --match-contract BridgeBotTest --gas-report
```

## 보안 고려사항

1. **오너 키 관리**: 오너 키는 안전하게 보관하고, 필요시 멀티시그 지갑 사용 권장
2. **자금 관리**: 브릿지에 필요한 최소한의 자금만 예치
3. **설정 검증**: 브릿지 설정 추가 시 토큰 주소와 체인 ID 검증 필수
4. **모니터링**: 브릿지 실행 상태를 정기적으로 모니터링

## 주의사항

- 토큰 승인(approve)이 충분한지 확인
- 브릿지 수수료를 고려하여 충분한 토큰 보유 필요
- 대상 체인의 브릿지 상태 확인
- 네트워크 상태에 따른 가스비 변동 고려

## 라이센스

MIT License
