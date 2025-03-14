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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmountForGasFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcGasFeeLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcGasFeeLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
		"889ad9e0": "calculateTokenAmountForGasFee(IPriceFeed,uint256,address,uint256)",
		"d449a832": "decimals(address)",
	},
	Bin: "0x610713610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c8063889ad9e01461004e578063b1db08d114610083578063d449a832146100a4575b5f5ffd5b61006161005c366004610479565b6100c9565b6040805193151584526020840192909252908201526060015b60405180910390f35b6100966100913660046104cc565b6101fe565b60405190815260200161007a565b6100b76100b236600461051e565b6102d3565b60405160ff909116815260200161007a565b5f5f5f5f876001600160a01b0316633afb1544886040518263ffffffff1660e01b81526004016100fb91815260200190565b606060405180830381865afa158015610116573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013a9190610539565b919550909250905083610156575f5f5f935093509350506101f4565b60405163e588566f60e01b81526001600160a01b0387811660048301525f91908a169063e588566f90602401606060405180830381865afa15801561019d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101c19190610539565b9196509093509050846101de575f5f5f94509450945050506101f4565b6101ef86838360126100918c6102d3565b935050505b9450945094915050565b5f845f0361023c57604051636a7fd60560e11b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff1610156102845761027f866102598587610585565b61026490600a610681565b61026e90886106a3565b61027891906106a3565b8890610353565b6102a7565b6102a76102786102948686610585565b61029f90600a610681565b87908961039a565b92509050806102c957604051630a69350960e41b815260040160405180910390fd5b5095945050505050565b5f6001600160a01b03821660011461034a57816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610321573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061034591906106c2565b61034d565b60125b92915050565b5f5f835f036103675750600190505f610393565b8383028385828161037a5761037a61068f565b041461038c575f5f9250925050610393565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036103ce578382816103c4576103c461068f565b049250505061044a565b8084116103e5576103e56003851502601118610451565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b634e487b715f52806020526024601cfd5b6001600160a01b0381168114610476575f5ffd5b50565b5f5f5f5f6080858703121561048c575f5ffd5b843561049781610462565b93506020850135925060408501356104ae81610462565b9396929550929360600135925050565b60ff81168114610476575f5ffd5b5f5f5f5f5f60a086880312156104e0575f5ffd5b8535945060208601359350604086013592506060860135610500816104be565b91506080860135610510816104be565b809150509295509295909350565b5f6020828403121561052e575f5ffd5b813561044a81610462565b5f5f5f6060848603121561054b575f5ffd5b8351801515811461055a575f5ffd5b602085015160409095015190969495509392505050565b634e487b7160e01b5f52601160045260245ffd5b60ff828116828216039081111561034d5761034d610571565b6001815b60018411156105d9578085048111156105bd576105bd610571565b60018416156105cb57908102905b60019390931c9280026105a2565b935093915050565b5f826105ef5750600161034d565b816105fb57505f61034d565b8160018114610611576002811461061b57610637565b600191505061034d565b60ff84111561062c5761062c610571565b50506001821b61034d565b5060208310610133831016604e8410600b841016171561065a575081810a61034d565b6106665f19848461059e565b805f190482111561067957610679610571565b029392505050565b5f61044a60ff8416836105e1565b634e487b7160e01b5f52601260045260245ffd5b5f826106bd57634e487b7160e01b5f52601260045260245ffd5b500490565b5f602082840312156106d2575f5ffd5b815161044a816104be56fea26469706673582212203e73dbdb631e3b6b2c07c2cef30127f0338907f6a935642fb62b8a743b7d03ba64736f6c634300081c0033",
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
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibCaller) CalculateTokenAmountForGasFee(opts *bind.CallOpts, feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	var out []interface{}
	err := _CalcGasFeeLib.contract.Call(opts, &out, "calculateTokenAmountForGasFee", feed, toChainID, token, nativeTokenAmount)

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

// CalculateTokenAmountForGasFee is a free data retrieval call binding the contract method 0x889ad9e0.
//
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibSession) CalculateTokenAmountForGasFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	return _CalcGasFeeLib.Contract.CalculateTokenAmountForGasFee(&_CalcGasFeeLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// CalculateTokenAmountForGasFee is a free data retrieval call binding the contract method 0x889ad9e0.
//
// Solidity: function calculateTokenAmountForGasFee(IPriceFeed feed, uint256 toChainID, address token, uint256 nativeTokenAmount) view returns(bool ok, uint256 tokenAmount, uint256 updatedAt)
func (_CalcGasFeeLib *CalcGasFeeLibCallerSession) CalculateTokenAmountForGasFee(feed common.Address, toChainID *big.Int, token common.Address, nativeTokenAmount *big.Int) (struct {
	Ok          bool
	TokenAmount *big.Int
	UpdatedAt   *big.Int
}, error) {
	return _CalcGasFeeLib.Contract.CalculateTokenAmountForGasFee(&_CalcGasFeeLib.CallOpts, feed, toChainID, token, nativeTokenAmount)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibCaller) Decimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _CalcGasFeeLib.contract.Call(opts, &out, "decimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibSession) Decimals(token common.Address) (uint8, error) {
	return _CalcGasFeeLib.Contract.Decimals(&_CalcGasFeeLib.CallOpts, token)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address token) view returns(uint8 _decimals)
func (_CalcGasFeeLib *CalcGasFeeLibCallerSession) Decimals(token common.Address) (uint8, error) {
	return _CalcGasFeeLib.Contract.Decimals(&_CalcGasFeeLib.CallOpts, token)
}
