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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"executeExtraCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"isWhitelistedTarget\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeWhitelistTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"methodID\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"ExtraCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TargetWhitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEFailedToReturnNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BEInvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BENotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BETargetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
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
	Bin: "0x6080346100aa57601f610efc38819003918201601f19168301916001600160401b038311848410176100ae5780849260409485528339810103126100aa57610052602061004b836100c2565b92016100c2565b6001600160a01b038216158015610099575b61008a5761007461007a926100d6565b5061014c565b50604051610cbc90816101e08239f35b63484b059f60e11b5f5260045ffd5b506001600160a01b03811615610064565b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100aa57565b6001600160a01b0381165f9081525f516020610edc5f395f51905f52602052604090205460ff16610147576001600160a01b03165f8181525f516020610edc5f395f51905f5260205260408120805460ff191660011790553391905f516020610e9c5f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f516020610ebc5f395f51905f52602052604090205460ff16610147576001600160a01b03165f8181525f516020610ebc5f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f516020610e9c5f395f51905f529080a460019056fe608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461056d5750806311735ea5146104e45780631eeed513146103a1578063248a9ca31461037b5780632f2ff15d1461034a57806336568abe1461030657806361c8c47914610267578063907caa001461022a57806391d14854146101e2578063a217fddf146101c8578063d04323c5146100e65763d547741f146100ac575f61000f565b346100e25760403660031901126100e25761001a6004356100cb6105d6565b906100dd6100d88261097a565b6109da565b610b6e565b5f80fd5b346100e25760603660031901126100e2576100ff6105c0565b602435906044356001600160a01b038116918282036100e25761012061098b565b821580156101b7575b6101a8576001600160a01b0316906001820361016b5750505f8080938193828215610162575bf11561015757005b6040513d5f823e3d90fd5b506108fc61014f565b60405163a9059cbb60e01b602082015261001a949093506101a39184916101959160248401610a12565b03601f1981018452836105ec565b610bee565b63b56bf2e760e01b5f5260045ffd5b506001600160a01b03811615610129565b346100e2575f3660031901126100e25760206040515f8152f35b346100e25760403660031901126100e2576101fb6105d6565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576001600160a01b0361024b6105c0565b165f526001602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576102806105c0565b61028861098b565b6001600160a01b031680156102f757805f52600160205260ff60405f2054166102e857805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b63484b059f60e11b5f5260045ffd5b346100e25760403660031901126100e25761031f6105d6565b336001600160a01b0382160361033b5761001a90600435610b6e565b63334bd91960e11b5f5260045ffd5b346100e25760403660031901126100e25761001a6004356103696105d6565b906103766100d88261097a565b610ae6565b346100e25760203660031901126100e257602061039960043561097a565b604051908152f35b60c03660031901126100e2576044356001600160a01b03811681036100e2576064356001600160a01b03811681036100e25760a435906001600160401b0382116100e257366023830112156100e2576004820135906001600160401b0382116100e25736602483850101116100e257335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff16156104ad575f516020610c675f395f51905f525c61049e5760209360246104849460015f516020610c675f395f51905f525d019160843591602435600435610661565b5f5f516020610c675f395f51905f525d6040519015158152f35b633ee5aeb560e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b346100e25760203660031901126100e2576104fd6105c0565b61050561098b565b6001600160a01b03165f8181526001602052604090205460ff161561055e57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b346100e25760203660031901126100e2576004359063ffffffff60e01b82168092036100e257602091637965db0b60e01b81149081156105af575b5015158152f35b6301ffc9a760e01b149050836105a8565b600435906001600160a01b03821682036100e257565b602435906001600160a01b03821682036100e257565b601f909101601f19168101906001600160401b0382119082101761060f57604052565b634e487b7160e01b5f52604160045260245ffd5b3d1561065c573d906001600160401b03821161060f5760405191610651601f8201601f1916602001846105ec565b82523d5f602084013e565b606090565b9295919495939093811561096f576014873560601c970195875f52600160205260ff60405f20541615610960576001600160a01b03811696600188149485156108db578434036108cc575b62030d405a10156108a8575f5b86156108a15785905b826014116100e2575f9283928d6040519260131981018885378301856013198201528360131991030193f1986106f6610623565b9286156107f5575b8a15968761077c575b5050915f516020610c475f395f51905f529560209260e0959460405197889660018060a01b0316875285870152604086015263ffffffff861b903516606085015215608084015260c060a084015280519182918260c0860152018484015e5f828201840152601f01601f19168101030190a490565b9294939192156107cf57505f80808088335af1610797610623565b50156107c0575f516020610c475f395f51905f529560e0946020935b9294955092819750610707565b6317e90b5960e21b5f5260045ffd5b9560e0946020936107f0885f516020610c475f395f51905f529a3390610a2d565b6107b3565b60405160205f81830163095ea7b360e01b8152856024850152816044850152604484526108236064856105ec565b83519082865af15f513d82610885575b505015610841575b506106fe565b61087f9061087960405163095ea7b360e01b60208201528560248201525f6044820152604481526108736064826105ec565b84610bee565b82610bee565b5f61083b565b90915061089957508a3b15155b5f80610833565b600114610892565b5f906106c2565b5a62030d3f198101908111156106b957634e487b7160e01b5f52601160045260245ffd5b63867deca760e01b5f5260045ffd5b6040516370a0823160e01b81523060048201526020816024818d5afa80156101575786915f9161092b575b501061091c57610917858b85610a2d565b6106ac565b632417a7f760e11b5f5260045ffd5b9150506020813d602011610958575b81610947602093836105ec565b810103126100e2578590515f610906565b3d915061093a565b63a6bec22b60e01b5f5260045ffd5b505050505050505f90565b5f525f602052600160405f20015490565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156109c357565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff16156109fc5750565b63e2517d3f60e01b5f523360045260245260445ffd5b6001600160a01b039091168152602081019190915260400190565b91909160205f60405193610a6485610a568582019363095ea7b360e01b85528960248401610a12565b03601f1981018752866105ec565b84519082855af15f513d82610ac1575b505015610a8057505050565b60405163095ea7b360e01b60208201526001600160a01b0390931660248401525f6044808501919091528352610abf926101a3906108796064826105ec565b565b909150610ade57506001600160a01b0381163b15155b5f80610a74565b600114610ad7565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610b68575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610b68575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b905f602091828151910182855af115610157575f513d610c3d57506001600160a01b0381163b155b610c1d5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415610c1656fe0ea5b5ea4f72c6e156b490f3d2837bcc82b91a7870d46d7beacbd65c0841fa819b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212207dfbd6498ac77009490faf45f3a35bbd742253052085a0857e4d645e3ca4c7fe64736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0ddae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

// BridgeExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeExecutorMetaData.ABI instead.
var BridgeExecutorABI = BridgeExecutorMetaData.ABI

// BridgeExecutorBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const BridgeExecutorBinRuntime = "608080604052600436101561001c575b50361561001a575f80fd5b005b5f3560e01c90816301ffc9a71461056d5750806311735ea5146104e45780631eeed513146103a1578063248a9ca31461037b5780632f2ff15d1461034a57806336568abe1461030657806361c8c47914610267578063907caa001461022a57806391d14854146101e2578063a217fddf146101c8578063d04323c5146100e65763d547741f146100ac575f61000f565b346100e25760403660031901126100e25761001a6004356100cb6105d6565b906100dd6100d88261097a565b6109da565b610b6e565b5f80fd5b346100e25760603660031901126100e2576100ff6105c0565b602435906044356001600160a01b038116918282036100e25761012061098b565b821580156101b7575b6101a8576001600160a01b0316906001820361016b5750505f8080938193828215610162575bf11561015757005b6040513d5f823e3d90fd5b506108fc61014f565b60405163a9059cbb60e01b602082015261001a949093506101a39184916101959160248401610a12565b03601f1981018452836105ec565b610bee565b63b56bf2e760e01b5f5260045ffd5b506001600160a01b03811615610129565b346100e2575f3660031901126100e25760206040515f8152f35b346100e25760403660031901126100e2576101fb6105d6565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576001600160a01b0361024b6105c0565b165f526001602052602060ff60405f2054166040519015158152f35b346100e25760203660031901126100e2576102806105c0565b61028861098b565b6001600160a01b031680156102f757805f52600160205260ff60405f2054166102e857805f52600160205260405f20600160ff198254161790557f3d5490c3a5f2e74f927ed836bbeb02e5cf4f5e601f134742ec4c00235710b0825f80a2005b6305a88b4b60e11b5f5260045ffd5b63484b059f60e11b5f5260045ffd5b346100e25760403660031901126100e25761031f6105d6565b336001600160a01b0382160361033b5761001a90600435610b6e565b63334bd91960e11b5f5260045ffd5b346100e25760403660031901126100e25761001a6004356103696105d6565b906103766100d88261097a565b610ae6565b346100e25760203660031901126100e257602061039960043561097a565b604051908152f35b60c03660031901126100e2576044356001600160a01b03811681036100e2576064356001600160a01b03811681036100e25760a435906001600160401b0382116100e257366023830112156100e2576004820135906001600160401b0382116100e25736602483850101116100e257335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff16156104ad575f516020610c675f395f51905f525c61049e5760209360246104849460015f516020610c675f395f51905f525d019160843591602435600435610661565b5f5f516020610c675f395f51905f525d6040519015158152f35b633ee5aeb560e01b5f5260045ffd5b63e2517d3f60e01b5f52336004527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6360245260445ffd5b346100e25760203660031901126100e2576104fd6105c0565b61050561098b565b6001600160a01b03165f8181526001602052604090205460ff161561055e57805f52600160205260405f2060ff1981541690557f09443fc4fa9f38ce51944d7db4f773a5ffb81b85d76c7119b37194b3d83125675f80a2005b63c70ebef560e01b5f5260045ffd5b346100e25760203660031901126100e2576004359063ffffffff60e01b82168092036100e257602091637965db0b60e01b81149081156105af575b5015158152f35b6301ffc9a760e01b149050836105a8565b600435906001600160a01b03821682036100e257565b602435906001600160a01b03821682036100e257565b601f909101601f19168101906001600160401b0382119082101761060f57604052565b634e487b7160e01b5f52604160045260245ffd5b3d1561065c573d906001600160401b03821161060f5760405191610651601f8201601f1916602001846105ec565b82523d5f602084013e565b606090565b9295919495939093811561096f576014873560601c970195875f52600160205260ff60405f20541615610960576001600160a01b03811696600188149485156108db578434036108cc575b62030d405a10156108a8575f5b86156108a15785905b826014116100e2575f9283928d6040519260131981018885378301856013198201528360131991030193f1986106f6610623565b9286156107f5575b8a15968761077c575b5050915f516020610c475f395f51905f529560209260e0959460405197889660018060a01b0316875285870152604086015263ffffffff861b903516606085015215608084015260c060a084015280519182918260c0860152018484015e5f828201840152601f01601f19168101030190a490565b9294939192156107cf57505f80808088335af1610797610623565b50156107c0575f516020610c475f395f51905f529560e0946020935b9294955092819750610707565b6317e90b5960e21b5f5260045ffd5b9560e0946020936107f0885f516020610c475f395f51905f529a3390610a2d565b6107b3565b60405160205f81830163095ea7b360e01b8152856024850152816044850152604484526108236064856105ec565b83519082865af15f513d82610885575b505015610841575b506106fe565b61087f9061087960405163095ea7b360e01b60208201528560248201525f6044820152604481526108736064826105ec565b84610bee565b82610bee565b5f61083b565b90915061089957508a3b15155b5f80610833565b600114610892565b5f906106c2565b5a62030d3f198101908111156106b957634e487b7160e01b5f52601160045260245ffd5b63867deca760e01b5f5260045ffd5b6040516370a0823160e01b81523060048201526020816024818d5afa80156101575786915f9161092b575b501061091c57610917858b85610a2d565b6106ac565b632417a7f760e11b5f5260045ffd5b9150506020813d602011610958575b81610947602093836105ec565b810103126100e2578590515f610906565b3d915061093a565b63a6bec22b60e01b5f5260045ffd5b505050505050505f90565b5f525f602052600160405f20015490565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156109c357565b63e2517d3f60e01b5f52336004525f60245260445ffd5b5f8181526020818152604080832033845290915290205460ff16156109fc5750565b63e2517d3f60e01b5f523360045260245260445ffd5b6001600160a01b039091168152602081019190915260400190565b91909160205f60405193610a6485610a568582019363095ea7b360e01b85528960248401610a12565b03601f1981018752866105ec565b84519082855af15f513d82610ac1575b505015610a8057505050565b60405163095ea7b360e01b60208201526001600160a01b0390931660248401525f6044808501919091528352610abf926101a3906108796064826105ec565b565b909150610ade57506001600160a01b0381163b15155b5f80610a74565b600114610ad7565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610b68575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff1615610b68575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b905f602091828151910182855af115610157575f513d610c3d57506001600160a01b0381163b155b610c1d5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415610c1656fe0ea5b5ea4f72c6e156b490f3d2837bcc82b91a7870d46d7beacbd65c0841fa819b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212207dfbd6498ac77009490faf45f3a35bbd742253052085a0857e4d645e3ca4c7fe64736f6c634300081c0033"

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
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
func (_BridgeExecutor *BridgeExecutorTransactor) ExecuteExtraCall(opts *bind.TransactOpts, fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.contract.Transact(opts, "executeExtraCall", fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
func (_BridgeExecutor *BridgeExecutorSession) ExecuteExtraCall(fromChainID *big.Int, index *big.Int, toToken common.Address, to common.Address, value *big.Int, extraData []byte) (*types.Transaction, error) {
	return _BridgeExecutor.Contract.ExecuteExtraCall(&_BridgeExecutor.TransactOpts, fromChainID, index, toToken, to, value, extraData)
}

// ExecuteExtraCall is a paid mutator transaction binding the contract method 0x1eeed513.
//
// Solidity: function executeExtraCall(uint256 fromChainID, uint256 index, address toToken, address to, uint256 value, bytes extraData) payable returns(bool success)
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
	ToToken        common.Address
	To             common.Address
	Value          *big.Int
	TargetContract common.Address
	MethodID       [4]byte
	Success        bool
	Reason         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExtraCallExecuted is a free log retrieval operation binding the contract event 0x0ea5b5ea4f72c6e156b490f3d2837bcc82b91a7870d46d7beacbd65c0841fa81.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bytes4 methodID, bool success, bytes reason)
func (_BridgeExecutor *BridgeExecutorFilterer) FilterExtraCallExecuted(opts *bind.FilterOpts, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (*BridgeExecutorExtraCallExecutedIterator, error) {

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

	logs, sub, err := _BridgeExecutor.contract.FilterLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeExecutorExtraCallExecutedIterator{contract: _BridgeExecutor.contract, event: "ExtraCallExecuted", logs: logs, sub: sub}, nil
}

// WatchExtraCallExecuted is a free log subscription operation binding the contract event 0x0ea5b5ea4f72c6e156b490f3d2837bcc82b91a7870d46d7beacbd65c0841fa81.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bytes4 methodID, bool success, bytes reason)
func (_BridgeExecutor *BridgeExecutorFilterer) WatchExtraCallExecuted(opts *bind.WatchOpts, sink chan<- *BridgeExecutorExtraCallExecuted, fromChainID []*big.Int, index []*big.Int, toToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeExecutor.contract.WatchLogs(opts, "ExtraCallExecuted", fromChainIDRule, indexRule, toTokenRule)
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

// ParseExtraCallExecuted is a log parse operation binding the contract event 0x0ea5b5ea4f72c6e156b490f3d2837bcc82b91a7870d46d7beacbd65c0841fa81.
//
// Solidity: event ExtraCallExecuted(uint256 indexed fromChainID, uint256 indexed index, address indexed toToken, address to, uint256 value, address targetContract, bytes4 methodID, bool success, bytes reason)
func (_BridgeExecutor *BridgeExecutorFilterer) ParseExtraCallExecuted(log types.Log) (*BridgeExecutorExtraCallExecuted, error) {
	event := new(BridgeExecutorExtraCallExecuted)
	if err := _BridgeExecutor.contract.UnpackLog(event, "ExtraCallExecuted", log); err != nil {
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
