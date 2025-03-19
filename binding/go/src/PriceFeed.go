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
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNativeTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPriceDataList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceInDollars\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_dollarDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updateNativeTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNoSource\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
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
		"fd47852b": "grantRoleBatch(bytes32,address[])",
		"91d14854": "hasRole(bytes32,address)",
		"943b24b2": "initialize(address,uint8)",
		"52d1902d": "proxiableUUID()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"82f51fa6": "revokeRoleBatch(bytes32,address[])",
		"01ffc9a7": "supportsInterface(bytes4)",
		"9409f47b": "updateNativeTokenPrice(uint256[],uint256[],uint256[])",
		"99e6440d": "updatePrice(address[],uint256[],uint256[])",
		"7519ab50": "updatedAt()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516121546100395f395f81816112140152818161123d015261138001526121545ff3fe608060405260043610610161575f3560e01c806391d14854116100cd578063a3246ad311610087578063d547741f11610062578063d547741f1461048e578063e588566f146104ad578063ec7221dd146104cc578063fd47852b146104e0575f5ffd5b8063a3246ad3146103f1578063ac41865a1461041d578063ad3cb1cc14610451575f5ffd5b806391d14854146103355780639409f47b14610354578063943b24b2146103735780639761cb221461039257806399e6440d146103bf578063a217fddf146103de575f5ffd5b80634f1ef2861161011e5780634f1ef2861461027f57806352d1902d1461029257806372279ba1146102a65780637519ab50146102d357806382f51fa6146102e85780638fb5a48214610307575f5ffd5b806301ffc9a714610165578063248a9ca3146101995780632f2ff15d146101c65780632fe3b6cf146101e757806336568abe146102125780633afb154414610231575b5f5ffd5b348015610170575f5ffd5b5061018461017f366004611a6a565b6104ff565b60405190151581526020015b60405180910390f35b3480156101a4575f5ffd5b506101b86101b3366004611a91565b610535565b604051908152602001610190565b3480156101d1575f5ffd5b506101e56101e0366004611ac3565b610555565b005b3480156101f2575f5ffd5b506032546102009060ff1681565b60405160ff9091168152602001610190565b34801561021d575f5ffd5b506101e561022c366004611ac3565b610577565b34801561023c575f5ffd5b5061026261024b366004611a91565b5f9081526037602052604090205460335481151592565b604080519315158452602084019290925290820152606001610190565b6101e561028d366004611b31565b6105af565b34801561029d575f5ffd5b506101b86105ce565b3480156102b1575f5ffd5b506102c56102c0366004611bd4565b6105e9565b604051610190929190611bed565b3480156102de575f5ffd5b506101b860335481565b3480156102f3575f5ffd5b506101e5610302366004611cb3565b610671565b348015610312575f5ffd5b50610326610321366004611cf6565b6106ba565b60405161019093929190611d6b565b348015610340575f5ffd5b5061018461034f366004611ac3565b610850565b34801561035f575f5ffd5b506101e561036e366004611e26565b610886565b34801561037e575f5ffd5b506101e561038d366004611ebf565b61093d565b34801561039d575f5ffd5b506103b16103ac366004611cf6565b610a84565b604051610190929190611ef4565b3480156103ca575f5ffd5b506101e56103d9366004611f74565b610c65565b3480156103e9575f5ffd5b506101b85f81565b3480156103fc575f5ffd5b5061041061040b366004611a91565b610d10565b6040516101909190611fa7565b348015610428575f5ffd5b5061043c610437366004611ff2565b610d29565b60408051928352602083019190915201610190565b34801561045c575f5ffd5b50610481604051806040016040528060058152602001640352e302e360dc1b81525081565b604051610190919061201a565b348015610499575f5ffd5b506101e56104a8366004611ac3565b61105d565b3480156104b8575f5ffd5b506102626104c7366004611bd4565b611079565b3480156104d7575f5ffd5b506103266110d8565b3480156104eb575f5ffd5b506101e56104fa366004611cb3565b6110f4565b5f6001600160e01b03198216637965db0b60e01b148061052f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f9081525f5160206120ff5f395f51905f52602052604090206001015490565b61055e82610535565b6105678161113d565b610571838361114a565b50505050565b6001600160a01b03811633146105a05760405163334bd91960e11b815260040160405180910390fd5b6105aa82826111ae565b505050565b6105b7611209565b6105c0826112af565b6105ca82826112b9565b5050565b5f6105d7611375565b505f5160206120df5f395f51905f5290565b604080516060810182525f808252602082018190529181018290526001600160a01b038084165f818152603660205260409020549091160361066c5750506001600160a01b038082165f90815260366020908152604091829020825160608101845281549094168452600181810154928501929092526002015491830191909152905b915091565b61067a82610535565b6106838161113d565b5f5b8251811015610571576106b1848483815181106106a4576106a461204f565b60200260200101516111ae565b50600101610685565b6060805f83516001600160401b038111156106d7576106d7611aed565b604051908082528060200260200182016040528015610700578160200160208202803683370190505b50925083516001600160401b0381111561071c5761071c611aed565b604051908082528060200260200182016040528015610745578160200160208202803683370190505b5091505f5b8451811015610844578481815181106107655761076561204f565b60200260200101516001600160a01b031660365f87848151811061078b5761078b61204f565b6020908102919091018101516001600160a01b039081168352908201929092526040015f2054160361083c5760018482815181106107cb576107cb61204f565b60200260200101901515908115158152505060365f8683815181106107f2576107f261204f565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206001015483828151811061082f5761082f61204f565b6020026020010181815250505b60010161074a565b50506033549193909250565b5f9182525f5160206120ff5f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b662aa82220aa27a960c91b61089a8161113d565b83518351811480156108ac5750825181145b6108c95760405163282e5b7160e11b815260040160405180910390fd5b5f5b81811015610931576109298682815181106108e8576108e861204f565b60200260200101518683815181106109025761090261204f565b602002602001015186848151811061091c5761091c61204f565b60200260200101516113be565b6001016108cb565b50504260335550505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156109815750825b90505f826001600160401b0316600114801561099c5750303b155b9050811580156109aa575080155b156109c85760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109f257845460ff60401b1916600160401b1785555b6109fa611436565b610a02611436565b610a0c5f8861114a565b50610a21662aa82220aa27a960c91b8861114a565b506032805460ff191660ff8816179055426033558315610a7b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60608082516001600160401b03811115610aa057610aa0611aed565b604051908082528060200260200182016040528015610ac9578160200160208202803683370190505b50915082516001600160401b03811115610ae557610ae5611aed565b604051908082528060200260200182016040528015610b4057816020015b610b2d60405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b815260200190600190039081610b035790505b5090505f5b8351811015610c5f57838181518110610b6057610b6061204f565b60200260200101516001600160a01b031660365f868481518110610b8657610b8661204f565b6020908102919091018101516001600160a01b039081168352908201929092526040015f20541603610c57576001838281518110610bc657610bc661204f565b60200260200101901515908115158152505060365f858381518110610bed57610bed61204f565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f2082516060810184528154909416845260018101549184019190915260020154908201528251839083908110610c4b57610c4b61204f565b60200260200101819052505b600101610b45565b50915091565b662aa82220aa27a960c91b610c798161113d565b8351835181148015610c8b5750825181145b610ca85760405163282e5b7160e11b815260040160405180910390fd5b5f5b8181101561093157610d08868281518110610cc757610cc761204f565b6020026020010151868381518110610ce157610ce161204f565b6020026020010151868481518110610cfb57610cfb61204f565b602002602001015161143e565b600101610caa565b5f81815260208190526040902060609061052f90611546565b5f5f826001600160a01b0316846001600160a01b031603610d4f57506001905043611056565b6040805160028082526060820183525f9260208301908036833701905050905084815f81518110610d8257610d8261204f565b60200260200101906001600160a01b031690816001600160a01b0316815250508381600181518110610db657610db661204f565b60200260200101906001600160a01b031690816001600160a01b031681525050606080610de2836106ba565b8251909650919350915082905f90610dfc57610dfc61204f565b60200260200101518790610e3457604051634ced742360e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b5081600181518110610e4857610e4861204f565b60200260200101518690610e7b57604051634ced742360e11b81526001600160a01b039091166004820152602401610e2b565b50604051636a24d41960e11b81526001600160a01b03881660048201525f90819073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9063d449a83290602401602060405180830381865af4158015610ed6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610efa9190612063565b604051636a24d41960e11b81526001600160a01b038a16600482015273__$c8dc1c3a159d88c2746a5586ef67caa4e3$__9063d449a83290602401602060405180830381865af4158015610f50573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f749190612063565b9150915073__$c8dc1c3a159d88c2746a5586ef67caa4e3$__63b1db08d16001855f81518110610fa657610fa661204f565b602002602001015186600181518110610fc157610fc161204f565b60209081029190910101516040516001600160e01b031960e086901b16815260048101939093526024830191909152604482015260ff80861660648301528416608482015260a401602060405180830381865af4158015611024573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611048919061207e565b603354909750955050505050505b9250929050565b61106682610535565b61106f8161113d565b61057183836111ae565b6001600160a01b038082165f81815260366020526040812054909283928392909116146110ad57505f9150819050806110d1565b5050506001600160a01b0381165f9081526036602052604090206001908101546033545b9193909250565b6060805f6110e96103216034611546565b925092509250909192565b6110fd82610535565b6111068161113d565b5f5b825181101561057157611134848483815181106111275761112761204f565b602002602001015161114a565b50600101611108565b6111478133611559565b50565b5f6111558383611592565b9050801561052f575f8381526020819052604090206111749083611633565b828490916111a65760405163d180cb3160e01b81526001600160a01b0390921660048301526024820152604401610e2b565b505092915050565b5f6111b98383611647565b9050801561052f575f8381526020819052604090206111d890836116c0565b828490916111a65760405162a95f1b60e31b81526001600160a01b0390921660048301526024820152604401610e2b565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061128f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112835f5160206120df5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156112ad5760405163703e46dd60e11b815260040160405180910390fd5b565b5f6105ca8161113d565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611313575060408051601f3d908101601f191682019092526113109181019061207e565b60015b61133b57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610e2b565b5f5160206120df5f395f51905f52811461136b57604051632a87526960e21b815260048101829052602401610e2b565b6105aa83836116d4565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112ad5760405163703e46dd60e11b815260040160405180910390fd5b815f036113f657604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b6044820152606401610e2b565b8042808211156114225760405163248fc98560e11b815260048101929092526024820152604401610e2b565b5050505f9182526037602052604090912055565b6112ad611729565b815f0361147657604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b6044820152606401610e2b565b8042808211156114a25760405163248fc98560e11b815260048101929092526024820152604401610e2b565b506114b09050603484611633565b50604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260368452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2910160405180910390a2505050565b60605f61155283611772565b9392505050565b6115638282610850565b6105ca5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610e2b565b5f5f5160206120ff5f395f51905f526115ab8484610850565b61162a575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556115e03390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061052f565b5f91505061052f565b5f611552836001600160a01b0384166117cb565b5f5f5160206120ff5f395f51905f526116608484610850565b1561162a575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061052f565b5f611552836001600160a01b038416611817565b6116dd826118f1565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611721576105aa8282611954565b6105ca6119c6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166112ad57604051631afcd79f60e31b815260040160405180910390fd5b6060815f018054806020026020016040519081016040528092919081815260200182805480156117bf57602002820191905f5260205f20905b8154815260200190600101908083116117ab575b50505050509050919050565b5f81815260018301602052604081205461181057508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561052f565b505f61052f565b5f818152600183016020526040812054801561162a575f611839600183612095565b85549091505f9061184c90600190612095565b90508082146118ab575f865f01828154811061186a5761186a61204f565b905f5260205f200154905080875f01848154811061188a5761188a61204f565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806118bc576118bc6120b4565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061052f565b806001600160a01b03163b5f0361192657604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610e2b565b5f5160206120df5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161197091906120c8565b5f60405180830381855af49150503d805f81146119a8576040519150601f19603f3d011682016040523d82523d5f602084013e6119ad565b606091505b50915091506119bd8583836119e5565b95945050505050565b34156112ad5760405163b398979f60e01b815260040160405180910390fd5b6060826119fa576119f582611a41565b611552565b8151158015611a1157506001600160a01b0384163b155b15611a3a57604051639996b31560e01b81526001600160a01b0385166004820152602401610e2b565b5092915050565b805115611a515780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f60208284031215611a7a575f5ffd5b81356001600160e01b031981168114611552575f5ffd5b5f60208284031215611aa1575f5ffd5b5035919050565b80356001600160a01b0381168114611abe575f5ffd5b919050565b5f5f60408385031215611ad4575f5ffd5b82359150611ae460208401611aa8565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611b2957611b29611aed565b604052919050565b5f5f60408385031215611b42575f5ffd5b611b4b83611aa8565b915060208301356001600160401b03811115611b65575f5ffd5b8301601f81018513611b75575f5ffd5b80356001600160401b03811115611b8e57611b8e611aed565b611ba1601f8201601f1916602001611b01565b818152866020838501011115611bb5575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215611be4575f5ffd5b61155282611aa8565b821515815260808101611552602083018480516001600160a01b0316825260208082015190830152604090810151910152565b5f6001600160401b03821115611c3857611c38611aed565b5060051b60200190565b5f82601f830112611c51575f5ffd5b8135611c64611c5f82611c20565b611b01565b8082825260208201915060208360051b860101925085831115611c85575f5ffd5b602085015b83811015611ca957611c9b81611aa8565b835260209283019201611c8a565b5095945050505050565b5f5f60408385031215611cc4575f5ffd5b8235915060208301356001600160401b03811115611ce0575f5ffd5b611cec85828601611c42565b9150509250929050565b5f60208284031215611d06575f5ffd5b81356001600160401b03811115611d1b575f5ffd5b611d2784828501611c42565b949350505050565b5f8151808452602084019350602083015f5b82811015611d615781511515865260209586019590910190600101611d41565b5093949350505050565b606081525f611d7d6060830186611d2f565b82810360208401528085518083526020830191506020870192505f5b81811015611db7578351835260209384019390920191600101611d99565b505060409390930193909352509392505050565b5f82601f830112611dda575f5ffd5b8135611de8611c5f82611c20565b8082825260208201915060208360051b860101925085831115611e09575f5ffd5b602085015b83811015611ca9578035835260209283019201611e0e565b5f5f5f60608486031215611e38575f5ffd5b83356001600160401b03811115611e4d575f5ffd5b611e5986828701611dcb565b93505060208401356001600160401b03811115611e74575f5ffd5b611e8086828701611dcb565b92505060408401356001600160401b03811115611e9b575f5ffd5b611ea786828701611dcb565b9150509250925092565b60ff81168114611147575f5ffd5b5f5f60408385031215611ed0575f5ffd5b611ed983611aa8565b91506020830135611ee981611eb1565b809150509250929050565b604081525f611f066040830185611d2f565b82810360208401528084518083526020830191506020860192505f5b81811015611f6857611f5283855180516001600160a01b0316825260208082015190830152604090810151910152565b6020939093019260609290920191600101611f22565b50909695505050505050565b5f5f5f60608486031215611f86575f5ffd5b83356001600160401b03811115611f9b575f5ffd5b611e5986828701611c42565b602080825282518282018190525f918401906040840190835b81811015611fe75783516001600160a01b0316835260209384019390920191600101611fc0565b509095945050505050565b5f5f60408385031215612003575f5ffd5b61200c83611aa8565b9150611ae460208401611aa8565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215612073575f5ffd5b815161155281611eb1565b5f6020828403121561208e575f5ffd5b5051919050565b8181038181111561052f57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a2646970667358221220c2a7523ee8b245a06b35c1a71cccc8deb593165eab6bece70ac824df619f9f4864736f6c634300081c0033",
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

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactor) GrantRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "grantRoleBatch", role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRoleBatch(&_PriceFeed.TransactOpts, role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactorSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.GrantRoleBatch(&_PriceFeed.TransactOpts, role, accounts)
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

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactor) RevokeRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "revokeRoleBatch", role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRoleBatch(&_PriceFeed.TransactOpts, role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_PriceFeed *PriceFeedTransactorSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RevokeRoleBatch(&_PriceFeed.TransactOpts, role, accounts)
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
