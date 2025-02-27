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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"feed\",\"type\":\"IPriceFeed\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"calculateAmountB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceB\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimalB\",\"type\":\"uint8\"}],\"name\":\"calculateAmountBWithPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedLibOverflow\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"0f126c6c": "calculateAmountB(IPriceFeed,address,address,uint256)",
		"b1db08d1": "calculateAmountBWithPrice(uint256,uint256,uint256,uint8,uint8)",
	},
	Bin: "0x6109c4610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061003f575f3560e01c80630f126c6c14610043578063b1db08d114610078575b5f5ffd5b610056610051366004610556565b610099565b6040805193151584526020840192909252908201526060015b60405180910390f35b61008b6100863660046105b2565b610274565b60405190815260200161006f565b5f5f5f846001600160a01b0316866001600160a01b0316036100c35750600191508290504261026a565b6040805160028082526060820183525f9260208301908036833701905050905086815f815181106100f6576100f6610618565b60200260200101906001600160a01b031690816001600160a01b031681525050858160018151811061012a5761012a610618565b6001600160a01b0392831660209182029290920101526040516347dad24160e11b81526060918291908b1690638fb5a4829061016a90869060040161062c565b5f60405180830381865afa158015610184573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101ab919081019061072b565b8251909650919350915082905f906101c5576101c5610618565b602002602001015180156101f05750816001815181106101e7576101e7610618565b60200260200101515b955085610208575f5f5f95509550955050505061026a565b5f5f6102148c8c610349565b61021e8d8c610349565b9150915061026289845f8151811061023857610238610618565b60200260200101518560018151811061025357610253610618565b60200260200101518585610274565b965050505050505b9450945094915050565b5f845f036102b257604051633a0b0f2960e01b815260206004820152600660248201526570726963654160d01b604482015260640160405180910390fd5b5f8360ff168360ff1610156102fa576102f5866102cf8587610815565b6102da90600a610917565b6102e49088610939565b6102ee9190610939565b8890610431565b61031d565b61031d6102ee61030a8686610815565b61031590600a610917565b879089610478565b925090508061033f57604051633961e4cf60e11b815260040160405180910390fd5b5095945050505050565b5f826001600160a01b031663167b78cd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610386573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103aa9190610958565b6001600160a01b0316826001600160a01b03161461042757816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103fe573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104229190610973565b61042a565b60125b9392505050565b5f5f835f036104455750600190505f610471565b8383028385828161045857610458610925565b041461046a575f5f9250925050610471565b6001925090505b9250929050565b5f838302815f1985870982811083820303915050805f036104ac578382816104a2576104a2610925565b049250505061042a565b8084116104c3576104c3600385150260111861052e565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b634e487b715f52806020526024601cfd5b6001600160a01b0381168114610553575f5ffd5b50565b5f5f5f5f60808587031215610569575f5ffd5b84356105748161053f565b935060208501356105848161053f565b925060408501356105948161053f565b9396929550929360600135925050565b60ff81168114610553575f5ffd5b5f5f5f5f5f60a086880312156105c6575f5ffd5b85359450602086013593506040860135925060608601356105e6816105a4565b915060808601356105f6816105a4565b809150509295509295909350565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b8181101561066c5783516001600160a01b0316835260209384019390920191600101610645565b509095945050505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156106a0576106a0610604565b604052919050565b5f67ffffffffffffffff8211156106c1576106c1610604565b5060051b60200190565b5f82601f8301126106da575f5ffd5b81516106ed6106e8826106a8565b610677565b8082825260208201915060208360051b86010192508583111561070e575f5ffd5b602085015b8381101561033f578051835260209283019201610713565b5f5f5f6060848603121561073d575f5ffd5b835167ffffffffffffffff811115610753575f5ffd5b8401601f81018613610763575f5ffd5b80516107716106e8826106a8565b8082825260208201915060208360051b850101925088831115610792575f5ffd5b6020840193505b828410156107c157835180151581146107b0575f5ffd5b825260209384019390910190610799565b80965050505050602084015167ffffffffffffffff8111156107e1575f5ffd5b6107ed868287016106cb565b604095909501519396949550929392505050565b634e487b7160e01b5f52601160045260245ffd5b60ff828116828216039081111561082e5761082e610801565b92915050565b6001815b600184111561086f5780850481111561085357610853610801565b600184161561086157908102905b60019390931c928002610838565b935093915050565b5f826108855750600161082e565b8161089157505f61082e565b81600181146108a757600281146108b1576108cd565b600191505061082e565b60ff8411156108c2576108c2610801565b50506001821b61082e565b5060208310610133831016604e8410600b84101617156108f0575081810a61082e565b6108fc5f198484610834565b805f190482111561090f5761090f610801565b029392505050565b5f61042a60ff841683610877565b634e487b7160e01b5f52601260045260245ffd5b5f8261095357634e487b7160e01b5f52601260045260245ffd5b500490565b5f60208284031215610968575f5ffd5b815161042a8161053f565b5f60208284031215610983575f5ffd5b815161042a816105a456fea2646970667358221220fa197f8c43377ad90c322f948378b36e79f6709e3a00769a73434c85cea6c5fe64736f6c634300081c0033",
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
