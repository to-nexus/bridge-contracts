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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"changeDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePriceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"PriceFeedDeadlineChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"PriceFeedNotValidPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageAleadyExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenStorageCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenStorageNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorManagerNotValidator\",\"type\":\"error\"}]",
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
		"4351e6b6": "initialize(uint8)",
		"facd743b": "isValidator(address)",
		"8da5cb5b": "owner()",
		"5fa7b584": "removeToken(address)",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
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
	Bin: "0x6080604052348015600e575f5ffd5b50611cfa8061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106101dc575f3560e01c8063715018a611610109578063cbae59581161009e578063e57d6fb71161006e578063e57d6fb714610420578063f2fde38b14610433578063f30589c314610446578063facd743b1461044e575f5ffd5b8063cbae5958146103df578063cf64f639146103f2578063d92fc67b14610405578063dee1f2af1461040d575f5ffd5b80638fb5a482116100d95780638fb5a482146103915780639300c926146103b1578063aed1d403146103c4578063b7f3358d146103cc575f5ffd5b8063715018a61461032b5780637aca97b51461033357806384b0196e146103465780638da5cb5b14610361575f5ffd5b806341976e091161017f5780635dbe47e81161014f5780635dbe47e8146102cd5780635fa7b584146102f05780636ff97f1d146103035780637101fcd314610318575f5ffd5b806341976e091461027257806342cde4e8146102925780634351e6b6146102a75780634f6ccce7146102ba575f5ffd5b806329dcb0cf116101ba57806329dcb0cf1461022d5780632fe3b6cf1461024457806333ddef071461024c57806340a141ff1461025f575f5ffd5b806311df9995146101e05780631327d3d8146102055780631d40f0d81461021a575b5f5ffd5b6101e8600181565b6040516001600160a01b0390911681526020015b60405180910390f35b6102186102133660046115ca565b610461565b005b6102186102283660046116bc565b61046f565b61023660665481565b6040519081526020016101fc565b610236600681565b61023661025a3660046115ca565b6104a9565b61021861026d3660046115ca565b610593565b6102856102803660046115ca565b61059d565b6040516101fc91906116f6565b60355460405160ff90911681526020016101fc565b6102186102b5366004611720565b61063f565b6101e86102c8366004611740565b610827565b6102e06102db3660046115ca565b610838565b60405190151581526020016101fc565b6102186102fe3660046115ca565b610843565b61030b610885565b6040516101fc9190611757565b6102186103263660046116bc565b6108e7565b6102186108fd565b610218610341366004611740565b610910565b61034e610954565b6040516101fc97969594939291906117d0565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101e8565b6103a461039f3660046116bc565b6109fd565b6040516101fc9190611866565b6102186103bf3660046116bc565b610b0e565b610236610b45565b6102186103da366004611720565b610b55565b6101e86103ed366004611740565b610b9f565b610218610400366004611920565b610bab565b610236610c3c565b61021861041b3660046119ae565b610c46565b61021861042e3660046119ae565b610c8e565b6102186104413660046115ca565b610d03565b61030b610d3d565b6102e061045c3660046115ca565b610d49565b61046c816001610d55565b50565b5f5b81518110156104a55761049d82828151811061048f5761048f6119de565b60200260200101515f610d55565b600101610471565b5050565b5f6104b382610838565b82906104e357604051630f2c982560e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b506001600160a01b038083165f90815260676020908152604091829020825160608101845281549094168452600181015491840191909152600201549082018190525f191480610544575060665481604001516105409190611a06565b4211155b8360665483604001516105579190611a06565b90916105875760405163e3ed050960e01b81526001600160a01b03909216600483015260248201526044016104da565b50506020015192915050565b61046c815f610d55565b6105c760405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b6105d082610838565b82906105fb57604051630f2c982560e01b81526001600160a01b0390911660048201526024016104da565b50506001600160a01b039081165f90815260676020908152604091829020825160608101845281549094168452600181015491840191909152600201549082015290565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156106845750825b90505f8267ffffffffffffffff1660011480156106a05750303b155b9050811580156106ae575080155b156106cc5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106f657845460ff60401b1916600160401b1785555b6107003387610e1e565b611c206066556107106001610e8e565b604080516060810190915260018152602081016103e86107326006600a611afc565b61073c9190611b07565b81525f1960209182015260015f526067815281517f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71d80546001600160a01b0319166001600160a01b039092169190911790558101517f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71e55604001517f6bee784efeb983674392298ab585b22866bedf00ebb0eea949d1e66f3f50e71f55831561081f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f6108328183610f23565b92915050565b5f6108328183610f35565b61084b610f56565b61085481610fb1565b6001600160a01b03165f90815260676020526040812080546001600160a01b03191681556001810182905560020155565b60605f6108915f611046565b90508067ffffffffffffffff8111156108ac576108ac6115e3565b6040519080825280602002602001820160405280156108d5578160200160208202803683370190505b5091506108e15f61104f565b91505090565b6108f4610228603361104f565b61046c81610b0e565b610905610f56565b61090e5f61105b565b565b610918610f56565b60668190556040518181527fca1d12600be991df47df34b798bb7437ed8c8422477a0ccda5bc128a78a13b83906020015b60405180910390a150565b5f60608082808083815f516020611ca55f395f51905f52805490915015801561097f57506001810154155b6109c35760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016104da565b6109cb6110cb565b6109d361118b565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6060815167ffffffffffffffff811115610a1957610a196115e3565b604051908082528060200260200182016040528015610a7457816020015b610a6160405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b815260200190600190039081610a375790505b5090505f5b8251811015610b085760675f848381518110610a9757610a976119de565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f2082516060810184528154909416845260018101549184019190915260020154908201528251839083908110610af557610af56119de565b6020908102919091010152600101610a79565b50919050565b5f5b81518110156104a557610b3d828281518110610b2e57610b2e6119de565b60200260200101516001610d55565b600101610b10565b5f610b506033611046565b905090565b610b5d610f56565b6035805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff90602001610949565b5f610832603383610f23565b8151835114610bcd5760405163282e5b7160e11b815260040160405180910390fd5b5f5b8351811015610c3657610c2e848281518110610bed57610bed6119de565b6020026020010151848381518110610c0757610c076119de565b6020026020010151848481518110610c2157610c216119de565b6020026020010151610c8e565b600101610bcf565b50505050565b5f610b505f611046565b610c4e610f56565b6001600160a01b038316610c755760405163620785c160e11b81526004016104da90611b26565b610c7e83610e8e565b610c898383836111c9565b505050565b610c9733610d49565b3390610cc2576040516338905e7160e11b81526001600160a01b0390911660048201526024016104da565b50610ccc83610838565b8390610cf757604051630f2c982560e01b81526001600160a01b0390911660048201526024016104da565b50610c898383836111c9565b610d0b610f56565b6001600160a01b038116610d3457604051631e4fbdf760e01b81525f60048201526024016104da565b61046c8161105b565b6060610b50603361104f565b5f610832603383610f35565b610d5d610f56565b8015610d9f57610d6e603383611306565b8290610d995760405163082cdf5560e21b81526001600160a01b0390911660048201526024016104da565b50610dd7565b610daa60338361131a565b8290610dd55760405163e491024560e01b81526001600160a01b0390911660048201526024016104da565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b610e2661132e565b610e2f82611377565b610e77604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b815250611388565b6035805460ff191660ff9290921691909117905550565b806001600160a01b038116610eb6576040516330de3edf60e11b81526004016104da90611b26565b610ec05f83611306565b8290610eeb576040516351eccfe160e11b81526001600160a01b0390911660048201526024016104da565b506040516001600160a01b038316907f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4905f90a25050565b5f610f2e838361139a565b9392505050565b6001600160a01b0381165f9081526001830160205260408120541515610f2e565b33610f887f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461090e5760405163118cdaa760e01b81523360048201526024016104da565b806001600160a01b038116610fd9576040516330de3edf60e11b81526004016104da90611b26565b610fe35f8361131a565b829061100e576040516340da71e560e01b81526001600160a01b0390911660048201526024016104da565b506040516001600160a01b038316907f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3905f90a25050565b5f610832825490565b60605f610f2e836113c0565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020611ca55f395f51905f529161110990611b45565b80601f016020809104026020016040519081016040528092919081815260200182805461113590611b45565b80156111805780601f1061115757610100808354040283529160200191611180565b820191905f5260205f20905b81548152906001019060200180831161116357829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020611ca55f395f51905f529161110990611b45565b6111d283610838565b83906111fd57604051630f2c982560e01b81526001600160a01b0390911660048201526024016104da565b50815f0361123657604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b60448201526064016104da565b5f198114806112455750428111155b8142909161126f5760405163248fc98560e11b8152600481019290925260248201526044016104da565b5050604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260678452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2910160405180910390a2505050565b5f610f2e836001600160a01b038416611419565b5f610f2e836001600160a01b038416611465565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661090e57604051631afcd79f60e31b815260040160405180910390fd5b61137f61132e565b61046c81611548565b61139061132e565b6104a58282611550565b5f825f0182815481106113af576113af6119de565b905f5260205f200154905092915050565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561140d57602002820191905f5260205f20905b8154815260200190600101908083116113f9575b50505050509050919050565b5f81815260018301602052604081205461145e57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610832565b505f610832565b5f818152600183016020526040812054801561153f575f611487600183611b77565b85549091505f9061149a90600190611b77565b90508082146114f9575f865f0182815481106114b8576114b86119de565b905f5260205f200154905080875f0184815481106114d8576114d86119de565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061150a5761150a611b8a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610832565b5f915050610832565b610d0b61132e565b61155861132e565b5f516020611ca55f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026115918482611be9565b50600381016115a08382611be9565b505f8082556001909101555050565b80356001600160a01b03811681146115c5575f5ffd5b919050565b5f602082840312156115da575f5ffd5b610f2e826115af565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611620576116206115e3565b604052919050565b5f67ffffffffffffffff821115611641576116416115e3565b5060051b60200190565b5f82601f83011261165a575f5ffd5b813561166d61166882611628565b6115f7565b8082825260208201915060208360051b86010192508583111561168e575f5ffd5b602085015b838110156116b2576116a4816115af565b835260209283019201611693565b5095945050505050565b5f602082840312156116cc575f5ffd5b813567ffffffffffffffff8111156116e2575f5ffd5b6116ee8482850161164b565b949350505050565b81516001600160a01b03168152602080830151908201526040808301519082015260608101610832565b5f60208284031215611730575f5ffd5b813560ff81168114610f2e575f5ffd5b5f60208284031215611750575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b818110156117975783516001600160a01b0316835260209384019390920191600101611770565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff60f81b8816815260e060208201525f6117ee60e08301896117a2565b828103604084015261180081896117a2565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611855578351835260209384019390920191600101611837565b50909b9a5050505050505050505050565b602080825282518282018190525f918401906040840190835b81811015611797576118af83855180516001600160a01b0316825260208082015190830152604090810151910152565b602093909301926060929092019160010161187f565b5f82601f8301126118d4575f5ffd5b81356118e261166882611628565b8082825260208201915060208360051b860101925085831115611903575f5ffd5b602085015b838110156116b2578035835260209283019201611908565b5f5f5f60608486031215611932575f5ffd5b833567ffffffffffffffff811115611948575f5ffd5b6119548682870161164b565b935050602084013567ffffffffffffffff811115611970575f5ffd5b61197c868287016118c5565b925050604084013567ffffffffffffffff811115611998575f5ffd5b6119a4868287016118c5565b9150509250925092565b5f5f5f606084860312156119c0575f5ffd5b6119c9846115af565b95602085013595506040909401359392505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610832576108326119f2565b6001815b6001841115611a5457808504811115611a3857611a386119f2565b6001841615611a4657908102905b60019390931c928002611a1d565b935093915050565b5f82611a6a57506001610832565b81611a7657505f610832565b8160018114611a8c5760028114611a9657611ab2565b6001915050610832565b60ff841115611aa757611aa76119f2565b50506001821b610832565b5060208310610133831016604e8410600b8410161715611ad5575081810a610832565b611ae15f198484611a19565b805f1904821115611af457611af46119f2565b029392505050565b5f610f2e8383611a5c565b5f82611b2157634e487b7160e01b5f52601260045260245ffd5b500490565b6020808252600590820152643a37b5b2b760d91b604082015260600190565b600181811c90821680611b5957607f821691505b602082108103610b0857634e487b7160e01b5f52602260045260245ffd5b81810381811115610832576108326119f2565b634e487b7160e01b5f52603160045260245ffd5b601f821115610c8957805f5260205f20601f840160051c81016020851015611bc35750805b601f840160051c820191505b81811015611be2575f8155600101611bcf565b5050505050565b815167ffffffffffffffff811115611c0357611c036115e3565b611c1781611c118454611b45565b84611b9e565b6020601f821160018114611c49575f8315611c325750848201515b5f19600385901b1c1916600184901b178455611be2565b5f84815260208120601f198516915b82811015611c785787850151825560209485019460019092019101611c58565b5084821015611c9557868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100a264697066735822122064ee2687ae1d7144a1d1a84e635b0116c0955a4bf3ba891523090aec0638867964736f6c634300081c0033",
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
// Solidity: function getValidPrice(address token) view returns(uint256 price)
func (_PriceFeedCross *PriceFeedCrossCaller) GetValidPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedCross.contract.Call(opts, &out, "getValidPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidPrice is a free data retrieval call binding the contract method 0x33ddef07.
//
// Solidity: function getValidPrice(address token) view returns(uint256 price)
func (_PriceFeedCross *PriceFeedCrossSession) GetValidPrice(token common.Address) (*big.Int, error) {
	return _PriceFeedCross.Contract.GetValidPrice(&_PriceFeedCross.CallOpts, token)
}

// GetValidPrice is a free data retrieval call binding the contract method 0x33ddef07.
//
// Solidity: function getValidPrice(address token) view returns(uint256 price)
func (_PriceFeedCross *PriceFeedCrossCallerSession) GetValidPrice(token common.Address) (*big.Int, error) {
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

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _threshold) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) Initialize(opts *bind.TransactOpts, _threshold uint8) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "initialize", _threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _threshold) returns()
func (_PriceFeedCross *PriceFeedCrossSession) Initialize(_threshold uint8) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.Initialize(&_PriceFeedCross.TransactOpts, _threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _threshold) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) Initialize(_threshold uint8) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.Initialize(&_PriceFeedCross.TransactOpts, _threshold)
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

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ResetValidators(&_PriceFeedCross.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeedCross *PriceFeedCrossTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeedCross.Contract.ResetValidators(&_PriceFeedCross.TransactOpts, validators)
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
