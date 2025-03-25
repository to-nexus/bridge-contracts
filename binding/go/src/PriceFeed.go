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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNativeTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPriceDataList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceInDollars\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_dollarDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updateNativeTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"NativeTokenPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalcAmountLibOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNoSource\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"ec7221dd": "allPrices()",
		"2fe3b6cf": "dollarDecimals()",
		"3afb1544": "getNativeTokenPrice(uint256)",
		"ac41865a": "getPrice(address,address)",
		"72279ba1": "getPriceData(address)",
		"9761cb22": "getPriceDataList(address[])",
		"8fb5a482": "getPrices(address[])",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"a3246ad3": "getRoleMembers(bytes32)",
		"e588566f": "getTokenPriceInDollars(address)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"b2b49e2e": "grantRoleBatch(bytes32[],address[])",
		"91d14854": "hasRole(bytes32,address)",
		"943b24b2": "initialize(address,uint8)",
		"52d1902d": "proxiableUUID()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"01ffc9a7": "supportsInterface(bytes4)",
		"9409f47b": "updateNativeTokenPrice(uint256[],uint256[],uint256[])",
		"99e6440d": "updatePrice(address[],uint256[],uint256[])",
		"7519ab50": "updatedAt()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516125e96100395f395f81816112310152818161125a015261139d01526125e95ff3fe608060405260043610610161575f3560e01c806391d14854116100cd578063a3246ad311610087578063b2b49e2e11610062578063b2b49e2e1461047b578063d547741f1461049a578063e588566f146104b9578063ec7221dd146104d8575f5ffd5b8063a3246ad3146103de578063ac41865a1461040a578063ad3cb1cc1461043e575f5ffd5b806391d14854146103225780639409f47b14610341578063943b24b2146103605780639761cb221461037f57806399e6440d146103ac578063a217fddf146103cb575f5ffd5b80633afb15441161011e5780633afb15441461024f5780634f1ef2861461028b57806352d1902d1461029e57806372279ba1146102b25780637519ab50146102df5780638fb5a482146102f4575f5ffd5b806301ffc9a714610165578063248a9ca3146101995780632f2ff15d146101c65780632fe3b6cf146101e757806336568abe14610211578063365d71e414610230575b5f5ffd5b348015610170575f5ffd5b5061018461017f366004611d36565b6104ec565b60405190151581526020015b60405180910390f35b3480156101a4575f5ffd5b506101b86101b3366004611d5d565b610522565b604051908152602001610190565b3480156101d1575f5ffd5b506101e56101e0366004611d8f565b610542565b005b3480156101f2575f5ffd5b505f546101ff9060ff1681565b60405160ff9091168152602001610190565b34801561021c575f5ffd5b506101e561022b366004611d8f565b610564565b34801561023b575f5ffd5b506101e561024a366004611e86565b61059c565b34801561025a575f5ffd5b5061026e610269366004611d5d565b61060d565b604080519315158452602084019290925290820152606001610190565b6101e5610299366004611f40565b610676565b3480156102a9575f5ffd5b506101b8610695565b3480156102bd575f5ffd5b506102d16102cc366004611fe3565b6106b0565b604051610190929190611ffc565b3480156102ea575f5ffd5b506101b860015481565b3480156102ff575f5ffd5b5061031361030e36600461202f565b610738565b604051610190939291906120a4565b34801561032d575f5ffd5b5061018461033c366004611d8f565b6108ce565b34801561034c575f5ffd5b506101e561035b36600461215f565b610904565b34801561036b575f5ffd5b506101e561037a3660046121f8565b6109d1565b34801561038a575f5ffd5b5061039e61039936600461202f565b610b5d565b60405161019092919061222d565b3480156103b7575f5ffd5b506101e56103c63660046122ad565b610d3e565b3480156103d6575f5ffd5b506101b85f81565b3480156103e9575f5ffd5b506103fd6103f8366004611d5d565b610dff565b60405161019091906122e0565b348015610415575f5ffd5b5061042961042436600461232b565b610e42565b60408051928352602083019190915201610190565b348015610449575f5ffd5b5061046e604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516101909190612353565b348015610486575f5ffd5b506101e5610495366004611e86565b611013565b3480156104a5575f5ffd5b506101e56104b4366004611d8f565b611084565b3480156104c4575f5ffd5b5061026e6104d3366004611fe3565b6110a0565b3480156104e3575f5ffd5b506103136110f9565b5f6001600160e01b03198216637965db0b60e01b148061051c57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f9081525f5160206125945f395f51905f52602052604090206001015490565b61054b82610522565b61055481611115565b61055e8383611122565b50505050565b6001600160a01b038116331461058d5760405163334bd91960e11b815260040160405180910390fd5b61059782826111a9565b505050565b80518251146105be5760405163031206d560e51b815260040160405180910390fd5b5f5b8151811015610597576106058382815181106105de576105de612388565b60200260200101518383815181106105f8576105f8612388565b6020026020010151611084565b6001016105c0565b5f81815260056020908152604080832081516060810183528154808252600183015494820194909452600290910154918101919091528291829190851461065d575f5f5f9350935093505061066f565b60018160200151600154935093509350505b9193909250565b61067e611226565b610687826112cc565b61069182826112d6565b5050565b5f61069e611392565b505f5160206125745f395f51905f5290565b604080516060810182525f808252602082018190529181018290526001600160a01b038084165f81815260046020526040902054909116036107335750506001600160a01b038082165f90815260046020908152604091829020825160608101845281549094168452600181810154928501929092526002015491830191909152905b915091565b6060805f83516001600160401b0381111561075557610755611db9565b60405190808252806020026020018201604052801561077e578160200160208202803683370190505b50925083516001600160401b0381111561079a5761079a611db9565b6040519080825280602002602001820160405280156107c3578160200160208202803683370190505b5091505f5b84518110156108c2578481815181106107e3576107e3612388565b60200260200101516001600160a01b031660045f87848151811061080957610809612388565b6020908102919091018101516001600160a01b039081168352908201929092526040015f205416036108ba57600184828151811061084957610849612388565b60200260200101901515908115158152505060045f86838151811061087057610870612388565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20600101548382815181106108ad576108ad612388565b6020026020010181815250505b6001016107c8565b50506001549193909250565b5f9182525f5160206125945f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a61092e81611115565b83518351811480156109405750825181145b61095d5760405163282e5b7160e11b815260040160405180910390fd5b5f5b818110156109c5576109bd86828151811061097c5761097c612388565b602002602001015186838151811061099657610996612388565b60200260200101518684815181106109b0576109b0612388565b60200260200101516113db565b60010161095f565b50504260015550505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610a155750825b90505f826001600160401b03166001148015610a305750303b155b905081158015610a3e575080155b15610a5c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610a8657845460ff60401b1916600160401b1785555b866001600160a01b038116610abf57604051631a1f881960e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610ac86114a2565b610ad1876114aa565b610afb7fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a88611122565b505f805460ff191660ff8816179055426001558315610b5457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60608082516001600160401b03811115610b7957610b79611db9565b604051908082528060200260200182016040528015610ba2578160200160208202803683370190505b50915082516001600160401b03811115610bbe57610bbe611db9565b604051908082528060200260200182016040528015610c1957816020015b610c0660405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b815260200190600190039081610bdc5790505b5090505f5b8351811015610d3857838181518110610c3957610c39612388565b60200260200101516001600160a01b031660045f868481518110610c5f57610c5f612388565b6020908102919091018101516001600160a01b039081168352908201929092526040015f20541603610d30576001838281518110610c9f57610c9f612388565b60200260200101901515908115158152505060045f858381518110610cc657610cc6612388565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f2082516060810184528154909416845260018101549184019190915260020154908201528251839083908110610d2457610d24612388565b60200260200101819052505b600101610c1e565b50915091565b7fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a610d6881611115565b8351835181148015610d7a5750825181145b610d975760405163282e5b7160e11b815260040160405180910390fd5b5f5b818110156109c557610df7868281518110610db657610db6612388565b6020026020010151868381518110610dd057610dd0612388565b6020026020010151868481518110610dea57610dea612388565b60200260200101516114bb565b600101610d99565b5f8181527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f006020819052604090912060609190610e3b906115a3565b9392505050565b5f5f826001600160a01b0316846001600160a01b031603610e685750600190504361100c565b6040805160028082526060820183525f9260208301908036833701905050905084815f81518110610e9b57610e9b612388565b60200260200101906001600160a01b031690816001600160a01b0316815250508381600181518110610ecf57610ecf612388565b60200260200101906001600160a01b031690816001600160a01b031681525050606080610efb83610738565b8251909650919350915082905f90610f1557610f15612388565b60200260200101518790610f4857604051634ced742360e11b81526001600160a01b039091166004820152602401610ab6565b5081600181518110610f5c57610f5c612388565b60200260200101518690610f8f57604051634ced742360e11b81526001600160a01b039091166004820152602401610ab6565b505f5f610f9b896115af565b610fa4896115af565b9092509050610ffe610fb783600a612493565b610fc29060016124a1565b845f81518110610fd457610fd4612388565b602002602001015185600181518110610fef57610fef612388565b6020026020010151858561162e565b600154909750955050505050505b9250929050565b80518251146110355760405163031206d560e51b815260040160405180910390fd5b5f5b81518110156105975761107c83828151811061105557611055612388565b602002602001015183838151811061106f5761106f612388565b6020026020010151610542565b600101611037565b61108d82610522565b61109681611115565b61055e83836111a9565b6001600160a01b038082165f8181526004602090815260408083208151606081018352815490961680875260018201549387019390935260020154908501529092839283921461065d575f5f5f9350935093505061066f565b6060805f61110a61030e60026115a3565b925092509250909192565b61111f81336116ff565b50565b5f61112d8383611738565b9050801561051c575f8381527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f006020819052604090912061116e90846117d9565b838590916111a05760405163d180cb3160e01b81526001600160a01b0390921660048301526024820152604401610ab6565b50505092915050565b5f6111b483836117ed565b9050801561051c575f8381527f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00602081905260409091206111f59084611866565b838590916111a05760405162a95f1b60e31b81526001600160a01b0390921660048301526024820152604401610ab6565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112ac57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112a05f5160206125745f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156112ca5760405163703e46dd60e11b815260040160405180910390fd5b565b5f61069181611115565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611330575060408051601f3d908101601f1916820190925261132d918101906124b8565b60015b61135857604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610ab6565b5f5160206125745f395f51905f52811461138857604051632a87526960e21b815260048101829052602401610ab6565b610597838361187a565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112ca5760405163703e46dd60e11b815260040160405180910390fd5b815f036113fb5760405163406571d160e01b815260040160405180910390fd5b8042808211156114275760405163248fc98560e11b815260048101929092526024820152604401610ab6565b50506040805160608101825284815260208082018581528284018581525f88815260058452859020935184559051600184015551600290920191909155815184815290810183905284917f291ab1bcc002ef9062ceee3c7ea71d95374e2886b7ff336cb31a1e63e264965891015b60405180910390a2505050565b6112ca6118cf565b6114b26118cf565b61111f81611918565b815f036114db5760405163406571d160e01b815260040160405180910390fd5b8042808211156115075760405163248fc98560e11b815260048101929092526024820152604401610ab6565b5061151590506002846117d9565b50604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260048452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c09101611495565b60605f610e3b83611932565b5f6001600160a01b03821660011461162657816001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115fd573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061162191906124cf565b61051c565b601292915050565b5f845f03611668576040516349671cd160e11b815260206004820152600660248201526570726963654160d01b6044820152606401610ab6565b5f8360ff168360ff1610156116b0576116ab8661168585876124ea565b61169090600a612493565b61169a9088612517565b6116a49190612517565b889061198b565b6116d3565b6116d36116a46116c086866124ea565b6116cb90600a612493565b8790896119d0565b92509050806116f55760405163c7ef3d4f60e01b815260040160405180910390fd5b5095945050505050565b61170982826108ce565b6106915760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610ab6565b5f5f5160206125945f395f51905f5261175184846108ce565b6117d0575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556117863390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061051c565b5f91505061051c565b5f610e3b836001600160a01b038416611a86565b5f5f5160206125945f395f51905f5261180684846108ce565b156117d0575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061051c565b5f610e3b836001600160a01b038416611ad2565b61188382611bac565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156118c7576105978282611c0f565b610691611c81565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166112ca57604051631afcd79f60e31b815260040160405180910390fd5b6119206118cf565b6119286114a2565b6106915f82611122565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561197f57602002820191905f5260205f20905b81548152602001906001019080831161196b575b50505050509050919050565b5f5f835f0361199f5750600190505f61100c565b838302838582816119b2576119b2612503565b04146119c4575f5f925092505061100c565b60019590945092505050565b5f838302815f1985870982811083820303915050805f03611a04578382816119fa576119fa612503565b0492505050610e3b565b808411611a1b57611a1b6003851502601118611ca0565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f818152600183016020526040812054611acb57508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561051c565b505f61051c565b5f81815260018301602052604081205480156117d0575f611af4600183612536565b85549091505f90611b0790600190612536565b9050808214611b66575f865f018281548110611b2557611b25612388565b905f5260205f200154905080875f018481548110611b4557611b45612388565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611b7757611b77612549565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061051c565b806001600160a01b03163b5f03611be157604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610ab6565b5f5160206125745f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051611c2b919061255d565b5f60405180830381855af49150503d805f8114611c63576040519150601f19603f3d011682016040523d82523d5f602084013e611c68565b606091505b5091509150611c78858383611cb1565b95945050505050565b34156112ca5760405163b398979f60e01b815260040160405180910390fd5b634e487b715f52806020526024601cfd5b606082611cc657611cc182611d0d565b610e3b565b8151158015611cdd57506001600160a01b0384163b155b15611d0657604051639996b31560e01b81526001600160a01b0385166004820152602401610ab6565b5080610e3b565b805115611d1d5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f60208284031215611d46575f5ffd5b81356001600160e01b031981168114610e3b575f5ffd5b5f60208284031215611d6d575f5ffd5b5035919050565b80356001600160a01b0381168114611d8a575f5ffd5b919050565b5f5f60408385031215611da0575f5ffd5b82359150611db060208401611d74565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611df557611df5611db9565b604052919050565b5f6001600160401b03821115611e1557611e15611db9565b5060051b60200190565b5f82601f830112611e2e575f5ffd5b8135611e41611e3c82611dfd565b611dcd565b8082825260208201915060208360051b860101925085831115611e62575f5ffd5b602085015b838110156116f557611e7881611d74565b835260209283019201611e67565b5f5f60408385031215611e97575f5ffd5b82356001600160401b03811115611eac575f5ffd5b8301601f81018513611ebc575f5ffd5b8035611eca611e3c82611dfd565b8082825260208201915060208360051b850101925087831115611eeb575f5ffd5b6020840193505b82841015611f0d578335825260209384019390910190611ef2565b945050505060208301356001600160401b03811115611f2a575f5ffd5b611f3685828601611e1f565b9150509250929050565b5f5f60408385031215611f51575f5ffd5b611f5a83611d74565b915060208301356001600160401b03811115611f74575f5ffd5b8301601f81018513611f84575f5ffd5b80356001600160401b03811115611f9d57611f9d611db9565b611fb0601f8201601f1916602001611dcd565b818152866020838501011115611fc4575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215611ff3575f5ffd5b610e3b82611d74565b821515815260808101610e3b602083018480516001600160a01b0316825260208082015190830152604090810151910152565b5f6020828403121561203f575f5ffd5b81356001600160401b03811115612054575f5ffd5b61206084828501611e1f565b949350505050565b5f8151808452602084019350602083015f5b8281101561209a578151151586526020958601959091019060010161207a565b5093949350505050565b606081525f6120b66060830186612068565b82810360208401528085518083526020830191506020870192505f5b818110156120f05783518352602093840193909201916001016120d2565b505060409390930193909352509392505050565b5f82601f830112612113575f5ffd5b8135612121611e3c82611dfd565b8082825260208201915060208360051b860101925085831115612142575f5ffd5b602085015b838110156116f5578035835260209283019201612147565b5f5f5f60608486031215612171575f5ffd5b83356001600160401b03811115612186575f5ffd5b61219286828701612104565b93505060208401356001600160401b038111156121ad575f5ffd5b6121b986828701612104565b92505060408401356001600160401b038111156121d4575f5ffd5b6121e086828701612104565b9150509250925092565b60ff8116811461111f575f5ffd5b5f5f60408385031215612209575f5ffd5b61221283611d74565b91506020830135612222816121ea565b809150509250929050565b604081525f61223f6040830185612068565b82810360208401528084518083526020830191506020860192505f5b818110156122a15761228b83855180516001600160a01b0316825260208082015190830152604090810151910152565b602093909301926060929092019160010161225b565b50909695505050505050565b5f5f5f606084860312156122bf575f5ffd5b83356001600160401b038111156122d4575f5ffd5b61219286828701611e1f565b602080825282518282018190525f918401906040840190835b818110156123205783516001600160a01b03168352602093840193909201916001016122f9565b509095945050505050565b5f5f6040838503121561233c575f5ffd5b61234583611d74565b9150611db060208401611d74565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b6001815b60018411156123eb578085048111156123cf576123cf61239c565b60018416156123dd57908102905b60019390931c9280026123b4565b935093915050565b5f826124015750600161051c565b8161240d57505f61051c565b8160018114612423576002811461242d57612449565b600191505061051c565b60ff84111561243e5761243e61239c565b50506001821b61051c565b5060208310610133831016604e8410600b841016171561246c575081810a61051c565b6124785f1984846123b0565b805f190482111561248b5761248b61239c565b029392505050565b5f610e3b60ff8416836123f3565b808202811582820484141761051c5761051c61239c565b5f602082840312156124c8575f5ffd5b5051919050565b5f602082840312156124df575f5ffd5b8151610e3b816121ea565b60ff828116828216039081111561051c5761051c61239c565b634e487b7160e01b5f52601260045260245ffd5b5f8261253157634e487b7160e01b5f52601260045260245ffd5b500490565b8181038181111561051c5761051c61239c565b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a26469706673582212208d7836710457a68a9fab212413bc51309666c537539312c949b53cf46ccc73b464736f6c634300081c0033",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PriceFeed *PriceFeedCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PriceFeed *PriceFeedSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PriceFeed.Contract.DEFAULTADMINROLE(&_PriceFeed.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PriceFeed *PriceFeedCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PriceFeed.Contract.DEFAULTADMINROLE(&_PriceFeed.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PriceFeed *PriceFeedCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PriceFeed *PriceFeedSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PriceFeed.Contract.UPGRADEINTERFACEVERSION(&_PriceFeed.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PriceFeed *PriceFeedCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PriceFeed.Contract.UPGRADEINTERFACEVERSION(&_PriceFeed.CallOpts)
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

// GetNativeTokenPrice is a free data retrieval call binding the contract method 0x3afb1544.
//
// Solidity: function getNativeTokenPrice(uint256 chainID) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) GetNativeTokenPrice(opts *bind.CallOpts, chainID *big.Int) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getNativeTokenPrice", chainID)

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

// GetNativeTokenPrice is a free data retrieval call binding the contract method 0x3afb1544.
//
// Solidity: function getNativeTokenPrice(uint256 chainID) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) GetNativeTokenPrice(chainID *big.Int) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetNativeTokenPrice(&_PriceFeed.CallOpts, chainID)
}

// GetNativeTokenPrice is a free data retrieval call binding the contract method 0x3afb1544.
//
// Solidity: function getNativeTokenPrice(uint256 chainID) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) GetNativeTokenPrice(chainID *big.Int) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetNativeTokenPrice(&_PriceFeed.CallOpts, chainID)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address tokenA, address tokenB) view returns(uint256 price, uint256 lastUpdate)
func (_PriceFeed *PriceFeedCaller) GetPrice(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Price      *big.Int
	LastUpdate *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getPrice", tokenA, tokenB)

	outstruct := new(struct {
		Price      *big.Int
		LastUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address tokenA, address tokenB) view returns(uint256 price, uint256 lastUpdate)
func (_PriceFeed *PriceFeedSession) GetPrice(tokenA common.Address, tokenB common.Address) (struct {
	Price      *big.Int
	LastUpdate *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrice(&_PriceFeed.CallOpts, tokenA, tokenB)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address tokenA, address tokenB) view returns(uint256 price, uint256 lastUpdate)
func (_PriceFeed *PriceFeedCallerSession) GetPrice(tokenA common.Address, tokenB common.Address) (struct {
	Price      *big.Int
	LastUpdate *big.Int
}, error) {
	return _PriceFeed.Contract.GetPrice(&_PriceFeed.CallOpts, tokenA, tokenB)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PriceFeed *PriceFeedCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PriceFeed *PriceFeedSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PriceFeed.Contract.GetRoleAdmin(&_PriceFeed.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PriceFeed *PriceFeedCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PriceFeed.Contract.GetRoleAdmin(&_PriceFeed.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_PriceFeed *PriceFeedCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_PriceFeed *PriceFeedSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _PriceFeed.Contract.GetRoleMembers(&_PriceFeed.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_PriceFeed *PriceFeedCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _PriceFeed.Contract.GetRoleMembers(&_PriceFeed.CallOpts, role)
}

// GetTokenPriceInDollars is a free data retrieval call binding the contract method 0xe588566f.
//
// Solidity: function getTokenPriceInDollars(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) GetTokenPriceInDollars(opts *bind.CallOpts, token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getTokenPriceInDollars", token)

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

// GetTokenPriceInDollars is a free data retrieval call binding the contract method 0xe588566f.
//
// Solidity: function getTokenPriceInDollars(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) GetTokenPriceInDollars(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetTokenPriceInDollars(&_PriceFeed.CallOpts, token)
}

// GetTokenPriceInDollars is a free data retrieval call binding the contract method 0xe588566f.
//
// Solidity: function getTokenPriceInDollars(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) GetTokenPriceInDollars(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetTokenPriceInDollars(&_PriceFeed.CallOpts, token)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PriceFeed *PriceFeedCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PriceFeed *PriceFeedSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PriceFeed.Contract.HasRole(&_PriceFeed.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PriceFeed *PriceFeedCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PriceFeed.Contract.HasRole(&_PriceFeed.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeed.Contract.ProxiableUUID(&_PriceFeed.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PriceFeed *PriceFeedCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PriceFeed.Contract.ProxiableUUID(&_PriceFeed.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PriceFeed *PriceFeedCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PriceFeed *PriceFeedSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PriceFeed.Contract.SupportsInterface(&_PriceFeed.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PriceFeed *PriceFeedCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PriceFeed.Contract.SupportsInterface(&_PriceFeed.CallOpts, interfaceId)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRole(&_PriceFeed.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRole(&_PriceFeed.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactor) GrantRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "grantRoleBatch", roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRoleBatch(&_PriceFeed.TransactOpts, roles, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xb2b49e2e.
//
// Solidity: function grantRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactorSession) GrantRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRoleBatch(&_PriceFeed.TransactOpts, roles, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "initialize", owner, _dollarDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedSession) Initialize(owner common.Address, _dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, owner, _dollarDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 _dollarDecimals) returns()
func (_PriceFeed *PriceFeedTransactorSession) Initialize(owner common.Address, _dollarDecimals uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, owner, _dollarDecimals)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PriceFeed *PriceFeedTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PriceFeed *PriceFeedSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceRole(&_PriceFeed.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PriceFeed *PriceFeedTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RenounceRole(&_PriceFeed.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRole(&_PriceFeed.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PriceFeed *PriceFeedTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRole(&_PriceFeed.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactor) RevokeRoleBatch(opts *bind.TransactOpts, roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "revokeRoleBatch", roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRoleBatch(&_PriceFeed.TransactOpts, roles, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x365d71e4.
//
// Solidity: function revokeRoleBatch(bytes32[] roles, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactorSession) RevokeRoleBatch(roles [][32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRoleBatch(&_PriceFeed.TransactOpts, roles, accounts)
}

// UpdateNativeTokenPrice is a paid mutator transaction binding the contract method 0x9409f47b.
//
// Solidity: function updateNativeTokenPrice(uint256[] chainIDs, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedTransactor) UpdateNativeTokenPrice(opts *bind.TransactOpts, chainIDs []*big.Int, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "updateNativeTokenPrice", chainIDs, prices, pricesAt)
}

// UpdateNativeTokenPrice is a paid mutator transaction binding the contract method 0x9409f47b.
//
// Solidity: function updateNativeTokenPrice(uint256[] chainIDs, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedSession) UpdateNativeTokenPrice(chainIDs []*big.Int, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdateNativeTokenPrice(&_PriceFeed.TransactOpts, chainIDs, prices, pricesAt)
}

// UpdateNativeTokenPrice is a paid mutator transaction binding the contract method 0x9409f47b.
//
// Solidity: function updateNativeTokenPrice(uint256[] chainIDs, uint256[] prices, uint256[] pricesAt) returns()
func (_PriceFeed *PriceFeedTransactorSession) UpdateNativeTokenPrice(chainIDs []*big.Int, prices []*big.Int, pricesAt []*big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpdateNativeTokenPrice(&_PriceFeed.TransactOpts, chainIDs, prices, pricesAt)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeToAndCall(&_PriceFeed.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PriceFeed *PriceFeedTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PriceFeed.Contract.UpgradeToAndCall(&_PriceFeed.TransactOpts, newImplementation, data)
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

// PriceFeedNativeTokenPriceUpdatedIterator is returned from FilterNativeTokenPriceUpdated and is used to iterate over the raw logs and unpacked data for NativeTokenPriceUpdated events raised by the PriceFeed contract.
type PriceFeedNativeTokenPriceUpdatedIterator struct {
	Event *PriceFeedNativeTokenPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedNativeTokenPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedNativeTokenPriceUpdated)
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
		it.Event = new(PriceFeedNativeTokenPriceUpdated)
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
func (it *PriceFeedNativeTokenPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedNativeTokenPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedNativeTokenPriceUpdated represents a NativeTokenPriceUpdated event raised by the PriceFeed contract.
type PriceFeedNativeTokenPriceUpdated struct {
	ChainID *big.Int
	Price   *big.Int
	PriceAt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNativeTokenPriceUpdated is a free log retrieval operation binding the contract event 0x291ab1bcc002ef9062ceee3c7ea71d95374e2886b7ff336cb31a1e63e2649658.
//
// Solidity: event NativeTokenPriceUpdated(uint256 indexed chainID, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) FilterNativeTokenPriceUpdated(opts *bind.FilterOpts, chainID []*big.Int) (*PriceFeedNativeTokenPriceUpdatedIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "NativeTokenPriceUpdated", chainIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedNativeTokenPriceUpdatedIterator{contract: _PriceFeed.contract, event: "NativeTokenPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchNativeTokenPriceUpdated is a free log subscription operation binding the contract event 0x291ab1bcc002ef9062ceee3c7ea71d95374e2886b7ff336cb31a1e63e2649658.
//
// Solidity: event NativeTokenPriceUpdated(uint256 indexed chainID, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) WatchNativeTokenPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedNativeTokenPriceUpdated, chainID []*big.Int) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "NativeTokenPriceUpdated", chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedNativeTokenPriceUpdated)
				if err := _PriceFeed.contract.UnpackLog(event, "NativeTokenPriceUpdated", log); err != nil {
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

// ParseNativeTokenPriceUpdated is a log parse operation binding the contract event 0x291ab1bcc002ef9062ceee3c7ea71d95374e2886b7ff336cb31a1e63e2649658.
//
// Solidity: event NativeTokenPriceUpdated(uint256 indexed chainID, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) ParseNativeTokenPriceUpdated(log types.Log) (*PriceFeedNativeTokenPriceUpdated, error) {
	event := new(PriceFeedNativeTokenPriceUpdated)
	if err := _PriceFeed.contract.UnpackLog(event, "NativeTokenPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the PriceFeed contract.
type PriceFeedPriceUpdatedIterator struct {
	Event *PriceFeedPriceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceUpdated)
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
		it.Event = new(PriceFeedPriceUpdated)
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
func (it *PriceFeedPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceUpdated represents a PriceUpdated event raised by the PriceFeed contract.
type PriceFeedPriceUpdated struct {
	Token   common.Address
	Price   *big.Int
	PriceAt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed token, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) FilterPriceUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceFeedPriceUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceUpdatedIterator{contract: _PriceFeed.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed token, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceUpdated)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address indexed token, uint256 price, uint256 priceAt)
func (_PriceFeed *PriceFeedFilterer) ParsePriceUpdated(log types.Log) (*PriceFeedPriceUpdated, error) {
	event := new(PriceFeedPriceUpdated)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PriceFeed contract.
type PriceFeedRoleAdminChangedIterator struct {
	Event *PriceFeedRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceFeedRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedRoleAdminChanged)
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
		it.Event = new(PriceFeedRoleAdminChanged)
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
func (it *PriceFeedRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedRoleAdminChanged represents a RoleAdminChanged event raised by the PriceFeed contract.
type PriceFeedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PriceFeed *PriceFeedFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PriceFeedRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedRoleAdminChangedIterator{contract: _PriceFeed.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PriceFeed *PriceFeedFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedRoleAdminChanged)
				if err := _PriceFeed.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PriceFeed *PriceFeedFilterer) ParseRoleAdminChanged(log types.Log) (*PriceFeedRoleAdminChanged, error) {
	event := new(PriceFeedRoleAdminChanged)
	if err := _PriceFeed.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PriceFeed contract.
type PriceFeedRoleGrantedIterator struct {
	Event *PriceFeedRoleGranted // Event containing the contract specifics and raw log

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
func (it *PriceFeedRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedRoleGranted)
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
		it.Event = new(PriceFeedRoleGranted)
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
func (it *PriceFeedRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedRoleGranted represents a RoleGranted event raised by the PriceFeed contract.
type PriceFeedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceFeedRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedRoleGrantedIterator{contract: _PriceFeed.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PriceFeedRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedRoleGranted)
				if err := _PriceFeed.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) ParseRoleGranted(log types.Log) (*PriceFeedRoleGranted, error) {
	event := new(PriceFeedRoleGranted)
	if err := _PriceFeed.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PriceFeed contract.
type PriceFeedRoleRevokedIterator struct {
	Event *PriceFeedRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PriceFeedRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedRoleRevoked)
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
		it.Event = new(PriceFeedRoleRevoked)
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
func (it *PriceFeedRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedRoleRevoked represents a RoleRevoked event raised by the PriceFeed contract.
type PriceFeedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceFeedRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedRoleRevokedIterator{contract: _PriceFeed.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PriceFeedRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedRoleRevoked)
				if err := _PriceFeed.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PriceFeed *PriceFeedFilterer) ParseRoleRevoked(log types.Log) (*PriceFeedRoleRevoked, error) {
	event := new(PriceFeedRoleRevoked)
	if err := _PriceFeed.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PriceFeed contract.
type PriceFeedUpgradedIterator struct {
	Event *PriceFeedUpgraded // Event containing the contract specifics and raw log

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
func (it *PriceFeedUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedUpgraded)
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
		it.Event = new(PriceFeedUpgraded)
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
func (it *PriceFeedUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedUpgraded represents a Upgraded event raised by the PriceFeed contract.
type PriceFeedUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PriceFeedUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedUpgradedIterator{contract: _PriceFeed.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PriceFeedUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedUpgraded)
				if err := _PriceFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PriceFeed *PriceFeedFilterer) ParseUpgraded(log types.Log) (*PriceFeedUpgraded, error) {
	event := new(PriceFeedUpgraded)
	if err := _PriceFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
