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

// PriceFeedCrossMetaData contains all meta data concerning the PriceFeedCross contract.
var PriceFeedCrossMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"changeDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"PriceFeedDeadlineChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"dee1f2af": "addToken(address,uint256,uint256)",
		"6ff97f1d": "allTokens()",
		"f30589c3": "allValidators()",
		"7aca97b5": "changeDeadline(uint256)",
		"b7f3358d": "changeThreshold(uint8)",
		"11df9995": "coin()",
		"5dbe47e8": "contains(address)",
		"29dcb0cf": "deadline()",
		"2fe3b6cf": "dollarDecimals()",
		"84b0196e": "eip712Domain()",
		"41976e09": "getPrice(address)",
		"8fb5a482": "getPrices(address[])",
		"33ddef07": "getValidPrice(address)",
		"8129fc1c": "initialize()",
		"facd743b": "isValidator(address)",
		"8da5cb5b": "owner()",
		"5fa7b584": "removeToken(address)",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"d92fc67b": "tokensLength()",
		"f2fde38b": "transferOwnership(address)",
		"e57d6fb7": "updatePrice(address,uint256,uint256)",
		"cf64f639": "updatePriceBatch(address[],uint256[],uint256[])",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50611c298061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106101d1575f3560e01c80638129fc1c116100fe578063cbae59581161009e578063e57d6fb71161006e578063e57d6fb714610418578063f2fde38b1461042b578063f30589c31461043e578063facd743b14610446575f5ffd5b8063cbae5958146103d7578063cf64f639146103ea578063d92fc67b146103fd578063dee1f2af14610405575f5ffd5b80638fb5a482116100d95780638fb5a482146103895780639300c926146103a9578063aed1d403146103bc578063b7f3358d146103c4575f5ffd5b80638129fc1c1461033657806384b0196e1461033e5780638da5cb5b14610359575f5ffd5b806341976e09116101745780635fa7b584116101445780635fa7b584146102f35780636ff97f1d14610306578063715018a61461031b5780637aca97b514610323575f5ffd5b806341976e091461028857806342cde4e8146102a85780634f6ccce7146102bd5780635dbe47e8146102d0575f5ffd5b806329dcb0cf116101af57806329dcb0cf1461022d5780632fe3b6cf1461024457806333ddef071461024d57806340a141ff14610275575f5ffd5b806311df9995146101d55780631327d3d8146102055780631d40f0d81461021a575b5f5ffd5b6066546101e8906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b610218610213366004611500565b610459565b005b6102186102283660046115f2565b610467565b61023660675481565b6040519081526020016101fc565b61023660685481565b61026061025b366004611500565b6104a1565b604080519283529015156020830152016101fc565b610218610283366004611500565b61054e565b61029b610296366004611500565b610558565b6040516101fc919061162c565b60355460405160ff90911681526020016101fc565b6101e86102cb366004611656565b6105fa565b6102e36102de366004611500565b61060b565b60405190151581526020016101fc565b610218610301366004611500565b610616565b61030e610658565b6040516101fc919061166d565b6102186106ba565b610218610331366004611656565b6106cd565b610218610711565b6103466108d0565b6040516101fc97969594939291906116e6565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101e8565b61039c6103973660046115f2565b610979565b6040516101fc919061177c565b6102186103b73660046115f2565b610a4b565b610236610a82565b6102186103d23660046117db565b610a92565b6101e86103e5366004611656565b610adc565b6102186103f8366004611856565b610ae8565b610236610b79565b6102186104133660046118e4565b610b83565b6102186104263660046118e4565b610bcb565b610218610439366004611500565b610c40565b61030e610c7a565b6102e3610454366004611500565b610c86565b610464816001610c92565b50565b5f5b815181101561049d5761049582828151811061048757610487611914565b60200260200101515f610c92565b600101610469565b5050565b5f5f6104ac8361060b565b83906104dc57604051630f2c982560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b506001600160a01b038084165f9081526069602090815260409182902082516060810184528154909416845260018101549184018290526002015491830182905293505f19148061053e5750606754816040015161053a919061193c565b4211155b1561054857600191505b50915091565b610464815f610c92565b61058260405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b61058b8261060b565b82906105b657604051630f2c982560e01b81526001600160a01b0390911660048201526024016104d3565b50506001600160a01b039081165f90815260696020908152604091829020825160608101845281549094168452600181015491840191909152600201549082015290565b5f6106058183610d5b565b92915050565b5f6106058183610d6d565b61061e610d8e565b61062781610de9565b6001600160a01b03165f90815260696020526040812080546001600160a01b03191681556001810182905560020155565b60605f6106645f610e7e565b90508067ffffffffffffffff81111561067f5761067f611519565b6040519080825280602002602001820160405280156106a8578160200160208202803683370190505b5091506106b45f610e87565b91505090565b6106c2610d8e565b6106cb5f610e93565b565b6106d5610d8e565b60678190556040518181527fca1d12600be991df47df34b798bb7437ed8c8422477a0ccda5bc128a78a13b83906020015b60405180910390a150565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156107565750825b90505f8267ffffffffffffffff1660011480156107725750303b155b905081158015610780575080155b1561079e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156107c857845460ff60401b1916600160401b1785555b6107d133610f03565b611c206067556006606855606680546001600160a01b03191660019081179091556107fb90610f6c565b60408051606081019091526066546001600160a01b0316815260685460208201906103e89061082b90600a611a32565b6108359190611a3d565b81525f196020918201526066546001600160a01b039081165f90815260698352604090819020845181546001600160a01b031916931692909217825591830151600182015591015160029091015583156108c957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050565b5f60608082808083815f516020611bd45f395f51905f5280549091501580156108fb57506001810154155b61093f5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016104d3565b610947611001565b61094f6110c1565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6060815167ffffffffffffffff81111561099557610995611519565b6040519080825280602002602001820160405280156109f057816020015b6109dd60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b8152602001906001900390816109b35790505b5090505f5b8251811015610a4557610a20838281518110610a1357610a13611914565b6020026020010151610558565b828281518110610a3257610a32611914565b60209081029190910101526001016109f5565b50919050565b5f5b815181101561049d57610a7a828281518110610a6b57610a6b611914565b60200260200101516001610c92565b600101610a4d565b5f610a8d6033610e7e565b905090565b610a9a610d8e565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff90602001610706565b5f610605603383610d5b565b8151835114610b0a5760405163282e5b7160e11b815260040160405180910390fd5b5f5b8351811015610b7357610b6b848281518110610b2a57610b2a611914565b6020026020010151848381518110610b4457610b44611914565b6020026020010151848481518110610b5e57610b5e611914565b6020026020010151610bcb565b600101610b0c565b50505050565b5f610a8d5f610e7e565b610b8b610d8e565b6001600160a01b038316610bb25760405163620785c160e11b81526004016104d390611a5c565b610bbb83610f6c565b610bc68383836110ff565b505050565b610bd433610c86565b3390610bff576040516338905e7160e11b81526001600160a01b0390911660048201526024016104d3565b50610c098361060b565b8390610c3457604051630f2c982560e01b81526001600160a01b0390911660048201526024016104d3565b50610bc68383836110ff565b610c48610d8e565b6001600160a01b038116610c7157604051631e4fbdf760e01b81525f60048201526024016104d3565b61046481610e93565b6060610a8d6033610e87565b5f610605603383610d6d565b610c9a610d8e565b8015610cdc57610cab60338361123c565b8290610cd65760405163082cdf5560e21b81526001600160a01b0390911660048201526024016104d3565b50610d14565b610ce7603383611250565b8290610d125760405163e491024560e01b81526001600160a01b0390911660048201526024016104d3565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b5f610d668383611264565b9392505050565b6001600160a01b0381165f9081526001830160205260408120541515610d66565b33610dc07f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106cb5760405163118cdaa760e01b81523360048201526024016104d3565b806001600160a01b038116610e11576040516330de3edf60e11b81526004016104d390611a5c565b610e1b5f83611250565b8290610e46576040516340da71e560e01b81526001600160a01b0390911660048201526024016104d3565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b5f610605825490565b60605f610d668361128a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610f0b6112e3565b610f148161132c565b610f5c604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b81525061133d565b506035805460ff19166003179055565b806001600160a01b038116610f94576040516330de3edf60e11b81526004016104d390611a5c565b610f9e5f8361123c565b8290610fc9576040516351eccfe160e11b81526001600160a01b0390911660048201526024016104d3565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020611bd45f395f51905f529161103f90611a7b565b80601f016020809104026020016040519081016040528092919081815260200182805461106b90611a7b565b80156110b65780601f1061108d576101008083540402835291602001916110b6565b820191905f5260205f20905b81548152906001019060200180831161109957829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020611bd45f395f51905f529161103f90611a7b565b6111088361060b565b839061113357604051630f2c982560e01b81526001600160a01b0390911660048201526024016104d3565b50815f0361116c57604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b60448201526064016104d3565b5f1981148061117b5750428111155b814290916111a55760405163248fc98560e11b8152600481019290925260248201526044016104d3565b5050604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260698452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2910160405180910390a2505050565b5f610d66836001600160a01b03841661134f565b5f610d66836001600160a01b03841661139b565b5f825f01828154811061127957611279611914565b905f5260205f200154905092915050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156112d757602002820191905f5260205f20905b8154815260200190600101908083116112c3575b50505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166106cb57604051631afcd79f60e31b815260040160405180910390fd5b6113346112e3565b6104648161147e565b6113456112e3565b61049d8282611486565b5f81815260018301602052604081205461139457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610605565b505f610605565b5f8181526001830160205260408120548015611475575f6113bd600183611aad565b85549091505f906113d090600190611aad565b905080821461142f575f865f0182815481106113ee576113ee611914565b905f5260205f200154905080875f01848154811061140e5761140e611914565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061144057611440611ac0565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610605565b5f915050610605565b610c486112e3565b61148e6112e3565b5f516020611bd45f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026114c78482611b18565b50600381016114d68382611b18565b505f8082556001909101555050565b80356001600160a01b03811681146114fb575f5ffd5b919050565b5f60208284031215611510575f5ffd5b610d66826114e5565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561155657611556611519565b604052919050565b5f67ffffffffffffffff82111561157757611577611519565b5060051b60200190565b5f82601f830112611590575f5ffd5b81356115a361159e8261155e565b61152d565b8082825260208201915060208360051b8601019250858311156115c4575f5ffd5b602085015b838110156115e8576115da816114e5565b8352602092830192016115c9565b5095945050505050565b5f60208284031215611602575f5ffd5b813567ffffffffffffffff811115611618575f5ffd5b61162484828501611581565b949350505050565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610605565b5f60208284031215611666575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b818110156116ad5783516001600160a01b0316835260209384019390920191600101611686565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f61170460e08301896116b8565b828103604084015261171681896116b8565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561176b57835183526020938401939092019160010161174d565b50909b9a5050505050505050505050565b602080825282518282018190525f918401906040840190835b818110156116ad576117c583855180516001600160a01b0316825260208082015190830152604090810151910152565b6020939093019260609290920191600101611795565b5f602082840312156117eb575f5ffd5b813560ff81168114610d66575f5ffd5b5f82601f83011261180a575f5ffd5b813561181861159e8261155e565b8082825260208201915060208360051b860101925085831115611839575f5ffd5b602085015b838110156115e857803583526020928301920161183e565b5f5f5f60608486031215611868575f5ffd5b833567ffffffffffffffff81111561187e575f5ffd5b61188a86828701611581565b935050602084013567ffffffffffffffff8111156118a6575f5ffd5b6118b2868287016117fb565b925050604084013567ffffffffffffffff8111156118ce575f5ffd5b6118da868287016117fb565b9150509250925092565b5f5f5f606084860312156118f6575f5ffd5b6118ff846114e5565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561060557610605611928565b6001815b600184111561198a5780850481111561196e5761196e611928565b600184161561197c57908102905b60019390931c928002611953565b935093915050565b5f826119a057506001610605565b816119ac57505f610605565b81600181146119c257600281146119cc576119e8565b6001915050610605565b60ff8411156119dd576119dd611928565b50506001821b610605565b5060208310610133831016604e8410600b8410161715611a0b575081810a610605565b611a175f19848461194f565b805f1904821115611a2a57611a2a611928565b029392505050565b5f610d668383611992565b5f82611a5757634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252600590820152643a37b5b2b760d91b604082015260600190565b600181811c90821680611a8f57607f821691505b602082108103610a4557634e487b7160e01b5f52602260045260245ffd5b8181038181111561060557610605611928565b634e487b7160e01b5f52603160045260245ffd5b601f821115610bc657805f5260205f20601f840160051c81016020851015611af95750805b601f840160051c820191505b818110156108c9575f8155600101611b05565b815167ffffffffffffffff811115611b3257611b32611519565b611b4681611b408454611a7b565b84611ad4565b6020601f821160018114611b78575f8315611b615750848201515b5f19600385901b1c1916600184901b1784556108c9565b5f84815260208120601f198516915b82811015611ba75787850151825560209485019460019092019101611b87565b5084821015611bc457868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100a26469706673582212206e69dc447e4473affb88eb7b9ead46cfeb80e1234c8492c69225a0f2fe86eb3664736f6c634300081c0033",
}

// PriceFeedCrossABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedCrossMetaData.ABI instead.
var PriceFeedCrossABI = PriceFeedCrossMetaData.ABI

// Deprecated: Use PriceFeedCrossMetaData.Sigs instead.
// PriceFeedCrossFuncSigs maps the 4-byte function signature to its string representation.
var PriceFeedCrossFuncSigs = PriceFeedCrossMetaData.Sigs

// PriceFeedCrossBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedCrossMetaData.Bin instead.
var PriceFeedCrossBin = PriceFeedCrossMetaData.Bin

// DeployPriceFeedCross deploys a new Ethereum contract, binding an instance of PriceFeedCross to it.
func DeployPriceFeedCross(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeedCross, error) {
	parsed, err := PriceFeedCrossMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedCrossBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeedCross{PriceFeedCrossCaller: PriceFeedCrossCaller{contract: contract}, PriceFeedCrossTransactor: PriceFeedCrossTransactor{contract: contract}, PriceFeedCrossFilterer: PriceFeedCrossFilterer{contract: contract}}, nil
}

// PriceFeedCross is an auto generated Go binding around an Ethereum contract.
type PriceFeedCross struct {
	PriceFeedCrossCaller     // Read-only binding to the contract
	PriceFeedCrossTransactor // Write-only binding to the contract
	PriceFeedCrossFilterer   // Log filterer for contract events
}

// PriceFeedCrossCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedCrossCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedCrossTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedCrossTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedCrossFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedCrossFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedCrossSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedCrossSession struct {
	Contract     *PriceFeedCross   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedCrossCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedCrossCallerSession struct {
	Contract *PriceFeedCrossCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PriceFeedCrossTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedCrossTransactorSession struct {
	Contract     *PriceFeedCrossTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PriceFeedCrossRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedCrossRaw struct {
	Contract *PriceFeedCross // Generic contract binding to access the raw methods on
}

// PriceFeedCrossCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedCrossCallerRaw struct {
	Contract *PriceFeedCrossCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedCrossTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedCrossTransactorRaw struct {
	Contract *PriceFeedCrossTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedCross creates a new instance of PriceFeedCross, bound to a specific deployed contract.
func NewPriceFeedCross(address common.Address, backend bind.ContractBackend) (*PriceFeedCross, error) {
	contract, err := bindPriceFeedCross(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCross{PriceFeedCrossCaller: PriceFeedCrossCaller{contract: contract}, PriceFeedCrossTransactor: PriceFeedCrossTransactor{contract: contract}, PriceFeedCrossFilterer: PriceFeedCrossFilterer{contract: contract}}, nil
}

// NewPriceFeedCrossCaller creates a new read-only instance of PriceFeedCross, bound to a specific deployed contract.
func NewPriceFeedCrossCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedCrossCaller, error) {
	contract, err := bindPriceFeedCross(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossCaller{contract: contract}, nil
}

// NewPriceFeedCrossTransactor creates a new write-only instance of PriceFeedCross, bound to a specific deployed contract.
func NewPriceFeedCrossTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedCrossTransactor, error) {
	contract, err := bindPriceFeedCross(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossTransactor{contract: contract}, nil
}

// NewPriceFeedCrossFilterer creates a new log filterer instance of PriceFeedCross, bound to a specific deployed contract.
func NewPriceFeedCrossFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedCrossFilterer, error) {
	contract, err := bindPriceFeedCross(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossFilterer{contract: contract}, nil
}

// bindPriceFeedCross binds a generic wrapper to an already deployed contract.
func bindPriceFeedCross(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedCrossMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedCross *PriceFeedCrossRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedCross.Contract.PriceFeedCrossCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedCross *PriceFeedCrossRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.PriceFeedCrossTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedCross *PriceFeedCrossRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.PriceFeedCrossTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedCross *PriceFeedCrossCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedCross.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedCross *PriceFeedCrossTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedCross *PriceFeedCrossTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.contract.Transact(opts, method, params...)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_PriceFeedCross *PriceFeedCrossCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_PriceFeedCross *PriceFeedCrossSession) AllTokens() ([]common.Address, error) {
	return _PriceFeedCross.Contract.AllTokens(&_PriceFeedCross.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[] tokens)
func (_PriceFeedCross *PriceFeedCrossCallerSession) AllTokens() ([]common.Address, error) {
	return _PriceFeedCross.Contract.AllTokens(&_PriceFeedCross.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeedCross *PriceFeedCrossCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeedCross *PriceFeedCrossSession) AllValidators() ([]common.Address, error) {
	return _PriceFeedCross.Contract.AllValidators(&_PriceFeedCross.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeedCross *PriceFeedCrossCallerSession) AllValidators() ([]common.Address, error) {
	return _PriceFeedCross.Contract.AllValidators(&_PriceFeedCross.CallOpts)
}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_PriceFeedCross *PriceFeedCrossCaller) Coin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "coin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_PriceFeedCross *PriceFeedCrossSession) Coin() (common.Address, error) {
	return _PriceFeedCross.Contract.Coin(&_PriceFeedCross.CallOpts)
}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Coin() (common.Address, error) {
	return _PriceFeedCross.Contract.Coin(&_PriceFeedCross.CallOpts)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossCaller) Contains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "contains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossSession) Contains(token common.Address) (bool, error) {
	return _PriceFeedCross.Contract.Contains(&_PriceFeedCross.CallOpts, token)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address token) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Contains(token common.Address) (bool, error) {
	return _PriceFeedCross.Contract.Contains(&_PriceFeedCross.CallOpts, token)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossSession) Deadline() (*big.Int, error) {
	return _PriceFeedCross.Contract.Deadline(&_PriceFeedCross.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Deadline() (*big.Int, error) {
	return _PriceFeedCross.Contract.Deadline(&_PriceFeedCross.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCaller) DollarDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "dollarDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossSession) DollarDecimals() (*big.Int, error) {
	return _PriceFeedCross.Contract.DollarDecimals(&_PriceFeedCross.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCallerSession) DollarDecimals() (*big.Int, error) {
	return _PriceFeedCross.Contract.DollarDecimals(&_PriceFeedCross.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PriceFeedCross *PriceFeedCrossCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PriceFeedCross *PriceFeedCrossSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PriceFeedCross.Contract.Eip712Domain(&_PriceFeedCross.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PriceFeedCross.Contract.Eip712Domain(&_PriceFeedCross.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns((address,uint256,uint256) data)
func (_PriceFeedCross *PriceFeedCrossCaller) GetPrice(opts *bind.CallOpts, token common.Address) (IPriceFeedPriceData, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(IPriceFeedPriceData), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceFeedPriceData)).(*IPriceFeedPriceData)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns((address,uint256,uint256) data)
func (_PriceFeedCross *PriceFeedCrossSession) GetPrice(token common.Address) (IPriceFeedPriceData, error) {
	return _PriceFeedCross.Contract.GetPrice(&_PriceFeedCross.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns((address,uint256,uint256) data)
func (_PriceFeedCross *PriceFeedCrossCallerSession) GetPrice(token common.Address) (IPriceFeedPriceData, error) {
	return _PriceFeedCross.Contract.GetPrice(&_PriceFeedCross.CallOpts, token)
}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] token) view returns((address,uint256,uint256)[] data)
func (_PriceFeedCross *PriceFeedCrossCaller) GetPrices(opts *bind.CallOpts, token []common.Address) ([]IPriceFeedPriceData, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "getPrices", token)

	if err != nil {
		return *new([]IPriceFeedPriceData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPriceFeedPriceData)).(*[]IPriceFeedPriceData)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] token) view returns((address,uint256,uint256)[] data)
func (_PriceFeedCross *PriceFeedCrossSession) GetPrices(token []common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedCross.Contract.GetPrices(&_PriceFeedCross.CallOpts, token)
}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] token) view returns((address,uint256,uint256)[] data)
func (_PriceFeedCross *PriceFeedCrossCallerSession) GetPrices(token []common.Address) ([]IPriceFeedPriceData, error) {
	return _PriceFeedCross.Contract.GetPrices(&_PriceFeedCross.CallOpts, token)
}

// GetValidPrice is a free data retrieval call binding the contract method 0x33ddef07.
//
// Solidity: function getValidPrice(address token) view returns(uint256 price, bool isValid)
func (_PriceFeedCross *PriceFeedCrossCaller) GetValidPrice(opts *bind.CallOpts, token common.Address) (struct {
	Price   *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "getValidPrice", token)

	outstruct := new(struct {
		Price   *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidPrice is a free data retrieval call binding the contract method 0x33ddef07.
//
// Solidity: function getValidPrice(address token) view returns(uint256 price, bool isValid)
func (_PriceFeedCross *PriceFeedCrossSession) GetValidPrice(token common.Address) (struct {
	Price   *big.Int
	IsValid bool
}, error) {
	return _PriceFeedCross.Contract.GetValidPrice(&_PriceFeedCross.CallOpts, token)
}

// GetValidPrice is a free data retrieval call binding the contract method 0x33ddef07.
//
// Solidity: function getValidPrice(address token) view returns(uint256 price, bool isValid)
func (_PriceFeedCross *PriceFeedCrossCallerSession) GetValidPrice(token common.Address) (struct {
	Price   *big.Int
	IsValid bool
}, error) {
	return _PriceFeedCross.Contract.GetValidPrice(&_PriceFeedCross.CallOpts, token)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossSession) IsValidator(validator common.Address) (bool, error) {
	return _PriceFeedCross.Contract.IsValidator(&_PriceFeedCross.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeedCross *PriceFeedCrossCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _PriceFeedCross.Contract.IsValidator(&_PriceFeedCross.CallOpts, validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedCross *PriceFeedCrossCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedCross *PriceFeedCrossSession) Owner() (common.Address, error) {
	return _PriceFeedCross.Contract.Owner(&_PriceFeedCross.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Owner() (common.Address, error) {
	return _PriceFeedCross.Contract.Owner(&_PriceFeedCross.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeedCross *PriceFeedCrossCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeedCross *PriceFeedCrossSession) Threshold() (uint8, error) {
	return _PriceFeedCross.Contract.Threshold(&_PriceFeedCross.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeedCross *PriceFeedCrossCallerSession) Threshold() (uint8, error) {
	return _PriceFeedCross.Contract.Threshold(&_PriceFeedCross.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_PriceFeedCross *PriceFeedCrossCaller) TokenByIndex(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "tokenByIndex", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_PriceFeedCross *PriceFeedCrossSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _PriceFeedCross.Contract.TokenByIndex(&_PriceFeedCross.CallOpts, i)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 i) view returns(address)
func (_PriceFeedCross *PriceFeedCrossCallerSession) TokenByIndex(i *big.Int) (common.Address, error) {
	return _PriceFeedCross.Contract.TokenByIndex(&_PriceFeedCross.CallOpts, i)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossSession) TokensLength() (*big.Int, error) {
	return _PriceFeedCross.Contract.TokensLength(&_PriceFeedCross.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCallerSession) TokensLength() (*big.Int, error) {
	return _PriceFeedCross.Contract.TokensLength(&_PriceFeedCross.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeedCross *PriceFeedCrossCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeedCross *PriceFeedCrossSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _PriceFeedCross.Contract.ValidatorByIndex(&_PriceFeedCross.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeedCross *PriceFeedCrossCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _PriceFeedCross.Contract.ValidatorByIndex(&_PriceFeedCross.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossSession) ValidatorLength() (*big.Int, error) {
	return _PriceFeedCross.Contract.ValidatorLength(&_PriceFeedCross.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeedCross *PriceFeedCrossCallerSession) ValidatorLength() (*big.Int, error) {
	return _PriceFeedCross.Contract.ValidatorLength(&_PriceFeedCross.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) AddToken(opts *bind.TransactOpts, token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "addToken", token, price, priceAt)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossSession) AddToken(token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.AddToken(&_PriceFeedCross.TransactOpts, token, price, priceAt)
}

// AddToken is a paid mutator transaction binding the contract method 0xdee1f2af.
//
// Solidity: function addToken(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) AddToken(token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.AddToken(&_PriceFeedCross.TransactOpts, token, price, priceAt)
}

// ChangeDeadline is a paid mutator transaction binding the contract method 0x7aca97b5.
//
// Solidity: function changeDeadline(uint256 _deadline) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) ChangeDeadline(opts *bind.TransactOpts, _deadline *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "changeDeadline", _deadline)
}

// ChangeDeadline is a paid mutator transaction binding the contract method 0x7aca97b5.
//
// Solidity: function changeDeadline(uint256 _deadline) returns()
func (_PriceFeedCross *PriceFeedCrossSession) ChangeDeadline(_deadline *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ChangeDeadline(&_PriceFeedCross.TransactOpts, _deadline)
}

// ChangeDeadline is a paid mutator transaction binding the contract method 0x7aca97b5.
//
// Solidity: function changeDeadline(uint256 _deadline) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) ChangeDeadline(_deadline *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ChangeDeadline(&_PriceFeedCross.TransactOpts, _deadline)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeedCross *PriceFeedCrossSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ChangeThreshold(&_PriceFeedCross.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ChangeThreshold(&_PriceFeedCross.TransactOpts, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedCross *PriceFeedCrossSession) Initialize() (*types.Transaction, error) {
	return _PriceFeedCross.Contract.Initialize(&_PriceFeedCross.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) Initialize() (*types.Transaction, error) {
	return _PriceFeedCross.Contract.Initialize(&_PriceFeedCross.TransactOpts)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_PriceFeedCross *PriceFeedCrossSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveToken(&_PriceFeedCross.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveToken(&_PriceFeedCross.TransactOpts, token)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveValidator(&_PriceFeedCross.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveValidator(&_PriceFeedCross.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveValidators(&_PriceFeedCross.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RemoveValidators(&_PriceFeedCross.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedCross *PriceFeedCrossSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RenounceOwnership(&_PriceFeedCross.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeedCross.Contract.RenounceOwnership(&_PriceFeedCross.TransactOpts)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.SetValidator(&_PriceFeedCross.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.SetValidator(&_PriceFeedCross.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.SetValidators(&_PriceFeedCross.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.SetValidators(&_PriceFeedCross.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedCross *PriceFeedCrossSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.TransferOwnership(&_PriceFeedCross.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.TransferOwnership(&_PriceFeedCross.TransactOpts, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe57d6fb7.
//
// Solidity: function updatePrice(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) UpdatePrice(opts *bind.TransactOpts, token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "updatePrice", token, price, priceAt)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe57d6fb7.
//
// Solidity: function updatePrice(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossSession) UpdatePrice(token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.UpdatePrice(&_PriceFeedCross.TransactOpts, token, price, priceAt)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0xe57d6fb7.
//
// Solidity: function updatePrice(address token, uint256 price, uint256 priceAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) UpdatePrice(token common.Address, price *big.Int, priceAt *big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.UpdatePrice(&_PriceFeedCross.TransactOpts, token, price, priceAt)
}

// UpdatePriceBatch is a paid mutator transaction binding the contract method 0xcf64f639.
//
// Solidity: function updatePriceBatch(address[] tokens, uint256[] _prices, uint256[] pricesAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) UpdatePriceBatch(opts *bind.TransactOpts, tokens []common.Address, _prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "updatePriceBatch", tokens, _prices, pricesAt)
}

// UpdatePriceBatch is a paid mutator transaction binding the contract method 0xcf64f639.
//
// Solidity: function updatePriceBatch(address[] tokens, uint256[] _prices, uint256[] pricesAt) returns()
func (_PriceFeedCross *PriceFeedCrossSession) UpdatePriceBatch(tokens []common.Address, _prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.UpdatePriceBatch(&_PriceFeedCross.TransactOpts, tokens, _prices, pricesAt)
}

// UpdatePriceBatch is a paid mutator transaction binding the contract method 0xcf64f639.
//
// Solidity: function updatePriceBatch(address[] tokens, uint256[] _prices, uint256[] pricesAt) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) UpdatePriceBatch(tokens []common.Address, _prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.UpdatePriceBatch(&_PriceFeedCross.TransactOpts, tokens, _prices, pricesAt)
}

// PriceFeedCrossEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the PriceFeedCross contract.
type PriceFeedCrossEIP712DomainChangedIterator struct {
	Event *PriceFeedCrossEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossEIP712DomainChanged)
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
		it.Event = new(PriceFeedCrossEIP712DomainChanged)
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
func (it *PriceFeedCrossEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossEIP712DomainChanged represents a EIP712DomainChanged event raised by the PriceFeedCross contract.
type PriceFeedCrossEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PriceFeedCrossEIP712DomainChangedIterator, error) {

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossEIP712DomainChangedIterator{contract: _PriceFeedCross.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossEIP712DomainChanged)
				if err := _PriceFeedCross.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseEIP712DomainChanged(log types.Log) (*PriceFeedCrossEIP712DomainChanged, error) {
	event := new(PriceFeedCrossEIP712DomainChanged)
	if err := _PriceFeedCross.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PriceFeedCross contract.
type PriceFeedCrossInitializedIterator struct {
	Event *PriceFeedCrossInitialized // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossInitialized)
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
		it.Event = new(PriceFeedCrossInitialized)
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
func (it *PriceFeedCrossInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossInitialized represents a Initialized event raised by the PriceFeedCross contract.
type PriceFeedCrossInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterInitialized(opts *bind.FilterOpts) (*PriceFeedCrossInitializedIterator, error) {

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossInitializedIterator{contract: _PriceFeedCross.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossInitialized) (event.Subscription, error) {

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossInitialized)
				if err := _PriceFeedCross.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseInitialized(log types.Log) (*PriceFeedCrossInitialized, error) {
	event := new(PriceFeedCrossInitialized)
	if err := _PriceFeedCross.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeedCross contract.
type PriceFeedCrossOwnershipTransferredIterator struct {
	Event *PriceFeedCrossOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossOwnershipTransferred)
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
		it.Event = new(PriceFeedCrossOwnershipTransferred)
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
func (it *PriceFeedCrossOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeedCross contract.
type PriceFeedCrossOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeedCrossOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossOwnershipTransferredIterator{contract: _PriceFeedCross.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossOwnershipTransferred)
				if err := _PriceFeedCross.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeedCrossOwnershipTransferred, error) {
	event := new(PriceFeedCrossOwnershipTransferred)
	if err := _PriceFeedCross.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossPriceFeedDeadlineChangedIterator is returned from FilterPriceFeedDeadlineChanged and is used to iterate over the raw logs and unpacked data for PriceFeedDeadlineChanged events raised by the PriceFeedCross contract.
type PriceFeedCrossPriceFeedDeadlineChangedIterator struct {
	Event *PriceFeedCrossPriceFeedDeadlineChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossPriceFeedDeadlineChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossPriceFeedDeadlineChanged)
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
		it.Event = new(PriceFeedCrossPriceFeedDeadlineChanged)
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
func (it *PriceFeedCrossPriceFeedDeadlineChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossPriceFeedDeadlineChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossPriceFeedDeadlineChanged represents a PriceFeedDeadlineChanged event raised by the PriceFeedCross contract.
type PriceFeedCrossPriceFeedDeadlineChanged struct {
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedDeadlineChanged is a free log retrieval operation binding the contract event 0xca1d12600be991df47df34b798bb7437ed8c8422477a0ccda5bc128a78a13b83.
//
// Solidity: event PriceFeedDeadlineChanged(uint256 deadline)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterPriceFeedDeadlineChanged(opts *bind.FilterOpts) (*PriceFeedCrossPriceFeedDeadlineChangedIterator, error) {

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "PriceFeedDeadlineChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossPriceFeedDeadlineChangedIterator{contract: _PriceFeedCross.contract, event: "PriceFeedDeadlineChanged", logs: logs, sub: sub}, nil
}

// WatchPriceFeedDeadlineChanged is a free log subscription operation binding the contract event 0xca1d12600be991df47df34b798bb7437ed8c8422477a0ccda5bc128a78a13b83.
//
// Solidity: event PriceFeedDeadlineChanged(uint256 deadline)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchPriceFeedDeadlineChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossPriceFeedDeadlineChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "PriceFeedDeadlineChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossPriceFeedDeadlineChanged)
				if err := _PriceFeedCross.contract.UnpackLog(event, "PriceFeedDeadlineChanged", log); err != nil {
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

// ParsePriceFeedDeadlineChanged is a log parse operation binding the contract event 0xca1d12600be991df47df34b798bb7437ed8c8422477a0ccda5bc128a78a13b83.
//
// Solidity: event PriceFeedDeadlineChanged(uint256 deadline)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParsePriceFeedDeadlineChanged(log types.Log) (*PriceFeedCrossPriceFeedDeadlineChanged, error) {
	event := new(PriceFeedCrossPriceFeedDeadlineChanged)
	if err := _PriceFeedCross.contract.UnpackLog(event, "PriceFeedDeadlineChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossPriceFeedPriceUpdatedIterator is returned from FilterPriceFeedPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedPriceUpdated events raised by the PriceFeedCross contract.
type PriceFeedCrossPriceFeedPriceUpdatedIterator struct {
	Event *PriceFeedCrossPriceFeedPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossPriceFeedPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossPriceFeedPriceUpdated)
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
		it.Event = new(PriceFeedCrossPriceFeedPriceUpdated)
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
func (it *PriceFeedCrossPriceFeedPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossPriceFeedPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossPriceFeedPriceUpdated represents a PriceFeedPriceUpdated event raised by the PriceFeedCross contract.
type PriceFeedCrossPriceFeedPriceUpdated struct {
	Token     common.Address
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedPriceUpdated is a free log retrieval operation binding the contract event 0xc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2.
//
// Solidity: event PriceFeedPriceUpdated(address indexed token, uint256 price, uint256 timestamp)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterPriceFeedPriceUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceFeedCrossPriceFeedPriceUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "PriceFeedPriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossPriceFeedPriceUpdatedIterator{contract: _PriceFeedCross.contract, event: "PriceFeedPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedPriceUpdated is a free log subscription operation binding the contract event 0xc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2.
//
// Solidity: event PriceFeedPriceUpdated(address indexed token, uint256 price, uint256 timestamp)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchPriceFeedPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossPriceFeedPriceUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "PriceFeedPriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossPriceFeedPriceUpdated)
				if err := _PriceFeedCross.contract.UnpackLog(event, "PriceFeedPriceUpdated", log); err != nil {
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

// ParsePriceFeedPriceUpdated is a log parse operation binding the contract event 0xc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2.
//
// Solidity: event PriceFeedPriceUpdated(address indexed token, uint256 price, uint256 timestamp)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParsePriceFeedPriceUpdated(log types.Log) (*PriceFeedCrossPriceFeedPriceUpdated, error) {
	event := new(PriceFeedCrossPriceFeedPriceUpdated)
	if err := _PriceFeedCross.contract.UnpackLog(event, "PriceFeedPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the PriceFeedCross contract.
type PriceFeedCrossThresholdChangedIterator struct {
	Event *PriceFeedCrossThresholdChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossThresholdChanged)
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
		it.Event = new(PriceFeedCrossThresholdChanged)
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
func (it *PriceFeedCrossThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossThresholdChanged represents a ThresholdChanged event raised by the PriceFeedCross contract.
type PriceFeedCrossThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*PriceFeedCrossThresholdChangedIterator, error) {

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossThresholdChangedIterator{contract: _PriceFeedCross.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossThresholdChanged)
				if err := _PriceFeedCross.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseThresholdChanged(log types.Log) (*PriceFeedCrossThresholdChanged, error) {
	event := new(PriceFeedCrossThresholdChanged)
	if err := _PriceFeedCross.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the PriceFeedCross contract.
type PriceFeedCrossTokenAddedIterator struct {
	Event *PriceFeedCrossTokenAdded // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossTokenAdded)
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
		it.Event = new(PriceFeedCrossTokenAdded)
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
func (it *PriceFeedCrossTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossTokenAdded represents a TokenAdded event raised by the PriceFeedCross contract.
type PriceFeedCrossTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*PriceFeedCrossTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossTokenAddedIterator{contract: _PriceFeedCross.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossTokenAdded)
				if err := _PriceFeedCross.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseTokenAdded(log types.Log) (*PriceFeedCrossTokenAdded, error) {
	event := new(PriceFeedCrossTokenAdded)
	if err := _PriceFeedCross.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the PriceFeedCross contract.
type PriceFeedCrossTokenRemovedIterator struct {
	Event *PriceFeedCrossTokenRemoved // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossTokenRemoved)
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
		it.Event = new(PriceFeedCrossTokenRemoved)
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
func (it *PriceFeedCrossTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossTokenRemoved represents a TokenRemoved event raised by the PriceFeedCross contract.
type PriceFeedCrossTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*PriceFeedCrossTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossTokenRemovedIterator{contract: _PriceFeedCross.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossTokenRemoved)
				if err := _PriceFeedCross.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseTokenRemoved(log types.Log) (*PriceFeedCrossTokenRemoved, error) {
	event := new(PriceFeedCrossTokenRemoved)
	if err := _PriceFeedCross.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedCrossValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the PriceFeedCross contract.
type PriceFeedCrossValidatorSetIterator struct {
	Event *PriceFeedCrossValidatorSet // Event containing the contract specifics and raw log

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
func (it *PriceFeedCrossValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedCrossValidatorSet)
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
		it.Event = new(PriceFeedCrossValidatorSet)
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
func (it *PriceFeedCrossValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedCrossValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedCrossValidatorSet represents a ValidatorSet event raised by the PriceFeedCross contract.
type PriceFeedCrossValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_PriceFeedCross *PriceFeedCrossFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*PriceFeedCrossValidatorSetIterator, error) {

	logs, sub, err := _PriceFeedCross.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &PriceFeedCrossValidatorSetIterator{contract: _PriceFeedCross.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_PriceFeedCross *PriceFeedCrossFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *PriceFeedCrossValidatorSet) (event.Subscription, error) {

	logs, sub, err := _PriceFeedCross.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedCrossValidatorSet)
				if err := _PriceFeedCross.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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

// ParseValidatorSet is a log parse operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_PriceFeedCross *PriceFeedCrossFilterer) ParseValidatorSet(log types.Log) (*PriceFeedCrossValidatorSet, error) {
	event := new(PriceFeedCrossValidatorSet)
	if err := _PriceFeedCross.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
