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

// HyperMintableERC20MetaData contains all meta data concerning the HyperMintableERC20 contract.
var HyperMintableERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CORE_SYSTEM_ADDRESS_PREFIX\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HYPERCORE_DEPLOYER_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreExtraWeiDecimals\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreSystemAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreTokenIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"coreTransferableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coreUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryLinker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hyperCoreDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialMinter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialLinker\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCoreTokenIndexSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLinkAuthority\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"setCoreTokenIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"}],\"name\":\"setHyperCoreDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToCore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coreRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToCoreFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"evmExtraWeiDecimals\",\"type\":\"int8\"}],\"name\":\"CoreTokenIndexSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousDeployer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDeployer\",\"type\":\"address\"}],\"name\":\"HyperCoreDeployerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CoreAmountBelowOneCoreWei\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"weiDecimals\",\"type\":\"uint8\"},{\"internalType\":\"int8\",\"name\":\"evmExtraWeiDecimals\",\"type\":\"int8\"}],\"name\":\"CoreDecimalsMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"evmContract\",\"type\":\"address\"}],\"name\":\"CoreLinkNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CoreRecipientIsSystemAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CoreRecipientZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CoreTokenIndexAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"CoreTokenIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"CoreTokenInfoUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HyperCoreLinkAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HyperMintableERC20CoreTokenIndexNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"d750809a": "CORE_SYSTEM_ADDRESS_PREFIX()",
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"3644e515": "DOMAIN_SEPARATOR()",
		"b83e14dc": "HYPERCORE_DEPLOYER_SLOT()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"9dc29fac": "burn(address,uint256)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"d83f0470": "coreExtraWeiDecimals()",
		"13f206f3": "coreSystemAddress()",
		"df31c6f1": "coreTokenIndex()",
		"868f9b34": "coreTransferableAmount(uint256)",
		"2eaf423b": "coreUnit()",
		"313ce567": "decimals()",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"84b0196e": "eip712Domain()",
		"28668468": "factoryLinker()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"b768259d": "hyperCoreDeployer()",
		"238b4bc5": "initialize(address,address,address,string,string,uint8)",
		"d54c5841": "isCoreTokenIndexSet()",
		"55c9cb07": "isLinkAuthority(address)",
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
		"60d4eab9": "setCoreTokenIndex(uint64)",
		"89e98291": "setHyperCoreDeployer(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"45f72615": "transferToCore(uint256)",
		"9438568d": "transferToCoreFor(address,uint256)",
	},
	Bin: "0x6080806040523460aa575f5160206131025f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b60405161305390816100af8239f35b6001600160401b0319166001600160401b039081175f5160206131025f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816301ffc9a7146120a057508063022d63fb1461208357806306fdde0314611fd9578063095ea7b314611fb35780630aa6220b14611f1057806313f206f314611ee457806318160ddd14611ebb578063238b4bc51461168b57806323b872dd146115b3578063248a9ca314611595578063286684681461155d5780632eaf423b146115435780632f2ff15d1461150c578063313ce567146114dd5780633644e515146114c357806336568abe146113c857806340c10f191461131557806345f72615146112df57806355c9cb07146112b257806360d4eab914610f1b578063634e93da14610e4b578063649a5ec714610ccb57806370a0823114610c875780637ecebe0014610c4357806384b0196e14610b1f57806384ef8ffc14610a2f578063868f9b3414610afb57806389e9829114610a345780638da5cb5b14610a2f57806391d14854146109da5780639438568d1461092757806395d89b41146108445780639dc29fac14610765578063a1eda53c14610700578063a217fddf146106e6578063a9059cbb146106b5578063b768259d14610681578063b83e14dc1461065a578063cc8463c814610630578063cefc14291461056c578063cf6eefb714610532578063d505accf146103e3578063d547741f14610396578063d54c584114610365578063d602b9fd14610303578063d750809a146102e4578063d83f0470146102b6578063dd62ed3e1461026f5763df31c6f114610237575f80fd5b3461026b575f36600319011261026b575f516020612d9e5f395f51905f52546040516001600160401b039091168152602090f35b5f80fd5b3461026b57604036600319011261026b57610288612131565b610299610293612147565b9161250d565b9060018060a01b03165f52602052602060405f2054604051908152f35b3461026b575f36600319011261026b5760205f516020612d9e5f395f51905f525460481c5f0b604051908152f35b3461026b575f36600319011261026b576040516001609d1b8152602090f35b3461026b575f36600319011261026b5761031b61261a565b65ffffffffffff61032a6125ef565b5f516020612e7e5f395f51905f5280546001600160d01b031916905591909116905061035257005b5f516020612ebe5f395f51905f525f80a1005b3461026b575f36600319011261026b57602060ff5f516020612d9e5f395f51905f525460401c166040519015158152f35b3461026b57604036600319011261026b576004356103b2612147565b81156103d457816103cd6103c86103d2946123f8565b6126c5565b612a5c565b005b631fe1e13d60e11b5f5260045ffd5b3461026b5760e036600319011261026b576103fc612131565b610404612147565b604435906064359260843560ff8116810361026b5784421161051f576104e46104ed9160018060a01b03841696875f525f516020612e5e5f395f51905f5260205260405f20908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c081526104b260e08261215d565b5190206104bd6129a7565b906040519161190160f01b83526002830152602282015260c43591604260a4359220612ae1565b90929192612b64565b6001600160a01b031684810361050857506103d29350612875565b84906325c0072360e11b5f5260045260245260445ffd5b8463313c898160e11b5f5260045260245ffd5b3461026b575f36600319011261026b57604065ffffffffffff6105536125ef565b83516001600160a01b0390921682529091166020820152f35b3461026b575f36600319011261026b576105846125ef565b506001600160a01b0316330361061d5765ffffffffffff6105a36125ef565b919091169081158015610613575b610600575f516020612f7e5f395f51905f52546105e191906105db906001600160a01b0316612a0e565b5061270b565b505f516020612e7e5f395f51905f5280546001600160d01b0319169055005b506319ca5ebb60e01b5f5260045260245ffd5b50428210156105b1565b636116401160e11b5f523360045260245ffd5b3461026b575f36600319011261026b57602061064a61259e565b65ffffffffffff60405191168152f35b3461026b575f36600319011261026b5760206040515f516020612dbe5f395f51905f528152f35b3461026b575f36600319011261026b575f516020612dbe5f395f51905f52546040516001600160a01b039091168152602090f35b3461026b57604036600319011261026b576106db6106d1612131565b60243590336127ac565b602060405160018152f35b3461026b575f36600319011261026b5760206040515f8152f35b3461026b575f36600319011261026b575f516020612f7e5f395f51905f52548060d01c908115158061075b575b156107515760a01c65ffffffffffff165b61074d60405192839283612215565b0390f35b50505f5f9061073e565b504282101561072d565b3461026b57604036600319011261026b5761077e612131565b6024359061078a612669565b6001600160a01b0316801561083157805f525f516020612d3e5f395f51905f5260205260405f2054828110610818576020835f945f516020612efe5f395f51905f52938587525f516020612d3e5f395f51905f528452036040862055805f516020612dde5f395f51905f5254035f516020612dde5f395f51905f5255604051908152a3602060405160018152f35b9063391434e360e21b5f5260045260245260445260645ffd5b634b637e8f60e11b5f525f60045260245ffd5b3461026b575f36600319011261026b576040515f5f516020612d7e5f395f51905f52546108708161222e565b808452906001811690811561090357506001146108ac575b61074d836108988185038261215d565b60405191829160208352602083019061210d565b5f516020612d7e5f395f51905f525f9081525f516020612f5e5f395f51905f52939250905b8082106108e957509091508101602001610898610888565b9192600181602092548385880101520191019092916108d1565b60ff191660208086019190915291151560051b840190910191506108989050610888565b3461026b57604036600319011261026b57610940612131565b6001600160a01b03811660243581156109cb5761095b6123bc565b916001600160a01b038316146109bc576109869061098061097a612416565b82612545565b90612563565b9081156109ad57816109a591846109a083602097336127ac565b6127ac565b604051908152f35b63e886ac9d60e01b5f5260045ffd5b634742251f60e11b5f5260045ffd5b63730d016b60e11b5f5260045ffd5b3461026b57604036600319011261026b576109f3612147565b6004355f525f516020612f1e5f395f51905f5260205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b6121e1565b3461026b57602036600319011261026b57610a4d612131565b610a5633612449565b15610ad75760ff5f516020612d9e5f395f51905f525460401c16610ac8575f516020612dbe5f395f51905f5280546001600160a01b039283166001600160a01b0319821681179092559091167fcae9891438942589688090dd5afab743aa8c874826bb838ba0c629ba75366e3e5f80a3005b63886a16ed60e01b5f5260045ffd5b63e2517d3f60e01b5f52336004525f516020612f9e5f395f51905f5260245260445ffd5b3461026b57602036600319011261026b5760206109a560043561098061097a612416565b3461026b575f36600319011261026b575f516020612e1e5f395f51905f52541580610c2d575b15610bf057610b94610b55612266565b610b5d612322565b6020610ba260405192610b70838561215d565b5f84525f368137604051958695600f60f81b875260e08588015260e087019061210d565b90858203604087015261210d565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b828110610bd957505050500390f35b835185528695509381019392810192600101610bca565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b505f516020612fde5f395f51905f525415610b45565b3461026b57602036600319011261026b576001600160a01b03610c64612131565b165f525f516020612e5e5f395f51905f52602052602060405f2054604051908152f35b3461026b57602036600319011261026b576001600160a01b03610ca8612131565b165f525f516020612d3e5f395f51905f52602052602060405f2054604051908152f35b3461026b57602036600319011261026b5760043565ffffffffffff81169081810361026b57610cf861261a565b610d0142612ab2565b9165ffffffffffff610d1161259e565b1680821115610e1057505f516020612f3e5f395f51905f529265ffffffffffff826206978080610d4b951091180262069780181690612857565b905f516020612f7e5f395f51905f52548060d01c80610dbb575b50505f516020612f7e5f395f51905f5280546001600160d01b031960d085901b166001600160a01b0390911665ffffffffffff60a01b60a085901b1617179055604051918291610db6919083612215565b0390a1005b421115610df9575f516020612e7e5f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b8380610d65565b505f516020612e9e5f395f51905f525f80a1610df2565b0365ffffffffffff8111610e37575f516020612f3e5f395f51905f5292610d4b9190612857565b634e487b7160e01b5f52601160045260245ffd5b3461026b57602036600319011261026b57610e64612131565b610e6c61261a565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020610ea9610e9b42612ab2565b610ea361259e565b90612857565b65ffffffffffff610eb86125ef565b5f516020612e7e5f395f51905f5280546001600160d01b0319166001600160a01b0390981697881785851660a01b179055919091169050610f05575b65ffffffffffff60405191168152a2005b5f516020612ebe5f395f51905f525f80a1610ef4565b3461026b57602036600319011261026b576004356001600160401b0381169081810361026b57610f4a33612449565b15610ad7575f516020612d9e5f395f51905f52549060ff8260401c166112a45763ffffffff8311611291575f809160405163ffffffff602082019216825260208152610f9760408261215d565b519061080c5afa3d15611289573d90610faf82612180565b91610fbd604051938461215d565b82523d5f602084013e5b8061127f575b1561126c5780518101602081019160208183031261026b576020810151906001600160401b03821161026b570190610100908290031261026b576040519161010083016001600160401b038111848210176112585760405260208201516001600160401b03811161026b5760209083010181601f8201121561026b57805161105481612180565b91611062604051938461215d565b818352836020838301011161026b57815f9260208093018386015e83010152835260408201516001600160401b03811161026b576020908301019080601f8301121561026b578151916001600160401b038311611258578260051b9060208201936110d0604051958661215d565b845260208085019282010192831161026b57602001905b828210611240575050506020830152611102606082016124d7565b6040830152611113608082016124eb565b606083015261112460a082016124eb565b80608084015261113660c083016124ff565b60a084015261010061114a60e084016124ff565b9260c08501938452015192835f0b840361026b5760e0019283526001600160a01b031630810361122a575060ff8360501c1660ff8251169183515f0b925f8482019485129112908015821691151617610e375760ff9051169083515f0b92810361121257835160ff60481b604882901b166001600160501b03198716881717600160401b175f516020612d9e5f395f51905f52556040515f9190910b815286907f2bd02f43024c0d42c7d60ce7b4c93d53b2d420af5fa93bee4f57d5de645de7c090602090a2005b633882970560e01b5f5260045260245260445260645ffd5b846359dbc94760e11b5f5260045260245260445ffd5b6020809161124d846124d7565b8152019101906110e7565b634e487b7160e01b5f52604160045260245ffd5b826372e4909160e01b5f5260045260245ffd5b5080511515610fcd565b606090610fc7565b82637e22f0cb60e11b5f5260045260245ffd5b62d94d4f60e11b5f5260045ffd5b3461026b57602036600319011261026b5760206112d56112d0612131565b612449565b6040519015158152f35b3461026b57602036600319011261026b57602061130360043561098061097a612416565b6109a58161130f6123bc565b336127ac565b3461026b57604036600319011261026b5761132e612131565b6024359061133a612669565b6001600160a01b03169081156113b5575f516020612dde5f395f51905f525490808201809211610e375760205f516020612efe5f395f51905f52915f935f516020612dde5f395f51905f52558484525f516020612d3e5f395f51905f52825260408420818154019055604051908152a3602060405160018152f35b63ec442f0560e01b5f525f60045260245ffd5b3461026b57604036600319011261026b576004356113e4612147565b8115806114a0575b611419575b336001600160a01b0382160361140a576103d291612a5c565b63334bd91960e11b5f5260045ffd5b6114216125ef565b906001600160a01b031615801590611490575b801561147e575b61146357505f516020612e7e5f395f51905f52805465ffffffffffff60a01b191690556113f1565b65ffffffffffff906319ca5ebb60e01b5f521660045260245ffd5b504265ffffffffffff8216101561143b565b5065ffffffffffff811615611434565b505f516020612f7e5f395f51905f52546001600160a01b038281169116146113ec565b3461026b575f36600319011261026b5760206109a56129a7565b3461026b575f36600319011261026b57602060ff5f516020612d9e5f395f51905f525460501c16604051908152f35b3461026b57604036600319011261026b57600435611528612147565b81156103d4578161153e6103c86103d2946123f8565b61276b565b3461026b575f36600319011261026b5760206109a5612416565b3461026b575f36600319011261026b575f516020612d9e5f395f51905f525460405160589190911c6001600160a01b03168152602090f35b3461026b57602036600319011261026b5760206109a56004356123f8565b3461026b57606036600319011261026b576115cc612131565b6115d4612147565b604435906115e18361250d565b335f9081526020919091526040902054925f198410611605575b6106db93506127ac565b828410611670576001600160a01b0381161561165d57331561164a576106db9361162e8261250d565b60018060a01b0333165f526020528360405f20910390556115fb565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8284637dc7a0d960e11b5f523360045260245260445260645ffd5b3461026b5760c036600319011261026b576116a4612131565b6116ac612147565b6044356001600160a01b03811680820361026b576064356001600160401b03811161026b576116df90369060040161219b565b926084356001600160401b03811161026b576116ff90369060040161219b565b9460a4359460ff8616860361026b575f516020612fbe5f395f51905f5254604081901c60ff161597906001600160401b03811680159081611eb3575b6001149081611ea9575b159081611ea0575b50611e91576001600160401b031981166001175f516020612fbe5f395f51905f525588611e69575b5061177e6128d8565b6117866128d8565b81516001600160401b038111611258576117ad5f516020612d1e5f395f51905f525461222e565b601f8111611e0d575b50806020601f8211600114611d91575f91611d86575b508160011b915f199060031b1c1916175f516020612d1e5f395f51905f52555b8051906001600160401b0382116112585781906118165f516020612d7e5f395f51905f525461222e565b601f8111611d1f575b50602090601f8311600114611ca1575f92611c96575b50508160011b915f199060031b1c1916175f516020612d7e5f395f51905f52555b61185e6128d8565b6040519061186d60408361215d565b60018252603160f81b60208301526118836128d8565b8051906001600160401b0382116112585781906118ad5f516020612d5e5f395f51905f525461222e565b601f8111611c2f575b50602090601f8311600114611bb1575f92611ba6575b50508160011b915f199060031b1c1916175f516020612d5e5f395f51905f52555b8051906001600160401b0382116112585781906119175f516020612dfe5f395f51905f525461222e565b601f8111611b3f575b50602090601f8311600114611ac1575f92611ab6575b50508160011b915f199060031b1c1916175f516020612dfe5f395f51905f52555b5f5f516020612e1e5f395f51905f52555f5f516020612fde5f395f51905f525561197f6128d8565b6119876128d8565b6001600160a01b03811615611aa3575f516020612e7e5f395f51905f5280546001600160d01b031690556119ba9061270b565b506001600160a01b038116611a93575b505f516020612d9e5f395f51905f528054600160581b600160f81b031916605884901b600160581b600160f81b0316179055611a83575b505f516020612d9e5f395f51905f52805460ff60501b191660509290921b60ff60501b16919091179055611a3157005b60ff60401b195f516020612fbe5f395f51905f5254165f516020612fbe5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b611a8c90612754565b5082611a01565b611a9c9061273d565b50846119ca565b636116401160e11b5f525f60045260245ffd5b015190508880611936565b5f516020612dfe5f395f51905f525f9081528281209350601f198516905b818110611b275750908460019594939210611b0f575b505050811b015f516020612dfe5f395f51905f5255611957565b01515f1960f88460031b161c19169055888080611af5565b92936020600181928786015181550195019301611adf565b5f516020612dfe5f395f51905f525f529091505f516020612ffe5f395f51905f52601f840160051c81019160208510611b9c575b90601f859493920160051c01905b818110611b8e5750611920565b5f8155849350600101611b81565b9091508190611b73565b0151905089806118cc565b5f516020612d5e5f395f51905f525f9081528281209350601f198516905b818110611c175750908460019594939210611bff575b505050811b015f516020612d5e5f395f51905f52556118ed565b01515f1960f88460031b161c19169055898080611be5565b92936020600181928786015181550195019301611bcf565b5f516020612d5e5f395f51905f525f529091505f516020612e3e5f395f51905f52601f840160051c81019160208510611c8c575b90601f859493920160051c01905b818110611c7e57506118b6565b5f8155849350600101611c71565b9091508190611c63565b015190508980611835565b5f516020612d7e5f395f51905f525f9081528281209350601f198516905b818110611d075750908460019594939210611cef575b505050811b015f516020612d7e5f395f51905f5255611856565b01515f1960f88460031b161c19169055898080611cd5565b92936020600181928786015181550195019301611cbf565b5f516020612d7e5f395f51905f525f529091505f516020612f5e5f395f51905f52601f840160051c81019160208510611d7c575b90601f859493920160051c01905b818110611d6e575061181f565b5f8155849350600101611d61565b9091508190611d53565b90508301518a6117cc565b5f516020612d1e5f395f51905f525f9081528181209250601f198416905b818110611df557509083600194939210611ddd575b5050811b015f516020612d1e5f395f51905f52556117ec565b8501515f1960f88460031b161c191690558a80611dc4565b9192602060018192868a015181550194019201611daf565b5f516020612d1e5f395f51905f525f525f516020612cfe5f395f51905f52601f830160051c81019160208410611e5f575b601f0160051c01905b818110611e5457506117b6565b5f8155600101611e47565b9091508190611e3e565b6001600160481b0319166001600160401b01175f516020612fbe5f395f51905f525588611775565b63f92ee8a960e01b5f5260045ffd5b9050158a61174d565b303b159150611745565b8a915061173b565b3461026b575f36600319011261026b5760205f516020612dde5f395f51905f5254604051908152f35b3461026b575f36600319011261026b576020611efe6123bc565b6040516001600160a01b039091168152f35b3461026b575f36600319011261026b57611f2861261a565b5f516020612f7e5f395f51905f52548060d01c80611f5e575b5f516020612f7e5f395f51905f5280546001600160a01b03169055005b421115611f9c575f516020612e7e5f395f51905f5280546001600160d01b031660309290921b6001600160d01b0319169190911790555b8080611f41565b505f516020612e9e5f395f51905f525f80a1611f95565b3461026b57604036600319011261026b576106db611fcf612131565b6024359033612875565b3461026b575f36600319011261026b576040515f5f516020612d1e5f395f51905f52546120058161222e565b8084529060018116908115610903575060011461202c5761074d836108988185038261215d565b5f516020612d1e5f395f51905f525f9081525f516020612cfe5f395f51905f52939250905b80821061206957509091508101602001610898610888565b919260018160209254838588010152019101909291612051565b3461026b575f36600319011261026b576020604051620697808152f35b3461026b57602036600319011261026b576004359063ffffffff60e01b821680920361026b576020916318a4c3c360e11b81149081156120e2575b5015158152f35b637965db0b60e01b8114915081156120fc575b50836120db565b6301ffc9a760e01b149050836120f5565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361026b57565b602435906001600160a01b038216820361026b57565b601f909101601f19168101906001600160401b0382119082101761125857604052565b6001600160401b03811161125857601f01601f191660200190565b81601f8201121561026b578035906121b282612180565b926121c0604051948561215d565b8284526020838301011161026b57815f926020809301838601378301015290565b3461026b575f36600319011261026b575f516020612f7e5f395f51905f52546040516001600160a01b039091168152602090f35b65ffffffffffff91821681529116602082015260400190565b90600182811c9216801561225c575b602083101461224857565b634e487b7160e01b5f52602260045260245ffd5b91607f169161223d565b604051905f825f516020612d5e5f395f51905f5254916122858361222e565b808352926001811690811561230357506001146122ab575b6122a99250038361215d565b565b505f516020612d5e5f395f51905f525f90815290915f516020612e3e5f395f51905f525b8183106122e75750509060206122a99282010161229d565b60209193508060019154838589010152019101909184926122cf565b602092506122a994915060ff191682840152151560051b82010161229d565b604051905f825f516020612dfe5f395f51905f5254916123418361222e565b80835292600181169081156123035750600114612364576122a99250038361215d565b505f516020612dfe5f395f51905f525f90815290915f516020612ffe5f395f51905f525b8183106123a05750509060206122a99282010161229d565b6020919350806001915483858901015201910190918492612388565b5f516020612d9e5f395f51905f525460ff8160401c16156123e9576001600160401b03166001609d1b1790565b6349aee5fd60e11b5f5260045ffd5b5f525f516020612f1e5f395f51905f52602052600160405f20015490565b5f516020612d9e5f395f51905f525460481c5f81810b13156124435760ff16604d8111610e3757600a0a90565b50600190565b6001600160a01b031680156124d2575f516020612f7e5f395f51905f52546001600160a01b03168114612443575f516020612d9e5f395f51905f525460581c6001600160a01b03168114908161249d575090565b5f9081527fd44a0d2099ad7d3a533e56e7ef83de2cd2546d692152b02f67f56eddc7787ef3602052604090205460ff16919050565b505f90565b51906001600160401b038216820361026b57565b51906001600160a01b038216820361026b57565b519060ff8216820361026b57565b6001600160a01b03165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b811561254f570690565b634e487b7160e01b5f52601260045260245ffd5b91908203918211610e3757565b5f516020612f7e5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b5f516020612f7e5f395f51905f52548060d01c80151590816125e5575b50156125cf5760a01c65ffffffffffff1690565b505f516020612e7e5f395f51905f525460d01c90565b905042115f6125bb565b5f516020612e7e5f395f51905f52546001600160a01b0381169160a09190911c65ffffffffffff1690565b335f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff161561265257565b63e2517d3f60e01b5f52336004525f60245260445ffd5b335f9081527f549fe2656c81d2947b3b913f0a53b9ea86c71e049f3a1b8aa23c09a8a05cb8d4602052604090205460ff16156126a157565b63e2517d3f60e01b5f52336004525f516020612ede5f395f51905f5260245260445ffd5b5f8181525f516020612f1e5f395f51905f526020908152604080832033845290915290205460ff16156126f55750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f516020612f7e5f395f51905f52546001600160a01b03166103d4578061273461273a92612570565b5f612903565b90565b61273a905f516020612ede5f395f51905f52612903565b61273a905f516020612f9e5f395f51905f52612903565b90811561277c575b61273a91612903565b5f516020612f7e5f395f51905f52546001600160a01b03166103d45761273a916127a582612570565b9150612773565b6001600160a01b0316908115610831576001600160a01b03169182156113b557815f525f516020612d3e5f395f51905f5260205260405f205481811061283e57815f516020612efe5f395f51905f5292602092855f525f516020612d3e5f395f51905f5284520360405f2055845f525f516020612d3e5f395f51905f52825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b9065ffffffffffff8091169116019065ffffffffffff8211610e3757565b916001600160a01b03831691821561165d576001600160a01b031692831561164a577f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925916128c460209261250d565b855f5282528060405f2055604051908152a3565b60ff5f516020612fbe5f395f51905f525460401c16156128f457565b631afcd79f60e31b5f5260045ffd5b5f8181525f516020612f1e5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff166129a1575f8181525f516020612f1e5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b6129af612bd8565b6129b7612c2f565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a08152612a0860c08261215d565b51902090565b5f516020612f7e5f395f51905f525461273a91906001600160a01b03808316911614612a3b575b5f612c61565b5f516020612f7e5f395f51905f5280546001600160a01b0319169055612a35565b9061273a91801580612a8f575b15612c61575f516020612f7e5f395f51905f5280546001600160a01b0319169055612c61565b505f516020612f7e5f395f51905f52546001600160a01b03838116911614612a69565b65ffffffffffff8111612aca5765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411612b59579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15612b4e575f516001600160a01b03811615612b4457905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f9160039190565b6004811015612bc45780612b76575050565b60018103612b8d5763f645eedf60e01b5f5260045ffd5b60028103612ba8575063fce698f760e01b5f5260045260245ffd5b600314612bb25750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b612be0612266565b8051908115612bf0576020012090565b50505f516020612e1e5f395f51905f52548015612c0a5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b612c37612322565b8051908115612c47576020012090565b50505f516020612fde5f395f51905f52548015612c0a5790565b5f8181525f516020612f1e5f395f51905f52602090815260408083206001600160a01b038616845290915290205460ff16156129a1575f8181525f516020612f1e5f395f51905f52602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fe2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace04c65afee183f44952959f664e5c2b8a5e4916480c324e2787b2215d744391d4008c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d103a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10042ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb00eef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984002b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051099f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800f1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b46a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aaeef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401733bac3dca102687aa08c854c5f9067fc424f98fd8e90e41ad6b73aecc59a4fdf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a264697066735822122009c2b8b725019e4c80098c4fc994c1b2750b23dc1008f9b0208b431af4ee19ef64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// HyperMintableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperMintableERC20MetaData.ABI instead.
var HyperMintableERC20ABI = HyperMintableERC20MetaData.ABI

// Deprecated: Use HyperMintableERC20MetaData.Sigs instead.
// HyperMintableERC20FuncSigs maps the 4-byte function signature to its string representation.
var HyperMintableERC20FuncSigs = HyperMintableERC20MetaData.Sigs

// HyperMintableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperMintableERC20MetaData.Bin instead.
var HyperMintableERC20Bin = HyperMintableERC20MetaData.Bin

// DeployHyperMintableERC20 deploys a new Ethereum contract, binding an instance of HyperMintableERC20 to it.
func DeployHyperMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HyperMintableERC20, error) {
	parsed, err := HyperMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperMintableERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperMintableERC20{HyperMintableERC20Caller: HyperMintableERC20Caller{contract: contract}, HyperMintableERC20Transactor: HyperMintableERC20Transactor{contract: contract}, HyperMintableERC20Filterer: HyperMintableERC20Filterer{contract: contract}}, nil
}

// HyperMintableERC20 is an auto generated Go binding around an Ethereum contract.
type HyperMintableERC20 struct {
	HyperMintableERC20Caller     // Read-only binding to the contract
	HyperMintableERC20Transactor // Write-only binding to the contract
	HyperMintableERC20Filterer   // Log filterer for contract events
}

// HyperMintableERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type HyperMintableERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperMintableERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperMintableERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperMintableERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperMintableERC20Session struct {
	Contract     *HyperMintableERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HyperMintableERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperMintableERC20CallerSession struct {
	Contract *HyperMintableERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HyperMintableERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperMintableERC20TransactorSession struct {
	Contract     *HyperMintableERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HyperMintableERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type HyperMintableERC20Raw struct {
	Contract *HyperMintableERC20 // Generic contract binding to access the raw methods on
}

// HyperMintableERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperMintableERC20CallerRaw struct {
	Contract *HyperMintableERC20Caller // Generic read-only contract binding to access the raw methods on
}

// HyperMintableERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperMintableERC20TransactorRaw struct {
	Contract *HyperMintableERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperMintableERC20 creates a new instance of HyperMintableERC20, bound to a specific deployed contract.
func NewHyperMintableERC20(address common.Address, backend bind.ContractBackend) (*HyperMintableERC20, error) {
	contract, err := bindHyperMintableERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20{HyperMintableERC20Caller: HyperMintableERC20Caller{contract: contract}, HyperMintableERC20Transactor: HyperMintableERC20Transactor{contract: contract}, HyperMintableERC20Filterer: HyperMintableERC20Filterer{contract: contract}}, nil
}

// NewHyperMintableERC20Caller creates a new read-only instance of HyperMintableERC20, bound to a specific deployed contract.
func NewHyperMintableERC20Caller(address common.Address, caller bind.ContractCaller) (*HyperMintableERC20Caller, error) {
	contract, err := bindHyperMintableERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20Caller{contract: contract}, nil
}

// NewHyperMintableERC20Transactor creates a new write-only instance of HyperMintableERC20, bound to a specific deployed contract.
func NewHyperMintableERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*HyperMintableERC20Transactor, error) {
	contract, err := bindHyperMintableERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20Transactor{contract: contract}, nil
}

// NewHyperMintableERC20Filterer creates a new log filterer instance of HyperMintableERC20, bound to a specific deployed contract.
func NewHyperMintableERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*HyperMintableERC20Filterer, error) {
	contract, err := bindHyperMintableERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20Filterer{contract: contract}, nil
}

// bindHyperMintableERC20 binds a generic wrapper to an already deployed contract.
func bindHyperMintableERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperMintableERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperMintableERC20 *HyperMintableERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperMintableERC20.Contract.HyperMintableERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperMintableERC20 *HyperMintableERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.HyperMintableERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperMintableERC20 *HyperMintableERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.HyperMintableERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperMintableERC20 *HyperMintableERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperMintableERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperMintableERC20 *HyperMintableERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperMintableERC20 *HyperMintableERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.contract.Transact(opts, method, params...)
}

// CORESYSTEMADDRESSPREFIX is a free data retrieval call binding the contract method 0xd750809a.
//
// Solidity: function CORE_SYSTEM_ADDRESS_PREFIX() view returns(uint160)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CORESYSTEMADDRESSPREFIX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "CORE_SYSTEM_ADDRESS_PREFIX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CORESYSTEMADDRESSPREFIX is a free data retrieval call binding the contract method 0xd750809a.
//
// Solidity: function CORE_SYSTEM_ADDRESS_PREFIX() view returns(uint160)
func (_HyperMintableERC20 *HyperMintableERC20Session) CORESYSTEMADDRESSPREFIX() (*big.Int, error) {
	return _HyperMintableERC20.Contract.CORESYSTEMADDRESSPREFIX(&_HyperMintableERC20.CallOpts)
}

// CORESYSTEMADDRESSPREFIX is a free data retrieval call binding the contract method 0xd750809a.
//
// Solidity: function CORE_SYSTEM_ADDRESS_PREFIX() view returns(uint160)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CORESYSTEMADDRESSPREFIX() (*big.Int, error) {
	return _HyperMintableERC20.Contract.CORESYSTEMADDRESSPREFIX(&_HyperMintableERC20.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _HyperMintableERC20.Contract.DEFAULTADMINROLE(&_HyperMintableERC20.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _HyperMintableERC20.Contract.DEFAULTADMINROLE(&_HyperMintableERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _HyperMintableERC20.Contract.DOMAINSEPARATOR(&_HyperMintableERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _HyperMintableERC20.Contract.DOMAINSEPARATOR(&_HyperMintableERC20.CallOpts)
}

// HYPERCOREDEPLOYERSLOT is a free data retrieval call binding the contract method 0xb83e14dc.
//
// Solidity: function HYPERCORE_DEPLOYER_SLOT() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Caller) HYPERCOREDEPLOYERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "HYPERCORE_DEPLOYER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HYPERCOREDEPLOYERSLOT is a free data retrieval call binding the contract method 0xb83e14dc.
//
// Solidity: function HYPERCORE_DEPLOYER_SLOT() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Session) HYPERCOREDEPLOYERSLOT() ([32]byte, error) {
	return _HyperMintableERC20.Contract.HYPERCOREDEPLOYERSLOT(&_HyperMintableERC20.CallOpts)
}

// HYPERCOREDEPLOYERSLOT is a free data retrieval call binding the contract method 0xb83e14dc.
//
// Solidity: function HYPERCORE_DEPLOYER_SLOT() view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) HYPERCOREDEPLOYERSLOT() ([32]byte, error) {
	return _HyperMintableERC20.Contract.HYPERCOREDEPLOYERSLOT(&_HyperMintableERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.Allowance(&_HyperMintableERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.Allowance(&_HyperMintableERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.BalanceOf(&_HyperMintableERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.BalanceOf(&_HyperMintableERC20.CallOpts, account)
}

// CoreExtraWeiDecimals is a free data retrieval call binding the contract method 0xd83f0470.
//
// Solidity: function coreExtraWeiDecimals() view returns(int8)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CoreExtraWeiDecimals(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "coreExtraWeiDecimals")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// CoreExtraWeiDecimals is a free data retrieval call binding the contract method 0xd83f0470.
//
// Solidity: function coreExtraWeiDecimals() view returns(int8)
func (_HyperMintableERC20 *HyperMintableERC20Session) CoreExtraWeiDecimals() (int8, error) {
	return _HyperMintableERC20.Contract.CoreExtraWeiDecimals(&_HyperMintableERC20.CallOpts)
}

// CoreExtraWeiDecimals is a free data retrieval call binding the contract method 0xd83f0470.
//
// Solidity: function coreExtraWeiDecimals() view returns(int8)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CoreExtraWeiDecimals() (int8, error) {
	return _HyperMintableERC20.Contract.CoreExtraWeiDecimals(&_HyperMintableERC20.CallOpts)
}

// CoreSystemAddress is a free data retrieval call binding the contract method 0x13f206f3.
//
// Solidity: function coreSystemAddress() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CoreSystemAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "coreSystemAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoreSystemAddress is a free data retrieval call binding the contract method 0x13f206f3.
//
// Solidity: function coreSystemAddress() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Session) CoreSystemAddress() (common.Address, error) {
	return _HyperMintableERC20.Contract.CoreSystemAddress(&_HyperMintableERC20.CallOpts)
}

// CoreSystemAddress is a free data retrieval call binding the contract method 0x13f206f3.
//
// Solidity: function coreSystemAddress() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CoreSystemAddress() (common.Address, error) {
	return _HyperMintableERC20.Contract.CoreSystemAddress(&_HyperMintableERC20.CallOpts)
}

// CoreTokenIndex is a free data retrieval call binding the contract method 0xdf31c6f1.
//
// Solidity: function coreTokenIndex() view returns(uint64)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CoreTokenIndex(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "coreTokenIndex")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CoreTokenIndex is a free data retrieval call binding the contract method 0xdf31c6f1.
//
// Solidity: function coreTokenIndex() view returns(uint64)
func (_HyperMintableERC20 *HyperMintableERC20Session) CoreTokenIndex() (uint64, error) {
	return _HyperMintableERC20.Contract.CoreTokenIndex(&_HyperMintableERC20.CallOpts)
}

// CoreTokenIndex is a free data retrieval call binding the contract method 0xdf31c6f1.
//
// Solidity: function coreTokenIndex() view returns(uint64)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CoreTokenIndex() (uint64, error) {
	return _HyperMintableERC20.Contract.CoreTokenIndex(&_HyperMintableERC20.CallOpts)
}

// CoreTransferableAmount is a free data retrieval call binding the contract method 0x868f9b34.
//
// Solidity: function coreTransferableAmount(uint256 amount) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CoreTransferableAmount(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "coreTransferableAmount", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoreTransferableAmount is a free data retrieval call binding the contract method 0x868f9b34.
//
// Solidity: function coreTransferableAmount(uint256 amount) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) CoreTransferableAmount(amount *big.Int) (*big.Int, error) {
	return _HyperMintableERC20.Contract.CoreTransferableAmount(&_HyperMintableERC20.CallOpts, amount)
}

// CoreTransferableAmount is a free data retrieval call binding the contract method 0x868f9b34.
//
// Solidity: function coreTransferableAmount(uint256 amount) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CoreTransferableAmount(amount *big.Int) (*big.Int, error) {
	return _HyperMintableERC20.Contract.CoreTransferableAmount(&_HyperMintableERC20.CallOpts, amount)
}

// CoreUnit is a free data retrieval call binding the contract method 0x2eaf423b.
//
// Solidity: function coreUnit() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) CoreUnit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "coreUnit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoreUnit is a free data retrieval call binding the contract method 0x2eaf423b.
//
// Solidity: function coreUnit() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) CoreUnit() (*big.Int, error) {
	return _HyperMintableERC20.Contract.CoreUnit(&_HyperMintableERC20.CallOpts)
}

// CoreUnit is a free data retrieval call binding the contract method 0x2eaf423b.
//
// Solidity: function coreUnit() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) CoreUnit() (*big.Int, error) {
	return _HyperMintableERC20.Contract.CoreUnit(&_HyperMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HyperMintableERC20 *HyperMintableERC20Session) Decimals() (uint8, error) {
	return _HyperMintableERC20.Contract.Decimals(&_HyperMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Decimals() (uint8, error) {
	return _HyperMintableERC20.Contract.Decimals(&_HyperMintableERC20.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Caller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Session) DefaultAdmin() (common.Address, error) {
	return _HyperMintableERC20.Contract.DefaultAdmin(&_HyperMintableERC20.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) DefaultAdmin() (common.Address, error) {
	return _HyperMintableERC20.Contract.DefaultAdmin(&_HyperMintableERC20.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20Caller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20Session) DefaultAdminDelay() (*big.Int, error) {
	return _HyperMintableERC20.Contract.DefaultAdminDelay(&_HyperMintableERC20.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _HyperMintableERC20.Contract.DefaultAdminDelay(&_HyperMintableERC20.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20Caller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20Session) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _HyperMintableERC20.Contract.DefaultAdminDelayIncreaseWait(&_HyperMintableERC20.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _HyperMintableERC20.Contract.DefaultAdminDelayIncreaseWait(&_HyperMintableERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "eip712Domain")

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
func (_HyperMintableERC20 *HyperMintableERC20Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _HyperMintableERC20.Contract.Eip712Domain(&_HyperMintableERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _HyperMintableERC20.Contract.Eip712Domain(&_HyperMintableERC20.CallOpts)
}

// FactoryLinker is a free data retrieval call binding the contract method 0x28668468.
//
// Solidity: function factoryLinker() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Caller) FactoryLinker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "factoryLinker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryLinker is a free data retrieval call binding the contract method 0x28668468.
//
// Solidity: function factoryLinker() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Session) FactoryLinker() (common.Address, error) {
	return _HyperMintableERC20.Contract.FactoryLinker(&_HyperMintableERC20.CallOpts)
}

// FactoryLinker is a free data retrieval call binding the contract method 0x28668468.
//
// Solidity: function factoryLinker() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) FactoryLinker() (common.Address, error) {
	return _HyperMintableERC20.Contract.FactoryLinker(&_HyperMintableERC20.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HyperMintableERC20.Contract.GetRoleAdmin(&_HyperMintableERC20.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _HyperMintableERC20.Contract.GetRoleAdmin(&_HyperMintableERC20.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HyperMintableERC20.Contract.HasRole(&_HyperMintableERC20.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _HyperMintableERC20.Contract.HasRole(&_HyperMintableERC20.CallOpts, role, account)
}

// HyperCoreDeployer is a free data retrieval call binding the contract method 0xb768259d.
//
// Solidity: function hyperCoreDeployer() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Caller) HyperCoreDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "hyperCoreDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HyperCoreDeployer is a free data retrieval call binding the contract method 0xb768259d.
//
// Solidity: function hyperCoreDeployer() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Session) HyperCoreDeployer() (common.Address, error) {
	return _HyperMintableERC20.Contract.HyperCoreDeployer(&_HyperMintableERC20.CallOpts)
}

// HyperCoreDeployer is a free data retrieval call binding the contract method 0xb768259d.
//
// Solidity: function hyperCoreDeployer() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) HyperCoreDeployer() (common.Address, error) {
	return _HyperMintableERC20.Contract.HyperCoreDeployer(&_HyperMintableERC20.CallOpts)
}

// IsCoreTokenIndexSet is a free data retrieval call binding the contract method 0xd54c5841.
//
// Solidity: function isCoreTokenIndexSet() view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Caller) IsCoreTokenIndexSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "isCoreTokenIndexSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCoreTokenIndexSet is a free data retrieval call binding the contract method 0xd54c5841.
//
// Solidity: function isCoreTokenIndexSet() view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) IsCoreTokenIndexSet() (bool, error) {
	return _HyperMintableERC20.Contract.IsCoreTokenIndexSet(&_HyperMintableERC20.CallOpts)
}

// IsCoreTokenIndexSet is a free data retrieval call binding the contract method 0xd54c5841.
//
// Solidity: function isCoreTokenIndexSet() view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) IsCoreTokenIndexSet() (bool, error) {
	return _HyperMintableERC20.Contract.IsCoreTokenIndexSet(&_HyperMintableERC20.CallOpts)
}

// IsLinkAuthority is a free data retrieval call binding the contract method 0x55c9cb07.
//
// Solidity: function isLinkAuthority(address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Caller) IsLinkAuthority(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "isLinkAuthority", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLinkAuthority is a free data retrieval call binding the contract method 0x55c9cb07.
//
// Solidity: function isLinkAuthority(address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) IsLinkAuthority(account common.Address) (bool, error) {
	return _HyperMintableERC20.Contract.IsLinkAuthority(&_HyperMintableERC20.CallOpts, account)
}

// IsLinkAuthority is a free data retrieval call binding the contract method 0x55c9cb07.
//
// Solidity: function isLinkAuthority(address account) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) IsLinkAuthority(account common.Address) (bool, error) {
	return _HyperMintableERC20.Contract.IsLinkAuthority(&_HyperMintableERC20.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20Session) Name() (string, error) {
	return _HyperMintableERC20.Contract.Name(&_HyperMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Name() (string, error) {
	return _HyperMintableERC20.Contract.Name(&_HyperMintableERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Nonces(opts *bind.CallOpts, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "nonces", owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) Nonces(owner_ common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.Nonces(&_HyperMintableERC20.CallOpts, owner_)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner_) view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Nonces(owner_ common.Address) (*big.Int, error) {
	return _HyperMintableERC20.Contract.Nonces(&_HyperMintableERC20.CallOpts, owner_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20Session) Owner() (common.Address, error) {
	return _HyperMintableERC20.Contract.Owner(&_HyperMintableERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Owner() (common.Address, error) {
	return _HyperMintableERC20.Contract.Owner(&_HyperMintableERC20.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_HyperMintableERC20 *HyperMintableERC20Caller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_HyperMintableERC20 *HyperMintableERC20Session) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20.Contract.PendingDefaultAdmin(&_HyperMintableERC20.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20.Contract.PendingDefaultAdmin(&_HyperMintableERC20.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_HyperMintableERC20 *HyperMintableERC20Caller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_HyperMintableERC20 *HyperMintableERC20Session) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20.Contract.PendingDefaultAdminDelay(&_HyperMintableERC20.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _HyperMintableERC20.Contract.PendingDefaultAdminDelay(&_HyperMintableERC20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HyperMintableERC20.Contract.SupportsInterface(&_HyperMintableERC20.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HyperMintableERC20.Contract.SupportsInterface(&_HyperMintableERC20.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20Session) Symbol() (string, error) {
	return _HyperMintableERC20.Contract.Symbol(&_HyperMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) Symbol() (string, error) {
	return _HyperMintableERC20.Contract.Symbol(&_HyperMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperMintableERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20Session) TotalSupply() (*big.Int, error) {
	return _HyperMintableERC20.Contract.TotalSupply(&_HyperMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HyperMintableERC20 *HyperMintableERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _HyperMintableERC20.Contract.TotalSupply(&_HyperMintableERC20.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.AcceptDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.AcceptDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Approve(&_HyperMintableERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Approve(&_HyperMintableERC20.TransactOpts, spender, value)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.BeginDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.BeginDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts, newAdmin)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Burn(&_HyperMintableERC20.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Burn(&_HyperMintableERC20.TransactOpts, _account, _amount)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.CancelDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.CancelDefaultAdminTransfer(&_HyperMintableERC20.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.ChangeDefaultAdminDelay(&_HyperMintableERC20.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.ChangeDefaultAdminDelay(&_HyperMintableERC20.TransactOpts, newDelay)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.GrantRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.GrantRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x238b4bc5.
//
// Solidity: function initialize(address initialOwner, address initialMinter, address initialLinker, string name_, string symbol_, uint8 decimals_) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialMinter common.Address, initialLinker common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "initialize", initialOwner, initialMinter, initialLinker, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x238b4bc5.
//
// Solidity: function initialize(address initialOwner, address initialMinter, address initialLinker, string name_, string symbol_, uint8 decimals_) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) Initialize(initialOwner common.Address, initialMinter common.Address, initialLinker common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Initialize(&_HyperMintableERC20.TransactOpts, initialOwner, initialMinter, initialLinker, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x238b4bc5.
//
// Solidity: function initialize(address initialOwner, address initialMinter, address initialLinker, string name_, string symbol_, uint8 decimals_) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Initialize(initialOwner common.Address, initialMinter common.Address, initialLinker common.Address, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Initialize(&_HyperMintableERC20.TransactOpts, initialOwner, initialMinter, initialLinker, name_, symbol_, decimals_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Mint(&_HyperMintableERC20.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Mint(&_HyperMintableERC20.TransactOpts, _account, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Permit(&_HyperMintableERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Permit(&_HyperMintableERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RenounceRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RenounceRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RevokeRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RevokeRole(&_HyperMintableERC20.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RollbackDefaultAdminDelay(&_HyperMintableERC20.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.RollbackDefaultAdminDelay(&_HyperMintableERC20.TransactOpts)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x60d4eab9.
//
// Solidity: function setCoreTokenIndex(uint64 index) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) SetCoreTokenIndex(opts *bind.TransactOpts, index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "setCoreTokenIndex", index)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x60d4eab9.
//
// Solidity: function setCoreTokenIndex(uint64 index) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) SetCoreTokenIndex(index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.SetCoreTokenIndex(&_HyperMintableERC20.TransactOpts, index)
}

// SetCoreTokenIndex is a paid mutator transaction binding the contract method 0x60d4eab9.
//
// Solidity: function setCoreTokenIndex(uint64 index) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) SetCoreTokenIndex(index uint64) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.SetCoreTokenIndex(&_HyperMintableERC20.TransactOpts, index)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0x89e98291.
//
// Solidity: function setHyperCoreDeployer(address finalizer) returns()
func (_HyperMintableERC20 *HyperMintableERC20Transactor) SetHyperCoreDeployer(opts *bind.TransactOpts, finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "setHyperCoreDeployer", finalizer)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0x89e98291.
//
// Solidity: function setHyperCoreDeployer(address finalizer) returns()
func (_HyperMintableERC20 *HyperMintableERC20Session) SetHyperCoreDeployer(finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.SetHyperCoreDeployer(&_HyperMintableERC20.TransactOpts, finalizer)
}

// SetHyperCoreDeployer is a paid mutator transaction binding the contract method 0x89e98291.
//
// Solidity: function setHyperCoreDeployer(address finalizer) returns()
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) SetHyperCoreDeployer(finalizer common.Address) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.SetHyperCoreDeployer(&_HyperMintableERC20.TransactOpts, finalizer)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Transfer(&_HyperMintableERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.Transfer(&_HyperMintableERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferFrom(&_HyperMintableERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferFrom(&_HyperMintableERC20.TransactOpts, from, to, value)
}

// TransferToCore is a paid mutator transaction binding the contract method 0x45f72615.
//
// Solidity: function transferToCore(uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) TransferToCore(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "transferToCore", amount)
}

// TransferToCore is a paid mutator transaction binding the contract method 0x45f72615.
//
// Solidity: function transferToCore(uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20Session) TransferToCore(amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferToCore(&_HyperMintableERC20.TransactOpts, amount)
}

// TransferToCore is a paid mutator transaction binding the contract method 0x45f72615.
//
// Solidity: function transferToCore(uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) TransferToCore(amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferToCore(&_HyperMintableERC20.TransactOpts, amount)
}

// TransferToCoreFor is a paid mutator transaction binding the contract method 0x9438568d.
//
// Solidity: function transferToCoreFor(address coreRecipient, uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20Transactor) TransferToCoreFor(opts *bind.TransactOpts, coreRecipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.contract.Transact(opts, "transferToCoreFor", coreRecipient, amount)
}

// TransferToCoreFor is a paid mutator transaction binding the contract method 0x9438568d.
//
// Solidity: function transferToCoreFor(address coreRecipient, uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20Session) TransferToCoreFor(coreRecipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferToCoreFor(&_HyperMintableERC20.TransactOpts, coreRecipient, amount)
}

// TransferToCoreFor is a paid mutator transaction binding the contract method 0x9438568d.
//
// Solidity: function transferToCoreFor(address coreRecipient, uint256 amount) returns(uint256 sent)
func (_HyperMintableERC20 *HyperMintableERC20TransactorSession) TransferToCoreFor(coreRecipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HyperMintableERC20.Contract.TransferToCoreFor(&_HyperMintableERC20.TransactOpts, coreRecipient, amount)
}

// HyperMintableERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HyperMintableERC20 contract.
type HyperMintableERC20ApprovalIterator struct {
	Event *HyperMintableERC20Approval // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20Approval)
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
		it.Event = new(HyperMintableERC20Approval)
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
func (it *HyperMintableERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20Approval represents a Approval event raised by the HyperMintableERC20 contract.
type HyperMintableERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HyperMintableERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20ApprovalIterator{contract: _HyperMintableERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20Approval)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseApproval(log types.Log) (*HyperMintableERC20Approval, error) {
	event := new(HyperMintableERC20Approval)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20CoreTokenIndexSetIterator is returned from FilterCoreTokenIndexSet and is used to iterate over the raw logs and unpacked data for CoreTokenIndexSet events raised by the HyperMintableERC20 contract.
type HyperMintableERC20CoreTokenIndexSetIterator struct {
	Event *HyperMintableERC20CoreTokenIndexSet // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20CoreTokenIndexSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20CoreTokenIndexSet)
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
		it.Event = new(HyperMintableERC20CoreTokenIndexSet)
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
func (it *HyperMintableERC20CoreTokenIndexSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20CoreTokenIndexSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20CoreTokenIndexSet represents a CoreTokenIndexSet event raised by the HyperMintableERC20 contract.
type HyperMintableERC20CoreTokenIndexSet struct {
	Index               uint64
	EvmExtraWeiDecimals int8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCoreTokenIndexSet is a free log retrieval operation binding the contract event 0x2bd02f43024c0d42c7d60ce7b4c93d53b2d420af5fa93bee4f57d5de645de7c0.
//
// Solidity: event CoreTokenIndexSet(uint64 indexed index, int8 evmExtraWeiDecimals)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterCoreTokenIndexSet(opts *bind.FilterOpts, index []uint64) (*HyperMintableERC20CoreTokenIndexSetIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "CoreTokenIndexSet", indexRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20CoreTokenIndexSetIterator{contract: _HyperMintableERC20.contract, event: "CoreTokenIndexSet", logs: logs, sub: sub}, nil
}

// WatchCoreTokenIndexSet is a free log subscription operation binding the contract event 0x2bd02f43024c0d42c7d60ce7b4c93d53b2d420af5fa93bee4f57d5de645de7c0.
//
// Solidity: event CoreTokenIndexSet(uint64 indexed index, int8 evmExtraWeiDecimals)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchCoreTokenIndexSet(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20CoreTokenIndexSet, index []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "CoreTokenIndexSet", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20CoreTokenIndexSet)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "CoreTokenIndexSet", log); err != nil {
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

// ParseCoreTokenIndexSet is a log parse operation binding the contract event 0x2bd02f43024c0d42c7d60ce7b4c93d53b2d420af5fa93bee4f57d5de645de7c0.
//
// Solidity: event CoreTokenIndexSet(uint64 indexed index, int8 evmExtraWeiDecimals)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseCoreTokenIndexSet(log types.Log) (*HyperMintableERC20CoreTokenIndexSet, error) {
	event := new(HyperMintableERC20CoreTokenIndexSet)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "CoreTokenIndexSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20DefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminDelayChangeCanceledIterator struct {
	Event *HyperMintableERC20DefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20DefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20DefaultAdminDelayChangeCanceled)
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
		it.Event = new(HyperMintableERC20DefaultAdminDelayChangeCanceled)
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
func (it *HyperMintableERC20DefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20DefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20DefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*HyperMintableERC20DefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20DefaultAdminDelayChangeCanceledIterator{contract: _HyperMintableERC20.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20DefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20DefaultAdminDelayChangeCanceled)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*HyperMintableERC20DefaultAdminDelayChangeCanceled, error) {
	event := new(HyperMintableERC20DefaultAdminDelayChangeCanceled)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20DefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminDelayChangeScheduledIterator struct {
	Event *HyperMintableERC20DefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20DefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20DefaultAdminDelayChangeScheduled)
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
		it.Event = new(HyperMintableERC20DefaultAdminDelayChangeScheduled)
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
func (it *HyperMintableERC20DefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20DefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20DefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*HyperMintableERC20DefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20DefaultAdminDelayChangeScheduledIterator{contract: _HyperMintableERC20.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20DefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20DefaultAdminDelayChangeScheduled)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*HyperMintableERC20DefaultAdminDelayChangeScheduled, error) {
	event := new(HyperMintableERC20DefaultAdminDelayChangeScheduled)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20DefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminTransferCanceledIterator struct {
	Event *HyperMintableERC20DefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20DefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20DefaultAdminTransferCanceled)
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
		it.Event = new(HyperMintableERC20DefaultAdminTransferCanceled)
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
func (it *HyperMintableERC20DefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20DefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20DefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*HyperMintableERC20DefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20DefaultAdminTransferCanceledIterator{contract: _HyperMintableERC20.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20DefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20DefaultAdminTransferCanceled)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseDefaultAdminTransferCanceled(log types.Log) (*HyperMintableERC20DefaultAdminTransferCanceled, error) {
	event := new(HyperMintableERC20DefaultAdminTransferCanceled)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20DefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminTransferScheduledIterator struct {
	Event *HyperMintableERC20DefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20DefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20DefaultAdminTransferScheduled)
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
		it.Event = new(HyperMintableERC20DefaultAdminTransferScheduled)
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
func (it *HyperMintableERC20DefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20DefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20DefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the HyperMintableERC20 contract.
type HyperMintableERC20DefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*HyperMintableERC20DefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20DefaultAdminTransferScheduledIterator{contract: _HyperMintableERC20.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20DefaultAdminTransferScheduled)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseDefaultAdminTransferScheduled(log types.Log) (*HyperMintableERC20DefaultAdminTransferScheduled, error) {
	event := new(HyperMintableERC20DefaultAdminTransferScheduled)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the HyperMintableERC20 contract.
type HyperMintableERC20EIP712DomainChangedIterator struct {
	Event *HyperMintableERC20EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20EIP712DomainChanged)
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
		it.Event = new(HyperMintableERC20EIP712DomainChanged)
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
func (it *HyperMintableERC20EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20EIP712DomainChanged represents a EIP712DomainChanged event raised by the HyperMintableERC20 contract.
type HyperMintableERC20EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*HyperMintableERC20EIP712DomainChangedIterator, error) {

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20EIP712DomainChangedIterator{contract: _HyperMintableERC20.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20EIP712DomainChanged)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseEIP712DomainChanged(log types.Log) (*HyperMintableERC20EIP712DomainChanged, error) {
	event := new(HyperMintableERC20EIP712DomainChanged)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20HyperCoreDeployerSetIterator is returned from FilterHyperCoreDeployerSet and is used to iterate over the raw logs and unpacked data for HyperCoreDeployerSet events raised by the HyperMintableERC20 contract.
type HyperMintableERC20HyperCoreDeployerSetIterator struct {
	Event *HyperMintableERC20HyperCoreDeployerSet // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20HyperCoreDeployerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20HyperCoreDeployerSet)
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
		it.Event = new(HyperMintableERC20HyperCoreDeployerSet)
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
func (it *HyperMintableERC20HyperCoreDeployerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20HyperCoreDeployerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20HyperCoreDeployerSet represents a HyperCoreDeployerSet event raised by the HyperMintableERC20 contract.
type HyperMintableERC20HyperCoreDeployerSet struct {
	PreviousDeployer common.Address
	NewDeployer      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHyperCoreDeployerSet is a free log retrieval operation binding the contract event 0xcae9891438942589688090dd5afab743aa8c874826bb838ba0c629ba75366e3e.
//
// Solidity: event HyperCoreDeployerSet(address indexed previousDeployer, address indexed newDeployer)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterHyperCoreDeployerSet(opts *bind.FilterOpts, previousDeployer []common.Address, newDeployer []common.Address) (*HyperMintableERC20HyperCoreDeployerSetIterator, error) {

	var previousDeployerRule []interface{}
	for _, previousDeployerItem := range previousDeployer {
		previousDeployerRule = append(previousDeployerRule, previousDeployerItem)
	}
	var newDeployerRule []interface{}
	for _, newDeployerItem := range newDeployer {
		newDeployerRule = append(newDeployerRule, newDeployerItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "HyperCoreDeployerSet", previousDeployerRule, newDeployerRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20HyperCoreDeployerSetIterator{contract: _HyperMintableERC20.contract, event: "HyperCoreDeployerSet", logs: logs, sub: sub}, nil
}

// WatchHyperCoreDeployerSet is a free log subscription operation binding the contract event 0xcae9891438942589688090dd5afab743aa8c874826bb838ba0c629ba75366e3e.
//
// Solidity: event HyperCoreDeployerSet(address indexed previousDeployer, address indexed newDeployer)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchHyperCoreDeployerSet(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20HyperCoreDeployerSet, previousDeployer []common.Address, newDeployer []common.Address) (event.Subscription, error) {

	var previousDeployerRule []interface{}
	for _, previousDeployerItem := range previousDeployer {
		previousDeployerRule = append(previousDeployerRule, previousDeployerItem)
	}
	var newDeployerRule []interface{}
	for _, newDeployerItem := range newDeployer {
		newDeployerRule = append(newDeployerRule, newDeployerItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "HyperCoreDeployerSet", previousDeployerRule, newDeployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20HyperCoreDeployerSet)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "HyperCoreDeployerSet", log); err != nil {
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

// ParseHyperCoreDeployerSet is a log parse operation binding the contract event 0xcae9891438942589688090dd5afab743aa8c874826bb838ba0c629ba75366e3e.
//
// Solidity: event HyperCoreDeployerSet(address indexed previousDeployer, address indexed newDeployer)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseHyperCoreDeployerSet(log types.Log) (*HyperMintableERC20HyperCoreDeployerSet, error) {
	event := new(HyperMintableERC20HyperCoreDeployerSet)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "HyperCoreDeployerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HyperMintableERC20 contract.
type HyperMintableERC20InitializedIterator struct {
	Event *HyperMintableERC20Initialized // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20Initialized)
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
		it.Event = new(HyperMintableERC20Initialized)
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
func (it *HyperMintableERC20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20Initialized represents a Initialized event raised by the HyperMintableERC20 contract.
type HyperMintableERC20Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterInitialized(opts *bind.FilterOpts) (*HyperMintableERC20InitializedIterator, error) {

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20InitializedIterator{contract: _HyperMintableERC20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20Initialized) (event.Subscription, error) {

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20Initialized)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseInitialized(log types.Log) (*HyperMintableERC20Initialized, error) {
	event := new(HyperMintableERC20Initialized)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleAdminChangedIterator struct {
	Event *HyperMintableERC20RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20RoleAdminChanged)
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
		it.Event = new(HyperMintableERC20RoleAdminChanged)
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
func (it *HyperMintableERC20RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20RoleAdminChanged represents a RoleAdminChanged event raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*HyperMintableERC20RoleAdminChangedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20RoleAdminChangedIterator{contract: _HyperMintableERC20.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20RoleAdminChanged)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseRoleAdminChanged(log types.Log) (*HyperMintableERC20RoleAdminChanged, error) {
	event := new(HyperMintableERC20RoleAdminChanged)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleGrantedIterator struct {
	Event *HyperMintableERC20RoleGranted // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20RoleGranted)
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
		it.Event = new(HyperMintableERC20RoleGranted)
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
func (it *HyperMintableERC20RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20RoleGranted represents a RoleGranted event raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HyperMintableERC20RoleGrantedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20RoleGrantedIterator{contract: _HyperMintableERC20.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20RoleGranted)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseRoleGranted(log types.Log) (*HyperMintableERC20RoleGranted, error) {
	event := new(HyperMintableERC20RoleGranted)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleRevokedIterator struct {
	Event *HyperMintableERC20RoleRevoked // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20RoleRevoked)
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
		it.Event = new(HyperMintableERC20RoleRevoked)
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
func (it *HyperMintableERC20RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20RoleRevoked represents a RoleRevoked event raised by the HyperMintableERC20 contract.
type HyperMintableERC20RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*HyperMintableERC20RoleRevokedIterator, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20RoleRevokedIterator{contract: _HyperMintableERC20.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20RoleRevoked)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseRoleRevoked(log types.Log) (*HyperMintableERC20RoleRevoked, error) {
	event := new(HyperMintableERC20RoleRevoked)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperMintableERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HyperMintableERC20 contract.
type HyperMintableERC20TransferIterator struct {
	Event *HyperMintableERC20Transfer // Event containing the contract specifics and raw log

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
func (it *HyperMintableERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperMintableERC20Transfer)
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
		it.Event = new(HyperMintableERC20Transfer)
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
func (it *HyperMintableERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperMintableERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperMintableERC20Transfer represents a Transfer event raised by the HyperMintableERC20 contract.
type HyperMintableERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HyperMintableERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HyperMintableERC20TransferIterator{contract: _HyperMintableERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HyperMintableERC20 *HyperMintableERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HyperMintableERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HyperMintableERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperMintableERC20Transfer)
				if err := _HyperMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_HyperMintableERC20 *HyperMintableERC20Filterer) ParseTransfer(log types.Log) (*HyperMintableERC20Transfer, error) {
	event := new(HyperMintableERC20Transfer)
	if err := _HyperMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
