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

// CalcGasFeeLibMetaData contains all meta data concerning the CalcGasFeeLib contract.
var CalcGasFeeLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmountForGasFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcGasFeeLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcGasFeeLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
		"889ad9e0": "calculateTokenAmountForGasFee(IPriceFeed,uint256,address,uint256)",
		"0fe74f3b": "decimals(IPriceFeed,address)",
	},
	Bin: "0x6107b8610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c80630fe74f3b1461004e578063889ad9e014610078578063b1db08d1146100a8575b5f5ffd5b61006161005c3660046104e1565b6100c9565b60405160ff90911681526020015b60405180910390f35b61008b610086366004610518565b6101b1565b60408051931515845260208401929092529082015260600161006f565b6100bb6100b636600461056b565b6102e7565b60405190815260200161006f565b5f826001600160a01b031663e1758bd86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610106573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061012a91906105bd565b6001600160a01b0316826001600160a01b0316146101a757816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa15801561017e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101a291906105d8565b6101aa565b60125b9392505050565b5f5f5f5f876001600160a01b0316633afb1544886040518263ffffffff1660e01b81526004016101e391815260200190565b606060405180830381865afa1580156101fe573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061022291906105f3565b91955090925090508361023e575f5f5f935093509350506102dd565b60405163e588566f60e01b81526001600160a01b0387811660048301525f91908a169063e588566f90602401606060405180830381865afa158015610285573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102a991906105f3565b9196509093509050846102c6575f5f5f94509450945050506102dd565b6102d886838360126100b68e8d6100c9565b935050505b9450945094915050565b5f845f0361032557604051636a7fd60560e11b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff16101561036d5761036886610342858761063f565b61034d90600a610741565b6103579088610763565b6103619190610763565b88906103bc565b610390565b61039061036161037d868661063f565b61038890600a610741565b879089610403565b92509050806103b257604051630a69350960e41b815260040160405180910390fd5b5095945050505050565b5f5f835f036103d05750600190505f6103fc565b838302838582816103e3576103e361074f565b04146103f5575f5f92509250506103fc565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036104375783828161042d5761042d61074f565b04925050506101aa565b80841161044e5761044e60038515026011186104b9565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6001600160a01b03811681146104de575f5ffd5b50565b5f5f604083850312156104f2575f5ffd5b82356104fd816104ca565b9150602083013561050d816104ca565b809150509250929050565b5f5f5f5f6080858703121561052b575f5ffd5b8435610536816104ca565b935060208501359250604085013561054d816104ca565b9396929550929360600135925050565b60ff811681146104de575f5ffd5b5f5f5f5f5f60a0868803121561057f575f5ffd5b853594506020860135935060408601359250606086013561059f8161055d565b915060808601356105af8161055d565b809150509295509295909350565b5f602082840312156105cd575f5ffd5b81516101aa816104ca565b5f602082840312156105e8575f5ffd5b81516101aa8161055d565b5f5f5f60608486031215610605575f5ffd5b83518015158114610614575f5ffd5b602085015160409095015190969495509392505050565b634e487b7160e01b5f52601160045260245ffd5b60ff82811682821603908111156106585761065861062b565b92915050565b6001815b60018411156106995780850481111561067d5761067d61062b565b600184161561068b57908102905b60019390931c928002610662565b935093915050565b5f826106af57506001610658565b816106bb57505f610658565b81600181146106d157600281146106db576106f7565b6001915050610658565b60ff8411156106ec576106ec61062b565b50506001821b610658565b5060208310610133831016604e8410600b841016171561071a575081810a610658565b6107265f19848461065e565b805f19048211156107395761073961062b565b029392505050565b5f6101aa60ff8416836106a1565b634e487b7160e01b5f52601260045260245ffd5b5f8261077d57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220a2d0045f1a40c146096de98dbd2628303c041e3ca569d8c79a517ba4fe1ef7f964736f6c634300081c0033",
}

// CalcGasFeeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use CalcGasFeeLibMetaData.ABI instead.
var CalcGasFeeLibABI = CalcGasFeeLibMetaData.ABI

// Deprecated: Use CalcGasFeeLibMetaData.Sigs instead.
// CalcGasFeeLibFuncSigs maps the 4-byte function signature to its string representation.
var CalcGasFeeLibFuncSigs = CalcGasFeeLibMetaData.Sigs

// CalcGasFeeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CalcGasFeeLibMetaData.Bin instead.
var CalcGasFeeLibBin = CalcGasFeeLibMetaData.Bin

// DeployCalcGasFeeLib deploys a new Ethereum contract, binding an instance of CalcGasFeeLib to it.
func DeployCalcGasFeeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CalcGasFeeLib, error) {
	parsed, err := CalcGasFeeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CalcGasFeeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CalcGasFeeLib{CalcGasFeeLibCaller: CalcGasFeeLibCaller{contract: contract}, CalcGasFeeLibTransactor: CalcGasFeeLibTransactor{contract: contract}, CalcGasFeeLibFilterer: CalcGasFeeLibFilterer{contract: contract}}, nil
}

// CalcGasFeeLib is an auto generated Go binding around an Ethereum contract.
type CalcGasFeeLib struct {
	CalcGasFeeLibCaller     // Read-only binding to the contract
	CalcGasFeeLibTransactor // Write-only binding to the contract
	CalcGasFeeLibFilterer   // Log filterer for contract events
}

// CalcGasFeeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalcGasFeeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcGasFeeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalcGasFeeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcGasFeeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalcGasFeeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalcGasFeeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalcGasFeeLibSession struct {
	Contract     *CalcGasFeeLib    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalcGasFeeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalcGasFeeLibCallerSession struct {
	Contract *CalcGasFeeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CalcGasFeeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalcGasFeeLibTransactorSession struct {
	Contract     *CalcGasFeeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CalcGasFeeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalcGasFeeLibRaw struct {
	Contract *CalcGasFeeLib // Generic contract binding to access the raw methods on
}

// CalcGasFeeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalcGasFeeLibCallerRaw struct {
	Contract *CalcGasFeeLibCaller // Generic read-only contract binding to access the raw methods on
}

// CalcGasFeeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalcGasFeeLibTransactorRaw struct {
	Contract *CalcGasFeeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalcGasFeeLib creates a new instance of CalcGasFeeLib, bound to a specific deployed contract.
func NewCalcGasFeeLib(address common.Address, backend bind.ContractBackend) (*CalcGasFeeLib, error) {
	contract, err := bindCalcGasFeeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CalcGasFeeLib{CalcGasFeeLibCaller: CalcGasFeeLibCaller{contract: contract}, CalcGasFeeLibTransactor: CalcGasFeeLibTransactor{contract: contract}, CalcGasFeeLibFilterer: CalcGasFeeLibFilterer{contract: contract}}, nil
}

// NewCalcGasFeeLibCaller creates a new read-only instance of CalcGasFeeLib, bound to a specific deployed contract.
func NewCalcGasFeeLibCaller(address common.Address, caller bind.ContractCaller) (*CalcGasFeeLibCaller, error) {
	contract, err := bindCalcGasFeeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalcGasFeeLibCaller{contract: contract}, nil
}

// NewCalcGasFeeLibTransactor creates a new write-only instance of CalcGasFeeLib, bound to a specific deployed contract.
func NewCalcGasFeeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*CalcGasFeeLibTransactor, error) {
	contract, err := bindCalcGasFeeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalcGasFeeLibTransactor{contract: contract}, nil
}

// NewCalcGasFeeLibFilterer creates a new log filterer instance of CalcGasFeeLib, bound to a specific deployed contract.
func NewCalcGasFeeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*CalcGasFeeLibFilterer, error) {
	contract, err := bindCalcGasFeeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalcGasFeeLibFilterer{contract: contract}, nil
}

// bindCalcGasFeeLib binds a generic wrapper to an already deployed contract.
func bindCalcGasFeeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CalcGasFeeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalcGasFeeLib *CalcGasFeeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalcGasFeeLib.Contract.CalcGasFeeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalcGasFeeLib *CalcGasFeeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalcGasFeeLib.Contract.CalcGasFeeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalcGasFeeLib *CalcGasFeeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalcGasFeeLib.Contract.CalcGasFeeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalcGasFeeLib *CalcGasFeeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalcGasFeeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalcGasFeeLib *CalcGasFeeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalcGasFeeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalcGasFeeLib *CalcGasFeeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalcGasFeeLib.Contract.contract.Transact(opts, method, params...)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcGasFeeLib *CalcGasFeeLibCaller) CalculateAmountBWithPrice(opts *bind.CallOpts, amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	var out []interface{}
	err := _CalcGasFeeLib.contract.Call(opts, &out, "calculateAmountBWithPrice", amountA, priceA, priceB, decimalA, decimalB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcGasFeeLib *CalcGasFeeLibSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _CalcGasFeeLib.Contract.CalculateAmountBWithPrice(&_CalcGasFeeLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_CalcGasFeeLib *CalcGasFeeLibCallerSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _CalcGasFeeLib.Contract.CalculateAmountBWithPrice(&_CalcGasFeeLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// CalculateTokenAmountForGasFee is a free data retrieval call binding the contract method 0x889ad9e0.
//
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibCaller) CalculateTokenAmountForGasFee(opts *bind.CallOpts, feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _CalcGasFeeLib.contract.Call(opts, &out, "calculateTokenAmountForGasFee", feed, toChainID, token, nativeTokenAmount)

	outstruct := new(struct {
		Ok        bool
		AmountB   *big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ok = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateTokenAmountForGasFee is a free data retrieval call binding the contract method 0x889ad9e0.
//
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibSession) CalculateTokenAmountForGasFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	return _CalcGasFeeLib.Contract.CalculateTokenAmountForGasFee(&_CalcGasFeeLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// CalculateTokenAmountForGasFee is a free data retrieval call binding the contract method 0x889ad9e0.
//
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibCallerSession) CalculateTokenAmountForGasFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	return _CalcGasFeeLib.Contract.CalculateTokenAmountForGasFee(&_CalcGasFeeLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibCaller) Decimals(opts *bind.CallOpts, feed common.Address, token common.Address) (uint8, error) {
	var out []interface{}
	err := _CalcGasFeeLib.contract.Call(opts, &out, "decimals", feed, token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibSession) Decimals(feed common.Address, token common.Address) (uint8, error) {
	return _CalcGasFeeLib.Contract.Decimals(&_CalcGasFeeLib.CallOpts, feed, token)
}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibCallerSession) Decimals(feed common.Address, token common.Address) (uint8, error) {
	return _CalcGasFeeLib.Contract.Decimals(&_CalcGasFeeLib.CallOpts, feed, token)
}
