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

// BridgeBotBridgeConfig is an auto generated low-level Go binding around an user-defined struct.
type BridgeBotBridgeConfig struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}

// BridgeBotMetaData contains all meta data concerning the BridgeBot contract.
var BridgeBotMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_adminDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addBridgeConfig\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBaseBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canExecuteBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"canExecute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBridgeBatch\",\"inputs\":[{\"name\":\"configIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutableConfigs\",\"inputs\":[{\"name\":\"maxConfigs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"executableConfigs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantExecutorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextConfigId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeExecutorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toggleBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeConfigAdded\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigToggled\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigUpdated\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeExecuted\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BridgeBotBridgeFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotInsufficientBalance\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BridgeBotNotTimeYet\",\"inputs\":[{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a03461012c57601f6123ef38819003918201601f19168301916001600160401b038311848410176101305780849260809460405283398101031261012c5761004781610144565b9061005460208201610144565b91606061006360408401610144565b92015165ffffffffffff8116810361012c576001600160a01b0384161561011957600180546001600160d01b031660d09290921b6001600160d01b0319169190911790556100b083610158565b5060016003556001600160a01b0316801561010a576001600160a01b0382161561010a576100ea926100e4916080526101a0565b506101a0565b506040516121a0908161024f82396080518181816103d10152611a5b0152f35b630508665f60e41b5f5260045ffd5b636116401160e11b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361012c57565b600254906001600160a01b038216610191576001600160a01b03199091166001600160a01b0382161760025561018e905f6101c6565b90565b631fe1e13d60e11b5f5260045ffd5b61018e907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e635b5f818152602081815260408083206001600160a01b038616845290915290205460ff16610248575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f9056fe60806040526004361015610010575b005b5f3560e01c806301e33667146113c857806301ffc9a714611358578063022d63fb1461133b57806307b18bde1461128657806307bd02651461125f5780630aa6220b146111cd578063248a9ca31461119b5780632f2ff15d1461115857806331f7d9641461113d57806336568abe146110525780634624e68014610f0157806349a396b914610e9857806350f760e914610da4578063634e93da14610cc8578063649a5ec714610b2f57806370d2ddf414610a4d57806376e95f161461093257806380a4d414146108b857806384ef8ffc146108905780638da5cb5b1461089057806391d148541461084857806399d726c71461082b578063a1eda53c146107c9578063a217fddf146107af578063b157607414610779578063bd5f0afb146106e7578063cc8463c8146106bd578063cefc142914610611578063cf6eefb7146105d8578063d172f2f014610562578063d547741f1461050b578063d602b9fd146104aa578063d9f66db11461042c578063e1068d8d14610400578063e78cea92146103bc578063e9cf510c146102c85763ed2859d90361000e57346102c45760a03660031901126102c4576004356101c7611462565b6044356001600160a01b038116908190036102c457608435916101e86117ee565b5f8481526004602052604090205461020a906001600160a01b031615156115f8565b6001600160a01b038116156102b55781156102b55782156102a6575f848152600460205260409081902080546001600160a01b03939093166001600160a01b0319938416178155600181018054909316939093179091556064356002830155600382019290925590517fe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b7659181906102a19082611586565b0390a2005b6319e9855160e11b5f5260045ffd5b630508665f60e41b5f5260045ffd5b5f80fd5b346102c45760203660031901126102c4576004356102e5816117a8565b905f5f91600554925b838110806103b3575b1561033a57610305816116a9565b50610319575b61031490611578565b6102ee565b91610332818461032c61031494896117da565b52611578565b92905061030b565b8483610345816117a8565b915f5b82811061039457836040518091602082016020835281518091526020604084019201905f5b81811061037b575050500390f35b825184528594506020938401939092019160010161036d565b806103a1600192846117da565b516103ac82876117da565b5201610348565b508183106102f7565b346102c4575f3660031901126102c4576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346102c45760203660031901126102c457604061041e6004356116a9565b825191151582526020820152f35b346102c45760203660031901126102c45761044561144c565b61044d6117ee565b6001600160a01b031680156102b557478061046457005b6020816104a15f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af161049b6114fb565b5061153a565b604051908152a2005b346102c4575f3660031901126102c4576104c26117ee565b600180546001600160d01b0319811690915560a01c65ffffffffffff166104e557005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051095f80a1005b346102c45760403660031901126102c457600435610527611462565b8115610553578161054e61054961000e945f525f602052600160405f20015490565b611899565b611fd5565b631fe1e13d60e11b5f5260045ffd5b346102c45760203660031901126102c4576004355f52600460205260c060405f2060018060a01b038154169060018060a01b0360018201541690600281015460038201549060ff6005600485015494015416936040519586526020860152604085015260608401526080830152151560a0820152f35b346102c4575f3660031901126102c457600154604080516001600160a01b038316815260a09290921c65ffffffffffff16602083015290f35b346102c4575f3660031901126102c4576001546001600160a01b031633036106aa5760015460a081901c65ffffffffffff16906001600160a01b0316811580156106a0575b61068d5760025461067a9190610674906001600160a01b0316611f9f565b50611f10565b50600180546001600160d01b0319169055005b506319ca5ebb60e01b5f5260045260245ffd5b5042821015610656565b636116401160e11b5f523360045260245ffd5b346102c4575f3660031901126102c45760206106d7611638565b65ffffffffffff60405191168152f35b346102c45760403660031901126102c457600435602435801515908181036102c4577f82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0916104a16020926107396117ee565b5f86815260048552604090205461075a906001600160a01b031615156115f8565b855f5260048452600560405f20019060ff801983541691151516179055565b346102c45760403660031901126102c45761079261183d565b61079a61197d565b6107a86024356004356119fc565b6001600355005b346102c4575f3660031901126102c45760206040515f8152f35b346102c4575f3660031901126102c4576002548060d01c9081151580610821575b156108175760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b50505f5f906107fb565b50428210156107ea565b346102c4575f3660031901126102c4576020600554604051908152f35b346102c45760403660031901126102c457610861611462565b6004355f525f60205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346102c4575f3660031901126102c4576002546040516001600160a01b039091168152602090f35b346102c45760203660031901126102c4576108d161144c565b6108d96117ee565b6001600160a01b038116156102b5575f51602061214b5f395f51905f525f9081526020527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d706a5461000e919061092d90611899565b611ef9565b346102c45760403660031901126102c45760043567ffffffffffffffff81116102c457610963903690600401611478565b9060243567ffffffffffffffff81116102c457610984903690600401611478565b929061098e61183d565b61099661197d565b838203610a10575f5b8281106109ad576001600355005b806109c46109be60019386886115d4565b356116a9565b50806109fc575b6109d6575b0161099f565b6109f76109e48286886115d4565b356109f08389876115d4565b35906119fc565b6109d0565b50610a088187856115d4565b3515156109cb565b60405162461bcd60e51b8152602060048201526015602482015274082e4e4c2f240d8cadccee8d040dad2e6dac2e8c6d605b1b6044820152606490fd5b346102c45760203660031901126102c4575f60a0604051610a6d816114a9565b82815282602082015282604082015282606082015282608082015201526004355f52600460205260c060405f20604051610aa6816114a9565b60018060a01b038254169182825260018060a01b03600182015416602083019081526002820154604084019081526003830154916060850192835260a060ff6005600487015496608089019788520154169501941515855260405195865260018060a01b03905116602086015251604085015251606084015251608083015251151560a0820152f35b346102c45760203660031901126102c45760043565ffffffffffff8116908181036102c457610b5c6117ee565b610b6542612013565b9165ffffffffffff610b75611638565b1680821115610c7a57507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9265ffffffffffff826206978080610bc295109118026206978018169061195f565b906002548060d01c80610c1e575b5050600280546001600160a01b031660a083901b65ffffffffffff60a01b161760d084901b6001600160d01b0319161790556040805165ffffffffffff9283168152919092166020820152a1005b421115610c5057600180546001600160d01b031660309290921b6001600160d01b0319169190911790555b8380610bd0565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec55f80a1610c49565b0365ffffffffffff8111610cb4577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92610bc2919061195f565b634e487b7160e01b5f52601160045260245ffd5b346102c45760203660031901126102c457610ce161144c565b610ce96117ee565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020610d26610d1842612013565b610d20611638565b9061195f565b600180546001600160d01b031981166001600160a01b039690961695861760a084811b65ffffffffffff60a01b169190911790925565ffffffffffff911c16610d7b575b65ffffffffffff60405191168152a2005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051095f80a1610d6a565b346102c45760403660031901126102c457610dbd61144c565b610dc5611462565b90610dce6117ee565b6001600160a01b031680156102b5576001600160a01b0382169182156102b5576040516370a0823160e01b815230600482015290602082602481865afa918215610e8d575f92610e59575b5081610e2157005b81610e507f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e562093602093866118d1565b604051908152a3005b9091506020813d602011610e85575b81610e75602093836114d9565b810103126102c457519084610e19565b3d9150610e68565b6040513d5f823e3d90fd5b346102c45760203660031901126102c45761000e610eb461144c565b610ebc6117ee565b5f51602061214b5f395f51905f525f9081526020527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d706a54610efc90611899565b611f88565b346102c45760803660031901126102c457610f1a61144c565b610f22611462565b9060643590610f2f6117ee565b6001600160a01b03169182156102b5576001600160a01b031680156102b55781156102a657602092600561100b92815494610f6986611578565b835560405193610f78856114a9565b84528684019182526044356040808601918252606086019283525f60808701818152600160a089018181528b84526004808e5294909320985189546001600160a01b039182166001600160a01b0319918216178b559751918a0180549290911691909716179095559151600287015591516003860155915190840155519101805460ff191691151560ff16919091179055565b805f5260048252807f607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb12361104760405f2060405191829182611586565b0390a2604051908152f35b346102c45760403660031901126102c45760043561106e611462565b811580611126575b6110a3575b336001600160a01b038216036110945761000e91611fd5565b63334bd91960e11b5f5260045ffd5b60015465ffffffffffff60a082901c16906001600160a01b031615801590611116575b8015611104575b6110e957506001805465ffffffffffff60a01b1916905561107b565b65ffffffffffff906319ca5ebb60e01b5f521660045260245ffd5b504265ffffffffffff821610156110cd565b5065ffffffffffff8116156110c6565b506002546001600160a01b03828116911614611076565b346102c4575f3660031901126102c457602060405160018152f35b346102c45760403660031901126102c457600435611174611462565b8115610553578161119661054961000e945f525f602052600160405f20015490565b611f46565b346102c45760203660031901126102c45760206111c56004355f525f602052600160405f20015490565b604051908152f35b346102c4575f3660031901126102c4576111e56117ee565b6002548060d01c80611203575b600280546001600160a01b03169055005b42111561123557600180546001600160d01b031660309290921b6001600160d01b0319169190911790555b80806111f2565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec55f80a161122e565b346102c4575f3660031901126102c45760206040515f51602061214b5f395f51905f528152f35b346102c45760403660031901126102c45761129f61144c565b602435906112ab6117ee565b6001600160a01b03169081156102b55780156102a6578047106112ff576020816104a15f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af161049b6114fb565b60405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606490fd5b346102c4575f3660031901126102c4576020604051620697808152f35b346102c45760203660031901126102c45760043563ffffffff60e01b81168091036102c4576020906318a4c3c360e11b811490811561139d575b506040519015158152f35b637965db0b60e01b8114915081156113b7575b5082611392565b6301ffc9a760e01b149050826113b0565b346102c45760603660031901126102c4576113e161144c565b6113e9611462565b90604435906113f66117ee565b6001600160a01b03169081156102b5576001600160a01b0383169283156102b55781156102a65781610e507f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e562093602093866118d1565b600435906001600160a01b03821682036102c457565b602435906001600160a01b03821682036102c457565b9181601f840112156102c45782359167ffffffffffffffff83116102c4576020808501948460051b0101116102c457565b60c0810190811067ffffffffffffffff8211176114c557604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176114c557604052565b3d15611535573d9067ffffffffffffffff82116114c5576040519161152a601f8201601f1916602001846114d9565b82523d5f602084013e565b606090565b1561154157565b60405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606490fd5b5f198114610cb45760010190565b81546001600160a01b039081168252600183015416602082015260028201546040820152600382015460608201526004820154608082015260059091015460ff16151560a082015260c00190565b91908110156115e45760051b0190565b634e487b7160e01b5f52603260045260245ffd5b156115ff57565b60405162461bcd60e51b8152602060048201526011602482015270436f6e666967206e6f742065786973747360781b6044820152606490fd5b6002548060d01c8015159081611667575b501561165d5760a01c65ffffffffffff1690565b5060015460d01c90565b905042115f611649565b811561167b570690565b634e487b7160e01b5f52601260045260245ffd5b91908203918211610cb457565b91908201809211610cb457565b5f52600460205260405f20906040516116c1816114a9565b82546001600160a01b03908116808352600185015490911660208301526002840154604083015260038401546060830181815260048601546080850181815260059097015460ff1615801560a0909601869052919492939092611787575b5061177c57156117715761173f6117396117529242611671565b4261168f565b935161174c835182611671565b9061168f565b80931192835f146117635750505f90565b61176e91519061169c565b90565b505090506001905f90565b50505090505f905f90565b9050155f61171f565b67ffffffffffffffff81116114c55760051b60200190565b906117b282611790565b6117bf60405191826114d9565b82815280926117d0601f1991611790565b0190602036910137565b80518210156115e45760209160051b010190565b335f9081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561182657565b63e2517d3f60e01b5f52336004525f60245260445ffd5b335f9081527fdae2aa361dfd1ca020a396615627d436107c35eff9fe7738a3512819782d7069602052604090205460ff161561187557565b63e2517d3f60e01b5f52336004525f51602061214b5f395f51905f5260245260445ffd5b5f8181526020818152604080832033845290915290205460ff16156118bb5750565b63e2517d3f60e01b5f523360045260245260445ffd5b60405163a9059cbb60e01b60208281019182526001600160a01b03909416602483015260448083019590955293815290925f9161190f6064826114d9565b519082855af115610e8d575f513d61195657506001600160a01b0381163b155b6119365750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b6001141561192f565b9065ffffffffffff8091169116019065ffffffffffff8211610cb457565b60026003541461198e576002600355565b633ee5aeb560e01b5f5260045ffd5b908160209103126102c4575180151581036102c45790565b9081526001600160a01b039182166020820152911660408201526060810191909152608081019190915260a081019190915260e060c082018190525f908201526101000190565b81156102a657805f52600460205260405f209060ff60058301541680611ee6575b15611eaa57611a2b816116a9565b9015611e98575081546001600160a01b03166001811493908415611e355747945b604051637838174760e11b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610e8d575f92611df1575b506002870197606089546064604051809681936337dba1f760e21b835260048301528960248301528a604483015260018060a01b03165afa908115610e8d575f5f945f93611dad575b50611af783611af2878b61169c565b61169c565b91828110611d9757508710611d5b57869415611c8a575b88549495506020946001600160a01b031660018103611c1957508954895460018b0154604051632fe9316f60e11b815298899788968795611b659591949293926001600160a01b03908116921690600488016119b5565b03925af1908115610e8d575f91611bea575b505b15611bdb5760807f96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee9142600486015560018060a01b03855416946001808060a01b039101541695546040519182526020820152336040820152426060820152a4565b6307c4732760e51b5f5260045ffd5b611c0c915060203d602011611c12575b611c0481836114d9565b81019061199d565b5f611b77565b503d611bfa565b8a5460018b0154604051632fe9316f60e11b815298899788965f96508795611c539591949293926001600160a01b031691600488016119b5565b03925af1908115610e8d575f91611c6b575b50611b79565b611c84915060203d602011611c1257611c0481836114d9565b5f611c65565b604051636eb1769f60e11b8152306004820152602481018490529450602085604481895afa8015610e8d5787955f91611d25575b50811115611b0e57935060205f9560446040518098819363095ea7b360e01b8352876004840152811960248401525af1938415610e8d578695602095611d08575b50859450611b0e565b611d1e90863d8811611c1257611c0481836114d9565b505f611cff565b9550506020853d602011611d53575b81611d41602093836114d9565b810103126102c4578087955190611cbe565b3d9150611d34565b60405162461bcd60e51b8152602060048201526014602482015273416d6f756e742062656c6f77206d696e696d756d60601b6044820152606490fd5b8263203b880360e11b5f5260045260245260445ffd5b92505092506060813d606011611de9575b81611dcb606093836114d9565b810103126102c457805160208201516040909201519193905f611ae3565b3d9150611dbe565b9091506020813d602011611e2d575b81611e0d602093836114d9565b810103126102c457516001600160a01b03811681036102c457905f611a9a565b3d9150611e00565b6040516370a0823160e01b8152306004820152602081602481855afa908115610e8d575f91611e66575b5094611a4c565b90506020813d602011611e90575b81611e81602093836114d9565b810103126102c457515f611e5f565b3d9150611e74565b6357c8d07f60e11b5f5260045260245ffd5b60405162461bcd60e51b8152602060048201526014602482015273436f6e666967206e6f7420617661696c61626c6560601b6044820152606490fd5b5081546001600160a01b03161515611a1d565b61176e905f51602061214b5f395f51905f52612042565b600254906001600160a01b038216610553576001600160a01b03199091166001600160a01b0382161760025561176e905f612042565b908115611f57575b61176e91612042565b600254916001600160a01b038316610553576001600160a01b03199092166001600160a01b03821617600255611f4e565b61176e905f51602061214b5f395f51905f526120ca565b60025461176e91906001600160a01b03808316911614611fc0575b5f6120ca565b600280546001600160a01b0319169055611fba565b9061176e91801580611ffc575b156120ca57600280546001600160a01b03191690556120ca565b506002546001600160a01b03838116911614611fe2565b65ffffffffffff811161202b5765ffffffffffff1690565b6306dfcc6560e41b5f52603060045260245260445ffd5b5f818152602081815260408083206001600160a01b038616845290915290205460ff166120c4575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19166001179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b5f818152602081815260408083206001600160a01b038616845290915290205460ff16156120c4575f818152602081815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a460019056fed8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63a26469706673582212206ee711909714dda8b552dd6c04c7d35ee34e956473496cdb4f59826fc172b0fe64736f6c634300081c0033",
}

// BridgeBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBotMetaData.ABI instead.
var BridgeBotABI = BridgeBotMetaData.ABI

// BridgeBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeBotMetaData.Bin instead.
var BridgeBotBin = BridgeBotMetaData.Bin

// DeployBridgeBot deploys a new Ethereum contract, binding an instance of BridgeBot to it.
func DeployBridgeBot(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _owner common.Address, _executor common.Address, _adminDelay *big.Int) (common.Address, *types.Transaction, *BridgeBot, error) {
	parsed, err := BridgeBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBotBin), backend, _bridge, _owner, _executor, _adminDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeBot{BridgeBotCaller: BridgeBotCaller{contract: contract}, BridgeBotTransactor: BridgeBotTransactor{contract: contract}, BridgeBotFilterer: BridgeBotFilterer{contract: contract}}, nil
}

// BridgeBot is an auto generated Go binding around an Ethereum contract.
type BridgeBot struct {
	BridgeBotCaller     // Read-only binding to the contract
	BridgeBotTransactor // Write-only binding to the contract
	BridgeBotFilterer   // Log filterer for contract events
}

// BridgeBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBotSession struct {
	Contract     *BridgeBot        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBotCallerSession struct {
	Contract *BridgeBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBotTransactorSession struct {
	Contract     *BridgeBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBotRaw struct {
	Contract *BridgeBot // Generic contract binding to access the raw methods on
}

// BridgeBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBotCallerRaw struct {
	Contract *BridgeBotCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBotTransactorRaw struct {
	Contract *BridgeBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBot creates a new instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBot(address common.Address, backend bind.ContractBackend) (*BridgeBot, error) {
	contract, err := bindBridgeBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBot{BridgeBotCaller: BridgeBotCaller{contract: contract}, BridgeBotTransactor: BridgeBotTransactor{contract: contract}, BridgeBotFilterer: BridgeBotFilterer{contract: contract}}, nil
}

// NewBridgeBotCaller creates a new read-only instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotCaller(address common.Address, caller bind.ContractCaller) (*BridgeBotCaller, error) {
	contract, err := bindBridgeBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBotCaller{contract: contract}, nil
}

// NewBridgeBotTransactor creates a new write-only instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBotTransactor, error) {
	contract, err := bindBridgeBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBotTransactor{contract: contract}, nil
}

// NewBridgeBotFilterer creates a new log filterer instance of BridgeBot, bound to a specific deployed contract.
func NewBridgeBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBotFilterer, error) {
	contract, err := bindBridgeBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBotFilterer{contract: contract}, nil
}

// bindBridgeBot binds a generic wrapper to an already deployed contract.
func bindBridgeBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBot *BridgeBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBot.Contract.BridgeBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBot *BridgeBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.Contract.BridgeBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBot *BridgeBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBot.Contract.BridgeBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBot *BridgeBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBot *BridgeBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBot *BridgeBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBot.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeBot.Contract.DEFAULTADMINROLE(&_BridgeBot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeBot.Contract.DEFAULTADMINROLE(&_BridgeBot.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotSession) EXECUTORROLE() ([32]byte, error) {
	return _BridgeBot.Contract.EXECUTORROLE(&_BridgeBot.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_BridgeBot *BridgeBotCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _BridgeBot.Contract.EXECUTORROLE(&_BridgeBot.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotSession) NATIVETOKEN() (common.Address, error) {
	return _BridgeBot.Contract.NATIVETOKEN(&_BridgeBot.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) NATIVETOKEN() (common.Address, error) {
	return _BridgeBot.Contract.NATIVETOKEN(&_BridgeBot.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotSession) Bridge() (common.Address, error) {
	return _BridgeBot.Contract.Bridge(&_BridgeBot.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) Bridge() (common.Address, error) {
	return _BridgeBot.Contract.Bridge(&_BridgeBot.CallOpts)
}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotCaller) BridgeConfigs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "bridgeConfigs", arg0)

	outstruct := new(struct {
		TokenAddress common.Address
		Recipient    common.Address
		ToChainID    *big.Int
		Interval     *big.Int
		LastExecuted *big.Int
		Enabled      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ToChainID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Interval = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastExecuted = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotSession) BridgeConfigs(arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	return _BridgeBot.Contract.BridgeConfigs(&_BridgeBot.CallOpts, arg0)
}

// BridgeConfigs is a free data retrieval call binding the contract method 0xd172f2f0.
//
// Solidity: function bridgeConfigs(uint256 ) view returns(address tokenAddress, address recipient, uint256 toChainID, uint256 interval, uint256 lastExecuted, bool enabled)
func (_BridgeBot *BridgeBotCallerSession) BridgeConfigs(arg0 *big.Int) (struct {
	TokenAddress common.Address
	Recipient    common.Address
	ToChainID    *big.Int
	Interval     *big.Int
	LastExecuted *big.Int
	Enabled      bool
}, error) {
	return _BridgeBot.Contract.BridgeConfigs(&_BridgeBot.CallOpts, arg0)
}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotCaller) CanExecuteBridge(opts *bind.CallOpts, configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "canExecuteBridge", configId)

	outstruct := new(struct {
		CanExecute        bool
		NextAvailableTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CanExecute = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NextAvailableTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotSession) CanExecuteBridge(configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	return _BridgeBot.Contract.CanExecuteBridge(&_BridgeBot.CallOpts, configId)
}

// CanExecuteBridge is a free data retrieval call binding the contract method 0xe1068d8d.
//
// Solidity: function canExecuteBridge(uint256 configId) view returns(bool canExecute, uint256 nextAvailableTime)
func (_BridgeBot *BridgeBotCallerSession) CanExecuteBridge(configId *big.Int) (struct {
	CanExecute        bool
	NextAvailableTime *big.Int
}, error) {
	return _BridgeBot.Contract.CanExecuteBridge(&_BridgeBot.CallOpts, configId)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BridgeBot *BridgeBotCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BridgeBot *BridgeBotSession) DefaultAdmin() (common.Address, error) {
	return _BridgeBot.Contract.DefaultAdmin(&_BridgeBot.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) DefaultAdmin() (common.Address, error) {
	return _BridgeBot.Contract.DefaultAdmin(&_BridgeBot.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BridgeBot *BridgeBotCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BridgeBot *BridgeBotSession) DefaultAdminDelay() (*big.Int, error) {
	return _BridgeBot.Contract.DefaultAdminDelay(&_BridgeBot.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_BridgeBot *BridgeBotCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BridgeBot.Contract.DefaultAdminDelay(&_BridgeBot.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BridgeBot *BridgeBotCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BridgeBot *BridgeBotSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BridgeBot.Contract.DefaultAdminDelayIncreaseWait(&_BridgeBot.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_BridgeBot *BridgeBotCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BridgeBot.Contract.DefaultAdminDelayIncreaseWait(&_BridgeBot.CallOpts)
}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotCaller) GetBridgeConfig(opts *bind.CallOpts, configId *big.Int) (BridgeBotBridgeConfig, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "getBridgeConfig", configId)

	if err != nil {
		return *new(BridgeBotBridgeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeBotBridgeConfig)).(*BridgeBotBridgeConfig)

	return out0, err

}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotSession) GetBridgeConfig(configId *big.Int) (BridgeBotBridgeConfig, error) {
	return _BridgeBot.Contract.GetBridgeConfig(&_BridgeBot.CallOpts, configId)
}

// GetBridgeConfig is a free data retrieval call binding the contract method 0x70d2ddf4.
//
// Solidity: function getBridgeConfig(uint256 configId) view returns((address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotCallerSession) GetBridgeConfig(configId *big.Int) (BridgeBotBridgeConfig, error) {
	return _BridgeBot.Contract.GetBridgeConfig(&_BridgeBot.CallOpts, configId)
}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotCaller) GetExecutableConfigs(opts *bind.CallOpts, maxConfigs *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "getExecutableConfigs", maxConfigs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotSession) GetExecutableConfigs(maxConfigs *big.Int) ([]*big.Int, error) {
	return _BridgeBot.Contract.GetExecutableConfigs(&_BridgeBot.CallOpts, maxConfigs)
}

// GetExecutableConfigs is a free data retrieval call binding the contract method 0xe9cf510c.
//
// Solidity: function getExecutableConfigs(uint256 maxConfigs) view returns(uint256[] executableConfigs)
func (_BridgeBot *BridgeBotCallerSession) GetExecutableConfigs(maxConfigs *big.Int) ([]*big.Int, error) {
	return _BridgeBot.Contract.GetExecutableConfigs(&_BridgeBot.CallOpts, maxConfigs)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBot *BridgeBotCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBot *BridgeBotSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeBot.Contract.GetRoleAdmin(&_BridgeBot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeBot *BridgeBotCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeBot.Contract.GetRoleAdmin(&_BridgeBot.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBot *BridgeBotCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBot *BridgeBotSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeBot.Contract.HasRole(&_BridgeBot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeBot *BridgeBotCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeBot.Contract.HasRole(&_BridgeBot.CallOpts, role, account)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotCaller) NextConfigId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "nextConfigId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotSession) NextConfigId() (*big.Int, error) {
	return _BridgeBot.Contract.NextConfigId(&_BridgeBot.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint256)
func (_BridgeBot *BridgeBotCallerSession) NextConfigId() (*big.Int, error) {
	return _BridgeBot.Contract.NextConfigId(&_BridgeBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotSession) Owner() (common.Address, error) {
	return _BridgeBot.Contract.Owner(&_BridgeBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBot *BridgeBotCallerSession) Owner() (common.Address, error) {
	return _BridgeBot.Contract.Owner(&_BridgeBot.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_BridgeBot *BridgeBotCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_BridgeBot *BridgeBotSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _BridgeBot.Contract.PendingDefaultAdmin(&_BridgeBot.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_BridgeBot *BridgeBotCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _BridgeBot.Contract.PendingDefaultAdmin(&_BridgeBot.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_BridgeBot *BridgeBotCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_BridgeBot *BridgeBotSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _BridgeBot.Contract.PendingDefaultAdminDelay(&_BridgeBot.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_BridgeBot *BridgeBotCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _BridgeBot.Contract.PendingDefaultAdminDelay(&_BridgeBot.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBot *BridgeBotCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeBot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBot *BridgeBotSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeBot.Contract.SupportsInterface(&_BridgeBot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeBot *BridgeBotCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeBot.Contract.SupportsInterface(&_BridgeBot.CallOpts, interfaceId)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BridgeBot.Contract.AcceptDefaultAdminTransfer(&_BridgeBot.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BridgeBot.Contract.AcceptDefaultAdminTransfer(&_BridgeBot.TransactOpts)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotTransactor) AddBridgeConfig(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "addBridgeConfig", tokenAddress, recipient, toChainID, interval)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotSession) AddBridgeConfig(tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.AddBridgeConfig(&_BridgeBot.TransactOpts, tokenAddress, recipient, toChainID, interval)
}

// AddBridgeConfig is a paid mutator transaction binding the contract method 0x4624e680.
//
// Solidity: function addBridgeConfig(address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns(uint256 configId)
func (_BridgeBot *BridgeBotTransactorSession) AddBridgeConfig(tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.AddBridgeConfig(&_BridgeBot.TransactOpts, tokenAddress, recipient, toChainID, interval)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BridgeBot *BridgeBotTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BridgeBot *BridgeBotSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.BeginDefaultAdminTransfer(&_BridgeBot.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_BridgeBot *BridgeBotTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.BeginDefaultAdminTransfer(&_BridgeBot.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BridgeBot.Contract.CancelDefaultAdminTransfer(&_BridgeBot.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_BridgeBot *BridgeBotTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BridgeBot.Contract.CancelDefaultAdminTransfer(&_BridgeBot.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BridgeBot *BridgeBotTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BridgeBot *BridgeBotSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ChangeDefaultAdminDelay(&_BridgeBot.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_BridgeBot *BridgeBotTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ChangeDefaultAdminDelay(&_BridgeBot.TransactOpts, newDelay)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) ExecuteBridge(opts *bind.TransactOpts, configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "executeBridge", configId, amount)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) ExecuteBridge(configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridge(&_BridgeBot.TransactOpts, configId, amount)
}

// ExecuteBridge is a paid mutator transaction binding the contract method 0xb1576074.
//
// Solidity: function executeBridge(uint256 configId, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) ExecuteBridge(configId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridge(&_BridgeBot.TransactOpts, configId, amount)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotTransactor) ExecuteBridgeBatch(opts *bind.TransactOpts, configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "executeBridgeBatch", configIds, amounts)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotSession) ExecuteBridgeBatch(configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridgeBatch(&_BridgeBot.TransactOpts, configIds, amounts)
}

// ExecuteBridgeBatch is a paid mutator transaction binding the contract method 0x76e95f16.
//
// Solidity: function executeBridgeBatch(uint256[] configIds, uint256[] amounts) returns()
func (_BridgeBot *BridgeBotTransactorSession) ExecuteBridgeBatch(configIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.ExecuteBridgeBatch(&_BridgeBot.TransactOpts, configIds, amounts)
}

// GrantExecutorRole is a paid mutator transaction binding the contract method 0x80a4d414.
//
// Solidity: function grantExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotTransactor) GrantExecutorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "grantExecutorRole", account)
}

// GrantExecutorRole is a paid mutator transaction binding the contract method 0x80a4d414.
//
// Solidity: function grantExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotSession) GrantExecutorRole(account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.GrantExecutorRole(&_BridgeBot.TransactOpts, account)
}

// GrantExecutorRole is a paid mutator transaction binding the contract method 0x80a4d414.
//
// Solidity: function grantExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotTransactorSession) GrantExecutorRole(account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.GrantExecutorRole(&_BridgeBot.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.GrantRole(&_BridgeBot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.GrantRole(&_BridgeBot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceRole(&_BridgeBot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceRole(&_BridgeBot.TransactOpts, role, account)
}

// RevokeExecutorRole is a paid mutator transaction binding the contract method 0x49a396b9.
//
// Solidity: function revokeExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotTransactor) RevokeExecutorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "revokeExecutorRole", account)
}

// RevokeExecutorRole is a paid mutator transaction binding the contract method 0x49a396b9.
//
// Solidity: function revokeExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotSession) RevokeExecutorRole(account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RevokeExecutorRole(&_BridgeBot.TransactOpts, account)
}

// RevokeExecutorRole is a paid mutator transaction binding the contract method 0x49a396b9.
//
// Solidity: function revokeExecutorRole(address account) returns()
func (_BridgeBot *BridgeBotTransactorSession) RevokeExecutorRole(account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RevokeExecutorRole(&_BridgeBot.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RevokeRole(&_BridgeBot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeBot *BridgeBotTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RevokeRole(&_BridgeBot.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BridgeBot *BridgeBotTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BridgeBot *BridgeBotSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BridgeBot.Contract.RollbackDefaultAdminDelay(&_BridgeBot.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_BridgeBot *BridgeBotTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BridgeBot.Contract.RollbackDefaultAdminDelay(&_BridgeBot.TransactOpts)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotTransactor) ToggleBridgeConfig(opts *bind.TransactOpts, configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "toggleBridgeConfig", configId, enabled)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotSession) ToggleBridgeConfig(configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.Contract.ToggleBridgeConfig(&_BridgeBot.TransactOpts, configId, enabled)
}

// ToggleBridgeConfig is a paid mutator transaction binding the contract method 0xbd5f0afb.
//
// Solidity: function toggleBridgeConfig(uint256 configId, bool enabled) returns()
func (_BridgeBot *BridgeBotTransactorSession) ToggleBridgeConfig(configId *big.Int, enabled bool) (*types.Transaction, error) {
	return _BridgeBot.Contract.ToggleBridgeConfig(&_BridgeBot.TransactOpts, configId, enabled)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotTransactor) UpdateBridgeConfig(opts *bind.TransactOpts, configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "updateBridgeConfig", configId, tokenAddress, recipient, toChainID, interval)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotSession) UpdateBridgeConfig(configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.UpdateBridgeConfig(&_BridgeBot.TransactOpts, configId, tokenAddress, recipient, toChainID, interval)
}

// UpdateBridgeConfig is a paid mutator transaction binding the contract method 0xed2859d9.
//
// Solidity: function updateBridgeConfig(uint256 configId, address tokenAddress, address recipient, uint256 toChainID, uint256 interval) returns()
func (_BridgeBot *BridgeBotTransactorSession) UpdateBridgeConfig(configId *big.Int, tokenAddress common.Address, recipient common.Address, toChainID *big.Int, interval *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.UpdateBridgeConfig(&_BridgeBot.TransactOpts, configId, tokenAddress, recipient, toChainID, interval)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawAllNative(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawAllNative", to)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotSession) WithdrawAllNative(to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllNative(&_BridgeBot.TransactOpts, to)
}

// WithdrawAllNative is a paid mutator transaction binding the contract method 0xd9f66db1.
//
// Solidity: function withdrawAllNative(address to) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawAllNative(to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllNative(&_BridgeBot.TransactOpts, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawAllTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawAllTokens", token, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotSession) WithdrawAllTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllTokens(&_BridgeBot.TransactOpts, token, to)
}

// WithdrawAllTokens is a paid mutator transaction binding the contract method 0x50f760e9.
//
// Solidity: function withdrawAllTokens(address token, address to) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawAllTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawAllTokens(&_BridgeBot.TransactOpts, token, to)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawNative(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawNative", to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawNative(&_BridgeBot.TransactOpts, to, amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawNative(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawNative(&_BridgeBot.TransactOpts, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "withdrawToken", token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawToken(&_BridgeBot.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_BridgeBot *BridgeBotTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeBot.Contract.WithdrawToken(&_BridgeBot.TransactOpts, token, to, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.Contract.Fallback(&_BridgeBot.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BridgeBot *BridgeBotTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BridgeBot.Contract.Fallback(&_BridgeBot.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotSession) Receive() (*types.Transaction, error) {
	return _BridgeBot.Contract.Receive(&_BridgeBot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeBot *BridgeBotTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeBot.Contract.Receive(&_BridgeBot.TransactOpts)
}

// BridgeBotBridgeConfigAddedIterator is returned from FilterBridgeConfigAdded and is used to iterate over the raw logs and unpacked data for BridgeConfigAdded events raised by the BridgeBot contract.
type BridgeBotBridgeConfigAddedIterator struct {
	Event *BridgeBotBridgeConfigAdded // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigAdded)
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
		it.Event = new(BridgeBotBridgeConfigAdded)
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
func (it *BridgeBotBridgeConfigAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigAdded represents a BridgeConfigAdded event raised by the BridgeBot contract.
type BridgeBotBridgeConfigAdded struct {
	ConfigId *big.Int
	Config   BridgeBotBridgeConfig
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigAdded is a free log retrieval operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigAdded(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigAddedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigAdded", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigAddedIterator{contract: _BridgeBot.contract, event: "BridgeConfigAdded", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigAdded is a free log subscription operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigAdded(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigAdded, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigAdded", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigAdded)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigAdded", log); err != nil {
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

// ParseBridgeConfigAdded is a log parse operation binding the contract event 0x607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123.
//
// Solidity: event BridgeConfigAdded(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigAdded(log types.Log) (*BridgeBotBridgeConfigAdded, error) {
	event := new(BridgeBotBridgeConfigAdded)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeConfigToggledIterator is returned from FilterBridgeConfigToggled and is used to iterate over the raw logs and unpacked data for BridgeConfigToggled events raised by the BridgeBot contract.
type BridgeBotBridgeConfigToggledIterator struct {
	Event *BridgeBotBridgeConfigToggled // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigToggled)
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
		it.Event = new(BridgeBotBridgeConfigToggled)
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
func (it *BridgeBotBridgeConfigToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigToggled represents a BridgeConfigToggled event raised by the BridgeBot contract.
type BridgeBotBridgeConfigToggled struct {
	ConfigId *big.Int
	Enabled  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigToggled is a free log retrieval operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigToggled(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigToggledIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigToggled", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigToggledIterator{contract: _BridgeBot.contract, event: "BridgeConfigToggled", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigToggled is a free log subscription operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigToggled(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigToggled, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigToggled", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigToggled)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigToggled", log); err != nil {
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

// ParseBridgeConfigToggled is a log parse operation binding the contract event 0x82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0.
//
// Solidity: event BridgeConfigToggled(uint256 indexed configId, bool enabled)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigToggled(log types.Log) (*BridgeBotBridgeConfigToggled, error) {
	event := new(BridgeBotBridgeConfigToggled)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeConfigUpdatedIterator is returned from FilterBridgeConfigUpdated and is used to iterate over the raw logs and unpacked data for BridgeConfigUpdated events raised by the BridgeBot contract.
type BridgeBotBridgeConfigUpdatedIterator struct {
	Event *BridgeBotBridgeConfigUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeConfigUpdated)
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
		it.Event = new(BridgeBotBridgeConfigUpdated)
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
func (it *BridgeBotBridgeConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeConfigUpdated represents a BridgeConfigUpdated event raised by the BridgeBot contract.
type BridgeBotBridgeConfigUpdated struct {
	ConfigId *big.Int
	Config   BridgeBotBridgeConfig
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeConfigUpdated is a free log retrieval operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeConfigUpdated(opts *bind.FilterOpts, configId []*big.Int) (*BridgeBotBridgeConfigUpdatedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeConfigUpdated", configIdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeConfigUpdatedIterator{contract: _BridgeBot.contract, event: "BridgeConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeConfigUpdated is a free log subscription operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeConfigUpdated, configId []*big.Int) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeConfigUpdated", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeConfigUpdated)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigUpdated", log); err != nil {
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

// ParseBridgeConfigUpdated is a log parse operation binding the contract event 0xe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b765.
//
// Solidity: event BridgeConfigUpdated(uint256 indexed configId, (address,address,uint256,uint256,uint256,bool) config)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeConfigUpdated(log types.Log) (*BridgeBotBridgeConfigUpdated, error) {
	event := new(BridgeBotBridgeConfigUpdated)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotBridgeExecutedIterator is returned from FilterBridgeExecuted and is used to iterate over the raw logs and unpacked data for BridgeExecuted events raised by the BridgeBot contract.
type BridgeBotBridgeExecutedIterator struct {
	Event *BridgeBotBridgeExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeBotBridgeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotBridgeExecuted)
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
		it.Event = new(BridgeBotBridgeExecuted)
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
func (it *BridgeBotBridgeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotBridgeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotBridgeExecuted represents a BridgeExecuted event raised by the BridgeBot contract.
type BridgeBotBridgeExecuted struct {
	ConfigId     *big.Int
	TokenAddress common.Address
	Amount       *big.Int
	Recipient    common.Address
	ToChainID    *big.Int
	Executor     common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgeExecuted is a free log retrieval operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) FilterBridgeExecuted(opts *bind.FilterOpts, configId []*big.Int, tokenAddress []common.Address, recipient []common.Address) (*BridgeBotBridgeExecutedIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "BridgeExecuted", configIdRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotBridgeExecutedIterator{contract: _BridgeBot.contract, event: "BridgeExecuted", logs: logs, sub: sub}, nil
}

// WatchBridgeExecuted is a free log subscription operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) WatchBridgeExecuted(opts *bind.WatchOpts, sink chan<- *BridgeBotBridgeExecuted, configId []*big.Int, tokenAddress []common.Address, recipient []common.Address) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "BridgeExecuted", configIdRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotBridgeExecuted)
				if err := _BridgeBot.contract.UnpackLog(event, "BridgeExecuted", log); err != nil {
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

// ParseBridgeExecuted is a log parse operation binding the contract event 0x96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee.
//
// Solidity: event BridgeExecuted(uint256 indexed configId, address indexed tokenAddress, uint256 amount, address indexed recipient, uint256 toChainID, address executor, uint256 timestamp)
func (_BridgeBot *BridgeBotFilterer) ParseBridgeExecuted(log types.Log) (*BridgeBotBridgeExecuted, error) {
	event := new(BridgeBotBridgeExecuted)
	if err := _BridgeBot.contract.UnpackLog(event, "BridgeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the BridgeBot contract.
type BridgeBotDefaultAdminDelayChangeCanceledIterator struct {
	Event *BridgeBotDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *BridgeBotDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotDefaultAdminDelayChangeCanceled)
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
		it.Event = new(BridgeBotDefaultAdminDelayChangeCanceled)
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
func (it *BridgeBotDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the BridgeBot contract.
type BridgeBotDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_BridgeBot *BridgeBotFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BridgeBotDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BridgeBotDefaultAdminDelayChangeCanceledIterator{contract: _BridgeBot.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_BridgeBot *BridgeBotFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BridgeBotDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotDefaultAdminDelayChangeCanceled)
				if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BridgeBotDefaultAdminDelayChangeCanceled, error) {
	event := new(BridgeBotDefaultAdminDelayChangeCanceled)
	if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the BridgeBot contract.
type BridgeBotDefaultAdminDelayChangeScheduledIterator struct {
	Event *BridgeBotDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *BridgeBotDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotDefaultAdminDelayChangeScheduled)
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
		it.Event = new(BridgeBotDefaultAdminDelayChangeScheduled)
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
func (it *BridgeBotDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the BridgeBot contract.
type BridgeBotDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_BridgeBot *BridgeBotFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BridgeBotDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BridgeBotDefaultAdminDelayChangeScheduledIterator{contract: _BridgeBot.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_BridgeBot *BridgeBotFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BridgeBotDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotDefaultAdminDelayChangeScheduled)
				if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BridgeBotDefaultAdminDelayChangeScheduled, error) {
	event := new(BridgeBotDefaultAdminDelayChangeScheduled)
	if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the BridgeBot contract.
type BridgeBotDefaultAdminTransferCanceledIterator struct {
	Event *BridgeBotDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *BridgeBotDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotDefaultAdminTransferCanceled)
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
		it.Event = new(BridgeBotDefaultAdminTransferCanceled)
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
func (it *BridgeBotDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the BridgeBot contract.
type BridgeBotDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_BridgeBot *BridgeBotFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BridgeBotDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BridgeBotDefaultAdminTransferCanceledIterator{contract: _BridgeBot.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_BridgeBot *BridgeBotFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BridgeBotDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotDefaultAdminTransferCanceled)
				if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BridgeBotDefaultAdminTransferCanceled, error) {
	event := new(BridgeBotDefaultAdminTransferCanceled)
	if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the BridgeBot contract.
type BridgeBotDefaultAdminTransferScheduledIterator struct {
	Event *BridgeBotDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *BridgeBotDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotDefaultAdminTransferScheduled)
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
		it.Event = new(BridgeBotDefaultAdminTransferScheduled)
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
func (it *BridgeBotDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the BridgeBot contract.
type BridgeBotDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_BridgeBot *BridgeBotFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BridgeBotDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotDefaultAdminTransferScheduledIterator{contract: _BridgeBot.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_BridgeBot *BridgeBotFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BridgeBotDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotDefaultAdminTransferScheduled)
				if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BridgeBotDefaultAdminTransferScheduled, error) {
	event := new(BridgeBotDefaultAdminTransferScheduled)
	if err := _BridgeBot.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotNativeWithdrawnIterator is returned from FilterNativeWithdrawn and is used to iterate over the raw logs and unpacked data for NativeWithdrawn events raised by the BridgeBot contract.
type BridgeBotNativeWithdrawnIterator struct {
	Event *BridgeBotNativeWithdrawn // Event containing the contract specifics and raw log

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
func (it *BridgeBotNativeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotNativeWithdrawn)
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
		it.Event = new(BridgeBotNativeWithdrawn)
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
func (it *BridgeBotNativeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotNativeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotNativeWithdrawn represents a NativeWithdrawn event raised by the BridgeBot contract.
type BridgeBotNativeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeWithdrawn is a free log retrieval operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) FilterNativeWithdrawn(opts *bind.FilterOpts, to []common.Address) (*BridgeBotNativeWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "NativeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotNativeWithdrawnIterator{contract: _BridgeBot.contract, event: "NativeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNativeWithdrawn is a free log subscription operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) WatchNativeWithdrawn(opts *bind.WatchOpts, sink chan<- *BridgeBotNativeWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "NativeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotNativeWithdrawn)
				if err := _BridgeBot.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
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

// ParseNativeWithdrawn is a log parse operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) ParseNativeWithdrawn(log types.Log) (*BridgeBotNativeWithdrawn, error) {
	event := new(BridgeBotNativeWithdrawn)
	if err := _BridgeBot.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeBot contract.
type BridgeBotRoleAdminChangedIterator struct {
	Event *BridgeBotRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeBotRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotRoleAdminChanged)
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
		it.Event = new(BridgeBotRoleAdminChanged)
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
func (it *BridgeBotRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeBot contract.
type BridgeBotRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeBot *BridgeBotFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeBotRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotRoleAdminChangedIterator{contract: _BridgeBot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeBot *BridgeBotFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeBotRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotRoleAdminChanged)
				if err := _BridgeBot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeBotRoleAdminChanged, error) {
	event := new(BridgeBotRoleAdminChanged)
	if err := _BridgeBot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeBot contract.
type BridgeBotRoleGrantedIterator struct {
	Event *BridgeBotRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeBotRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotRoleGranted)
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
		it.Event = new(BridgeBotRoleGranted)
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
func (it *BridgeBotRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotRoleGranted represents a RoleGranted event raised by the BridgeBot contract.
type BridgeBotRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBot *BridgeBotFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeBotRoleGrantedIterator, error) {

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

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotRoleGrantedIterator{contract: _BridgeBot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBot *BridgeBotFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeBotRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotRoleGranted)
				if err := _BridgeBot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseRoleGranted(log types.Log) (*BridgeBotRoleGranted, error) {
	event := new(BridgeBotRoleGranted)
	if err := _BridgeBot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeBot contract.
type BridgeBotRoleRevokedIterator struct {
	Event *BridgeBotRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeBotRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotRoleRevoked)
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
		it.Event = new(BridgeBotRoleRevoked)
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
func (it *BridgeBotRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotRoleRevoked represents a RoleRevoked event raised by the BridgeBot contract.
type BridgeBotRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBot *BridgeBotFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeBotRoleRevokedIterator, error) {

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

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotRoleRevokedIterator{contract: _BridgeBot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeBot *BridgeBotFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeBotRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotRoleRevoked)
				if err := _BridgeBot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BridgeBot *BridgeBotFilterer) ParseRoleRevoked(log types.Log) (*BridgeBotRoleRevoked, error) {
	event := new(BridgeBotRoleRevoked)
	if err := _BridgeBot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBotTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the BridgeBot contract.
type BridgeBotTokenWithdrawnIterator struct {
	Event *BridgeBotTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *BridgeBotTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotTokenWithdrawn)
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
		it.Event = new(BridgeBotTokenWithdrawn)
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
func (it *BridgeBotTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotTokenWithdrawn represents a TokenWithdrawn event raised by the BridgeBot contract.
type BridgeBotTokenWithdrawn struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeBotTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotTokenWithdrawnIterator{contract: _BridgeBot.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BridgeBotTokenWithdrawn, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotTokenWithdrawn)
				if err := _BridgeBot.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed to, uint256 amount)
func (_BridgeBot *BridgeBotFilterer) ParseTokenWithdrawn(log types.Log) (*BridgeBotTokenWithdrawn, error) {
	event := new(BridgeBotTokenWithdrawn)
	if err := _BridgeBot.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
