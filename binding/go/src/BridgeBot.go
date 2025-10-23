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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBridgeConfig\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBaseBridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeConfigs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canExecuteBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"canExecute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBridge\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeBridgeBatch\",\"inputs\":[{\"name\":\"configIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutableConfigs\",\"inputs\":[{\"name\":\"maxConfigs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"executableConfigs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextConfigId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toggleBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBridgeConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAllTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNative\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeConfigAdded\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigToggled\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeConfigUpdated\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBridgeBot.BridgeConfig\",\"components\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeExecuted\",\"inputs\":[{\"name\":\"configId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toChainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BridgeBotBridgeFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotCanNotZeroValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BridgeBotInsufficientBalance\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BridgeBotNotTimeYet\",\"inputs\":[{\"name\":\"nextAvailableTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a0346100ff57601f6116db38819003918201601f19168301916001600160401b038311848410176101035780849260409485528339810103126100ff5761004681610117565b906001600160a01b039061005c90602001610117565b169081156100ec575f80546001600160a01b031981168417825560405193916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3600180556001600160a01b031680156100dd576080526115af908161012c823960805181818161039401526110db0152f35b630508665f60e41b5f5260045ffd5b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100ff5756fe60806040526004361015610010575b005b5f3560e01c806301e3366714610b4857806307b18bde14610a9357806331f7d96414610a785780634624e6801461092557806350f760e91461083157806370d2ddf41461074f578063715018a6146106f857806376e95f16146105e65780638da5cb5b146105bf57806399d726c7146105a2578063b157607414610575578063bd5f0afb146104e3578063d172f2f01461046d578063d9f66db1146103ef578063e1068d8d146103c3578063e78cea921461037f578063e9cf510c1461028b578063ed2859d91461016e5763f2fde38b0361000e573461016a57602036600319011261016a576100fe610bcc565b610106610f49565b6001600160a01b03168015610157575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b3461016a5760a036600319011261016a5760043561018a610be2565b6044356001600160a01b0381169081900361016a57608435916101ab610f49565b5f848152600260205260409020546101cd906001600160a01b03161515610d8c565b6001600160a01b0381161561027c57811561027c57821561026d575f8481526002602081905260409182902080546001600160a01b03949094166001600160a01b03199485161781556001810180549094169490941790925560643591830191909155600382019290925590517fe983b12d31b0781e2aa16f6b2bbac09cf920528457b5fe0c4a48a4b7b639b7659181906102689082610d1a565b0390a2005b6319e9855160e11b5f5260045ffd5b630508665f60e41b5f5260045ffd5b3461016a57602036600319011261016a576004356102a881610f03565b905f5f91600354925b83811080610376575b156102fd576102c881610e04565b506102dc575b6102d790610cf8565b6102b1565b916102f581846102ef6102d79489610f35565b52610cf8565b9290506102ce565b848361030881610f03565b915f5b82811061035757836040518091602082016020835281518091526020604084019201905f5b81811061033e575050500390f35b8251845285945060209384019390920191600101610330565b8061036460019284610f35565b5161036f8287610f35565b520161030b565b508183106102ba565b3461016a575f36600319011261016a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461016a57602036600319011261016a5760406103e1600435610e04565b825191151582526020820152f35b3461016a57602036600319011261016a57610408610bcc565b610410610f49565b6001600160a01b0316801561027c57478061042757005b6020816104645f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af161045e610c7b565b50610cba565b604051908152a2005b3461016a57602036600319011261016a576004355f52600260205260c060405f2060018060a01b038154169060018060a01b0360018201541690600281015460038201549060ff6005600485015494015416936040519586526020860152604085015260608401526080830152151560a0820152f35b3461016a57604036600319011261016a576004356024358015159081810361016a577f82f1ecf86a9817521d2294bafc22d903bcad5c99e954156dfe15a17d381465e091610464602092610535610f49565b5f868152600285526040902054610556906001600160a01b03161515610d8c565b855f5260028452600560405f20019060ff801983541691151516179055565b3461016a57604036600319011261016a5761058e610ffd565b61059c60243560043561107c565b60018055005b3461016a575f36600319011261016a576020600354604051908152f35b3461016a575f36600319011261016a575f546040516001600160a01b039091168152602090f35b3461016a57604036600319011261016a5760043567ffffffffffffffff811161016a57610617903690600401610bf8565b9060243567ffffffffffffffff811161016a57610638903690600401610bf8565b9290610642610ffd565b8382036106bb575f5b8281106106585760018055005b8061066f6106696001938688610d68565b35610e04565b50806106a7575b610681575b0161064b565b6106a261068f828688610d68565b3561069b838987610d68565b359061107c565b61067b565b506106b3818785610d68565b351515610676565b60405162461bcd60e51b8152602060048201526015602482015274082e4e4c2f240d8cadccee8d040dad2e6dac2e8c6d605b1b6044820152606490fd5b3461016a575f36600319011261016a57610710610f49565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461016a57602036600319011261016a575f60a060405161076f81610c29565b82815282602082015282604082015282606082015282608082015201526004355f52600260205260c060405f206040516107a881610c29565b60018060a01b038254169182825260018060a01b03600182015416602083019081526002820154604084019081526003830154916060850192835260a060ff6005600487015496608089019788520154169501941515855260405195865260018060a01b03905116602086015251604085015251606084015251608083015251151560a0820152f35b3461016a57604036600319011261016a5761084a610bcc565b610852610be2565b9061085b610f49565b6001600160a01b0316801561027c576001600160a01b03821691821561027c576040516370a0823160e01b815230600482015290602082602481865afa91821561091a575f926108e6575b50816108ae57005b816108dd7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e56209360209386610f6f565b604051908152a3005b9091506020813d602011610912575b8161090260209383610c59565b8101031261016a575190846108a6565b3d91506108f5565b6040513d5f823e3d90fd5b3461016a57608036600319011261016a5761093e610bcc565b610946610be2565b9060643590610953610f49565b6001600160a01b031691821561027c576001600160a01b0316801561027c57811561026d576020926005610a31926003549461098e86610cf8565b6003556040519361099e85610c29565b84528684019182526044356040808601918252606086019283525f60808701818152600160a089018181528b84526002808e5294909320985189546001600160a01b039182166001600160a01b0319918216178b559751918a0180549290911691909716179095559151908601559051600385015590516004840155519101805460ff191691151560ff16919091179055565b805f5260028252807f607fee93225368fbfdfbb5f502b83308d35c011b90d3dd40178380cd0d3cb123610a6d60405f2060405191829182610d1a565b0390a2604051908152f35b3461016a575f36600319011261016a57602060405160018152f35b3461016a57604036600319011261016a57610aac610bcc565b60243590610ab8610f49565b6001600160a01b031690811561027c57801561026d57804710610b0c576020816104645f8080807fc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed50497895af161045e610c7b565b60405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b6044820152606490fd5b3461016a57606036600319011261016a57610b61610bcc565b610b69610be2565b9060443590610b76610f49565b6001600160a01b031690811561027c576001600160a01b03831692831561027c57811561026d57816108dd7f8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e56209360209386610f6f565b600435906001600160a01b038216820361016a57565b602435906001600160a01b038216820361016a57565b9181601f8401121561016a5782359167ffffffffffffffff831161016a576020808501948460051b01011161016a57565b60c0810190811067ffffffffffffffff821117610c4557604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff821117610c4557604052565b3d15610cb5573d9067ffffffffffffffff8211610c455760405191610caa601f8201601f191660200184610c59565b82523d5f602084013e565b606090565b15610cc157565b60405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606490fd5b5f198114610d065760010190565b634e487b7160e01b5f52601160045260245ffd5b81546001600160a01b039081168252600183015416602082015260028201546040820152600382015460608201526004820154608082015260059091015460ff16151560a082015260c00190565b9190811015610d785760051b0190565b634e487b7160e01b5f52603260045260245ffd5b15610d9357565b60405162461bcd60e51b8152602060048201526011602482015270436f6e666967206e6f742065786973747360781b6044820152606490fd5b8115610dd6570690565b634e487b7160e01b5f52601260045260245ffd5b91908203918211610d0657565b91908201809211610d0657565b5f52600260205260405f2090604051610e1c81610c29565b82546001600160a01b03908116808352600185015490911660208301526002840154604083015260038401546060830181815260048601546080850181815260059097015460ff1615801560a0909601869052919492939092610ee2575b50610ed75715610ecc57610e9a610e94610ead9242610dcc565b42610dea565b9351610ea7835182610dcc565b90610dea565b80931192835f14610ebe5750505f90565b610ec9915190610df7565b90565b505090506001905f90565b50505090505f905f90565b9050155f610e7a565b67ffffffffffffffff8111610c455760051b60200190565b90610f0d82610eeb565b610f1a6040519182610c59565b8281528092610f2b601f1991610eeb565b0190602036910137565b8051821015610d785760209160051b010190565b5f546001600160a01b03163303610f5c57565b63118cdaa760e01b5f523360045260245ffd5b60405163a9059cbb60e01b60208281019182526001600160a01b03909416602483015260448083019590955293815290925f91610fad606482610c59565b519082855af11561091a575f513d610ff457506001600160a01b0381163b155b610fd45750565b635274afe760e01b5f9081526001600160a01b0391909116600452602490fd5b60011415610fcd565b60026001541461100e576002600155565b633ee5aeb560e01b5f5260045ffd5b9081602091031261016a5751801515810361016a5790565b9081526001600160a01b039182166020820152911660408201526060810191909152608081019190915260a081019190915260e060c082018190525f908201526101000190565b811561026d57805f52600260205260405f209060ff60058301541680611566575b1561152a576110ab81610e04565b9015611518575081546001600160a01b031660018114939084156114b55747945b604051637838174760e11b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316602082600481845afa91821561091a575f92611471575b506002870197606089546064604051809681936337dba1f760e21b835260048301528960248301528a604483015260018060a01b03165afa90811561091a575f5f945f9361142d575b5061117783611172878b610df7565b610df7565b91828110611417575087106113db5786941561130a575b88549495506020946001600160a01b03166001810361129957508954895460018b0154604051632fe9316f60e11b8152988997889687956111e59591949293926001600160a01b0390811692169060048801611035565b03925af190811561091a575f9161126a575b505b1561125b5760807f96c3158688b0338fd59b23dd6fceb0f7b812847dac9bc3b620784f4da36b68ee9142600486015560018060a01b03855416946001808060a01b039101541695546040519182526020820152336040820152426060820152a4565b6307c4732760e51b5f5260045ffd5b61128c915060203d602011611292575b6112848183610c59565b81019061101d565b5f6111f7565b503d61127a565b8a5460018b0154604051632fe9316f60e11b815298899788965f965087956112d39591949293926001600160a01b03169160048801611035565b03925af190811561091a575f916112eb575b506111f9565b611304915060203d602011611292576112848183610c59565b5f6112e5565b604051636eb1769f60e11b8152306004820152602481018490529450602085604481895afa801561091a5787955f916113a5575b5081111561118e57935060205f9560446040518098819363095ea7b360e01b8352876004840152811960248401525af193841561091a578695602095611388575b5085945061118e565b61139e90863d8811611292576112848183610c59565b505f61137f565b9550506020853d6020116113d3575b816113c160209383610c59565b8101031261016a57808795519061133e565b3d91506113b4565b60405162461bcd60e51b8152602060048201526014602482015273416d6f756e742062656c6f77206d696e696d756d60601b6044820152606490fd5b8263203b880360e11b5f5260045260245260445ffd5b92505092506060813d606011611469575b8161144b60609383610c59565b8101031261016a57805160208201516040909201519193905f611163565b3d915061143e565b9091506020813d6020116114ad575b8161148d60209383610c59565b8101031261016a57516001600160a01b038116810361016a57905f61111a565b3d9150611480565b6040516370a0823160e01b8152306004820152602081602481855afa90811561091a575f916114e6575b50946110cc565b90506020813d602011611510575b8161150160209383610c59565b8101031261016a57515f6114df565b3d91506114f4565b6357c8d07f60e11b5f5260045260245ffd5b60405162461bcd60e51b8152602060048201526014602482015273436f6e666967206e6f7420617661696c61626c6560601b6044820152606490fd5b5081546001600160a01b0316151561109d56fea264697066735822122083e4eb090766de61c58833b245160f4db4284dce552696f071a6db82fb8d280464736f6c634300081c0033",
}

// BridgeBotABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBotMetaData.ABI instead.
var BridgeBotABI = BridgeBotMetaData.ABI

// BridgeBotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeBotMetaData.Bin instead.
var BridgeBotBin = BridgeBotMetaData.Bin

// DeployBridgeBot deploys a new Ethereum contract, binding an instance of BridgeBot to it.
func DeployBridgeBot(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _owner common.Address) (common.Address, *types.Transaction, *BridgeBot, error) {
	parsed, err := BridgeBotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBotBin), backend, _bridge, _owner)
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
