// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BSCBridgeV2MetaData contains all meta data concerning the BSCBridgeV2 contract.
var BSCBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deadWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alreadyTransferred\",\"type\":\"bool\"}],\"name\":\"burnCrossToDeadWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross_\",\"type\":\"address\"}],\"name\":\"reinitializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeInvalidDeadWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"a10bab0b": "bridgeExecutor()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
		"86e9e175": "burnCrossToDeadWallet(address,uint256,bool)",
		"b7f3358d": "changeThreshold(uint8)",
		"670e6268": "createToken(uint256,address,string,uint8)",
		"4be13f83": "crossMintableERC20Code()",
		"91cca3db": "dev()",
		"f698da25": "domainSeparator()",
		"84b0196e": "eip712Domain()",
		"1938e0f2": "finalizeBridge((uint256,uint256,address,address,uint256,bytes),uint8[],bytes32[],bytes32[])",
		"88d67d6d": "finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[],uint8[][],bytes32[][],bytes32[][])",
		"d5717fc5": "getNextFinalizeIndex(uint256)",
		"ae6893f8": "getNextInitiateIndex(uint256)",
		"b33eb36e": "getPendingArguments(uint256,uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"814914b5": "getTokenPair(uint256,address)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"b2b49e2e": "grantRoleBatch(bytes32[],address[])",
		"91d14854": "hasRole(bytes32,address)",
		"89232a00": "initialize(address,address,uint8)",
		"91cf6d3e": "initializedAt()",
		"761fe2d8": "isTokenFinalizePaused(uint256,address)",
		"aa20ed47": "manualReleasePending(uint256,uint256)",
		"48a00ed8": "manualReleasePendingWithRecipient(uint256,uint256,address)",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"1313996b": "permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"edbbea39": "registerToken(uint256,bool,address,address)",
		"7b429ad8": "reinitializeBSCBridge(uint256,address)",
		"4ee078ba": "releasePending(uint256,uint256)",
		"9f9b4f1c": "releasePendingBatch(uint256[],uint256[])",
		"8ae87c5c": "removePending(uint256,uint256)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"cba25e4b": "setBridgeExecutor(address)",
		"0b1d8d06": "setBridgeVerifier(address)",
		"b8aa837e": "setChainPause(uint256,bool)",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"bedb86fb": "setPause(bool)",
		"bfbf6765": "setTokenPause(uint256,address,bool,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f516020615fce5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615f0790816100c78239608051818181610d390152610ef20152f35b6001600160401b0319166001600160401b039081175f516020615fce5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b61002061386a565b005b5f3560e01c806301ffc9a7146103a15780630b1d8d061461039c5780631313996b146103975780631938e0f214610392578063248a9ca31461038d5780632f2ff15d1461038857806336568abe14610383578063365d71e41461037e57806342cde4e81461037957806348a00ed8146103745780634be13f831461036f5780634d5d00561461036a5780634ee078ba146103655780634f1ef28614610360578063502cc5c01461035b57806352d1902d146103565780635b605f5c146103515780635c975abb1461034c5780635fd262de14610347578063670e626814610342578063761fe2d81461033d57806379214874146103385780637b429ad814610333578063814914b51461032e57806384b0196e1461032957806386e9e1751461032457806388d67d6d1461031f57806389232a001461031a5780638ae87c5c1461031557806391cca3db1461031057806391cf6d3e1461030b57806391d14854146103065780639f9b4f1c14610301578063a10bab0b146102fc578063a217fddf146102f7578063a3246ad3146102f2578063aa1bd0bc146102ed578063aa20ed47146102e8578063ad3cb1cc146102e3578063ae6893f8146102de578063b2b49e2e146102d9578063b33eb36e146102d4578063b7f3358d146102cf578063b8aa837e146102ca578063bedb86fb146102c5578063bfbf6765146102c0578063cba25e4b146102bb578063cf56118e146102b6578063d477f05f146102b1578063d547741f146102ac578063d5717fc5146102a7578063e2af6cd7146102a2578063edbbea391461029d578063f0702e8e14610298578063f4509643146102935763f698da250361000e5761270b565b61266a565b612642565b6124ea565b612471565b612438565b612409565b6123aa565b612336565b6122cb565b6121f2565b612108565b61206c565b611fe8565b611f57565b611e49565b611e10565b611dc9565b611d3e565b611cf1565b611c75565b611c19565b611bf1565b611ade565b611a9c565b611a7f565b611a57565b6119ee565b6118e9565b6117df565b611628565b611574565b611427565b611351565b6112e1565b61125b565b6111e4565b6110dc565b6110ae565b610fcf565b610ee0565b610e50565b610cf7565b610b97565b610b2a565b610ad6565b6109ac565b610980565b61090d565b61081d565b6107e2565b6107b1565b610700565b6104ba565b610418565b346103f75760203660031901126103f75760043563ffffffff60e01b81168091036103f757602090637965db0b60e01b81149081156103e6575b506040519015158152f35b6301ffc9a760e01b1490505f6103db565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103f757565b346103f75760203660031901126103f75760043561043581610407565b61043e33614914565b6001600160a01b0316610452811515612777565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103f7578235916001600160401b0383116103f7576020808501948460051b0101116103f757565b60403660031901126103f7576004356001600160401b0381116103f7576104e590369060040161048a565b602435916001600160401b0383116103f757366023840112156103f7576004830135916001600160401b0383116103f75736602460e08502860101116103f7576024610020940191612996565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761056257604052565b610532565b60e081019081106001600160401b0382111761056257604052565b606081019081106001600160401b0382111761056257604052565b60c081019081106001600160401b0382111761056257604052565b601f909101601f19168101906001600160401b0382119082101761056257604052565b604051906105ea6080836105b8565b565b604051906105ea60e0836105b8565b604051906105ea610100836105b8565b604051906105ea6060836105b8565b6001600160401b0381116105625760051b60200190565b60ff8116036103f757565b9080601f830112156103f75781356106538161061a565b9261066160405194856105b8565b81845260208085019260051b8201019283116103f757602001905b8282106106895750505090565b60208091833561069881610631565b81520191019061067c565b9080601f830112156103f75781356106ba8161061a565b926106c860405194856105b8565b81845260208085019260051b8201019283116103f757602001905b8282106106f05750505090565b81358152602091820191016106e3565b60803660031901126103f7576004356001600160401b0381116103f75760c060031982360301126103f7576024356001600160401b0381116103f75761074a90369060040161063c565b906044356001600160401b0381116103f75761076a9036906004016106a3565b60643591906001600160401b0383116103f7576107ad9361079261079b9436906004016106a3565b92600401612ad4565b60405191829182901515815260200190565b0390f35b346103f75760203660031901126103f75760206107cf600435612e2c565b604051908152f35b35906105ea82610407565b346103f75760403660031901126103f75761002060243560043561080582610407565b61081861081182612e2c565b3390614b34565b614315565b346103f75760403660031901126103f75760043560243561083d81610407565b336001600160a01b038216036108565761002091614375565b63334bd91960e11b5f5260045ffd5b9060406003198301126103f7576004356001600160401b0381116103f75782610890916004016106a3565b91602435906001600160401b0382116103f757806023830112156103f75781600401356108bc8161061a565b926108ca60405194856105b8565b8184526024602085019260051b8201019283116103f757602401905b8282106108f35750505090565b60208091833561090281610407565b8152019101906108e6565b346103f75761091b36610865565b906109298151835114612e4a565b5f5b8251811015610020578061096f61094460019385612e60565b51838060a01b036109558488612e60565b51169061096a3361096583612e2c565b614b34565b614375565b500161092b565b5f9103126103f757565b346103f7575f3660031901126103f757602060ff5f516020615c125f395f51905f525416604051908152f35b346103f75760603660031901126103f7576024356004356044356109cf81610407565b6109d83361499c565b6109e06138b4565b815f5260076020526109fe836109f98160405f206150a3565b612e74565b80610a098484614f35565b916001600160a01b031615610ac2575b8151905f516020615c325f395f51905f526020840191610a7b835195610a71610a5f6040830197885160018060a01b03169986608086019b8c519260a088015194613c55565b610a6881611ebc565b600181146143d4565b51935194516103fb565b94516040516001600160a01b0390961695918291610a9c9142919084612e0e565b0390a45f516020615bd25f395f51905f525f80a35f5f516020615db25f395f51905f525d005b5060608101516001600160a01b0316610a19565b346103f7575f3660031901126103f7575f546040516001600160a01b039091168152602090f35b9181601f840112156103f7578235916001600160401b0383116103f757602083818601950101116103f757565b6101c03660031901126103f757602435600435610b4682610407565b604435610b5281610407565b6064359060a43560843560c4356001600160401b0381116103f757610b7b903690600401610afd565b94909360e03660e31901126103f7576107ad9761079b97612e8e565b346103f75760403660031901126103f757610c76602435600435610bb961388d565b610bc16138b4565b805f526007602052610bda826109f98160405f206150a3565b610c716040610c01610bfc85610bef86612d00565b905f5260205260405f2090565b613242565b610c5e610c2182516080610c17868301516103fb565b9101519087613aef565b50610c2b81611ebc565b610c3481611ebc565b83516020810182905290600190610c5883604081015b03601f1981018552846105b8565b1461327b565b01518015908115610c7e575b42916132a7565b6143f7565b6100206138e9565b4281109150610c6a565b6001600160401b03811161056257601f01601f191660200190565b929192610caf82610c88565b91610cbd60405193846105b8565b8294818452818301116103f7578281602093845f960137010152565b9080601f830112156103f757816020610cf493359101610ca3565b90565b60403660031901126103f757600435610d0f81610407565b6024356001600160401b0381116103f757610d2e903690600401610cd9565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610e2e575b50610e1f57610d7233614914565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610dee575b50610dbb57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615cd25f395f51905f528303610dda57610020925061552f565b632a87526960e21b5f52600483905260245ffd5b610e1191945060203d602011610e18575b610e0981836105b8565b8101906147b1565b925f610d9a565b503d610dff565b63703e46dd60e11b5f5260045ffd5b5f516020615cd25f395f51905f52546001600160a01b0316141590505f610d64565b346103f75760603660031901126103f757602435600435604435610e733361499c565b815f526007602052610e8c836109f98160405f206150a3565b4201804211610edb5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b6130ae565b346103f7575f3660031901126103f7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e1f5760206040515f516020615cd25f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610faf5750505090565b909192602060e082610fc46001948851610f37565b019401929101610fa2565b346103f75760203660031901126103f757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061109557505061101e925003836105b8565b61102882516132fb565b915f5b81518110156110875760019061106b61106661104686612d0e565b6110606110538588612e60565b516001600160a01b031690565b9061334a565b61335f565b6110758287612e60565b526110808186612e60565b500161102b565b604051806107ad8682610f8c565b8454835260019485019487945060209093019201611009565b346103f7575f3660031901126103f757602060ff5f516020615d725f395f51905f5254166040519015158152f35b60e03660031901126103f7576024356004356110f782610407565b6044359161110483610407565b6064359260c435916084359160a435916001600160401b0385116103f7576111b59661116a61113a6111ab973690600401610afd565b95909661114561388d565b6001600160a01b03851694849061115c878d61449f565b6111646138b4565b8b61453f565b9390926040519861117a8a610546565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610ca3565b60e08201526146da565b5f5f516020615db25f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103f75760803660031901126103f75760243560043561120482610407565b604435906001600160401b0382116103f757366023830112156103f7576107ad9261123c61124f933690602481600401359101610ca3565b906064359261124a84610631565b6133e7565b604051918291826111d1565b346103f75760403660031901126103f757602060ff61129160243560043561128282610407565b5f526009845260405f2061334a565b54166040519015158152f35b90602080835192838152019201905f5b8181106112ba5750505090565b82518452602093840193909201916001016112ad565b906020610cf492818152019061129d565b346103f75760203660031901126103f7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061133b576107ad8561132f818703826105b8565b604051918291826112d0565b8254845260209093019260019283019201611318565b346103f75760403660031901126103f75760043560243561137181610407565b61137a33614914565b5f516020615e325f395f51905f52549160ff8360401c168015611403575b6113f4576001600160401b03199092166002175f516020615e325f395f51905f52556113cb916113c66134b0565b613513565b6113d36134d1565b604051600281525f516020615c525f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b0384161015611398565b60e0810192916105ea9190610f37565b346103f75760403660031901126103f7576107ad61146660243560043561144d82610407565b6114556132c5565b505f52600660205260405f2061334a565b60046040519161147583610567565b80546001600160a01b0316835260018101546114ca906114c1906114a461149b826103fb565b602088016130f7565b6114b860a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611417565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161154a9061153c610cf497959693600f60f81b865260e0602087015260e08601906114f1565b9084820360408601526114f1565b60608301949094526001600160a01b031660808201525f60a082015280830360c09091015261129d565b346103f7575f3660031901126103f7575f516020615c925f395f51905f52541580611608575b156115cb576115a76147c0565b6115af61487a565b906107ad6115bb6135e9565b6040519384933091469186611515565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615e525f395f51905f52541561159a565b801515036103f757565b346103f75760603660031901126103f75760043561164581610407565b6024356044356116548161161e565b61165c61388d565b6116646138b4565b6001600160a01b0383165f9081526066602052604090206116909061168b905b5460ff1690565b613604565b61169b8215156134ec565b611752576116b681836116af6065546103fb565b3390614b80565b335b5f516020615eb25f395f51905f526116d1606454614bc9565b926064549261173b6116e46065546103fb565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a46117466138e9565b60405160018152602090f35b61175b33614914565b306116b8565b9080601f830112156103f75781356117788161061a565b9261178660405194856105b8565b81845260208085019260051b820101918383116103f75760208201905b8382106117b257505050505090565b81356001600160401b0381116103f7576020916117d4878480948801016106a3565b8152019101906117a3565b60803660031901126103f7576004356001600160401b0381116103f75761180a90369060040161048a565b90602435906001600160401b0382116103f757366023830112156103f75781600401356118368161061a565b9261184460405194856105b8565b8184526024602085019260051b820101903682116103f75760248101925b8284106118ba5750506044359150506001600160401b0381116103f75761188d903690600401611761565b606435926001600160401b0384116103f7576107ad946118b461079b953690600401611761565b9361361a565b83356001600160401b0381116103f7576020916118de83926024369187010161063c565b815201930192611862565b346103f75760603660031901126103f75760043561190681610407565b6024359061191382610407565b60443561191f81610631565b5f516020615e325f395f51905f52549260ff604085901c1615611941565b1590565b936001600160401b0316801590816119e6575b60011490816119dc575b1590816119d3575b506113f4575f516020615e325f395f51905f5280546001600160401b031916600117905561199892846119c657614be2565b61199e57005b6119a66134d1565b604051600181525f516020615c525f395f51905f529080602081016113ef565b6119ce6134b0565b614be2565b9050155f611966565b303b15915061195e565b859150611954565b346103f75760403660031901126103f757602435600435611a0e33614914565b611a166138b4565b611a208282614f35565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615db25f395f51905f525d005b346103f7575f3660031901126103f7576033546040516001600160a01b039091168152602090f35b346103f7575f3660031901126103f7576020603454604051908152f35b346103f75760403660031901126103f757602060ff611291602435600435611ac382610407565b5f525f516020615d525f395f51905f52845260405f2061334a565b346103f75760403660031901126103f7576004356001600160401b0381116103f757611b0e9036906004016106a3565b6024356001600160401b0381116103f757611b2d9036906004016106a3565b90611b3b81518351146127ac565b5f5b82518110156100205780611be3611b5660019385612e60565b51611b618387612e60565b5190611b6b61388d565b611b736138b4565b805f526007602052611b8c826109f98160405f206150a3565b610c716040611ba1610bfc85610bef86612d00565b610c5e611bb782516080610c17868301516103fb565b50611bc181611ebc565b611bca81611ebc565b835160208101829052908a90610c588360408101610c4a565b611beb6138e9565b01611b3d565b346103f7575f3660031901126103f7576035546040516001600160a01b039091168152602090f35b346103f7575f3660031901126103f75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611c565750505090565b82516001600160a01b0316845260209384019390920191600101611c49565b346103f75760203660031901126103f7576004355f525f516020615cf25f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611cdb576107ad85611ccf818703826105b8565b60405191829182611c33565b8254845260209093019260019283019201611cb8565b346103f75760203660031901126103f7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611d3133614914565b80600155604051908152a1005b346103f75760403660031901126103f757602435600435611d5e3361499c565b611d673361499c565b611d6f6138b4565b805f526007602052611d88826109f98160405f206150a3565b611d9282826143f7565b5f516020615bd25f395f51905f525f80a35f5f516020615db25f395f51905f525d005b60405190611dc46020836105b8565b5f8252565b346103f7575f3660031901126103f7576107ad604051611dea6040826105b8565b60058152640352e302e360dc1b60208201526040519182916020835260208301906114f1565b346103f75760203660031901126103f7576004355f526004602052600160405f20015460018101809111610edb57602090604051908152f35b346103f757611e5736610865565b90611e658151835114612e4a565b5f5b82518110156100205780611ea1611e8060019385612e60565b51838060a01b03611e918488612e60565b5116906108183361096583612e2c565b5001611e67565b634e487b7160e01b5f52602160045260245ffd5b600a1115611ec657565b611ea8565b90600a821015611ec65752565b6020815260606040611f3d60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c06101208601526101408501906114f1565b93611f4f602082015183860190611ecb565b015191015290565b346103f75760403660031901126103f757600435602435905f60408051611f7d81610582565b611f856136b0565b815282602082015201525f52600860205260405f20905f526020526107ad60405f20600760405191611fb683610582565b611fbf8161313e565b8352611fd560ff60068301541660208501613236565b0154604082015260405191829182611ed8565b346103f75760203660031901126103f7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff60043561202a81610631565b61203333614914565b1661203f8115156136e0565b8060ff195f516020615c125f395f51905f525416175f516020615c125f395f51905f5255604051908152a1005b346103f75760403660031901126103f75760043560243561208c8161161e565b61209533614a24565b6120aa825f52600360205260405f2054151590565b156120f55760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526120ea81600360405f2001613502565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103f75760203660031901126103f7576004356121258161161e565b61212e33614a24565b1561218c5761213b61388d565b600160ff195f516020615d725f395f51905f525416175f516020615d725f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615d725f395f51905f525460ff8116156121e35760ff19165f516020615d725f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103f75760803660031901126103f75760243560043561221282610407565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356122418161161e565b60643561224d8161161e565b61225633614a24565b845f5260056020526122ba816122b5855f20986122858161228060018060a01b038216809d6150a3565b6136f6565b885f5260066020526122a586600161229f848b5f2061334a565b0161371e565b885f526009602052865f2061334a565b613502565b8251911515825215156020820152a3005b346103f75760203660031901126103f7576004356122e881610407565b6122f133614914565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103f7575f3660031901126103f757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b818110612394576107ad8561132f818703826105b8565b825484526020909301926001928301920161237d565b346103f75760203660031901126103f7576004356123c781610407565b6123d033614914565b6001600160a01b03166123e4811515612777565b603380546001600160a01b031916821790555f516020615d125f395f51905f525f80a2005b346103f75760403660031901126103f75761002060243560043561242c82610407565b61096a61081182612e2c565b346103f75760203660031901126103f7576004355f526004602052600260405f20015460018101809111610edb57602090604051908152f35b346103f75760203660031901126103f75760043561248e81610407565b61249733614914565b6001600160a01b03166124ab81151561373b565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346103f75760803660031901126103f75760043560243561250a8161161e565b6044359161251783610407565b5f516020615df25f395f51905f526126016064359361253585610407565b61253e33614aac565b61254984151561373b565b6001600160a01b0386169461079b9061256387151561373b565b6001600160a01b038116976125fc9061257d8a151561373b565b612586886158aa565b612606575b6125af816125aa816125a58c5f52600560205260405f2090565b615309565b614ffc565b6125ce6125ba6105ec565b936125c583866130f7565b602085016130f7565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526125f788612d0e565b61334a565b615024565b0390a4005b61263d6126116105db565b8981525f60208201525f60408201525f60608201526126388a5f52600460205260405f2090565b614fd1565b61258b565b346103f7575f3660031901126103f7576032546040516001600160a01b039091168152602090f35b346103f75760403660031901126103f75760243560043561268a82610407565b61269333614914565b805f5260056020525f60046126ce60408320946126bd8161228060018060a01b03821680996156b5565b84845260066020526040842061334a565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103f7575f3660031901126103f7576127236159f5565b61272b615a4c565b6040519060208201925f516020615e725f395f51905f528452604083015260608201524660808201523060a082015260a0815261276960c0826105b8565b519020604051908152602090f35b1561277e57565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b156127b357565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156127f85760051b8101359060fe19813603018212156103f7570190565b6127c2565b35610cf481610407565b903590601e19813603018212156103f757018035906001600160401b0382116103f7576020019181360383136103f757565b91908110156127f85760e0020190565b908160209103126103f75751610cf48161161e565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936128cf97939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161285e565b9480356128db81610407565b6001600160a01b031660e085015260208101356128f781610407565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561292c81610631565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561297a573d9061296182610c88565b9161296f60405193846105b8565b82523d5f602084013e565b606090565b604090610cf49392815281602082015201906114f1565b90919392936129a68584146127ac565b5f945b8386106129b857505050509050565b6129c38685856127d6565b3560206129db816129d58a89896127d6565b016127fd565b6129eb60606129d58b8a8a6127d6565b92612a618a86888a8c612a456080612a048784866127d6565b013595612a3d612a338260a0612a1b82888a6127d6565b01359560c0612a2b83838b6127d6565b0135976127d6565b60e0810190612807565b969095612839565b946040519a8b998a996326ae802b60e11b8b5260048b0161287e565b03815f305af19081612aa8575b50612a9d5785612a7c612950565b60405163f495148b60e01b8152918291612a99916004840161297f565b0390fd5b6001909501946129a9565b612ac89060203d8111612acd575b612ac081836105b8565b810190612849565b612a6e565b503d612ab6565b90612bbe939291612ae361388d565b612aeb6138b4565b803592612b00845f52600560205260405f2090565b91612b35612b236040830194612b1d612b18876127fd565b6103fb565b906138fb565b612b2f612b18866127fd565b90612d1c565b612b40343415612d44565b612bd4612b60865f526004602052600260405f2001600181540180915590565b96612b73602084013589819a8214612d62565b612b7f612b18866127fd565b93606084019688612b8f896127fd565b96612bcc8c60808901359e8f60a08b019b612baa8d8d612807565b929091604051978896602088019a8b612d80565b03601f1981018352826105b8565b51902061393e565b612be787612be1856127fd565b87613b43565b919092600184612bf681611ebc565b14612cc3575b50612c0683611ebc565b60018303612c605750505090612c32612c2c5f516020615c325f395f51905f52936127fd565b916127fd565b6040516001600160a01b03909216958291612c4f91429184612e0e565b0390a45b612c5b6138e9565b600190565b612c9d8394612c98612cbb947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612ca3956141eb565b6127fd565b926127fd565b9260405193849360018060a01b031698429185612de4565b0390a4612c53565b612cf991935088612cd3866127fd565b91612cf1612cea612ce38a6127fd565b9288612807565b3691610ca3565b928a8a613c55565b915f612bfc565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612d245750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612d4c5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612d6b575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610cf497959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161285e565b6105ea93606092969593608083019760018060a01b03168352602083015260408201520190611ecb565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615d525f395f51905f52602052600160405f20015490565b15612e5157565b63031206d560e51b5f5260045ffd5b80518210156127f85760209160051b010190565b15612e7c5750565b6373a1390160e11b5f5260045260245ffd5b959394612f10919597939297612ea261388d565b612ee76001600160a01b038816612eb9818b61449f565b612ec16138b4565b612ed1612b18612b1860e46127fd565b811490612ee1612b1860e46127fd565b9161306b565b612f08612ef8612b186101046127fd565b6001600160a01b038b1614613098565b83878961453f565b9161012435612f4381612f2c86612f2787876130c2565b6130c2565b811015612f3d87612f2788886130c2565b906130cf565b612f51612b186032546103fb565b90612f5d6101046127fd565b906101443592612f6e6101646130ed565b92610184356101a43591833b156103f757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af19889156130665761302d6111ab9861303693612c539c61304c575b5061302461300f6101046127fd565b916130186105fb565b9c8d5260208d016130f7565b60408b016130f7565b606089016130f7565b608087015260a086015260c08501523691610ca3565b8061305a5f613060936105b8565b80610976565b5f613000565b612945565b15613074575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b1561309f57565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610edb57565b156130d8575050565b63943892eb60e01b5f5260045260245260445ffd5b35610cf481610631565b6001600160a01b039091169052565b90600182811c92168015613134575b602083101461312057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613115565b9060405161314b8161059d565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916131a382613106565b808552916001811690811561320f57506001146131d0575b505060a092916131cc9103846105b8565b0152565b5f908152602081209092505b8183106131f35750508101602001816131cc6131bb565b60209193508060019154838589010152019101909184926131dc565b60ff191660208681019190915292151560051b850190920192508391506131cc90506131bb565b600a821015611ec65752565b9060405161324f81610582565b60406007829461325e8161313e565b845261327460ff60068301541660208601613236565b0154910152565b156132835750565b60405162461bcd60e51b815260206004820152908190612a999060248301906114f1565b156132b0575050565b637a88099160e11b5f5260045260245260445ffd5b604051906132d282610567565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906133058261061a565b61331260405191826105b8565b8281528092613323601f199161061a565b01905f5b82811061333357505050565b60209061333e6132c5565b82828501015201613327565b9060018060a01b03165f5260205260405f2090565b9060405161336c81610567565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916133b7906133ae906114b8565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103f75751610cf481610407565b5f5490949392906001600160a01b038116156134a157604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906134509060848701906114f1565b931691015203925af18015613066576105ea925f91613472575b508094613751565b613494915060203d60201161349a575b61348c81836105b8565b8101906133d2565b5f61346a565b503d613482565b6315aeca0d60e11b5f5260045ffd5b5f516020615e325f395f51905f52805460ff60401b1916600160401b179055565b5f516020615e325f395f51905f52805460ff60401b19169055565b156134f357565b63ff5a231360e01b5f5260045ffd5b9060ff801983541691151516179055565b9061351f8215156134ec565b6001600160a01b03169081156135da57606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b604051906135f86020836105b8565b5f808352366020840137565b1561360b57565b635e0c1f8960e01b5f5260045ffd5b91929493909384845114806136a6575b8061369c575b613639906127ac565b5f5b85811015613690578060051b8401359060be19853603018212156103f7576136896001926136698389612e60565b51613674848c612e60565b51906136808589612e60565b51928901612ad4565b500161363b565b50945050505050600190565b5081518514613630565b508486511461362a565b604051906136bd8261059d565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156136e757565b63f0f15b9160e01b5f5260045ffd5b156136fe5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b1561374257565b6302bfb33d60e51b5f5260045ffd5b91909161375d33614aac565b61376881151561373b565b6001600160a01b038316916137f69061378284151561373b565b6001600160a01b038116946125fc9061379c87151561373b565b6137a5856158aa565b613816575b6137c4816125aa816125a5895f52600560205260405f2090565b6137cf6125ba6105ec565b5f60408401525f60608401525f60808401525f60a08401525f60c08401526125f785612d0e565b6040515f81525f516020615df25f395f51905f529080602081015b0390a4565b6138486138216105db565b8681525f60208201525f60408201525f6060820152612638875f52600460205260405f2090565b6137aa565b916138669183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b0316330361387e57565b63c82cebcb60e01b5f5260045ffd5b60ff5f516020615d725f395f51905f5254166138a557565b63d93c066560e01b5f5260045ffd5b5f516020615db25f395f51905f525c6138da5760015f516020615db25f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615db25f395f51905f525d565b610cf4916001600160a01b0316906150a3565b1561391557565b63b3f07a3960e01b5f5260045ffd5b1561392c5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a9c575b6139579061390e565b6139786139725f516020615c125f395f51905f525460ff1690565b60ff1690565b956139868488811015613924565b5f945f935f5b8681106139a757505050505050506105ea9192811015613924565b613a046139d3836139b6615758565b6042916040519161190160f01b8352600283015260228201522090565b6139e76139e08489612e60565b5160ff1690565b6139f18487612e60565b51906139fd8589612e60565b51926150b6565b6001600160a01b038181169088161080613a35575b613a27575b5060010161398c565b600198890198909650613a1e565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615d525f395f51905f52602052613a97611684827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561334a565b613a19565b508551831461394e565b15613aad57565b6330d45fb160e01b5f5260045ffd5b908160209103126103f75751600a8110156103f75790565b6001600160a01b039091168152602081019190915260400190565b9150613b3060ff91613b1f5f94613b1160325460018060a01b03161515613aa6565b5f52600960205260405f2090565b6001600160a01b039091169061334a565b5416613b3b57600191565b506002905f90565b9190915f92613b84611684613b74613b5f612b186032546103fb565b94613b116001600160a01b0387161515613aa6565b6001600160a01b0384169061334a565b613c0b5791602091613bae95935f604051809881958294633f4bdec560e01b845260048401613ad4565b03925af1928315613066575f93613bda575b50600183613bcd81611ebc565b03613bd457565b60019150565b613bfd91935060203d602011613c04575b613bf581836105b8565b810190613abc565b915f613bc0565b503d613beb565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610cf4929101906114f1565b938093959295613c6486612d0e565b96613c876001613c7c818060a01b038416809b61334a565b015460a01c60ff1690565b93613c96612b186035546103fb565b9560188151101580613f33575b80613f29575b613ce4575b5050613cba94506150ce565b92613cc484611ebc565b60018414613cd3575b50505090565b613cdc9261517e565b5f8080613ccd565b602081015160601c5f808092604051613d1181612bbe602082019462483e5560e91b8652602483016111d1565b51908b5afa613d1e612950565b9080613f1d575b613f11575b5015613cae57909192935060018914159182613ee0575b5f905f808a8c89613d7a8a612bbe869a8c8814613eda5787965b6040519586948c6020870199631eeed51360e01b8b5260248801613c16565b51918c5af1613d87612950565b9080613ece575b613eb9575b50897f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613dca8682901515815260200190565b0390a3613e5d575090859291613de2575b8480613cae565b8315613dfd57613df883613cba96309084614b80565b613ddb565b91602090613e20956040519687928392632770a7eb60e21b845260048401613ad4565b03815f8b5af191821561306657613cba948693613e3e575b50613ddb565b613e569060203d602011612acd57612ac081836105b8565b505f613e38565b96613e6d9396929891945061517e565b83613e7d575b5050505050600190565b15613e9557613e8b93614b80565b5f80808080613e73565b5090613eb39250613ea4611db5565b916001600160a01b0316615216565b50613e8b565b9150915060406020820151910151915f613d93565b50604081511015613d8e565b86613d5b565b613eec868989876150ce565b613ef581611ebc565b60018103613f035750613d41565b995050505050505050505090565b6020915001515f613d2a565b50602081511015613d25565b50863b1515613ca9565b506001600160a01b0387161515613ca3565b15613f4d5750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103f75760405190613f788261059d565b819380358352602081013560208401526040810135613f9681610407565b6040840152613fa7606082016107d7565b60608401526080818101359084015260a0810135916001600160401b0383116103f75760a092613fd79201610cd9565b910152565b818110613fe7575050565b5f8155600101613fdc565b90601f8211613fff575050565b6105ea915f516020615bf25f395f51905f525f5260205f20906020601f840160051c83019310614037575b601f0160051c0190613fdc565b909150819061402a565b9190601f811161405057505050565b6105ea925f5260205f20906020601f840160051c8301931061403757601f0160051c0190613fdc565b8160011b915f199060031b1c19161790565b90600a811015611ec65760ff80198354169116179055565b8151805182556020810151600183015560408101519192916140d1906001600160a01b03166002850161278d565b60608101516140ec906001600160a01b03166003850161278d565b6080810151600484015560a00151805160058401916001600160401b038211610562576141238261411d8554613106565b85614041565b602090601f831160011461417b5782600795936040959361414c935f92614170575b5050614079565b90555b614169602082015161416081611ebc565b6006860161408b565b0151910155565b015190505f80614145565b90601f1983169161418f855f5260205f2090565b925f5b8181106141d357509260019285926007989660409896106141bb575b505050811b01905561414f565b01515f1960f88460031b161c191690555f80806141ae565b92936020600181928786015181550195019301614192565b9160806142856142766002936142718761426c61422561386699833595865f52600760205261422a60405f20602087013594858092615913565b613f45565b156142925761425e61423e600154426130c2565b915b61425361424b61060b565b963690613f5f565b865260208601613236565b6040840152610bef85612d00565b6140a3565b612d0e565b611060612b18604088016127fd565b93013592019182546130c2565b61425e5f91614240565b906142a7825f61527f565b91826142b05750565b5f80525f516020615cf25f395f51905f526020526001600160a01b03166142f7817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615913565b156142ff5750565b63d180cb3160e01b5f526004525f60245260445ffd5b919091614322838261527f565b928361432c575050565b815f525f516020615cf25f395f51905f5260205261435760405f209160018060a01b03168092615913565b15614360575050565b63d180cb3160e01b5f5260045260245260445ffd5b919091614382838261531c565b928361438c575050565b815f525f516020615cf25f395f51905f526020526143b760405f209160018060a01b031680926156b5565b156143c0575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143dc5750565b63290cd55f60e01b5f52600a811015611ec65760045260245ffd5b9061440191614f35565b60018060a01b036060820151168151905f516020615c325f395f51905f526020840191825194614451610a5f6040830196875160018060a01b03169885608086019a8b519260a088015194613c55565b519251935194516040516001600160a01b03909616959182916138119142919084612e0e565b1561447f5750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526144c38161228060405f2060018060a01b038316906150a3565b825f52600460205260ff600360405f200154166144fb5760ff60016144ef836125f76105ea9697612d0e565b015460a81c1615614477565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103f7578051916040602083015192015190565b1561453057565b6358e8878560e01b5f5260045ffd5b826060916145c897959693614559611066613b7484612d0e565b61456961193d6040830151151590565b61466a575b5061457d612b186032546103fb565b916145926001600160a01b0384161515613aa6565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115613066575f955f905f9361462c575b509082916105ea94939681988515159586614621575b505084614616575b50508261460b575b5050614529565b101590505f80614604565b101592505f806145fc565b101594505f806145f4565b90506146579196506105ea93925060603d606011614663575b61464f81836105b8565b81019061450e565b919692939192916145de565b503d614645565b60c061467c91015184808210156130cf565b5f61456e565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916146d591908601906114f1565b930152565b6146e48151614bc9565b6146ee8251612d0e565b5f516020615eb25f395f51905f52614720612b1860016147196020880195611060612b1888516103fb565b01546103fb565b9380519061381161473185516103fb565b91604081019061474182516103fb565b9061476a6080820196875160a084019485519861476460c087019a8b51906130c2565b936153ee565b614780614779825199516103fb565b93516103fb565b9460e061479060608401516103fb565b9751935191519201519260405197889760018060a01b03169c429689614682565b908160209103126103f7575190565b604051905f825f516020615bf25f395f51905f5254916147df83613106565b808352926001811690811561485b5750600114614803575b6105ea925003836105b8565b505f516020615bf25f395f51905f525f90815290915f516020615cb25f395f51905f525b81831061483f5750509060206105ea928201016147f7565b6020919350806001915483858901015201910190918492614827565b602092506105ea94915060ff191682840152151560051b8201016147f7565b604051905f825f516020615c725f395f51905f52549161489983613106565b808352926001811690811561485b57506001146148bc576105ea925003836105b8565b505f516020615c725f395f51905f525f90815290915f516020615e925f395f51905f525b8183106148f85750509060206105ea928201016147f7565b60209193508060019154838589010152019101909184926148e0565b5f516020615dd25f395f51905f525f525f516020615d525f395f51905f5260205260ff614961827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61334a565b54161561496b5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615dd25f395f51905f52602452604490fd5b5f516020615e125f395f51905f525f525f516020615d525f395f51905f5260205260ff6149e9827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361334a565b5416156149f35750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e125f395f51905f52602452604490fd5b5f516020615d325f395f51905f525f525f516020615d525f395f51905f5260205260ff614a71827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661334a565b541615614a7b5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d325f395f51905f52602452604490fd5b5f516020615d925f395f51905f525f525f516020615d525f395f51905f5260205260ff614af9827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61334a565b541615614b035750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d925f395f51905f52602452604490fd5b90815f525f516020615d525f395f51905f5260205260ff614b588260405f2061334a565b541615614b63575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105ea91614bc46084836105b8565b6155d1565b5f526004602052600160405f2001600181540180915590565b91614beb615629565b614bff6001600160a01b0384161515612777565b6001600160a01b03821692614c15841515612777565b60ff821615614e8757614c8190614c2a615629565b614c32615629565b614c3a615629565b60ff195f516020615d725f395f51905f5254165f516020615d725f395f51905f5255614c64615629565b614c6c615629565b614c74615629565b614c7c615629565b61429c565b50614c8a615629565b614c9860ff821615156136e0565b60408051614ca682826105b8565b60098152682b30b634b230ba37b960b91b6020820152614cc8825192836105b8565b60058252640312e302e360dc1b6020830152614ce2615629565b614cea615629565b8051906001600160401b03821161056257614d1b82614d165f516020615bf25f395f51905f5254613106565b613ff2565b602090601f8311600114614df35792614d4983614dd9979694614d5d94614daf975f92614170575050614079565b5f516020615bf25f395f51905f5255615adc565b614d725f5f516020615c925f395f51905f5255565b614d875f5f516020615e525f395f51905f5255565b60ff1660ff195f516020615c125f395f51905f525416175f516020615c125f395f51905f5255565b614db7615654565b603380546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615d125f395f51905f525f80a26105ea43603455565b5f516020615bf25f395f51905f525f52601f19831691905f516020615cb25f395f51905f52925f5b818110614e6f575093614d5d93614daf969360019383614dd99b9a9810614e57575b505050811b015f516020615bf25f395f51905f5255615adc565b01515f1960f88460031b161c191690555f8080614e3d565b92936020600181928786015181550195019301614e1b565b6338854f1160e21b5f5260045ffd5b91908203918211610edb57565b60075f9182815582600182015582600282015582600382015582600482015560058101614ed08154613106565b9081614ee3575b50508260068201550155565b601f8211600114614efa57849055505b5f80614ed7565b614f20614f30926001601f614f12855f5260205f2090565b920160051c82019101613fdc565b5f81815260208120918190559055565b614ef3565b9190614f3f6136b0565b50825f526007602052614f558160405f206156b5565b15614fbf57614fba6105ea91845f52600860205260405f20815f52602052610bef614f8260405f2061313e565b95614f9f614f8f82612d0e565b611060612b1860408b01516103fb565b614fb3600260808a01519201918254614e96565b9055612d00565b614ea3565b637f11bea960e01b5f5260045260245ffd5b600360606105ea93805184556020810151600185015560408101516002850155015115159101613502565b156150045750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c060049161503d60018060a01b038251168561278d565b602081015161508890600186019061505e906001600160a01b03168261278d565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061371e565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b91610cf493916150c5936157ac565b9092919261582e565b6001600160a01b03169291906001841461515d5781156151545761511d92156151285760405163a9059cbb60e01b602082015291615115918391612bbe9160248401613ad4565b6005926151c7565b15610cf45750600190565b6040516340c10f1960e01b60208201529161514c918391612bbe9160248401613ad4565b6006926151c7565b50505050600190565b90615170935061193d9250613ea4611db5565b61517957600190565b600590565b90615193915f52600660205260405f2061334a565b600181015460a01c60ff16156151b5576003018054918203918211610edb5755565b6004018054918201809211610edb5755565b81516001600160a01b03909116915f9182916020018285620186a0f16151eb612950565b901561521057805161520757503b1561520357600190565b5f90565b60209150015190565b50505f90565b8147106152705782516001600160a01b03909116925f9182916020018486620186a0f190615242612950565b91156152695715615255575b5050600190565b805161520757503b15615203575f8061524e565b5050505f90565b632b60b36f60e21b5f5260045ffd5b805f525f516020615d525f395f51905f5260205260ff6152a28360405f2061334a565b541661521057805f525f516020615d525f395f51905f526020526152c98260405f2061334a565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610cf4916001600160a01b031690615913565b805f525f516020615d525f395f51905f5260205260ff61533f8360405f2061334a565b54161561521057805f525f516020615d525f395f51905f526020526153678260405f2061334a565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b156153ae57505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156153df57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361545c57506105ea945061542461541682866130c2565b34143490612f3d84886130c2565b80615430575b5061598f565b615451615456916154426033546103fb565b9061544b611db5565b91615216565b6153d8565b5f61542a565b90615468343415612d44565b61547d61547582876130c2565b308489614b80565b80615511575b5061549961193d6001613c7c866125f787612d0e565b6154a9575b506105ea935061598f565b604051632770a7eb60e21b8152602081806154c8883060048401613ad4565b03815f885af1918215613066576105ea966154ec9387935f916154f2575b506153a4565b5f61549e565b61550b915060203d602011612acd57612ac081836105b8565b5f6154e6565b61552990615523612b186033546103fb565b8761595a565b5f615483565b90813b156155b0575f516020615cd25f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a280511561559857615595916159d8565b50565b5050346155a157565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b905f602091828151910182855af115612945575f513d61562057506001600160a01b0381163b155b6156005750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156155f9565b60ff5f516020615e325f395f51905f525460401c161561564557565b631afcd79f60e31b5f5260045ffd5b61565c615629565b62015180600155565b80548210156127f8575f5260205f2001905f90565b805480156156a1575f1901906156908282615665565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615750575f198401848111610edb5783545f19810194908511610edb575f95858361570397610bef9503615709575b50505061567a565b55600190565b6157396157339161572a6157206157479588615665565b90549060031b1c90565b92839187615665565b9061384d565b85905f5260205260405f2090565b555f80806156fb565b505050505f90565b6157606159f5565b615768615a4c565b6040519060208201925f516020615e725f395f51905f528452604083015260608201524660808201523060a082015260a081526157a660c0826105b8565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615819579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15613066575f516001600160a01b0381161561580f57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611ec657565b61583781615824565b80615840575050565b61584981615824565b600181036158605763f645eedf60e01b5f5260045ffd5b61586981615824565b60028103615884575063fce698f760e01b5f5260045260245ffd5b80615890600392615824565b146158985750565b6335e2f38360e21b5f5260045260245ffd5b5f8181526003602052604090205461590e57600254600160401b811015610562576158f76158e18260018594016002556002615665565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61591d82826150a3565b61521057805490600160401b82101561056257826159456158e1846001809601855584615665565b90558054925f520160205260405f2055600190565b614bc46105ea939261598160405194859263a9059cbb60e01b602085015260248401613ad4565b03601f1981018452836105b8565b906159a4915f52600660205260405f2061334a565b600181015460a01c60ff16156159c6576003018054918201809211610edb5755565b6004018054918203918211610edb5755565b5f80610cf493602081519101845af46159ef612950565b91615a7e565b6159fd6147c0565b8051908115615a0d576020012090565b50505f516020615c925f395f51905f52548015615a275790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615a5461487a565b8051908115615a64576020012090565b50505f516020615e525f395f51905f52548015615a275790565b90615aa25750805115615a9357805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615ad3575b615ab3575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615aab565b80519091906001600160401b03811161056257615b1d81615b0a5f516020615c725f395f51905f5254613106565b5f516020615c725f395f51905f52614041565b602092601f8211600114615b5057615b3f929382915f92614170575050614079565b5f516020615c725f395f51905f5255565b5f516020615c725f395f51905f525f52601f198216935f516020615e925f395f51905f52915f5b868110615bb95750836001959610615ba1575b505050811b015f516020615c725f395f51905f5255565b01515f1960f88460031b161c191690555f8080615b8a565b91926020600181928685015181550194019201615b7756feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a26469706673582212208e3a77885b6216d79da8edc5a587eb28cca0b9393c40bef7960752dddc11fea164736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeV2MetaData.ABI instead.
var BSCBridgeV2ABI = BSCBridgeV2MetaData.ABI

// BSCBridgeV2BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BSCBridgeV2BinRuntime = "60806040526004361015610022575b3615610018575f80fd5b61002061386a565b005b5f3560e01c806301ffc9a7146103a15780630b1d8d061461039c5780631313996b146103975780631938e0f214610392578063248a9ca31461038d5780632f2ff15d1461038857806336568abe14610383578063365d71e41461037e57806342cde4e81461037957806348a00ed8146103745780634be13f831461036f5780634d5d00561461036a5780634ee078ba146103655780634f1ef28614610360578063502cc5c01461035b57806352d1902d146103565780635b605f5c146103515780635c975abb1461034c5780635fd262de14610347578063670e626814610342578063761fe2d81461033d57806379214874146103385780637b429ad814610333578063814914b51461032e57806384b0196e1461032957806386e9e1751461032457806388d67d6d1461031f57806389232a001461031a5780638ae87c5c1461031557806391cca3db1461031057806391cf6d3e1461030b57806391d14854146103065780639f9b4f1c14610301578063a10bab0b146102fc578063a217fddf146102f7578063a3246ad3146102f2578063aa1bd0bc146102ed578063aa20ed47146102e8578063ad3cb1cc146102e3578063ae6893f8146102de578063b2b49e2e146102d9578063b33eb36e146102d4578063b7f3358d146102cf578063b8aa837e146102ca578063bedb86fb146102c5578063bfbf6765146102c0578063cba25e4b146102bb578063cf56118e146102b6578063d477f05f146102b1578063d547741f146102ac578063d5717fc5146102a7578063e2af6cd7146102a2578063edbbea391461029d578063f0702e8e14610298578063f4509643146102935763f698da250361000e5761270b565b61266a565b612642565b6124ea565b612471565b612438565b612409565b6123aa565b612336565b6122cb565b6121f2565b612108565b61206c565b611fe8565b611f57565b611e49565b611e10565b611dc9565b611d3e565b611cf1565b611c75565b611c19565b611bf1565b611ade565b611a9c565b611a7f565b611a57565b6119ee565b6118e9565b6117df565b611628565b611574565b611427565b611351565b6112e1565b61125b565b6111e4565b6110dc565b6110ae565b610fcf565b610ee0565b610e50565b610cf7565b610b97565b610b2a565b610ad6565b6109ac565b610980565b61090d565b61081d565b6107e2565b6107b1565b610700565b6104ba565b610418565b346103f75760203660031901126103f75760043563ffffffff60e01b81168091036103f757602090637965db0b60e01b81149081156103e6575b506040519015158152f35b6301ffc9a760e01b1490505f6103db565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103f757565b346103f75760203660031901126103f75760043561043581610407565b61043e33614914565b6001600160a01b0316610452811515612777565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103f7578235916001600160401b0383116103f7576020808501948460051b0101116103f757565b60403660031901126103f7576004356001600160401b0381116103f7576104e590369060040161048a565b602435916001600160401b0383116103f757366023840112156103f7576004830135916001600160401b0383116103f75736602460e08502860101116103f7576024610020940191612996565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761056257604052565b610532565b60e081019081106001600160401b0382111761056257604052565b606081019081106001600160401b0382111761056257604052565b60c081019081106001600160401b0382111761056257604052565b601f909101601f19168101906001600160401b0382119082101761056257604052565b604051906105ea6080836105b8565b565b604051906105ea60e0836105b8565b604051906105ea610100836105b8565b604051906105ea6060836105b8565b6001600160401b0381116105625760051b60200190565b60ff8116036103f757565b9080601f830112156103f75781356106538161061a565b9261066160405194856105b8565b81845260208085019260051b8201019283116103f757602001905b8282106106895750505090565b60208091833561069881610631565b81520191019061067c565b9080601f830112156103f75781356106ba8161061a565b926106c860405194856105b8565b81845260208085019260051b8201019283116103f757602001905b8282106106f05750505090565b81358152602091820191016106e3565b60803660031901126103f7576004356001600160401b0381116103f75760c060031982360301126103f7576024356001600160401b0381116103f75761074a90369060040161063c565b906044356001600160401b0381116103f75761076a9036906004016106a3565b60643591906001600160401b0383116103f7576107ad9361079261079b9436906004016106a3565b92600401612ad4565b60405191829182901515815260200190565b0390f35b346103f75760203660031901126103f75760206107cf600435612e2c565b604051908152f35b35906105ea82610407565b346103f75760403660031901126103f75761002060243560043561080582610407565b61081861081182612e2c565b3390614b34565b614315565b346103f75760403660031901126103f75760043560243561083d81610407565b336001600160a01b038216036108565761002091614375565b63334bd91960e11b5f5260045ffd5b9060406003198301126103f7576004356001600160401b0381116103f75782610890916004016106a3565b91602435906001600160401b0382116103f757806023830112156103f75781600401356108bc8161061a565b926108ca60405194856105b8565b8184526024602085019260051b8201019283116103f757602401905b8282106108f35750505090565b60208091833561090281610407565b8152019101906108e6565b346103f75761091b36610865565b906109298151835114612e4a565b5f5b8251811015610020578061096f61094460019385612e60565b51838060a01b036109558488612e60565b51169061096a3361096583612e2c565b614b34565b614375565b500161092b565b5f9103126103f757565b346103f7575f3660031901126103f757602060ff5f516020615c125f395f51905f525416604051908152f35b346103f75760603660031901126103f7576024356004356044356109cf81610407565b6109d83361499c565b6109e06138b4565b815f5260076020526109fe836109f98160405f206150a3565b612e74565b80610a098484614f35565b916001600160a01b031615610ac2575b8151905f516020615c325f395f51905f526020840191610a7b835195610a71610a5f6040830197885160018060a01b03169986608086019b8c519260a088015194613c55565b610a6881611ebc565b600181146143d4565b51935194516103fb565b94516040516001600160a01b0390961695918291610a9c9142919084612e0e565b0390a45f516020615bd25f395f51905f525f80a35f5f516020615db25f395f51905f525d005b5060608101516001600160a01b0316610a19565b346103f7575f3660031901126103f7575f546040516001600160a01b039091168152602090f35b9181601f840112156103f7578235916001600160401b0383116103f757602083818601950101116103f757565b6101c03660031901126103f757602435600435610b4682610407565b604435610b5281610407565b6064359060a43560843560c4356001600160401b0381116103f757610b7b903690600401610afd565b94909360e03660e31901126103f7576107ad9761079b97612e8e565b346103f75760403660031901126103f757610c76602435600435610bb961388d565b610bc16138b4565b805f526007602052610bda826109f98160405f206150a3565b610c716040610c01610bfc85610bef86612d00565b905f5260205260405f2090565b613242565b610c5e610c2182516080610c17868301516103fb565b9101519087613aef565b50610c2b81611ebc565b610c3481611ebc565b83516020810182905290600190610c5883604081015b03601f1981018552846105b8565b1461327b565b01518015908115610c7e575b42916132a7565b6143f7565b6100206138e9565b4281109150610c6a565b6001600160401b03811161056257601f01601f191660200190565b929192610caf82610c88565b91610cbd60405193846105b8565b8294818452818301116103f7578281602093845f960137010152565b9080601f830112156103f757816020610cf493359101610ca3565b90565b60403660031901126103f757600435610d0f81610407565b6024356001600160401b0381116103f757610d2e903690600401610cd9565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610e2e575b50610e1f57610d7233614914565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610dee575b50610dbb57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615cd25f395f51905f528303610dda57610020925061552f565b632a87526960e21b5f52600483905260245ffd5b610e1191945060203d602011610e18575b610e0981836105b8565b8101906147b1565b925f610d9a565b503d610dff565b63703e46dd60e11b5f5260045ffd5b5f516020615cd25f395f51905f52546001600160a01b0316141590505f610d64565b346103f75760603660031901126103f757602435600435604435610e733361499c565b815f526007602052610e8c836109f98160405f206150a3565b4201804211610edb5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b6130ae565b346103f7575f3660031901126103f7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e1f5760206040515f516020615cd25f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610faf5750505090565b909192602060e082610fc46001948851610f37565b019401929101610fa2565b346103f75760203660031901126103f757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061109557505061101e925003836105b8565b61102882516132fb565b915f5b81518110156110875760019061106b61106661104686612d0e565b6110606110538588612e60565b516001600160a01b031690565b9061334a565b61335f565b6110758287612e60565b526110808186612e60565b500161102b565b604051806107ad8682610f8c565b8454835260019485019487945060209093019201611009565b346103f7575f3660031901126103f757602060ff5f516020615d725f395f51905f5254166040519015158152f35b60e03660031901126103f7576024356004356110f782610407565b6044359161110483610407565b6064359260c435916084359160a435916001600160401b0385116103f7576111b59661116a61113a6111ab973690600401610afd565b95909661114561388d565b6001600160a01b03851694849061115c878d61449f565b6111646138b4565b8b61453f565b9390926040519861117a8a610546565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610ca3565b60e08201526146da565b5f5f516020615db25f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103f75760803660031901126103f75760243560043561120482610407565b604435906001600160401b0382116103f757366023830112156103f7576107ad9261123c61124f933690602481600401359101610ca3565b906064359261124a84610631565b6133e7565b604051918291826111d1565b346103f75760403660031901126103f757602060ff61129160243560043561128282610407565b5f526009845260405f2061334a565b54166040519015158152f35b90602080835192838152019201905f5b8181106112ba5750505090565b82518452602093840193909201916001016112ad565b906020610cf492818152019061129d565b346103f75760203660031901126103f7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061133b576107ad8561132f818703826105b8565b604051918291826112d0565b8254845260209093019260019283019201611318565b346103f75760403660031901126103f75760043560243561137181610407565b61137a33614914565b5f516020615e325f395f51905f52549160ff8360401c168015611403575b6113f4576001600160401b03199092166002175f516020615e325f395f51905f52556113cb916113c66134b0565b613513565b6113d36134d1565b604051600281525f516020615c525f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b0384161015611398565b60e0810192916105ea9190610f37565b346103f75760403660031901126103f7576107ad61146660243560043561144d82610407565b6114556132c5565b505f52600660205260405f2061334a565b60046040519161147583610567565b80546001600160a01b0316835260018101546114ca906114c1906114a461149b826103fb565b602088016130f7565b6114b860a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611417565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161154a9061153c610cf497959693600f60f81b865260e0602087015260e08601906114f1565b9084820360408601526114f1565b60608301949094526001600160a01b031660808201525f60a082015280830360c09091015261129d565b346103f7575f3660031901126103f7575f516020615c925f395f51905f52541580611608575b156115cb576115a76147c0565b6115af61487a565b906107ad6115bb6135e9565b6040519384933091469186611515565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615e525f395f51905f52541561159a565b801515036103f757565b346103f75760603660031901126103f75760043561164581610407565b6024356044356116548161161e565b61165c61388d565b6116646138b4565b6001600160a01b0383165f9081526066602052604090206116909061168b905b5460ff1690565b613604565b61169b8215156134ec565b611752576116b681836116af6065546103fb565b3390614b80565b335b5f516020615eb25f395f51905f526116d1606454614bc9565b926064549261173b6116e46065546103fb565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a46117466138e9565b60405160018152602090f35b61175b33614914565b306116b8565b9080601f830112156103f75781356117788161061a565b9261178660405194856105b8565b81845260208085019260051b820101918383116103f75760208201905b8382106117b257505050505090565b81356001600160401b0381116103f7576020916117d4878480948801016106a3565b8152019101906117a3565b60803660031901126103f7576004356001600160401b0381116103f75761180a90369060040161048a565b90602435906001600160401b0382116103f757366023830112156103f75781600401356118368161061a565b9261184460405194856105b8565b8184526024602085019260051b820101903682116103f75760248101925b8284106118ba5750506044359150506001600160401b0381116103f75761188d903690600401611761565b606435926001600160401b0384116103f7576107ad946118b461079b953690600401611761565b9361361a565b83356001600160401b0381116103f7576020916118de83926024369187010161063c565b815201930192611862565b346103f75760603660031901126103f75760043561190681610407565b6024359061191382610407565b60443561191f81610631565b5f516020615e325f395f51905f52549260ff604085901c1615611941565b1590565b936001600160401b0316801590816119e6575b60011490816119dc575b1590816119d3575b506113f4575f516020615e325f395f51905f5280546001600160401b031916600117905561199892846119c657614be2565b61199e57005b6119a66134d1565b604051600181525f516020615c525f395f51905f529080602081016113ef565b6119ce6134b0565b614be2565b9050155f611966565b303b15915061195e565b859150611954565b346103f75760403660031901126103f757602435600435611a0e33614914565b611a166138b4565b611a208282614f35565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615db25f395f51905f525d005b346103f7575f3660031901126103f7576033546040516001600160a01b039091168152602090f35b346103f7575f3660031901126103f7576020603454604051908152f35b346103f75760403660031901126103f757602060ff611291602435600435611ac382610407565b5f525f516020615d525f395f51905f52845260405f2061334a565b346103f75760403660031901126103f7576004356001600160401b0381116103f757611b0e9036906004016106a3565b6024356001600160401b0381116103f757611b2d9036906004016106a3565b90611b3b81518351146127ac565b5f5b82518110156100205780611be3611b5660019385612e60565b51611b618387612e60565b5190611b6b61388d565b611b736138b4565b805f526007602052611b8c826109f98160405f206150a3565b610c716040611ba1610bfc85610bef86612d00565b610c5e611bb782516080610c17868301516103fb565b50611bc181611ebc565b611bca81611ebc565b835160208101829052908a90610c588360408101610c4a565b611beb6138e9565b01611b3d565b346103f7575f3660031901126103f7576035546040516001600160a01b039091168152602090f35b346103f7575f3660031901126103f75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611c565750505090565b82516001600160a01b0316845260209384019390920191600101611c49565b346103f75760203660031901126103f7576004355f525f516020615cf25f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611cdb576107ad85611ccf818703826105b8565b60405191829182611c33565b8254845260209093019260019283019201611cb8565b346103f75760203660031901126103f7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611d3133614914565b80600155604051908152a1005b346103f75760403660031901126103f757602435600435611d5e3361499c565b611d673361499c565b611d6f6138b4565b805f526007602052611d88826109f98160405f206150a3565b611d9282826143f7565b5f516020615bd25f395f51905f525f80a35f5f516020615db25f395f51905f525d005b60405190611dc46020836105b8565b5f8252565b346103f7575f3660031901126103f7576107ad604051611dea6040826105b8565b60058152640352e302e360dc1b60208201526040519182916020835260208301906114f1565b346103f75760203660031901126103f7576004355f526004602052600160405f20015460018101809111610edb57602090604051908152f35b346103f757611e5736610865565b90611e658151835114612e4a565b5f5b82518110156100205780611ea1611e8060019385612e60565b51838060a01b03611e918488612e60565b5116906108183361096583612e2c565b5001611e67565b634e487b7160e01b5f52602160045260245ffd5b600a1115611ec657565b611ea8565b90600a821015611ec65752565b6020815260606040611f3d60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c06101208601526101408501906114f1565b93611f4f602082015183860190611ecb565b015191015290565b346103f75760403660031901126103f757600435602435905f60408051611f7d81610582565b611f856136b0565b815282602082015201525f52600860205260405f20905f526020526107ad60405f20600760405191611fb683610582565b611fbf8161313e565b8352611fd560ff60068301541660208501613236565b0154604082015260405191829182611ed8565b346103f75760203660031901126103f7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff60043561202a81610631565b61203333614914565b1661203f8115156136e0565b8060ff195f516020615c125f395f51905f525416175f516020615c125f395f51905f5255604051908152a1005b346103f75760403660031901126103f75760043560243561208c8161161e565b61209533614a24565b6120aa825f52600360205260405f2054151590565b156120f55760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526120ea81600360405f2001613502565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103f75760203660031901126103f7576004356121258161161e565b61212e33614a24565b1561218c5761213b61388d565b600160ff195f516020615d725f395f51905f525416175f516020615d725f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615d725f395f51905f525460ff8116156121e35760ff19165f516020615d725f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103f75760803660031901126103f75760243560043561221282610407565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356122418161161e565b60643561224d8161161e565b61225633614a24565b845f5260056020526122ba816122b5855f20986122858161228060018060a01b038216809d6150a3565b6136f6565b885f5260066020526122a586600161229f848b5f2061334a565b0161371e565b885f526009602052865f2061334a565b613502565b8251911515825215156020820152a3005b346103f75760203660031901126103f7576004356122e881610407565b6122f133614914565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103f7575f3660031901126103f757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b818110612394576107ad8561132f818703826105b8565b825484526020909301926001928301920161237d565b346103f75760203660031901126103f7576004356123c781610407565b6123d033614914565b6001600160a01b03166123e4811515612777565b603380546001600160a01b031916821790555f516020615d125f395f51905f525f80a2005b346103f75760403660031901126103f75761002060243560043561242c82610407565b61096a61081182612e2c565b346103f75760203660031901126103f7576004355f526004602052600260405f20015460018101809111610edb57602090604051908152f35b346103f75760203660031901126103f75760043561248e81610407565b61249733614914565b6001600160a01b03166124ab81151561373b565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346103f75760803660031901126103f75760043560243561250a8161161e565b6044359161251783610407565b5f516020615df25f395f51905f526126016064359361253585610407565b61253e33614aac565b61254984151561373b565b6001600160a01b0386169461079b9061256387151561373b565b6001600160a01b038116976125fc9061257d8a151561373b565b612586886158aa565b612606575b6125af816125aa816125a58c5f52600560205260405f2090565b615309565b614ffc565b6125ce6125ba6105ec565b936125c583866130f7565b602085016130f7565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526125f788612d0e565b61334a565b615024565b0390a4005b61263d6126116105db565b8981525f60208201525f60408201525f60608201526126388a5f52600460205260405f2090565b614fd1565b61258b565b346103f7575f3660031901126103f7576032546040516001600160a01b039091168152602090f35b346103f75760403660031901126103f75760243560043561268a82610407565b61269333614914565b805f5260056020525f60046126ce60408320946126bd8161228060018060a01b03821680996156b5565b84845260066020526040842061334a565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103f7575f3660031901126103f7576127236159f5565b61272b615a4c565b6040519060208201925f516020615e725f395f51905f528452604083015260608201524660808201523060a082015260a0815261276960c0826105b8565b519020604051908152602090f35b1561277e57565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b156127b357565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156127f85760051b8101359060fe19813603018212156103f7570190565b6127c2565b35610cf481610407565b903590601e19813603018212156103f757018035906001600160401b0382116103f7576020019181360383136103f757565b91908110156127f85760e0020190565b908160209103126103f75751610cf48161161e565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936128cf97939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161285e565b9480356128db81610407565b6001600160a01b031660e085015260208101356128f781610407565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561292c81610631565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561297a573d9061296182610c88565b9161296f60405193846105b8565b82523d5f602084013e565b606090565b604090610cf49392815281602082015201906114f1565b90919392936129a68584146127ac565b5f945b8386106129b857505050509050565b6129c38685856127d6565b3560206129db816129d58a89896127d6565b016127fd565b6129eb60606129d58b8a8a6127d6565b92612a618a86888a8c612a456080612a048784866127d6565b013595612a3d612a338260a0612a1b82888a6127d6565b01359560c0612a2b83838b6127d6565b0135976127d6565b60e0810190612807565b969095612839565b946040519a8b998a996326ae802b60e11b8b5260048b0161287e565b03815f305af19081612aa8575b50612a9d5785612a7c612950565b60405163f495148b60e01b8152918291612a99916004840161297f565b0390fd5b6001909501946129a9565b612ac89060203d8111612acd575b612ac081836105b8565b810190612849565b612a6e565b503d612ab6565b90612bbe939291612ae361388d565b612aeb6138b4565b803592612b00845f52600560205260405f2090565b91612b35612b236040830194612b1d612b18876127fd565b6103fb565b906138fb565b612b2f612b18866127fd565b90612d1c565b612b40343415612d44565b612bd4612b60865f526004602052600260405f2001600181540180915590565b96612b73602084013589819a8214612d62565b612b7f612b18866127fd565b93606084019688612b8f896127fd565b96612bcc8c60808901359e8f60a08b019b612baa8d8d612807565b929091604051978896602088019a8b612d80565b03601f1981018352826105b8565b51902061393e565b612be787612be1856127fd565b87613b43565b919092600184612bf681611ebc565b14612cc3575b50612c0683611ebc565b60018303612c605750505090612c32612c2c5f516020615c325f395f51905f52936127fd565b916127fd565b6040516001600160a01b03909216958291612c4f91429184612e0e565b0390a45b612c5b6138e9565b600190565b612c9d8394612c98612cbb947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612ca3956141eb565b6127fd565b926127fd565b9260405193849360018060a01b031698429185612de4565b0390a4612c53565b612cf991935088612cd3866127fd565b91612cf1612cea612ce38a6127fd565b9288612807565b3691610ca3565b928a8a613c55565b915f612bfc565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612d245750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612d4c5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612d6b575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610cf497959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161285e565b6105ea93606092969593608083019760018060a01b03168352602083015260408201520190611ecb565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615d525f395f51905f52602052600160405f20015490565b15612e5157565b63031206d560e51b5f5260045ffd5b80518210156127f85760209160051b010190565b15612e7c5750565b6373a1390160e11b5f5260045260245ffd5b959394612f10919597939297612ea261388d565b612ee76001600160a01b038816612eb9818b61449f565b612ec16138b4565b612ed1612b18612b1860e46127fd565b811490612ee1612b1860e46127fd565b9161306b565b612f08612ef8612b186101046127fd565b6001600160a01b038b1614613098565b83878961453f565b9161012435612f4381612f2c86612f2787876130c2565b6130c2565b811015612f3d87612f2788886130c2565b906130cf565b612f51612b186032546103fb565b90612f5d6101046127fd565b906101443592612f6e6101646130ed565b92610184356101a43591833b156103f757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af19889156130665761302d6111ab9861303693612c539c61304c575b5061302461300f6101046127fd565b916130186105fb565b9c8d5260208d016130f7565b60408b016130f7565b606089016130f7565b608087015260a086015260c08501523691610ca3565b8061305a5f613060936105b8565b80610976565b5f613000565b612945565b15613074575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b1561309f57565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610edb57565b156130d8575050565b63943892eb60e01b5f5260045260245260445ffd5b35610cf481610631565b6001600160a01b039091169052565b90600182811c92168015613134575b602083101461312057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613115565b9060405161314b8161059d565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916131a382613106565b808552916001811690811561320f57506001146131d0575b505060a092916131cc9103846105b8565b0152565b5f908152602081209092505b8183106131f35750508101602001816131cc6131bb565b60209193508060019154838589010152019101909184926131dc565b60ff191660208681019190915292151560051b850190920192508391506131cc90506131bb565b600a821015611ec65752565b9060405161324f81610582565b60406007829461325e8161313e565b845261327460ff60068301541660208601613236565b0154910152565b156132835750565b60405162461bcd60e51b815260206004820152908190612a999060248301906114f1565b156132b0575050565b637a88099160e11b5f5260045260245260445ffd5b604051906132d282610567565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906133058261061a565b61331260405191826105b8565b8281528092613323601f199161061a565b01905f5b82811061333357505050565b60209061333e6132c5565b82828501015201613327565b9060018060a01b03165f5260205260405f2090565b9060405161336c81610567565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916133b7906133ae906114b8565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103f75751610cf481610407565b5f5490949392906001600160a01b038116156134a157604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906134509060848701906114f1565b931691015203925af18015613066576105ea925f91613472575b508094613751565b613494915060203d60201161349a575b61348c81836105b8565b8101906133d2565b5f61346a565b503d613482565b6315aeca0d60e11b5f5260045ffd5b5f516020615e325f395f51905f52805460ff60401b1916600160401b179055565b5f516020615e325f395f51905f52805460ff60401b19169055565b156134f357565b63ff5a231360e01b5f5260045ffd5b9060ff801983541691151516179055565b9061351f8215156134ec565b6001600160a01b03169081156135da57606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b604051906135f86020836105b8565b5f808352366020840137565b1561360b57565b635e0c1f8960e01b5f5260045ffd5b91929493909384845114806136a6575b8061369c575b613639906127ac565b5f5b85811015613690578060051b8401359060be19853603018212156103f7576136896001926136698389612e60565b51613674848c612e60565b51906136808589612e60565b51928901612ad4565b500161363b565b50945050505050600190565b5081518514613630565b508486511461362a565b604051906136bd8261059d565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156136e757565b63f0f15b9160e01b5f5260045ffd5b156136fe5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b1561374257565b6302bfb33d60e51b5f5260045ffd5b91909161375d33614aac565b61376881151561373b565b6001600160a01b038316916137f69061378284151561373b565b6001600160a01b038116946125fc9061379c87151561373b565b6137a5856158aa565b613816575b6137c4816125aa816125a5895f52600560205260405f2090565b6137cf6125ba6105ec565b5f60408401525f60608401525f60808401525f60a08401525f60c08401526125f785612d0e565b6040515f81525f516020615df25f395f51905f529080602081015b0390a4565b6138486138216105db565b8681525f60208201525f60408201525f6060820152612638875f52600460205260405f2090565b6137aa565b916138669183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b0316330361387e57565b63c82cebcb60e01b5f5260045ffd5b60ff5f516020615d725f395f51905f5254166138a557565b63d93c066560e01b5f5260045ffd5b5f516020615db25f395f51905f525c6138da5760015f516020615db25f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615db25f395f51905f525d565b610cf4916001600160a01b0316906150a3565b1561391557565b63b3f07a3960e01b5f5260045ffd5b1561392c5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a9c575b6139579061390e565b6139786139725f516020615c125f395f51905f525460ff1690565b60ff1690565b956139868488811015613924565b5f945f935f5b8681106139a757505050505050506105ea9192811015613924565b613a046139d3836139b6615758565b6042916040519161190160f01b8352600283015260228201522090565b6139e76139e08489612e60565b5160ff1690565b6139f18487612e60565b51906139fd8589612e60565b51926150b6565b6001600160a01b038181169088161080613a35575b613a27575b5060010161398c565b600198890198909650613a1e565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615d525f395f51905f52602052613a97611684827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561334a565b613a19565b508551831461394e565b15613aad57565b6330d45fb160e01b5f5260045ffd5b908160209103126103f75751600a8110156103f75790565b6001600160a01b039091168152602081019190915260400190565b9150613b3060ff91613b1f5f94613b1160325460018060a01b03161515613aa6565b5f52600960205260405f2090565b6001600160a01b039091169061334a565b5416613b3b57600191565b506002905f90565b9190915f92613b84611684613b74613b5f612b186032546103fb565b94613b116001600160a01b0387161515613aa6565b6001600160a01b0384169061334a565b613c0b5791602091613bae95935f604051809881958294633f4bdec560e01b845260048401613ad4565b03925af1928315613066575f93613bda575b50600183613bcd81611ebc565b03613bd457565b60019150565b613bfd91935060203d602011613c04575b613bf581836105b8565b810190613abc565b915f613bc0565b503d613beb565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610cf4929101906114f1565b938093959295613c6486612d0e565b96613c876001613c7c818060a01b038416809b61334a565b015460a01c60ff1690565b93613c96612b186035546103fb565b9560188151101580613f33575b80613f29575b613ce4575b5050613cba94506150ce565b92613cc484611ebc565b60018414613cd3575b50505090565b613cdc9261517e565b5f8080613ccd565b602081015160601c5f808092604051613d1181612bbe602082019462483e5560e91b8652602483016111d1565b51908b5afa613d1e612950565b9080613f1d575b613f11575b5015613cae57909192935060018914159182613ee0575b5f905f808a8c89613d7a8a612bbe869a8c8814613eda5787965b6040519586948c6020870199631eeed51360e01b8b5260248801613c16565b51918c5af1613d87612950565b9080613ece575b613eb9575b50897f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613dca8682901515815260200190565b0390a3613e5d575090859291613de2575b8480613cae565b8315613dfd57613df883613cba96309084614b80565b613ddb565b91602090613e20956040519687928392632770a7eb60e21b845260048401613ad4565b03815f8b5af191821561306657613cba948693613e3e575b50613ddb565b613e569060203d602011612acd57612ac081836105b8565b505f613e38565b96613e6d9396929891945061517e565b83613e7d575b5050505050600190565b15613e9557613e8b93614b80565b5f80808080613e73565b5090613eb39250613ea4611db5565b916001600160a01b0316615216565b50613e8b565b9150915060406020820151910151915f613d93565b50604081511015613d8e565b86613d5b565b613eec868989876150ce565b613ef581611ebc565b60018103613f035750613d41565b995050505050505050505090565b6020915001515f613d2a565b50602081511015613d25565b50863b1515613ca9565b506001600160a01b0387161515613ca3565b15613f4d5750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103f75760405190613f788261059d565b819380358352602081013560208401526040810135613f9681610407565b6040840152613fa7606082016107d7565b60608401526080818101359084015260a0810135916001600160401b0383116103f75760a092613fd79201610cd9565b910152565b818110613fe7575050565b5f8155600101613fdc565b90601f8211613fff575050565b6105ea915f516020615bf25f395f51905f525f5260205f20906020601f840160051c83019310614037575b601f0160051c0190613fdc565b909150819061402a565b9190601f811161405057505050565b6105ea925f5260205f20906020601f840160051c8301931061403757601f0160051c0190613fdc565b8160011b915f199060031b1c19161790565b90600a811015611ec65760ff80198354169116179055565b8151805182556020810151600183015560408101519192916140d1906001600160a01b03166002850161278d565b60608101516140ec906001600160a01b03166003850161278d565b6080810151600484015560a00151805160058401916001600160401b038211610562576141238261411d8554613106565b85614041565b602090601f831160011461417b5782600795936040959361414c935f92614170575b5050614079565b90555b614169602082015161416081611ebc565b6006860161408b565b0151910155565b015190505f80614145565b90601f1983169161418f855f5260205f2090565b925f5b8181106141d357509260019285926007989660409896106141bb575b505050811b01905561414f565b01515f1960f88460031b161c191690555f80806141ae565b92936020600181928786015181550195019301614192565b9160806142856142766002936142718761426c61422561386699833595865f52600760205261422a60405f20602087013594858092615913565b613f45565b156142925761425e61423e600154426130c2565b915b61425361424b61060b565b963690613f5f565b865260208601613236565b6040840152610bef85612d00565b6140a3565b612d0e565b611060612b18604088016127fd565b93013592019182546130c2565b61425e5f91614240565b906142a7825f61527f565b91826142b05750565b5f80525f516020615cf25f395f51905f526020526001600160a01b03166142f7817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615913565b156142ff5750565b63d180cb3160e01b5f526004525f60245260445ffd5b919091614322838261527f565b928361432c575050565b815f525f516020615cf25f395f51905f5260205261435760405f209160018060a01b03168092615913565b15614360575050565b63d180cb3160e01b5f5260045260245260445ffd5b919091614382838261531c565b928361438c575050565b815f525f516020615cf25f395f51905f526020526143b760405f209160018060a01b031680926156b5565b156143c0575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143dc5750565b63290cd55f60e01b5f52600a811015611ec65760045260245ffd5b9061440191614f35565b60018060a01b036060820151168151905f516020615c325f395f51905f526020840191825194614451610a5f6040830196875160018060a01b03169885608086019a8b519260a088015194613c55565b519251935194516040516001600160a01b03909616959182916138119142919084612e0e565b1561447f5750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526144c38161228060405f2060018060a01b038316906150a3565b825f52600460205260ff600360405f200154166144fb5760ff60016144ef836125f76105ea9697612d0e565b015460a81c1615614477565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103f7578051916040602083015192015190565b1561453057565b6358e8878560e01b5f5260045ffd5b826060916145c897959693614559611066613b7484612d0e565b61456961193d6040830151151590565b61466a575b5061457d612b186032546103fb565b916145926001600160a01b0384161515613aa6565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115613066575f955f905f9361462c575b509082916105ea94939681988515159586614621575b505084614616575b50508261460b575b5050614529565b101590505f80614604565b101592505f806145fc565b101594505f806145f4565b90506146579196506105ea93925060603d606011614663575b61464f81836105b8565b81019061450e565b919692939192916145de565b503d614645565b60c061467c91015184808210156130cf565b5f61456e565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916146d591908601906114f1565b930152565b6146e48151614bc9565b6146ee8251612d0e565b5f516020615eb25f395f51905f52614720612b1860016147196020880195611060612b1888516103fb565b01546103fb565b9380519061381161473185516103fb565b91604081019061474182516103fb565b9061476a6080820196875160a084019485519861476460c087019a8b51906130c2565b936153ee565b614780614779825199516103fb565b93516103fb565b9460e061479060608401516103fb565b9751935191519201519260405197889760018060a01b03169c429689614682565b908160209103126103f7575190565b604051905f825f516020615bf25f395f51905f5254916147df83613106565b808352926001811690811561485b5750600114614803575b6105ea925003836105b8565b505f516020615bf25f395f51905f525f90815290915f516020615cb25f395f51905f525b81831061483f5750509060206105ea928201016147f7565b6020919350806001915483858901015201910190918492614827565b602092506105ea94915060ff191682840152151560051b8201016147f7565b604051905f825f516020615c725f395f51905f52549161489983613106565b808352926001811690811561485b57506001146148bc576105ea925003836105b8565b505f516020615c725f395f51905f525f90815290915f516020615e925f395f51905f525b8183106148f85750509060206105ea928201016147f7565b60209193508060019154838589010152019101909184926148e0565b5f516020615dd25f395f51905f525f525f516020615d525f395f51905f5260205260ff614961827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61334a565b54161561496b5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615dd25f395f51905f52602452604490fd5b5f516020615e125f395f51905f525f525f516020615d525f395f51905f5260205260ff6149e9827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361334a565b5416156149f35750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e125f395f51905f52602452604490fd5b5f516020615d325f395f51905f525f525f516020615d525f395f51905f5260205260ff614a71827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661334a565b541615614a7b5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d325f395f51905f52602452604490fd5b5f516020615d925f395f51905f525f525f516020615d525f395f51905f5260205260ff614af9827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61334a565b541615614b035750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615d925f395f51905f52602452604490fd5b90815f525f516020615d525f395f51905f5260205260ff614b588260405f2061334a565b541615614b63575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105ea91614bc46084836105b8565b6155d1565b5f526004602052600160405f2001600181540180915590565b91614beb615629565b614bff6001600160a01b0384161515612777565b6001600160a01b03821692614c15841515612777565b60ff821615614e8757614c8190614c2a615629565b614c32615629565b614c3a615629565b60ff195f516020615d725f395f51905f5254165f516020615d725f395f51905f5255614c64615629565b614c6c615629565b614c74615629565b614c7c615629565b61429c565b50614c8a615629565b614c9860ff821615156136e0565b60408051614ca682826105b8565b60098152682b30b634b230ba37b960b91b6020820152614cc8825192836105b8565b60058252640312e302e360dc1b6020830152614ce2615629565b614cea615629565b8051906001600160401b03821161056257614d1b82614d165f516020615bf25f395f51905f5254613106565b613ff2565b602090601f8311600114614df35792614d4983614dd9979694614d5d94614daf975f92614170575050614079565b5f516020615bf25f395f51905f5255615adc565b614d725f5f516020615c925f395f51905f5255565b614d875f5f516020615e525f395f51905f5255565b60ff1660ff195f516020615c125f395f51905f525416175f516020615c125f395f51905f5255565b614db7615654565b603380546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615d125f395f51905f525f80a26105ea43603455565b5f516020615bf25f395f51905f525f52601f19831691905f516020615cb25f395f51905f52925f5b818110614e6f575093614d5d93614daf969360019383614dd99b9a9810614e57575b505050811b015f516020615bf25f395f51905f5255615adc565b01515f1960f88460031b161c191690555f8080614e3d565b92936020600181928786015181550195019301614e1b565b6338854f1160e21b5f5260045ffd5b91908203918211610edb57565b60075f9182815582600182015582600282015582600382015582600482015560058101614ed08154613106565b9081614ee3575b50508260068201550155565b601f8211600114614efa57849055505b5f80614ed7565b614f20614f30926001601f614f12855f5260205f2090565b920160051c82019101613fdc565b5f81815260208120918190559055565b614ef3565b9190614f3f6136b0565b50825f526007602052614f558160405f206156b5565b15614fbf57614fba6105ea91845f52600860205260405f20815f52602052610bef614f8260405f2061313e565b95614f9f614f8f82612d0e565b611060612b1860408b01516103fb565b614fb3600260808a01519201918254614e96565b9055612d00565b614ea3565b637f11bea960e01b5f5260045260245ffd5b600360606105ea93805184556020810151600185015560408101516002850155015115159101613502565b156150045750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c060049161503d60018060a01b038251168561278d565b602081015161508890600186019061505e906001600160a01b03168261278d565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061371e565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b91610cf493916150c5936157ac565b9092919261582e565b6001600160a01b03169291906001841461515d5781156151545761511d92156151285760405163a9059cbb60e01b602082015291615115918391612bbe9160248401613ad4565b6005926151c7565b15610cf45750600190565b6040516340c10f1960e01b60208201529161514c918391612bbe9160248401613ad4565b6006926151c7565b50505050600190565b90615170935061193d9250613ea4611db5565b61517957600190565b600590565b90615193915f52600660205260405f2061334a565b600181015460a01c60ff16156151b5576003018054918203918211610edb5755565b6004018054918201809211610edb5755565b81516001600160a01b03909116915f9182916020018285620186a0f16151eb612950565b901561521057805161520757503b1561520357600190565b5f90565b60209150015190565b50505f90565b8147106152705782516001600160a01b03909116925f9182916020018486620186a0f190615242612950565b91156152695715615255575b5050600190565b805161520757503b15615203575f8061524e565b5050505f90565b632b60b36f60e21b5f5260045ffd5b805f525f516020615d525f395f51905f5260205260ff6152a28360405f2061334a565b541661521057805f525f516020615d525f395f51905f526020526152c98260405f2061334a565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610cf4916001600160a01b031690615913565b805f525f516020615d525f395f51905f5260205260ff61533f8360405f2061334a565b54161561521057805f525f516020615d525f395f51905f526020526153678260405f2061334a565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b156153ae57505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156153df57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361545c57506105ea945061542461541682866130c2565b34143490612f3d84886130c2565b80615430575b5061598f565b615451615456916154426033546103fb565b9061544b611db5565b91615216565b6153d8565b5f61542a565b90615468343415612d44565b61547d61547582876130c2565b308489614b80565b80615511575b5061549961193d6001613c7c866125f787612d0e565b6154a9575b506105ea935061598f565b604051632770a7eb60e21b8152602081806154c8883060048401613ad4565b03815f885af1918215613066576105ea966154ec9387935f916154f2575b506153a4565b5f61549e565b61550b915060203d602011612acd57612ac081836105b8565b5f6154e6565b61552990615523612b186033546103fb565b8761595a565b5f615483565b90813b156155b0575f516020615cd25f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a280511561559857615595916159d8565b50565b5050346155a157565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b905f602091828151910182855af115612945575f513d61562057506001600160a01b0381163b155b6156005750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156155f9565b60ff5f516020615e325f395f51905f525460401c161561564557565b631afcd79f60e31b5f5260045ffd5b61565c615629565b62015180600155565b80548210156127f8575f5260205f2001905f90565b805480156156a1575f1901906156908282615665565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615750575f198401848111610edb5783545f19810194908511610edb575f95858361570397610bef9503615709575b50505061567a565b55600190565b6157396157339161572a6157206157479588615665565b90549060031b1c90565b92839187615665565b9061384d565b85905f5260205260405f2090565b555f80806156fb565b505050505f90565b6157606159f5565b615768615a4c565b6040519060208201925f516020615e725f395f51905f528452604083015260608201524660808201523060a082015260a081526157a660c0826105b8565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615819579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15613066575f516001600160a01b0381161561580f57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611ec657565b61583781615824565b80615840575050565b61584981615824565b600181036158605763f645eedf60e01b5f5260045ffd5b61586981615824565b60028103615884575063fce698f760e01b5f5260045260245ffd5b80615890600392615824565b146158985750565b6335e2f38360e21b5f5260045260245ffd5b5f8181526003602052604090205461590e57600254600160401b811015610562576158f76158e18260018594016002556002615665565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61591d82826150a3565b61521057805490600160401b82101561056257826159456158e1846001809601855584615665565b90558054925f520160205260405f2055600190565b614bc46105ea939261598160405194859263a9059cbb60e01b602085015260248401613ad4565b03601f1981018452836105b8565b906159a4915f52600660205260405f2061334a565b600181015460a01c60ff16156159c6576003018054918201809211610edb5755565b6004018054918203918211610edb5755565b5f80610cf493602081519101845af46159ef612950565b91615a7e565b6159fd6147c0565b8051908115615a0d576020012090565b50505f516020615c925f395f51905f52548015615a275790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615a5461487a565b8051908115615a64576020012090565b50505f516020615e525f395f51905f52548015615a275790565b90615aa25750805115615a9357805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615ad3575b615ab3575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615aab565b80519091906001600160401b03811161056257615b1d81615b0a5f516020615c725f395f51905f5254613106565b5f516020615c725f395f51905f52614041565b602092601f8211600114615b5057615b3f929382915f92614170575050614079565b5f516020615c725f395f51905f5255565b5f516020615c725f395f51905f525f52601f198216935f516020615e925f395f51905f52915f5b868110615bb95750836001959610615ba1575b505050811b015f516020615c725f395f51905f5255565b01515f1960f88460031b161c191690555f8080615b8a565b91926020600181928685015181550194019201615b7756feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a26469706673582212208e3a77885b6216d79da8edc5a587eb28cca0b9393c40bef7960752dddc11fea164736f6c634300081c0033"

// Deprecated: Use BSCBridgeV2MetaData.Sigs instead.
// BSCBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var BSCBridgeV2FuncSigs = BSCBridgeV2MetaData.Sigs

// BSCBridgeV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BSCBridgeV2MetaData.Bin instead.
var BSCBridgeV2Bin = BSCBridgeV2MetaData.Bin

// DeployBSCBridgeV2 deploys a new Ethereum contract, binding an instance of BSCBridgeV2 to it.
func DeployBSCBridgeV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCBridgeV2, error) {
	parsed, err := BSCBridgeV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BSCBridgeV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCBridgeV2{BSCBridgeV2Caller: BSCBridgeV2Caller{contract: contract}, BSCBridgeV2Transactor: BSCBridgeV2Transactor{contract: contract}, BSCBridgeV2Filterer: BSCBridgeV2Filterer{contract: contract}}, nil
}

// BSCBridgeV2 is an auto generated Go binding around an Ethereum contract.
type BSCBridgeV2 struct {
	BSCBridgeV2Caller     // Read-only binding to the contract
	BSCBridgeV2Transactor // Write-only binding to the contract
	BSCBridgeV2Filterer   // Log filterer for contract events
}

// BSCBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BSCBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCBridgeV2Session struct {
	Contract     *BSCBridgeV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCBridgeV2CallerSession struct {
	Contract *BSCBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BSCBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCBridgeV2TransactorSession struct {
	Contract     *BSCBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BSCBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BSCBridgeV2Raw struct {
	Contract *BSCBridgeV2 // Generic contract binding to access the raw methods on
}

// BSCBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCBridgeV2CallerRaw struct {
	Contract *BSCBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// BSCBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCBridgeV2TransactorRaw struct {
	Contract *BSCBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCBridgeV2 creates a new instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2(address common.Address, backend bind.ContractBackend) (*BSCBridgeV2, error) {
	contract, err := bindBSCBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2{BSCBridgeV2Caller: BSCBridgeV2Caller{contract: contract}, BSCBridgeV2Transactor: BSCBridgeV2Transactor{contract: contract}, BSCBridgeV2Filterer: BSCBridgeV2Filterer{contract: contract}}, nil
}

// NewBSCBridgeV2Caller creates a new read-only instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*BSCBridgeV2Caller, error) {
	contract, err := bindBSCBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Caller{contract: contract}, nil
}

// NewBSCBridgeV2Transactor creates a new write-only instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BSCBridgeV2Transactor, error) {
	contract, err := bindBSCBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Transactor{contract: contract}, nil
}

// NewBSCBridgeV2Filterer creates a new log filterer instance of BSCBridgeV2, bound to a specific deployed contract.
func NewBSCBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BSCBridgeV2Filterer, error) {
	contract, err := bindBSCBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2Filterer{contract: contract}, nil
}

// bindBSCBridgeV2 binds a generic wrapper to an already deployed contract.
func bindBSCBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BSCBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeV2.Contract.BSCBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BSCBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeV2 *BSCBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BSCBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeV2 *BSCBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeV2 *BSCBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeV2 *BSCBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DEFAULTADMINROLE(&_BSCBridgeV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DEFAULTADMINROLE(&_BSCBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_BSCBridgeV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridgeV2.Contract.UPGRADEINTERFACEVERSION(&_BSCBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllChainIDs(&_BSCBridgeV2.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllChainIDs(&_BSCBridgeV2.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllPendingIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridgeV2.Contract.AllPendingIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2Session) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.AllTokenPairs(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.AllTokenPairs(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeExecutor() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecutor(&_BSCBridgeV2.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) BridgeExecutor() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecutor(&_BSCBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeVerifier() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeVerifier(&_BSCBridgeV2.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) BridgeVerifier() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeVerifier(&_BSCBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridgeV2.Contract.CrossMintableERC20Code(&_BSCBridgeV2.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridgeV2.Contract.CrossMintableERC20Code(&_BSCBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) Dev() (common.Address, error) {
	return _BSCBridgeV2.Contract.Dev(&_BSCBridgeV2.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Dev() (common.Address, error) {
	return _BSCBridgeV2.Contract.Dev(&_BSCBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) DomainSeparator() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DomainSeparator(&_BSCBridgeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) DomainSeparator() ([32]byte, error) {
	return _BSCBridgeV2.Contract.DomainSeparator(&_BSCBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridgeV2 *BSCBridgeV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridgeV2.Contract.Eip712Domain(&_BSCBridgeV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridgeV2.Contract.Eip712Domain(&_BSCBridgeV2.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextFinalizeIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextFinalizeIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextInitiateIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridgeV2.Contract.GetNextInitiateIndex(&_BSCBridgeV2.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Session) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridgeV2.Contract.GetPendingArguments(&_BSCBridgeV2.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridgeV2.Contract.GetPendingArguments(&_BSCBridgeV2.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridgeV2.Contract.GetRoleAdmin(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridgeV2.Contract.GetRoleAdmin(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2Session) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridgeV2.Contract.GetRoleMembers(&_BSCBridgeV2.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridgeV2.Contract.GetRoleMembers(&_BSCBridgeV2.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Caller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2Session) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.GetTokenPair(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridgeV2.Contract.GetTokenPair(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.HasRole(&_BSCBridgeV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.HasRole(&_BSCBridgeV2.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Caller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2Session) InitializedAt() (*big.Int, error) {
	return _BSCBridgeV2.Contract.InitializedAt(&_BSCBridgeV2.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) InitializedAt() (*big.Int, error) {
	return _BSCBridgeV2.Contract.InitializedAt(&_BSCBridgeV2.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.IsTokenFinalizePaused(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridgeV2.Contract.IsTokenFinalizePaused(&_BSCBridgeV2.CallOpts, remoteChainID, token)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) Paused() (bool, error) {
	return _BSCBridgeV2.Contract.Paused(&_BSCBridgeV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Paused() (bool, error) {
	return _BSCBridgeV2.Contract.Paused(&_BSCBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2Session) ProxiableUUID() ([32]byte, error) {
	return _BSCBridgeV2.Contract.ProxiableUUID(&_BSCBridgeV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _BSCBridgeV2.Contract.ProxiableUUID(&_BSCBridgeV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridgeV2.Contract.SupportsInterface(&_BSCBridgeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridgeV2.Contract.SupportsInterface(&_BSCBridgeV2.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2Caller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2Session) Threshold() (uint8, error) {
	return _BSCBridgeV2.Contract.Threshold(&_BSCBridgeV2.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) Threshold() (uint8, error) {
	return _BSCBridgeV2.Contract.Threshold(&_BSCBridgeV2.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) BurnCrossToDeadWallet(opts *bind.TransactOpts, deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "burnCrossToDeadWallet", deadWallet, amount, alreadyTransferred)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) BurnCrossToDeadWallet(deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BurnCrossToDeadWallet(&_BSCBridgeV2.TransactOpts, deadWallet, amount, alreadyTransferred)
}

// BurnCrossToDeadWallet is a paid mutator transaction binding the contract method 0x86e9e175.
//
// Solidity: function burnCrossToDeadWallet(address deadWallet, uint256 amount, bool alreadyTransferred) returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) BurnCrossToDeadWallet(deadWallet common.Address, amount *big.Int, alreadyTransferred bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.BurnCrossToDeadWallet(&_BSCBridgeV2.TransactOpts, deadWallet, amount, alreadyTransferred)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ChangeThreshold(&_BSCBridgeV2.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ChangeThreshold(&_BSCBridgeV2.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2Session) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.CreateToken(&_BSCBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.CreateToken(&_BSCBridgeV2.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridge(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridge(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridgeBatch(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.FinalizeBridgeBatch(&_BSCBridgeV2.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.GrantRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) Initialize(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "initialize", owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Initialize(&_BSCBridgeV2.TransactOpts, owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Initialize(&_BSCBridgeV2.TransactOpts, owner, dev_, threshold_)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePendingWithRecipient(&_BSCBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ManualReleasePendingWithRecipient(&_BSCBridgeV2.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Transactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2Session) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeToken(&_BSCBridgeV2.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeTokenBatch(&_BSCBridgeV2.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.PermitBridgeTokenBatch(&_BSCBridgeV2.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RegisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RegisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ReinitializeBSCBridge(opts *bind.TransactOpts, crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "reinitializeBSCBridge", crossChainID_, cross_)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ReinitializeBSCBridge(crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReinitializeBSCBridge(&_BSCBridgeV2.TransactOpts, crossChainID_, cross_)
}

// ReinitializeBSCBridge is a paid mutator transaction binding the contract method 0x7b429ad8.
//
// Solidity: function reinitializeBSCBridge(uint256 crossChainID_, address cross_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ReinitializeBSCBridge(crossChainID_ *big.Int, cross_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReinitializeBSCBridge(&_BSCBridgeV2.TransactOpts, crossChainID_, cross_)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "releasePendingBatch", remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePendingBatch(&_BSCBridgeV2.TransactOpts, remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.ReleasePendingBatch(&_BSCBridgeV2.TransactOpts, remoteChainIDs, indexes)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RemovePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RemovePending(&_BSCBridgeV2.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RenounceRole(&_BSCBridgeV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RenounceRole(&_BSCBridgeV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRole(&_BSCBridgeV2.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.RevokeRoleBatch(&_BSCBridgeV2.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecutor(&_BSCBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecutor(&_BSCBridgeV2.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeVerifier(&_BSCBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeVerifier(&_BSCBridgeV2.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetChainPause(&_BSCBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetChainPause(&_BSCBridgeV2.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetCrossMintableERC20Code(&_BSCBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetCrossMintableERC20Code(&_BSCBridgeV2.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetDev(&_BSCBridgeV2.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetDev(&_BSCBridgeV2.TransactOpts, dev_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPause(&_BSCBridgeV2.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetPause(&_BSCBridgeV2.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelay(&_BSCBridgeV2.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelay(&_BSCBridgeV2.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelayExpiration(&_BSCBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetVerificationDelayExpiration(&_BSCBridgeV2.TransactOpts, remoteChainID, index, delay)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UnregisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UnregisterToken(&_BSCBridgeV2.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UpgradeToAndCall(&_BSCBridgeV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.UpgradeToAndCall(&_BSCBridgeV2.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) Receive() (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Receive(&_BSCBridgeV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) Receive() (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.Receive(&_BSCBridgeV2.TransactOpts)
}

// BSCBridgeV2BridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecutorSetIterator struct {
	Event *BSCBridgeV2BridgeExecutorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2BridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeExecutorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2BridgeExecutorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2BridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeExecutorSet represents a BridgeExecutorSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*BSCBridgeV2BridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeExecutorSetIterator{contract: _BSCBridgeV2.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeExecutorSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeExecutorSet is a log parse operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeExecutorSet(log types.Log) (*BSCBridgeV2BridgeExecutorSet, error) {
	event := new(BSCBridgeV2BridgeExecutorSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeFinalizedIterator struct {
	Event *BSCBridgeV2BridgeFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2BridgeFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeFinalized represents a BridgeFinalized event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeFinalized struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeFinalized is a free log retrieval operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeV2BridgeFinalizedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeFinalizedIterator{contract: _BSCBridgeV2.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeFinalized)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeFinalized is a log parse operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeFinalized(log types.Log) (*BSCBridgeV2BridgeFinalized, error) {
	event := new(BSCBridgeV2BridgeFinalized)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeInitiatedIterator struct {
	Event *BSCBridgeV2BridgeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2BridgeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeInitiated represents a BridgeInitiated event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeInitiated struct {
	ToChainID  *big.Int
	Index      *big.Int
	FromToken  common.Address
	ToToken    common.Address
	From       common.Address
	To         common.Address
	Value      *big.Int
	NetworkFee *big.Int
	ExFee      *big.Int
	ExtraData  []byte
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*BSCBridgeV2BridgeInitiatedIterator, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeInitiatedIterator{contract: _BSCBridgeV2.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeInitiated)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeInitiated is a log parse operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeInitiated(log types.Log) (*BSCBridgeV2BridgeInitiated, error) {
	event := new(BSCBridgeV2BridgeInitiated)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgePendingIterator struct {
	Event *BSCBridgeV2BridgePending // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2BridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgePending)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2BridgePending)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2BridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgePending represents a BridgePending event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgePending struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Time        *big.Int
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgePending is a free log retrieval operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeV2BridgePendingIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgePendingIterator{contract: _BSCBridgeV2.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgePending)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgePending is a log parse operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgePending(log types.Log) (*BSCBridgeV2BridgePending, error) {
	event := new(BSCBridgeV2BridgePending)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeVerifierSetIterator struct {
	Event *BSCBridgeV2BridgeVerifierSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2BridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeVerifierSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2BridgeVerifierSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2BridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeVerifierSet represents a BridgeVerifierSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*BSCBridgeV2BridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeVerifierSetIterator{contract: _BSCBridgeV2.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeVerifierSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeVerifierSet is a log parse operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeVerifierSet(log types.Log) (*BSCBridgeV2BridgeVerifierSet, error) {
	event := new(BSCBridgeV2BridgeVerifierSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ChainPauseSetIterator struct {
	Event *BSCBridgeV2ChainPauseSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2ChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ChainPauseSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2ChainPauseSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2ChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ChainPauseSet represents a ChainPauseSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BSCBridgeV2ChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ChainPauseSetIterator{contract: _BSCBridgeV2.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ChainPauseSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChainPauseSet is a log parse operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseChainPauseSet(log types.Log) (*BSCBridgeV2ChainPauseSet, error) {
	event := new(BSCBridgeV2ChainPauseSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2CrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2CrossMintableERC20CodeSetIterator struct {
	Event *BSCBridgeV2CrossMintableERC20CodeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2CrossMintableERC20CodeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2CrossMintableERC20CodeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2CrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2CrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2CrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*BSCBridgeV2CrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2CrossMintableERC20CodeSetIterator{contract: _BSCBridgeV2.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2CrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2CrossMintableERC20CodeSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseCrossMintableERC20CodeSet(log types.Log) (*BSCBridgeV2CrossMintableERC20CodeSet, error) {
	event := new(BSCBridgeV2CrossMintableERC20CodeSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2DevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2DevSetIterator struct {
	Event *BSCBridgeV2DevSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2DevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2DevSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2DevSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2DevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2DevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2DevSet represents a DevSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2DevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*BSCBridgeV2DevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2DevSetIterator{contract: _BSCBridgeV2.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2DevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2DevSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDevSet is a log parse operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseDevSet(log types.Log) (*BSCBridgeV2DevSet, error) {
	event := new(BSCBridgeV2DevSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2EIP712DomainChangedIterator struct {
	Event *BSCBridgeV2EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BSCBridgeV2EIP712DomainChangedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2EIP712DomainChangedIterator{contract: _BSCBridgeV2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2EIP712DomainChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseEIP712DomainChanged(log types.Log) (*BSCBridgeV2EIP712DomainChanged, error) {
	event := new(BSCBridgeV2EIP712DomainChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ExtraCallExecutedIterator struct {
	Event *BSCBridgeV2ExtraCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2ExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ExtraCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2ExtraCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2ExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ExtraCallExecuted represents a ExtraCallExecuted event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ExtraCallExecuted struct {
	FromChainID *big.Int
	Index       *big.Int
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BSCBridgeV2ExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ExtraCallExecutedIterator{contract: _BSCBridgeV2.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ExtraCallExecuted, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ExtraCallExecuted)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtraCallExecuted is a log parse operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseExtraCallExecuted(log types.Log) (*BSCBridgeV2ExtraCallExecuted, error) {
	event := new(BSCBridgeV2ExtraCallExecuted)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BSCBridgeV2 contract.
type BSCBridgeV2InitializedIterator struct {
	Event *BSCBridgeV2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Initialized represents a Initialized event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*BSCBridgeV2InitializedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2InitializedIterator{contract: _BSCBridgeV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Initialized) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Initialized)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseInitialized(log types.Log) (*BSCBridgeV2Initialized, error) {
	event := new(BSCBridgeV2Initialized)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ManualReleasedIterator struct {
	Event *BSCBridgeV2ManualReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2ManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ManualReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2ManualReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2ManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ManualReleased represents a ManualReleased event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgeV2ManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ManualReleasedIterator{contract: _BSCBridgeV2.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ManualReleased)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManualReleased is a log parse operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseManualReleased(log types.Log) (*BSCBridgeV2ManualReleased, error) {
	event := new(BSCBridgeV2ManualReleased)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BSCBridgeV2 contract.
type BSCBridgeV2PausedIterator struct {
	Event *BSCBridgeV2Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Paused represents a Paused event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterPaused(opts *bind.FilterOpts) (*BSCBridgeV2PausedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2PausedIterator{contract: _BSCBridgeV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Paused) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Paused)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParsePaused(log types.Log) (*BSCBridgeV2Paused, error) {
	event := new(BSCBridgeV2Paused)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2PendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the BSCBridgeV2 contract.
type BSCBridgeV2PendingRemovedIterator struct {
	Event *BSCBridgeV2PendingRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2PendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2PendingRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2PendingRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2PendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2PendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2PendingRemoved represents a PendingRemoved event raised by the BSCBridgeV2 contract.
type BSCBridgeV2PendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgeV2PendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2PendingRemovedIterator{contract: _BSCBridgeV2.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2PendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2PendingRemoved)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingRemoved is a log parse operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParsePendingRemoved(log types.Log) (*BSCBridgeV2PendingRemoved, error) {
	event := new(BSCBridgeV2PendingRemoved)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleAdminChangedIterator struct {
	Event *BSCBridgeV2RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleAdminChanged represents a RoleAdminChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BSCBridgeV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleAdminChangedIterator{contract: _BSCBridgeV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleAdminChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleAdminChanged(log types.Log) (*BSCBridgeV2RoleAdminChanged, error) {
	event := new(BSCBridgeV2RoleAdminChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleGrantedIterator struct {
	Event *BSCBridgeV2RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleGranted represents a RoleGranted event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleGrantedIterator{contract: _BSCBridgeV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleGranted)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleGranted(log types.Log) (*BSCBridgeV2RoleGranted, error) {
	event := new(BSCBridgeV2RoleGranted)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleRevokedIterator struct {
	Event *BSCBridgeV2RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2RoleRevoked represents a RoleRevoked event raised by the BSCBridgeV2 contract.
type BSCBridgeV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2RoleRevokedIterator{contract: _BSCBridgeV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2RoleRevoked)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseRoleRevoked(log types.Log) (*BSCBridgeV2RoleRevoked, error) {
	event := new(BSCBridgeV2RoleRevoked)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2ThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BSCBridgeV2 contract.
type BSCBridgeV2ThresholdChangedIterator struct {
	Event *BSCBridgeV2ThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2ThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2ThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2ThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2ThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2ThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2ThresholdChanged represents a ThresholdChanged event raised by the BSCBridgeV2 contract.
type BSCBridgeV2ThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BSCBridgeV2ThresholdChangedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2ThresholdChangedIterator{contract: _BSCBridgeV2.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2ThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2ThresholdChanged)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdChanged is a log parse operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseThresholdChanged(log types.Log) (*BSCBridgeV2ThresholdChanged, error) {
	event := new(BSCBridgeV2ThresholdChanged)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2TokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairRegisteredIterator struct {
	Event *BSCBridgeV2TokenPairRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2TokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2TokenPairRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2TokenPairRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2TokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2TokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2TokenPairRegistered represents a TokenPairRegistered event raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BSCBridgeV2TokenPairRegisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2TokenPairRegisteredIterator{contract: _BSCBridgeV2.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2TokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2TokenPairRegistered)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairRegistered is a log parse operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseTokenPairRegistered(log types.Log) (*BSCBridgeV2TokenPairRegistered, error) {
	event := new(BSCBridgeV2TokenPairRegistered)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2TokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairUnregisteredIterator struct {
	Event *BSCBridgeV2TokenPairUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2TokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2TokenPairUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2TokenPairUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2TokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2TokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2TokenPairUnregistered represents a TokenPairUnregistered event raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*BSCBridgeV2TokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2TokenPairUnregisteredIterator{contract: _BSCBridgeV2.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2TokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2TokenPairUnregistered)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairUnregistered is a log parse operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseTokenPairUnregistered(log types.Log) (*BSCBridgeV2TokenPairUnregistered, error) {
	event := new(BSCBridgeV2TokenPairUnregistered)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2TokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPauseSetIterator struct {
	Event *BSCBridgeV2TokenPauseSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2TokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2TokenPauseSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2TokenPauseSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2TokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2TokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2TokenPauseSet represents a TokenPauseSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2TokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BSCBridgeV2TokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2TokenPauseSetIterator{contract: _BSCBridgeV2.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2TokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2TokenPauseSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPauseSet is a log parse operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseTokenPauseSet(log types.Log) (*BSCBridgeV2TokenPauseSet, error) {
	event := new(BSCBridgeV2TokenPauseSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BSCBridgeV2 contract.
type BSCBridgeV2UnpausedIterator struct {
	Event *BSCBridgeV2Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Unpaused represents a Unpaused event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*BSCBridgeV2UnpausedIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2UnpausedIterator{contract: _BSCBridgeV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Unpaused)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseUnpaused(log types.Log) (*BSCBridgeV2Unpaused, error) {
	event := new(BSCBridgeV2Unpaused)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BSCBridgeV2 contract.
type BSCBridgeV2UpgradedIterator struct {
	Event *BSCBridgeV2Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2Upgraded represents a Upgraded event raised by the BSCBridgeV2 contract.
type BSCBridgeV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BSCBridgeV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2UpgradedIterator{contract: _BSCBridgeV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2Upgraded)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseUpgraded(log types.Log) (*BSCBridgeV2Upgraded, error) {
	event := new(BSCBridgeV2Upgraded)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2VerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelayExpirationSetIterator struct {
	Event *BSCBridgeV2VerificationDelayExpirationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2VerificationDelayExpirationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2VerificationDelayExpirationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2VerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2VerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BSCBridgeV2VerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2VerificationDelayExpirationSetIterator{contract: _BSCBridgeV2.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2VerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2VerificationDelayExpirationSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerificationDelayExpirationSet is a log parse operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseVerificationDelayExpirationSet(log types.Log) (*BSCBridgeV2VerificationDelayExpirationSet, error) {
	event := new(BSCBridgeV2VerificationDelayExpirationSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2VerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelaySetIterator struct {
	Event *BSCBridgeV2VerificationDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCBridgeV2VerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2VerificationDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCBridgeV2VerificationDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCBridgeV2VerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2VerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2VerificationDelaySet represents a VerificationDelaySet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2VerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*BSCBridgeV2VerificationDelaySetIterator, error) {

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2VerificationDelaySetIterator{contract: _BSCBridgeV2.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2VerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2VerificationDelaySet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerificationDelaySet is a log parse operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseVerificationDelaySet(log types.Log) (*BSCBridgeV2VerificationDelaySet, error) {
	event := new(BSCBridgeV2VerificationDelaySet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
