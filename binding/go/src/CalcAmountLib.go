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

// CalcAmountLibMetaData contains all meta data concerning the CalcAmountLib contract.
var CalcAmountLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmountForNetworkFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcAmountLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
		"81d889b4": "calculateTokenAmountForNetworkFee(IPriceFeed,uint256,address,uint256)",
		"d449a832": "decimals(address)",
	},
	Bin: "0x610c1c61004d600b8282823980515f1a6073146041577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c806381d889b41461004e578063b1db08d114610080578063d449a832146100b0575b5f5ffd5b6100686004803603810190610063919061063f565b6100e0565b604051610077939291906106cc565b60405180910390f35b61009a60048036038101906100959190610737565b61023b565b6040516100a791906107ae565b60405180910390f35b6100ca60048036038101906100c591906107c7565b61035a565b6040516100d79190610801565b60405180910390f35b5f5f5f5f8773ffffffffffffffffffffffffffffffffffffffff16633afb1544886040518263ffffffff1660e01b815260040161011d9190610829565b606060405180830381865afa158015610138573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061015c9190610880565b8094508193508296505050508361017c575f5f5f93509350935050610231565b5f8873ffffffffffffffffffffffffffffffffffffffff1663e588566f886040518263ffffffff1660e01b81526004016101b691906108df565b606060405180830381865afa1580156101d1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101f59190610880565b80955081935082975050505084610216575f5f5f9450945094505050610231565b61022c86838360126102278c61035a565b61023b565b935050505b9450945094915050565b5f5f850361027e576040517f92ce39a200000000000000000000000000000000000000000000000000000000815260040161027590610952565b60405180910390fd5b5f8360ff168360ff1610156102d2576102cd86848661029d919061099d565b600a6102a99190610b00565b876102b49190610b77565b6102be9190610b77565b8861040b90919063ffffffff16565b610311565b61031061030185856102e4919061099d565b600a6102f09190610b00565b88886104599092919063ffffffff16565b8861040b90919063ffffffff16565b5b809350819250505080610350576040517fc7ef3d4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5095945050505050565b5f600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610401578173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103d8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103fc9190610bbb565b610404565b60125b9050919050565b5f5f5f84036104205760015f91509150610452565b5f83850290508385828161043757610436610b4a565b5b0414610449575f5f9250925050610452565b60018192509250505b9250929050565b5f5f83850290505f5f19858709828110838203039150505f81036104915783828161048757610486610b4a565b5b0492505050610537565b8084116104b0576104af6104aa5f86146012601161053e565b610557565b5b5f8486880990508281118203915080830392505f855f038616905080860495508084049350600181825f0304019050808302841793505f600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b5f61054884610568565b82841802821890509392505050565b634e487b715f52806020526024601cfd5b5f8115159050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105a082610577565b9050919050565b5f6105b182610596565b9050919050565b6105c1816105a7565b81146105cb575f5ffd5b50565b5f813590506105dc816105b8565b92915050565b5f819050919050565b6105f4816105e2565b81146105fe575f5ffd5b50565b5f8135905061060f816105eb565b92915050565b61061e81610596565b8114610628575f5ffd5b50565b5f8135905061063981610615565b92915050565b5f5f5f5f6080858703121561065757610656610573565b5b5f610664878288016105ce565b945050602061067587828801610601565b93505060406106868782880161062b565b925050606061069787828801610601565b91505092959194509250565b5f8115159050919050565b6106b7816106a3565b82525050565b6106c6816105e2565b82525050565b5f6060820190506106df5f8301866106ae565b6106ec60208301856106bd565b6106f960408301846106bd565b949350505050565b5f60ff82169050919050565b61071681610701565b8114610720575f5ffd5b50565b5f813590506107318161070d565b92915050565b5f5f5f5f5f60a086880312156107505761074f610573565b5b5f61075d88828901610601565b955050602061076e88828901610601565b945050604061077f88828901610601565b935050606061079088828901610723565b92505060806107a188828901610723565b9150509295509295909350565b5f6020820190506107c15f8301846106bd565b92915050565b5f602082840312156107dc576107db610573565b5b5f6107e98482850161062b565b91505092915050565b6107fb81610701565b82525050565b5f6020820190506108145f8301846107f2565b92915050565b610823816105e2565b82525050565b5f60208201905061083c5f83018461081a565b92915050565b61084b816106a3565b8114610855575f5ffd5b50565b5f8151905061086681610842565b92915050565b5f8151905061087a816105eb565b92915050565b5f5f5f6060848603121561089757610896610573565b5b5f6108a486828701610858565b93505060206108b58682870161086c565b92505060406108c68682870161086c565b9150509250925092565b6108d981610596565b82525050565b5f6020820190506108f25f8301846108d0565b92915050565b5f82825260208201905092915050565b7f70726963654100000000000000000000000000000000000000000000000000005f82015250565b5f61093c6006836108f8565b915061094782610908565b602082019050919050565b5f6020820190508181035f83015261096981610930565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6109a782610701565b91506109b283610701565b9250828203905060ff8111156109cb576109ca610970565b5b92915050565b5f8160011c9050919050565b5f5f8291508390505b6001851115610a2657808604811115610a0257610a01610970565b5b6001851615610a115780820291505b8081029050610a1f856109d1565b94506109e6565b94509492505050565b5f82610a3e5760019050610af9565b81610a4b575f9050610af9565b8160018114610a615760028114610a6b57610a9a565b6001915050610af9565b60ff841115610a7d57610a7c610970565b5b8360020a915084821115610a9457610a93610970565b5b50610af9565b5060208310610133831016604e8410600b8410161715610acf5782820a905083811115610aca57610ac9610970565b5b610af9565b610adc84848460016109dd565b92509050818404811115610af357610af2610970565b5b81810290505b9392505050565b5f610b0a826105e2565b9150610b1583610701565b9250610b427fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610a2f565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610b81826105e2565b9150610b8c836105e2565b925082610b9c57610b9b610b4a565b5b828204905092915050565b5f81519050610bb58161070d565b92915050565b5f60208284031215610bd057610bcf610573565b5b5f610bdd84828501610ba7565b9150509291505056fea2646970667358221220f5fe0f7140be42f501e7362a547645bcbe88fa97dffe7c4c833086fc45d5ffb464736f6c634300081c0033",
}

// CalcAmountLibABI is the input ABI used to generate the binding from.
// Deprecated: Use CalcAmountLibMetaData.ABI instead.
var CalcAmountLibABI = CalcAmountLibMetaData.ABI

// Deprecated: Use CalcAmountLibMetaData.Sigs instead.
// CalcAmountLibFuncSigs maps the 4-byte function signature to its string representation.
var CalcAmountLibFuncSigs = CalcAmountLibMetaData.Sigs

// CalcAmountLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CalcAmountLibMetaData.Bin instead.
var CalcAmountLibBin = CalcAmountLibMetaData.Bin

// DeployCalcAmountLib deploys a new Ethereum contract, binding an instance of CalcAmountLib to it.
func DeployCalcAmountLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CalcAmountLib, error) {
	parsed, err := CalcAmountLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CalcAmountLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CalcAmountLib{CalcAmountLibCaller: CalcAmountLibCaller{contract: contract}, CalcAmountLibTransactor: CalcAmountLibTransactor{contract: contract}, CalcAmountLibFilterer: CalcAmountLibFilterer{contract: contract}}, nil
}

// CalcAmountLib is an auto generated Go binding around an Ethereum contract.
type CalcAmountLib struct {
	CalcAmountLibCaller     // Read-only binding to the contract
	CalcAmountLibTransactor // Write-only binding to the contract
	CalcAmountLibFilterer   // Log filterer for contract events
}

// CalcAmountLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalcAmountLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcAmountLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalcAmountLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcAmountLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalcAmountLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcAmountLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalcAmountLibSession struct {
	Contract     *CalcAmountLib    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalcAmountLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalcAmountLibCallerSession struct {
	Contract *CalcAmountLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CalcAmountLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalcAmountLibTransactorSession struct {
	Contract     *CalcAmountLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CalcAmountLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalcAmountLibRaw struct {
	Contract *CalcAmountLib // Generic contract binding to access the raw methods on
}

// CalcAmountLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalcAmountLibCallerRaw struct {
	Contract *CalcAmountLibCaller // Generic read-only contract binding to access the raw methods on
}

// CalcAmountLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalcAmountLibTransactorRaw struct {
	Contract *CalcAmountLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalcAmountLib creates a new instance of CalcAmountLib, bound to a specific deployed contract.
func NewCalcAmountLib(address common.Address, backend bind.ContractBackend) (*CalcAmountLib, error) {
	contract, err := bindCalcAmountLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CalcAmountLib{CalcAmountLibCaller: CalcAmountLibCaller{contract: contract}, CalcAmountLibTransactor: CalcAmountLibTransactor{contract: contract}, CalcAmountLibFilterer: CalcAmountLibFilterer{contract: contract}}, nil
}

// NewCalcAmountLibCaller creates a new read-only instance of CalcAmountLib, bound to a specific deployed contract.
func NewCalcAmountLibCaller(address common.Address, caller bind.ContractCaller) (*CalcAmountLibCaller, error) {
	contract, err := bindCalcAmountLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalcAmountLibCaller{contract: contract}, nil
}

// NewCalcAmountLibTransactor creates a new write-only instance of CalcAmountLib, bound to a specific deployed contract.
func NewCalcAmountLibTransactor(address common.Address, transactor bind.ContractTransactor) (*CalcAmountLibTransactor, error) {
	contract, err := bindCalcAmountLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalcAmountLibTransactor{contract: contract}, nil
}

// NewCalcAmountLibFilterer creates a new log filterer instance of CalcAmountLib, bound to a specific deployed contract.
func NewCalcAmountLibFilterer(address common.Address, filterer bind.ContractFilterer) (*CalcAmountLibFilterer, error) {
	contract, err := bindCalcAmountLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalcAmountLibFilterer{contract: contract}, nil
}

// bindCalcAmountLib binds a generic wrapper to an already deployed contract.
func bindCalcAmountLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CalcAmountLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalcAmountLib *CalcAmountLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalcAmountLib.Contract.CalcAmountLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalcAmountLib *CalcAmountLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalcAmountLib.Contract.CalcAmountLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalcAmountLib *CalcAmountLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalcAmountLib.Contract.CalcAmountLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalcAmountLib *CalcAmountLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalcAmountLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalcAmountLib *CalcAmountLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalcAmountLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalcAmountLib *CalcAmountLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalcAmountLib.Contract.contract.Transact(opts, method, params...)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcAmountLib *CalcAmountLibCaller) CalculateAmountBWithPrice(opts *bind.CallOpts, amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	var out []interface{}
	err := _CalcAmountLib.contract.Call(opts, &out, "calculateAmountBWithPrice", amountA, priceA, priceB, decimalA, decimalB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcAmountLib *CalcAmountLibSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _CalcAmountLib.Contract.CalculateAmountBWithPrice(&_CalcAmountLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcAmountLib *CalcAmountLibCallerSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _CalcAmountLib.Contract.CalculateAmountBWithPrice(&_CalcAmountLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// CalculateTokenAmountForNetworkFee is a free data retrieval call binding the contract method 0x81d889b4.
//
// Solidity: function calculateTokenAmountForNetworkFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcAmountLib *CalcAmountLibCaller) CalculateTokenAmountForNetworkFee(opts *bind.CallOpts, feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	var out []interface{}
	err := _CalcAmountLib.contract.Call(opts, &out, "calculateTokenAmountForNetworkFee", feed, toChainID, token, nativeTokenAmount)

	outstruct := new(struct {
		Ok          bool
		TokenAmount *big.Int
		UpdatedAt   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ok = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.TokenAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateTokenAmountForNetworkFee is a free data retrieval call binding the contract method 0x81d889b4.
//
// Solidity: function calculateTokenAmountForNetworkFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcAmountLib *CalcAmountLibSession) CalculateTokenAmountForNetworkFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	return _CalcAmountLib.Contract.CalculateTokenAmountForNetworkFee(&_CalcAmountLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// CalculateTokenAmountForNetworkFee is a free data retrieval call binding the contract method 0x81d889b4.
//
// Solidity: function calculateTokenAmountForNetworkFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcAmountLib *CalcAmountLibCallerSession) CalculateTokenAmountForNetworkFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	return _CalcAmountLib.Contract.CalculateTokenAmountForNetworkFee(&_CalcAmountLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcAmountLib *CalcAmountLibCaller) Decimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _CalcAmountLib.contract.Call(opts, &out, "decimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcAmountLib *CalcAmountLibSession) Decimals(token common.Address) (uint8, error) {
	return _CalcAmountLib.Contract.Decimals(&_CalcAmountLib.CallOpts, token)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcAmountLib *CalcAmountLibCallerSession) Decimals(token common.Address) (uint8, error) {
	return _CalcAmountLib.Contract.Decimals(&_CalcAmountLib.CallOpts, token)
}
