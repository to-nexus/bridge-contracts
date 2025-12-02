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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecuter\",\"outputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossSupplyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"bscChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeCrossBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"_bridgeExecuter\",\"type\":\"address\"}],\"name\":\"setBridgeExecuter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"setCrossSupplyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BridgeExecuterCallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecuter\",\"type\":\"address\"}],\"name\":\"BridgeExecuterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crossSupplyLimit\",\"type\":\"uint256\"}],\"name\":\"CrossSupplyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
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
		"b8886e1a": "setCrossSupplyLimit(uint256)",
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
	Bin: "0x60a080604052346100bb57306080525f5160206160595f395f51905f525460ff8160401c166100ac576002600160401b03196001600160401b03821601610059575b604051615f9990816100c0823960805181610edb0152f35b6001600160401b0319166001600160401b039081175f5160206160595f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b6100206135ee565b005b5f3560e01c806301ffc9a7146103b15780630b1d8d06146103ac5780631313996b146103a75780631938e0f2146103a2578063248a9ca31461039d5780632f2ff15d1461039857806336568abe14610393578063365d71e41461038e57806342cde4e81461038957806348a00ed8146103845780634be13f831461037f5780634d5d00561461037a5780634ee078ba146103755780634f1ef28614610370578063502cc5c01461036b57806352d1902d146103665780635b605f5c146103615780635c975abb1461035c5780635fd262de14610357578063670e62681461035257806369ceb2f11461034d5780637921487414610348578063814914b514610343578063815f1de61461033e57806384b0196e1461033957806388d67d6d1461033457806389232a001461032f5780638ae87c5c1461032a57806391cca3db1461032557806391cf6d3e1461032057806391d148541461031b57806394ddc8c6146103165780639f9b4f1c14610311578063a217fddf1461030c578063a3246ad314610307578063aa1bd0bc14610302578063aa20ed47146102fd578063ad3cb1cc146102f8578063ae6893f8146102f3578063b2b49e2e146102ee578063b33eb36e146102e9578063b7f3358d146102e4578063b8886e1a146102df578063b8aa837e146102da578063bc8cadd6146102d5578063bedb86fb146102d0578063cd959706146102cb578063cf56118e146102c6578063d477f05f146102c1578063d547741f146102bc578063d5717fc5146102b7578063dfcbce15146102b2578063e2af6cd7146102ad578063edbbea39146102a8578063f0702e8e146102a3578063f45096431461029e5763f698da250361000e57612531565b612491565b612469565b61241f565b61236d565b61229f565b612266565b612237565b6121d9565b612165565b61214b565b612062565b61203a565b611f93565b611f47565b611ec4565b611e33565b611d26565b611ced565b611ca6565b611c1d565b611bd1565b611b55565b611af9565b6119dd565b611935565b6118dd565b6118c0565b611898565b611830565b61170e565b611604565b6114dc565b6113ef565b611325565b6112a5565b611244565b6111cd565b6110c5565b611097565b610fb8565b610ec9565b610e3a565b610cfb565b610b92565b610b25565b610ad1565b6109b3565b610987565b61091a565b61082a565b6107f1565b6107c0565b61070f565b6104c9565b610428565b346104075760203660031901126104075760043563ffffffff60e01b811680910361040757602090637965db0b60e01b81149081156103f6575b506040519015158152f35b6301ffc9a760e01b1490505f6103eb565b5f80fd5b6001600160a01b031690565b6001600160a01b0381160361040757565b346104075760203660031901126104075760043561044581610417565b61044d61363d565b6001600160a01b031661046181151561259d565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f84011215610407578235916001600160401b038311610407576020808501948460051b01011161040757565b6040366003190112610407576004356001600160401b038111610407576104f4903690600401610499565b602435916001600160401b0383116104075736602384011215610407576004830135916001600160401b0383116104075736602460e08502860101116104075760246100209401916127de565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761057157604052565b610541565b60e081019081106001600160401b0382111761057157604052565b606081019081106001600160401b0382111761057157604052565b60c081019081106001600160401b0382111761057157604052565b601f909101601f19168101906001600160401b0382119082101761057157604052565b604051906105fa610100836105c7565b565b604051906105fa6060836105c7565b604051906105fa6080836105c7565b604051906105fa60e0836105c7565b6001600160401b0381116105715760051b60200190565b60ff81160361040757565b9080601f8301121561040757813561066281610629565b9261067060405194856105c7565b81845260208085019260051b82010192831161040757602001905b8282106106985750505090565b6020809183356106a781610640565b81520191019061068b565b9080601f830112156104075781356106c981610629565b926106d760405194856105c7565b81845260208085019260051b82010192831161040757602001905b8282106106ff5750505090565b81358152602091820191016106f2565b6080366003190112610407576004356001600160401b0381116104075760c06003198236030112610407576024356001600160401b0381116104075761075990369060040161064b565b906044356001600160401b038111610407576107799036906004016106b2565b60643591906001600160401b038311610407576107bc936107a16107aa9436906004016106b2565b9260040161291c565b60405191829182901515815260200190565b0390f35b346104075760203660031901126104075760206107de600435612c60565b604051908152f35b35906105fa82610417565b346104075760403660031901126104075761002060243560043561081482610417565b61082561082082612c60565b613825565b6142d9565b346104075760403660031901126104075760043560243561084a81610417565b336001600160a01b038216036108635761002091614339565b63334bd91960e11b5f5260045ffd5b906040600319830112610407576004356001600160401b038111610407578261089d916004016106b2565b91602435906001600160401b03821161040757806023830112156104075781600401356108c981610629565b926108d760405194856105c7565b8184526024602085019260051b82010192831161040757602401905b8282106109005750505090565b60208091833561090f81610417565b8152019101906108f3565b346104075761092836610872565b906109368151835114612c7e565b5f5b8251811015610020578061097661095160019385612c94565b51838060a01b036109628488612c94565b51169061097161082082612c60565b614339565b5001610938565b5f91031261040757565b34610407575f36600319011261040757602060ff5f516020615ce45f395f51905f525416604051908152f35b34610407576060366003190112610407576024356004356044356109d681610417565b6109de6136b7565b6109e661388f565b815f526007602052610a04836109ff8160405f20614e9f565b612ca8565b80610a0f8484614a92565b916001600160a01b031615610abd575b8151915f516020615d045f395f51905f526040820192610a6c610a5a855160018060a01b03169660808601978489519160a089015193613bac565b610a6381611d98565b60018114614398565b8251602090930151935194516040516001600160a01b0390961695918291610a979142919084612c42565b0390a45f516020615ca45f395f51905f525f80a35f5f516020615e645f395f51905f525d005b5060608101516001600160a01b0316610a1f565b34610407575f366003190112610407575f546040516001600160a01b039091168152602090f35b9181601f84011215610407578235916001600160401b038311610407576020838186019501011161040757565b6101c036600319011261040757602435600435610b4182610417565b604435610b4d81610417565b6064359060a43560843560c4356001600160401b03811161040757610b76903690600401610af8565b94909360e03660e3190112610407576107bc976107aa97612cc2565b3461040757604036600319011261040757610c7a602435600435610bb4613868565b610bbc61388f565b805f526007602052610bd5826109ff8160405f20614e9f565b610c756040610c05610c0085610bf3865f52600860205260405f2090565b905f5260205260405f2090565b613076565b610c62610c2582516080610c1b8683015161040b565b9101519087613a8d565b50610c2f81611d98565b610c3881611d98565b83516020810182905290600190610c5c83604081015b03601f1981018552846105c7565b146130af565b01518015908115610c82575b42916130db565b6143bb565b6100206138c4565b4281109150610c6e565b6001600160401b03811161057157601f01601f191660200190565b929192610cb382610c8c565b91610cc160405193846105c7565b829481845281830111610407578281602093845f960137010152565b9080601f8301121561040757816020610cf893359101610ca7565b90565b604036600319011261040757600435610d1381610417565b6024356001600160401b03811161040757610d32903690600401610cdd565b903061aaaa625c0eb760891b01148015610e0f575b610e0057610d5361363d565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610dcf575b50610d9c57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d845f395f51905f528303610dbb576100209250615341565b632a87526960e21b5f52600483905260245ffd5b610df291945060203d602011610df9575b610dea81836105c7565b810190613b82565b925f610d7b565b503d610de0565b63703e46dd60e11b5f5260045ffd5b505f516020615d845f395f51905f52546001600160a01b031661aaaa625c0eb760891b011415610d47565b3461040757606036600319011261040757602435600435604435610e5c6136b7565b815f526007602052610e75836109ff8160405f20614e9f565b4201804211610ec45760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612ee2565b34610407575f366003190112610407577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e005760206040515f516020615d845f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610f985750505090565b909192602060e082610fad6001948851610f20565b019401929101610f8b565b3461040757602036600319011261040757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061107e575050611007925003836105c7565b611011825161312f565b915f5b81518110156110705760019061105461104f61102f86612b42565b61104961103c8588612c94565b516001600160a01b031690565b9061317e565b613193565b61105e8287612c94565b526110698186612c94565b5001611014565b604051806107bc8682610f75565b8454835260019485019487945060209093019201610ff2565b34610407575f36600319011261040757602060ff5f516020615e245f395f51905f5254166040519015158152f35b60e0366003190112610407576024356004356110e082610417565b604435916110ed83610417565b6064359260c435916084359160a435916001600160401b0385116104075761119e96611153611123611194973690600401610af8565b95909661112e613868565b6001600160a01b038516948490611145878d614465565b61114d61388f565b8b61451b565b939092604051986111638a610555565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610ca7565b60e08201526146c6565b5f5f516020615e645f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b34610407576080366003190112610407576024356004356111ed82610417565b604435906001600160401b0382116104075736602383011215610407576107bc92611225611238933690602481600401359101610ca7565b906064359261123384610640565b61321b565b604051918291826111ba565b34610407575f366003190112610407576020606454604051908152f35b90602080835192838152019201905f5b81811061127e5750505090565b8251845260209384019390920191600101611271565b906020610cf8928181520190611261565b34610407576020366003190112610407576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106112ff576107bc856112f3818703826105c7565b60405191829182611294565b82548452602090930192600192830192016112dc565b60e0810192916105fa9190610f20565b34610407576040366003190112610407576107bc61136460243560043561134b82610417565b6113536130f9565b505f52600660205260405f2061317e565b60046040519161137383610576565b80546001600160a01b0316835260018101546113c8906113bf906113a26113998261040b565b60208801612f2b565b6113b660a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611315565b346104075760203660031901126104075760043561140c81610417565b61141461363d565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916114b2906114a4610cf897959693600f60f81b865260e0602087015260e0860190611459565b908482036040860152611459565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611261565b34610407575f366003190112610407575f516020615d445f395f51905f52541580611570575b156115335761150f6147c4565b61151761487e565b906107bc6115236132e4565b604051938493309146918661147d565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615f045f395f51905f525415611502565b9080601f8301121561040757813561159d81610629565b926115ab60405194856105c7565b81845260208085019260051b820101918383116104075760208201905b8382106115d757505050505090565b81356001600160401b038111610407576020916115f9878480948801016106b2565b8152019101906115c8565b6080366003190112610407576004356001600160401b0381116104075761162f903690600401610499565b90602435906001600160401b038211610407573660238301121561040757816004013561165b81610629565b9261166960405194856105c7565b8184526024602085019260051b820101903682116104075760248101925b8284106116df5750506044359150506001600160401b038111610407576116b2903690600401611586565b606435926001600160401b038411610407576107bc946116d96107aa953690600401611586565b936132ff565b83356001600160401b0381116104075760209161170383926024369187010161064b565b815201930192611687565b346104075760603660031901126104075760043561172b81610417565b6024359061173882610417565b60443561174481610640565b5f516020615ee45f395f51905f52549260ff604085901c1615611766565b1590565b936001600160401b031680159081611828575b600114908161181e575b159081611815575b50611806576117a6928461179d613395565b6117f957614918565b6117ac57005b6117ca5f516020615ee45f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6118016133b6565b614918565b63f92ee8a960e01b5f5260045ffd5b9050155f61178b565b303b159150611783565b859150611779565b346104075760403660031901126104075760243560043561184f61363d565b61185761388f565b6118618282614a92565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615e645f395f51905f525d005b34610407575f366003190112610407576034546040516001600160a01b039091168152602090f35b34610407575f366003190112610407576020603554604051908152f35b3461040757604036600319011261040757602060ff61191f60243560043561190482610417565b5f525f516020615e045f395f51905f52845260405f2061317e565b54166040519015158152f35b8015150361040757565b346104075760603660031901126104075760243560043561195582610417565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea60206044356119848161192b565b61198c613731565b835f52600582526119d28160016119cc60405f20986119bc816119b7858060a01b038216809d614e9f565b6133d7565b885f526006875260405f2061317e565b016133ff565b6040519015158152a3005b34610407576040366003190112610407576004356001600160401b03811161040757611a0d9036906004016106b2565b6024356001600160401b03811161040757611a2c9036906004016106b2565b90611a3a81518351146125f4565b5f5b82518110156100205780611aeb611a5560019385612c94565b51611a608387612c94565b5190611a6a613868565b611a7261388f565b805f526007602052611a8b826109ff8160405f20614e9f565b610c756040611aa9610c0085610bf3865f52600860205260405f2090565b610c62611abf82516080610c1b8683015161040b565b50611ac981611d98565b611ad281611d98565b835160208101829052908a90610c5c8360408101610c4e565b611af36138c4565b01611a3c565b34610407575f3660031901126104075760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611b365750505090565b82516001600160a01b0316845260209384019390920191600101611b29565b34610407576020366003190112610407576004355f525f516020615da45f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611bbb576107bc85611baf818703826105c7565b60405191829182611b13565b8254845260209093019260019283019201611b98565b34610407576020366003190112610407577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611c1061363d565b80600155604051908152a1005b3461040757604036600319011261040757602435600435611c3c6136b7565b611c446136b7565b611c4c61388f565b805f526007602052611c65826109ff8160405f20614e9f565b611c6f82826143bb565b5f516020615ca45f395f51905f525f80a35f5f516020615e645f395f51905f525d005b60405190611ca16020836105c7565b5f8252565b34610407575f366003190112610407576107bc604051611cc76040826105c7565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611459565b34610407576020366003190112610407576004355f526004602052600160405f20015460018101809111610ec457602090604051908152f35b3461040757611d3436610872565b90611d428151835114612c7e565b5f5b82518110156100205780611d7d611d5d60019385612c94565b51838060a01b03611d6e8488612c94565b51169061082561082082612c60565b5001611d44565b634e487b7160e01b5f52602160045260245ffd5b600a1115611da257565b611d84565b90600a821015611da25752565b6020815260606040611e1960a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611459565b93611e2b602082015183860190611da7565b015191015290565b3461040757604036600319011261040757600435602435905f60408051611e5981610591565b611e6161341c565b815282602082015201525f52600860205260405f20905f526020526107bc60405f20600760405191611e9283610591565b611e9b81612f72565b8352611eb160ff6006830154166020850161306a565b0154604082015260405191829182611db4565b34610407576020366003190112610407577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611f0681610640565b611f0e61363d565b16611f1a81151561344c565b8060ff195f516020615ce45f395f51905f525416175f516020615ce45f395f51905f5255604051908152a1005b34610407576020366003190112610407577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435611f8661363d565b80606455604051908152a1005b3461040757604036600319011261040757600435602435611fb38161192b565b611fbb613731565b611fd0825f52600360205260405f2054151590565b156120275760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f526004825261201c81600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b34610407575f366003190112610407576033546040516001600160a01b039091168152602090f35b346104075760203660031901126104075760043561207f8161192b565b612087613731565b156120e557612094613868565b600160ff195f516020615e245f395f51905f525416175f516020615e245f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615e245f395f51905f525460ff81161561213c5760ff19165f516020615e245f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b34610407575f3660031901126104075760206107de613462565b34610407575f36600319011261040757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106121c3576107bc856112f3818703826105c7565b82548452602090930192600192830192016121ac565b34610407576020366003190112610407576004356121f681610417565b6121fe61363d565b6001600160a01b031661221281151561259d565b603480546001600160a01b031916821790555f516020615dc45f395f51905f525f80a2005b346104075760403660031901126104075761002060243560043561225a82610417565b61097161082082612c60565b34610407576020366003190112610407576004355f526004602052600260405f20015460018101809111610ec457602090604051908152f35b346104075760c0366003190112610407576004356122bc81610417565b602435906122c982610417565b6044356122d581610640565b6084356064356122e482610417565b60a435925f516020615ee45f395f51905f5254956123096117628860ff9060401c1690565b966001600160401b031680159081612365575b600114908161235b575b159081612352575b50611806576117a69587612340613395565b156134825761234d6133b6565b613482565b9050155f61232e565b303b159150612326565b88915061231c565b346104075760203660031901126104075760043561238a81610417565b61239261363d565b5f546001600160a01b03168061240d57506123b76001600160a01b03821615156135ac565b5f80546001600160a01b0319166001600160a01b0392831617908190556123e6906123e19061040b565b61040b565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b34610407576080366003190112610407576100206004356024356124428161192b565b6044359061244f82610417565b6064359261245c84610417565b6124646137ab565b614d1f565b34610407575f366003190112610407576032546040516001600160a01b039091168152602090f35b34610407576040366003190112610407576024356004356124b182610417565b6124b961363d565b805f5260056020525f60046124f460408320946124e3816119b760018060a01b0382168099615634565b84845260066020526040842061317e565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b34610407575f36600319011261040757612549615a6f565b612551615ac6565b6040519060208201925f516020615f245f395f51905f528452604083015260608201524660808201523060a082015260a0815261258f60c0826105c7565b519020604051908152602090f35b156125a457565b638219abe360e01b5f5260045ffd5b603480546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b156125fb57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156126405760051b8101359060fe1981360301821215610407570190565b61260a565b35610cf881610417565b903590601e198136030182121561040757018035906001600160401b0382116104075760200191813603831361040757565b91908110156126405760e0020190565b908160209103126104075751610cf88161192b565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361271797939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c08601916126a6565b94803561272381610417565b6001600160a01b031660e0850152602081013561273f81610417565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561277481610640565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d156127c2573d906127a982610c8c565b916127b760405193846105c7565b82523d5f602084013e565b606090565b604090610cf8939281528160208201520190611459565b90919392936127ee8584146125f4565b5f945b83861061280057505050509050565b61280b86858561261e565b3560206128238161281d8a898961261e565b01612645565b612833606061281d8b8a8a61261e565b926128a98a86888a8c61288d608061284c87848661261e565b01359561288561287b8260a061286382888a61261e565b01359560c061287383838b61261e565b01359761261e565b60e081019061264f565b969095612681565b946040519a8b998a996326ae802b60e11b8b5260048b016126c6565b03815f305af190816128f0575b506128e557856128c4612798565b60405163f495148b60e01b81529182916128e191600484016127c7565b0390fd5b6001909501946127f1565b6129109060203d8111612915575b61290881836105c7565b810190612691565b6128b6565b503d6128fe565b90612a0193929161292b613868565b61293361388f565b803592612948845f52600560205260405f2090565b9161297861296660408301946129606123e187612645565b906138d6565b6129726123e186612645565b90612b50565b612983343415612b78565b612a176129a3865f526004602052600260405f2001600181540180915590565b966129b6602084013589819a8214612b96565b6129c26123e186612645565b936060840196886129d289612645565b96612a0f8c60808901359e8f60a08b019b6129ed8d8d61264f565b929091604051978896602088019a8b612bb4565b03601f1981018352826105c7565b519020613919565b612a2a87612a2485612645565b87613af0565b919092600184612a3981611d98565b14612b06575b50612a4983611d98565b60018303612aa35750505090612a75612a6f5f516020615d045f395f51905f5293612645565b91612645565b6040516001600160a01b03909216958291612a9291429184612c42565b0390a45b612a9e6138c4565b600190565b612ae08394612adb612afe947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612ae6956141a6565b612645565b92612645565b9260405193849360018060a01b031698429185612c18565b0390a4612a96565b612b3b91935088612b1686612645565b91612b34612b2d612b268a612645565b928861264f565b3691610ca7565b9289613bac565b915f612a3f565b5f52600660205260405f2090565b15612b585750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612b805750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b9f575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610cf897959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c082015201916126a6565b6105fa93606092969593608083019760018060a01b03168352602083015260408201520190611da7565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615e045f395f51905f52602052600160405f20015490565b15612c8557565b63031206d560e51b5f5260045ffd5b80518210156126405760209160051b010190565b15612cb05750565b6373a1390160e11b5f5260045260245ffd5b959394612d44919597939297612cd6613868565b612d1b6001600160a01b038816612ced818b614465565b612cf561388f565b612d056123e16123e160e4612645565b811490612d156123e160e4612645565b91612e9f565b612d3c612d2c6123e1610104612645565b6001600160a01b038b1614612ecc565b83878961451b565b9161012435612d7781612d6086612d5b8787612ef6565b612ef6565b811015612d7187612d5b8888612ef6565b90612f03565b612d856123e160325461040b565b90612d91610104612645565b906101443592612da2610164612f21565b92610184356101a43591833b1561040757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612e9a57612e6161119498612e6a93612a969c612e80575b50612e58612e43610104612645565b91612e4c6105ea565b9c8d5260208d01612f2b565b60408b01612f2b565b60608901612f2b565b608087015260a086015260c08501523691610ca7565b80612e8e5f612e94936105c7565b8061097d565b5f612e34565b61278d565b15612ea8575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612ed357565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ec457565b15612f0c575050565b63943892eb60e01b5f5260045260245260445ffd5b35610cf881610640565b6001600160a01b039091169052565b90600182811c92168015612f68575b6020831014612f5457565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612f49565b90604051612f7f816105ac565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91612fd782612f3a565b80855291600181169081156130435750600114613004575b505060a092916130009103846105c7565b0152565b5f908152602081209092505b818310613027575050810160200181613000612fef565b6020919350806001915483858901015201910190918492613010565b60ff191660208681019190915292151560051b850190920192508391506130009050612fef565b600a821015611da25752565b9060405161308381610591565b60406007829461309281612f72565b84526130a860ff6006830154166020860161306a565b0154910152565b156130b75750565b60405162461bcd60e51b8152602060048201529081906128e1906024830190611459565b156130e4575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061310682610576565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061313982610629565b61314660405191826105c7565b8281528092613157601f1991610629565b01905f5b82811061316757505050565b6020906131726130f9565b8282850101520161315b565b9060018060a01b03165f5260205260405f2090565b906040516131a081610576565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916131eb906131e2906113b6565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126104075751610cf881610417565b5f5490949392906001600160a01b038116156132d557604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90613284906084870190611459565b931691015203925af18015612e9a576105fa925f916132a6575b5080946135c2565b6132c8915060203d6020116132ce575b6132c081836105c7565b810190613206565b5f61329e565b503d6132b6565b6315aeca0d60e11b5f5260045ffd5b604051906132f36020836105c7565b5f808352366020840137565b919294939093848451148061338b575b80613381575b61331e906125f4565b5f5b85811015613375578060051b8401359060be19853603018212156104075761336e60019261334e8389612c94565b51613359848c612c94565b51906133658589612c94565b5192890161291c565b5001613320565b50945050505050600190565b5081518514613315565b508486511461330f565b5f516020615ee45f395f51905f5280546001600160401b0319166001179055565b5f516020615ee45f395f51905f52805460ff60401b1916600160401b179055565b156133df5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b60405190613429826105ac565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b1561345357565b63f0f15b9160e01b5f5260045ffd5b6065545f52600660205260405f2060015f52602052600460405f20015490565b919492939094841561359d576001600160a01b0382161561358e576134a56153e3565b6134b96001600160a01b038416151561259d565b6001600160a01b038616906134cf82151561259d565b60ff81161561357f576105fa9661353661354392613531613563976134f26153e3565b6134fa6153e3565b6135026153e3565b60ff195f516020615e245f395f51905f5254165f516020615e245f395f51905f525561352c6153e3565b61540e565b61541a565b61353e6155d3565b6125b3565b5f516020615dc45f395f51905f525f80a261355d43603555565b83614c15565b8061356f575b50606555565b6135799082614e1b565b5f613569565b6338854f1160e21b5f5260045ffd5b6324cd8cef60e01b5f5260045ffd5b63110d132360e21b5f5260045ffd5b156135b357565b6302bfb33d60e51b5f5260045ffd5b905f6105fa93926124646137ab565b916135ea9183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361360257565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b5f516020615e845f395f51905f525f525f516020615e045f395f51905f5260205260ff61368a337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61317e565b54161561369357565b63e2517d3f60e01b5f52336004525f516020615e845f395f51905f5260245260445ffd5b5f516020615ec45f395f51905f525f525f516020615e045f395f51905f5260205260ff613704337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361317e565b54161561370d57565b63e2517d3f60e01b5f52336004525f516020615ec45f395f51905f5260245260445ffd5b5f516020615de45f395f51905f525f525f516020615e045f395f51905f5260205260ff61377e337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661317e565b54161561378757565b63e2517d3f60e01b5f52336004525f516020615de45f395f51905f5260245260445ffd5b5f516020615e445f395f51905f525f525f516020615e045f395f51905f5260205260ff6137f8337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61317e565b54161561380157565b63e2517d3f60e01b5f52336004525f516020615e445f395f51905f5260245260445ffd5b805f525f516020615e045f395f51905f5260205260ff6138483360405f2061317e565b5416156138525750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615e245f395f51905f52541661388057565b63d93c066560e01b5f5260045ffd5b5f516020615e645f395f51905f525c6138b55760015f516020615e645f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615e645f395f51905f525d565b610cf8916001600160a01b031690614e9f565b156138f057565b63b3f07a3960e01b5f5260045ffd5b156139075750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a83575b613932906138e9565b61395361394d5f516020615ce45f395f51905f525460ff1690565b60ff1690565b9561396184888110156138ff565b5f945f935f5b86811061398257505050505050506105fa91928110156138ff565b6139df6139ae836139916156d7565b6042916040519161190160f01b8352600283015260228201522090565b6139c26139bb8489612c94565b5160ff1690565b6139cc8487612c94565b51906139d88589612c94565b5192614eb2565b6001600160a01b038181169088161080613a10575b613a02575b50600101613967565b6001988901989096506139f9565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615e045f395f51905f52602052613a7e613a77827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561317e565b61317e565b5460ff1690565b6139f4565b5085518314613929565b9091906001600160a01b03831660011480613ae5575b613ab9575b90613ab592600192614ee2565b9091565b90613ac2613462565b818101809111610ec45760645410613ada5790613aa8565b505050600990600190565b506065548114613aa3565b9091906001600160a01b03831660011480613b38575b613b17575b90613ab5925f92614ee2565b90613b20613462565b818101809111610ec45760645410613ada5790613b0b565b506065548114613b06565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610cf892910190611459565b90816020910312610407575190565b6001600160a01b039091168152602081019190915260400190565b92919382613bb985612b42565b95613bdc6001613bd1818060a01b038716809a61317e565b015460a01c60ff1690565b92601481511180613ede575b613c21575b50613bf793614fcc565b92613c0184611d98565b60018414613c10575b50505090565b613c1992614e68565b5f8080613c0a565b9150613c5993602083015160601c6020613c3f6123e160335461040b565b916040518098819262483e5560e91b8352600483016111ba565b0381845afa938415612e9a578888958b985f91613ebf575b50613c7f5750505093613bed565b92809585988896936002613c9d6001985f52600460205260405f2090565b01549687921480159889613e84575b5060209596505f14613e79578c613ceb84925b613ccd6123e160335461040b565b95604051998a9889978896631eeed51360e01b885260048801613b43565b03925af1908115612e9a575f91613e5a575b50613e4a579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613d58898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613e42575b50613d74575b92613bf793614fcc565b9050613d846123e160335461040b565b92604051916370a0823160e01b835260208380613da488600483016111ba565b03818b5afa8015612e9a5786935f91613e23575b5080613dc8575b50909350613d6a565b6020919350613dec956040519687928392632770a7eb60e21b845260048401613b91565b03815f8b5af1918215612e9a57613bf794869315613dbf57613e1c9060203d6020116129155761290881836105c7565b505f613dbf565b613e3c915060203d602011610df957610dea81836105c7565b5f613db8565b90505f613d64565b50505050509091612a9e92614e68565b613e73915060203d6020116129155761290881836105c7565b5f613cfd565b8c613ceb5f92613cbf565b939250945094613e9393614fcc565b613e9c81611d98565b60018103613eb15750848884848a948a613cac565b995050505050505050505090565b613ed8915060203d6020116129155761290881836105c7565b5f613c71565b506033546001600160a01b0390613ef8906123e19061040b565b161515613be8565b15613f085750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126104075760405190613f33826105ac565b819380358352602081013560208401526040810135613f5181610417565b6040840152613f62606082016107e6565b60608401526080818101359084015260a0810135916001600160401b0383116104075760a092613f929201610cdd565b910152565b818110613fa2575050565b5f8155600101613f97565b90601f8211613fba575050565b6105fa915f516020615cc45f395f51905f525f5260205f20906020601f840160051c83019310613ff2575b601f0160051c0190613f97565b9091508190613fe5565b9190601f811161400b57505050565b6105fa925f5260205f20906020601f840160051c83019310613ff257601f0160051c0190613f97565b8160011b915f199060031b1c19161790565b90600a811015611da25760ff80198354169116179055565b81518051825560208101516001830155604081015191929161408c906001600160a01b0316600285016125d5565b60608101516140a7906001600160a01b0316600385016125d5565b6080810151600484015560a00151805160058401916001600160401b038211610571576140de826140d88554612f3a565b85613ffc565b602090601f831160011461413657826007959360409593614107935f9261412b575b5050614034565b90555b614124602082015161411b81611d98565b60068601614046565b0151910155565b015190505f80614100565b90601f1983169161414a855f5260205f2090565b925f5b81811061418e5750926001928592600798966040989610614176575b505050811b01905561410a565b01515f1960f88460031b161c191690555f8080614169565b9293602060018192878601518155019501930161414d565b91608061424961423a600293614235876142306141e06135ea99833595865f5260076020526141e560405f20602087013594858092615944565b613f00565b15614256576142196141f960015442612ef6565b915b61420e6142066105fc565b963690613f1a565b86526020860161306a565b6040840152610bf3855f52600860205260405f2090565b61405e565b612b42565b6110496123e160408801612645565b9301359201918254612ef6565b6142195f916141fb565b9061426b825f61508b565b91826142745750565b5f80525f516020615da45f395f51905f526020526001600160a01b03166142bb817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615944565b156142c35750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916142e6838261508b565b92836142f0575050565b815f525f516020615da45f395f51905f5260205261431b60405f209160018060a01b03168092615944565b15614324575050565b63d180cb3160e01b5f5260045260245260445ffd5b919091614346838261512e565b9283614350575050565b815f525f516020615da45f395f51905f5260205261437b60405f209160018060a01b03168092615634565b15614384575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143a05750565b63290cd55f60e01b5f52600a811015611da25760045260245ffd5b906143c591614a92565b60018060a01b036060820151168151915f516020615d045f395f51905f52604082019261440d610a5a855160018060a01b03169660808601978489519160a089015193613bac565b8251602090930151935194516040516001600160a01b03909616959182916144389142919084612c42565b0390a4565b156144455750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f526005602052614489816119b760405f2060018060a01b03831690614e9f565b825f52600460205260ff600360405f200154166144c15760ff60016144b583613a726105fa9697612b42565b015460a81c161561443d565b82636fc24b7960e11b5f5260045260245ffd5b156144db57565b6330d45fb160e01b5f5260045ffd5b90816060910312610407578051916040602083015192015190565b1561450c57565b6358e8878560e01b5f5260045ffd5b826060916145b49795969361454561104f61453584612b42565b6001600160a01b0384169061317e565b6145556117626040830151151590565b614656575b506145696123e160325461040b565b9161457e6001600160a01b03841615156144d4565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612e9a575f955f905f93614618575b509082916105fa9493968198851515958661460d575b505084614602575b5050826145f7575b5050614505565b101590505f806145f0565b101592505f806145e8565b101594505f806145e0565b90506146439196506105fa93925060603d60601161464f575b61463b81836105c7565b8101906144ea565b919692939192916145ca565b503d614631565b60c06146689101518480821015612f03565b5f61455a565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916146c19190860190611459565b930152565b6146e481515f526004602052600160405f2001600181540180915590565b6146ee8251612b42565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086147336123e1600161472c60208801956110496123e1885161040b565b015461040b565b93805190614438614744855161040b565b916040810190614754825161040b565b9061477d6080820196875160a084019485519861477760c087019a8b5190612ef6565b93615200565b61479361478c8251995161040b565b935161040b565b9460e06147a3606084015161040b565b9751935191519201519260405197889760018060a01b03169c42968961466e565b604051905f825f516020615cc45f395f51905f5254916147e383612f3a565b808352926001811690811561485f5750600114614807575b6105fa925003836105c7565b505f516020615cc45f395f51905f525f90815290915f516020615d645f395f51905f525b8183106148435750509060206105fa928201016147fb565b602091935080600191548385890101520191019091849261482b565b602092506105fa94915060ff191682840152151560051b8201016147fb565b604051905f825f516020615d245f395f51905f52549161489d83612f3a565b808352926001811690811561485f57506001146148c0576105fa925003836105c7565b505f516020615d245f395f51905f525f90815290915f516020615f445f395f51905f525b8183106148fc5750509060206105fa928201016147fb565b60209193508060019154838589010152019101909184926148e4565b91906149226153e3565b6149366001600160a01b038416151561259d565b6001600160a01b0381169261494c84151561259d565b60ff83161561357f576149d9926149bf6149c5926149686153e3565b6149706153e3565b6149786153e3565b60ff195f516020615e245f395f51905f5254165f516020615e245f395f51905f52556149a26153e3565b6149aa6153e3565b6149b26153e3565b6149ba6153e3565b614260565b5061541a565b6149cd6153e3565b620151806001556125b3565b5f516020615dc45f395f51905f525f80a26105fa43603555565b91908203918211610ec457565b60075f9182815582600182015582600282015582600382015582600482015560058101614a2d8154612f3a565b9081614a40575b50508260068201550155565b601f8211600114614a5757849055505b5f80614a34565b614a7d614a8d926001601f614a6f855f5260205f2090565b920160051c82019101613f97565b5f81815260208120918190559055565b614a50565b9190614a9c61341c565b50825f526007602052614ab28160405f20615634565b15614b2557614b206105fa91845f52600860205260405f20815f52602052610bf3614adf60405f20612f72565b95614afc614aec82612b42565b6110496123e160408b015161040b565b614b10600260808a015192019182546149f3565b90555f52600860205260405f2090565b614a00565b637f11bea960e01b5f5260045260245ffd5b600360606105fa938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b15614b765750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c0600491614baf60018060a01b03825116856125d5565b6020810151614bfa906001860190614bd0906001600160a01b0316826125d5565b6040830151815460ff60a01b191690151560a01b60ff60a01b1617815560608301511515906133ff565b6080810151600285015560a081015160038501550151910155565b600190614c238115156135ac565b614c2c826135ac565b614cc4828060a01b03841693614c438515156135ac565b614c4c836158db565b614ce3575b614c7584614c7081614c6b875f52600560205260405f2090565b61511b565b614b6e565b614c94614c8061061a565b91614c8b8684612f2b565b60208301612f2b565b5f60408201525f60608201525f60808201525f60a08201525f60c0820152614cbf84613a7285612b42565b614b96565b6040515f81525f516020615ea45f395f51905f52908060208101614438565b614d1a614cee61060b565b8481525f60208201525f60408201525f6060820152614d15855f52600460205260405f2090565b614b37565b614c51565b906144385f516020615ea45f395f51905f5291949394614d408415156135ac565b6001600160a01b038616946107aa90614d5a8715156135ac565b6001600160a01b03811697614cbf90614d748a15156135ac565b614d7d886158db565b614de4575b614d9c81614c7081614c6b8c5f52600560205260405f2090565b614dbb614da761061a565b93614db28386612f2b565b60208501612f2b565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a7288612b42565b614e16614def61060b565b8981525f60208201525f60408201525f6060820152614d158a5f52600460205260405f2090565b614d82565b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614e56576003018054918203918211610ec45755565b6004018054918201809211610ec45755565b90614e7d915f52600660205260405f2061317e565b600181015460a01c60ff1615614e56576003018054918203918211610ec45755565b6001915f520160205260405f2054151590565b91610cf89391614ec19361572b565b909291926157ad565b908160209103126104075751600a8110156104075790565b905f93614f2d6060614f2661104f614f16614f016123e160325461040b565b976142356001600160a01b038a1615156144d4565b6001600160a01b0387169061317e565b0151151590565b614fc057614fb85791602091614f5b95935f604051809881958294633f4bdec560e01b845260048401613b91565b03925af1928315612e9a575f93614f87575b50600183614f7a81611d98565b03614f8157565b60019150565b614faa91935060203d602011614fb1575b614fa281836105c7565b810190614eca565b915f614f6d565b503d614f98565b505050600191565b50505050506002905f90565b6001600160a01b03169291906001841461505b5781156150525761501b92156150265760405163a9059cbb60e01b602082015291615013918391612a019160248401613b91565b600592615829565b15610cf85750600190565b6040516340c10f1960e01b60208201529161504a918391612a019160248401613b91565b600692615829565b50505050600190565b9061507d9350611762925061506e611c92565b916001600160a01b0316615872565b61508657600190565b600590565b805f525f516020615e045f395f51905f5260205260ff6150ae8360405f2061317e565b541661511557805f525f516020615e045f395f51905f526020526150d58260405f2061317e565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610cf8916001600160a01b031690615944565b805f525f516020615e045f395f51905f5260205260ff6151518360405f2061317e565b54161561511557805f525f516020615e045f395f51905f526020526151798260405f2061317e565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b156151c057505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156151f157565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361526e57506105fa94506152366152288286612ef6565b34143490612d718488612ef6565b80615242575b50615a09565b6152636152689161525460345461040b565b9061525d611c92565b91615872565b6151ea565b5f61523c565b9061527a343415612b78565b61528f6152878287612ef6565b30848961598b565b80615323575b506152ab6117626001613bd186613a7287612b42565b6152bb575b506105fa9350615a09565b604051632770a7eb60e21b8152602081806152da883060048401613b91565b03815f885af1918215612e9a576105fa966152fe9387935f91615304575b506151b6565b5f6152b0565b61531d915060203d6020116129155761290881836105c7565b5f6152f8565b61533b906153356123e160345461040b565b876159d4565b5f615295565b90813b156153c2575f516020615d845f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156153aa576153a791615a52565b50565b5050346153b357565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b60ff5f516020615ee45f395f51905f525460401c16156153ff57565b631afcd79f60e31b5f5260045ffd5b6153a7906149aa6153e3565b906154236153e3565b61543160ff8316151561344c565b60409182519261544181856105c7565b60098452682b30b634b230ba37b960b91b6020850152615463815191826105c7565b60058152640312e302e360dc1b602082015261547d6153e3565b6154856153e3565b83516001600160401b038111610571576154b5816154b05f516020615cc45f395f51905f5254612f3a565b613fad565b6020601f821160011461554457816154f293926154de926105fa97985f9261412b575050614034565b5f516020615cc45f395f51905f5255615bae565b6155075f5f516020615d445f395f51905f5255565b61551c5f5f516020615f045f395f51905f5255565b60ff1660ff195f516020615ce45f395f51905f525416175f516020615ce45f395f51905f5255565b5f516020615cc45f395f51905f525f52601f198216955f516020615d645f395f51905f52965f5b8181106155bb57509660019284926154f296956105fa999a106155a3575b505050811b015f516020615cc45f395f51905f5255615bae565b01515f1960f88460031b161c191690555f8080615589565b8383015189556001909801976020938401930161556b565b6155db6153e3565b62015180600155565b8054821015612640575f5260205f2001905f90565b80548015615620575f19019061560f82826155e4565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146156cf575f198401848111610ec45783545f19810194908511610ec4575f95858361568297610bf39503615688575b5050506155f9565b55600190565b6156b86156b2916156a961569f6156c695886155e4565b90549060031b1c90565b928391876155e4565b906135d1565b85905f5260205260405f2090565b555f808061567a565b505050505f90565b6156df615a6f565b6156e7615ac6565b6040519060208201925f516020615f245f395f51905f528452604083015260608201524660808201523060a082015260a0815261572560c0826105c7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615798579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612e9a575f516001600160a01b0381161561578e57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611da257565b6157b6816157a3565b806157bf575050565b6157c8816157a3565b600181036157df5763f645eedf60e01b5f5260045ffd5b6157e8816157a3565b60028103615803575063fce698f760e01b5f5260045260245ffd5b8061580f6003926157a3565b146158175750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f161584d612798565b901561511557805161586957503b1561586557600190565b5f90565b60209150015190565b8147106158cc5782516001600160a01b03909116925f9182916020018486620186a0f19061589e612798565b91156158c557156158b1575b5050600190565b805161586957503b15615865575f806158aa565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461593f57600254600160401b8110156105715761592861591282600185940160025560026155e4565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61594e8282614e9f565b61511557805490600160401b82101561057157826159766159128460018096018555846155e4565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105fa916159cf6084836105c7565b615af8565b6159cf6105fa93926159fb60405194859263a9059cbb60e01b602085015260248401613b91565b03601f1981018452836105c7565b90615a1e915f52600660205260405f2061317e565b600181015460a01c60ff1615615a40576003018054918201809211610ec45755565b6004018054918203918211610ec45755565b5f80610cf893602081519101845af4615a69612798565b91615b50565b615a776147c4565b8051908115615a87576020012090565b50505f516020615d445f395f51905f52548015615aa15790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615ace61487e565b8051908115615ade576020012090565b50505f516020615f045f395f51905f52548015615aa15790565b905f602091828151910182855af11561278d575f513d615b4757506001600160a01b0381163b155b615b275750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615b20565b90615b745750805115615b6557805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615ba5575b615b85575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b7d565b80519091906001600160401b03811161057157615bef81615bdc5f516020615d245f395f51905f5254612f3a565b5f516020615d245f395f51905f52613ffc565b602092601f8211600114615c2257615c11929382915f9261412b575050614034565b5f516020615d245f395f51905f5255565b5f516020615d245f395f51905f525f52601f198216935f516020615f445f395f51905f52915f5b868110615c8b5750836001959610615c73575b505050811b015f516020615d245f395f51905f5255565b01515f1960f88460031b161c191690555f8080615c5c565b91926020600181928685015181550194019201615c4956feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220ac5a176b8fa3baf378781320d9569bf7dd3d87417d6f8862a4bdfbc9d9225fbe64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// CrossBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossBridgeMetaData.ABI instead.
var CrossBridgeABI = CrossBridgeMetaData.ABI

// CrossBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossBridgeBinRuntime = "60806040526004361015610022575b3615610018575f80fd5b6100206135ee565b005b5f3560e01c806301ffc9a7146103b15780630b1d8d06146103ac5780631313996b146103a75780631938e0f2146103a2578063248a9ca31461039d5780632f2ff15d1461039857806336568abe14610393578063365d71e41461038e57806342cde4e81461038957806348a00ed8146103845780634be13f831461037f5780634d5d00561461037a5780634ee078ba146103755780634f1ef28614610370578063502cc5c01461036b57806352d1902d146103665780635b605f5c146103615780635c975abb1461035c5780635fd262de14610357578063670e62681461035257806369ceb2f11461034d5780637921487414610348578063814914b514610343578063815f1de61461033e57806384b0196e1461033957806388d67d6d1461033457806389232a001461032f5780638ae87c5c1461032a57806391cca3db1461032557806391cf6d3e1461032057806391d148541461031b57806394ddc8c6146103165780639f9b4f1c14610311578063a217fddf1461030c578063a3246ad314610307578063aa1bd0bc14610302578063aa20ed47146102fd578063ad3cb1cc146102f8578063ae6893f8146102f3578063b2b49e2e146102ee578063b33eb36e146102e9578063b7f3358d146102e4578063b8886e1a146102df578063b8aa837e146102da578063bc8cadd6146102d5578063bedb86fb146102d0578063cd959706146102cb578063cf56118e146102c6578063d477f05f146102c1578063d547741f146102bc578063d5717fc5146102b7578063dfcbce15146102b2578063e2af6cd7146102ad578063edbbea39146102a8578063f0702e8e146102a3578063f45096431461029e5763f698da250361000e57612531565b612491565b612469565b61241f565b61236d565b61229f565b612266565b612237565b6121d9565b612165565b61214b565b612062565b61203a565b611f93565b611f47565b611ec4565b611e33565b611d26565b611ced565b611ca6565b611c1d565b611bd1565b611b55565b611af9565b6119dd565b611935565b6118dd565b6118c0565b611898565b611830565b61170e565b611604565b6114dc565b6113ef565b611325565b6112a5565b611244565b6111cd565b6110c5565b611097565b610fb8565b610ec9565b610e3a565b610cfb565b610b92565b610b25565b610ad1565b6109b3565b610987565b61091a565b61082a565b6107f1565b6107c0565b61070f565b6104c9565b610428565b346104075760203660031901126104075760043563ffffffff60e01b811680910361040757602090637965db0b60e01b81149081156103f6575b506040519015158152f35b6301ffc9a760e01b1490505f6103eb565b5f80fd5b6001600160a01b031690565b6001600160a01b0381160361040757565b346104075760203660031901126104075760043561044581610417565b61044d61363d565b6001600160a01b031661046181151561259d565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f84011215610407578235916001600160401b038311610407576020808501948460051b01011161040757565b6040366003190112610407576004356001600160401b038111610407576104f4903690600401610499565b602435916001600160401b0383116104075736602384011215610407576004830135916001600160401b0383116104075736602460e08502860101116104075760246100209401916127de565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761057157604052565b610541565b60e081019081106001600160401b0382111761057157604052565b606081019081106001600160401b0382111761057157604052565b60c081019081106001600160401b0382111761057157604052565b601f909101601f19168101906001600160401b0382119082101761057157604052565b604051906105fa610100836105c7565b565b604051906105fa6060836105c7565b604051906105fa6080836105c7565b604051906105fa60e0836105c7565b6001600160401b0381116105715760051b60200190565b60ff81160361040757565b9080601f8301121561040757813561066281610629565b9261067060405194856105c7565b81845260208085019260051b82010192831161040757602001905b8282106106985750505090565b6020809183356106a781610640565b81520191019061068b565b9080601f830112156104075781356106c981610629565b926106d760405194856105c7565b81845260208085019260051b82010192831161040757602001905b8282106106ff5750505090565b81358152602091820191016106f2565b6080366003190112610407576004356001600160401b0381116104075760c06003198236030112610407576024356001600160401b0381116104075761075990369060040161064b565b906044356001600160401b038111610407576107799036906004016106b2565b60643591906001600160401b038311610407576107bc936107a16107aa9436906004016106b2565b9260040161291c565b60405191829182901515815260200190565b0390f35b346104075760203660031901126104075760206107de600435612c60565b604051908152f35b35906105fa82610417565b346104075760403660031901126104075761002060243560043561081482610417565b61082561082082612c60565b613825565b6142d9565b346104075760403660031901126104075760043560243561084a81610417565b336001600160a01b038216036108635761002091614339565b63334bd91960e11b5f5260045ffd5b906040600319830112610407576004356001600160401b038111610407578261089d916004016106b2565b91602435906001600160401b03821161040757806023830112156104075781600401356108c981610629565b926108d760405194856105c7565b8184526024602085019260051b82010192831161040757602401905b8282106109005750505090565b60208091833561090f81610417565b8152019101906108f3565b346104075761092836610872565b906109368151835114612c7e565b5f5b8251811015610020578061097661095160019385612c94565b51838060a01b036109628488612c94565b51169061097161082082612c60565b614339565b5001610938565b5f91031261040757565b34610407575f36600319011261040757602060ff5f516020615ce45f395f51905f525416604051908152f35b34610407576060366003190112610407576024356004356044356109d681610417565b6109de6136b7565b6109e661388f565b815f526007602052610a04836109ff8160405f20614e9f565b612ca8565b80610a0f8484614a92565b916001600160a01b031615610abd575b8151915f516020615d045f395f51905f526040820192610a6c610a5a855160018060a01b03169660808601978489519160a089015193613bac565b610a6381611d98565b60018114614398565b8251602090930151935194516040516001600160a01b0390961695918291610a979142919084612c42565b0390a45f516020615ca45f395f51905f525f80a35f5f516020615e645f395f51905f525d005b5060608101516001600160a01b0316610a1f565b34610407575f366003190112610407575f546040516001600160a01b039091168152602090f35b9181601f84011215610407578235916001600160401b038311610407576020838186019501011161040757565b6101c036600319011261040757602435600435610b4182610417565b604435610b4d81610417565b6064359060a43560843560c4356001600160401b03811161040757610b76903690600401610af8565b94909360e03660e3190112610407576107bc976107aa97612cc2565b3461040757604036600319011261040757610c7a602435600435610bb4613868565b610bbc61388f565b805f526007602052610bd5826109ff8160405f20614e9f565b610c756040610c05610c0085610bf3865f52600860205260405f2090565b905f5260205260405f2090565b613076565b610c62610c2582516080610c1b8683015161040b565b9101519087613a8d565b50610c2f81611d98565b610c3881611d98565b83516020810182905290600190610c5c83604081015b03601f1981018552846105c7565b146130af565b01518015908115610c82575b42916130db565b6143bb565b6100206138c4565b4281109150610c6e565b6001600160401b03811161057157601f01601f191660200190565b929192610cb382610c8c565b91610cc160405193846105c7565b829481845281830111610407578281602093845f960137010152565b9080601f8301121561040757816020610cf893359101610ca7565b90565b604036600319011261040757600435610d1381610417565b6024356001600160401b03811161040757610d32903690600401610cdd565b903061aaaa625c0eb760891b01148015610e0f575b610e0057610d5361363d565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610dcf575b50610d9c57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d845f395f51905f528303610dbb576100209250615341565b632a87526960e21b5f52600483905260245ffd5b610df291945060203d602011610df9575b610dea81836105c7565b810190613b82565b925f610d7b565b503d610de0565b63703e46dd60e11b5f5260045ffd5b505f516020615d845f395f51905f52546001600160a01b031661aaaa625c0eb760891b011415610d47565b3461040757606036600319011261040757602435600435604435610e5c6136b7565b815f526007602052610e75836109ff8160405f20614e9f565b4201804211610ec45760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612ee2565b34610407575f366003190112610407577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e005760206040515f516020615d845f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610f985750505090565b909192602060e082610fad6001948851610f20565b019401929101610f8b565b3461040757602036600319011261040757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061107e575050611007925003836105c7565b611011825161312f565b915f5b81518110156110705760019061105461104f61102f86612b42565b61104961103c8588612c94565b516001600160a01b031690565b9061317e565b613193565b61105e8287612c94565b526110698186612c94565b5001611014565b604051806107bc8682610f75565b8454835260019485019487945060209093019201610ff2565b34610407575f36600319011261040757602060ff5f516020615e245f395f51905f5254166040519015158152f35b60e0366003190112610407576024356004356110e082610417565b604435916110ed83610417565b6064359260c435916084359160a435916001600160401b0385116104075761119e96611153611123611194973690600401610af8565b95909661112e613868565b6001600160a01b038516948490611145878d614465565b61114d61388f565b8b61451b565b939092604051986111638a610555565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610ca7565b60e08201526146c6565b5f5f516020615e645f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b34610407576080366003190112610407576024356004356111ed82610417565b604435906001600160401b0382116104075736602383011215610407576107bc92611225611238933690602481600401359101610ca7565b906064359261123384610640565b61321b565b604051918291826111ba565b34610407575f366003190112610407576020606454604051908152f35b90602080835192838152019201905f5b81811061127e5750505090565b8251845260209384019390920191600101611271565b906020610cf8928181520190611261565b34610407576020366003190112610407576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106112ff576107bc856112f3818703826105c7565b60405191829182611294565b82548452602090930192600192830192016112dc565b60e0810192916105fa9190610f20565b34610407576040366003190112610407576107bc61136460243560043561134b82610417565b6113536130f9565b505f52600660205260405f2061317e565b60046040519161137383610576565b80546001600160a01b0316835260018101546113c8906113bf906113a26113998261040b565b60208801612f2b565b6113b660a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611315565b346104075760203660031901126104075760043561140c81610417565b61141461363d565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916114b2906114a4610cf897959693600f60f81b865260e0602087015260e0860190611459565b908482036040860152611459565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611261565b34610407575f366003190112610407575f516020615d445f395f51905f52541580611570575b156115335761150f6147c4565b61151761487e565b906107bc6115236132e4565b604051938493309146918661147d565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615f045f395f51905f525415611502565b9080601f8301121561040757813561159d81610629565b926115ab60405194856105c7565b81845260208085019260051b820101918383116104075760208201905b8382106115d757505050505090565b81356001600160401b038111610407576020916115f9878480948801016106b2565b8152019101906115c8565b6080366003190112610407576004356001600160401b0381116104075761162f903690600401610499565b90602435906001600160401b038211610407573660238301121561040757816004013561165b81610629565b9261166960405194856105c7565b8184526024602085019260051b820101903682116104075760248101925b8284106116df5750506044359150506001600160401b038111610407576116b2903690600401611586565b606435926001600160401b038411610407576107bc946116d96107aa953690600401611586565b936132ff565b83356001600160401b0381116104075760209161170383926024369187010161064b565b815201930192611687565b346104075760603660031901126104075760043561172b81610417565b6024359061173882610417565b60443561174481610640565b5f516020615ee45f395f51905f52549260ff604085901c1615611766565b1590565b936001600160401b031680159081611828575b600114908161181e575b159081611815575b50611806576117a6928461179d613395565b6117f957614918565b6117ac57005b6117ca5f516020615ee45f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b6118016133b6565b614918565b63f92ee8a960e01b5f5260045ffd5b9050155f61178b565b303b159150611783565b859150611779565b346104075760403660031901126104075760243560043561184f61363d565b61185761388f565b6118618282614a92565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615e645f395f51905f525d005b34610407575f366003190112610407576034546040516001600160a01b039091168152602090f35b34610407575f366003190112610407576020603554604051908152f35b3461040757604036600319011261040757602060ff61191f60243560043561190482610417565b5f525f516020615e045f395f51905f52845260405f2061317e565b54166040519015158152f35b8015150361040757565b346104075760603660031901126104075760243560043561195582610417565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea60206044356119848161192b565b61198c613731565b835f52600582526119d28160016119cc60405f20986119bc816119b7858060a01b038216809d614e9f565b6133d7565b885f526006875260405f2061317e565b016133ff565b6040519015158152a3005b34610407576040366003190112610407576004356001600160401b03811161040757611a0d9036906004016106b2565b6024356001600160401b03811161040757611a2c9036906004016106b2565b90611a3a81518351146125f4565b5f5b82518110156100205780611aeb611a5560019385612c94565b51611a608387612c94565b5190611a6a613868565b611a7261388f565b805f526007602052611a8b826109ff8160405f20614e9f565b610c756040611aa9610c0085610bf3865f52600860205260405f2090565b610c62611abf82516080610c1b8683015161040b565b50611ac981611d98565b611ad281611d98565b835160208101829052908a90610c5c8360408101610c4e565b611af36138c4565b01611a3c565b34610407575f3660031901126104075760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611b365750505090565b82516001600160a01b0316845260209384019390920191600101611b29565b34610407576020366003190112610407576004355f525f516020615da45f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611bbb576107bc85611baf818703826105c7565b60405191829182611b13565b8254845260209093019260019283019201611b98565b34610407576020366003190112610407577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611c1061363d565b80600155604051908152a1005b3461040757604036600319011261040757602435600435611c3c6136b7565b611c446136b7565b611c4c61388f565b805f526007602052611c65826109ff8160405f20614e9f565b611c6f82826143bb565b5f516020615ca45f395f51905f525f80a35f5f516020615e645f395f51905f525d005b60405190611ca16020836105c7565b5f8252565b34610407575f366003190112610407576107bc604051611cc76040826105c7565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611459565b34610407576020366003190112610407576004355f526004602052600160405f20015460018101809111610ec457602090604051908152f35b3461040757611d3436610872565b90611d428151835114612c7e565b5f5b82518110156100205780611d7d611d5d60019385612c94565b51838060a01b03611d6e8488612c94565b51169061082561082082612c60565b5001611d44565b634e487b7160e01b5f52602160045260245ffd5b600a1115611da257565b611d84565b90600a821015611da25752565b6020815260606040611e1960a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611459565b93611e2b602082015183860190611da7565b015191015290565b3461040757604036600319011261040757600435602435905f60408051611e5981610591565b611e6161341c565b815282602082015201525f52600860205260405f20905f526020526107bc60405f20600760405191611e9283610591565b611e9b81612f72565b8352611eb160ff6006830154166020850161306a565b0154604082015260405191829182611db4565b34610407576020366003190112610407577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611f0681610640565b611f0e61363d565b16611f1a81151561344c565b8060ff195f516020615ce45f395f51905f525416175f516020615ce45f395f51905f5255604051908152a1005b34610407576020366003190112610407577f146cf9ff7459bfbdd8c5d5fb95ef19a728c1d2cae21b455499a7124ff7d8129f6020600435611f8661363d565b80606455604051908152a1005b3461040757604036600319011261040757600435602435611fb38161192b565b611fbb613731565b611fd0825f52600360205260405f2054151590565b156120275760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f526004825261201c81600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b34610407575f366003190112610407576033546040516001600160a01b039091168152602090f35b346104075760203660031901126104075760043561207f8161192b565b612087613731565b156120e557612094613868565b600160ff195f516020615e245f395f51905f525416175f516020615e245f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615e245f395f51905f525460ff81161561213c5760ff19165f516020615e245f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b34610407575f3660031901126104075760206107de613462565b34610407575f36600319011261040757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106121c3576107bc856112f3818703826105c7565b82548452602090930192600192830192016121ac565b34610407576020366003190112610407576004356121f681610417565b6121fe61363d565b6001600160a01b031661221281151561259d565b603480546001600160a01b031916821790555f516020615dc45f395f51905f525f80a2005b346104075760403660031901126104075761002060243560043561225a82610417565b61097161082082612c60565b34610407576020366003190112610407576004355f526004602052600260405f20015460018101809111610ec457602090604051908152f35b346104075760c0366003190112610407576004356122bc81610417565b602435906122c982610417565b6044356122d581610640565b6084356064356122e482610417565b60a435925f516020615ee45f395f51905f5254956123096117628860ff9060401c1690565b966001600160401b031680159081612365575b600114908161235b575b159081612352575b50611806576117a69587612340613395565b156134825761234d6133b6565b613482565b9050155f61232e565b303b159150612326565b88915061231c565b346104075760203660031901126104075760043561238a81610417565b61239261363d565b5f546001600160a01b03168061240d57506123b76001600160a01b03821615156135ac565b5f80546001600160a01b0319166001600160a01b0392831617908190556123e6906123e19061040b565b61040b565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b34610407576080366003190112610407576100206004356024356124428161192b565b6044359061244f82610417565b6064359261245c84610417565b6124646137ab565b614d1f565b34610407575f366003190112610407576032546040516001600160a01b039091168152602090f35b34610407576040366003190112610407576024356004356124b182610417565b6124b961363d565b805f5260056020525f60046124f460408320946124e3816119b760018060a01b0382168099615634565b84845260066020526040842061317e565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b34610407575f36600319011261040757612549615a6f565b612551615ac6565b6040519060208201925f516020615f245f395f51905f528452604083015260608201524660808201523060a082015260a0815261258f60c0826105c7565b519020604051908152602090f35b156125a457565b638219abe360e01b5f5260045ffd5b603480546001600160a01b0319166001600160a01b0392909216919091179055565b80546001600160a01b0319166001600160a01b03909216919091179055565b156125fb57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156126405760051b8101359060fe1981360301821215610407570190565b61260a565b35610cf881610417565b903590601e198136030182121561040757018035906001600160401b0382116104075760200191813603831361040757565b91908110156126405760e0020190565b908160209103126104075751610cf88161192b565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361271797939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c08601916126a6565b94803561272381610417565b6001600160a01b031660e0850152602081013561273f81610417565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561277481610640565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d156127c2573d906127a982610c8c565b916127b760405193846105c7565b82523d5f602084013e565b606090565b604090610cf8939281528160208201520190611459565b90919392936127ee8584146125f4565b5f945b83861061280057505050509050565b61280b86858561261e565b3560206128238161281d8a898961261e565b01612645565b612833606061281d8b8a8a61261e565b926128a98a86888a8c61288d608061284c87848661261e565b01359561288561287b8260a061286382888a61261e565b01359560c061287383838b61261e565b01359761261e565b60e081019061264f565b969095612681565b946040519a8b998a996326ae802b60e11b8b5260048b016126c6565b03815f305af190816128f0575b506128e557856128c4612798565b60405163f495148b60e01b81529182916128e191600484016127c7565b0390fd5b6001909501946127f1565b6129109060203d8111612915575b61290881836105c7565b810190612691565b6128b6565b503d6128fe565b90612a0193929161292b613868565b61293361388f565b803592612948845f52600560205260405f2090565b9161297861296660408301946129606123e187612645565b906138d6565b6129726123e186612645565b90612b50565b612983343415612b78565b612a176129a3865f526004602052600260405f2001600181540180915590565b966129b6602084013589819a8214612b96565b6129c26123e186612645565b936060840196886129d289612645565b96612a0f8c60808901359e8f60a08b019b6129ed8d8d61264f565b929091604051978896602088019a8b612bb4565b03601f1981018352826105c7565b519020613919565b612a2a87612a2485612645565b87613af0565b919092600184612a3981611d98565b14612b06575b50612a4983611d98565b60018303612aa35750505090612a75612a6f5f516020615d045f395f51905f5293612645565b91612645565b6040516001600160a01b03909216958291612a9291429184612c42565b0390a45b612a9e6138c4565b600190565b612ae08394612adb612afe947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612ae6956141a6565b612645565b92612645565b9260405193849360018060a01b031698429185612c18565b0390a4612a96565b612b3b91935088612b1686612645565b91612b34612b2d612b268a612645565b928861264f565b3691610ca7565b9289613bac565b915f612a3f565b5f52600660205260405f2090565b15612b585750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612b805750565b63943892eb60e01b5f525f60045260245260445ffd5b15612b9f575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610cf897959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c082015201916126a6565b6105fa93606092969593608083019760018060a01b03168352602083015260408201520190611da7565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615e045f395f51905f52602052600160405f20015490565b15612c8557565b63031206d560e51b5f5260045ffd5b80518210156126405760209160051b010190565b15612cb05750565b6373a1390160e11b5f5260045260245ffd5b959394612d44919597939297612cd6613868565b612d1b6001600160a01b038816612ced818b614465565b612cf561388f565b612d056123e16123e160e4612645565b811490612d156123e160e4612645565b91612e9f565b612d3c612d2c6123e1610104612645565b6001600160a01b038b1614612ecc565b83878961451b565b9161012435612d7781612d6086612d5b8787612ef6565b612ef6565b811015612d7187612d5b8888612ef6565b90612f03565b612d856123e160325461040b565b90612d91610104612645565b906101443592612da2610164612f21565b92610184356101a43591833b1561040757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612e9a57612e6161119498612e6a93612a969c612e80575b50612e58612e43610104612645565b91612e4c6105ea565b9c8d5260208d01612f2b565b60408b01612f2b565b60608901612f2b565b608087015260a086015260c08501523691610ca7565b80612e8e5f612e94936105c7565b8061097d565b5f612e34565b61278d565b15612ea8575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612ed357565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ec457565b15612f0c575050565b63943892eb60e01b5f5260045260245260445ffd5b35610cf881610640565b6001600160a01b039091169052565b90600182811c92168015612f68575b6020831014612f5457565b634e487b7160e01b5f52602260045260245ffd5b91607f1691612f49565b90604051612f7f816105ac565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91612fd782612f3a565b80855291600181169081156130435750600114613004575b505060a092916130009103846105c7565b0152565b5f908152602081209092505b818310613027575050810160200181613000612fef565b6020919350806001915483858901015201910190918492613010565b60ff191660208681019190915292151560051b850190920192508391506130009050612fef565b600a821015611da25752565b9060405161308381610591565b60406007829461309281612f72565b84526130a860ff6006830154166020860161306a565b0154910152565b156130b75750565b60405162461bcd60e51b8152602060048201529081906128e1906024830190611459565b156130e4575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061310682610576565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061313982610629565b61314660405191826105c7565b8281528092613157601f1991610629565b01905f5b82811061316757505050565b6020906131726130f9565b8282850101520161315b565b9060018060a01b03165f5260205260405f2090565b906040516131a081610576565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916131eb906131e2906113b6565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126104075751610cf881610417565b5f5490949392906001600160a01b038116156132d557604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90613284906084870190611459565b931691015203925af18015612e9a576105fa925f916132a6575b5080946135c2565b6132c8915060203d6020116132ce575b6132c081836105c7565b810190613206565b5f61329e565b503d6132b6565b6315aeca0d60e11b5f5260045ffd5b604051906132f36020836105c7565b5f808352366020840137565b919294939093848451148061338b575b80613381575b61331e906125f4565b5f5b85811015613375578060051b8401359060be19853603018212156104075761336e60019261334e8389612c94565b51613359848c612c94565b51906133658589612c94565b5192890161291c565b5001613320565b50945050505050600190565b5081518514613315565b508486511461330f565b5f516020615ee45f395f51905f5280546001600160401b0319166001179055565b5f516020615ee45f395f51905f52805460ff60401b1916600160401b179055565b156133df5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b60405190613429826105ac565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b1561345357565b63f0f15b9160e01b5f5260045ffd5b6065545f52600660205260405f2060015f52602052600460405f20015490565b919492939094841561359d576001600160a01b0382161561358e576134a56153e3565b6134b96001600160a01b038416151561259d565b6001600160a01b038616906134cf82151561259d565b60ff81161561357f576105fa9661353661354392613531613563976134f26153e3565b6134fa6153e3565b6135026153e3565b60ff195f516020615e245f395f51905f5254165f516020615e245f395f51905f525561352c6153e3565b61540e565b61541a565b61353e6155d3565b6125b3565b5f516020615dc45f395f51905f525f80a261355d43603555565b83614c15565b8061356f575b50606555565b6135799082614e1b565b5f613569565b6338854f1160e21b5f5260045ffd5b6324cd8cef60e01b5f5260045ffd5b63110d132360e21b5f5260045ffd5b156135b357565b6302bfb33d60e51b5f5260045ffd5b905f6105fa93926124646137ab565b916135ea9183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361360257565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b5f516020615e845f395f51905f525f525f516020615e045f395f51905f5260205260ff61368a337fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61317e565b54161561369357565b63e2517d3f60e01b5f52336004525f516020615e845f395f51905f5260245260445ffd5b5f516020615ec45f395f51905f525f525f516020615e045f395f51905f5260205260ff613704337fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361317e565b54161561370d57565b63e2517d3f60e01b5f52336004525f516020615ec45f395f51905f5260245260445ffd5b5f516020615de45f395f51905f525f525f516020615e045f395f51905f5260205260ff61377e337f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661317e565b54161561378757565b63e2517d3f60e01b5f52336004525f516020615de45f395f51905f5260245260445ffd5b5f516020615e445f395f51905f525f525f516020615e045f395f51905f5260205260ff6137f8337f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61317e565b54161561380157565b63e2517d3f60e01b5f52336004525f516020615e445f395f51905f5260245260445ffd5b805f525f516020615e045f395f51905f5260205260ff6138483360405f2061317e565b5416156138525750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615e245f395f51905f52541661388057565b63d93c066560e01b5f5260045ffd5b5f516020615e645f395f51905f525c6138b55760015f516020615e645f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615e645f395f51905f525d565b610cf8916001600160a01b031690614e9f565b156138f057565b63b3f07a3960e01b5f5260045ffd5b156139075750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a83575b613932906138e9565b61395361394d5f516020615ce45f395f51905f525460ff1690565b60ff1690565b9561396184888110156138ff565b5f945f935f5b86811061398257505050505050506105fa91928110156138ff565b6139df6139ae836139916156d7565b6042916040519161190160f01b8352600283015260228201522090565b6139c26139bb8489612c94565b5160ff1690565b6139cc8487612c94565b51906139d88589612c94565b5192614eb2565b6001600160a01b038181169088161080613a10575b613a02575b50600101613967565b6001988901989096506139f9565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615e045f395f51905f52602052613a7e613a77827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561317e565b61317e565b5460ff1690565b6139f4565b5085518314613929565b9091906001600160a01b03831660011480613ae5575b613ab9575b90613ab592600192614ee2565b9091565b90613ac2613462565b818101809111610ec45760645410613ada5790613aa8565b505050600990600190565b506065548114613aa3565b9091906001600160a01b03831660011480613b38575b613b17575b90613ab5925f92614ee2565b90613b20613462565b818101809111610ec45760645410613ada5790613b0b565b506065548114613b06565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610cf892910190611459565b90816020910312610407575190565b6001600160a01b039091168152602081019190915260400190565b92919382613bb985612b42565b95613bdc6001613bd1818060a01b038716809a61317e565b015460a01c60ff1690565b92601481511180613ede575b613c21575b50613bf793614fcc565b92613c0184611d98565b60018414613c10575b50505090565b613c1992614e68565b5f8080613c0a565b9150613c5993602083015160601c6020613c3f6123e160335461040b565b916040518098819262483e5560e91b8352600483016111ba565b0381845afa938415612e9a578888958b985f91613ebf575b50613c7f5750505093613bed565b92809585988896936002613c9d6001985f52600460205260405f2090565b01549687921480159889613e84575b5060209596505f14613e79578c613ceb84925b613ccd6123e160335461040b565b95604051998a9889978896631eeed51360e01b885260048801613b43565b03925af1908115612e9a575f91613e5a575b50613e4a579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613d58898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613e42575b50613d74575b92613bf793614fcc565b9050613d846123e160335461040b565b92604051916370a0823160e01b835260208380613da488600483016111ba565b03818b5afa8015612e9a5786935f91613e23575b5080613dc8575b50909350613d6a565b6020919350613dec956040519687928392632770a7eb60e21b845260048401613b91565b03815f8b5af1918215612e9a57613bf794869315613dbf57613e1c9060203d6020116129155761290881836105c7565b505f613dbf565b613e3c915060203d602011610df957610dea81836105c7565b5f613db8565b90505f613d64565b50505050509091612a9e92614e68565b613e73915060203d6020116129155761290881836105c7565b5f613cfd565b8c613ceb5f92613cbf565b939250945094613e9393614fcc565b613e9c81611d98565b60018103613eb15750848884848a948a613cac565b995050505050505050505090565b613ed8915060203d6020116129155761290881836105c7565b5f613c71565b506033546001600160a01b0390613ef8906123e19061040b565b161515613be8565b15613f085750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126104075760405190613f33826105ac565b819380358352602081013560208401526040810135613f5181610417565b6040840152613f62606082016107e6565b60608401526080818101359084015260a0810135916001600160401b0383116104075760a092613f929201610cdd565b910152565b818110613fa2575050565b5f8155600101613f97565b90601f8211613fba575050565b6105fa915f516020615cc45f395f51905f525f5260205f20906020601f840160051c83019310613ff2575b601f0160051c0190613f97565b9091508190613fe5565b9190601f811161400b57505050565b6105fa925f5260205f20906020601f840160051c83019310613ff257601f0160051c0190613f97565b8160011b915f199060031b1c19161790565b90600a811015611da25760ff80198354169116179055565b81518051825560208101516001830155604081015191929161408c906001600160a01b0316600285016125d5565b60608101516140a7906001600160a01b0316600385016125d5565b6080810151600484015560a00151805160058401916001600160401b038211610571576140de826140d88554612f3a565b85613ffc565b602090601f831160011461413657826007959360409593614107935f9261412b575b5050614034565b90555b614124602082015161411b81611d98565b60068601614046565b0151910155565b015190505f80614100565b90601f1983169161414a855f5260205f2090565b925f5b81811061418e5750926001928592600798966040989610614176575b505050811b01905561410a565b01515f1960f88460031b161c191690555f8080614169565b9293602060018192878601518155019501930161414d565b91608061424961423a600293614235876142306141e06135ea99833595865f5260076020526141e560405f20602087013594858092615944565b613f00565b15614256576142196141f960015442612ef6565b915b61420e6142066105fc565b963690613f1a565b86526020860161306a565b6040840152610bf3855f52600860205260405f2090565b61405e565b612b42565b6110496123e160408801612645565b9301359201918254612ef6565b6142195f916141fb565b9061426b825f61508b565b91826142745750565b5f80525f516020615da45f395f51905f526020526001600160a01b03166142bb817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615944565b156142c35750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916142e6838261508b565b92836142f0575050565b815f525f516020615da45f395f51905f5260205261431b60405f209160018060a01b03168092615944565b15614324575050565b63d180cb3160e01b5f5260045260245260445ffd5b919091614346838261512e565b9283614350575050565b815f525f516020615da45f395f51905f5260205261437b60405f209160018060a01b03168092615634565b15614384575050565b62a95f1b60e31b5f5260045260245260445ffd5b156143a05750565b63290cd55f60e01b5f52600a811015611da25760045260245ffd5b906143c591614a92565b60018060a01b036060820151168151915f516020615d045f395f51905f52604082019261440d610a5a855160018060a01b03169660808601978489519160a089015193613bac565b8251602090930151935194516040516001600160a01b03909616959182916144389142919084612c42565b0390a4565b156144455750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f526005602052614489816119b760405f2060018060a01b03831690614e9f565b825f52600460205260ff600360405f200154166144c15760ff60016144b583613a726105fa9697612b42565b015460a81c161561443d565b82636fc24b7960e11b5f5260045260245ffd5b156144db57565b6330d45fb160e01b5f5260045ffd5b90816060910312610407578051916040602083015192015190565b1561450c57565b6358e8878560e01b5f5260045ffd5b826060916145b49795969361454561104f61453584612b42565b6001600160a01b0384169061317e565b6145556117626040830151151590565b614656575b506145696123e160325461040b565b9161457e6001600160a01b03841615156144d4565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612e9a575f955f905f93614618575b509082916105fa9493968198851515958661460d575b505084614602575b5050826145f7575b5050614505565b101590505f806145f0565b101592505f806145e8565b101594505f806145e0565b90506146439196506105fa93925060603d60601161464f575b61463b81836105c7565b8101906144ea565b919692939192916145ca565b503d614631565b60c06146689101518480821015612f03565b5f61455a565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e0916146c19190860190611459565b930152565b6146e481515f526004602052600160405f2001600181540180915590565b6146ee8251612b42565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086147336123e1600161472c60208801956110496123e1885161040b565b015461040b565b93805190614438614744855161040b565b916040810190614754825161040b565b9061477d6080820196875160a084019485519861477760c087019a8b5190612ef6565b93615200565b61479361478c8251995161040b565b935161040b565b9460e06147a3606084015161040b565b9751935191519201519260405197889760018060a01b03169c42968961466e565b604051905f825f516020615cc45f395f51905f5254916147e383612f3a565b808352926001811690811561485f5750600114614807575b6105fa925003836105c7565b505f516020615cc45f395f51905f525f90815290915f516020615d645f395f51905f525b8183106148435750509060206105fa928201016147fb565b602091935080600191548385890101520191019091849261482b565b602092506105fa94915060ff191682840152151560051b8201016147fb565b604051905f825f516020615d245f395f51905f52549161489d83612f3a565b808352926001811690811561485f57506001146148c0576105fa925003836105c7565b505f516020615d245f395f51905f525f90815290915f516020615f445f395f51905f525b8183106148fc5750509060206105fa928201016147fb565b60209193508060019154838589010152019101909184926148e4565b91906149226153e3565b6149366001600160a01b038416151561259d565b6001600160a01b0381169261494c84151561259d565b60ff83161561357f576149d9926149bf6149c5926149686153e3565b6149706153e3565b6149786153e3565b60ff195f516020615e245f395f51905f5254165f516020615e245f395f51905f52556149a26153e3565b6149aa6153e3565b6149b26153e3565b6149ba6153e3565b614260565b5061541a565b6149cd6153e3565b620151806001556125b3565b5f516020615dc45f395f51905f525f80a26105fa43603555565b91908203918211610ec457565b60075f9182815582600182015582600282015582600382015582600482015560058101614a2d8154612f3a565b9081614a40575b50508260068201550155565b601f8211600114614a5757849055505b5f80614a34565b614a7d614a8d926001601f614a6f855f5260205f2090565b920160051c82019101613f97565b5f81815260208120918190559055565b614a50565b9190614a9c61341c565b50825f526007602052614ab28160405f20615634565b15614b2557614b206105fa91845f52600860205260405f20815f52602052610bf3614adf60405f20612f72565b95614afc614aec82612b42565b6110496123e160408b015161040b565b614b10600260808a015192019182546149f3565b90555f52600860205260405f2090565b614a00565b637f11bea960e01b5f5260045260245ffd5b600360606105fa938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b15614b765750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c0600491614baf60018060a01b03825116856125d5565b6020810151614bfa906001860190614bd0906001600160a01b0316826125d5565b6040830151815460ff60a01b191690151560a01b60ff60a01b1617815560608301511515906133ff565b6080810151600285015560a081015160038501550151910155565b600190614c238115156135ac565b614c2c826135ac565b614cc4828060a01b03841693614c438515156135ac565b614c4c836158db565b614ce3575b614c7584614c7081614c6b875f52600560205260405f2090565b61511b565b614b6e565b614c94614c8061061a565b91614c8b8684612f2b565b60208301612f2b565b5f60408201525f60608201525f60808201525f60a08201525f60c0820152614cbf84613a7285612b42565b614b96565b6040515f81525f516020615ea45f395f51905f52908060208101614438565b614d1a614cee61060b565b8481525f60208201525f60408201525f6060820152614d15855f52600460205260405f2090565b614b37565b614c51565b906144385f516020615ea45f395f51905f5291949394614d408415156135ac565b6001600160a01b038616946107aa90614d5a8715156135ac565b6001600160a01b03811697614cbf90614d748a15156135ac565b614d7d886158db565b614de4575b614d9c81614c7081614c6b8c5f52600560205260405f2090565b614dbb614da761061a565b93614db28386612f2b565b60208501612f2b565b84151560408401525f60608401525f60808401525f60a08401525f60c0840152613a7288612b42565b614e16614def61060b565b8981525f60208201525f60408201525f6060820152614d158a5f52600460205260405f2090565b614d82565b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614e56576003018054918203918211610ec45755565b6004018054918201809211610ec45755565b90614e7d915f52600660205260405f2061317e565b600181015460a01c60ff1615614e56576003018054918203918211610ec45755565b6001915f520160205260405f2054151590565b91610cf89391614ec19361572b565b909291926157ad565b908160209103126104075751600a8110156104075790565b905f93614f2d6060614f2661104f614f16614f016123e160325461040b565b976142356001600160a01b038a1615156144d4565b6001600160a01b0387169061317e565b0151151590565b614fc057614fb85791602091614f5b95935f604051809881958294633f4bdec560e01b845260048401613b91565b03925af1928315612e9a575f93614f87575b50600183614f7a81611d98565b03614f8157565b60019150565b614faa91935060203d602011614fb1575b614fa281836105c7565b810190614eca565b915f614f6d565b503d614f98565b505050600191565b50505050506002905f90565b6001600160a01b03169291906001841461505b5781156150525761501b92156150265760405163a9059cbb60e01b602082015291615013918391612a019160248401613b91565b600592615829565b15610cf85750600190565b6040516340c10f1960e01b60208201529161504a918391612a019160248401613b91565b600692615829565b50505050600190565b9061507d9350611762925061506e611c92565b916001600160a01b0316615872565b61508657600190565b600590565b805f525f516020615e045f395f51905f5260205260ff6150ae8360405f2061317e565b541661511557805f525f516020615e045f395f51905f526020526150d58260405f2061317e565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610cf8916001600160a01b031690615944565b805f525f516020615e045f395f51905f5260205260ff6151518360405f2061317e565b54161561511557805f525f516020615e045f395f51905f526020526151798260405f2061317e565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b156151c057505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156151f157565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361526e57506105fa94506152366152288286612ef6565b34143490612d718488612ef6565b80615242575b50615a09565b6152636152689161525460345461040b565b9061525d611c92565b91615872565b6151ea565b5f61523c565b9061527a343415612b78565b61528f6152878287612ef6565b30848961598b565b80615323575b506152ab6117626001613bd186613a7287612b42565b6152bb575b506105fa9350615a09565b604051632770a7eb60e21b8152602081806152da883060048401613b91565b03815f885af1918215612e9a576105fa966152fe9387935f91615304575b506151b6565b5f6152b0565b61531d915060203d6020116129155761290881836105c7565b5f6152f8565b61533b906153356123e160345461040b565b876159d4565b5f615295565b90813b156153c2575f516020615d845f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156153aa576153a791615a52565b50565b5050346153b357565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b60ff5f516020615ee45f395f51905f525460401c16156153ff57565b631afcd79f60e31b5f5260045ffd5b6153a7906149aa6153e3565b906154236153e3565b61543160ff8316151561344c565b60409182519261544181856105c7565b60098452682b30b634b230ba37b960b91b6020850152615463815191826105c7565b60058152640312e302e360dc1b602082015261547d6153e3565b6154856153e3565b83516001600160401b038111610571576154b5816154b05f516020615cc45f395f51905f5254612f3a565b613fad565b6020601f821160011461554457816154f293926154de926105fa97985f9261412b575050614034565b5f516020615cc45f395f51905f5255615bae565b6155075f5f516020615d445f395f51905f5255565b61551c5f5f516020615f045f395f51905f5255565b60ff1660ff195f516020615ce45f395f51905f525416175f516020615ce45f395f51905f5255565b5f516020615cc45f395f51905f525f52601f198216955f516020615d645f395f51905f52965f5b8181106155bb57509660019284926154f296956105fa999a106155a3575b505050811b015f516020615cc45f395f51905f5255615bae565b01515f1960f88460031b161c191690555f8080615589565b8383015189556001909801976020938401930161556b565b6155db6153e3565b62015180600155565b8054821015612640575f5260205f2001905f90565b80548015615620575f19019061560f82826155e4565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f146156cf575f198401848111610ec45783545f19810194908511610ec4575f95858361568297610bf39503615688575b5050506155f9565b55600190565b6156b86156b2916156a961569f6156c695886155e4565b90549060031b1c90565b928391876155e4565b906135d1565b85905f5260205260405f2090565b555f808061567a565b505050505f90565b6156df615a6f565b6156e7615ac6565b6040519060208201925f516020615f245f395f51905f528452604083015260608201524660808201523060a082015260a0815261572560c0826105c7565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411615798579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612e9a575f516001600160a01b0381161561578e57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611da257565b6157b6816157a3565b806157bf575050565b6157c8816157a3565b600181036157df5763f645eedf60e01b5f5260045ffd5b6157e8816157a3565b60028103615803575063fce698f760e01b5f5260045260245ffd5b8061580f6003926157a3565b146158175750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f161584d612798565b901561511557805161586957503b1561586557600190565b5f90565b60209150015190565b8147106158cc5782516001600160a01b03909116925f9182916020018486620186a0f19061589e612798565b91156158c557156158b1575b5050600190565b805161586957503b15615865575f806158aa565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461593f57600254600160401b8110156105715761592861591282600185940160025560026155e4565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61594e8282614e9f565b61511557805490600160401b82101561057157826159766159128460018096018555846155e4565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105fa916159cf6084836105c7565b615af8565b6159cf6105fa93926159fb60405194859263a9059cbb60e01b602085015260248401613b91565b03601f1981018452836105c7565b90615a1e915f52600660205260405f2061317e565b600181015460a01c60ff1615615a40576003018054918201809211610ec45755565b6004018054918203918211610ec45755565b5f80610cf893602081519101845af4615a69612798565b91615b50565b615a776147c4565b8051908115615a87576020012090565b50505f516020615d445f395f51905f52548015615aa15790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615ace61487e565b8051908115615ade576020012090565b50505f516020615f045f395f51905f52548015615aa15790565b905f602091828151910182855af11561278d575f513d615b4757506001600160a01b0381163b155b615b275750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415615b20565b90615b745750805115615b6557805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615ba5575b615b85575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b7d565b80519091906001600160401b03811161057157615bef81615bdc5f516020615d245f395f51905f5254612f3a565b5f516020615d245f395f51905f52613ffc565b602092601f8211600114615c2257615c11929382915f9261412b575050614034565b5f516020615d245f395f51905f5255565b5f516020615d245f395f51905f525f52601f198216935f516020615f445f395f51905f52915f5b868110615c8b5750836001959610615c73575b505050811b015f516020615d245f395f51905f5255565b01515f1960f88460031b161c191690555f8080615c5c565b91926020600181928685015181550194019201615c4956feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220ac5a176b8fa3baf378781320d9569bf7dd3d87417d6f8862a4bdfbc9d9225fbe64736f6c634300081c0033"

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

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_CrossBridge *CrossBridgeCaller) BridgeExecuter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossBridge.contract.Call(opts, &out, "bridgeExecuter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_CrossBridge *CrossBridgeSession) BridgeExecuter() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecuter(&_CrossBridge.CallOpts)
}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_CrossBridge *CrossBridgeCallerSession) BridgeExecuter() (common.Address, error) {
	return _CrossBridge.Contract.BridgeExecuter(&_CrossBridge.CallOpts)
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

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_CrossBridge *CrossBridgeTransactor) SetBridgeExecuter(opts *bind.TransactOpts, _bridgeExecuter common.Address) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setBridgeExecuter", _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_CrossBridge *CrossBridgeSession) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecuter(&_CrossBridge.TransactOpts, _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetBridgeExecuter(&_CrossBridge.TransactOpts, _bridgeExecuter)
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

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.contract.Transact(opts, "setTokenPause", remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_CrossBridge *CrossBridgeTransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _CrossBridge.Contract.SetTokenPause(&_CrossBridge.TransactOpts, remoteChainID, token, pause)
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

// CrossBridgeBridgeExecuterCallFailedIterator is returned from FilterBridgeExecuterCallFailed and is used to iterate over the raw logs and unpacked data for BridgeExecuterCallFailed events raised by the CrossBridge contract.
type CrossBridgeBridgeExecuterCallFailedIterator struct {
	Event *CrossBridgeBridgeExecuterCallFailed // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeExecuterCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeExecuterCallFailed)
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
		it.Event = new(CrossBridgeBridgeExecuterCallFailed)
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
func (it *CrossBridgeBridgeExecuterCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeExecuterCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeExecuterCallFailed represents a BridgeExecuterCallFailed event raised by the CrossBridge contract.
type CrossBridgeBridgeExecuterCallFailed struct {
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
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeExecuterCallFailed(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*CrossBridgeBridgeExecuterCallFailedIterator, error) {

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

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeExecuterCallFailedIterator{contract: _CrossBridge.contract, event: "BridgeExecuterCallFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterCallFailed is a free log subscription operation binding the contract event 0x9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be1.
//
// Solidity: event BridgeExecuterCallFailed(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, bytes reason)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeExecuterCallFailed(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeExecuterCallFailed, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeExecuterCallFailed)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeExecuterCallFailed(log types.Log) (*CrossBridgeBridgeExecuterCallFailed, error) {
	event := new(CrossBridgeBridgeExecuterCallFailed)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossBridgeBridgeExecuterSetIterator is returned from FilterBridgeExecuterSet and is used to iterate over the raw logs and unpacked data for BridgeExecuterSet events raised by the CrossBridge contract.
type CrossBridgeBridgeExecuterSetIterator struct {
	Event *CrossBridgeBridgeExecuterSet // Event containing the contract specifics and raw log

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
func (it *CrossBridgeBridgeExecuterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossBridgeBridgeExecuterSet)
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
		it.Event = new(CrossBridgeBridgeExecuterSet)
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
func (it *CrossBridgeBridgeExecuterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossBridgeBridgeExecuterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossBridgeBridgeExecuterSet represents a BridgeExecuterSet event raised by the CrossBridge contract.
type CrossBridgeBridgeExecuterSet struct {
	BridgeExecuter common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuterSet is a free log retrieval operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_CrossBridge *CrossBridgeFilterer) FilterBridgeExecuterSet(opts *bind.FilterOpts, bridgeExecuter []common.Address) (*CrossBridgeBridgeExecuterSetIterator, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeBridgeExecuterSetIterator{contract: _CrossBridge.contract, event: "BridgeExecuterSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterSet is a free log subscription operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_CrossBridge *CrossBridgeFilterer) WatchBridgeExecuterSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeBridgeExecuterSet, bridgeExecuter []common.Address) (event.Subscription, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossBridgeBridgeExecuterSet)
				if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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
func (_CrossBridge *CrossBridgeFilterer) ParseBridgeExecuterSet(log types.Log) (*CrossBridgeBridgeExecuterSet, error) {
	event := new(CrossBridgeBridgeExecuterSet)
	if err := _CrossBridge.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_CrossBridge *CrossBridgeFilterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*CrossBridgeCrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _CrossBridge.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &CrossBridgeCrossMintableERC20CodeSetIterator{contract: _CrossBridge.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_CrossBridge *CrossBridgeFilterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *CrossBridgeCrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _CrossBridge.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
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
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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
