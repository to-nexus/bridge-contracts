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

// BridgeExecutorMetaData contains all meta data concerning the BridgeExecutor contract.
var BridgeExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"addWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"executeExtraCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"consumed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isMethodCheckEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"isWhitelistedMethod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isWhitelistedTarget\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReturnDataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"removeWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMaxReturnDataSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMethodCheckEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MaxReturnDataSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MethodCheckEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEFailedToReturnNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidMaxReturnDataSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEMethodNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BENotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"a8582dc2": "addWhitelistMethods(address,bytes4[])",
		"61c8c479": "addWhitelistTarget(address)",
		"1eeed513": "executeExtraCall(uint256,uint256,address,address,uint256,bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"ab779650": "isMethodCheckEnabled(address)",
		"b8e25592": "isWhitelistedMethod(address,bytes4)",
		"907caa00": "isWhitelistedTarget(address)",
		"a1b8442a": "maxReturnDataSize()",
		"d04323c5": "recoverToken(address,uint256,address)",
		"4f8249a2": "removeWhitelistMethods(address,bytes4[])",
		"11735ea5": "removeWhitelistTarget(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"8a7fdc75": "setMaxReturnDataSize(uint256)",
		"47239f71": "setMethodCheckEnabled(address,bool)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x6080346100b957601f61166e38819003918201601f19168301916001600160401b038311848410176100bd5780849260409485528339810103126100b957610052602061004b836100d1565b92016100d1565b610400600355906001600160a01b038116156100aa578061007561007b926100e5565b5061015b565b506001600160a01b03811661009a575b60405161136c90816102828239f35b6100a3906101ee565b505f61008b565b63484b059f60e11b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100b957565b6001600160a01b0381165f9081525f51602061164e5f395f51905f52602052604090205460ff16610156576001600160a01b03165f8181525f51602061164e5f395f51905f5260205260408120805460ff191660011790553391905f5160206115ee5f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f51602061160e5f395f51905f52602052604090205460ff16610156576001600160a01b03165f8181525f51602061160e5f395f51905f5260205260408120805460ff191660011790553391907fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775905f5160206115ee5f395f51905f529080a4600190565b6001600160a01b0381165f9081525f51602061162e5f395f51905f52602052604090205460ff16610156576001600160a01b03165f8181525f51602061162e5f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f5160206115ee5f395f51905f529080a460019056fe6080604052600436101561001a575b3615610018575f80fd5b005b5f3560e01c806301ffc9a71461014957806311735ea5146101445780631eeed5131461013f578063248a9ca31461013a5780632f2ff15d1461013557806336568abe1461013057806347239f711461012b5780634f8249a21461012657806361c8c479146101215780638a7fdc751461011c578063907caa001461011757806391d1485414610112578063a1b8442a1461010d578063a217fddf14610108578063a8582dc214610103578063ab779650146100fe578063b8e25592146100f9578063d04323c5146100f45763d547741f0361000e57610907565b610863565b610818565b6107d5565b610727565b61070d565b6106f0565b6106af565b61066f565b610603565b610565565b6104af565b6103be565b610376565b61033d565b610317565b61028a565b6101cb565b610164565b6001600160e01b031981160361016057565b5f80fd5b346101605760203660031901126101605760206004356101838161014e565b63ffffffff60e01b16637965db0b60e01b81149081156101a9575b506040519015158152f35b6301ffc9a760e01b1490505f61019e565b6001600160a01b0381160361016057565b34610160576020366003190112610160576004356101e8816101ba565b6101f0610ea9565b6001600160a01b03165f8181526001602052604090205460ff161561024957805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b9060609260209183526040828401528051918291826040860152018484015e5f828201840152601f01601f1916010190565b60c0366003190112610160576044356024356004356102a8836101ba565b6064356102b4816101ba565b60a4359360843591906001600160401b0386116101605736602387011215610160576004860135946001600160401b038611610160573660248789010111610160576024610303970194610b27565b9061031360405192839283610258565b0390f35b34610160576020366003190112610160576020610335600435610e5e565b604051908152f35b3461016057604036600319011261016057610018602435600435610360826101ba565b61037161036c82610e5e565b610f85565b61119a565b3461016057604036600319011261016057600435602435610396816101ba565b336001600160a01b038216036103af5761001891611210565b63334bd91960e11b5f5260045ffd5b34610160576040366003190112610160576004356103db816101ba565b602435908115158092036101605760207f5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e391610415610ea9565b6001600160a01b03169261042a841515610e6f565b835f526001825260405f20805461ff008360081b169061ff001916179055604051908152a2005b604060031982011261016057600435610469816101ba565b916024356001600160401b0381116101605782602382011215610160576004810135926001600160401b0384116101605760248460051b83010111610160576024019190565b34610160576104bd36610451565b90916104c7610ea9565b6001600160a01b03166104db811515610e6f565b5f5b8281106104e657005b600190825f52600260205261051360405f20610503838789610e85565b359061050e8261014e565b610992565b805460ff19169055610526818587610e85565b356105308161014e565b6001600160e01b031916837f5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a5f80a3016104dd565b3461016057602036600319011261016057600435610582816101ba565b61058a610ea9565b6001600160a01b031661059e811515610e6f565b805f52600160205260ff60405f2054166105f457805f5260016020526105ce60405f20600160ff19825416179055565b7f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b346101605760203660031901126101605760043561061f610ea9565b604081106106605760407f3d613bec426ed2a2fb9a9244f66939a8de8f29c84fa65bfabde181e0996b68f891600354908060035582519182526020820152a1005b6308de317760e21b5f5260045ffd5b346101605760203660031901126101605760043561068c816101ba565b60018060a01b03165f526001602052602060ff60405f2054166040519015158152f35b3461016057604036600319011261016057602060ff6106e46024356004356106d6826101ba565b5f525f845260405f2061093b565b54166040519015158152f35b34610160575f366003190112610160576020600354604051908152f35b34610160575f3660031901126101605760206040515f8152f35b346101605761073536610451565b909161073f610ea9565b6001600160a01b0316610753811515610e6f565b5f5b82811061075e57005b600190825f52600260205261078b61077e60405f2061050384888a610e85565b805460ff19166001179055565b610796818587610e85565b356107a08161014e565b6001600160e01b031916837f17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd5f80a301610755565b34610160576020366003190112610160576004356107f2816101ba565b60018060a01b03165f526001602052602060ff60405f205460081c166040519015158152f35b3461016057604036600319011261016057602060ff6106e460043561083c816101ba565b602435906108498261014e565b6001600160a01b03165f9081526002855260409020610992565b3461016057606036600319011261016057600435610880816101ba565b60243560443561088f816101ba565b610897610ea9565b6001600160a01b03811692831515806108f5575b156108e6576001600160a01b0316600181036108dc5750505f80808093610018955af16108d6610ae2565b50610b11565b6100189350611165565b63b56bf2e760e01b5f5260045ffd5b506001600160a01b03811615156108ab565b346101605760403660031901126101605761001860243560043561092a826101ba565b61093661036c82610e5e565b611210565b9060018060a01b03165f5260205260405f2090565b1561095757565b6308ba07a960e21b5f5260045ffd5b1561096d57565b6309b7a22f60e01b5f5260045ffd5b1561098357565b63a6bec22b60e01b5f5260045ffd5b9063ffffffff60e01b165f5260205260405f2090565b156109af57565b630850473d60e31b5f5260045ffd5b156109c557565b63867deca760e01b5f5260045ffd5b634e487b7160e01b5f52604160045260245ffd5b601f909101601f19168101906001600160401b03821190821017610a0b57604052565b6109d4565b90816020910312610160575190565b6001600160a01b03909116815260200190565b6040513d5f823e3d90fd5b91908203918211610a4a57565b634e487b7160e01b5f52601160045260245ffd5b909291928360141161016057831161016057601401916013190190565b6001600160401b038111610a0b57601f01601f191660200190565b929192610aa282610a7b565b91610ab060405193846109e8565b829481845281830111610160578281602093845f960137010152565b15610ad357565b63e2ef9c2b60e01b5f5260045ffd5b3d15610b0c573d90610af382610a7b565b91610b0160405193846109e8565b82523d5f602084013e565b606090565b15610b1857565b6317e90b5960e21b5f5260045ffd5b5050949390919294610b37610f17565b5f5160206112f75f395f51905f525c610e4f5760015f5160206112f75f395f51905f525d610b66841515610950565b610b736018831015610966565b853560601c5f81815260016020526040902090968791610ba39054610b9a60ff821661097c565b60081c60ff1690565b610e16575b6001600160a01b0383166001811494909187918615610dfb57610bcc8334146109be565b8615610d8d57610c239291610c0e610c07610c1593610bec473490610a3d565b985b8b159687610d7c575b50508a15610d7457808d94610a5e565b3691610a96565b908c611121565b9a9091610d64575b50610acc565b8315610cf7575047905b8315610cd75780821115610ccf57610c4491610a3d565b848110610cc057505f5b935b8480821115610cb757610c6291610a3d565b915b82610c7a575b50505050610c76610fbb565b9190565b15610ca85750610c9f915f918291829182916001600160a01b03165af16108d6610ae2565b5f808080610c6a565b91610cb292611165565b610c9f565b50505f91610c64565b610cca9085610a3d565b610c4e565b50505f610c44565b9080821115610cef57610ce991610a3d565b93610c50565b50505f610ce9565b602060405180926370a0823160e01b82528180610d173060048301610a1f565b03915afa908115610d5f575f91610d30575b5090610c2d565b610d52915060203d602011610d58575b610d4a81836109e8565b810190610a10565b5f610d29565b503d610d40565b610a32565b610d6e9086611033565b5f610c1d565b805f94610a5e565b610d86918c6110f8565b8e8d610bf7565b6040516370a0823160e01b8152919450915060208180610db03060048301610a1f565b0381865afa918215610d5f578a94610c0e610c07610c1593610c23968d965f91610ddc575b5098610bee565b610df5915060203d602011610d5857610d4a81836109e8565b5f610dd5565b610e0534156109be565b610e1183303389610fcd565b610bcc565b610e4a610e45610e3e601484013561050e8660018060a01b03165f52600260205260405f2090565b5460ff1690565b6109a8565b610ba8565b633ee5aeb560e01b5f5260045ffd5b5f525f602052600160405f20015490565b15610e7657565b63484b059f60e11b5f5260045ffd5b9190811015610e955760051b0190565b634e487b7160e01b5f52603260045260245ffd5b5f5160206113175f395f51905f525f90815260205260ff610eea337f7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ec61093b565b541615610ef357565b63e2517d3f60e01b5f52336004525f5160206113175f395f51905f5260245260445ffd5b5f5160206112d75f395f51905f525f90815260205260ff610f58337fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d706961093b565b541615610f6157565b63e2517d3f60e01b5f52336004525f5160206112d75f395f51905f5260245260445ffd5b805f525f60205260ff610f9b3360405f2061093b565b541615610fa55750565b63e2517d3f60e01b5f523360045260245260445ffd5b5f5f5160206112f75f395f51905f525d565b6040516323b872dd60e01b60208201526001600160a01b039283166024820152929091166044830152606480830193909352918152611016916110116084836109e8565b61127e565b565b6001600160a01b039091168152602081019190915260400190565b6040519060205f81840163095ea7b360e01b815261106785611059848960248401611018565b03601f1981018752866109e8565b84519082855af15f51903d816110cc575b501590505b61108657505050565b60405163095ea7b360e01b60208201526001600160a01b039390931660248401525f604480850191909152835261101692611011906110c66064826109e8565b8261127e565b151590506110ec575061107d6001600160a01b0382163b15155b5f611078565b600161107d91146110e6565b91909160205f60405193611067856110598582019363095ea7b360e01b85528960248401611018565b9290915f8060209260035496604051968583519301915af1933d9080821161115d575b50808452805f8386013e601f01601f1916830101604052565b90505f611144565b611011611016939261118c60405194859263a9059cbb60e01b602085015260248401611018565b03601f1981018452836109e8565b805f525f60205260ff6111b08360405f2061093b565b541661120a57805f525f6020526111ca8260405f2061093b565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b50505f90565b805f525f60205260ff6112268360405f2061093b565b54161561120a57805f525f6020526112418260405f2061093b565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b905f602091828151910182855af115610a32575f513d6112cd57506001600160a01b0381163b155b6112ad5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b600114156112a656fed8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e639b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a264697066735822122000301955764226535a26d1828c7babdfc3fcceaf75314326ed3e7d1a8d27968464736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d7d7ffb7a348e1c6a02869081a26547b49160dd3df72d1d75a570eb9b698292ecdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// BridgeExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeExecutorMetaData.ABI instead.
var BridgeExecutorABI = BridgeExecutorMetaData.ABI

// Deprecated: Use BridgeExecutorMetaData.Sigs instead.
// BridgeExecutorFuncSigs maps the 4-byte function signature to its string representation.
var BridgeExecutorFuncSigs = BridgeExecutorMetaData.Sigs

// BridgeExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeExecutorMetaData.Bin instead.
var BridgeExecutorBin = BridgeExecutorMetaData.Bin

// DeployBridgeExecutor deploys a new Ethereum contract, binding an instance of BridgeExecutor to it.
func DeployBridgeExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, executor common.Address) (common.Address, *types.Transaction, *BridgeExecutor, error) {
	parsed, err := BridgeExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeExecutorBin), backend, owner, executor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeExecutor{BridgeExecutorCaller: BridgeExecutorCaller{contract: contract}, BridgeExecutorTransactor: BridgeExecutorTransactor{contract: contract}, BridgeExecutorFilterer: BridgeExecutorFilterer{contract: contract}}, nil
}

// BridgeExecutor is an auto generated Go binding around an Ethereum contract.
type BridgeExecutor struct {
	BridgeExecutorCaller     // Read-only binding to the contract
	BridgeExecutorTransactor // Write-only binding to the contract
	BridgeExecutorFilterer   // Log filterer for contract events
}

// BridgeExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeExecutorSession struct {
	Contract     *BridgeExecutor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeExecutorCallerSession struct {
	Contract *BridgeExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeExecutorTransactorSession struct {
	Contract     *BridgeExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeExecutorRaw struct {
	Contract *BridgeExecutor // Generic contract binding to access the raw methods on
}

// BridgeExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeExecutorCallerRaw struct {
	Contract *BridgeExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeExecutorTransactorRaw struct {
	Contract *BridgeExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeExecutor creates a new instance of BridgeExecutor, bound to a specific deployed contract.
func NewBridgeExecutor(address common.Address, backend bind.ContractBackend) (*BridgeExecutor, error) {
	contract, err := bindBridgeExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutor{BridgeExecutorCaller: BridgeExecutorCaller{contract: contract}, BridgeExecutorTransactor: BridgeExecutorTransactor{contract: contract}, BridgeExecutorFilterer: BridgeExecutorFilterer{contract: contract}}, nil
}

// NewBridgeExecutorCaller creates a new read-only instance of BridgeExecutor, bound to a specific deployed contract.
func NewBridgeExecutorCaller(address common.Address, caller bind.ContractCaller) (*BridgeExecutorCaller, error) {
	contract, err := bindBridgeExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorCaller{contract: contract}, nil
}

// NewBridgeExecutorTransactor creates a new write-only instance of BridgeExecutor, bound to a specific deployed contract.
func NewBridgeExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeExecutorTransactor, error) {
	contract, err := bindBridgeExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorTransactor{contract: contract}, nil
}

// NewBridgeExecutorFilterer creates a new log filterer instance of BridgeExecutor, bound to a specific deployed contract.
func NewBridgeExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeExecutorFilterer, error) {
	contract, err := bindBridgeExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorFilterer{contract: contract}, nil
}

// bindBridgeExecutor binds a generic wrapper to an already deployed contract.
func bindBridgeExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeExecutor *BridgeExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeExecutor.Contract.BridgeExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeExecutor *BridgeExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.BridgeExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeExecutor *BridgeExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.BridgeExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeExecutor *BridgeExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeExecutor *BridgeExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeExecutor *BridgeExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeExecutor.Contract.DEFAULTADMINROLE(&_BridgeExecutor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeExecutor.Contract.DEFAULTADMINROLE(&_BridgeExecutor.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeExecutor.Contract.GetRoleAdmin(&_BridgeExecutor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecutor *BridgeExecutorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeExecutor.Contract.GetRoleAdmin(&_BridgeExecutor.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecutor *BridgeExecutorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeExecutor.Contract.HasRole(&_BridgeExecutor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeExecutor.Contract.HasRole(&_BridgeExecutor.CallOpts, role, account)
}

// IsMethodCheckEnabled is a free data retrieval call binding the contract method 0xab779650.
//
// Solidity: function isMethodCheckEnabled(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCaller) IsMethodCheckEnabled(opts *bind.CallOpts, target common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "isMethodCheckEnabled", target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMethodCheckEnabled is a free data retrieval call binding the contract method 0xab779650.
//
// Solidity: function isMethodCheckEnabled(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorSession) IsMethodCheckEnabled(target common.Address) (bool, error) {
	return _BridgeExecutor.Contract.IsMethodCheckEnabled(&_BridgeExecutor.CallOpts, target)
}

// IsMethodCheckEnabled is a free data retrieval call binding the contract method 0xab779650.
//
// Solidity: function isMethodCheckEnabled(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCallerSession) IsMethodCheckEnabled(target common.Address) (bool, error) {
	return _BridgeExecutor.Contract.IsMethodCheckEnabled(&_BridgeExecutor.CallOpts, target)
}

// IsWhitelistedMethod is a free data retrieval call binding the contract method 0xb8e25592.
//
// Solidity: function isWhitelistedMethod(address target, bytes4 methodID) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCaller) IsWhitelistedMethod(opts *bind.CallOpts, target common.Address, methodID [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "isWhitelistedMethod", target, methodID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedMethod is a free data retrieval call binding the contract method 0xb8e25592.
//
// Solidity: function isWhitelistedMethod(address target, bytes4 methodID) view returns(bool)
func (_BridgeExecutor *BridgeExecutorSession) IsWhitelistedMethod(target common.Address, methodID [4]byte) (bool, error) {
	return _BridgeExecutor.Contract.IsWhitelistedMethod(&_BridgeExecutor.CallOpts, target, methodID)
}

// IsWhitelistedMethod is a free data retrieval call binding the contract method 0xb8e25592.
//
// Solidity: function isWhitelistedMethod(address target, bytes4 methodID) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCallerSession) IsWhitelistedMethod(target common.Address, methodID [4]byte) (bool, error) {
	return _BridgeExecutor.Contract.IsWhitelistedMethod(&_BridgeExecutor.CallOpts, target, methodID)
}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCaller) IsWhitelistedTarget(opts *bind.CallOpts, target common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "isWhitelistedTarget", target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorSession) IsWhitelistedTarget(target common.Address) (bool, error) {
	return _BridgeExecutor.Contract.IsWhitelistedTarget(&_BridgeExecutor.CallOpts, target)
}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCallerSession) IsWhitelistedTarget(target common.Address) (bool, error) {
	return _BridgeExecutor.Contract.IsWhitelistedTarget(&_BridgeExecutor.CallOpts, target)
}

// MaxReturnDataSize is a free data retrieval call binding the contract method 0xa1b8442a.
//
// Solidity: function maxReturnDataSize() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorCaller) MaxReturnDataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "maxReturnDataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReturnDataSize is a free data retrieval call binding the contract method 0xa1b8442a.
//
// Solidity: function maxReturnDataSize() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorSession) MaxReturnDataSize() (*big.Int, error) {
	return _BridgeExecutor.Contract.MaxReturnDataSize(&_BridgeExecutor.CallOpts)
}

// MaxReturnDataSize is a free data retrieval call binding the contract method 0xa1b8442a.
//
// Solidity: function maxReturnDataSize() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorCallerSession) MaxReturnDataSize() (*big.Int, error) {
	return _BridgeExecutor.Contract.MaxReturnDataSize(&_BridgeExecutor.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecutor *BridgeExecutorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeExecutor.Contract.SupportsInterface(&_BridgeExecutor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecutor *BridgeExecutorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeExecutor.Contract.SupportsInterface(&_BridgeExecutor.CallOpts, interfaceId)
}

// AddWhitelistMethods is a paid mutator transaction binding the contract method 0xa8582dc2.
//
// Solidity: function addWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) AddWhitelistMethods(opts *bind.TransactOpts, target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "addWhitelistMethods", target, methodIDs)
}

// AddWhitelistMethods is a paid mutator transaction binding the contract method 0xa8582dc2.
//
// Solidity: function addWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorSession) AddWhitelistMethods(target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.AddWhitelistMethods(&_BridgeExecutor.TransactOpts, target, methodIDs)
}

// AddWhitelistMethods is a paid mutator transaction binding the contract method 0xa8582dc2.
//
// Solidity: function addWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) AddWhitelistMethods(target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.AddWhitelistMethods(&_BridgeExecutor.TransactOpts, target, methodIDs)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) AddWhitelistTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "addWhitelistTarget", target)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorSession) AddWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.AddWhitelistTarget(&_BridgeExecutor.TransactOpts, target)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) AddWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.AddWhitelistTarget(&_BridgeExecutor.TransactOpts, target)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 , uint256 , address toToken, address to, uint256 value, bytes extraData) payable returns(uint256 consumed, bytes returnData)
func (_BridgeExecutor *BridgeExecutorTransactor) ExecuteExtraCall(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "executeExtraCall", arg0, arg1, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 , uint256 , address toToken, address to, uint256 value, bytes extraData) payable returns(uint256 consumed, bytes returnData)
func (_BridgeExecutor *BridgeExecutorSession) ExecuteExtraCall(arg0 *big.Int, arg1 *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, arg0, arg1, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 , uint256 , address toToken, address to, uint256 value, bytes extraData) payable returns(uint256 consumed, bytes returnData)
func (_BridgeExecutor *BridgeExecutorTransactorSession) ExecuteExtraCall(arg0 *big.Int, arg1 *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, arg0, arg1, toToken, to, value, extraData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.GrantRole(&_BridgeExecutor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.GrantRole(&_BridgeExecutor.TransactOpts, role, account)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) RecoverToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "recoverToken", token, amount, recipient)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecutor *BridgeExecutorSession) RecoverToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RecoverToken(&_BridgeExecutor.TransactOpts, token, amount, recipient)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) RecoverToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RecoverToken(&_BridgeExecutor.TransactOpts, token, amount, recipient)
}

// RemoveWhitelistMethods is a paid mutator transaction binding the contract method 0x4f8249a2.
//
// Solidity: function removeWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) RemoveWhitelistMethods(opts *bind.TransactOpts, target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "removeWhitelistMethods", target, methodIDs)
}

// RemoveWhitelistMethods is a paid mutator transaction binding the contract method 0x4f8249a2.
//
// Solidity: function removeWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorSession) RemoveWhitelistMethods(target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RemoveWhitelistMethods(&_BridgeExecutor.TransactOpts, target, methodIDs)
}

// RemoveWhitelistMethods is a paid mutator transaction binding the contract method 0x4f8249a2.
//
// Solidity: function removeWhitelistMethods(address target, bytes4[] methodIDs) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) RemoveWhitelistMethods(target common.Address, methodIDs [][4]byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RemoveWhitelistMethods(&_BridgeExecutor.TransactOpts, target, methodIDs)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) RemoveWhitelistTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "removeWhitelistTarget", target)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorSession) RemoveWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RemoveWhitelistTarget(&_BridgeExecutor.TransactOpts, target)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) RemoveWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RemoveWhitelistTarget(&_BridgeExecutor.TransactOpts, target)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecutor *BridgeExecutorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RenounceRole(&_BridgeExecutor.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RenounceRole(&_BridgeExecutor.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RevokeRole(&_BridgeExecutor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.RevokeRole(&_BridgeExecutor.TransactOpts, role, account)
}

// SetMaxReturnDataSize is a paid mutator transaction binding the contract method 0x8a7fdc75.
//
// Solidity: function setMaxReturnDataSize(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) SetMaxReturnDataSize(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "setMaxReturnDataSize", value)
}

// SetMaxReturnDataSize is a paid mutator transaction binding the contract method 0x8a7fdc75.
//
// Solidity: function setMaxReturnDataSize(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorSession) SetMaxReturnDataSize(value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetMaxReturnDataSize(&_BridgeExecutor.TransactOpts, value)
}

// SetMaxReturnDataSize is a paid mutator transaction binding the contract method 0x8a7fdc75.
//
// Solidity: function setMaxReturnDataSize(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) SetMaxReturnDataSize(value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetMaxReturnDataSize(&_BridgeExecutor.TransactOpts, value)
}

// SetMethodCheckEnabled is a paid mutator transaction binding the contract method 0x47239f71.
//
// Solidity: function setMethodCheckEnabled(address target, bool enabled) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) SetMethodCheckEnabled(opts *bind.TransactOpts, target common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "setMethodCheckEnabled", target, enabled)
}

// SetMethodCheckEnabled is a paid mutator transaction binding the contract method 0x47239f71.
//
// Solidity: function setMethodCheckEnabled(address target, bool enabled) returns()
func (_BridgeExecutor *BridgeExecutorSession) SetMethodCheckEnabled(target common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetMethodCheckEnabled(&_BridgeExecutor.TransactOpts, target, enabled)
}

// SetMethodCheckEnabled is a paid mutator transaction binding the contract method 0x47239f71.
//
// Solidity: function setMethodCheckEnabled(address target, bool enabled) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) SetMethodCheckEnabled(target common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetMethodCheckEnabled(&_BridgeExecutor.TransactOpts, target, enabled)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecutor *BridgeExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecutor *BridgeExecutorSession) Receive() (*types.Transaction, error) {
	return _BridgeExecutor.Contract.Receive(&_BridgeExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeExecutor.Contract.Receive(&_BridgeExecutor.TransactOpts)
}

// BridgeExecutorMaxReturnDataSizeSetIterator is returned from FilterMaxReturnDataSizeSet and is used to iterate over the raw logs and unpacked data for MaxReturnDataSizeSet events raised by the BridgeExecutor contract.
type BridgeExecutorMaxReturnDataSizeSetIterator struct {
	Event *BridgeExecutorMaxReturnDataSizeSet // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorMaxReturnDataSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorMaxReturnDataSizeSet)
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
		it.Event = new(BridgeExecutorMaxReturnDataSizeSet)
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
func (it *BridgeExecutorMaxReturnDataSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorMaxReturnDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorMaxReturnDataSizeSet represents a MaxReturnDataSizeSet event raised by the BridgeExecutor contract.
type BridgeExecutorMaxReturnDataSizeSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMaxReturnDataSizeSet is a free log retrieval operation binding the contract event 0x3d613bec426ed2a2fb9a9244f66939a8de8f29c84fa65bfabde181e0996b68f8.
//
// Solidity: event MaxReturnDataSizeSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterMaxReturnDataSizeSet(opts *bind.FilterOpts) (*BridgeExecutorMaxReturnDataSizeSetIterator, error) {

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "MaxReturnDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorMaxReturnDataSizeSetIterator{contract: _BridgeExecutor.contract, event: "MaxReturnDataSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxReturnDataSizeSet is a free log subscription operation binding the contract event 0x3d613bec426ed2a2fb9a9244f66939a8de8f29c84fa65bfabde181e0996b68f8.
//
// Solidity: event MaxReturnDataSizeSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchMaxReturnDataSizeSet(opts *bind.WatchOpts, sink chan<- *BridgeExecutorMaxReturnDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "MaxReturnDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorMaxReturnDataSizeSet)
				if err := _BridgeExecutor.contract.UnpackLog(event, "MaxReturnDataSizeSet", log); err != nil {
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

// ParseMaxReturnDataSizeSet is a log parse operation binding the contract event 0x3d613bec426ed2a2fb9a9244f66939a8de8f29c84fa65bfabde181e0996b68f8.
//
// Solidity: event MaxReturnDataSizeSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseMaxReturnDataSizeSet(log types.Log) (*BridgeExecutorMaxReturnDataSizeSet, error) {
	event := new(BridgeExecutorMaxReturnDataSizeSet)
	if err := _BridgeExecutor.contract.UnpackLog(event, "MaxReturnDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorMethodCheckEnabledIterator is returned from FilterMethodCheckEnabled and is used to iterate over the raw logs and unpacked data for MethodCheckEnabled events raised by the BridgeExecutor contract.
type BridgeExecutorMethodCheckEnabledIterator struct {
	Event *BridgeExecutorMethodCheckEnabled // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorMethodCheckEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorMethodCheckEnabled)
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
		it.Event = new(BridgeExecutorMethodCheckEnabled)
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
func (it *BridgeExecutorMethodCheckEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorMethodCheckEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorMethodCheckEnabled represents a MethodCheckEnabled event raised by the BridgeExecutor contract.
type BridgeExecutorMethodCheckEnabled struct {
	Target  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMethodCheckEnabled is a free log retrieval operation binding the contract event 0x5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e3.
//
// Solidity: event MethodCheckEnabled(address indexed target, bool enabled)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterMethodCheckEnabled(opts *bind.FilterOpts, target []common.Address) (*BridgeExecutorMethodCheckEnabledIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "MethodCheckEnabled", targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorMethodCheckEnabledIterator{contract: _BridgeExecutor.contract, event: "MethodCheckEnabled", logs: logs, sub: sub}, nil
}

// WatchMethodCheckEnabled is a free log subscription operation binding the contract event 0x5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e3.
//
// Solidity: event MethodCheckEnabled(address indexed target, bool enabled)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchMethodCheckEnabled(opts *bind.WatchOpts, sink chan<- *BridgeExecutorMethodCheckEnabled, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "MethodCheckEnabled", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorMethodCheckEnabled)
				if err := _BridgeExecutor.contract.UnpackLog(event, "MethodCheckEnabled", log); err != nil {
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

// ParseMethodCheckEnabled is a log parse operation binding the contract event 0x5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e3.
//
// Solidity: event MethodCheckEnabled(address indexed target, bool enabled)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseMethodCheckEnabled(log types.Log) (*BridgeExecutorMethodCheckEnabled, error) {
	event := new(BridgeExecutorMethodCheckEnabled)
	if err := _BridgeExecutor.contract.UnpackLog(event, "MethodCheckEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorMethodRemovedFromWhitelistIterator is returned from FilterMethodRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for MethodRemovedFromWhitelist events raised by the BridgeExecutor contract.
type BridgeExecutorMethodRemovedFromWhitelistIterator struct {
	Event *BridgeExecutorMethodRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorMethodRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorMethodRemovedFromWhitelist)
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
		it.Event = new(BridgeExecutorMethodRemovedFromWhitelist)
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
func (it *BridgeExecutorMethodRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorMethodRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorMethodRemovedFromWhitelist represents a MethodRemovedFromWhitelist event raised by the BridgeExecutor contract.
type BridgeExecutorMethodRemovedFromWhitelist struct {
	Target   common.Address
	MethodID [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMethodRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a.
//
// Solidity: event MethodRemovedFromWhitelist(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterMethodRemovedFromWhitelist(opts *bind.FilterOpts, target []common.Address, methodID [][4]byte) (*BridgeExecutorMethodRemovedFromWhitelistIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var methodIDRule []interface{}
	for _, methodIDItem := range methodID {
		methodIDRule = append(methodIDRule, methodIDItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "MethodRemovedFromWhitelist", targetRule, methodIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorMethodRemovedFromWhitelistIterator{contract: _BridgeExecutor.contract, event: "MethodRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchMethodRemovedFromWhitelist is a free log subscription operation binding the contract event 0x5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a.
//
// Solidity: event MethodRemovedFromWhitelist(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchMethodRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *BridgeExecutorMethodRemovedFromWhitelist, target []common.Address, methodID [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var methodIDRule []interface{}
	for _, methodIDItem := range methodID {
		methodIDRule = append(methodIDRule, methodIDItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "MethodRemovedFromWhitelist", targetRule, methodIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorMethodRemovedFromWhitelist)
				if err := _BridgeExecutor.contract.UnpackLog(event, "MethodRemovedFromWhitelist", log); err != nil {
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

// ParseMethodRemovedFromWhitelist is a log parse operation binding the contract event 0x5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a.
//
// Solidity: event MethodRemovedFromWhitelist(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseMethodRemovedFromWhitelist(log types.Log) (*BridgeExecutorMethodRemovedFromWhitelist, error) {
	event := new(BridgeExecutorMethodRemovedFromWhitelist)
	if err := _BridgeExecutor.contract.UnpackLog(event, "MethodRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorMethodWhitelistedIterator is returned from FilterMethodWhitelisted and is used to iterate over the raw logs and unpacked data for MethodWhitelisted events raised by the BridgeExecutor contract.
type BridgeExecutorMethodWhitelistedIterator struct {
	Event *BridgeExecutorMethodWhitelisted // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorMethodWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorMethodWhitelisted)
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
		it.Event = new(BridgeExecutorMethodWhitelisted)
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
func (it *BridgeExecutorMethodWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorMethodWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorMethodWhitelisted represents a MethodWhitelisted event raised by the BridgeExecutor contract.
type BridgeExecutorMethodWhitelisted struct {
	Target   common.Address
	MethodID [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMethodWhitelisted is a free log retrieval operation binding the contract event 0x17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd.
//
// Solidity: event MethodWhitelisted(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterMethodWhitelisted(opts *bind.FilterOpts, target []common.Address, methodID [][4]byte) (*BridgeExecutorMethodWhitelistedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var methodIDRule []interface{}
	for _, methodIDItem := range methodID {
		methodIDRule = append(methodIDRule, methodIDItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "MethodWhitelisted", targetRule, methodIDRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorMethodWhitelistedIterator{contract: _BridgeExecutor.contract, event: "MethodWhitelisted", logs: logs, sub: sub}, nil
}

// WatchMethodWhitelisted is a free log subscription operation binding the contract event 0x17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd.
//
// Solidity: event MethodWhitelisted(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchMethodWhitelisted(opts *bind.WatchOpts, sink chan<- *BridgeExecutorMethodWhitelisted, target []common.Address, methodID [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var methodIDRule []interface{}
	for _, methodIDItem := range methodID {
		methodIDRule = append(methodIDRule, methodIDItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "MethodWhitelisted", targetRule, methodIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorMethodWhitelisted)
				if err := _BridgeExecutor.contract.UnpackLog(event, "MethodWhitelisted", log); err != nil {
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

// ParseMethodWhitelisted is a log parse operation binding the contract event 0x17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd.
//
// Solidity: event MethodWhitelisted(address indexed target, bytes4 indexed methodID)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseMethodWhitelisted(log types.Log) (*BridgeExecutorMethodWhitelisted, error) {
	event := new(BridgeExecutorMethodWhitelisted)
	if err := _BridgeExecutor.contract.UnpackLog(event, "MethodWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeExecutor contract.
type BridgeExecutorRoleAdminChangedIterator struct {
	Event *BridgeExecutorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorRoleAdminChanged)
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
		it.Event = new(BridgeExecutorRoleAdminChanged)
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
func (it *BridgeExecutorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeExecutor contract.
type BridgeExecutorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeExecutorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorRoleAdminChangedIterator{contract: _BridgeExecutor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeExecutorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorRoleAdminChanged)
				if err := _BridgeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeExecutor *BridgeExecutorFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeExecutorRoleAdminChanged, error) {
	event := new(BridgeExecutorRoleAdminChanged)
	if err := _BridgeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeExecutor contract.
type BridgeExecutorRoleGrantedIterator struct {
	Event *BridgeExecutorRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorRoleGranted)
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
		it.Event = new(BridgeExecutorRoleGranted)
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
func (it *BridgeExecutorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorRoleGranted represents a RoleGranted event raised by the BridgeExecutor contract.
type BridgeExecutorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeExecutorRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorRoleGrantedIterator{contract: _BridgeExecutor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeExecutorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorRoleGranted)
				if err := _BridgeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeExecutor *BridgeExecutorFilterer) ParseRoleGranted(log types.Log) (*BridgeExecutorRoleGranted, error) {
	event := new(BridgeExecutorRoleGranted)
	if err := _BridgeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeExecutor contract.
type BridgeExecutorRoleRevokedIterator struct {
	Event *BridgeExecutorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorRoleRevoked)
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
		it.Event = new(BridgeExecutorRoleRevoked)
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
func (it *BridgeExecutorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorRoleRevoked represents a RoleRevoked event raised by the BridgeExecutor contract.
type BridgeExecutorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeExecutorRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorRoleRevokedIterator{contract: _BridgeExecutor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeExecutorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorRoleRevoked)
				if err := _BridgeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeExecutor *BridgeExecutorFilterer) ParseRoleRevoked(log types.Log) (*BridgeExecutorRoleRevoked, error) {
	event := new(BridgeExecutorRoleRevoked)
	if err := _BridgeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorTargetRemovedFromWhitelistIterator is returned from FilterTargetRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for TargetRemovedFromWhitelist events raised by the BridgeExecutor contract.
type BridgeExecutorTargetRemovedFromWhitelistIterator struct {
	Event *BridgeExecutorTargetRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorTargetRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorTargetRemovedFromWhitelist)
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
		it.Event = new(BridgeExecutorTargetRemovedFromWhitelist)
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
func (it *BridgeExecutorTargetRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorTargetRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorTargetRemovedFromWhitelist represents a TargetRemovedFromWhitelist event raised by the BridgeExecutor contract.
type BridgeExecutorTargetRemovedFromWhitelist struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d8312567.
//
// Solidity: event TargetRemovedFromWhitelist(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterTargetRemovedFromWhitelist(opts *bind.FilterOpts, target []common.Address) (*BridgeExecutorTargetRemovedFromWhitelistIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "TargetRemovedFromWhitelist", targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorTargetRemovedFromWhitelistIterator{contract: _BridgeExecutor.contract, event: "TargetRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchTargetRemovedFromWhitelist is a free log subscription operation binding the contract event 0x09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d8312567.
//
// Solidity: event TargetRemovedFromWhitelist(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchTargetRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *BridgeExecutorTargetRemovedFromWhitelist, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "TargetRemovedFromWhitelist", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorTargetRemovedFromWhitelist)
				if err := _BridgeExecutor.contract.UnpackLog(event, "TargetRemovedFromWhitelist", log); err != nil {
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

// ParseTargetRemovedFromWhitelist is a log parse operation binding the contract event 0x09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d8312567.
//
// Solidity: event TargetRemovedFromWhitelist(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseTargetRemovedFromWhitelist(log types.Log) (*BridgeExecutorTargetRemovedFromWhitelist, error) {
	event := new(BridgeExecutorTargetRemovedFromWhitelist)
	if err := _BridgeExecutor.contract.UnpackLog(event, "TargetRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecutorTargetWhitelistedIterator is returned from FilterTargetWhitelisted and is used to iterate over the raw logs and unpacked data for TargetWhitelisted events raised by the BridgeExecutor contract.
type BridgeExecutorTargetWhitelistedIterator struct {
	Event *BridgeExecutorTargetWhitelisted // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorTargetWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorTargetWhitelisted)
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
		it.Event = new(BridgeExecutorTargetWhitelisted)
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
func (it *BridgeExecutorTargetWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorTargetWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorTargetWhitelisted represents a TargetWhitelisted event raised by the BridgeExecutor contract.
type BridgeExecutorTargetWhitelisted struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetWhitelisted is a free log retrieval operation binding the contract event 0x3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b082.
//
// Solidity: event TargetWhitelisted(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterTargetWhitelisted(opts *bind.FilterOpts, target []common.Address) (*BridgeExecutorTargetWhitelistedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "TargetWhitelisted", targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorTargetWhitelistedIterator{contract: _BridgeExecutor.contract, event: "TargetWhitelisted", logs: logs, sub: sub}, nil
}

// WatchTargetWhitelisted is a free log subscription operation binding the contract event 0x3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b082.
//
// Solidity: event TargetWhitelisted(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchTargetWhitelisted(opts *bind.WatchOpts, sink chan<- *BridgeExecutorTargetWhitelisted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "TargetWhitelisted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorTargetWhitelisted)
				if err := _BridgeExecutor.contract.UnpackLog(event, "TargetWhitelisted", log); err != nil {
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

// ParseTargetWhitelisted is a log parse operation binding the contract event 0x3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b082.
//
// Solidity: event TargetWhitelisted(address indexed target)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseTargetWhitelisted(log types.Log) (*BridgeExecutorTargetWhitelisted, error) {
	event := new(BridgeExecutorTargetWhitelisted)
	if err := _BridgeExecutor.contract.UnpackLog(event, "TargetWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
