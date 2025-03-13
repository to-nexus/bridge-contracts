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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenAmountForGasFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcGasFeeLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcGasFeeLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
		"889ad9e0": "calculateTokenAmountForGasFee(IPriceFeed,uint256,address,uint256)",
		"0fe74f3b": "decimals(IPriceFeed,address)",
	},
	Bin: "0x610cdb61004d600b8282823980515f1a6073146041577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c80630fe74f3b1461004e578063889ad9e01461007e578063b1db08d1146100b0575b5f5ffd5b61006860048036038101906100639190610679565b6100e0565b60405161007591906106d2565b60405180910390f35b6100986004803603810190610093919061071e565b6101fd565b6040516100a7939291906107ab565b60405180910390f35b6100ca60048036038101906100c5919061080a565b610359565b6040516100d79190610881565b60405180910390f35b5f8273ffffffffffffffffffffffffffffffffffffffff1663e1758bd86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561012a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061014e91906108ae565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146101f2578173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101ed91906108ed565b6101f5565b60125b905092915050565b5f5f5f5f8773ffffffffffffffffffffffffffffffffffffffff16633afb1544886040518263ffffffff1660e01b815260040161023a9190610927565b606060405180830381865afa158015610255573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610279919061097e565b80945081935082965050505083610299575f5f5f9350935093505061034f565b5f8873ffffffffffffffffffffffffffffffffffffffff1663e588566f886040518263ffffffff1660e01b81526004016102d391906109dd565b606060405180830381865afa1580156102ee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610312919061097e565b80955081935082975050505084610333575f5f5f945094509450505061034f565b61034a86838360126103458e8d6100e0565b610359565b935050505b9450945094915050565b5f5f850361039c576040517fd4ffac0a00000000000000000000000000000000000000000000000000000000815260040161039390610a50565b60405180910390fd5b5f8360ff168360ff1610156103f0576103eb8684866103bb9190610a9b565b600a6103c79190610bfe565b876103d29190610c75565b6103dc9190610c75565b8861047890919063ffffffff16565b61042f565b61042e61041f85856104029190610a9b565b600a61040e9190610bfe565b88886104c69092919063ffffffff16565b8861047890919063ffffffff16565b5b80935081925050508061046e576040517fa693509000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5095945050505050565b5f5f5f840361048d5760015f915091506104bf565b5f8385029050838582816104a4576104a3610c48565b5b04146104b6575f5f92509250506104bf565b60018192509250505b9250929050565b5f5f83850290505f5f19858709828110838203039150505f81036104fe578382816104f4576104f3610c48565b5b04925050506105a4565b80841161051d5761051c6105175f8614601260116105ab565b6105c4565b5b5f8486880990508281118203915080830392505f855f038616905080860495508084049350600181825f0304019050808302841793505f600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b5f6105b5846105d5565b82841802821890509392505050565b634e487b715f52806020526024601cfd5b5f8115159050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61060d826105e4565b9050919050565b5f61061e82610603565b9050919050565b61062e81610614565b8114610638575f5ffd5b50565b5f8135905061064981610625565b92915050565b61065881610603565b8114610662575f5ffd5b50565b5f813590506106738161064f565b92915050565b5f5f6040838503121561068f5761068e6105e0565b5b5f61069c8582860161063b565b92505060206106ad85828601610665565b9150509250929050565b5f60ff82169050919050565b6106cc816106b7565b82525050565b5f6020820190506106e55f8301846106c3565b92915050565b5f819050919050565b6106fd816106eb565b8114610707575f5ffd5b50565b5f81359050610718816106f4565b92915050565b5f5f5f5f60808587031215610736576107356105e0565b5b5f6107438782880161063b565b94505060206107548782880161070a565b935050604061076587828801610665565b92505060606107768782880161070a565b91505092959194509250565b5f8115159050919050565b61079681610782565b82525050565b6107a5816106eb565b82525050565b5f6060820190506107be5f83018661078d565b6107cb602083018561079c565b6107d8604083018461079c565b949350505050565b6107e9816106b7565b81146107f3575f5ffd5b50565b5f81359050610804816107e0565b92915050565b5f5f5f5f5f60a08688031215610823576108226105e0565b5b5f6108308882890161070a565b95505060206108418882890161070a565b94505060406108528882890161070a565b9350506060610863888289016107f6565b9250506080610874888289016107f6565b9150509295509295909350565b5f6020820190506108945f83018461079c565b92915050565b5f815190506108a88161064f565b92915050565b5f602082840312156108c3576108c26105e0565b5b5f6108d08482850161089a565b91505092915050565b5f815190506108e7816107e0565b92915050565b5f60208284031215610902576109016105e0565b5b5f61090f848285016108d9565b91505092915050565b610921816106eb565b82525050565b5f60208201905061093a5f830184610918565b92915050565b61094981610782565b8114610953575f5ffd5b50565b5f8151905061096481610940565b92915050565b5f81519050610978816106f4565b92915050565b5f5f5f60608486031215610995576109946105e0565b5b5f6109a286828701610956565b93505060206109b38682870161096a565b92505060406109c48682870161096a565b9150509250925092565b6109d781610603565b82525050565b5f6020820190506109f05f8301846109ce565b92915050565b5f82825260208201905092915050565b7f70726963654100000000000000000000000000000000000000000000000000005f82015250565b5f610a3a6006836109f6565b9150610a4582610a06565b602082019050919050565b5f6020820190508181035f830152610a6781610a2e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610aa5826106b7565b9150610ab0836106b7565b9250828203905060ff811115610ac957610ac8610a6e565b5b92915050565b5f8160011c9050919050565b5f5f8291508390505b6001851115610b2457808604811115610b0057610aff610a6e565b5b6001851615610b0f5780820291505b8081029050610b1d85610acf565b9450610ae4565b94509492505050565b5f82610b3c5760019050610bf7565b81610b49575f9050610bf7565b8160018114610b5f5760028114610b6957610b98565b6001915050610bf7565b60ff841115610b7b57610b7a610a6e565b5b8360020a915084821115610b9257610b91610a6e565b5b50610bf7565b5060208310610133831016604e8410600b8410161715610bcd5782820a905083811115610bc857610bc7610a6e565b5b610bf7565b610bda8484846001610adb565b92509050818404811115610bf157610bf0610a6e565b5b81810290505b9392505050565b5f610c08826106eb565b9150610c13836106b7565b9250610c407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610b2d565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610c7f826106eb565b9150610c8a836106eb565b925082610c9a57610c99610c48565b5b82820490509291505056fea2646970667358221220476fe9dab25ad7ddaab5c31c351efd44bda216c48f96767d23c2307567dcc5a064736f6c634300081c0033",
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
