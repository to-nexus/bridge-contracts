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

// IBridgeFeeStationGasInfo is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFeeStationGasInfo struct {
	ChainID  *big.Int
	GasPrice *big.Int
	GasToken common.Address
}

// BridgeFeeStationMetaData contains all meta data concerning the BridgeFeeStation contract.
var BridgeFeeStationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"}],\"name\":\"getGasInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"gasToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFeeStation.GasInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setExFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"setFinalizeBridgeGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"setGasInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minimumAmountList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exFeeRateList\",\"type\":\"uint256[]\"}],\"name\":\"setTokenFeeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exFeeRate\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeStationExchangeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizeBridgeGas\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeStationFinalizeBridgeGasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeStationGasPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPriceFeed\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"BridgeFeeStationPriceFeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeFeeStationCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BridgeFeeStationCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeStationChainAleadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeFeeStationInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"96ce0795": "denominator()",
		"ae766389": "estimateFee(uint256,address,uint256)",
		"a3fc93b3": "getGasInfo(uint256)",
		"09585b1c": "getTokenFee(uint256,address)",
		"8da5cb5b": "owner()",
		"ce2e5c66": "removePriceFeed()",
		"1b034c1c": "removeTokenFee(address)",
		"715018a6": "renounceOwnership()",
		"005d0f10": "setExFeeRate(uint256)",
		"f289d3ba": "setFinalizeBridgeGas(uint256)",
		"3f0cf10d": "setGasInfo(uint256,address,uint256)",
		"724e78da": "setPriceFeed(address)",
		"dffa229e": "setTokenFee(address,uint256,uint256)",
		"6731d860": "setTokenFeeBatch(address[],uint256[],uint256[])",
		"f2fde38b": "transferOwnership(address)",
		"1f96131e": "updateGasPrice(uint256,uint256)",
	},
	Bin: "0x608060405234801561000f575f5ffd5b50604051610fc2380380610fc283398101604081905261002e916100fc565b338061005457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61005d816100ad565b50805f036100a257604051631c622e5b60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b604482015260640161004b565b60035560025561011e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f5f6040838503121561010d575f5ffd5b505080516020909101519092909150565b610e978061012b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100fa575f3560e01c80638da5cb5b11610093578063ce2e5c6611610063578063ce2e5c6614610228578063dffa229e14610230578063f289d3ba14610243578063f2fde38b14610256575f5ffd5b80638da5cb5b146101ad57806396ce0795146101c7578063a3fc93b3146101d7578063ae76638914610215575f5ffd5b80633f0cf10d116100ce5780633f0cf10d1461016c5780636731d8601461017f578063715018a614610192578063724e78da1461019a575f5ffd5b80625d0f10146100fe57806309585b1c146101135780631b034c1c146101465780631f96131e14610159575b5f5ffd5b61011161010c366004610b1d565b610269565b005b610126610121366004610b48565b6102ad565b604080519384526020840192909252908201526060015b60405180910390f35b610111610154366004610b76565b610335565b610111610167366004610b91565b6103d6565b61011161017a366004610bb1565b610431565b61011161018d366004610cb8565b610509565b6101116105af565b6101116101a8366004610b76565b6105c2565b5f546040516001600160a01b03909116815260200161013d565b6040516103e8815260200161013d565b6101ea6101e5366004610b1d565b610656565b604080518251815260208084015190820152918101516001600160a01b03169082015260600161013d565b610126610223366004610bb1565b6106c1565b6101116106ef565b61011161023e366004610da6565b610733565b610111610251366004610b1d565b6107d2565b610111610264366004610b76565b610853565b61027161088d565b60028190556040518181527f0312f690e5eab8afc0fc83c2e2558848b2debb98c72fe871573d8950080deb71906020015b60405180910390a150565b5f5f5f6102ba85856108b9565b506001600160a01b038086165f81815260046020908152604091829020825160608101845281549095168086526001820154928601929092526002015491840191909152929450909114610311575f60025461031c565b806020015181604001515b90945091506001820161032d575f91505b509250925092565b61033d61088d565b6001600160a01b038116610381576040516304d45a0560e51b81526020600482015260056024820152643a37b5b2b760d91b60448201526064015b60405180910390fd5b6001600160a01b038082165f81815260046020526040902054909116036103d3576001600160a01b0381165f90815260046020526040812080546001600160a01b031916815560018101829055600201555b50565b6103de61088d565b5f82815260056020526040908190206001018290555182907f5c39f52fe1fce1c47c4175fc8ba058e40f59adc2571a3c282c4def80c5c12cbe906104259084815260200190565b60405180910390a25050565b61043961088d565b5f8381526005602052604090205483901561046a576040516302be268960e21b815260040161037891815260200190565b506001600160a01b0382166104ad57604051631c622e5b60e11b815260206004820152600860248201526733b0b9aa37b5b2b760c11b6044820152606401610378565b6040805160608101825284815260208082019384526001600160a01b039485168284019081525f968752600590915291909420935184559051600184015551600290920180546001600160a01b03191692909116919091179055565b61051161088d565b81518351148015610523575080518351145b61054057604051631582cf0d60e21b815260040160405180910390fd5b5f5b83518110156105a9576105a184828151811061056057610560610dd8565b602002602001015184838151811061057a5761057a610dd8565b602002602001015184848151811061059457610594610dd8565b6020026020010151610733565b600101610542565b50505050565b6105b761088d565b6105c05f610a06565b565b6105ca61088d565b6001600160a01b03811661060d576040516304d45a0560e51b81526020600482015260096024820152681c1c9a58d95199595960ba1b6044820152606401610378565b600180546001600160a01b0319166001600160a01b0383169081179091556040517fcde95eaba5b290eeb4ecebcb3aa3a398f3098f54a7f754664d87876b10fe6a98905f90a250565b61068060405180606001604052805f81526020015f81526020015f6001600160a01b031681525090565b505f90815260056020908152604091829020825160608101845281548152600182015492810192909252600201546001600160a01b03169181019190915290565b5f5f5f5f6106cf87876102ad565b919550935090506106e385826103e8610a55565b91505093509350939050565b6106f761088d565b600180546001600160a01b03191690556040515f907fcde95eaba5b290eeb4ecebcb3aa3a398f3098f54a7f754664d87876b10fe6a98908290a2565b61073b61088d565b6001600160a01b03831661077a576040516304d45a0560e51b81526020600482015260056024820152643a37b5b2b760d91b6044820152606401610378565b604080516060810182526001600160a01b0394851680825260208083019586528284019485525f9182526004905291909120905181546001600160a01b03191694169390931783559051600183015551600290910155565b6107da61088d565b805f0361081e57604051631c622e5b60e11b815260206004820152601160248201527066696e616c697a6542726964676547617360781b6044820152606401610378565b60038190556040518181527f570de5cf44bf6056ffb0b1a5be86c62c41ef8c92f81427627f5eadce4cffca7a906020016102a2565b61085b61088d565b6001600160a01b03811661088457604051631e4fbdf760e01b81525f6004820152602401610378565b6103d381610a06565b5f546001600160a01b031633146105c05760405163118cdaa760e01b8152336004820152602401610378565b5f828152600560209081526040808320815160608101835281548152600180830154948201949094526002909101546001600160a01b0390811692820192909252915483929116158061090d575080518514155b1561091e575f5f92509250506109ff565b600154604082015160035460208401515f9373__$79c24c681325f76413d25d3c06c8219b6e$__93630f126c6c936001600160a01b039092169290918a916109669190610dec565b6040516001600160e01b031960e087901b1681526001600160a01b0394851660048201529284166024840152921660448201526064810191909152608401606060405180830381865af41580156109bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109e39190610e15565b90955093509050806109fc575f5f9350935050506109ff565b50505b9250929050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f838302815f1985870982811083820303915050805f03610a8957838281610a7f57610a7f610e4d565b0492505050610b05565b808411610aa057610aa06003851502601118610b0c565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150505b9392505050565b634e487b715f52806020526024601cfd5b5f60208284031215610b2d575f5ffd5b5035919050565b6001600160a01b03811681146103d3575f5ffd5b5f5f60408385031215610b59575f5ffd5b823591506020830135610b6b81610b34565b809150509250929050565b5f60208284031215610b86575f5ffd5b8135610b0581610b34565b5f5f60408385031215610ba2575f5ffd5b50508035926020909101359150565b5f5f5f60608486031215610bc3575f5ffd5b833592506020840135610bd581610b34565b929592945050506040919091013590565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c2357610c23610be6565b604052919050565b5f67ffffffffffffffff821115610c4457610c44610be6565b5060051b60200190565b5f82601f830112610c5d575f5ffd5b8135610c70610c6b82610c2b565b610bfa565b8082825260208201915060208360051b860101925085831115610c91575f5ffd5b602085015b83811015610cae578035835260209283019201610c96565b5095945050505050565b5f5f5f60608486031215610cca575f5ffd5b833567ffffffffffffffff811115610ce0575f5ffd5b8401601f81018613610cf0575f5ffd5b8035610cfe610c6b82610c2b565b8082825260208201915060208360051b850101925088831115610d1f575f5ffd5b6020840193505b82841015610d4a578335610d3981610b34565b825260209384019390910190610d26565b9550505050602084013567ffffffffffffffff811115610d68575f5ffd5b610d7486828701610c4e565b925050604084013567ffffffffffffffff811115610d90575f5ffd5b610d9c86828701610c4e565b9150509250925092565b5f5f5f60608486031215610db8575f5ffd5b8335610dc381610b34565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b8082028115828204841417610e0f57634e487b7160e01b5f52601160045260245ffd5b92915050565b5f5f5f60608486031215610e27575f5ffd5b83518015158114610e36575f5ffd5b602085015160409095015190969495509392505050565b634e487b7160e01b5f52601260045260245ffdfea2646970667358221220c356dbe557da2ed0f98cad2be3b1242872e52fabc9d3d093f140afc40d6b612364736f6c634300081c0033",
}

// BridgeFeeStationABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeFeeStationMetaData.ABI instead.
var BridgeFeeStationABI = BridgeFeeStationMetaData.ABI

// Deprecated: Use BridgeFeeStationMetaData.Sigs instead.
// BridgeFeeStationFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFeeStationFuncSigs = BridgeFeeStationMetaData.Sigs

// BridgeFeeStationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeFeeStationMetaData.Bin instead.
var BridgeFeeStationBin = BridgeFeeStationMetaData.Bin

// DeployBridgeFeeStation deploys a new Ethereum contract, binding an instance of BridgeFeeStation to it.
func DeployBridgeFeeStation(auth *bind.TransactOpts, backend bind.ContractBackend, exFeeRate *big.Int, finalizeBridgeGas *big.Int) (common.Address, *types.Transaction, *BridgeFeeStation, error) {
	parsed, err := BridgeFeeStationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeFeeStationBin), backend, exFeeRate, finalizeBridgeGas)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeFeeStation{BridgeFeeStationCaller: BridgeFeeStationCaller{contract: contract}, BridgeFeeStationTransactor: BridgeFeeStationTransactor{contract: contract}, BridgeFeeStationFilterer: BridgeFeeStationFilterer{contract: contract}}, nil
}

// BridgeFeeStation is an auto generated Go binding around an Ethereum contract.
type BridgeFeeStation struct {
	BridgeFeeStationCaller     // Read-only binding to the contract
	BridgeFeeStationTransactor // Write-only binding to the contract
	BridgeFeeStationFilterer   // Log filterer for contract events
}

// BridgeFeeStationCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeFeeStationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeStationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeFeeStationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeStationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFeeStationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFeeStationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeFeeStationSession struct {
	Contract     *BridgeFeeStation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeFeeStationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeFeeStationCallerSession struct {
	Contract *BridgeFeeStationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BridgeFeeStationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeFeeStationTransactorSession struct {
	Contract     *BridgeFeeStationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BridgeFeeStationRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeFeeStationRaw struct {
	Contract *BridgeFeeStation // Generic contract binding to access the raw methods on
}

// BridgeFeeStationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeFeeStationCallerRaw struct {
	Contract *BridgeFeeStationCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeFeeStationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeFeeStationTransactorRaw struct {
	Contract *BridgeFeeStationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeFeeStation creates a new instance of BridgeFeeStation, bound to a specific deployed contract.
func NewBridgeFeeStation(address common.Address, backend bind.ContractBackend) (*BridgeFeeStation, error) {
	contract, err := bindBridgeFeeStation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStation{BridgeFeeStationCaller: BridgeFeeStationCaller{contract: contract}, BridgeFeeStationTransactor: BridgeFeeStationTransactor{contract: contract}, BridgeFeeStationFilterer: BridgeFeeStationFilterer{contract: contract}}, nil
}

// NewBridgeFeeStationCaller creates a new read-only instance of BridgeFeeStation, bound to a specific deployed contract.
func NewBridgeFeeStationCaller(address common.Address, caller bind.ContractCaller) (*BridgeFeeStationCaller, error) {
	contract, err := bindBridgeFeeStation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationCaller{contract: contract}, nil
}

// NewBridgeFeeStationTransactor creates a new write-only instance of BridgeFeeStation, bound to a specific deployed contract.
func NewBridgeFeeStationTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeFeeStationTransactor, error) {
	contract, err := bindBridgeFeeStation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationTransactor{contract: contract}, nil
}

// NewBridgeFeeStationFilterer creates a new log filterer instance of BridgeFeeStation, bound to a specific deployed contract.
func NewBridgeFeeStationFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFeeStationFilterer, error) {
	contract, err := bindBridgeFeeStation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationFilterer{contract: contract}, nil
}

// bindBridgeFeeStation binds a generic wrapper to an already deployed contract.
func bindBridgeFeeStation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeFeeStationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFeeStation *BridgeFeeStationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFeeStation.Contract.BridgeFeeStationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFeeStation *BridgeFeeStationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.BridgeFeeStationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFeeStation *BridgeFeeStationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.BridgeFeeStationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeFeeStation *BridgeFeeStationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeFeeStation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeFeeStation *BridgeFeeStationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeFeeStation *BridgeFeeStationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.contract.Transact(opts, method, params...)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeFeeStation *BridgeFeeStationCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeFeeStation.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeFeeStation *BridgeFeeStationSession) Denominator() (*big.Int, error) {
	return _BridgeFeeStation.Contract.Denominator(&_BridgeFeeStation.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() pure returns(uint256)
func (_BridgeFeeStation *BridgeFeeStationCallerSession) Denominator() (*big.Int, error) {
	return _BridgeFeeStation.Contract.Denominator(&_BridgeFeeStation.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BridgeFeeStation *BridgeFeeStationCaller) EstimateFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	var out []interface{}
	err := _BridgeFeeStation.contract.Call(opts, &out, "estimateFee", remoteChainID, token, value)

	outstruct := new(struct {
		MinimumAmount *big.Int
		GasFee        *big.Int
		ExFee         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BridgeFeeStation *BridgeFeeStationSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BridgeFeeStation.Contract.EstimateFee(&_BridgeFeeStation.CallOpts, remoteChainID, token, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0xae766389.
//
// Solidity: function estimateFee(uint256 remoteChainID, address token, uint256 value) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFee)
func (_BridgeFeeStation *BridgeFeeStationCallerSession) EstimateFee(remoteChainID *big.Int, token common.Address, value *big.Int) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFee         *big.Int
}, error) {
	return _BridgeFeeStation.Contract.EstimateFee(&_BridgeFeeStation.CallOpts, remoteChainID, token, value)
}

// GetGasInfo is a free data retrieval call binding the contract method 0xa3fc93b3.
//
// Solidity: function getGasInfo(uint256 remoteChainID) view returns((uint256,uint256,address))
func (_BridgeFeeStation *BridgeFeeStationCaller) GetGasInfo(opts *bind.CallOpts, remoteChainID *big.Int) (IBridgeFeeStationGasInfo, error) {
	var out []interface{}
	err := _BridgeFeeStation.contract.Call(opts, &out, "getGasInfo", remoteChainID)

	if err != nil {
		return *new(IBridgeFeeStationGasInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeFeeStationGasInfo)).(*IBridgeFeeStationGasInfo)

	return out0, err

}

// GetGasInfo is a free data retrieval call binding the contract method 0xa3fc93b3.
//
// Solidity: function getGasInfo(uint256 remoteChainID) view returns((uint256,uint256,address))
func (_BridgeFeeStation *BridgeFeeStationSession) GetGasInfo(remoteChainID *big.Int) (IBridgeFeeStationGasInfo, error) {
	return _BridgeFeeStation.Contract.GetGasInfo(&_BridgeFeeStation.CallOpts, remoteChainID)
}

// GetGasInfo is a free data retrieval call binding the contract method 0xa3fc93b3.
//
// Solidity: function getGasInfo(uint256 remoteChainID) view returns((uint256,uint256,address))
func (_BridgeFeeStation *BridgeFeeStationCallerSession) GetGasInfo(remoteChainID *big.Int) (IBridgeFeeStationGasInfo, error) {
	return _BridgeFeeStation.Contract.GetGasInfo(&_BridgeFeeStation.CallOpts, remoteChainID)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationCaller) GetTokenFee(opts *bind.CallOpts, remoteChainID *big.Int, token common.Address) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFeeRate     *big.Int
}, error) {
	var out []interface{}
	err := _BridgeFeeStation.contract.Call(opts, &out, "getTokenFee", remoteChainID, token)

	outstruct := new(struct {
		MinimumAmount *big.Int
		GasFee        *big.Int
		ExFeeRate     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFeeRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationSession) GetTokenFee(remoteChainID *big.Int, token common.Address) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFeeRate     *big.Int
}, error) {
	return _BridgeFeeStation.Contract.GetTokenFee(&_BridgeFeeStation.CallOpts, remoteChainID, token)
}

// GetTokenFee is a free data retrieval call binding the contract method 0x09585b1c.
//
// Solidity: function getTokenFee(uint256 remoteChainID, address token) view returns(uint256 minimumAmount, uint256 gasFee, uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationCallerSession) GetTokenFee(remoteChainID *big.Int, token common.Address) (struct {
	MinimumAmount *big.Int
	GasFee        *big.Int
	ExFeeRate     *big.Int
}, error) {
	return _BridgeFeeStation.Contract.GetTokenFee(&_BridgeFeeStation.CallOpts, remoteChainID, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeStation *BridgeFeeStationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeFeeStation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeStation *BridgeFeeStationSession) Owner() (common.Address, error) {
	return _BridgeFeeStation.Contract.Owner(&_BridgeFeeStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeFeeStation *BridgeFeeStationCallerSession) Owner() (common.Address, error) {
	return _BridgeFeeStation.Contract.Owner(&_BridgeFeeStation.CallOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) RemovePriceFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "removePriceFeed")
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeStation *BridgeFeeStationSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RemovePriceFeed(&_BridgeFeeStation.TransactOpts)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xce2e5c66.
//
// Solidity: function removePriceFeed() returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) RemovePriceFeed() (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RemovePriceFeed(&_BridgeFeeStation.TransactOpts)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) RemoveTokenFee(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "removeTokenFee", token)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) RemoveTokenFee(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RemoveTokenFee(&_BridgeFeeStation.TransactOpts, token)
}

// RemoveTokenFee is a paid mutator transaction binding the contract method 0x1b034c1c.
//
// Solidity: function removeTokenFee(address token) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) RemoveTokenFee(token common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RemoveTokenFee(&_BridgeFeeStation.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeStation *BridgeFeeStationSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RenounceOwnership(&_BridgeFeeStation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.RenounceOwnership(&_BridgeFeeStation.TransactOpts)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetExFeeRate(opts *bind.TransactOpts, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setExFeeRate", exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetExFeeRate(&_BridgeFeeStation.TransactOpts, exFeeRate)
}

// SetExFeeRate is a paid mutator transaction binding the contract method 0x005d0f10.
//
// Solidity: function setExFeeRate(uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetExFeeRate(exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetExFeeRate(&_BridgeFeeStation.TransactOpts, exFeeRate)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetFinalizeBridgeGas(opts *bind.TransactOpts, finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setFinalizeBridgeGas", finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetFinalizeBridgeGas(&_BridgeFeeStation.TransactOpts, finalizeBridgeGas)
}

// SetFinalizeBridgeGas is a paid mutator transaction binding the contract method 0xf289d3ba.
//
// Solidity: function setFinalizeBridgeGas(uint256 finalizeBridgeGas) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetFinalizeBridgeGas(finalizeBridgeGas *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetFinalizeBridgeGas(&_BridgeFeeStation.TransactOpts, finalizeBridgeGas)
}

// SetGasInfo is a paid mutator transaction binding the contract method 0x3f0cf10d.
//
// Solidity: function setGasInfo(uint256 remoteChainID, address gasToken, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetGasInfo(opts *bind.TransactOpts, remoteChainID *big.Int, gasToken common.Address, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setGasInfo", remoteChainID, gasToken, gasPrice)
}

// SetGasInfo is a paid mutator transaction binding the contract method 0x3f0cf10d.
//
// Solidity: function setGasInfo(uint256 remoteChainID, address gasToken, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetGasInfo(remoteChainID *big.Int, gasToken common.Address, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetGasInfo(&_BridgeFeeStation.TransactOpts, remoteChainID, gasToken, gasPrice)
}

// SetGasInfo is a paid mutator transaction binding the contract method 0x3f0cf10d.
//
// Solidity: function setGasInfo(uint256 remoteChainID, address gasToken, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetGasInfo(remoteChainID *big.Int, gasToken common.Address, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetGasInfo(&_BridgeFeeStation.TransactOpts, remoteChainID, gasToken, gasPrice)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetPriceFeed(&_BridgeFeeStation.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetPriceFeed(&_BridgeFeeStation.TransactOpts, priceFeed)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumAmount, uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, minimumAmount *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setTokenFee", token, minimumAmount, exFeeRate)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumAmount, uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetTokenFee(token common.Address, minimumAmount *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetTokenFee(&_BridgeFeeStation.TransactOpts, token, minimumAmount, exFeeRate)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0xdffa229e.
//
// Solidity: function setTokenFee(address token, uint256 minimumAmount, uint256 exFeeRate) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetTokenFee(token common.Address, minimumAmount *big.Int, exFeeRate *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetTokenFee(&_BridgeFeeStation.TransactOpts, token, minimumAmount, exFeeRate)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumAmountList, uint256[] exFeeRateList) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) SetTokenFeeBatch(opts *bind.TransactOpts, tokenList []common.Address, minimumAmountList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "setTokenFeeBatch", tokenList, minimumAmountList, exFeeRateList)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumAmountList, uint256[] exFeeRateList) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) SetTokenFeeBatch(tokenList []common.Address, minimumAmountList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetTokenFeeBatch(&_BridgeFeeStation.TransactOpts, tokenList, minimumAmountList, exFeeRateList)
}

// SetTokenFeeBatch is a paid mutator transaction binding the contract method 0x6731d860.
//
// Solidity: function setTokenFeeBatch(address[] tokenList, uint256[] minimumAmountList, uint256[] exFeeRateList) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) SetTokenFeeBatch(tokenList []common.Address, minimumAmountList []*big.Int, exFeeRateList []*big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.SetTokenFeeBatch(&_BridgeFeeStation.TransactOpts, tokenList, minimumAmountList, exFeeRateList)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.TransferOwnership(&_BridgeFeeStation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.TransferOwnership(&_BridgeFeeStation.TransactOpts, newOwner)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactor) UpdateGasPrice(opts *bind.TransactOpts, remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.contract.Transact(opts, "updateGasPrice", remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.UpdateGasPrice(&_BridgeFeeStation.TransactOpts, remoteChainID, gasPrice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x1f96131e.
//
// Solidity: function updateGasPrice(uint256 remoteChainID, uint256 gasPrice) returns()
func (_BridgeFeeStation *BridgeFeeStationTransactorSession) UpdateGasPrice(remoteChainID *big.Int, gasPrice *big.Int) (*types.Transaction, error) {
	return _BridgeFeeStation.Contract.UpdateGasPrice(&_BridgeFeeStation.TransactOpts, remoteChainID, gasPrice)
}

// BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator is returned from FilterBridgeFeeStationExchangeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeStationExchangeFeeUpdated events raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator struct {
	Event *BridgeFeeStationBridgeFeeStationExchangeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeStationBridgeFeeStationExchangeFeeUpdated)
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
		it.Event = new(BridgeFeeStationBridgeFeeStationExchangeFeeUpdated)
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
func (it *BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeStationBridgeFeeStationExchangeFeeUpdated represents a BridgeFeeStationExchangeFeeUpdated event raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationExchangeFeeUpdated struct {
	ExFeeRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeStationExchangeFeeUpdated is a free log retrieval operation binding the contract event 0x0312f690e5eab8afc0fc83c2e2558848b2debb98c72fe871573d8950080deb71.
//
// Solidity: event BridgeFeeStationExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationFilterer) FilterBridgeFeeStationExchangeFeeUpdated(opts *bind.FilterOpts) (*BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator, error) {

	logs, sub, err := _BridgeFeeStation.contract.FilterLogs(opts, "BridgeFeeStationExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationBridgeFeeStationExchangeFeeUpdatedIterator{contract: _BridgeFeeStation.contract, event: "BridgeFeeStationExchangeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeStationExchangeFeeUpdated is a free log subscription operation binding the contract event 0x0312f690e5eab8afc0fc83c2e2558848b2debb98c72fe871573d8950080deb71.
//
// Solidity: event BridgeFeeStationExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationFilterer) WatchBridgeFeeStationExchangeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeStationBridgeFeeStationExchangeFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeFeeStation.contract.WatchLogs(opts, "BridgeFeeStationExchangeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeStationBridgeFeeStationExchangeFeeUpdated)
				if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationExchangeFeeUpdated", log); err != nil {
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

// ParseBridgeFeeStationExchangeFeeUpdated is a log parse operation binding the contract event 0x0312f690e5eab8afc0fc83c2e2558848b2debb98c72fe871573d8950080deb71.
//
// Solidity: event BridgeFeeStationExchangeFeeUpdated(uint256 exFeeRate)
func (_BridgeFeeStation *BridgeFeeStationFilterer) ParseBridgeFeeStationExchangeFeeUpdated(log types.Log) (*BridgeFeeStationBridgeFeeStationExchangeFeeUpdated, error) {
	event := new(BridgeFeeStationBridgeFeeStationExchangeFeeUpdated)
	if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationExchangeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator is returned from FilterBridgeFeeStationFinalizeBridgeGasSet and is used to iterate over the raw logs and unpacked data for BridgeFeeStationFinalizeBridgeGasSet events raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator struct {
	Event *BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet // Event containing the contract specifics and raw log

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
func (it *BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet)
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
		it.Event = new(BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet)
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
func (it *BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet represents a BridgeFeeStationFinalizeBridgeGasSet event raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet struct {
	FinalizeBridgeGas *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeStationFinalizeBridgeGasSet is a free log retrieval operation binding the contract event 0x570de5cf44bf6056ffb0b1a5be86c62c41ef8c92f81427627f5eadce4cffca7a.
//
// Solidity: event BridgeFeeStationFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeStation *BridgeFeeStationFilterer) FilterBridgeFeeStationFinalizeBridgeGasSet(opts *bind.FilterOpts) (*BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator, error) {

	logs, sub, err := _BridgeFeeStation.contract.FilterLogs(opts, "BridgeFeeStationFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSetIterator{contract: _BridgeFeeStation.contract, event: "BridgeFeeStationFinalizeBridgeGasSet", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeStationFinalizeBridgeGasSet is a free log subscription operation binding the contract event 0x570de5cf44bf6056ffb0b1a5be86c62c41ef8c92f81427627f5eadce4cffca7a.
//
// Solidity: event BridgeFeeStationFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeStation *BridgeFeeStationFilterer) WatchBridgeFeeStationFinalizeBridgeGasSet(opts *bind.WatchOpts, sink chan<- *BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet) (event.Subscription, error) {

	logs, sub, err := _BridgeFeeStation.contract.WatchLogs(opts, "BridgeFeeStationFinalizeBridgeGasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet)
				if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationFinalizeBridgeGasSet", log); err != nil {
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

// ParseBridgeFeeStationFinalizeBridgeGasSet is a log parse operation binding the contract event 0x570de5cf44bf6056ffb0b1a5be86c62c41ef8c92f81427627f5eadce4cffca7a.
//
// Solidity: event BridgeFeeStationFinalizeBridgeGasSet(uint256 finalizeBridgeGas)
func (_BridgeFeeStation *BridgeFeeStationFilterer) ParseBridgeFeeStationFinalizeBridgeGasSet(log types.Log) (*BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet, error) {
	event := new(BridgeFeeStationBridgeFeeStationFinalizeBridgeGasSet)
	if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationFinalizeBridgeGasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator is returned from FilterBridgeFeeStationGasPriceUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeStationGasPriceUpdated events raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator struct {
	Event *BridgeFeeStationBridgeFeeStationGasPriceUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeStationBridgeFeeStationGasPriceUpdated)
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
		it.Event = new(BridgeFeeStationBridgeFeeStationGasPriceUpdated)
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
func (it *BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeStationBridgeFeeStationGasPriceUpdated represents a BridgeFeeStationGasPriceUpdated event raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationGasPriceUpdated struct {
	RemoteChainID *big.Int
	GasPrice      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeStationGasPriceUpdated is a free log retrieval operation binding the contract event 0x5c39f52fe1fce1c47c4175fc8ba058e40f59adc2571a3c282c4def80c5c12cbe.
//
// Solidity: event BridgeFeeStationGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeStation *BridgeFeeStationFilterer) FilterBridgeFeeStationGasPriceUpdated(opts *bind.FilterOpts, remoteChainID []*big.Int) (*BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.FilterLogs(opts, "BridgeFeeStationGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationBridgeFeeStationGasPriceUpdatedIterator{contract: _BridgeFeeStation.contract, event: "BridgeFeeStationGasPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeStationGasPriceUpdated is a free log subscription operation binding the contract event 0x5c39f52fe1fce1c47c4175fc8ba058e40f59adc2571a3c282c4def80c5c12cbe.
//
// Solidity: event BridgeFeeStationGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeStation *BridgeFeeStationFilterer) WatchBridgeFeeStationGasPriceUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeStationBridgeFeeStationGasPriceUpdated, remoteChainID []*big.Int) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.WatchLogs(opts, "BridgeFeeStationGasPriceUpdated", remoteChainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeStationBridgeFeeStationGasPriceUpdated)
				if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationGasPriceUpdated", log); err != nil {
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

// ParseBridgeFeeStationGasPriceUpdated is a log parse operation binding the contract event 0x5c39f52fe1fce1c47c4175fc8ba058e40f59adc2571a3c282c4def80c5c12cbe.
//
// Solidity: event BridgeFeeStationGasPriceUpdated(uint256 indexed remoteChainID, uint256 gasPrice)
func (_BridgeFeeStation *BridgeFeeStationFilterer) ParseBridgeFeeStationGasPriceUpdated(log types.Log) (*BridgeFeeStationBridgeFeeStationGasPriceUpdated, error) {
	event := new(BridgeFeeStationBridgeFeeStationGasPriceUpdated)
	if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationGasPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator is returned from FilterBridgeFeeStationPriceFeedUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeStationPriceFeedUpdated events raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator struct {
	Event *BridgeFeeStationBridgeFeeStationPriceFeedUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeStationBridgeFeeStationPriceFeedUpdated)
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
		it.Event = new(BridgeFeeStationBridgeFeeStationPriceFeedUpdated)
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
func (it *BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeStationBridgeFeeStationPriceFeedUpdated represents a BridgeFeeStationPriceFeedUpdated event raised by the BridgeFeeStation contract.
type BridgeFeeStationBridgeFeeStationPriceFeedUpdated struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeStationPriceFeedUpdated is a free log retrieval operation binding the contract event 0xcde95eaba5b290eeb4ecebcb3aa3a398f3098f54a7f754664d87876b10fe6a98.
//
// Solidity: event BridgeFeeStationPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeStation *BridgeFeeStationFilterer) FilterBridgeFeeStationPriceFeedUpdated(opts *bind.FilterOpts, priceFeed []common.Address) (*BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.FilterLogs(opts, "BridgeFeeStationPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationBridgeFeeStationPriceFeedUpdatedIterator{contract: _BridgeFeeStation.contract, event: "BridgeFeeStationPriceFeedUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeStationPriceFeedUpdated is a free log subscription operation binding the contract event 0xcde95eaba5b290eeb4ecebcb3aa3a398f3098f54a7f754664d87876b10fe6a98.
//
// Solidity: event BridgeFeeStationPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeStation *BridgeFeeStationFilterer) WatchBridgeFeeStationPriceFeedUpdated(opts *bind.WatchOpts, sink chan<- *BridgeFeeStationBridgeFeeStationPriceFeedUpdated, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.WatchLogs(opts, "BridgeFeeStationPriceFeedUpdated", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeStationBridgeFeeStationPriceFeedUpdated)
				if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationPriceFeedUpdated", log); err != nil {
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

// ParseBridgeFeeStationPriceFeedUpdated is a log parse operation binding the contract event 0xcde95eaba5b290eeb4ecebcb3aa3a398f3098f54a7f754664d87876b10fe6a98.
//
// Solidity: event BridgeFeeStationPriceFeedUpdated(address indexed priceFeed)
func (_BridgeFeeStation *BridgeFeeStationFilterer) ParseBridgeFeeStationPriceFeedUpdated(log types.Log) (*BridgeFeeStationBridgeFeeStationPriceFeedUpdated, error) {
	event := new(BridgeFeeStationBridgeFeeStationPriceFeedUpdated)
	if err := _BridgeFeeStation.contract.UnpackLog(event, "BridgeFeeStationPriceFeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeStationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeFeeStation contract.
type BridgeFeeStationOwnershipTransferredIterator struct {
	Event *BridgeFeeStationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeFeeStationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeStationOwnershipTransferred)
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
		it.Event = new(BridgeFeeStationOwnershipTransferred)
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
func (it *BridgeFeeStationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeStationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeStationOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeFeeStation contract.
type BridgeFeeStationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeFeeStation *BridgeFeeStationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeFeeStationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFeeStationOwnershipTransferredIterator{contract: _BridgeFeeStation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeFeeStation *BridgeFeeStationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeFeeStationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeFeeStation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeStationOwnershipTransferred)
				if err := _BridgeFeeStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeFeeStation *BridgeFeeStationFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeFeeStationOwnershipTransferred, error) {
	event := new(BridgeFeeStationOwnershipTransferred)
	if err := _BridgeFeeStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
