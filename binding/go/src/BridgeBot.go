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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXECUTOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBridgeConfig\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBaseBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canExecuteBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"canExecute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBridgeBatch\",\"inputs\":[{\"name\":\"configIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutableConfigs\",\"inputs\":[{\"name\":\"maxConfigs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"executableConfigs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantExecutorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextConfigId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeExecutorRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toggleBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeConfigAdded\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigToggled\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigUpdated\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeExecuted\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BridgeBotBridgeFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotInsufficientBalance\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BridgeBotNotTimeYet\",\"inputs\":[{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a03461012f57601f611ef838819003918201601f19168301916001600160401b038311848410176101335780849260609460405283398101031261012f5761004781610147565b61005f604061005860208501610147565b9301610147565b906001600160a01b038316801561011c575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a360016002556001600160a01b0316801561010d576001600160a01b0382161561010d576100ed926100e7916080526100e78161015b565b506101d1565b50604051611c33908161026582396080518181816103ff015261171f0152f35b630508665f60e41b5f5260045ffd5b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361012f57565b6001600160a01b0381165f9081525f516020611eb85f395f51905f52602052604090205460ff166101cc576001600160a01b03165f8181525f516020611eb85f395f51905f5260205260408120805460ff191660011790553391905f516020611e985f395f51905f528180a4600190565b505f90565b6001600160a01b0381165f9081525f516020611ed85f395f51905f52602052604090205460ff166101cc576001600160a01b03165f8181525f516020611ed85f395f51905f5260205260408120805460ff191660011790553391907fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63905f516020611e985f395f51905f529080a460019056fe60806040526004361015610010575b005b5f3560e01c806301e3366714610e8357806301ffc9a714610e2d57806307b18bde14610d7857806307bd026514610d51578063248a9ca314610d1e5780632f2ff15d14610ce057806331f7d96414610cc557806336568abe14610c815780634624e68014610b2b57806349a396b914610ac257806350f760e9146109ce57806370d2ddf4146108ec578063715018a61461089557806376e95f161461077a57806380a4d414146107005780638da5cb5b146106d957806391d148541461069057806399d726c714610673578063a217fddf14610659578063b157607414610623578063bd5f0afb14610591578063d172f2f01461051b578063d547741f146104d8578063d9f66db11461045a578063e1068d8d1461042e578063e78cea92146103ea578063e9cf510c146102f6578063ed2859d9146101dc5763f2fde38b0361000e57346101d85760203660031901126101d85761016c610f07565b610174611284565b6001600160a01b031680156101c5575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b346101d85760a03660031901126101d8576004356101f8610f1d565b6044356001600160a01b038116908190036101d85760843591610219611284565b5f8481526003602052604090205461023b906001600160a01b031615156110c7565b6001600160a01b038116156102e75781156102e75782156102d8575f8481526003602081905260409182902080546001600160a01b03949094166001600160a01b03199485161781556001810180549094169490941790925560643560028401559082019290925590517fe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b7659181906102d39082611055565b0390a2005b6319e9855160e11b5f5260045ffd5b630508665f60e41b5f5260045ffd5b346101d85760203660031901126101d8576004356103138161123e565b905f5f91600454925b838110806103e1575b15610368576103338161113f565b50610347575b61034290611033565b61031c565b91610360818461035a6103429489611270565b52611033565b929050610339565b84836103738161123e565b915f5b8281106103c257836040518091602082016020835281518091526020604084019201905f5b8181106103a9575050500390f35b825184528594506020938401939092019160010161039b565b806103cf60019284611270565b516103da8287611270565b5201610376565b50818310610325565b346101d8575f3660031901126101d8576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101d85760203660031901126101d857604061044c60043561113f565b825191151582526020820152f35b346101d85760203660031901126101d857610473610f07565b61047b611284565b6001600160a01b031680156102e757478061049257005b6020816104cf5f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af16104c9610fb6565b50610ff5565b604051908152a2005b346101d85760403660031901126101d85761000e6004356104f7610f1d565b90610516610511825f526001602052600160405f20015490565b6113d0565b6115bf565b346101d85760203660031901126101d8576004355f52600360205260c060405f2060018060a01b038154169060018060a01b0360018201541690600281015460038201549060ff6005600485015494015416936040519586526020860152604085015260608401526080830152151560a0820152f35b346101d85760403660031901126101d857600435602435801515908181036101d8577f82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e0916104cf6020926105e3611284565b5f868152600385526040902054610604906001600160a01b031615156110c7565b855f5260038452600560405f20019060ff801983541691151516179055565b346101d85760403660031901126101d85761063c611387565b610644611643565b6106526024356004356116c0565b6001600255005b346101d8575f3660031901126101d85760206040515f8152f35b346101d8575f3660031901126101d8576020600454604051908152f35b346101d85760403660031901126101d8576106a9610f1d565b6004355f52600160205260405f209060018060a01b03165f52602052602060ff60405f2054166040519015158152f35b346101d8575f3660031901126101d8575f546040516001600160a01b039091168152602090f35b346101d85760203660031901126101d857610719610f07565b610721611338565b6001600160a01b038116156102e7575f516020611bbe5f395f51905f525f5260016020527fc9ee83ecf8e561e5df8e9e3a8d108a689296832fb5d541ca3c450b10eaabf8e25461000e9190610775906113d0565b61140a565b346101d85760403660031901126101d85760043567ffffffffffffffff81116101d8576107ab903690600401610f33565b9060243567ffffffffffffffff81116101d8576107cc903690600401610f33565b92906107d6611387565b6107de611643565b838203610858575f5b8281106107f5576001600255005b8061080c61080660019386886110a3565b3561113f565b5080610844575b61081e575b016107e7565b61083f61082c8286886110a3565b356108388389876110a3565b35906116c0565b610818565b506108508187856110a3565b351515610813565b60405162461bcd60e51b8152602060048201526015602482015274082e4e4c2f240d8cadccee8d040dad2e6dac2e8c6d605b1b6044820152606490fd5b346101d8575f3660031901126101d8576108ad611284565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346101d85760203660031901126101d8575f60a060405161090c81610f64565b82815282602082015282604082015282606082015282608082015201526004355f52600360205260c060405f2060405161094581610f64565b60018060a01b038254169182825260018060a01b03600182015416602083019081526002820154604084019081526003830154916060850192835260a060ff6005600487015496608089019788520154169501941515855260405195865260018060a01b03905116602086015251604085015251606084015251608083015251151560a0820152f35b346101d85760403660031901126101d8576109e7610f07565b6109ef610f1d565b906109f8611284565b6001600160a01b031680156102e7576001600160a01b0382169182156102e7576040516370a0823160e01b815230600482015290602082602481865afa918215610ab7575f92610a83575b5081610a4b57005b81610a7a7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e562093602093866112aa565b604051908152a3005b9091506020813d602011610aaf575b81610a9f60209383610f94565b810103126101d857519084610a43565b3d9150610a92565b6040513d5f823e3d90fd5b346101d85760203660031901126101d85761000e610ade610f07565b610ae6611338565b5f516020611bbe5f395f51905f525f5260016020527fc9ee83ecf8e561e5df8e9e3a8d108a689296832fb5d541ca3c450b10eaabf8e254610b26906113d0565b61152e565b346101d85760803660031901126101d857610b44610f07565b610b4c610f1d565b9060643590610b59611284565b6001600160a01b03169182156102e7576001600160a01b031680156102e75781156102d8576020926005610c3a9260045494610b9486611033565b60045560405193610ba485610f64565b84528684019182526044356040808601918252606086019283525f60808701818152600160a089018181528b84526003808e5294909320985189546001600160a01b039182166001600160a01b0319918216178b559751918a0180549290911691909716179095559151600287015591519185019190915590516004840155519101805460ff191691151560ff16919091179055565b805f5260038252807f607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123610c7660405f2060405191829182611055565b0390a2604051908152f35b346101d85760403660031901126101d857610c9a610f1d565b336001600160a01b03821603610cb65761000e906004356115bf565b63334bd91960e11b5f5260045ffd5b346101d8575f3660031901126101d857602060405160018152f35b346101d85760403660031901126101d85761000e600435610cff610f1d565b90610d19610511825f526001602052600160405f20015490565b6114a2565b346101d85760203660031901126101d8576020610d496004355f526001602052600160405f20015490565b604051908152f35b346101d8575f3660031901126101d85760206040515f516020611bbe5f395f51905f528152f35b346101d85760403660031901126101d857610d91610f07565b60243590610d9d611284565b6001600160a01b03169081156102e75780156102d857804710610df1576020816104cf5f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af16104c9610fb6565b60405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606490fd5b346101d85760203660031901126101d85760043563ffffffff60e01b81168091036101d857602090637965db0b60e01b8114908115610e72575b506040519015158152f35b6301ffc9a760e01b14905082610e67565b346101d85760603660031901126101d857610e9c610f07565b610ea4610f1d565b9060443590610eb1611284565b6001600160a01b03169081156102e7576001600160a01b0383169283156102e75781156102d85781610a7a7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e562093602093866112aa565b600435906001600160a01b03821682036101d857565b602435906001600160a01b03821682036101d857565b9181601f840112156101d85782359167ffffffffffffffff83116101d8576020808501948460051b0101116101d857565b60c0810190811067ffffffffffffffff821117610f8057604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff821117610f8057604052565b3d15610ff0573d9067ffffffffffffffff8211610f805760405191610fe5601f8201601f191660200184610f94565b82523d5f602084013e565b606090565b15610ffc57565b60405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606490fd5b5f1981146110415760010190565b634e487b7160e01b5f52601160045260245ffd5b81546001600160a01b039081168252600183015416602082015260028201546040820152600382015460608201526004820154608082015260059091015460ff16151560a082015260c00190565b91908110156110b35760051b0190565b634e487b7160e01b5f52603260045260245ffd5b156110ce57565b60405162461bcd60e51b8152602060048201526011602482015270436f6e666967206e6f742065786973747360781b6044820152606490fd5b8115611111570690565b634e487b7160e01b5f52601260045260245ffd5b9190820391821161104157565b9190820180921161104157565b5f52600360205260405f209060405161115781610f64565b82546001600160a01b03908116808352600185015490911660208301526002840154604083015260038401546060830181815260048601546080850181815260059097015460ff1615801560a090960186905291949293909261121d575b506112125715611207576111d56111cf6111e89242611107565b42611125565b93516111e2835182611107565b90611125565b80931192835f146111f95750505f90565b611204915190611132565b90565b505090506001905f90565b50505090505f905f90565b9050155f6111b5565b67ffffffffffffffff8111610f805760051b60200190565b9061124882611226565b6112556040519182610f94565b8281528092611266601f1991611226565b0190602036910137565b80518210156110b35760209160051b010190565b5f546001600160a01b0316330361129757565b63118cdaa760e01b5f523360045260245ffd5b60405163a9059cbb60e01b60208281019182526001600160a01b03909416602483015260448083019590955293815290925f916112e8606482610f94565b519082855af115610ab7575f513d61132f57506001600160a01b0381163b155b61130f5750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415611308565b335f9081527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49602052604090205460ff161561137057565b63e2517d3f60e01b5f52336004525f60245260445ffd5b335f9081525f516020611bde5f395f51905f52602052604090205460ff16156113ac57565b63e2517d3f60e01b5f52336004525f516020611bbe5f395f51905f5260245260445ffd5b5f81815260016020908152604080832033845290915290205460ff16156113f45750565b63e2517d3f60e01b5f523360045260245260445ffd5b6001600160a01b0381165f9081525f516020611bde5f395f51905f52602052604090205460ff1661149d576001600160a01b03165f8181525f516020611bde5f395f51905f5260205260408120805460ff191660011790553391905f516020611bbe5f395f51905f52907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b505f90565b5f8181526001602090815260408083206001600160a01b038616845290915290205460ff16611528575f8181526001602081815260408084206001600160a01b0396909616808552959091528220805460ff19169091179055339291907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50505f90565b6001600160a01b0381165f9081525f516020611bde5f395f51905f52602052604090205460ff161561149d576001600160a01b03165f8181525f516020611bde5f395f51905f5260205260408120805460ff191690553391905f516020611bbe5f395f51905f52907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b5f8181526001602090815260408083206001600160a01b038616845290915290205460ff1615611528575f8181526001602090815260408083206001600160a01b0395909516808452949091528120805460ff19169055339291907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b60028054146116525760028055565b633ee5aeb560e01b5f5260045ffd5b908160209103126101d8575180151581036101d85790565b9081526001600160a01b039182166020820152911660408201526060810191909152608081019190915260a081019190915260e060c082018190525f908201526101000190565b81156102d857805f52600360205260405f209060ff60058301541680611baa575b15611b6e576116ef8161113f565b9015611b5c575081546001600160a01b03166001811493908415611af95747945b604051637838174760e11b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa918215610ab7575f92611ab5575b506002870197606089546064604051809681936337dba1f760e21b835260048301528960248301528a604483015260018060a01b03165afa908115610ab7575f5f945f93611a71575b506117bb836117b6878b611132565b611132565b91828110611a5b57508710611a1f5786941561194e575b88549495506020946001600160a01b0316600181036118dd57508954895460018b0154604051632fe9316f60e11b8152988997889687956118299591949293926001600160a01b0390811692169060048801611679565b03925af1908115610ab7575f916118ae575b505b1561189f5760807f96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee9142600486015560018060a01b03855416946001808060a01b039101541695546040519182526020820152336040820152426060820152a4565b6307c4732760e51b5f5260045ffd5b6118d0915060203d6020116118d6575b6118c88183610f94565b810190611661565b5f61183b565b503d6118be565b8a5460018b0154604051632fe9316f60e11b815298899788965f965087956119179591949293926001600160a01b03169160048801611679565b03925af1908115610ab7575f9161192f575b5061183d565b611948915060203d6020116118d6576118c88183610f94565b5f611929565b604051636eb1769f60e11b8152306004820152602481018490529450602085604481895afa8015610ab75787955f916119e9575b508111156117d257935060205f9560446040518098819363095ea7b360e01b8352876004840152811960248401525af1938415610ab75786956020956119cc575b508594506117d2565b6119e290863d88116118d6576118c88183610f94565b505f6119c3565b9550506020853d602011611a17575b81611a0560209383610f94565b810103126101d8578087955190611982565b3d91506119f8565b60405162461bcd60e51b8152602060048201526014602482015273416d6f756e742062656c6f77206d696e696d756d60601b6044820152606490fd5b8263203b880360e11b5f5260045260245260445ffd5b92505092506060813d606011611aad575b81611a8f60609383610f94565b810103126101d857805160208201516040909201519193905f6117a7565b3d9150611a82565b9091506020813d602011611af1575b81611ad160209383610f94565b810103126101d857516001600160a01b03811681036101d857905f61175e565b3d9150611ac4565b6040516370a0823160e01b8152306004820152602081602481855afa908115610ab7575f91611b2a575b5094611710565b90506020813d602011611b54575b81611b4560209383610f94565b810103126101d857515f611b23565b3d9150611b38565b6357c8d07f60e11b5f5260045260245ffd5b60405162461bcd60e51b8152602060048201526014602482015273436f6e666967206e6f7420617661696c61626c6560601b6044820152606490fd5b5081546001600160a01b031615156116e156fed8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63c9ee83ecf8e561e5df8e9e3a8d108a689296832fb5d541ca3c450b10eaabf8e1a264697066735822122061a195821bff59740401063cb63e50b0120f83c5a7efc9a09da76937773c62de64736f6c634300081c00332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0da6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49c9ee83ecf8e561e5df8e9e3a8d108a689296832fb5d541ca3c450b10eaabf8e1",
}

// BridgeBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBotMetaData.ABI instead.
var BridgeBotABI = BridgeBotMetaData.ABI

// BridgeBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeBotMetaData.Bin instead.
var BridgeBotBin = BridgeBotMetaData.Bin

// DeployBridgeBot deploys a new Ethereum contract, binding an instance of BridgeBot to it.
func DeployBridgeBot(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _owner common.Address, _executor common.Address) (common.Address, *types.Transaction, *BridgeBot, error) {
	parsed, err := BridgeBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBotBin), backend, _bridge, _owner, _executor)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceOwnership(&_BridgeBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBot *BridgeBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceOwnership(&_BridgeBot.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeBot *BridgeBotTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeBot *BridgeBotSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceRole(&_BridgeBot.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeBot *BridgeBotTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.RenounceRole(&_BridgeBot.TransactOpts, role, callerConfirmation)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.TransferOwnership(&_BridgeBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBot *BridgeBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBot.Contract.TransferOwnership(&_BridgeBot.TransactOpts, newOwner)
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

// BridgeBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeBot contract.
type BridgeBotOwnershipTransferredIterator struct {
	Event *BridgeBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBotOwnershipTransferred)
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
		it.Event = new(BridgeBotOwnershipTransferred)
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
func (it *BridgeBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBotOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeBot contract.
type BridgeBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBotOwnershipTransferredIterator{contract: _BridgeBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBotOwnershipTransferred)
				if err := _BridgeBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBot *BridgeBotFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeBotOwnershipTransferred, error) {
	event := new(BridgeBotOwnershipTransferred)
	if err := _BridgeBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
