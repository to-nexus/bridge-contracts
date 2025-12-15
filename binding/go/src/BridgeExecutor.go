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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"addWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"executeExtraCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isMethodCheckEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"isWhitelistedMethod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isWhitelistedTarget\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCallGasReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"methodIDs\",\"type\":\"bytes4[]\"}],\"name\":\"removeWhitelistMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMethodCheckEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setPostCallGasReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MethodCheckEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"}],\"name\":\"MethodWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"PostCallGasReserveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEFailedToReturnNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidGasReserve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEMethodNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BENotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
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
	Bin: "0x6080346100b357601f6114a638819003918201601f19168301916001600160401b038311848410176100b75780849260409485528339810103126100b357610052602061004b836100cb565b92016100cb565b62030d406003556001600160a01b0382161515806100a1575b156100925761007c610082926100df565b50610155565b5060405161125d90816101e98239f35b63484b059f60e11b5f5260045ffd5b506001600160a01b038116151561006b565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100b357565b6001600160a01b0381165f9081525f5160206114865f395f51905f52602052604090205460ff16610150576001600160a01b03165f8181525f5160206114865f395f51905f5260205260408120805460ff191660011790553391905f5160206114465f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f5160206114665f395f51905f52602052604090205460ff16610150576001600160a01b03165f8181525f5160206114665f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f5160206114465f395f51905f529080a460019056fe608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461090e5750806311735ea5146108855780631eeed5131461074a578063230f4cda1461072d578063248a9ca3146107075780632f2ff15d146106d657806336568abe1461069257806347239f71146106085780634bc51672146105865780634f8249a2146104dc57806361c8c4791461044c578063907caa001461040f57806391d14854146103c7578063a217fddf146103ad578063a8582dc2146102ed578063ab779650146102ad578063b8e2559214610252578063d04323c5146101335763d547741f146100f9575f61000f565b3461012f57604036600319011261012f5761001a60043561011861098d565b9061012a61012582610dce565b610e67565b6110c6565b5f80fd5b3461012f57606036600319011261012f5761014c610961565b602435906044356001600160a01b0381169182820361012f5761016d610e18565b82151580610240575b15610231576001600160a01b031690600182036101b65750505f80809381935af161019f610a60565b50156101a757005b6317e90b5960e21b5f5260045ffd5b602092506101e06101ee5f93956040519283918783019563a9059cbb60e01b87526024840161115f565b03601f198101835282610a29565b519082855af115610226575f513d61021d5750803b155b61020b57005b635274afe760e01b5f5260045260245ffd5b60011415610205565b6040513d5f823e3d90fd5b63b56bf2e760e01b5f5260045ffd5b506001600160a01b0381161515610176565b3461012f57604036600319011261012f5761026b610961565b6024359063ffffffff60e01b821680920361012f5760018060a01b03165f52600260205260405f20905f52602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f576001600160a01b036102ce610961565b165f526001602052602060ff60405f205460081c166040519015158152f35b3461012f576102fb366109a3565b91610304610e18565b6001600160a01b0316801561039e575f5b83811061031e57005b600190825f52600260205260405f2063ffffffff60e01b610348610343848989610ddf565b610e03565b165f5260205260405f208260ff1982541617905563ffffffff60e01b610372610343838888610ddf565b16837f17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd5f80a301610315565b63484b059f60e11b5f5260045ffd5b3461012f575f36600319011261012f5760206040515f8152f35b3461012f57604036600319011261012f576103e061098d565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f576001600160a01b03610430610961565b165f526001602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f57610465610961565b61046d610e18565b6001600160a01b0316801561039e57805f52600160205260ff60405f2054166104cd57805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b3461012f576104ea366109a3565b916104f3610e18565b6001600160a01b0316801561039e575f5b83811061050d57005b600190825f52600260205260405f2063ffffffff60e01b610532610343848989610ddf565b165f5260205260405f2060ff19815416905563ffffffff60e01b61055a610343838888610ddf565b16837f5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a5f80a301610504565b3461012f57602036600319011261012f576004356105a2610e18565b61c350811015806105fb575b156105ec5760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91600354908060035582519182526020820152a1005b6315e4ec0160e01b5f5260045ffd5b50620f42408111156105ae565b3461012f57604036600319011261012f57610621610961565b6024359081151580920361012f57610637610e18565b6001600160a01b031690811561039e5760207f5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e391835f526001825260405f20805461ff008360081b169061ff001916179055604051908152a2005b3461012f57604036600319011261012f576106ab61098d565b336001600160a01b038216036106c75761001a906004356110c6565b63334bd91960e11b5f5260045ffd5b3461012f57604036600319011261012f5761001a6004356106f561098d565b9061070261012582610dce565b611044565b3461012f57602036600319011261012f576020610725600435610dce565b604051908152f35b3461012f575f36600319011261012f576020600354604051908152f35b60c036600319011261012f576044356001600160a01b038116810361012f57610771610977565b5060a4356001600160401b03811161012f573660238201121561012f5760048101356001600160401b03811161012f57366024828401011161012f57335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff161561084e575f5160206112085f395f51905f525c61083f5760409260246108219360015f5160206112085f395f51905f525d019060843590602435600435610a9e565b5f5f5160206112085f395f51905f525d825191151582526020820152f35b633ee5aeb560e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b3461012f57602036600319011261012f5761089e610961565b6108a6610e18565b6001600160a01b03165f8181526001602052604090205460ff16156108ff57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b3461012f57602036600319011261012f576004359063ffffffff60e01b821680920361012f57602091637965db0b60e01b8114908115610950575b5015158152f35b6301ffc9a760e01b14905083610949565b600435906001600160a01b038216820361012f57565b606435906001600160a01b038216820361012f57565b602435906001600160a01b038216820361012f57565b604060031982011261012f576004356001600160a01b038116810361012f57916024356001600160401b03811161012f578260238201121561012f576004810135926001600160401b03841161012f5760248460051b8301011161012f576024019190565b91908203918211610a1557565b634e487b7160e01b5f52601160045260245ffd5b601f909101601f19168101906001600160401b03821190821017610a4c57604052565b634e487b7160e01b5f52604160045260245ffd5b3d15610a99573d906001600160401b038211610a4c5760405191610a8e601f8201601f191660200184610a29565b82523d5f602084013e565b606090565b9091929395948415610dc15760188110610daf57863560601c5f81815260016020819052604090912080546001600160a01b0388169092149992968793926014840135929060ff1615610d95575460081c60ff1680610d6b575b610d48578a15610cef57610b0c3447610a08565b98803403610cc7575b600354805a1015610cb757505f935b8c15610cb157815b8760141161012f575f60a0968f9382938a7fc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e409c6040519360146013198301910185378301856013198201528360131991030193f19a610b89610a60565b968215610ca0575b508115610c8857475b80821015610c79575f5b92610c4d575b505080821115610c4357610bbd91610a08565b809b5b81610c18575b505050602060405193849263ffffffff60e01b168352891515828401528b6040840152608060608401528051918291826080860152018484015e5f828201840152601f01601f19168101030190a49190565b15610c3257610c28915033610e9f565b505b5f8a81610bc6565b610c3d913390610feb565b50610c2a565b50505f809b610bc0565b909150818110610c7157610c6a91610c6491610a08565b82610a08565b5f80610baa565b505080610c6a565b610c838183610a08565b610ba4565b610c923085610ecb565b905f03610b9a57505f610b9a565b610caa9085610f4d565b505f610b91565b5f610b2c565b610cc1905a610a08565b93610b24565b50505050505050505050905034610cdf575b5f905f90565b610ce93433610e9f565b50610cd9565b34610d38575b610cff3082610ecb565b90158015610d2f575b610d1d5798610d18818684610feb565b610b15575b5050505050505050505090505f905f90565b50898110610d08565b610d423433610e9f565b50610cf5565b5050505050505050509080610d62575b610cdf575f905f90565b50341515610d58565b50875f52600260205260405f2063ffffffff60e01b83165f5260205260ff60405f20541615610af8565b505050505050505050509080610d6257610cdf575f905f90565b5050505050905034610cdf575f905f90565b505050505090505f905f90565b5f525f602052600160405f20015490565b9190811015610def5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160e01b03198116810361012f5790565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615610e5057565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff1615610e895750565b63e2517d3f60e01b5f523360045260245260445ffd5b8115610ec4575f918291829182916001600160a01b03165af1610ec0610a60565b5090565b5050600190565b6040516370a0823160e01b602082019081526001600160a01b0390931660248083019190915281525f92839291610f03604482610a29565b51916001600160a01b03165afa610f18610a60565b90158015610f42575b610f3b5760208180518101031261012f5760200151600191565b505f905f90565b506020815110610f21565b90610f7960405163095ea7b360e01b6020820152610f73816101e05f866024840161115f565b836111a1565b610ec45760405163095ea7b360e01b6020820152610fa290610f73816101e08560248301611146565b15610fe557610fe291610fdd5f92610fcf60405194859263095ea7b360e01b60208501526024840161115f565b03601f198101845283610a29565b6111a1565b90565b50505f90565b9190610ff882828561117a565b61103c5760405163095ea7b360e01b602082015261102790611021816101e08560248301611146565b846111a1565b1561103557610fe29261117a565b5050505f90565b505050600190565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610fe5575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610fe5575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001600160a01b0390911681525f602082015260400190565b6001600160a01b039091168152602081019190915260400190565b610fdd610fe29392610fcf60405194859263095ea7b360e01b60208501526024840161115f565b81516001600160a01b03909116915f91829160200182855af1906111c3610a60565b9115610fe55781519081156111fe575060208110156111e25750505f90565b816020918101031261012f5760200151801515810361012f5790565b9150503b15159056fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212200a362b979d1e198ea3254d72e191b7a9b76906669b9796b1884a0e076bcc20dc64736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0ddae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// BridgeExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeExecutorMetaData.ABI instead.
var BridgeExecutorABI = BridgeExecutorMetaData.ABI

// BridgeExecutorBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeExecutorBinRuntime = "608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461090e5750806311735ea5146108855780631eeed5131461074a578063230f4cda1461072d578063248a9ca3146107075780632f2ff15d146106d657806336568abe1461069257806347239f71146106085780634bc51672146105865780634f8249a2146104dc57806361c8c4791461044c578063907caa001461040f57806391d14854146103c7578063a217fddf146103ad578063a8582dc2146102ed578063ab779650146102ad578063b8e2559214610252578063d04323c5146101335763d547741f146100f9575f61000f565b3461012f57604036600319011261012f5761001a60043561011861098d565b9061012a61012582610dce565b610e67565b6110c6565b5f80fd5b3461012f57606036600319011261012f5761014c610961565b602435906044356001600160a01b0381169182820361012f5761016d610e18565b82151580610240575b15610231576001600160a01b031690600182036101b65750505f80809381935af161019f610a60565b50156101a757005b6317e90b5960e21b5f5260045ffd5b602092506101e06101ee5f93956040519283918783019563a9059cbb60e01b87526024840161115f565b03601f198101835282610a29565b519082855af115610226575f513d61021d5750803b155b61020b57005b635274afe760e01b5f5260045260245ffd5b60011415610205565b6040513d5f823e3d90fd5b63b56bf2e760e01b5f5260045ffd5b506001600160a01b0381161515610176565b3461012f57604036600319011261012f5761026b610961565b6024359063ffffffff60e01b821680920361012f5760018060a01b03165f52600260205260405f20905f52602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f576001600160a01b036102ce610961565b165f526001602052602060ff60405f205460081c166040519015158152f35b3461012f576102fb366109a3565b91610304610e18565b6001600160a01b0316801561039e575f5b83811061031e57005b600190825f52600260205260405f2063ffffffff60e01b610348610343848989610ddf565b610e03565b165f5260205260405f208260ff1982541617905563ffffffff60e01b610372610343838888610ddf565b16837f17cce9eb3d9badc6086fac9c927183f6e8bd32f8e122d6bc9a8ba5f5919395dd5f80a301610315565b63484b059f60e11b5f5260045ffd5b3461012f575f36600319011261012f5760206040515f8152f35b3461012f57604036600319011261012f576103e061098d565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f576001600160a01b03610430610961565b165f526001602052602060ff60405f2054166040519015158152f35b3461012f57602036600319011261012f57610465610961565b61046d610e18565b6001600160a01b0316801561039e57805f52600160205260ff60405f2054166104cd57805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b3461012f576104ea366109a3565b916104f3610e18565b6001600160a01b0316801561039e575f5b83811061050d57005b600190825f52600260205260405f2063ffffffff60e01b610532610343848989610ddf565b165f5260205260405f2060ff19815416905563ffffffff60e01b61055a610343838888610ddf565b16837f5a0866eb654c49b51bd173c271b0025fdfebd26fd53182c118d53b4fd3f42d2a5f80a301610504565b3461012f57602036600319011261012f576004356105a2610e18565b61c350811015806105fb575b156105ec5760407f18c087ee0200bbe9d388bef52ea98caa28702660805c632add5ffd538a66356e91600354908060035582519182526020820152a1005b6315e4ec0160e01b5f5260045ffd5b50620f42408111156105ae565b3461012f57604036600319011261012f57610621610961565b6024359081151580920361012f57610637610e18565b6001600160a01b031690811561039e5760207f5579fe57ff4f4b388db4f445bcef31a9cf6a04181411c7e592279df0854b04e391835f526001825260405f20805461ff008360081b169061ff001916179055604051908152a2005b3461012f57604036600319011261012f576106ab61098d565b336001600160a01b038216036106c75761001a906004356110c6565b63334bd91960e11b5f5260045ffd5b3461012f57604036600319011261012f5761001a6004356106f561098d565b9061070261012582610dce565b611044565b3461012f57602036600319011261012f576020610725600435610dce565b604051908152f35b3461012f575f36600319011261012f576020600354604051908152f35b60c036600319011261012f576044356001600160a01b038116810361012f57610771610977565b5060a4356001600160401b03811161012f573660238201121561012f5760048101356001600160401b03811161012f57366024828401011161012f57335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff161561084e575f5160206112085f395f51905f525c61083f5760409260246108219360015f5160206112085f395f51905f525d019060843590602435600435610a9e565b5f5f5160206112085f395f51905f525d825191151582526020820152f35b633ee5aeb560e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b3461012f57602036600319011261012f5761089e610961565b6108a6610e18565b6001600160a01b03165f8181526001602052604090205460ff16156108ff57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b3461012f57602036600319011261012f576004359063ffffffff60e01b821680920361012f57602091637965db0b60e01b8114908115610950575b5015158152f35b6301ffc9a760e01b14905083610949565b600435906001600160a01b038216820361012f57565b606435906001600160a01b038216820361012f57565b602435906001600160a01b038216820361012f57565b604060031982011261012f576004356001600160a01b038116810361012f57916024356001600160401b03811161012f578260238201121561012f576004810135926001600160401b03841161012f5760248460051b8301011161012f576024019190565b91908203918211610a1557565b634e487b7160e01b5f52601160045260245ffd5b601f909101601f19168101906001600160401b03821190821017610a4c57604052565b634e487b7160e01b5f52604160045260245ffd5b3d15610a99573d906001600160401b038211610a4c5760405191610a8e601f8201601f191660200184610a29565b82523d5f602084013e565b606090565b9091929395948415610dc15760188110610daf57863560601c5f81815260016020819052604090912080546001600160a01b0388169092149992968793926014840135929060ff1615610d95575460081c60ff1680610d6b575b610d48578a15610cef57610b0c3447610a08565b98803403610cc7575b600354805a1015610cb757505f935b8c15610cb157815b8760141161012f575f60a0968f9382938a7fc4ab19c1c99c1dcca2d85419198714f88b7232b2ce23cc52e3ff60c7a1ef8e409c6040519360146013198301910185378301856013198201528360131991030193f19a610b89610a60565b968215610ca0575b508115610c8857475b80821015610c79575f5b92610c4d575b505080821115610c4357610bbd91610a08565b809b5b81610c18575b505050602060405193849263ffffffff60e01b168352891515828401528b6040840152608060608401528051918291826080860152018484015e5f828201840152601f01601f19168101030190a49190565b15610c3257610c28915033610e9f565b505b5f8a81610bc6565b610c3d913390610feb565b50610c2a565b50505f809b610bc0565b909150818110610c7157610c6a91610c6491610a08565b82610a08565b5f80610baa565b505080610c6a565b610c838183610a08565b610ba4565b610c923085610ecb565b905f03610b9a57505f610b9a565b610caa9085610f4d565b505f610b91565b5f610b2c565b610cc1905a610a08565b93610b24565b50505050505050505050905034610cdf575b5f905f90565b610ce93433610e9f565b50610cd9565b34610d38575b610cff3082610ecb565b90158015610d2f575b610d1d5798610d18818684610feb565b610b15575b5050505050505050505090505f905f90565b50898110610d08565b610d423433610e9f565b50610cf5565b5050505050505050509080610d62575b610cdf575f905f90565b50341515610d58565b50875f52600260205260405f2063ffffffff60e01b83165f5260205260ff60405f20541615610af8565b505050505050505050509080610d6257610cdf575f905f90565b5050505050905034610cdf575f905f90565b505050505090505f905f90565b5f525f602052600160405f20015490565b9190811015610def5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160e01b03198116810361012f5790565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615610e5057565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff1615610e895750565b63e2517d3f60e01b5f523360045260245260445ffd5b8115610ec4575f918291829182916001600160a01b03165af1610ec0610a60565b5090565b5050600190565b6040516370a0823160e01b602082019081526001600160a01b0390931660248083019190915281525f92839291610f03604482610a29565b51916001600160a01b03165afa610f18610a60565b90158015610f42575b610f3b5760208180518101031261012f5760200151600191565b505f905f90565b506020815110610f21565b90610f7960405163095ea7b360e01b6020820152610f73816101e05f866024840161115f565b836111a1565b610ec45760405163095ea7b360e01b6020820152610fa290610f73816101e08560248301611146565b15610fe557610fe291610fdd5f92610fcf60405194859263095ea7b360e01b60208501526024840161115f565b03601f198101845283610a29565b6111a1565b90565b50505f90565b9190610ff882828561117a565b61103c5760405163095ea7b360e01b602082015261102790611021816101e08560248301611146565b846111a1565b1561103557610fe29261117a565b5050505f90565b505050600190565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610fe5575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610fe5575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6001600160a01b0390911681525f602082015260400190565b6001600160a01b039091168152602081019190915260400190565b610fdd610fe29392610fcf60405194859263095ea7b360e01b60208501526024840161115f565b81516001600160a01b03909116915f91829160200182855af1906111c3610a60565b9115610fe55781519081156111fe575060208110156111e25750505f90565b816020918101031261012f5760200151801515810361012f5790565b9150503b15159056fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212200a362b979d1e198ea3254d72e191b7a9b76906669b9796b1884a0e076bcc20dc64736f6c634300081c0033"

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
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address , uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorTransactor) ExecuteExtraCall(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, toToken common.Address, arg3 common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "executeExtraCall", fromChainID, index, toToken, arg3, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address , uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, arg3 common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, fromChainID, index, toToken, arg3, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address , uint256 value, bytes extraData) payable returns(bool success, uint256 remaining)
func (_BridgeExecutor *BridgeExecutorTransactorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, arg3 common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, fromChainID, index, toToken, arg3, value, extraData)
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
