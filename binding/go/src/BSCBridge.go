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

// BSCBridgeMetaData contains all meta data concerning the BSCBridge contract.
var BSCBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenFinalizePauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"a10bab0b": "bridgeExecutor()",
		"5fd262de": "bridgeToken(uint256,address,address,uint256,uint256,uint256,bytes)",
		"f0702e8e": "bridgeVerifier()",
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
		"3aefddbf": "initializeBSCBridge(address,address,uint8,uint256,address,uint256)",
		"91cf6d3e": "initializedAt()",
		"761fe2d8": "isTokenFinalizePaused(uint256,address)",
		"aa20ed47": "manualReleasePending(uint256,uint256)",
		"48a00ed8": "manualReleasePendingWithRecipient(uint256,uint256,address)",
		"5c975abb": "paused()",
		"4d5d0056": "permitBridgeToken(uint256,address,address,uint256,uint256,uint256,bytes,(address,address,uint256,uint256,uint8,bytes32,bytes32))",
		"1313996b": "permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[],(address,address,uint256,uint256,uint8,bytes32,bytes32)[])",
		"52d1902d": "proxiableUUID()",
		"edbbea39": "registerToken(uint256,bool,address,address)",
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
	Bin: "0x60a080604052346100c257306080525f516020615f105f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615e4990816100c78239608051818181610e4c01526110030152f35b6001600160401b0319166001600160401b039081175f516020615f105f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b61002061359a565b005b5f3560e01c806301ffc9a7146103915780630b1d8d061461038c5780631313996b146103875780631938e0f214610382578063248a9ca31461037d5780632f2ff15d1461037857806336568abe14610373578063365d71e41461036e5780633aefddbf1461036957806342cde4e81461036457806348a00ed81461035f5780634be13f831461035a5780634d5d0056146103555780634ee078ba146103505780634f1ef2861461034b578063502cc5c01461034657806352d1902d146103415780635b605f5c1461033c5780635c975abb146103375780635fd262de14610332578063670e62681461032d578063761fe2d8146103285780637921487414610323578063814914b51461031e57806384b0196e1461031957806388d67d6d1461031457806389232a001461030f5780638ae87c5c1461030a57806391cca3db1461030557806391cf6d3e1461030057806391d14854146102fb5780639f9b4f1c146102f6578063a10bab0b146102f1578063a217fddf146102ec578063a3246ad3146102e7578063aa1bd0bc146102e2578063aa20ed47146102dd578063ad3cb1cc146102d8578063ae6893f8146102d3578063b2b49e2e146102ce578063b33eb36e146102c9578063b7f3358d146102c4578063b8aa837e146102bf578063bedb86fb146102ba578063bfbf6765146102b5578063cba25e4b146102b0578063cf56118e146102ab578063d477f05f146102a6578063d547741f146102a1578063d5717fc51461029c578063e2af6cd714610297578063edbbea3914610292578063f0702e8e1461028d578063f4509643146102885763f698da250361000e576124ea565b61244a565b612422565b6123d8565b612326565b6122ed565b6122be565b612260565b6121ec565b612182565b6120aa565b611fc1565b611f26565b611e99565b611e08565b611cfb565b611cc2565b611c7b565b611bf2565b611ba6565b611b2a565b611ace565b611aa6565b611993565b611951565b611934565b61190c565b6118a4565b6117f1565b6116e7565b6115bf565b611472565b6113f2565b61136c565b6112f5565b6111ed565b6111bf565b6110e0565b610ff1565b610f62565b610e0a565b610caa565b610c3d565b610be9565b610acb565b610a9f565b61095d565b6108fa565b61080a565b6107d1565b6107a0565b6106ef565b6104a9565b610408565b346103e75760203660031901126103e75760043563ffffffff60e01b81168091036103e757602090637965db0b60e01b81149081156103d6575b506040519015158152f35b6301ffc9a760e01b1490505f6103cb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103e757565b346103e75760203660031901126103e757600435610425816103f7565b61042d6135e9565b6001600160a01b0316610441811515612556565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103e7578235916001600160401b0383116103e7576020808501948460051b0101116103e757565b60403660031901126103e7576004356001600160401b0381116103e7576104d4903690600401610479565b602435916001600160401b0383116103e757366023840112156103e7576004830135916001600160401b0383116103e75736602460e08502860101116103e7576024610020940191612797565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761055157604052565b610521565b60e081019081106001600160401b0382111761055157604052565b606081019081106001600160401b0382111761055157604052565b60c081019081106001600160401b0382111761055157604052565b601f909101601f19168101906001600160401b0382119082101761055157604052565b604051906105da610100836105a7565b565b604051906105da6060836105a7565b604051906105da6080836105a7565b604051906105da60e0836105a7565b6001600160401b0381116105515760051b60200190565b60ff8116036103e757565b9080601f830112156103e757813561064281610609565b9261065060405194856105a7565b81845260208085019260051b8201019283116103e757602001905b8282106106785750505090565b60208091833561068781610620565b81520191019061066b565b9080601f830112156103e75781356106a981610609565b926106b760405194856105a7565b81845260208085019260051b8201019283116103e757602001905b8282106106df5750505090565b81358152602091820191016106d2565b60803660031901126103e7576004356001600160401b0381116103e75760c060031982360301126103e7576024356001600160401b0381116103e75761073990369060040161062b565b906044356001600160401b0381116103e757610759903690600401610692565b60643591906001600160401b0383116103e75761079c9361078161078a943690600401610692565b926004016128d5565b60405191829182901515815260200190565b0390f35b346103e75760203660031901126103e75760206107be600435612c27565b604051908152f35b35906105da826103f7565b346103e75760403660031901126103e7576100206024356004356107f4826103f7565b61080561080082612c27565b6137d1565b61430d565b346103e75760403660031901126103e75760043560243561082a816103f7565b336001600160a01b03821603610843576100209161436d565b63334bd91960e11b5f5260045ffd5b9060406003198301126103e7576004356001600160401b0381116103e7578261087d91600401610692565b91602435906001600160401b0382116103e757806023830112156103e75781600401356108a981610609565b926108b760405194856105a7565b8184526024602085019260051b8201019283116103e757602401905b8282106108e05750505090565b6020809183356108ef816103f7565b8152019101906108d3565b346103e75761090836610852565b906109168151835114612c45565b5f5b8251811015610020578061095661093160019385612c5b565b51838060a01b036109428488612c5b565b51169061095161080082612c27565b61436d565b5001610918565b346103e75760c03660031901126103e75760043561097a816103f7565b60243590610987826103f7565b60443561099381610620565b6084356064356109a2826103f7565b60a435925f516020615d945f395f51905f5254956109cb6109c78860ff9060401c1690565b1590565b966001600160401b031680159081610a8d575b6001149081610a83575b159081610a7a575b50610a6b57610a0b9587610a02612c6f565b610a5e57612cb1565b610a1157005b610a2f5f516020615d945f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a66612c90565b612cb1565b63f92ee8a960e01b5f5260045ffd5b9050155f6109f0565b303b1591506109e8565b8891506109de565b5f9103126103e757565b346103e7575f3660031901126103e757602060ff5f516020615b945f395f51905f525416604051908152f35b346103e75760603660031901126103e757602435600435604435610aee816103f7565b610af6613663565b610afe61383b565b815f526007602052610b1c83610b178160405f20614e6d565b612dcf565b80610b278484614dd1565b916001600160a01b031615610bd5575b8151915f516020615bb45f395f51905f526040820192610b84610b72855160018060a01b03169660808601978489519160a089015193613c0a565b610b7b81611d6d565b600181146147db565b8251602090930151935194516040516001600160a01b0390961695918291610baf9142919084612c09565b0390a45f516020615b545f395f51905f525f80a35f5f516020615d145f395f51905f525d005b5060608101516001600160a01b0316610b37565b346103e7575f3660031901126103e7575f546040516001600160a01b039091168152602090f35b9181601f840112156103e7578235916001600160401b0383116103e757602083818601950101116103e757565b6101c03660031901126103e757602435600435610c59826103f7565b604435610c65816103f7565b6064359060a43560843560c4356001600160401b0381116103e757610c8e903690600401610c10565b94909360e03660e31901126103e75761079c9761078a97612de9565b346103e75760403660031901126103e757610d89602435600435610ccc613814565b610cd461383b565b805f526007602052610ced82610b178160405f20614e6d565b610d846040610d14610d0f85610d0286612afb565b905f5260205260405f2090565b61319d565b610d71610d3482516080610d2a868301516103eb565b9101519087613a82565b50610d3e81611d6d565b610d4781611d6d565b83516020810182905290600190610d6b83604081015b03601f1981018552846105a7565b146131d6565b01518015908115610d91575b4291613202565b6147fe565b610020613870565b4281109150610d7d565b6001600160401b03811161055157601f01601f191660200190565b929192610dc282610d9b565b91610dd060405193846105a7565b8294818452818301116103e7578281602093845f960137010152565b9080601f830112156103e757816020610e0793359101610db6565b90565b60403660031901126103e757600435610e22816103f7565b6024356001600160401b0381116103e757610e41903690600401610dec565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f40575b50610f3157610e846135e9565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610f00575b50610ecd57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615c345f395f51905f528303610eec57610020925061545a565b632a87526960e21b5f52600483905260245ffd5b610f2391945060203d602011610f2a575b610f1b81836105a7565b810190614bdc565b925f610eac565b503d610f11565b63703e46dd60e11b5f5260045ffd5b5f516020615c345f395f51905f52546001600160a01b0316141590505f610e77565b346103e75760603660031901126103e757602435600435604435610f84613663565b815f526007602052610f9d83610b178160405f20614e6d565b4201804211610fec5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b613009565b346103e7575f3660031901126103e7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f315760206040515f516020615c345f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110c05750505090565b909192602060e0826110d56001948851611048565b0194019291016110b3565b346103e75760203660031901126103e757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106111a657505061112f925003836105a7565b6111398251613256565b915f5b81518110156111985760019061117c61117761115786612b09565b6111716111648588612c5b565b516001600160a01b031690565b906132a5565b6132ba565b6111868287612c5b565b526111918186612c5b565b500161113c565b6040518061079c868261109d565b845483526001948501948794506020909301920161111a565b346103e7575f3660031901126103e757602060ff5f516020615cd45f395f51905f5254166040519015158152f35b60e03660031901126103e757602435600435611208826103f7565b60443591611215836103f7565b6064359260c435916084359160a435916001600160401b0385116103e7576112c69661127b61124b6112bc973690600401610c10565b959096611256613814565b6001600160a01b03851694849061126d878d6148a3565b61127561383b565b8b614943565b9390926040519861128b8a610535565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610db6565b60e0820152614ade565b5f5f516020615d145f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103e75760803660031901126103e757602435600435611315826103f7565b604435906001600160401b0382116103e757366023830112156103e75761079c9261134d611360933690602481600401359101610db6565b906064359261135b84610620565b613342565b604051918291826112e2565b346103e75760403660031901126103e757602060ff6113a2602435600435611393826103f7565b5f526009845260405f206132a5565b54166040519015158152f35b90602080835192838152019201905f5b8181106113cb5750505090565b82518452602093840193909201916001016113be565b906020610e079281815201906113ae565b346103e75760203660031901126103e7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061144c5761079c85611440818703826105a7565b604051918291826113e1565b8254845260209093019260019283019201611429565b60e0810192916105da9190611048565b346103e75760403660031901126103e75761079c6114b1602435600435611498826103f7565b6114a0613220565b505f52600660205260405f206132a5565b6004604051916114c083610556565b80546001600160a01b0316835260018101546115159061150c906114ef6114e6826103eb565b60208801613052565b61150360a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611462565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161159590611587610e0797959693600f60f81b865260e0602087015260e086019061153c565b90848203604086015261153c565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526113ae565b346103e7575f3660031901126103e7575f516020615bf45f395f51905f52541580611653575b15611616576115f2614beb565b6115fa614ca5565b9061079c61160661340b565b6040519384933091469186611560565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615db45f395f51905f5254156115e5565b9080601f830112156103e757813561168081610609565b9261168e60405194856105a7565b81845260208085019260051b820101918383116103e75760208201905b8382106116ba57505050505090565b81356001600160401b0381116103e7576020916116dc87848094880101610692565b8152019101906116ab565b60803660031901126103e7576004356001600160401b0381116103e757611712903690600401610479565b90602435906001600160401b0382116103e757366023830112156103e757816004013561173e81610609565b9261174c60405194856105a7565b8184526024602085019260051b820101903682116103e75760248101925b8284106117c25750506044359150506001600160401b0381116103e757611795903690600401611669565b606435926001600160401b0384116103e75761079c946117bc61078a953690600401611669565b93613426565b83356001600160401b0381116103e7576020916117e683926024369187010161062b565b81520193019261176a565b346103e75760603660031901126103e75760043561180e816103f7565b6024359061181b826103f7565b60443561182781610620565b5f516020615d945f395f51905f52549260ff604085901c1615936001600160401b03168015908161189c575b6001149081611892575b159081611889575b50610a6b57610a0b9284611877612c6f565b156143cc57611884612c90565b6143cc565b9050155f611865565b303b15915061185d565b859150611853565b346103e75760403660031901126103e7576024356004356118c36135e9565b6118cb61383b565b6118d58282614dd1565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615d145f395f51905f525d005b346103e7575f3660031901126103e7576033546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e7576020603454604051908152f35b346103e75760403660031901126103e757602060ff6113a2602435600435611978826103f7565b5f525f516020615cb45f395f51905f52845260405f206132a5565b346103e75760403660031901126103e7576004356001600160401b0381116103e7576119c3903690600401610692565b6024356001600160401b0381116103e7576119e2903690600401610692565b906119f081518351146125ad565b5f5b82518110156100205780611a98611a0b60019385612c5b565b51611a168387612c5b565b5190611a20613814565b611a2861383b565b805f526007602052611a4182610b178160405f20614e6d565b610d846040611a56610d0f85610d0286612afb565b610d71611a6c82516080610d2a868301516103eb565b50611a7681611d6d565b611a7f81611d6d565b835160208101829052908a90610d6b8360408101610d5d565b611aa0613870565b016119f2565b346103e7575f3660031901126103e7576035546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611b0b5750505090565b82516001600160a01b0316845260209384019390920191600101611afe565b346103e75760203660031901126103e7576004355f525f516020615c545f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611b905761079c85611b84818703826105a7565b60405191829182611ae8565b8254845260209093019260019283019201611b6d565b346103e75760203660031901126103e7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611be56135e9565b80600155604051908152a1005b346103e75760403660031901126103e757602435600435611c11613663565b611c19613663565b611c2161383b565b805f526007602052611c3a82610b178160405f20614e6d565b611c4482826147fe565b5f516020615b545f395f51905f525f80a35f5f516020615d145f395f51905f525d005b60405190611c766020836105a7565b5f8252565b346103e7575f3660031901126103e75761079c604051611c9c6040826105a7565b60058152640352e302e360dc1b602082015260405191829160208352602083019061153c565b346103e75760203660031901126103e7576004355f526004602052600160405f20015460018101809111610fec57602090604051908152f35b346103e757611d0936610852565b90611d178151835114612c45565b5f5b82518110156100205780611d52611d3260019385612c5b565b51838060a01b03611d438488612c5b565b51169061080561080082612c27565b5001611d19565b634e487b7160e01b5f52602160045260245ffd5b600a1115611d7757565b611d59565b90600a821015611d775752565b6020815260606040611dee60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061153c565b93611e00602082015183860190611d7c565b015191015290565b346103e75760403660031901126103e757600435602435905f60408051611e2e81610571565b611e366134bc565b815282602082015201525f52600860205260405f20905f5260205261079c60405f20600760405191611e6783610571565b611e7081613099565b8352611e8660ff60068301541660208501613191565b0154604082015260405191829182611d89565b346103e75760203660031901126103e7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611edb81610620565b611ee36135e9565b16611eef8115156134ec565b8060ff195f516020615b945f395f51905f525416175f516020615b945f395f51905f5255604051908152a1005b801515036103e757565b346103e75760403660031901126103e757600435602435611f4681611f1c565b611f4e6136dd565b611f63825f52600360205260405f2054151590565b15611fae5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611fa381600360405f2001613502565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103e75760203660031901126103e757600435611fde81611f1c565b611fe66136dd565b1561204457611ff3613814565b600160ff195f516020615cd45f395f51905f525416175f516020615cd45f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615cd45f395f51905f525460ff81161561209b5760ff19165f516020615cd45f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103e75760803660031901126103e7576024356004356120ca826103f7565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356120f981611f1c565b60643561210581611f1c565b61210d6136dd565b845f5260056020526121718161216c855f209861213c8161213760018060a01b038216809d614e6d565b613513565b885f52600660205261215c866001612156848b5f206132a5565b0161353b565b885f526009602052865f206132a5565b613502565b8251911515825215156020820152a3005b346103e75760203660031901126103e75760043561219f816103f7565b6121a76135e9565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103e7575f3660031901126103e757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b81811061224a5761079c85611440818703826105a7565b8254845260209093019260019283019201612233565b346103e75760203660031901126103e75760043561227d816103f7565b6122856135e9565b6001600160a01b0316612299811515612556565b603380546001600160a01b031916821790555f516020615c745f395f51905f525f80a2005b346103e75760403660031901126103e7576100206024356004356122e1826103f7565b61095161080082612c27565b346103e75760203660031901126103e7576004355f526004602052600260405f20015460018101809111610fec57602090604051908152f35b346103e75760203660031901126103e757600435612343816103f7565b61234b6135e9565b5f546001600160a01b0316806123c657506123706001600160a01b0382161515613558565b5f80546001600160a01b0319166001600160a01b03928316179081905561239f9061239a906103eb565b6103eb565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103e75760803660031901126103e7576100206004356024356123fb81611f1c565b60443590612408826103f7565b60643592612415846103f7565b61241d613757565b614689565b346103e7575f3660031901126103e7576032546040516001600160a01b039091168152602090f35b346103e75760403660031901126103e75760243560043561246a826103f7565b6124726135e9565b805f5260056020525f60046124ad604083209461249c8161213760018060a01b0382168099615549565b8484526006602052604084206132a5565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103e7575f3660031901126103e75761250261591f565b61250a615976565b6040519060208201925f516020615dd45f395f51905f528452604083015260608201524660808201523060a082015260a0815261254860c0826105a7565b519020604051908152602090f35b1561255d57565b638219abe360e01b5f5260045ffd5b603380546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b156125b457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125f95760051b8101359060fe19813603018212156103e7570190565b6125c3565b35610e07816103f7565b903590601e19813603018212156103e757018035906001600160401b0382116103e7576020019181360383136103e757565b91908110156125f95760e0020190565b908160209103126103e75751610e0781611f1c565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936126d097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161265f565b9480356126dc816103f7565b6001600160a01b031660e085015260208101356126f8816103f7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561272d81610620565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561277b573d9061276282610d9b565b9161277060405193846105a7565b82523d5f602084013e565b606090565b604090610e0793928152816020820152019061153c565b90919392936127a78584146125ad565b5f945b8386106127b957505050509050565b6127c48685856125d7565b3560206127dc816127d68a89896125d7565b016125fe565b6127ec60606127d68b8a8a6125d7565b926128628a86888a8c61284660806128058784866125d7565b01359561283e6128348260a061281c82888a6125d7565b01359560c061282c83838b6125d7565b0135976125d7565b60e0810190612608565b96909561263a565b946040519a8b998a996326ae802b60e11b8b5260048b0161267f565b03815f305af190816128a9575b5061289e578561287d612751565b60405163f495148b60e01b815291829161289a9160048401612780565b0390fd5b6001909501946127aa565b6128c99060203d81116128ce575b6128c181836105a7565b81019061264a565b61286f565b503d6128b7565b906129ba9392916128e4613814565b6128ec61383b565b803592612901845f52600560205260405f2090565b9161293161291f604083019461291961239a876125fe565b90613882565b61292b61239a866125fe565b90612b17565b61293c343415612b3f565b6129d061295c865f526004602052600260405f2001600181540180915590565b9661296f602084013589819a8214612b5d565b61297b61239a866125fe565b9360608401968861298b896125fe565b966129c88c60808901359e8f60a08b019b6129a68d8d612608565b929091604051978896602088019a8b612b7b565b03601f1981018352826105a7565b5190206138c5565b6129e3876129dd856125fe565b87613ad6565b9190926001846129f281611d6d565b14612abf575b50612a0283611d6d565b60018303612a5c5750505090612a2e612a285f516020615bb45f395f51905f52936125fe565b916125fe565b6040516001600160a01b03909216958291612a4b91429184612c09565b0390a45b612a57613870565b600190565b612a998394612a94612ab7947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a9f956141e3565b6125fe565b926125fe565b9260405193849360018060a01b031698429185612bdf565b0390a4612a4f565b612af491935088612acf866125fe565b91612aed612ae6612adf8a6125fe565b9288612608565b3691610db6565b9289613c0a565b915f6129f8565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612b1f5750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612b475750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b66575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0797959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161265f565b6105da93606092969593608083019760018060a01b03168352602083015260408201520190611d7c565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615cb45f395f51905f52602052600160405f20015490565b15612c4c57565b63031206d560e51b5f5260045ffd5b80518210156125f95760209160051b010190565b5f516020615d945f395f51905f5280546001600160401b0319166001179055565b5f516020615d945f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612dc0576001600160a01b03841615612db157612cd46150cb565b612ce86001600160a01b0384161515612556565b6001600160a01b03811692612cfe841515612556565b60ff831615612da257612d6b92612d59612d5e92612d1a6150cb565b612d226150cb565b612d2a6150cb565b60ff195f516020615cd45f395f51905f5254165f516020615cd45f395f51905f5255612d546150cb565b6150f6565b615105565b612d666152be565b61256c565b5f516020615c745f395f51905f525f80a2612d8543603455565b612d8f8184614579565b81612d9957505050565b6105da92614792565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612dd75750565b6373a1390160e11b5f5260045260245ffd5b959394612e6b919597939297612dfd613814565b612e426001600160a01b038816612e14818b6148a3565b612e1c61383b565b612e2c61239a61239a60e46125fe565b811490612e3c61239a60e46125fe565b91612fc6565b612e63612e5361239a6101046125fe565b6001600160a01b038b1614612ff3565b838789614943565b9161012435612e9e81612e8786612e82878761301d565b61301d565b811015612e9887612e82888861301d565b9061302a565b612eac61239a6032546103eb565b90612eb86101046125fe565b906101443592612ec9610164613048565b92610184356101a43591833b156103e757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612fc157612f886112bc98612f9193612a4f9c612fa7575b50612f7f612f6a6101046125fe565b91612f736105ca565b9c8d5260208d01613052565b60408b01613052565b60608901613052565b608087015260a086015260c08501523691610db6565b80612fb55f612fbb936105a7565b80610a95565b5f612f5b565b612746565b15612fcf575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612ffa57565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610fec57565b15613033575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e0781610620565b6001600160a01b039091169052565b90600182811c9216801561308f575b602083101461307b57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613070565b906040516130a68161058c565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916130fe82613061565b808552916001811690811561316a575060011461312b575b505060a092916131279103846105a7565b0152565b5f908152602081209092505b81831061314e575050810160200181613127613116565b6020919350806001915483858901015201910190918492613137565b60ff191660208681019190915292151560051b850190920192508391506131279050613116565b600a821015611d775752565b906040516131aa81610571565b6040600782946131b981613099565b84526131cf60ff60068301541660208601613191565b0154910152565b156131de5750565b60405162461bcd60e51b81526020600482015290819061289a90602483019061153c565b1561320b575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061322d82610556565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061326082610609565b61326d60405191826105a7565b828152809261327e601f1991610609565b01905f5b82811061328e57505050565b602090613299613220565b82828501015201613282565b9060018060a01b03165f5260205260405f2090565b906040516132c781610556565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916133129061330990611503565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103e75751610e07816103f7565b5f5490949392906001600160a01b038116156133fc57604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906133ab90608487019061153c565b931691015203925af18015612fc1576105da925f916133cd575b50809461356e565b6133ef915060203d6020116133f5575b6133e781836105a7565b81019061332d565b5f6133c5565b503d6133dd565b6315aeca0d60e11b5f5260045ffd5b6040519061341a6020836105a7565b5f808352366020840137565b91929493909384845114806134b2575b806134a8575b613445906125ad565b5f5b8581101561349c578060051b8401359060be19853603018212156103e7576134956001926134758389612c5b565b51613480848c612c5b565b519061348c8589612c5b565b519289016128d5565b5001613447565b50945050505050600190565b508151851461343c565b5084865114613436565b604051906134c98261058c565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134f357565b63f0f15b9160e01b5f5260045ffd5b9060ff801983541691151516179055565b1561351b5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b1561355f57565b6302bfb33d60e51b5f5260045ffd5b905f6105da939261241d613757565b916135969183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b031633036135ae57565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba37b960691b6044820152606490fd5b5f516020615d345f395f51905f525f525f516020615cb45f395f51905f5260205260ff613636337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c6132a5565b54161561363f57565b63e2517d3f60e01b5f52336004525f516020615d345f395f51905f5260245260445ffd5b5f516020615d745f395f51905f525f525f516020615cb45f395f51905f5260205260ff6136b0337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b1436132a5565b5416156136b957565b63e2517d3f60e01b5f52336004525f516020615d745f395f51905f5260245260445ffd5b5f516020615c945f395f51905f525f525f516020615cb45f395f51905f5260205260ff61372a337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb4566132a5565b54161561373357565b63e2517d3f60e01b5f52336004525f516020615c945f395f51905f5260245260445ffd5b5f516020615cf45f395f51905f525f525f516020615cb45f395f51905f5260205260ff6137a4337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f6132a5565b5416156137ad57565b63e2517d3f60e01b5f52336004525f516020615cf45f395f51905f5260245260445ffd5b805f525f516020615cb45f395f51905f5260205260ff6137f43360405f206132a5565b5416156137fe5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615cd45f395f51905f52541661382c57565b63d93c066560e01b5f5260045ffd5b5f516020615d145f395f51905f525c6138615760015f516020615d145f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615d145f395f51905f525d565b610e07916001600160a01b031690614e6d565b1561389c57565b63b3f07a3960e01b5f5260045ffd5b156138b35750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a2f575b6138de90613895565b6138ff6138f95f516020615b945f395f51905f525460ff1690565b60ff1690565b9561390d84888110156138ab565b5f945f935f5b86811061392e57505050505050506105da91928110156138ab565b61398b61395a8361393d6155ec565b6042916040519161190160f01b8352600283015260228201522090565b61396e6139678489612c5b565b5160ff1690565b6139788487612c5b565b51906139848589612c5b565b5192614e80565b6001600160a01b0381811690881610806139bc575b6139ae575b50600101613913565b6001988901989096506139a5565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615cb45f395f51905f52602052613a2a613a23827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa56132a5565b6132a5565b5460ff1690565b6139a0565b50855183146138d5565b15613a4057565b6330d45fb160e01b5f5260045ffd5b908160209103126103e75751600a8110156103e75790565b6001600160a01b039091168152602081019190915260400190565b9150613ac360ff91613ab25f94613aa460325460018060a01b03161515613a39565b5f52600960205260405f2090565b6001600160a01b03909116906132a5565b5416613ace57600191565b506002905f90565b9190915f92613b17613a23613b07613af261239a6032546103eb565b94613aa46001600160a01b0387161515613a39565b6001600160a01b038416906132a5565b613b9e5791602091613b4195935f604051809881958294633f4bdec560e01b845260048401613a67565b03925af1928315612fc1575f93613b6d575b50600183613b6081611d6d565b03613b6757565b60019150565b613b9091935060203d602011613b97575b613b8881836105a7565b810190613a4f565b915f613b53565b503d613b7e565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610e079291019061153c565b6001600160a01b03918216815291166020820152604081019190915260600190565b9291909382613c1885612b09565b95613c3b6001613c30818060a01b038416809a6132a5565b015460a01c60ff1690565b92601881511180613f1b575b613c81575b5092613c5793614e98565b92613c6184611d6d565b60018414613c70575b50505090565b613c7992614f57565b5f8080613c6a565b613cb99250602081015160601c906020613c9f61239a6035546103eb565b926040518096819262483e5560e91b8352600483016112e2565b0381855afa8015612fc15787945f91613efc575b50613cd9575b50613c4c565b8286929394506002613cf38a5f52600460205260405f2090565b015491868960018d14159687613ec3575b5060209392915084908c613d46613d1f61239a6035546103eb565b948a15613ebd575f935b604051631eeed51360e01b81529a8b988997889660048801613ba9565b03925af1918215612fc1575f92613e9c575b50877f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613d8f8682901515815260200190565b0390a3613e8d57908491613da5575b5f80613cd3565b90508115613e2057613ddf92602085613dc261239a6035546103eb565b6040516323b872dd60e01b81529687928392309060048501613be8565b03815f8b5af1918215612fc157613c57948693613e01575b505b909350613d9e565b613e199060203d6020116128ce576128c181836105a7565b505f613df7565b613e5092602085613e3561239a6035546103eb565b604051632770a7eb60e21b8152968792839260048401613a67565b03815f8b5af1918215612fc157613c57948693613e6e575b50613df9565b613e869060203d6020116128ce576128c181836105a7565b505f613e68565b505050509091612a5792614f57565b613eb691925060203d6020116128ce576128c181836105a7565b905f613d58565b83613d29565b909192949550613ed293614e98565b613edb81611d6d565b60018103613eef5786929186898793613d04565b9850505050505050505090565b613f15915060203d6020116128ce576128c181836105a7565b5f613ccd565b506035546001600160a01b0390613f359061239a906103eb565b161515613c47565b15613f455750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103e75760405190613f708261058c565b819380358352602081013560208401526040810135613f8e816103f7565b6040840152613f9f606082016107c6565b60608401526080818101359084015260a0810135916001600160401b0383116103e75760a092613fcf9201610dec565b910152565b818110613fdf575050565b5f8155600101613fd4565b90601f8211613ff7575050565b6105da915f516020615b745f395f51905f525f5260205f20906020601f840160051c8301931061402f575b601f0160051c0190613fd4565b9091508190614022565b9190601f811161404857505050565b6105da925f5260205f20906020601f840160051c8301931061402f57601f0160051c0190613fd4565b8160011b915f199060031b1c19161790565b90600a811015611d775760ff80198354169116179055565b8151805182556020810151600183015560408101519192916140c9906001600160a01b03166002850161258e565b60608101516140e4906001600160a01b03166003850161258e565b6080810151600484015560a00151805160058401916001600160401b0382116105515761411b826141158554613061565b85614039565b602090601f831160011461417357826007959360409593614144935f92614168575b5050614071565b90555b614161602082015161415881611d6d565b60068601614083565b0151910155565b015190505f8061413d565b90601f19831691614187855f5260205f2090565b925f5b8181106141cb57509260019285926007989660409896106141b3575b505050811b019055614147565b01515f1960f88460031b161c191690555f80806141a6565b9293602060018192878601518155019501930161418a565b91608061427d61426e6002936142698761426461421d61359699833595865f52600760205261422260405f20602087013594858092615859565b613f3d565b1561428a576142566142366001544261301d565b915b61424b6142436105dc565b963690613f57565b865260208601613191565b6040840152610d0285612afb565b61409b565b612b09565b61117161239a604088016125fe565b930135920191825461301d565b6142565f91614238565b9061429f825f614fa0565b91826142a85750565b5f80525f516020615c545f395f51905f526020526001600160a01b03166142ef817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615859565b156142f75750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161431a8382614fa0565b9283614324575050565b815f525f516020615c545f395f51905f5260205261434f60405f209160018060a01b03168092615859565b15614358575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161437a8382615043565b9283614384575050565b815f525f516020615c545f395f51905f526020526143af60405f209160018060a01b03168092615549565b156143b8575050565b62a95f1b60e31b5f5260045260245260445ffd5b91906143d66150cb565b6143ea6001600160a01b0384161515612556565b6001600160a01b03811692614400841515612556565b60ff831615612da25761448d926144736144799261441c6150cb565b6144246150cb565b61442c6150cb565b60ff195f516020615cd45f395f51905f5254165f516020615cd45f395f51905f52556144566150cb565b61445e6150cb565b6144666150cb565b61446e6150cb565b614294565b50615105565b6144816150cb565b6201518060015561256c565b5f516020615c745f395f51905f525f80a26105da43603455565b600360606105da93805184556020810151600185015560408101516002850155015115159101613502565b156144da5750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c060049161451360018060a01b038251168561258e565b602081015161455e906001860190614534906001600160a01b03168261258e565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061353b565b6080810151600285015560a081015160038501550151910155565b90600191614588811515613558565b61462b838060a01b0383169261459f841515613558565b6145a885613558565b6145b1836157f0565b61464d575b6145da816145d5816145d0875f52600560205260405f2090565b615030565b6144d2565b6146266145e56105fa565b916145f08184613052565b6145fd8760208501613052565b86151560408401525f60608401525f60808401525f60a08401525f60c0840152613a1e85612b09565b6144fa565b60405183151581525f516020615d545f395f51905f529080602081015b0390a4565b6146846146586105eb565b8481525f60208201525f60408201525f606082015261467f855f52600460205260405f2090565b6144a7565b6145b6565b906146485f516020615d545f395f51905f52919493946146aa841515613558565b6001600160a01b0386169461078a906146c4871515613558565b6001600160a01b03811697614626906146de8a1515613558565b6146e7886157f0565b61474e575b614706816145d5816145d08c5f52600560205260405f2090565b6147256147116105fa565b9361471c8386613052565b60208501613052565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a1e88612b09565b6147806147596105eb565b8981525f60208201525f60408201525f606082015261467f8a5f52600460205260405f2090565b6146ec565b91908203918211610fec57565b906147a7915f52600660205260405f206132a5565b600181015460a01c60ff16156147c9576003018054918201809211610fec5755565b6004018054918203918211610fec5755565b156147e35750565b63290cd55f60e01b5f52600a811015611d775760045260245ffd5b9061480891614dd1565b60018060a01b036060820151168151915f516020615bb45f395f51905f526040820192614850610b72855160018060a01b03169660808601978489519160a089015193613c0a565b8251602090930151935194516040516001600160a01b03909616959182916146489142919084612c09565b156148835750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526148c78161213760405f2060018060a01b03831690614e6d565b825f52600460205260ff600360405f200154166148ff5760ff60016148f383613a1e6105da9697612b09565b015460a81c161561487b565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103e7578051916040602083015192015190565b1561493457565b6358e8878560e01b5f5260045ffd5b826060916149cc9795969361495d611177613b0784612b09565b61496d6109c76040830151151590565b614a6e575b5061498161239a6032546103eb565b916149966001600160a01b0384161515613a39565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612fc1575f955f905f93614a30575b509082916105da94939681988515159586614a25575b505084614a1a575b505082614a0f575b505061492d565b101590505f80614a08565b101592505f80614a00565b101594505f806149f8565b9050614a5b9196506105da93925060603d606011614a67575b614a5381836105a7565b810190614912565b919692939192916149e2565b503d614a49565b60c0614a80910151848082101561302a565b5f614972565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e091614ad9919086019061153c565b930152565b614afc81515f526004602052600160405f2001600181540180915590565b614b068251612b09565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708614b4b61239a6001614b44602088019561117161239a88516103eb565b01546103eb565b93805190614648614b5c85516103eb565b916040810190614b6c82516103eb565b90614b956080820196875160a0840194855198614b8f60c087019a8b519061301d565b93615319565b614bab614ba4825199516103eb565b93516103eb565b9460e0614bbb60608401516103eb565b9751935191519201519260405197889760018060a01b03169c429689614a86565b908160209103126103e7575190565b604051905f825f516020615b745f395f51905f525491614c0a83613061565b8083529260018116908115614c865750600114614c2e575b6105da925003836105a7565b505f516020615b745f395f51905f525f90815290915f516020615c145f395f51905f525b818310614c6a5750509060206105da92820101614c22565b6020919350806001915483858901015201910190918492614c52565b602092506105da94915060ff191682840152151560051b820101614c22565b604051905f825f516020615bd45f395f51905f525491614cc483613061565b8083529260018116908115614c865750600114614ce7576105da925003836105a7565b505f516020615bd45f395f51905f525f90815290915f516020615df45f395f51905f525b818310614d235750509060206105da92820101614c22565b6020919350806001915483858901015201910190918492614d0b565b60075f9182815582600182015582600282015582600382015582600482015560058101614d6c8154613061565b9081614d7f575b50508260068201550155565b601f8211600114614d9657849055505b5f80614d73565b614dbc614dcc926001601f614dae855f5260205f2090565b920160051c82019101613fd4565b5f81815260208120918190559055565b614d8f565b9190614ddb6134bc565b50825f526007602052614df18160405f20615549565b15614e5b57614e566105da91845f52600860205260405f20815f52602052610d02614e1e60405f20613099565b95614e3b614e2b82612b09565b61117161239a60408b01516103eb565b614e4f600260808a01519201918254614785565b9055612afb565b614d3f565b637f11bea960e01b5f5260045260245ffd5b6001915f520160205260405f2054151590565b91610e079391614e8f93615640565b909291926156c2565b6001600160a01b031692919060018414614f27578115614f1e57614ee79215614ef25760405163a9059cbb60e01b602082015291614edf9183916129ba9160248401613a67565b60059261573e565b15610e075750600190565b6040516340c10f1960e01b602082015291614f169183916129ba9160248401613a67565b60069261573e565b50505050600190565b90614f4993506109c79250614f3a611c67565b916001600160a01b0316615787565b614f5257600190565b600590565b90614f6c915f52600660205260405f206132a5565b600181015460a01c60ff1615614f8e576003018054918203918211610fec5755565b6004018054918201809211610fec5755565b805f525f516020615cb45f395f51905f5260205260ff614fc38360405f206132a5565b541661502a57805f525f516020615cb45f395f51905f52602052614fea8260405f206132a5565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610e07916001600160a01b031690615859565b805f525f516020615cb45f395f51905f5260205260ff6150668360405f206132a5565b54161561502a57805f525f516020615cb45f395f51905f5260205261508e8260405f206132a5565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615d945f395f51905f525460401c16156150e757565b631afcd79f60e31b5f5260045ffd5b6151029061445e6150cb565b50565b9061510e6150cb565b61511c60ff831615156134ec565b60409182519261512c81856105a7565b60098452682b30b634b230ba37b960b91b602085015261514e815191826105a7565b60058152640312e302e360dc1b60208201526151686150cb565b6151706150cb565b83516001600160401b038111610551576151a08161519b5f516020615b745f395f51905f5254613061565b613fea565b6020601f821160011461522f57816151dd93926151c9926105da97985f92614168575050614071565b5f516020615b745f395f51905f52556159a8565b6151f25f5f516020615bf45f395f51905f5255565b6152075f5f516020615db45f395f51905f5255565b60ff1660ff195f516020615b945f395f51905f525416175f516020615b945f395f51905f5255565b5f516020615b745f395f51905f525f52601f198216955f516020615c145f395f51905f52965f5b8181106152a657509660019284926151dd96956105da999a1061528e575b505050811b015f516020615b745f395f51905f52556159a8565b01515f1960f88460031b161c191690555f8080615274565b83830151895560019098019760209384019301615256565b6152c66150cb565b62015180600155565b156152d957505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561530a57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361538757506105da945061534f615341828661301d565b34143490612e98848861301d565b8061535b575b50614792565b61537c6153819161536d6033546103eb565b90615376611c67565b91615787565b615303565b5f615355565b90615393343415612b3f565b6153a86153a0828761301d565b3084896158a0565b8061543c575b506153c46109c76001613c3086613a1e87612b09565b6153d4575b506105da9350614792565b604051632770a7eb60e21b8152602081806153f3883060048401613a67565b03815f885af1918215612fc1576105da966154179387935f9161541d575b506152cf565b5f6153c9565b615436915060203d6020116128ce576128c181836105a7565b5f615411565b6154549061544e61239a6033546103eb565b876158db565b5f6153ae565b90813b156154d8575f516020615c345f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156154c05761510291615902565b5050346154c957565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125f9575f5260205f2001905f90565b80548015615535575f19019061552482826154f9565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146155e4575f198401848111610fec5783545f19810194908511610fec575f95858361559797610d02950361559d575b50505061550e565b55600190565b6155cd6155c7916155be6155b46155db95886154f9565b90549060031b1c90565b928391876154f9565b9061357d565b85905f5260205260405f2090565b555f808061558f565b505050505f90565b6155f461591f565b6155fc615976565b6040519060208201925f516020615dd45f395f51905f528452604083015260608201524660808201523060a082015260a0815261563a60c0826105a7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116156ad579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612fc1575f516001600160a01b038116156156a357905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611d7757565b6156cb816156b8565b806156d4575050565b6156dd816156b8565b600181036156f45763f645eedf60e01b5f5260045ffd5b6156fd816156b8565b60028103615718575063fce698f760e01b5f5260045260245ffd5b806157246003926156b8565b1461572c5750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f1615762612751565b901561502a57805161577e57503b1561577a57600190565b5f90565b60209150015190565b8147106157e15782516001600160a01b03909116925f9182916020018486620186a0f1906157b3612751565b91156157da57156157c6575b5050600190565b805161577e57503b1561577a575f806157bf565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461585457600254600160401b8110156105515761583d61582782600185940160025560026154f9565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6158638282614e6d565b61502a57805490600160401b821015610551578261588b6158278460018096018555846154f9565b90558054925f520160205260405f2055600190565b906158d6906158c86105da956040519586936323b872dd60e01b602086015260248501613be8565b03601f1981018452836105a7565b615a9d565b6158d66105da93926158c860405194859263a9059cbb60e01b602085015260248401613a67565b5f80610e0793602081519101845af4615919612751565b91615af5565b615927614beb565b8051908115615937576020012090565b50505f516020615bf45f395f51905f525480156159515790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61597e614ca5565b805190811561598e576020012090565b50505f516020615db45f395f51905f525480156159515790565b80519091906001600160401b038111610551576159e9816159d65f516020615bd45f395f51905f5254613061565b5f516020615bd45f395f51905f52614039565b602092601f8211600114615a1c57615a0b929382915f92614168575050614071565b5f516020615bd45f395f51905f5255565b5f516020615bd45f395f51905f525f52601f198216935f516020615df45f395f51905f52915f5b868110615a855750836001959610615a6d575b505050811b015f516020615bd45f395f51905f5255565b01515f1960f88460031b161c191690555f8080615a56565b91926020600181928685015181550194019201615a43565b905f602091828151910182855af115612746575f513d615aec57506001600160a01b0381163b155b615acc5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615ac5565b90615b195750805115615b0a57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b4a575b615b2a575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b2256feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a26469706673582212206fd1c69178c1c2b6bb4f0885f9e71000d79127662e03d80605d8b3cd1a8a7c4b64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeMetaData.ABI instead.
var BSCBridgeABI = BSCBridgeMetaData.ABI

// BSCBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BSCBridgeBinRuntime = "60806040526004361015610022575b3615610018575f80fd5b61002061359a565b005b5f3560e01c806301ffc9a7146103915780630b1d8d061461038c5780631313996b146103875780631938e0f214610382578063248a9ca31461037d5780632f2ff15d1461037857806336568abe14610373578063365d71e41461036e5780633aefddbf1461036957806342cde4e81461036457806348a00ed81461035f5780634be13f831461035a5780634d5d0056146103555780634ee078ba146103505780634f1ef2861461034b578063502cc5c01461034657806352d1902d146103415780635b605f5c1461033c5780635c975abb146103375780635fd262de14610332578063670e62681461032d578063761fe2d8146103285780637921487414610323578063814914b51461031e57806384b0196e1461031957806388d67d6d1461031457806389232a001461030f5780638ae87c5c1461030a57806391cca3db1461030557806391cf6d3e1461030057806391d14854146102fb5780639f9b4f1c146102f6578063a10bab0b146102f1578063a217fddf146102ec578063a3246ad3146102e7578063aa1bd0bc146102e2578063aa20ed47146102dd578063ad3cb1cc146102d8578063ae6893f8146102d3578063b2b49e2e146102ce578063b33eb36e146102c9578063b7f3358d146102c4578063b8aa837e146102bf578063bedb86fb146102ba578063bfbf6765146102b5578063cba25e4b146102b0578063cf56118e146102ab578063d477f05f146102a6578063d547741f146102a1578063d5717fc51461029c578063e2af6cd714610297578063edbbea3914610292578063f0702e8e1461028d578063f4509643146102885763f698da250361000e576124ea565b61244a565b612422565b6123d8565b612326565b6122ed565b6122be565b612260565b6121ec565b612182565b6120aa565b611fc1565b611f26565b611e99565b611e08565b611cfb565b611cc2565b611c7b565b611bf2565b611ba6565b611b2a565b611ace565b611aa6565b611993565b611951565b611934565b61190c565b6118a4565b6117f1565b6116e7565b6115bf565b611472565b6113f2565b61136c565b6112f5565b6111ed565b6111bf565b6110e0565b610ff1565b610f62565b610e0a565b610caa565b610c3d565b610be9565b610acb565b610a9f565b61095d565b6108fa565b61080a565b6107d1565b6107a0565b6106ef565b6104a9565b610408565b346103e75760203660031901126103e75760043563ffffffff60e01b81168091036103e757602090637965db0b60e01b81149081156103d6575b506040519015158152f35b6301ffc9a760e01b1490505f6103cb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103e757565b346103e75760203660031901126103e757600435610425816103f7565b61042d6135e9565b6001600160a01b0316610441811515612556565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103e7578235916001600160401b0383116103e7576020808501948460051b0101116103e757565b60403660031901126103e7576004356001600160401b0381116103e7576104d4903690600401610479565b602435916001600160401b0383116103e757366023840112156103e7576004830135916001600160401b0383116103e75736602460e08502860101116103e7576024610020940191612797565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761055157604052565b610521565b60e081019081106001600160401b0382111761055157604052565b606081019081106001600160401b0382111761055157604052565b60c081019081106001600160401b0382111761055157604052565b601f909101601f19168101906001600160401b0382119082101761055157604052565b604051906105da610100836105a7565b565b604051906105da6060836105a7565b604051906105da6080836105a7565b604051906105da60e0836105a7565b6001600160401b0381116105515760051b60200190565b60ff8116036103e757565b9080601f830112156103e757813561064281610609565b9261065060405194856105a7565b81845260208085019260051b8201019283116103e757602001905b8282106106785750505090565b60208091833561068781610620565b81520191019061066b565b9080601f830112156103e75781356106a981610609565b926106b760405194856105a7565b81845260208085019260051b8201019283116103e757602001905b8282106106df5750505090565b81358152602091820191016106d2565b60803660031901126103e7576004356001600160401b0381116103e75760c060031982360301126103e7576024356001600160401b0381116103e75761073990369060040161062b565b906044356001600160401b0381116103e757610759903690600401610692565b60643591906001600160401b0383116103e75761079c9361078161078a943690600401610692565b926004016128d5565b60405191829182901515815260200190565b0390f35b346103e75760203660031901126103e75760206107be600435612c27565b604051908152f35b35906105da826103f7565b346103e75760403660031901126103e7576100206024356004356107f4826103f7565b61080561080082612c27565b6137d1565b61430d565b346103e75760403660031901126103e75760043560243561082a816103f7565b336001600160a01b03821603610843576100209161436d565b63334bd91960e11b5f5260045ffd5b9060406003198301126103e7576004356001600160401b0381116103e7578261087d91600401610692565b91602435906001600160401b0382116103e757806023830112156103e75781600401356108a981610609565b926108b760405194856105a7565b8184526024602085019260051b8201019283116103e757602401905b8282106108e05750505090565b6020809183356108ef816103f7565b8152019101906108d3565b346103e75761090836610852565b906109168151835114612c45565b5f5b8251811015610020578061095661093160019385612c5b565b51838060a01b036109428488612c5b565b51169061095161080082612c27565b61436d565b5001610918565b346103e75760c03660031901126103e75760043561097a816103f7565b60243590610987826103f7565b60443561099381610620565b6084356064356109a2826103f7565b60a435925f516020615d945f395f51905f5254956109cb6109c78860ff9060401c1690565b1590565b966001600160401b031680159081610a8d575b6001149081610a83575b159081610a7a575b50610a6b57610a0b9587610a02612c6f565b610a5e57612cb1565b610a1157005b610a2f5f516020615d945f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a66612c90565b612cb1565b63f92ee8a960e01b5f5260045ffd5b9050155f6109f0565b303b1591506109e8565b8891506109de565b5f9103126103e757565b346103e7575f3660031901126103e757602060ff5f516020615b945f395f51905f525416604051908152f35b346103e75760603660031901126103e757602435600435604435610aee816103f7565b610af6613663565b610afe61383b565b815f526007602052610b1c83610b178160405f20614e6d565b612dcf565b80610b278484614dd1565b916001600160a01b031615610bd5575b8151915f516020615bb45f395f51905f526040820192610b84610b72855160018060a01b03169660808601978489519160a089015193613c0a565b610b7b81611d6d565b600181146147db565b8251602090930151935194516040516001600160a01b0390961695918291610baf9142919084612c09565b0390a45f516020615b545f395f51905f525f80a35f5f516020615d145f395f51905f525d005b5060608101516001600160a01b0316610b37565b346103e7575f3660031901126103e7575f546040516001600160a01b039091168152602090f35b9181601f840112156103e7578235916001600160401b0383116103e757602083818601950101116103e757565b6101c03660031901126103e757602435600435610c59826103f7565b604435610c65816103f7565b6064359060a43560843560c4356001600160401b0381116103e757610c8e903690600401610c10565b94909360e03660e31901126103e75761079c9761078a97612de9565b346103e75760403660031901126103e757610d89602435600435610ccc613814565b610cd461383b565b805f526007602052610ced82610b178160405f20614e6d565b610d846040610d14610d0f85610d0286612afb565b905f5260205260405f2090565b61319d565b610d71610d3482516080610d2a868301516103eb565b9101519087613a82565b50610d3e81611d6d565b610d4781611d6d565b83516020810182905290600190610d6b83604081015b03601f1981018552846105a7565b146131d6565b01518015908115610d91575b4291613202565b6147fe565b610020613870565b4281109150610d7d565b6001600160401b03811161055157601f01601f191660200190565b929192610dc282610d9b565b91610dd060405193846105a7565b8294818452818301116103e7578281602093845f960137010152565b9080601f830112156103e757816020610e0793359101610db6565b90565b60403660031901126103e757600435610e22816103f7565b6024356001600160401b0381116103e757610e41903690600401610dec565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f40575b50610f3157610e846135e9565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610f00575b50610ecd57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615c345f395f51905f528303610eec57610020925061545a565b632a87526960e21b5f52600483905260245ffd5b610f2391945060203d602011610f2a575b610f1b81836105a7565b810190614bdc565b925f610eac565b503d610f11565b63703e46dd60e11b5f5260045ffd5b5f516020615c345f395f51905f52546001600160a01b0316141590505f610e77565b346103e75760603660031901126103e757602435600435604435610f84613663565b815f526007602052610f9d83610b178160405f20614e6d565b4201804211610fec5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b613009565b346103e7575f3660031901126103e7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f315760206040515f516020615c345f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110c05750505090565b909192602060e0826110d56001948851611048565b0194019291016110b3565b346103e75760203660031901126103e757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106111a657505061112f925003836105a7565b6111398251613256565b915f5b81518110156111985760019061117c61117761115786612b09565b6111716111648588612c5b565b516001600160a01b031690565b906132a5565b6132ba565b6111868287612c5b565b526111918186612c5b565b500161113c565b6040518061079c868261109d565b845483526001948501948794506020909301920161111a565b346103e7575f3660031901126103e757602060ff5f516020615cd45f395f51905f5254166040519015158152f35b60e03660031901126103e757602435600435611208826103f7565b60443591611215836103f7565b6064359260c435916084359160a435916001600160401b0385116103e7576112c69661127b61124b6112bc973690600401610c10565b959096611256613814565b6001600160a01b03851694849061126d878d6148a3565b61127561383b565b8b614943565b9390926040519861128b8a610535565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610db6565b60e0820152614ade565b5f5f516020615d145f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103e75760803660031901126103e757602435600435611315826103f7565b604435906001600160401b0382116103e757366023830112156103e75761079c9261134d611360933690602481600401359101610db6565b906064359261135b84610620565b613342565b604051918291826112e2565b346103e75760403660031901126103e757602060ff6113a2602435600435611393826103f7565b5f526009845260405f206132a5565b54166040519015158152f35b90602080835192838152019201905f5b8181106113cb5750505090565b82518452602093840193909201916001016113be565b906020610e079281815201906113ae565b346103e75760203660031901126103e7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061144c5761079c85611440818703826105a7565b604051918291826113e1565b8254845260209093019260019283019201611429565b60e0810192916105da9190611048565b346103e75760403660031901126103e75761079c6114b1602435600435611498826103f7565b6114a0613220565b505f52600660205260405f206132a5565b6004604051916114c083610556565b80546001600160a01b0316835260018101546115159061150c906114ef6114e6826103eb565b60208801613052565b61150360a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611462565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161159590611587610e0797959693600f60f81b865260e0602087015260e086019061153c565b90848203604086015261153c565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526113ae565b346103e7575f3660031901126103e7575f516020615bf45f395f51905f52541580611653575b15611616576115f2614beb565b6115fa614ca5565b9061079c61160661340b565b6040519384933091469186611560565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615db45f395f51905f5254156115e5565b9080601f830112156103e757813561168081610609565b9261168e60405194856105a7565b81845260208085019260051b820101918383116103e75760208201905b8382106116ba57505050505090565b81356001600160401b0381116103e7576020916116dc87848094880101610692565b8152019101906116ab565b60803660031901126103e7576004356001600160401b0381116103e757611712903690600401610479565b90602435906001600160401b0382116103e757366023830112156103e757816004013561173e81610609565b9261174c60405194856105a7565b8184526024602085019260051b820101903682116103e75760248101925b8284106117c25750506044359150506001600160401b0381116103e757611795903690600401611669565b606435926001600160401b0384116103e75761079c946117bc61078a953690600401611669565b93613426565b83356001600160401b0381116103e7576020916117e683926024369187010161062b565b81520193019261176a565b346103e75760603660031901126103e75760043561180e816103f7565b6024359061181b826103f7565b60443561182781610620565b5f516020615d945f395f51905f52549260ff604085901c1615936001600160401b03168015908161189c575b6001149081611892575b159081611889575b50610a6b57610a0b9284611877612c6f565b156143cc57611884612c90565b6143cc565b9050155f611865565b303b15915061185d565b859150611853565b346103e75760403660031901126103e7576024356004356118c36135e9565b6118cb61383b565b6118d58282614dd1565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615d145f395f51905f525d005b346103e7575f3660031901126103e7576033546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e7576020603454604051908152f35b346103e75760403660031901126103e757602060ff6113a2602435600435611978826103f7565b5f525f516020615cb45f395f51905f52845260405f206132a5565b346103e75760403660031901126103e7576004356001600160401b0381116103e7576119c3903690600401610692565b6024356001600160401b0381116103e7576119e2903690600401610692565b906119f081518351146125ad565b5f5b82518110156100205780611a98611a0b60019385612c5b565b51611a168387612c5b565b5190611a20613814565b611a2861383b565b805f526007602052611a4182610b178160405f20614e6d565b610d846040611a56610d0f85610d0286612afb565b610d71611a6c82516080610d2a868301516103eb565b50611a7681611d6d565b611a7f81611d6d565b835160208101829052908a90610d6b8360408101610d5d565b611aa0613870565b016119f2565b346103e7575f3660031901126103e7576035546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611b0b5750505090565b82516001600160a01b0316845260209384019390920191600101611afe565b346103e75760203660031901126103e7576004355f525f516020615c545f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611b905761079c85611b84818703826105a7565b60405191829182611ae8565b8254845260209093019260019283019201611b6d565b346103e75760203660031901126103e7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611be56135e9565b80600155604051908152a1005b346103e75760403660031901126103e757602435600435611c11613663565b611c19613663565b611c2161383b565b805f526007602052611c3a82610b178160405f20614e6d565b611c4482826147fe565b5f516020615b545f395f51905f525f80a35f5f516020615d145f395f51905f525d005b60405190611c766020836105a7565b5f8252565b346103e7575f3660031901126103e75761079c604051611c9c6040826105a7565b60058152640352e302e360dc1b602082015260405191829160208352602083019061153c565b346103e75760203660031901126103e7576004355f526004602052600160405f20015460018101809111610fec57602090604051908152f35b346103e757611d0936610852565b90611d178151835114612c45565b5f5b82518110156100205780611d52611d3260019385612c5b565b51838060a01b03611d438488612c5b565b51169061080561080082612c27565b5001611d19565b634e487b7160e01b5f52602160045260245ffd5b600a1115611d7757565b611d59565b90600a821015611d775752565b6020815260606040611dee60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061153c565b93611e00602082015183860190611d7c565b015191015290565b346103e75760403660031901126103e757600435602435905f60408051611e2e81610571565b611e366134bc565b815282602082015201525f52600860205260405f20905f5260205261079c60405f20600760405191611e6783610571565b611e7081613099565b8352611e8660ff60068301541660208501613191565b0154604082015260405191829182611d89565b346103e75760203660031901126103e7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611edb81610620565b611ee36135e9565b16611eef8115156134ec565b8060ff195f516020615b945f395f51905f525416175f516020615b945f395f51905f5255604051908152a1005b801515036103e757565b346103e75760403660031901126103e757600435602435611f4681611f1c565b611f4e6136dd565b611f63825f52600360205260405f2054151590565b15611fae5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611fa381600360405f2001613502565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103e75760203660031901126103e757600435611fde81611f1c565b611fe66136dd565b1561204457611ff3613814565b600160ff195f516020615cd45f395f51905f525416175f516020615cd45f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615cd45f395f51905f525460ff81161561209b5760ff19165f516020615cd45f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103e75760803660031901126103e7576024356004356120ca826103f7565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356120f981611f1c565b60643561210581611f1c565b61210d6136dd565b845f5260056020526121718161216c855f209861213c8161213760018060a01b038216809d614e6d565b613513565b885f52600660205261215c866001612156848b5f206132a5565b0161353b565b885f526009602052865f206132a5565b613502565b8251911515825215156020820152a3005b346103e75760203660031901126103e75760043561219f816103f7565b6121a76135e9565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b346103e7575f3660031901126103e757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b81811061224a5761079c85611440818703826105a7565b8254845260209093019260019283019201612233565b346103e75760203660031901126103e75760043561227d816103f7565b6122856135e9565b6001600160a01b0316612299811515612556565b603380546001600160a01b031916821790555f516020615c745f395f51905f525f80a2005b346103e75760403660031901126103e7576100206024356004356122e1826103f7565b61095161080082612c27565b346103e75760203660031901126103e7576004355f526004602052600260405f20015460018101809111610fec57602090604051908152f35b346103e75760203660031901126103e757600435612343816103f7565b61234b6135e9565b5f546001600160a01b0316806123c657506123706001600160a01b0382161515613558565b5f80546001600160a01b0319166001600160a01b03928316179081905561239f9061239a906103eb565b6103eb565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103e75760803660031901126103e7576100206004356024356123fb81611f1c565b60443590612408826103f7565b60643592612415846103f7565b61241d613757565b614689565b346103e7575f3660031901126103e7576032546040516001600160a01b039091168152602090f35b346103e75760403660031901126103e75760243560043561246a826103f7565b6124726135e9565b805f5260056020525f60046124ad604083209461249c8161213760018060a01b0382168099615549565b8484526006602052604084206132a5565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103e7575f3660031901126103e75761250261591f565b61250a615976565b6040519060208201925f516020615dd45f395f51905f528452604083015260608201524660808201523060a082015260a0815261254860c0826105a7565b519020604051908152602090f35b1561255d57565b638219abe360e01b5f5260045ffd5b603380546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b156125b457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125f95760051b8101359060fe19813603018212156103e7570190565b6125c3565b35610e07816103f7565b903590601e19813603018212156103e757018035906001600160401b0382116103e7576020019181360383136103e757565b91908110156125f95760e0020190565b908160209103126103e75751610e0781611f1c565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936126d097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161265f565b9480356126dc816103f7565b6001600160a01b031660e085015260208101356126f8816103f7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561272d81610620565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561277b573d9061276282610d9b565b9161277060405193846105a7565b82523d5f602084013e565b606090565b604090610e0793928152816020820152019061153c565b90919392936127a78584146125ad565b5f945b8386106127b957505050509050565b6127c48685856125d7565b3560206127dc816127d68a89896125d7565b016125fe565b6127ec60606127d68b8a8a6125d7565b926128628a86888a8c61284660806128058784866125d7565b01359561283e6128348260a061281c82888a6125d7565b01359560c061282c83838b6125d7565b0135976125d7565b60e0810190612608565b96909561263a565b946040519a8b998a996326ae802b60e11b8b5260048b0161267f565b03815f305af190816128a9575b5061289e578561287d612751565b60405163f495148b60e01b815291829161289a9160048401612780565b0390fd5b6001909501946127aa565b6128c99060203d81116128ce575b6128c181836105a7565b81019061264a565b61286f565b503d6128b7565b906129ba9392916128e4613814565b6128ec61383b565b803592612901845f52600560205260405f2090565b9161293161291f604083019461291961239a876125fe565b90613882565b61292b61239a866125fe565b90612b17565b61293c343415612b3f565b6129d061295c865f526004602052600260405f2001600181540180915590565b9661296f602084013589819a8214612b5d565b61297b61239a866125fe565b9360608401968861298b896125fe565b966129c88c60808901359e8f60a08b019b6129a68d8d612608565b929091604051978896602088019a8b612b7b565b03601f1981018352826105a7565b5190206138c5565b6129e3876129dd856125fe565b87613ad6565b9190926001846129f281611d6d565b14612abf575b50612a0283611d6d565b60018303612a5c5750505090612a2e612a285f516020615bb45f395f51905f52936125fe565b916125fe565b6040516001600160a01b03909216958291612a4b91429184612c09565b0390a45b612a57613870565b600190565b612a998394612a94612ab7947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a9f956141e3565b6125fe565b926125fe565b9260405193849360018060a01b031698429185612bdf565b0390a4612a4f565b612af491935088612acf866125fe565b91612aed612ae6612adf8a6125fe565b9288612608565b3691610db6565b9289613c0a565b915f6129f8565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612b1f5750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612b475750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b66575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0797959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161265f565b6105da93606092969593608083019760018060a01b03168352602083015260408201520190611d7c565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615cb45f395f51905f52602052600160405f20015490565b15612c4c57565b63031206d560e51b5f5260045ffd5b80518210156125f95760209160051b010190565b5f516020615d945f395f51905f5280546001600160401b0319166001179055565b5f516020615d945f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612dc0576001600160a01b03841615612db157612cd46150cb565b612ce86001600160a01b0384161515612556565b6001600160a01b03811692612cfe841515612556565b60ff831615612da257612d6b92612d59612d5e92612d1a6150cb565b612d226150cb565b612d2a6150cb565b60ff195f516020615cd45f395f51905f5254165f516020615cd45f395f51905f5255612d546150cb565b6150f6565b615105565b612d666152be565b61256c565b5f516020615c745f395f51905f525f80a2612d8543603455565b612d8f8184614579565b81612d9957505050565b6105da92614792565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612dd75750565b6373a1390160e11b5f5260045260245ffd5b959394612e6b919597939297612dfd613814565b612e426001600160a01b038816612e14818b6148a3565b612e1c61383b565b612e2c61239a61239a60e46125fe565b811490612e3c61239a60e46125fe565b91612fc6565b612e63612e5361239a6101046125fe565b6001600160a01b038b1614612ff3565b838789614943565b9161012435612e9e81612e8786612e82878761301d565b61301d565b811015612e9887612e82888861301d565b9061302a565b612eac61239a6032546103eb565b90612eb86101046125fe565b906101443592612ec9610164613048565b92610184356101a43591833b156103e757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612fc157612f886112bc98612f9193612a4f9c612fa7575b50612f7f612f6a6101046125fe565b91612f736105ca565b9c8d5260208d01613052565b60408b01613052565b60608901613052565b608087015260a086015260c08501523691610db6565b80612fb55f612fbb936105a7565b80610a95565b5f612f5b565b612746565b15612fcf575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612ffa57565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610fec57565b15613033575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e0781610620565b6001600160a01b039091169052565b90600182811c9216801561308f575b602083101461307b57565b634e487b7160e01b5f52602260045260245ffd5b91607f1691613070565b906040516130a68161058c565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916130fe82613061565b808552916001811690811561316a575060011461312b575b505060a092916131279103846105a7565b0152565b5f908152602081209092505b81831061314e575050810160200181613127613116565b6020919350806001915483858901015201910190918492613137565b60ff191660208681019190915292151560051b850190920192508391506131279050613116565b600a821015611d775752565b906040516131aa81610571565b6040600782946131b981613099565b84526131cf60ff60068301541660208601613191565b0154910152565b156131de5750565b60405162461bcd60e51b81526020600482015290819061289a90602483019061153c565b1561320b575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061322d82610556565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061326082610609565b61326d60405191826105a7565b828152809261327e601f1991610609565b01905f5b82811061328e57505050565b602090613299613220565b82828501015201613282565b9060018060a01b03165f5260205260405f2090565b906040516132c781610556565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916133129061330990611503565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103e75751610e07816103f7565b5f5490949392906001600160a01b038116156133fc57604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906133ab90608487019061153c565b931691015203925af18015612fc1576105da925f916133cd575b50809461356e565b6133ef915060203d6020116133f5575b6133e781836105a7565b81019061332d565b5f6133c5565b503d6133dd565b6315aeca0d60e11b5f5260045ffd5b6040519061341a6020836105a7565b5f808352366020840137565b91929493909384845114806134b2575b806134a8575b613445906125ad565b5f5b8581101561349c578060051b8401359060be19853603018212156103e7576134956001926134758389612c5b565b51613480848c612c5b565b519061348c8589612c5b565b519289016128d5565b5001613447565b50945050505050600190565b508151851461343c565b5084865114613436565b604051906134c98261058c565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134f357565b63f0f15b9160e01b5f5260045ffd5b9060ff801983541691151516179055565b1561351b5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b1561355f57565b6302bfb33d60e51b5f5260045ffd5b905f6105da939261241d613757565b916135969183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b031633036135ae57565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba37b960691b6044820152606490fd5b5f516020615d345f395f51905f525f525f516020615cb45f395f51905f5260205260ff613636337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c6132a5565b54161561363f57565b63e2517d3f60e01b5f52336004525f516020615d345f395f51905f5260245260445ffd5b5f516020615d745f395f51905f525f525f516020615cb45f395f51905f5260205260ff6136b0337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b1436132a5565b5416156136b957565b63e2517d3f60e01b5f52336004525f516020615d745f395f51905f5260245260445ffd5b5f516020615c945f395f51905f525f525f516020615cb45f395f51905f5260205260ff61372a337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb4566132a5565b54161561373357565b63e2517d3f60e01b5f52336004525f516020615c945f395f51905f5260245260445ffd5b5f516020615cf45f395f51905f525f525f516020615cb45f395f51905f5260205260ff6137a4337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f6132a5565b5416156137ad57565b63e2517d3f60e01b5f52336004525f516020615cf45f395f51905f5260245260445ffd5b805f525f516020615cb45f395f51905f5260205260ff6137f43360405f206132a5565b5416156137fe5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615cd45f395f51905f52541661382c57565b63d93c066560e01b5f5260045ffd5b5f516020615d145f395f51905f525c6138615760015f516020615d145f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615d145f395f51905f525d565b610e07916001600160a01b031690614e6d565b1561389c57565b63b3f07a3960e01b5f5260045ffd5b156138b35750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a2f575b6138de90613895565b6138ff6138f95f516020615b945f395f51905f525460ff1690565b60ff1690565b9561390d84888110156138ab565b5f945f935f5b86811061392e57505050505050506105da91928110156138ab565b61398b61395a8361393d6155ec565b6042916040519161190160f01b8352600283015260228201522090565b61396e6139678489612c5b565b5160ff1690565b6139788487612c5b565b51906139848589612c5b565b5192614e80565b6001600160a01b0381811690881610806139bc575b6139ae575b50600101613913565b6001988901989096506139a5565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615cb45f395f51905f52602052613a2a613a23827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa56132a5565b6132a5565b5460ff1690565b6139a0565b50855183146138d5565b15613a4057565b6330d45fb160e01b5f5260045ffd5b908160209103126103e75751600a8110156103e75790565b6001600160a01b039091168152602081019190915260400190565b9150613ac360ff91613ab25f94613aa460325460018060a01b03161515613a39565b5f52600960205260405f2090565b6001600160a01b03909116906132a5565b5416613ace57600191565b506002905f90565b9190915f92613b17613a23613b07613af261239a6032546103eb565b94613aa46001600160a01b0387161515613a39565b6001600160a01b038416906132a5565b613b9e5791602091613b4195935f604051809881958294633f4bdec560e01b845260048401613a67565b03925af1928315612fc1575f93613b6d575b50600183613b6081611d6d565b03613b6757565b60019150565b613b9091935060203d602011613b97575b613b8881836105a7565b810190613a4f565b915f613b53565b503d613b7e565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610e079291019061153c565b6001600160a01b03918216815291166020820152604081019190915260600190565b9291909382613c1885612b09565b95613c3b6001613c30818060a01b038416809a6132a5565b015460a01c60ff1690565b92601881511180613f1b575b613c81575b5092613c5793614e98565b92613c6184611d6d565b60018414613c70575b50505090565b613c7992614f57565b5f8080613c6a565b613cb99250602081015160601c906020613c9f61239a6035546103eb565b926040518096819262483e5560e91b8352600483016112e2565b0381855afa8015612fc15787945f91613efc575b50613cd9575b50613c4c565b8286929394506002613cf38a5f52600460205260405f2090565b015491868960018d14159687613ec3575b5060209392915084908c613d46613d1f61239a6035546103eb565b948a15613ebd575f935b604051631eeed51360e01b81529a8b988997889660048801613ba9565b03925af1918215612fc1575f92613e9c575b50877f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613d8f8682901515815260200190565b0390a3613e8d57908491613da5575b5f80613cd3565b90508115613e2057613ddf92602085613dc261239a6035546103eb565b6040516323b872dd60e01b81529687928392309060048501613be8565b03815f8b5af1918215612fc157613c57948693613e01575b505b909350613d9e565b613e199060203d6020116128ce576128c181836105a7565b505f613df7565b613e5092602085613e3561239a6035546103eb565b604051632770a7eb60e21b8152968792839260048401613a67565b03815f8b5af1918215612fc157613c57948693613e6e575b50613df9565b613e869060203d6020116128ce576128c181836105a7565b505f613e68565b505050509091612a5792614f57565b613eb691925060203d6020116128ce576128c181836105a7565b905f613d58565b83613d29565b909192949550613ed293614e98565b613edb81611d6d565b60018103613eef5786929186898793613d04565b9850505050505050505090565b613f15915060203d6020116128ce576128c181836105a7565b5f613ccd565b506035546001600160a01b0390613f359061239a906103eb565b161515613c47565b15613f455750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103e75760405190613f708261058c565b819380358352602081013560208401526040810135613f8e816103f7565b6040840152613f9f606082016107c6565b60608401526080818101359084015260a0810135916001600160401b0383116103e75760a092613fcf9201610dec565b910152565b818110613fdf575050565b5f8155600101613fd4565b90601f8211613ff7575050565b6105da915f516020615b745f395f51905f525f5260205f20906020601f840160051c8301931061402f575b601f0160051c0190613fd4565b9091508190614022565b9190601f811161404857505050565b6105da925f5260205f20906020601f840160051c8301931061402f57601f0160051c0190613fd4565b8160011b915f199060031b1c19161790565b90600a811015611d775760ff80198354169116179055565b8151805182556020810151600183015560408101519192916140c9906001600160a01b03166002850161258e565b60608101516140e4906001600160a01b03166003850161258e565b6080810151600484015560a00151805160058401916001600160401b0382116105515761411b826141158554613061565b85614039565b602090601f831160011461417357826007959360409593614144935f92614168575b5050614071565b90555b614161602082015161415881611d6d565b60068601614083565b0151910155565b015190505f8061413d565b90601f19831691614187855f5260205f2090565b925f5b8181106141cb57509260019285926007989660409896106141b3575b505050811b019055614147565b01515f1960f88460031b161c191690555f80806141a6565b9293602060018192878601518155019501930161418a565b91608061427d61426e6002936142698761426461421d61359699833595865f52600760205261422260405f20602087013594858092615859565b613f3d565b1561428a576142566142366001544261301d565b915b61424b6142436105dc565b963690613f57565b865260208601613191565b6040840152610d0285612afb565b61409b565b612b09565b61117161239a604088016125fe565b930135920191825461301d565b6142565f91614238565b9061429f825f614fa0565b91826142a85750565b5f80525f516020615c545f395f51905f526020526001600160a01b03166142ef817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615859565b156142f75750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161431a8382614fa0565b9283614324575050565b815f525f516020615c545f395f51905f5260205261434f60405f209160018060a01b03168092615859565b15614358575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161437a8382615043565b9283614384575050565b815f525f516020615c545f395f51905f526020526143af60405f209160018060a01b03168092615549565b156143b8575050565b62a95f1b60e31b5f5260045260245260445ffd5b91906143d66150cb565b6143ea6001600160a01b0384161515612556565b6001600160a01b03811692614400841515612556565b60ff831615612da25761448d926144736144799261441c6150cb565b6144246150cb565b61442c6150cb565b60ff195f516020615cd45f395f51905f5254165f516020615cd45f395f51905f52556144566150cb565b61445e6150cb565b6144666150cb565b61446e6150cb565b614294565b50615105565b6144816150cb565b6201518060015561256c565b5f516020615c745f395f51905f525f80a26105da43603455565b600360606105da93805184556020810151600185015560408101516002850155015115159101613502565b156144da5750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c060049161451360018060a01b038251168561258e565b602081015161455e906001860190614534906001600160a01b03168261258e565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061353b565b6080810151600285015560a081015160038501550151910155565b90600191614588811515613558565b61462b838060a01b0383169261459f841515613558565b6145a885613558565b6145b1836157f0565b61464d575b6145da816145d5816145d0875f52600560205260405f2090565b615030565b6144d2565b6146266145e56105fa565b916145f08184613052565b6145fd8760208501613052565b86151560408401525f60608401525f60808401525f60a08401525f60c0840152613a1e85612b09565b6144fa565b60405183151581525f516020615d545f395f51905f529080602081015b0390a4565b6146846146586105eb565b8481525f60208201525f60408201525f606082015261467f855f52600460205260405f2090565b6144a7565b6145b6565b906146485f516020615d545f395f51905f52919493946146aa841515613558565b6001600160a01b0386169461078a906146c4871515613558565b6001600160a01b03811697614626906146de8a1515613558565b6146e7886157f0565b61474e575b614706816145d5816145d08c5f52600560205260405f2090565b6147256147116105fa565b9361471c8386613052565b60208501613052565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a1e88612b09565b6147806147596105eb565b8981525f60208201525f60408201525f606082015261467f8a5f52600460205260405f2090565b6146ec565b91908203918211610fec57565b906147a7915f52600660205260405f206132a5565b600181015460a01c60ff16156147c9576003018054918201809211610fec5755565b6004018054918203918211610fec5755565b156147e35750565b63290cd55f60e01b5f52600a811015611d775760045260245ffd5b9061480891614dd1565b60018060a01b036060820151168151915f516020615bb45f395f51905f526040820192614850610b72855160018060a01b03169660808601978489519160a089015193613c0a565b8251602090930151935194516040516001600160a01b03909616959182916146489142919084612c09565b156148835750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526148c78161213760405f2060018060a01b03831690614e6d565b825f52600460205260ff600360405f200154166148ff5760ff60016148f383613a1e6105da9697612b09565b015460a81c161561487b565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103e7578051916040602083015192015190565b1561493457565b6358e8878560e01b5f5260045ffd5b826060916149cc9795969361495d611177613b0784612b09565b61496d6109c76040830151151590565b614a6e575b5061498161239a6032546103eb565b916149966001600160a01b0384161515613a39565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612fc1575f955f905f93614a30575b509082916105da94939681988515159586614a25575b505084614a1a575b505082614a0f575b505061492d565b101590505f80614a08565b101592505f80614a00565b101594505f806149f8565b9050614a5b9196506105da93925060603d606011614a67575b614a5381836105a7565b810190614912565b919692939192916149e2565b503d614a49565b60c0614a80910151848082101561302a565b5f614972565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e091614ad9919086019061153c565b930152565b614afc81515f526004602052600160405f2001600181540180915590565b614b068251612b09565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708614b4b61239a6001614b44602088019561117161239a88516103eb565b01546103eb565b93805190614648614b5c85516103eb565b916040810190614b6c82516103eb565b90614b956080820196875160a0840194855198614b8f60c087019a8b519061301d565b93615319565b614bab614ba4825199516103eb565b93516103eb565b9460e0614bbb60608401516103eb565b9751935191519201519260405197889760018060a01b03169c429689614a86565b908160209103126103e7575190565b604051905f825f516020615b745f395f51905f525491614c0a83613061565b8083529260018116908115614c865750600114614c2e575b6105da925003836105a7565b505f516020615b745f395f51905f525f90815290915f516020615c145f395f51905f525b818310614c6a5750509060206105da92820101614c22565b6020919350806001915483858901015201910190918492614c52565b602092506105da94915060ff191682840152151560051b820101614c22565b604051905f825f516020615bd45f395f51905f525491614cc483613061565b8083529260018116908115614c865750600114614ce7576105da925003836105a7565b505f516020615bd45f395f51905f525f90815290915f516020615df45f395f51905f525b818310614d235750509060206105da92820101614c22565b6020919350806001915483858901015201910190918492614d0b565b60075f9182815582600182015582600282015582600382015582600482015560058101614d6c8154613061565b9081614d7f575b50508260068201550155565b601f8211600114614d9657849055505b5f80614d73565b614dbc614dcc926001601f614dae855f5260205f2090565b920160051c82019101613fd4565b5f81815260208120918190559055565b614d8f565b9190614ddb6134bc565b50825f526007602052614df18160405f20615549565b15614e5b57614e566105da91845f52600860205260405f20815f52602052610d02614e1e60405f20613099565b95614e3b614e2b82612b09565b61117161239a60408b01516103eb565b614e4f600260808a01519201918254614785565b9055612afb565b614d3f565b637f11bea960e01b5f5260045260245ffd5b6001915f520160205260405f2054151590565b91610e079391614e8f93615640565b909291926156c2565b6001600160a01b031692919060018414614f27578115614f1e57614ee79215614ef25760405163a9059cbb60e01b602082015291614edf9183916129ba9160248401613a67565b60059261573e565b15610e075750600190565b6040516340c10f1960e01b602082015291614f169183916129ba9160248401613a67565b60069261573e565b50505050600190565b90614f4993506109c79250614f3a611c67565b916001600160a01b0316615787565b614f5257600190565b600590565b90614f6c915f52600660205260405f206132a5565b600181015460a01c60ff1615614f8e576003018054918203918211610fec5755565b6004018054918201809211610fec5755565b805f525f516020615cb45f395f51905f5260205260ff614fc38360405f206132a5565b541661502a57805f525f516020615cb45f395f51905f52602052614fea8260405f206132a5565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610e07916001600160a01b031690615859565b805f525f516020615cb45f395f51905f5260205260ff6150668360405f206132a5565b54161561502a57805f525f516020615cb45f395f51905f5260205261508e8260405f206132a5565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615d945f395f51905f525460401c16156150e757565b631afcd79f60e31b5f5260045ffd5b6151029061445e6150cb565b50565b9061510e6150cb565b61511c60ff831615156134ec565b60409182519261512c81856105a7565b60098452682b30b634b230ba37b960b91b602085015261514e815191826105a7565b60058152640312e302e360dc1b60208201526151686150cb565b6151706150cb565b83516001600160401b038111610551576151a08161519b5f516020615b745f395f51905f5254613061565b613fea565b6020601f821160011461522f57816151dd93926151c9926105da97985f92614168575050614071565b5f516020615b745f395f51905f52556159a8565b6151f25f5f516020615bf45f395f51905f5255565b6152075f5f516020615db45f395f51905f5255565b60ff1660ff195f516020615b945f395f51905f525416175f516020615b945f395f51905f5255565b5f516020615b745f395f51905f525f52601f198216955f516020615c145f395f51905f52965f5b8181106152a657509660019284926151dd96956105da999a1061528e575b505050811b015f516020615b745f395f51905f52556159a8565b01515f1960f88460031b161c191690555f8080615274565b83830151895560019098019760209384019301615256565b6152c66150cb565b62015180600155565b156152d957505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561530a57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361538757506105da945061534f615341828661301d565b34143490612e98848861301d565b8061535b575b50614792565b61537c6153819161536d6033546103eb565b90615376611c67565b91615787565b615303565b5f615355565b90615393343415612b3f565b6153a86153a0828761301d565b3084896158a0565b8061543c575b506153c46109c76001613c3086613a1e87612b09565b6153d4575b506105da9350614792565b604051632770a7eb60e21b8152602081806153f3883060048401613a67565b03815f885af1918215612fc1576105da966154179387935f9161541d575b506152cf565b5f6153c9565b615436915060203d6020116128ce576128c181836105a7565b5f615411565b6154549061544e61239a6033546103eb565b876158db565b5f6153ae565b90813b156154d8575f516020615c345f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156154c05761510291615902565b5050346154c957565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125f9575f5260205f2001905f90565b80548015615535575f19019061552482826154f9565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146155e4575f198401848111610fec5783545f19810194908511610fec575f95858361559797610d02950361559d575b50505061550e565b55600190565b6155cd6155c7916155be6155b46155db95886154f9565b90549060031b1c90565b928391876154f9565b9061357d565b85905f5260205260405f2090565b555f808061558f565b505050505f90565b6155f461591f565b6155fc615976565b6040519060208201925f516020615dd45f395f51905f528452604083015260608201524660808201523060a082015260a0815261563a60c0826105a7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116156ad579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612fc1575f516001600160a01b038116156156a357905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611d7757565b6156cb816156b8565b806156d4575050565b6156dd816156b8565b600181036156f45763f645eedf60e01b5f5260045ffd5b6156fd816156b8565b60028103615718575063fce698f760e01b5f5260045260245ffd5b806157246003926156b8565b1461572c5750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f1615762612751565b901561502a57805161577e57503b1561577a57600190565b5f90565b60209150015190565b8147106157e15782516001600160a01b03909116925f9182916020018486620186a0f1906157b3612751565b91156157da57156157c6575b5050600190565b805161577e57503b1561577a575f806157bf565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461585457600254600160401b8110156105515761583d61582782600185940160025560026154f9565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6158638282614e6d565b61502a57805490600160401b821015610551578261588b6158278460018096018555846154f9565b90558054925f520160205260405f2055600190565b906158d6906158c86105da956040519586936323b872dd60e01b602086015260248501613be8565b03601f1981018452836105a7565b615a9d565b6158d66105da93926158c860405194859263a9059cbb60e01b602085015260248401613a67565b5f80610e0793602081519101845af4615919612751565b91615af5565b615927614beb565b8051908115615937576020012090565b50505f516020615bf45f395f51905f525480156159515790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61597e614ca5565b805190811561598e576020012090565b50505f516020615db45f395f51905f525480156159515790565b80519091906001600160401b038111610551576159e9816159d65f516020615bd45f395f51905f5254613061565b5f516020615bd45f395f51905f52614039565b602092601f8211600114615a1c57615a0b929382915f92614168575050614071565b5f516020615bd45f395f51905f5255565b5f516020615bd45f395f51905f525f52601f198216935f516020615df45f395f51905f52915f5b868110615a855750836001959610615a6d575b505050811b015f516020615bd45f395f51905f5255565b01515f1960f88460031b161c191690555f8080615a56565b91926020600181928685015181550194019201615a43565b905f602091828151910182855af115612746575f513d615aec57506001600160a01b0381163b155b615acc5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615ac5565b90615b195750805115615b0a57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b4a575b615b2a575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b2256feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a26469706673582212206fd1c69178c1c2b6bb4f0885f9e71000d79127662e03d80605d8b3cd1a8a7c4b64736f6c634300081c0033"

// Deprecated: Use BSCBridgeMetaData.Sigs instead.
// BSCBridgeFuncSigs maps the 4-byte function signature to its string representation.
var BSCBridgeFuncSigs = BSCBridgeMetaData.Sigs

// BSCBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BSCBridgeMetaData.Bin instead.
var BSCBridgeBin = BSCBridgeMetaData.Bin

// DeployBSCBridge deploys a new Ethereum contract, binding an instance of BSCBridge to it.
func DeployBSCBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCBridge, error) {
	parsed, err := BSCBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BSCBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCBridge{BSCBridgeCaller: BSCBridgeCaller{contract: contract}, BSCBridgeTransactor: BSCBridgeTransactor{contract: contract}, BSCBridgeFilterer: BSCBridgeFilterer{contract: contract}}, nil
}

// BSCBridge is an auto generated Go binding around an Ethereum contract.
type BSCBridge struct {
	BSCBridgeCaller     // Read-only binding to the contract
	BSCBridgeTransactor // Write-only binding to the contract
	BSCBridgeFilterer   // Log filterer for contract events
}

// BSCBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCBridgeSession struct {
	Contract     *BSCBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCBridgeCallerSession struct {
	Contract *BSCBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BSCBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCBridgeTransactorSession struct {
	Contract     *BSCBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BSCBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCBridgeRaw struct {
	Contract *BSCBridge // Generic contract binding to access the raw methods on
}

// BSCBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCBridgeCallerRaw struct {
	Contract *BSCBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BSCBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCBridgeTransactorRaw struct {
	Contract *BSCBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCBridge creates a new instance of BSCBridge, bound to a specific deployed contract.
func NewBSCBridge(address common.Address, backend bind.ContractBackend) (*BSCBridge, error) {
	contract, err := bindBSCBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCBridge{BSCBridgeCaller: BSCBridgeCaller{contract: contract}, BSCBridgeTransactor: BSCBridgeTransactor{contract: contract}, BSCBridgeFilterer: BSCBridgeFilterer{contract: contract}}, nil
}

// NewBSCBridgeCaller creates a new read-only instance of BSCBridge, bound to a specific deployed contract.
func NewBSCBridgeCaller(address common.Address, caller bind.ContractCaller) (*BSCBridgeCaller, error) {
	contract, err := bindBSCBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeCaller{contract: contract}, nil
}

// NewBSCBridgeTransactor creates a new write-only instance of BSCBridge, bound to a specific deployed contract.
func NewBSCBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCBridgeTransactor, error) {
	contract, err := bindBSCBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeTransactor{contract: contract}, nil
}

// NewBSCBridgeFilterer creates a new log filterer instance of BSCBridge, bound to a specific deployed contract.
func NewBSCBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCBridgeFilterer, error) {
	contract, err := bindBSCBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeFilterer{contract: contract}, nil
}

// bindBSCBridge binds a generic wrapper to an already deployed contract.
func bindBSCBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BSCBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridge *BSCBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridge.Contract.BSCBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridge *BSCBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridge.Contract.BSCBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridge *BSCBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridge.Contract.BSCBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridge *BSCBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridge *BSCBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridge *BSCBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridge *BSCBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridge *BSCBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridge.Contract.DEFAULTADMINROLE(&_BSCBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BSCBridge *BSCBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BSCBridge.Contract.DEFAULTADMINROLE(&_BSCBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridge *BSCBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridge *BSCBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridge.Contract.UPGRADEINTERFACEVERSION(&_BSCBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BSCBridge *BSCBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BSCBridge.Contract.UPGRADEINTERFACEVERSION(&_BSCBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridge *BSCBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridge *BSCBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridge.Contract.AllChainIDs(&_BSCBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_BSCBridge *BSCBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _BSCBridge.Contract.AllChainIDs(&_BSCBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridge *BSCBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridge *BSCBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridge.Contract.AllPendingIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_BSCBridge *BSCBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _BSCBridge.Contract.AllPendingIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridge *BSCBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridge *BSCBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridge.Contract.AllTokenPairs(&_BSCBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_BSCBridge *BSCBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _BSCBridge.Contract.AllTokenPairs(&_BSCBridge.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridge *BSCBridgeCaller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridge *BSCBridgeSession) BridgeExecutor() (common.Address, error) {
	return _BSCBridge.Contract.BridgeExecutor(&_BSCBridge.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_BSCBridge *BSCBridgeCallerSession) BridgeExecutor() (common.Address, error) {
	return _BSCBridge.Contract.BridgeExecutor(&_BSCBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridge *BSCBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridge *BSCBridgeSession) BridgeVerifier() (common.Address, error) {
	return _BSCBridge.Contract.BridgeVerifier(&_BSCBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_BSCBridge *BSCBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _BSCBridge.Contract.BridgeVerifier(&_BSCBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridge *BSCBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridge *BSCBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridge.Contract.CrossMintableERC20Code(&_BSCBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_BSCBridge *BSCBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _BSCBridge.Contract.CrossMintableERC20Code(&_BSCBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridge *BSCBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridge *BSCBridgeSession) Dev() (common.Address, error) {
	return _BSCBridge.Contract.Dev(&_BSCBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_BSCBridge *BSCBridgeCallerSession) Dev() (common.Address, error) {
	return _BSCBridge.Contract.Dev(&_BSCBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridge *BSCBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridge *BSCBridgeSession) DomainSeparator() ([32]byte, error) {
	return _BSCBridge.Contract.DomainSeparator(&_BSCBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BSCBridge *BSCBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _BSCBridge.Contract.DomainSeparator(&_BSCBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridge *BSCBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_BSCBridge *BSCBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridge.Contract.Eip712Domain(&_BSCBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BSCBridge *BSCBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BSCBridge.Contract.Eip712Domain(&_BSCBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridge.Contract.GetNextFinalizeIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridge.Contract.GetNextFinalizeIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridge.Contract.GetNextInitiateIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_BSCBridge *BSCBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _BSCBridge.Contract.GetNextInitiateIndex(&_BSCBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridge *BSCBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridge *BSCBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridge.Contract.GetPendingArguments(&_BSCBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_BSCBridge *BSCBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _BSCBridge.Contract.GetPendingArguments(&_BSCBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridge *BSCBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridge *BSCBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridge.Contract.GetRoleAdmin(&_BSCBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BSCBridge *BSCBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BSCBridge.Contract.GetRoleAdmin(&_BSCBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridge *BSCBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridge *BSCBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridge.Contract.GetRoleMembers(&_BSCBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_BSCBridge *BSCBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _BSCBridge.Contract.GetRoleMembers(&_BSCBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridge *BSCBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridge *BSCBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridge.Contract.GetTokenPair(&_BSCBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_BSCBridge *BSCBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _BSCBridge.Contract.GetTokenPair(&_BSCBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridge *BSCBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridge *BSCBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridge.Contract.HasRole(&_BSCBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BSCBridge *BSCBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BSCBridge.Contract.HasRole(&_BSCBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridge *BSCBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridge *BSCBridgeSession) InitializedAt() (*big.Int, error) {
	return _BSCBridge.Contract.InitializedAt(&_BSCBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_BSCBridge *BSCBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _BSCBridge.Contract.InitializedAt(&_BSCBridge.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridge *BSCBridgeCaller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridge *BSCBridgeSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridge.Contract.IsTokenFinalizePaused(&_BSCBridge.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_BSCBridge *BSCBridgeCallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _BSCBridge.Contract.IsTokenFinalizePaused(&_BSCBridge.CallOpts, remoteChainID, token)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridge *BSCBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridge *BSCBridgeSession) Paused() (bool, error) {
	return _BSCBridge.Contract.Paused(&_BSCBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BSCBridge *BSCBridgeCallerSession) Paused() (bool, error) {
	return _BSCBridge.Contract.Paused(&_BSCBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridge *BSCBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridge *BSCBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _BSCBridge.Contract.ProxiableUUID(&_BSCBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BSCBridge *BSCBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BSCBridge.Contract.ProxiableUUID(&_BSCBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridge *BSCBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridge *BSCBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridge.Contract.SupportsInterface(&_BSCBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BSCBridge *BSCBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BSCBridge.Contract.SupportsInterface(&_BSCBridge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridge *BSCBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridge *BSCBridgeSession) Threshold() (uint8, error) {
	return _BSCBridge.Contract.Threshold(&_BSCBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_BSCBridge *BSCBridgeCallerSession) Threshold() (uint8, error) {
	return _BSCBridge.Contract.Threshold(&_BSCBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridge *BSCBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.BridgeToken(&_BSCBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.BridgeToken(&_BSCBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.ChangeThreshold(&_BSCBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.ChangeThreshold(&_BSCBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridge *BSCBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridge *BSCBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.CreateToken(&_BSCBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_BSCBridge *BSCBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.CreateToken(&_BSCBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridge *BSCBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.FinalizeBridge(&_BSCBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.FinalizeBridge(&_BSCBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridge *BSCBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.FinalizeBridgeBatch(&_BSCBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.FinalizeBridgeBatch(&_BSCBridge.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.GrantRole(&_BSCBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.GrantRole(&_BSCBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.GrantRoleBatch(&_BSCBridge.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.GrantRoleBatch(&_BSCBridge.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "initialize", owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.Initialize(&_BSCBridge.TransactOpts, owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_BSCBridge *BSCBridgeTransactorSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _BSCBridge.Contract.Initialize(&_BSCBridge.TransactOpts, owner, dev_, threshold_)
}

// InitializeBSCBridge is a paid mutator transaction binding the contract method 0x3aefddbf.
//
// Solidity: function initializeBSCBridge(address owner, address dev_, uint8 threshold_, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_BSCBridge *BSCBridgeTransactor) InitializeBSCBridge(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "initializeBSCBridge", owner, dev_, threshold_, crossChainID, cross, crossInitialSupply)
}

// InitializeBSCBridge is a paid mutator transaction binding the contract method 0x3aefddbf.
//
// Solidity: function initializeBSCBridge(address owner, address dev_, uint8 threshold_, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_BSCBridge *BSCBridgeSession) InitializeBSCBridge(owner common.Address, dev_ common.Address, threshold_ uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.InitializeBSCBridge(&_BSCBridge.TransactOpts, owner, dev_, threshold_, crossChainID, cross, crossInitialSupply)
}

// InitializeBSCBridge is a paid mutator transaction binding the contract method 0x3aefddbf.
//
// Solidity: function initializeBSCBridge(address owner, address dev_, uint8 threshold_, uint256 crossChainID, address cross, uint256 crossInitialSupply) returns()
func (_BSCBridge *BSCBridgeTransactorSession) InitializeBSCBridge(owner common.Address, dev_ common.Address, threshold_ uint8, crossChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.InitializeBSCBridge(&_BSCBridge.TransactOpts, owner, dev_, threshold_, crossChainID, cross, crossInitialSupply)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ManualReleasePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ManualReleasePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridge *BSCBridgeTransactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridge *BSCBridgeSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.ManualReleasePendingWithRecipient(&_BSCBridge.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_BSCBridge *BSCBridgeTransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.ManualReleasePendingWithRecipient(&_BSCBridge.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridge *BSCBridgeSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.Contract.PermitBridgeToken(&_BSCBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_BSCBridge *BSCBridgeTransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.Contract.PermitBridgeToken(&_BSCBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridge *BSCBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridge *BSCBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.Contract.PermitBridgeTokenBatch(&_BSCBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_BSCBridge *BSCBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _BSCBridge.Contract.PermitBridgeTokenBatch(&_BSCBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridge *BSCBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridge *BSCBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RegisterToken(&_BSCBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_BSCBridge *BSCBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RegisterToken(&_BSCBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ReleasePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ReleasePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridge *BSCBridgeTransactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "releasePendingBatch", remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridge *BSCBridgeSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ReleasePendingBatch(&_BSCBridge.TransactOpts, remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_BSCBridge *BSCBridgeTransactorSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.ReleasePendingBatch(&_BSCBridge.TransactOpts, remoteChainIDs, indexes)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.RemovePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_BSCBridge *BSCBridgeTransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.RemovePending(&_BSCBridge.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridge *BSCBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridge *BSCBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RenounceRole(&_BSCBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BSCBridge *BSCBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RenounceRole(&_BSCBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RevokeRole(&_BSCBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BSCBridge *BSCBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RevokeRole(&_BSCBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RevokeRoleBatch(&_BSCBridge.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_BSCBridge *BSCBridgeTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.RevokeRoleBatch(&_BSCBridge.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridge *BSCBridgeTransactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridge *BSCBridgeSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeExecutor(&_BSCBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeExecutor(&_BSCBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridge *BSCBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridge *BSCBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeVerifier(&_BSCBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeVerifier(&_BSCBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridge *BSCBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridge *BSCBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetChainPause(&_BSCBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetChainPause(&_BSCBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridge *BSCBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridge *BSCBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetCrossMintableERC20Code(&_BSCBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetCrossMintableERC20Code(&_BSCBridge.TransactOpts, _crossMintableERC20Code)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridge *BSCBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridge *BSCBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetDev(&_BSCBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetDev(&_BSCBridge.TransactOpts, dev_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridge *BSCBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridge *BSCBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetPause(&_BSCBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetPause(&_BSCBridge.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridge *BSCBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridge *BSCBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetTokenPause(&_BSCBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetTokenPause(&_BSCBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridge *BSCBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridge *BSCBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetVerificationDelay(&_BSCBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetVerificationDelay(&_BSCBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridge *BSCBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridge *BSCBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetVerificationDelayExpiration(&_BSCBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetVerificationDelayExpiration(&_BSCBridge.TransactOpts, remoteChainID, index, delay)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridge *BSCBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridge *BSCBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.UnregisterToken(&_BSCBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_BSCBridge *BSCBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.UnregisterToken(&_BSCBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridge *BSCBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridge *BSCBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.UpgradeToAndCall(&_BSCBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BSCBridge *BSCBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BSCBridge.Contract.UpgradeToAndCall(&_BSCBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridge *BSCBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridge *BSCBridgeSession) Receive() (*types.Transaction, error) {
	return _BSCBridge.Contract.Receive(&_BSCBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BSCBridge *BSCBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _BSCBridge.Contract.Receive(&_BSCBridge.TransactOpts)
}

// BSCBridgeBridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the BSCBridge contract.
type BSCBridgeBridgeExecutorSetIterator struct {
	Event *BSCBridgeBridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeExecutorSet)
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
		it.Event = new(BSCBridgeBridgeExecutorSet)
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
func (it *BSCBridgeBridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeExecutorSet represents a BridgeExecutorSet event raised by the BSCBridge contract.
type BSCBridgeBridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*BSCBridgeBridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeExecutorSetIterator{contract: _BSCBridge.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeExecutorSet)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeExecutorSet(log types.Log) (*BSCBridgeBridgeExecutorSet, error) {
	event := new(BSCBridgeBridgeExecutorSet)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the BSCBridge contract.
type BSCBridgeBridgeFinalizedIterator struct {
	Event *BSCBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeFinalized)
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
		it.Event = new(BSCBridgeBridgeFinalized)
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
func (it *BSCBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeFinalized represents a BridgeFinalized event raised by the BSCBridge contract.
type BSCBridgeBridgeFinalized struct {
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
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeFinalizedIterator{contract: _BSCBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeFinalized)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeFinalized(log types.Log) (*BSCBridgeBridgeFinalized, error) {
	event := new(BSCBridgeBridgeFinalized)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the BSCBridge contract.
type BSCBridgeBridgeInitiatedIterator struct {
	Event *BSCBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeInitiated)
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
		it.Event = new(BSCBridgeBridgeInitiated)
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
func (it *BSCBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeInitiated represents a BridgeInitiated event raised by the BSCBridge contract.
type BSCBridgeBridgeInitiated struct {
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
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*BSCBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeInitiatedIterator{contract: _BSCBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeInitiated)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeInitiated(log types.Log) (*BSCBridgeBridgeInitiated, error) {
	event := new(BSCBridgeBridgeInitiated)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the BSCBridge contract.
type BSCBridgeBridgePendingIterator struct {
	Event *BSCBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgePending)
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
		it.Event = new(BSCBridgeBridgePending)
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
func (it *BSCBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgePending represents a BridgePending event raised by the BSCBridge contract.
type BSCBridgeBridgePending struct {
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
func (_BSCBridge *BSCBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeBridgePendingIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgePendingIterator{contract: _BSCBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgePending)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseBridgePending(log types.Log) (*BSCBridgeBridgePending, error) {
	event := new(BSCBridgeBridgePending)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the BSCBridge contract.
type BSCBridgeBridgeVerifierSetIterator struct {
	Event *BSCBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeVerifierSet)
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
		it.Event = new(BSCBridgeBridgeVerifierSet)
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
func (it *BSCBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the BSCBridge contract.
type BSCBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*BSCBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeVerifierSetIterator{contract: _BSCBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeVerifierSet)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*BSCBridgeBridgeVerifierSet, error) {
	event := new(BSCBridgeBridgeVerifierSet)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the BSCBridge contract.
type BSCBridgeChainPauseSetIterator struct {
	Event *BSCBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeChainPauseSet)
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
		it.Event = new(BSCBridgeChainPauseSet)
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
func (it *BSCBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeChainPauseSet represents a ChainPauseSet event raised by the BSCBridge contract.
type BSCBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridge *BSCBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BSCBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeChainPauseSetIterator{contract: _BSCBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_BSCBridge *BSCBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeChainPauseSet)
				if err := _BSCBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseChainPauseSet(log types.Log) (*BSCBridgeChainPauseSet, error) {
	event := new(BSCBridgeChainPauseSet)
	if err := _BSCBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the BSCBridge contract.
type BSCBridgeCrossMintableERC20CodeSetIterator struct {
	Event *BSCBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(BSCBridgeCrossMintableERC20CodeSet)
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
func (it *BSCBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the BSCBridge contract.
type BSCBridgeCrossMintableERC20CodeSet struct {
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BSCBridge *BSCBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*BSCBridgeCrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeCrossMintableERC20CodeSetIterator{contract: _BSCBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BSCBridge *BSCBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeCrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeCrossMintableERC20CodeSet)
				if err := _BSCBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BSCBridge *BSCBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*BSCBridgeCrossMintableERC20CodeSet, error) {
	event := new(BSCBridgeCrossMintableERC20CodeSet)
	if err := _BSCBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeDevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the BSCBridge contract.
type BSCBridgeDevSetIterator struct {
	Event *BSCBridgeDevSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeDevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeDevSet)
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
		it.Event = new(BSCBridgeDevSet)
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
func (it *BSCBridgeDevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeDevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeDevSet represents a DevSet event raised by the BSCBridge contract.
type BSCBridgeDevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridge *BSCBridgeFilterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*BSCBridgeDevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeDevSetIterator{contract: _BSCBridge.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_BSCBridge *BSCBridgeFilterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeDevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeDevSet)
				if err := _BSCBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseDevSet(log types.Log) (*BSCBridgeDevSet, error) {
	event := new(BSCBridgeDevSet)
	if err := _BSCBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BSCBridge contract.
type BSCBridgeEIP712DomainChangedIterator struct {
	Event *BSCBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeEIP712DomainChanged)
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
		it.Event = new(BSCBridgeEIP712DomainChanged)
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
func (it *BSCBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the BSCBridge contract.
type BSCBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridge *BSCBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BSCBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeEIP712DomainChangedIterator{contract: _BSCBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BSCBridge *BSCBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeEIP712DomainChanged)
				if err := _BSCBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*BSCBridgeEIP712DomainChanged, error) {
	event := new(BSCBridgeEIP712DomainChanged)
	if err := _BSCBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BSCBridge contract.
type BSCBridgeExtraCallExecutedIterator struct {
	Event *BSCBridgeExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *BSCBridgeExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeExtraCallExecuted)
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
		it.Event = new(BSCBridgeExtraCallExecuted)
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
func (it *BSCBridgeExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeExtraCallExecuted represents a ExtraCallExecuted event raised by the BSCBridge contract.
type BSCBridgeExtraCallExecuted struct {
	FromChainID *big.Int
	Index       *big.Int
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_BSCBridge *BSCBridgeFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BSCBridgeExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeExtraCallExecutedIterator{contract: _BSCBridge.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_BSCBridge *BSCBridgeFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BSCBridgeExtraCallExecuted, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeExtraCallExecuted)
				if err := _BSCBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseExtraCallExecuted(log types.Log) (*BSCBridgeExtraCallExecuted, error) {
	event := new(BSCBridgeExtraCallExecuted)
	if err := _BSCBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BSCBridge contract.
type BSCBridgeInitializedIterator struct {
	Event *BSCBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BSCBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeInitialized)
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
		it.Event = new(BSCBridgeInitialized)
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
func (it *BSCBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeInitialized represents a Initialized event raised by the BSCBridge contract.
type BSCBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridge *BSCBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BSCBridgeInitializedIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeInitializedIterator{contract: _BSCBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BSCBridge *BSCBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BSCBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeInitialized)
				if err := _BSCBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseInitialized(log types.Log) (*BSCBridgeInitialized, error) {
	event := new(BSCBridgeInitialized)
	if err := _BSCBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the BSCBridge contract.
type BSCBridgeManualReleasedIterator struct {
	Event *BSCBridgeManualReleased // Event containing the contract specifics and raw log

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
func (it *BSCBridgeManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeManualReleased)
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
		it.Event = new(BSCBridgeManualReleased)
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
func (it *BSCBridgeManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeManualReleased represents a ManualReleased event raised by the BSCBridge contract.
type BSCBridgeManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridge *BSCBridgeFilterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgeManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeManualReleasedIterator{contract: _BSCBridge.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridge *BSCBridgeFilterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *BSCBridgeManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeManualReleased)
				if err := _BSCBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseManualReleased(log types.Log) (*BSCBridgeManualReleased, error) {
	event := new(BSCBridgeManualReleased)
	if err := _BSCBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BSCBridge contract.
type BSCBridgePausedIterator struct {
	Event *BSCBridgePaused // Event containing the contract specifics and raw log

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
func (it *BSCBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgePaused)
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
		it.Event = new(BSCBridgePaused)
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
func (it *BSCBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgePaused represents a Paused event raised by the BSCBridge contract.
type BSCBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridge *BSCBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BSCBridgePausedIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgePausedIterator{contract: _BSCBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BSCBridge *BSCBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BSCBridgePaused) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgePaused)
				if err := _BSCBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParsePaused(log types.Log) (*BSCBridgePaused, error) {
	event := new(BSCBridgePaused)
	if err := _BSCBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgePendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the BSCBridge contract.
type BSCBridgePendingRemovedIterator struct {
	Event *BSCBridgePendingRemoved // Event containing the contract specifics and raw log

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
func (it *BSCBridgePendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgePendingRemoved)
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
		it.Event = new(BSCBridgePendingRemoved)
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
func (it *BSCBridgePendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgePendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgePendingRemoved represents a PendingRemoved event raised by the BSCBridge contract.
type BSCBridgePendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridge *BSCBridgeFilterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*BSCBridgePendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgePendingRemovedIterator{contract: _BSCBridge.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_BSCBridge *BSCBridgeFilterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *BSCBridgePendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgePendingRemoved)
				if err := _BSCBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParsePendingRemoved(log types.Log) (*BSCBridgePendingRemoved, error) {
	event := new(BSCBridgePendingRemoved)
	if err := _BSCBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BSCBridge contract.
type BSCBridgeRoleAdminChangedIterator struct {
	Event *BSCBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeRoleAdminChanged)
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
		it.Event = new(BSCBridgeRoleAdminChanged)
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
func (it *BSCBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the BSCBridge contract.
type BSCBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridge *BSCBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BSCBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeRoleAdminChangedIterator{contract: _BSCBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BSCBridge *BSCBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeRoleAdminChanged)
				if err := _BSCBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BSCBridgeRoleAdminChanged, error) {
	event := new(BSCBridgeRoleAdminChanged)
	if err := _BSCBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BSCBridge contract.
type BSCBridgeRoleGrantedIterator struct {
	Event *BSCBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BSCBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeRoleGranted)
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
		it.Event = new(BSCBridgeRoleGranted)
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
func (it *BSCBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeRoleGranted represents a RoleGranted event raised by the BSCBridge contract.
type BSCBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridge *BSCBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeRoleGrantedIterator{contract: _BSCBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridge *BSCBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BSCBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeRoleGranted)
				if err := _BSCBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseRoleGranted(log types.Log) (*BSCBridgeRoleGranted, error) {
	event := new(BSCBridgeRoleGranted)
	if err := _BSCBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BSCBridge contract.
type BSCBridgeRoleRevokedIterator struct {
	Event *BSCBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BSCBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeRoleRevoked)
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
		it.Event = new(BSCBridgeRoleRevoked)
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
func (it *BSCBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeRoleRevoked represents a RoleRevoked event raised by the BSCBridge contract.
type BSCBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridge *BSCBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BSCBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeRoleRevokedIterator{contract: _BSCBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BSCBridge *BSCBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BSCBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeRoleRevoked)
				if err := _BSCBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseRoleRevoked(log types.Log) (*BSCBridgeRoleRevoked, error) {
	event := new(BSCBridgeRoleRevoked)
	if err := _BSCBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the BSCBridge contract.
type BSCBridgeThresholdChangedIterator struct {
	Event *BSCBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BSCBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeThresholdChanged)
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
		it.Event = new(BSCBridgeThresholdChanged)
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
func (it *BSCBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeThresholdChanged represents a ThresholdChanged event raised by the BSCBridge contract.
type BSCBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridge *BSCBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BSCBridgeThresholdChangedIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeThresholdChangedIterator{contract: _BSCBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_BSCBridge *BSCBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BSCBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeThresholdChanged)
				if err := _BSCBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseThresholdChanged(log types.Log) (*BSCBridgeThresholdChanged, error) {
	event := new(BSCBridgeThresholdChanged)
	if err := _BSCBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeTokenFinalizePauseSetIterator is returned from FilterTokenFinalizePauseSet and is used to iterate over the raw logs and unpacked data for TokenFinalizePauseSet events raised by the BSCBridge contract.
type BSCBridgeTokenFinalizePauseSetIterator struct {
	Event *BSCBridgeTokenFinalizePauseSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeTokenFinalizePauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeTokenFinalizePauseSet)
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
		it.Event = new(BSCBridgeTokenFinalizePauseSet)
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
func (it *BSCBridgeTokenFinalizePauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeTokenFinalizePauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeTokenFinalizePauseSet represents a TokenFinalizePauseSet event raised by the BSCBridge contract.
type BSCBridgeTokenFinalizePauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenFinalizePauseSet is a free log retrieval operation binding the contract event 0x02c5bc0a5f43e2797484ce130ba7fd2ade9dfa8e41f4a78240c0b08817727188.
//
// Solidity: event TokenFinalizePauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_BSCBridge *BSCBridgeFilterer) FilterTokenFinalizePauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BSCBridgeTokenFinalizePauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "TokenFinalizePauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeTokenFinalizePauseSetIterator{contract: _BSCBridge.contract, event: "TokenFinalizePauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenFinalizePauseSet is a free log subscription operation binding the contract event 0x02c5bc0a5f43e2797484ce130ba7fd2ade9dfa8e41f4a78240c0b08817727188.
//
// Solidity: event TokenFinalizePauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_BSCBridge *BSCBridgeFilterer) WatchTokenFinalizePauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeTokenFinalizePauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "TokenFinalizePauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeTokenFinalizePauseSet)
				if err := _BSCBridge.contract.UnpackLog(event, "TokenFinalizePauseSet", log); err != nil {
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

// ParseTokenFinalizePauseSet is a log parse operation binding the contract event 0x02c5bc0a5f43e2797484ce130ba7fd2ade9dfa8e41f4a78240c0b08817727188.
//
// Solidity: event TokenFinalizePauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
func (_BSCBridge *BSCBridgeFilterer) ParseTokenFinalizePauseSet(log types.Log) (*BSCBridgeTokenFinalizePauseSet, error) {
	event := new(BSCBridgeTokenFinalizePauseSet)
	if err := _BSCBridge.contract.UnpackLog(event, "TokenFinalizePauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the BSCBridge contract.
type BSCBridgeTokenPairRegisteredIterator struct {
	Event *BSCBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *BSCBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeTokenPairRegistered)
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
		it.Event = new(BSCBridgeTokenPairRegistered)
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
func (it *BSCBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the BSCBridge contract.
type BSCBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridge *BSCBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*BSCBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeTokenPairRegisteredIterator{contract: _BSCBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_BSCBridge *BSCBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *BSCBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeTokenPairRegistered)
				if err := _BSCBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*BSCBridgeTokenPairRegistered, error) {
	event := new(BSCBridgeTokenPairRegistered)
	if err := _BSCBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the BSCBridge contract.
type BSCBridgeTokenPairUnregisteredIterator struct {
	Event *BSCBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *BSCBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeTokenPairUnregistered)
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
		it.Event = new(BSCBridgeTokenPairUnregistered)
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
func (it *BSCBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the BSCBridge contract.
type BSCBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BSCBridge *BSCBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*BSCBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeTokenPairUnregisteredIterator{contract: _BSCBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_BSCBridge *BSCBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *BSCBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeTokenPairUnregistered)
				if err := _BSCBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*BSCBridgeTokenPairUnregistered, error) {
	event := new(BSCBridgeTokenPairUnregistered)
	if err := _BSCBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the BSCBridge contract.
type BSCBridgeTokenPauseSetIterator struct {
	Event *BSCBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeTokenPauseSet)
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
		it.Event = new(BSCBridgeTokenPauseSet)
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
func (it *BSCBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeTokenPauseSet represents a TokenPauseSet event raised by the BSCBridge contract.
type BSCBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridge *BSCBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*BSCBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeTokenPauseSetIterator{contract: _BSCBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_BSCBridge *BSCBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeTokenPauseSet)
				if err := _BSCBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseTokenPauseSet(log types.Log) (*BSCBridgeTokenPauseSet, error) {
	event := new(BSCBridgeTokenPauseSet)
	if err := _BSCBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BSCBridge contract.
type BSCBridgeUnpausedIterator struct {
	Event *BSCBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BSCBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeUnpaused)
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
		it.Event = new(BSCBridgeUnpaused)
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
func (it *BSCBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeUnpaused represents a Unpaused event raised by the BSCBridge contract.
type BSCBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridge *BSCBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BSCBridgeUnpausedIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeUnpausedIterator{contract: _BSCBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BSCBridge *BSCBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BSCBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeUnpaused)
				if err := _BSCBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseUnpaused(log types.Log) (*BSCBridgeUnpaused, error) {
	event := new(BSCBridgeUnpaused)
	if err := _BSCBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BSCBridge contract.
type BSCBridgeUpgradedIterator struct {
	Event *BSCBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BSCBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeUpgraded)
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
		it.Event = new(BSCBridgeUpgraded)
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
func (it *BSCBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeUpgraded represents a Upgraded event raised by the BSCBridge contract.
type BSCBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridge *BSCBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BSCBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeUpgradedIterator{contract: _BSCBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BSCBridge *BSCBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BSCBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeUpgraded)
				if err := _BSCBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseUpgraded(log types.Log) (*BSCBridgeUpgraded, error) {
	event := new(BSCBridgeUpgraded)
	if err := _BSCBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the BSCBridge contract.
type BSCBridgeVerificationDelayExpirationSetIterator struct {
	Event *BSCBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeVerificationDelayExpirationSet)
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
		it.Event = new(BSCBridgeVerificationDelayExpirationSet)
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
func (it *BSCBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the BSCBridge contract.
type BSCBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BSCBridge *BSCBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*BSCBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeVerificationDelayExpirationSetIterator{contract: _BSCBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_BSCBridge *BSCBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeVerificationDelayExpirationSet)
				if err := _BSCBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*BSCBridgeVerificationDelayExpirationSet, error) {
	event := new(BSCBridgeVerificationDelayExpirationSet)
	if err := _BSCBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the BSCBridge contract.
type BSCBridgeVerificationDelaySetIterator struct {
	Event *BSCBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeVerificationDelaySet)
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
		it.Event = new(BSCBridgeVerificationDelaySet)
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
func (it *BSCBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the BSCBridge contract.
type BSCBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridge *BSCBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*BSCBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &BSCBridgeVerificationDelaySetIterator{contract: _BSCBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_BSCBridge *BSCBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *BSCBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeVerificationDelaySet)
				if err := _BSCBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_BSCBridge *BSCBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*BSCBridgeVerificationDelaySet, error) {
	event := new(BSCBridgeVerificationDelaySet)
	if err := _BSCBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
