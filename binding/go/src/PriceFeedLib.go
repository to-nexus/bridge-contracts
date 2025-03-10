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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"calculateAmountB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"0f126c6c": "calculateAmountB(IPriceFeed,address,address,uint256)",
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
		"0fe74f3b": "decimals(IPriceFeed,address)",
	},
	Bin: "0x610a2b610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c80630f126c6c1461004e5780630fe74f3b14610083578063b1db08d1146100a8575b5f5ffd5b61006161005c366004610586565b6100c9565b6040805193151584526020840192909252908201526060015b60405180910390f35b6100966100913660046105d4565b6102a4565b60405160ff909116815260200161007a565b6100bb6100b6366004610619565b61038c565b60405190815260200161007a565b5f5f5f846001600160a01b0316866001600160a01b0316036100f35750600191508290504261029a565b6040805160028082526060820183525f9260208301908036833701905050905086815f815181106101265761012661067f565b60200260200101906001600160a01b031690816001600160a01b031681525050858160018151811061015a5761015a61067f565b6001600160a01b0392831660209182029290920101526040516347dad24160e11b81526060918291908b1690638fb5a4829061019a908690600401610693565b5f60405180830381865afa1580156101b4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101db9190810190610792565b8251909650919350915082905f906101f5576101f561067f565b602002602001015180156102205750816001815181106102175761021761067f565b60200260200101515b955085610238575f5f5f95509550955050505061029a565b5f5f6102448c8c6102a4565b61024e8d8c6102a4565b9150915061029289845f815181106102685761026861067f565b6020026020010151856001815181106102835761028361067f565b6020026020010151858561038c565b965050505050505b9450945094915050565b5f826001600160a01b031663e1758bd86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103059190610868565b6001600160a01b0316826001600160a01b03161461038257816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610359573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037d9190610883565b610385565b60125b9392505050565b5f845f036103ca57604051633a0b0f2960e01b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff1610156104125761040d866103e785876108b2565b6103f290600a6109b4565b6103fc90886109d6565b61040691906109d6565b8890610461565b610435565b61043561040661042286866108b2565b61042d90600a6109b4565b8790896104a8565b925090508061045757604051633961e4cf60e11b815260040160405180910390fd5b5095945050505050565b5f5f835f036104755750600190505f6104a1565b83830283858281610488576104886109c2565b041461049a575f5f92509250506104a1565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036104dc578382816104d2576104d26109c2565b0492505050610385565b8084116104f3576104f3600385150260111861055e565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6001600160a01b0381168114610583575f5ffd5b50565b5f5f5f5f60808587031215610599575f5ffd5b84356105a48161056f565b935060208501356105b48161056f565b925060408501356105c48161056f565b9396929550929360600135925050565b5f5f604083850312156105e5575f5ffd5b82356105f08161056f565b915060208301356106008161056f565b809150509250929050565b60ff81168114610583575f5ffd5b5f5f5f5f5f60a0868803121561062d575f5ffd5b853594506020860135935060408601359250606086013561064d8161060b565b9150608086013561065d8161060b565b809150509295509295909350565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b818110156106d35783516001600160a01b03168352602093840193909201916001016106ac565b509095945050505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156107075761070761066b565b604052919050565b5f67ffffffffffffffff8211156107285761072861066b565b5060051b60200190565b5f82601f830112610741575f5ffd5b815161075461074f8261070f565b6106de565b8082825260208201915060208360051b860101925085831115610775575f5ffd5b602085015b8381101561045757805183526020928301920161077a565b5f5f5f606084860312156107a4575f5ffd5b835167ffffffffffffffff8111156107ba575f5ffd5b8401601f810186136107ca575f5ffd5b80516107d861074f8261070f565b8082825260208201915060208360051b8501019250888311156107f9575f5ffd5b6020840193505b828410156108285783518015158114610817575f5ffd5b825260209384019390910190610800565b80965050505050602084015167ffffffffffffffff811115610848575f5ffd5b61085486828701610732565b604095909501519396949550929392505050565b5f60208284031215610878575f5ffd5b81516103858161056f565b5f60208284031215610893575f5ffd5b81516103858161060b565b634e487b7160e01b5f52601160045260245ffd5b60ff82811682821603908111156108cb576108cb61089e565b92915050565b6001815b600184111561090c578085048111156108f0576108f061089e565b60018416156108fe57908102905b60019390931c9280026108d5565b935093915050565b5f82610922575060016108cb565b8161092e57505f6108cb565b8160018114610944576002811461094e5761096a565b60019150506108cb565b60ff84111561095f5761095f61089e565b50506001821b6108cb565b5060208310610133831016604e8410600b841016171561098d575081810a6108cb565b6109995f1984846108d1565b805f19048211156109ac576109ac61089e565b029392505050565b5f61038560ff841683610914565b634e487b7160e01b5f52601260045260245ffd5b5f826109f057634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220b4d9493c22afbc03a02d5b3741168c215781e9ac0f51f73d3f2bdacfc8afd67e64736f6c634300081c0033",
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

// CalculateAmountB is a free data retrieval call binding the contract method 0x0f126c6c.
//
// Solidity: function calculateAmountB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_PriceFeedLib *PriceFeedLibCaller) CalculateAmountB(opts *bind.CallOpts, feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "calculateAmountB", feed, tokenA, tokenB, amountA)

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

// CalculateAmountB is a free data retrieval call binding the contract method 0x0f126c6c.
//
// Solidity: function calculateAmountB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_PriceFeedLib *PriceFeedLibSession) CalculateAmountB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeedLib.Contract.CalculateAmountB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// CalculateAmountB is a free data retrieval call binding the contract method 0x0f126c6c.
//
// Solidity: function calculateAmountB(IPriceFeed feed, address tokenA, address tokenB, uint256 amountA) view returns(bool ok, uint256 amountB, uint256 updatedAt)
func (_PriceFeedLib *PriceFeedLibCallerSession) CalculateAmountB(feed common.Address, tokenA common.Address, tokenB common.Address, amountA *big.Int) (struct {
	Ok        bool
	AmountB   *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeedLib.Contract.CalculateAmountB(&_PriceFeedLib.CallOpts, feed, tokenA, tokenB, amountA)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCaller) CalculateAmountBWithPrice(opts *bind.CallOpts, amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "calculateAmountBWithPrice", amountA, priceA, priceB, decimalA, decimalB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.CalculateAmountBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// CalculateAmountBWithPrice is a free data retrieval call binding the contract method 0xb1db08d1.
//
// Solidity: function calculateAmountBWithPrice(uint256 amountA, uint256 priceA, uint256 priceB, uint8 decimalA, uint8 decimalB) pure returns(uint256 amountB)
func (_PriceFeedLib *PriceFeedLibCallerSession) CalculateAmountBWithPrice(amountA *big.Int, priceA *big.Int, priceB *big.Int, decimalA uint8, decimalB uint8) (*big.Int, error) {
	return _PriceFeedLib.Contract.CalculateAmountBWithPrice(&_PriceFeedLib.CallOpts, amountA, priceA, priceB, decimalA, decimalB)
}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_PriceFeedLib *PriceFeedLibCaller) Decimals(opts *bind.CallOpts, feed common.Address, token common.Address) (uint8, error) {
	var out []interface{}
	err := _PriceFeedLib.contract.Call(opts, &out, "decimals", feed, token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_PriceFeedLib *PriceFeedLibSession) Decimals(feed common.Address, token common.Address) (uint8, error) {
	return _PriceFeedLib.Contract.Decimals(&_PriceFeedLib.CallOpts, feed, token)
}

// Decimals is a free data retrieval call binding the contract method 0x0fe74f3b.
//
// Solidity: function decimals(IPriceFeed feed, address token) view returns(uint8 _decimals)
func (_PriceFeedLib *PriceFeedLibCallerSession) Decimals(feed common.Address, token common.Address) (uint8, error) {
	return _PriceFeedLib.Contract.Decimals(&_PriceFeedLib.CallOpts, feed, token)
}
