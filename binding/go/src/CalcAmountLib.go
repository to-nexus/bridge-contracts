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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmountForNetworkFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcAmountLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"81d889b4": "calculateTokenAmountForNetworkFee(IPriceFeed,uint256,address,uint256)",
	},
	Bin: "0x610645610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610610034575f3560e01c806381d889b414610038575b5f5ffd5b61004b610046366004610421565b61006c565b60408051931515845260208401929092529082015260600160405180910390f35b5f5f5f5f876001600160a01b0316633afb1544886040518263ffffffff1660e01b815260040161009e91815260200190565b606060405180830381865afa1580156100b9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100dd9190610466565b9195509092509050836100f9575f5f5f9350935093505061019c565b60405163e588566f60e01b81526001600160a01b0387811660048301525f91908a169063e588566f90602401606060405180830381865afa158015610140573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101649190610466565b919650909350905084610181575f5f5f945094509450505061019c565b61019786838360126101928c6101a6565b610226565b935050505b9450945094915050565b5f6001600160a01b03821660011461021d57816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101f4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610218919061049e565b610220565b60125b92915050565b5f845f03610264576040516349671cd160e11b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff1610156102ac576102a78661028185876104d2565b61028c90600a6105ce565b61029690886105f0565b6102a091906105f0565b88906102fb565b6102cf565b6102cf6102a06102bc86866104d2565b6102c790600a6105ce565b879089610342565b92509050806102f15760405163c7ef3d4f60e01b815260040160405180910390fd5b5095945050505050565b5f5f835f0361030f5750600190505f61033b565b83830283858281610322576103226105dc565b0414610334575f5f925092505061033b565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036103765783828161036c5761036c6105dc565b04925050506103f2565b80841161038d5761038d60038515026011186103f9565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b634e487b715f52806020526024601cfd5b6001600160a01b038116811461041e575f5ffd5b50565b5f5f5f5f60808587031215610434575f5ffd5b843561043f8161040a565b93506020850135925060408501356104568161040a565b9396929550929360600135925050565b5f5f5f60608486031215610478575f5ffd5b83518015158114610487575f5ffd5b602085015160409095015190969495509392505050565b5f602082840312156104ae575f5ffd5b815160ff811681146103f2575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b60ff8281168282160390811115610220576102206104be565b6001815b60018411156105265780850481111561050a5761050a6104be565b600184161561051857908102905b60019390931c9280026104ef565b935093915050565b5f8261053c57506001610220565b8161054857505f610220565b816001811461055e576002811461056857610584565b6001915050610220565b60ff841115610579576105796104be565b50506001821b610220565b5060208310610133831016604e8410600b84101617156105a7575081810a610220565b6105b35f1984846104eb565b805f19048211156105c6576105c66104be565b029392505050565b5f6103f260ff84168361052e565b634e487b7160e01b5f52601260045260245ffd5b5f8261060a57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220b14efe3c0e12fa4644e76505a459ccc11936d568d122e79ad28438edb2f8b4dc64736f6c634300081c0033",
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
