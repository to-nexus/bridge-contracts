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

// CrossBridgeMetaData contains all meta data concerning the CrossBridge contract.
var CrossBridgeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecutor\",\"outputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupplyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"bscChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeCrossBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isTokenFinalizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecutor\",\"name\":\"_bridgeExecutor\",\"type\":\"address\"}],\"name\":\"setBridgeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"setCrossSupplyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecutor\",\"type\":\"address\"}],\"name\":\"BridgeExecutorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"oldCode\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"newCode\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"CrossSupplyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatePause\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalizePause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
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
		"cd959706": "crossSupply()",
		"69ceb2f1": "crossSupplyLimit()",
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
		"dfcbce15": "initializeCrossBridge(address,address,uint8,uint256,address,uint256)",
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
		"b8886e1a": "setCrossSupplyLimit(uint256)",
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
	Bin: "0x60a080604052346100bb57306080525f516020615fef5f395f51905f525460ff8160401c166100ac576002600160401b03196001600160401b03821601610059575b604051615f2f90816100c0823960805181610eed0152f35b6001600160401b0319166001600160401b039081175f516020615fef5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b61002061363c565b005b5f3560e01c806301ffc9a7146103c15780630b1d8d06146103bc5780631313996b146103b75780631938e0f2146103b2578063248a9ca3146103ad5780632f2ff15d146103a857806336568abe146103a3578063365d71e41461039e57806342cde4e81461039957806348a00ed8146103945780634be13f831461038f5780634d5d00561461038a5780634ee078ba146103855780634f1ef28614610380578063502cc5c01461037b57806352d1902d146103765780635b605f5c146103715780635c975abb1461036c5780635fd262de14610367578063670e62681461036257806369ceb2f11461035d578063761fe2d8146103585780637921487414610353578063814914b51461034e57806384b0196e1461034957806388d67d6d1461034457806389232a001461033f5780638ae87c5c1461033a57806391cca3db1461033557806391cf6d3e1461033057806391d148541461032b5780639f9b4f1c14610326578063a10bab0b14610321578063a217fddf1461031c578063a3246ad314610317578063aa1bd0bc14610312578063aa20ed471461030d578063ad3cb1cc14610308578063ae6893f814610303578063b2b49e2e146102fe578063b33eb36e146102f9578063b7f3358d146102f4578063b8886e1a146102ef578063b8aa837e146102ea578063bedb86fb146102e5578063bfbf6765146102e0578063cba25e4b146102db578063cd959706146102d6578063cf56118e146102d1578063d477f05f146102cc578063d547741f146102c7578063d5717fc5146102c2578063dfcbce15146102bd578063e2af6cd7146102b8578063edbbea39146102b3578063f0702e8e146102ae578063f4509643146102a95763f698da250361000e5761255a565b6124ba565b612492565b612448565b6123d0565b612302565b6122c9565b61229a565b61223c565b6121c8565b6121ae565b612144565b61206c565b611f83565b611ee8565b611e92565b611e0f565b611d7e565b611c71565b611c38565b611bf1565b611b68565b611b1c565b611aa0565b611a44565b611a1c565b611909565b6118c7565b6118aa565b611882565b61181a565b6116f8565b6115ee565b6114c6565b611379565b6112f9565b611273565b611256565b6111df565b6110d7565b6110a9565b610fca565b610edb565b610e4c565b610d0d565b610bad565b610b40565b610aec565b6109c3565b610997565b61092a565b61083a565b610801565b6107d0565b61071f565b6104d9565b610438565b346104175760203660031901126104175760043563ffffffff60e01b811680910361041757602090637965db0b60e01b8114908115610406575b506040519015158152f35b6301ffc9a760e01b1490505f6103fb565b5f80fd5b6001600160a01b031690565b6001600160a01b0381160361041757565b346104175760203660031901126104175760043561045581610427565b61045d61365f565b6001600160a01b03166104718115156125c6565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f84011215610417578235916001600160401b038311610417576020808501948460051b01011161041757565b6040366003190112610417576004356001600160401b038111610417576105049036906004016104a9565b602435916001600160401b0383116104175736602384011215610417576004830135916001600160401b0383116104175736602460e0850286010111610417576024610020940191612807565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761058157604052565b610551565b60e081019081106001600160401b0382111761058157604052565b606081019081106001600160401b0382111761058157604052565b60c081019081106001600160401b0382111761058157604052565b601f909101601f19168101906001600160401b0382119082101761058157604052565b6040519061060a610100836105d7565b565b6040519061060a6060836105d7565b6040519061060a6080836105d7565b6040519061060a60e0836105d7565b6001600160401b0381116105815760051b60200190565b60ff81160361041757565b9080601f8301121561041757813561067281610639565b9261068060405194856105d7565b81845260208085019260051b82010192831161041757602001905b8282106106a85750505090565b6020809183356106b781610650565b81520191019061069b565b9080601f830112156104175781356106d981610639565b926106e760405194856105d7565b81845260208085019260051b82010192831161041757602001905b82821061070f5750505090565b8135815260209182019101610702565b6080366003190112610417576004356001600160401b0381116104175760c06003198236030112610417576024356001600160401b0381116104175761076990369060040161065b565b906044356001600160401b038111610417576107899036906004016106c2565b60643591906001600160401b038311610417576107cc936107b16107ba9436906004016106c2565b92600401612945565b60405191829182901515815260200190565b0390f35b346104175760203660031901126104175760206107ee600435612c9d565b604051908152f35b359061060a82610427565b346104175760403660031901126104175761002060243560043561082482610427565b61083561083082612c9d565b613847565b61427f565b346104175760403660031901126104175760043560243561085a81610427565b336001600160a01b0382160361087357610020916142df565b63334bd91960e11b5f5260045ffd5b906040600319830112610417576004356001600160401b03811161041757826108ad916004016106c2565b91602435906001600160401b03821161041757806023830112156104175781600401356108d981610639565b926108e760405194856105d7565b8184526024602085019260051b82010192831161041757602401905b8282106109105750505090565b60208091833561091f81610427565b815201910190610903565b346104175761093836610882565b906109468151835114612cbb565b5f5b8251811015610020578061098661096160019385612cd1565b51838060a01b036109728488612cd1565b51169061098161083082612c9d565b6142df565b5001610948565b5f91031261041757565b34610417575f36600319011261041757602060ff5f516020615c7a5f395f51905f525416604051908152f35b34610417576060366003190112610417576024356004356044356109e681610427565b6109ee6136d9565b6109f66138b1565b815f526007602052610a1483610a0f8160405f20614e42565b612ce5565b80610a1f8484614a4a565b916001600160a01b031615610ad8575b8151905f516020615c9a5f395f51905f526020840191610a91835195610a87610a756040830197885160018060a01b03169986608086019b8c519260a088015194613bbf565b610a7e81611ce3565b6001811461433e565b519351945161041b565b94516040516001600160a01b0390961695918291610ab29142919084612c7f565b0390a45f516020615c3a5f395f51905f525f80a35f5f516020615dfa5f395f51905f525d005b5060608101516001600160a01b0316610a2f565b34610417575f366003190112610417575f546040516001600160a01b039091168152602090f35b9181601f84011215610417578235916001600160401b038311610417576020838186019501011161041757565b6101c036600319011261041757602435600435610b5c82610427565b604435610b6881610427565b6064359060a43560843560c4356001600160401b03811161041757610b91903690600401610b13565b94909360e03660e3190112610417576107cc976107ba97612cff565b3461041757604036600319011261041757610c8c602435600435610bcf61388a565b610bd76138b1565b805f526007602052610bf082610a0f8160405f20614e42565b610c876040610c17610c1285610c0586612b71565b905f5260205260405f2090565b6130b3565b610c74610c3782516080610c2d8683015161041b565b9101519087613aaf565b50610c4181611ce3565b610c4a81611ce3565b83516020810182905290600190610c6e83604081015b03601f1981018552846105d7565b146130ec565b01518015908115610c94575b4291613118565b614361565b6100206138e6565b4281109150610c80565b6001600160401b03811161058157601f01601f191660200190565b929192610cc582610c9e565b91610cd360405193846105d7565b829481845281830111610417578281602093845f960137010152565b9080601f8301121561041757816020610d0a93359101610cb9565b90565b604036600319011261041757600435610d2581610427565b6024356001600160401b03811161041757610d44903690600401610cef565b903061aaaa625c0eb760891b01148015610e21575b610e1257610d6561365f565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610de1575b50610dae57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d1a5f395f51905f528303610dcd5761002092506153d2565b632a87526960e21b5f52600483905260245ffd5b610e0491945060203d602011610e0b575b610dfc81836105d7565b81019061476d565b925f610d8d565b503d610df2565b63703e46dd60e11b5f5260045ffd5b505f516020615d1a5f395f51905f52546001600160a01b031661aaaa625c0eb760891b011415610d59565b3461041757606036600319011261041757602435600435604435610e6e6136d9565b815f526007602052610e8783610a0f8160405f20614e42565b4201804211610ed65760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612f1f565b34610417575f366003190112610417577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e125760206040515f516020615d1a5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610faa5750505090565b909192602060e082610fbf6001948851610f32565b019401929101610f9d565b3461041757602036600319011261041757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b818110611090575050611019925003836105d7565b611023825161316c565b915f5b81518110156110825760019061106661106161104186612b7f565b61105b61104e8588612cd1565b516001600160a01b031690565b906131bb565b6131d0565b6110708287612cd1565b5261107b8186612cd1565b5001611026565b604051806107cc8682610f87565b8454835260019485019487945060209093019201611004565b34610417575f36600319011261041757602060ff5f516020615dba5f395f51905f5254166040519015158152f35b60e0366003190112610417576024356004356110f282610427565b604435916110ff83610427565b6064359260c435916084359160a435916001600160401b038511610417576111b0966111656111356111a6973690600401610b13565b95909661114061388a565b6001600160a01b038516948490611157878d61440e565b61115f6138b1565b8b6144c4565b939092604051986111758a610565565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610cb9565b60e082015261466f565b5f5f516020615dfa5f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b34610417576080366003190112610417576024356004356111ff82610427565b604435906001600160401b0382116104175736602383011215610417576107cc9261123761124a933690602481600401359101610cb9565b906064359261124584610650565b613258565b604051918291826111cc565b34610417575f366003190112610417576020606454604051908152f35b3461041757604036600319011261041757602060ff6112a960243560043561129a82610427565b5f526009845260405f206131bb565b54166040519015158152f35b90602080835192838152019201905f5b8181106112d25750505090565b82518452602093840193909201916001016112c5565b906020610d0a9281815201906112b5565b34610417576020366003190112610417576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b818110611353576107cc85611347818703826105d7565b604051918291826112e8565b8254845260209093019260019283019201611330565b60e08101929161060a9190610f32565b34610417576040366003190112610417576107cc6113b860243560043561139f82610427565b6113a7613136565b505f52600660205260405f206131bb565b6004604051916113c783610586565b80546001600160a01b03168352600181015461141c90611413906113f66113ed8261041b565b60208801612f68565b61140a60a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611369565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161149c9061148e610d0a97959693600f60f81b865260e0602087015260e0860190611443565b908482036040860152611443565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526112b5565b34610417575f366003190112610417575f516020615cda5f395f51905f5254158061155a575b1561151d576114f961477c565b611501614836565b906107cc61150d613321565b6040519384933091469186611467565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615e9a5f395f51905f5254156114ec565b9080601f8301121561041757813561158781610639565b9261159560405194856105d7565b81845260208085019260051b820101918383116104175760208201905b8382106115c157505050505090565b81356001600160401b038111610417576020916115e3878480948801016106c2565b8152019101906115b2565b6080366003190112610417576004356001600160401b038111610417576116199036906004016104a9565b90602435906001600160401b038211610417573660238301121561041757816004013561164581610639565b9261165360405194856105d7565b8184526024602085019260051b820101903682116104175760248101925b8284106116c95750506044359150506001600160401b0381116104175761169c903690600401611570565b606435926001600160401b038411610417576107cc946116c36107ba953690600401611570565b9361333c565b83356001600160401b038111610417576020916116ed83926024369187010161065b565b815201930192611671565b346104175760603660031901126104175760043561171581610427565b6024359061172282610427565b60443561172e81610650565b5f516020615e7a5f395f51905f52549260ff604085901c1615611750565b1590565b936001600160401b031680159081611812575b6001149081611808575b1590816117ff575b506117f05761179092846117876133d2565b6117e3576148d0565b61179657005b6117b45f516020615e7a5f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6117eb6133f3565b6148d0565b63f92ee8a960e01b5f5260045ffd5b9050155f611775565b303b15915061176d565b859150611763565b346104175760403660031901126104175760243560043561183961365f565b6118416138b1565b61184b8282614a4a565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615dfa5f395f51905f525d005b34610417575f366003190112610417576033546040516001600160a01b039091168152602090f35b34610417575f366003190112610417576020603454604051908152f35b3461041757604036600319011261041757602060ff6112a96024356004356118ee82610427565b5f525f516020615d9a5f395f51905f52845260405f206131bb565b34610417576040366003190112610417576004356001600160401b038111610417576119399036906004016106c2565b6024356001600160401b038111610417576119589036906004016106c2565b90611966815183511461261d565b5f5b82518110156100205780611a0e61198160019385612cd1565b5161198c8387612cd1565b519061199661388a565b61199e6138b1565b805f5260076020526119b782610a0f8160405f20614e42565b610c8760406119cc610c1285610c0586612b71565b610c746119e282516080610c2d8683015161041b565b506119ec81611ce3565b6119f581611ce3565b835160208101829052908a90610c6e8360408101610c60565b611a166138e6565b01611968565b34610417575f366003190112610417576035546040516001600160a01b039091168152602090f35b34610417575f3660031901126104175760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611a815750505090565b82516001600160a01b0316845260209384019390920191600101611a74565b34610417576020366003190112610417576004355f525f516020615d3a5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611b06576107cc85611afa818703826105d7565b60405191829182611a5e565b8254845260209093019260019283019201611ae3565b34610417576020366003190112610417577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611b5b61365f565b80600155604051908152a1005b3461041757604036600319011261041757602435600435611b876136d9565b611b8f6136d9565b611b976138b1565b805f526007602052611bb082610a0f8160405f20614e42565b611bba8282614361565b5f516020615c3a5f395f51905f525f80a35f5f516020615dfa5f395f51905f525d005b60405190611bec6020836105d7565b5f8252565b34610417575f366003190112610417576107cc604051611c126040826105d7565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611443565b34610417576020366003190112610417576004355f526004602052600160405f20015460018101809111610ed657602090604051908152f35b3461041757611c7f36610882565b90611c8d8151835114612cbb565b5f5b82518110156100205780611cc8611ca860019385612cd1565b51838060a01b03611cb98488612cd1565b51169061083561083082612c9d565b5001611c8f565b634e487b7160e01b5f52602160045260245ffd5b600a1115611ced57565b611ccf565b90600a821015611ced5752565b6020815260606040611d6460a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611443565b93611d76602082015183860190611cf2565b015191015290565b3461041757604036600319011261041757600435602435905f60408051611da4816105a1565b611dac613414565b815282602082015201525f52600860205260405f20905f526020526107cc60405f20600760405191611ddd836105a1565b611de681612faf565b8352611dfc60ff600683015416602085016130a7565b0154604082015260405191829182611cff565b34610417576020366003190112610417577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611e5181610650565b611e5961365f565b16611e65811515613444565b8060ff195f516020615c7a5f395f51905f525416175f516020615c7a5f395f51905f5255604051908152a1005b34610417576020366003190112610417577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435611ed161365f565b80606455604051908152a1005b8015150361041757565b3461041757604036600319011261041757600435602435611f0881611ede565b611f10613753565b611f25825f52600360205260405f2054151590565b15611f705760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611f6581600360405f200161345a565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b3461041757602036600319011261041757600435611fa081611ede565b611fa8613753565b1561200657611fb561388a565b600160ff195f516020615dba5f395f51905f525416175f516020615dba5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615dba5f395f51905f525460ff81161561205d5760ff19165f516020615dba5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346104175760803660031901126104175760243560043561208c82610427565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356120bb81611ede565b6064356120c781611ede565b6120cf613753565b845f5260056020526121338161212e855f20986120fe816120f960018060a01b038216809d614e42565b61346b565b885f52600660205261211e866001612118848b5f206131bb565b01613493565b885f526009602052865f206131bb565b61345a565b8251911515825215156020820152a3005b346104175760203660031901126104175760043561216181610427565b61216961365f565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b34610417575f3660031901126104175760206107ee6134b0565b34610417575f36600319011261041757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b818110612226576107cc85611347818703826105d7565b825484526020909301926001928301920161220f565b346104175760203660031901126104175760043561225981610427565b61226161365f565b6001600160a01b03166122758115156125c6565b603380546001600160a01b031916821790555f516020615d5a5f395f51905f525f80a2005b34610417576040366003190112610417576100206024356004356122bd82610427565b61098161083082612c9d565b34610417576020366003190112610417576004355f526004602052600260405f20015460018101809111610ed657602090604051908152f35b346104175760c03660031901126104175760043561231f81610427565b6024359061232c82610427565b60443561233881610650565b60843560643561234782610427565b60a435925f516020615e7a5f395f51905f52549561236c61174c8860ff9060401c1690565b966001600160401b0316801590816123c8575b60011490816123be575b1590816123b5575b506117f05761179095876123a36133d2565b156134d0576123b06133f3565b6134d0565b9050155f612391565b303b159150612389565b88915061237f565b34610417576020366003190112610417576004356123ed81610427565b6123f561365f565b6001600160a01b03166124098115156135fa565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346104175760803660031901126104175761002060043560243561246b81611ede565b6044359061247882610427565b6064359261248584610427565b61248d6137cd565b614cc2565b34610417575f366003190112610417576032546040516001600160a01b039091168152602090f35b34610417576040366003190112610417576024356004356124da82610427565b6124e261365f565b805f5260056020525f600461251d604083209461250c816120f960018060a01b03821680996156c5565b8484526006602052604084206131bb565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b34610417575f36600319011261041757612572615a5d565b61257a615ab4565b6040519060208201925f516020615eba5f395f51905f528452604083015260608201524660808201523060a082015260a081526125b860c0826105d7565b519020604051908152602090f35b156125cd57565b638219abe360e01b5f5260045ffd5b603380546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b1561262457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156126695760051b8101359060fe1981360301821215610417570190565b612633565b35610d0a81610427565b903590601e198136030182121561041757018035906001600160401b0382116104175760200191813603831361041757565b91908110156126695760e0020190565b908160209103126104175751610d0a81611ede565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361274097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c08601916126cf565b94803561274c81610427565b6001600160a01b031660e0850152602081013561276881610427565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561279d81610650565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d156127eb573d906127d282610c9e565b916127e060405193846105d7565b82523d5f602084013e565b606090565b604090610d0a939281528160208201520190611443565b909193929361281785841461261d565b5f945b83861061282957505050509050565b612834868585612647565b35602061284c816128468a8989612647565b0161266e565b61285c60606128468b8a8a612647565b926128d28a86888a8c6128b66080612875878486612647565b0135956128ae6128a48260a061288c82888a612647565b01359560c061289c83838b612647565b013597612647565b60e0810190612678565b9690956126aa565b946040519a8b998a996326ae802b60e11b8b5260048b016126ef565b03815f305af19081612919575b5061290e57856128ed6127c1565b60405163f495148b60e01b815291829161290a91600484016127f0565b0390fd5b60019095019461281a565b6129399060203d811161293e575b61293181836105d7565b8101906126ba565b6128df565b503d612927565b90612a2f93929161295461388a565b61295c6138b1565b803592612971845f52600560205260405f2090565b916129a6612994604083019461298e6129898761266e565b61041b565b906138f8565b6129a06129898661266e565b90612b8d565b6129b1343415612bb5565b612a456129d1865f526004602052600260405f2001600181540180915590565b966129e4602084013589819a8214612bd3565b6129f06129898661266e565b93606084019688612a008961266e565b96612a3d8c60808901359e8f60a08b019b612a1b8d8d612678565b929091604051978896602088019a8b612bf1565b03601f1981018352826105d7565b51902061393b565b612a5887612a528561266e565b87613b12565b919092600184612a6781611ce3565b14612b34575b50612a7783611ce3565b60018303612ad15750505090612aa3612a9d5f516020615c9a5f395f51905f529361266e565b9161266e565b6040516001600160a01b03909216958291612ac091429184612c7f565b0390a45b612acc6138e6565b600190565b612b0e8394612b09612b2c947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612b1495614155565b61266e565b9261266e565b9260405193849360018060a01b031698429185612c55565b0390a4612ac4565b612b6a91935088612b448661266e565b91612b62612b5b612b548a61266e565b9288612678565b3691610cb9565b928a8a613bbf565b915f612a6d565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612b955750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612bbd5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612bdc575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610d0a97959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c082015201916126cf565b61060a93606092969593608083019760018060a01b03168352602083015260408201520190611cf2565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615d9a5f395f51905f52602052600160405f20015490565b15612cc257565b63031206d560e51b5f5260045ffd5b80518210156126695760209160051b010190565b15612ced5750565b6373a1390160e11b5f5260045260245ffd5b959394612d81919597939297612d1361388a565b612d586001600160a01b038816612d2a818b61440e565b612d326138b1565b612d4261298961298960e461266e565b811490612d5261298960e461266e565b91612edc565b612d79612d6961298961010461266e565b6001600160a01b038b1614612f09565b8387896144c4565b9161012435612db481612d9d86612d988787612f33565b612f33565b811015612dae87612d988888612f33565b90612f40565b612dc261298960325461041b565b90612dce61010461266e565b906101443592612ddf610164612f5e565b92610184356101a43591833b1561041757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612ed757612e9e6111a698612ea793612ac49c612ebd575b50612e95612e8061010461266e565b91612e896105fa565b9c8d5260208d01612f68565b60408b01612f68565b60608901612f68565b608087015260a086015260c08501523691610cb9565b80612ecb5f612ed1936105d7565b8061098d565b5f612e71565b6127b6565b15612ee5575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612f1057565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ed657565b15612f49575050565b63943892eb60e01b5f5260045260245260445ffd5b35610d0a81610650565b6001600160a01b039091169052565b90600182811c92168015612fa5575b6020831014612f9157565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612f86565b90604051612fbc816105bc565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f9161301482612f77565b80855291600181169081156130805750600114613041575b505060a0929161303d9103846105d7565b0152565b5f908152602081209092505b81831061306457505081016020018161303d61302c565b602091935080600191548385890101520191019091849261304d565b60ff191660208681019190915292151560051b8501909201925083915061303d905061302c565b600a821015611ced5752565b906040516130c0816105a1565b6040600782946130cf81612faf565b84526130e560ff600683015416602086016130a7565b0154910152565b156130f45750565b60405162461bcd60e51b81526020600482015290819061290a906024830190611443565b15613121575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061314382610586565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061317682610639565b61318360405191826105d7565b8281528092613194601f1991610639565b01905f5b8281106131a457505050565b6020906131af613136565b82828501015201613198565b9060018060a01b03165f5260205260405f2090565b906040516131dd81610586565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132289061321f9061140a565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126104175751610d0a81610427565b5f5490949392906001600160a01b0381161561331257604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906132c1906084870190611443565b931691015203925af18015612ed75761060a925f916132e3575b508094613610565b613305915060203d60201161330b575b6132fd81836105d7565b810190613243565b5f6132db565b503d6132f3565b6315aeca0d60e11b5f5260045ffd5b604051906133306020836105d7565b5f808352366020840137565b91929493909384845114806133c8575b806133be575b61335b9061261d565b5f5b858110156133b2578060051b8401359060be1985360301821215610417576133ab60019261338b8389612cd1565b51613396848c612cd1565b51906133a28589612cd1565b51928901612945565b500161335d565b50945050505050600190565b5081518514613352565b508486511461334c565b5f516020615e7a5f395f51905f5280546001600160401b0319166001179055565b5f516020615e7a5f395f51905f52805460ff60401b1916600160401b179055565b60405190613421826105bc565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b1561344b57565b63f0f15b9160e01b5f5260045ffd5b9060ff801983541691151516179055565b156134735750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6065545f52600660205260405f2060015f52602052600460405f20015490565b91949293909484156135eb576001600160a01b038216156135dc576134f3615474565b6135076001600160a01b03841615156125c6565b6001600160a01b0386169061351d8215156125c6565b60ff8116156135cd5761060a966135846135919261357f6135b197613540615474565b613548615474565b613550615474565b60ff195f516020615dba5f395f51905f5254165f516020615dba5f395f51905f525561357a615474565b61549f565b6154ab565b61358c615664565b6125dc565b5f516020615d5a5f395f51905f525f80a26135ab43603455565b83614bb8565b806135bd575b50606555565b6135c79082614dbe565b5f6135b7565b6338854f1160e21b5f5260045ffd5b6324cd8cef60e01b5f5260045ffd5b63110d132360e21b5f5260045ffd5b1561360157565b6302bfb33d60e51b5f5260045ffd5b905f61060a939261248d6137cd565b916136389183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b0316330361365057565b63c82cebcb60e01b5f5260045ffd5b5f516020615e1a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff6136ac337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c6131bb565b5416156136b557565b63e2517d3f60e01b5f52336004525f516020615e1a5f395f51905f5260245260445ffd5b5f516020615e5a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff613726337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b1436131bb565b54161561372f57565b63e2517d3f60e01b5f52336004525f516020615e5a5f395f51905f5260245260445ffd5b5f516020615d7a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff6137a0337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb4566131bb565b5416156137a957565b63e2517d3f60e01b5f52336004525f516020615d7a5f395f51905f5260245260445ffd5b5f516020615dda5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff61381a337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f6131bb565b54161561382357565b63e2517d3f60e01b5f52336004525f516020615dda5f395f51905f5260245260445ffd5b805f525f516020615d9a5f395f51905f5260205260ff61386a3360405f206131bb565b5416156138745750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615dba5f395f51905f5254166138a257565b63d93c066560e01b5f5260045ffd5b5f516020615dfa5f395f51905f525c6138d75760015f516020615dfa5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615dfa5f395f51905f525d565b610d0a916001600160a01b031690614e42565b1561391257565b63b3f07a3960e01b5f5260045ffd5b156139295750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613aa5575b6139549061390b565b61397561396f5f516020615c7a5f395f51905f525460ff1690565b60ff1690565b956139838488811015613921565b5f945f935f5b8681106139a4575050505050505061060a9192811015613921565b613a016139d0836139b3615768565b6042916040519161190160f01b8352600283015260228201522090565b6139e46139dd8489612cd1565b5160ff1690565b6139ee8487612cd1565b51906139fa8589612cd1565b5192614e55565b6001600160a01b038181169088161080613a32575b613a24575b50600101613989565b600198890198909650613a1b565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615d9a5f395f51905f52602052613aa0613a99827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa56131bb565b6131bb565b5460ff1690565b613a16565b508551831461394b565b9091906001600160a01b03831660011480613b07575b613adb575b90613ad792600192614e85565b9091565b90613ae46134b0565b818101809111610ed65760645410613afc5790613aca565b505050600990600190565b506065548114613ac5565b9091906001600160a01b03831660011480613b5a575b613b39575b90613ad7925f92614e85565b90613b426134b0565b818101809111610ed65760645410613afc5790613b2d565b506065548114613b28565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610d0a92910190611443565b6001600160a01b039091168152602081019190915260400190565b938093959295613bce86612b7f565b96613bf16001613be6818060a01b038416809b6131bb565b015460a01c60ff1690565b93613c0061298960355461041b565b9560188151101580613e9d575b80613e93575b613c4e575b5050613c249450614f71565b92613c2e84611ce3565b60018414613c3d575b50505090565b613c4692614e0b565b5f8080613c37565b602081015160601c5f808092604051613c7b81612a2f602082019462483e5560e91b8652602483016111cc565b51908b5afa613c886127c1565b9080613e87575b613e7b575b5015613c1857909192935060018914159182613e4a575b5f905f808a8c89613ce48a612a2f869a8c8814613e445787965b6040519586948c6020870199631eeed51360e01b8b5260248801613b65565b51918c5af1613cf16127c1565b9080613e38575b613e23575b50897f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613d348682901515815260200190565b0390a3613dc7575090859291613d4c575b8480613c18565b8315613d6757613d6283613c24963090846150d9565b613d45565b91602090613d8a956040519687928392632770a7eb60e21b845260048401613ba4565b03815f8b5af1918215612ed757613c24948693613da8575b50613d45565b613dc09060203d60201161293e5761293181836105d7565b505f613da2565b96613dd793969298919450614e0b565b83613de7575b5050505050600190565b15613dff57613df5936150d9565b5f80808080613ddd565b5090613e1d9250613e0e611bdd565b916001600160a01b0316615070565b50613df5565b9150915060406020820151910151915f613cfd565b50604081511015613cf8565b86613cc5565b613e5686898987614f71565b613e5f81611ce3565b60018103613e6d5750613cab565b995050505050505050505090565b6020915001515f613c94565b50602081511015613c8f565b50863b1515613c13565b506001600160a01b0387161515613c0d565b15613eb75750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126104175760405190613ee2826105bc565b819380358352602081013560208401526040810135613f0081610427565b6040840152613f11606082016107f6565b60608401526080818101359084015260a0810135916001600160401b0383116104175760a092613f419201610cef565b910152565b818110613f51575050565b5f8155600101613f46565b90601f8211613f69575050565b61060a915f516020615c5a5f395f51905f525f5260205f20906020601f840160051c83019310613fa1575b601f0160051c0190613f46565b9091508190613f94565b9190601f8111613fba57505050565b61060a925f5260205f20906020601f840160051c83019310613fa157601f0160051c0190613f46565b8160011b915f199060031b1c19161790565b90600a811015611ced5760ff80198354169116179055565b81518051825560208101516001830155604081015191929161403b906001600160a01b0316600285016125fe565b6060810151614056906001600160a01b0316600385016125fe565b6080810151600484015560a00151805160058401916001600160401b0382116105815761408d826140878554612f77565b85613fab565b602090601f83116001146140e5578260079593604095936140b6935f926140da575b5050613fe3565b90555b6140d360208201516140ca81611ce3565b60068601613ff5565b0151910155565b015190505f806140af565b90601f198316916140f9855f5260205f2090565b925f5b81811061413d5750926001928592600798966040989610614125575b505050811b0190556140b9565b01515f1960f88460031b161c191690555f8080614118565b929360206001819287860151815501950193016140fc565b9160806141ef6141e06002936141db876141d661418f61363899833595865f52600760205261419460405f2060208701359485809261597b565b613eaf565b156141fc576141c86141a860015442612f33565b915b6141bd6141b561060c565b963690613ec9565b8652602086016130a7565b6040840152610c0585612b71565b61400d565b612b7f565b61105b6129896040880161266e565b9301359201918254612f33565b6141c85f916141aa565b90614211825f615122565b918261421a5750565b5f80525f516020615d3a5f395f51905f526020526001600160a01b0316614261817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d61597b565b156142695750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161428c8382615122565b9283614296575050565b815f525f516020615d3a5f395f51905f526020526142c160405f209160018060a01b0316809261597b565b156142ca575050565b63d180cb3160e01b5f5260045260245260445ffd5b9190916142ec83826151bf565b92836142f6575050565b815f525f516020615d3a5f395f51905f5260205261432160405f209160018060a01b031680926156c5565b1561432a575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143465750565b63290cd55f60e01b5f52600a811015611ced5760045260245ffd5b9061436b91614a4a565b60018060a01b036060820151168151905f516020615c9a5f395f51905f5260208401918251946143bb610a756040830196875160018060a01b03169885608086019a8b519260a088015194613bbf565b519251935194516040516001600160a01b03909616959182916143e19142919084612c7f565b0390a4565b156143ee5750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f526005602052614432816120f960405f2060018060a01b03831690614e42565b825f52600460205260ff600360405f2001541661446a5760ff600161445e83613a9461060a9697612b7f565b015460a81c16156143e6565b82636fc24b7960e11b5f5260045260245ffd5b1561448457565b6330d45fb160e01b5f5260045ffd5b90816060910312610417578051916040602083015192015190565b156144b557565b6358e8878560e01b5f5260045ffd5b8260609161455d979596936144ee6110616144de84612b7f565b6001600160a01b038416906131bb565b6144fe61174c6040830151151590565b6145ff575b5061451261298960325461041b565b916145276001600160a01b038416151561447d565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612ed7575f955f905f936145c1575b5090829161060a949396819885151595866145b6575b5050846145ab575b5050826145a0575b50506144ae565b101590505f80614599565b101592505f80614591565b101594505f80614589565b90506145ec91965061060a93925060603d6060116145f8575b6145e481836105d7565b810190614493565b91969293919291614573565b503d6145da565b60c06146119101518480821015612f40565b5f614503565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161466a9190860190611443565b930152565b61468d81515f526004602052600160405f2001600181540180915590565b6146978251612b7f565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086146dc61298960016146d5602088019561105b612989885161041b565b015461041b565b938051906143e16146ed855161041b565b9160408101906146fd825161041b565b906147266080820196875160a084019485519861472060c087019a8b5190612f33565b93615291565b61473c6147358251995161041b565b935161041b565b9460e061474c606084015161041b565b9751935191519201519260405197889760018060a01b03169c429689614617565b90816020910312610417575190565b604051905f825f516020615c5a5f395f51905f52549161479b83612f77565b808352926001811690811561481757506001146147bf575b61060a925003836105d7565b505f516020615c5a5f395f51905f525f90815290915f516020615cfa5f395f51905f525b8183106147fb57505090602061060a928201016147b3565b60209193508060019154838589010152019101909184926147e3565b6020925061060a94915060ff191682840152151560051b8201016147b3565b604051905f825f516020615cba5f395f51905f52549161485583612f77565b808352926001811690811561481757506001146148785761060a925003836105d7565b505f516020615cba5f395f51905f525f90815290915f516020615eda5f395f51905f525b8183106148b457505090602061060a928201016147b3565b602091935080600191548385890101520191019091849261489c565b91906148da615474565b6148ee6001600160a01b03841615156125c6565b6001600160a01b038116926149048415156125c6565b60ff8316156135cd576149919261497761497d92614920615474565b614928615474565b614930615474565b60ff195f516020615dba5f395f51905f5254165f516020615dba5f395f51905f525561495a615474565b614962615474565b61496a615474565b614972615474565b614206565b506154ab565b614985615474565b620151806001556125dc565b5f516020615d5a5f395f51905f525f80a261060a43603455565b91908203918211610ed657565b60075f91828155826001820155826002820155826003820155826004820155600581016149e58154612f77565b90816149f8575b50508260068201550155565b601f8211600114614a0f57849055505b5f806149ec565b614a35614a45926001601f614a27855f5260205f2090565b920160051c82019101613f46565b5f81815260208120918190559055565b614a08565b9190614a54613414565b50825f526007602052614a6a8160405f206156c5565b15614ad457614acf61060a91845f52600860205260405f20815f52602052610c05614a9760405f20612faf565b95614ab4614aa482612b7f565b61105b61298960408b015161041b565b614ac8600260808a015192019182546149ab565b9055612b71565b6149b8565b637f11bea960e01b5f5260045260245ffd5b6003606061060a9380518455602081015160018501556040810151600285015501511515910161345a565b15614b195750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c0600491614b5260018060a01b03825116856125fe565b6020810151614b9d906001860190614b73906001600160a01b0316826125fe565b6040830151815460ff60a01b191690151560a01b60ff60a01b161781556060830151151590613493565b6080810151600285015560a081015160038501550151910155565b600190614bc68115156135fa565b614bcf826135fa565b614c67828060a01b03841693614be68515156135fa565b614bef83615912565b614c86575b614c1884614c1381614c0e875f52600560205260405f2090565b6151ac565b614b11565b614c37614c2361062a565b91614c2e8684612f68565b60208301612f68565b5f60408201525f60608201525f60808201525f60a08201525f60c0820152614c6284613a9485612b7f565b614b39565b6040515f81525f516020615e3a5f395f51905f529080602081016143e1565b614cbd614c9161061b565b8481525f60208201525f60408201525f6060820152614cb8855f52600460205260405f2090565b614ae6565b614bf4565b906143e15f516020615e3a5f395f51905f5291949394614ce38415156135fa565b6001600160a01b038616946107ba90614cfd8715156135fa565b6001600160a01b03811697614c6290614d178a15156135fa565b614d2088615912565b614d87575b614d3f81614c1381614c0e8c5f52600560205260405f2090565b614d5e614d4a61062a565b93614d558386612f68565b60208501612f68565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a9488612b7f565b614db9614d9261061b565b8981525f60208201525f60408201525f6060820152614cb88a5f52600460205260405f2090565b614d25565b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614df9576003018054918203918211610ed65755565b6004018054918201809211610ed65755565b90614e20915f52600660205260405f206131bb565b600181015460a01c60ff1615614df9576003018054918203918211610ed65755565b6001915f520160205260405f2054151590565b91610d0a9391614e64936157bc565b9092919261583e565b908160209103126104175751600a8110156104175790565b905f93614ed2613a99614ec2614e9f61298960325461041b565b95614eb46001600160a01b038816151561447d565b5f52600960205260405f2090565b6001600160a01b038516906131bb565b614f6557614f5d5791602091614f0095935f604051809881958294633f4bdec560e01b845260048401613ba4565b03925af1928315612ed7575f93614f2c575b50600183614f1f81611ce3565b03614f2657565b60019150565b614f4f91935060203d602011614f56575b614f4781836105d7565b810190614e6d565b915f614f12565b503d614f3d565b505050600191565b50505050506002905f90565b6001600160a01b031692919060018414615000578115614ff757614fc09215614fcb5760405163a9059cbb60e01b602082015291614fb8918391612a2f9160248401613ba4565b600592615021565b15610d0a5750600190565b6040516340c10f1960e01b602082015291614fef918391612a2f9160248401613ba4565b600692615021565b50505050600190565b90615013935061174c9250613e0e611bdd565b61501c57600190565b600590565b81516001600160a01b03909116915f9182916020018285620186a0f16150456127c1565b901561506a57805161506157503b1561505d57600190565b5f90565b60209150015190565b50505f90565b8147106150ca5782516001600160a01b03909116925f9182916020018486620186a0f19061509c6127c1565b91156150c357156150af575b5050600190565b805161506157503b1561505d575f806150a8565b5050505f90565b632b60b36f60e21b5f5260045ffd5b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815261060a9161511d6084836105d7565b6158ba565b805f525f516020615d9a5f395f51905f5260205260ff6151458360405f206131bb565b541661506a57805f525f516020615d9a5f395f51905f5260205261516c8260405f206131bb565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610d0a916001600160a01b03169061597b565b805f525f516020615d9a5f395f51905f5260205260ff6151e28360405f206131bb565b54161561506a57805f525f516020615d9a5f395f51905f5260205261520a8260405f206131bb565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b1561525157505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561528257565b63559d141b60e11b5f5260045ffd5b90936001600160a01b03851692600184036152ff575061060a94506152c76152b98286612f33565b34143490612dae8488612f33565b806152d3575b506159f7565b6152f46152f9916152e560335461041b565b906152ee611bdd565b91615070565b61527b565b5f6152cd565b9061530b343415612bb5565b6153206153188287612f33565b3084896150d9565b806153b4575b5061533c61174c6001613be686613a9487612b7f565b61534c575b5061060a93506159f7565b604051632770a7eb60e21b81526020818061536b883060048401613ba4565b03815f885af1918215612ed75761060a9661538f9387935f91615395575b50615247565b5f615341565b6153ae915060203d60201161293e5761293181836105d7565b5f615389565b6153cc906153c661298960335461041b565b876159c2565b5f615326565b90813b15615453575f516020615d1a5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a280511561543b5761543891615a40565b50565b50503461544457565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b60ff5f516020615e7a5f395f51905f525460401c161561549057565b631afcd79f60e31b5f5260045ffd5b61543890614962615474565b906154b4615474565b6154c260ff83161515613444565b6040918251926154d281856105d7565b60098452682b30b634b230ba37b960b91b60208501526154f4815191826105d7565b60058152640312e302e360dc1b602082015261550e615474565b615516615474565b83516001600160401b03811161058157615546816155415f516020615c5a5f395f51905f5254612f77565b613f5c565b6020601f82116001146155d55781615583939261556f9261060a97985f926140da575050613fe3565b5f516020615c5a5f395f51905f5255615b44565b6155985f5f516020615cda5f395f51905f5255565b6155ad5f5f516020615e9a5f395f51905f5255565b60ff1660ff195f516020615c7a5f395f51905f525416175f516020615c7a5f395f51905f5255565b5f516020615c5a5f395f51905f525f52601f198216955f516020615cfa5f395f51905f52965f5b81811061564c5750966001928492615583969561060a999a10615634575b505050811b015f516020615c5a5f395f51905f5255615b44565b01515f1960f88460031b161c191690555f808061561a565b838301518955600190980197602093840193016155fc565b61566c615474565b62015180600155565b8054821015612669575f5260205f2001905f90565b805480156156b1575f1901906156a08282615675565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615760575f198401848111610ed65783545f19810194908511610ed6575f95858361571397610c059503615719575b50505061568a565b55600190565b6157496157439161573a6157306157579588615675565b90549060031b1c90565b92839187615675565b9061361f565b85905f5260205260405f2090565b555f808061570b565b505050505f90565b615770615a5d565b615778615ab4565b6040519060208201925f516020615eba5f395f51905f528452604083015260608201524660808201523060a082015260a081526157b660c0826105d7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615829579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612ed7575f516001600160a01b0381161561581f57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611ced57565b61584781615834565b80615850575050565b61585981615834565b600181036158705763f645eedf60e01b5f5260045ffd5b61587981615834565b60028103615894575063fce698f760e01b5f5260045260245ffd5b806158a0600392615834565b146158a85750565b6335e2f38360e21b5f5260045260245ffd5b905f602091828151910182855af1156127b6575f513d61590957506001600160a01b0381163b155b6158e95750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156158e2565b5f8181526003602052604090205461597657600254600160401b8110156105815761595f6159498260018594016002556002615675565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6159858282614e42565b61506a57805490600160401b82101561058157826159ad615949846001809601855584615675565b90558054925f520160205260405f2055600190565b61511d61060a93926159e960405194859263a9059cbb60e01b602085015260248401613ba4565b03601f1981018452836105d7565b90615a0c915f52600660205260405f206131bb565b600181015460a01c60ff1615615a2e576003018054918201809211610ed65755565b6004018054918203918211610ed65755565b5f80610d0a93602081519101845af4615a576127c1565b91615ae6565b615a6561477c565b8051908115615a75576020012090565b50505f516020615cda5f395f51905f52548015615a8f5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615abc614836565b8051908115615acc576020012090565b50505f516020615e9a5f395f51905f52548015615a8f5790565b90615b0a5750805115615afb57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b3b575b615b1b575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b13565b80519091906001600160401b03811161058157615b8581615b725f516020615cba5f395f51905f5254612f77565b5f516020615cba5f395f51905f52613fab565b602092601f8211600114615bb857615ba7929382915f926140da575050613fe3565b5f516020615cba5f395f51905f5255565b5f516020615cba5f395f51905f525f52601f198216935f516020615eda5f395f51905f52915f5b868110615c215750836001959610615c09575b505050811b015f516020615cba5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615bf2565b91926020600181928685015181550194019201615bdf56feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220962babe63ee1eefa7adc382e96a57fc4274abfa7e2220a1e211af8d80836259664736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// CrossBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeMetaData.ABI instead.
var CrossBridgeABI = CrossBridgeMetaData.ABI

// CrossBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossBridgeBinRuntime = "60806040526004361015610022575b3615610018575f80fd5b61002061363c565b005b5f3560e01c806301ffc9a7146103c15780630b1d8d06146103bc5780631313996b146103b75780631938e0f2146103b2578063248a9ca3146103ad5780632f2ff15d146103a857806336568abe146103a3578063365d71e41461039e57806342cde4e81461039957806348a00ed8146103945780634be13f831461038f5780634d5d00561461038a5780634ee078ba146103855780634f1ef28614610380578063502cc5c01461037b57806352d1902d146103765780635b605f5c146103715780635c975abb1461036c5780635fd262de14610367578063670e62681461036257806369ceb2f11461035d578063761fe2d8146103585780637921487414610353578063814914b51461034e57806384b0196e1461034957806388d67d6d1461034457806389232a001461033f5780638ae87c5c1461033a57806391cca3db1461033557806391cf6d3e1461033057806391d148541461032b5780639f9b4f1c14610326578063a10bab0b14610321578063a217fddf1461031c578063a3246ad314610317578063aa1bd0bc14610312578063aa20ed471461030d578063ad3cb1cc14610308578063ae6893f814610303578063b2b49e2e146102fe578063b33eb36e146102f9578063b7f3358d146102f4578063b8886e1a146102ef578063b8aa837e146102ea578063bedb86fb146102e5578063bfbf6765146102e0578063cba25e4b146102db578063cd959706146102d6578063cf56118e146102d1578063d477f05f146102cc578063d547741f146102c7578063d5717fc5146102c2578063dfcbce15146102bd578063e2af6cd7146102b8578063edbbea39146102b3578063f0702e8e146102ae578063f4509643146102a95763f698da250361000e5761255a565b6124ba565b612492565b612448565b6123d0565b612302565b6122c9565b61229a565b61223c565b6121c8565b6121ae565b612144565b61206c565b611f83565b611ee8565b611e92565b611e0f565b611d7e565b611c71565b611c38565b611bf1565b611b68565b611b1c565b611aa0565b611a44565b611a1c565b611909565b6118c7565b6118aa565b611882565b61181a565b6116f8565b6115ee565b6114c6565b611379565b6112f9565b611273565b611256565b6111df565b6110d7565b6110a9565b610fca565b610edb565b610e4c565b610d0d565b610bad565b610b40565b610aec565b6109c3565b610997565b61092a565b61083a565b610801565b6107d0565b61071f565b6104d9565b610438565b346104175760203660031901126104175760043563ffffffff60e01b811680910361041757602090637965db0b60e01b8114908115610406575b506040519015158152f35b6301ffc9a760e01b1490505f6103fb565b5f80fd5b6001600160a01b031690565b6001600160a01b0381160361041757565b346104175760203660031901126104175760043561045581610427565b61045d61365f565b6001600160a01b03166104718115156125c6565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f84011215610417578235916001600160401b038311610417576020808501948460051b01011161041757565b6040366003190112610417576004356001600160401b038111610417576105049036906004016104a9565b602435916001600160401b0383116104175736602384011215610417576004830135916001600160401b0383116104175736602460e0850286010111610417576024610020940191612807565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761058157604052565b610551565b60e081019081106001600160401b0382111761058157604052565b606081019081106001600160401b0382111761058157604052565b60c081019081106001600160401b0382111761058157604052565b601f909101601f19168101906001600160401b0382119082101761058157604052565b6040519061060a610100836105d7565b565b6040519061060a6060836105d7565b6040519061060a6080836105d7565b6040519061060a60e0836105d7565b6001600160401b0381116105815760051b60200190565b60ff81160361041757565b9080601f8301121561041757813561067281610639565b9261068060405194856105d7565b81845260208085019260051b82010192831161041757602001905b8282106106a85750505090565b6020809183356106b781610650565b81520191019061069b565b9080601f830112156104175781356106d981610639565b926106e760405194856105d7565b81845260208085019260051b82010192831161041757602001905b82821061070f5750505090565b8135815260209182019101610702565b6080366003190112610417576004356001600160401b0381116104175760c06003198236030112610417576024356001600160401b0381116104175761076990369060040161065b565b906044356001600160401b038111610417576107899036906004016106c2565b60643591906001600160401b038311610417576107cc936107b16107ba9436906004016106c2565b92600401612945565b60405191829182901515815260200190565b0390f35b346104175760203660031901126104175760206107ee600435612c9d565b604051908152f35b359061060a82610427565b346104175760403660031901126104175761002060243560043561082482610427565b61083561083082612c9d565b613847565b61427f565b346104175760403660031901126104175760043560243561085a81610427565b336001600160a01b0382160361087357610020916142df565b63334bd91960e11b5f5260045ffd5b906040600319830112610417576004356001600160401b03811161041757826108ad916004016106c2565b91602435906001600160401b03821161041757806023830112156104175781600401356108d981610639565b926108e760405194856105d7565b8184526024602085019260051b82010192831161041757602401905b8282106109105750505090565b60208091833561091f81610427565b815201910190610903565b346104175761093836610882565b906109468151835114612cbb565b5f5b8251811015610020578061098661096160019385612cd1565b51838060a01b036109728488612cd1565b51169061098161083082612c9d565b6142df565b5001610948565b5f91031261041757565b34610417575f36600319011261041757602060ff5f516020615c7a5f395f51905f525416604051908152f35b34610417576060366003190112610417576024356004356044356109e681610427565b6109ee6136d9565b6109f66138b1565b815f526007602052610a1483610a0f8160405f20614e42565b612ce5565b80610a1f8484614a4a565b916001600160a01b031615610ad8575b8151905f516020615c9a5f395f51905f526020840191610a91835195610a87610a756040830197885160018060a01b03169986608086019b8c519260a088015194613bbf565b610a7e81611ce3565b6001811461433e565b519351945161041b565b94516040516001600160a01b0390961695918291610ab29142919084612c7f565b0390a45f516020615c3a5f395f51905f525f80a35f5f516020615dfa5f395f51905f525d005b5060608101516001600160a01b0316610a2f565b34610417575f366003190112610417575f546040516001600160a01b039091168152602090f35b9181601f84011215610417578235916001600160401b038311610417576020838186019501011161041757565b6101c036600319011261041757602435600435610b5c82610427565b604435610b6881610427565b6064359060a43560843560c4356001600160401b03811161041757610b91903690600401610b13565b94909360e03660e3190112610417576107cc976107ba97612cff565b3461041757604036600319011261041757610c8c602435600435610bcf61388a565b610bd76138b1565b805f526007602052610bf082610a0f8160405f20614e42565b610c876040610c17610c1285610c0586612b71565b905f5260205260405f2090565b6130b3565b610c74610c3782516080610c2d8683015161041b565b9101519087613aaf565b50610c4181611ce3565b610c4a81611ce3565b83516020810182905290600190610c6e83604081015b03601f1981018552846105d7565b146130ec565b01518015908115610c94575b4291613118565b614361565b6100206138e6565b4281109150610c80565b6001600160401b03811161058157601f01601f191660200190565b929192610cc582610c9e565b91610cd360405193846105d7565b829481845281830111610417578281602093845f960137010152565b9080601f8301121561041757816020610d0a93359101610cb9565b90565b604036600319011261041757600435610d2581610427565b6024356001600160401b03811161041757610d44903690600401610cef565b903061aaaa625c0eb760891b01148015610e21575b610e1257610d6561365f565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610de1575b50610dae57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d1a5f395f51905f528303610dcd5761002092506153d2565b632a87526960e21b5f52600483905260245ffd5b610e0491945060203d602011610e0b575b610dfc81836105d7565b81019061476d565b925f610d8d565b503d610df2565b63703e46dd60e11b5f5260045ffd5b505f516020615d1a5f395f51905f52546001600160a01b031661aaaa625c0eb760891b011415610d59565b3461041757606036600319011261041757602435600435604435610e6e6136d9565b815f526007602052610e8783610a0f8160405f20614e42565b4201804211610ed65760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612f1f565b34610417575f366003190112610417577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e125760206040515f516020615d1a5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610faa5750505090565b909192602060e082610fbf6001948851610f32565b019401929101610f9d565b3461041757602036600319011261041757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b818110611090575050611019925003836105d7565b611023825161316c565b915f5b81518110156110825760019061106661106161104186612b7f565b61105b61104e8588612cd1565b516001600160a01b031690565b906131bb565b6131d0565b6110708287612cd1565b5261107b8186612cd1565b5001611026565b604051806107cc8682610f87565b8454835260019485019487945060209093019201611004565b34610417575f36600319011261041757602060ff5f516020615dba5f395f51905f5254166040519015158152f35b60e0366003190112610417576024356004356110f282610427565b604435916110ff83610427565b6064359260c435916084359160a435916001600160401b038511610417576111b0966111656111356111a6973690600401610b13565b95909661114061388a565b6001600160a01b038516948490611157878d61440e565b61115f6138b1565b8b6144c4565b939092604051986111758a610565565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610cb9565b60e082015261466f565b5f5f516020615dfa5f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b34610417576080366003190112610417576024356004356111ff82610427565b604435906001600160401b0382116104175736602383011215610417576107cc9261123761124a933690602481600401359101610cb9565b906064359261124584610650565b613258565b604051918291826111cc565b34610417575f366003190112610417576020606454604051908152f35b3461041757604036600319011261041757602060ff6112a960243560043561129a82610427565b5f526009845260405f206131bb565b54166040519015158152f35b90602080835192838152019201905f5b8181106112d25750505090565b82518452602093840193909201916001016112c5565b906020610d0a9281815201906112b5565b34610417576020366003190112610417576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b818110611353576107cc85611347818703826105d7565b604051918291826112e8565b8254845260209093019260019283019201611330565b60e08101929161060a9190610f32565b34610417576040366003190112610417576107cc6113b860243560043561139f82610427565b6113a7613136565b505f52600660205260405f206131bb565b6004604051916113c783610586565b80546001600160a01b03168352600181015461141c90611413906113f66113ed8261041b565b60208801612f68565b61140a60a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611369565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161149c9061148e610d0a97959693600f60f81b865260e0602087015260e0860190611443565b908482036040860152611443565b60608301949094526001600160a01b031660808201525f60a082015280830360c0909101526112b5565b34610417575f366003190112610417575f516020615cda5f395f51905f5254158061155a575b1561151d576114f961477c565b611501614836565b906107cc61150d613321565b6040519384933091469186611467565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615e9a5f395f51905f5254156114ec565b9080601f8301121561041757813561158781610639565b9261159560405194856105d7565b81845260208085019260051b820101918383116104175760208201905b8382106115c157505050505090565b81356001600160401b038111610417576020916115e3878480948801016106c2565b8152019101906115b2565b6080366003190112610417576004356001600160401b038111610417576116199036906004016104a9565b90602435906001600160401b038211610417573660238301121561041757816004013561164581610639565b9261165360405194856105d7565b8184526024602085019260051b820101903682116104175760248101925b8284106116c95750506044359150506001600160401b0381116104175761169c903690600401611570565b606435926001600160401b038411610417576107cc946116c36107ba953690600401611570565b9361333c565b83356001600160401b038111610417576020916116ed83926024369187010161065b565b815201930192611671565b346104175760603660031901126104175760043561171581610427565b6024359061172282610427565b60443561172e81610650565b5f516020615e7a5f395f51905f52549260ff604085901c1615611750565b1590565b936001600160401b031680159081611812575b6001149081611808575b1590816117ff575b506117f05761179092846117876133d2565b6117e3576148d0565b61179657005b6117b45f516020615e7a5f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6117eb6133f3565b6148d0565b63f92ee8a960e01b5f5260045ffd5b9050155f611775565b303b15915061176d565b859150611763565b346104175760403660031901126104175760243560043561183961365f565b6118416138b1565b61184b8282614a4a565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615dfa5f395f51905f525d005b34610417575f366003190112610417576033546040516001600160a01b039091168152602090f35b34610417575f366003190112610417576020603454604051908152f35b3461041757604036600319011261041757602060ff6112a96024356004356118ee82610427565b5f525f516020615d9a5f395f51905f52845260405f206131bb565b34610417576040366003190112610417576004356001600160401b038111610417576119399036906004016106c2565b6024356001600160401b038111610417576119589036906004016106c2565b90611966815183511461261d565b5f5b82518110156100205780611a0e61198160019385612cd1565b5161198c8387612cd1565b519061199661388a565b61199e6138b1565b805f5260076020526119b782610a0f8160405f20614e42565b610c8760406119cc610c1285610c0586612b71565b610c746119e282516080610c2d8683015161041b565b506119ec81611ce3565b6119f581611ce3565b835160208101829052908a90610c6e8360408101610c60565b611a166138e6565b01611968565b34610417575f366003190112610417576035546040516001600160a01b039091168152602090f35b34610417575f3660031901126104175760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611a815750505090565b82516001600160a01b0316845260209384019390920191600101611a74565b34610417576020366003190112610417576004355f525f516020615d3a5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611b06576107cc85611afa818703826105d7565b60405191829182611a5e565b8254845260209093019260019283019201611ae3565b34610417576020366003190112610417577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611b5b61365f565b80600155604051908152a1005b3461041757604036600319011261041757602435600435611b876136d9565b611b8f6136d9565b611b976138b1565b805f526007602052611bb082610a0f8160405f20614e42565b611bba8282614361565b5f516020615c3a5f395f51905f525f80a35f5f516020615dfa5f395f51905f525d005b60405190611bec6020836105d7565b5f8252565b34610417575f366003190112610417576107cc604051611c126040826105d7565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611443565b34610417576020366003190112610417576004355f526004602052600160405f20015460018101809111610ed657602090604051908152f35b3461041757611c7f36610882565b90611c8d8151835114612cbb565b5f5b82518110156100205780611cc8611ca860019385612cd1565b51838060a01b03611cb98488612cd1565b51169061083561083082612c9d565b5001611c8f565b634e487b7160e01b5f52602160045260245ffd5b600a1115611ced57565b611ccf565b90600a821015611ced5752565b6020815260606040611d6460a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611443565b93611d76602082015183860190611cf2565b015191015290565b3461041757604036600319011261041757600435602435905f60408051611da4816105a1565b611dac613414565b815282602082015201525f52600860205260405f20905f526020526107cc60405f20600760405191611ddd836105a1565b611de681612faf565b8352611dfc60ff600683015416602085016130a7565b0154604082015260405191829182611cff565b34610417576020366003190112610417577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611e5181610650565b611e5961365f565b16611e65811515613444565b8060ff195f516020615c7a5f395f51905f525416175f516020615c7a5f395f51905f5255604051908152a1005b34610417576020366003190112610417577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435611ed161365f565b80606455604051908152a1005b8015150361041757565b3461041757604036600319011261041757600435602435611f0881611ede565b611f10613753565b611f25825f52600360205260405f2054151590565b15611f705760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f5260048252611f6581600360405f200161345a565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b3461041757602036600319011261041757600435611fa081611ede565b611fa8613753565b1561200657611fb561388a565b600160ff195f516020615dba5f395f51905f525416175f516020615dba5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615dba5f395f51905f525460ff81161561205d5760ff19165f516020615dba5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346104175760803660031901126104175760243560043561208c82610427565b7fa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c60406044356120bb81611ede565b6064356120c781611ede565b6120cf613753565b845f5260056020526121338161212e855f20986120fe816120f960018060a01b038216809d614e42565b61346b565b885f52600660205261211e866001612118848b5f206131bb565b01613493565b885f526009602052865f206131bb565b61345a565b8251911515825215156020820152a3005b346104175760203660031901126104175760043561216181610427565b61216961365f565b603580546001600160a01b0319166001600160a01b039290921691821790557f0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc55f80a2005b34610417575f3660031901126104175760206107ee6134b0565b34610417575f36600319011261041757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b818110612226576107cc85611347818703826105d7565b825484526020909301926001928301920161220f565b346104175760203660031901126104175760043561225981610427565b61226161365f565b6001600160a01b03166122758115156125c6565b603380546001600160a01b031916821790555f516020615d5a5f395f51905f525f80a2005b34610417576040366003190112610417576100206024356004356122bd82610427565b61098161083082612c9d565b34610417576020366003190112610417576004355f526004602052600260405f20015460018101809111610ed657602090604051908152f35b346104175760c03660031901126104175760043561231f81610427565b6024359061232c82610427565b60443561233881610650565b60843560643561234782610427565b60a435925f516020615e7a5f395f51905f52549561236c61174c8860ff9060401c1690565b966001600160401b0316801590816123c8575b60011490816123be575b1590816123b5575b506117f05761179095876123a36133d2565b156134d0576123b06133f3565b6134d0565b9050155f612391565b303b159150612389565b88915061237f565b34610417576020366003190112610417576004356123ed81610427565b6123f561365f565b6001600160a01b03166124098115156135fa565b5f548160018060a01b0382167f52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e5f80a36001600160a01b031916175f55005b346104175760803660031901126104175761002060043560243561246b81611ede565b6044359061247882610427565b6064359261248584610427565b61248d6137cd565b614cc2565b34610417575f366003190112610417576032546040516001600160a01b039091168152602090f35b34610417576040366003190112610417576024356004356124da82610427565b6124e261365f565b805f5260056020525f600461251d604083209461250c816120f960018060a01b03821680996156c5565b8484526006602052604084206131bb565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b34610417575f36600319011261041757612572615a5d565b61257a615ab4565b6040519060208201925f516020615eba5f395f51905f528452604083015260608201524660808201523060a082015260a081526125b860c0826105d7565b519020604051908152602090f35b156125cd57565b638219abe360e01b5f5260045ffd5b603380546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b1561262457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156126695760051b8101359060fe1981360301821215610417570190565b612633565b35610d0a81610427565b903590601e198136030182121561041757018035906001600160401b0382116104175760200191813603831361041757565b91908110156126695760e0020190565b908160209103126104175751610d0a81611ede565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361274097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c08601916126cf565b94803561274c81610427565b6001600160a01b031660e0850152602081013561276881610427565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561279d81610650565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d156127eb573d906127d282610c9e565b916127e060405193846105d7565b82523d5f602084013e565b606090565b604090610d0a939281528160208201520190611443565b909193929361281785841461261d565b5f945b83861061282957505050509050565b612834868585612647565b35602061284c816128468a8989612647565b0161266e565b61285c60606128468b8a8a612647565b926128d28a86888a8c6128b66080612875878486612647565b0135956128ae6128a48260a061288c82888a612647565b01359560c061289c83838b612647565b013597612647565b60e0810190612678565b9690956126aa565b946040519a8b998a996326ae802b60e11b8b5260048b016126ef565b03815f305af19081612919575b5061290e57856128ed6127c1565b60405163f495148b60e01b815291829161290a91600484016127f0565b0390fd5b60019095019461281a565b6129399060203d811161293e575b61293181836105d7565b8101906126ba565b6128df565b503d612927565b90612a2f93929161295461388a565b61295c6138b1565b803592612971845f52600560205260405f2090565b916129a6612994604083019461298e6129898761266e565b61041b565b906138f8565b6129a06129898661266e565b90612b8d565b6129b1343415612bb5565b612a456129d1865f526004602052600260405f2001600181540180915590565b966129e4602084013589819a8214612bd3565b6129f06129898661266e565b93606084019688612a008961266e565b96612a3d8c60808901359e8f60a08b019b612a1b8d8d612678565b929091604051978896602088019a8b612bf1565b03601f1981018352826105d7565b51902061393b565b612a5887612a528561266e565b87613b12565b919092600184612a6781611ce3565b14612b34575b50612a7783611ce3565b60018303612ad15750505090612aa3612a9d5f516020615c9a5f395f51905f529361266e565b9161266e565b6040516001600160a01b03909216958291612ac091429184612c7f565b0390a45b612acc6138e6565b600190565b612b0e8394612b09612b2c947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612b1495614155565b61266e565b9261266e565b9260405193849360018060a01b031698429185612c55565b0390a4612ac4565b612b6a91935088612b448661266e565b91612b62612b5b612b548a61266e565b9288612678565b3691610cb9565b928a8a613bbf565b915f612a6d565b5f52600860205260405f2090565b5f52600660205260405f2090565b15612b955750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612bbd5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612bdc575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610d0a97959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c082015201916126cf565b61060a93606092969593608083019760018060a01b03168352602083015260408201520190611cf2565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615d9a5f395f51905f52602052600160405f20015490565b15612cc257565b63031206d560e51b5f5260045ffd5b80518210156126695760209160051b010190565b15612ced5750565b6373a1390160e11b5f5260045260245ffd5b959394612d81919597939297612d1361388a565b612d586001600160a01b038816612d2a818b61440e565b612d326138b1565b612d4261298961298960e461266e565b811490612d5261298960e461266e565b91612edc565b612d79612d6961298961010461266e565b6001600160a01b038b1614612f09565b8387896144c4565b9161012435612db481612d9d86612d988787612f33565b612f33565b811015612dae87612d988888612f33565b90612f40565b612dc261298960325461041b565b90612dce61010461266e565b906101443592612ddf610164612f5e565b92610184356101a43591833b1561041757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612ed757612e9e6111a698612ea793612ac49c612ebd575b50612e95612e8061010461266e565b91612e896105fa565b9c8d5260208d01612f68565b60408b01612f68565b60608901612f68565b608087015260a086015260c08501523691610cb9565b80612ecb5f612ed1936105d7565b8061098d565b5f612e71565b6127b6565b15612ee5575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612f1057565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ed657565b15612f49575050565b63943892eb60e01b5f5260045260245260445ffd5b35610d0a81610650565b6001600160a01b039091169052565b90600182811c92168015612fa5575b6020831014612f9157565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612f86565b90604051612fbc816105bc565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f9161301482612f77565b80855291600181169081156130805750600114613041575b505060a0929161303d9103846105d7565b0152565b5f908152602081209092505b81831061306457505081016020018161303d61302c565b602091935080600191548385890101520191019091849261304d565b60ff191660208681019190915292151560051b8501909201925083915061303d905061302c565b600a821015611ced5752565b906040516130c0816105a1565b6040600782946130cf81612faf565b84526130e560ff600683015416602086016130a7565b0154910152565b156130f45750565b60405162461bcd60e51b81526020600482015290819061290a906024830190611443565b15613121575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061314382610586565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061317682610639565b61318360405191826105d7565b8281528092613194601f1991610639565b01905f5b8281106131a457505050565b6020906131af613136565b82828501015201613198565b9060018060a01b03165f5260205260405f2090565b906040516131dd81610586565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132289061321f9061140a565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126104175751610d0a81610427565b5f5490949392906001600160a01b0381161561331257604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff906132c1906084870190611443565b931691015203925af18015612ed75761060a925f916132e3575b508094613610565b613305915060203d60201161330b575b6132fd81836105d7565b810190613243565b5f6132db565b503d6132f3565b6315aeca0d60e11b5f5260045ffd5b604051906133306020836105d7565b5f808352366020840137565b91929493909384845114806133c8575b806133be575b61335b9061261d565b5f5b858110156133b2578060051b8401359060be1985360301821215610417576133ab60019261338b8389612cd1565b51613396848c612cd1565b51906133a28589612cd1565b51928901612945565b500161335d565b50945050505050600190565b5081518514613352565b508486511461334c565b5f516020615e7a5f395f51905f5280546001600160401b0319166001179055565b5f516020615e7a5f395f51905f52805460ff60401b1916600160401b179055565b60405190613421826105bc565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b1561344b57565b63f0f15b9160e01b5f5260045ffd5b9060ff801983541691151516179055565b156134735750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6065545f52600660205260405f2060015f52602052600460405f20015490565b91949293909484156135eb576001600160a01b038216156135dc576134f3615474565b6135076001600160a01b03841615156125c6565b6001600160a01b0386169061351d8215156125c6565b60ff8116156135cd5761060a966135846135919261357f6135b197613540615474565b613548615474565b613550615474565b60ff195f516020615dba5f395f51905f5254165f516020615dba5f395f51905f525561357a615474565b61549f565b6154ab565b61358c615664565b6125dc565b5f516020615d5a5f395f51905f525f80a26135ab43603455565b83614bb8565b806135bd575b50606555565b6135c79082614dbe565b5f6135b7565b6338854f1160e21b5f5260045ffd5b6324cd8cef60e01b5f5260045ffd5b63110d132360e21b5f5260045ffd5b1561360157565b6302bfb33d60e51b5f5260045ffd5b905f61060a939261248d6137cd565b916136389183549060031b91821b915f19901b19161790565b9055565b6035546001600160a01b0316330361365057565b63c82cebcb60e01b5f5260045ffd5b5f516020615e1a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff6136ac337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c6131bb565b5416156136b557565b63e2517d3f60e01b5f52336004525f516020615e1a5f395f51905f5260245260445ffd5b5f516020615e5a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff613726337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b1436131bb565b54161561372f57565b63e2517d3f60e01b5f52336004525f516020615e5a5f395f51905f5260245260445ffd5b5f516020615d7a5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff6137a0337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb4566131bb565b5416156137a957565b63e2517d3f60e01b5f52336004525f516020615d7a5f395f51905f5260245260445ffd5b5f516020615dda5f395f51905f525f525f516020615d9a5f395f51905f5260205260ff61381a337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f6131bb565b54161561382357565b63e2517d3f60e01b5f52336004525f516020615dda5f395f51905f5260245260445ffd5b805f525f516020615d9a5f395f51905f5260205260ff61386a3360405f206131bb565b5416156138745750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615dba5f395f51905f5254166138a257565b63d93c066560e01b5f5260045ffd5b5f516020615dfa5f395f51905f525c6138d75760015f516020615dfa5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615dfa5f395f51905f525d565b610d0a916001600160a01b031690614e42565b1561391257565b63b3f07a3960e01b5f5260045ffd5b156139295750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613aa5575b6139549061390b565b61397561396f5f516020615c7a5f395f51905f525460ff1690565b60ff1690565b956139838488811015613921565b5f945f935f5b8681106139a4575050505050505061060a9192811015613921565b613a016139d0836139b3615768565b6042916040519161190160f01b8352600283015260228201522090565b6139e46139dd8489612cd1565b5160ff1690565b6139ee8487612cd1565b51906139fa8589612cd1565b5192614e55565b6001600160a01b038181169088161080613a32575b613a24575b50600101613989565b600198890198909650613a1b565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615d9a5f395f51905f52602052613aa0613a99827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa56131bb565b6131bb565b5460ff1690565b613a16565b508551831461394b565b9091906001600160a01b03831660011480613b07575b613adb575b90613ad792600192614e85565b9091565b90613ae46134b0565b818101809111610ed65760645410613afc5790613aca565b505050600990600190565b506065548114613ac5565b9091906001600160a01b03831660011480613b5a575b613b39575b90613ad7925f92614e85565b90613b426134b0565b818101809111610ed65760645410613afc5790613b2d565b506065548114613b28565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610d0a92910190611443565b6001600160a01b039091168152602081019190915260400190565b938093959295613bce86612b7f565b96613bf16001613be6818060a01b038416809b6131bb565b015460a01c60ff1690565b93613c0061298960355461041b565b9560188151101580613e9d575b80613e93575b613c4e575b5050613c249450614f71565b92613c2e84611ce3565b60018414613c3d575b50505090565b613c4692614e0b565b5f8080613c37565b602081015160601c5f808092604051613c7b81612a2f602082019462483e5560e91b8652602483016111cc565b51908b5afa613c886127c1565b9080613e87575b613e7b575b5015613c1857909192935060018914159182613e4a575b5f905f808a8c89613ce48a612a2f869a8c8814613e445787965b6040519586948c6020870199631eeed51360e01b8b5260248801613b65565b51918c5af1613cf16127c1565b9080613e38575b613e23575b50897f2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f60405180613d348682901515815260200190565b0390a3613dc7575090859291613d4c575b8480613c18565b8315613d6757613d6283613c24963090846150d9565b613d45565b91602090613d8a956040519687928392632770a7eb60e21b845260048401613ba4565b03815f8b5af1918215612ed757613c24948693613da8575b50613d45565b613dc09060203d60201161293e5761293181836105d7565b505f613da2565b96613dd793969298919450614e0b565b83613de7575b5050505050600190565b15613dff57613df5936150d9565b5f80808080613ddd565b5090613e1d9250613e0e611bdd565b916001600160a01b0316615070565b50613df5565b9150915060406020820151910151915f613cfd565b50604081511015613cf8565b86613cc5565b613e5686898987614f71565b613e5f81611ce3565b60018103613e6d5750613cab565b995050505050505050505090565b6020915001515f613c94565b50602081511015613c8f565b50863b1515613c13565b506001600160a01b0387161515613c0d565b15613eb75750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126104175760405190613ee2826105bc565b819380358352602081013560208401526040810135613f0081610427565b6040840152613f11606082016107f6565b60608401526080818101359084015260a0810135916001600160401b0383116104175760a092613f419201610cef565b910152565b818110613f51575050565b5f8155600101613f46565b90601f8211613f69575050565b61060a915f516020615c5a5f395f51905f525f5260205f20906020601f840160051c83019310613fa1575b601f0160051c0190613f46565b9091508190613f94565b9190601f8111613fba57505050565b61060a925f5260205f20906020601f840160051c83019310613fa157601f0160051c0190613f46565b8160011b915f199060031b1c19161790565b90600a811015611ced5760ff80198354169116179055565b81518051825560208101516001830155604081015191929161403b906001600160a01b0316600285016125fe565b6060810151614056906001600160a01b0316600385016125fe565b6080810151600484015560a00151805160058401916001600160401b0382116105815761408d826140878554612f77565b85613fab565b602090601f83116001146140e5578260079593604095936140b6935f926140da575b5050613fe3565b90555b6140d360208201516140ca81611ce3565b60068601613ff5565b0151910155565b015190505f806140af565b90601f198316916140f9855f5260205f2090565b925f5b81811061413d5750926001928592600798966040989610614125575b505050811b0190556140b9565b01515f1960f88460031b161c191690555f8080614118565b929360206001819287860151815501950193016140fc565b9160806141ef6141e06002936141db876141d661418f61363899833595865f52600760205261419460405f2060208701359485809261597b565b613eaf565b156141fc576141c86141a860015442612f33565b915b6141bd6141b561060c565b963690613ec9565b8652602086016130a7565b6040840152610c0585612b71565b61400d565b612b7f565b61105b6129896040880161266e565b9301359201918254612f33565b6141c85f916141aa565b90614211825f615122565b918261421a5750565b5f80525f516020615d3a5f395f51905f526020526001600160a01b0316614261817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d61597b565b156142695750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161428c8382615122565b9283614296575050565b815f525f516020615d3a5f395f51905f526020526142c160405f209160018060a01b0316809261597b565b156142ca575050565b63d180cb3160e01b5f5260045260245260445ffd5b9190916142ec83826151bf565b92836142f6575050565b815f525f516020615d3a5f395f51905f5260205261432160405f209160018060a01b031680926156c5565b1561432a575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143465750565b63290cd55f60e01b5f52600a811015611ced5760045260245ffd5b9061436b91614a4a565b60018060a01b036060820151168151905f516020615c9a5f395f51905f5260208401918251946143bb610a756040830196875160018060a01b03169885608086019a8b519260a088015194613bbf565b519251935194516040516001600160a01b03909616959182916143e19142919084612c7f565b0390a4565b156143ee5750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f526005602052614432816120f960405f2060018060a01b03831690614e42565b825f52600460205260ff600360405f2001541661446a5760ff600161445e83613a9461060a9697612b7f565b015460a81c16156143e6565b82636fc24b7960e11b5f5260045260245ffd5b1561448457565b6330d45fb160e01b5f5260045ffd5b90816060910312610417578051916040602083015192015190565b156144b557565b6358e8878560e01b5f5260045ffd5b8260609161455d979596936144ee6110616144de84612b7f565b6001600160a01b038416906131bb565b6144fe61174c6040830151151590565b6145ff575b5061451261298960325461041b565b916145276001600160a01b038416151561447d565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612ed7575f955f905f936145c1575b5090829161060a949396819885151595866145b6575b5050846145ab575b5050826145a0575b50506144ae565b101590505f80614599565b101592505f80614591565b101594505f80614589565b90506145ec91965061060a93925060603d6060116145f8575b6145e481836105d7565b810190614493565b91969293919291614573565b503d6145da565b60c06146119101518480821015612f40565b5f614503565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161466a9190860190611443565b930152565b61468d81515f526004602052600160405f2001600181540180915590565b6146978251612b7f565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086146dc61298960016146d5602088019561105b612989885161041b565b015461041b565b938051906143e16146ed855161041b565b9160408101906146fd825161041b565b906147266080820196875160a084019485519861472060c087019a8b5190612f33565b93615291565b61473c6147358251995161041b565b935161041b565b9460e061474c606084015161041b565b9751935191519201519260405197889760018060a01b03169c429689614617565b90816020910312610417575190565b604051905f825f516020615c5a5f395f51905f52549161479b83612f77565b808352926001811690811561481757506001146147bf575b61060a925003836105d7565b505f516020615c5a5f395f51905f525f90815290915f516020615cfa5f395f51905f525b8183106147fb57505090602061060a928201016147b3565b60209193508060019154838589010152019101909184926147e3565b6020925061060a94915060ff191682840152151560051b8201016147b3565b604051905f825f516020615cba5f395f51905f52549161485583612f77565b808352926001811690811561481757506001146148785761060a925003836105d7565b505f516020615cba5f395f51905f525f90815290915f516020615eda5f395f51905f525b8183106148b457505090602061060a928201016147b3565b602091935080600191548385890101520191019091849261489c565b91906148da615474565b6148ee6001600160a01b03841615156125c6565b6001600160a01b038116926149048415156125c6565b60ff8316156135cd576149919261497761497d92614920615474565b614928615474565b614930615474565b60ff195f516020615dba5f395f51905f5254165f516020615dba5f395f51905f525561495a615474565b614962615474565b61496a615474565b614972615474565b614206565b506154ab565b614985615474565b620151806001556125dc565b5f516020615d5a5f395f51905f525f80a261060a43603455565b91908203918211610ed657565b60075f91828155826001820155826002820155826003820155826004820155600581016149e58154612f77565b90816149f8575b50508260068201550155565b601f8211600114614a0f57849055505b5f806149ec565b614a35614a45926001601f614a27855f5260205f2090565b920160051c82019101613f46565b5f81815260208120918190559055565b614a08565b9190614a54613414565b50825f526007602052614a6a8160405f206156c5565b15614ad457614acf61060a91845f52600860205260405f20815f52602052610c05614a9760405f20612faf565b95614ab4614aa482612b7f565b61105b61298960408b015161041b565b614ac8600260808a015192019182546149ab565b9055612b71565b6149b8565b637f11bea960e01b5f5260045260245ffd5b6003606061060a9380518455602081015160018501556040810151600285015501511515910161345a565b15614b195750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c0600491614b5260018060a01b03825116856125fe565b6020810151614b9d906001860190614b73906001600160a01b0316826125fe565b6040830151815460ff60a01b191690151560a01b60ff60a01b161781556060830151151590613493565b6080810151600285015560a081015160038501550151910155565b600190614bc68115156135fa565b614bcf826135fa565b614c67828060a01b03841693614be68515156135fa565b614bef83615912565b614c86575b614c1884614c1381614c0e875f52600560205260405f2090565b6151ac565b614b11565b614c37614c2361062a565b91614c2e8684612f68565b60208301612f68565b5f60408201525f60608201525f60808201525f60a08201525f60c0820152614c6284613a9485612b7f565b614b39565b6040515f81525f516020615e3a5f395f51905f529080602081016143e1565b614cbd614c9161061b565b8481525f60208201525f60408201525f6060820152614cb8855f52600460205260405f2090565b614ae6565b614bf4565b906143e15f516020615e3a5f395f51905f5291949394614ce38415156135fa565b6001600160a01b038616946107ba90614cfd8715156135fa565b6001600160a01b03811697614c6290614d178a15156135fa565b614d2088615912565b614d87575b614d3f81614c1381614c0e8c5f52600560205260405f2090565b614d5e614d4a61062a565b93614d558386612f68565b60208501612f68565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a9488612b7f565b614db9614d9261061b565b8981525f60208201525f60408201525f6060820152614cb88a5f52600460205260405f2090565b614d25565b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614df9576003018054918203918211610ed65755565b6004018054918201809211610ed65755565b90614e20915f52600660205260405f206131bb565b600181015460a01c60ff1615614df9576003018054918203918211610ed65755565b6001915f520160205260405f2054151590565b91610d0a9391614e64936157bc565b9092919261583e565b908160209103126104175751600a8110156104175790565b905f93614ed2613a99614ec2614e9f61298960325461041b565b95614eb46001600160a01b038816151561447d565b5f52600960205260405f2090565b6001600160a01b038516906131bb565b614f6557614f5d5791602091614f0095935f604051809881958294633f4bdec560e01b845260048401613ba4565b03925af1928315612ed7575f93614f2c575b50600183614f1f81611ce3565b03614f2657565b60019150565b614f4f91935060203d602011614f56575b614f4781836105d7565b810190614e6d565b915f614f12565b503d614f3d565b505050600191565b50505050506002905f90565b6001600160a01b031692919060018414615000578115614ff757614fc09215614fcb5760405163a9059cbb60e01b602082015291614fb8918391612a2f9160248401613ba4565b600592615021565b15610d0a5750600190565b6040516340c10f1960e01b602082015291614fef918391612a2f9160248401613ba4565b600692615021565b50505050600190565b90615013935061174c9250613e0e611bdd565b61501c57600190565b600590565b81516001600160a01b03909116915f9182916020018285620186a0f16150456127c1565b901561506a57805161506157503b1561505d57600190565b5f90565b60209150015190565b50505f90565b8147106150ca5782516001600160a01b03909116925f9182916020018486620186a0f19061509c6127c1565b91156150c357156150af575b5050600190565b805161506157503b1561505d575f806150a8565b5050505f90565b632b60b36f60e21b5f5260045ffd5b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815261060a9161511d6084836105d7565b6158ba565b805f525f516020615d9a5f395f51905f5260205260ff6151458360405f206131bb565b541661506a57805f525f516020615d9a5f395f51905f5260205261516c8260405f206131bb565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610d0a916001600160a01b03169061597b565b805f525f516020615d9a5f395f51905f5260205260ff6151e28360405f206131bb565b54161561506a57805f525f516020615d9a5f395f51905f5260205261520a8260405f206131bb565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b1561525157505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561528257565b63559d141b60e11b5f5260045ffd5b90936001600160a01b03851692600184036152ff575061060a94506152c76152b98286612f33565b34143490612dae8488612f33565b806152d3575b506159f7565b6152f46152f9916152e560335461041b565b906152ee611bdd565b91615070565b61527b565b5f6152cd565b9061530b343415612bb5565b6153206153188287612f33565b3084896150d9565b806153b4575b5061533c61174c6001613be686613a9487612b7f565b61534c575b5061060a93506159f7565b604051632770a7eb60e21b81526020818061536b883060048401613ba4565b03815f885af1918215612ed75761060a9661538f9387935f91615395575b50615247565b5f615341565b6153ae915060203d60201161293e5761293181836105d7565b5f615389565b6153cc906153c661298960335461041b565b876159c2565b5f615326565b90813b15615453575f516020615d1a5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a280511561543b5761543891615a40565b50565b50503461544457565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b60ff5f516020615e7a5f395f51905f525460401c161561549057565b631afcd79f60e31b5f5260045ffd5b61543890614962615474565b906154b4615474565b6154c260ff83161515613444565b6040918251926154d281856105d7565b60098452682b30b634b230ba37b960b91b60208501526154f4815191826105d7565b60058152640312e302e360dc1b602082015261550e615474565b615516615474565b83516001600160401b03811161058157615546816155415f516020615c5a5f395f51905f5254612f77565b613f5c565b6020601f82116001146155d55781615583939261556f9261060a97985f926140da575050613fe3565b5f516020615c5a5f395f51905f5255615b44565b6155985f5f516020615cda5f395f51905f5255565b6155ad5f5f516020615e9a5f395f51905f5255565b60ff1660ff195f516020615c7a5f395f51905f525416175f516020615c7a5f395f51905f5255565b5f516020615c5a5f395f51905f525f52601f198216955f516020615cfa5f395f51905f52965f5b81811061564c5750966001928492615583969561060a999a10615634575b505050811b015f516020615c5a5f395f51905f5255615b44565b01515f1960f88460031b161c191690555f808061561a565b838301518955600190980197602093840193016155fc565b61566c615474565b62015180600155565b8054821015612669575f5260205f2001905f90565b805480156156b1575f1901906156a08282615675565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615760575f198401848111610ed65783545f19810194908511610ed6575f95858361571397610c059503615719575b50505061568a565b55600190565b6157496157439161573a6157306157579588615675565b90549060031b1c90565b92839187615675565b9061361f565b85905f5260205260405f2090565b555f808061570b565b505050505f90565b615770615a5d565b615778615ab4565b6040519060208201925f516020615eba5f395f51905f528452604083015260608201524660808201523060a082015260a081526157b660c0826105d7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615829579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612ed7575f516001600160a01b0381161561581f57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611ced57565b61584781615834565b80615850575050565b61585981615834565b600181036158705763f645eedf60e01b5f5260045ffd5b61587981615834565b60028103615894575063fce698f760e01b5f5260045260245ffd5b806158a0600392615834565b146158a85750565b6335e2f38360e21b5f5260045260245ffd5b905f602091828151910182855af1156127b6575f513d61590957506001600160a01b0381163b155b6158e95750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156158e2565b5f8181526003602052604090205461597657600254600160401b8110156105815761595f6159498260018594016002556002615675565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b6159858282614e42565b61506a57805490600160401b82101561058157826159ad615949846001809601855584615675565b90558054925f520160205260405f2055600190565b61511d61060a93926159e960405194859263a9059cbb60e01b602085015260248401613ba4565b03601f1981018452836105d7565b90615a0c915f52600660205260405f206131bb565b600181015460a01c60ff1615615a2e576003018054918201809211610ed65755565b6004018054918203918211610ed65755565b5f80610d0a93602081519101845af4615a576127c1565b91615ae6565b615a6561477c565b8051908115615a75576020012090565b50505f516020615cda5f395f51905f52548015615a8f5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615abc614836565b8051908115615acc576020012090565b50505f516020615e9a5f395f51905f52548015615a8f5790565b90615b0a5750805115615afb57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b3b575b615b1b575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b13565b80519091906001600160401b03811161058157615b8581615b725f516020615cba5f395f51905f5254612f77565b5f516020615cba5f395f51905f52613fab565b602092601f8211600114615bb857615ba7929382915f926140da575050613fe3565b5f516020615cba5f395f51905f5255565b5f516020615cba5f395f51905f525f52601f198216935f516020615eda5f395f51905f52915f5b868110615c215750836001959610615c09575b505050811b015f516020615cba5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615bf2565b91926020600181928685015181550194019201615bdf56feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220962babe63ee1eefa7adc382e96a57fc4274abfa7e2220a1e211af8d80836259664736f6c634300081c0033"

// Deprecated: Use CrossBridgeMetaData.Sigs instead.
// CrossBridgeFuncSigs maps the 4-byte function signature to its string representation.
var CrossBridgeFuncSigs = CrossBridgeMetaData.Sigs

// CrossBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossBridgeMetaData.Bin instead.
var CrossBridgeBin = CrossBridgeMetaData.Bin

// DeployCrossBridge deploys a new Ethereum contract, binding an instance of CrossBridge to it.
func DeployCrossBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossBridge, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// CrossBridge is an auto generated Go binding around an Ethereum contract.
type CrossBridge struct {
	CrossBridgeCaller     // Read-only binding to the contract
	CrossBridgeTransactor // Write-only binding to the contract
	CrossBridgeFilterer   // Log filterer for contract events
}

// CrossBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossBridgeSession struct {
	Contract     *CrossBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossBridgeCallerSession struct {
	Contract *CrossBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrossBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossBridgeTransactorSession struct {
	Contract     *CrossBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossBridgeRaw struct {
	Contract *CrossBridge // Generic contract binding to access the raw methods on
}

// CrossBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossBridgeCallerRaw struct {
	Contract *CrossBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossBridgeTransactorRaw struct {
	Contract *CrossBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossBridge creates a new instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridge(address common.Address, backend bind.ContractBackend) (*CrossBridge, error) {
	contract, err := bindCrossBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossBridge{CrossBridgeCaller: CrossBridgeCaller{contract: contract}, CrossBridgeTransactor: CrossBridgeTransactor{contract: contract}, CrossBridgeFilterer: CrossBridgeFilterer{contract: contract}}, nil
}

// NewCrossBridgeCaller creates a new read-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeCaller(address common.Address, caller bind.ContractCaller) (*CrossBridgeCaller, error) {
	contract, err := bindCrossBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCaller{contract: contract}, nil
}

// NewCrossBridgeTransactor creates a new write-only instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossBridgeTransactor, error) {
	contract, err := bindCrossBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTransactor{contract: contract}, nil
}

// NewCrossBridgeFilterer creates a new log filterer instance of CrossBridge, bound to a specific deployed contract.
func NewCrossBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossBridgeFilterer, error) {
	contract, err := bindCrossBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeFilterer{contract: contract}, nil
}

// bindCrossBridge binds a generic wrapper to an already deployed contract.
func bindCrossBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.CrossBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.CrossBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossBridge *CrossBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossBridge *CrossBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossBridge *CrossBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridge.Contract.DEFAULTADMINROLE(&_CrossBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossBridge.Contract.DEFAULTADMINROLE(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossBridge *CrossBridgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossBridge.Contract.UPGRADEINTERFACEVERSION(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllChainIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allChainIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllChainIDs is a free data retrieval call binding the contract method 0xcf56118e.
//
// Solidity: function allChainIDs() view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllChainIDs() ([]*big.Int, error) {
	return _CrossBridge.Contract.AllChainIDs(&_CrossBridge.CallOpts)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCaller) AllPendingIndex(opts *bind.CallOpts, remoteChainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allPendingIndex", remoteChainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllPendingIndex is a free data retrieval call binding the contract method 0x79214874.
//
// Solidity: function allPendingIndex(uint256 remoteChainID) view returns(uint256[])
func (_CrossBridge *CrossBridgeCallerSession) AllPendingIndex(remoteChainID *big.Int) ([]*big.Int, error) {
	return _CrossBridge.Contract.AllPendingIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCaller) AllTokenPairs(opts *bind.CallOpts, remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "allTokenPairs", remoteChainID)

	if err != nil {
		return *new([]IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBridgeRegistryTokenPair)).(*[]IBridgeRegistryTokenPair)

	return out0, err

}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// AllTokenPairs is a free data retrieval call binding the contract method 0x5b605f5c.
//
// Solidity: function allTokenPairs(uint256 remoteChainID) view returns((address,address,bool,bool,uint256,uint256,uint256)[])
func (_CrossBridge *CrossBridgeCallerSession) AllTokenPairs(remoteChainID *big.Int) ([]IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.AllTokenPairs(&_CrossBridge.CallOpts, remoteChainID)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeExecutor() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecutor(&_CrossBridge.CallOpts)
}

// BridgeExecutor is a free data retrieval call binding the contract method 0xa10bab0b.
//
// Solidity: function bridgeExecutor() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeExecutor() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecutor(&_CrossBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeVerifier() (common.Address, error) {
	return _CrossBridge.Contract.BridgeVerifier(&_CrossBridge.CallOpts)
}

// BridgeVerifier is a free data retrieval call binding the contract method 0xf0702e8e.
//
// Solidity: function bridgeVerifier() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeVerifier() (common.Address, error) {
	return _CrossBridge.Contract.BridgeVerifier(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCaller) CrossMintableERC20Code(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossMintableERC20Code")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// CrossMintableERC20Code is a free data retrieval call binding the contract method 0x4be13f83.
//
// Solidity: function crossMintableERC20Code() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) CrossMintableERC20Code() (common.Address, error) {
	return _CrossBridge.Contract.CrossMintableERC20Code(&_CrossBridge.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) CrossSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) CrossSupply() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupply(&_CrossBridge.CallOpts)
}

// CrossSupply is a free data retrieval call binding the contract method 0xcd959706.
//
// Solidity: function crossSupply() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) CrossSupply() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupply(&_CrossBridge.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) CrossSupplyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "crossSupplyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupplyLimit(&_CrossBridge.CallOpts)
}

// CrossSupplyLimit is a free data retrieval call binding the contract method 0x69ceb2f1.
//
// Solidity: function crossSupplyLimit() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) CrossSupplyLimit() (*big.Int, error) {
	return _CrossBridge.Contract.CrossSupplyLimit(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) Dev() (common.Address, error) {
	return _CrossBridge.Contract.Dev(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossBridge.Contract.DomainSeparator(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossBridge *CrossBridgeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossBridge *CrossBridgeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossBridge.Contract.Eip712Domain(&_CrossBridge.CallOpts)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextFinalizeIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextFinalizeIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextFinalizeIndex is a free data retrieval call binding the contract method 0xd5717fc5.
//
// Solidity: function getNextFinalizeIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextFinalizeIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextFinalizeIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) GetNextInitiateIndex(opts *bind.CallOpts, remoteChainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getNextInitiateIndex", remoteChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetNextInitiateIndex is a free data retrieval call binding the contract method 0xae6893f8.
//
// Solidity: function getNextInitiateIndex(uint256 remoteChainID) view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) GetNextInitiateIndex(remoteChainID *big.Int) (*big.Int, error) {
	return _CrossBridge.Contract.GetNextInitiateIndex(&_CrossBridge.CallOpts, remoteChainID)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeCaller) GetPendingArguments(opts *bind.CallOpts, remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getPendingArguments", remoteChainID, index)

	if err != nil {
		return *new(IBridgeRegistryPendingData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryPendingData)).(*IBridgeRegistryPendingData)

	return out0, err

}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetPendingArguments is a free data retrieval call binding the contract method 0xb33eb36e.
//
// Solidity: function getPendingArguments(uint256 remoteChainID, uint256 index) view returns(((uint256,uint256,address,address,uint256,bytes),uint8,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetPendingArguments(remoteChainID *big.Int, index *big.Int) (IBridgeRegistryPendingData, error) {
	return _CrossBridge.Contract.GetPendingArguments(&_CrossBridge.CallOpts, remoteChainID, index)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridge.Contract.GetRoleAdmin(&_CrossBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossBridge.Contract.GetRoleAdmin(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossBridge *CrossBridgeCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossBridge.Contract.GetRoleMembers(&_CrossBridge.CallOpts, role)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCaller) GetTokenPair(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "getTokenPair", remoteChainID, token)

	if err != nil {
		return *new(IBridgeRegistryTokenPair), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRegistryTokenPair)).(*IBridgeRegistryTokenPair)

	return out0, err

}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// GetTokenPair is a free data retrieval call binding the contract method 0x814914b5.
//
// Solidity: function getTokenPair(uint256 remoteChainID, address token) view returns((address,address,bool,bool,uint256,uint256,uint256))
func (_CrossBridge *CrossBridgeCallerSession) GetTokenPair(remoteChainID *big.Int, token common.Address) (IBridgeRegistryTokenPair, error) {
	return _CrossBridge.Contract.GetTokenPair(&_CrossBridge.CallOpts, remoteChainID, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossBridge.Contract.HasRole(&_CrossBridge.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossBridge *CrossBridgeCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossBridge.Contract.InitializedAt(&_CrossBridge.CallOpts)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) IsTokenFinalizePaused(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "isTokenFinalizePaused", remoteChainID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridge.Contract.IsTokenFinalizePaused(&_CrossBridge.CallOpts, remoteChainID, token)
}

// IsTokenFinalizePaused is a free data retrieval call binding the contract method 0x761fe2d8.
//
// Solidity: function isTokenFinalizePaused(uint256 remoteChainID, address token) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) IsTokenFinalizePaused(remoteChainID *big.Int, token common.Address) (bool, error) {
	return _CrossBridge.Contract.IsTokenFinalizePaused(&_CrossBridge.CallOpts, remoteChainID, token)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) Paused() (bool, error) {
	return _CrossBridge.Contract.Paused(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossBridge *CrossBridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossBridge.Contract.ProxiableUUID(&_CrossBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridge.Contract.SupportsInterface(&_CrossBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossBridge *CrossBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossBridge.Contract.SupportsInterface(&_CrossBridge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossBridge *CrossBridgeCallerSession) Threshold() (uint8, error) {
	return _CrossBridge.Contract.Threshold(&_CrossBridge.CallOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) BridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "bridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x5fd262de.
//
// Solidity: function bridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) BridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.BridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.ChangeThreshold(&_CrossBridge.TransactOpts, threshold_)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactor) CreateToken(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "createToken", remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateToken is a paid mutator transaction binding the contract method 0x670e6268.
//
// Solidity: function createToken(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossBridge *CrossBridgeTransactorSession) CreateToken(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.CreateToken(&_CrossBridge.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridge(opts *bind.TransactOpts, args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridge", args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridge is a paid mutator transaction binding the contract method 0x1938e0f2.
//
// Solidity: function finalizeBridge((uint256,uint256,address,address,uint256,bytes) args, uint8[] v, bytes32[] r, bytes32[] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridge(args IBridgeRegistryFinalizeArguments, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridge(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) FinalizeBridgeBatch(opts *bind.TransactOpts, args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "finalizeBridgeBatch", args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// FinalizeBridgeBatch is a paid mutator transaction binding the contract method 0x88d67d6d.
//
// Solidity: function finalizeBridgeBatch((uint256,uint256,address,address,uint256,bytes)[] args, uint8[][] v, bytes32[][] r, bytes32[][] s) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) FinalizeBridgeBatch(args []IBridgeRegistryFinalizeArguments, v [][]uint8, r [][][32]byte, s [][][32]byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.FinalizeBridgeBatch(&_CrossBridge.TransactOpts, args, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRole(&_CrossBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRole(&_CrossBridge.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.GrantRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "initialize", owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, owner, dev_, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x89232a00.
//
// Solidity: function initialize(address owner, address dev_, uint8 threshold_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) Initialize(owner common.Address, dev_ common.Address, threshold_ uint8) (*types.Transaction, error) {
	return _CrossBridge.Contract.Initialize(&_CrossBridge.TransactOpts, owner, dev_, threshold_)
}

// InitializeCrossBridge is a paid mutator transaction binding the contract method 0xdfcbce15.
//
// Solidity: function initializeCrossBridge(address owner, address dev_, uint8 threshold_, uint256 bscChainID, address cross, uint256 crossInitialSupply) returns()
func (_CrossBridge *CrossBridgeTransactor) InitializeCrossBridge(opts *bind.TransactOpts, owner common.Address, dev_ common.Address, threshold_ uint8, bscChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "initializeCrossBridge", owner, dev_, threshold_, bscChainID, cross, crossInitialSupply)
}

// InitializeCrossBridge is a paid mutator transaction binding the contract method 0xdfcbce15.
//
// Solidity: function initializeCrossBridge(address owner, address dev_, uint8 threshold_, uint256 bscChainID, address cross, uint256 crossInitialSupply) returns()
func (_CrossBridge *CrossBridgeSession) InitializeCrossBridge(owner common.Address, dev_ common.Address, threshold_ uint8, bscChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.InitializeCrossBridge(&_CrossBridge.TransactOpts, owner, dev_, threshold_, bscChainID, cross, crossInitialSupply)
}

// InitializeCrossBridge is a paid mutator transaction binding the contract method 0xdfcbce15.
//
// Solidity: function initializeCrossBridge(address owner, address dev_, uint8 threshold_, uint256 bscChainID, address cross, uint256 crossInitialSupply) returns()
func (_CrossBridge *CrossBridgeTransactorSession) InitializeCrossBridge(owner common.Address, dev_ common.Address, threshold_ uint8, bscChainID *big.Int, cross common.Address, crossInitialSupply *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.InitializeCrossBridge(&_CrossBridge.TransactOpts, owner, dev_, threshold_, bscChainID, cross, crossInitialSupply)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) ManualReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualReleasePending", remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePending is a paid mutator transaction binding the contract method 0xaa20ed47.
//
// Solidity: function manualReleasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ManualReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeTransactor) ManualReleasePendingWithRecipient(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "manualReleasePendingWithRecipient", remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePendingWithRecipient(&_CrossBridge.TransactOpts, remoteChainID, index, recipient)
}

// ManualReleasePendingWithRecipient is a paid mutator transaction binding the contract method 0x48a00ed8.
//
// Solidity: function manualReleasePendingWithRecipient(uint256 remoteChainID, uint256 index, address recipient) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ManualReleasePendingWithRecipient(remoteChainID *big.Int, index *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.ManualReleasePendingWithRecipient(&_CrossBridge.TransactOpts, remoteChainID, index, recipient)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeToken(opts *bind.TransactOpts, toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeToken", toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeToken is a paid mutator transaction binding the contract method 0x4d5d0056.
//
// Solidity: function permitBridgeToken(uint256 toChainID, address fromToken, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, (address,address,uint256,uint256,uint8,bytes32,bytes32) permitArgs) payable returns(bool)
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeToken(toChainID *big.Int, fromToken common.Address, to common.Address, value *big.Int, networkFee *big.Int, exFee *big.Int, extraData []byte, permitArgs IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeToken(&_CrossBridge.TransactOpts, toChainID, fromToken, to, value, networkFee, exFee, extraData, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactor) PermitBridgeTokenBatch(opts *bind.TransactOpts, args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "permitBridgeTokenBatch", args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// PermitBridgeTokenBatch is a paid mutator transaction binding the contract method 0x1313996b.
//
// Solidity: function permitBridgeTokenBatch((uint256,address,address,address,uint256,uint256,uint256,bytes)[] args, (address,address,uint256,uint256,uint8,bytes32,bytes32)[] permitArgs) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) PermitBridgeTokenBatch(args []IBaseBridgeBridgeTokenArguments, permitArgs []IBaseBridgePermitArguments) (*types.Transaction, error) {
	return _CrossBridge.Contract.PermitBridgeTokenBatch(&_CrossBridge.TransactOpts, args, permitArgs)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeTransactor) RegisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "registerToken", remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// RegisterToken is a paid mutator transaction binding the contract method 0xedbbea39.
//
// Solidity: function registerToken(uint256 remoteChainID, bool isOrigin, address localToken, address remoteToken) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RegisterToken(remoteChainID *big.Int, isOrigin bool, localToken common.Address, remoteToken common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RegisterToken(&_CrossBridge.TransactOpts, remoteChainID, isOrigin, localToken, remoteToken)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) ReleasePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "releasePending", remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ReleasePending is a paid mutator transaction binding the contract method 0x4ee078ba.
//
// Solidity: function releasePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ReleasePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_CrossBridge *CrossBridgeTransactor) ReleasePendingBatch(opts *bind.TransactOpts, remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "releasePendingBatch", remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_CrossBridge *CrossBridgeSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePendingBatch(&_CrossBridge.TransactOpts, remoteChainIDs, indexes)
}

// ReleasePendingBatch is a paid mutator transaction binding the contract method 0x9f9b4f1c.
//
// Solidity: function releasePendingBatch(uint256[] remoteChainIDs, uint256[] indexes) returns()
func (_CrossBridge *CrossBridgeTransactorSession) ReleasePendingBatch(remoteChainIDs []*big.Int, indexes []*big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.ReleasePendingBatch(&_CrossBridge.TransactOpts, remoteChainIDs, indexes)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactor) RemovePending(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "removePending", remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RemovePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RemovePending is a paid mutator transaction binding the contract method 0x8ae87c5c.
//
// Solidity: function removePending(uint256 remoteChainID, uint256 index) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RemovePending(remoteChainID *big.Int, index *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.RemovePending(&_CrossBridge.TransactOpts, remoteChainID, index)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceRole(&_CrossBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RenounceRole(&_CrossBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRole(&_CrossBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRole(&_CrossBridge.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_CrossBridge *CrossBridgeTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.RevokeRoleBatch(&_CrossBridge.TransactOpts, roles, accounts)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeTransactor) SetBridgeExecutor(opts *bind.TransactOpts, _bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setBridgeExecutor", _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecutor(&_CrossBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeExecutor is a paid mutator transaction binding the contract method 0xcba25e4b.
//
// Solidity: function setBridgeExecutor(address _bridgeExecutor) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetBridgeExecutor(_bridgeExecutor common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecutor(&_CrossBridge.TransactOpts, _bridgeExecutor)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeTransactor) SetBridgeVerifier(opts *bind.TransactOpts, _bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setBridgeVerifier", _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeVerifier(&_CrossBridge.TransactOpts, _bridgeVerifier)
}

// SetBridgeVerifier is a paid mutator transaction binding the contract method 0x0b1d8d06.
//
// Solidity: function setBridgeVerifier(address _bridgeVerifier) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetBridgeVerifier(_bridgeVerifier common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeVerifier(&_CrossBridge.TransactOpts, _bridgeVerifier)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetChainPause(opts *bind.TransactOpts, remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setChainPause", remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetChainPause(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetChainPause is a paid mutator transaction binding the contract method 0xb8aa837e.
//
// Solidity: function setChainPause(uint256 remoteChainID, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetChainPause(remoteChainID *big.Int, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetChainPause(&_CrossBridge.TransactOpts, remoteChainID, pause)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossMintableERC20Code(opts *bind.TransactOpts, _crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossMintableERC20Code", _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossMintableERC20Code is a paid mutator transaction binding the contract method 0xe2af6cd7.
//
// Solidity: function setCrossMintableERC20Code(address _crossMintableERC20Code) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossMintableERC20Code(_crossMintableERC20Code common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossMintableERC20Code(&_CrossBridge.TransactOpts, _crossMintableERC20Code)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeTransactor) SetCrossSupplyLimit(opts *bind.TransactOpts, _crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setCrossSupplyLimit", _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeSession) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossSupplyLimit(&_CrossBridge.TransactOpts, _crossSupplyLimit)
}

// SetCrossSupplyLimit is a paid mutator transaction binding the contract method 0xb8886e1a.
//
// Solidity: function setCrossSupplyLimit(uint256 _crossSupplyLimit) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetCrossSupplyLimit(_crossSupplyLimit *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetCrossSupplyLimit(&_CrossBridge.TransactOpts, _crossSupplyLimit)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactor) SetDev(opts *bind.TransactOpts, dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setDev", dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address dev_) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetDev(dev_ common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetDev(&_CrossBridge.TransactOpts, dev_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetPause(&_CrossBridge.TransactOpts, set)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0xbfbf6765.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool initiatePause, bool finalizePause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, initiatePause bool, finalizePause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, initiatePause, finalizePause)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelay", delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, delay)
}

// SetVerificationDelay is a paid mutator transaction binding the contract method 0xaa1bd0bc.
//
// Solidity: function setVerificationDelay(uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelay(delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelay(&_CrossBridge.TransactOpts, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactor) SetVerificationDelayExpiration(opts *bind.TransactOpts, remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setVerificationDelayExpiration", remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelayExpiration(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// SetVerificationDelayExpiration is a paid mutator transaction binding the contract method 0x502cc5c0.
//
// Solidity: function setVerificationDelayExpiration(uint256 remoteChainID, uint256 index, uint256 delay) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetVerificationDelayExpiration(remoteChainID *big.Int, index *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetVerificationDelayExpiration(&_CrossBridge.TransactOpts, remoteChainID, index, delay)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactor) UnregisterToken(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "unregisterToken", remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UnregisterToken is a paid mutator transaction binding the contract method 0xf4509643.
//
// Solidity: function unregisterToken(uint256 remoteChainID, address token) returns()
func (_CrossBridge *CrossBridgeTransactorSession) UnregisterToken(remoteChainID *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.UnregisterToken(&_CrossBridge.TransactOpts, remoteChainID, token)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossBridge.Contract.UpgradeToAndCall(&_CrossBridge.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrossBridge *CrossBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _CrossBridge.Contract.Receive(&_CrossBridge.TransactOpts)
}

// CrossBridgeBridgeExecutorSetIterator is returned from FilterBridgeExecutorSet and is used to iterate over the raw logs and unpacked data for BridgeExecutorSet events raised by the CrossBridge contract.
type CrossBridgeBridgeExecutorSetIterator struct {
	Event *CrossBridgeBridgeExecutorSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeExecutorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeExecutorSet)
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
		it.Event = new(CrossBridgeBridgeExecutorSet)
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
func (it *CrossBridgeBridgeExecutorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeExecutorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeExecutorSet represents a BridgeExecutorSet event raised by the CrossBridge contract.
type CrossBridgeBridgeExecutorSet struct {
	BridgeExecutor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecutorSet is a free log retrieval operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeExecutorSet(opts *bind.FilterOpts, bridgeExecutor []common.Address) (*CrossBridgeBridgeExecutorSetIterator, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeExecutorSetIterator{contract: _CrossBridge.contract, event: "BridgeExecutorSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecutorSet is a free log subscription operation binding the contract event 0x0d023ad617d5258aaf577e5785aaa6950179b87058f8600c8617d005856b1fc5.
//
// Solidity: event BridgeExecutorSet(address indexed bridgeExecutor)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeExecutorSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeExecutorSet, bridgeExecutor []common.Address) (event.Subscription, error) {

	var bridgeExecutorRule []interface{}
	for _, bridgeExecutorItem := range bridgeExecutor {
		bridgeExecutorRule = append(bridgeExecutorRule, bridgeExecutorItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeExecutorSet", bridgeExecutorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeExecutorSet)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeExecutorSet(log types.Log) (*CrossBridgeBridgeExecutorSet, error) {
	event := new(CrossBridgeBridgeExecutorSet)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecutorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeFinalizedIterator is returned from FilterBridgeFinalized and is used to iterate over the raw logs and unpacked data for BridgeFinalized events raised by the CrossBridge contract.
type CrossBridgeBridgeFinalizedIterator struct {
	Event *CrossBridgeBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeFinalized)
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
		it.Event = new(CrossBridgeBridgeFinalized)
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
func (it *CrossBridgeBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeFinalized represents a BridgeFinalized event raised by the CrossBridge contract.
type CrossBridgeBridgeFinalized struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeFinalized(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgeFinalizedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeFinalizedIterator{contract: _CrossBridge.contract, event: "BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgeFinalized is a free log subscription operation binding the contract event 0x94ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6b.
//
// Solidity: event BridgeFinalized(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeFinalized(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeFinalized, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeFinalized", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeFinalized)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeFinalized(log types.Log) (*CrossBridgeBridgeFinalized, error) {
	event := new(CrossBridgeBridgeFinalized)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the CrossBridge contract.
type CrossBridgeBridgeInitiatedIterator struct {
	Event *CrossBridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeInitiated)
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
		it.Event = new(CrossBridgeBridgeInitiated)
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
func (it *CrossBridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeInitiated represents a BridgeInitiated event raised by the CrossBridge contract.
type CrossBridgeBridgeInitiated struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*CrossBridgeBridgeInitiatedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeInitiatedIterator{contract: _CrossBridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708.
//
// Solidity: event BridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address fromToken, address toToken, address indexed from, address to, uint256 value, uint256 networkFee, uint256 exFee, bytes extraData, uint256 time)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeInitiated)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeInitiated(log types.Log) (*CrossBridgeBridgeInitiated, error) {
	event := new(CrossBridgeBridgeInitiated)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgePendingIterator is returned from FilterBridgePending and is used to iterate over the raw logs and unpacked data for BridgePending events raised by the CrossBridge contract.
type CrossBridgeBridgePendingIterator struct {
	Event *CrossBridgeBridgePending // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgePending)
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
		it.Event = new(CrossBridgeBridgePending)
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
func (it *CrossBridgeBridgePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgePending represents a BridgePending event raised by the CrossBridge contract.
type CrossBridgeBridgePending struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgePending(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgePendingIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgePendingIterator{contract: _CrossBridge.contract, event: "BridgePending", logs: logs, sub: sub}, nil
}

// WatchBridgePending is a free log subscription operation binding the contract event 0x17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f.
//
// Solidity: event BridgePending(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, uint256 time, uint8 status)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgePending(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgePending, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgePending", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgePending)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgePending(log types.Log) (*CrossBridgeBridgePending, error) {
	event := new(CrossBridgeBridgePending)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeVerifierSetIterator is returned from FilterBridgeVerifierSet and is used to iterate over the raw logs and unpacked data for BridgeVerifierSet events raised by the CrossBridge contract.
type CrossBridgeBridgeVerifierSetIterator struct {
	Event *CrossBridgeBridgeVerifierSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeVerifierSet)
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
		it.Event = new(CrossBridgeBridgeVerifierSet)
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
func (it *CrossBridgeBridgeVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeVerifierSet represents a BridgeVerifierSet event raised by the CrossBridge contract.
type CrossBridgeBridgeVerifierSet struct {
	BridgeVerifier common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeVerifierSet is a free log retrieval operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeVerifierSet(opts *bind.FilterOpts, bridgeVerifier []common.Address) (*CrossBridgeBridgeVerifierSetIterator, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeVerifierSetIterator{contract: _CrossBridge.contract, event: "BridgeVerifierSet", logs: logs, sub: sub}, nil
}

// WatchBridgeVerifierSet is a free log subscription operation binding the contract event 0xa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f3.
//
// Solidity: event BridgeVerifierSet(address indexed bridgeVerifier)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeVerifierSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeVerifierSet, bridgeVerifier []common.Address) (event.Subscription, error) {

	var bridgeVerifierRule []interface{}
	for _, bridgeVerifierItem := range bridgeVerifier {
		bridgeVerifierRule = append(bridgeVerifierRule, bridgeVerifierItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeVerifierSet", bridgeVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeVerifierSet)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeVerifierSet(log types.Log) (*CrossBridgeBridgeVerifierSet, error) {
	event := new(CrossBridgeBridgeVerifierSet)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeVerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeChainPauseSetIterator is returned from FilterChainPauseSet and is used to iterate over the raw logs and unpacked data for ChainPauseSet events raised by the CrossBridge contract.
type CrossBridgeChainPauseSetIterator struct {
	Event *CrossBridgeChainPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeChainPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeChainPauseSet)
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
		it.Event = new(CrossBridgeChainPauseSet)
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
func (it *CrossBridgeChainPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeChainPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeChainPauseSet represents a ChainPauseSet event raised by the CrossBridge contract.
type CrossBridgeChainPauseSet struct {
	RemoteChainID *big.Int
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChainPauseSet is a free log retrieval operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) FilterChainPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int) (*CrossBridgeChainPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeChainPauseSetIterator{contract: _CrossBridge.contract, event: "ChainPauseSet", logs: logs, sub: sub}, nil
}

// WatchChainPauseSet is a free log subscription operation binding the contract event 0x41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c675.
//
// Solidity: event ChainPauseSet(uint256 indexed remoteChainID, bool pause)
func (_CrossBridge *CrossBridgeFilterer) WatchChainPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeChainPauseSet, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ChainPauseSet", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeChainPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseChainPauseSet(log types.Log) (*CrossBridgeChainPauseSet, error) {
	event := new(CrossBridgeChainPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "ChainPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossMintableERC20CodeSetIterator is returned from FilterCrossMintableERC20CodeSet and is used to iterate over the raw logs and unpacked data for CrossMintableERC20CodeSet events raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSetIterator struct {
	Event *CrossBridgeCrossMintableERC20CodeSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
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
		it.Event = new(CrossBridgeCrossMintableERC20CodeSet)
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
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossMintableERC20CodeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossMintableERC20CodeSet represents a CrossMintableERC20CodeSet event raised by the CrossBridge contract.
type CrossBridgeCrossMintableERC20CodeSet struct {
	OldCode common.Address
	NewCode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, oldCode []common.Address, newCode []common.Address) (*CrossBridgeCrossMintableERC20CodeSetIterator, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossMintableERC20CodeSetIterator{contract: _CrossBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0x52f716a2d4db9b34b5f7edf912076509c41b1f8ab09146f9fe82953598c03a7e.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed oldCode, address indexed newCode)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossMintableERC20CodeSet, oldCode []common.Address, newCode []common.Address) (event.Subscription, error) {

	var oldCodeRule []interface{}
	for _, oldCodeItem := range oldCode {
		oldCodeRule = append(oldCodeRule, oldCodeItem)
	}
	var newCodeRule []interface{}
	for _, newCodeItem := range newCode {
		newCodeRule = append(newCodeRule, newCodeItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", oldCodeRule, newCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossMintableERC20CodeSet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseCrossMintableERC20CodeSet(log types.Log) (*CrossBridgeCrossMintableERC20CodeSet, error) {
	event := new(CrossBridgeCrossMintableERC20CodeSet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossMintableERC20CodeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeCrossSupplyLimitSetIterator is returned from FilterCrossSupplyLimitSet and is used to iterate over the raw logs and unpacked data for CrossSupplyLimitSet events raised by the CrossBridge contract.
type CrossBridgeCrossSupplyLimitSetIterator struct {
	Event *CrossBridgeCrossSupplyLimitSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeCrossSupplyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeCrossSupplyLimitSet)
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
		it.Event = new(CrossBridgeCrossSupplyLimitSet)
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
func (it *CrossBridgeCrossSupplyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeCrossSupplyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeCrossSupplyLimitSet represents a CrossSupplyLimitSet event raised by the CrossBridge contract.
type CrossBridgeCrossSupplyLimitSet struct {
	CrossSupplyLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCrossSupplyLimitSet is a free log retrieval operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossSupplyLimitSet(opts *bind.FilterOpts) (*CrossBridgeCrossSupplyLimitSetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossSupplyLimitSetIterator{contract: _CrossBridge.contract, event: "CrossSupplyLimitSet", logs: logs, sub: sub}, nil
}

// WatchCrossSupplyLimitSet is a free log subscription operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossSupplyLimitSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossSupplyLimitSet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossSupplyLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeCrossSupplyLimitSet)
				if err := _CrossBridge.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
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

// ParseCrossSupplyLimitSet is a log parse operation binding the contract event 0x146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f.
//
// Solidity: event CrossSupplyLimitSet(uint256 crossSupplyLimit)
func (_CrossBridge *CrossBridgeFilterer) ParseCrossSupplyLimitSet(log types.Log) (*CrossBridgeCrossSupplyLimitSet, error) {
	event := new(CrossBridgeCrossSupplyLimitSet)
	if err := _CrossBridge.contract.UnpackLog(event, "CrossSupplyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeDevSetIterator is returned from FilterDevSet and is used to iterate over the raw logs and unpacked data for DevSet events raised by the CrossBridge contract.
type CrossBridgeDevSetIterator struct {
	Event *CrossBridgeDevSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeDevSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeDevSet)
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
		it.Event = new(CrossBridgeDevSet)
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
func (it *CrossBridgeDevSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeDevSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeDevSet represents a DevSet event raised by the CrossBridge contract.
type CrossBridgeDevSet struct {
	Dev common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDevSet is a free log retrieval operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridge *CrossBridgeFilterer) FilterDevSet(opts *bind.FilterOpts, dev []common.Address) (*CrossBridgeDevSetIterator, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeDevSetIterator{contract: _CrossBridge.contract, event: "DevSet", logs: logs, sub: sub}, nil
}

// WatchDevSet is a free log subscription operation binding the contract event 0x76ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc50871.
//
// Solidity: event DevSet(address indexed dev)
func (_CrossBridge *CrossBridgeFilterer) WatchDevSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeDevSet, dev []common.Address) (event.Subscription, error) {

	var devRule []interface{}
	for _, devItem := range dev {
		devRule = append(devRule, devItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "DevSet", devRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeDevSet)
				if err := _CrossBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseDevSet(log types.Log) (*CrossBridgeDevSet, error) {
	event := new(CrossBridgeDevSet)
	if err := _CrossBridge.contract.UnpackLog(event, "DevSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChangedIterator struct {
	Event *CrossBridgeEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeEIP712DomainChanged)
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
		it.Event = new(CrossBridgeEIP712DomainChanged)
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
func (it *CrossBridgeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossBridge contract.
type CrossBridgeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossBridgeEIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeEIP712DomainChangedIterator{contract: _CrossBridge.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossBridge *CrossBridgeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeEIP712DomainChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseEIP712DomainChanged(log types.Log) (*CrossBridgeEIP712DomainChanged, error) {
	event := new(CrossBridgeEIP712DomainChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the CrossBridge contract.
type CrossBridgeExtraCallExecutedIterator struct {
	Event *CrossBridgeExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeExtraCallExecuted)
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
		it.Event = new(CrossBridgeExtraCallExecuted)
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
func (it *CrossBridgeExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeExtraCallExecuted represents a ExtraCallExecuted event raised by the CrossBridge contract.
type CrossBridgeExtraCallExecuted struct {
	FromChainID *big.Int
	Index       *big.Int
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_CrossBridge *CrossBridgeFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*CrossBridgeExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeExtraCallExecutedIterator{contract: _CrossBridge.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0x2cfbbc2e6584ac8d7e60c764d69b616146f799d79cf04ee599ccf80ce8cd073f.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, bool success)
func (_CrossBridge *CrossBridgeFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *CrossBridgeExtraCallExecuted, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeExtraCallExecuted)
				if err := _CrossBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseExtraCallExecuted(log types.Log) (*CrossBridgeExtraCallExecuted, error) {
	event := new(CrossBridgeExtraCallExecuted)
	if err := _CrossBridge.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossBridge contract.
type CrossBridgeInitializedIterator struct {
	Event *CrossBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *CrossBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeInitialized)
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
		it.Event = new(CrossBridgeInitialized)
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
func (it *CrossBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeInitialized represents a Initialized event raised by the CrossBridge contract.
type CrossBridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossBridgeInitializedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeInitializedIterator{contract: _CrossBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossBridge *CrossBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeInitialized)
				if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseInitialized(log types.Log) (*CrossBridgeInitialized, error) {
	event := new(CrossBridgeInitialized)
	if err := _CrossBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeManualReleasedIterator is returned from FilterManualReleased and is used to iterate over the raw logs and unpacked data for ManualReleased events raised by the CrossBridge contract.
type CrossBridgeManualReleasedIterator struct {
	Event *CrossBridgeManualReleased // Event containing the contract specifics and raw log

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
func (it *CrossBridgeManualReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeManualReleased)
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
		it.Event = new(CrossBridgeManualReleased)
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
func (it *CrossBridgeManualReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeManualReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeManualReleased represents a ManualReleased event raised by the CrossBridge contract.
type CrossBridgeManualReleased struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterManualReleased is a free log retrieval operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterManualReleased(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgeManualReleasedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeManualReleasedIterator{contract: _CrossBridge.contract, event: "ManualReleased", logs: logs, sub: sub}, nil
}

// WatchManualReleased is a free log subscription operation binding the contract event 0xb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1.
//
// Solidity: event ManualReleased(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchManualReleased(opts *bind.WatchOpts, sink chan<- *CrossBridgeManualReleased, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ManualReleased", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeManualReleased)
				if err := _CrossBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseManualReleased(log types.Log) (*CrossBridgeManualReleased, error) {
	event := new(CrossBridgeManualReleased)
	if err := _CrossBridge.contract.UnpackLog(event, "ManualReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossBridge contract.
type CrossBridgePausedIterator struct {
	Event *CrossBridgePaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePaused)
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
		it.Event = new(CrossBridgePaused)
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
func (it *CrossBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePaused represents a Paused event raised by the CrossBridge contract.
type CrossBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossBridgePausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgePausedIterator{contract: _CrossBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossBridgePaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePaused(log types.Log) (*CrossBridgePaused, error) {
	event := new(CrossBridgePaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgePendingRemovedIterator is returned from FilterPendingRemoved and is used to iterate over the raw logs and unpacked data for PendingRemoved events raised by the CrossBridge contract.
type CrossBridgePendingRemovedIterator struct {
	Event *CrossBridgePendingRemoved // Event containing the contract specifics and raw log

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
func (it *CrossBridgePendingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgePendingRemoved)
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
		it.Event = new(CrossBridgePendingRemoved)
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
func (it *CrossBridgePendingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgePendingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgePendingRemoved represents a PendingRemoved event raised by the CrossBridge contract.
type CrossBridgePendingRemoved struct {
	RemoteChainID *big.Int
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingRemoved is a free log retrieval operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) FilterPendingRemoved(opts *bind.FilterOpts, remoteChainID []*big.Int, index []*big.Int) (*CrossBridgePendingRemovedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgePendingRemovedIterator{contract: _CrossBridge.contract, event: "PendingRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingRemoved is a free log subscription operation binding the contract event 0xe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf3.
//
// Solidity: event PendingRemoved(uint256 indexed remoteChainID, uint256 indexed index)
func (_CrossBridge *CrossBridgeFilterer) WatchPendingRemoved(opts *bind.WatchOpts, sink chan<- *CrossBridgePendingRemoved, remoteChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "PendingRemoved", remoteChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgePendingRemoved)
				if err := _CrossBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParsePendingRemoved(log types.Log) (*CrossBridgePendingRemoved, error) {
	event := new(CrossBridgePendingRemoved)
	if err := _CrossBridge.contract.UnpackLog(event, "PendingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossBridge contract.
type CrossBridgeRoleAdminChangedIterator struct {
	Event *CrossBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleAdminChanged)
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
		it.Event = new(CrossBridgeRoleAdminChanged)
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
func (it *CrossBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the CrossBridge contract.
type CrossBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleAdminChangedIterator{contract: _CrossBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleAdminChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*CrossBridgeRoleAdminChanged, error) {
	event := new(CrossBridgeRoleAdminChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossBridge contract.
type CrossBridgeRoleGrantedIterator struct {
	Event *CrossBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleGranted)
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
		it.Event = new(CrossBridgeRoleGranted)
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
func (it *CrossBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleGranted represents a RoleGranted event raised by the CrossBridge contract.
type CrossBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleGrantedIterator{contract: _CrossBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleGranted)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleGranted(log types.Log) (*CrossBridgeRoleGranted, error) {
	event := new(CrossBridgeRoleGranted)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossBridge contract.
type CrossBridgeRoleRevokedIterator struct {
	Event *CrossBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeRoleRevoked)
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
		it.Event = new(CrossBridgeRoleRevoked)
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
func (it *CrossBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeRoleRevoked represents a RoleRevoked event raised by the CrossBridge contract.
type CrossBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeRoleRevokedIterator{contract: _CrossBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossBridge *CrossBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeRoleRevoked)
				if err := _CrossBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseRoleRevoked(log types.Log) (*CrossBridgeRoleRevoked, error) {
	event := new(CrossBridgeRoleRevoked)
	if err := _CrossBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossBridge contract.
type CrossBridgeThresholdChangedIterator struct {
	Event *CrossBridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossBridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeThresholdChanged)
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
		it.Event = new(CrossBridgeThresholdChanged)
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
func (it *CrossBridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeThresholdChanged represents a ThresholdChanged event raised by the CrossBridge contract.
type CrossBridgeThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossBridgeThresholdChangedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeThresholdChangedIterator{contract: _CrossBridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossBridge *CrossBridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossBridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeThresholdChanged)
				if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseThresholdChanged(log types.Log) (*CrossBridgeThresholdChanged, error) {
	event := new(CrossBridgeThresholdChanged)
	if err := _CrossBridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairRegisteredIterator is returned from FilterTokenPairRegistered and is used to iterate over the raw logs and unpacked data for TokenPairRegistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairRegisteredIterator struct {
	Event *CrossBridgeTokenPairRegistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPairRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairRegistered)
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
		it.Event = new(CrossBridgeTokenPairRegistered)
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
func (it *CrossBridgeTokenPairRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairRegistered represents a TokenPairRegistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairRegistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairRegistered is a free log retrieval operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairRegistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (*CrossBridgeTokenPairRegisteredIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairRegisteredIterator{contract: _CrossBridge.contract, event: "TokenPairRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairRegistered is a free log subscription operation binding the contract event 0xe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db.
//
// Solidity: event TokenPairRegistered(uint256 indexed remoteChainID, address indexed localToken, address indexed remoteToken, bool isOrigin)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairRegistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairRegistered, remoteChainID []*big.Int, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairRegistered", remoteChainIDRule, localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairRegistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairRegistered(log types.Log) (*CrossBridgeTokenPairRegistered, error) {
	event := new(CrossBridgeTokenPairRegistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPairUnregisteredIterator is returned from FilterTokenPairUnregistered and is used to iterate over the raw logs and unpacked data for TokenPairUnregistered events raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregisteredIterator struct {
	Event *CrossBridgeTokenPairUnregistered // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPairUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPairUnregistered)
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
		it.Event = new(CrossBridgeTokenPairUnregistered)
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
func (it *CrossBridgeTokenPairUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPairUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPairUnregistered represents a TokenPairUnregistered event raised by the CrossBridge contract.
type CrossBridgeTokenPairUnregistered struct {
	RemoteChainID *big.Int
	LocalToken    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPairUnregistered is a free log retrieval operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPairUnregistered(opts *bind.FilterOpts, remoteChainID []*big.Int, localToken []common.Address) (*CrossBridgeTokenPairUnregisteredIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPairUnregisteredIterator{contract: _CrossBridge.contract, event: "TokenPairUnregistered", logs: logs, sub: sub}, nil
}

// WatchTokenPairUnregistered is a free log subscription operation binding the contract event 0xa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d.
//
// Solidity: event TokenPairUnregistered(uint256 indexed remoteChainID, address indexed localToken)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPairUnregistered(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPairUnregistered, remoteChainID []*big.Int, localToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPairUnregistered", remoteChainIDRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPairUnregistered)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPairUnregistered(log types.Log) (*CrossBridgeTokenPairUnregistered, error) {
	event := new(CrossBridgeTokenPairUnregistered)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPairUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeTokenPauseSetIterator is returned from FilterTokenPauseSet and is used to iterate over the raw logs and unpacked data for TokenPauseSet events raised by the CrossBridge contract.
type CrossBridgeTokenPauseSetIterator struct {
	Event *CrossBridgeTokenPauseSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeTokenPauseSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeTokenPauseSet)
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
		it.Event = new(CrossBridgeTokenPauseSet)
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
func (it *CrossBridgeTokenPauseSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeTokenPauseSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeTokenPauseSet represents a TokenPauseSet event raised by the CrossBridge contract.
type CrossBridgeTokenPauseSet struct {
	RemoteChainID *big.Int
	Token         common.Address
	InitiatePause bool
	FinalizePause bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridge *CrossBridgeFilterer) FilterTokenPauseSet(opts *bind.FilterOpts, remoteChainID []*big.Int, token []common.Address) (*CrossBridgeTokenPauseSetIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeTokenPauseSetIterator{contract: _CrossBridge.contract, event: "TokenPauseSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0xa8b29da12d479a45d43650e31c679be7f1266725b6852395ff9cb31357d1bd1c.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool initiatePause, bool finalizePause)
func (_CrossBridge *CrossBridgeFilterer) WatchTokenPauseSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeTokenPauseSet, remoteChainID []*big.Int, token []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "TokenPauseSet", remoteChainIDRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeTokenPauseSet)
				if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseTokenPauseSet(log types.Log) (*CrossBridgeTokenPauseSet, error) {
	event := new(CrossBridgeTokenPauseSet)
	if err := _CrossBridge.contract.UnpackLog(event, "TokenPauseSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossBridge contract.
type CrossBridgeUnpausedIterator struct {
	Event *CrossBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUnpaused)
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
		it.Event = new(CrossBridgeUnpaused)
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
func (it *CrossBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUnpaused represents a Unpaused event raised by the CrossBridge contract.
type CrossBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossBridgeUnpausedIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUnpausedIterator{contract: _CrossBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossBridge *CrossBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUnpaused)
				if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUnpaused(log types.Log) (*CrossBridgeUnpaused, error) {
	event := new(CrossBridgeUnpaused)
	if err := _CrossBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossBridge contract.
type CrossBridgeUpgradedIterator struct {
	Event *CrossBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeUpgraded)
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
		it.Event = new(CrossBridgeUpgraded)
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
func (it *CrossBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeUpgraded represents a Upgraded event raised by the CrossBridge contract.
type CrossBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeUpgradedIterator{contract: _CrossBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossBridge *CrossBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeUpgraded)
				if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseUpgraded(log types.Log) (*CrossBridgeUpgraded, error) {
	event := new(CrossBridgeUpgraded)
	if err := _CrossBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeVerificationDelayExpirationSetIterator is returned from FilterVerificationDelayExpirationSet and is used to iterate over the raw logs and unpacked data for VerificationDelayExpirationSet events raised by the CrossBridge contract.
type CrossBridgeVerificationDelayExpirationSetIterator struct {
	Event *CrossBridgeVerificationDelayExpirationSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeVerificationDelayExpirationSet)
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
		it.Event = new(CrossBridgeVerificationDelayExpirationSet)
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
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeVerificationDelayExpirationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeVerificationDelayExpirationSet represents a VerificationDelayExpirationSet event raised by the CrossBridge contract.
type CrossBridgeVerificationDelayExpirationSet struct {
	FromChainID *big.Int
	Index       *big.Int
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelayExpirationSet is a free log retrieval operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) FilterVerificationDelayExpirationSet(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int) (*CrossBridgeVerificationDelayExpirationSetIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeVerificationDelayExpirationSetIterator{contract: _CrossBridge.contract, event: "VerificationDelayExpirationSet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelayExpirationSet is a free log subscription operation binding the contract event 0xe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db7985365949.
//
// Solidity: event VerificationDelayExpirationSet(uint256 indexed fromChainID, uint256 indexed index, uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) WatchVerificationDelayExpirationSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeVerificationDelayExpirationSet, fromChainID []*big.Int, index []*big.Int) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "VerificationDelayExpirationSet", fromChainIDRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeVerificationDelayExpirationSet)
				if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseVerificationDelayExpirationSet(log types.Log) (*CrossBridgeVerificationDelayExpirationSet, error) {
	event := new(CrossBridgeVerificationDelayExpirationSet)
	if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelayExpirationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeVerificationDelaySetIterator is returned from FilterVerificationDelaySet and is used to iterate over the raw logs and unpacked data for VerificationDelaySet events raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySetIterator struct {
	Event *CrossBridgeVerificationDelaySet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeVerificationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeVerificationDelaySet)
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
		it.Event = new(CrossBridgeVerificationDelaySet)
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
func (it *CrossBridgeVerificationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeVerificationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeVerificationDelaySet represents a VerificationDelaySet event raised by the CrossBridge contract.
type CrossBridgeVerificationDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVerificationDelaySet is a free log retrieval operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) FilterVerificationDelaySet(opts *bind.FilterOpts) (*CrossBridgeVerificationDelaySetIterator, error) {

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return &CrossBridgeVerificationDelaySetIterator{contract: _CrossBridge.contract, event: "VerificationDelaySet", logs: logs, sub: sub}, nil
}

// WatchVerificationDelaySet is a free log subscription operation binding the contract event 0xef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b3.
//
// Solidity: event VerificationDelaySet(uint256 delay)
func (_CrossBridge *CrossBridgeFilterer) WatchVerificationDelaySet(opts *bind.WatchOpts, sink chan<- *CrossBridgeVerificationDelaySet) (event.Subscription, error) {

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "VerificationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeVerificationDelaySet)
				if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseVerificationDelaySet(log types.Log) (*CrossBridgeVerificationDelaySet, error) {
	event := new(CrossBridgeVerificationDelaySet)
	if err := _CrossBridge.contract.UnpackLog(event, "VerificationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
