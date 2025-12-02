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
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allPendingIndex\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"allTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeExecuter\",\"outputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeVerifier\",\"outputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deadWallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"alreadyTransferred\",\"type\":\"bool\"}],\"name\":\"burnCrossToDeadWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossMintableERC20Code\",\"outputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8[][]\",\"name\":\"v\",\"type\":\"uint8[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"r\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"s\",\"type\":\"bytes32[][]\"}],\"name\":\"finalizeBridgeBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextFinalizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getNextInitiateIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPendingArguments\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeRegistry.FinalizeArguments\",\"name\":\"args\",\"type\":\"tuple\"},{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.PendingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPair\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRegistry.TokenPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"manualReleasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"manualReleasePendingWithRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments\",\"name\":\"permitArgs\",\"type\":\"tuple\"}],\"name\":\"permitBridgeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseBridge.BridgeTokenArguments[]\",\"name\":\"args\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"contractIERC20Permit\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIBaseBridge.PermitArguments[]\",\"name\":\"permitArgs\",\"type\":\"tuple[]\"}],\"name\":\"permitBridgeTokenBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"cross_\",\"type\":\"address\"}],\"name\":\"reinitializeBSCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"releasePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"remoteChainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"releasePendingBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removePending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeExecuter\",\"name\":\"_bridgeExecuter\",\"type\":\"address\"}],\"name\":\"setBridgeExecuter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeVerifier\",\"name\":\"_bridgeVerifier\",\"type\":\"address\"}],\"name\":\"setBridgeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setChainPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20Code\",\"name\":\"_crossMintableERC20Code\",\"type\":\"address\"}],\"name\":\"setCrossMintableERC20Code\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"dev_\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setTokenPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setVerificationDelayExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unregisterToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BridgeExecuterCallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeExecuter\",\"type\":\"address\"}],\"name\":\"BridgeExecuterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BridgePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeVerifier\",\"type\":\"address\"}],\"name\":\"BridgeVerifierSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"ChainPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20CodeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dev\",\"type\":\"address\"}],\"name\":\"DevSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ManualReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"PendingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"TokenPairRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"TokenPairUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"TokenPauseSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelayExpirationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"VerificationDelaySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BSCBridgeInvalidDeadWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeBurnFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeFailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"BaseBridgeFailedPermitBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumConst.FinalizeStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BaseBridgeFailedRelease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeInvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"BaseBridgeInvalidPermitToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeMismatchPermitAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseBridgeNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delayExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseBridgeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeNotMatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BaseBridgeVerifierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryChainPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"}],\"name\":\"RegistryExistERC20Code\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryExistToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryFactoryNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RegistryNotExistIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RegistryTokenPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorThresholdCanNotZero\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cf56118e": "allChainIDs()",
		"79214874": "allPendingIndex(uint256)",
		"5b605f5c": "allTokenPairs(uint256)",
		"bc8cadd6": "bridgeExecuter()",
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
	Bin: "0x60a080604052346100c257306080525f51602061604b5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051615f8490816100c78239608051818181610d270152610ee00152f35b6001600160401b0319166001600160401b039081175f51602061604b5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610022575b3615610018575f80fd5b61002061381c565b005b5f3560e01c806301ffc9a7146103915780630b1d8d061461038c5780631313996b146103875780631938e0f214610382578063248a9ca31461037d5780632f2ff15d1461037857806336568abe14610373578063365d71e41461036e57806342cde4e81461036957806348a00ed8146103645780634be13f831461035f5780634d5d00561461035a5780634ee078ba146103555780634f1ef28614610350578063502cc5c01461034b57806352d1902d146103465780635b605f5c146103415780635c975abb1461033c5780635fd262de14610337578063670e626814610332578063792148741461032d5780637b429ad814610328578063814914b514610323578063815f1de61461031e57806384b0196e1461031957806386e9e1751461031457806388d67d6d1461030f57806389232a001461030a5780638ae87c5c1461030557806391cca3db1461030057806391cf6d3e146102fb57806391d14854146102f657806394ddc8c6146102f15780639f9b4f1c146102ec578063a217fddf146102e7578063a3246ad3146102e2578063aa1bd0bc146102dd578063aa20ed47146102d8578063ad3cb1cc146102d3578063ae6893f8146102ce578063b2b49e2e146102c9578063b33eb36e146102c4578063b7f3358d146102bf578063b8aa837e146102ba578063bc8cadd6146102b5578063bedb86fb146102b0578063cf56118e146102ab578063d477f05f146102a6578063d547741f146102a1578063d5717fc51461029c578063e2af6cd714610297578063edbbea3914610292578063f0702e8e1461028d578063f4509643146102885763f698da250361000e576126e2565b612641565b612619565b6124c1565b61240e565b6123d5565b6123a6565b612347565b6122d3565b6121e9565b6121c1565b612119565b612095565b612004565b611ef6565b611ebd565b611e76565b611deb565b611d9e565b611d22565b611cc6565b611baa565b611b01565b611ab3565b611a96565b611a6e565b611a05565b611900565b6117f6565b61163f565b61158b565b61149d565b6113d3565b6112fd565b61128d565b6111d2565b6110ca565b61109c565b610fbd565b610ece565b610e3e565b610ce5565b610b7c565b610b0f565b610abb565b61099c565b610970565b6108fd565b61080d565b6107d2565b6107a1565b6106f0565b6104aa565b610408565b346103e75760203660031901126103e75760043563ffffffff60e01b81168091036103e757602090637965db0b60e01b81149081156103d6575b506040519015158152f35b6301ffc9a760e01b1490505f6103cb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103e757565b346103e75760203660031901126103e757600435610425816103f7565b61042e3361496d565b6001600160a01b031661044281151561274e565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103e7578235916001600160401b0383116103e7576020808501948460051b0101116103e757565b60403660031901126103e7576004356001600160401b0381116103e7576104d590369060040161047a565b602435916001600160401b0383116103e757366023840112156103e7576004830135916001600160401b0383116103e75736602460e08502860101116103e757602461002094019161296d565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761055257604052565b610522565b60e081019081106001600160401b0382111761055257604052565b606081019081106001600160401b0382111761055257604052565b60c081019081106001600160401b0382111761055257604052565b601f909101601f19168101906001600160401b0382119082101761055257604052565b604051906105da6080836105a8565b565b604051906105da60e0836105a8565b604051906105da610100836105a8565b604051906105da6060836105a8565b6001600160401b0381116105525760051b60200190565b60ff8116036103e757565b9080601f830112156103e75781356106438161060a565b9261065160405194856105a8565b81845260208085019260051b8201019283116103e757602001905b8282106106795750505090565b60208091833561068881610621565b81520191019061066c565b9080601f830112156103e75781356106aa8161060a565b926106b860405194856105a8565b81845260208085019260051b8201019283116103e757602001905b8282106106e05750505090565b81358152602091820191016106d3565b60803660031901126103e7576004356001600160401b0381116103e75760c060031982360301126103e7576024356001600160401b0381116103e75761073a90369060040161062c565b906044356001600160401b0381116103e75761075a903690600401610693565b60643591906001600160401b0383116103e75761079d9361078261078b943690600401610693565b92600401612aab565b60405191829182901515815260200190565b0390f35b346103e75760203660031901126103e75760206107bf600435612def565b604051908152f35b35906105da826103f7565b346103e75760403660031901126103e7576100206024356004356107f5826103f7565b61080861080182612def565b3390614b8d565b614370565b346103e75760403660031901126103e75760043560243561082d816103f7565b336001600160a01b0382160361084657610020916143d0565b63334bd91960e11b5f5260045ffd5b9060406003198301126103e7576004356001600160401b0381116103e7578261088091600401610693565b91602435906001600160401b0382116103e757806023830112156103e75781600401356108ac8161060a565b926108ba60405194856105a8565b8184526024602085019260051b8201019283116103e757602401905b8282106108e35750505090565b6020809183356108f2816103f7565b8152019101906108d6565b346103e75761090b36610855565b906109198151835114612e0d565b5f5b8251811015610020578061095f61093460019385612e23565b51838060a01b036109458488612e23565b51169061095a3361095583612def565b614b8d565b6143d0565b500161091b565b5f9103126103e757565b346103e7575f3660031901126103e757602060ff5f516020615c8f5f395f51905f525416604051908152f35b346103e75760603660031901126103e7576024356004356044356109bf816103f7565b6109c8336149f5565b6109d0613892565b815f5260076020526109ee836109e98160405f20615111565b612e37565b806109f98484614f8e565b916001600160a01b031615610aa7575b8151915f516020615caf5f395f51905f526040820192610a56610a44855160018060a01b03169660808601978489519160a089015193613c48565b610a4d81611f69565b6001811461442f565b8251602090930151935194516040516001600160a01b0390961695918291610a819142919084612dd1565b0390a45f516020615c4f5f395f51905f525f80a35f5f516020615e2f5f395f51905f525d005b5060608101516001600160a01b0316610a09565b346103e7575f3660031901126103e7575f546040516001600160a01b039091168152602090f35b9181601f840112156103e7578235916001600160401b0383116103e757602083818601950101116103e757565b6101c03660031901126103e757602435600435610b2b826103f7565b604435610b37816103f7565b6064359060a43560843560c4356001600160401b0381116103e757610b60903690600401610ae2565b94909360e03660e31901126103e75761079d9761078b97612e51565b346103e75760403660031901126103e757610c64602435600435610b9e61386b565b610ba6613892565b805f526007602052610bbf826109e98160405f20615111565b610c5f6040610bef610bea85610bdd865f52600860205260405f2090565b905f5260205260405f2090565b613205565b610c4c610c0f82516080610c05868301516103eb565b9101519087613acd565b50610c1981611f69565b610c2281611f69565b83516020810182905290600190610c4683604081015b03601f1981018552846105a8565b1461323e565b01518015908115610c6c575b429161326a565b614452565b6100206138c7565b4281109150610c58565b6001600160401b03811161055257601f01601f191660200190565b929192610c9d82610c76565b91610cab60405193846105a8565b8294818452818301116103e7578281602093845f960137010152565b9080601f830112156103e757816020610ce293359101610c91565b90565b60403660031901126103e757600435610cfd816103f7565b6024356001600160401b0381116103e757610d1c903690600401610cc7565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610e1c575b50610e0d57610d603361496d565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610ddc575b50610da957634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d4f5f395f51905f528303610dc85761002092506154fa565b632a87526960e21b5f52600483905260245ffd5b610dff91945060203d602011610e06575b610df781836105a8565b810190613c39565b925f610d88565b503d610ded565b63703e46dd60e11b5f5260045ffd5b5f516020615d4f5f395f51905f52546001600160a01b0316141590505f610d52565b346103e75760603660031901126103e757602435600435604435610e61336149f5565b815f526007602052610e7a836109e98160405f20615111565b4201804211610ec95760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b613071565b346103e7575f3660031901126103e7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e0d5760206040515f516020615d4f5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610f9d5750505090565b909192602060e082610fb26001948851610f25565b019401929101610f90565b346103e75760203660031901126103e757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061108357505061100c925003836105a8565b61101682516132be565b915f5b81518110156110755760019061105961105461103486612cd1565b61104e6110418588612e23565b516001600160a01b031690565b9061330d565b613322565b6110638287612e23565b5261106e8186612e23565b5001611019565b6040518061079d8682610f7a565b8454835260019485019487945060209093019201610ff7565b346103e7575f3660031901126103e757602060ff5f516020615def5f395f51905f5254166040519015158152f35b60e03660031901126103e7576024356004356110e5826103f7565b604435916110f2836103f7565b6064359260c435916084359160a435916001600160401b0385116103e7576111a396611158611128611199973690600401610ae2565b95909661113361386b565b6001600160a01b03851694849061114a878d6144f7565b611152613892565b8b614597565b939092604051986111688a610536565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610c91565b60e0820152614742565b5f5f516020615e2f5f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103e75760803660031901126103e7576024356004356111f2826103f7565b604435906001600160401b0382116103e757366023830112156103e75761079d9261122a61123d933690602481600401359101610c91565b906064359261123884610621565b6133aa565b604051918291826111bf565b90602080835192838152019201905f5b8181106112665750505090565b8251845260209384019390920191600101611259565b906020610ce2928181520190611249565b346103e75760203660031901126103e7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106112e75761079d856112db818703826105a8565b6040519182918261127c565b82548452602090930192600192830192016112c4565b346103e75760403660031901126103e75760043560243561131d816103f7565b6113263361496d565b5f516020615eaf5f395f51905f52549160ff8360401c1680156113af575b6113a0576001600160401b03199092166002175f516020615eaf5f395f51905f525561137791611372613473565b6134c5565b61137f613494565b604051600281525f516020615ccf5f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b0384161015611344565b60e0810192916105da9190610f25565b346103e75760403660031901126103e75761079d6114126024356004356113f9826103f7565b611401613288565b505f52600660205260405f2061330d565b60046040519161142183610557565b80546001600160a01b0316835260018101546114769061146d90611450611447826103eb565b602088016130ba565b61146460a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c0820152604051918291826113c3565b346103e75760203660031901126103e7576004356114ba816103f7565b6114c33361496d565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161156190611553610ce297959693600f60f81b865260e0602087015260e0860190611508565b908482036040860152611508565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611249565b346103e7575f3660031901126103e7575f516020615d0f5f395f51905f5254158061161f575b156115e2576115be614819565b6115c66148d3565b9061079d6115d261359b565b604051938493309146918661152c565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615ecf5f395f51905f5254156115b1565b801515036103e757565b346103e75760603660031901126103e75760043561165c816103f7565b60243560443561166b81611635565b61167361386b565b61167b613892565b6001600160a01b0383165f9081526066602052604090206116a7906116a2905b5460ff1690565b6135b6565b6116b28215156134af565b611769576116cd81836116c66065546103eb565b3390614bd9565b335b5f516020615f2f5f395f51905f526116e8606454614c22565b92606454926117526116fb6065546103eb565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a461175d6138c7565b60405160018152602090f35b6117723361496d565b306116cf565b9080601f830112156103e757813561178f8161060a565b9261179d60405194856105a8565b81845260208085019260051b820101918383116103e75760208201905b8382106117c957505050505090565b81356001600160401b0381116103e7576020916117eb87848094880101610693565b8152019101906117ba565b60803660031901126103e7576004356001600160401b0381116103e75761182190369060040161047a565b90602435906001600160401b0382116103e757366023830112156103e757816004013561184d8161060a565b9261185b60405194856105a8565b8184526024602085019260051b820101903682116103e75760248101925b8284106118d15750506044359150506001600160401b0381116103e7576118a4903690600401611778565b606435926001600160401b0384116103e75761079d946118cb61078b953690600401611778565b936135cc565b83356001600160401b0381116103e7576020916118f583926024369187010161062c565b815201930192611879565b346103e75760603660031901126103e75760043561191d816103f7565b6024359061192a826103f7565b60443561193681610621565b5f516020615eaf5f395f51905f52549260ff604085901c1615611958565b1590565b936001600160401b0316801590816119fd575b60011490816119f3575b1590816119ea575b506113a0575f516020615eaf5f395f51905f5280546001600160401b03191660011790556119af92846119dd57614c3b565b6119b557005b6119bd613494565b604051600181525f516020615ccf5f395f51905f5290806020810161139b565b6119e5613473565b614c3b565b9050155f61197d565b303b159150611975565b85915061196b565b346103e75760403660031901126103e757602435600435611a253361496d565b611a2d613892565b611a378282614f8e565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615e2f5f395f51905f525d005b346103e7575f3660031901126103e7576034546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e7576020603554604051908152f35b346103e75760403660031901126103e757602060ff611af5602435600435611ada826103f7565b5f525f516020615dcf5f395f51905f52845260405f2061330d565b54166040519015158152f35b346103e75760603660031901126103e757602435600435611b21826103f7565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea6020604435611b5081611635565b611b5933614a7d565b835f5260058252611b9f816001611b9960405f2098611b8981611b84858060a01b038216809d615111565b613662565b885f526006875260405f2061330d565b0161368a565b6040519015158152a3005b346103e75760403660031901126103e7576004356001600160401b0381116103e757611bda903690600401610693565b6024356001600160401b0381116103e757611bf9903690600401610693565b90611c078151835114612783565b5f5b82518110156100205780611cb8611c2260019385612e23565b51611c2d8387612e23565b5190611c3761386b565b611c3f613892565b805f526007602052611c58826109e98160405f20615111565b610c5f6040611c76610bea85610bdd865f52600860205260405f2090565b610c4c611c8c82516080610c05868301516103eb565b50611c9681611f69565b611c9f81611f69565b835160208101829052908a90610c468360408101610c38565b611cc06138c7565b01611c09565b346103e7575f3660031901126103e75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611d035750505090565b82516001600160a01b0316845260209384019390920191600101611cf6565b346103e75760203660031901126103e7576004355f525f516020615d6f5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611d885761079d85611d7c818703826105a8565b60405191829182611ce0565b8254845260209093019260019283019201611d65565b346103e75760203660031901126103e7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611dde3361496d565b80600155604051908152a1005b346103e75760403660031901126103e757602435600435611e0b336149f5565b611e14336149f5565b611e1c613892565b805f526007602052611e35826109e98160405f20615111565b611e3f8282614452565b5f516020615c4f5f395f51905f525f80a35f5f516020615e2f5f395f51905f525d005b60405190611e716020836105a8565b5f8252565b346103e7575f3660031901126103e75761079d604051611e976040826105a8565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611508565b346103e75760203660031901126103e7576004355f526004602052600160405f20015460018101809111610ec957602090604051908152f35b346103e757611f0436610855565b90611f128151835114612e0d565b5f5b82518110156100205780611f4e611f2d60019385612e23565b51838060a01b03611f3e8488612e23565b5116906108083361095583612def565b5001611f14565b634e487b7160e01b5f52602160045260245ffd5b600a1115611f7357565b611f55565b90600a821015611f735752565b6020815260606040611fea60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611508565b93611ffc602082015183860190611f78565b015191015290565b346103e75760403660031901126103e757600435602435905f6040805161202a81610572565b6120326136a7565b815282602082015201525f52600860205260405f20905f5260205261079d60405f2060076040519161206383610572565b61206c81613101565b835261208260ff600683015416602085016131f9565b0154604082015260405191829182611f85565b346103e75760203660031901126103e7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff6004356120d781610621565b6120e03361496d565b166120ec8115156136d7565b8060ff195f516020615c8f5f395f51905f525416175f516020615c8f5f395f51905f5255604051908152a1005b346103e75760403660031901126103e75760043560243561213981611635565b61214233614a7d565b612157825f52600360205260405f2054151590565b156121ae5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526121a381600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103e7575f3660031901126103e7576033546040516001600160a01b039091168152602090f35b346103e75760203660031901126103e75760043561220681611635565b61220f33614a7d565b1561226d5761221c61386b565b600160ff195f516020615def5f395f51905f525416175f516020615def5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615def5f395f51905f525460ff8116156122c45760ff19165f516020615def5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103e7575f3660031901126103e757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106123315761079d856112db818703826105a8565b825484526020909301926001928301920161231a565b346103e75760203660031901126103e757600435612364816103f7565b61236d3361496d565b6001600160a01b031661238181151561274e565b603480546001600160a01b031916821790555f516020615d8f5f395f51905f525f80a2005b346103e75760403660031901126103e7576100206024356004356123c9826103f7565b61095a61080182612def565b346103e75760203660031901126103e7576004355f526004602052600260405f20015460018101809111610ec957602090604051908152f35b346103e75760203660031901126103e75760043561242b816103f7565b6124343361496d565b5f546001600160a01b0316806124af57506124596001600160a01b03821615156136ed565b5f80546001600160a01b0319166001600160a01b03928316179081905561248890612483906103eb565b6103eb565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103e75760803660031901126103e7576004356024356124e181611635565b604435916124ee836103f7565b5f516020615e6f5f395f51905f526125d86064359361250c856103f7565b61251533614b05565b6125208415156136ed565b6001600160a01b0386169461078b9061253a8715156136ed565b6001600160a01b038116976125d3906125548a15156136ed565b61255d88615927565b6125dd575b612586816125818161257c8c5f52600560205260405f2090565b6152d4565b61506a565b6125a56125916105dc565b9361259c83866130ba565b602085016130ba565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526125ce88612cd1565b61330d565b615092565b0390a4005b6126146125e86105cb565b8981525f60208201525f60408201525f606082015261260f8a5f52600460205260405f2090565b615033565b612562565b346103e7575f3660031901126103e7576032546040516001600160a01b039091168152602090f35b346103e75760403660031901126103e757602435600435612661826103f7565b61266a3361496d565b805f5260056020525f60046126a5604083209461269481611b8460018060a01b0382168099615680565b84845260066020526040842061330d565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103e7575f3660031901126103e7576126fa615a72565b612702615ac9565b6040519060208201925f516020615eef5f395f51905f528452604083015260608201524660808201523060a082015260a0815261274060c0826105a8565b519020604051908152602090f35b1561275557565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b1561278a57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156127cf5760051b8101359060fe19813603018212156103e7570190565b612799565b35610ce2816103f7565b903590601e19813603018212156103e757018035906001600160401b0382116103e7576020019181360383136103e757565b91908110156127cf5760e0020190565b908160209103126103e75751610ce281611635565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936128a697939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c0860191612835565b9480356128b2816103f7565b6001600160a01b031660e085015260208101356128ce816103f7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561290381610621565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d15612951573d9061293882610c76565b9161294660405193846105a8565b82523d5f602084013e565b606090565b604090610ce2939281528160208201520190611508565b909193929361297d858414612783565b5f945b83861061298f57505050509050565b61299a8685856127ad565b3560206129b2816129ac8a89896127ad565b016127d4565b6129c260606129ac8b8a8a6127ad565b92612a388a86888a8c612a1c60806129db8784866127ad565b013595612a14612a0a8260a06129f282888a6127ad565b01359560c0612a0283838b6127ad565b0135976127ad565b60e08101906127de565b969095612810565b946040519a8b998a996326ae802b60e11b8b5260048b01612855565b03815f305af19081612a7f575b50612a745785612a53612927565b60405163f495148b60e01b8152918291612a709160048401612956565b0390fd5b600190950194612980565b612a9f9060203d8111612aa4575b612a9781836105a8565b810190612820565b612a45565b503d612a8d565b90612b90939291612aba61386b565b612ac2613892565b803592612ad7845f52600560205260405f2090565b91612b07612af56040830194612aef612483876127d4565b906138d9565b612b01612483866127d4565b90612cdf565b612b12343415612d07565b612ba6612b32865f526004602052600260405f2001600181540180915590565b96612b45602084013589819a8214612d25565b612b51612483866127d4565b93606084019688612b61896127d4565b96612b9e8c60808901359e8f60a08b019b612b7c8d8d6127de565b929091604051978896602088019a8b612d43565b03601f1981018352826105a8565b51902061391c565b612bb987612bb3856127d4565b87613b1b565b919092600184612bc881611f69565b14612c95575b50612bd883611f69565b60018303612c325750505090612c04612bfe5f516020615caf5f395f51905f52936127d4565b916127d4565b6040516001600160a01b03909216958291612c2191429184612dd1565b0390a45b612c2d6138c7565b600190565b612c6f8394612c6a612c8d947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612c7595614242565b6127d4565b926127d4565b9260405193849360018060a01b031698429185612da7565b0390a4612c25565b612cca91935088612ca5866127d4565b91612cc3612cbc612cb58a6127d4565b92886127de565b3691610c91565b9289613c48565b915f612bce565b5f52600660205260405f2090565b15612ce75750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612d0f5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612d2e575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610ce297959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c08201520191612835565b6105da93606092969593608083019760018060a01b03168352602083015260408201520190611f78565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615dcf5f395f51905f52602052600160405f20015490565b15612e1457565b63031206d560e51b5f5260045ffd5b80518210156127cf5760209160051b010190565b15612e3f5750565b6373a1390160e11b5f5260045260245ffd5b959394612ed3919597939297612e6561386b565b612eaa6001600160a01b038816612e7c818b6144f7565b612e84613892565b612e9461248361248360e46127d4565b811490612ea461248360e46127d4565b9161302e565b612ecb612ebb6124836101046127d4565b6001600160a01b038b161461305b565b838789614597565b9161012435612f0681612eef86612eea8787613085565b613085565b811015612f0087612eea8888613085565b90613092565b612f146124836032546103eb565b90612f206101046127d4565b906101443592612f316101646130b0565b92610184356101a43591833b156103e757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af198891561302957612ff061119998612ff993612c259c61300f575b50612fe7612fd26101046127d4565b91612fdb6105eb565b9c8d5260208d016130ba565b60408b016130ba565b606089016130ba565b608087015260a086015260c08501523691610c91565b8061301d5f613023936105a8565b80610966565b5f612fc3565b61291c565b15613037575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b1561306257565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ec957565b1561309b575050565b63943892eb60e01b5f5260045260245260445ffd5b35610ce281610621565b6001600160a01b039091169052565b90600182811c921680156130f7575b60208310146130e357565b634e487b7160e01b5f52602260045260245ffd5b91607f16916130d8565b9060405161310e8161058d565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91613166826130c9565b80855291600181169081156131d25750600114613193575b505060a0929161318f9103846105a8565b0152565b5f908152602081209092505b8183106131b657505081016020018161318f61317e565b602091935080600191548385890101520191019091849261319f565b60ff191660208681019190915292151560051b8501909201925083915061318f905061317e565b600a821015611f735752565b9060405161321281610572565b60406007829461322181613101565b845261323760ff600683015416602086016131f9565b0154910152565b156132465750565b60405162461bcd60e51b815260206004820152908190612a70906024830190611508565b15613273575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061329582610557565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906132c88261060a565b6132d560405191826105a8565b82815280926132e6601f199161060a565b01905f5b8281106132f657505050565b602090613301613288565b828285010152016132ea565b9060018060a01b03165f5260205260405f2090565b9060405161332f81610557565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c09160049161337a9061337190611464565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103e75751610ce2816103f7565b5f5490949392906001600160a01b0381161561346457604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90613413906084870190611508565b931691015203925af18015613029576105da925f91613435575b508094613703565b613457915060203d60201161345d575b61344f81836105a8565b810190613395565b5f61342d565b503d613445565b6315aeca0d60e11b5f5260045ffd5b5f516020615eaf5f395f51905f52805460ff60401b1916600160401b179055565b5f516020615eaf5f395f51905f52805460ff60401b19169055565b156134b657565b63ff5a231360e01b5f5260045ffd5b906134d18215156134af565b6001600160a01b031690811561358c57606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b604051906135aa6020836105a8565b5f808352366020840137565b156135bd57565b635e0c1f8960e01b5f5260045ffd5b9192949390938484511480613658575b8061364e575b6135eb90612783565b5f5b85811015613642578060051b8401359060be19853603018212156103e75761363b60019261361b8389612e23565b51613626848c612e23565b51906136328589612e23565b51928901612aab565b50016135ed565b50945050505050600190565b50815185146135e2565b50848651146135dc565b1561366a5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b604051906136b48261058d565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156136de57565b63f0f15b9160e01b5f5260045ffd5b156136f457565b6302bfb33d60e51b5f5260045ffd5b91909161370f33614b05565b61371a8115156136ed565b6001600160a01b038316916137a8906137348415156136ed565b6001600160a01b038116946125d39061374e8715156136ed565b61375785615927565b6137c8575b613776816125818161257c895f52600560205260405f2090565b6137816125916105dc565b5f60408401525f60608401525f60808401525f60a08401525f60c08401526125ce85612cd1565b6040515f81525f516020615e6f5f395f51905f529080602081015b0390a4565b6137fa6137d36105cb565b8681525f60208201525f60408201525f606082015261260f875f52600460205260405f2090565b61375c565b916138189183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361383057565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b60ff5f516020615def5f395f51905f52541661388357565b63d93c066560e01b5f5260045ffd5b5f516020615e2f5f395f51905f525c6138b85760015f516020615e2f5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615e2f5f395f51905f525d565b610ce2916001600160a01b031690615111565b156138f357565b63b3f07a3960e01b5f5260045ffd5b1561390a5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a7a575b613935906138ec565b6139566139505f516020615c8f5f395f51905f525460ff1690565b60ff1690565b956139648488811015613902565b5f945f935f5b86811061398557505050505050506105da9192811015613902565b6139e26139b183613994615723565b6042916040519161190160f01b8352600283015260228201522090565b6139c56139be8489612e23565b5160ff1690565b6139cf8487612e23565b51906139db8589612e23565b5192615124565b6001600160a01b038181169088161080613a13575b613a05575b5060010161396a565b6001988901989096506139fc565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615dcf5f395f51905f52602052613a7561169b827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561330d565b6139f7565b508551831461392c565b15613a8b57565b6330d45fb160e01b5f5260045ffd5b908160209103126103e75751600a8110156103e75790565b6001600160a01b039091168152602081019190915260400190565b9150613b08611054606092613af75f95613af260325460018060a01b03161515613a84565b612cd1565b6001600160a01b039091169061330d565b0151613b1357600191565b506002905f90565b9190915f92613b686060613b61611054613b51613b3c6124836032546103eb565b96613af26001600160a01b0389161515613a84565b6001600160a01b0386169061330d565b0151151590565b613bef5791602091613b9295935f604051809881958294633f4bdec560e01b845260048401613ab2565b03925af1928315613029575f93613bbe575b50600183613bb181611f69565b03613bb857565b60019150565b613be191935060203d602011613be8575b613bd981836105a8565b810190613a9a565b915f613ba4565b503d613bcf565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610ce292910190611508565b908160209103126103e7575190565b92919382613c5585612cd1565b95613c786001613c6d818060a01b038716809a61330d565b015460a01c60ff1690565b92601481511180613f7a575b613cbd575b50613c939361513c565b92613c9d84611f69565b60018414613cac575b50505090565b613cb5926151fb565b5f8080613ca6565b9150613cf593602083015160601c6020613cdb6124836033546103eb565b916040518098819262483e5560e91b8352600483016111bf565b0381845afa938415613029578888958b985f91613f5b575b50613d1b5750505093613c89565b92809585988896936002613d396001985f52600460205260405f2090565b01549687921480159889613f20575b5060209596505f14613f15578c613d8784925b613d696124836033546103eb565b95604051998a9889978896631eeed51360e01b885260048801613bfa565b03925af1908115613029575f91613ef6575b50613ee6579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613df4898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613ede575b50613e10575b92613c939361513c565b9050613e206124836033546103eb565b92604051916370a0823160e01b835260208380613e4088600483016111bf565b03818b5afa80156130295786935f91613ebf575b5080613e64575b50909350613e06565b6020919350613e88956040519687928392632770a7eb60e21b845260048401613ab2565b03815f8b5af191821561302957613c9394869315613e5b57613eb89060203d602011612aa457612a9781836105a8565b505f613e5b565b613ed8915060203d602011610e0657610df781836105a8565b5f613e54565b90505f613e00565b50505050509091612c2d926151fb565b613f0f915060203d602011612aa457612a9781836105a8565b5f613d99565b8c613d875f92613d5b565b939250945094613f2f9361513c565b613f3881611f69565b60018103613f4d5750848884848a948a613d48565b995050505050505050505090565b613f74915060203d602011612aa457612a9781836105a8565b5f613d0d565b506033546001600160a01b0390613f9490612483906103eb565b161515613c84565b15613fa45750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103e75760405190613fcf8261058d565b819380358352602081013560208401526040810135613fed816103f7565b6040840152613ffe606082016107c7565b60608401526080818101359084015260a0810135916001600160401b0383116103e75760a09261402e9201610cc7565b910152565b81811061403e575050565b5f8155600101614033565b90601f8211614056575050565b6105da915f516020615c6f5f395f51905f525f5260205f20906020601f840160051c8301931061408e575b601f0160051c0190614033565b9091508190614081565b9190601f81116140a757505050565b6105da925f5260205f20906020601f840160051c8301931061408e57601f0160051c0190614033565b8160011b915f199060031b1c19161790565b90600a811015611f735760ff80198354169116179055565b815180518255602081015160018301556040810151919291614128906001600160a01b031660028501612764565b6060810151614143906001600160a01b031660038501612764565b6080810151600484015560a00151805160058401916001600160401b0382116105525761417a8261417485546130c9565b85614098565b602090601f83116001146141d2578260079593604095936141a3935f926141c7575b50506140d0565b90555b6141c060208201516141b781611f69565b600686016140e2565b0151910155565b015190505f8061419c565b90601f198316916141e6855f5260205f2090565b925f5b81811061422a5750926001928592600798966040989610614212575b505050811b0190556141a6565b01515f1960f88460031b161c191690555f8080614205565b929360206001819287860151815501950193016141e9565b9160806142e06142d1600293613af2876142cc61427c61381899833595865f52600760205261428160405f20602087013594858092615990565b613f9c565b156142ed576142b561429560015442613085565b915b6142aa6142a26105fb565b963690613fb6565b8652602086016131f9565b6040840152610bdd855f52600860205260405f2090565b6140fa565b61104e612483604088016127d4565b9301359201918254613085565b6142b55f91614297565b90614302825f615244565b918261430b5750565b5f80525f516020615d6f5f395f51905f526020526001600160a01b0316614352817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615990565b1561435a5750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161437d8382615244565b9283614387575050565b815f525f516020615d6f5f395f51905f526020526143b260405f209160018060a01b03168092615990565b156143bb575050565b63d180cb3160e01b5f5260045260245260445ffd5b9190916143dd83826152e7565b92836143e7575050565b815f525f516020615d6f5f395f51905f5260205261441260405f209160018060a01b03168092615680565b1561441b575050565b62a95f1b60e31b5f5260045260245260445ffd5b156144375750565b63290cd55f60e01b5f52600a811015611f735760045260245ffd5b9061445c91614f8e565b60018060a01b036060820151168151915f516020615caf5f395f51905f5260408201926144a4610a44855160018060a01b03169660808601978489519160a089015193613c48565b8251602090930151935194516040516001600160a01b03909616959182916137c39142919084612dd1565b156144d75750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f52600560205261451b81611b8460405f2060018060a01b03831690615111565b825f52600460205260ff600360405f200154166145535760ff6001614547836125ce6105da9697612cd1565b015460a81c16156144cf565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103e7578051916040602083015192015190565b1561458857565b6358e8878560e01b5f5260045ffd5b82606091614630979596936145c16110546145b184612cd1565b6001600160a01b0384169061330d565b6145d16119546040830151151590565b6146d2575b506145e56124836032546103eb565b916145fa6001600160a01b0384161515613a84565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115613029575f955f905f93614694575b509082916105da94939681988515159586614689575b50508461467e575b505082614673575b5050614581565b101590505f8061466c565b101592505f80614664565b101594505f8061465c565b90506146bf9196506105da93925060603d6060116146cb575b6146b781836105a8565b810190614566565b91969293919291614646565b503d6146ad565b60c06146e49101518480821015613092565b5f6145d6565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161473d9190860190611508565b930152565b61474c8151614c22565b6147568251612cd1565b5f516020615f2f5f395f51905f526147886124836001614781602088019561104e61248388516103eb565b01546103eb565b938051906137c361479985516103eb565b9160408101906147a982516103eb565b906147d26080820196875160a08401948551986147cc60c087019a8b5190613085565b936153b9565b6147e86147e1825199516103eb565b93516103eb565b9460e06147f860608401516103eb565b9751935191519201519260405197889760018060a01b03169c4296896146ea565b604051905f825f516020615c6f5f395f51905f525491614838836130c9565b80835292600181169081156148b4575060011461485c575b6105da925003836105a8565b505f516020615c6f5f395f51905f525f90815290915f516020615d2f5f395f51905f525b8183106148985750509060206105da92820101614850565b6020919350806001915483858901015201910190918492614880565b602092506105da94915060ff191682840152151560051b820101614850565b604051905f825f516020615cef5f395f51905f5254916148f2836130c9565b80835292600181169081156148b45750600114614915576105da925003836105a8565b505f516020615cef5f395f51905f525f90815290915f516020615f0f5f395f51905f525b8183106149515750509060206105da92820101614850565b6020919350806001915483858901015201910190918492614939565b5f516020615e4f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff6149ba827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61330d565b5416156149c45750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e4f5f395f51905f52602452604490fd5b5f516020615e8f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614a42827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361330d565b541615614a4c5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e8f5f395f51905f52602452604490fd5b5f516020615daf5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614aca827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661330d565b541615614ad45750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615daf5f395f51905f52602452604490fd5b5f516020615e0f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614b52827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61330d565b541615614b5c5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e0f5f395f51905f52602452604490fd5b90815f525f516020615dcf5f395f51905f5260205260ff614bb18260405f2061330d565b541615614bbc575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105da91614c1d6084836105a8565b61559c565b5f526004602052600160405f2001600181540180915590565b91614c446155f4565b614c586001600160a01b038416151561274e565b6001600160a01b03821692614c6e84151561274e565b60ff821615614ee057614cda90614c836155f4565b614c8b6155f4565b614c936155f4565b60ff195f516020615def5f395f51905f5254165f516020615def5f395f51905f5255614cbd6155f4565b614cc56155f4565b614ccd6155f4565b614cd56155f4565b6142f7565b50614ce36155f4565b614cf160ff821615156136d7565b60408051614cff82826105a8565b60098152682b30b634b230ba37b960b91b6020820152614d21825192836105a8565b60058252640312e302e360dc1b6020830152614d3b6155f4565b614d436155f4565b8051906001600160401b03821161055257614d7482614d6f5f516020615c6f5f395f51905f52546130c9565b614049565b602090601f8311600114614e4c5792614da283614e32979694614db694614e08975f926141c75750506140d0565b5f516020615c6f5f395f51905f5255615b59565b614dcb5f5f516020615d0f5f395f51905f5255565b614de05f5f516020615ecf5f395f51905f5255565b60ff1660ff195f516020615c8f5f395f51905f525416175f516020615c8f5f395f51905f5255565b614e1061561f565b603480546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615d8f5f395f51905f525f80a26105da43603555565b5f516020615c6f5f395f51905f525f52601f19831691905f516020615d2f5f395f51905f52925f5b818110614ec8575093614db693614e08969360019383614e329b9a9810614eb0575b505050811b015f516020615c6f5f395f51905f5255615b59565b01515f1960f88460031b161c191690555f8080614e96565b92936020600181928786015181550195019301614e74565b6338854f1160e21b5f5260045ffd5b91908203918211610ec957565b60075f9182815582600182015582600282015582600382015582600482015560058101614f2981546130c9565b9081614f3c575b50508260068201550155565b601f8211600114614f5357849055505b5f80614f30565b614f79614f89926001601f614f6b855f5260205f2090565b920160051c82019101614033565b5f81815260208120918190559055565b614f4c565b9190614f986136a7565b50825f526007602052614fae8160405f20615680565b156150215761501c6105da91845f52600860205260405f20815f52602052610bdd614fdb60405f20613101565b95614ff8614fe882612cd1565b61104e61248360408b01516103eb565b61500c600260808a01519201918254614eef565b90555f52600860205260405f2090565b614efc565b637f11bea960e01b5f5260045260245ffd5b600360606105da938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156150725750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916150ab60018060a01b0382511685612764565b60208101516150f69060018601906150cc906001600160a01b031682612764565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061368a565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b91610ce2939161513393615777565b909291926157f9565b6001600160a01b0316929190600184146151cb5781156151c25761518b92156151965760405163a9059cbb60e01b602082015291615183918391612b909160248401613ab2565b600592615875565b15610ce25750600190565b6040516340c10f1960e01b6020820152916151ba918391612b909160248401613ab2565b600692615875565b50505050600190565b906151ed935061195492506151de611e62565b916001600160a01b03166158be565b6151f657600190565b600590565b90615210915f52600660205260405f2061330d565b600181015460a01c60ff1615615232576003018054918203918211610ec95755565b6004018054918201809211610ec95755565b805f525f516020615dcf5f395f51905f5260205260ff6152678360405f2061330d565b54166152ce57805f525f516020615dcf5f395f51905f5260205261528e8260405f2061330d565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610ce2916001600160a01b031690615990565b805f525f516020615dcf5f395f51905f5260205260ff61530a8360405f2061330d565b5416156152ce57805f525f516020615dcf5f395f51905f526020526153328260405f2061330d565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b1561537957505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156153aa57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361542757506105da94506153ef6153e18286613085565b34143490612f008488613085565b806153fb575b50615a0c565b61541c6154219161540d6034546103eb565b90615416611e62565b916158be565b6153a3565b5f6153f5565b90615433343415612d07565b6154486154408287613085565b308489614bd9565b806154dc575b506154646119546001613c6d866125ce87612cd1565b615474575b506105da9350615a0c565b604051632770a7eb60e21b815260208180615493883060048401613ab2565b03815f885af1918215613029576105da966154b79387935f916154bd575b5061536f565b5f615469565b6154d6915060203d602011612aa457612a9781836105a8565b5f6154b1565b6154f4906154ee6124836034546103eb565b876159d7565b5f61544e565b90813b1561557b575f516020615d4f5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156155635761556091615a55565b50565b50503461556c57565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b905f602091828151910182855af11561291c575f513d6155eb57506001600160a01b0381163b155b6155cb5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156155c4565b60ff5f516020615eaf5f395f51905f525460401c161561561057565b631afcd79f60e31b5f5260045ffd5b6156276155f4565b62015180600155565b80548210156127cf575f5260205f2001905f90565b8054801561566c575f19019061565b8282615630565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f1461571b575f198401848111610ec95783545f19810194908511610ec9575f9585836156ce97610bdd95036156d4575b505050615645565b55600190565b6157046156fe916156f56156eb6157129588615630565b90549060031b1c90565b92839187615630565b906137ff565b85905f5260205260405f2090565b555f80806156c6565b505050505f90565b61572b615a72565b615733615ac9565b6040519060208201925f516020615eef5f395f51905f528452604083015260608201524660808201523060a082015260a0815261577160c0826105a8565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116157e4579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15613029575f516001600160a01b038116156157da57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611f7357565b615802816157ef565b8061580b575050565b615814816157ef565b6001810361582b5763f645eedf60e01b5f5260045ffd5b615834816157ef565b6002810361584f575063fce698f760e01b5f5260045260245ffd5b8061585b6003926157ef565b146158635750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f1615899612927565b90156152ce5780516158b557503b156158b157600190565b5f90565b60209150015190565b8147106159185782516001600160a01b03909116925f9182916020018486620186a0f1906158ea612927565b911561591157156158fd575b5050600190565b80516158b557503b156158b1575f806158f6565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461598b57600254600160401b8110156105525761597461595e8260018594016002556002615630565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61599a8282615111565b6152ce57805490600160401b82101561055257826159c261595e846001809601855584615630565b90558054925f520160205260405f2055600190565b614c1d6105da93926159fe60405194859263a9059cbb60e01b602085015260248401613ab2565b03601f1981018452836105a8565b90615a21915f52600660205260405f2061330d565b600181015460a01c60ff1615615a43576003018054918201809211610ec95755565b6004018054918203918211610ec95755565b5f80610ce293602081519101845af4615a6c612927565b91615afb565b615a7a614819565b8051908115615a8a576020012090565b50505f516020615d0f5f395f51905f52548015615aa45790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615ad16148d3565b8051908115615ae1576020012090565b50505f516020615ecf5f395f51905f52548015615aa45790565b90615b1f5750805115615b1057805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b50575b615b30575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b28565b80519091906001600160401b03811161055257615b9a81615b875f516020615cef5f395f51905f52546130c9565b5f516020615cef5f395f51905f52614098565b602092601f8211600114615bcd57615bbc929382915f926141c75750506140d0565b5f516020615cef5f395f51905f5255565b5f516020615cef5f395f51905f525f52601f198216935f516020615f0f5f395f51905f52915f5b868110615c365750836001959610615c1e575b505050811b015f516020615cef5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615c07565b91926020600181928685015181550194019201615bf456feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a26469706673582212201c1b4463b28a449840fe9ae5bd74fc12e3d7143120dfcba65aa3cb91b31ab00d64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// BSCBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeV2MetaData.ABI instead.
var BSCBridgeV2ABI = BSCBridgeV2MetaData.ABI

// BSCBridgeV2BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BSCBridgeV2BinRuntime = "60806040526004361015610022575b3615610018575f80fd5b61002061381c565b005b5f3560e01c806301ffc9a7146103915780630b1d8d061461038c5780631313996b146103875780631938e0f214610382578063248a9ca31461037d5780632f2ff15d1461037857806336568abe14610373578063365d71e41461036e57806342cde4e81461036957806348a00ed8146103645780634be13f831461035f5780634d5d00561461035a5780634ee078ba146103555780634f1ef28614610350578063502cc5c01461034b57806352d1902d146103465780635b605f5c146103415780635c975abb1461033c5780635fd262de14610337578063670e626814610332578063792148741461032d5780637b429ad814610328578063814914b514610323578063815f1de61461031e57806384b0196e1461031957806386e9e1751461031457806388d67d6d1461030f57806389232a001461030a5780638ae87c5c1461030557806391cca3db1461030057806391cf6d3e146102fb57806391d14854146102f657806394ddc8c6146102f15780639f9b4f1c146102ec578063a217fddf146102e7578063a3246ad3146102e2578063aa1bd0bc146102dd578063aa20ed47146102d8578063ad3cb1cc146102d3578063ae6893f8146102ce578063b2b49e2e146102c9578063b33eb36e146102c4578063b7f3358d146102bf578063b8aa837e146102ba578063bc8cadd6146102b5578063bedb86fb146102b0578063cf56118e146102ab578063d477f05f146102a6578063d547741f146102a1578063d5717fc51461029c578063e2af6cd714610297578063edbbea3914610292578063f0702e8e1461028d578063f4509643146102885763f698da250361000e576126e2565b612641565b612619565b6124c1565b61240e565b6123d5565b6123a6565b612347565b6122d3565b6121e9565b6121c1565b612119565b612095565b612004565b611ef6565b611ebd565b611e76565b611deb565b611d9e565b611d22565b611cc6565b611baa565b611b01565b611ab3565b611a96565b611a6e565b611a05565b611900565b6117f6565b61163f565b61158b565b61149d565b6113d3565b6112fd565b61128d565b6111d2565b6110ca565b61109c565b610fbd565b610ece565b610e3e565b610ce5565b610b7c565b610b0f565b610abb565b61099c565b610970565b6108fd565b61080d565b6107d2565b6107a1565b6106f0565b6104aa565b610408565b346103e75760203660031901126103e75760043563ffffffff60e01b81168091036103e757602090637965db0b60e01b81149081156103d6575b506040519015158152f35b6301ffc9a760e01b1490505f6103cb565b5f80fd5b6001600160a01b031690565b6001600160a01b038116036103e757565b346103e75760203660031901126103e757600435610425816103f7565b61042e3361496d565b6001600160a01b031661044281151561274e565b603280546001600160a01b031916821790557fa019822918678fc46796568754f842027eb216950ae7c9bc8b6db399678f14f35f80a2005b9181601f840112156103e7578235916001600160401b0383116103e7576020808501948460051b0101116103e757565b60403660031901126103e7576004356001600160401b0381116103e7576104d590369060040161047a565b602435916001600160401b0383116103e757366023840112156103e7576004830135916001600160401b0383116103e75736602460e08502860101116103e757602461002094019161296d565b634e487b7160e01b5f52604160045260245ffd5b61010081019081106001600160401b0382111761055257604052565b610522565b60e081019081106001600160401b0382111761055257604052565b606081019081106001600160401b0382111761055257604052565b60c081019081106001600160401b0382111761055257604052565b601f909101601f19168101906001600160401b0382119082101761055257604052565b604051906105da6080836105a8565b565b604051906105da60e0836105a8565b604051906105da610100836105a8565b604051906105da6060836105a8565b6001600160401b0381116105525760051b60200190565b60ff8116036103e757565b9080601f830112156103e75781356106438161060a565b9261065160405194856105a8565b81845260208085019260051b8201019283116103e757602001905b8282106106795750505090565b60208091833561068881610621565b81520191019061066c565b9080601f830112156103e75781356106aa8161060a565b926106b860405194856105a8565b81845260208085019260051b8201019283116103e757602001905b8282106106e05750505090565b81358152602091820191016106d3565b60803660031901126103e7576004356001600160401b0381116103e75760c060031982360301126103e7576024356001600160401b0381116103e75761073a90369060040161062c565b906044356001600160401b0381116103e75761075a903690600401610693565b60643591906001600160401b0383116103e75761079d9361078261078b943690600401610693565b92600401612aab565b60405191829182901515815260200190565b0390f35b346103e75760203660031901126103e75760206107bf600435612def565b604051908152f35b35906105da826103f7565b346103e75760403660031901126103e7576100206024356004356107f5826103f7565b61080861080182612def565b3390614b8d565b614370565b346103e75760403660031901126103e75760043560243561082d816103f7565b336001600160a01b0382160361084657610020916143d0565b63334bd91960e11b5f5260045ffd5b9060406003198301126103e7576004356001600160401b0381116103e7578261088091600401610693565b91602435906001600160401b0382116103e757806023830112156103e75781600401356108ac8161060a565b926108ba60405194856105a8565b8184526024602085019260051b8201019283116103e757602401905b8282106108e35750505090565b6020809183356108f2816103f7565b8152019101906108d6565b346103e75761090b36610855565b906109198151835114612e0d565b5f5b8251811015610020578061095f61093460019385612e23565b51838060a01b036109458488612e23565b51169061095a3361095583612def565b614b8d565b6143d0565b500161091b565b5f9103126103e757565b346103e7575f3660031901126103e757602060ff5f516020615c8f5f395f51905f525416604051908152f35b346103e75760603660031901126103e7576024356004356044356109bf816103f7565b6109c8336149f5565b6109d0613892565b815f5260076020526109ee836109e98160405f20615111565b612e37565b806109f98484614f8e565b916001600160a01b031615610aa7575b8151915f516020615caf5f395f51905f526040820192610a56610a44855160018060a01b03169660808601978489519160a089015193613c48565b610a4d81611f69565b6001811461442f565b8251602090930151935194516040516001600160a01b0390961695918291610a819142919084612dd1565b0390a45f516020615c4f5f395f51905f525f80a35f5f516020615e2f5f395f51905f525d005b5060608101516001600160a01b0316610a09565b346103e7575f3660031901126103e7575f546040516001600160a01b039091168152602090f35b9181601f840112156103e7578235916001600160401b0383116103e757602083818601950101116103e757565b6101c03660031901126103e757602435600435610b2b826103f7565b604435610b37816103f7565b6064359060a43560843560c4356001600160401b0381116103e757610b60903690600401610ae2565b94909360e03660e31901126103e75761079d9761078b97612e51565b346103e75760403660031901126103e757610c64602435600435610b9e61386b565b610ba6613892565b805f526007602052610bbf826109e98160405f20615111565b610c5f6040610bef610bea85610bdd865f52600860205260405f2090565b905f5260205260405f2090565b613205565b610c4c610c0f82516080610c05868301516103eb565b9101519087613acd565b50610c1981611f69565b610c2281611f69565b83516020810182905290600190610c4683604081015b03601f1981018552846105a8565b1461323e565b01518015908115610c6c575b429161326a565b614452565b6100206138c7565b4281109150610c58565b6001600160401b03811161055257601f01601f191660200190565b929192610c9d82610c76565b91610cab60405193846105a8565b8294818452818301116103e7578281602093845f960137010152565b9080601f830112156103e757816020610ce293359101610c91565b90565b60403660031901126103e757600435610cfd816103f7565b6024356001600160401b0381116103e757610d1c903690600401610cc7565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610e1c575b50610e0d57610d603361496d565b6040516352d1902d60e01b8152916020836004816001600160a01b0386165afa5f9381610ddc575b50610da957634c9c8ce360e01b5f526001600160a01b03821660045260245ffd5b905f516020615d4f5f395f51905f528303610dc85761002092506154fa565b632a87526960e21b5f52600483905260245ffd5b610dff91945060203d602011610e06575b610df781836105a8565b810190613c39565b925f610d88565b503d610ded565b63703e46dd60e11b5f5260045ffd5b5f516020615d4f5f395f51905f52546001600160a01b0316141590505f610d52565b346103e75760603660031901126103e757602435600435604435610e61336149f5565b815f526007602052610e7a836109e98160405f20615111565b4201804211610ec95760207fe11d2f45f48f5ac157e44f5f1b7bac1d69219e5d4e904e17eb60db798536594991835f526008825260405f20855f52825280600760405f200155604051908152a3005b613071565b346103e7575f3660031901126103e7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610e0d5760206040515f516020615d4f5f395f51905f528152f35b60c0809160018060a01b03815116845260018060a01b0360208201511660208501526040810151151560408501526060810151151560608501526080810151608085015260a081015160a08501520151910152565b60206040818301928281528451809452019201905f5b818110610f9d5750505090565b909192602060e082610fb26001948851610f25565b019401929101610f90565b346103e75760203660031901126103e757600435805f52600560205260405f2090604051808360208295549384815201905f5260205f20925f5b81811061108357505061100c925003836105a8565b61101682516132be565b915f5b81518110156110755760019061105961105461103486612cd1565b61104e6110418588612e23565b516001600160a01b031690565b9061330d565b613322565b6110638287612e23565b5261106e8186612e23565b5001611019565b6040518061079d8682610f7a565b8454835260019485019487945060209093019201610ff7565b346103e7575f3660031901126103e757602060ff5f516020615def5f395f51905f5254166040519015158152f35b60e03660031901126103e7576024356004356110e5826103f7565b604435916110f2836103f7565b6064359260c435916084359160a435916001600160401b0385116103e7576111a396611158611128611199973690600401610ae2565b95909661113361386b565b6001600160a01b03851694849061114a878d6144f7565b611152613892565b8b614597565b939092604051986111688a610536565b895260208901523360408901526001600160a01b03166060880152608087015260a086015260c08501523691610c91565b60e0820152614742565b5f5f516020615e2f5f395f51905f525d60405160018152602090f35b6001600160a01b03909116815260200190565b346103e75760803660031901126103e7576024356004356111f2826103f7565b604435906001600160401b0382116103e757366023830112156103e75761079d9261122a61123d933690602481600401359101610c91565b906064359261123884610621565b6133aa565b604051918291826111bf565b90602080835192838152019201905f5b8181106112665750505090565b8251845260209384019390920191600101611259565b906020610ce2928181520190611249565b346103e75760203660031901126103e7576004355f52600760205260405f206040519081602082549182815201915f5260205f20905f5b8181106112e75761079d856112db818703826105a8565b6040519182918261127c565b82548452602090930192600192830192016112c4565b346103e75760403660031901126103e75760043560243561131d816103f7565b6113263361496d565b5f516020615eaf5f395f51905f52549160ff8360401c1680156113af575b6113a0576001600160401b03199092166002175f516020615eaf5f395f51905f525561137791611372613473565b6134c5565b61137f613494565b604051600281525f516020615ccf5f395f51905f529080602081015b0390a1005b63f92ee8a960e01b5f5260045ffd5b5060026001600160401b0384161015611344565b60e0810192916105da9190610f25565b346103e75760403660031901126103e75761079d6114126024356004356113f9826103f7565b611401613288565b505f52600660205260405f2061330d565b60046040519161142183610557565b80546001600160a01b0316835260018101546114769061146d90611450611447826103eb565b602088016130ba565b61146460a082901c60ff1615156040880152565b60a81c60ff1690565b15156060850152565b60028101546080840152600381015460a0840152015460c0820152604051918291826113c3565b346103e75760203660031901126103e7576004356114ba816103f7565b6114c33361496d565b603380546001600160a01b0319166001600160a01b039290921691821790557ff1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d5f80a2005b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9161156190611553610ce297959693600f60f81b865260e0602087015260e0860190611508565b908482036040860152611508565b60608301949094526001600160a01b031660808201525f60a082015280830360c090910152611249565b346103e7575f3660031901126103e7575f516020615d0f5f395f51905f5254158061161f575b156115e2576115be614819565b6115c66148d3565b9061079d6115d261359b565b604051938493309146918661152c565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020615ecf5f395f51905f5254156115b1565b801515036103e757565b346103e75760603660031901126103e75760043561165c816103f7565b60243560443561166b81611635565b61167361386b565b61167b613892565b6001600160a01b0383165f9081526066602052604090206116a7906116a2905b5460ff1690565b6135b6565b6116b28215156134af565b611769576116cd81836116c66065546103eb565b3390614bd9565b335b5f516020615f2f5f395f51905f526116e8606454614c22565b92606454926117526116fb6065546103eb565b604080516001600160a01b039283168152600160208201529882169089015260608801939093525f6080880181905260a0880181905261010060c089018190528801524260e0880152911694908190610120820190565b0390a461175d6138c7565b60405160018152602090f35b6117723361496d565b306116cf565b9080601f830112156103e757813561178f8161060a565b9261179d60405194856105a8565b81845260208085019260051b820101918383116103e75760208201905b8382106117c957505050505090565b81356001600160401b0381116103e7576020916117eb87848094880101610693565b8152019101906117ba565b60803660031901126103e7576004356001600160401b0381116103e75761182190369060040161047a565b90602435906001600160401b0382116103e757366023830112156103e757816004013561184d8161060a565b9261185b60405194856105a8565b8184526024602085019260051b820101903682116103e75760248101925b8284106118d15750506044359150506001600160401b0381116103e7576118a4903690600401611778565b606435926001600160401b0384116103e75761079d946118cb61078b953690600401611778565b936135cc565b83356001600160401b0381116103e7576020916118f583926024369187010161062c565b815201930192611879565b346103e75760603660031901126103e75760043561191d816103f7565b6024359061192a826103f7565b60443561193681610621565b5f516020615eaf5f395f51905f52549260ff604085901c1615611958565b1590565b936001600160401b0316801590816119fd575b60011490816119f3575b1590816119ea575b506113a0575f516020615eaf5f395f51905f5280546001600160401b03191660011790556119af92846119dd57614c3b565b6119b557005b6119bd613494565b604051600181525f516020615ccf5f395f51905f5290806020810161139b565b6119e5613473565b614c3b565b9050155f61197d565b303b159150611975565b85915061196b565b346103e75760403660031901126103e757602435600435611a253361496d565b611a2d613892565b611a378282614f8e565b507fe52d2dff3f36d32a1e8e3b6b1516a5f16949d1b73fe2ce4c281d878d21e2eaf35f80a35f5f516020615e2f5f395f51905f525d005b346103e7575f3660031901126103e7576034546040516001600160a01b039091168152602090f35b346103e7575f3660031901126103e7576020603554604051908152f35b346103e75760403660031901126103e757602060ff611af5602435600435611ada826103f7565b5f525f516020615dcf5f395f51905f52845260405f2061330d565b54166040519015158152f35b346103e75760603660031901126103e757602435600435611b21826103f7565b7f493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea6020604435611b5081611635565b611b5933614a7d565b835f5260058252611b9f816001611b9960405f2098611b8981611b84858060a01b038216809d615111565b613662565b885f526006875260405f2061330d565b0161368a565b6040519015158152a3005b346103e75760403660031901126103e7576004356001600160401b0381116103e757611bda903690600401610693565b6024356001600160401b0381116103e757611bf9903690600401610693565b90611c078151835114612783565b5f5b82518110156100205780611cb8611c2260019385612e23565b51611c2d8387612e23565b5190611c3761386b565b611c3f613892565b805f526007602052611c58826109e98160405f20615111565b610c5f6040611c76610bea85610bdd865f52600860205260405f2090565b610c4c611c8c82516080610c05868301516103eb565b50611c9681611f69565b611c9f81611f69565b835160208101829052908a90610c468360408101610c38565b611cc06138c7565b01611c09565b346103e7575f3660031901126103e75760206040515f8152f35b60206040818301928281528451809452019201905f5b818110611d035750505090565b82516001600160a01b0316845260209384019390920191600101611cf6565b346103e75760203660031901126103e7576004355f525f516020615d6f5f395f51905f5260205260405f206040519081602082549182815201915f5260205f20905f5b818110611d885761079d85611d7c818703826105a8565b60405191829182611ce0565b8254845260209093019260019283019201611d65565b346103e75760203660031901126103e7577fef10cb70bdfa4ec907a771a1e5b478dd1da8b5514c41fc074405d9c5ea52d8b36020600435611dde3361496d565b80600155604051908152a1005b346103e75760403660031901126103e757602435600435611e0b336149f5565b611e14336149f5565b611e1c613892565b805f526007602052611e35826109e98160405f20615111565b611e3f8282614452565b5f516020615c4f5f395f51905f525f80a35f5f516020615e2f5f395f51905f525d005b60405190611e716020836105a8565b5f8252565b346103e7575f3660031901126103e75761079d604051611e976040826105a8565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611508565b346103e75760203660031901126103e7576004355f526004602052600160405f20015460018101809111610ec957602090604051908152f35b346103e757611f0436610855565b90611f128151835114612e0d565b5f5b82518110156100205780611f4e611f2d60019385612e23565b51838060a01b03611f3e8488612e23565b5116906108083361095583612def565b5001611f14565b634e487b7160e01b5f52602160045260245ffd5b600a1115611f7357565b611f55565b90600a821015611f735752565b6020815260606040611fea60a0855184602087015280516080870152602081015182870152600180831b03848201511660c0870152600180831b03858201511660e08701526080810151610100870152015160c0610120860152610140850190611508565b93611ffc602082015183860190611f78565b015191015290565b346103e75760403660031901126103e757600435602435905f6040805161202a81610572565b6120326136a7565b815282602082015201525f52600860205260405f20905f5260205261079d60405f2060076040519161206383610572565b61206c81613101565b835261208260ff600683015416602085016131f9565b0154604082015260405191829182611f85565b346103e75760203660031901126103e7577f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff602060ff6004356120d781610621565b6120e03361496d565b166120ec8115156136d7565b8060ff195f516020615c8f5f395f51905f525416175f516020615c8f5f395f51905f5255604051908152a1005b346103e75760403660031901126103e75760043560243561213981611635565b61214233614a7d565b612157825f52600360205260405f2054151590565b156121ae5760207f41ee956671dd2884b868e36b1976ae869e7977d8cc9476d99563a0c0c1f3c67591835f52600482526121a381600360405f20019060ff801983541691151516179055565b6040519015158152a2005b5063ac7dbbfd60e01b5f5260045260245ffd5b346103e7575f3660031901126103e7576033546040516001600160a01b039091168152602090f35b346103e75760203660031901126103e75760043561220681611635565b61220f33614a7d565b1561226d5761221c61386b565b600160ff195f516020615def5f395f51905f525416175f516020615def5f395f51905f52557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b5f516020615def5f395f51905f525460ff8116156122c45760ff19165f516020615def5f395f51905f52557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b638dfc202b60e01b5f5260045ffd5b346103e7575f3660031901126103e757604051600280548083525f91825260208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91905b8181106123315761079d856112db818703826105a8565b825484526020909301926001928301920161231a565b346103e75760203660031901126103e757600435612364816103f7565b61236d3361496d565b6001600160a01b031661238181151561274e565b603480546001600160a01b031916821790555f516020615d8f5f395f51905f525f80a2005b346103e75760403660031901126103e7576100206024356004356123c9826103f7565b61095a61080182612def565b346103e75760203660031901126103e7576004355f526004602052600260405f20015460018101809111610ec957602090604051908152f35b346103e75760203660031901126103e75760043561242b816103f7565b6124343361496d565b5f546001600160a01b0316806124af57506124596001600160a01b03821615156136ed565b5f80546001600160a01b0319166001600160a01b03928316179081905561248890612483906103eb565b6103eb565b167fedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee5f80a2005b63f6c75f8f60e01b5f5260045260245ffd5b346103e75760803660031901126103e7576004356024356124e181611635565b604435916124ee836103f7565b5f516020615e6f5f395f51905f526125d86064359361250c856103f7565b61251533614b05565b6125208415156136ed565b6001600160a01b0386169461078b9061253a8715156136ed565b6001600160a01b038116976125d3906125548a15156136ed565b61255d88615927565b6125dd575b612586816125818161257c8c5f52600560205260405f2090565b6152d4565b61506a565b6125a56125916105dc565b9361259c83866130ba565b602085016130ba565b84151560408401525f60608401525f60808401525f60a08401525f60c08401526125ce88612cd1565b61330d565b615092565b0390a4005b6126146125e86105cb565b8981525f60208201525f60408201525f606082015261260f8a5f52600460205260405f2090565b615033565b612562565b346103e7575f3660031901126103e7576032546040516001600160a01b039091168152602090f35b346103e75760403660031901126103e757602435600435612661826103f7565b61266a3361496d565b805f5260056020525f60046126a5604083209461269481611b8460018060a01b0382168099615680565b84845260066020526040842061330d565b82815582600182015582600282015582600382015501557fa7353976a33b36069eee140579faabe142bd14e7f8dd5496ac9199bd862edd4d5f80a3005b346103e7575f3660031901126103e7576126fa615a72565b612702615ac9565b6040519060208201925f516020615eef5f395f51905f528452604083015260608201524660808201523060a082015260a0815261274060c0826105a8565b519020604051908152602090f35b1561275557565b638219abe360e01b5f5260045ffd5b80546001600160a01b0319166001600160a01b03909216919091179055565b1561278a57565b6329f517fd60e01b5f5260045ffd5b634e487b7160e01b5f52603260045260245ffd5b91908110156127cf5760051b8101359060fe19813603018212156103e7570190565b612799565b35610ce2816103f7565b903590601e19813603018212156103e757018035906001600160401b0382116103e7576020019181360383136103e757565b91908110156127cf5760e0020190565b908160209103126103e75751610ce281611635565b908060209392818452848401375f828201840152601f01601f1916010190565b9792936101a0979460c097936128a697939c9b9c8b5260018060a01b031660208b015260018060a01b031660408a01526060890152608088015260a08701526101c0848701526101c0860191612835565b9480356128b2816103f7565b6001600160a01b031660e085015260208101356128ce816103f7565b6001600160a01b03166101008501526040810135610120850152606081013561014085015260ff608082013561290381610621565b1661016085015260a08101356101808501520135910152565b6040513d5f823e3d90fd5b3d15612951573d9061293882610c76565b9161294660405193846105a8565b82523d5f602084013e565b606090565b604090610ce2939281528160208201520190611508565b909193929361297d858414612783565b5f945b83861061298f57505050509050565b61299a8685856127ad565b3560206129b2816129ac8a89896127ad565b016127d4565b6129c260606129ac8b8a8a6127ad565b92612a388a86888a8c612a1c60806129db8784866127ad565b013595612a14612a0a8260a06129f282888a6127ad565b01359560c0612a0283838b6127ad565b0135976127ad565b60e08101906127de565b969095612810565b946040519a8b998a996326ae802b60e11b8b5260048b01612855565b03815f305af19081612a7f575b50612a745785612a53612927565b60405163f495148b60e01b8152918291612a709160048401612956565b0390fd5b600190950194612980565b612a9f9060203d8111612aa4575b612a9781836105a8565b810190612820565b612a45565b503d612a8d565b90612b90939291612aba61386b565b612ac2613892565b803592612ad7845f52600560205260405f2090565b91612b07612af56040830194612aef612483876127d4565b906138d9565b612b01612483866127d4565b90612cdf565b612b12343415612d07565b612ba6612b32865f526004602052600260405f2001600181540180915590565b96612b45602084013589819a8214612d25565b612b51612483866127d4565b93606084019688612b61896127d4565b96612b9e8c60808901359e8f60a08b019b612b7c8d8d6127de565b929091604051978896602088019a8b612d43565b03601f1981018352826105a8565b51902061391c565b612bb987612bb3856127d4565b87613b1b565b919092600184612bc881611f69565b14612c95575b50612bd883611f69565b60018303612c325750505090612c04612bfe5f516020615caf5f395f51905f52936127d4565b916127d4565b6040516001600160a01b03909216958291612c2191429184612dd1565b0390a45b612c2d6138c7565b600190565b612c6f8394612c6a612c8d947f17b82bf8aab2947c6cef6b829c29275d00f666ec463cc32341c423e9c856229f9896612c7595614242565b6127d4565b926127d4565b9260405193849360018060a01b031698429185612da7565b0390a4612c25565b612cca91935088612ca5866127d4565b91612cc3612cbc612cb58a6127d4565b92886127de565b3691610c91565b9289613c48565b915f612bce565b5f52600660205260405f2090565b15612ce75750565b633142cb1160e11b5f9081526001600160a01b0391909116600452602490fd5b15612d0f5750565b63943892eb60e01b5f525f60045260245260445ffd5b15612d2e575050565b63a6ab65c560e01b5f5260045260245260445ffd5b92610ce297959260e095927fd03fbc82682decb107729b5ed42c725f2c3d2fc7a69eb5ffa0daffc7b772191486526020860152604085015260018060a01b0316606084015260018060a01b0316608083015260a08201528160c08201520191612835565b6105da93606092969593608083019760018060a01b03168352602083015260408201520190611f78565b604091949392606082019560018060a01b0316825260208201520152565b5f525f516020615dcf5f395f51905f52602052600160405f20015490565b15612e1457565b63031206d560e51b5f5260045ffd5b80518210156127cf5760209160051b010190565b15612e3f5750565b6373a1390160e11b5f5260045260245ffd5b959394612ed3919597939297612e6561386b565b612eaa6001600160a01b038816612e7c818b6144f7565b612e84613892565b612e9461248361248360e46127d4565b811490612ea461248360e46127d4565b9161302e565b612ecb612ebb6124836101046127d4565b6001600160a01b038b161461305b565b838789614597565b9161012435612f0681612eef86612eea8787613085565b613085565b811015612f0087612eea8888613085565b90613092565b612f146124836032546103eb565b90612f206101046127d4565b906101443592612f316101646130b0565b92610184356101a43591833b156103e757604051637f40ec1760e11b81526001600160a01b03808f16600483015290911660248201523060448201526064810194909452608484019590955260ff90931660a483015260c482019390935260e4810191909152905f908290818381610104810103925af198891561302957612ff061119998612ff993612c259c61300f575b50612fe7612fd26101046127d4565b91612fdb6105eb565b9c8d5260208d016130ba565b60408b016130ba565b606089016130ba565b608087015260a086015260c08501523691610c91565b8061301d5f613023936105a8565b80610966565b5f612fc3565b61291c565b15613037575050565b6313f7c32b60e31b5f9081526001600160a01b039182166004529116602452604490fd5b1561306257565b630672aec160e01b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b91908201809211610ec957565b1561309b575050565b63943892eb60e01b5f5260045260245260445ffd5b35610ce281610621565b6001600160a01b039091169052565b90600182811c921680156130f7575b60208310146130e357565b634e487b7160e01b5f52602260045260245ffd5b91607f16916130d8565b9060405161310e8161058d565b825481526001830154602082015260028301546001600160a01b0390811660408084019190915260038501549091166060830152600484015460808301525160059093018054919391849183915f91613166826130c9565b80855291600181169081156131d25750600114613193575b505060a0929161318f9103846105a8565b0152565b5f908152602081209092505b8183106131b657505081016020018161318f61317e565b602091935080600191548385890101520191019091849261319f565b60ff191660208681019190915292151560051b8501909201925083915061318f905061317e565b600a821015611f735752565b9060405161321281610572565b60406007829461322181613101565b845261323760ff600683015416602086016131f9565b0154910152565b156132465750565b60405162461bcd60e51b815260206004820152908190612a70906024830190611508565b15613273575050565b637a88099160e11b5f5260045260245260445ffd5b6040519061329582610557565b5f60c0838281528260208201528260408201528260608201528260808201528260a08201520152565b906132c88261060a565b6132d560405191826105a8565b82815280926132e6601f199161060a565b01905f5b8281106132f657505050565b602090613301613288565b828285010152016132ea565b9060018060a01b03165f5260205260405f2090565b9060405161332f81610557565b82546001600160a01b0390811682526001840154908116602083015260a081901c60ff16151560408301529092839160c09160049161337a9061337190611464565b15156060860152565b60028101546080850152600381015460a08501520154910152565b908160209103126103e75751610ce2816103f7565b5f5490949392906001600160a01b0381161561346457604051637c469ea160e11b8152600481018781526001600160a01b03858116602484015260806044840152919560209487949093169284925f92849290919060609060ff90613413906084870190611508565b931691015203925af18015613029576105da925f91613435575b508094613703565b613457915060203d60201161345d575b61344f81836105a8565b810190613395565b5f61342d565b503d613445565b6315aeca0d60e11b5f5260045ffd5b5f516020615eaf5f395f51905f52805460ff60401b1916600160401b179055565b5f516020615eaf5f395f51905f52805460ff60401b19169055565b156134b657565b63ff5a231360e01b5f5260045ffd5b906134d18215156134af565b6001600160a01b031690811561358c57606455606580546001600160a01b03191691909117905560666020527f7bff7289460e40a7871440a623b3e58f37e81e8af7f28690a5cdfb556e92f2ad805460ff1990811660019081179092557fda2e410554560bb4f7657c5a29466761154072b10a25bb76ed04ee6f7e37a55c805482168317905561dead60901b5f527fd88e03f995b4ae15110a81bb3b46f829091e220d775e2d8ab911a127f494a91780549091169091179055565b63643b674560e01b5f5260045ffd5b604051906135aa6020836105a8565b5f808352366020840137565b156135bd57565b635e0c1f8960e01b5f5260045ffd5b9192949390938484511480613658575b8061364e575b6135eb90612783565b5f5b85811015613642578060051b8401359060be19853603018212156103e75761363b60019261361b8389612e23565b51613626848c612e23565b51906136328589612e23565b51928901612aab565b50016135ed565b50945050505050600190565b50815185146135e2565b50848651146135dc565b1561366a5750565b63153096f360e11b5f9081526001600160a01b0391909116600452602490fd5b805460ff60a81b191691151560a81b60ff60a81b16919091179055565b604051906136b48261058d565b606060a0835f81525f60208201525f60408201525f838201525f60808201520152565b156136de57565b63f0f15b9160e01b5f5260045ffd5b156136f457565b6302bfb33d60e51b5f5260045ffd5b91909161370f33614b05565b61371a8115156136ed565b6001600160a01b038316916137a8906137348415156136ed565b6001600160a01b038116946125d39061374e8715156136ed565b61375785615927565b6137c8575b613776816125818161257c895f52600560205260405f2090565b6137816125916105dc565b5f60408401525f60608401525f60808401525f60a08401525f60c08401526125ce85612cd1565b6040515f81525f516020615e6f5f395f51905f529080602081015b0390a4565b6137fa6137d36105cb565b8681525f60208201525f60408201525f606082015261260f875f52600460205260405f2090565b61375c565b916138189183549060031b91821b915f19901b19161790565b9055565b6033546001600160a01b0316330361383057565b60405162461bcd60e51b815260206004820152601360248201527227b7363c90213934b233b2a2bc32b1baba32b960691b6044820152606490fd5b60ff5f516020615def5f395f51905f52541661388357565b63d93c066560e01b5f5260045ffd5b5f516020615e2f5f395f51905f525c6138b85760015f516020615e2f5f395f51905f525d565b633ee5aeb560e01b5f5260045ffd5b5f5f516020615e2f5f395f51905f525d565b610ce2916001600160a01b031690615111565b156138f357565b63b3f07a3960e01b5f5260045ffd5b1561390a5750565b631aebd17960e11b5f5260045260245ffd5b9392938151918351831480613a7a575b613935906138ec565b6139566139505f516020615c8f5f395f51905f525460ff1690565b60ff1690565b956139648488811015613902565b5f945f935f5b86811061398557505050505050506105da9192811015613902565b6139e26139b183613994615723565b6042916040519161190160f01b8352600283015260228201522090565b6139c56139be8489612e23565b5160ff1690565b6139cf8487612e23565b51906139db8589612e23565b5192615124565b6001600160a01b038181169088161080613a13575b613a05575b5060010161396a565b6001988901989096506139fc565b507f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989265f525f516020615dcf5f395f51905f52602052613a7561169b827fb323bca1072ad6c43c8a5019bad43672f13be3697b760574a9a2f43b558a3aa561330d565b6139f7565b508551831461392c565b15613a8b57565b6330d45fb160e01b5f5260045ffd5b908160209103126103e75751600a8110156103e75790565b6001600160a01b039091168152602081019190915260400190565b9150613b08611054606092613af75f95613af260325460018060a01b03161515613a84565b612cd1565b6001600160a01b039091169061330d565b0151613b1357600191565b506002905f90565b9190915f92613b686060613b61611054613b51613b3c6124836032546103eb565b96613af26001600160a01b0389161515613a84565b6001600160a01b0386169061330d565b0151151590565b613bef5791602091613b9295935f604051809881958294633f4bdec560e01b845260048401613ab2565b03925af1928315613029575f93613bbe575b50600183613bb181611f69565b03613bb857565b60019150565b613be191935060203d602011613be8575b613bd981836105a8565b810190613a9a565b915f613ba4565b503d613bcf565b505050506002905f90565b90815260208101919091526001600160a01b03918216604082015291166060820152608081019190915260c060a08201819052610ce292910190611508565b908160209103126103e7575190565b92919382613c5585612cd1565b95613c786001613c6d818060a01b038716809a61330d565b015460a01c60ff1690565b92601481511180613f7a575b613cbd575b50613c939361513c565b92613c9d84611f69565b60018414613cac575b50505090565b613cb5926151fb565b5f8080613ca6565b9150613cf593602083015160601c6020613cdb6124836033546103eb565b916040518098819262483e5560e91b8352600483016111bf565b0381845afa938415613029578888958b985f91613f5b575b50613d1b5750505093613c89565b92809585988896936002613d396001985f52600460205260405f2090565b01549687921480159889613f20575b5060209596505f14613f15578c613d8784925b613d696124836033546103eb565b95604051998a9889978896631eeed51360e01b885260048801613bfa565b03925af1908115613029575f91613ef6575b50613ee6579087869392887f9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be160405180613df4898c83909160809260018060a01b031682526020820152606060408201525f60608201520190565b0390a483159081613ede575b50613e10575b92613c939361513c565b9050613e206124836033546103eb565b92604051916370a0823160e01b835260208380613e4088600483016111bf565b03818b5afa80156130295786935f91613ebf575b5080613e64575b50909350613e06565b6020919350613e88956040519687928392632770a7eb60e21b845260048401613ab2565b03815f8b5af191821561302957613c9394869315613e5b57613eb89060203d602011612aa457612a9781836105a8565b505f613e5b565b613ed8915060203d602011610e0657610df781836105a8565b5f613e54565b90505f613e00565b50505050509091612c2d926151fb565b613f0f915060203d602011612aa457612a9781836105a8565b5f613d99565b8c613d875f92613d5b565b939250945094613f2f9361513c565b613f3881611f69565b60018103613f4d5750848884848a948a613d48565b995050505050505050505090565b613f74915060203d602011612aa457612a9781836105a8565b5f613d0d565b506033546001600160a01b0390613f9490612483906103eb565b161515613c84565b15613fa45750565b6307a5c91d60e31b5f5260045260245ffd5b919060c0838203126103e75760405190613fcf8261058d565b819380358352602081013560208401526040810135613fed816103f7565b6040840152613ffe606082016107c7565b60608401526080818101359084015260a0810135916001600160401b0383116103e75760a09261402e9201610cc7565b910152565b81811061403e575050565b5f8155600101614033565b90601f8211614056575050565b6105da915f516020615c6f5f395f51905f525f5260205f20906020601f840160051c8301931061408e575b601f0160051c0190614033565b9091508190614081565b9190601f81116140a757505050565b6105da925f5260205f20906020601f840160051c8301931061408e57601f0160051c0190614033565b8160011b915f199060031b1c19161790565b90600a811015611f735760ff80198354169116179055565b815180518255602081015160018301556040810151919291614128906001600160a01b031660028501612764565b6060810151614143906001600160a01b031660038501612764565b6080810151600484015560a00151805160058401916001600160401b0382116105525761417a8261417485546130c9565b85614098565b602090601f83116001146141d2578260079593604095936141a3935f926141c7575b50506140d0565b90555b6141c060208201516141b781611f69565b600686016140e2565b0151910155565b015190505f8061419c565b90601f198316916141e6855f5260205f2090565b925f5b81811061422a5750926001928592600798966040989610614212575b505050811b0190556141a6565b01515f1960f88460031b161c191690555f8080614205565b929360206001819287860151815501950193016141e9565b9160806142e06142d1600293613af2876142cc61427c61381899833595865f52600760205261428160405f20602087013594858092615990565b613f9c565b156142ed576142b561429560015442613085565b915b6142aa6142a26105fb565b963690613fb6565b8652602086016131f9565b6040840152610bdd855f52600860205260405f2090565b6140fa565b61104e612483604088016127d4565b9301359201918254613085565b6142b55f91614297565b90614302825f615244565b918261430b5750565b5f80525f516020615d6f5f395f51905f526020526001600160a01b0316614352817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d615990565b1561435a5750565b63d180cb3160e01b5f526004525f60245260445ffd5b91909161437d8382615244565b9283614387575050565b815f525f516020615d6f5f395f51905f526020526143b260405f209160018060a01b03168092615990565b156143bb575050565b63d180cb3160e01b5f5260045260245260445ffd5b9190916143dd83826152e7565b92836143e7575050565b815f525f516020615d6f5f395f51905f5260205261441260405f209160018060a01b03168092615680565b1561441b575050565b62a95f1b60e31b5f5260045260245260445ffd5b156144375750565b63290cd55f60e01b5f52600a811015611f735760045260245ffd5b9061445c91614f8e565b60018060a01b036060820151168151915f516020615caf5f395f51905f5260408201926144a4610a44855160018060a01b03169660808601978489519160a089015193613c48565b8251602090930151935194516040516001600160a01b03909616959182916137c39142919084612dd1565b156144d75750565b6338384f6f60e11b5f9081526001600160a01b0391909116600452602490fd5b9190825f52600560205261451b81611b8460405f2060018060a01b03831690615111565b825f52600460205260ff600360405f200154166145535760ff6001614547836125ce6105da9697612cd1565b015460a81c16156144cf565b82636fc24b7960e11b5f5260045260245ffd5b908160609103126103e7578051916040602083015192015190565b1561458857565b6358e8878560e01b5f5260045ffd5b82606091614630979596936145c16110546145b184612cd1565b6001600160a01b0384169061330d565b6145d16119546040830151151590565b6146d2575b506145e56124836032546103eb565b916145fa6001600160a01b0384161515613a84565b604051998a94859384936337dba1f760e21b8552600485016040919493926060820195825260018060a01b031660208201520152565b03915afa908115613029575f955f905f93614694575b509082916105da94939681988515159586614689575b50508461467e575b505082614673575b5050614581565b101590505f8061466c565b101592505f80614664565b101594505f8061465c565b90506146bf9196506105da93925060603d6060116146cb575b6146b781836105a8565b810190614566565b91969293919291614646565b503d6146ad565b60c06146e49101518480821015613092565b5f6145d6565b6001600160a01b039182168152918116602083015290911660408201526060810191909152608081019190915260a081019190915261010060c082018190529093929160e09161473d9190860190611508565b930152565b61474c8151614c22565b6147568251612cd1565b5f516020615f2f5f395f51905f526147886124836001614781602088019561104e61248388516103eb565b01546103eb565b938051906137c361479985516103eb565b9160408101906147a982516103eb565b906147d26080820196875160a08401948551986147cc60c087019a8b5190613085565b936153b9565b6147e86147e1825199516103eb565b93516103eb565b9460e06147f860608401516103eb565b9751935191519201519260405197889760018060a01b03169c4296896146ea565b604051905f825f516020615c6f5f395f51905f525491614838836130c9565b80835292600181169081156148b4575060011461485c575b6105da925003836105a8565b505f516020615c6f5f395f51905f525f90815290915f516020615d2f5f395f51905f525b8183106148985750509060206105da92820101614850565b6020919350806001915483858901015201910190918492614880565b602092506105da94915060ff191682840152151560051b820101614850565b604051905f825f516020615cef5f395f51905f5254916148f2836130c9565b80835292600181169081156148b45750600114614915576105da925003836105a8565b505f516020615cef5f395f51905f525f90815290915f516020615f0f5f395f51905f525b8183106149515750509060206105da92820101614850565b6020919350806001915483858901015201910190918492614939565b5f516020615e4f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff6149ba827fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c61330d565b5416156149c45750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e4f5f395f51905f52602452604490fd5b5f516020615e8f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614a42827fd56b761303e15a4cbea84f37eba80a5289f49fda0b4884e30ac1d73eb9b8b14361330d565b541615614a4c5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e8f5f395f51905f52602452604490fd5b5f516020615daf5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614aca827f448256db8f8fb95ee3eaaf89c1051414494e85cebb6057fcf996cc3d0ccfb45661330d565b541615614ad45750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615daf5f395f51905f52602452604490fd5b5f516020615e0f5f395f51905f525f525f516020615dcf5f395f51905f5260205260ff614b52827f50e48e6bd160f41806a4807f6d26cc6f1f57f630bdda6fa6bb2608373aa1449f61330d565b541615614b5c5750565b63e2517d3f60e01b5f9081526001600160a01b03919091166004525f516020615e0f5f395f51905f52602452604490fd5b90815f525f516020615dcf5f395f51905f5260205260ff614bb18260405f2061330d565b541615614bbc575050565b63e2517d3f60e01b5f5260018060a01b031660045260245260445ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526105da91614c1d6084836105a8565b61559c565b5f526004602052600160405f2001600181540180915590565b91614c446155f4565b614c586001600160a01b038416151561274e565b6001600160a01b03821692614c6e84151561274e565b60ff821615614ee057614cda90614c836155f4565b614c8b6155f4565b614c936155f4565b60ff195f516020615def5f395f51905f5254165f516020615def5f395f51905f5255614cbd6155f4565b614cc56155f4565b614ccd6155f4565b614cd56155f4565b6142f7565b50614ce36155f4565b614cf160ff821615156136d7565b60408051614cff82826105a8565b60098152682b30b634b230ba37b960b91b6020820152614d21825192836105a8565b60058252640312e302e360dc1b6020830152614d3b6155f4565b614d436155f4565b8051906001600160401b03821161055257614d7482614d6f5f516020615c6f5f395f51905f52546130c9565b614049565b602090601f8311600114614e4c5792614da283614e32979694614db694614e08975f926141c75750506140d0565b5f516020615c6f5f395f51905f5255615b59565b614dcb5f5f516020615d0f5f395f51905f5255565b614de05f5f516020615ecf5f395f51905f5255565b60ff1660ff195f516020615c8f5f395f51905f525416175f516020615c8f5f395f51905f5255565b614e1061561f565b603480546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020615d8f5f395f51905f525f80a26105da43603555565b5f516020615c6f5f395f51905f525f52601f19831691905f516020615d2f5f395f51905f52925f5b818110614ec8575093614db693614e08969360019383614e329b9a9810614eb0575b505050811b015f516020615c6f5f395f51905f5255615b59565b01515f1960f88460031b161c191690555f8080614e96565b92936020600181928786015181550195019301614e74565b6338854f1160e21b5f5260045ffd5b91908203918211610ec957565b60075f9182815582600182015582600282015582600382015582600482015560058101614f2981546130c9565b9081614f3c575b50508260068201550155565b601f8211600114614f5357849055505b5f80614f30565b614f79614f89926001601f614f6b855f5260205f2090565b920160051c82019101614033565b5f81815260208120918190559055565b614f4c565b9190614f986136a7565b50825f526007602052614fae8160405f20615680565b156150215761501c6105da91845f52600860205260405f20815f52602052610bdd614fdb60405f20613101565b95614ff8614fe882612cd1565b61104e61248360408b01516103eb565b61500c600260808a01519201918254614eef565b90555f52600860205260405f2090565b614efc565b637f11bea960e01b5f5260045260245ffd5b600360606105da938051845560208101516001850155604081015160028501550151151591019060ff801983541691151516179055565b156150725750565b6311dd05f360e31b5f9081526001600160a01b0391909116600452602490fd5b9060c06004916150ab60018060a01b0382511685612764565b60208101516150f69060018601906150cc906001600160a01b031682612764565b6040830151815460ff60a01b191690151560a01b60ff60a01b16178155606083015115159061368a565b6080810151600285015560a081015160038501550151910155565b6001915f520160205260405f2054151590565b91610ce2939161513393615777565b909291926157f9565b6001600160a01b0316929190600184146151cb5781156151c25761518b92156151965760405163a9059cbb60e01b602082015291615183918391612b909160248401613ab2565b600592615875565b15610ce25750600190565b6040516340c10f1960e01b6020820152916151ba918391612b909160248401613ab2565b600692615875565b50505050600190565b906151ed935061195492506151de611e62565b916001600160a01b03166158be565b6151f657600190565b600590565b90615210915f52600660205260405f2061330d565b600181015460a01c60ff1615615232576003018054918203918211610ec95755565b6004018054918201809211610ec95755565b805f525f516020615dcf5f395f51905f5260205260ff6152678360405f2061330d565b54166152ce57805f525f516020615dcf5f395f51905f5260205261528e8260405f2061330d565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b610ce2916001600160a01b031690615990565b805f525f516020615dcf5f395f51905f5260205260ff61530a8360405f2061330d565b5416156152ce57805f525f516020615dcf5f395f51905f526020526153328260405f2061330d565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b1561537957505050565b639ac2f56d60e01b5f9081526001600160a01b039182166004529116602452604491909152606490fd5b156153aa57565b63559d141b60e11b5f5260045ffd5b90936001600160a01b038516926001840361542757506105da94506153ef6153e18286613085565b34143490612f008488613085565b806153fb575b50615a0c565b61541c6154219161540d6034546103eb565b90615416611e62565b916158be565b6153a3565b5f6153f5565b90615433343415612d07565b6154486154408287613085565b308489614bd9565b806154dc575b506154646119546001613c6d866125ce87612cd1565b615474575b506105da9350615a0c565b604051632770a7eb60e21b815260208180615493883060048401613ab2565b03815f885af1918215613029576105da966154b79387935f916154bd575b5061536f565b5f615469565b6154d6915060203d602011612aa457612a9781836105a8565b5f6154b1565b6154f4906154ee6124836034546103eb565b876159d7565b5f61544e565b90813b1561557b575f516020615d4f5f395f51905f5280546001600160a01b0319166001600160a01b0384169081179091557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a28051156155635761556091615a55565b50565b50503461556c57565b63b398979f60e01b5f5260045ffd5b50634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b905f602091828151910182855af11561291c575f513d6155eb57506001600160a01b0381163b155b6155cb5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156155c4565b60ff5f516020615eaf5f395f51905f525460401c161561561057565b631afcd79f60e31b5f5260045ffd5b6156276155f4565b62015180600155565b80548210156127cf575f5260205f2001905f90565b8054801561566c575f19019061565b8282615630565b8154905f199060031b1b1916905555565b634e487b7160e01b5f52603160045260245ffd5b6001810191805f528260205260405f2054928315155f1461571b575f198401848111610ec95783545f19810194908511610ec9575f9585836156ce97610bdd95036156d4575b505050615645565b55600190565b6157046156fe916156f56156eb6157129588615630565b90549060031b1c90565b92839187615630565b906137ff565b85905f5260205260405f2090565b555f80806156c6565b505050505f90565b61572b615a72565b615733615ac9565b6040519060208201925f516020615eef5f395f51905f528452604083015260608201524660808201523060a082015260a0815261577160c0826105a8565b51902090565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116157e4579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15613029575f516001600160a01b038116156157da57905f905f90565b505f906001905f90565b5050505f9160039190565b60041115611f7357565b615802816157ef565b8061580b575050565b615814816157ef565b6001810361582b5763f645eedf60e01b5f5260045ffd5b615834816157ef565b6002810361584f575063fce698f760e01b5f5260045260245ffd5b8061585b6003926157ef565b146158635750565b6335e2f38360e21b5f5260045260245ffd5b81516001600160a01b03909116915f9182916020018285620186a0f1615899612927565b90156152ce5780516158b557503b156158b157600190565b5f90565b60209150015190565b8147106159185782516001600160a01b03909116925f9182916020018486620186a0f1906158ea612927565b911561591157156158fd575b5050600190565b80516158b557503b156158b1575f806158f6565b5050505f90565b632b60b36f60e21b5f5260045ffd5b5f8181526003602052604090205461598b57600254600160401b8110156105525761597461595e8260018594016002556002615630565b819391549060031b91821b915f19901b19161790565b9055600254905f52600360205260405f2055600190565b505f90565b61599a8282615111565b6152ce57805490600160401b82101561055257826159c261595e846001809601855584615630565b90558054925f520160205260405f2055600190565b614c1d6105da93926159fe60405194859263a9059cbb60e01b602085015260248401613ab2565b03601f1981018452836105a8565b90615a21915f52600660205260405f2061330d565b600181015460a01c60ff1615615a43576003018054918201809211610ec95755565b6004018054918203918211610ec95755565b5f80610ce293602081519101845af4615a6c612927565b91615afb565b615a7a614819565b8051908115615a8a576020012090565b50505f516020615d0f5f395f51905f52548015615aa45790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b615ad16148d3565b8051908115615ae1576020012090565b50505f516020615ecf5f395f51905f52548015615aa45790565b90615b1f5750805115615b1057805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580615b50575b615b30575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15615b28565b80519091906001600160401b03811161055257615b9a81615b875f516020615cef5f395f51905f52546130c9565b5f516020615cef5f395f51905f52614098565b602092601f8211600114615bcd57615bbc929382915f926141c75750506140d0565b5f516020615cef5f395f51905f5255565b5f516020615cef5f395f51905f525f52601f198216935f516020615f0f5f395f51905f52915f5b868110615c365750836001959610615c1e575b505050811b015f516020615cef5f395f51905f5255565b01515f1960f88460031b161c191690555f8080615c07565b91926020600181928685015181550194019201615bf456feb4e7fb3fec39fae9d12157cf1027fca258064a3bd6401233228a2ecaa3b45db1a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd91320094ee0e8d51fc110ff38d48510c267aeced8608d1c96f14d63e93085dba231d6bc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0076ba73ac0846c1aa28fd3ec5fe6d17993af0d779afde941634cad517abc5087197667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330021d1167972f621f75904fb065136bc8b53c7ba1c60ccd3a7758fbee465851e9c9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775e51c83f65984f32d3377139ea3ea559ee255ee6e9c8f91a6e0a0b299d834e5db0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea09f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1018b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b7518586689aac70df719039c2a1fef16cfe30db5bb5c02c4b6c3d1a73ff2280708a26469706673582212201c1b4463b28a449840fe9ae5bd74fc12e3d7143120dfcba65aa3cb91b31ab00d64736f6c634300081c0033"

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

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Caller) BridgeExecuter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeV2.contract.Call(opts, &out, "bridgeExecuter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2Session) BridgeExecuter() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecuter(&_BSCBridgeV2.CallOpts)
}

// BridgeExecuter is a free data retrieval call binding the contract method 0xbc8cadd6.
//
// Solidity: function bridgeExecuter() view returns(address)
func (_BSCBridgeV2 *BSCBridgeV2CallerSession) BridgeExecuter() (common.Address, error) {
	return _BSCBridgeV2.Contract.BridgeExecuter(&_BSCBridgeV2.CallOpts)
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

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetBridgeExecuter(opts *bind.TransactOpts, _bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setBridgeExecuter", _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecuter(&_BSCBridgeV2.TransactOpts, _bridgeExecuter)
}

// SetBridgeExecuter is a paid mutator transaction binding the contract method 0x815f1de6.
//
// Solidity: function setBridgeExecuter(address _bridgeExecuter) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetBridgeExecuter(_bridgeExecuter common.Address) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetBridgeExecuter(&_BSCBridgeV2.TransactOpts, _bridgeExecuter)
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

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Transactor) SetTokenPause(opts *bind.TransactOpts, remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.contract.Transact(opts, "setTokenPause", remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2Session) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, pause)
}

// SetTokenPause is a paid mutator transaction binding the contract method 0x94ddc8c6.
//
// Solidity: function setTokenPause(uint256 remoteChainID, address token, bool pause) returns()
func (_BSCBridgeV2 *BSCBridgeV2TransactorSession) SetTokenPause(remoteChainID *big.Int, token common.Address, pause bool) (*types.Transaction, error) {
	return _BSCBridgeV2.Contract.SetTokenPause(&_BSCBridgeV2.TransactOpts, remoteChainID, token, pause)
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

// BSCBridgeV2BridgeExecuterCallFailedIterator is returned from FilterBridgeExecuterCallFailed and is used to iterate over the raw logs and unpacked data for BridgeExecuterCallFailed events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecuterCallFailedIterator struct {
	Event *BSCBridgeV2BridgeExecuterCallFailed // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeExecuterCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeExecuterCallFailed)
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
		it.Event = new(BSCBridgeV2BridgeExecuterCallFailed)
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
func (it *BSCBridgeV2BridgeExecuterCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeExecuterCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeExecuterCallFailed represents a BridgeExecuterCallFailed event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecuterCallFailed struct {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeExecuterCallFailed(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BSCBridgeV2BridgeExecuterCallFailedIterator, error) {

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

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeExecuterCallFailedIterator{contract: _BSCBridgeV2.contract, event: "BridgeExecuterCallFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterCallFailed is a free log subscription operation binding the contract event 0x9a7fcefef1f11328528152e408a4d5942f6e17ef00b9c8e556bf8ed6aa279be1.
//
// Solidity: event BridgeExecuterCallFailed(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, bytes reason)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeExecuterCallFailed(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeExecuterCallFailed, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeExecuterCallFailed", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeExecuterCallFailed)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeExecuterCallFailed(log types.Log) (*BSCBridgeV2BridgeExecuterCallFailed, error) {
	event := new(BSCBridgeV2BridgeExecuterCallFailed)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecuterCallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeV2BridgeExecuterSetIterator is returned from FilterBridgeExecuterSet and is used to iterate over the raw logs and unpacked data for BridgeExecuterSet events raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecuterSetIterator struct {
	Event *BSCBridgeV2BridgeExecuterSet // Event containing the contract specifics and raw log

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
func (it *BSCBridgeV2BridgeExecuterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeV2BridgeExecuterSet)
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
		it.Event = new(BSCBridgeV2BridgeExecuterSet)
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
func (it *BSCBridgeV2BridgeExecuterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeV2BridgeExecuterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeV2BridgeExecuterSet represents a BridgeExecuterSet event raised by the BSCBridgeV2 contract.
type BSCBridgeV2BridgeExecuterSet struct {
	BridgeExecuter common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuterSet is a free log retrieval operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterBridgeExecuterSet(opts *bind.FilterOpts, bridgeExecuter []common.Address) (*BSCBridgeV2BridgeExecuterSetIterator, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2BridgeExecuterSetIterator{contract: _BSCBridgeV2.contract, event: "BridgeExecuterSet", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuterSet is a free log subscription operation binding the contract event 0xf1a422c3bb184dcd08557083889de04eafe4ec097c96e03c58b4d52e9a729f6d.
//
// Solidity: event BridgeExecuterSet(address indexed bridgeExecuter)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchBridgeExecuterSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2BridgeExecuterSet, bridgeExecuter []common.Address) (event.Subscription, error) {

	var bridgeExecuterRule []interface{}
	for _, bridgeExecuterItem := range bridgeExecuter {
		bridgeExecuterRule = append(bridgeExecuterRule, bridgeExecuterItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "BridgeExecuterSet", bridgeExecuterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeV2BridgeExecuterSet)
				if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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
func (_BSCBridgeV2 *BSCBridgeV2Filterer) ParseBridgeExecuterSet(log types.Log) (*BSCBridgeV2BridgeExecuterSet, error) {
	event := new(BSCBridgeV2BridgeExecuterSet)
	if err := _BSCBridgeV2.contract.UnpackLog(event, "BridgeExecuterSet", log); err != nil {
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
	Code common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCrossMintableERC20CodeSet is a free log retrieval operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) FilterCrossMintableERC20CodeSet(opts *bind.FilterOpts, code []common.Address) (*BSCBridgeV2CrossMintableERC20CodeSetIterator, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.FilterLogs(opts, "CrossMintableERC20CodeSet", codeRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeV2CrossMintableERC20CodeSetIterator{contract: _BSCBridgeV2.contract, event: "CrossMintableERC20CodeSet", logs: logs, sub: sub}, nil
}

// WatchCrossMintableERC20CodeSet is a free log subscription operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
func (_BSCBridgeV2 *BSCBridgeV2Filterer) WatchCrossMintableERC20CodeSet(opts *bind.WatchOpts, sink chan<- *BSCBridgeV2CrossMintableERC20CodeSet, code []common.Address) (event.Subscription, error) {

	var codeRule []interface{}
	for _, codeItem := range code {
		codeRule = append(codeRule, codeItem)
	}

	logs, sub, err := _BSCBridgeV2.contract.WatchLogs(opts, "CrossMintableERC20CodeSet", codeRule)
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

// ParseCrossMintableERC20CodeSet is a log parse operation binding the contract event 0xedd032701363d0aea82c2dd885af304cc0627d71af78c87d1a2159ef3f9d2cee.
//
// Solidity: event CrossMintableERC20CodeSet(address indexed code)
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
	Pause         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseSet is a free log retrieval operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// WatchTokenPauseSet is a free log subscription operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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

// ParseTokenPauseSet is a log parse operation binding the contract event 0x493160def6e3ff3605efa5011d5fe909dfb92d4d4bc921682b8022e559dd7bea.
//
// Solidity: event TokenPauseSet(uint256 indexed remoteChainID, address indexed token, bool pause)
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
