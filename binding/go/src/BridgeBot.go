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

// BridgeBotBridgeConfig is an auto generated low-level Go binding around an user-defined struct.
type BridgeBotBridgeConfig struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}

// BridgeBotMetaData contains all meta data concerning the BridgeBot contract.
var BridgeBotMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBridgeConfig\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBaseBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canExecuteBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"canExecute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBridgeBatch\",\"inputs\":[{\"name\":\"configIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutableConfigs\",\"inputs\":[{\"name\":\"maxConfigs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"executableConfigs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextConfigId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toggleBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeConfigAdded\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigToggled\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigUpdated\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeExecuted\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BridgeBotBridgeFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotInsufficientBalance\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BridgeBotNotTimeYet\",\"inputs\":[{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// BridgeBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBotMetaData.ABI instead.
var BridgeBotABI = BridgeBotMetaData.ABI

// BridgeBot is an auto generated Go binding around an Ethereum contract.
type BridgeBot struct {
	BridgeBotCaller     // Read-only binding to the contract
	BridgeBotTransactor // Write-only binding to the contract
	BridgeBotFilterer   // Log filterer for contract events
}

// BridgeBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBotSession struct {
	Contract     *BridgeBot        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBotCallerSession struct {
	Contract *BridgeBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBotTransactorSession struct {
	Contract     *BridgeBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBotRaw struct {
	Contract *BridgeBot // Generic contract binding to access the raw methods on
}

// BridgeBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBotCallerRaw struct {
	Contract *BridgeBotCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBotTransactorRaw struct {
	Contract *BridgeBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBot creates a new instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBot(address common.Address, backend bind.ContractBackend) (*BridgeBot, error) {
	contract, err := bindBridgeBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBot{BridgeBotCaller: BridgeBotCaller{contract: contract}, BridgeBotTransactor: BridgeBotTransactor{contract: contract}, BridgeBotFilterer: BridgeBotFilterer{contract: contract}}, nil
}

// NewBridgeBotCaller creates a new read-only instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotCaller(address common.Address, caller bind.ContractCaller) (*BridgeBotCaller, error) {
	contract, err := bindBridgeBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBotCaller{contract: contract}, nil
}

// NewBridgeBotTransactor creates a new write-only instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBotTransactor, error) {
	contract, err := bindBridgeBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBotTransactor{contract: contract}, nil
}

// NewBridgeBotFilterer creates a new log filterer instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBotFilterer, error) {
	contract, err := bindBridgeBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBotFilterer{contract: contract}, nil
}

// bindBridgeBot binds a generic wrapper to an already deployed contract.
func bindBridgeBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBot *BridgeBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBot.Contract.BridgeBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBot *BridgeBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.Contract.BridgeBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBot *BridgeBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBot.Contract.BridgeBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBot *BridgeBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBot *BridgeBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBot *BridgeBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBot.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotSession) NATIVETOKEN() (common.Address, error) {
	return _BridgeBot.Contract.NATIVETOKEN(&_BridgeBot.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) NATIVETOKEN() (common.Address, error) {
	return _BridgeBot.Contract.NATIVETOKEN(&_BridgeBot.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotSession) Bridge() (common.Address, error) {
	return _BridgeBot.Contract.Bridge(&_BridgeBot.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) Bridge() (common.Address, error) {
	return _BridgeBot.Contract.Bridge(&_BridgeBot.CallOpts)
}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotCaller) BridgeConfigs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "bridgeConfigs", arg0)

	outstruct := new(struct {
		TokenAddress common.Address
		Recipient    common.Address
		ToChainID    *big.Int
		Interval     *big.Int
		LastExecuted *big.Int
		Enabled      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ToChainID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Interval = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastExecuted = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotSession) BridgeConfigs(arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	return _BridgeBot.Contract.BridgeConfigs(&_BridgeBot.CallOpts, arg0)
}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotCallerSession) BridgeConfigs(arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	return _BridgeBot.Contract.BridgeConfigs(&_BridgeBot.CallOpts, arg0)
}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotCaller) CanExecuteBridge(opts *bind.CallOpts, configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "canExecuteBridge", configId)

	outstruct := new(struct {
		CanExecute        bool
		NextAvailableTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanExecute = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NextAvailableTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotSession) CanExecuteBridge(configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	return _BridgeBot.Contract.CanExecuteBridge(&_BridgeBot.CallOpts, configId)
}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotCallerSession) CanExecuteBridge(configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	return _BridgeBot.Contract.CanExecuteBridge(&_BridgeBot.CallOpts, configId)
}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotCaller) GetBridgeConfig(opts *bind.CallOpts, configId *big.Int) (BridgeBotBridgeConfig, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "getBridgeConfig", configId)

	if err != nil {
		return *new(BridgeBotBridgeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeBotBridgeConfig)).(*BridgeBotBridgeConfig)

	return out0, err

}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotSession) GetBridgeConfig(configId *big.Int) (BridgeBotBridgeConfig, error) {
	return _BridgeBot.Contract.GetBridgeConfig(&_BridgeBot.CallOpts, configId)
}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotCallerSession) GetBridgeConfig(configId *big.Int) (BridgeBotBridgeConfig, error) {
	return _BridgeBot.Contract.GetBridgeConfig(&_BridgeBot.CallOpts, configId)
}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotCaller) GetExecutableConfigs(opts *bind.CallOpts, maxConfigs *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "getExecutableConfigs", maxConfigs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotSession) GetExecutableConfigs(maxConfigs *big.Int) ([]*big.Int, error) {
	return _BridgeBot.Contract.GetExecutableConfigs(&_BridgeBot.CallOpts, maxConfigs)
}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotCallerSession) GetExecutableConfigs(maxConfigs *big.Int) ([]*big.Int, error) {
	return _BridgeBot.Contract.GetExecutableConfigs(&_BridgeBot.CallOpts, maxConfigs)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotCaller) NextConfigId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "nextConfigId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotSession) NextConfigId() (*big.Int, error) {
	return _BridgeBot.Contract.NextConfigId(&_BridgeBot.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotCallerSession) NextConfigId() (*big.Int, error) {
	return _BridgeBot.Contract.NextConfigId(&_BridgeBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotSession) Owner() (common.Address, error) {
	return _BridgeBot.Contract.Owner(&_BridgeBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) Owner() (common.Address, error) {
	return _BridgeBot.Contract.Owner(&_BridgeBot.CallOpts)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotTransactor) AddBridgeConfig(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "addBridgeConfig", tokenAddress, recipient, toChainID, interval)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotSession) AddBridgeConfig(tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.AddBridgeConfig(&_BridgeBot.TransactOpts, tokenAddress, recipient, toChainID, interval)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotTransactorSession) AddBridgeConfig(tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.AddBridgeConfig(&_BridgeBot.TransactOpts, tokenAddress, recipient, toChainID, interval)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) ExecuteBridge(opts *bind.TransactOpts, configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "executeBridge", configId, amount)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) ExecuteBridge(configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridge(&_BridgeBot.TransactOpts, configId, amount)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) ExecuteBridge(configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridge(&_BridgeBot.TransactOpts, configId, amount)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotTransactor) ExecuteBridgeBatch(opts *bind.TransactOpts, configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "executeBridgeBatch", configIds, amounts)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotSession) ExecuteBridgeBatch(configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridgeBatch(&_BridgeBot.TransactOpts, configIds, amounts)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotTransactorSession) ExecuteBridgeBatch(configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridgeBatch(&_BridgeBot.TransactOpts, configIds, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceOwnership(&_BridgeBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceOwnership(&_BridgeBot.TransactOpts)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotTransactor) ToggleBridgeConfig(opts *bind.TransactOpts, configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "toggleBridgeConfig", configId, enabled)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotSession) ToggleBridgeConfig(configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.Contract.ToggleBridgeConfig(&_BridgeBot.TransactOpts, configId, enabled)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotTransactorSession) ToggleBridgeConfig(configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.Contract.ToggleBridgeConfig(&_BridgeBot.TransactOpts, configId, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.TransferOwnership(&_BridgeBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.TransferOwnership(&_BridgeBot.TransactOpts, newOwner)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotTransactor) UpdateBridgeConfig(opts *bind.TransactOpts, configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "updateBridgeConfig", configId, tokenAddress, recipient, toChainID, interval)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotSession) UpdateBridgeConfig(configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.UpdateBridgeConfig(&_BridgeBot.TransactOpts, configId, tokenAddress, recipient, toChainID, interval)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotTransactorSession) UpdateBridgeConfig(configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.UpdateBridgeConfig(&_BridgeBot.TransactOpts, configId, tokenAddress, recipient, toChainID, interval)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawAllNative(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawAllNative", to)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotSession) WithdrawAllNative(to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllNative(&_BridgeBot.TransactOpts, to)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawAllNative(to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllNative(&_BridgeBot.TransactOpts, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawAllTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawAllTokens", token, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotSession) WithdrawAllTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllTokens(&_BridgeBot.TransactOpts, token, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawAllTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllTokens(&_BridgeBot.TransactOpts, token, to)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawNative", to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawNative(&_BridgeBot.TransactOpts, to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawNative(&_BridgeBot.TransactOpts, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawToken", token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawToken(&_BridgeBot.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawToken(&_BridgeBot.TransactOpts, token, to, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.Contract.Fallback(&_BridgeBot.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.Contract.Fallback(&_BridgeBot.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotSession) Receive() (*types.Transaction, error) {
	return _BridgeBot.Contract.Receive(&_BridgeBot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBot.Contract.Receive(&_BridgeBot.TransactOpts)
}

// BridgeBotBridgeConfigAddedIterator is returned from FilterBridgeConfigAdded and is used to iterate over the raw logs and unpacked data for BridgeConfigAdded events raised by the BridgeBot contract.
type BridgeBotBridgeConfigAddedIterator struct {
	Event *BridgeBotBridgeConfigAdded // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigAdded)
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
		it.Event = new(BridgeBotBridgeConfigAdded)
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
func (it *BridgeBotBridgeConfigAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigAdded represents a BridgeConfigAdded event raised by the BridgeBot contract.
type BridgeBotBridgeConfigAdded struct {
	ConfigId *big.Int
	Config   BridgeBotBridgeConfig
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigAdded is a free log retrieval operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigAdded(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigAddedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigAdded", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigAddedIterator{contract: _BridgeBot.contract, event: "BridgeConfigAdded", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigAdded is a free log subscription operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigAdded(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigAdded, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigAdded", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigAdded)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigAdded", log); err != nil {
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

// ParseBridgeConfigAdded is a log parse operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigAdded(log types.Log) (*BridgeBotBridgeConfigAdded, error) {
	event := new(BridgeBotBridgeConfigAdded)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeConfigToggledIterator is returned from FilterBridgeConfigToggled and is used to iterate over the raw logs and unpacked data for BridgeConfigToggled events raised by the BridgeBot contract.
type BridgeBotBridgeConfigToggledIterator struct {
	Event *BridgeBotBridgeConfigToggled // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigToggled)
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
		it.Event = new(BridgeBotBridgeConfigToggled)
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
func (it *BridgeBotBridgeConfigToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigToggled represents a BridgeConfigToggled event raised by the BridgeBot contract.
type BridgeBotBridgeConfigToggled struct {
	ConfigId *big.Int
	Enabled  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigToggled is a free log retrieval operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigToggled(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigToggledIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigToggled", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigToggledIterator{contract: _BridgeBot.contract, event: "BridgeConfigToggled", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigToggled is a free log subscription operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigToggled(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigToggled, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigToggled", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigToggled)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigToggled", log); err != nil {
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

// ParseBridgeConfigToggled is a log parse operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigToggled(log types.Log) (*BridgeBotBridgeConfigToggled, error) {
	event := new(BridgeBotBridgeConfigToggled)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeConfigUpdatedIterator is returned from FilterBridgeConfigUpdated and is used to iterate over the raw logs and unpacked data for BridgeConfigUpdated events raised by the BridgeBot contract.
type BridgeBotBridgeConfigUpdatedIterator struct {
	Event *BridgeBotBridgeConfigUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigUpdated)
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
		it.Event = new(BridgeBotBridgeConfigUpdated)
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
func (it *BridgeBotBridgeConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigUpdated represents a BridgeConfigUpdated event raised by the BridgeBot contract.
type BridgeBotBridgeConfigUpdated struct {
	ConfigId *big.Int
	Config   BridgeBotBridgeConfig
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigUpdated is a free log retrieval operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigUpdated(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigUpdatedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigUpdated", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigUpdatedIterator{contract: _BridgeBot.contract, event: "BridgeConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigUpdated is a free log subscription operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigUpdated, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigUpdated", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigUpdated)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigUpdated", log); err != nil {
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

// ParseBridgeConfigUpdated is a log parse operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigUpdated(log types.Log) (*BridgeBotBridgeConfigUpdated, error) {
	event := new(BridgeBotBridgeConfigUpdated)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeExecutedIterator is returned from FilterBridgeExecuted and is used to iterate over the raw logs and unpacked data for BridgeExecuted events raised by the BridgeBot contract.
type BridgeBotBridgeExecutedIterator struct {
	Event *BridgeBotBridgeExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeExecuted)
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
		it.Event = new(BridgeBotBridgeExecuted)
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
func (it *BridgeBotBridgeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeExecuted represents a BridgeExecuted event raised by the BridgeBot contract.
type BridgeBotBridgeExecuted struct {
	ConfigId     *big.Int
	TokenAddress common.Address
	Amount       *big.Int
	Recipient    common.Address
	ToChainID    *big.Int
	Executor     common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuted is a free log retrieval operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeExecuted(opts *bind.FilterOpts, configId []*big.Int, tokenAddress []common.Address, recipient []common.Address) (*BridgeBotBridgeExecutedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeExecuted", configIdRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeExecutedIterator{contract: _BridgeBot.contract, event: "BridgeExecuted", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuted is a free log subscription operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeExecuted(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeExecuted, configId []*big.Int, tokenAddress []common.Address, recipient []common.Address) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeExecuted", configIdRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeExecuted)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeExecuted", log); err != nil {
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

// ParseBridgeExecuted is a log parse operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeExecuted(log types.Log) (*BridgeBotBridgeExecuted, error) {
	event := new(BridgeBotBridgeExecuted)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotNativeWithdrawnIterator is returned from FilterNativeWithdrawn and is used to iterate over the raw logs and unpacked data for NativeWithdrawn events raised by the BridgeBot contract.
type BridgeBotNativeWithdrawnIterator struct {
	Event *BridgeBotNativeWithdrawn // Event containing the contract specifics and raw log

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
func (it *BridgeBotNativeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotNativeWithdrawn)
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
		it.Event = new(BridgeBotNativeWithdrawn)
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
func (it *BridgeBotNativeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotNativeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotNativeWithdrawn represents a NativeWithdrawn event raised by the BridgeBot contract.
type BridgeBotNativeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeWithdrawn is a free log retrieval operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) FilterNativeWithdrawn(opts *bind.FilterOpts, to []common.Address) (*BridgeBotNativeWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "NativeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotNativeWithdrawnIterator{contract: _BridgeBot.contract, event: "NativeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNativeWithdrawn is a free log subscription operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) WatchNativeWithdrawn(opts *bind.WatchOpts, sink chan<- *BridgeBotNativeWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "NativeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotNativeWithdrawn)
				if err := _BridgeBot.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
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

// ParseNativeWithdrawn is a log parse operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) ParseNativeWithdrawn(log types.Log) (*BridgeBotNativeWithdrawn, error) {
	event := new(BridgeBotNativeWithdrawn)
	if err := _BridgeBot.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeBot contract.
type BridgeBotOwnershipTransferredIterator struct {
	Event *BridgeBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotOwnershipTransferred)
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
		it.Event = new(BridgeBotOwnershipTransferred)
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
func (it *BridgeBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeBot contract.
type BridgeBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotOwnershipTransferredIterator{contract: _BridgeBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotOwnershipTransferred)
				if err := _BridgeBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeBotOwnershipTransferred, error) {
	event := new(BridgeBotOwnershipTransferred)
	if err := _BridgeBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the BridgeBot contract.
type BridgeBotTokenWithdrawnIterator struct {
	Event *BridgeBotTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *BridgeBotTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotTokenWithdrawn)
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
		it.Event = new(BridgeBotTokenWithdrawn)
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
func (it *BridgeBotTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotTokenWithdrawn represents a TokenWithdrawn event raised by the BridgeBot contract.
type BridgeBotTokenWithdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeBotTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotTokenWithdrawnIterator{contract: _BridgeBot.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BridgeBotTokenWithdrawn, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotTokenWithdrawn)
				if err := _BridgeBot.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) ParseTokenWithdrawn(log types.Log) (*BridgeBotTokenWithdrawn, error) {
	event := new(BridgeBotTokenWithdrawn)
	if err := _BridgeBot.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
