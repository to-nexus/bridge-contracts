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

// CrossMintableERC20V2MetaData contains all meta data concerning the CrossMintableERC20V2 contract.
var CrossMintableERC20V2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"3644e515": "DOMAIN_SEPARATOR()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"9dc29fac": "burn(address,uint256)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"313ce567": "decimals()",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"84b0196e": "eip712Domain()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"8420ce99": "initialize(address,address,string,string,uint8)",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"8da5cb5b": "owner()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"01ffc9a7": "supportsInterface(bytes4)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x6080806040523460aa575f51602061273d5f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b60405161268e90816100af8239f35b6001600160401b0319166001600160401b039081175f51602061273d5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a71461188c57508063022d63fb1461186f57806306fdde03146117c5578063095ea7b31461179f5780630aa6220b146116fc57806318160ddd146116d357806323b872dd146115fb578063248a9ca3146115dd5780632f2ff15d146115a6578063313ce5671461157a5780633644e5151461155857806336568abe1461145d57806340c10f19146113aa578063634e93da146112da578063649a5ec71461115a57806370a08231146111165780637ecebe00146110d25780638420ce99146108ee57806384b0196e146107ca57806384ef8ffc146107c55780638da5cb5b146107c557806391d148541461077057806395d89b411461068d5780639dc29fac146105ae578063a1eda53c14610549578063a217fddf1461052f578063a9059cbb146104fe578063cc8463c8146104d4578063cefc142914610410578063cf6eefb7146103d6578063d505accf1461028c578063d547741f1461023f578063d602b9fd146101dd5763dd62ed3e14610192575f80fd5b346101d95760403660031901126101d9576101ab61191d565b6101bc6101b6611933565b91611bca565b9060018060a01b03165f52602052602060405f2054604051908152f35b5f80fd5b346101d9575f3660031901126101d9576101f5611c7e565b65ffffffffffff610204611c53565b5f5160206124d95f395f51905f5280546001600160d01b031916905591909116905061022c57005b5f5160206125195f395f51905f525f80a1005b346101d95760403660031901126101d95760043561025b611933565b811561027d578161027661027161027b94611bac565b611d29565b612008565b005b631fe1e13d60e11b5f5260045ffd5b346101d95760e03660031901126101d9576102a561191d565b6102ad611933565b60443590606435926102bd6119c1565b8442116103c3576103886103919160018060a01b03841696875f525f5160206124b95f395f51905f5260205260405f20908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c0815261035660e082611949565b519020610361611f53565b906040519161190160f01b83526002830152602282015260c43591604260a435922061215c565b909291926121df565b6001600160a01b03168481036103ac575061027b9350611ef0565b84906325c0072360e11b5f5260045260245260445ffd5b8463313c898160e11b5f5260045260245ffd5b346101d9575f3660031901126101d957604065ffffffffffff6103f7611c53565b83516001600160a01b0390921682529091166020820152f35b346101d9575f3660031901126101d957610428611c53565b506001600160a01b031633036104c15765ffffffffffff610447611c53565b9190911690811580156104b7575b6104a4575f5160206125d95f395f51905f5254610485919061047f906001600160a01b0316611fba565b50611e66565b505f5160206124d95f395f51905f5280546001600160d01b0319169055005b506319ca5ebb60e01b5f5260045260245ffd5b5042821015610455565b636116401160e11b5f523360045260245ffd5b346101d9575f3660031901126101d95760206104ee611c02565b65ffffffffffff60405191168152f35b346101d95760403660031901126101d95761052461051a61191d565b6024359033611d6f565b602060405160018152f35b346101d9575f3660031901126101d95760206040515f8152f35b346101d9575f3660031901126101d9575f5160206125d95f395f51905f52548060d01c90811515806105a4575b1561059a5760a01c65ffffffffffff165b61059660405192839283611a05565b0390f35b50505f5f90610587565b5042821015610576565b346101d95760403660031901126101d9576105c761191d565b602435906105d3611ccd565b6001600160a01b0316801561067a57805f525f5160206123d95f395f51905f5260205260405f2054828110610661576020835f945f5160206125595f395f51905f52938587525f5160206123d95f395f51905f528452036040862055805f5160206124395f395f51905f5254035f5160206124395f395f51905f5255604051908152a3602060405160018152f35b9063391434e360e21b5f5260045260245260445260645ffd5b634b637e8f60e11b5f525f60045260245ffd5b346101d9575f3660031901126101d9576040515f5f5160206124195f395f51905f52546106b981611a1e565b808452906001811690811561074c57506001146106f5575b610596836106e181850382611949565b6040519182916020835260208301906118f9565b5f5160206124195f395f51905f525f9081525f5160206125b95f395f51905f52939250905b808210610732575090915081016020016106e16106d1565b91926001816020925483858801015201910190929161071a565b60ff191660208086019190915291151560051b840190910191506106e190506106d1565b346101d95760403660031901126101d957610789611933565b6004355f525f5160206125795f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b6119d1565b346101d9575f3660031901126101d9575f5160206124795f395f51905f525415806108d8575b1561089b5761083f610800611a56565b610808611b12565b602061084d6040519261081b8385611949565b5f84525f368137604051958695600f60f81b875260e08588015260e08701906118f9565b9085820360408701526118f9565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b82811061088457505050500390f35b835185528695509381019392810192600101610875565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f5160206126195f395f51905f5254156107f0565b346101d95760a03660031901126101d95761090761191d565b61090f611933565b906044356001600160401b0381116101d95761092f90369060040161196c565b6064356001600160401b0381116101d95761094e90369060040161196c565b926109576119c1565b5f5160206125f95f395f51905f5254604081901c60ff161595919391906001600160401b038116801590816110ca575b60011490816110c0575b1590816110b7575b506110a8576001600160401b031981166001175f5160206125f95f395f51905f525586611080575b506109ca61208d565b6109d261208d565b81516001600160401b038111610da9576109f95f5160206123b95f395f51905f5254611a1e565b601f8111611024575b50806020601f8211600114610fa8575f91610f9d575b508160011b915f199060031b1c1916175f5160206123b95f395f51905f52555b8051906001600160401b038211610da9578190610a625f5160206124195f395f51905f5254611a1e565b601f8111610f36575b50602090601f8311600114610eb8575f92610ead575b50508160011b915f199060031b1c1916175f5160206124195f395f51905f52555b610aaa61208d565b60405190610ab9604083611949565b60018252603160f81b6020830152610acf61208d565b8051906001600160401b038211610da9578190610af95f5160206123f95f395f51905f5254611a1e565b601f8111610e46575b50602090601f8311600114610dc8575f92610dbd575b50508160011b915f199060031b1c1916175f5160206123f95f395f51905f52555b8051906001600160401b038211610da9578190610b635f5160206124595f395f51905f5254611a1e565b601f8111610d42575b50602090601f8311600114610cc4575f92610cb9575b50508160011b915f199060031b1c1916175f5160206124595f395f51905f52555b5f5f5160206124795f395f51905f52555f5f5160206126195f395f51905f5255610bcb61208d565b610bd361208d565b6001600160a01b03831615610ca6575f5160206124d95f395f51905f5280546001600160d01b0316905560ff92610c0990611e66565b506001600160a01b038116610c96575b501660ff195f5160206123995f395f51905f525416175f5160206123995f395f51905f5255610c4457005b60ff60401b195f5160206125f95f395f51905f5254165f5160206125f95f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b610c9f90611e98565b5083610c19565b636116401160e11b5f525f60045260245ffd5b015190508680610b82565b5f5160206124595f395f51905f525f9081528281209350601f198516905b818110610d2a5750908460019594939210610d12575b505050811b015f5160206124595f395f51905f5255610ba3565b01515f1960f88460031b161c19169055868080610cf8565b92936020600181928786015181550195019301610ce2565b5f5160206124595f395f51905f525f529091505f5160206126395f395f51905f52601f840160051c81019160208510610d9f575b90601f859493920160051c01905b818110610d915750610b6c565b5f8155849350600101610d84565b9091508190610d76565b634e487b7160e01b5f52604160045260245ffd5b015190508780610b18565b5f5160206123f95f395f51905f525f9081528281209350601f198516905b818110610e2e5750908460019594939210610e16575b505050811b015f5160206123f95f395f51905f5255610b39565b01515f1960f88460031b161c19169055878080610dfc565b92936020600181928786015181550195019301610de6565b5f5160206123f95f395f51905f525f529091505f5160206124995f395f51905f52601f840160051c81019160208510610ea3575b90601f859493920160051c01905b818110610e955750610b02565b5f8155849350600101610e88565b9091508190610e7a565b015190508780610a81565b5f5160206124195f395f51905f525f9081528281209350601f198516905b818110610f1e5750908460019594939210610f06575b505050811b015f5160206124195f395f51905f5255610aa2565b01515f1960f88460031b161c19169055878080610eec565b92936020600181928786015181550195019301610ed6565b5f5160206124195f395f51905f525f529091505f5160206125b95f395f51905f52601f840160051c81019160208510610f93575b90601f859493920160051c01905b818110610f855750610a6b565b5f8155849350600101610f78565b9091508190610f6a565b905083015188610a18565b5f5160206123b95f395f51905f525f9081528181209250601f198416905b81811061100c57509083600194939210610ff4575b5050811b015f5160206123b95f395f51905f5255610a38565b8501515f1960f88460031b161c191690558880610fdb565b9192602060018192868a015181550194019201610fc6565b5f5160206123b95f395f51905f525f525f5160206123795f395f51905f52601f830160051c81019160208410611076575b601f0160051c01905b81811061106b5750610a02565b5f815560010161105e565b9091508190611055565b6001600160481b0319166001600160401b01175f5160206125f95f395f51905f5255866109c1565b63f92ee8a960e01b5f5260045ffd5b90501588610999565b303b159150610991565b889150610987565b346101d95760203660031901126101d9576001600160a01b036110f361191d565b165f525f5160206124b95f395f51905f52602052602060405f2054604051908152f35b346101d95760203660031901126101d9576001600160a01b0361113761191d565b165f525f5160206123d95f395f51905f52602052602060405f2054604051908152f35b346101d95760203660031901126101d95760043565ffffffffffff8116908181036101d957611187611c7e565b6111904261205e565b9165ffffffffffff6111a0611c02565b168082111561129f57505f5160206125995f395f51905f529265ffffffffffff8262069780806111da951091180262069780181690611e1a565b905f5160206125d95f395f51905f52548060d01c8061124a575b50505f5160206125d95f395f51905f5280546001600160d01b031960d085901b166001600160a01b0390911665ffffffffffff60a01b60a085901b1617179055604051918291611245919083611a05565b0390a1005b421115611288575f5160206124d95f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b83806111f4565b505f5160206124f95f395f51905f525f80a1611281565b0365ffffffffffff81116112c6575f5160206125995f395f51905f52926111da9190611e1a565b634e487b7160e01b5f52601160045260245ffd5b346101d95760203660031901126101d9576112f361191d565b6112fb611c7e565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6602061133861132a4261205e565b611332611c02565b90611e1a565b65ffffffffffff611347611c53565b5f5160206124d95f395f51905f5280546001600160d01b0319166001600160a01b0390981697881785851660a01b179055919091169050611394575b65ffffffffffff60405191168152a2005b5f5160206125195f395f51905f525f80a1611383565b346101d95760403660031901126101d9576113c361191d565b602435906113cf611ccd565b6001600160a01b031690811561144a575f5160206124395f395f51905f5254908082018092116112c65760205f5160206125595f395f51905f52915f935f5160206124395f395f51905f52558484525f5160206123d95f395f51905f52825260408420818154019055604051908152a3602060405160018152f35b63ec442f0560e01b5f525f60045260245ffd5b346101d95760403660031901126101d957600435611479611933565b811580611535575b6114ae575b336001600160a01b0382160361149f5761027b91612008565b63334bd91960e11b5f5260045ffd5b6114b6611c53565b906001600160a01b031615801590611525575b8015611513575b6114f857505f5160206124d95f395f51905f52805465ffffffffffff60a01b19169055611486565b65ffffffffffff906319ca5ebb60e01b5f521660045260245ffd5b504265ffffffffffff821610156114d0565b5065ffffffffffff8116156114c9565b505f5160206125d95f395f51905f52546001600160a01b03828116911614611481565b346101d9575f3660031901126101d9576020611572611f53565b604051908152f35b346101d9575f3660031901126101d957602060ff5f5160206123995f395f51905f525416604051908152f35b346101d95760403660031901126101d9576004356115c2611933565b811561027d57816115d861027161027b94611bac565b611eaf565b346101d95760203660031901126101d9576020611572600435611bac565b346101d95760603660031901126101d95761161461191d565b61161c611933565b6044359061162983611bca565b335f9081526020919091526040902054925f19841061164d575b6105249350611d6f565b8284106116b8576001600160a01b038116156116a5573315611692576105249361167682611bca565b60018060a01b0333165f526020528360405f2091039055611643565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8284637dc7a0d960e11b5f523360045260245260445260645ffd5b346101d9575f3660031901126101d95760205f5160206124395f395f51905f5254604051908152f35b346101d9575f3660031901126101d957611714611c7e565b5f5160206125d95f395f51905f52548060d01c8061174a575b5f5160206125d95f395f51905f5280546001600160a01b03169055005b421115611788575f5160206124d95f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b808061172d565b505f5160206124f95f395f51905f525f80a1611781565b346101d95760403660031901126101d9576105246117bb61191d565b6024359033611ef0565b346101d9575f3660031901126101d9576040515f5f5160206123b95f395f51905f52546117f181611a1e565b808452906001811690811561074c575060011461181857610596836106e181850382611949565b5f5160206123b95f395f51905f525f9081525f5160206123795f395f51905f52939250905b808210611855575090915081016020016106e16106d1565b91926001816020925483858801015201910190929161183d565b346101d9575f3660031901126101d9576020604051620697808152f35b346101d95760203660031901126101d9576004359063ffffffff60e01b82168092036101d9576020916318a4c3c360e11b81149081156118ce575b5015158152f35b637965db0b60e01b8114915081156118e8575b50836118c7565b6301ffc9a760e01b149050836118e1565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b03821682036101d957565b602435906001600160a01b03821682036101d957565b601f909101601f19168101906001600160401b03821190821017610da957604052565b81601f820112156101d9578035906001600160401b038211610da957604051926119a0601f8401601f191660200185611949565b828452602083830101116101d957815f926020809301838601378301015290565b6084359060ff821682036101d957565b346101d9575f3660031901126101d9575f5160206125d95f395f51905f52546040516001600160a01b039091168152602090f35b65ffffffffffff91821681529116602082015260400190565b90600182811c92168015611a4c575b6020831014611a3857565b634e487b7160e01b5f52602260045260245ffd5b91607f1691611a2d565b604051905f825f5160206123f95f395f51905f525491611a7583611a1e565b8083529260018116908115611af35750600114611a9b575b611a9992500383611949565b565b505f5160206123f95f395f51905f525f90815290915f5160206124995f395f51905f525b818310611ad7575050906020611a9992820101611a8d565b6020919350806001915483858901015201910190918492611abf565b60209250611a9994915060ff191682840152151560051b820101611a8d565b604051905f825f5160206124595f395f51905f525491611b3183611a1e565b8083529260018116908115611af35750600114611b5457611a9992500383611949565b505f5160206124595f395f51905f525f90815290915f5160206126395f395f51905f525b818310611b90575050906020611a9992820101611a8d565b6020919350806001915483858901015201910190918492611b78565b5f525f5160206125795f395f51905f52602052600160405f20015490565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b5f5160206125d95f395f51905f52548060d01c8015159081611c49575b5015611c335760a01c65ffffffffffff1690565b505f5160206124d95f395f51905f525460d01c90565b905042115f611c1f565b5f5160206124d95f395f51905f52546001600160a01b0381169160a09190911c65ffffffffffff1690565b335f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff1615611cb657565b63e2517d3f60e01b5f52336004525f60245260445ffd5b335f9081527f549fe2656c81d2947b3b913f0a53b9ea86c71e049f3a1b8aa23c09a8a05cb8d4602052604090205460ff1615611d0557565b63e2517d3f60e01b5f52336004525f5160206125395f395f51905f5260245260445ffd5b5f8181525f5160206125795f395f51905f526020908152604080832033845290915290205460ff1615611d595750565b63e2517d3f60e01b5f523360045260245260445ffd5b6001600160a01b031690811561067a576001600160a01b031691821561144a57815f525f5160206123d95f395f51905f5260205260405f2054818110611e0157815f5160206125595f395f51905f5292602092855f525f5160206123d95f395f51905f5284520360405f2055845f525f5160206123d95f395f51905f52825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b9065ffffffffffff8091169116019065ffffffffffff82116112c657565b5f5160206125d95f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b5f5160206125d95f395f51905f52546001600160a01b031661027d5780611e8f611e9592611e38565b5f6120b8565b90565b611e95905f5160206125395f395f51905f526120b8565b908115611ec0575b611e95916120b8565b5f5160206125d95f395f51905f52546001600160a01b031661027d57611e9591611ee982611e38565b9150611eb7565b916001600160a01b0383169182156116a5576001600160a01b0316928315611692577f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591611f3f602092611bca565b855f5282528060405f2055604051908152a3565b611f5b612253565b611f636122aa565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152611fb460c082611949565b51902090565b5f5160206125d95f395f51905f5254611e9591906001600160a01b03808316911614611fe7575b5f6122dc565b5f5160206125d95f395f51905f5280546001600160a01b0319169055611fe1565b90611e959180158061203b575b156122dc575f5160206125d95f395f51905f5280546001600160a01b03191690556122dc565b505f5160206125d95f395f51905f52546001600160a01b03838116911614612015565b65ffffffffffff81116120765765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b60ff5f5160206125f95f395f51905f525460401c16156120a957565b631afcd79f60e31b5f5260045ffd5b5f8181525f5160206125795f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16612156575f8181525f5160206125795f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384116121d4579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa156121c9575f516001600160a01b038116156121bf57905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f9160039190565b600481101561223f57806121f1575050565b600181036122085763f645eedf60e01b5f5260045ffd5b60028103612223575063fce698f760e01b5f5260045260245ffd5b60031461222d5750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b61225b611a56565b805190811561226b576020012090565b50505f5160206124795f395f51905f525480156122855790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6122b2611b12565b80519081156122c2576020012090565b50505f5160206126195f395f51905f525480156122855790565b5f8181525f5160206125795f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615612156575f8181525f5160206125795f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fe2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab09581ab046f3e11ba7ecf3d09df64f89f69c3a2a7a7e6f470b2498e2d7b1f570052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0452c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00eef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984002b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051099f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aaeef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a2646970667358221220389afd1ddb80ac3d5940c34a988583fa7ac9caf5961c30deed7098a652a1f16164736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// CrossMintableERC20V2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossMintableERC20V2MetaData.ABI instead.
var CrossMintableERC20V2ABI = CrossMintableERC20V2MetaData.ABI

// Deprecated: Use CrossMintableERC20V2MetaData.Sigs instead.
// CrossMintableERC20V2FuncSigs maps the 4-byte function signature to its string representation.
var CrossMintableERC20V2FuncSigs = CrossMintableERC20V2MetaData.Sigs

// CrossMintableERC20V2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossMintableERC20V2MetaData.Bin instead.
var CrossMintableERC20V2Bin = CrossMintableERC20V2MetaData.Bin

// DeployCrossMintableERC20V2 deploys a new Ethereum contract, binding an instance of CrossMintableERC20V2 to it.
func DeployCrossMintableERC20V2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossMintableERC20V2, error) {
	parsed, err := CrossMintableERC20V2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossMintableERC20V2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossMintableERC20V2{CrossMintableERC20V2Caller: CrossMintableERC20V2Caller{contract: contract}, CrossMintableERC20V2Transactor: CrossMintableERC20V2Transactor{contract: contract}, CrossMintableERC20V2Filterer: CrossMintableERC20V2Filterer{contract: contract}}, nil
}

// CrossMintableERC20V2 is an auto generated Go binding around an Ethereum contract.
type CrossMintableERC20V2 struct {
	CrossMintableERC20V2Caller     // Read-only binding to the contract
	CrossMintableERC20V2Transactor // Write-only binding to the contract
	CrossMintableERC20V2Filterer   // Log filterer for contract events
}

// CrossMintableERC20V2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossMintableERC20V2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20V2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossMintableERC20V2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20V2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossMintableERC20V2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20V2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossMintableERC20V2Session struct {
	Contract     *CrossMintableERC20V2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossMintableERC20V2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossMintableERC20V2CallerSession struct {
	Contract *CrossMintableERC20V2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CrossMintableERC20V2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossMintableERC20V2TransactorSession struct {
	Contract     *CrossMintableERC20V2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CrossMintableERC20V2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossMintableERC20V2Raw struct {
	Contract *CrossMintableERC20V2 // Generic contract binding to access the raw methods on
}

// CrossMintableERC20V2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossMintableERC20V2CallerRaw struct {
	Contract *CrossMintableERC20V2Caller // Generic read-only contract binding to access the raw methods on
}

// CrossMintableERC20V2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossMintableERC20V2TransactorRaw struct {
	Contract *CrossMintableERC20V2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossMintableERC20V2 creates a new instance of CrossMintableERC20V2, bound to a specific deployed contract.
func NewCrossMintableERC20V2(address common.Address, backend bind.ContractBackend) (*CrossMintableERC20V2, error) {
	contract, err := bindCrossMintableERC20V2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2{CrossMintableERC20V2Caller: CrossMintableERC20V2Caller{contract: contract}, CrossMintableERC20V2Transactor: CrossMintableERC20V2Transactor{contract: contract}, CrossMintableERC20V2Filterer: CrossMintableERC20V2Filterer{contract: contract}}, nil
}

// NewCrossMintableERC20V2Caller creates a new read-only instance of CrossMintableERC20V2, bound to a specific deployed contract.
func NewCrossMintableERC20V2Caller(address common.Address, caller bind.ContractCaller) (*CrossMintableERC20V2Caller, error) {
	contract, err := bindCrossMintableERC20V2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2Caller{contract: contract}, nil
}

// NewCrossMintableERC20V2Transactor creates a new write-only instance of CrossMintableERC20V2, bound to a specific deployed contract.
func NewCrossMintableERC20V2Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossMintableERC20V2Transactor, error) {
	contract, err := bindCrossMintableERC20V2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2Transactor{contract: contract}, nil
}

// NewCrossMintableERC20V2Filterer creates a new log filterer instance of CrossMintableERC20V2, bound to a specific deployed contract.
func NewCrossMintableERC20V2Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossMintableERC20V2Filterer, error) {
	contract, err := bindCrossMintableERC20V2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2Filterer{contract: contract}, nil
}

// bindCrossMintableERC20V2 binds a generic wrapper to an already deployed contract.
func bindCrossMintableERC20V2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossMintableERC20V2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20V2 *CrossMintableERC20V2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20V2.Contract.CrossMintableERC20V2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20V2 *CrossMintableERC20V2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.CrossMintableERC20V2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20V2 *CrossMintableERC20V2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.CrossMintableERC20V2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20V2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.DEFAULTADMINROLE(&_CrossMintableERC20V2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.DEFAULTADMINROLE(&_CrossMintableERC20V2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.DOMAINSEPARATOR(&_CrossMintableERC20V2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.DOMAINSEPARATOR(&_CrossMintableERC20V2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.Allowance(&_CrossMintableERC20V2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.Allowance(&_CrossMintableERC20V2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.BalanceOf(&_CrossMintableERC20V2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.BalanceOf(&_CrossMintableERC20V2.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Decimals() (uint8, error) {
	return _CrossMintableERC20V2.Contract.Decimals(&_CrossMintableERC20V2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Decimals() (uint8, error) {
	return _CrossMintableERC20V2.Contract.Decimals(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) DefaultAdmin() (common.Address, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdmin(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) DefaultAdmin() (common.Address, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdmin(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) DefaultAdminDelay() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdminDelay(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdminDelay(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdminDelayIncreaseWait(&_CrossMintableERC20V2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.DefaultAdminDelayIncreaseWait(&_CrossMintableERC20V2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.Eip712Domain(&_CrossMintableERC20V2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.Eip712Domain(&_CrossMintableERC20V2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.GetRoleAdmin(&_CrossMintableERC20V2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossMintableERC20V2.Contract.GetRoleAdmin(&_CrossMintableERC20V2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossMintableERC20V2.Contract.HasRole(&_CrossMintableERC20V2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossMintableERC20V2.Contract.HasRole(&_CrossMintableERC20V2.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Name() (string, error) {
	return _CrossMintableERC20V2.Contract.Name(&_CrossMintableERC20V2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Name() (string, error) {
	return _CrossMintableERC20V2.Contract.Name(&_CrossMintableERC20V2.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Nonces(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "nonces", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Nonces(owner_ common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.Nonces(&_CrossMintableERC20V2.CallOpts, owner_)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Nonces(owner_ common.Address) (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.Nonces(&_CrossMintableERC20V2.CallOpts, owner_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Owner() (common.Address, error) {
	return _CrossMintableERC20V2.Contract.Owner(&_CrossMintableERC20V2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Owner() (common.Address, error) {
	return _CrossMintableERC20V2.Contract.Owner(&_CrossMintableERC20V2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.PendingDefaultAdmin(&_CrossMintableERC20V2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.PendingDefaultAdmin(&_CrossMintableERC20V2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.PendingDefaultAdminDelay(&_CrossMintableERC20V2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossMintableERC20V2.Contract.PendingDefaultAdminDelay(&_CrossMintableERC20V2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossMintableERC20V2.Contract.SupportsInterface(&_CrossMintableERC20V2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossMintableERC20V2.Contract.SupportsInterface(&_CrossMintableERC20V2.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Symbol() (string, error) {
	return _CrossMintableERC20V2.Contract.Symbol(&_CrossMintableERC20V2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) Symbol() (string, error) {
	return _CrossMintableERC20V2.Contract.Symbol(&_CrossMintableERC20V2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20V2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) TotalSupply() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.TotalSupply(&_CrossMintableERC20V2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20V2 *CrossMintableERC20V2CallerSession) TotalSupply() (*big.Int, error) {
	return _CrossMintableERC20V2.Contract.TotalSupply(&_CrossMintableERC20V2.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.AcceptDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.AcceptDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Approve(&_CrossMintableERC20V2.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Approve(&_CrossMintableERC20V2.TransactOpts, spender, value)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.BeginDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.BeginDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts, newAdmin)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Burn(&_CrossMintableERC20V2.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Burn(&_CrossMintableERC20V2.TransactOpts, _account, _amount)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.CancelDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.CancelDefaultAdminTransfer(&_CrossMintableERC20V2.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.ChangeDefaultAdminDelay(&_CrossMintableERC20V2.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.ChangeDefaultAdminDelay(&_CrossMintableERC20V2.TransactOpts, newDelay)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.GrantRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.GrantRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address initialOwner, address initialMinter, string name_, string symbol_, uint8 decimals_) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialMinter common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "initialize", initialOwner, initialMinter, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address initialOwner, address initialMinter, string name_, string symbol_, uint8 decimals_) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Initialize(initialOwner common.Address, initialMinter common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Initialize(&_CrossMintableERC20V2.TransactOpts, initialOwner, initialMinter, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8420ce99.
//
// Solidity: function initialize(address initialOwner, address initialMinter, string name_, string symbol_, uint8 decimals_) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Initialize(initialOwner common.Address, initialMinter common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Initialize(&_CrossMintableERC20V2.TransactOpts, initialOwner, initialMinter, name_, symbol_, decimals_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Mint(&_CrossMintableERC20V2.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Mint(&_CrossMintableERC20V2.TransactOpts, _account, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Permit(&_CrossMintableERC20V2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Permit(&_CrossMintableERC20V2.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RenounceRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RenounceRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RevokeRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RevokeRole(&_CrossMintableERC20V2.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RollbackDefaultAdminDelay(&_CrossMintableERC20V2.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.RollbackDefaultAdminDelay(&_CrossMintableERC20V2.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Transfer(&_CrossMintableERC20V2.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.Transfer(&_CrossMintableERC20V2.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.TransferFrom(&_CrossMintableERC20V2.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20V2 *CrossMintableERC20V2TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20V2.Contract.TransferFrom(&_CrossMintableERC20V2.TransactOpts, from, to, value)
}

// CrossMintableERC20V2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2ApprovalIterator struct {
	Event *CrossMintableERC20V2Approval // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2Approval)
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
		it.Event = new(CrossMintableERC20V2Approval)
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
func (it *CrossMintableERC20V2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2Approval represents a Approval event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossMintableERC20V2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2ApprovalIterator{contract: _CrossMintableERC20V2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2Approval)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseApproval(log types.Log) (*CrossMintableERC20V2Approval, error) {
	event := new(CrossMintableERC20V2Approval)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator struct {
	Event *CrossMintableERC20V2DefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2DefaultAdminDelayChangeCanceled)
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
		it.Event = new(CrossMintableERC20V2DefaultAdminDelayChangeCanceled)
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
func (it *CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2DefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2DefaultAdminDelayChangeCanceledIterator{contract: _CrossMintableERC20V2.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2DefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2DefaultAdminDelayChangeCanceled)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossMintableERC20V2DefaultAdminDelayChangeCanceled, error) {
	event := new(CrossMintableERC20V2DefaultAdminDelayChangeCanceled)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator struct {
	Event *CrossMintableERC20V2DefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2DefaultAdminDelayChangeScheduled)
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
		it.Event = new(CrossMintableERC20V2DefaultAdminDelayChangeScheduled)
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
func (it *CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2DefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2DefaultAdminDelayChangeScheduledIterator{contract: _CrossMintableERC20V2.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2DefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2DefaultAdminDelayChangeScheduled)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossMintableERC20V2DefaultAdminDelayChangeScheduled, error) {
	event := new(CrossMintableERC20V2DefaultAdminDelayChangeScheduled)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2DefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminTransferCanceledIterator struct {
	Event *CrossMintableERC20V2DefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2DefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2DefaultAdminTransferCanceled)
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
		it.Event = new(CrossMintableERC20V2DefaultAdminTransferCanceled)
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
func (it *CrossMintableERC20V2DefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2DefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2DefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossMintableERC20V2DefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2DefaultAdminTransferCanceledIterator{contract: _CrossMintableERC20V2.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2DefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2DefaultAdminTransferCanceled)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseDefaultAdminTransferCanceled(log types.Log) (*CrossMintableERC20V2DefaultAdminTransferCanceled, error) {
	event := new(CrossMintableERC20V2DefaultAdminTransferCanceled)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2DefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminTransferScheduledIterator struct {
	Event *CrossMintableERC20V2DefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2DefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2DefaultAdminTransferScheduled)
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
		it.Event = new(CrossMintableERC20V2DefaultAdminTransferScheduled)
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
func (it *CrossMintableERC20V2DefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2DefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2DefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2DefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossMintableERC20V2DefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2DefaultAdminTransferScheduledIterator{contract: _CrossMintableERC20V2.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2DefaultAdminTransferScheduled)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseDefaultAdminTransferScheduled(log types.Log) (*CrossMintableERC20V2DefaultAdminTransferScheduled, error) {
	event := new(CrossMintableERC20V2DefaultAdminTransferScheduled)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2EIP712DomainChangedIterator struct {
	Event *CrossMintableERC20V2EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2EIP712DomainChanged)
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
		it.Event = new(CrossMintableERC20V2EIP712DomainChanged)
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
func (it *CrossMintableERC20V2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2EIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossMintableERC20V2EIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2EIP712DomainChangedIterator{contract: _CrossMintableERC20V2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2EIP712DomainChanged)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseEIP712DomainChanged(log types.Log) (*CrossMintableERC20V2EIP712DomainChanged, error) {
	event := new(CrossMintableERC20V2EIP712DomainChanged)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2InitializedIterator struct {
	Event *CrossMintableERC20V2Initialized // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2Initialized)
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
		it.Event = new(CrossMintableERC20V2Initialized)
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
func (it *CrossMintableERC20V2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2Initialized represents a Initialized event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CrossMintableERC20V2InitializedIterator, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2InitializedIterator{contract: _CrossMintableERC20V2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2Initialized) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2Initialized)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseInitialized(log types.Log) (*CrossMintableERC20V2Initialized, error) {
	event := new(CrossMintableERC20V2Initialized)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleAdminChangedIterator struct {
	Event *CrossMintableERC20V2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2RoleAdminChanged)
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
		it.Event = new(CrossMintableERC20V2RoleAdminChanged)
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
func (it *CrossMintableERC20V2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2RoleAdminChanged represents a RoleAdminChanged event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossMintableERC20V2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2RoleAdminChangedIterator{contract: _CrossMintableERC20V2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2RoleAdminChanged)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseRoleAdminChanged(log types.Log) (*CrossMintableERC20V2RoleAdminChanged, error) {
	event := new(CrossMintableERC20V2RoleAdminChanged)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleGrantedIterator struct {
	Event *CrossMintableERC20V2RoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2RoleGranted)
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
		it.Event = new(CrossMintableERC20V2RoleGranted)
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
func (it *CrossMintableERC20V2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2RoleGranted represents a RoleGranted event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossMintableERC20V2RoleGrantedIterator, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2RoleGrantedIterator{contract: _CrossMintableERC20V2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2RoleGranted)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseRoleGranted(log types.Log) (*CrossMintableERC20V2RoleGranted, error) {
	event := new(CrossMintableERC20V2RoleGranted)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleRevokedIterator struct {
	Event *CrossMintableERC20V2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2RoleRevoked)
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
		it.Event = new(CrossMintableERC20V2RoleRevoked)
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
func (it *CrossMintableERC20V2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2RoleRevoked represents a RoleRevoked event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossMintableERC20V2RoleRevokedIterator, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2RoleRevokedIterator{contract: _CrossMintableERC20V2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2RoleRevoked)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseRoleRevoked(log types.Log) (*CrossMintableERC20V2RoleRevoked, error) {
	event := new(CrossMintableERC20V2RoleRevoked)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20V2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2TransferIterator struct {
	Event *CrossMintableERC20V2Transfer // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20V2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20V2Transfer)
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
		it.Event = new(CrossMintableERC20V2Transfer)
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
func (it *CrossMintableERC20V2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20V2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20V2Transfer represents a Transfer event raised by the CrossMintableERC20V2 contract.
type CrossMintableERC20V2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossMintableERC20V2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20V2TransferIterator{contract: _CrossMintableERC20V2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20V2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossMintableERC20V2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20V2Transfer)
				if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20V2 *CrossMintableERC20V2Filterer) ParseTransfer(log types.Log) (*CrossMintableERC20V2Transfer, error) {
	event := new(CrossMintableERC20V2Transfer)
	if err := _CrossMintableERC20V2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
