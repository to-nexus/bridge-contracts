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

// PriceFeedLibMetaData contains all meta data concerning the PriceFeedLib contract.
var PriceFeedLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"}],\"name\":\"allPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"prices\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"convertAtoB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"convertAtoBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"999d5cf4": "allPrices(IPriceFeed)",
		"126e5b3a": "convertAtoB(IPriceFeed,address,address,uint256)",
		"957bb22e": "convertAtoBWithPrice(uint256,uint256,uint256,uint8,uint8)",
	},
	Bin: "0x610b0c610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c8063126e5b3a1461004e578063957bb22e1461007b578063999d5cf41461009c575b5f5ffd5b61006161005c3660046105c0565b6100bc565b604080519283529015156020830152015b60405180910390f35b61008e61008936600461061c565b610208565b604051908152602001610072565b6100af6100aa36600461066e565b6102dd565b6040516100729190610689565b5f5f836001600160a01b0316856001600160a01b0316036100e2575081905060016101ff565b6040516333ddef0760e01b81526001600160a01b0386811660048301525f9182918916906333ddef07906024016040805180830381865afa158015610129573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061014d91906106ed565b6040516333ddef0760e01b81526001600160a01b0389811660048301529294509092505f918291908b16906333ddef07906024016040805180830381865afa15801561019b573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101bf91906106ed565b915091508280156101cd5750805b94505f5f6101db8c8c6103b3565b6101e58d8c6103b3565b915091506101f68987868585610208565b97505050505050505b94509492505050565b5f845f0361024657604051633a0b0f2960e01b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff16101561028e57610289866102638587610734565b61026e90600a610830565b6102789088610852565b6102829190610852565b889061049b565b6102b1565b6102b161028261029e8686610734565b6102a990600a610830565b8790896104e2565b92509050806102d357604051633961e4cf60e11b815260040160405180910390fd5b5095945050505050565b6060816001600160a01b0316638fb5a482836001600160a01b0316636ff97f1d6040518163ffffffff1660e01b81526004015f60405180830381865afa158015610329573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103509190810190610902565b6040518263ffffffff1660e01b815260040161036c91906109a1565b5f60405180830381865afa158015610386573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103ad91908101906109e1565b92915050565b5f826001600160a01b03166311df99956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103f0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104149190610aa0565b6001600160a01b0316826001600160a01b03161461049157816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610468573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061048c9190610abb565b610494565b60125b9392505050565b5f5f835f036104af5750600190505f6104db565b838302838582816104c2576104c261083e565b04146104d4575f5f92509250506104db565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036105165783828161050c5761050c61083e565b0492505050610494565b80841161052d5761052d6003851502601118610598565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6001600160a01b03811681146105bd575f5ffd5b50565b5f5f5f5f608085870312156105d3575f5ffd5b84356105de816105a9565b935060208501356105ee816105a9565b925060408501356105fe816105a9565b9396929550929360600135925050565b60ff811681146105bd575f5ffd5b5f5f5f5f5f60a08688031215610630575f5ffd5b85359450602086013593506040860135925060608601356106508161060e565b915060808601356106608161060e565b809150509295509295909350565b5f6020828403121561067e575f5ffd5b8135610494816105a9565b602080825282518282018190525f918401906040840190835b818110156106e257835180516001600160a01b031684526020808201518186015260409182015191850191909152909301926060909201916001016106a2565b509095945050505050565b5f5f604083850312156106fe575f5ffd5b825160208401519092508015158114610715575f5ffd5b809150509250929050565b634e487b7160e01b5f52601160045260245ffd5b60ff82811682821603908111156103ad576103ad610720565b6001815b60018411156107885780850481111561076c5761076c610720565b600184161561077a57908102905b60019390931c928002610751565b935093915050565b5f8261079e575060016103ad565b816107aa57505f6103ad565b81600181146107c057600281146107ca576107e6565b60019150506103ad565b60ff8411156107db576107db610720565b50506001821b6103ad565b5060208310610133831016604e8410600b8410161715610809575081810a6103ad565b6108155f19848461074d565b805f190482111561082857610828610720565b029392505050565b5f61049460ff841683610790565b634e487b7160e01b5f52601260045260245ffd5b5f8261086c57634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156108a8576108a8610871565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108d7576108d7610871565b604052919050565b5f67ffffffffffffffff8211156108f8576108f8610871565b5060051b60200190565b5f60208284031215610912575f5ffd5b815167ffffffffffffffff811115610928575f5ffd5b8201601f81018413610938575f5ffd5b805161094b610946826108df565b6108ae565b8082825260208201915060208360051b85010192508683111561096c575f5ffd5b6020840193505b82841015610997578351610986816105a9565b825260209384019390910190610973565b9695505050505050565b602080825282518282018190525f918401906040840190835b818110156106e25783516001600160a01b03168352602093840193909201916001016109ba565b5f602082840312156109f1575f5ffd5b815167ffffffffffffffff811115610a07575f5ffd5b8201601f81018413610a17575f5ffd5b8051610a25610946826108df565b80828252602082019150602060608402850101925086831115610a46575f5ffd5b6020840193505b828410156109975760608488031215610a64575f5ffd5b610a6c610885565b8451610a77816105a9565b815260208581015181830152604080870151908301529083526060909401939190910190610a4d565b5f60208284031215610ab0575f5ffd5b8151610494816105a9565b5f60208284031215610acb575f5ffd5b81516104948161060e56fea2646970667358221220c1123a3bad0bb6b185dc4e67bb49eecbe3b820e6b48071082e94bb9ee2fcc27964736f6c634300081c0033",
}

// PriceFeedLibABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedLibMetaData.ABI instead.
var PriceFeedLibABI = PriceFeedLibMetaData.ABI

// Deprecated: Use PriceFeedLibMetaData.Sigs instead.
// PriceFeedLibFuncSigs maps the 4-byte function signature to its string representation.
var PriceFeedLibFuncSigs = PriceFeedLibMetaData.Sigs

// PriceFeedLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedLibMetaData.Bin instead.
var PriceFeedLibBin = PriceFeedLibMetaData.Bin

// DeployPriceFeedLib deploys a new Ethereum contract, binding an instance of PriceFeedLib to it.
func DeployPriceFeedLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeedLib, error) {
	parsed, err := PriceFeedLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeedLib{PriceFeedLibCaller: PriceFeedLibCaller{contract: contract}, PriceFeedLibTransactor: PriceFeedLibTransactor{contract: contract}, PriceFeedLibFilterer: PriceFeedLibFilterer{contract: contract}}, nil
}

// PriceFeedLib is an auto generated Go binding around an Ethereum contract.
type PriceFeedLib struct {
	PriceFeedLibCaller     // Read-only binding to the contract
	PriceFeedLibTransactor // Write-only binding to the contract
	PriceFeedLibFilterer   // Log filterer for contract events
}

// PriceFeedLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedLibSession struct {
	Contract     *PriceFeedLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedLibCallerSession struct {
	Contract *PriceFeedLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PriceFeedLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedLibTransactorSession struct {
	Contract     *PriceFeedLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PriceFeedLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedLibRaw struct {
	Contract *PriceFeedLib // Generic contract binding to access the raw methods on
}

// PriceFeedLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedLibCallerRaw struct {
	Contract *PriceFeedLibCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedLibTransactorRaw struct {
	Contract *PriceFeedLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedLib creates a new instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLib(address common.Address, backend bind.ContractBackend) (*PriceFeedLib, error) {
	contract, err := bindPriceFeedLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLib{PriceFeedLibCaller: PriceFeedLibCaller{contract: contract}, PriceFeedLibTransactor: PriceFeedLibTransactor{contract: contract}, PriceFeedLibFilterer: PriceFeedLibFilterer{contract: contract}}, nil
}

// NewPriceFeedLibCaller creates a new read-only instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedLibCaller, error) {
	contract, err := bindPriceFeedLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibCaller{contract: contract}, nil
}

// NewPriceFeedLibTransactor creates a new write-only instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedLibTransactor, error) {
	contract, err := bindPriceFeedLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibTransactor{contract: contract}, nil
}

// NewPriceFeedLibFilterer creates a new log filterer instance of PriceFeedLib, bound to a specific deployed contract.
func NewPriceFeedLibFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedLibFilterer, error) {
	contract, err := bindPriceFeedLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedLibFilterer{contract: contract}, nil
}

// bindPriceFeedLib binds a generic wrapper to an already deployed contract.
func bindPriceFeedLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedLib *PriceFeedLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedLib.Contract.PriceFeedLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedLib *PriceFeedLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.PriceFeedLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedLib *PriceFeedLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.PriceFeedLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedLib *PriceFeedLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedLib *PriceFeedLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedLib *PriceFeedLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedLib.Contract.contract.Transact(opts, method, params...)
}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibCaller) AllPrices(opts *bind.CallOpts, feed common.Address) ([]IPriceFeedPriceData, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "allPrices", feed)

	if err != nil {
		return *new([]IPriceFeedPriceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPriceFeedPriceData)).(*[]IPriceFeedPriceData)

	return out0, err

}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibSession) AllPrices(feed common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedLib.Contract.AllPrices(&_PriceFeedLib.CallOpts, feed)
}

// AllPrices is a free data retrieval call binding the contract method 0x999d5cf4.
//
// Solidity: function allPrices(IPriceFeed feed) view returns((address,uint256,uint256)[] prices)
func (_PriceFeedLib *PriceFeedLibCallerSession) AllPrices(feed common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedLib.Contract.AllPrices(&_PriceFeedLib.CallOpts, feed)
}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibCaller) ConvertAtoB(opts *bind.CallOpts, feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "convertAtoB", feed, tokenA, tokenB, amountA)

	outstruct := new(struct {
		AmountB *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountB = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibSession) ConvertAtoB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	return _PriceFeedLib.Contract.ConvertAtoB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// ConvertAtoB is a free data retrieval call binding the contract method 0x126e5b3a.
//
// Solidity: function convertAtoB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(uint256 amountB, bool isValid)
func (_PriceFeedLib *PriceFeedLibCallerSession) ConvertAtoB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	AmountB *big.Int
	IsValid bool
}, error) {
	return _PriceFeedLib.Contract.ConvertAtoB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCaller) ConvertAtoBWithPrice(opts *bind.CallOpts, amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "convertAtoBWithPrice", amountA, priceA, priceB, decimalA, decimalB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibSession) ConvertAtoBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.ConvertAtoBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// ConvertAtoBWithPrice is a free data retrieval call binding the contract method 0x957bb22e.
//
// Solidity: function convertAtoBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCallerSession) ConvertAtoBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.ConvertAtoBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}
