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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"addWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"executeExtraCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isMethodCheckEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"isWhitelistedMethod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isWhitelistedTarget\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"removeWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMethodCheckEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MethodCheckEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEFailedToReturnNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEMethodNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BENotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
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
		"230f4cda": "postCallGasReserve()",
		"d04323c5": "recoverToken(address,uint256,address)",
		"4f8249a2": "removeWhitelistMethods(address,bytes4[])",
		"11735ea5": "removeWhitelistTarget(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"47239f71": "setMethodCheckEnabled(address,bool)",
		"4bc51672": "setPostCallGasReserve(uint256)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x6080346100b357601f61171138819003918201601f19168301916001600160401b038311848410176100b75780849260409485528339810103126100b357610052602061004b836100cb565b92016100cb565b62030d406003556001600160a01b0382161515806100a1575b156100925761007c610082926100df565b50610155565b506040516114c890816101e98239f35b63484b059f60e11b5f5260045ffd5b506001600160a01b038116151561006b565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100b357565b6001600160a01b0381165f9081525f5160206116f15f395f51905f52602052604090205460ff16610150576001600160a01b03165f8181525f5160206116f15f395f51905f5260205260408120805460ff191660011790553391905f5160206116b15f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f5160206116d15f395f51905f52602052604090205460ff16610150576001600160a01b03165f8181525f5160206116d15f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f5160206116b15f395f51905f529080a460019056fe6080604052600436101561001a575b3615610018575f80fd5b005b5f3560e01c806301ffc9a71461014957806311735ea5146101445780631eeed5131461013f578063230f4cda1461013a578063248a9ca3146101355780632f2ff15d1461013057806336568abe1461012b57806347239f71146101265780634bc51672146101215780634f8249a21461011c57806361c8c47914610117578063907caa001461011257806391d148541461010d578063a217fddf14610108578063a8582dc214610103578063ab779650146100fe578063b8e25592146100f9578063d04323c5146100f45763d547741f0361000e5761095e565b610847565b6107fc565b6107b9565b61070b565b6106f1565b6106b0565b610670565b6105d2565b61051c565b61043c565b6103a9565b610361565b610328565b610302565b6102e5565b610258565b6101cb565b610164565b6001600160e01b031981160361016057565b5f80fd5b346101605760203660031901126101605760206004356101838161014e565b63ffffffff60e01b16637965db0b60e01b81149081156101a9575b506040519015158152f35b6301ffc9a760e01b1490505f61019e565b6001600160a01b0381160361016057565b34610160576020366003190112610160576004356101e8816101ba565b6101f0610fe3565b6001600160a01b03165f8181526001602052604090205460ff161561024957805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b60c036600319011261016057604435602435600435610276836101ba565b606435610282816101ba565b60a4359360843591906001600160401b0386116101605736602387011215610160576004860135946001600160401b0386116101605736602487890101116101605760246102d19701946109a7565b604080519215158352602083019190915290f35b34610160575f366003190112610160576020600354604051908152f35b34610160576020366003190112610160576020610320600435610f53565b604051908152f35b346101605760403660031901126101605761001860243560043561034b826101ba565b61035c61035782610f53565b611036565b6112aa565b3461016057604036600319011261016057600435602435610381816101ba565b336001600160a01b0382160361039a576100189161131a565b63334bd91960e11b5f5260045ffd5b34610160576040366003190112610160576004356103c6816101ba565b602435908115158092036101605760207f5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e391610400610fe3565b6001600160a01b031692610415841515610f64565b835f526001825260405f20805461ff008360081b169061ff001916179055604051908152a2005b3461016057602036600319011261016057600435610458610fe3565b61c350811015806104b1575b156104a25760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91600354908060035582519182526020820152a1005b6315e4ec0160e01b5f5260045ffd5b50620f4240811115610464565b6040600319820112610160576004356104d6816101ba565b916024356001600160401b0381116101605782602382011215610160576004810135926001600160401b0384116101605760248460051b83010111610160576024019190565b346101605761052a366104be565b9091610534610fe3565b6001600160a01b0316610548811515610f64565b5f5b82811061055357005b600190825f52600260205261058060405f20610570838789610f7a565b359061057b8261014e565b610a6b565b805460ff19169055610593818587610f7a565b3561059d8161014e565b6001600160e01b031916837f5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a5f80a30161054a565b34610160576020366003190112610160576004356105ef816101ba565b6105f7610fe3565b6001600160a01b031661060b811515610f64565b805f52600160205260ff60405f20541661066157805f52600160205261063b60405f20600160ff19825416179055565b7f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b346101605760203660031901126101605760043561068d816101ba565b60018060a01b03165f526001602052602060ff60405f2054166040519015158152f35b3461016057604036600319011261016057602060ff6106e56024356004356106d7826101ba565b5f525f845260405f20610992565b54166040519015158152f35b34610160575f3660031901126101605760206040515f8152f35b3461016057610719366104be565b9091610723610fe3565b6001600160a01b0316610737811515610f64565b5f5b82811061074257005b600190825f52600260205261076f61076260405f2061057084888a610f7a565b805460ff19166001179055565b61077a818587610f7a565b356107848161014e565b6001600160e01b031916837f17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd5f80a301610739565b34610160576020366003190112610160576004356107d6816101ba565b60018060a01b03165f526001602052602060ff60405f205460081c166040519015158152f35b3461016057604036600319011261016057602060ff6106e5600435610820816101ba565b6024359061082d8261014e565b6001600160a01b03165f9081526002855260409020610a6b565b3461016057606036600319011261016057600435610864816101ba565b60243590604435610874816101ba565b61087c610fe3565b6001600160a01b038116918215158061094c575b1561093d576001600160a01b031690600182036108c25750505f80806100189481945af16108bc610f9e565b50610fcd565b602092506108ec6108fa5f93956040519283918783019563a9059cbb60e01b875260248401611255565b03601f198101835282610ad8565b519082855af115610932575f513d6109295750803b155b61091757005b635274afe760e01b5f5260045260245ffd5b60011415610911565b6040513d5f823e3d90fd5b63b56bf2e760e01b5f5260045ffd5b506001600160a01b0381161515610890565b3461016057604036600319011261016057610018602435600435610981826101ba565b61098d61035782610f53565b61131a565b9060018060a01b03165f5260205260405f2090565b5f5160206114535f395f51905f525f90815260205295949392919060ff6109ee337fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069610992565b541615610a47575f5160206114735f395f51905f525c610a3857610a229660015f5160206114735f395f51905f525d610ba8565b91905f5f5160206114735f395f51905f525d9190565b633ee5aeb560e01b5f5260045ffd5b63e2517d3f60e01b5f52336004525f5160206114535f395f51905f5260245260445ffd5b9063ffffffff60e01b165f5260205260405f2090565b634e487b7160e01b5f52601160045260245ffd5b91908203918211610aa257565b610a81565b909291928360141161016057831161016057601401916013190190565b634e487b7160e01b5f52604160045260245ffd5b601f909101601f19168101906001600160401b03821190821017610afb57604052565b610ac4565b6001600160401b038111610afb57601f01601f191660200190565b929192610b2782610b00565b91610b356040519384610ad8565b829481845281830111610160578281602093845f960137010152565b906175308201809211610aa257565b9260209260a0959263ffffffff60e01b1685521515838501526040840152608060608401528051918291826080860152018484015e5f828201840152601f01601f1916010190565b91909293969594968015610f455760188610610f2357873560601c5f8181526001602081905260409091206001600160a01b0385169091149960148101359792949291610bfd610bf9825460ff1690565b1590565b610e97575460081c60ff1680610ef2575b610ede578a15610e5657610c223447610a95565b98843403610e3957610c6191610c5a610c538e935b600354805a1015610e2957505f945b15610e2157808994610aa7565b3691610b1b565b9187611213565b9390988b15610e11575b8b15610df35747905b8c15610dd95780821115610dd157610c8b91610a95565b818110610dc257505f5b80821115610db957610ca691610a95565b915b829a83610ce5575b505050505f5160206114335f395f51905f5291610cde87928a60405194859460018060a01b03169985610b60565b0390a49190565b939a9315610d6a5750610d17915082610cff600354610b51565b805a1115610d6057610d11905a610a95565b9161109a565b15610d3b5750845f5160206114335f395f51905f5291610cde5f995b928294610cb0565b965f5160206114335f395f51905f5291610cde8792610d5a8b3361106c565b50610d33565b506175309161109a565b9082610d769183611270565b15610d9857505050845f5160206114335f395f51905f5291610cde5f99610d33565b879299610d5a5f5160206114335f395f51905f529593610cde9333906111e3565b50505f91610ca8565b610dcc9082610a95565b610c95565b50505f610c8b565b9080821115610deb57610dcc91610a95565b50505f610c95565b610dfd3085611136565b9015610e0a575b90610c74565b505f610e04565b610e1b86856111b1565b50610c6b565b805f94610aa7565b610e33905a610a95565b94610c46565b505050505050505050509050610e4f343361106c565b505f905f90565b34610ece575b610e663084611136565b90158015610ec5575b610eb05798610e82610bf98688876111e3565b610e9757610c6191610c5a610c538e93610c37565b50509650935093505050610eac9394506110c4565b9091565b97505096509350505050610eac9394506110c4565b50848110610e6f565b610ed8343361106c565b50610e5c565b509650935093505050610eac9394506110c4565b50610f1e610bf9610f178a61057b8960018060a01b03165f52600260205260405f2090565b5460ff1690565b610c0e565b939450915050610eac93945034156110c457610f3f343361106c565b506110c4565b50505050505090505f905f90565b5f525f602052600160405f20015490565b15610f6b57565b63484b059f60e11b5f5260045ffd5b9190811015610f8a5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b3d15610fc8573d90610faf82610b00565b91610fbd6040519384610ad8565b82523d5f602084013e565b606090565b15610fd457565b6317e90b5960e21b5f5260045ffd5b5f80805260205260ff611016337fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5610992565b54161561101f57565b63e2517d3f60e01b5f52336004525f60245260445ffd5b805f525f60205260ff61104c3360405f20610992565b5416156110565750565b63e2517d3f60e01b5f523360045260245260445ffd5b8115611093575f918291829182916001600160a01b0316617530f161108f610f9e565b5090565b5050600190565b9181156110bc575f928392839283926001600160a01b031690f161108f610f9e565b505050600190565b909291831561112c57836001600160a01b03831660010361110a576110ef9250610cff600354610b51565b611103576110fd823361106c565b505f9190565b5f91508190565b6111149183611270565b61112457826110fd9133906111e3565b505f91508190565b505090505f905f90565b6040516370a0823160e01b602082019081526001600160a01b0390931660248083019190915281525f9283929161116e604482610ad8565b51916001600160a01b03165afa611183610f9e565b901580156111a6575b610e4f576020818051810103126101605760200151600191565b50602081511061118c565b6111bc5f83836113c4565b611093576111ca8282611388565b156111dd576111da915f916113c4565b90565b50505f90565b91906111f08282856113c4565b6110bc576111fe8184611388565b1561120c576111da926113c4565b5050505f90565b925f929183926040519560208451940192f1913d610400811161124c575b806020918452805f8386013e601f01601f1916830101604052565b50610400611231565b6001600160a01b039091168152602081019190915260400190565b6112a56111da939261129760405194859263a9059cbb60e01b602085015260248401611255565b03601f198101845283610ad8565b6113eb565b805f525f60205260ff6112c08360405f20610992565b54166111dd57805f525f6020526112da8260405f20610992565b805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b805f525f60205260ff6113308360405f20610992565b5416156111dd57805f525f60205261134b8260405f20610992565b805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b60405163095ea7b360e01b60208201526001600160a01b0390921660248301525f60448084019190915282526111da91906112a5606483610ad8565b6112a56111da939261129760405194859263095ea7b360e01b602085015260248401611255565b905f602091828151910182855af1903d5f51908361140a575b50505090565b9192509061142857506001600160a01b03163b15155b5f8080611404565b600191501461142056fec4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e639b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212200a8026bdde3db9430e06f3c2741587f63a8395d1e7e44607c38ecf8f2b458be164736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0ddae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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
func DeployBridgeExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, bridge common.Address) (common.Address, *types.Transaction, *BridgeExecutor, error) {
	parsed, err := BridgeExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeExecutorBin), backend, owner, bridge)
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

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorCaller) PostCallGasReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeExecutor.contract.Call(opts, &out, "postCallGasReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorSession) PostCallGasReserve() (*big.Int, error) {
	return _BridgeExecutor.Contract.PostCallGasReserve(&_BridgeExecutor.CallOpts)
}

// PostCallGasReserve is a free data retrieval call binding the contract method 0x230f4cda.
//
// Solidity: function postCallGasReserve() view returns(uint256)
func (_BridgeExecutor *BridgeExecutorCallerSession) PostCallGasReserve() (*big.Int, error) {
	return _BridgeExecutor.Contract.PostCallGasReserve(&_BridgeExecutor.CallOpts)
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
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorTransactor) ExecuteExtraCall(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "executeExtraCall", fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorTransactorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, fromChainID, index, toToken, to, value, extraData)
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

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorTransactor) SetPostCallGasReserve(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "setPostCallGasReserve", value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetPostCallGasReserve(&_BridgeExecutor.TransactOpts, value)
}

// SetPostCallGasReserve is a paid mutator transaction binding the contract method 0x4bc51672.
//
// Solidity: function setPostCallGasReserve(uint256 value) returns()
func (_BridgeExecutor *BridgeExecutorTransactorSession) SetPostCallGasReserve(value *big.Int) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.SetPostCallGasReserve(&_BridgeExecutor.TransactOpts, value)
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

// BridgeExecutorExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BridgeExecutor contract.
type BridgeExecutorExtraCallExecutedIterator struct {
	Event *BridgeExecutorExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorExtraCallExecuted)
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
		it.Event = new(BridgeExecutorExtraCallExecuted)
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
func (it *BridgeExecutorExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorExtraCallExecuted represents a ExtraCallExecuted event raised by the BridgeExecutor contract.
type BridgeExecutorExtraCallExecuted struct {
	FromChainID    *big.Int
	Index          *big.Int
	TargetContract common.Address
	MethodID       [4]byte
	Success        bool
	Remaining      *big.Int
	ReturnData     []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 remaining, bytes returnData)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (*BridgeExecutorExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorExtraCallExecutedIterator{contract: _BridgeExecutor.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 remaining, bytes returnData)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BridgeExecutorExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, targetContract []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorExtraCallExecuted)
				if err := _BridgeExecutor.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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

// ParseExtraCallExecuted is a log parse operation binding the contract event 0xc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e40.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed targetContract, bytes4 methodID, bool success, uint256 remaining, bytes returnData)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseExtraCallExecuted(log types.Log) (*BridgeExecutorExtraCallExecuted, error) {
	event := new(BridgeExecutorExtraCallExecuted)
	if err := _BridgeExecutor.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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

// BridgeExecutorPostCallGasReserveSetIterator is returned from FilterPostCallGasReserveSet and is used to iterate over the raw logs and unpacked data for PostCallGasReserveSet events raised by the BridgeExecutor contract.
type BridgeExecutorPostCallGasReserveSetIterator struct {
	Event *BridgeExecutorPostCallGasReserveSet // Event containing the contract specifics and raw log

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
func (it *BridgeExecutorPostCallGasReserveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecutorPostCallGasReserveSet)
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
		it.Event = new(BridgeExecutorPostCallGasReserveSet)
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
func (it *BridgeExecutorPostCallGasReserveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecutorPostCallGasReserveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecutorPostCallGasReserveSet represents a PostCallGasReserveSet event raised by the BridgeExecutor contract.
type BridgeExecutorPostCallGasReserveSet struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostCallGasReserveSet is a free log retrieval operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterPostCallGasReserveSet(opts *bind.FilterOpts) (*BridgeExecutorPostCallGasReserveSetIterator, error) {

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorPostCallGasReserveSetIterator{contract: _BridgeExecutor.contract, event: "PostCallGasReserveSet", logs: logs, sub: sub}, nil
}

// WatchPostCallGasReserveSet is a free log subscription operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchPostCallGasReserveSet(opts *bind.WatchOpts, sink chan<- *BridgeExecutorPostCallGasReserveSet) (event.Subscription, error) {

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "PostCallGasReserveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecutorPostCallGasReserveSet)
				if err := _BridgeExecutor.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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

// ParsePostCallGasReserveSet is a log parse operation binding the contract event 0x18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e.
//
// Solidity: event PostCallGasReserveSet(uint256 oldValue, uint256 newValue)
func (_BridgeExecutor *BridgeExecutorFilterer) ParsePostCallGasReserveSet(log types.Log) (*BridgeExecutorPostCallGasReserveSet, error) {
	event := new(BridgeExecutorPostCallGasReserveSet)
	if err := _BridgeExecutor.contract.UnpackLog(event, "PostCallGasReserveSet", log); err != nil {
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
