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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cross\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInitialSupply\",\"type\":\"uint256\"}],\"name\":\"initializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
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
	Bin: "0x60a080604052346100c257306080525f516020615c975f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615bd090816100c78239608051818181610e4c01526110110152f35b6001600160401b0319166001600160401b039081175f516020615c975f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146103545780630b1d8d061461034f5780631313996b1461034a5780631938e0f214610345578063248a9ca3146103405780632f2ff15d1461033b57806336568abe14610336578063365d71e4146103315780633aefddbf1461032c57806342cde4e81461032757806348a00ed8146103225780634be13f831461031d5780634d5d0056146103185780634ee078ba146103135780634f1ef2861461030e578063502cc5c01461030957806352d1902d146103045780635b605f5c146102ff5780635c975abb146102fa5780635fd262de146102f5578063670e6268146102f057806379214874146102eb578063814914b5146102e657806384b0196e146102e157806388d67d6d146102dc57806389232a00146102d75780638ae87c5c146102d257806391cca3db146102cd57806391cf6d3e146102c857806391d14854146102c357806394ddc8c6146102be5780639f9b4f1c146102b9578063a217fddf146102b4578063a3246ad3146102af578063aa1bd0bc146102aa578063aa20ed47146102a5578063ad3cb1cc146102a0578063ae6893f81461029b578063b2b49e2e14610296578063b33eb36e14610291578063b7f3358d1461028c578063b8aa837e14610287578063bedb86fb14610282578063cf56118e1461027d578063d477f05f14610278578063d547741f14610273578063d5717fc51461026e578063e2af6cd714610269578063edbbea3914610264578063f0702e8e1461025f578063f45096431461025a5763f698da2514610255575f80fd5b6124c8565b612418565b6123f0565b6123a6565b6122f4565b6122bb565b61228c565b61222e565b6121ba565b6120d1565b61202a565b611fa7565b611f16565b611e09565b611dd0565b611d89565b611cdf565b611c93565b611c17565b611bbb565b611a88565b6119aa565b611942565b611925565b6118fd565b611895565b6117d3565b6116ca565b6115a2565b611450565b6113d0565b611307565b611212565b6111e4565b6110ee565b610fff565b610f62565b610e0a565b610c8e565b610c21565b610bcd565b610a91565b610a65565b610914565b6108b1565b6107c1565b610788565b610757565b6106a6565b610460565b6103bf565b346103aa5760203660031901126103aa5760043563ffffffff60e01b81168091036103aa57602090637965db0b60e01b8114908115610399575b506040519015158152f35b6301ffc9a760e01b1490505f61038e565b5f80fd5b6001600160a01b038116036103aa57565b346103aa5760203660031901126103aa576004356103dc816103ae565b6103e4613501565b6001600160a01b03166103f8811515612547565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103aa578235916001600160401b0383116103aa576020808501948460051b0101116103aa57565b60403660031901126103aa576004356001600160401b0381116103aa5761048b903690600401610430565b602435916001600160401b0383116103aa57366023840112156103aa578260040135916001600160401b0383116103aa5736602460e08502860101116103aa5760246104d8940191612747565b005b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761050a57604052565b6104da565b60e081019081106001600160401b0382111761050a57604052565b606081019081106001600160401b0382111761050a57604052565b60c081019081106001600160401b0382111761050a57604052565b90601f801991011681019081106001600160401b0382111761050a57604052565b6040519061059161010083610560565b565b60405190610591606083610560565b60405190610591608083610560565b6040519061059160e083610560565b6001600160401b03811161050a5760051b60200190565b60ff8116036103aa57565b9080601f830112156103aa5781356105f9816105c0565b926106076040519485610560565b81845260208085019260051b8201019283116103aa57602001905b82821061062f5750505090565b60208091833561063e816105d7565b815201910190610622565b9080601f830112156103aa578135610660816105c0565b9261066e6040519485610560565b81845260208085019260051b8201019283116103aa57602001905b8282106106965750505090565b8135815260209182019101610689565b60803660031901126103aa576004356001600160401b0381116103aa5760c060031982360301126103aa576024356001600160401b0381116103aa576106f09036906004016105e2565b906044356001600160401b0381116103aa57610710903690600401610649565b90606435916001600160401b0383116103aa5761075393610738610741943690600401610649565b92600401612885565b60405190151581529081906020820190565b0390f35b346103aa5760203660031901126103aa576020610775600435612ba0565b604051908152f35b3590610591826103ae565b346103aa5760403660031901126103aa576104d86024356004356107ab826103ae565b6107bc6107b782612ba0565b6136bd565b613fb1565b346103aa5760403660031901126103aa576004356024356107e1816103ae565b336001600160a01b038216036107fa576104d891614011565b63334bd91960e11b5f5260045ffd5b9060406003198301126103aa576004356001600160401b0381116103aa578261083491600401610649565b91602435906001600160401b0382116103aa57806023830112156103aa578160040135610860816105c0565b9261086e6040519485610560565b8184526024602085019260051b8201019283116103aa57602401905b8282106108975750505090565b6020809183356108a6816103ae565b81520191019061088a565b346103aa576108bf36610809565b906108cd8151835114612bbe565b5f5b82518110156104d8578061090d6108e860019385612bd4565b51838060a01b036108f98488612bd4565b5116906109086107b782612ba0565b614011565b50016108cf565b346103aa5760c03660031901126103aa57600435610931816103ae565b6024359061093e826103ae565b60443561094a816105d7565b608435606435610959826103ae565b60a435925f516020615b5b5f395f51905f5254956001600160401b0361099a61098d6109898a60ff9060401c1690565b1590565b986001600160401b031690565b1680159081610a53575b6001149081610a49575b159081610a40575b50610a31576109d195876109c8612be8565b610a2457612c36565b6109d757005b6109f55f516020615b5b5f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a2c612c15565b612c36565b63f92ee8a960e01b5f5260045ffd5b9050155f6109b6565b303b1591506109ae565b8891506109a4565b5f9103126103aa57565b346103aa575f3660031901126103aa57602060ff5f5160206159fb5f395f51905f525416604051908152f35b346103aa5760603660031901126103aa57602435600435604435610ab4816103ae565b610abc613570565b610ac461372a565b815f526007602052610af083610aeb8160405f206001915f520160205260405f2054151590565b612d6c565b80610afb8484614b71565b916001600160a01b031615610bb9575b8151915f516020615a1b5f395f51905f526040820192610b52610b4060018060a01b0386511696836080870198895192613aae565b610b4981611e7b565b600181146144f6565b825160209384015194519551604080516001600160a01b03948516815295860191909152429085015294169391606090a47fb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db15f80a35f5f516020615b3b5f395f51905f525d005b5060608101516001600160a01b0316610b0b565b346103aa575f3660031901126103aa575f546040516001600160a01b039091168152602090f35b9181601f840112156103aa578235916001600160401b0383116103aa57602083818601950101116103aa57565b6101c03660031901126103aa57602435600435610c3d826103ae565b604435610c49816103ae565b6064359060a43560843560c4356001600160401b0381116103aa57610c72903690600401610bf4565b94909360e03660e31901126103aa576107539761074197612d86565b346103aa5760403660031901126103aa57610d89602435600435610cb0613703565b610cb861372a565b805f526007602052610cdf82610aeb8160405f206001915f520160205260405f2054151590565b610d846040610d0f610d0a85610cfd865f52600860205260405f2090565b905f5260205260405f2090565b613147565b805182810151610d7191610d34916080906001600160a01b03165b910151908761394f565b50610d3e81611e7b565b610d4781611e7b565b83516020810182905290600190610d6b83604081015b03601f198101855284610560565b14613180565b01518015908115610d91575b42916131ac565b614519565b6104d861375f565b4281109150610d7d565b6001600160401b03811161050a57601f01601f191660200190565b929192610dc282610d9b565b91610dd06040519384610560565b8294818452818301116103aa578281602093845f960137010152565b9080601f830112156103aa57816020610e0793359101610db6565b90565b60403660031901126103aa57600435610e22816103ae565b6024356001600160401b0381116103aa57610e41903690600401610dec565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f40575b50610f3157610e84613501565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610f00575b50610ecd57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615a9b5f395f51905f528303610eec576104d89250615289565b632a87526960e21b5f52600483905260245ffd5b610f2391945060203d602011610f2a575b610f1b8183610560565b810190614969565b925f610eac565b503d610f11565b63703e46dd60e11b5f5260045ffd5b5f516020615a9b5f395f51905f52546001600160a01b0316141590505f610e77565b346103aa5760603660031901126103aa57602435600435604435610f84613570565b815f526007602052610fab83610aeb8160405f206001915f520160205260405f2054151590565b4201804211610ffa5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612fc2565b346103aa575f3660031901126103aa577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f315760206040515f516020615a9b5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110ce5750505090565b909192602060e0826110e36001948851611056565b0194019291016110c1565b346103aa5760203660031901126103aa57600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106111cb57505061113d92500383610560565b6111478251613200565b915f5b81518110156111bd576001906111a161119c61116e865f52600660205260405f2090565b61118861117b8588612bd4565b516001600160a01b031690565b60018060a01b03165f5260205260405f2090565b61324f565b6111ab8287612bd4565b526111b68186612bd4565b500161114a565b6040518061075386826110ab565b8454835260019485019487945060209093019201611128565b346103aa575f3660031901126103aa57602060ff5f516020615b1b5f395f51905f5254166040519015158152f35b60e03660031901126103aa5760243560043561122d826103ae565b6044359161123a836103ae565b606435926084359060a4359060c435936001600160401b0385116103aa576112eb966112a06112706112e1973690600401610bf4565b95909661127b613703565b6001600160a01b038516948490611292878d6145c4565b61129a61372a565b8b61467b565b939092604051986112b08a6104ee565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610db6565b60e0820152614841565b5f5f516020615b3b5f395f51905f525d60405160018152602090f35b346103aa5760803660031901126103aa57602435600435611327826103ae565b604435906001600160401b0382116103aa57366023830112156103aa576107539261135f611372933690602481600401359101610db6565b906064359261136d846105d7565b6132d7565b6040516001600160a01b0390911681529081906020820190565b90602080835192838152019201905f5b8181106113a95750505090565b825184526020938401939092019160010161139c565b906020610e0792818152019061138c565b346103aa5760203660031901126103aa576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061142a576107538561141e81870382610560565b604051918291826113bf565b8254845260209093019260019283019201611407565b60e0810192916105919190611056565b346103aa5760403660031901126103aa5761075361149f602435600435611476826103ae565b61147e6131ca565b505f52600660205260405f209060018060a01b03165f5260205260405f2090565b6004604051916114ae8361050f565b80546001600160a01b039081168452600182015490811660208501526114f8906114ef906114e660a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611440565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916115789061156a610e0797959693600f60f81b865260e0602087015260e086019061151f565b90848203604086015261151f565b60608301949094526001600160a01b031660808201525f60a082015280830360c09091015261138c565b346103aa575f3660031901126103aa575f516020615a5b5f395f51905f52541580611636575b156115f9576115d5614978565b6115dd614a32565b906107536115e96133a0565b6040519384933091469186611543565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615b7b5f395f51905f5254156115c8565b9080601f830112156103aa578135611663816105c0565b926116716040519485610560565b81845260208085019260051b820101918383116103aa5760208201905b83821061169d57505050505090565b81356001600160401b0381116103aa576020916116bf87848094880101610649565b81520191019061168e565b60803660031901126103aa576004356001600160401b0381116103aa576116f5903690600401610430565b90602435906001600160401b0382116103aa57366023830112156103aa578160040135611721816105c0565b9261172f6040519485610560565b8184526024602085019260051b820101903682116103aa5760248101925b8284106117a457505050506044356001600160401b0381116103aa5761177790369060040161164c565b606435926001600160401b0384116103aa576107539461179e61074195369060040161164c565b936133bb565b83356001600160401b0381116103aa576020916117c88392602436918701016105e2565b81520193019261174d565b346103aa5760603660031901126103aa576004356117f0816103ae565b602435906117fd826103ae565b604435611809816105d7565b5f516020615b5b5f395f51905f5254926001600160401b0361183a60ff604087901c1615956001600160401b031690565b168015908161188d575b6001149081611883575b15908161187a575b50610a31576109d19284611868612be8565b1561407057611875612c15565b614070565b9050155f611856565b303b15915061184e565b859150611844565b346103aa5760403660031901126103aa576024356004356118b4613501565b6118bc61372a565b6118c68282614b71565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615b3b5f395f51905f525d005b346103aa575f3660031901126103aa576033546040516001600160a01b039091168152602090f35b346103aa575f3660031901126103aa576020603454604051908152f35b346103aa5760403660031901126103aa57602060ff611994602435600435611969826103ae565b5f525f516020615afb5f395f51905f52845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b801515036103aa57565b346103aa5760603660031901126103aa576024356004356119ca826103ae565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea60206044356119f9816119a0565b611a016135df565b835f5260058252611a7d816001611a5f60405f2098611a3f81611a3a858060a01b038216809d6001915f520160205260405f2054151590565b613451565b885f526006875260405f209060018060a01b03165f5260205260405f2090565b01805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6040519015158152a3005b346103aa5760403660031901126103aa576004356001600160401b0381116103aa57611ab8903690600401610649565b6024356001600160401b0381116103aa57611ad7903690600401610649565b90611ae5815183511461255d565b5f5b82518110156104d85780611bad611b0060019385612bd4565b51611b0b8387612bd4565b5190611b15613703565b611b1d61372a565b805f526007602052611b4482610aeb8160405f206001915f520160205260405f2054151590565b610d846040611b62610d0a85610cfd865f52600860205260405f2090565b805182810151610d7191611b81916080906001600160a01b0316610d2a565b50611b8b81611e7b565b611b9481611e7b565b835160208101829052908a90610d6b8360408101610d5d565b611bb561375f565b01611ae7565b346103aa575f3660031901126103aa5760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611bf85750505090565b82516001600160a01b0316845260209384019390920191600101611beb565b346103aa5760203660031901126103aa576004355f525f516020615abb5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611c7d5761075385611c7181870382610560565b60405191829182611bd5565b8254845260209093019260019283019201611c5a565b346103aa5760203660031901126103aa577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611cd2613501565b80600155604051908152a1005b346103aa5760403660031901126103aa57602435600435611cfe613570565b611d06613570565b611d0e61372a565b805f526007602052611d3582610aeb8160405f206001915f520160205260405f2054151590565b611d3f8282614519565b7fb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db15f80a35f5f516020615b3b5f395f51905f525d005b60405190611d84602083610560565b5f8252565b346103aa575f3660031901126103aa57610753604051611daa604082610560565b60058152640352e302e360dc1b602082015260405191829160208352602083019061151f565b346103aa5760203660031901126103aa576004355f526004602052600160405f20015460018101809111610ffa57602090604051908152f35b346103aa57611e1736610809565b90611e258151835114612bbe565b5f5b82518110156104d85780611e60611e4060019385612bd4565b51838060a01b03611e518488612bd4565b5116906107bc6107b782612ba0565b5001611e27565b634e487b7160e01b5f52602160045260245ffd5b600a1115611e8557565b611e67565b90600a821015611e855752565b6020815260606040611efc60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061151f565b93611f0e602082015183860190611e8a565b015191015290565b346103aa5760403660031901126103aa57600435602435905f60408051611f3c8161052a565b611f44613479565b815282602082015201525f52600860205260405f20905f5260205261075360405f20600760405191611f758361052a565b611f7e81613043565b8352611f9460ff6006830154166020850161313b565b0154604082015260405191829182611e97565b346103aa5760203660031901126103aa577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611fe9816105d7565b611ff1613501565b16611ffd8115156134a9565b8060ff195f5160206159fb5f395f51905f525416175f5160206159fb5f395f51905f5255604051908152a1005b346103aa5760403660031901126103aa5760043560243561204a816119a0565b6120526135df565b612067825f52600360205260405f2054151590565b156120be5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526120b381600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103aa5760203660031901126103aa576004356120ee816119a0565b6120f66135df565b1561215457612103613703565b600160ff195f516020615b1b5f395f51905f525416175f516020615b1b5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615b1b5f395f51905f525460ff8116156121ab5760ff19165f516020615b1b5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103aa575f3660031901126103aa5760405180602060025491828152019060025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace905f5b818110612218576107538561141e81870382610560565b8254845260209093019260019283019201612201565b346103aa5760203660031901126103aa5760043561224b816103ae565b612253613501565b6001600160a01b0316612267811515612547565b603380546001600160a01b031916821790555f516020615adb5f395f51905f525f80a2005b346103aa5760403660031901126103aa576104d86024356004356122af826103ae565b6109086107b782612ba0565b346103aa5760203660031901126103aa576004355f526004602052600260405f20015460018101809111610ffa57602090604051908152f35b346103aa5760203660031901126103aa57600435612311816103ae565b612319613501565b5f546001600160a01b031680612394575061233e6001600160a01b03821615156134bf565b5f80546001600160a01b0319166001600160a01b0392831690811790915561236d905b6001600160a01b031690565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103aa5760803660031901126103aa576104d86004356024356123c9816119a0565b604435906123d6826103ae565b606435926123e3846103ae565b6123eb61364e565b614373565b346103aa575f3660031901126103aa576032546040516001600160a01b039091168152602090f35b346103aa5760403660031901126103aa57602435600435612438826103ae565b612440613501565b805f5260056020525f600461248b604083209461246a81611a3a60018060a01b0382168099615378565b8484526006602052604084209060018060a01b03165f5260205260405f2090565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103aa575f3660031901126103aa576124e06156d8565b6124e86157df565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261253960c082610560565b519020604051908152602090f35b1561254e57565b638219abe360e01b5f5260045ffd5b1561256457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125a95760051b8101359060fe19813603018212156103aa570190565b612573565b35610e07816103ae565b903590601e19813603018212156103aa57018035906001600160401b0382116103aa576020019181360383136103aa57565b91908110156125a95760e0020190565b908160209103126103aa5751610e07816119a0565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361268097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161260f565b94803561268c816103ae565b6001600160a01b031660e085015260208101356126a8816103ae565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff60808201356126dd816105d7565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561272b573d9061271282610d9b565b916127206040519384610560565b82523d5f602084013e565b606090565b604090610e0793928152816020820152019061151f565b909193929361275785841461255d565b5f945b83861061276957505050509050565b612774868585612587565b35602061278c816127868a8989612587565b016125ae565b61279c60606127868b8a8a612587565b926128128a86888a8c6127f660806127b5878486612587565b0135956127ee6127e48260a06127cc82888a612587565b01359560c06127dc83838b612587565b013597612587565b60e08101906125b8565b9690956125ea565b946040519a8b998a996326ae802b60e11b8b5260048b0161262f565b03815f305af19081612859575b5061284e578561282d612701565b60405163f495148b60e01b815291829161284a9160048401612730565b0390fd5b60019095019461275a565b6128799060203d811161287e575b6128718183610560565b8101906125fa565b61281f565b503d612867565b90929192612891613703565b61289961372a565b8135916128ae835f52600560205260405f2090565b906128f76128e560408301936128c6612361866125ae565b6001600160a01b03165f90815260019091016020526040902054151590565b6128f1612361856125ae565b90612aae565b612902343415612ad6565b612998612922855f526004602052600260405f2001600181540180915590565b9561293560208401358881998214612af4565b612941612361856125ae565b97606084019587612951886125ae565b9a6129908b61298260808a01359e8f9061296e60a08d018d6125b8565b929091604051978896602088019a8b612b12565b03601f198101835282610560565b5190206137a1565b6129ab866129a5846125ae565b866139b1565b90916001836129b981611e7b565b14612a89575b6129c883611e7b565b60018303612a2657505050906129f46129ee5f516020615a1b5f395f51905f52936125ae565b916125ae565b604080516001600160a01b039283168152602081019790975242908701521693606090a45b612a2161375f565b600190565b612a638394612a5e612a81947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a6995613e83565b6125ae565b926125ae565b9260405193849360018060a01b031698429185612b76565b0390a4612a19565b9150612aa887612a98856125ae565b612aa1876125ae565b9088613aae565b916129bf565b15612ab65750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612ade5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612afd575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0797959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161260f565b61059193606092969593608083019760018060a01b03168352602083015260408201520190611e8a565b5f525f516020615afb5f395f51905f52602052600160405f20015490565b15612bc557565b63031206d560e51b5f5260045ffd5b80518210156125a95760209160051b010190565b60016001600160401b03195f516020615b5b5f395f51905f525416175f516020615b5b5f395f51905f5255565b5f516020615b5b5f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612d5d576001600160a01b03841615612d4e57612c59614edb565b612c6d6001600160a01b0384161515612547565b6001600160a01b03811692612c83841515612547565b60ff831615612d3f57612d0892612cde612ce392612c9f614edb565b612ca7614edb565b612caf614edb565b60ff195f516020615b1b5f395f51905f5254165f516020615b1b5f395f51905f5255612cd9614edb565b614f06565b614f15565b612ceb6150db565b60018060a01b03166001600160601b0360a01b6033541617603355565b5f516020615adb5f395f51905f525f80a2612d2243603455565b612d2c81846141c2565b81612d3657505050565b610591926144a1565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612d745750565b6373a1390160e11b5f5260045260245ffd5b959394612e08919597939297612d9a613703565b612ddf6001600160a01b038816612db1818b6145c4565b612db961372a565b612dc961236161236160e46125ae565b811490612dd961236160e46125ae565b91612f7f565b612e00612df06123616101046125ae565b6001600160a01b038b1614612fac565b83878961467b565b9161012435612e3b81612e2486612e1f8787612fd6565b612fd6565b811015612e3587612e1f8888612fd6565b90612fe3565b603254612e50906001600160a01b0316612361565b90612e5c6101046125ae565b906101443592612e6d610164613001565b92610184356101a43591833b156103aa57604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612f7a57612f3a6112e198612f4a93612a199c612f60575b50612f2a612f0e6101046125ae565b91612f17610581565b9c8d526001600160a01b031660208d0152565b6001600160a01b031660408b0152565b6001600160a01b03166060890152565b608087015260a086015260c08501523691610db6565b80612f6e5f612f7493610560565b80610a5b565b5f612eff565b6126f6565b15612f88575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612fb357565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ffa57565b15612fec575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e07816105d7565b90600182811c92168015613039575b602083101461302557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161301a565b9060405161305081610545565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916130a88261300b565b808552916001811690811561311457506001146130d5575b505060a092916130d1910384610560565b0152565b5f908152602081209092505b8183106130f85750508101602001816130d16130c0565b60209193508060019154838589010152019101909184926130e1565b60ff191660208681019190915292151560051b850190920192508391506130d190506130c0565b600a821015611e855752565b906040516131548161052a565b60406007829461316381613043565b845261317960ff6006830154166020860161313b565b0154910152565b156131885750565b60405162461bcd60e51b81526020600482015290819061284a90602483019061151f565b156131b5575050565b637a88099160e11b5f5260045260245260445ffd5b604051906131d78261050f565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061320a826105c0565b6132176040519182610560565b8281528092613228601f19916105c0565b01905f5b82811061323857505050565b6020906132436131ca565b8282850101520161322c565b9060405161325c8161050f565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132a79061329e906114e6565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103aa5751610e07816103ae565b5f5490949392906001600160a01b0381161561339157604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff9061334090608487019061151f565b931691015203925af18015612f7a57610591925f91613362575b5080946134d5565b613384915060203d60201161338a575b61337c8183610560565b8101906132c2565b5f61335a565b503d613372565b6315aeca0d60e11b5f5260045ffd5b604051906133af602083610560565b5f808352366020840137565b9192949390938484511480613447575b8061343d575b6133da9061255d565b5f5b85811015613431578060051b8401359060be19853603018212156103aa5761342a60019261340a8389612bd4565b51613415848c612bd4565b51906134218589612bd4565b51928901612885565b50016133dc565b50945050505050600190565b50815185146133d1565b50848651146133cb565b156134595750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b6040519061348682610545565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134b057565b63f0f15b9160e01b5f5260045ffd5b156134c657565b6302bfb33d60e51b5f5260045ffd5b905f61059193926123eb61364e565b916134fd9183549060031b91821b915f19901b19161790565b9055565b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff161561353957565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b335f9081527fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143602052604090205460ff16156135a857565b63e2517d3f60e01b5f52336004527f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0960245260445ffd5b335f9081527f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456602052604090205460ff161561361757565b63e2517d3f60e01b5f52336004527f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92960245260445ffd5b335f9081527f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f602052604090205460ff161561368657565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b5f8181525f516020615afb5f395f51905f526020908152604080832033845290915290205460ff16156136ed5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615b1b5f395f51905f52541661371b57565b63d93c066560e01b5f5260045ffd5b5f516020615b3b5f395f51905f525c6137505760015f516020615b3b5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615b3b5f395f51905f525d565b1561377857565b63b3f07a3960e01b5f5260045ffd5b1561378f5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613917575b6137ba90613771565b6137db6137d55f5160206159fb5f395f51905f525460ff1690565b60ff1690565b956137e98488811015613787565b5f945f935f5b86811061380a57505050505050506105919192811015613787565b6138676138368361381961541b565b6042916040519161190160f01b8352600283015260228201522090565b61384a6138438489612bd4565b5160ff1690565b6138548487612bd4565b51906138608589612bd4565b5192614c26565b6001600160a01b038181169088161080613898575b61388a575b506001016137ef565b600198890198909650613881565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615afb5f395f51905f5260205261391261390b827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa55b9060018060a01b03165f5260205260405f2090565b5460ff1690565b61387c565b50855183146137b1565b1561392857565b6330d45fb160e01b5f5260045ffd5b908160209103126103aa5751600a8110156103aa5790565b915061399e61119c6060926139825f9561397460325460018060a01b03161515613921565b5f52600660205260405f2090565b6001600160a01b039091165f9081526020919091526040902090565b01516139a957600191565b506002905f90565b9190915f92613a0f6060613a0861119c6139ed6139d861236160325460018060a01b031690565b966139746001600160a01b0389161515613921565b6001600160a01b0386165f9081526020919091526040902090565b0151151590565b613aa357604051633f4bdec560e01b81526001600160a01b039190911660048201526024810192909252909290602090849060449082905f905af1928315612f7a575f93613a72575b50600183613a6581611e7b565b03613a6c57565b60019150565b613a9591935060203d602011613a9c575b613a8d8183610560565b810190613937565b915f613a58565b503d613a83565b505050506002905f90565b9291905f906001600160a01b031660018103613af9575050610989613ae59183613ad6611d75565b916001600160a01b0316614c8d565b613af257612a2191614cf6565b5050600590565b929094939181613b095750505050565b819293949550613b356001613b2a876138f6885f52600660205260405f2090565b015460a01c60ff1690565b15613b9d5760405163a9059cbb60e01b60208201526001600160a01b0390911660248201526044810191909152613b6f8160648101612982565b92613b7d6005945b82614c3e565b613b88575b50505090565b60019350613b9592614d43565b5f8080613b82565b6040516340c10f1960e01b60208201526001600160a01b0390911660248201526044810191909152613bd28160648101612982565b92613b7d600694613b77565b15613be65750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103aa5760405190613c1182610545565b819380358352602081013560208401526040810135613c2f816103ae565b6040840152613c406060820161077d565b60608401526080810135608084015260a0810135916001600160401b0383116103aa5760a092613c709201610dec565b910152565b818110613c80575050565b5f8155600101613c75565b90601f8211613c98575050565b610591915f5160206159db5f395f51905f525f5260205f20906020601f840160051c83019310613cd0575b601f0160051c0190613c75565b9091508190613cc3565b9190601f8111613ce957505050565b610591925f5260205f20906020601f840160051c83019310613cd057601f0160051c0190613c75565b90600a811015611e855760ff80198354169116179055565b8151805182556020810151600183015560408101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a00151805191929160058401916001600160401b03821161050a57613dae82613da8855461300b565b85613cda565b602090601f8311600114613e1357826007959360409593613de4935f92613e08575b50508160011b915f199060031b1c19161790565b90555b613e016020820151613df881611e7b565b60068601613d12565b0151910155565b015190505f80613dd0565b90601f19831691613e27855f5260205f2090565b925f5b818110613e6b5750926001928592600798966040989610613e53575b505050811b019055613de7565b01515f1960f88460031b161c191690555f8080613e46565b92936020600181928786015181550195019301613e2a565b916080613f21613f1260029361397487613f0d613ebd6134fd99833595865f526007602052613ec260405f206020870135948580926155e9565b613bde565b15613f2e57613ef6613ed660015442612fd6565b915b613eeb613ee3610593565b963690613bf8565b86526020860161313b565b6040840152610cfd855f52600860205260405f2090565b613d2a565b611188612361604088016125ae565b9301359201918254612fd6565b613ef65f91613ed8565b90613f43825f614d86565b9182613f4c5750565b5f80525f516020615abb5f395f51905f526020526001600160a01b0316613f93817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d6155e9565b15613f9b5750565b63d180cb3160e01b5f526004525f60245260445ffd5b919091613fbe8382614d86565b9283613fc8575050565b815f525f516020615abb5f395f51905f52602052613ff360405f209160018060a01b031680926155e9565b15613ffc575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161401e8382614e3b565b9283614028575050565b815f525f516020615abb5f395f51905f5260205261405360405f209160018060a01b03168092615378565b1561405c575050565b62a95f1b60e31b5f5260045260245260445ffd5b919061407a614edb565b61408e6001600160a01b0384161515612547565b6001600160a01b038116926140a4841515612547565b60ff831615612d3f576141499261411761411d926140c0614edb565b6140c8614edb565b6140d0614edb565b60ff195f516020615b1b5f395f51905f5254165f516020615b1b5f395f51905f52556140fa614edb565b614102614edb565b61410a614edb565b614112614edb565b613f38565b50614f15565b614125614edb565b6201518060015560018060a01b03166001600160601b0360a01b6033541617603355565b5f516020615adb5f395f51905f525f80a261059143603455565b60036060610591938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156141a25750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b906001916141d18115156134bf565b614302838060a01b038316926141e88415156134bf565b6141f1856134bf565b6141fa83615580565b614337575b6142238161421e81614219875f52600560205260405f2090565b614e28565b61419a565b61427c61422e6105b1565b6001600160a01b0383168152916001600160a01b038716602084015286151560408401525f60608401525f60808401525f60a08401525f60c08401526138f6855f52600660205260405f2090565b815181546001600160a01b0319166001600160a01b0391821617825560208301516001830180546040860151606087015190151560a01b60ff60a01b16939094166001600160b01b0319909116179190911791151560a81b60ff60a81b169190911790559060049060c0906080810151600285015560a081015160038501550151910155565b60405183151581527fe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db9080602081015b0390a4565b61436e6143426105a2565b8481525f60208201525f60408201525f6060820152614369855f52600460205260405f2090565b614163565b6141ff565b906143327fe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db919493946143a78415156134bf565b6001600160a01b03861694610741906143c18715156134bf565b6001600160a01b0381169761427c906143db8a15156134bf565b6143e488615580565b61445d575b6144038161421e816142198c5f52600560205260405f2090565b61442b61440e6105b1565b6001600160a01b0383168152936001600160a01b03166020850152565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526138f6885f52600660205260405f2090565b61448f6144686105a2565b8981525f60208201525f60408201525f60608201526143698a5f52600460205260405f2090565b6143e9565b91908203918211610ffa57565b5f9081526006602090815260408083206001600160a01b03909416835292905220600181015460a01c60ff16156144e4576003018054918201809211610ffa5755565b6004018054918203918211610ffa5755565b156144fe5750565b63290cd55f60e01b5f52600a811015611e855760045260245ffd5b9061452391614b71565b60018060a01b036060820151168151915f516020615a1b5f395f51905f526040820192614565610b4060018060a01b0386511696836080870198895192613aae565b825160209384015194519551604080516001600160a01b039485168152958601919091524290850152941693918060608101614332565b156145a45750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526145f681611a3a60405f2060018060a01b038316906001915f520160205260405f2054151590565b825f52600460205260ff600360405f200154166146375760ff600161462b836138f661059196975f52600660205260405f2090565b015460a81c161561459c565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103aa578051916040602083015192015190565b1561466c57565b6358e8878560e01b5f5260045ffd5b8260609161472f979596936146b961119c61469e845f52600660205260405f2090565b6001600160a01b0384165f9081526020919091526040902090565b6146c96109896040830151151590565b6147d1575b506032546146e4906001600160a01b0316612361565b916146f96001600160a01b0384161515613921565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612f7a575f955f905f93614793575b5090829161059194939681988515159586614788575b50508461477d575b505082614772575b5050614665565b101590505f8061476b565b101592505f80614763565b101594505f8061475b565b90506147be91965061059193925060603d6060116147ca575b6147b68183610560565b81019061464a565b91969293919291614745565b503d6147ac565b60c06147e39101518480821015612fe3565b5f6146ce565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161483c919086019061151f565b930152565b61485f81515f526004602052600160405f2001600181540180915590565b61487282515f52600660205260405f2090565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086148c461236160016148b66020880195611188612361885160018060a01b031690565b01546001600160a01b031690565b8451835191959091614332906001600160a01b0316604083018051919390916001600160a01b0316906149156080820196875160a084019485519861490f60c087019a8b5190612fd6565b93615136565b80519751614934906001600160a01b031693516001600160a01b031690565b606082015190959060e0906001600160a01b03169751935191519201519260405197889760018060a01b03169c4296896147e9565b908160209103126103aa575190565b604051905f825f5160206159db5f395f51905f5254916149978361300b565b8083529260018116908115614a1357506001146149bb575b61059192500383610560565b505f5160206159db5f395f51905f525f90815290915f516020615a7b5f395f51905f525b8183106149f7575050906020610591928201016149af565b60209193508060019154838589010152019101909184926149df565b6020925061059194915060ff191682840152151560051b8201016149af565b604051905f825f516020615a3b5f395f51905f525491614a518361300b565b8083529260018116908115614a135750600114614a745761059192500383610560565b505f516020615a3b5f395f51905f525f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b818310614ac3575050906020610591928201016149af565b6020919350806001915483858901015201910190918492614aab565b60075f9182815582600182015582600282015582600382015582600482015560058101614b0c815461300b565b9081614b1f575b50508260068201550155565b601f8211600114614b3657849055505b5f80614b13565b614b5c614b6c926001601f614b4e855f5260205f2090565b920160051c82019101613c75565b5f81815260208120918190559055565b614b2f565b9190614b7b613479565b50825f526007602052614b918160405f20615378565b15614c1457614c0f61059191845f52600860205260405f20815f52602052610cfd614bbe60405f20613043565b95614beb614bd4825f52600660205260405f2090565b6040890151611188906001600160a01b0316612361565b614bff600260808a01519201918254614494565b90555f52600860205260405f2090565b614adf565b637f11bea960e01b5f5260045260245ffd5b91610e079391614c3593615482565b90929192615504565b81516001600160a01b03909116915f9182916020018285620186a0f1614c62612701565b9015614c87578051614c7e57503b15614c7a57600190565b5f90565b60209150015190565b50505f90565b814710614ce75782516001600160a01b03909116925f9182916020018486620186a0f190614cb9612701565b9115614ce05715614ccc575b5050600190565b8051614c7e57503b15614c7a575f80614cc5565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614d31576003018054918203918211610ffa5755565b6004018054918201809211610ffa5755565b5f9081526006602090815260408083206001600160a01b03909416835292905220600181015460a01c60ff1615614d31576003018054918203918211610ffa5755565b5f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16614c87575f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b03861684529091529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610e07916001600160a01b0316906155e9565b5f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615614c87575f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b03861684529091529020805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615b5b5f395f51905f525460401c1615614ef757565b631afcd79f60e31b5f5260045ffd5b614f1290614102614edb565b50565b90614f1e614edb565b614f2c60ff831615156134a9565b604091825192614f3c8185610560565b60098452682b30b634b230ba37b960b91b6020850152614f5e81519182610560565b60058152640312e302e360dc1b6020820152614f78614edb565b614f80614edb565b83516001600160401b03811161050a57614fb081614fab5f5160206159db5f395f51905f525461300b565b613c8b565b6020601f821160011461504c5781614ffa9392614fe69261059197985f92613e085750508160011b915f199060031b1c19161790565b5f5160206159db5f395f51905f5255615811565b61500f5f5f516020615a5b5f395f51905f5255565b6150245f5f516020615b7b5f395f51905f5255565b60ff1660ff195f5160206159fb5f395f51905f525416175f5160206159fb5f395f51905f5255565b5f5160206159db5f395f51905f525f52601f198216955f516020615a7b5f395f51905f52965f5b8181106150c35750966001928492614ffa9695610591999a106150ab575b505050811b015f5160206159db5f395f51905f5255615811565b01515f1960f88460031b161c191690555f8080615091565b83830151895560019098019760209384019301615073565b6150e3614edb565b62015180600155565b156150f657505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561512757565b63559d141b60e11b5f5260045ffd5b90936001600160a01b03851692600184036151a65750610591945061516c61515e8286612fd6565b34143490612e358488612fd6565b80615178575b506144a1565b6033546151a09161519b916001600160a01b031690615195611d75565b91614c8d565b615120565b5f615172565b906151b2343415612ad6565b6151c76151bf8287612fd6565b308489615636565b80615263575b506151ec6109896001613b2a866138f6875f52600660205260405f2090565b6151fc575b5061059193506144a1565b604051632770a7eb60e21b8152306004820152602481018590526020816044815f885af1918215612f7a576105919661523e9387935f91615244575b506150ec565b5f6151f1565b61525d915060203d60201161287e576128718183610560565b5f615238565b603354615283919061527d906001600160a01b0316612361565b8761567f565b5f6151cd565b90813b15615307575f516020615a9b5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156152ef57614f12916156bb565b5050346152f857565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125a9575f5260205f2001905f90565b80548015615364575f1901906153538282615328565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615413575f198401848111610ffa5783545f19810194908511610ffa575f9585836153c697610cfd95036153cc575b50505061533d565b55600190565b6153fc6153f6916153ed6153e361540a9588615328565b90549060031b1c90565b92839187615328565b906134e4565b85905f5260205260405f2090565b555f80806153be565b505050505f90565b6154236156d8565b61542b6157df565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261547c60c082610560565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116154ef579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612f7a575f516001600160a01b038116156154e557905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611e8557565b61550d816154fa565b80615516575050565b61551f816154fa565b600181036155365763f645eedf60e01b5f5260045ffd5b61553f816154fa565b6002810361555a575063fce698f760e01b5f5260045260245ffd5b806155666003926154fa565b1461556e5750565b6335e2f38360e21b5f5260045260245ffd5b5f818152600360205260409020546155e457600254600160401b81101561050a576155cd6155b78260018594016002556002615328565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b5f828152600182016020526040902054614c8757805490600160401b82101561050a57826156216155b7846001809601855584615328565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105919161567a608483610560565b615924565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526105919161567a606483610560565b5f80610e0793602081519101845af46156d2612701565b9161597c565b6040515f5160206159db5f395f51905f5254905f816156f68461300b565b9182825260208201946001811690815f146157c3575060011461576b575b61572092500382610560565b5190811561572c572090565b50505f516020615a5b5f395f51905f525480156157465790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b505f5160206159db5f395f51905f525f90815290915f516020615a7b5f395f51905f525b8183106157a757505090602061572092820101615714565b602091935080600191548385880101520191019091839261578f565b60ff191686525061572092151560051b82016020019050615714565b6157e7614a32565b80519081156157f7576020012090565b50505f516020615b7b5f395f51905f525480156157465790565b9081516001600160401b03811161050a576158508161583d5f516020615a3b5f395f51905f525461300b565b5f516020615a3b5f395f51905f52613cda565b602092601f82116001146158905761587f929382915f92613e085750508160011b915f199060031b1c19161790565b5f516020615a3b5f395f51905f5255565b5f516020615a3b5f395f51905f525f52601f198216937f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75915f5b86811061590c57508360019596106158f4575b505050811b015f516020615a3b5f395f51905f5255565b01515f1960f88460031b161c191690555f80806158dd565b919260206001819286850151815501940192016158ca565b905f602091828151910182855af1156126f6575f513d61597357506001600160a01b0381163b155b6159535750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b6001141561594c565b906159a0575080511561599157805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806159d1575b6159b1575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156159a956fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101a2646970667358221220dd2946cd6b9e9bf63ead699bdaa7434d723e0b864cfb6417972e7a22cdf00c6264736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeMetaData.ABI instead.
var BSCBridgeABI = BSCBridgeMetaData.ABI

// BSCBridgeBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BSCBridgeBinRuntime = "60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146103545780630b1d8d061461034f5780631313996b1461034a5780631938e0f214610345578063248a9ca3146103405780632f2ff15d1461033b57806336568abe14610336578063365d71e4146103315780633aefddbf1461032c57806342cde4e81461032757806348a00ed8146103225780634be13f831461031d5780634d5d0056146103185780634ee078ba146103135780634f1ef2861461030e578063502cc5c01461030957806352d1902d146103045780635b605f5c146102ff5780635c975abb146102fa5780635fd262de146102f5578063670e6268146102f057806379214874146102eb578063814914b5146102e657806384b0196e146102e157806388d67d6d146102dc57806389232a00146102d75780638ae87c5c146102d257806391cca3db146102cd57806391cf6d3e146102c857806391d14854146102c357806394ddc8c6146102be5780639f9b4f1c146102b9578063a217fddf146102b4578063a3246ad3146102af578063aa1bd0bc146102aa578063aa20ed47146102a5578063ad3cb1cc146102a0578063ae6893f81461029b578063b2b49e2e14610296578063b33eb36e14610291578063b7f3358d1461028c578063b8aa837e14610287578063bedb86fb14610282578063cf56118e1461027d578063d477f05f14610278578063d547741f14610273578063d5717fc51461026e578063e2af6cd714610269578063edbbea3914610264578063f0702e8e1461025f578063f45096431461025a5763f698da2514610255575f80fd5b6124c8565b612418565b6123f0565b6123a6565b6122f4565b6122bb565b61228c565b61222e565b6121ba565b6120d1565b61202a565b611fa7565b611f16565b611e09565b611dd0565b611d89565b611cdf565b611c93565b611c17565b611bbb565b611a88565b6119aa565b611942565b611925565b6118fd565b611895565b6117d3565b6116ca565b6115a2565b611450565b6113d0565b611307565b611212565b6111e4565b6110ee565b610fff565b610f62565b610e0a565b610c8e565b610c21565b610bcd565b610a91565b610a65565b610914565b6108b1565b6107c1565b610788565b610757565b6106a6565b610460565b6103bf565b346103aa5760203660031901126103aa5760043563ffffffff60e01b81168091036103aa57602090637965db0b60e01b8114908115610399575b506040519015158152f35b6301ffc9a760e01b1490505f61038e565b5f80fd5b6001600160a01b038116036103aa57565b346103aa5760203660031901126103aa576004356103dc816103ae565b6103e4613501565b6001600160a01b03166103f8811515612547565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103aa578235916001600160401b0383116103aa576020808501948460051b0101116103aa57565b60403660031901126103aa576004356001600160401b0381116103aa5761048b903690600401610430565b602435916001600160401b0383116103aa57366023840112156103aa578260040135916001600160401b0383116103aa5736602460e08502860101116103aa5760246104d8940191612747565b005b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761050a57604052565b6104da565b60e081019081106001600160401b0382111761050a57604052565b606081019081106001600160401b0382111761050a57604052565b60c081019081106001600160401b0382111761050a57604052565b90601f801991011681019081106001600160401b0382111761050a57604052565b6040519061059161010083610560565b565b60405190610591606083610560565b60405190610591608083610560565b6040519061059160e083610560565b6001600160401b03811161050a5760051b60200190565b60ff8116036103aa57565b9080601f830112156103aa5781356105f9816105c0565b926106076040519485610560565b81845260208085019260051b8201019283116103aa57602001905b82821061062f5750505090565b60208091833561063e816105d7565b815201910190610622565b9080601f830112156103aa578135610660816105c0565b9261066e6040519485610560565b81845260208085019260051b8201019283116103aa57602001905b8282106106965750505090565b8135815260209182019101610689565b60803660031901126103aa576004356001600160401b0381116103aa5760c060031982360301126103aa576024356001600160401b0381116103aa576106f09036906004016105e2565b906044356001600160401b0381116103aa57610710903690600401610649565b90606435916001600160401b0383116103aa5761075393610738610741943690600401610649565b92600401612885565b60405190151581529081906020820190565b0390f35b346103aa5760203660031901126103aa576020610775600435612ba0565b604051908152f35b3590610591826103ae565b346103aa5760403660031901126103aa576104d86024356004356107ab826103ae565b6107bc6107b782612ba0565b6136bd565b613fb1565b346103aa5760403660031901126103aa576004356024356107e1816103ae565b336001600160a01b038216036107fa576104d891614011565b63334bd91960e11b5f5260045ffd5b9060406003198301126103aa576004356001600160401b0381116103aa578261083491600401610649565b91602435906001600160401b0382116103aa57806023830112156103aa578160040135610860816105c0565b9261086e6040519485610560565b8184526024602085019260051b8201019283116103aa57602401905b8282106108975750505090565b6020809183356108a6816103ae565b81520191019061088a565b346103aa576108bf36610809565b906108cd8151835114612bbe565b5f5b82518110156104d8578061090d6108e860019385612bd4565b51838060a01b036108f98488612bd4565b5116906109086107b782612ba0565b614011565b50016108cf565b346103aa5760c03660031901126103aa57600435610931816103ae565b6024359061093e826103ae565b60443561094a816105d7565b608435606435610959826103ae565b60a435925f516020615b5b5f395f51905f5254956001600160401b0361099a61098d6109898a60ff9060401c1690565b1590565b986001600160401b031690565b1680159081610a53575b6001149081610a49575b159081610a40575b50610a31576109d195876109c8612be8565b610a2457612c36565b6109d757005b6109f55f516020615b5b5f395f51905f52805460ff60401b19169055565b604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b610a2c612c15565b612c36565b63f92ee8a960e01b5f5260045ffd5b9050155f6109b6565b303b1591506109ae565b8891506109a4565b5f9103126103aa57565b346103aa575f3660031901126103aa57602060ff5f5160206159fb5f395f51905f525416604051908152f35b346103aa5760603660031901126103aa57602435600435604435610ab4816103ae565b610abc613570565b610ac461372a565b815f526007602052610af083610aeb8160405f206001915f520160205260405f2054151590565b612d6c565b80610afb8484614b71565b916001600160a01b031615610bb9575b8151915f516020615a1b5f395f51905f526040820192610b52610b4060018060a01b0386511696836080870198895192613aae565b610b4981611e7b565b600181146144f6565b825160209384015194519551604080516001600160a01b03948516815295860191909152429085015294169391606090a47fb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db15f80a35f5f516020615b3b5f395f51905f525d005b5060608101516001600160a01b0316610b0b565b346103aa575f3660031901126103aa575f546040516001600160a01b039091168152602090f35b9181601f840112156103aa578235916001600160401b0383116103aa57602083818601950101116103aa57565b6101c03660031901126103aa57602435600435610c3d826103ae565b604435610c49816103ae565b6064359060a43560843560c4356001600160401b0381116103aa57610c72903690600401610bf4565b94909360e03660e31901126103aa576107539761074197612d86565b346103aa5760403660031901126103aa57610d89602435600435610cb0613703565b610cb861372a565b805f526007602052610cdf82610aeb8160405f206001915f520160205260405f2054151590565b610d846040610d0f610d0a85610cfd865f52600860205260405f2090565b905f5260205260405f2090565b613147565b805182810151610d7191610d34916080906001600160a01b03165b910151908761394f565b50610d3e81611e7b565b610d4781611e7b565b83516020810182905290600190610d6b83604081015b03601f198101855284610560565b14613180565b01518015908115610d91575b42916131ac565b614519565b6104d861375f565b4281109150610d7d565b6001600160401b03811161050a57601f01601f191660200190565b929192610dc282610d9b565b91610dd06040519384610560565b8294818452818301116103aa578281602093845f960137010152565b9080601f830112156103aa57816020610e0793359101610db6565b90565b60403660031901126103aa57600435610e22816103ae565b6024356001600160401b0381116103aa57610e41903690600401610dec565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f40575b50610f3157610e84613501565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610f00575b50610ecd57634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615a9b5f395f51905f528303610eec576104d89250615289565b632a87526960e21b5f52600483905260245ffd5b610f2391945060203d602011610f2a575b610f1b8183610560565b810190614969565b925f610eac565b503d610f11565b63703e46dd60e11b5f5260045ffd5b5f516020615a9b5f395f51905f52546001600160a01b0316141590505f610e77565b346103aa5760603660031901126103aa57602435600435604435610f84613570565b815f526007602052610fab83610aeb8160405f206001915f520160205260405f2054151590565b4201804211610ffa5760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b612fc2565b346103aa575f3660031901126103aa577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610f315760206040515f516020615a9b5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b8181106110ce5750505090565b909192602060e0826110e36001948851611056565b0194019291016110c1565b346103aa5760203660031901126103aa57600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b8181106111cb57505061113d92500383610560565b6111478251613200565b915f5b81518110156111bd576001906111a161119c61116e865f52600660205260405f2090565b61118861117b8588612bd4565b516001600160a01b031690565b60018060a01b03165f5260205260405f2090565b61324f565b6111ab8287612bd4565b526111b68186612bd4565b500161114a565b6040518061075386826110ab565b8454835260019485019487945060209093019201611128565b346103aa575f3660031901126103aa57602060ff5f516020615b1b5f395f51905f5254166040519015158152f35b60e03660031901126103aa5760243560043561122d826103ae565b6044359161123a836103ae565b606435926084359060a4359060c435936001600160401b0385116103aa576112eb966112a06112706112e1973690600401610bf4565b95909661127b613703565b6001600160a01b038516948490611292878d6145c4565b61129a61372a565b8b61467b565b939092604051986112b08a6104ee565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610db6565b60e0820152614841565b5f5f516020615b3b5f395f51905f525d60405160018152602090f35b346103aa5760803660031901126103aa57602435600435611327826103ae565b604435906001600160401b0382116103aa57366023830112156103aa576107539261135f611372933690602481600401359101610db6565b906064359261136d846105d7565b6132d7565b6040516001600160a01b0390911681529081906020820190565b90602080835192838152019201905f5b8181106113a95750505090565b825184526020938401939092019160010161139c565b906020610e0792818152019061138c565b346103aa5760203660031901126103aa576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b81811061142a576107538561141e81870382610560565b604051918291826113bf565b8254845260209093019260019283019201611407565b60e0810192916105919190611056565b346103aa5760403660031901126103aa5761075361149f602435600435611476826103ae565b61147e6131ca565b505f52600660205260405f209060018060a01b03165f5260205260405f2090565b6004604051916114ae8361050f565b80546001600160a01b039081168452600182015490811660208501526114f8906114ef906114e660a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c082015260405191829182611440565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b916115789061156a610e0797959693600f60f81b865260e0602087015260e086019061151f565b90848203604086015261151f565b60608301949094526001600160a01b031660808201525f60a082015280830360c09091015261138c565b346103aa575f3660031901126103aa575f516020615a5b5f395f51905f52541580611636575b156115f9576115d5614978565b6115dd614a32565b906107536115e96133a0565b6040519384933091469186611543565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615b7b5f395f51905f5254156115c8565b9080601f830112156103aa578135611663816105c0565b926116716040519485610560565b81845260208085019260051b820101918383116103aa5760208201905b83821061169d57505050505090565b81356001600160401b0381116103aa576020916116bf87848094880101610649565b81520191019061168e565b60803660031901126103aa576004356001600160401b0381116103aa576116f5903690600401610430565b90602435906001600160401b0382116103aa57366023830112156103aa578160040135611721816105c0565b9261172f6040519485610560565b8184526024602085019260051b820101903682116103aa5760248101925b8284106117a457505050506044356001600160401b0381116103aa5761177790369060040161164c565b606435926001600160401b0384116103aa576107539461179e61074195369060040161164c565b936133bb565b83356001600160401b0381116103aa576020916117c88392602436918701016105e2565b81520193019261174d565b346103aa5760603660031901126103aa576004356117f0816103ae565b602435906117fd826103ae565b604435611809816105d7565b5f516020615b5b5f395f51905f5254926001600160401b0361183a60ff604087901c1615956001600160401b031690565b168015908161188d575b6001149081611883575b15908161187a575b50610a31576109d19284611868612be8565b1561407057611875612c15565b614070565b9050155f611856565b303b15915061184e565b859150611844565b346103aa5760403660031901126103aa576024356004356118b4613501565b6118bc61372a565b6118c68282614b71565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615b3b5f395f51905f525d005b346103aa575f3660031901126103aa576033546040516001600160a01b039091168152602090f35b346103aa575f3660031901126103aa576020603454604051908152f35b346103aa5760403660031901126103aa57602060ff611994602435600435611969826103ae565b5f525f516020615afb5f395f51905f52845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b801515036103aa57565b346103aa5760603660031901126103aa576024356004356119ca826103ae565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea60206044356119f9816119a0565b611a016135df565b835f5260058252611a7d816001611a5f60405f2098611a3f81611a3a858060a01b038216809d6001915f520160205260405f2054151590565b613451565b885f526006875260405f209060018060a01b03165f5260205260405f2090565b01805460ff60a81b191691151560a81b60ff60a81b16919091179055565b6040519015158152a3005b346103aa5760403660031901126103aa576004356001600160401b0381116103aa57611ab8903690600401610649565b6024356001600160401b0381116103aa57611ad7903690600401610649565b90611ae5815183511461255d565b5f5b82518110156104d85780611bad611b0060019385612bd4565b51611b0b8387612bd4565b5190611b15613703565b611b1d61372a565b805f526007602052611b4482610aeb8160405f206001915f520160205260405f2054151590565b610d846040611b62610d0a85610cfd865f52600860205260405f2090565b805182810151610d7191611b81916080906001600160a01b0316610d2a565b50611b8b81611e7b565b611b9481611e7b565b835160208101829052908a90610d6b8360408101610d5d565b611bb561375f565b01611ae7565b346103aa575f3660031901126103aa5760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611bf85750505090565b82516001600160a01b0316845260209384019390920191600101611beb565b346103aa5760203660031901126103aa576004355f525f516020615abb5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611c7d5761075385611c7181870382610560565b60405191829182611bd5565b8254845260209093019260019283019201611c5a565b346103aa5760203660031901126103aa577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611cd2613501565b80600155604051908152a1005b346103aa5760403660031901126103aa57602435600435611cfe613570565b611d06613570565b611d0e61372a565b805f526007602052611d3582610aeb8160405f206001915f520160205260405f2054151590565b611d3f8282614519565b7fb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db15f80a35f5f516020615b3b5f395f51905f525d005b60405190611d84602083610560565b5f8252565b346103aa575f3660031901126103aa57610753604051611daa604082610560565b60058152640352e302e360dc1b602082015260405191829160208352602083019061151f565b346103aa5760203660031901126103aa576004355f526004602052600160405f20015460018101809111610ffa57602090604051908152f35b346103aa57611e1736610809565b90611e258151835114612bbe565b5f5b82518110156104d85780611e60611e4060019385612bd4565b51838060a01b03611e518488612bd4565b5116906107bc6107b782612ba0565b5001611e27565b634e487b7160e01b5f52602160045260245ffd5b600a1115611e8557565b611e67565b90600a821015611e855752565b6020815260606040611efc60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c061012086015261014085019061151f565b93611f0e602082015183860190611e8a565b015191015290565b346103aa5760403660031901126103aa57600435602435905f60408051611f3c8161052a565b611f44613479565b815282602082015201525f52600860205260405f20905f5260205261075360405f20600760405191611f758361052a565b611f7e81613043565b8352611f9460ff6006830154166020850161313b565b0154604082015260405191829182611e97565b346103aa5760203660031901126103aa577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff600435611fe9816105d7565b611ff1613501565b16611ffd8115156134a9565b8060ff195f5160206159fb5f395f51905f525416175f5160206159fb5f395f51905f5255604051908152a1005b346103aa5760403660031901126103aa5760043560243561204a816119a0565b6120526135df565b612067825f52600360205260405f2054151590565b156120be5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526120b381600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103aa5760203660031901126103aa576004356120ee816119a0565b6120f66135df565b1561215457612103613703565b600160ff195f516020615b1b5f395f51905f525416175f516020615b1b5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615b1b5f395f51905f525460ff8116156121ab5760ff19165f516020615b1b5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103aa575f3660031901126103aa5760405180602060025491828152019060025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace905f5b818110612218576107538561141e81870382610560565b8254845260209093019260019283019201612201565b346103aa5760203660031901126103aa5760043561224b816103ae565b612253613501565b6001600160a01b0316612267811515612547565b603380546001600160a01b031916821790555f516020615adb5f395f51905f525f80a2005b346103aa5760403660031901126103aa576104d86024356004356122af826103ae565b6109086107b782612ba0565b346103aa5760203660031901126103aa576004355f526004602052600260405f20015460018101809111610ffa57602090604051908152f35b346103aa5760203660031901126103aa57600435612311816103ae565b612319613501565b5f546001600160a01b031680612394575061233e6001600160a01b03821615156134bf565b5f80546001600160a01b0319166001600160a01b0392831690811790915561236d905b6001600160a01b031690565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103aa5760803660031901126103aa576104d86004356024356123c9816119a0565b604435906123d6826103ae565b606435926123e3846103ae565b6123eb61364e565b614373565b346103aa575f3660031901126103aa576032546040516001600160a01b039091168152602090f35b346103aa5760403660031901126103aa57602435600435612438826103ae565b612440613501565b805f5260056020525f600461248b604083209461246a81611a3a60018060a01b0382168099615378565b8484526006602052604084209060018060a01b03165f5260205260405f2090565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103aa575f3660031901126103aa576124e06156d8565b6124e86157df565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261253960c082610560565b519020604051908152602090f35b1561254e57565b638219abe360e01b5f5260045ffd5b1561256457565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156125a95760051b8101359060fe19813603018212156103aa570190565b612573565b35610e07816103ae565b903590601e19813603018212156103aa57018035906001600160401b0382116103aa576020019181360383136103aa57565b91908110156125a95760e0020190565b908160209103126103aa5751610e07816119a0565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c0979361268097939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c086019161260f565b94803561268c816103ae565b6001600160a01b031660e085015260208101356126a8816103ae565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff60808201356126dd816105d7565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d1561272b573d9061271282610d9b565b916127206040519384610560565b82523d5f602084013e565b606090565b604090610e0793928152816020820152019061151f565b909193929361275785841461255d565b5f945b83861061276957505050509050565b612774868585612587565b35602061278c816127868a8989612587565b016125ae565b61279c60606127868b8a8a612587565b926128128a86888a8c6127f660806127b5878486612587565b0135956127ee6127e48260a06127cc82888a612587565b01359560c06127dc83838b612587565b013597612587565b60e08101906125b8565b9690956125ea565b946040519a8b998a996326ae802b60e11b8b5260048b0161262f565b03815f305af19081612859575b5061284e578561282d612701565b60405163f495148b60e01b815291829161284a9160048401612730565b0390fd5b60019095019461275a565b6128799060203d811161287e575b6128718183610560565b8101906125fa565b61281f565b503d612867565b90929192612891613703565b61289961372a565b8135916128ae835f52600560205260405f2090565b906128f76128e560408301936128c6612361866125ae565b6001600160a01b03165f90815260019091016020526040902054151590565b6128f1612361856125ae565b90612aae565b612902343415612ad6565b612998612922855f526004602052600260405f2001600181540180915590565b9561293560208401358881998214612af4565b612941612361856125ae565b97606084019587612951886125ae565b9a6129908b61298260808a01359e8f9061296e60a08d018d6125b8565b929091604051978896602088019a8b612b12565b03601f198101835282610560565b5190206137a1565b6129ab866129a5846125ae565b866139b1565b90916001836129b981611e7b565b14612a89575b6129c883611e7b565b60018303612a2657505050906129f46129ee5f516020615a1b5f395f51905f52936125ae565b916125ae565b604080516001600160a01b039283168152602081019790975242908701521693606090a45b612a2161375f565b600190565b612a638394612a5e612a81947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612a6995613e83565b6125ae565b926125ae565b9260405193849360018060a01b031698429185612b76565b0390a4612a19565b9150612aa887612a98856125ae565b612aa1876125ae565b9088613aae565b916129bf565b15612ab65750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612ade5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612afd575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610e0797959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c0820152019161260f565b61059193606092969593608083019760018060a01b03168352602083015260408201520190611e8a565b5f525f516020615afb5f395f51905f52602052600160405f20015490565b15612bc557565b63031206d560e51b5f5260045ffd5b80518210156125a95760209160051b010190565b60016001600160401b03195f516020615b5b5f395f51905f525416175f516020615b5b5f395f51905f5255565b5f516020615b5b5f395f51905f52805460ff60401b1916600160401b179055565b9190949392948515612d5d576001600160a01b03841615612d4e57612c59614edb565b612c6d6001600160a01b0384161515612547565b6001600160a01b03811692612c83841515612547565b60ff831615612d3f57612d0892612cde612ce392612c9f614edb565b612ca7614edb565b612caf614edb565b60ff195f516020615b1b5f395f51905f5254165f516020615b1b5f395f51905f5255612cd9614edb565b614f06565b614f15565b612ceb6150db565b60018060a01b03166001600160601b0360a01b6033541617603355565b5f516020615adb5f395f51905f525f80a2612d2243603455565b612d2c81846141c2565b81612d3657505050565b610591926144a1565b6338854f1160e21b5f5260045ffd5b63643b674560e01b5f5260045ffd5b63ff5a231360e01b5f5260045ffd5b15612d745750565b6373a1390160e11b5f5260045260245ffd5b959394612e08919597939297612d9a613703565b612ddf6001600160a01b038816612db1818b6145c4565b612db961372a565b612dc961236161236160e46125ae565b811490612dd961236160e46125ae565b91612f7f565b612e00612df06123616101046125ae565b6001600160a01b038b1614612fac565b83878961467b565b9161012435612e3b81612e2486612e1f8787612fd6565b612fd6565b811015612e3587612e1f8888612fd6565b90612fe3565b603254612e50906001600160a01b0316612361565b90612e5c6101046125ae565b906101443592612e6d610164613001565b92610184356101a43591833b156103aa57604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af1988915612f7a57612f3a6112e198612f4a93612a199c612f60575b50612f2a612f0e6101046125ae565b91612f17610581565b9c8d526001600160a01b031660208d0152565b6001600160a01b031660408b0152565b6001600160a01b03166060890152565b608087015260a086015260c08501523691610db6565b80612f6e5f612f7493610560565b80610a5b565b5f612eff565b6126f6565b15612f88575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b15612fb357565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ffa57565b15612fec575050565b63943892eb60e01b5f5260045260245260445ffd5b35610e07816105d7565b90600182811c92168015613039575b602083101461302557565b634e487b7160e01b5f52602260045260245ffd5b91607f169161301a565b9060405161305081610545565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f916130a88261300b565b808552916001811690811561311457506001146130d5575b505060a092916130d1910384610560565b0152565b5f908152602081209092505b8183106130f85750508101602001816130d16130c0565b60209193508060019154838589010152019101909184926130e1565b60ff191660208681019190915292151560051b850190920192508391506130d190506130c0565b600a821015611e855752565b906040516131548161052a565b60406007829461316381613043565b845261317960ff6006830154166020860161313b565b0154910152565b156131885750565b60405162461bcd60e51b81526020600482015290819061284a90602483019061151f565b156131b5575050565b637a88099160e11b5f5260045260245260445ffd5b604051906131d78261050f565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b9061320a826105c0565b6132176040519182610560565b8281528092613228601f19916105c0565b01905f5b82811061323857505050565b6020906132436131ca565b8282850101520161322c565b9060405161325c8161050f565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c0916004916132a79061329e906114e6565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103aa5751610e07816103ae565b5f5490949392906001600160a01b0381161561339157604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff9061334090608487019061151f565b931691015203925af18015612f7a57610591925f91613362575b5080946134d5565b613384915060203d60201161338a575b61337c8183610560565b8101906132c2565b5f61335a565b503d613372565b6315aeca0d60e11b5f5260045ffd5b604051906133af602083610560565b5f808352366020840137565b9192949390938484511480613447575b8061343d575b6133da9061255d565b5f5b85811015613431578060051b8401359060be19853603018212156103aa5761342a60019261340a8389612bd4565b51613415848c612bd4565b51906134218589612bd4565b51928901612885565b50016133dc565b50945050505050600190565b50815185146133d1565b50848651146133cb565b156134595750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b6040519061348682610545565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156134b057565b63f0f15b9160e01b5f5260045ffd5b156134c657565b6302bfb33d60e51b5f5260045ffd5b905f61059193926123eb61364e565b916134fd9183549060031b91821b915f19901b19161790565b9055565b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff161561353957565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b335f9081527fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b143602052604090205460ff16156135a857565b63e2517d3f60e01b5f52336004527f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0960245260445ffd5b335f9081527f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb456602052604090205460ff161561361757565b63e2517d3f60e01b5f52336004527f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92960245260445ffd5b335f9081527f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f602052604090205460ff161561368657565b63e2517d3f60e01b5f52336004527f21d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c60245260445ffd5b5f8181525f516020615afb5f395f51905f526020908152604080832033845290915290205460ff16156136ed5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60ff5f516020615b1b5f395f51905f52541661371b57565b63d93c066560e01b5f5260045ffd5b5f516020615b3b5f395f51905f525c6137505760015f516020615b3b5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615b3b5f395f51905f525d565b1561377857565b63b3f07a3960e01b5f5260045ffd5b1561378f5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613917575b6137ba90613771565b6137db6137d55f5160206159fb5f395f51905f525460ff1690565b60ff1690565b956137e98488811015613787565b5f945f935f5b86811061380a57505050505050506105919192811015613787565b6138676138368361381961541b565b6042916040519161190160f01b8352600283015260228201522090565b61384a6138438489612bd4565b5160ff1690565b6138548487612bd4565b51906138608589612bd4565b5192614c26565b6001600160a01b038181169088161080613898575b61388a575b506001016137ef565b600198890198909650613881565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615afb5f395f51905f5260205261391261390b827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa55b9060018060a01b03165f5260205260405f2090565b5460ff1690565b61387c565b50855183146137b1565b1561392857565b6330d45fb160e01b5f5260045ffd5b908160209103126103aa5751600a8110156103aa5790565b915061399e61119c6060926139825f9561397460325460018060a01b03161515613921565b5f52600660205260405f2090565b6001600160a01b039091165f9081526020919091526040902090565b01516139a957600191565b506002905f90565b9190915f92613a0f6060613a0861119c6139ed6139d861236160325460018060a01b031690565b966139746001600160a01b0389161515613921565b6001600160a01b0386165f9081526020919091526040902090565b0151151590565b613aa357604051633f4bdec560e01b81526001600160a01b039190911660048201526024810192909252909290602090849060449082905f905af1928315612f7a575f93613a72575b50600183613a6581611e7b565b03613a6c57565b60019150565b613a9591935060203d602011613a9c575b613a8d8183610560565b810190613937565b915f613a58565b503d613a83565b505050506002905f90565b9291905f906001600160a01b031660018103613af9575050610989613ae59183613ad6611d75565b916001600160a01b0316614c8d565b613af257612a2191614cf6565b5050600590565b929094939181613b095750505050565b819293949550613b356001613b2a876138f6885f52600660205260405f2090565b015460a01c60ff1690565b15613b9d5760405163a9059cbb60e01b60208201526001600160a01b0390911660248201526044810191909152613b6f8160648101612982565b92613b7d6005945b82614c3e565b613b88575b50505090565b60019350613b9592614d43565b5f8080613b82565b6040516340c10f1960e01b60208201526001600160a01b0390911660248201526044810191909152613bd28160648101612982565b92613b7d600694613b77565b15613be65750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103aa5760405190613c1182610545565b819380358352602081013560208401526040810135613c2f816103ae565b6040840152613c406060820161077d565b60608401526080810135608084015260a0810135916001600160401b0383116103aa5760a092613c709201610dec565b910152565b818110613c80575050565b5f8155600101613c75565b90601f8211613c98575050565b610591915f5160206159db5f395f51905f525f5260205f20906020601f840160051c83019310613cd0575b601f0160051c0190613c75565b9091508190613cc3565b9190601f8111613ce957505050565b610591925f5260205f20906020601f840160051c83019310613cd057601f0160051c0190613c75565b90600a811015611e855760ff80198354169116179055565b8151805182556020810151600183015560408101516002830180546001600160a01b039283166001600160a01b03199182161790915560608301516003850180549190931691161790556080810151600483015560a00151805191929160058401916001600160401b03821161050a57613dae82613da8855461300b565b85613cda565b602090601f8311600114613e1357826007959360409593613de4935f92613e08575b50508160011b915f199060031b1c19161790565b90555b613e016020820151613df881611e7b565b60068601613d12565b0151910155565b015190505f80613dd0565b90601f19831691613e27855f5260205f2090565b925f5b818110613e6b5750926001928592600798966040989610613e53575b505050811b019055613de7565b01515f1960f88460031b161c191690555f8080613e46565b92936020600181928786015181550195019301613e2a565b916080613f21613f1260029361397487613f0d613ebd6134fd99833595865f526007602052613ec260405f206020870135948580926155e9565b613bde565b15613f2e57613ef6613ed660015442612fd6565b915b613eeb613ee3610593565b963690613bf8565b86526020860161313b565b6040840152610cfd855f52600860205260405f2090565b613d2a565b611188612361604088016125ae565b9301359201918254612fd6565b613ef65f91613ed8565b90613f43825f614d86565b9182613f4c5750565b5f80525f516020615abb5f395f51905f526020526001600160a01b0316613f93817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d6155e9565b15613f9b5750565b63d180cb3160e01b5f526004525f60245260445ffd5b919091613fbe8382614d86565b9283613fc8575050565b815f525f516020615abb5f395f51905f52602052613ff360405f209160018060a01b031680926155e9565b15613ffc575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161401e8382614e3b565b9283614028575050565b815f525f516020615abb5f395f51905f5260205261405360405f209160018060a01b03168092615378565b1561405c575050565b62a95f1b60e31b5f5260045260245260445ffd5b919061407a614edb565b61408e6001600160a01b0384161515612547565b6001600160a01b038116926140a4841515612547565b60ff831615612d3f576141499261411761411d926140c0614edb565b6140c8614edb565b6140d0614edb565b60ff195f516020615b1b5f395f51905f5254165f516020615b1b5f395f51905f52556140fa614edb565b614102614edb565b61410a614edb565b614112614edb565b613f38565b50614f15565b614125614edb565b6201518060015560018060a01b03166001600160601b0360a01b6033541617603355565b5f516020615adb5f395f51905f525f80a261059143603455565b60036060610591938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156141a25750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b906001916141d18115156134bf565b614302838060a01b038316926141e88415156134bf565b6141f1856134bf565b6141fa83615580565b614337575b6142238161421e81614219875f52600560205260405f2090565b614e28565b61419a565b61427c61422e6105b1565b6001600160a01b0383168152916001600160a01b038716602084015286151560408401525f60608401525f60808401525f60a08401525f60c08401526138f6855f52600660205260405f2090565b815181546001600160a01b0319166001600160a01b0391821617825560208301516001830180546040860151606087015190151560a01b60ff60a01b16939094166001600160b01b0319909116179190911791151560a81b60ff60a81b169190911790559060049060c0906080810151600285015560a081015160038501550151910155565b60405183151581527fe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db9080602081015b0390a4565b61436e6143426105a2565b8481525f60208201525f60408201525f6060820152614369855f52600460205260405f2090565b614163565b6141ff565b906143327fe51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db919493946143a78415156134bf565b6001600160a01b03861694610741906143c18715156134bf565b6001600160a01b0381169761427c906143db8a15156134bf565b6143e488615580565b61445d575b6144038161421e816142198c5f52600560205260405f2090565b61442b61440e6105b1565b6001600160a01b0383168152936001600160a01b03166020850152565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526138f6885f52600660205260405f2090565b61448f6144686105a2565b8981525f60208201525f60408201525f60608201526143698a5f52600460205260405f2090565b6143e9565b91908203918211610ffa57565b5f9081526006602090815260408083206001600160a01b03909416835292905220600181015460a01c60ff16156144e4576003018054918201809211610ffa5755565b6004018054918203918211610ffa5755565b156144fe5750565b63290cd55f60e01b5f52600a811015611e855760045260245ffd5b9061452391614b71565b60018060a01b036060820151168151915f516020615a1b5f395f51905f526040820192614565610b4060018060a01b0386511696836080870198895192613aae565b825160209384015194519551604080516001600160a01b039485168152958601919091524290850152941693918060608101614332565b156145a45750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f5260056020526145f681611a3a60405f2060018060a01b038316906001915f520160205260405f2054151590565b825f52600460205260ff600360405f200154166146375760ff600161462b836138f661059196975f52600660205260405f2090565b015460a81c161561459c565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103aa578051916040602083015192015190565b1561466c57565b6358e8878560e01b5f5260045ffd5b8260609161472f979596936146b961119c61469e845f52600660205260405f2090565b6001600160a01b0384165f9081526020919091526040902090565b6146c96109896040830151151590565b6147d1575b506032546146e4906001600160a01b0316612361565b916146f96001600160a01b0384161515613921565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115612f7a575f955f905f93614793575b5090829161059194939681988515159586614788575b50508461477d575b505082614772575b5050614665565b101590505f8061476b565b101592505f80614763565b101594505f8061475b565b90506147be91965061059193925060603d6060116147ca575b6147b68183610560565b81019061464a565b91969293919291614745565b503d6147ac565b60c06147e39101518480821015612fe3565b5f6146ce565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161483c919086019061151f565b930152565b61485f81515f526004602052600160405f2001600181540180915590565b61487282515f52600660205260405f2090565b7f18586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff22807086148c461236160016148b66020880195611188612361885160018060a01b031690565b01546001600160a01b031690565b8451835191959091614332906001600160a01b0316604083018051919390916001600160a01b0316906149156080820196875160a084019485519861490f60c087019a8b5190612fd6565b93615136565b80519751614934906001600160a01b031693516001600160a01b031690565b606082015190959060e0906001600160a01b03169751935191519201519260405197889760018060a01b03169c4296896147e9565b908160209103126103aa575190565b604051905f825f5160206159db5f395f51905f5254916149978361300b565b8083529260018116908115614a1357506001146149bb575b61059192500383610560565b505f5160206159db5f395f51905f525f90815290915f516020615a7b5f395f51905f525b8183106149f7575050906020610591928201016149af565b60209193508060019154838589010152019101909184926149df565b6020925061059194915060ff191682840152151560051b8201016149af565b604051905f825f516020615a3b5f395f51905f525491614a518361300b565b8083529260018116908115614a135750600114614a745761059192500383610560565b505f516020615a3b5f395f51905f525f90815290917f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b755b818310614ac3575050906020610591928201016149af565b6020919350806001915483858901015201910190918492614aab565b60075f9182815582600182015582600282015582600382015582600482015560058101614b0c815461300b565b9081614b1f575b50508260068201550155565b601f8211600114614b3657849055505b5f80614b13565b614b5c614b6c926001601f614b4e855f5260205f2090565b920160051c82019101613c75565b5f81815260208120918190559055565b614b2f565b9190614b7b613479565b50825f526007602052614b918160405f20615378565b15614c1457614c0f61059191845f52600860205260405f20815f52602052610cfd614bbe60405f20613043565b95614beb614bd4825f52600660205260405f2090565b6040890151611188906001600160a01b0316612361565b614bff600260808a01519201918254614494565b90555f52600860205260405f2090565b614adf565b637f11bea960e01b5f5260045260245ffd5b91610e079391614c3593615482565b90929192615504565b81516001600160a01b03909116915f9182916020018285620186a0f1614c62612701565b9015614c87578051614c7e57503b15614c7a57600190565b5f90565b60209150015190565b50505f90565b814710614ce75782516001600160a01b03909116925f9182916020018486620186a0f190614cb9612701565b9115614ce05715614ccc575b5050600190565b8051614c7e57503b15614c7a575f80614cc5565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f908152600660209081526040808320600180855292529091209081015460a01c60ff1615614d31576003018054918203918211610ffa5755565b6004018054918201809211610ffa5755565b5f9081526006602090815260408083206001600160a01b03909416835292905220600181015460a01c60ff1615614d31576003018054918203918211610ffa5755565b5f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16614c87575f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b03861684529091529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b610e07916001600160a01b0316906155e9565b5f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615614c87575f8181525f516020615afb5f395f51905f52602090815260408083206001600160a01b03861684529091529020805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60ff5f516020615b5b5f395f51905f525460401c1615614ef757565b631afcd79f60e31b5f5260045ffd5b614f1290614102614edb565b50565b90614f1e614edb565b614f2c60ff831615156134a9565b604091825192614f3c8185610560565b60098452682b30b634b230ba37b960b91b6020850152614f5e81519182610560565b60058152640312e302e360dc1b6020820152614f78614edb565b614f80614edb565b83516001600160401b03811161050a57614fb081614fab5f5160206159db5f395f51905f525461300b565b613c8b565b6020601f821160011461504c5781614ffa9392614fe69261059197985f92613e085750508160011b915f199060031b1c19161790565b5f5160206159db5f395f51905f5255615811565b61500f5f5f516020615a5b5f395f51905f5255565b6150245f5f516020615b7b5f395f51905f5255565b60ff1660ff195f5160206159fb5f395f51905f525416175f5160206159fb5f395f51905f5255565b5f5160206159db5f395f51905f525f52601f198216955f516020615a7b5f395f51905f52965f5b8181106150c35750966001928492614ffa9695610591999a106150ab575b505050811b015f5160206159db5f395f51905f5255615811565b01515f1960f88460031b161c191690555f8080615091565b83830151895560019098019760209384019301615073565b6150e3614edb565b62015180600155565b156150f657505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b1561512757565b63559d141b60e11b5f5260045ffd5b90936001600160a01b03851692600184036151a65750610591945061516c61515e8286612fd6565b34143490612e358488612fd6565b80615178575b506144a1565b6033546151a09161519b916001600160a01b031690615195611d75565b91614c8d565b615120565b5f615172565b906151b2343415612ad6565b6151c76151bf8287612fd6565b308489615636565b80615263575b506151ec6109896001613b2a866138f6875f52600660205260405f2090565b6151fc575b5061059193506144a1565b604051632770a7eb60e21b8152306004820152602481018590526020816044815f885af1918215612f7a576105919661523e9387935f91615244575b506150ec565b5f6151f1565b61525d915060203d60201161287e576128718183610560565b5f615238565b603354615283919061527d906001600160a01b0316612361565b8761567f565b5f6151cd565b90813b15615307575f516020615a9b5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156152ef57614f12916156bb565b5050346152f857565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b80548210156125a9575f5260205f2001905f90565b80548015615364575f1901906153538282615328565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f14615413575f198401848111610ffa5783545f19810194908511610ffa575f9585836153c697610cfd95036153cc575b50505061533d565b55600190565b6153fc6153f6916153ed6153e361540a9588615328565b90549060031b1c90565b92839187615328565b906134e4565b85905f5260205260405f2090565b555f80806153be565b505050505f90565b6154236156d8565b61542b6157df565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815261547c60c082610560565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116154ef579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612f7a575f516001600160a01b038116156154e557905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611e8557565b61550d816154fa565b80615516575050565b61551f816154fa565b600181036155365763f645eedf60e01b5f5260045ffd5b61553f816154fa565b6002810361555a575063fce698f760e01b5f5260045260245ffd5b806155666003926154fa565b1461556e5750565b6335e2f38360e21b5f5260045260245ffd5b5f818152600360205260409020546155e457600254600160401b81101561050a576155cd6155b78260018594016002556002615328565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b5f828152600182016020526040902054614c8757805490600160401b82101561050a57826156216155b7846001809601855584615328565b90558054925f520160205260405f2055600190565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105919161567a608483610560565b615924565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526105919161567a606483610560565b5f80610e0793602081519101845af46156d2612701565b9161597c565b6040515f5160206159db5f395f51905f5254905f816156f68461300b565b9182825260208201946001811690815f146157c3575060011461576b575b61572092500382610560565b5190811561572c572090565b50505f516020615a5b5f395f51905f525480156157465790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b505f5160206159db5f395f51905f525f90815290915f516020615a7b5f395f51905f525b8183106157a757505090602061572092820101615714565b602091935080600191548385880101520191019091839261578f565b60ff191686525061572092151560051b82016020019050615714565b6157e7614a32565b80519081156157f7576020012090565b50505f516020615b7b5f395f51905f525480156157465790565b9081516001600160401b03811161050a576158508161583d5f516020615a3b5f395f51905f525461300b565b5f516020615a3b5f395f51905f52613cda565b602092601f82116001146158905761587f929382915f92613e085750508160011b915f199060031b1c19161790565b5f516020615a3b5f395f51905f5255565b5f516020615a3b5f395f51905f525f52601f198216937f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75915f5b86811061590c57508360019596106158f4575b505050811b015f516020615a3b5f395f51905f5255565b01515f1960f88460031b161c191690555f80806158dd565b919260206001819286850151815501940192016158ca565b905f602091828151910182855af1156126f6575f513d61597357506001600160a01b0381163b155b6159535750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b6001141561594c565b906159a0575080511561599157805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806159d1575b6159b1575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156159a956fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6ba16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d101a2646970667358221220dd2946cd6b9e9bf63ead699bdaa7434d723e0b864cfb6417972e7a22cdf00c6264736f6c634300081c0033"

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
