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
	ABI: "[{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold_\",\"type\":\"uint8\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDollarPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPriceDataList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_dollarDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"resetRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"set\",\"type\":\"bool\"}],\"name\":\"setRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceFeedPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNoSource\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleAlreadyGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleNotGranted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ValidatorInsufficientSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ValidatorInvalidSignatures\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"ec7221dd": "allPrices()",
		"b7f3358d": "changeThreshold(uint8)",
		"2fe3b6cf": "dollarDecimals()",
		"84b0196e": "eip712Domain()",
		"e4eeca6f": "getDollarPrice(address)",
		"ac41865a": "getPrice(address,address)",
		"72279ba1": "getPriceData(address)",
		"9761cb22": "getPriceDataList(address[])",
		"8fb5a482": "getPrices(address[])",
		"a3246ad3": "getRoleMembers(bytes32)",
		"91d14854": "hasRole(bytes32,address)",
		"4351e6b6": "initialize(uint8)",
		"e1758bd8": "nativeToken()",
		"8da5cb5b": "owner()",
		"52d1902d": "proxiableUUID()",
		"715018a6": "renounceOwnership()",
		"2d87b7ee": "resetRole(bytes32,address[])",
		"d4bf502a": "setRole(bytes32,address[],bool)",
		"42cde4e8": "threshold()",
		"f2fde38b": "transferOwnership(address)",
		"99e6440d": "updatePrice(address[],uint256[],uint256[])",
		"7519ab50": "updatedAt()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a0604052306080523480156012575f5ffd5b506080516125f56100395f395f81816112370152818161126001526113a401526125f55ff3fe608060405260043610610147575f3560e01c806391d14854116100b3578063b7f3358d1161006d578063b7f3358d1461041a578063d4bf502a14610439578063e1758bd814610458578063e4eeca6f1461046b578063ec7221dd146104a7578063f2fde38b146104bb575f5ffd5b806391d14854146103025780639761cb221461033157806399e6440d1461035e578063a3246ad31461037d578063ac41865a146103a9578063ad3cb1cc146103dd575f5ffd5b8063715018a611610104578063715018a61461020757806372279ba11461021b5780637519ab501461024857806384b0196e1461025d5780638da5cb5b146102845780638fb5a482146102d4575f5ffd5b80632d87b7ee1461014b5780632fe3b6cf1461016c57806342cde4e81461019c5780634351e6b6146101b35780634f1ef286146101d257806352d1902d146101e5575b5f5ffd5b348015610156575f5ffd5b5061016a610165366004611d43565b6104da565b005b348015610177575f5ffd5b506064546101859060ff1681565b60405160ff90911681526020015b60405180910390f35b3480156101a7575f5ffd5b5060325460ff16610185565b3480156101be575f5ffd5b5061016a6101cd366004611d94565b61050c565b61016a6101e0366004611daf565b610709565b3480156101f0575f5ffd5b506101f9610724565b604051908152602001610193565b348015610212575f5ffd5b5061016a61073f565b348015610226575f5ffd5b5061023a610235366004611e52565b610752565b604051610193929190611e6b565b348015610253575f5ffd5b506101f960655481565b348015610268575f5ffd5b506102716107da565b6040516101939796959493929190611f06565b34801561028f575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610193565b3480156102df575f5ffd5b506102f36102ee366004611f75565b610888565b60405161019393929190611fe0565b34801561030d575f5ffd5b5061032161031c366004612015565b610a1e565b6040519015158152602001610193565b34801561033c575f5ffd5b5061035061034b366004611f75565b610a3e565b60405161019392919061203f565b348015610369575f5ffd5b5061016a61037836600461211a565b610c1f565b348015610388575f5ffd5b5061039c6103973660046121a5565b610d0b565b60405161019391906121bc565b3480156103b4575f5ffd5b506103c86103c3366004612207565b610d24565b60408051928352602083019190915201610193565b3480156103e8575f5ffd5b5061040d604051806040016040528060058152602001640352e302e360dc1b81525081565b604051610193919061222f565b348015610425575f5ffd5b5061016a610434366004611d94565b61105f565b348015610444575f5ffd5b5061016a610453366004612241565b6110af565b348015610463575f5ffd5b5060016102bc565b348015610476575f5ffd5b5061048a610485366004611e52565b6110ec565b604080519315158452602084019290925290820152606001610193565b3480156104b2575f5ffd5b506102f3611141565b3480156104c6575f5ffd5b5061016a6104d5366004611e52565b61115d565b6104fc826104f65f5f8681526020019081526020015f2061119a565b5f6110af565b610508828260016110af565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156105505750825b90505f826001600160401b0316600114801561056b5750303b155b905081158015610579575080155b156105975760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156105c157845460ff60401b1916600160401b1785555b6105ca336111ad565b6105d35f6111be565b6105db611224565b6064805460ff191660ff8816908117909155426065556040805160608101909152600181529060208201906103e89061061590600a612392565b61061f91906123a0565b81524260209182015260015f526068815281517f82eaf0fca2207f91f5027fcf68136c84edb7e928c081c42aa5bbc2a771c7d37680546001600160a01b0319166001600160a01b039092169190911790558101517f82eaf0fca2207f91f5027fcf68136c84edb7e928c081c42aa5bbc2a771c7d37755604001517f82eaf0fca2207f91f5027fcf68136c84edb7e928c081c42aa5bbc2a771c7d37855831561070157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b61071161122c565b61071a826112d0565b61050882826112d8565b5f61072d611399565b505f5160206125a05f395f51905f5290565b6107476113e2565b6107505f61143d565b565b604080516060810182525f808252602082018190529181018290526001600160a01b038084165f81815260686020526040902054909116036107d55750506001600160a01b038082165f90815260686020908152604091829020825160608101845281549094168452600181810154928501929092526002015491830191909152905b915091565b5f60608082808083815f5160206125805f395f51905f52805490915015801561080557506001810154155b61084e5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b6108566114ad565b61085e61156d565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b6060805f83516001600160401b038111156108a5576108a5611c51565b6040519080825280602002602001820160405280156108ce578160200160208202803683370190505b50925083516001600160401b038111156108ea576108ea611c51565b604051908082528060200260200182016040528015610913578160200160208202803683370190505b5091505f5b8451811015610a1257848181518110610933576109336123bf565b60200260200101516001600160a01b031660685f878481518110610959576109596123bf565b6020908102919091018101516001600160a01b039081168352908201929092526040015f20541603610a0a576001848281518110610999576109996123bf565b60200260200101901515908115158152505060685f8683815181106109c0576109c06123bf565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20600101548382815181106109fd576109fd6123bf565b6020026020010181815250505b600101610918565b50506065549193909250565b5f828152602081905260408120610a3590836115ab565b90505b92915050565b60608082516001600160401b03811115610a5a57610a5a611c51565b604051908082528060200260200182016040528015610a83578160200160208202803683370190505b50915082516001600160401b03811115610a9f57610a9f611c51565b604051908082528060200260200182016040528015610afa57816020015b610ae760405180606001604052805f6001600160a01b031681526020015f81526020015f81525090565b815260200190600190039081610abd5790505b5090505f5b8351811015610c1957838181518110610b1a57610b1a6123bf565b60200260200101516001600160a01b031660685f868481518110610b4057610b406123bf565b6020908102919091018101516001600160a01b039081168352908201929092526040015f20541603610c11576001838281518110610b8057610b806123bf565b60200260200101901515908115158152505060685f858381518110610ba757610ba76123bf565b6020908102919091018101516001600160a01b0390811683528282019390935260409182015f2082516060810184528154909416845260018101549184019190915260020154908201528251839083908110610c0557610c056123bf565b60200260200101819052505b600101610aff565b50915091565b610c35682b20a624a220aa27a960b91b33610a1e565b682b20a624a220aa27a960b91b339091610c745760405163698bd5ff60e01b815260048101929092526001600160a01b03166024820152604401610845565b50508151835114610c985760405163282e5b7160e11b815260040160405180910390fd5b5f5b8351811015610d0157610cf9848281518110610cb857610cb86123bf565b6020026020010151848381518110610cd257610cd26123bf565b6020026020010151848481518110610cec57610cec6123bf565b60200260200101516115cc565b600101610c9a565b5050426065555050565b5f818152602081905260409020606090610a389061119a565b5f5f826001600160a01b0316846001600160a01b031603610d4a57506001905043611058565b6040805160028082526060820183525f9260208301908036833701905050905084815f81518110610d7d57610d7d6123bf565b60200260200101906001600160a01b031690816001600160a01b0316815250508381600181518110610db157610db16123bf565b60200260200101906001600160a01b031690816001600160a01b031681525050606080610ddd83610888565b8251909650919350915082905f90610df757610df76123bf565b60200260200101518790610e2a57604051634ced742360e11b81526001600160a01b039091166004820152602401610845565b5081600181518110610e3e57610e3e6123bf565b60200260200101518690610e7157604051634ced742360e11b81526001600160a01b039091166004820152602401610845565b50604051630fe74f3b60e01b81523060048201526001600160a01b03881660248201525f90819073__$79c24c681325f76413d25d3c06c8219b6e$__90630fe74f3b90604401602060405180830381865af4158015610ed2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef691906123d3565b604051630fe74f3b60e01b81523060048201526001600160a01b038a16602482015273__$79c24c681325f76413d25d3c06c8219b6e$__90630fe74f3b90604401602060405180830381865af4158015610f52573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f7691906123d3565b9150915073__$79c24c681325f76413d25d3c06c8219b6e$__63b1db08d16001855f81518110610fa857610fa86123bf565b602002602001015186600181518110610fc357610fc36123bf565b60209081029190910101516040516001600160e01b031960e086901b16815260048101939093526024830191909152604482015260ff80861660648301528416608482015260a401602060405180830381865af4158015611026573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061104a91906123ee565b606554909750955050505050505b9250929050565b6110676113e2565b6032805460ff191660ff83169081179091556040519081527f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff9060200160405180910390a150565b5f5b82518110156110e6576110de848483815181106110d0576110d06123bf565b6020026020010151846116e1565b6001016110b1565b50505050565b6001600160a01b038082165f8181526068602052604081205490928392839290911614611136576001600160a01b0384165f90815260686020526040902060019081015490935091505b506065549193909250565b6060805f6111526102ee606661119a565b925092509250909192565b6111656113e2565b6001600160a01b03811661118e57604051631e4fbdf760e01b81525f6004820152602401610845565b6111978161143d565b50565b60605f6111a683611811565b9392505050565b6111b561186a565b611197816118b3565b6111c661186a565b61120e604051806040016040528060098152602001682b30b634b230ba37b960b91b815250604051806040016040528060058152602001640312e302e360dc1b8152506118bb565b6032805460ff191660ff92909216919091179055565b61075061186a565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112b257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112a65f5160206125a05f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156107505760405163703e46dd60e11b815260040160405180910390fd5b6111976113e2565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611332575060408051601f3d908101601f1916820190925261132f918101906123ee565b60015b61135a57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610845565b5f5160206125a05f395f51905f52811461138a57604051632a87526960e21b815260048101829052602401610845565b61139483836118cd565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107505760405163703e46dd60e11b815260040160405180910390fd5b336114147f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107505760405163118cdaa760e01b8152336004820152602401610845565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206125805f395f51905f52916114eb90612405565b80601f016020809104026020016040519081016040528092919081815260200182805461151790612405565b80156115625780601f1061153957610100808354040283529160200191611562565b820191905f5260205f20905b81548152906001019060200180831161154557829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206125805f395f51905f52916114eb90612405565b6001600160a01b0381165f9081526001830160205260408120541515610a35565b815f0361160457604051637ae6ed6d60e01b8152602060048201526005602482015264707269636560d81b6044820152606401610845565b5f198114806116135750428111155b8142909161163d5760405163248fc98560e11b815260048101929092526024820152604401610845565b5061164b9050606684611922565b50604080516060810182526001600160a01b0385811680835260208084018781528486018781525f84815260688452879020955186546001600160a01b03191695169490941785555160018501559151600290930192909255825185815290810184905290917fc4552bb17fb4b7a4ff96683f5fc7545352dafb23ed81de0ac8b12709a82bfde2910160405180910390a2505050565b6116e96113e2565b6001600160a01b03821661171057604051635989b9d360e01b815260040160405180910390fd5b8261172e57604051630e1b39f960e31b815260040160405180910390fd5b8015611785575f83815260208190526040902061174b9083611922565b8383909161177e57604051631b753c1b60e21b815260048101929092526001600160a01b03166024820152604401610845565b50506117d2565b5f83815260208190526040902061179c9083611936565b838390916117cf57604051633a401ef360e21b815260048101929092526001600160a01b03166024820152604401610845565b50505b80151583836001600160a01b03167f8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f360405160405180910390a4505050565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561185e57602002820191905f5260205f20905b81548152602001906001019080831161184a575b50505050509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661075057604051631afcd79f60e31b815260040160405180910390fd5b61116561186a565b6118c361186a565b610508828261194a565b6118d6826119a9565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561191a576113948282611a0c565b610508611a7e565b5f610a35836001600160a01b038416611a9d565b5f610a35836001600160a01b038416611ae9565b61195261186a565b5f5160206125805f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10261198b8482612488565b506003810161199a8382612488565b505f8082556001909101555050565b806001600160a01b03163b5f036119de57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610845565b5f5160206125a05f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051611a289190612542565b5f60405180830381855af49150503d805f8114611a60576040519150601f19603f3d011682016040523d82523d5f602084013e611a65565b606091505b5091509150611a75858383611bcc565b95945050505050565b34156107505760405163b398979f60e01b815260040160405180910390fd5b5f818152600183016020526040812054611ae257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a38565b505f610a38565b5f8181526001830160205260408120548015611bc3575f611b0b600183612558565b85549091505f90611b1e90600190612558565b9050808214611b7d575f865f018281548110611b3c57611b3c6123bf565b905f5260205f200154905080875f018481548110611b5c57611b5c6123bf565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611b8e57611b8e61256b565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a38565b5f915050610a38565b606082611be157611bdc82611c28565b6111a6565b8151158015611bf857506001600160a01b0384163b155b15611c2157604051639996b31560e01b81526001600160a01b0385166004820152602401610845565b5092915050565b805115611c385780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611c8d57611c8d611c51565b604052919050565b5f6001600160401b03821115611cad57611cad611c51565b5060051b60200190565b80356001600160a01b0381168114611ccd575f5ffd5b919050565b5f82601f830112611ce1575f5ffd5b8135611cf4611cef82611c95565b611c65565b8082825260208201915060208360051b860101925085831115611d15575f5ffd5b602085015b83811015611d3957611d2b81611cb7565b835260209283019201611d1a565b5095945050505050565b5f5f60408385031215611d54575f5ffd5b8235915060208301356001600160401b03811115611d70575f5ffd5b611d7c85828601611cd2565b9150509250929050565b60ff81168114611197575f5ffd5b5f60208284031215611da4575f5ffd5b81356111a681611d86565b5f5f60408385031215611dc0575f5ffd5b611dc983611cb7565b915060208301356001600160401b03811115611de3575f5ffd5b8301601f81018513611df3575f5ffd5b80356001600160401b03811115611e0c57611e0c611c51565b611e1f601f8201601f1916602001611c65565b818152866020838501011115611e33575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215611e62575f5ffd5b610a3582611cb7565b8215158152608081016111a6602083018480516001600160a01b0316825260208082015190830152604090810151910152565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b82811015611efc578151865260209586019590910190600101611ede565b5093949350505050565b60ff60f81b8816815260e060208201525f611f2460e0830189611e9e565b8281036040840152611f368189611e9e565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050611f678185611ecc565b9a9950505050505050505050565b5f60208284031215611f85575f5ffd5b81356001600160401b03811115611f9a575f5ffd5b611fa684828501611cd2565b949350505050565b5f8151808452602084019350602083015f5b82811015611efc5781511515865260209586019590910190600101611fc0565b606081525f611ff26060830186611fae565b82810360208401526120048186611ecc565b915050826040830152949350505050565b5f5f60408385031215612026575f5ffd5b8235915061203660208401611cb7565b90509250929050565b604081525f6120516040830185611fae565b82810360208401528084518083526020830191506020860192505f5b818110156120b35761209d83855180516001600160a01b0316825260208082015190830152604090810151910152565b602093909301926060929092019160010161206d565b50909695505050505050565b5f82601f8301126120ce575f5ffd5b81356120dc611cef82611c95565b8082825260208201915060208360051b8601019250858311156120fd575f5ffd5b602085015b83811015611d39578035835260209283019201612102565b5f5f5f6060848603121561212c575f5ffd5b83356001600160401b03811115612141575f5ffd5b61214d86828701611cd2565b93505060208401356001600160401b03811115612168575f5ffd5b612174868287016120bf565b92505060408401356001600160401b0381111561218f575f5ffd5b61219b868287016120bf565b9150509250925092565b5f602082840312156121b5575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b818110156121fc5783516001600160a01b03168352602093840193909201916001016121d5565b509095945050505050565b5f5f60408385031215612218575f5ffd5b61222183611cb7565b915061203660208401611cb7565b602081525f610a356020830184611e9e565b5f5f5f60608486031215612253575f5ffd5b8335925060208401356001600160401b0381111561226f575f5ffd5b61227b86828701611cd2565b92505060408401358015158114612290575f5ffd5b809150509250925092565b634e487b7160e01b5f52601160045260245ffd5b6001815b60018411156122ea578085048111156122ce576122ce61229b565b60018416156122dc57908102905b60019390931c9280026122b3565b935093915050565b5f8261230057506001610a38565b8161230c57505f610a38565b8160018114612322576002811461232c57612348565b6001915050610a38565b60ff84111561233d5761233d61229b565b50506001821b610a38565b5060208310610133831016604e8410600b841016171561236b575081810a610a38565b6123775f1984846122af565b805f190482111561238a5761238a61229b565b029392505050565b5f610a3560ff8416836122f2565b5f826123ba57634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156123e3575f5ffd5b81516111a681611d86565b5f602082840312156123fe575f5ffd5b5051919050565b600181811c9082168061241957607f821691505b60208210810361243757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561139457805f5260205f20601f840160051c810160208510156124625750805b601f840160051c820191505b81811015612481575f815560010161246e565b5050505050565b81516001600160401b038111156124a1576124a1611c51565b6124b5816124af8454612405565b8461243d565b6020601f8211600181146124e7575f83156124d05750848201515b5f19600385901b1c1916600184901b178455612481565b5f84815260208120601f198516915b8281101561251657878501518255602094850194600190920191016124f6565b508482101561253357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b81810381811115610a3857610a3861229b565b634e487b7160e01b5f52603160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122012875b294cc5d073ac08a9bcfb79e7f4e7e283f78f86fb690fdde2b64e04004164736f6c634300081c0033",
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

// GetDollarPrice is a free data retrieval call binding the contract method 0xe4eeca6f.
//
// Solidity: function getDollarPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCaller) GetDollarPrice(opts *bind.CallOpts, token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "getDollarPrice", token)

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

// GetDollarPrice is a free data retrieval call binding the contract method 0xe4eeca6f.
//
// Solidity: function getDollarPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedSession) GetDollarPrice(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetDollarPrice(&_PriceFeed.CallOpts, token)
}

// GetDollarPrice is a free data retrieval call binding the contract method 0xe4eeca6f.
//
// Solidity: function getDollarPrice(address token) view returns(bool exist, uint256 price, uint256 updatedAt_)
func (_PriceFeed *PriceFeedCallerSession) GetDollarPrice(token common.Address) (struct {
	Exist     bool
	Price     *big.Int
	UpdatedAt *big.Int
}, error) {
	return _PriceFeed.Contract.GetDollarPrice(&_PriceFeed.CallOpts, token)
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

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() pure returns(address)
func (_PriceFeed *PriceFeedCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() pure returns(address)
func (_PriceFeed *PriceFeedSession) NativeToken() (common.Address, error) {
	return _PriceFeed.Contract.NativeToken(&_PriceFeed.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() pure returns(address)
func (_PriceFeed *PriceFeedCallerSession) NativeToken() (common.Address, error) {
	return _PriceFeed.Contract.NativeToken(&_PriceFeed.CallOpts)
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

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_PriceFeed *PriceFeedTransactor) ResetRole(opts *bind.TransactOpts, role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "resetRole", role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_PriceFeed *PriceFeedSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.ResetRole(&_PriceFeed.TransactOpts, role, newAccounts)
}

// ResetRole is a paid mutator transaction binding the contract method 0x2d87b7ee.
//
// Solidity: function resetRole(bytes32 role, address[] newAccounts) returns()
func (_PriceFeed *PriceFeedTransactorSession) ResetRole(role [32]byte, newAccounts []common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.ResetRole(&_PriceFeed.TransactOpts, role, newAccounts)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_PriceFeed *PriceFeedTransactor) SetRole(opts *bind.TransactOpts, role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "setRole", role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_PriceFeed *PriceFeedSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetRole(&_PriceFeed.TransactOpts, role, accounts, set)
}

// SetRole is a paid mutator transaction binding the contract method 0xd4bf502a.
//
// Solidity: function setRole(bytes32 role, address[] accounts, bool set) returns()
func (_PriceFeed *PriceFeedTransactorSession) SetRole(role [32]byte, accounts []common.Address, set bool) (*types.Transaction, error) {
	return _PriceFeed.Contract.SetRole(&_PriceFeed.TransactOpts, role, accounts, set)
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

// PriceFeedRoleUpdatedIterator is returned from FilterRoleUpdated and is used to iterate over the raw logs and unpacked data for RoleUpdated events raised by the PriceFeed contract.
type PriceFeedRoleUpdatedIterator struct {
	Event *PriceFeedRoleUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedRoleUpdated)
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
		it.Event = new(PriceFeedRoleUpdated)
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
func (it *PriceFeedRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedRoleUpdated represents a RoleUpdated event raised by the PriceFeed contract.
type PriceFeedRoleUpdated struct {
	Account common.Address
	Role    [32]byte
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleUpdated is a free log retrieval operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_PriceFeed *PriceFeedFilterer) FilterRoleUpdated(opts *bind.FilterOpts, account []common.Address, role [][32]byte, status []bool) (*PriceFeedRoleUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedRoleUpdatedIterator{contract: _PriceFeed.contract, event: "RoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleUpdated is a free log subscription operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_PriceFeed *PriceFeedFilterer) WatchRoleUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedRoleUpdated, account []common.Address, role [][32]byte, status []bool) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "RoleUpdated", accountRule, roleRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedRoleUpdated)
				if err := _PriceFeed.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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

// ParseRoleUpdated is a log parse operation binding the contract event 0x8737757d0849a3cc97f98ff695ca69b8c01a8e5cb9e21b72d5e776a9dd5b06f3.
//
// Solidity: event RoleUpdated(address indexed account, bytes32 indexed role, bool indexed status)
func (_PriceFeed *PriceFeedFilterer) ParseRoleUpdated(log types.Log) (*PriceFeedRoleUpdated, error) {
	event := new(PriceFeedRoleUpdated)
	if err := _PriceFeed.contract.UnpackLog(event, "RoleUpdated", log); err != nil {
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
