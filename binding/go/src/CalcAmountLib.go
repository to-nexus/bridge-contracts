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
	Bin: "0x610713610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c806381d889b41461004e578063b1db08d114610083578063d449a832146100a4575b5f5ffd5b61006161005c366004610479565b6100c9565b6040805193151584526020840192909252908201526060015b60405180910390f35b6100966100913660046104cc565b6101fe565b60405190815260200161007a565b6100b76100b236600461051e565b6102d3565b60405160ff909116815260200161007a565b5f5f5f5f876001600160a01b0316633afb1544886040518263ffffffff1660e01b81526004016100fb91815260200190565b606060405180830381865afa158015610116573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013a9190610539565b919550909250905083610156575f5f5f935093509350506101f4565b60405163e588566f60e01b81526001600160a01b0387811660048301525f91908a169063e588566f90602401606060405180830381865afa15801561019d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101c19190610539565b9196509093509050846101de575f5f5f94509450945050506101f4565b6101ef86838360126100918c6102d3565b935050505b9450945094915050565b5f845f0361023c576040516349671cd160e11b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff1610156102845761027f866102598587610585565b61026490600a610681565b61026e90886106a3565b61027891906106a3565b8890610353565b6102a7565b6102a76102786102948686610585565b61029f90600a610681565b87908961039a565b92509050806102c95760405163c7ef3d4f60e01b815260040160405180910390fd5b5095945050505050565b5f6001600160a01b03821660011461034a57816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610321573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061034591906106c2565b61034d565b60125b92915050565b5f5f835f036103675750600190505f610393565b8383028385828161037a5761037a61068f565b041461038c575f5f9250925050610393565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036103ce578382816103c4576103c461068f565b049250505061044a565b8084116103e5576103e56003851502601118610451565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b634e487b715f52806020526024601cfd5b6001600160a01b0381168114610476575f5ffd5b50565b5f5f5f5f6080858703121561048c575f5ffd5b843561049781610462565b93506020850135925060408501356104ae81610462565b9396929550929360600135925050565b60ff81168114610476575f5ffd5b5f5f5f5f5f60a086880312156104e0575f5ffd5b8535945060208601359350604086013592506060860135610500816104be565b91506080860135610510816104be565b809150509295509295909350565b5f6020828403121561052e575f5ffd5b813561044a81610462565b5f5f5f6060848603121561054b575f5ffd5b8351801515811461055a575f5ffd5b602085015160409095015190969495509392505050565b634e487b7160e01b5f52601160045260245ffd5b60ff828116828216039081111561034d5761034d610571565b6001815b60018411156105d9578085048111156105bd576105bd610571565b60018416156105cb57908102905b60019390931c9280026105a2565b935093915050565b5f826105ef5750600161034d565b816105fb57505f61034d565b8160018114610611576002811461061b57610637565b600191505061034d565b60ff84111561062c5761062c610571565b50506001821b61034d565b5060208310610133831016604e8410600b841016171561065a575081810a61034d565b6106665f19848461059e565b805f190482111561067957610679610571565b029392505050565b5f61044a60ff8416836105e1565b634e487b7160e01b5f52601260045260245ffd5b5f826106bd57634e487b7160e01b5f52601260045260245ffd5b500490565b5f602082840312156106d2575f5ffd5b815161044a816104be56fea2646970667358221220e97f61997576bf562b4d812c6dc5b813b639030da6d559665b8452f54376cc6c64736f6c634300081c0033",
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
