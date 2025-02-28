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

// IPriceFeedPriceData is an auto generated low-level Go binding around an user-defined struct.
type IPriceFeedPriceData struct {
	Token       common.Address
	Price       *big.Int
	LastUpdated *big.Int
}

// PriceFeedMetaData contains all meta data concerning the PriceFeed contract.
var PriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"allPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPriceDataList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_dollarDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"resetValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"setValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validators\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ValidatorSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNotExistToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerAlreadyExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"errValidatorManagerInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotExistValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errValidatorManagerNotValidator\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ec7221dd": "allPrices()",
		"f30589c3": "allValidators()",
		"b7f3358d": "changeThreshold(uint8)",
		"2fe3b6cf": "dollarDecimals()",
		"84b0196e": "eip712Domain()",
		"41976e09": "getPrice(address)",
		"72279ba1": "getPriceData(address)",
		"9761cb22": "getPriceDataList(address[])",
		"8fb5a482": "getPrices(address[])",
		"4351e6b6": "initialize(uint8)",
		"facd743b": "isValidator(address)",
		"167b78cd": "nativeCoin()",
		"8da5cb5b": "owner()",
		"40a141ff": "removeValidator(address)",
		"1d40f0d8": "removeValidators(address[])",
		"715018a6": "renounceOwnership()",
		"7101fcd3": "resetValidators(address[])",
		"1327d3d8": "setValidator(address)",
		"9300c926": "setValidators(address[])",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"99e6440d": "updatePrice(address[],uint256[],uint256[])",
		"7519ab50": "updatedAt()",
		"cbae5958": "validatorByIndex(uint256)",
		"aed1d403": "validatorLength()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50611b9b8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061016d575f3560e01c806384b0196e116100d9578063aed1d40311610093578063ec7221dd1161006e578063ec7221dd1461036e578063f2fde38b14610376578063f30589c314610389578063facd743b1461039e575f5ffd5b8063aed1d40314610340578063b7f3358d14610348578063cbae59581461035b575f5ffd5b806384b0196e1461028c5780638da5cb5b146102a75780638fb5a482146102d75780639300c926146102f95780639761cb221461030c57806399e6440d1461032d575f5ffd5b806342cde4e81161012a57806342cde4e81461021b5780634351e6b6146102265780637101fcd314610239578063715018a61461024c57806372279ba1146102545780637519ab5014610275575f5ffd5b80631327d3d814610171578063167b78cd146101865780631d40f0d8146101a65780632fe3b6cf146101b957806340a141ff146101d857806341976e09146101eb575b5f5ffd5b61018461017f366004611426565b6103c1565b005b60015b6040516001600160a01b0390911681526020015b60405180910390f35b6101846101b4366004611516565b6103cf565b6033546101c69060ff1681565b60405160ff909116815260200161019d565b6101846101e6366004611426565b610409565b6101fe6101f9366004611426565b610413565b60408051931515845260208401929092529082015260600161019d565b60025460ff166101c6565b61018461023436600461154f565b610468565b610184610247366004611516565b61065d565b610184610672565b610267610262366004611426565b610685565b60405161019d92919061156f565b61027e60345481565b60405190815260200161019d565b61029461070d565b60405161019d979695949392919061160a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610189565b6102ea6102e5366004611516565b6107bb565b60405161019d939291906116ab565b610184610307366004611516565b610951565b61031f61031a366004611516565b610988565b60405161019d9291906116e0565b61018461033b3660046117bb565b610b69565b61027e610c33565b61018461035636600461154f565b610c42565b610189610369366004611846565b610c92565b6102ea610ca3565b610184610384366004611426565b610cbf565b610391610cf9565b60405161019d919061185d565b6103b16103ac366004611426565b610d04565b604051901515815260200161019d565b6103cc816001610d0f565b50565b5f5b8151811015610405576103fd8282815181106103ef576103ef6118a8565b60200260200101515f610d0f565b6001016103d1565b5050565b6103cc815f610d0f565b6001600160a01b038082165f818152603760205260408120549092839283929091161461045d576001600160a01b0384165f90815260376020526040902060019081015490935091505b506034549193909250565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156104ac5750825b90505f826001600160401b031660011480156104c75750303b155b9050811580156104d5575080155b156104f35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561051d57845460ff60401b1916600160401b1785555b61052633610dd6565b61052f5f610de7565b6033805460ff191660ff8816908117909155426034556040805160608101909152600181529060208201906103e89061056990600a6119b3565b61057391906119c1565b81524260209182015260015f526037815281517f8e0cc0f1f0504b4cb44a23b328568106915b169e79003737a7b094503cdbeeb080546001600160a01b0319166001600160a01b039092169190911790558101517f8e0cc0f1f0504b4cb44a23b328568106915b169e79003737a7b094503cdbeeb155604001517f8e0cc0f1f0504b4cb44a23b328568106915b169e79003737a7b094503cdbeeb255831561065557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6106696101b45f610e4d565b6103cc81610951565b61067a610e60565b6106835f610ebb565b565b604080516060810182525f808252602082018190529181018290526001600160a01b038084165f81815260376020526040902054909116036107085750506001600160a01b038082165f90815260376020908152604091829020825160608101845281549094168452600181810154928501929092526002015491830191909152905b915091565b5f60608082808083815f516020611b465f395f51905f52805490915015801561073857506001810154155b6107815760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b610789610f2b565b610791610feb565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6060805f83516001600160401b038111156107d8576107d861143f565b604051908082528060200260200182016040528015610801578160200160208202803683370190505b50925083516001600160401b0381111561081d5761081d61143f565b604051908082528060200260200182016040528015610846578160200160208202803683370190505b5091505f5b845181101561094557848181518110610866576108666118a8565b60200260200101516001600160a01b031660375f87848151811061088c5761088c6118a8565b6020908102919091018101516001600160a01b039081168352908201929092526040015f2054160361093d5760018482815181106108cc576108cc6118a8565b60200260200101901515908115158152505060375f8683815181106108f3576108f36118a8565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2060010154838281518110610930576109306118a8565b6020026020010181815250505b60010161084b565b50506034549193909250565b5f5b815181101561040557610980828281518110610971576109716118a8565b60200260200101516001610d0f565b600101610953565b60608082516001600160401b038111156109a4576109a461143f565b6040519080825280602002602001820160405280156109cd578160200160208202803683370190505b50915082516001600160401b038111156109e9576109e961143f565b604051908082528060200260200182016040528015610a4457816020015b610a3160405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b815260200190600190039081610a075790505b5090505f5b8351811015610b6357838181518110610a6457610a646118a8565b60200260200101516001600160a01b031660375f868481518110610a8a57610a8a6118a8565b6020908102919091018101516001600160a01b039081168352908201929092526040015f20541603610b5b576001838281518110610aca57610aca6118a8565b60200260200101901515908115158152505060375f858381518110610af157610af16118a8565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f2082516060810184528154909416845260018101549184019190915260020154908201528251839083908110610b4f57610b4f6118a8565b60200260200101819052505b600101610a49565b50915091565b610b7233610d04565b3390610b9d5760405163845a09e760e01b81526001600160a01b039091166004820152602401610778565b508151835114610bc05760405163282e5b7160e11b815260040160405180910390fd5b5f5b8351811015610c2957610c21848281518110610be057610be06118a8565b6020026020010151848381518110610bfa57610bfa6118a8565b6020026020010151848481518110610c1457610c146118a8565b6020026020010151611029565b600101610bc2565b5050426034555050565b5f610c3d5f61113e565b905090565b610c4a610e60565b6002805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9060200160405180910390a150565b5f610c9d8183611147565b92915050565b6060805f610cb46102e56035610e4d565b925092509250909192565b610cc7610e60565b6001600160a01b038116610cf057604051631e4fbdf760e01b81525f6004820152602401610778565b6103cc81610ebb565b6060610c3d5f610e4d565b5f610c9d8183611152565b610d17610e60565b8015610d5857610d275f83611173565b8290610d52576040516329a04e7760e21b81526001600160a01b039091166004820152602401610778565b50610d8f565b610d625f83611187565b8290610d8d5760405163fdbc594760e01b81526001600160a01b039091166004820152602401610778565b505b604080516001600160a01b038416815282151560208201527f6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d910160405180910390a15050565b610dde61119b565b6103cc816111e4565b610def61119b565b610e37604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506111ec565b6002805460ff191660ff92909216919091179055565b60605f610e59836111fe565b9392505050565b33610e927f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106835760405163118cdaa760e01b8152336004820152602401610778565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f516020611b465f395f51905f5291610f69906119e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610f95906119e0565b8015610fe05780601f10610fb757610100808354040283529160200191610fe0565b820191905f5260205f20905b815481529060010190602001808311610fc357829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f516020611b465f395f51905f5291610f69906119e0565b815f0361106157604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b6044820152606401610778565b5f198114806110705750428111155b8142909161109a5760405163248fc98560e11b815260048101929092526024820152604401610778565b506110a89050603584611173565b50604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260378452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2910160405180910390a2505050565b5f610c9d825490565b5f610e598383611257565b6001600160a01b0381165f9081526001830160205260408120541515610e59565b5f610e59836001600160a01b03841661127d565b5f610e59836001600160a01b0384166112c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661068357604051631afcd79f60e31b815260040160405180910390fd5b610cc761119b565b6111f461119b565b61040582826113ac565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561124b57602002820191905f5260205f20905b815481526020019060010190808311611237575b50505050509050919050565b5f825f01828154811061126c5761126c6118a8565b905f5260205f200154905092915050565b5f8181526001830160205260408120546112c257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610c9d565b505f610c9d565b5f81815260018301602052604081205480156113a3575f6112eb600183611a18565b85549091505f906112fe90600190611a18565b905080821461135d575f865f01828154811061131c5761131c6118a8565b905f5260205f200154905080875f01848154811061133c5761133c6118a8565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061136e5761136e611a2b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610c9d565b5f915050610c9d565b6113b461119b565b5f516020611b465f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026113ed8482611a8b565b50600381016113fc8382611a8b565b505f8082556001909101555050565b80356001600160a01b0381168114611421575f5ffd5b919050565b5f60208284031215611436575f5ffd5b610e598261140b565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561147b5761147b61143f565b604052919050565b5f6001600160401b0382111561149b5761149b61143f565b5060051b60200190565b5f82601f8301126114b4575f5ffd5b81356114c76114c282611483565b611453565b8082825260208201915060208360051b8601019250858311156114e8575f5ffd5b602085015b8381101561150c576114fe8161140b565b8352602092830192016114ed565b5095945050505050565b5f60208284031215611526575f5ffd5b81356001600160401b0381111561153b575f5ffd5b611547848285016114a5565b949350505050565b5f6020828403121561155f575f5ffd5b813560ff81168114610e59575f5ffd5b821515815260808101610e59602083018480516001600160a01b0316825260208082015190830152604090810151910152565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b828110156116005781518652602095860195909101906001016115e2565b5093949350505050565b60ff60f81b8816815260e060208201525f61162860e08301896115a2565b828103604084015261163a81896115a2565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061166b81856115d0565b9a9950505050505050505050565b5f8151808452602084019350602083015f5b82811015611600578151151586526020958601959091019060010161168b565b606081525f6116bd6060830186611679565b82810360208401526116cf81866115d0565b915050826040830152949350505050565b604081525f6116f26040830185611679565b82810360208401528084518083526020830191506020860192505f5b818110156117545761173e83855180516001600160a01b0316825260208082015190830152604090810151910152565b602093909301926060929092019160010161170e565b50909695505050505050565b5f82601f83011261176f575f5ffd5b813561177d6114c282611483565b8082825260208201915060208360051b86010192508583111561179e575f5ffd5b602085015b8381101561150c5780358352602092830192016117a3565b5f5f5f606084860312156117cd575f5ffd5b83356001600160401b038111156117e2575f5ffd5b6117ee868287016114a5565b93505060208401356001600160401b03811115611809575f5ffd5b61181586828701611760565b92505060408401356001600160401b03811115611830575f5ffd5b61183c86828701611760565b9150509250925092565b5f60208284031215611856575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b8181101561189d5783516001600160a01b0316835260209384019390920191600101611876565b509095945050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b6001815b600184111561190b578085048111156118ef576118ef6118bc565b60018416156118fd57908102905b60019390931c9280026118d4565b935093915050565b5f8261192157506001610c9d565b8161192d57505f610c9d565b8160018114611943576002811461194d57611969565b6001915050610c9d565b60ff84111561195e5761195e6118bc565b50506001821b610c9d565b5060208310610133831016604e8410600b841016171561198c575081810a610c9d565b6119985f1984846118d0565b805f19048211156119ab576119ab6118bc565b029392505050565b5f610e5960ff841683611913565b5f826119db57634e487b7160e01b5f52601260045260245ffd5b500490565b600181811c908216806119f457607f821691505b602082108103611a1257634e487b7160e01b5f52602260045260245ffd5b50919050565b81810381811115610c9d57610c9d6118bc565b634e487b7160e01b5f52603160045260245ffd5b601f821115611a8657805f5260205f20601f840160051c81016020851015611a645750805b601f840160051c820191505b81811015611a83575f8155600101611a70565b50505b505050565b81516001600160401b03811115611aa457611aa461143f565b611ab881611ab284546119e0565b84611a3f565b6020601f821160018114611aea575f8315611ad35750848201515b5f19600385901b1c1916600184901b178455611a83565b5f84815260208120601f198516915b82811015611b195787850151825560209485019460019092019101611af9565b5084821015611b3657868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100a2646970667358221220bf56389143e79d2f1a18f0b24c12ef710a42ec9f6562a4e9137478cfeb80587164736f6c634300081c0033",
}

// PriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedMetaData.ABI instead.
var PriceFeedABI = PriceFeedMetaData.ABI

// Deprecated: Use PriceFeedMetaData.Sigs instead.
// PriceFeedFuncSigs maps the 4-byte function signature to its string representation.
var PriceFeedFuncSigs = PriceFeedMetaData.Sigs

// PriceFeedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedMetaData.Bin instead.
var PriceFeedBin = PriceFeedMetaData.Bin

// DeployPriceFeed deploys a new Ethereum contract, binding an instance of PriceFeed to it.
func DeployPriceFeed(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeed, error) {
	parsed, err := PriceFeedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeed{PriceFeedCaller: PriceFeedCaller{contract: contract}, PriceFeedTransactor: PriceFeedTransactor{contract: contract}, PriceFeedFilterer: PriceFeedFilterer{contract: contract}}, nil
}

// PriceFeed is an auto generated Go binding around an Ethereum contract.
type PriceFeed struct {
	PriceFeedCaller     // Read-only binding to the contract
	PriceFeedTransactor // Write-only binding to the contract
	PriceFeedFilterer   // Log filterer for contract events
}

// PriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedSession struct {
	Contract     *PriceFeed        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedCallerSession struct {
	Contract *PriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedTransactorSession struct {
	Contract     *PriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedRaw struct {
	Contract *PriceFeed // Generic contract binding to access the raw methods on
}

// PriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedCallerRaw struct {
	Contract *PriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedTransactorRaw struct {
	Contract *PriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeed creates a new instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeed(address common.Address, backend bind.ContractBackend) (*PriceFeed, error) {
	contract, err := bindPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeed{PriceFeedCaller: PriceFeedCaller{contract: contract}, PriceFeedTransactor: PriceFeedTransactor{contract: contract}, PriceFeedFilterer: PriceFeedFilterer{contract: contract}}, nil
}

// NewPriceFeedCaller creates a new read-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedCaller, error) {
	contract, err := bindPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedCaller{contract: contract}, nil
}

// NewPriceFeedTransactor creates a new write-only instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedTransactor, error) {
	contract, err := bindPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTransactor{contract: contract}, nil
}

// NewPriceFeedFilterer creates a new log filterer instance of PriceFeed, bound to a specific deployed contract.
func NewPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedFilterer, error) {
	contract, err := bindPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedFilterer{contract: contract}, nil
}

// bindPriceFeed binds a generic wrapper to an already deployed contract.
func bindPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.PriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.PriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeed *PriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeed *PriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeed *PriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeed.Contract.contract.Transact(opts, method, params...)
}

// AllPrices is a free data retrieval call binding the contract method 0xec7221dd.
//
// Solidity: function allPrices() view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) AllPrices(opts *bind.CallOpts) (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "allPrices")

	outstruct := new(struct {
		Exist     []bool
		Prices    []*big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	outstruct.Prices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllPrices is a free data retrieval call binding the contract method 0xec7221dd.
//
// Solidity: function allPrices() view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) AllPrices() (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.AllPrices(&_PriceFeed.CallOpts)
}

// AllPrices is a free data retrieval call binding the contract method 0xec7221dd.
//
// Solidity: function allPrices() view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) AllPrices() (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.AllPrices(&_PriceFeed.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeed *PriceFeedCaller) AllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeed *PriceFeedSession) AllValidators() ([]common.Address, error) {
	return _PriceFeed.Contract.AllValidators(&_PriceFeed.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns(address[])
func (_PriceFeed *PriceFeedCallerSession) AllValidators() ([]common.Address, error) {
	return _PriceFeed.Contract.AllValidators(&_PriceFeed.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint8)
func (_PriceFeed *PriceFeedCaller) DollarDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "dollarDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint8)
func (_PriceFeed *PriceFeedSession) DollarDecimals() (uint8, error) {
	return _PriceFeed.Contract.DollarDecimals(&_PriceFeed.CallOpts)
}

// DollarDecimals is a free data retrieval call binding the contract method 0x2fe3b6cf.
//
// Solidity: function dollarDecimals() view returns(uint8)
func (_PriceFeed *PriceFeedCallerSession) DollarDecimals() (uint8, error) {
	return _PriceFeed.Contract.DollarDecimals(&_PriceFeed.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PriceFeed *PriceFeedCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "eip712Domain")

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
func (_PriceFeed *PriceFeedSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PriceFeed.Contract.Eip712Domain(&_PriceFeed.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_PriceFeed *PriceFeedCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _PriceFeed.Contract.Eip712Domain(&_PriceFeed.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) GetPrice(opts *bind.CallOpts, token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getPrice", token)

	outstruct := new(struct {
		Exist     bool
		Price     *big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) GetPrice(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrice(&_PriceFeed.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) GetPrice(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrice(&_PriceFeed.CallOpts, token)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address token) view returns(bool exist, (address,uint256,uint256) data)
func (_PriceFeed *PriceFeedCaller) GetPriceData(opts *bind.CallOpts, token common.Address) (struct {
	Exist bool
	Data  IPriceFeedPriceData
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getPriceData", token)

	outstruct := new(struct {
		Exist bool
		Data  IPriceFeedPriceData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Data = *abi.ConvertType(out[1], new(IPriceFeedPriceData)).(*IPriceFeedPriceData)

	return *outstruct, err

}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address token) view returns(bool exist, (address,uint256,uint256) data)
func (_PriceFeed *PriceFeedSession) GetPriceData(token common.Address) (struct {
	Exist bool
	Data  IPriceFeedPriceData
}, error) {
	return _PriceFeed.Contract.GetPriceData(&_PriceFeed.CallOpts, token)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address token) view returns(bool exist, (address,uint256,uint256) data)
func (_PriceFeed *PriceFeedCallerSession) GetPriceData(token common.Address) (struct {
	Exist bool
	Data  IPriceFeedPriceData
}, error) {
	return _PriceFeed.Contract.GetPriceData(&_PriceFeed.CallOpts, token)
}

// GetPriceDataList is a free data retrieval call binding the contract method 0x9761cb22.
//
// Solidity: function getPriceDataList(address[] tokens) view returns(bool[] exist, (address,uint256,uint256)[] list)
func (_PriceFeed *PriceFeedCaller) GetPriceDataList(opts *bind.CallOpts, tokens []common.Address) (struct {
	Exist []bool
	List  []IPriceFeedPriceData
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getPriceDataList", tokens)

	outstruct := new(struct {
		Exist []bool
		List  []IPriceFeedPriceData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	outstruct.List = *abi.ConvertType(out[1], new([]IPriceFeedPriceData)).(*[]IPriceFeedPriceData)

	return *outstruct, err

}

// GetPriceDataList is a free data retrieval call binding the contract method 0x9761cb22.
//
// Solidity: function getPriceDataList(address[] tokens) view returns(bool[] exist, (address,uint256,uint256)[] list)
func (_PriceFeed *PriceFeedSession) GetPriceDataList(tokens []common.Address) (struct {
	Exist []bool
	List  []IPriceFeedPriceData
}, error) {
	return _PriceFeed.Contract.GetPriceDataList(&_PriceFeed.CallOpts, tokens)
}

// GetPriceDataList is a free data retrieval call binding the contract method 0x9761cb22.
//
// Solidity: function getPriceDataList(address[] tokens) view returns(bool[] exist, (address,uint256,uint256)[] list)
func (_PriceFeed *PriceFeedCallerSession) GetPriceDataList(tokens []common.Address) (struct {
	Exist []bool
	List  []IPriceFeedPriceData
}, error) {
	return _PriceFeed.Contract.GetPriceDataList(&_PriceFeed.CallOpts, tokens)
}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] tokens) view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) GetPrices(opts *bind.CallOpts, tokens []common.Address) (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getPrices", tokens)

	outstruct := new(struct {
		Exist     []bool
		Prices    []*big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exist = *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	outstruct.Prices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] tokens) view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) GetPrices(tokens []common.Address) (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrices(&_PriceFeed.CallOpts, tokens)
}

// GetPrices is a free data retrieval call binding the contract method 0x8fb5a482.
//
// Solidity: function getPrices(address[] tokens) view returns(bool[] exist, uint256[] prices, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) GetPrices(tokens []common.Address) (struct {
	Exist     []bool
	Prices    []*big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrices(&_PriceFeed.CallOpts, tokens)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeed *PriceFeedCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeed *PriceFeedSession) IsValidator(validator common.Address) (bool, error) {
	return _PriceFeed.Contract.IsValidator(&_PriceFeed.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_PriceFeed *PriceFeedCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _PriceFeed.Contract.IsValidator(&_PriceFeed.CallOpts, validator)
}

// NativeCoin is a free data retrieval call binding the contract method 0x167b78cd.
//
// Solidity: function nativeCoin() pure returns(address)
func (_PriceFeed *PriceFeedCaller) NativeCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "nativeCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeCoin is a free data retrieval call binding the contract method 0x167b78cd.
//
// Solidity: function nativeCoin() pure returns(address)
func (_PriceFeed *PriceFeedSession) NativeCoin() (common.Address, error) {
	return _PriceFeed.Contract.NativeCoin(&_PriceFeed.CallOpts)
}

// NativeCoin is a free data retrieval call binding the contract method 0x167b78cd.
//
// Solidity: function nativeCoin() pure returns(address)
func (_PriceFeed *PriceFeedCallerSession) NativeCoin() (common.Address, error) {
	return _PriceFeed.Contract.NativeCoin(&_PriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedSession) Owner() (common.Address, error) {
	return _PriceFeed.Contract.Owner(&_PriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeed *PriceFeedCallerSession) Owner() (common.Address, error) {
	return _PriceFeed.Contract.Owner(&_PriceFeed.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeed *PriceFeedCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeed *PriceFeedSession) Threshold() (uint8, error) {
	return _PriceFeed.Contract.Threshold(&_PriceFeed.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_PriceFeed *PriceFeedCallerSession) Threshold() (uint8, error) {
	return _PriceFeed.Contract.Threshold(&_PriceFeed.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeed *PriceFeedCaller) UpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "updatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeed *PriceFeedSession) UpdatedAt() (*big.Int, error) {
	return _PriceFeed.Contract.UpdatedAt(&_PriceFeed.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_PriceFeed *PriceFeedCallerSession) UpdatedAt() (*big.Int, error) {
	return _PriceFeed.Contract.UpdatedAt(&_PriceFeed.CallOpts)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeed *PriceFeedCaller) ValidatorByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "validatorByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeed *PriceFeedSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _PriceFeed.Contract.ValidatorByIndex(&_PriceFeed.CallOpts, index)
}

// ValidatorByIndex is a free data retrieval call binding the contract method 0xcbae5958.
//
// Solidity: function validatorByIndex(uint256 index) view returns(address)
func (_PriceFeed *PriceFeedCallerSession) ValidatorByIndex(index *big.Int) (common.Address, error) {
	return _PriceFeed.Contract.ValidatorByIndex(&_PriceFeed.CallOpts, index)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeed *PriceFeedCaller) ValidatorLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "validatorLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeed *PriceFeedSession) ValidatorLength() (*big.Int, error) {
	return _PriceFeed.Contract.ValidatorLength(&_PriceFeed.CallOpts)
}

// ValidatorLength is a free data retrieval call binding the contract method 0xaed1d403.
//
// Solidity: function validatorLength() view returns(uint256)
func (_PriceFeed *PriceFeedCallerSession) ValidatorLength() (*big.Int, error) {
	return _PriceFeed.Contract.ValidatorLength(&_PriceFeed.CallOpts)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeed *PriceFeedTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeed *PriceFeedSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.ChangeThreshold(&_PriceFeed.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_PriceFeed *PriceFeedTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.ChangeThreshold(&_PriceFeed.TransactOpts, threshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedTransactor) Initialize(opts *bind.TransactOpts, _dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "initialize", _dollarDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedSession) Initialize(_dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, _dollarDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedTransactorSession) Initialize(_dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, _dollarDecimals)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeed *PriceFeedTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeed *PriceFeedSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveValidator(&_PriceFeed.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_PriceFeed *PriceFeedTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveValidator(&_PriceFeed.TransactOpts, validator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactor) RemoveValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "removeValidators", validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveValidators(&_PriceFeed.TransactOpts, validators)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactorSession) RemoveValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveValidators(&_PriceFeed.TransactOpts, validators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceOwnership(&_PriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceFeed *PriceFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceOwnership(&_PriceFeed.TransactOpts)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactor) ResetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "resetValidators", validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.ResetValidators(&_PriceFeed.TransactOpts, validators)
}

// ResetValidators is a paid mutator transaction binding the contract method 0x7101fcd3.
//
// Solidity: function resetValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactorSession) ResetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.ResetValidators(&_PriceFeed.TransactOpts, validators)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeed *PriceFeedTransactor) SetValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "setValidator", validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeed *PriceFeedSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetValidator(&_PriceFeed.TransactOpts, validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address validator) returns()
func (_PriceFeed *PriceFeedTransactorSession) SetValidator(validator common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetValidator(&_PriceFeed.TransactOpts, validator)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactor) SetValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "setValidators", validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetValidators(&_PriceFeed.TransactOpts, validators)
}

// SetValidators is a paid mutator transaction binding the contract method 0x9300c926.
//
// Solidity: function setValidators(address[] validators) returns()
func (_PriceFeed *PriceFeedTransactorSession) SetValidators(validators []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetValidators(&_PriceFeed.TransactOpts, validators)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.TransferOwnership(&_PriceFeed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PriceFeed *PriceFeedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.TransferOwnership(&_PriceFeed.TransactOpts, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x99e6440d.
//
// Solidity: function updatePrice(address[] tokens, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedTransactor) UpdatePrice(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "updatePrice", tokens, prices, pricesAt)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x99e6440d.
//
// Solidity: function updatePrice(address[] tokens, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedSession) UpdatePrice(tokens []common.Address, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdatePrice(&_PriceFeed.TransactOpts, tokens, prices, pricesAt)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x99e6440d.
//
// Solidity: function updatePrice(address[] tokens, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedTransactorSession) UpdatePrice(tokens []common.Address, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdatePrice(&_PriceFeed.TransactOpts, tokens, prices, pricesAt)
}

// PriceFeedEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the PriceFeed contract.
type PriceFeedEIP712DomainChangedIterator struct {
	Event *PriceFeedEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedEIP712DomainChanged)
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
		it.Event = new(PriceFeedEIP712DomainChanged)
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
func (it *PriceFeedEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedEIP712DomainChanged represents a EIP712DomainChanged event raised by the PriceFeed contract.
type PriceFeedEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PriceFeed *PriceFeedFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PriceFeedEIP712DomainChangedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedEIP712DomainChangedIterator{contract: _PriceFeed.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_PriceFeed *PriceFeedFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedEIP712DomainChanged)
				if err := _PriceFeed.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParseEIP712DomainChanged(log types.Log) (*PriceFeedEIP712DomainChanged, error) {
	event := new(PriceFeedEIP712DomainChanged)
	if err := _PriceFeed.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PriceFeed contract.
type PriceFeedInitializedIterator struct {
	Event *PriceFeedInitialized // Event containing the contract specifics and raw log

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
func (it *PriceFeedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedInitialized)
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
		it.Event = new(PriceFeedInitialized)
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
func (it *PriceFeedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedInitialized represents a Initialized event raised by the PriceFeed contract.
type PriceFeedInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PriceFeed *PriceFeedFilterer) FilterInitialized(opts *bind.FilterOpts) (*PriceFeedInitializedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PriceFeedInitializedIterator{contract: _PriceFeed.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PriceFeed *PriceFeedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PriceFeedInitialized) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedInitialized)
				if err := _PriceFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParseInitialized(log types.Log) (*PriceFeedInitialized, error) {
	event := new(PriceFeedInitialized)
	if err := _PriceFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceFeed contract.
type PriceFeedOwnershipTransferredIterator struct {
	Event *PriceFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedOwnershipTransferred)
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
		it.Event = new(PriceFeedOwnershipTransferred)
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
func (it *PriceFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedOwnershipTransferred represents a OwnershipTransferred event raised by the PriceFeed contract.
type PriceFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeed *PriceFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedOwnershipTransferredIterator{contract: _PriceFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PriceFeed *PriceFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedOwnershipTransferred)
				if err := _PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParseOwnershipTransferred(log types.Log) (*PriceFeedOwnershipTransferred, error) {
	event := new(PriceFeedOwnershipTransferred)
	if err := _PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedPriceFeedPriceUpdatedIterator is returned from FilterPriceFeedPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceFeedPriceUpdated events raised by the PriceFeed contract.
type PriceFeedPriceFeedPriceUpdatedIterator struct {
	Event *PriceFeedPriceFeedPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceFeedPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceFeedPriceUpdated)
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
		it.Event = new(PriceFeedPriceFeedPriceUpdated)
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
func (it *PriceFeedPriceFeedPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceFeedPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceFeedPriceUpdated represents a PriceFeedPriceUpdated event raised by the PriceFeed contract.
type PriceFeedPriceFeedPriceUpdated struct {
	Token     common.Address
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedPriceUpdated is a free log retrieval operation binding the contract event 0xc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2.
//
// Solidity: event PriceFeedPriceUpdated(address indexed token, uint256 price, uint256 timestamp)
func (_PriceFeed *PriceFeedFilterer) FilterPriceFeedPriceUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceFeedPriceFeedPriceUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceFeedPriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceFeedPriceUpdatedIterator{contract: _PriceFeed.contract, event: "PriceFeedPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceFeedPriceUpdated is a free log subscription operation binding the contract event 0xc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2.
//
// Solidity: event PriceFeedPriceUpdated(address indexed token, uint256 price, uint256 timestamp)
func (_PriceFeed *PriceFeedFilterer) WatchPriceFeedPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceFeedPriceUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceFeedPriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceFeedPriceUpdated)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedPriceUpdated", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParsePriceFeedPriceUpdated(log types.Log) (*PriceFeedPriceFeedPriceUpdated, error) {
	event := new(PriceFeedPriceFeedPriceUpdated)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceFeedPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the PriceFeed contract.
type PriceFeedThresholdChangedIterator struct {
	Event *PriceFeedThresholdChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedThresholdChanged)
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
		it.Event = new(PriceFeedThresholdChanged)
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
func (it *PriceFeedThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedThresholdChanged represents a ThresholdChanged event raised by the PriceFeed contract.
type PriceFeedThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_PriceFeed *PriceFeedFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*PriceFeedThresholdChangedIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedThresholdChangedIterator{contract: _PriceFeed.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_PriceFeed *PriceFeedFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedThresholdChanged)
				if err := _PriceFeed.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParseThresholdChanged(log types.Log) (*PriceFeedThresholdChanged, error) {
	event := new(PriceFeedThresholdChanged)
	if err := _PriceFeed.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedValidatorSetIterator is returned from FilterValidatorSet and is used to iterate over the raw logs and unpacked data for ValidatorSet events raised by the PriceFeed contract.
type PriceFeedValidatorSetIterator struct {
	Event *PriceFeedValidatorSet // Event containing the contract specifics and raw log

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
func (it *PriceFeedValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedValidatorSet)
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
		it.Event = new(PriceFeedValidatorSet)
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
func (it *PriceFeedValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedValidatorSet represents a ValidatorSet event raised by the PriceFeed contract.
type PriceFeedValidatorSet struct {
	Validators common.Address
	Status     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSet is a free log retrieval operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_PriceFeed *PriceFeedFilterer) FilterValidatorSet(opts *bind.FilterOpts) (*PriceFeedValidatorSetIterator, error) {

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return &PriceFeedValidatorSetIterator{contract: _PriceFeed.contract, event: "ValidatorSet", logs: logs, sub: sub}, nil
}

// WatchValidatorSet is a free log subscription operation binding the contract event 0x6b3b7d0d26ec99d080840dca1323c7147d1868adc66a4290afb8101d7908320d.
//
// Solidity: event ValidatorSet(address validators, bool status)
func (_PriceFeed *PriceFeedFilterer) WatchValidatorSet(opts *bind.WatchOpts, sink chan<- *PriceFeedValidatorSet) (event.Subscription, error) {

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "ValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedValidatorSet)
				if err := _PriceFeed.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
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
func (_PriceFeed *PriceFeedFilterer) ParseValidatorSet(log types.Log) (*PriceFeedValidatorSet, error) {
	event := new(PriceFeedValidatorSet)
	if err := _PriceFeed.contract.UnpackLog(event, "ValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
