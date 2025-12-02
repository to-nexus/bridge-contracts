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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecuter\",\"outputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"_bridgeExecuter\",\"type\":\"address\"}],\"name\":\"setBridgeExecuter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BridgeExecuterCallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecuter\",\"type\":\"address\"}],\"name\":\"BridgeExecuterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"bc8cadd6": "bridgeExecuter()",
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
		"815f1de6": "setBridgeExecuter(address)",
		"0b1d8d06": "setBridgeVerifier(address)",
		"b8aa837e": "setChainPause(uint256,bool)",
		"e2af6cd7": "setCrossMintableERC20Code(address)",
		"d477f05f": "setDev(address)",
		"bedb86fb": "setPause(bool)",
		"94ddc8c6": "setTokenPause(uint256,address,bool)",
		"aa1bd0bc": "setVerificationDelay(uint256)",
		"502cc5c0": "setVerificationDelayExpiration(uint256,uint256,uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"42cde4e8": "threshold()",
		"f4509643": "unregisterToken(uint256,address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f516020615ee35f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615e1c90816100c78239608051818181610e450152610ffc0152f35b6001600160401b0319166001600160401b039081175f516020615ee35f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b610020613523565b005b5f3560e01c806301ffc9a7146103815780630b1d8d061461037c5780631313996b146103775780631938e0f214610372578063248a9ca31461036d5780632f2ff15d1461036857806336568abe14610363578063365d71e41461035e5780633aefddbf1461035957806342cde4e81461035457806348a00ed81461034f5780634be13f831461034a5780634d5d0056146103455780634ee078ba146103405780634f1ef2861461033b578063502cc5c01461033657806352d1902d146103315780635b605f5c1461032c5780635c975abb146103275780635fd262de14610322578063670e62681461031d5780637921487414610318578063814914b514610313578063815f1de61461030e57806384b0196e1461030957806388d67d6d1461030457806389232a00146102ff5780638ae87c5c146102fa57806391cca3db146102f557806391cf6d3e146102f057806391d14854146102eb57806394ddc8c6146102e65780639f9b4f1c146102e1578063a217fddf146102dc578063a3246ad3146102d7578063aa1bd0bc146102d2578063aa20ed47146102cd578063ad3cb1cc146102c8578063ae6893f8146102c3578063b2b49e2e146102be578063b33eb36e146102b9578063b7f3358d146102b4578063b8aa837e146102af578063bc8cadd6146102aa578063bedb86fb146102a5578063cf56118e146102a0578063d477f05f1461029b578063d547741f14610296578063d5717fc514610291578063e2af6cd71461028c578063edbbea3914610287578063f0702e8e14610282578063f45096431461027d5763f698da250361000e57612492565b6123f2565b6123ca565b612380565b6122ce565b612295565b612266565b612208565b612194565b6120ab565b612083565b611fdc565b611f59565b611ec8565b611dbb565b611d82565b611d3b565b611cb2565b611c66565b611bea565b611b8e565b611a72565b6119ca565b611972565b611955565b61192d565b6118c5565b611812565b611708565b6115e0565b6114f3565b611429565b6113a9565b6112ee565b6111e6565b6111b8565b6110d9565b610fea565b610f5b565b610e03565b610c9a565b610c2d565b610bd9565b610abb565b610a8f565b61094d565b6108ea565b6107fa565b6107c1565b610790565b6106df565b610499565b6103f8565b346103d75760203660031901126103d75760043563ffffffff60e01b81168091036103d757602090637965db0b60e01b81149081156103c6575b506040519015158152f35b6301ffc9a760e01b1490505f6103bb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103d757565b346103d75760203660031901126103d757600435610415816103e7565b61041d613572565b6001600160a01b03166104318115156124fe565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103d7578235916001600160401b0383116103d7576020808501948460051b0101116103d757565b60403660031901126103d7576004356001600160401b0381116103d7576104c4903690600401610469565b602435916001600160401b0383116103d757366023840112156103d7576004830135916001600160401b0383116103d75736602460e08502860101116103d757602461002094019161273f565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761054157604052565b610511565b60e081019081106001600160401b0382111761054157604052565b606081019081106001600160401b0382111761054157604052565b60c081019081106001600160401b0382111761054157604052565b601f909101601f19168101906001600160401b0382119082101761054157604052565b604051906105ca61010083610597565b565b604051906105ca606083610597565b604051906105ca608083610597565b604051906105ca60e083610597565b6001600160401b0381116105415760051b60200190565b60ff8116036103d757565b9080601f830112156103d7578135610632816105f9565b926106406040519485610597565b81845260208085019260051b8201019283116103d757602001905b8282106106685750505090565b60208091833561067781610610565b81520191019061065b565b9080601f830112156103d7578135610699816105f9565b926106a76040519485610597565b81845260208085019260051b8201019283116103d757602001905b8282106106cf5750505090565b81358152602091820191016106c2565b60803660031901126103d7576004356001600160401b0381116103d75760c060031982360301126103d7576024356001600160401b0381116103d75761072990369060040161061b565b906044356001600160401b0381116103d757610749903690600401610682565b60643591906001600160401b0383116103d75761078c9361077161077a943690600401610682565b9260040161287d565b60405191829182901515815260200190565b0390f35b346103d75760203660031901126103d75760206107ae600435612bc1565b604051908152f35b35906105ca826103e7565b346103d75760403660031901126103d7576100206024356004356107e4826103e7565b6107f56107f082612bc1565b61375a565b6142ae565b346103d75760403660031901126103d75760043560243561081a816103e7565b336001600160a01b03821603610833576100209161430e565b63334bd91960e11b5f5260045ffd5b9060406003198301126103d7576004356001600160401b0381116103d7578261086d91600401610682565b91602435906001600160401b0382116103d757806023830112156103d7578160040135610899816105f9565b926108a76040519485610597565b8184526024602085019260051b8201019283116103d757602401905b8282106108d05750505090565b6020809183356108df816103e7565b8152019101906108c3565b346103d7576108f836610842565b906109068151835114612bdf565b5f5b8251811015610020578061094661092160019385612bf5565b51838060a01b036109328488612bf5565b5116906109416107f082612bc1565b61430e565b5001610908565b346103d75760c03660031901126103d75760043561096a816103e7565b60243590610977826103e7565b60443561098381610610565b608435606435610992826103e7565b60a435925f516020615d675f395f51905f5254956109bb6109b78860ff9060401c1690565b1590565b966001600160401b031680159081610a7d575b6001149081610a73575b159081610a6a575b50610a5b576109fb95876109f2612c09565b610a4e57612c4b565b610a0157005b610a1f5f516020615d675f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a56612c2a565b612c4b565b63f92ee8a960e01b5f5260045ffd5b9050155f6109e0565b303b1591506109d8565b8891506109ce565b5f9103126103d757565b346103d7575f3660031901126103d757602060ff5f516020615b675f395f51905f525416604051908152f35b346103d75760603660031901126103d757602435600435604435610ade816103e7565b610ae66135ec565b610aee6137c4565b815f526007602052610b0c83610b078160405f20614e24565b612d69565b80610b178484614d7f565b916001600160a01b031615610bc5575b8151915f516020615b875f395f51905f526040820192610b74610b62855160018060a01b03169660808601978489519160a089015193613b86565b610b6b81611e2d565b60018114614788565b8251602090930151935194516040516001600160a01b0390961695918291610b9f9142919084612ba3565b0390a45f516020615b275f395f51905f525f80a35f5f516020615ce75f395f51905f525d005b5060608101516001600160a01b0316610b27565b346103d7575f3660031901126103d7575f546040516001600160a01b039091168152602090f35b9181601f840112156103d7578235916001600160401b0383116103d757602083818601950101116103d757565b6101c03660031901126103d757602435600435610c49826103e7565b604435610c55816103e7565b6064359060a43560843560c4356001600160401b0381116103d757610c7e903690600401610c00565b94909360e03660e31901126103d75761078c9761077a97612d83565b346103d75760403660031901126103d757610d82602435600435610cbc61379d565b610cc46137c4565b805f526007602052610cdd82610b078160405f20614e24565b610d7d6040610d0d610d0885610cfb865f52600860205260405f2090565b905f5260205260405f2090565b613137565b610d6a610d2d82516080610d23868301516103db565b9101519087613a0b565b50610d3781611e2d565b610d4081611e2d565b83516020810182905290600190610d6483604081015b03601f198101855284610597565b14613170565b01518015908115610d8a575b429161319c565b6147ab565b6100206137f9565b4281109150610d76565b6001600160401b03811161054157601f01601f191660200190565b929192610dbb82610d94565b91610dc96040519384610597565b8294818452818301116103d7578281602093845f960137010152565b9080601f830112156103d757816020610e0093359101610daf565b90565b60403660031901126103d757600435610e1b816103e7565b6024356001600160401b0381116103d757610e3a903690600401610de5565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f39575b50610f2a57610e7d613572565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610ef9575b50610ec657634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615c075f395f51905f528303610ee5576100209250615411565b632a87526960e21b5f52600483905260245ffd5b610f1c91945060203d602011610f23575b610f148183610597565b810190613b77565b925f610ea5565b503d610f0a565b63703e46dd60e11b5f5260045ffd5b5f516020615c075f395f51905f52546001600160a01b0316141590505f610e70565b346103d75760603660031901126103d757602435600435604435610f7d6135ec565b815f526007602052610f9683610b078160405f20614e24565b4201804211610fe55760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612fa3565b346103d7575f3660031901126103d7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f2a5760206040515f516020615c075f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110b95750505090565b909192602060e0826110ce6001948851611041565b0194019291016110ac565b346103d75760203660031901126103d757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061119f57505061112892500383610597565b61113282516131f0565b915f5b81518110156111915760019061117561117061115086612aa3565b61116a61115d8588612bf5565b516001600160a01b031690565b9061323f565b613254565b61117f8287612bf5565b5261118a8186612bf5565b5001611135565b6040518061078c8682611096565b8454835260019485019487945060209093019201611113565b346103d7575f3660031901126103d757602060ff5f516020615ca75f395f51905f5254166040519015158152f35b60e03660031901126103d757602435600435611201826103e7565b6044359161120e836103e7565b6064359260c435916084359160a435916001600160401b0385116103d7576112bf966112746112446112b5973690600401610c00565b95909661124f61379d565b6001600160a01b038516948490611266878d614850565b61126e6137c4565b8b6148f0565b939092604051986112848a610525565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610daf565b60e0820152614a9b565b5f5f516020615ce75f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103d75760803660031901126103d75760243560043561130e826103e7565b604435906001600160401b0382116103d757366023830112156103d75761078c92611346611359933690602481600401359101610daf565b906064359261135484610610565b6132dc565b604051918291826112db565b90602080835192838152019201905f5b8181106113825750505090565b8251845260209384019390920191600101611375565b906020610e00928181520190611365565b346103d75760203660031901126103d7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106114035761078c856113f781870382610597565b60405191829182611398565b82548452602090930192600192830192016113e0565b60e0810192916105ca9190611041565b346103d75760403660031901126103d75761078c61146860243560043561144f826103e7565b6114576131ba565b505f52600660205260405f2061323f565b60046040519161147783610546565b80546001600160a01b0316835260018101546114cc906114c3906114a661149d826103db565b60208801612fec565b6114ba60a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611419565b346103d75760203660031901126103d757600435611510816103e7565b611518613572565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916115b6906115a8610e0097959693600f60f81b865260e0602087015260e086019061155d565b90848203604086015261155d565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611365565b346103d7575f3660031901126103d7575f516020615bc75f395f51905f52541580611674575b1561163757611613614b99565b61161b614c53565b9061078c6116276133a5565b6040519384933091469186611581565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615d875f395f51905f525415611606565b9080601f830112156103d75781356116a1816105f9565b926116af6040519485610597565b81845260208085019260051b820101918383116103d75760208201905b8382106116db57505050505090565b81356001600160401b0381116103d7576020916116fd87848094880101610682565b8152019101906116cc565b60803660031901126103d7576004356001600160401b0381116103d757611733903690600401610469565b90602435906001600160401b0382116103d757366023830112156103d757816004013561175f816105f9565b9261176d6040519485610597565b8184526024602085019260051b820101903682116103d75760248101925b8284106117e35750506044359150506001600160401b0381116103d7576117b690369060040161168a565b606435926001600160401b0384116103d75761078c946117dd61077a95369060040161168a565b936133c0565b83356001600160401b0381116103d75760209161180783926024369187010161061b565b81520193019261178b565b346103d75760603660031901126103d75760043561182f816103e7565b6024359061183c826103e7565b60443561184881610610565b5f516020615d675f395f51905f52549260ff604085901c1615936001600160401b0316801590816118bd575b60011490816118b3575b1590816118aa575b50610a5b576109fb9284611898612c09565b1561436d576118a5612c2a565b61436d565b9050155f611886565b303b15915061187e565b859150611874565b346103d75760403660031901126103d7576024356004356118e4613572565b6118ec6137c4565b6118f68282614d7f565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615ce75f395f51905f525d005b346103d7575f3660031901126103d7576034546040516001600160a01b039091168152602090f35b346103d7575f3660031901126103d7576020603554604051908152f35b346103d75760403660031901126103d757602060ff6119b4602435600435611999826103e7565b5f525f516020615c875f395f51905f52845260405f2061323f565b54166040519015158152f35b801515036103d757565b346103d75760603660031901126103d7576024356004356119ea826103e7565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea6020604435611a19816119c0565b611a21613666565b835f5260058252611a67816001611a6160405f2098611a5181611a4c858060a01b038216809d614e24565b613456565b885f526006875260405f2061323f565b0161347e565b6040519015158152a3005b346103d75760403660031901126103d7576004356001600160401b0381116103d757611aa2903690600401610682565b6024356001600160401b0381116103d757611ac1903690600401610682565b90611acf8151835114612555565b5f5b82518110156100205780611b80611aea60019385612bf5565b51611af58387612bf5565b5190611aff61379d565b611b076137c4565b805f526007602052611b2082610b078160405f20614e24565b610d7d6040611b3e610d0885610cfb865f52600860205260405f2090565b610d6a611b5482516080610d23868301516103db565b50611b5e81611e2d565b611b6781611e2d565b835160208101829052908a90610d648360408101610d56565b611b886137f9565b01611ad1565b346103d7575f3660031901126103d75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611bcb5750505090565b82516001600160a01b0316845260209384019390920191600101611bbe565b346103d75760203660031901126103d7576004355f525f516020615c275f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611c505761078c85611c4481870382610597565b60405191829182611ba8565b8254845260209093019260019283019201611c2d565b346103d75760203660031901126103d7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611ca5613572565b80600155604051908152a1005b346103d75760403660031901126103d757602435600435611cd16135ec565b611cd96135ec565b611ce16137c4565b805f526007602052611cfa82610b078160405f20614e24565b611d0482826147ab565b5f516020615b275f395f51905f525f80a35f5f516020615ce75f395f51905f525d005b60405190611d36602083610597565b5f8252565b346103d7575f3660031901126103d75761078c604051611d5c604082610597565b60058152640352e302e360dc1b602082015260405191829160208352602083019061155d565b346103d75760203660031901126103d7576004355f526004602052600160405f20015460018101809111610fe557602090604051908152f35b346103d757611dc936610842565b90611dd78151835114612bdf565b5f5b82518110156100205780611e12611df260019385612bf5565b51838060a01b03611e038488612bf5565b5116906107f56107f082612bc1565b5001611dd9565b634e487b7160e01b5f52602160045260245ffd5b600a1115611e3757565b611e19565b90600a821015611e375752565b6020815260606040611eae60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061155d565b93611ec0602082015183860190611e3c565b015191015290565b346103d75760403660031901126103d757600435602435905f60408051611eee81610561565b611ef661349b565b815282602082015201525f52600860205260405f20905f5260205261078c60405f20600760405191611f2783610561565b611f3081613033565b8352611f4660ff6006830154166020850161312b565b0154604082015260405191829182611e49565b346103d75760203660031901126103d7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611f9b81610610565b611fa3613572565b16611faf8115156134cb565b8060ff195f516020615b675f395f51905f525416175f516020615b675f395f51905f5255604051908152a1005b346103d75760403660031901126103d757600435602435611ffc816119c0565b612004613666565b612019825f52600360205260405f2054151590565b156120705760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f526004825261206581600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103d7575f3660031901126103d7576033546040516001600160a01b039091168152602090f35b346103d75760203660031901126103d7576004356120c8816119c0565b6120d0613666565b1561212e576120dd61379d565b600160ff195f516020615ca75f395f51905f525416175f516020615ca75f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615ca75f395f51905f525460ff8116156121855760ff19165f516020615ca75f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103d7575f3660031901126103d757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106121f25761078c856113f781870382610597565b82548452602090930192600192830192016121db565b346103d75760203660031901126103d757600435612225816103e7565b61222d613572565b6001600160a01b03166122418115156124fe565b603480546001600160a01b031916821790555f516020615c475f395f51905f525f80a2005b346103d75760403660031901126103d757610020602435600435612289826103e7565b6109416107f082612bc1565b346103d75760203660031901126103d7576004355f526004602052600260405f20015460018101809111610fe557602090604051908152f35b346103d75760203660031901126103d7576004356122eb816103e7565b6122f3613572565b5f546001600160a01b03168061236e57506123186001600160a01b03821615156134e1565b5f80546001600160a01b0319166001600160a01b03928316179081905561234790612342906103db565b6103db565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103d75760803660031901126103d7576100206004356024356123a3816119c0565b604435906123b0826103e7565b606435926123bd846103e7565b6123c56136e0565b614636565b346103d7575f3660031901126103d7576032546040516001600160a01b039091168152602090f35b346103d75760403660031901126103d757602435600435612412826103e7565b61241a613572565b805f5260056020525f6004612455604083209461244481611a4c60018060a01b0382168099615500565b84845260066020526040842061323f565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103d7575f3660031901126103d7576124aa6158f2565b6124b2615949565b6040519060208201925f516020615da75f395f51905f528452604083015260608201524660808201523060a082015260a081526124f060c082610597565b519020604051908152602090f35b1561250557565b638219abe360e01b5f5260045ffd5b603480546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b1561255c57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125a15760051b8101359060fe19813603018212156103d7570190565b61256b565b35610e00816103e7565b903590601e19813603018212156103d757018035906001600160401b0382116103d7576020019181360383136103d757565b91908110156125a15760e0020190565b908160209103126103d75751610e00816119c0565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361267897939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c0860191612607565b948035612684816103e7565b6001600160a01b031660e085015260208101356126a0816103e7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff60808201356126d581610610565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d15612723573d9061270a82610d94565b916127186040519384610597565b82523d5f602084013e565b606090565b604090610e0093928152816020820152019061155d565b909193929361274f858414612555565b5f945b83861061276157505050509050565b61276c86858561257f565b3560206127848161277e8a898961257f565b016125a6565b612794606061277e8b8a8a61257f565b9261280a8a86888a8c6127ee60806127ad87848661257f565b0135956127e66127dc8260a06127c482888a61257f565b01359560c06127d483838b61257f565b01359761257f565b60e08101906125b0565b9690956125e2565b946040519a8b998a996326ae802b60e11b8b5260048b01612627565b03815f305af19081612851575b5061284657856128256126f9565b60405163f495148b60e01b81529182916128429160048401612728565b0390fd5b600190950194612752565b6128719060203d8111612876575b6128698183610597565b8101906125f2565b612817565b503d61285f565b9061296293929161288c61379d565b6128946137c4565b8035926128a9845f52600560205260405f2090565b916128d96128c760408301946128c1612342876125a6565b9061380b565b6128d3612342866125a6565b90612ab1565b6128e4343415612ad9565b612978612904865f526004602052600260405f2001600181540180915590565b96612917602084013589819a8214612af7565b612923612342866125a6565b93606084019688612933896125a6565b966129708c60808901359e8f60a08b019b61294e8d8d6125b0565b929091604051978896602088019a8b612b15565b03601f198101835282610597565b51902061384e565b61298b87612985856125a6565b87613a59565b91909260018461299a81611e2d565b14612a67575b506129aa83611e2d565b60018303612a0457505050906129d66129d05f516020615b875f395f51905f52936125a6565b916125a6565b6040516001600160a01b039092169582916129f391429184612ba3565b0390a45b6129ff6137f9565b600190565b612a418394612a3c612a5f947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a4795614180565b6125a6565b926125a6565b9260405193849360018060a01b031698429185612b79565b0390a46129f7565b612a9c91935088612a77866125a6565b91612a95612a8e612a878a6125a6565b92886125b0565b3691610daf565b9289613b86565b915f6129a0565b5f52600660205260405f2090565b15612ab95750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612ae15750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b00575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0097959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c08201520191612607565b6105ca93606092969593608083019760018060a01b03168352602083015260408201520190611e3c565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615c875f395f51905f52602052600160405f20015490565b15612be657565b63031206d560e51b5f5260045ffd5b80518210156125a15760209160051b010190565b5f516020615d675f395f51905f5280546001600160401b0319166001179055565b5f516020615d675f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612d5a576001600160a01b03841615612d4b57612c6e615082565b612c826001600160a01b03841615156124fe565b6001600160a01b03811692612c988415156124fe565b60ff831615612d3c57612d0592612cf3612cf892612cb4615082565b612cbc615082565b612cc4615082565b60ff195f516020615ca75f395f51905f5254165f516020615ca75f395f51905f5255612cee615082565b6150ad565b6150bc565b612d00615275565b612514565b5f516020615c475f395f51905f525f80a2612d1f43603555565b612d298184614526565b81612d3357505050565b6105ca9261473f565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612d715750565b6373a1390160e11b5f5260045260245ffd5b959394612e05919597939297612d9761379d565b612ddc6001600160a01b038816612dae818b614850565b612db66137c4565b612dc661234261234260e46125a6565b811490612dd661234260e46125a6565b91612f60565b612dfd612ded6123426101046125a6565b6001600160a01b038b1614612f8d565b8387896148f0565b9161012435612e3881612e2186612e1c8787612fb7565b612fb7565b811015612e3287612e1c8888612fb7565b90612fc4565b612e466123426032546103db565b90612e526101046125a6565b906101443592612e63610164612fe2565b92610184356101a43591833b156103d757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612f5b57612f226112b598612f2b936129f79c612f41575b50612f19612f046101046125a6565b91612f0d6105ba565b9c8d5260208d01612fec565b60408b01612fec565b60608901612fec565b608087015260a086015260c08501523691610daf565b80612f4f5f612f5593610597565b80610a85565b5f612ef5565b6126ee565b15612f69575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612f9457565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610fe557565b15612fcd575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e0081610610565b6001600160a01b039091169052565b90600182811c92168015613029575b602083101461301557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161300a565b906040516130408161057c565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f9161309882612ffb565b808552916001811690811561310457506001146130c5575b505060a092916130c1910384610597565b0152565b5f908152602081209092505b8183106130e85750508101602001816130c16130b0565b60209193508060019154838589010152019101909184926130d1565b60ff191660208681019190915292151560051b850190920192508391506130c190506130b0565b600a821015611e375752565b9060405161314481610561565b60406007829461315381613033565b845261316960ff6006830154166020860161312b565b0154910152565b156131785750565b60405162461bcd60e51b81526020600482015290819061284290602483019061155d565b156131a5575050565b637a88099160e11b5f5260045260245260445ffd5b604051906131c782610546565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906131fa826105f9565b6132076040519182610597565b8281528092613218601f19916105f9565b01905f5b82811061322857505050565b6020906132336131ba565b8282850101520161321c565b9060018060a01b03165f5260205260405f2090565b9060405161326181610546565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132ac906132a3906114ba565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103d75751610e00816103e7565b5f5490949392906001600160a01b0381161561339657604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff9061334590608487019061155d565b931691015203925af18015612f5b576105ca925f91613367575b5080946134f7565b613389915060203d60201161338f575b6133818183610597565b8101906132c7565b5f61335f565b503d613377565b6315aeca0d60e11b5f5260045ffd5b604051906133b4602083610597565b5f808352366020840137565b919294939093848451148061344c575b80613442575b6133df90612555565b5f5b85811015613436578060051b8401359060be19853603018212156103d75761342f60019261340f8389612bf5565b5161341a848c612bf5565b51906134268589612bf5565b5192890161287d565b50016133e1565b50945050505050600190565b50815185146133d6565b50848651146133d0565b1561345e5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b604051906134a88261057c565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134d257565b63f0f15b9160e01b5f5260045ffd5b156134e857565b6302bfb33d60e51b5f5260045ffd5b905f6105ca93926123c56136e0565b9161351f9183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361353757565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b5f516020615d075f395f51905f525f525f516020615c875f395f51905f5260205260ff6135bf337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61323f565b5416156135c857565b63e2517d3f60e01b5f52336004525f516020615d075f395f51905f5260245260445ffd5b5f516020615d475f395f51905f525f525f516020615c875f395f51905f5260205260ff613639337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361323f565b54161561364257565b63e2517d3f60e01b5f52336004525f516020615d475f395f51905f5260245260445ffd5b5f516020615c675f395f51905f525f525f516020615c875f395f51905f5260205260ff6136b3337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661323f565b5416156136bc57565b63e2517d3f60e01b5f52336004525f516020615c675f395f51905f5260245260445ffd5b5f516020615cc75f395f51905f525f525f516020615c875f395f51905f5260205260ff61372d337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61323f565b54161561373657565b63e2517d3f60e01b5f52336004525f516020615cc75f395f51905f5260245260445ffd5b805f525f516020615c875f395f51905f5260205260ff61377d3360405f2061323f565b5416156137875750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615ca75f395f51905f5254166137b557565b63d93c066560e01b5f5260045ffd5b5f516020615ce75f395f51905f525c6137ea5760015f516020615ce75f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615ce75f395f51905f525d565b610e00916001600160a01b031690614e24565b1561382557565b63b3f07a3960e01b5f5260045ffd5b1561383c5750565b631aebd17960e11b5f5260045260245ffd5b93929381519183518314806139b8575b6138679061381e565b6138886138825f516020615b675f395f51905f525460ff1690565b60ff1690565b956138968488811015613834565b5f945f935f5b8681106138b757505050505050506105ca9192811015613834565b6139146138e3836138c66155a3565b6042916040519161190160f01b8352600283015260228201522090565b6138f76138f08489612bf5565b5160ff1690565b6139018487612bf5565b519061390d8589612bf5565b5192614e37565b6001600160a01b038181169088161080613945575b613937575b5060010161389c565b60019889019890965061392e565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615c875f395f51905f526020526139b36139ac827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561323f565b61323f565b5460ff1690565b613929565b508551831461385e565b156139c957565b6330d45fb160e01b5f5260045ffd5b908160209103126103d75751600a8110156103d75790565b6001600160a01b039091168152602081019190915260400190565b9150613a46611170606092613a355f95613a3060325460018060a01b031615156139c2565b612aa3565b6001600160a01b039091169061323f565b0151613a5157600191565b506002905f90565b9190915f92613aa66060613a9f611170613a8f613a7a6123426032546103db565b96613a306001600160a01b03891615156139c2565b6001600160a01b0386169061323f565b0151151590565b613b2d5791602091613ad095935f604051809881958294633f4bdec560e01b8452600484016139f0565b03925af1928315612f5b575f93613afc575b50600183613aef81611e2d565b03613af657565b60019150565b613b1f91935060203d602011613b26575b613b178183610597565b8101906139d8565b915f613ae2565b503d613b0d565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610e009291019061155d565b908160209103126103d7575190565b92919382613b9385612aa3565b95613bb66001613bab818060a01b038716809a61323f565b015460a01c60ff1690565b92601481511180613eb8575b613bfb575b50613bd193614e4f565b92613bdb84611e2d565b60018414613bea575b50505090565b613bf392614f0e565b5f8080613be4565b9150613c3393602083015160601c6020613c196123426033546103db565b916040518098819262483e5560e91b8352600483016112db565b0381845afa938415612f5b578888958b985f91613e99575b50613c595750505093613bc7565b92809585988896936002613c776001985f52600460205260405f2090565b01549687921480159889613e5e575b5060209596505f14613e53578c613cc584925b613ca76123426033546103db565b95604051998a9889978896631eeed51360e01b885260048801613b38565b03925af1908115612f5b575f91613e34575b50613e24579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613d32898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613e1c575b50613d4e575b92613bd193614e4f565b9050613d5e6123426033546103db565b92604051916370a0823160e01b835260208380613d7e88600483016112db565b03818b5afa8015612f5b5786935f91613dfd575b5080613da2575b50909350613d44565b6020919350613dc6956040519687928392632770a7eb60e21b8452600484016139f0565b03815f8b5af1918215612f5b57613bd194869315613d9957613df69060203d602011612876576128698183610597565b505f613d99565b613e16915060203d602011610f2357610f148183610597565b5f613d92565b90505f613d3e565b505050505090916129ff92614f0e565b613e4d915060203d602011612876576128698183610597565b5f613cd7565b8c613cc55f92613c99565b939250945094613e6d93614e4f565b613e7681611e2d565b60018103613e8b5750848884848a948a613c86565b995050505050505050505090565b613eb2915060203d602011612876576128698183610597565b5f613c4b565b506033546001600160a01b0390613ed290612342906103db565b161515613bc2565b15613ee25750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103d75760405190613f0d8261057c565b819380358352602081013560208401526040810135613f2b816103e7565b6040840152613f3c606082016107b6565b60608401526080818101359084015260a0810135916001600160401b0383116103d75760a092613f6c9201610de5565b910152565b818110613f7c575050565b5f8155600101613f71565b90601f8211613f94575050565b6105ca915f516020615b475f395f51905f525f5260205f20906020601f840160051c83019310613fcc575b601f0160051c0190613f71565b9091508190613fbf565b9190601f8111613fe557505050565b6105ca925f5260205f20906020601f840160051c83019310613fcc57601f0160051c0190613f71565b8160011b915f199060031b1c19161790565b90600a811015611e375760ff80198354169116179055565b815180518255602081015160018301556040810151919291614066906001600160a01b031660028501612536565b6060810151614081906001600160a01b031660038501612536565b6080810151600484015560a00151805160058401916001600160401b038211610541576140b8826140b28554612ffb565b85613fd6565b602090601f8311600114614110578260079593604095936140e1935f92614105575b505061400e565b90555b6140fe60208201516140f581611e2d565b60068601614020565b0151910155565b015190505f806140da565b90601f19831691614124855f5260205f2090565b925f5b8181106141685750926001928592600798966040989610614150575b505050811b0190556140e4565b01515f1960f88460031b161c191690555f8080614143565b92936020600181928786015181550195019301614127565b91608061421e61420f600293613a308761420a6141ba61351f99833595865f5260076020526141bf60405f20602087013594858092615810565b613eda565b1561422b576141f36141d360015442612fb7565b915b6141e86141e06105cc565b963690613ef4565b86526020860161312b565b6040840152610cfb855f52600860205260405f2090565b614038565b61116a612342604088016125a6565b9301359201918254612fb7565b6141f35f916141d5565b90614240825f614f57565b91826142495750565b5f80525f516020615c275f395f51905f526020526001600160a01b0316614290817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615810565b156142985750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916142bb8382614f57565b92836142c5575050565b815f525f516020615c275f395f51905f526020526142f060405f209160018060a01b03168092615810565b156142f9575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161431b8382614ffa565b9283614325575050565b815f525f516020615c275f395f51905f5260205261435060405f209160018060a01b03168092615500565b15614359575050565b62a95f1b60e31b5f5260045260245260445ffd5b9190614377615082565b61438b6001600160a01b03841615156124fe565b6001600160a01b038116926143a18415156124fe565b60ff831615612d3c5761442e9261441461441a926143bd615082565b6143c5615082565b6143cd615082565b60ff195f516020615ca75f395f51905f5254165f516020615ca75f395f51905f52556143f7615082565b6143ff615082565b614407615082565b61440f615082565b614235565b506150bc565b614422615082565b62015180600155612514565b5f516020615c475f395f51905f525f80a26105ca43603555565b600360606105ca938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156144875750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916144c060018060a01b0382511685612536565b602081015161450b9060018601906144e1906001600160a01b031682612536565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061347e565b6080810151600285015560a081015160038501550151910155565b906001916145358115156134e1565b6145d8838060a01b0383169261454c8415156134e1565b614555856134e1565b61455e836157a7565b6145fa575b614587816145828161457d875f52600560205260405f2090565b614fe7565b61447f565b6145d36145926105ea565b9161459d8184612fec565b6145aa8760208501612fec565b86151560408401525f60608401525f60808401525f60a08401525f60c08401526139a785612aa3565b6144a7565b60405183151581525f516020615d275f395f51905f529080602081015b0390a4565b6146316146056105db565b8481525f60208201525f60408201525f606082015261462c855f52600460205260405f2090565b614448565b614563565b906145f55f516020615d275f395f51905f52919493946146578415156134e1565b6001600160a01b0386169461077a906146718715156134e1565b6001600160a01b038116976145d39061468b8a15156134e1565b614694886157a7565b6146fb575b6146b3816145828161457d8c5f52600560205260405f2090565b6146d26146be6105ea565b936146c98386612fec565b60208501612fec565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526139a788612aa3565b61472d6147066105db565b8981525f60208201525f60408201525f606082015261462c8a5f52600460205260405f2090565b614699565b91908203918211610fe557565b90614754915f52600660205260405f2061323f565b600181015460a01c60ff1615614776576003018054918201809211610fe55755565b6004018054918203918211610fe55755565b156147905750565b63290cd55f60e01b5f52600a811015611e375760045260245ffd5b906147b591614d7f565b60018060a01b036060820151168151915f516020615b875f395f51905f5260408201926147fd610b62855160018060a01b03169660808601978489519160a089015193613b86565b8251602090930151935194516040516001600160a01b03909616959182916145f59142919084612ba3565b156148305750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f52600560205261487481611a4c60405f2060018060a01b03831690614e24565b825f52600460205260ff600360405f200154166148ac5760ff60016148a0836139a76105ca9697612aa3565b015460a81c1615614828565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103d7578051916040602083015192015190565b156148e157565b6358e8878560e01b5f5260045ffd5b826060916149899795969361491a61117061490a84612aa3565b6001600160a01b0384169061323f565b61492a6109b76040830151151590565b614a2b575b5061493e6123426032546103db565b916149536001600160a01b03841615156139c2565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612f5b575f955f905f936149ed575b509082916105ca949396819885151595866149e2575b5050846149d7575b5050826149cc575b50506148da565b101590505f806149c5565b101592505f806149bd565b101594505f806149b5565b9050614a189196506105ca93925060603d606011614a24575b614a108183610597565b8101906148bf565b9196929391929161499f565b503d614a06565b60c0614a3d9101518480821015612fc4565b5f61492f565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e091614a96919086019061155d565b930152565b614ab981515f526004602052600160405f2001600181540180915590565b614ac38251612aa3565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708614b086123426001614b01602088019561116a61234288516103db565b01546103db565b938051906145f5614b1985516103db565b916040810190614b2982516103db565b90614b526080820196875160a0840194855198614b4c60c087019a8b5190612fb7565b936152d0565b614b68614b61825199516103db565b93516103db565b9460e0614b7860608401516103db565b9751935191519201519260405197889760018060a01b03169c429689614a43565b604051905f825f516020615b475f395f51905f525491614bb883612ffb565b8083529260018116908115614c345750600114614bdc575b6105ca92500383610597565b505f516020615b475f395f51905f525f90815290915f516020615be75f395f51905f525b818310614c185750509060206105ca92820101614bd0565b6020919350806001915483858901015201910190918492614c00565b602092506105ca94915060ff191682840152151560051b820101614bd0565b604051905f825f516020615ba75f395f51905f525491614c7283612ffb565b8083529260018116908115614c345750600114614c95576105ca92500383610597565b505f516020615ba75f395f51905f525f90815290915f516020615dc75f395f51905f525b818310614cd15750509060206105ca92820101614bd0565b6020919350806001915483858901015201910190918492614cb9565b60075f9182815582600182015582600282015582600382015582600482015560058101614d1a8154612ffb565b9081614d2d575b50508260068201550155565b601f8211600114614d4457849055505b5f80614d21565b614d6a614d7a926001601f614d5c855f5260205f2090565b920160051c82019101613f71565b5f81815260208120918190559055565b614d3d565b9190614d8961349b565b50825f526007602052614d9f8160405f20615500565b15614e1257614e0d6105ca91845f52600860205260405f20815f52602052610cfb614dcc60405f20613033565b95614de9614dd982612aa3565b61116a61234260408b01516103db565b614dfd600260808a01519201918254614732565b90555f52600860205260405f2090565b614ced565b637f11bea960e01b5f5260045260245ffd5b6001915f520160205260405f2054151590565b91610e009391614e46936155f7565b90929192615679565b6001600160a01b031692919060018414614ede578115614ed557614e9e9215614ea95760405163a9059cbb60e01b602082015291614e9691839161296291602484016139f0565b6005926156f5565b15610e005750600190565b6040516340c10f1960e01b602082015291614ecd91839161296291602484016139f0565b6006926156f5565b50505050600190565b90614f0093506109b79250614ef1611d27565b916001600160a01b031661573e565b614f0957600190565b600590565b90614f23915f52600660205260405f2061323f565b600181015460a01c60ff1615614f45576003018054918203918211610fe55755565b6004018054918201809211610fe55755565b805f525f516020615c875f395f51905f5260205260ff614f7a8360405f2061323f565b5416614fe157805f525f516020615c875f395f51905f52602052614fa18260405f2061323f565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610e00916001600160a01b031690615810565b805f525f516020615c875f395f51905f5260205260ff61501d8360405f2061323f565b541615614fe157805f525f516020615c875f395f51905f526020526150458260405f2061323f565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615d675f395f51905f525460401c161561509e57565b631afcd79f60e31b5f5260045ffd5b6150b9906143ff615082565b50565b906150c5615082565b6150d360ff831615156134cb565b6040918251926150e38185610597565b60098452682b30b634b230ba37b960b91b602085015261510581519182610597565b60058152640312e302e360dc1b602082015261511f615082565b615127615082565b83516001600160401b03811161054157615157816151525f516020615b475f395f51905f5254612ffb565b613f87565b6020601f82116001146151e657816151949392615180926105ca97985f9261410557505061400e565b5f516020615b475f395f51905f525561597b565b6151a95f5f516020615bc75f395f51905f5255565b6151be5f5f516020615d875f395f51905f5255565b60ff1660ff195f516020615b675f395f51905f525416175f516020615b675f395f51905f5255565b5f516020615b475f395f51905f525f52601f198216955f516020615be75f395f51905f52965f5b81811061525d575096600192849261519496956105ca999a10615245575b505050811b015f516020615b475f395f51905f525561597b565b01515f1960f88460031b161c191690555f808061522b565b8383015189556001909801976020938401930161520d565b61527d615082565b62015180600155565b1561529057505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156152c157565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361533e57506105ca94506153066152f88286612fb7565b34143490612e328488612fb7565b80615312575b5061473f565b615333615338916153246034546103db565b9061532d611d27565b9161573e565b6152ba565b5f61530c565b9061534a343415612ad9565b61535f6153578287612fb7565b308489615857565b806153f3575b5061537b6109b76001613bab866139a787612aa3565b61538b575b506105ca935061473f565b604051632770a7eb60e21b8152602081806153aa8830600484016139f0565b03815f885af1918215612f5b576105ca966153ce9387935f916153d4575b50615286565b5f615380565b6153ed915060203d602011612876576128698183610597565b5f6153c8565b61540b906154056123426034546103db565b876158a0565b5f615365565b90813b1561548f575f516020615c075f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115615477576150b9916158d5565b50503461548057565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125a1575f5260205f2001905f90565b805480156154ec575f1901906154db82826154b0565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f1461559b575f198401848111610fe55783545f19810194908511610fe5575f95858361554e97610cfb9503615554575b5050506154c5565b55600190565b61558461557e9161557561556b61559295886154b0565b90549060031b1c90565b928391876154b0565b90613506565b85905f5260205260405f2090565b555f8080615546565b505050505f90565b6155ab6158f2565b6155b3615949565b6040519060208201925f516020615da75f395f51905f528452604083015260608201524660808201523060a082015260a081526155f160c082610597565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615664579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612f5b575f516001600160a01b0381161561565a57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611e3757565b6156828161566f565b8061568b575050565b6156948161566f565b600181036156ab5763f645eedf60e01b5f5260045ffd5b6156b48161566f565b600281036156cf575063fce698f760e01b5f5260045260245ffd5b806156db60039261566f565b146156e35750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f16157196126f9565b9015614fe157805161573557503b1561573157600190565b5f90565b60209150015190565b8147106157985782516001600160a01b03909116925f9182916020018486620186a0f19061576a6126f9565b9115615791571561577d575b5050600190565b805161573557503b15615731575f80615776565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461580b57600254600160401b811015610541576157f46157de82600185940160025560026154b0565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61581a8282614e24565b614fe157805490600160401b82101561054157826158426157de8460018096018555846154b0565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105ca9161589b608483610597565b615a70565b61589b6105ca93926158c760405194859263a9059cbb60e01b6020850152602484016139f0565b03601f198101845283610597565b5f80610e0093602081519101845af46158ec6126f9565b91615ac8565b6158fa614b99565b805190811561590a576020012090565b50505f516020615bc75f395f51905f525480156159245790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615951614c53565b8051908115615961576020012090565b50505f516020615d875f395f51905f525480156159245790565b80519091906001600160401b038111610541576159bc816159a95f516020615ba75f395f51905f5254612ffb565b5f516020615ba75f395f51905f52613fd6565b602092601f82116001146159ef576159de929382915f9261410557505061400e565b5f516020615ba75f395f51905f5255565b5f516020615ba75f395f51905f525f52601f198216935f516020615dc75f395f51905f52915f5b868110615a585750836001959610615a40575b505050811b015f516020615ba75f395f51905f5255565b01515f1960f88460031b161c191690555f8080615a29565b91926020600181928685015181550194019201615a16565b905f602091828151910182855af1156126ee575f513d615abf57506001600160a01b0381163b155b615a9f5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615a98565b90615aec5750805115615add57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b1d575b615afd575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615af556feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a26469706673582212204646ca45fdf6dc0ae23dafd2c733d4ade835b388d0da3c3464656eb735f3409364736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeMetaData.ABI instead.
var BSCBridgeABI = BSCBridgeMetaData.ABI

// BSCBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BSCBridgeBinRuntime = "60806040526004361015610022575b3615610018575f80fd5b610020613523565b005b5f3560e01c806301ffc9a7146103815780630b1d8d061461037c5780631313996b146103775780631938e0f214610372578063248a9ca31461036d5780632f2ff15d1461036857806336568abe14610363578063365d71e41461035e5780633aefddbf1461035957806342cde4e81461035457806348a00ed81461034f5780634be13f831461034a5780634d5d0056146103455780634ee078ba146103405780634f1ef2861461033b578063502cc5c01461033657806352d1902d146103315780635b605f5c1461032c5780635c975abb146103275780635fd262de14610322578063670e62681461031d5780637921487414610318578063814914b514610313578063815f1de61461030e57806384b0196e1461030957806388d67d6d1461030457806389232a00146102ff5780638ae87c5c146102fa57806391cca3db146102f557806391cf6d3e146102f057806391d14854146102eb57806394ddc8c6146102e65780639f9b4f1c146102e1578063a217fddf146102dc578063a3246ad3146102d7578063aa1bd0bc146102d2578063aa20ed47146102cd578063ad3cb1cc146102c8578063ae6893f8146102c3578063b2b49e2e146102be578063b33eb36e146102b9578063b7f3358d146102b4578063b8aa837e146102af578063bc8cadd6146102aa578063bedb86fb146102a5578063cf56118e146102a0578063d477f05f1461029b578063d547741f14610296578063d5717fc514610291578063e2af6cd71461028c578063edbbea3914610287578063f0702e8e14610282578063f45096431461027d5763f698da250361000e57612492565b6123f2565b6123ca565b612380565b6122ce565b612295565b612266565b612208565b612194565b6120ab565b612083565b611fdc565b611f59565b611ec8565b611dbb565b611d82565b611d3b565b611cb2565b611c66565b611bea565b611b8e565b611a72565b6119ca565b611972565b611955565b61192d565b6118c5565b611812565b611708565b6115e0565b6114f3565b611429565b6113a9565b6112ee565b6111e6565b6111b8565b6110d9565b610fea565b610f5b565b610e03565b610c9a565b610c2d565b610bd9565b610abb565b610a8f565b61094d565b6108ea565b6107fa565b6107c1565b610790565b6106df565b610499565b6103f8565b346103d75760203660031901126103d75760043563ffffffff60e01b81168091036103d757602090637965db0b60e01b81149081156103c6575b506040519015158152f35b6301ffc9a760e01b1490505f6103bb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103d757565b346103d75760203660031901126103d757600435610415816103e7565b61041d613572565b6001600160a01b03166104318115156124fe565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103d7578235916001600160401b0383116103d7576020808501948460051b0101116103d757565b60403660031901126103d7576004356001600160401b0381116103d7576104c4903690600401610469565b602435916001600160401b0383116103d757366023840112156103d7576004830135916001600160401b0383116103d75736602460e08502860101116103d757602461002094019161273f565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761054157604052565b610511565b60e081019081106001600160401b0382111761054157604052565b606081019081106001600160401b0382111761054157604052565b60c081019081106001600160401b0382111761054157604052565b601f909101601f19168101906001600160401b0382119082101761054157604052565b604051906105ca61010083610597565b565b604051906105ca606083610597565b604051906105ca608083610597565b604051906105ca60e083610597565b6001600160401b0381116105415760051b60200190565b60ff8116036103d757565b9080601f830112156103d7578135610632816105f9565b926106406040519485610597565b81845260208085019260051b8201019283116103d757602001905b8282106106685750505090565b60208091833561067781610610565b81520191019061065b565b9080601f830112156103d7578135610699816105f9565b926106a76040519485610597565b81845260208085019260051b8201019283116103d757602001905b8282106106cf5750505090565b81358152602091820191016106c2565b60803660031901126103d7576004356001600160401b0381116103d75760c060031982360301126103d7576024356001600160401b0381116103d75761072990369060040161061b565b906044356001600160401b0381116103d757610749903690600401610682565b60643591906001600160401b0383116103d75761078c9361077161077a943690600401610682565b9260040161287d565b60405191829182901515815260200190565b0390f35b346103d75760203660031901126103d75760206107ae600435612bc1565b604051908152f35b35906105ca826103e7565b346103d75760403660031901126103d7576100206024356004356107e4826103e7565b6107f56107f082612bc1565b61375a565b6142ae565b346103d75760403660031901126103d75760043560243561081a816103e7565b336001600160a01b03821603610833576100209161430e565b63334bd91960e11b5f5260045ffd5b9060406003198301126103d7576004356001600160401b0381116103d7578261086d91600401610682565b91602435906001600160401b0382116103d757806023830112156103d7578160040135610899816105f9565b926108a76040519485610597565b8184526024602085019260051b8201019283116103d757602401905b8282106108d05750505090565b6020809183356108df816103e7565b8152019101906108c3565b346103d7576108f836610842565b906109068151835114612bdf565b5f5b8251811015610020578061094661092160019385612bf5565b51838060a01b036109328488612bf5565b5116906109416107f082612bc1565b61430e565b5001610908565b346103d75760c03660031901126103d75760043561096a816103e7565b60243590610977826103e7565b60443561098381610610565b608435606435610992826103e7565b60a435925f516020615d675f395f51905f5254956109bb6109b78860ff9060401c1690565b1590565b966001600160401b031680159081610a7d575b6001149081610a73575b159081610a6a575b50610a5b576109fb95876109f2612c09565b610a4e57612c4b565b610a0157005b610a1f5f516020615d675f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a56612c2a565b612c4b565b63f92ee8a960e01b5f5260045ffd5b9050155f6109e0565b303b1591506109d8565b8891506109ce565b5f9103126103d757565b346103d7575f3660031901126103d757602060ff5f516020615b675f395f51905f525416604051908152f35b346103d75760603660031901126103d757602435600435604435610ade816103e7565b610ae66135ec565b610aee6137c4565b815f526007602052610b0c83610b078160405f20614e24565b612d69565b80610b178484614d7f565b916001600160a01b031615610bc5575b8151915f516020615b875f395f51905f526040820192610b74610b62855160018060a01b03169660808601978489519160a089015193613b86565b610b6b81611e2d565b60018114614788565b8251602090930151935194516040516001600160a01b0390961695918291610b9f9142919084612ba3565b0390a45f516020615b275f395f51905f525f80a35f5f516020615ce75f395f51905f525d005b5060608101516001600160a01b0316610b27565b346103d7575f3660031901126103d7575f546040516001600160a01b039091168152602090f35b9181601f840112156103d7578235916001600160401b0383116103d757602083818601950101116103d757565b6101c03660031901126103d757602435600435610c49826103e7565b604435610c55816103e7565b6064359060a43560843560c4356001600160401b0381116103d757610c7e903690600401610c00565b94909360e03660e31901126103d75761078c9761077a97612d83565b346103d75760403660031901126103d757610d82602435600435610cbc61379d565b610cc46137c4565b805f526007602052610cdd82610b078160405f20614e24565b610d7d6040610d0d610d0885610cfb865f52600860205260405f2090565b905f5260205260405f2090565b613137565b610d6a610d2d82516080610d23868301516103db565b9101519087613a0b565b50610d3781611e2d565b610d4081611e2d565b83516020810182905290600190610d6483604081015b03601f198101855284610597565b14613170565b01518015908115610d8a575b429161319c565b6147ab565b6100206137f9565b4281109150610d76565b6001600160401b03811161054157601f01601f191660200190565b929192610dbb82610d94565b91610dc96040519384610597565b8294818452818301116103d7578281602093845f960137010152565b9080601f830112156103d757816020610e0093359101610daf565b90565b60403660031901126103d757600435610e1b816103e7565b6024356001600160401b0381116103d757610e3a903690600401610de5565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f39575b50610f2a57610e7d613572565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610ef9575b50610ec657634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615c075f395f51905f528303610ee5576100209250615411565b632a87526960e21b5f52600483905260245ffd5b610f1c91945060203d602011610f23575b610f148183610597565b810190613b77565b925f610ea5565b503d610f0a565b63703e46dd60e11b5f5260045ffd5b5f516020615c075f395f51905f52546001600160a01b0316141590505f610e70565b346103d75760603660031901126103d757602435600435604435610f7d6135ec565b815f526007602052610f9683610b078160405f20614e24565b4201804211610fe55760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612fa3565b346103d7575f3660031901126103d7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f2a5760206040515f516020615c075f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110b95750505090565b909192602060e0826110ce6001948851611041565b0194019291016110ac565b346103d75760203660031901126103d757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061119f57505061112892500383610597565b61113282516131f0565b915f5b81518110156111915760019061117561117061115086612aa3565b61116a61115d8588612bf5565b516001600160a01b031690565b9061323f565b613254565b61117f8287612bf5565b5261118a8186612bf5565b5001611135565b6040518061078c8682611096565b8454835260019485019487945060209093019201611113565b346103d7575f3660031901126103d757602060ff5f516020615ca75f395f51905f5254166040519015158152f35b60e03660031901126103d757602435600435611201826103e7565b6044359161120e836103e7565b6064359260c435916084359160a435916001600160401b0385116103d7576112bf966112746112446112b5973690600401610c00565b95909661124f61379d565b6001600160a01b038516948490611266878d614850565b61126e6137c4565b8b6148f0565b939092604051986112848a610525565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610daf565b60e0820152614a9b565b5f5f516020615ce75f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103d75760803660031901126103d75760243560043561130e826103e7565b604435906001600160401b0382116103d757366023830112156103d75761078c92611346611359933690602481600401359101610daf565b906064359261135484610610565b6132dc565b604051918291826112db565b90602080835192838152019201905f5b8181106113825750505090565b8251845260209384019390920191600101611375565b906020610e00928181520190611365565b346103d75760203660031901126103d7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106114035761078c856113f781870382610597565b60405191829182611398565b82548452602090930192600192830192016113e0565b60e0810192916105ca9190611041565b346103d75760403660031901126103d75761078c61146860243560043561144f826103e7565b6114576131ba565b505f52600660205260405f2061323f565b60046040519161147783610546565b80546001600160a01b0316835260018101546114cc906114c3906114a661149d826103db565b60208801612fec565b6114ba60a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611419565b346103d75760203660031901126103d757600435611510816103e7565b611518613572565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916115b6906115a8610e0097959693600f60f81b865260e0602087015260e086019061155d565b90848203604086015261155d565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611365565b346103d7575f3660031901126103d7575f516020615bc75f395f51905f52541580611674575b1561163757611613614b99565b61161b614c53565b9061078c6116276133a5565b6040519384933091469186611581565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615d875f395f51905f525415611606565b9080601f830112156103d75781356116a1816105f9565b926116af6040519485610597565b81845260208085019260051b820101918383116103d75760208201905b8382106116db57505050505090565b81356001600160401b0381116103d7576020916116fd87848094880101610682565b8152019101906116cc565b60803660031901126103d7576004356001600160401b0381116103d757611733903690600401610469565b90602435906001600160401b0382116103d757366023830112156103d757816004013561175f816105f9565b9261176d6040519485610597565b8184526024602085019260051b820101903682116103d75760248101925b8284106117e35750506044359150506001600160401b0381116103d7576117b690369060040161168a565b606435926001600160401b0384116103d75761078c946117dd61077a95369060040161168a565b936133c0565b83356001600160401b0381116103d75760209161180783926024369187010161061b565b81520193019261178b565b346103d75760603660031901126103d75760043561182f816103e7565b6024359061183c826103e7565b60443561184881610610565b5f516020615d675f395f51905f52549260ff604085901c1615936001600160401b0316801590816118bd575b60011490816118b3575b1590816118aa575b50610a5b576109fb9284611898612c09565b1561436d576118a5612c2a565b61436d565b9050155f611886565b303b15915061187e565b859150611874565b346103d75760403660031901126103d7576024356004356118e4613572565b6118ec6137c4565b6118f68282614d7f565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615ce75f395f51905f525d005b346103d7575f3660031901126103d7576034546040516001600160a01b039091168152602090f35b346103d7575f3660031901126103d7576020603554604051908152f35b346103d75760403660031901126103d757602060ff6119b4602435600435611999826103e7565b5f525f516020615c875f395f51905f52845260405f2061323f565b54166040519015158152f35b801515036103d757565b346103d75760603660031901126103d7576024356004356119ea826103e7565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea6020604435611a19816119c0565b611a21613666565b835f5260058252611a67816001611a6160405f2098611a5181611a4c858060a01b038216809d614e24565b613456565b885f526006875260405f2061323f565b0161347e565b6040519015158152a3005b346103d75760403660031901126103d7576004356001600160401b0381116103d757611aa2903690600401610682565b6024356001600160401b0381116103d757611ac1903690600401610682565b90611acf8151835114612555565b5f5b82518110156100205780611b80611aea60019385612bf5565b51611af58387612bf5565b5190611aff61379d565b611b076137c4565b805f526007602052611b2082610b078160405f20614e24565b610d7d6040611b3e610d0885610cfb865f52600860205260405f2090565b610d6a611b5482516080610d23868301516103db565b50611b5e81611e2d565b611b6781611e2d565b835160208101829052908a90610d648360408101610d56565b611b886137f9565b01611ad1565b346103d7575f3660031901126103d75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611bcb5750505090565b82516001600160a01b0316845260209384019390920191600101611bbe565b346103d75760203660031901126103d7576004355f525f516020615c275f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611c505761078c85611c4481870382610597565b60405191829182611ba8565b8254845260209093019260019283019201611c2d565b346103d75760203660031901126103d7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611ca5613572565b80600155604051908152a1005b346103d75760403660031901126103d757602435600435611cd16135ec565b611cd96135ec565b611ce16137c4565b805f526007602052611cfa82610b078160405f20614e24565b611d0482826147ab565b5f516020615b275f395f51905f525f80a35f5f516020615ce75f395f51905f525d005b60405190611d36602083610597565b5f8252565b346103d7575f3660031901126103d75761078c604051611d5c604082610597565b60058152640352e302e360dc1b602082015260405191829160208352602083019061155d565b346103d75760203660031901126103d7576004355f526004602052600160405f20015460018101809111610fe557602090604051908152f35b346103d757611dc936610842565b90611dd78151835114612bdf565b5f5b82518110156100205780611e12611df260019385612bf5565b51838060a01b03611e038488612bf5565b5116906107f56107f082612bc1565b5001611dd9565b634e487b7160e01b5f52602160045260245ffd5b600a1115611e3757565b611e19565b90600a821015611e375752565b6020815260606040611eae60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061155d565b93611ec0602082015183860190611e3c565b015191015290565b346103d75760403660031901126103d757600435602435905f60408051611eee81610561565b611ef661349b565b815282602082015201525f52600860205260405f20905f5260205261078c60405f20600760405191611f2783610561565b611f3081613033565b8352611f4660ff6006830154166020850161312b565b0154604082015260405191829182611e49565b346103d75760203660031901126103d7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611f9b81610610565b611fa3613572565b16611faf8115156134cb565b8060ff195f516020615b675f395f51905f525416175f516020615b675f395f51905f5255604051908152a1005b346103d75760403660031901126103d757600435602435611ffc816119c0565b612004613666565b612019825f52600360205260405f2054151590565b156120705760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f526004825261206581600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103d7575f3660031901126103d7576033546040516001600160a01b039091168152602090f35b346103d75760203660031901126103d7576004356120c8816119c0565b6120d0613666565b1561212e576120dd61379d565b600160ff195f516020615ca75f395f51905f525416175f516020615ca75f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615ca75f395f51905f525460ff8116156121855760ff19165f516020615ca75f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103d7575f3660031901126103d757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106121f25761078c856113f781870382610597565b82548452602090930192600192830192016121db565b346103d75760203660031901126103d757600435612225816103e7565b61222d613572565b6001600160a01b03166122418115156124fe565b603480546001600160a01b031916821790555f516020615c475f395f51905f525f80a2005b346103d75760403660031901126103d757610020602435600435612289826103e7565b6109416107f082612bc1565b346103d75760203660031901126103d7576004355f526004602052600260405f20015460018101809111610fe557602090604051908152f35b346103d75760203660031901126103d7576004356122eb816103e7565b6122f3613572565b5f546001600160a01b03168061236e57506123186001600160a01b03821615156134e1565b5f80546001600160a01b0319166001600160a01b03928316179081905561234790612342906103db565b6103db565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103d75760803660031901126103d7576100206004356024356123a3816119c0565b604435906123b0826103e7565b606435926123bd846103e7565b6123c56136e0565b614636565b346103d7575f3660031901126103d7576032546040516001600160a01b039091168152602090f35b346103d75760403660031901126103d757602435600435612412826103e7565b61241a613572565b805f5260056020525f6004612455604083209461244481611a4c60018060a01b0382168099615500565b84845260066020526040842061323f565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103d7575f3660031901126103d7576124aa6158f2565b6124b2615949565b6040519060208201925f516020615da75f395f51905f528452604083015260608201524660808201523060a082015260a081526124f060c082610597565b519020604051908152602090f35b1561250557565b638219abe360e01b5f5260045ffd5b603480546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b1561255c57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125a15760051b8101359060fe19813603018212156103d7570190565b61256b565b35610e00816103e7565b903590601e19813603018212156103d757018035906001600160401b0382116103d7576020019181360383136103d757565b91908110156125a15760e0020190565b908160209103126103d75751610e00816119c0565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361267897939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c0860191612607565b948035612684816103e7565b6001600160a01b031660e085015260208101356126a0816103e7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff60808201356126d581610610565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d15612723573d9061270a82610d94565b916127186040519384610597565b82523d5f602084013e565b606090565b604090610e0093928152816020820152019061155d565b909193929361274f858414612555565b5f945b83861061276157505050509050565b61276c86858561257f565b3560206127848161277e8a898961257f565b016125a6565b612794606061277e8b8a8a61257f565b9261280a8a86888a8c6127ee60806127ad87848661257f565b0135956127e66127dc8260a06127c482888a61257f565b01359560c06127d483838b61257f565b01359761257f565b60e08101906125b0565b9690956125e2565b946040519a8b998a996326ae802b60e11b8b5260048b01612627565b03815f305af19081612851575b5061284657856128256126f9565b60405163f495148b60e01b81529182916128429160048401612728565b0390fd5b600190950194612752565b6128719060203d8111612876575b6128698183610597565b8101906125f2565b612817565b503d61285f565b9061296293929161288c61379d565b6128946137c4565b8035926128a9845f52600560205260405f2090565b916128d96128c760408301946128c1612342876125a6565b9061380b565b6128d3612342866125a6565b90612ab1565b6128e4343415612ad9565b612978612904865f526004602052600260405f2001600181540180915590565b96612917602084013589819a8214612af7565b612923612342866125a6565b93606084019688612933896125a6565b966129708c60808901359e8f60a08b019b61294e8d8d6125b0565b929091604051978896602088019a8b612b15565b03601f198101835282610597565b51902061384e565b61298b87612985856125a6565b87613a59565b91909260018461299a81611e2d565b14612a67575b506129aa83611e2d565b60018303612a0457505050906129d66129d05f516020615b875f395f51905f52936125a6565b916125a6565b6040516001600160a01b039092169582916129f391429184612ba3565b0390a45b6129ff6137f9565b600190565b612a418394612a3c612a5f947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a4795614180565b6125a6565b926125a6565b9260405193849360018060a01b031698429185612b79565b0390a46129f7565b612a9c91935088612a77866125a6565b91612a95612a8e612a878a6125a6565b92886125b0565b3691610daf565b9289613b86565b915f6129a0565b5f52600660205260405f2090565b15612ab95750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612ae15750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b00575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0097959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c08201520191612607565b6105ca93606092969593608083019760018060a01b03168352602083015260408201520190611e3c565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615c875f395f51905f52602052600160405f20015490565b15612be657565b63031206d560e51b5f5260045ffd5b80518210156125a15760209160051b010190565b5f516020615d675f395f51905f5280546001600160401b0319166001179055565b5f516020615d675f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612d5a576001600160a01b03841615612d4b57612c6e615082565b612c826001600160a01b03841615156124fe565b6001600160a01b03811692612c988415156124fe565b60ff831615612d3c57612d0592612cf3612cf892612cb4615082565b612cbc615082565b612cc4615082565b60ff195f516020615ca75f395f51905f5254165f516020615ca75f395f51905f5255612cee615082565b6150ad565b6150bc565b612d00615275565b612514565b5f516020615c475f395f51905f525f80a2612d1f43603555565b612d298184614526565b81612d3357505050565b6105ca9261473f565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612d715750565b6373a1390160e11b5f5260045260245ffd5b959394612e05919597939297612d9761379d565b612ddc6001600160a01b038816612dae818b614850565b612db66137c4565b612dc661234261234260e46125a6565b811490612dd661234260e46125a6565b91612f60565b612dfd612ded6123426101046125a6565b6001600160a01b038b1614612f8d565b8387896148f0565b9161012435612e3881612e2186612e1c8787612fb7565b612fb7565b811015612e3287612e1c8888612fb7565b90612fc4565b612e466123426032546103db565b90612e526101046125a6565b906101443592612e63610164612fe2565b92610184356101a43591833b156103d757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612f5b57612f226112b598612f2b936129f79c612f41575b50612f19612f046101046125a6565b91612f0d6105ba565b9c8d5260208d01612fec565b60408b01612fec565b60608901612fec565b608087015260a086015260c08501523691610daf565b80612f4f5f612f5593610597565b80610a85565b5f612ef5565b6126ee565b15612f69575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612f9457565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610fe557565b15612fcd575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e0081610610565b6001600160a01b039091169052565b90600182811c92168015613029575b602083101461301557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161300a565b906040516130408161057c565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f9161309882612ffb565b808552916001811690811561310457506001146130c5575b505060a092916130c1910384610597565b0152565b5f908152602081209092505b8183106130e85750508101602001816130c16130b0565b60209193508060019154838589010152019101909184926130d1565b60ff191660208681019190915292151560051b850190920192508391506130c190506130b0565b600a821015611e375752565b9060405161314481610561565b60406007829461315381613033565b845261316960ff6006830154166020860161312b565b0154910152565b156131785750565b60405162461bcd60e51b81526020600482015290819061284290602483019061155d565b156131a5575050565b637a88099160e11b5f5260045260245260445ffd5b604051906131c782610546565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906131fa826105f9565b6132076040519182610597565b8281528092613218601f19916105f9565b01905f5b82811061322857505050565b6020906132336131ba565b8282850101520161321c565b9060018060a01b03165f5260205260405f2090565b9060405161326181610546565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132ac906132a3906114ba565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103d75751610e00816103e7565b5f5490949392906001600160a01b0381161561339657604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff9061334590608487019061155d565b931691015203925af18015612f5b576105ca925f91613367575b5080946134f7565b613389915060203d60201161338f575b6133818183610597565b8101906132c7565b5f61335f565b503d613377565b6315aeca0d60e11b5f5260045ffd5b604051906133b4602083610597565b5f808352366020840137565b919294939093848451148061344c575b80613442575b6133df90612555565b5f5b85811015613436578060051b8401359060be19853603018212156103d75761342f60019261340f8389612bf5565b5161341a848c612bf5565b51906134268589612bf5565b5192890161287d565b50016133e1565b50945050505050600190565b50815185146133d6565b50848651146133d0565b1561345e5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b604051906134a88261057c565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134d257565b63f0f15b9160e01b5f5260045ffd5b156134e857565b6302bfb33d60e51b5f5260045ffd5b905f6105ca93926123c56136e0565b9161351f9183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361353757565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b5f516020615d075f395f51905f525f525f516020615c875f395f51905f5260205260ff6135bf337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61323f565b5416156135c857565b63e2517d3f60e01b5f52336004525f516020615d075f395f51905f5260245260445ffd5b5f516020615d475f395f51905f525f525f516020615c875f395f51905f5260205260ff613639337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361323f565b54161561364257565b63e2517d3f60e01b5f52336004525f516020615d475f395f51905f5260245260445ffd5b5f516020615c675f395f51905f525f525f516020615c875f395f51905f5260205260ff6136b3337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661323f565b5416156136bc57565b63e2517d3f60e01b5f52336004525f516020615c675f395f51905f5260245260445ffd5b5f516020615cc75f395f51905f525f525f516020615c875f395f51905f5260205260ff61372d337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61323f565b54161561373657565b63e2517d3f60e01b5f52336004525f516020615cc75f395f51905f5260245260445ffd5b805f525f516020615c875f395f51905f5260205260ff61377d3360405f2061323f565b5416156137875750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615ca75f395f51905f5254166137b557565b63d93c066560e01b5f5260045ffd5b5f516020615ce75f395f51905f525c6137ea5760015f516020615ce75f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615ce75f395f51905f525d565b610e00916001600160a01b031690614e24565b1561382557565b63b3f07a3960e01b5f5260045ffd5b1561383c5750565b631aebd17960e11b5f5260045260245ffd5b93929381519183518314806139b8575b6138679061381e565b6138886138825f516020615b675f395f51905f525460ff1690565b60ff1690565b956138968488811015613834565b5f945f935f5b8681106138b757505050505050506105ca9192811015613834565b6139146138e3836138c66155a3565b6042916040519161190160f01b8352600283015260228201522090565b6138f76138f08489612bf5565b5160ff1690565b6139018487612bf5565b519061390d8589612bf5565b5192614e37565b6001600160a01b038181169088161080613945575b613937575b5060010161389c565b60019889019890965061392e565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615c875f395f51905f526020526139b36139ac827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561323f565b61323f565b5460ff1690565b613929565b508551831461385e565b156139c957565b6330d45fb160e01b5f5260045ffd5b908160209103126103d75751600a8110156103d75790565b6001600160a01b039091168152602081019190915260400190565b9150613a46611170606092613a355f95613a3060325460018060a01b031615156139c2565b612aa3565b6001600160a01b039091169061323f565b0151613a5157600191565b506002905f90565b9190915f92613aa66060613a9f611170613a8f613a7a6123426032546103db565b96613a306001600160a01b03891615156139c2565b6001600160a01b0386169061323f565b0151151590565b613b2d5791602091613ad095935f604051809881958294633f4bdec560e01b8452600484016139f0565b03925af1928315612f5b575f93613afc575b50600183613aef81611e2d565b03613af657565b60019150565b613b1f91935060203d602011613b26575b613b178183610597565b8101906139d8565b915f613ae2565b503d613b0d565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610e009291019061155d565b908160209103126103d7575190565b92919382613b9385612aa3565b95613bb66001613bab818060a01b038716809a61323f565b015460a01c60ff1690565b92601481511180613eb8575b613bfb575b50613bd193614e4f565b92613bdb84611e2d565b60018414613bea575b50505090565b613bf392614f0e565b5f8080613be4565b9150613c3393602083015160601c6020613c196123426033546103db565b916040518098819262483e5560e91b8352600483016112db565b0381845afa938415612f5b578888958b985f91613e99575b50613c595750505093613bc7565b92809585988896936002613c776001985f52600460205260405f2090565b01549687921480159889613e5e575b5060209596505f14613e53578c613cc584925b613ca76123426033546103db565b95604051998a9889978896631eeed51360e01b885260048801613b38565b03925af1908115612f5b575f91613e34575b50613e24579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613d32898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613e1c575b50613d4e575b92613bd193614e4f565b9050613d5e6123426033546103db565b92604051916370a0823160e01b835260208380613d7e88600483016112db565b03818b5afa8015612f5b5786935f91613dfd575b5080613da2575b50909350613d44565b6020919350613dc6956040519687928392632770a7eb60e21b8452600484016139f0565b03815f8b5af1918215612f5b57613bd194869315613d9957613df69060203d602011612876576128698183610597565b505f613d99565b613e16915060203d602011610f2357610f148183610597565b5f613d92565b90505f613d3e565b505050505090916129ff92614f0e565b613e4d915060203d602011612876576128698183610597565b5f613cd7565b8c613cc55f92613c99565b939250945094613e6d93614e4f565b613e7681611e2d565b60018103613e8b5750848884848a948a613c86565b995050505050505050505090565b613eb2915060203d602011612876576128698183610597565b5f613c4b565b506033546001600160a01b0390613ed290612342906103db565b161515613bc2565b15613ee25750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103d75760405190613f0d8261057c565b819380358352602081013560208401526040810135613f2b816103e7565b6040840152613f3c606082016107b6565b60608401526080818101359084015260a0810135916001600160401b0383116103d75760a092613f6c9201610de5565b910152565b818110613f7c575050565b5f8155600101613f71565b90601f8211613f94575050565b6105ca915f516020615b475f395f51905f525f5260205f20906020601f840160051c83019310613fcc575b601f0160051c0190613f71565b9091508190613fbf565b9190601f8111613fe557505050565b6105ca925f5260205f20906020601f840160051c83019310613fcc57601f0160051c0190613f71565b8160011b915f199060031b1c19161790565b90600a811015611e375760ff80198354169116179055565b815180518255602081015160018301556040810151919291614066906001600160a01b031660028501612536565b6060810151614081906001600160a01b031660038501612536565b6080810151600484015560a00151805160058401916001600160401b038211610541576140b8826140b28554612ffb565b85613fd6565b602090601f8311600114614110578260079593604095936140e1935f92614105575b505061400e565b90555b6140fe60208201516140f581611e2d565b60068601614020565b0151910155565b015190505f806140da565b90601f19831691614124855f5260205f2090565b925f5b8181106141685750926001928592600798966040989610614150575b505050811b0190556140e4565b01515f1960f88460031b161c191690555f8080614143565b92936020600181928786015181550195019301614127565b91608061421e61420f600293613a308761420a6141ba61351f99833595865f5260076020526141bf60405f20602087013594858092615810565b613eda565b1561422b576141f36141d360015442612fb7565b915b6141e86141e06105cc565b963690613ef4565b86526020860161312b565b6040840152610cfb855f52600860205260405f2090565b614038565b61116a612342604088016125a6565b9301359201918254612fb7565b6141f35f916141d5565b90614240825f614f57565b91826142495750565b5f80525f516020615c275f395f51905f526020526001600160a01b0316614290817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615810565b156142985750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916142bb8382614f57565b92836142c5575050565b815f525f516020615c275f395f51905f526020526142f060405f209160018060a01b03168092615810565b156142f9575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161431b8382614ffa565b9283614325575050565b815f525f516020615c275f395f51905f5260205261435060405f209160018060a01b03168092615500565b15614359575050565b62a95f1b60e31b5f5260045260245260445ffd5b9190614377615082565b61438b6001600160a01b03841615156124fe565b6001600160a01b038116926143a18415156124fe565b60ff831615612d3c5761442e9261441461441a926143bd615082565b6143c5615082565b6143cd615082565b60ff195f516020615ca75f395f51905f5254165f516020615ca75f395f51905f52556143f7615082565b6143ff615082565b614407615082565b61440f615082565b614235565b506150bc565b614422615082565b62015180600155612514565b5f516020615c475f395f51905f525f80a26105ca43603555565b600360606105ca938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156144875750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916144c060018060a01b0382511685612536565b602081015161450b9060018601906144e1906001600160a01b031682612536565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061347e565b6080810151600285015560a081015160038501550151910155565b906001916145358115156134e1565b6145d8838060a01b0383169261454c8415156134e1565b614555856134e1565b61455e836157a7565b6145fa575b614587816145828161457d875f52600560205260405f2090565b614fe7565b61447f565b6145d36145926105ea565b9161459d8184612fec565b6145aa8760208501612fec565b86151560408401525f60608401525f60808401525f60a08401525f60c08401526139a785612aa3565b6144a7565b60405183151581525f516020615d275f395f51905f529080602081015b0390a4565b6146316146056105db565b8481525f60208201525f60408201525f606082015261462c855f52600460205260405f2090565b614448565b614563565b906145f55f516020615d275f395f51905f52919493946146578415156134e1565b6001600160a01b0386169461077a906146718715156134e1565b6001600160a01b038116976145d39061468b8a15156134e1565b614694886157a7565b6146fb575b6146b3816145828161457d8c5f52600560205260405f2090565b6146d26146be6105ea565b936146c98386612fec565b60208501612fec565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526139a788612aa3565b61472d6147066105db565b8981525f60208201525f60408201525f606082015261462c8a5f52600460205260405f2090565b614699565b91908203918211610fe557565b90614754915f52600660205260405f2061323f565b600181015460a01c60ff1615614776576003018054918201809211610fe55755565b6004018054918203918211610fe55755565b156147905750565b63290cd55f60e01b5f52600a811015611e375760045260245ffd5b906147b591614d7f565b60018060a01b036060820151168151915f516020615b875f395f51905f5260408201926147fd610b62855160018060a01b03169660808601978489519160a089015193613b86565b8251602090930151935194516040516001600160a01b03909616959182916145f59142919084612ba3565b156148305750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f52600560205261487481611a4c60405f2060018060a01b03831690614e24565b825f52600460205260ff600360405f200154166148ac5760ff60016148a0836139a76105ca9697612aa3565b015460a81c1615614828565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103d7578051916040602083015192015190565b156148e157565b6358e8878560e01b5f5260045ffd5b826060916149899795969361491a61117061490a84612aa3565b6001600160a01b0384169061323f565b61492a6109b76040830151151590565b614a2b575b5061493e6123426032546103db565b916149536001600160a01b03841615156139c2565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612f5b575f955f905f936149ed575b509082916105ca949396819885151595866149e2575b5050846149d7575b5050826149cc575b50506148da565b101590505f806149c5565b101592505f806149bd565b101594505f806149b5565b9050614a189196506105ca93925060603d606011614a24575b614a108183610597565b8101906148bf565b9196929391929161499f565b503d614a06565b60c0614a3d9101518480821015612fc4565b5f61492f565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e091614a96919086019061155d565b930152565b614ab981515f526004602052600160405f2001600181540180915590565b614ac38251612aa3565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708614b086123426001614b01602088019561116a61234288516103db565b01546103db565b938051906145f5614b1985516103db565b916040810190614b2982516103db565b90614b526080820196875160a0840194855198614b4c60c087019a8b5190612fb7565b936152d0565b614b68614b61825199516103db565b93516103db565b9460e0614b7860608401516103db565b9751935191519201519260405197889760018060a01b03169c429689614a43565b604051905f825f516020615b475f395f51905f525491614bb883612ffb565b8083529260018116908115614c345750600114614bdc575b6105ca92500383610597565b505f516020615b475f395f51905f525f90815290915f516020615be75f395f51905f525b818310614c185750509060206105ca92820101614bd0565b6020919350806001915483858901015201910190918492614c00565b602092506105ca94915060ff191682840152151560051b820101614bd0565b604051905f825f516020615ba75f395f51905f525491614c7283612ffb565b8083529260018116908115614c345750600114614c95576105ca92500383610597565b505f516020615ba75f395f51905f525f90815290915f516020615dc75f395f51905f525b818310614cd15750509060206105ca92820101614bd0565b6020919350806001915483858901015201910190918492614cb9565b60075f9182815582600182015582600282015582600382015582600482015560058101614d1a8154612ffb565b9081614d2d575b50508260068201550155565b601f8211600114614d4457849055505b5f80614d21565b614d6a614d7a926001601f614d5c855f5260205f2090565b920160051c82019101613f71565b5f81815260208120918190559055565b614d3d565b9190614d8961349b565b50825f526007602052614d9f8160405f20615500565b15614e1257614e0d6105ca91845f52600860205260405f20815f52602052610cfb614dcc60405f20613033565b95614de9614dd982612aa3565b61116a61234260408b01516103db565b614dfd600260808a01519201918254614732565b90555f52600860205260405f2090565b614ced565b637f11bea960e01b5f5260045260245ffd5b6001915f520160205260405f2054151590565b91610e009391614e46936155f7565b90929192615679565b6001600160a01b031692919060018414614ede578115614ed557614e9e9215614ea95760405163a9059cbb60e01b602082015291614e9691839161296291602484016139f0565b6005926156f5565b15610e005750600190565b6040516340c10f1960e01b602082015291614ecd91839161296291602484016139f0565b6006926156f5565b50505050600190565b90614f0093506109b79250614ef1611d27565b916001600160a01b031661573e565b614f0957600190565b600590565b90614f23915f52600660205260405f2061323f565b600181015460a01c60ff1615614f45576003018054918203918211610fe55755565b6004018054918201809211610fe55755565b805f525f516020615c875f395f51905f5260205260ff614f7a8360405f2061323f565b5416614fe157805f525f516020615c875f395f51905f52602052614fa18260405f2061323f565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610e00916001600160a01b031690615810565b805f525f516020615c875f395f51905f5260205260ff61501d8360405f2061323f565b541615614fe157805f525f516020615c875f395f51905f526020526150458260405f2061323f565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615d675f395f51905f525460401c161561509e57565b631afcd79f60e31b5f5260045ffd5b6150b9906143ff615082565b50565b906150c5615082565b6150d360ff831615156134cb565b6040918251926150e38185610597565b60098452682b30b634b230ba37b960b91b602085015261510581519182610597565b60058152640312e302e360dc1b602082015261511f615082565b615127615082565b83516001600160401b03811161054157615157816151525f516020615b475f395f51905f5254612ffb565b613f87565b6020601f82116001146151e657816151949392615180926105ca97985f9261410557505061400e565b5f516020615b475f395f51905f525561597b565b6151a95f5f516020615bc75f395f51905f5255565b6151be5f5f516020615d875f395f51905f5255565b60ff1660ff195f516020615b675f395f51905f525416175f516020615b675f395f51905f5255565b5f516020615b475f395f51905f525f52601f198216955f516020615be75f395f51905f52965f5b81811061525d575096600192849261519496956105ca999a10615245575b505050811b015f516020615b475f395f51905f525561597b565b01515f1960f88460031b161c191690555f808061522b565b8383015189556001909801976020938401930161520d565b61527d615082565b62015180600155565b1561529057505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156152c157565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361533e57506105ca94506153066152f88286612fb7565b34143490612e328488612fb7565b80615312575b5061473f565b615333615338916153246034546103db565b9061532d611d27565b9161573e565b6152ba565b5f61530c565b9061534a343415612ad9565b61535f6153578287612fb7565b308489615857565b806153f3575b5061537b6109b76001613bab866139a787612aa3565b61538b575b506105ca935061473f565b604051632770a7eb60e21b8152602081806153aa8830600484016139f0565b03815f885af1918215612f5b576105ca966153ce9387935f916153d4575b50615286565b5f615380565b6153ed915060203d602011612876576128698183610597565b5f6153c8565b61540b906154056123426034546103db565b876158a0565b5f615365565b90813b1561548f575f516020615c075f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2805115615477576150b9916158d5565b50503461548057565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125a1575f5260205f2001905f90565b805480156154ec575f1901906154db82826154b0565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f1461559b575f198401848111610fe55783545f19810194908511610fe5575f95858361554e97610cfb9503615554575b5050506154c5565b55600190565b61558461557e9161557561556b61559295886154b0565b90549060031b1c90565b928391876154b0565b90613506565b85905f5260205260405f2090565b555f8080615546565b505050505f90565b6155ab6158f2565b6155b3615949565b6040519060208201925f516020615da75f395f51905f528452604083015260608201524660808201523060a082015260a081526155f160c082610597565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615664579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612f5b575f516001600160a01b0381161561565a57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611e3757565b6156828161566f565b8061568b575050565b6156948161566f565b600181036156ab5763f645eedf60e01b5f5260045ffd5b6156b48161566f565b600281036156cf575063fce698f760e01b5f5260045260245ffd5b806156db60039261566f565b146156e35750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f16157196126f9565b9015614fe157805161573557503b1561573157600190565b5f90565b60209150015190565b8147106157985782516001600160a01b03909116925f9182916020018486620186a0f19061576a6126f9565b9115615791571561577d575b5050600190565b805161573557503b15615731575f80615776565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461580b57600254600160401b811015610541576157f46157de82600185940160025560026154b0565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61581a8282614e24565b614fe157805490600160401b82101561054157826158426157de8460018096018555846154b0565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105ca9161589b608483610597565b615a70565b61589b6105ca93926158c760405194859263a9059cbb60e01b6020850152602484016139f0565b03601f198101845283610597565b5f80610e0093602081519101845af46158ec6126f9565b91615ac8565b6158fa614b99565b805190811561590a576020012090565b50505f516020615bc75f395f51905f525480156159245790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615951614c53565b8051908115615961576020012090565b50505f516020615d875f395f51905f525480156159245790565b80519091906001600160401b038111610541576159bc816159a95f516020615ba75f395f51905f5254612ffb565b5f516020615ba75f395f51905f52613fd6565b602092601f82116001146159ef576159de929382915f9261410557505061400e565b5f516020615ba75f395f51905f5255565b5f516020615ba75f395f51905f525f52601f198216935f516020615dc75f395f51905f52915f5b868110615a585750836001959610615a40575b505050811b015f516020615ba75f395f51905f5255565b01515f1960f88460031b161c191690555f8080615a29565b91926020600181928685015181550194019201615a16565b905f602091828151910182855af1156126ee575f513d615abf57506001600160a01b0381163b155b615a9f5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615a98565b90615aec5750805115615add57805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b1d575b615afd575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615af556feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a26469706673582212204646ca45fdf6dc0ae23dafd2c733d4ade835b388d0da3c3464656eb735f3409364736f6c634300081c0033"

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

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridge *BSCBridgeCaller) BridgeExecuter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridge.contract.Call(opts, &out, "bridgeExecuter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridge *BSCBridgeSession) BridgeExecuter() (common.Address, error) {
	return _BSCBridge.Contract.BridgeExecuter(&_BSCBridge.CallOpts)
}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridge *BSCBridgeCallerSession) BridgeExecuter() (common.Address, error) {
	return _BSCBridge.Contract.BridgeExecuter(&_BSCBridge.CallOpts)
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

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridge *BSCBridgeTransactor) SetBridgeExecuter(opts *bind.TransactOpts, _bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setBridgeExecuter", _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridge *BSCBridgeSession) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeExecuter(&_BSCBridge.TransactOpts, _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetBridgeExecuter(&_BSCBridge.TransactOpts, _bridgeExecuter)
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

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridge *BSCBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridge *BSCBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetTokenPause(&_BSCBridge.TransactOpts, remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridge *BSCBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridge.Contract.SetTokenPause(&_BSCBridge.TransactOpts, remoteChainID, token, pause)
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

// BSCBridgeBridgeExecuterCallFailedIterator is returned from FilterBridgeExecuterCallFailed and is used to iterate over the raw logs and unpacked data for BridgeExecuterCallFailed events raised by the BSCBridge contract.
type BSCBridgeBridgeExecuterCallFailedIterator struct {
	Event *BSCBridgeBridgeExecuterCallFailed // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeExecuterCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeExecuterCallFailed)
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
		it.Event = new(BSCBridgeBridgeExecuterCallFailed)
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
func (it *BSCBridgeBridgeExecuterCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeExecuterCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeExecuterCallFailed represents a BridgeExecuterCallFailed event raised by the BSCBridge contract.
type BSCBridgeBridgeExecuterCallFailed struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	Reason      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuterCallFailed is a free log retrieval operation binding the contract event 0x9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be1.
//
// Solidity: event BridgeExecuterCallFailed(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, bytes reason)
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeExecuterCallFailed(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeBridgeExecuterCallFailedIterator, error) {

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

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeExecuterCallFailedIterator{contract: _BSCBridge.contract, event: "BridgeExecuterCallFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterCallFailed is a free log subscription operation binding the contract event 0x9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be1.
//
// Solidity: event BridgeExecuterCallFailed(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, bytes reason)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeExecuterCallFailed(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeExecuterCallFailed, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeExecuterCallFailed)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
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

// ParseBridgeExecuterCallFailed is a log parse operation binding the contract event 0x9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be1.
//
// Solidity: event BridgeExecuterCallFailed(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, bytes reason)
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeExecuterCallFailed(log types.Log) (*BSCBridgeBridgeExecuterCallFailed, error) {
	event := new(BSCBridgeBridgeExecuterCallFailed)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeBridgeExecuterSetIterator is returned from FilterBridgeExecuterSet and is used to iterate over the raw logs and unpacked data for BridgeExecuterSet events raised by the BSCBridge contract.
type BSCBridgeBridgeExecuterSetIterator struct {
	Event *BSCBridgeBridgeExecuterSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeBridgeExecuterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeBridgeExecuterSet)
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
		it.Event = new(BSCBridgeBridgeExecuterSet)
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
func (it *BSCBridgeBridgeExecuterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeBridgeExecuterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeBridgeExecuterSet represents a BridgeExecuterSet event raised by the BSCBridge contract.
type BSCBridgeBridgeExecuterSet struct {
	BridgeExecuter common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuterSet is a free log retrieval operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_BSCBridge *BSCBridgeFilterer) FilterBridgeExecuterSet(opts *bind.FilterOpts, bridgeExecuter []common.Address) (*BSCBridgeBridgeExecuterSetIterator, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _BSCBridge.contract.FilterLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeBridgeExecuterSetIterator{contract: _BSCBridge.contract, event: "BridgeExecuterSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterSet is a free log subscription operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_BSCBridge *BSCBridgeFilterer) WatchBridgeExecuterSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeBridgeExecuterSet, bridgeExecuter []common.Address) (event.Subscription, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _BSCBridge.contract.WatchLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeBridgeExecuterSet)
				if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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

// ParseBridgeExecuterSet is a log parse operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_BSCBridge *BSCBridgeFilterer) ParseBridgeExecuterSet(log types.Log) (*BSCBridgeBridgeExecuterSet, error) {
	event := new(BSCBridgeBridgeExecuterSet)
	if err := _BSCBridge.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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
