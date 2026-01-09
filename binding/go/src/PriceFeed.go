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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChains\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chains\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dollarDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getNativeTokenPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData\",\"name\":\"data\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPriceDataList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceFeed.PriceData[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"exist\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceInDollars\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"dollarDecimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeNativeTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roles\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updateNativeTokenPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricesAt\",\"type\":\"uint256[]\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"NativeTokenPriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"NativeTokenPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceAt\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CalcAmountLibCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedInvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"name\":\"PriceFeedInvalidPriceAt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceFeedNoSource\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"RoleManagerDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoleManagerMissmatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"7f3f94b1": "allChains()",
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
		"6c237716": "removeNativeTokenPrice(uint256)",
		"f4961fa3": "removePrice(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"365d71e4": "revokeRoleBatch(bytes32[],address[])",
		"01ffc9a7": "supportsInterface(bytes4)",
		"9409f47b": "updateNativeTokenPrice(uint256[],uint256[],uint256[])",
		"99e6440d": "updatePrice(address[],uint256[],uint256[])",
		"7519ab50": "updatedAt()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f51602061222e5f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161216790816100c78239608051818181610cf60152610dc50152f35b6001600160401b0319166001600160401b039081175f51602061222e5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461110d57508063248a9ca3146110e75780632f2ff15d146110bb5780632fe3b6cf1461109c57806336568abe14611058578063365d71e414610ffe5780633afb154414610fdf5780634f1ef28614610d4a57806352d1902d14610ce45780636c23771614610c7e57806372279ba114610bd45780637519ab5014610bb75780637f3f94b114610b4d5780638fb5a48214610b1757806391d1485414610ac25780639409f47b14610962578063943b24b2146107cf5780639761cb221461062657806399e6440d1461048c578063a217fddf14610472578063a3246ad3146103b8578063ac41865a14610380578063ad3cb1cc14610322578063b2b49e2e146102b4578063d547741f1461027c578063e588566f14610239578063ec7221dd146101c85763f4961fa314610150575f80fd5b346101c45760203660031901126101c457610169611176565b610171611773565b6001600160a01b031661018381611d21565b50805f5260066020525f60026040822082815582600182015501557fb0e09cc2a1c729c1ad17b9273a92c742a5c22cac4285ff849736991d196bbf995f80a2005b5f80fd5b346101c4575f3660031901126101c457604051600254808252602082019060025f5260205f20905f5b8181106102235761021f6102108661020b818803826111bb565b611533565b604093919351938493846113a9565b0390f35b82548452602090930192600192830192016101f1565b346101c45760203660031901126101c45761021f61025d610258611176565b611720565b6040805193151584526020840192909252908201529081906060820190565b346101c45760403660031901126101c4576102b260043561029b611160565b906102ad6102a882611434565b6117e2565b611901565b005b346101c4576102c236611263565b908051825103610313575f5b82518110156102b2578061030c6102e760019385611483565b51838060a01b036102f88488611483565b5116906103076102a882611434565b6118a1565b50016102ce565b63031206d560e51b5f5260045ffd5b346101c4575f3660031901126101c457604080519061034181836111bb565b600582526020820191640352e302e360dc1b83528151928391602083525180918160208501528484015e5f828201840152601f01601f19168101030190f35b346101c45760403660031901126101c45760406103ac61039e611176565b6103a6611160565b90611657565b82519182526020820152f35b346101c45760203660031901126101c4576004355f525f5160206120d25f395f51905f5260205260405f206040518154808252602082019081935f5260205f20905f5b81811061045c57505050816104119103826111bb565b604051918291602083019060208452518091526040830191905f5b81811061043a575050500390f35b82516001600160a01b031684528594506020938401939092019160010161042c565b82548452602090930192600192830192016103fb565b346101c4575f3660031901126101c45760206040515f8152f35b346101c45760603660031901126101c4576004356001600160401b0381116101c4576104bc9036906004016111f5565b6024356001600160401b0381116101c4576104db9036906004016113d7565b906044356001600160401b0381116101c4576104fb9036906004016113d7565b91610504611773565b815191815183148061061c575b1561060d575f5b8381106105255742600155005b6001600160a01b036105378284611483565b5116906105448185611483565b51916105508288611483565b519280156105fe5760407fb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c09160019561058d428242811115611960565b806105f95750425b61059e85611faf565b5082516105aa8161118c565b85815260026020820184815285830190848252885f526006602052865f20938b8060a01b039051168b8060a01b0319855416178455518a8401555191015582519182526020820152a201610518565b610595565b63406571d160e01b5f5260045ffd5b63282e5b7160e11b5f5260045ffd5b5083518314610511565b346101c45760203660031901126101c4576004356001600160401b0381116101c4576106569036906004016111f5565b6106608151611501565b90805161066c816111de565b9061067a60405192836111bb565b808252610689601f19916111de565b015f5b8181106107b85750505f5b8251811015610756576001906001600160a01b036106b58286611483565b51165f526006602052818060a01b0360405f205416828060a01b036106da8387611483565b5116146106e8575b01610697565b816106f38287611483565b52818060a01b036107048286611483565b51165f52600660205260405f206002604051916107208361118c565b848060a01b038154168352848101546020840152015460408201526107458285611483565b526107508184611483565b506106e2565b61076f8483604051928392604084526040840190611374565b8281036020840152602080835192838152019201905f5b818110610794575050500390f35b9193509160206060826107aa600194885161131f565b019401910191849392610786565b6020906107c36114e3565b8282860101520161068c565b346101c45760403660031901126101c4576107e8611176565b60243560ff81168091036101c4575f5160206121125f395f51905f5254604081901c60ff161592906001600160401b0381168015908161095a575b6001149081610950575b159081610947575b50610938576001600160401b031981166001175f5160206121125f395f51905f525583610910575b506001600160a01b03811680156108fe57506108989061087b611e7d565b610883611e7d565b61088b611e7d565b610893611e7d565b611828565b5060ff195f5416175f55426001556108ac57005b60ff60401b195f5160206121125f395f51905f5254165f5160206121125f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b631a1f881960e11b5f5260045260245ffd5b6001600160481b0319166001600160401b01175f5160206121125f395f51905f52558361085d565b63f92ee8a960e01b5f5260045ffd5b90501585610835565b303b15915061082d565b859150610823565b346101c45760603660031901126101c4576004356001600160401b0381116101c4576109929036906004016113d7565b6024356001600160401b0381116101c4576109b19036906004016113d7565b906044356001600160401b0381116101c4576109d19036906004016113d7565b916109da611773565b8151918151831480610ab8575b1561060d575f5b8381106109fb5742600155005b610a058183611483565b5190610a118185611483565b5191610a1d8288611483565b519280156105fe5760407f291ab1bcc002ef9062ceee3c7ea71d95374e2886b7ff336cb31a1e63e264965891600195610a5a428242811115611960565b80610ab35750425b610a6b85611f5a565b508251610a778161118c565b85815260026020820184815285830190848252885f526007602052865f2093518455518a8401555191015582519182526020820152a2016109ee565b610a62565b50835183146109e7565b346101c45760403660031901126101c457610adb611160565b6004355f525f5160206120f25f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346101c45760203660031901126101c4576004356001600160401b0381116101c45761021061020b61021f9236906004016111f5565b346101c4575f3660031901126101c457604051600454808252602082019060045f5260205f20905f5b818110610ba15761021f85610b8d818703826111bb565b604051918291602083526020830190611341565b8254845260209093019260019283019201610b76565b346101c4575f3660031901126101c4576020600154604051908152f35b346101c45760203660031901126101c4576080610bef611176565b5f90610bf96114e3565b6001600160a01b039182165f8181526006602052604090205491929091168114610c37575b50610c35906040519215158352602083019061131f565bf35b9150506001905f526006602052610c3560405f20600260405191610c5a8361118c565b80546001600160a01b03168352600181015460208401520154604082015290610c1e565b346101c45760203660031901126101c457600435610c9a611773565b610ca381611c4c565b50805f5260076020525f60026040822082815582600182015501557fc7e360acbed133f1a9ebcf3f9d880638d323312c946698627544553572c180bb5f80a2005b346101c4575f3660031901126101c4577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610d3b5760206040515f5160206120b25f395f51905f528152f35b63703e46dd60e11b5f5260045ffd5b60403660031901126101c457610d5e611176565b602435906001600160401b0382116101c457366023830112156101c457816004013590610d8a82611304565b91610d9860405193846111bb565b808352602083019336602483830101116101c457815f926024602093018737840101526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610fbd575b50610d3b57335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff1615610f86576040516352d1902d60e01b81526001600160a01b0382169390602081600481885afa5f9181610f52575b50610e6d5784634c9c8ce360e01b5f5260045260245ffd5b805f5160206120b25f395f51905f52869203610f405750823b15610f2e575f5160206120b25f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2825115610f15575f80916102b2945190845af43d15610f0d573d91610ef183611304565b92610eff60405194856111bb565b83523d5f602085013e612053565b606091612053565b50505034610f1f57005b63b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f5260045260245ffd5b632a87526960e21b5f5260045260245ffd5b9091506020813d602011610f7e575b81610f6e602093836111bb565b810103126101c457519086610e55565b3d9150610f61565b63e2517d3f60e01b5f52336004527fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177560245260445ffd5b5f5160206120b25f395f51905f52546001600160a01b03161415905084610df0565b346101c45760203660031901126101c45761021f61025d600435611497565b346101c45761100c36611263565b908051825103610313575f5b82518110156102b2578061105161103160019385611483565b51838060a01b036110428488611483565b5116906102ad6102a882611434565b5001611018565b346101c45760403660031901126101c457611071611160565b336001600160a01b0382160361108d576102b290600435611901565b63334bd91960e11b5f5260045ffd5b346101c4575f3660031901126101c457602060ff5f5416604051908152f35b346101c45760403660031901126101c4576102b26004356110da611160565b906103076102a882611434565b346101c45760203660031901126101c4576020611105600435611434565b604051908152f35b346101c45760203660031901126101c4576004359063ffffffff60e01b82168092036101c457602091637965db0b60e01b811490811561114f575b5015158152f35b6301ffc9a760e01b14905083611148565b602435906001600160a01b03821682036101c457565b600435906001600160a01b03821682036101c457565b606081019081106001600160401b038211176111a757604052565b634e487b7160e01b5f52604160045260245ffd5b601f909101601f19168101906001600160401b038211908210176111a757604052565b6001600160401b0381116111a75760051b60200190565b9080601f830112156101c45781359061120d826111de565b9261121b60405194856111bb565b82845260208085019360051b8201019182116101c457602001915b8183106112435750505090565b82356001600160a01b03811681036101c457815260209283019201611236565b9060406003198301126101c4576004356001600160401b0381116101c457826023820112156101c45780600401359061129b826111de565b916112a960405193846111bb565b8083526024602084019160051b830101918583116101c457602401905b8282106112f457509193602435925090506001600160401b0382116101c4576112f1916004016111f5565b90565b81358152602091820191016112c6565b6001600160401b0381116111a757601f01601f191660200190565b80516001600160a01b0316825260208082015190830152604090810151910152565b90602080835192838152019201905f5b81811061135e5750505090565b8251845260209384019390920191600101611351565b90602080835192838152019201905f5b8181106113915750505090565b82511515845260209384019390920191600101611384565b9392916113d2906113c4604093606088526060880190611374565b908682036020880152611341565b930152565b9080601f830112156101c45781356113ee816111de565b926113fc60405194856111bb565b81845260208085019260051b8201019283116101c457602001905b8282106114245750505090565b8135815260209182019101611417565b5f525f5160206120f25f395f51905f52602052600160405f20015490565b80511561145f5760200190565b634e487b7160e01b5f52603260045260245ffd5b80516001101561145f5760400190565b805182101561145f5760209160051b010190565b90815f52600760205260405f20916040516114b18161118c565b835490818152604060026001870154968760208501520154910152036114da5760018054909291565b5f915081908190565b604051906114f08261118c565b5f6040838281528260208201520152565b9061150b826111de565b61151860405191826111bb565b8281528092611529601f19916111de565b0190602036910137565b9061153e8251611501565b9180519061154b826111de565b9161155960405193846111bb565b808352611568601f19916111de565b01366020840137815f5b82518110156115ff576001906001600160a01b036115908286611483565b51165f526006602052818060a01b0360405f205416828060a01b036115b58387611483565b5116146115c3575b01611572565b816115ce8289611483565b52818060a01b036115df8286611483565b51165f5260066020528160405f2001546115f98287611483565b526115bd565b509291505060015490565b60ff16604d811161161b57600a0a90565b634e487b7160e01b5f52601160045260245ffd5b156116375750565b634ced742360e11b5f9081526001600160a01b0391909116600452602490fd5b906001600160a01b0380831690821681811461170c576116da611705946116d5856116b86116c897966116e096604051916116936060846111bb565b6002835260403660208501376116a883611452565b526116b282611473565b52611533565b5097906116d0856116c883611452565b51151561162f565b611473565b61197e565b9161197e565b916116ea8261160a565b906116fe6116f782611452565b5191611473565b5191611a29565b9060015490565b50505061171b6117059161197e565b61160a565b60018060a01b031690815f52600660205260405f20916040516117428161118c565b60018060a01b0384541690818152604060026001870154968760208501520154910152036114da5760018054909291565b335f9081527f30d0252998c0e58f26c68d721d21aa6a1e2f173bf671484c59546e6750e12a58602052604090205460ff16156117ab57565b63e2517d3f60e01b5f52336004527fc6823861ee2bb2198ce6b1fd6faf4c8f44f745bc804aca4a762f67e0d507fd8a60245260445ffd5b5f8181525f5160206120f25f395f51905f526020908152604080832033845290915290205460ff16156118125750565b63e2517d3f60e01b5f523360045260245260445ffd5b90611833825f611af7565b918261183c5750565b5f80525f5160206120d25f395f51905f526020526001600160a01b0316611883817f2945b4696643b83a462ccb7d71bd065c2c9aff76eab9a316df70754985bf692d611fff565b1561188b5750565b63d180cb3160e01b5f526004525f60245260445ffd5b9190916118ae8382611af7565b92836118b8575050565b815f525f5160206120d25f395f51905f526020526118e360405f209160018060a01b03168092611fff565b156118ec575050565b63d180cb3160e01b5f5260045260245260445ffd5b91909161190e8382611b9b565b9283611918575050565b815f525f5160206120d25f395f51905f5260205261194360405f209160018060a01b03168092611dcc565b1561194c575050565b62a95f1b60e31b5f5260045260245260445ffd5b15611969575050565b63248fc98560e11b5f5260045260245260445ffd5b6001600160a01b0316600181036119955750601290565b60206004916040519283809263313ce56760e01b82525afa9081156119f7575f916119be575090565b90506020813d6020116119ef575b816119d9602093836111bb565b810103126101c4575160ff811681036101c45790565b3d91506119cc565b6040513d5f823e3d90fd5b9060ff8091169116039060ff821161161b57565b8181029291811591840414171561161b57565b929091938215611ac8578415611a995760ff811660ff831690808214611a8b571115611a715791611a6561171b611a6b936112f1979695611a02565b90611a16565b90611ea8565b93611a6561171b6112f196611a8594611a02565b91611ea8565b505050506112f19291611ea8565b6040516349671cd160e11b8152602060048201526006602482015265383934b1b2a160d11b6044820152606490fd5b6040516349671cd160e11b815260206004820152600660248201526570726963654160d01b6044820152606490fd5b5f8181525f5160206120f25f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16611b95575f8181525f5160206120f25f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f8181525f5160206120f25f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615611b95575f8181525f5160206120f25f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b805482101561145f575f5260205f2001905f90565b5f818152600560205260409020548015611b95575f19810181811161161b576004545f1981019190821161161b57818103611cd3575b5050506004548015611cbf575f1901611c9c816004611c37565b8154905f199060031b1b191690556004555f5260056020525f6040812055600190565b634e487b7160e01b5f52603160045260245ffd5b611d0b611ce4611cf5936004611c37565b90549060031b1c9283926004611c37565b819391549060031b91821b915f19901b19161790565b90555f52600560205260405f20555f8080611c82565b5f818152600360205260409020548015611b95575f19810181811161161b576002545f1981019190821161161b57818103611d94575b5050506002548015611cbf575f1901611d71816002611c37565b8154905f199060031b1b191690556002555f5260036020525f6040812055600190565b611db6611da5611cf5936002611c37565b90549060031b1c9283926002611c37565b90555f52600360205260405f20555f8080611d57565b906001820191815f528260205260405f20548015155f14611e75575f19810181811161161b5782545f1981019190821161161b57818103611e40575b50505080548015611cbf575f190190611e218282611c37565b8154905f199060031b1b19169055555f526020525f6040812055600190565b611e60611e50611cf59386611c37565b90549060031b1c92839286611c37565b90555f528360205260405f20555f8080611e08565b505050505f90565b60ff5f5160206121125f395f51905f525460401c1615611e9957565b631afcd79f60e31b5f5260045ffd5b91818302915f1981850993838086109503948086039514611f385784831115611f205790829109815f0382168092046002816003021880820260020302808202600203028082026002030280820260020302808202600203028091026002030293600183805f03040190848311900302920304170290565b82634e487b715f52156003026011186020526024601cfd5b505080925015611f46570490565b634e487b7160e01b5f52601260045260245ffd5b805f52600560205260405f2054155f14611faa57600454600160401b8110156111a757611f93611cf58260018594016004556004611c37565b9055600454905f52600560205260405f2055600190565b505f90565b805f52600360205260405f2054155f14611faa57600254600160401b8110156111a757611fe8611cf58260018594016002556002611c37565b9055600254905f52600360205260405f2055600190565b6001810190825f528160205260405f2054155f1461204c578054600160401b8110156111a757612039611cf5826001879401855584611c37565b905554915f5260205260405f2055600190565b5050505f90565b90612077575080511561206857805190602001fd5b63d6bda27560e01b5f5260045ffd5b815115806120a8575b612088575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561208056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f0002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220ae9e82202b0ed43ad27f3f1535e27645dc483c5911a4f627c758c737d2e32ef164736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint256[] chains)
func (_PriceFeed *PriceFeedCaller) AllChains(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PriceFeed.contract.Call(opts, &out, "allChains")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint256[] chains)
func (_PriceFeed *PriceFeedSession) AllChains() ([]*big.Int, error) {
	return _PriceFeed.Contract.AllChains(&_PriceFeed.CallOpts)
}

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint256[] chains)
func (_PriceFeed *PriceFeedCallerSession) AllChains() ([]*big.Int, error) {
	return _PriceFeed.Contract.AllChains(&_PriceFeed.CallOpts)
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
// Solidity: function initialize(address owner, uint8 dollarDecimals_) returns()
func (_PriceFeed *PriceFeedTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, dollarDecimals_ uint8) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "initialize", owner, dollarDecimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 dollarDecimals_) returns()
func (_PriceFeed *PriceFeedSession) Initialize(owner common.Address, dollarDecimals_ uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, owner, dollarDecimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 dollarDecimals_) returns()
func (_PriceFeed *PriceFeedTransactorSession) Initialize(owner common.Address, dollarDecimals_ uint8) (*types.Transaction, error) {
	return _PriceFeed.Contract.Initialize(&_PriceFeed.TransactOpts, owner, dollarDecimals_)
}

// RemoveNativeTokenPrice is a paid mutator transaction binding the contract method 0x6c237716.
//
// Solidity: function removeNativeTokenPrice(uint256 chainID) returns()
func (_PriceFeed *PriceFeedTransactor) RemoveNativeTokenPrice(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "removeNativeTokenPrice", chainID)
}

// RemoveNativeTokenPrice is a paid mutator transaction binding the contract method 0x6c237716.
//
// Solidity: function removeNativeTokenPrice(uint256 chainID) returns()
func (_PriceFeed *PriceFeedSession) RemoveNativeTokenPrice(chainID *big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveNativeTokenPrice(&_PriceFeed.TransactOpts, chainID)
}

// RemoveNativeTokenPrice is a paid mutator transaction binding the contract method 0x6c237716.
//
// Solidity: function removeNativeTokenPrice(uint256 chainID) returns()
func (_PriceFeed *PriceFeedTransactorSession) RemoveNativeTokenPrice(chainID *big.Int) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemoveNativeTokenPrice(&_PriceFeed.TransactOpts, chainID)
}

// RemovePrice is a paid mutator transaction binding the contract method 0xf4961fa3.
//
// Solidity: function removePrice(address token) returns()
func (_PriceFeed *PriceFeedTransactor) RemovePrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PriceFeed.contract.Transact(opts, "removePrice", token)
}

// RemovePrice is a paid mutator transaction binding the contract method 0xf4961fa3.
//
// Solidity: function removePrice(address token) returns()
func (_PriceFeed *PriceFeedSession) RemovePrice(token common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemovePrice(&_PriceFeed.TransactOpts, token)
}

// RemovePrice is a paid mutator transaction binding the contract method 0xf4961fa3.
//
// Solidity: function removePrice(address token) returns()
func (_PriceFeed *PriceFeedTransactorSession) RemovePrice(token common.Address) (*types.Transaction, error) {
	return _PriceFeed.Contract.RemovePrice(&_PriceFeed.TransactOpts, token)
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

// PriceFeedNativeTokenPriceRemovedIterator is returned from FilterNativeTokenPriceRemoved and is used to iterate over the raw logs and unpacked data for NativeTokenPriceRemoved events raised by the PriceFeed contract.
type PriceFeedNativeTokenPriceRemovedIterator struct {
	Event *PriceFeedNativeTokenPriceRemoved // Event containing the contract specifics and raw log

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
func (it *PriceFeedNativeTokenPriceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedNativeTokenPriceRemoved)
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
		it.Event = new(PriceFeedNativeTokenPriceRemoved)
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
func (it *PriceFeedNativeTokenPriceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedNativeTokenPriceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedNativeTokenPriceRemoved represents a NativeTokenPriceRemoved event raised by the PriceFeed contract.
type PriceFeedNativeTokenPriceRemoved struct {
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNativeTokenPriceRemoved is a free log retrieval operation binding the contract event 0xc7e360acbed133f1a9ebcf3f9d880638d323312c946698627544553572c180bb.
//
// Solidity: event NativeTokenPriceRemoved(uint256 indexed chainID)
func (_PriceFeed *PriceFeedFilterer) FilterNativeTokenPriceRemoved(opts *bind.FilterOpts, chainID []*big.Int) (*PriceFeedNativeTokenPriceRemovedIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "NativeTokenPriceRemoved", chainIDRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedNativeTokenPriceRemovedIterator{contract: _PriceFeed.contract, event: "NativeTokenPriceRemoved", logs: logs, sub: sub}, nil
}

// WatchNativeTokenPriceRemoved is a free log subscription operation binding the contract event 0xc7e360acbed133f1a9ebcf3f9d880638d323312c946698627544553572c180bb.
//
// Solidity: event NativeTokenPriceRemoved(uint256 indexed chainID)
func (_PriceFeed *PriceFeedFilterer) WatchNativeTokenPriceRemoved(opts *bind.WatchOpts, sink chan<- *PriceFeedNativeTokenPriceRemoved, chainID []*big.Int) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "NativeTokenPriceRemoved", chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedNativeTokenPriceRemoved)
				if err := _PriceFeed.contract.UnpackLog(event, "NativeTokenPriceRemoved", log); err != nil {
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

// ParseNativeTokenPriceRemoved is a log parse operation binding the contract event 0xc7e360acbed133f1a9ebcf3f9d880638d323312c946698627544553572c180bb.
//
// Solidity: event NativeTokenPriceRemoved(uint256 indexed chainID)
func (_PriceFeed *PriceFeedFilterer) ParseNativeTokenPriceRemoved(log types.Log) (*PriceFeedNativeTokenPriceRemoved, error) {
	event := new(PriceFeedNativeTokenPriceRemoved)
	if err := _PriceFeed.contract.UnpackLog(event, "NativeTokenPriceRemoved", log); err != nil {
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

// PriceFeedPriceRemovedIterator is returned from FilterPriceRemoved and is used to iterate over the raw logs and unpacked data for PriceRemoved events raised by the PriceFeed contract.
type PriceFeedPriceRemovedIterator struct {
	Event *PriceFeedPriceRemoved // Event containing the contract specifics and raw log

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
func (it *PriceFeedPriceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedPriceRemoved)
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
		it.Event = new(PriceFeedPriceRemoved)
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
func (it *PriceFeedPriceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedPriceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedPriceRemoved represents a PriceRemoved event raised by the PriceFeed contract.
type PriceFeedPriceRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceRemoved is a free log retrieval operation binding the contract event 0xb0e09cc2a1c729c1ad17b9273a92c742a5c22cac4285ff849736991d196bbf99.
//
// Solidity: event PriceRemoved(address indexed token)
func (_PriceFeed *PriceFeedFilterer) FilterPriceRemoved(opts *bind.FilterOpts, token []common.Address) (*PriceFeedPriceRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.FilterLogs(opts, "PriceRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedPriceRemovedIterator{contract: _PriceFeed.contract, event: "PriceRemoved", logs: logs, sub: sub}, nil
}

// WatchPriceRemoved is a free log subscription operation binding the contract event 0xb0e09cc2a1c729c1ad17b9273a92c742a5c22cac4285ff849736991d196bbf99.
//
// Solidity: event PriceRemoved(address indexed token)
func (_PriceFeed *PriceFeedFilterer) WatchPriceRemoved(opts *bind.WatchOpts, sink chan<- *PriceFeedPriceRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceFeed.contract.WatchLogs(opts, "PriceRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedPriceRemoved)
				if err := _PriceFeed.contract.UnpackLog(event, "PriceRemoved", log); err != nil {
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

// ParsePriceRemoved is a log parse operation binding the contract event 0xb0e09cc2a1c729c1ad17b9273a92c742a5c22cac4285ff849736991d196bbf99.
//
// Solidity: event PriceRemoved(address indexed token)
func (_PriceFeed *PriceFeedFilterer) ParsePriceRemoved(log types.Log) (*PriceFeedPriceRemoved, error) {
	event := new(PriceFeedPriceRemoved)
	if err := _PriceFeed.contract.UnpackLog(event, "PriceRemoved", log); err != nil {
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
