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

// BridgeExecuterMetaData contains all meta data concerning the BridgeExecuter contract.
var BridgeExecuterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"executeExtraCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isWhitelistedTarget\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEFailedToReturnNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BENotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"61c8c479": "addWhitelistTarget(address)",
		"1eeed513": "executeExtraCall(uint256,uint256,address,address,uint256,bytes)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"907caa00": "isWhitelistedTarget(address)",
		"d04323c5": "recoverToken(address,uint256,address)",
		"11735ea5": "removeWhitelistTarget(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x6080346100aa57601f610e3d38819003918201601f19168301916001600160401b038311848410176100ae5780849260409485528339810103126100aa57610052602061004b836100c2565b92016100c2565b6001600160a01b038216158015610099575b61008a5761007461007a926100d6565b5061014c565b50604051610bfd90816101e08239f35b63484b059f60e11b5f5260045ffd5b506001600160a01b03811615610064565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100aa57565b6001600160a01b0381165f9081525f516020610e1d5f395f51905f52602052604090205460ff16610147576001600160a01b03165f8181525f516020610e1d5f395f51905f5260205260408120805460ff191660011790553391905f516020610ddd5f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f516020610dfd5f395f51905f52602052604090205460ff16610147576001600160a01b03165f8181525f516020610dfd5f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f516020610ddd5f395f51905f529080a460019056fe608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461080c5750806311735ea5146107835780631eeed51314610386578063248a9ca3146103605780632f2ff15d1461032f57806336568abe146102eb57806361c8c4791461024c578063907caa001461020f57806391d14854146101c7578063a217fddf146101ad578063d04323c5146100e65763d547741f146100ac575f61000f565b346100e25760403660031901126100e25761001a6004356100cb610875565b906100dd6100d88261091b565b61097b565b610aef565b5f80fd5b346100e25760603660031901126100e2576100ff61085f565b602435906044356001600160a01b038116918282036100e25761012061092c565b821561019e576001600160a01b031690816101615750505f8080938193828215610158575bf11561014d57005b6040513d5f823e3d90fd5b506108fc610145565b60405163a9059cbb60e01b602082015261001a9490935061019991849161018b9160248401610900565b03601f19810184528361088b565b610b6f565b63b56bf2e760e01b5f5260045ffd5b346100e2575f3660031901126100e25760206040515f8152f35b346100e25760403660031901126100e2576101e0610875565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576001600160a01b0361023061085f565b165f526001602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e25761026561085f565b61026d61092c565b6001600160a01b031680156102dc57805f52600160205260ff60405f2054166102cd57805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b63484b059f60e11b5f5260045ffd5b346100e25760403660031901126100e257610304610875565b336001600160a01b038216036103205761001a90600435610aef565b63334bd91960e11b5f5260045ffd5b346100e25760403660031901126100e25761001a60043561034e610875565b9061035b6100d88261091b565b610a67565b346100e25760203660031901126100e257602061037e60043561091b565b604051908152f35b60c03660031901126100e2576044356001600160a01b038116908181036100e2576064356001600160a01b03811692908390036100e25760a435926084356001600160401b0385116100e257366023860112156100e2576004850135936001600160401b0385116100e25736602486880101116100e257335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff161561074c57602486013560601c95865f52600160205260ff60405f2054161561073d576001851495861561061657833403610607575b861561060157835b816014116100e2576040515f93849390601319820190603801833781018360131982015281601319910301918a5af1906104a46108c2565b91151595866105dc571561055457505f80808085335af16104c36108c2565b50156105455760209586925b60405194855283850152604084015284606084015260a0608084015280519182918260a08601520160c084015e5f60c082840101527fde3a7c1aa1e50f7a139705ce364660696c6cdaeb48e996766d24b47f5a115b896024359260c08160043594601f80199101168101030190a4604051908152f35b6317e90b5960e21b5f5260045ffd5b8661056291979392976109b3565b60405163a9059cbb60e01b815260208180610581853360048401610900565b03815f895af1610599575b50946020959186926104cf565b6020969196813d6020116105d4575b816105b56020938361088b565b810103126100e257519182151583036100e2579095909150602061058c565b3d91506105a8565b9287916020988995156105f1575b50506104cf565b6105fa916109b3565b88826105ea565b5f61046c565b63867deca760e01b5f5260045ffd5b6040516370a0823160e01b81523060048201526020816024818a5afa801561014d5785915f91610708575b50106106f95760405160205f8a61067b8461066d8a8683019463095ea7b360e01b865260248401610900565b03601f19810186528561088b565b83519082885af15f513d826106dd575b505015610699575b50610464565b6106d7906106d160405163095ea7b360e01b60208201528b60248201525f6044820152604481526106cb60648261088b565b86610b6f565b84610b6f565b88610693565b9091506106f15750863b15155b8a8061068b565b6001146106ea565b632417a7f760e11b5f5260045ffd5b9150506020813d602011610735575b816107246020938361088b565b810103126100e2578490518a610641565b3d9150610717565b63a6bec22b60e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b346100e25760203660031901126100e25761079c61085f565b6107a461092c565b6001600160a01b03165f8181526001602052604090205460ff16156107fd57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b346100e25760203660031901126100e2576004359063ffffffff60e01b82168092036100e257602091637965db0b60e01b811490811561084e575b5015158152f35b6301ffc9a760e01b14905083610847565b600435906001600160a01b03821682036100e257565b602435906001600160a01b03821682036100e257565b601f909101601f19168101906001600160401b038211908210176108ae57604052565b634e487b7160e01b5f52604160045260245ffd5b3d156108fb573d906001600160401b0382116108ae57604051916108f0601f8201601f19166020018461088b565b82523d5f602084013e565b606090565b6001600160a01b039091168152602081019190915260400190565b5f525f602052600160405f20015490565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561096457565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff161561099d5750565b63e2517d3f60e01b5f523360045260245260445ffd5b6040519060205f8184019463095ea7b360e01b865260018060a01b031694856024860152816044860152604485526109ec60648661088b565b84519082855af15f513d82610a42575b505015610a0857505050565b610199610a40936040519063095ea7b360e01b602083015260248201525f604482015260448152610a3a60648261088b565b82610b6f565b565b909150610a5f57506001600160a01b0381163b15155b5f806109fc565b600114610a58565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610ae9575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610ae9575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b905f602091828151910182855af11561014d575f513d610bbe57506001600160a01b0381163b155b610b9e5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415610b9756fea2646970667358221220779780712387d0d55cc9b0ce9e608deb8921f6c506d9144a362e6233cb1b2afb64736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0ddae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// BridgeExecuterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeExecuterMetaData.ABI instead.
var BridgeExecuterABI = BridgeExecuterMetaData.ABI

// BridgeExecuterBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeExecuterBinRuntime = "608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461080c5750806311735ea5146107835780631eeed51314610386578063248a9ca3146103605780632f2ff15d1461032f57806336568abe146102eb57806361c8c4791461024c578063907caa001461020f57806391d14854146101c7578063a217fddf146101ad578063d04323c5146100e65763d547741f146100ac575f61000f565b346100e25760403660031901126100e25761001a6004356100cb610875565b906100dd6100d88261091b565b61097b565b610aef565b5f80fd5b346100e25760603660031901126100e2576100ff61085f565b602435906044356001600160a01b038116918282036100e25761012061092c565b821561019e576001600160a01b031690816101615750505f8080938193828215610158575bf11561014d57005b6040513d5f823e3d90fd5b506108fc610145565b60405163a9059cbb60e01b602082015261001a9490935061019991849161018b9160248401610900565b03601f19810184528361088b565b610b6f565b63b56bf2e760e01b5f5260045ffd5b346100e2575f3660031901126100e25760206040515f8152f35b346100e25760403660031901126100e2576101e0610875565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576001600160a01b0361023061085f565b165f526001602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e25761026561085f565b61026d61092c565b6001600160a01b031680156102dc57805f52600160205260ff60405f2054166102cd57805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b63484b059f60e11b5f5260045ffd5b346100e25760403660031901126100e257610304610875565b336001600160a01b038216036103205761001a90600435610aef565b63334bd91960e11b5f5260045ffd5b346100e25760403660031901126100e25761001a60043561034e610875565b9061035b6100d88261091b565b610a67565b346100e25760203660031901126100e257602061037e60043561091b565b604051908152f35b60c03660031901126100e2576044356001600160a01b038116908181036100e2576064356001600160a01b03811692908390036100e25760a435926084356001600160401b0385116100e257366023860112156100e2576004850135936001600160401b0385116100e25736602486880101116100e257335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff161561074c57602486013560601c95865f52600160205260ff60405f2054161561073d576001851495861561061657833403610607575b861561060157835b816014116100e2576040515f93849390601319820190603801833781018360131982015281601319910301918a5af1906104a46108c2565b91151595866105dc571561055457505f80808085335af16104c36108c2565b50156105455760209586925b60405194855283850152604084015284606084015260a0608084015280519182918260a08601520160c084015e5f60c082840101527fde3a7c1aa1e50f7a139705ce364660696c6cdaeb48e996766d24b47f5a115b896024359260c08160043594601f80199101168101030190a4604051908152f35b6317e90b5960e21b5f5260045ffd5b8661056291979392976109b3565b60405163a9059cbb60e01b815260208180610581853360048401610900565b03815f895af1610599575b50946020959186926104cf565b6020969196813d6020116105d4575b816105b56020938361088b565b810103126100e257519182151583036100e2579095909150602061058c565b3d91506105a8565b9287916020988995156105f1575b50506104cf565b6105fa916109b3565b88826105ea565b5f61046c565b63867deca760e01b5f5260045ffd5b6040516370a0823160e01b81523060048201526020816024818a5afa801561014d5785915f91610708575b50106106f95760405160205f8a61067b8461066d8a8683019463095ea7b360e01b865260248401610900565b03601f19810186528561088b565b83519082885af15f513d826106dd575b505015610699575b50610464565b6106d7906106d160405163095ea7b360e01b60208201528b60248201525f6044820152604481526106cb60648261088b565b86610b6f565b84610b6f565b88610693565b9091506106f15750863b15155b8a8061068b565b6001146106ea565b632417a7f760e11b5f5260045ffd5b9150506020813d602011610735575b816107246020938361088b565b810103126100e2578490518a610641565b3d9150610717565b63a6bec22b60e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b346100e25760203660031901126100e25761079c61085f565b6107a461092c565b6001600160a01b03165f8181526001602052604090205460ff16156107fd57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b346100e25760203660031901126100e2576004359063ffffffff60e01b82168092036100e257602091637965db0b60e01b811490811561084e575b5015158152f35b6301ffc9a760e01b14905083610847565b600435906001600160a01b03821682036100e257565b602435906001600160a01b03821682036100e257565b601f909101601f19168101906001600160401b038211908210176108ae57604052565b634e487b7160e01b5f52604160045260245ffd5b3d156108fb573d906001600160401b0382116108ae57604051916108f0601f8201601f19166020018461088b565b82523d5f602084013e565b606090565b6001600160a01b039091168152602081019190915260400190565b5f525f602052600160405f20015490565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561096457565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff161561099d5750565b63e2517d3f60e01b5f523360045260245260445ffd5b6040519060205f8184019463095ea7b360e01b865260018060a01b031694856024860152816044860152604485526109ec60648661088b565b84519082855af15f513d82610a42575b505015610a0857505050565b610199610a40936040519063095ea7b360e01b602083015260248201525f604482015260448152610a3a60648261088b565b82610b6f565b565b909150610a5f57506001600160a01b0381163b15155b5f806109fc565b600114610a58565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610ae9575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610ae9575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b905f602091828151910182855af11561014d575f513d610bbe57506001600160a01b0381163b155b610b9e5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415610b9756fea2646970667358221220779780712387d0d55cc9b0ce9e608deb8921f6c506d9144a362e6233cb1b2afb64736f6c634300081c0033"

// Deprecated: Use BridgeExecuterMetaData.Sigs instead.
// BridgeExecuterFuncSigs maps the 4-byte function signature to its string representation.
var BridgeExecuterFuncSigs = BridgeExecuterMetaData.Sigs

// BridgeExecuterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeExecuterMetaData.Bin instead.
var BridgeExecuterBin = BridgeExecuterMetaData.Bin

// DeployBridgeExecuter deploys a new Ethereum contract, binding an instance of BridgeExecuter to it.
func DeployBridgeExecuter(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address, bridge common.Address) (common.Address, *types.Transaction, *BridgeExecuter, error) {
	parsed, err := BridgeExecuterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeExecuterBin), backend, owner, bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeExecuter{BridgeExecuterCaller: BridgeExecuterCaller{contract: contract}, BridgeExecuterTransactor: BridgeExecuterTransactor{contract: contract}, BridgeExecuterFilterer: BridgeExecuterFilterer{contract: contract}}, nil
}

// BridgeExecuter is an auto generated Go binding around an Ethereum contract.
type BridgeExecuter struct {
	BridgeExecuterCaller     // Read-only binding to the contract
	BridgeExecuterTransactor // Write-only binding to the contract
	BridgeExecuterFilterer   // Log filterer for contract events
}

// BridgeExecuterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeExecuterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecuterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeExecuterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecuterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeExecuterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeExecuterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeExecuterSession struct {
	Contract     *BridgeExecuter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeExecuterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeExecuterCallerSession struct {
	Contract *BridgeExecuterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BridgeExecuterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeExecuterTransactorSession struct {
	Contract     *BridgeExecuterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BridgeExecuterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeExecuterRaw struct {
	Contract *BridgeExecuter // Generic contract binding to access the raw methods on
}

// BridgeExecuterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeExecuterCallerRaw struct {
	Contract *BridgeExecuterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeExecuterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeExecuterTransactorRaw struct {
	Contract *BridgeExecuterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeExecuter creates a new instance of BridgeExecuter, bound to a specific deployed contract.
func NewBridgeExecuter(address common.Address, backend bind.ContractBackend) (*BridgeExecuter, error) {
	contract, err := bindBridgeExecuter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuter{BridgeExecuterCaller: BridgeExecuterCaller{contract: contract}, BridgeExecuterTransactor: BridgeExecuterTransactor{contract: contract}, BridgeExecuterFilterer: BridgeExecuterFilterer{contract: contract}}, nil
}

// NewBridgeExecuterCaller creates a new read-only instance of BridgeExecuter, bound to a specific deployed contract.
func NewBridgeExecuterCaller(address common.Address, caller bind.ContractCaller) (*BridgeExecuterCaller, error) {
	contract, err := bindBridgeExecuter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterCaller{contract: contract}, nil
}

// NewBridgeExecuterTransactor creates a new write-only instance of BridgeExecuter, bound to a specific deployed contract.
func NewBridgeExecuterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeExecuterTransactor, error) {
	contract, err := bindBridgeExecuter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterTransactor{contract: contract}, nil
}

// NewBridgeExecuterFilterer creates a new log filterer instance of BridgeExecuter, bound to a specific deployed contract.
func NewBridgeExecuterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeExecuterFilterer, error) {
	contract, err := bindBridgeExecuter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterFilterer{contract: contract}, nil
}

// bindBridgeExecuter binds a generic wrapper to an already deployed contract.
func bindBridgeExecuter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeExecuterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeExecuter *BridgeExecuterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeExecuter.Contract.BridgeExecuterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeExecuter *BridgeExecuterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.BridgeExecuterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeExecuter *BridgeExecuterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.BridgeExecuterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeExecuter *BridgeExecuterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeExecuter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeExecuter *BridgeExecuterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeExecuter *BridgeExecuterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeExecuter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeExecuter.Contract.DEFAULTADMINROLE(&_BridgeExecuter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeExecuter.Contract.DEFAULTADMINROLE(&_BridgeExecuter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeExecuter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeExecuter.Contract.GetRoleAdmin(&_BridgeExecuter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeExecuter *BridgeExecuterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeExecuter.Contract.GetRoleAdmin(&_BridgeExecuter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeExecuter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecuter *BridgeExecuterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeExecuter.Contract.HasRole(&_BridgeExecuter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeExecuter.Contract.HasRole(&_BridgeExecuter.CallOpts, role, account)
}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCaller) IsWhitelistedTarget(opts *bind.CallOpts, target common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeExecuter.contract.Call(opts, &out, "isWhitelistedTarget", target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecuter *BridgeExecuterSession) IsWhitelistedTarget(target common.Address) (bool, error) {
	return _BridgeExecuter.Contract.IsWhitelistedTarget(&_BridgeExecuter.CallOpts, target)
}

// IsWhitelistedTarget is a free data retrieval call binding the contract method 0x907caa00.
//
// Solidity: function isWhitelistedTarget(address target) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCallerSession) IsWhitelistedTarget(target common.Address) (bool, error) {
	return _BridgeExecuter.Contract.IsWhitelistedTarget(&_BridgeExecuter.CallOpts, target)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeExecuter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecuter *BridgeExecuterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeExecuter.Contract.SupportsInterface(&_BridgeExecuter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeExecuter *BridgeExecuterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeExecuter.Contract.SupportsInterface(&_BridgeExecuter.CallOpts, interfaceId)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) AddWhitelistTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "addWhitelistTarget", target)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterSession) AddWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.AddWhitelistTarget(&_BridgeExecuter.TransactOpts, target)
}

// AddWhitelistTarget is a paid mutator transaction binding the contract method 0x61c8c479.
//
// Solidity: function addWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) AddWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.AddWhitelistTarget(&_BridgeExecuter.TransactOpts, target)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
func (_BridgeExecuter *BridgeExecuterTransactor) ExecuteExtraCall(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "executeExtraCall", fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
func (_BridgeExecuter *BridgeExecuterSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.ExecuteExtraCall(&_BridgeExecuter.TransactOpts, fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
func (_BridgeExecuter *BridgeExecuterTransactorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.ExecuteExtraCall(&_BridgeExecuter.TransactOpts, fromChainID, index, toToken, to, value, extraData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.GrantRole(&_BridgeExecuter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.GrantRole(&_BridgeExecuter.TransactOpts, role, account)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) RecoverToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "recoverToken", token, amount, recipient)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecuter *BridgeExecuterSession) RecoverToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RecoverToken(&_BridgeExecuter.TransactOpts, token, amount, recipient)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xd04323c5.
//
// Solidity: function recoverToken(address token, uint256 amount, address recipient) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) RecoverToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RecoverToken(&_BridgeExecuter.TransactOpts, token, amount, recipient)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) RemoveWhitelistTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "removeWhitelistTarget", target)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterSession) RemoveWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RemoveWhitelistTarget(&_BridgeExecuter.TransactOpts, target)
}

// RemoveWhitelistTarget is a paid mutator transaction binding the contract method 0x11735ea5.
//
// Solidity: function removeWhitelistTarget(address target) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) RemoveWhitelistTarget(target common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RemoveWhitelistTarget(&_BridgeExecuter.TransactOpts, target)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecuter *BridgeExecuterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RenounceRole(&_BridgeExecuter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RenounceRole(&_BridgeExecuter.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RevokeRole(&_BridgeExecuter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeExecuter.Contract.RevokeRole(&_BridgeExecuter.TransactOpts, role, account)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecuter *BridgeExecuterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeExecuter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecuter *BridgeExecuterSession) Receive() (*types.Transaction, error) {
	return _BridgeExecuter.Contract.Receive(&_BridgeExecuter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeExecuter *BridgeExecuterTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeExecuter.Contract.Receive(&_BridgeExecuter.TransactOpts)
}

// BridgeExecuterExtraCallExecutedIterator is returned from FilterExtraCallExecuted and is used to iterate over the raw logs and unpacked data for ExtraCallExecuted events raised by the BridgeExecuter contract.
type BridgeExecuterExtraCallExecutedIterator struct {
	Event *BridgeExecuterExtraCallExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterExtraCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterExtraCallExecuted)
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
		it.Event = new(BridgeExecuterExtraCallExecuted)
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
func (it *BridgeExecuterExtraCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterExtraCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterExtraCallExecuted represents a ExtraCallExecuted event raised by the BridgeExecuter contract.
type BridgeExecuterExtraCallExecuted struct {
	FromChainID    *big.Int
	Index          *big.Int
	ToToken        common.Address
	To             common.Address
	Value          *big.Int
	TargetContract common.Address
	Success        bool
	Reason         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0xde3a7c1aa1e50f7a139705ce364660696c6cdaeb48e996766d24b47f5a115b89.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bool success, bytes reason)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BridgeExecuterExtraCallExecutedIterator, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterExtraCallExecutedIterator{contract: _BridgeExecuter.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0xde3a7c1aa1e50f7a139705ce364660696c6cdaeb48e996766d24b47f5a115b89.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bool success, bytes reason)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BridgeExecuterExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

	var fromChainIDRule []interface{}
	for _, fromChainIDItem := range fromChainID {
		fromChainIDRule = append(fromChainIDRule, fromChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterExtraCallExecuted)
				if err := _BridgeExecuter.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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

// ParseExtraCallExecuted is a log parse operation binding the contract event 0xde3a7c1aa1e50f7a139705ce364660696c6cdaeb48e996766d24b47f5a115b89.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bool success, bytes reason)
func (_BridgeExecuter *BridgeExecuterFilterer) ParseExtraCallExecuted(log types.Log) (*BridgeExecuterExtraCallExecuted, error) {
	event := new(BridgeExecuterExtraCallExecuted)
	if err := _BridgeExecuter.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeExecuter contract.
type BridgeExecuterRoleAdminChangedIterator struct {
	Event *BridgeExecuterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterRoleAdminChanged)
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
		it.Event = new(BridgeExecuterRoleAdminChanged)
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
func (it *BridgeExecuterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeExecuter contract.
type BridgeExecuterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeExecuterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterRoleAdminChangedIterator{contract: _BridgeExecuter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeExecuterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterRoleAdminChanged)
				if err := _BridgeExecuter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeExecuter *BridgeExecuterFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeExecuterRoleAdminChanged, error) {
	event := new(BridgeExecuterRoleAdminChanged)
	if err := _BridgeExecuter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeExecuter contract.
type BridgeExecuterRoleGrantedIterator struct {
	Event *BridgeExecuterRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterRoleGranted)
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
		it.Event = new(BridgeExecuterRoleGranted)
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
func (it *BridgeExecuterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterRoleGranted represents a RoleGranted event raised by the BridgeExecuter contract.
type BridgeExecuterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeExecuterRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterRoleGrantedIterator{contract: _BridgeExecuter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeExecuterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterRoleGranted)
				if err := _BridgeExecuter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeExecuter *BridgeExecuterFilterer) ParseRoleGranted(log types.Log) (*BridgeExecuterRoleGranted, error) {
	event := new(BridgeExecuterRoleGranted)
	if err := _BridgeExecuter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeExecuter contract.
type BridgeExecuterRoleRevokedIterator struct {
	Event *BridgeExecuterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterRoleRevoked)
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
		it.Event = new(BridgeExecuterRoleRevoked)
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
func (it *BridgeExecuterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterRoleRevoked represents a RoleRevoked event raised by the BridgeExecuter contract.
type BridgeExecuterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeExecuterRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterRoleRevokedIterator{contract: _BridgeExecuter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeExecuterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterRoleRevoked)
				if err := _BridgeExecuter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeExecuter *BridgeExecuterFilterer) ParseRoleRevoked(log types.Log) (*BridgeExecuterRoleRevoked, error) {
	event := new(BridgeExecuterRoleRevoked)
	if err := _BridgeExecuter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuterTargetRemovedFromWhitelistIterator is returned from FilterTargetRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for TargetRemovedFromWhitelist events raised by the BridgeExecuter contract.
type BridgeExecuterTargetRemovedFromWhitelistIterator struct {
	Event *BridgeExecuterTargetRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterTargetRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterTargetRemovedFromWhitelist)
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
		it.Event = new(BridgeExecuterTargetRemovedFromWhitelist)
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
func (it *BridgeExecuterTargetRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterTargetRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterTargetRemovedFromWhitelist represents a TargetRemovedFromWhitelist event raised by the BridgeExecuter contract.
type BridgeExecuterTargetRemovedFromWhitelist struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d8312567.
//
// Solidity: event TargetRemovedFromWhitelist(address indexed target)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterTargetRemovedFromWhitelist(opts *bind.FilterOpts, target []common.Address) (*BridgeExecuterTargetRemovedFromWhitelistIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "TargetRemovedFromWhitelist", targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterTargetRemovedFromWhitelistIterator{contract: _BridgeExecuter.contract, event: "TargetRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchTargetRemovedFromWhitelist is a free log subscription operation binding the contract event 0x09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d8312567.
//
// Solidity: event TargetRemovedFromWhitelist(address indexed target)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchTargetRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *BridgeExecuterTargetRemovedFromWhitelist, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "TargetRemovedFromWhitelist", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterTargetRemovedFromWhitelist)
				if err := _BridgeExecuter.contract.UnpackLog(event, "TargetRemovedFromWhitelist", log); err != nil {
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
func (_BridgeExecuter *BridgeExecuterFilterer) ParseTargetRemovedFromWhitelist(log types.Log) (*BridgeExecuterTargetRemovedFromWhitelist, error) {
	event := new(BridgeExecuterTargetRemovedFromWhitelist)
	if err := _BridgeExecuter.contract.UnpackLog(event, "TargetRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeExecuterTargetWhitelistedIterator is returned from FilterTargetWhitelisted and is used to iterate over the raw logs and unpacked data for TargetWhitelisted events raised by the BridgeExecuter contract.
type BridgeExecuterTargetWhitelistedIterator struct {
	Event *BridgeExecuterTargetWhitelisted // Event containing the contract specifics and raw log

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
func (it *BridgeExecuterTargetWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeExecuterTargetWhitelisted)
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
		it.Event = new(BridgeExecuterTargetWhitelisted)
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
func (it *BridgeExecuterTargetWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeExecuterTargetWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeExecuterTargetWhitelisted represents a TargetWhitelisted event raised by the BridgeExecuter contract.
type BridgeExecuterTargetWhitelisted struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetWhitelisted is a free log retrieval operation binding the contract event 0x3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b082.
//
// Solidity: event TargetWhitelisted(address indexed target)
func (_BridgeExecuter *BridgeExecuterFilterer) FilterTargetWhitelisted(opts *bind.FilterOpts, target []common.Address) (*BridgeExecuterTargetWhitelistedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecuter.contract.FilterLogs(opts, "TargetWhitelisted", targetRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecuterTargetWhitelistedIterator{contract: _BridgeExecuter.contract, event: "TargetWhitelisted", logs: logs, sub: sub}, nil
}

// WatchTargetWhitelisted is a free log subscription operation binding the contract event 0x3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b082.
//
// Solidity: event TargetWhitelisted(address indexed target)
func (_BridgeExecuter *BridgeExecuterFilterer) WatchTargetWhitelisted(opts *bind.WatchOpts, sink chan<- *BridgeExecuterTargetWhitelisted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _BridgeExecuter.contract.WatchLogs(opts, "TargetWhitelisted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeExecuterTargetWhitelisted)
				if err := _BridgeExecuter.contract.UnpackLog(event, "TargetWhitelisted", log); err != nil {
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
func (_BridgeExecuter *BridgeExecuterFilterer) ParseTargetWhitelisted(log types.Log) (*BridgeExecuterTargetWhitelisted, error) {
	event := new(BridgeExecuterTargetWhitelisted)
	if err := _BridgeExecuter.contract.UnpackLog(event, "TargetWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
