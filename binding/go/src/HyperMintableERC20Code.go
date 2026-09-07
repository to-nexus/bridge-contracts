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

// HyperMintableERC20CodeMetaData contains all meta data concerning the HyperMintableERC20Code contract.
var HyperMintableERC20CodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"computeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"computeTokenAddressWithName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createCrossMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"createHyperMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTokenAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beacon_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isHyperMintableERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"setCoreTokenIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"}],\"name\":\"setHyperCoreDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"HyperMintableERC20Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"HyperMintableERC20CodeUnknownToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HyperMintableERC20CodeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"59659e90": "beacon()",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"e59fc21a": "computeTokenAddress(uint256,address,string,uint8,address)",
		"08882984": "computeTokenAddressWithName(uint256,address,string,string,uint8,address)",
		"f88d3d42": "createCrossMintableERC20(uint256,address,string,uint8)",
		"af83959c": "createHyperMintableERC20(uint256,address,string,string,uint8,address)",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"f8c8765e": "initialize(address,address,address,address)",
		"d2b0c67a": "isHyperMintableERC20(address)",
		"8da5cb5b": "owner()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"52d1902d": "proxiableUUID()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"3c77c1ac": "setCoreTokenIndex(address,uint64)",
		"f5998587": "setHyperCoreDeployer(address,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"b7e1917c": "tokenAdmin()",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
	Bin: "0x60a080604052346100c257306080525f5160206123245f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161225d90816100c78239608051818181610cd60152610d7a0152f35b6001600160401b0319166001600160401b039081175f5160206123245f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a71461125157508063022d63fb1461123457806308882984146111f75780630aa6220b14611154578063248a9ca31461112e5780632f2ff15d146110e857806336568abe14610fed5780633c77c1ac14610f4b5780634f1ef28614610d2a57806352d1902d14610cc357806359659e9014610c8e578063634e93da14610bbc578063649a5ec714610a3957806384ef8ffc14610a345780638da5cb5b14610a3457806391d14854146109de578063a1eda53c1461097c578063a217fddf14610960578063ad3cb1cc14610913578063af83959c146108eb578063b7e1917c146108b6578063cc8463c81461088b578063cefc1429146107c7578063cf6eefb71461078c578063d2b0c67a1461075a578063d547741f14610708578063d602b9fd146106a3578063e59fc21a146105bd578063f599858714610512578063f88d3d42146103d35763f8c8765e14610173575f80fd5b346103d05760803660031901126103d05761018c6112d4565b6101946112be565b90604435906001600160a01b038216908183036103ce576064356001600160a01b038116908190036103ca575f5160206122085f395f51905f5254604081901c60ff161595906001600160401b038116801590816103c2575b60011490816103b8575b1590816103af575b506103a0576001600160401b031981166001175f5160206122085f395f51905f525586610378575b506001600160a01b0316801561036957811561036957610245611a64565b61024d611a64565b610255611a64565b6001600160a01b03831615610355575f5160206121085f395f51905f5280546001600160d01b031690556102e492919061028e836118f3565b5060018060a01b03195f5160206120a85f395f51905f525416175f5160206120a85f395f51905f525560018060a01b03195f5160206120c85f395f51905f525416175f5160206120c85f395f51905f5255611922565b50610345575b506102f25780f35b60ff60401b195f5160206122085f395f51905f5254165f5160206122085f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a180f35b61034e90611939565b505f6102ea565b636116401160e11b87526004879052602487fd5b6363a2dfe160e01b8752600487fd5b6001600160481b0319166001600160401b01175f5160206122085f395f51905f52555f610227565b63f92ee8a960e01b8852600488fd5b9050155f6101ff565b303b1591506101f7565b8891506101ed565b8580fd5b845b80fd5b50346103d05760803660031901126103d0576103ed6112be565b6044356001600160401b03811161050e5761040c903690600401611372565b610414611390565b905f5160206121885f395f51905f5284525f5160206121685f395f51905f526020526040842060018060a01b0333165f5260205260ff60405f205416156104ea57906020936104d893926104ce600187604051936c021b937b9b990213934b233b29609d1b828601526104a6602d8683519885850199808b8585015e820190838201520301601f1981018752866112ea565b6040519586915180918484015e8101600f60fb1b838201520301601e198101855201836112ea565b339360043561183f565b6040516001600160a01b039091168152f35b63e2517d3f60e01b8452336004525f5160206121885f395f51905f52602452604484fd5b8280fd5b50346103d05760403660031901126103d0578061052d6112d4565b6105356112be565b9061053e61176b565b6105548160ff61054d826114d4565b541661150c565b6001600160a01b031690813b156105b9576040516389e9829160e01b81526001600160a01b0390911660048201529082908290602490829084905af180156105ae5761059d5750f35b816105a7916112ea565b6103d05780f35b6040513d84823e3d90fd5b5050fd5b50346103d05760a03660031901126103d0576105d76112be565b6044356001600160401b03811161050e576105f6903690600401611372565b906105ff611390565b608435916001600160a01b03831683036103ce5791610696916020956104d89594604051610645816106378b82019460043586611498565b03601f1981018352826112ea565b51902094610691600189604051936c021b937b9b990213934b233b29609d1b828601526104a6602d8683519885850199808b8585015e820190838201520301601f1981018752866112ea565b6115de565b83815191012030916116f3565b50346103d057806003193601126103d0576106bc61171c565b65ffffffffffff6106cb611585565b5f5160206121085f395f51905f5280546001600160d01b03191690559190911690506106f45780f35b5f5160206121485f395f51905f528180a180f35b50346103d05760403660031901126103d0576004356107256112be565b90801561074b57908161074261073d610747946114b6565b6117c7565b6119df565b5080f35b631fe1e13d60e11b8352600483fd5b50346103d05760203660031901126103d057602060ff61078061077b6112d4565b6114d4565b54166040519015158152f35b50346103d057806003193601126103d057604065ffffffffffff6107ae611585565b83516001600160a01b0390921682529091166020820152f35b50346103d057806003193601126103d0576107e0611585565b506001600160a01b031633036108785765ffffffffffff6107ff611585565b9091168015801561086e575b61085c57505f5160206121e85f395f51905f525461083c9190610836906001600160a01b0316611991565b506118f3565b505f5160206121085f395f51905f5280546001600160d01b031916905580f35b6319ca5ebb60e01b8352600452602482fd5b504281101561080b565b636116401160e11b815233600452602490fd5b50346103d057806003193601126103d05760206108a6611534565b65ffffffffffff60405191168152f35b50346103d057806003193601126103d0575f5160206120a85f395f51905f52546040516001600160a01b039091168152602090f35b50346103d05760206104d86108ff366113a0565b9461090e94919493929361176b565b61183f565b50346103d057806003193601126103d0575061095c6040516109366040826112ea565b60058152640352e302e360dc1b6020820152604051918291602083526020830190611474565b0390f35b50346103d057806003193601126103d057602090604051908152f35b50346103d057806003193601126103d0575f5160206121e85f395f51905f52548060d01c91821515806109d4575b156109cb575060a01c65ffffffffffff165b61095c6040519283928361145b565b915050806109bc565b50428310156109aa565b50346103d05760403660031901126103d05760406109fa6112be565b9160043581525f5160206121685f395f51905f52602052209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b611427565b50346103d05760203660031901126103d05760043565ffffffffffff811680820361050e57610a6661171c565b610a6f42611a35565b9065ffffffffffff610a7f611534565b1680821115610b8057505f5160206121a85f395f51905f52929165ffffffffffff826206978080610aba95109118026206978018169061180d565b905f5160206121e85f395f51905f52548060d01c80610b2b575b50505f5160206121e85f395f51905f5280546001600160d01b031960d085901b166001600160a01b0390911665ffffffffffff60a01b60a085901b1617179055604051918291610b2591908361145b565b0390a180f35b421115610b69575f5160206121085f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b5f80610ad4565b505f5160206121285f395f51905f528480a1610b62565b0365ffffffffffff8111610ba8575f5160206121a85f395f51905f529291610aba919061180d565b634e487b7160e01b84526011600452602484fd5b50346103d05760203660031901126103d057610bd66112d4565b610bde61171c565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020610c1b610c0d42611a35565b610c15611534565b9061180d565b65ffffffffffff610c2a611585565b5f5160206121085f395f51905f5280546001600160d01b0319166001600160a01b0390981697881785851660a01b179055919091169050610c78575b65ffffffffffff60405191168152a280f35b5f5160206121485f395f51905f528580a1610c66565b50346103d057806003193601126103d0575f5160206120c85f395f51905f52546040516001600160a01b039091168152602090f35b50346103d057806003193601126103d0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610d1b5760206040515f5160206120e85f395f51905f528152f35b63703e46dd60e11b8152600490fd5b5060403660031901126103d057610d3f6112d4565b906024356001600160401b038111610f475736602382011215610f4757610d7090369060248160040135910161133c565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016308114908115610f25575b50610f1657610db261176b565b6040516352d1902d60e01b8152926001600160a01b0381169190602085600481865afa80958596610ee2575b50610df757634c9c8ce360e01b84526004839052602484fd5b9091845f5160206120e85f395f51905f528103610ed05750813b15610ebe575f5160206120e85f395f51905f5280546001600160a01b031916821790557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8480a28151839015610ea4578083602061074795519101845af43d15610e9c573d91610e8083611321565b92610e8e60405194856112ea565b83523d85602085013e611bcf565b606091611bcf565b50505034610eaf5780f35b63b398979f60e01b8152600490fd5b634c9c8ce360e01b8452600452602483fd5b632a87526960e21b8552600452602484fd5b9095506020813d602011610f0e575b81610efe602093836112ea565b810103126103ce5751945f610dde565b3d9150610ef1565b63703e46dd60e11b8252600482fd5b5f5160206120e85f395f51905f52546001600160a01b0316141590505f610da5565b5080fd5b5034610fe9576040366003190112610fe957610f656112d4565b6024356001600160401b0381169190829003610fe957610f8361176b565b610f928160ff61054d826114d4565b6001600160a01b031690813b15610fe9575f916024839260405194859384926360d4eab960e01b845260048401525af18015610fde57610fd0575080f35b610fdc91505f906112ea565b005b6040513d5f823e3d90fd5b5f80fd5b34610fe9576040366003190112610fe9576004356110096112be565b8115806110c5575b61103e575b336001600160a01b0382160361102f57610fdc916119df565b63334bd91960e11b5f5260045ffd5b611046611585565b906001600160a01b0316158015906110b5575b80156110a3575b61108857505f5160206121085f395f51905f52805465ffffffffffff60a01b19169055611016565b65ffffffffffff906319ca5ebb60e01b5f521660045260245ffd5b504265ffffffffffff82161015611060565b5065ffffffffffff811615611059565b505f5160206121e85f395f51905f52546001600160a01b03828116911614611011565b34610fe9576040366003190112610fe9576004356111046112be565b811561111f578161111a61073d610fdc946114b6565b611950565b631fe1e13d60e11b5f5260045ffd5b34610fe9576020366003190112610fe957602061114c6004356114b6565b604051908152f35b34610fe9575f366003190112610fe95761116c61171c565b5f5160206121e85f395f51905f52548060d01c806111a2575b5f5160206121e85f395f51905f5280546001600160a01b03169055005b4211156111e0575f5160206121085f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b8080611185565b505f5160206121285f395f51905f525f80a16111d9565b34610fe95760206104d861122b610696610637611213366113a0565b95919360409891959398519283918c83019586611498565b519020946115de565b34610fe9575f366003190112610fe9576020604051620697808152f35b34610fe9576020366003190112610fe9576004359063ffffffff60e01b8216809203610fe9576020916318a4c3c360e11b8114908115611293575b5015158152f35b637965db0b60e01b8114915081156112ad575b508361128c565b6301ffc9a760e01b149050836112a6565b602435906001600160a01b0382168203610fe957565b600435906001600160a01b0382168203610fe957565b601f909101601f19168101906001600160401b0382119082101761130d57604052565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161130d57601f01601f191660200190565b92919261134882611321565b9161135660405193846112ea565b829481845281830111610fe9578281602093845f960137010152565b9080601f83011215610fe95781602061138d9335910161133c565b90565b6064359060ff82168203610fe957565b9060c0600319830112610fe957600435916024356001600160a01b0381168103610fe957916044356001600160401b038111610fe957826113e391600401611372565b91606435906001600160401b038211610fe95761140291600401611372565b9060843560ff81168103610fe9579060a4356001600160a01b0381168103610fe95790565b34610fe9575f366003190112610fe9575f5160206121e85f395f51905f52546040516001600160a01b039091168152602090f35b65ffffffffffff91821681529116602082015260400190565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b90815260609190911b6001600160601b031916602082015260340190565b5f525f5160206121685f395f51905f52602052600160405f20015490565b6001600160a01b03165f9081527fd55591ae453b55973909532d0252f41e2a7c543bd5a578ad3871068e0c28b4026020526040902090565b156115145750565b63bbfb7dd160e01b5f9081526001600160a01b0391909116600452602490fd5b5f5160206121e85f395f51905f52548060d01c801515908161157b575b50156115655760a01c65ffffffffffff1690565b505f5160206121085f395f51905f525460d01c90565b905042115f611551565b5f5160206121085f395f51905f52546001600160a01b0381169160a09190911c65ffffffffffff1690565b5f5160206121e85f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b5f5160206120a85f395f51905f525460405163238b4bc560e01b60208201526001600160a01b0391821660248201529416604485015230606485015260c06084850152611665928492909160ff916116509161163e9060e4870190611474565b8581036023190160a487015290611474565b911660c483015203601f1981018352826112ea565b61138d61047a9160206040519161167e828601846112ea565b84835281830194611c2e86396106376116c460018060a01b035f5160206120c85f395f51905f525416926040519283918683019586526040808401526060830190611474565b6040519586945180918587015e840190838201905f8252519283915e01015f815203601f1981018352826112ea565b91600b92604051926040840152602083015281520160ff8153605590206001600160a01b031690565b335f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff161561175457565b63e2517d3f60e01b5f52336004525f60245260445ffd5b335f9081527fb16e88c42fd4e48df2dd6a2eabd6bc9aec654ec170056b470819f8892cc6431c602052604090205460ff16156117a357565b63e2517d3f60e01b5f52336004525f5160206121c85f395f51905f5260245260445ffd5b5f8181525f5160206121685f395f51905f526020908152604080832033845290915290205460ff16156117f75750565b63e2517d3f60e01b5f523360045260245260445ffd5b9065ffffffffffff8091169116019065ffffffffffff821161182b57565b634e487b7160e01b5f52601160045260245ffd5b9594909361185f9293604051602081019061122b816106378a8d86611498565b8051156118e4576020815191015ff5923d1519841516610fde576001600160a01b0384169182156118d55760207f58e6d5c50dd4cd668a4caf443654b5156401f414eba85b21cb4a4b284604f4b4916118b7876114d4565b805460ff191660011790556040519485526001600160a01b031693a3565b63b06ebf3d60e01b5f5260045ffd5b631328927760e21b5f5260045ffd5b5f5160206121e85f395f51905f52546001600160a01b031661111f578061191c61138d926115b0565b5f611a8f565b61138d905f5160206121c85f395f51905f52611a8f565b61138d905f5160206121885f395f51905f52611a8f565b908115611961575b61138d91611a8f565b5f5160206121e85f395f51905f52546001600160a01b031661111f5761138d9161198a826115b0565b9150611958565b5f5160206121e85f395f51905f525461138d91906001600160a01b038083169116146119be575b5f611b33565b5f5160206121e85f395f51905f5280546001600160a01b03191690556119b8565b9061138d91801580611a12575b15611b33575f5160206121e85f395f51905f5280546001600160a01b0319169055611b33565b505f5160206121e85f395f51905f52546001600160a01b038381169116146119ec565b65ffffffffffff8111611a4d5765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b60ff5f5160206122085f395f51905f525460401c1615611a8057565b631afcd79f60e31b5f5260045ffd5b5f8181525f5160206121685f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16611b2d575f8181525f5160206121685f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f8181525f5160206121685f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff1615611b2d575f8181525f5160206121685f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b90611bf35750805115611be457805190602001fd5b63d6bda27560e01b5f5260045ffd5b81511580611c24575b611c04575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b15611bfc56fe60a08060405261047a80380380916100178285610292565b833981016040828203126101eb5761002e826102c9565b602083015190926001600160401b0382116101eb57019080601f830112156101eb57815161005b816102dd565b926100696040519485610292565b8184526020840192602083830101116101eb57815f926020809301855e84010152823b15610274577fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b038516908117909155604051635c60da1b60e01b8152909190602081600481865afa9081156101f7575f9161023a575b50803b1561021a5750817f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e5f80a282511561020257602060049260405193848092635c60da1b60e01b82525afa9182156101f7575f926101ae575b505f809161018a945190845af43d156101a6573d9161016e836102dd565b9261017c6040519485610292565b83523d5f602085013e6102f8565b505b608052604051610123908161035782396080518160180152f35b6060916102f8565b9291506020833d6020116101ef575b816101ca60209383610292565b810103126101eb575f80916101e161018a956102c9565b9394509150610150565b5f80fd5b3d91506101bd565b6040513d5f823e3d90fd5b505050341561018c5763b398979f60e01b5f5260045ffd5b634c9c8ce360e01b5f9081526001600160a01b0391909116600452602490fd5b90506020813d60201161026c575b8161025560209383610292565b810103126101eb57610266906102c9565b5f6100f5565b3d9150610248565b631933b43b60e21b5f9081526001600160a01b038416600452602490fd5b601f909101601f19168101906001600160401b038211908210176102b557604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101eb57565b6001600160401b0381116102b557601f01601f191660200190565b9061031c575080511561030d57805190602001fd5b63d6bda27560e01b5f5260045ffd5b8151158061034d575b61032d575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b1561032556fe60806040819052635c60da1b60e01b81526020906004817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801560a2575f901560d1575060203d602011609c575b6080601f8201601f1916810191906001600160401b0383119083101760885760849160405260800160ad565b60d1565b634e487b7160e01b5f52604160045260245ffd5b503d6058565b6040513d5f823e3d90fd5b602090607f19011260cd576080516001600160a01b038116810360cd5790565b5f80fd5b5f8091368280378136915af43d5f803e1560e9573d5ff35b3d5ffdfea26469706673582212205f4b66a0285020beaa693f75b815b8398189fb1defde86e535fbf956d2e8b8df64736f6c634300081c0033d55591ae453b55973909532d0252f41e2a7c543bd5a578ad3871068e0c28b400d55591ae453b55973909532d0252f41e2a7c543bd5a578ad3871068e0c28b401360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbceef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984002b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680052ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9ba49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775eef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122088978564be2b30ce4cbc0061e27261d3944572b692653a691355cab929ece11f64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// HyperMintableERC20CodeABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperMintableERC20CodeMetaData.ABI instead.
var HyperMintableERC20CodeABI = HyperMintableERC20CodeMetaData.ABI

// Deprecated: Use HyperMintableERC20CodeMetaData.Sigs instead.
// HyperMintableERC20CodeFuncSigs maps the 4-byte function signature to its string representation.
var HyperMintableERC20CodeFuncSigs = HyperMintableERC20CodeMetaData.Sigs

// HyperMintableERC20CodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperMintableERC20CodeMetaData.Bin instead.
var HyperMintableERC20CodeBin = HyperMintableERC20CodeMetaData.Bin

// DeployHyperMintableERC20Code deploys a new Ethereum contract, binding an instance of HyperMintableERC20Code to it.
func DeployHyperMintableERC20Code(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HyperMintableERC20Code, error) {
	parsed, err := HyperMintableERC20CodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperMintableERC20CodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperMintableERC20Code{HyperMintableERC20CodeCaller: HyperMintableERC20CodeCaller{contract: contract}, HyperMintableERC20CodeTransactor: HyperMintableERC20CodeTransactor{contract: contract}, HyperMintableERC20CodeFilterer: HyperMintableERC20CodeFilterer{contract: contract}}, nil
}

// HyperMintableERC20Code is an auto generated Go binding around an Ethereum contract.
type HyperMintableERC20Code struct {
	HyperMintableERC20CodeCaller     // Read-only binding to the contract
	HyperMintableERC20CodeTransactor // Write-only binding to the contract
	HyperMintableERC20CodeFilterer   // Log filterer for contract events
}

// HyperMintableERC20CodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperMintableERC20CodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20CodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperMintableERC20CodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20CodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperMintableERC20CodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20CodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperMintableERC20CodeSession struct {
	Contract     *HyperMintableERC20Code // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HyperMintableERC20CodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperMintableERC20CodeCallerSession struct {
	Contract *HyperMintableERC20CodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// HyperMintableERC20CodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperMintableERC20CodeTransactorSession struct {
	Contract     *HyperMintableERC20CodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// HyperMintableERC20CodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperMintableERC20CodeRaw struct {
	Contract *HyperMintableERC20Code // Generic contract binding to access the raw methods on
}

// HyperMintableERC20CodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperMintableERC20CodeCallerRaw struct {
	Contract *HyperMintableERC20CodeCaller // Generic read-only contract binding to access the raw methods on
}

// HyperMintableERC20CodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperMintableERC20CodeTransactorRaw struct {
	Contract *HyperMintableERC20CodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperMintableERC20Code creates a new instance of HyperMintableERC20Code, bound to a specific deployed contract.
func NewHyperMintableERC20Code(address common.Address, backend bind.ContractBackend) (*HyperMintableERC20Code, error) {
	contract, err := bindHyperMintableERC20Code(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20Code{HyperMintableERC20CodeCaller: HyperMintableERC20CodeCaller{contract: contract}, HyperMintableERC20CodeTransactor: HyperMintableERC20CodeTransactor{contract: contract}, HyperMintableERC20CodeFilterer: HyperMintableERC20CodeFilterer{contract: contract}}, nil
}

// NewHyperMintableERC20CodeCaller creates a new read-only instance of HyperMintableERC20Code, bound to a specific deployed contract.
func NewHyperMintableERC20CodeCaller(address common.Address, caller bind.ContractCaller) (*HyperMintableERC20CodeCaller, error) {
	contract, err := bindHyperMintableERC20Code(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeCaller{contract: contract}, nil
}

// NewHyperMintableERC20CodeTransactor creates a new write-only instance of HyperMintableERC20Code, bound to a specific deployed contract.
func NewHyperMintableERC20CodeTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperMintableERC20CodeTransactor, error) {
	contract, err := bindHyperMintableERC20Code(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeTransactor{contract: contract}, nil
}

// NewHyperMintableERC20CodeFilterer creates a new log filterer instance of HyperMintableERC20Code, bound to a specific deployed contract.
func NewHyperMintableERC20CodeFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperMintableERC20CodeFilterer, error) {
	contract, err := bindHyperMintableERC20Code(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeFilterer{contract: contract}, nil
}

// bindHyperMintableERC20Code binds a generic wrapper to an already deployed contract.
func bindHyperMintableERC20Code(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperMintableERC20CodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperMintableERC20Code *HyperMintableERC20CodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperMintableERC20Code.Contract.HyperMintableERC20CodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperMintableERC20Code *HyperMintableERC20CodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.HyperMintableERC20CodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperMintableERC20Code *HyperMintableERC20CodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.HyperMintableERC20CodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperMintableERC20Code.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.DEFAULTADMINROLE(&_HyperMintableERC20Code.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.DEFAULTADMINROLE(&_HyperMintableERC20Code.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _HyperMintableERC20Code.Contract.UPGRADEINTERFACEVERSION(&_HyperMintableERC20Code.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _HyperMintableERC20Code.Contract.UPGRADEINTERFACEVERSION(&_HyperMintableERC20Code.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) Beacon() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.Beacon(&_HyperMintableERC20Code.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) Beacon() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.Beacon(&_HyperMintableERC20Code.CallOpts)
}

// ComputeTokenAddress is a free data retrieval call binding the contract method 0xe59fc21a.
//
// Solidity: function computeTokenAddress(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) ComputeTokenAddress(opts *bind.CallOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8, minter common.Address) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "computeTokenAddress", remoteChainID, remoteToken, symbol, decimals, minter)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenAddress is a free data retrieval call binding the contract method 0xe59fc21a.
//
// Solidity: function computeTokenAddress(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) ComputeTokenAddress(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8, minter common.Address) (common.Address, error) {
	return _HyperMintableERC20Code.Contract.ComputeTokenAddress(&_HyperMintableERC20Code.CallOpts, remoteChainID, remoteToken, symbol, decimals, minter)
}

// ComputeTokenAddress is a free data retrieval call binding the contract method 0xe59fc21a.
//
// Solidity: function computeTokenAddress(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) ComputeTokenAddress(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8, minter common.Address) (common.Address, error) {
	return _HyperMintableERC20Code.Contract.ComputeTokenAddress(&_HyperMintableERC20Code.CallOpts, remoteChainID, remoteToken, symbol, decimals, minter)
}

// ComputeTokenAddressWithName is a free data retrieval call binding the contract method 0x08882984.
//
// Solidity: function computeTokenAddressWithName(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) ComputeTokenAddressWithName(opts *bind.CallOpts, remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "computeTokenAddressWithName", remoteChainID, remoteToken, name_, symbol_, decimals, minter)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenAddressWithName is a free data retrieval call binding the contract method 0x08882984.
//
// Solidity: function computeTokenAddressWithName(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) ComputeTokenAddressWithName(remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (common.Address, error) {
	return _HyperMintableERC20Code.Contract.ComputeTokenAddressWithName(&_HyperMintableERC20Code.CallOpts, remoteChainID, remoteToken, name_, symbol_, decimals, minter)
}

// ComputeTokenAddressWithName is a free data retrieval call binding the contract method 0x08882984.
//
// Solidity: function computeTokenAddressWithName(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) ComputeTokenAddressWithName(remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (common.Address, error) {
	return _HyperMintableERC20Code.Contract.ComputeTokenAddressWithName(&_HyperMintableERC20Code.CallOpts, remoteChainID, remoteToken, name_, symbol_, decimals, minter)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) DefaultAdmin() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdmin(&_HyperMintableERC20Code.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) DefaultAdmin() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdmin(&_HyperMintableERC20Code.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) DefaultAdminDelay() (*big.Int, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdminDelay(&_HyperMintableERC20Code.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdminDelay(&_HyperMintableERC20Code.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdminDelayIncreaseWait(&_HyperMintableERC20Code.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _HyperMintableERC20Code.Contract.DefaultAdminDelayIncreaseWait(&_HyperMintableERC20Code.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.GetRoleAdmin(&_HyperMintableERC20Code.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.GetRoleAdmin(&_HyperMintableERC20Code.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HyperMintableERC20Code.Contract.HasRole(&_HyperMintableERC20Code.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HyperMintableERC20Code.Contract.HasRole(&_HyperMintableERC20Code.CallOpts, role, account)
}

// IsHyperMintableERC20 is a free data retrieval call binding the contract method 0xd2b0c67a.
//
// Solidity: function isHyperMintableERC20(address token) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) IsHyperMintableERC20(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "isHyperMintableERC20", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHyperMintableERC20 is a free data retrieval call binding the contract method 0xd2b0c67a.
//
// Solidity: function isHyperMintableERC20(address token) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) IsHyperMintableERC20(token common.Address) (bool, error) {
	return _HyperMintableERC20Code.Contract.IsHyperMintableERC20(&_HyperMintableERC20Code.CallOpts, token)
}

// IsHyperMintableERC20 is a free data retrieval call binding the contract method 0xd2b0c67a.
//
// Solidity: function isHyperMintableERC20(address token) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) IsHyperMintableERC20(token common.Address) (bool, error) {
	return _HyperMintableERC20Code.Contract.IsHyperMintableERC20(&_HyperMintableERC20Code.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) Owner() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.Owner(&_HyperMintableERC20Code.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) Owner() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.Owner(&_HyperMintableERC20Code.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20Code.Contract.PendingDefaultAdmin(&_HyperMintableERC20Code.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20Code.Contract.PendingDefaultAdmin(&_HyperMintableERC20Code.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20Code.Contract.PendingDefaultAdminDelay(&_HyperMintableERC20Code.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20Code.Contract.PendingDefaultAdminDelay(&_HyperMintableERC20Code.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) ProxiableUUID() ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.ProxiableUUID(&_HyperMintableERC20Code.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _HyperMintableERC20Code.Contract.ProxiableUUID(&_HyperMintableERC20Code.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HyperMintableERC20Code.Contract.SupportsInterface(&_HyperMintableERC20Code.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HyperMintableERC20Code.Contract.SupportsInterface(&_HyperMintableERC20Code.CallOpts, interfaceId)
}

// TokenAdmin is a free data retrieval call binding the contract method 0xb7e1917c.
//
// Solidity: function tokenAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCaller) TokenAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20Code.contract.Call(opts, &out, "tokenAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAdmin is a free data retrieval call binding the contract method 0xb7e1917c.
//
// Solidity: function tokenAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) TokenAdmin() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.TokenAdmin(&_HyperMintableERC20Code.CallOpts)
}

// TokenAdmin is a free data retrieval call binding the contract method 0xb7e1917c.
//
// Solidity: function tokenAdmin() view returns(address)
func (_HyperMintableERC20Code *HyperMintableERC20CodeCallerSession) TokenAdmin() (common.Address, error) {
	return _HyperMintableERC20Code.Contract.TokenAdmin(&_HyperMintableERC20Code.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.AcceptDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.AcceptDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.BeginDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.BeginDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CancelDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CancelDefaultAdminTransfer(&_HyperMintableERC20Code.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.ChangeDefaultAdminDelay(&_HyperMintableERC20Code.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.ChangeDefaultAdminDelay(&_HyperMintableERC20Code.TransactOpts, newDelay)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) CreateCrossMintableERC20(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "createCrossMintableERC20", remoteChainID, remoteToken, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) CreateCrossMintableERC20(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CreateCrossMintableERC20(&_HyperMintableERC20Code.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) CreateCrossMintableERC20(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CreateCrossMintableERC20(&_HyperMintableERC20Code.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateHyperMintableERC20 is a paid mutator transaction binding the contract method 0xaf83959c.
//
// Solidity: function createHyperMintableERC20(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) CreateHyperMintableERC20(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "createHyperMintableERC20", remoteChainID, remoteToken, name_, symbol_, decimals, minter)
}

// CreateHyperMintableERC20 is a paid mutator transaction binding the contract method 0xaf83959c.
//
// Solidity: function createHyperMintableERC20(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) CreateHyperMintableERC20(remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CreateHyperMintableERC20(&_HyperMintableERC20Code.TransactOpts, remoteChainID, remoteToken, name_, symbol_, decimals, minter)
}

// CreateHyperMintableERC20 is a paid mutator transaction binding the contract method 0xaf83959c.
//
// Solidity: function createHyperMintableERC20(uint256 remoteChainID, address remoteToken, string name_, string symbol_, uint8 decimals, address minter) returns(address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) CreateHyperMintableERC20(remoteChainID *big.Int, remoteToken common.Address, name_ string, symbol_ string, decimals uint8, minter common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.CreateHyperMintableERC20(&_HyperMintableERC20Code.TransactOpts, remoteChainID, remoteToken, name_, symbol_, decimals, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.GrantRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.GrantRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address initialTokenAdmin, address initialBridge, address beacon_) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialTokenAdmin common.Address, initialBridge common.Address, beacon_ common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "initialize", initialOwner, initialTokenAdmin, initialBridge, beacon_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address initialTokenAdmin, address initialBridge, address beacon_) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) Initialize(initialOwner common.Address, initialTokenAdmin common.Address, initialBridge common.Address, beacon_ common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.Initialize(&_HyperMintableERC20Code.TransactOpts, initialOwner, initialTokenAdmin, initialBridge, beacon_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address initialOwner, address initialTokenAdmin, address initialBridge, address beacon_) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) Initialize(initialOwner common.Address, initialTokenAdmin common.Address, initialBridge common.Address, beacon_ common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.Initialize(&_HyperMintableERC20Code.TransactOpts, initialOwner, initialTokenAdmin, initialBridge, beacon_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RenounceRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RenounceRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RevokeRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RevokeRole(&_HyperMintableERC20Code.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RollbackDefaultAdminDelay(&_HyperMintableERC20Code.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.RollbackDefaultAdminDelay(&_HyperMintableERC20Code.TransactOpts)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x3c77c1ac.
//
// Solidity: function setCoreTokenIndex(address token, uint64 index) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) SetCoreTokenIndex(opts *bind.TransactOpts, token common.Address, index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "setCoreTokenIndex", token, index)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x3c77c1ac.
//
// Solidity: function setCoreTokenIndex(address token, uint64 index) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) SetCoreTokenIndex(token common.Address, index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.SetCoreTokenIndex(&_HyperMintableERC20Code.TransactOpts, token, index)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x3c77c1ac.
//
// Solidity: function setCoreTokenIndex(address token, uint64 index) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) SetCoreTokenIndex(token common.Address, index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.SetCoreTokenIndex(&_HyperMintableERC20Code.TransactOpts, token, index)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0xf5998587.
//
// Solidity: function setHyperCoreDeployer(address token, address finalizer) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) SetHyperCoreDeployer(opts *bind.TransactOpts, token common.Address, finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "setHyperCoreDeployer", token, finalizer)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0xf5998587.
//
// Solidity: function setHyperCoreDeployer(address token, address finalizer) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) SetHyperCoreDeployer(token common.Address, finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.SetHyperCoreDeployer(&_HyperMintableERC20Code.TransactOpts, token, finalizer)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0xf5998587.
//
// Solidity: function setHyperCoreDeployer(address token, address finalizer) returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) SetHyperCoreDeployer(token common.Address, finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.SetHyperCoreDeployer(&_HyperMintableERC20Code.TransactOpts, token, finalizer)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HyperMintableERC20Code.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.UpgradeToAndCall(&_HyperMintableERC20Code.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_HyperMintableERC20Code *HyperMintableERC20CodeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _HyperMintableERC20Code.Contract.UpgradeToAndCall(&_HyperMintableERC20Code.TransactOpts, newImplementation, data)
}

// HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator struct {
	Event *HyperMintableERC20CodeDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeDefaultAdminDelayChangeCanceled)
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
		it.Event = new(HyperMintableERC20CodeDefaultAdminDelayChangeCanceled)
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
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeDefaultAdminDelayChangeCanceledIterator{contract: _HyperMintableERC20Code.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeDefaultAdminDelayChangeCanceled)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*HyperMintableERC20CodeDefaultAdminDelayChangeCanceled, error) {
	event := new(HyperMintableERC20CodeDefaultAdminDelayChangeCanceled)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator struct {
	Event *HyperMintableERC20CodeDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeDefaultAdminDelayChangeScheduled)
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
		it.Event = new(HyperMintableERC20CodeDefaultAdminDelayChangeScheduled)
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
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeDefaultAdminDelayChangeScheduledIterator{contract: _HyperMintableERC20Code.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeDefaultAdminDelayChangeScheduled)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*HyperMintableERC20CodeDefaultAdminDelayChangeScheduled, error) {
	event := new(HyperMintableERC20CodeDefaultAdminDelayChangeScheduled)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminTransferCanceledIterator struct {
	Event *HyperMintableERC20CodeDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeDefaultAdminTransferCanceled)
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
		it.Event = new(HyperMintableERC20CodeDefaultAdminTransferCanceled)
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
func (it *HyperMintableERC20CodeDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*HyperMintableERC20CodeDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeDefaultAdminTransferCanceledIterator{contract: _HyperMintableERC20Code.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeDefaultAdminTransferCanceled)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*HyperMintableERC20CodeDefaultAdminTransferCanceled, error) {
	event := new(HyperMintableERC20CodeDefaultAdminTransferCanceled)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminTransferScheduledIterator struct {
	Event *HyperMintableERC20CodeDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeDefaultAdminTransferScheduled)
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
		it.Event = new(HyperMintableERC20CodeDefaultAdminTransferScheduled)
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
func (it *HyperMintableERC20CodeDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*HyperMintableERC20CodeDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeDefaultAdminTransferScheduledIterator{contract: _HyperMintableERC20Code.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeDefaultAdminTransferScheduled)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*HyperMintableERC20CodeDefaultAdminTransferScheduled, error) {
	event := new(HyperMintableERC20CodeDefaultAdminTransferScheduled)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeHyperMintableERC20CreatedIterator is returned from FilterHyperMintableERC20Created and is used to iterate over the raw logs and unpacked data for HyperMintableERC20Created events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeHyperMintableERC20CreatedIterator struct {
	Event *HyperMintableERC20CodeHyperMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeHyperMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeHyperMintableERC20Created)
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
		it.Event = new(HyperMintableERC20CodeHyperMintableERC20Created)
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
func (it *HyperMintableERC20CodeHyperMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeHyperMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeHyperMintableERC20Created represents a HyperMintableERC20Created event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeHyperMintableERC20Created struct {
	RemoteChainID *big.Int
	RemoteToken   common.Address
	TokenAddress  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHyperMintableERC20Created is a free log retrieval operation binding the contract event 0x58e6d5c50dd4cd668a4caf443654b5156401f414eba85b21cb4a4b284604f4b4.
//
// Solidity: event HyperMintableERC20Created(uint256 indexed remoteChainID, address indexed remoteToken, address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterHyperMintableERC20Created(opts *bind.FilterOpts, remoteChainID []*big.Int, remoteToken []common.Address) (*HyperMintableERC20CodeHyperMintableERC20CreatedIterator, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "HyperMintableERC20Created", remoteChainIDRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeHyperMintableERC20CreatedIterator{contract: _HyperMintableERC20Code.contract, event: "HyperMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchHyperMintableERC20Created is a free log subscription operation binding the contract event 0x58e6d5c50dd4cd668a4caf443654b5156401f414eba85b21cb4a4b284604f4b4.
//
// Solidity: event HyperMintableERC20Created(uint256 indexed remoteChainID, address indexed remoteToken, address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchHyperMintableERC20Created(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeHyperMintableERC20Created, remoteChainID []*big.Int, remoteToken []common.Address) (event.Subscription, error) {

	var remoteChainIDRule []interface{}
	for _, remoteChainIDItem := range remoteChainID {
		remoteChainIDRule = append(remoteChainIDRule, remoteChainIDItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "HyperMintableERC20Created", remoteChainIDRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeHyperMintableERC20Created)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "HyperMintableERC20Created", log); err != nil {
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

// ParseHyperMintableERC20Created is a log parse operation binding the contract event 0x58e6d5c50dd4cd668a4caf443654b5156401f414eba85b21cb4a4b284604f4b4.
//
// Solidity: event HyperMintableERC20Created(uint256 indexed remoteChainID, address indexed remoteToken, address tokenAddress)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseHyperMintableERC20Created(log types.Log) (*HyperMintableERC20CodeHyperMintableERC20Created, error) {
	event := new(HyperMintableERC20CodeHyperMintableERC20Created)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "HyperMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeInitializedIterator struct {
	Event *HyperMintableERC20CodeInitialized // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeInitialized)
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
		it.Event = new(HyperMintableERC20CodeInitialized)
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
func (it *HyperMintableERC20CodeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeInitialized represents a Initialized event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterInitialized(opts *bind.FilterOpts) (*HyperMintableERC20CodeInitializedIterator, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeInitializedIterator{contract: _HyperMintableERC20Code.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeInitialized) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeInitialized)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseInitialized(log types.Log) (*HyperMintableERC20CodeInitialized, error) {
	event := new(HyperMintableERC20CodeInitialized)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleAdminChangedIterator struct {
	Event *HyperMintableERC20CodeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeRoleAdminChanged)
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
		it.Event = new(HyperMintableERC20CodeRoleAdminChanged)
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
func (it *HyperMintableERC20CodeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeRoleAdminChanged represents a RoleAdminChanged event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HyperMintableERC20CodeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeRoleAdminChangedIterator{contract: _HyperMintableERC20Code.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeRoleAdminChanged)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseRoleAdminChanged(log types.Log) (*HyperMintableERC20CodeRoleAdminChanged, error) {
	event := new(HyperMintableERC20CodeRoleAdminChanged)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleGrantedIterator struct {
	Event *HyperMintableERC20CodeRoleGranted // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeRoleGranted)
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
		it.Event = new(HyperMintableERC20CodeRoleGranted)
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
func (it *HyperMintableERC20CodeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeRoleGranted represents a RoleGranted event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HyperMintableERC20CodeRoleGrantedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeRoleGrantedIterator{contract: _HyperMintableERC20Code.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeRoleGranted)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseRoleGranted(log types.Log) (*HyperMintableERC20CodeRoleGranted, error) {
	event := new(HyperMintableERC20CodeRoleGranted)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleRevokedIterator struct {
	Event *HyperMintableERC20CodeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeRoleRevoked)
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
		it.Event = new(HyperMintableERC20CodeRoleRevoked)
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
func (it *HyperMintableERC20CodeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeRoleRevoked represents a RoleRevoked event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HyperMintableERC20CodeRoleRevokedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeRoleRevokedIterator{contract: _HyperMintableERC20Code.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeRoleRevoked)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseRoleRevoked(log types.Log) (*HyperMintableERC20CodeRoleRevoked, error) {
	event := new(HyperMintableERC20CodeRoleRevoked)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CodeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeUpgradedIterator struct {
	Event *HyperMintableERC20CodeUpgraded // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CodeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CodeUpgraded)
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
		it.Event = new(HyperMintableERC20CodeUpgraded)
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
func (it *HyperMintableERC20CodeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CodeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CodeUpgraded represents a Upgraded event raised by the HyperMintableERC20Code contract.
type HyperMintableERC20CodeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*HyperMintableERC20CodeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CodeUpgradedIterator{contract: _HyperMintableERC20Code.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CodeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _HyperMintableERC20Code.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CodeUpgraded)
				if err := _HyperMintableERC20Code.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_HyperMintableERC20Code *HyperMintableERC20CodeFilterer) ParseUpgraded(log types.Log) (*HyperMintableERC20CodeUpgraded, error) {
	event := new(HyperMintableERC20CodeUpgraded)
	if err := _HyperMintableERC20Code.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
