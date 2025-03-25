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

// CrossCheckBlockCheckBlock is an auto generated low-level Go binding around an user-defined struct.
type CrossCheckBlockCheckBlock struct {
	Nonce     uint64
	Start     uint64
	End       uint64
	CreatedAt uint64
	RootHash  [32]byte
	Proposer  common.Address
}

// CrossCheckMetaData contains all meta data concerning the CrossCheck contract.
var CrossCheckMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeThreshold\",\"inputs\":[{\"name\":\"threshold_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckBlock\",\"inputs\":[{\"name\":\"startBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCrossCheckBlock.CheckBlock\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckBlockByBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCrossCheckBlock.CheckBlock\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckBlockByNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCrossCheckBlock.CheckBlock\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextBlockInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"nextNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nextStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRoleBatch\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"validatorThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRoleBatch\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCheckpoint\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CrossCheckInitialized\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blocksPerCheck\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"validatorThreshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewCheckBlock\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CrossCheckBlockExists\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckBlockNotFound\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidChainID\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CrossCheckNotNextBlock\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RoleManagerAlreadyHasRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RoleManagerDoesNotHaveRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorInsufficientSignature\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorInvalidSignatures\",\"inputs\":[{\"name\":\"vLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotAuthorized\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ValidatorThresholdCanNotZero\",\"inputs\":[]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610042575f5ffd5b5061005161005660201b60201c565b6101b6565b5f61006561015460201b60201c565b9050805f0160089054906101000a900460ff16156100af576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101515767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051610148919061019d565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f67ffffffffffffffff82169050919050565b6101978161017b565b82525050565b5f6020820190506101b05f83018461018e565b92915050565b608051614d996101dc5f395f818161255b015281816125b0015261276f0152614d995ff3fe6080604052600436106101d7575f3560e01c80638456cb5911610101578063aea3679f11610094578063d547741f11610063578063d547741f1461069e578063f5b541a6146106c6578063f698da25146106f0578063fd47852b1461071a576101d7565b8063aea3679f146105d4578063b7f3358d14610610578063c49baebe14610638578063cf7910b614610662576101d7565b80639f3e14cf116100d05780639f3e14cf14610508578063a217fddf14610544578063a3246ad31461056e578063ad3cb1cc146105aa576101d7565b80638456cb591461044a57806384b0196e14610460578063871a782e1461049057806391d14854146104cc576101d7565b80633f4ba83a116101795780634f1ef286116101485780634f1ef286146103b257806352d1902d146103ce5780635c975abb146103f857806382f51fa614610422576101d7565b80633f4ba83a1461031f57806342cde4e8146103355780634351e6b61461035f5780634d58190714610387576101d7565b806317894ff4116101b557806317894ff414610269578063248a9ca3146102935780632f2ff15d146102cf57806336568abe146102f7576101d7565b806301ffc9a7146101db57806307a9eef21461021757806307e2da961461023f575b5f5ffd5b3480156101e6575f5ffd5b5061020160048036038101906101fc91906138d8565b610742565b60405161020e919061391d565b60405180910390f35b348015610222575f5ffd5b5061023d60048036038101906102389190613c08565b6107bb565b005b34801561024a575f5ffd5b506102536109e4565b6040516102609190613cf8565b60405180910390f35b348015610274575f5ffd5b5061027d6109ea565b60405161028a9190613cf8565b60405180910390f35b34801561029e575f5ffd5b506102b960048036038101906102b49190613d11565b6109f0565b6040516102c69190613d4b565b60405180910390f35b3480156102da575f5ffd5b506102f560048036038101906102f09190613dbe565b610a1a565b005b348015610302575f5ffd5b5061031d60048036038101906103189190613dbe565b610a3c565b005b34801561032a575f5ffd5b50610333610ab7565b005b348015610340575f5ffd5b50610349610aec565b6040516103569190613e0b565b60405180910390f35b34801561036a575f5ffd5b5061038560048036038101906103809190613e24565b610b0e565b005b348015610392575f5ffd5b5061039b610ce9565b6040516103a9929190613e4f565b60405180910390f35b6103cc60048036038101906103c79190613f26565b610d7c565b005b3480156103d9575f5ffd5b506103e2610d9b565b6040516103ef9190613d4b565b60405180910390f35b348015610403575f5ffd5b5061040c610dcc565b604051610419919061391d565b60405180910390f35b34801561042d575f5ffd5b5061044860048036038101906104439190614040565b610dee565b005b348015610455575f5ffd5b5061045e610e44565b005b34801561046b575f5ffd5b50610474610e79565b60405161048797969594939291906141fa565b60405180910390f35b34801561049b575f5ffd5b506104b660048036038101906104b191906142a6565b610f82565b6040516104c3919061438a565b60405180910390f35b3480156104d7575f5ffd5b506104f260048036038101906104ed9190613dbe565b6110ce565b6040516104ff919061391d565b60405180910390f35b348015610513575f5ffd5b5061052e600480360381019061052991906142a6565b61113f565b60405161053b919061438a565b60405180910390f35b34801561054f575f5ffd5b506105586116e4565b6040516105659190613d4b565b60405180910390f35b348015610579575f5ffd5b50610594600480360381019061058f9190613d11565b6116ea565b6040516105a1919061444b565b60405180910390f35b3480156105b5575f5ffd5b506105be611719565b6040516105cb919061446b565b60405180910390f35b3480156105df575f5ffd5b506105fa60048036038101906105f591906144e0565b611752565b604051610607919061391d565b60405180910390f35b34801561061b575f5ffd5b5061063660048036038101906106319190613e24565b6117cc565b005b348015610643575f5ffd5b5061064c611876565b6040516106599190613d4b565b60405180910390f35b34801561066d575f5ffd5b50610688600480360381019061068391906142a6565b61189a565b604051610695919061438a565b60405180910390f35b3480156106a9575f5ffd5b506106c460048036038101906106bf9190613dbe565b611bc4565b005b3480156106d1575f5ffd5b506106da611be6565b6040516106e79190613d4b565b60405180910390f35b3480156106fb575f5ffd5b50610704611c0a565b6040516107119190613d4b565b60405180910390f35b348015610725575f5ffd5b50610740600480360381019061073b9190614040565b611c18565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806107b457506107b382611c6e565b5b9050919050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989266107e581611cd7565b6107ed611ceb565b5f6107f6611d2c565b90505f878781019061080891906145de565b9050620956e28160800151146108595780608001516040517fa5acee440000000000000000000000000000000000000000000000000000000081526004016108509190613cf8565b60405180910390fd5b5f5f610863610ce9565b915091505f5f90506001620151808301039050835f01518314158061088c575083602001518214155b8061089b575083604001518114155b156108e657835f015184602001516040517f99a5875b0000000000000000000000000000000000000000000000000000000081526004016108dd929190613e4f565b60405180910390fd5b5f7ffe27049a99724ab477714cb3e8105a7b7d274fe7e5bee5cf905d63758b108490855f0151866020015187604001518860600151896080015160405160200161093596959493929190614609565b604051602081830303815290604052805190602001209050610959818b8b8b611d33565b61097586865f0151876020015188604001518960600151611f5e565b8460200151855f01518773ffffffffffffffffffffffffffffffffffffffff167fbb765a5098e4a96564f8633892153a94a950dadf5b371b7b46fa9fa170e034c2886040015189606001516040516109ce929190614668565b60405180910390a4505050505050505050505050565b60025481565b60035481565b5f5f6109fa612242565b9050805f015f8481526020019081526020015f2060010154915050919050565b610a23826109f0565b610a2c81611cd7565b610a368383612269565b50505050565b610a44611d2c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610aa8576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ab282826122fa565b505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929610ae181611cd7565b610ae961238b565b50565b5f5f610af66123f9565b9050805f015f9054906101000a900460ff1691505090565b5f610b17612420565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610b5f5750825b90505f60018367ffffffffffffffff16148015610b9257505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610ba0575080155b15610bd7576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610c24576001855f0160086101000a81548160ff0219169083151502179055505b610c2c612447565b610c34612451565b610c3d33612463565b610c4686612477565b7fb80103afdb269ad46586734e4733e996c5fd0ae2882ba36f5f8f2033ae8be277620956e26201518088604051610c7f9392919061468f565b60405180910390a18315610ce1575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610cd89190614706565b60405180910390a15b505050505050565b5f5f5f5f5f60025481526020019081526020015f2090505f815f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115610d77576001815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff160192506001815f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff160191505b509091565b610d84612559565b610d8d8261263f565b610d97828261264f565b5050565b5f610da461276d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b5f5f610dd66127f4565b9050805f015f9054906101000a900460ff1691505090565b610df7826109f0565b610e0081611cd7565b5f5f90505b8251811015610e3e57610e3284848381518110610e2557610e2461471f565b5b60200260200101516122fa565b50806001019050610e05565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929610e6e81611cd7565b610e7661281b565b50565b5f6060805f5f5f60605f610e8b61288a565b90505f5f1b815f0154148015610ea657505f5f1b8160010154145b610ee5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610edc90614796565b60405180910390fd5b610eed6128b1565b610ef561294f565b46305f5f1b5f67ffffffffffffffff811115610f1457610f136139a7565b5b604051908082528060200260200182016040528015610f425781602001602082028036833780820191505090505b507f0f0000000000000000000000000000000000000000000000000000000000000095949392919097509750975097509750975097505090919293949596565b610f8a613804565b5f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b5f5f6110d8612242565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b611147613804565b5f5f5f8481526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1611156112c3575f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506116df565b5f820361140e575f5f5f81526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506116df565b5f5f90505f600160035461142291906147e1565b90505b80821161166d575f6002828461143b9190614814565b6114459190614874565b90505f5f5f60015f8581526020019081526020015f205481526020019081526020015f20905085815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16111580156114bd5750805f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168611155b1561162b575f5f825f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509450505050506116df565b805f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1686101561165f57600182039250611666565b6001820193505b5050611425565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250925050505b919050565b5f5f1b81565b60605f6116f56129ed565b9050611711815f015f8581526020019081526020015f20612a14565b915050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f5f61175d8661113f565b90505f816060015167ffffffffffffffff16036117b157856040517f96fab5950000000000000000000000000000000000000000000000000000000081526004016117a89190613cf8565b60405180910390fd5b6117c18585836080015186612a33565b915050949350505050565b5f5f1b6117d881611cd7565b5f8260ff1603611814576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61181d6123f9565b905082815f015f6101000a81548160ff021916908360ff1602179055507f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff836040516118699190613e0b565b60405180910390a1505050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6118a2613804565b5f82036119ed575f5f5f81526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050611bbf565b5f60015f8481526020019081526020015f205490505f8103611a7e576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250915050611bbf565b5f5f8281526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509150505b919050565b611bcd826109f0565b611bd681611cd7565b611be083836122fa565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b5f611c13612a4b565b905090565b611c21826109f0565b611c2a81611cd7565b5f5f90505b8251811015611c6857611c5c84848381518110611c4f57611c4e61471f565b5b6020026020010151612269565b50806001019050611c2f565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b611ce881611ce3611d2c565b612a59565b50565b611cf3610dcc565b15611d2a576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f33905090565b5f611d3c6123f9565b90505f84519050835181148015611d535750825181145b8185518551909192611d9d576040517f9089eb84000000000000000000000000000000000000000000000000000000008152600401611d94939291906148a4565b60405180910390fd5b505050815f015f9054906101000a900460ff1660ff168110158190611df8576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611def9190613cf8565b60405180910390fd5b505f5f90505f5f90505f5f90505b83811015611efa575f611e7c898381518110611e2557611e2461471f565b5b6020026020010151898481518110611e4057611e3f61471f565b5b6020026020010151898581518110611e5b57611e5a61471f565b5b6020026020010151611e6c8e612aaa565b612ac3909392919063ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16108015611edf5750611ede7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926826110ce565b5b15611eee578360010193508092505b50806001019050611e06565b50835f015f9054906101000a900460ff1660ff168210158290611f53576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611f4a9190613cf8565b60405180910390fd5b505050505050505050565b8282111580611f6e57505f5f1b81145b15611fa5576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016841180611fc6575067ffffffffffffffff801683115b80611fda575067ffffffffffffffff801682115b15612011576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8581526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff16111561208657826040517fb7aa1d9700000000000000000000000000000000000000000000000000000000815260040161207d9190613cf8565b60405180910390fd5b5f6040518060c001604052808667ffffffffffffffff1681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681526020018381526020018773ffffffffffffffffffffffffffffffffffffffff168152509050805f5f8681526020019081526020015f205f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f0160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151815f0160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001015560a0820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050508360015f8781526020019081526020015f208190555060035f81546001019190508190555083600281905550505050505050565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b5f6122748383612af1565b905080156122f4575f6122856129ed565b90506122ab83825f015f8781526020019081526020015f20612be990919063ffffffff16565b838590916122f0576040517fd180cb310000000000000000000000000000000000000000000000000000000081526004016122e79291906148d9565b60405180910390fd5b5050505b92915050565b5f6123058383612c16565b90508015612385575f6123166129ed565b905061233c83825f015f8781526020019081526020015f20612d0e90919063ffffffff16565b83859091612381576040517f054af8d80000000000000000000000000000000000000000000000000000000081526004016123789291906148d9565b60405180910390fd5b5050505b92915050565b612393612d3b565b5f61239c6127f4565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6123e1611d2c565b6040516123ee9190614900565b60405180910390a150565b5f7f303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b61244f612d7b565b565b612459612d7b565b612461612dbb565b565b61246b612d7b565b61247481612deb565b50565b61247f612d7b565b5f8160ff16036124bb576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61252f6040518060400160405280600981526020017f56616c696461746f7200000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e30000000000000000000000000000000000000000000000000000000815250612e0b565b5f6125386123f9565b905081815f015f6101000a81548160ff021916908360ff1602179055505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061260657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125ed612e21565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561263d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f1b61264b81611cd7565b5050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156126b757506040513d601f19601f820116820180604052508101906126b4919061492d565b60015b6126f857816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016126ef9190614900565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461275e57806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016127559190613d4b565b60405180910390fd5b6127688383612e74565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146127f2576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b612823611ceb565b5f61282c6127f4565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612872611d2c565b60405161287f9190614900565b60405180910390a150565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100905090565b60605f6128bc61288a565b90508060020180546128cd90614985565b80601f01602080910402602001604051908101604052809291908181526020018280546128f990614985565b80156129445780601f1061291b57610100808354040283529160200191612944565b820191905f5260205f20905b81548152906001019060200180831161292757829003601f168201915b505050505091505090565b60605f61295a61288a565b905080600301805461296b90614985565b80601f016020809104026020016040519081016040528092919081815260200182805461299790614985565b80156129e25780601f106129b9576101008083540402835291602001916129e2565b820191905f5260205f20905b8154815290600101906020018083116129c557829003601f168201915b505050505091505090565b5f7f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00905090565b60605f612a22835f01612ee6565b905060608190508092505050919050565b5f82612a40868685612f3f565b149050949350505050565b5f612a54612f92565b905090565b612a6382826110ce565b612aa65780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612a9d9291906148d9565b60405180910390fd5b5050565b5f612abc612ab6612a4b565b83612ff5565b9050919050565b5f5f5f5f612ad388888888613035565b925092509250612ae3828261311c565b829350505050949350505050565b5f5f612afb612242565b9050612b0784846110ce565b612bde576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612b7a611d2c565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612be3565b5f9150505b92915050565b5f612c0e835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61327e565b905092915050565b5f5f612c20612242565b9050612c2c84846110ce565b15612d03575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612c9f611d2c565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050612d08565b5f9150505b92915050565b5f612d33835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6132e5565b905092915050565b612d43610dcc565b612d79576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612d836133e1565b612db9576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612dc3612d7b565b5f612dcc6127f4565b90505f815f015f6101000a81548160ff02191690831515021790555050565b612df3612d7b565b612dfb6133ff565b612e075f5f1b82612269565b5050565b612e13612d7b565b612e1d8282613409565b5050565b5f612e4d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61345a565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b612e7d82613463565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115612ed957612ed3828261352c565b50612ee2565b612ee16135ac565b5b5050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612f3357602002820191905f5260205f20905b815481526020019060010190808311612f1f575b50505050509050919050565b5f5f8290505f5f90505b85859050811015612f8657612f7782878784818110612f6b57612f6a61471f565b5b905060200201356135e8565b91508080600101915050612f49565b50809150509392505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612fbc613612565b612fc4613688565b4630604051602001612fda9594939291906149b5565b60405160208183030381529060405280519060200120905090565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c1115613071575f600385925092509250613112565b5f6001888888886040515f81526020016040526040516130949493929190614a06565b6020604051602081039080840390855afa1580156130b4573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613105575f60015f5f1b93509350935050613112565b805f5f5f1b935093509350505b9450945094915050565b5f600381111561312f5761312e614a49565b5b82600381111561314257613141614a49565b5b031561327a576001600381111561315c5761315b614a49565b5b82600381111561316f5761316e614a49565b5b036131a6576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260038111156131ba576131b9614a49565b5b8260038111156131cd576131cc614a49565b5b0361321157805f1c6040517ffce698f70000000000000000000000000000000000000000000000000000000081526004016132089190613cf8565b60405180910390fd5b60038081111561322457613223614a49565b5b82600381111561323757613236614a49565b5b0361327957806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016132709190613d4b565b60405180910390fd5b5b5050565b5f61328983836136ff565b6132db57825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f2081905550600190506132df565b5f90505b92915050565b5f5f836001015f8481526020019081526020015f205490505f81146133d6575f60018261331291906147e1565b90505f6001865f018054905061332891906147e1565b905080821461338e575f865f0182815481106133475761334661471f565b5b905f5260205f200154905080875f0184815481106133685761336761471f565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f018054806133a1576133a0614a76565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506133db565b5f9150505b92915050565b5f6133ea612420565b5f0160089054906101000a900460ff16905090565b613407612d7b565b565b613411612d7b565b5f61341a61288a565b90508281600201908161342d9190614c3a565b508181600301908161343f9190614c3a565b505f5f1b815f01819055505f5f1b8160010181905550505050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b036134be57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016134b59190614900565b60405180910390fd5b806134ea7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61345a565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516135559190614d4d565b5f60405180830381855af49150503d805f811461358d576040519150601f19603f3d011682016040523d82523d5f602084013e613592565b606091505b50915091506135a285838361371f565b9250505092915050565b5f3411156135e6576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8183106135ff576135fa82846137ac565b61360a565b61360983836137ac565b5b905092915050565b5f5f61361c61288a565b90505f6136276128b1565b90505f8151111561364357808051906020012092505050613685565b5f825f015490505f5f1b811461365e57809350505050613685565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f61369261288a565b90505f61369d61294f565b90505f815111156136b9578080519060200120925050506136fc565b5f826001015490505f5f1b81146136d5578093505050506136fc565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f836001015f8481526020019081526020015f20541415905092915050565b6060826137345761372f826137c0565b6137a4565b5f825114801561375a57505f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561379c57836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016137939190614900565b60405180910390fd5b8190506137a5565b5b9392505050565b5f825f528160205260405f20905092915050565b5f815111156137d25780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6138b781613883565b81146138c1575f5ffd5b50565b5f813590506138d2816138ae565b92915050565b5f602082840312156138ed576138ec61387b565b5b5f6138fa848285016138c4565b91505092915050565b5f8115159050919050565b61391781613903565b82525050565b5f6020820190506139305f83018461390e565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261395757613956613936565b5b8235905067ffffffffffffffff8111156139745761397361393a565b5b6020830191508360018202830111156139905761398f61393e565b5b9250929050565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6139dd82613997565b810181811067ffffffffffffffff821117156139fc576139fb6139a7565b5b80604052505050565b5f613a0e613872565b9050613a1a82826139d4565b919050565b5f67ffffffffffffffff821115613a3957613a386139a7565b5b602082029050602081019050919050565b5f60ff82169050919050565b613a5f81613a4a565b8114613a69575f5ffd5b50565b5f81359050613a7a81613a56565b92915050565b5f613a92613a8d84613a1f565b613a05565b90508083825260208201905060208402830185811115613ab557613ab461393e565b5b835b81811015613ade5780613aca8882613a6c565b845260208401935050602081019050613ab7565b5050509392505050565b5f82601f830112613afc57613afb613936565b5b8135613b0c848260208601613a80565b91505092915050565b5f67ffffffffffffffff821115613b2f57613b2e6139a7565b5b602082029050602081019050919050565b5f819050919050565b613b5281613b40565b8114613b5c575f5ffd5b50565b5f81359050613b6d81613b49565b92915050565b5f613b85613b8084613b15565b613a05565b90508083825260208201905060208402830185811115613ba857613ba761393e565b5b835b81811015613bd15780613bbd8882613b5f565b845260208401935050602081019050613baa565b5050509392505050565b5f82601f830112613bef57613bee613936565b5b8135613bff848260208601613b73565b91505092915050565b5f5f5f5f5f60808688031215613c2157613c2061387b565b5b5f86013567ffffffffffffffff811115613c3e57613c3d61387f565b5b613c4a88828901613942565b9550955050602086013567ffffffffffffffff811115613c6d57613c6c61387f565b5b613c7988828901613ae8565b935050604086013567ffffffffffffffff811115613c9a57613c9961387f565b5b613ca688828901613bdb565b925050606086013567ffffffffffffffff811115613cc757613cc661387f565b5b613cd388828901613bdb565b9150509295509295909350565b5f819050919050565b613cf281613ce0565b82525050565b5f602082019050613d0b5f830184613ce9565b92915050565b5f60208284031215613d2657613d2561387b565b5b5f613d3384828501613b5f565b91505092915050565b613d4581613b40565b82525050565b5f602082019050613d5e5f830184613d3c565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613d8d82613d64565b9050919050565b613d9d81613d83565b8114613da7575f5ffd5b50565b5f81359050613db881613d94565b92915050565b5f5f60408385031215613dd457613dd361387b565b5b5f613de185828601613b5f565b9250506020613df285828601613daa565b9150509250929050565b613e0581613a4a565b82525050565b5f602082019050613e1e5f830184613dfc565b92915050565b5f60208284031215613e3957613e3861387b565b5b5f613e4684828501613a6c565b91505092915050565b5f604082019050613e625f830185613ce9565b613e6f6020830184613ce9565b9392505050565b5f5ffd5b5f67ffffffffffffffff821115613e9457613e936139a7565b5b613e9d82613997565b9050602081019050919050565b828183375f83830152505050565b5f613eca613ec584613e7a565b613a05565b905082815260208101848484011115613ee657613ee5613e76565b5b613ef1848285613eaa565b509392505050565b5f82601f830112613f0d57613f0c613936565b5b8135613f1d848260208601613eb8565b91505092915050565b5f5f60408385031215613f3c57613f3b61387b565b5b5f613f4985828601613daa565b925050602083013567ffffffffffffffff811115613f6a57613f6961387f565b5b613f7685828601613ef9565b9150509250929050565b5f67ffffffffffffffff821115613f9a57613f996139a7565b5b602082029050602081019050919050565b5f613fbd613fb884613f80565b613a05565b90508083825260208201905060208402830185811115613fe057613fdf61393e565b5b835b818110156140095780613ff58882613daa565b845260208401935050602081019050613fe2565b5050509392505050565b5f82601f83011261402757614026613936565b5b8135614037848260208601613fab565b91505092915050565b5f5f604083850312156140565761405561387b565b5b5f61406385828601613b5f565b925050602083013567ffffffffffffffff8111156140845761408361387f565b5b61409085828601614013565b9150509250929050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6140ce8161409a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614106826140d4565b61411081856140de565b93506141208185602086016140ee565b61412981613997565b840191505092915050565b61413d81613d83565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61417581613ce0565b82525050565b5f614186838361416c565b60208301905092915050565b5f602082019050919050565b5f6141a882614143565b6141b2818561414d565b93506141bd8361415d565b805f5b838110156141ed5781516141d4888261417b565b97506141df83614192565b9250506001810190506141c0565b5085935050505092915050565b5f60e08201905061420d5f83018a6140c5565b818103602083015261421f81896140fc565b9050818103604083015261423381886140fc565b90506142426060830187613ce9565b61424f6080830186614134565b61425c60a0830185613d3c565b81810360c083015261426e818461419e565b905098975050505050505050565b61428581613ce0565b811461428f575f5ffd5b50565b5f813590506142a08161427c565b92915050565b5f602082840312156142bb576142ba61387b565b5b5f6142c884828501614292565b91505092915050565b5f67ffffffffffffffff82169050919050565b6142ed816142d1565b82525050565b6142fc81613b40565b82525050565b61430b81613d83565b82525050565b60c082015f8201516143255f8501826142e4565b50602082015161433860208501826142e4565b50604082015161434b60408501826142e4565b50606082015161435e60608501826142e4565b50608082015161437160808501826142f3565b5060a082015161438460a0850182614302565b50505050565b5f60c08201905061439d5f830184614311565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6143d78383614302565b60208301905092915050565b5f602082019050919050565b5f6143f9826143a3565b61440381856143ad565b935061440e836143bd565b805f5b8381101561443e57815161442588826143cc565b9750614430836143e3565b925050600181019050614411565b5085935050505092915050565b5f6020820190508181035f83015261446381846143ef565b905092915050565b5f6020820190508181035f83015261448381846140fc565b905092915050565b5f5f83601f8401126144a05761449f613936565b5b8235905067ffffffffffffffff8111156144bd576144bc61393a565b5b6020830191508360208202830111156144d9576144d861393e565b5b9250929050565b5f5f5f5f606085870312156144f8576144f761387b565b5b5f61450587828801614292565b945050602085013567ffffffffffffffff8111156145265761452561387f565b5b6145328782880161448b565b9350935050604061454587828801613b5f565b91505092959194509250565b5f5ffd5b5f60a0828403121561456a57614569614551565b5b61457460a0613a05565b90505f61458384828501614292565b5f83015250602061459684828501614292565b60208301525060406145aa84828501614292565b60408301525060606145be84828501613b5f565b60608301525060806145d284828501614292565b60808301525092915050565b5f60a082840312156145f3576145f261387b565b5b5f61460084828501614555565b91505092915050565b5f60c08201905061461c5f830189613d3c565b6146296020830188613ce9565b6146366040830187613ce9565b6146436060830186613ce9565b6146506080830185613d3c565b61465d60a0830184613ce9565b979650505050505050565b5f60408201905061467b5f830185613ce9565b6146886020830184613d3c565b9392505050565b5f6060820190506146a25f830186613ce9565b6146af6020830185613ce9565b6146bc6040830184613dfc565b949350505050565b5f819050919050565b5f819050919050565b5f6146f06146eb6146e6846146c4565b6146cd565b6142d1565b9050919050565b614700816146d6565b82525050565b5f6020820190506147195f8301846146f7565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4549503731323a20556e696e697469616c697a656400000000000000000000005f82015250565b5f6147806015836140de565b915061478b8261474c565b602082019050919050565b5f6020820190508181035f8301526147ad81614774565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6147eb82613ce0565b91506147f683613ce0565b925082820390508181111561480e5761480d6147b4565b5b92915050565b5f61481e82613ce0565b915061482983613ce0565b9250828201905080821115614841576148406147b4565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61487e82613ce0565b915061488983613ce0565b92508261489957614898614847565b5b828204905092915050565b5f6060820190506148b75f830186613ce9565b6148c46020830185613ce9565b6148d16040830184613ce9565b949350505050565b5f6040820190506148ec5f830185614134565b6148f96020830184613d3c565b9392505050565b5f6020820190506149135f830184614134565b92915050565b5f8151905061492781613b49565b92915050565b5f602082840312156149425761494161387b565b5b5f61494f84828501614919565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061499c57607f821691505b6020821081036149af576149ae614958565b5b50919050565b5f60a0820190506149c85f830188613d3c565b6149d56020830187613d3c565b6149e26040830186613d3c565b6149ef6060830185613ce9565b6149fc6080830184614134565b9695505050505050565b5f608082019050614a195f830187613d3c565b614a266020830186613dfc565b614a336040830185613d3c565b614a406060830184613d3c565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302614aff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614ac4565b614b098683614ac4565b95508019841693508086168417925050509392505050565b5f614b3b614b36614b3184613ce0565b6146cd565b613ce0565b9050919050565b5f819050919050565b614b5483614b21565b614b68614b6082614b42565b848454614ad0565b825550505050565b5f5f905090565b614b7f614b70565b614b8a818484614b4b565b505050565b5b81811015614bad57614ba25f82614b77565b600181019050614b90565b5050565b601f821115614bf257614bc381614aa3565b614bcc84614ab5565b81016020851015614bdb578190505b614bef614be785614ab5565b830182614b8f565b50505b505050565b5f82821c905092915050565b5f614c125f1984600802614bf7565b1980831691505092915050565b5f614c2a8383614c03565b9150826002028217905092915050565b614c43826140d4565b67ffffffffffffffff811115614c5c57614c5b6139a7565b5b614c668254614985565b614c71828285614bb1565b5f60209050601f831160018114614ca2575f8415614c90578287015190505b614c9a8582614c1f565b865550614d01565b601f198416614cb086614aa3565b5f5b82811015614cd757848901518255600182019150602085019450602081019050614cb2565b86831015614cf45784890151614cf0601f891682614c03565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f81905092915050565b5f614d2782614d09565b614d318185614d13565b9350614d418185602086016140ee565b80840191505092915050565b5f614d588284614d1d565b91508190509291505056fea26469706673582212202f82185f41b39990e018183f5e104cce4b3501255bac628e689d60b1fe78b0bc64736f6c634300081c0033",
}

// CrossCheckABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossCheckMetaData.ABI instead.
var CrossCheckABI = CrossCheckMetaData.ABI

// CrossCheckBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossCheckBinRuntime = "6080604052600436106101d7575f3560e01c80638456cb5911610101578063aea3679f11610094578063d547741f11610063578063d547741f1461069e578063f5b541a6146106c6578063f698da25146106f0578063fd47852b1461071a576101d7565b8063aea3679f146105d4578063b7f3358d14610610578063c49baebe14610638578063cf7910b614610662576101d7565b80639f3e14cf116100d05780639f3e14cf14610508578063a217fddf14610544578063a3246ad31461056e578063ad3cb1cc146105aa576101d7565b80638456cb591461044a57806384b0196e14610460578063871a782e1461049057806391d14854146104cc576101d7565b80633f4ba83a116101795780634f1ef286116101485780634f1ef286146103b257806352d1902d146103ce5780635c975abb146103f857806382f51fa614610422576101d7565b80633f4ba83a1461031f57806342cde4e8146103355780634351e6b61461035f5780634d58190714610387576101d7565b806317894ff4116101b557806317894ff414610269578063248a9ca3146102935780632f2ff15d146102cf57806336568abe146102f7576101d7565b806301ffc9a7146101db57806307a9eef21461021757806307e2da961461023f575b5f5ffd5b3480156101e6575f5ffd5b5061020160048036038101906101fc91906138d8565b610742565b60405161020e919061391d565b60405180910390f35b348015610222575f5ffd5b5061023d60048036038101906102389190613c08565b6107bb565b005b34801561024a575f5ffd5b506102536109e4565b6040516102609190613cf8565b60405180910390f35b348015610274575f5ffd5b5061027d6109ea565b60405161028a9190613cf8565b60405180910390f35b34801561029e575f5ffd5b506102b960048036038101906102b49190613d11565b6109f0565b6040516102c69190613d4b565b60405180910390f35b3480156102da575f5ffd5b506102f560048036038101906102f09190613dbe565b610a1a565b005b348015610302575f5ffd5b5061031d60048036038101906103189190613dbe565b610a3c565b005b34801561032a575f5ffd5b50610333610ab7565b005b348015610340575f5ffd5b50610349610aec565b6040516103569190613e0b565b60405180910390f35b34801561036a575f5ffd5b5061038560048036038101906103809190613e24565b610b0e565b005b348015610392575f5ffd5b5061039b610ce9565b6040516103a9929190613e4f565b60405180910390f35b6103cc60048036038101906103c79190613f26565b610d7c565b005b3480156103d9575f5ffd5b506103e2610d9b565b6040516103ef9190613d4b565b60405180910390f35b348015610403575f5ffd5b5061040c610dcc565b604051610419919061391d565b60405180910390f35b34801561042d575f5ffd5b5061044860048036038101906104439190614040565b610dee565b005b348015610455575f5ffd5b5061045e610e44565b005b34801561046b575f5ffd5b50610474610e79565b60405161048797969594939291906141fa565b60405180910390f35b34801561049b575f5ffd5b506104b660048036038101906104b191906142a6565b610f82565b6040516104c3919061438a565b60405180910390f35b3480156104d7575f5ffd5b506104f260048036038101906104ed9190613dbe565b6110ce565b6040516104ff919061391d565b60405180910390f35b348015610513575f5ffd5b5061052e600480360381019061052991906142a6565b61113f565b60405161053b919061438a565b60405180910390f35b34801561054f575f5ffd5b506105586116e4565b6040516105659190613d4b565b60405180910390f35b348015610579575f5ffd5b50610594600480360381019061058f9190613d11565b6116ea565b6040516105a1919061444b565b60405180910390f35b3480156105b5575f5ffd5b506105be611719565b6040516105cb919061446b565b60405180910390f35b3480156105df575f5ffd5b506105fa60048036038101906105f591906144e0565b611752565b604051610607919061391d565b60405180910390f35b34801561061b575f5ffd5b5061063660048036038101906106319190613e24565b6117cc565b005b348015610643575f5ffd5b5061064c611876565b6040516106599190613d4b565b60405180910390f35b34801561066d575f5ffd5b50610688600480360381019061068391906142a6565b61189a565b604051610695919061438a565b60405180910390f35b3480156106a9575f5ffd5b506106c460048036038101906106bf9190613dbe565b611bc4565b005b3480156106d1575f5ffd5b506106da611be6565b6040516106e79190613d4b565b60405180910390f35b3480156106fb575f5ffd5b50610704611c0a565b6040516107119190613d4b565b60405180910390f35b348015610725575f5ffd5b50610740600480360381019061073b9190614040565b611c18565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806107b457506107b382611c6e565b5b9050919050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989266107e581611cd7565b6107ed611ceb565b5f6107f6611d2c565b90505f878781019061080891906145de565b9050620956e28160800151146108595780608001516040517fa5acee440000000000000000000000000000000000000000000000000000000081526004016108509190613cf8565b60405180910390fd5b5f5f610863610ce9565b915091505f5f90506001620151808301039050835f01518314158061088c575083602001518214155b8061089b575083604001518114155b156108e657835f015184602001516040517f99a5875b0000000000000000000000000000000000000000000000000000000081526004016108dd929190613e4f565b60405180910390fd5b5f7ffe27049a99724ab477714cb3e8105a7b7d274fe7e5bee5cf905d63758b108490855f0151866020015187604001518860600151896080015160405160200161093596959493929190614609565b604051602081830303815290604052805190602001209050610959818b8b8b611d33565b61097586865f0151876020015188604001518960600151611f5e565b8460200151855f01518773ffffffffffffffffffffffffffffffffffffffff167fbb765a5098e4a96564f8633892153a94a950dadf5b371b7b46fa9fa170e034c2886040015189606001516040516109ce929190614668565b60405180910390a4505050505050505050505050565b60025481565b60035481565b5f5f6109fa612242565b9050805f015f8481526020019081526020015f2060010154915050919050565b610a23826109f0565b610a2c81611cd7565b610a368383612269565b50505050565b610a44611d2c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610aa8576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ab282826122fa565b505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929610ae181611cd7565b610ae961238b565b50565b5f5f610af66123f9565b9050805f015f9054906101000a900460ff1691505090565b5f610b17612420565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610b5f5750825b90505f60018367ffffffffffffffff16148015610b9257505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610ba0575080155b15610bd7576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610c24576001855f0160086101000a81548160ff0219169083151502179055505b610c2c612447565b610c34612451565b610c3d33612463565b610c4686612477565b7fb80103afdb269ad46586734e4733e996c5fd0ae2882ba36f5f8f2033ae8be277620956e26201518088604051610c7f9392919061468f565b60405180910390a18315610ce1575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610cd89190614706565b60405180910390a15b505050505050565b5f5f5f5f5f60025481526020019081526020015f2090505f815f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115610d77576001815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff160192506001815f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff160191505b509091565b610d84612559565b610d8d8261263f565b610d97828261264f565b5050565b5f610da461276d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b5f5f610dd66127f4565b9050805f015f9054906101000a900460ff1691505090565b610df7826109f0565b610e0081611cd7565b5f5f90505b8251811015610e3e57610e3284848381518110610e2557610e2461471f565b5b60200260200101516122fa565b50806001019050610e05565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b929610e6e81611cd7565b610e7661281b565b50565b5f6060805f5f5f60605f610e8b61288a565b90505f5f1b815f0154148015610ea657505f5f1b8160010154145b610ee5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610edc90614796565b60405180910390fd5b610eed6128b1565b610ef561294f565b46305f5f1b5f67ffffffffffffffff811115610f1457610f136139a7565b5b604051908082528060200260200182016040528015610f425781602001602082028036833780820191505090505b507f0f0000000000000000000000000000000000000000000000000000000000000095949392919097509750975097509750975097505090919293949596565b610f8a613804565b5f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b5f5f6110d8612242565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b611147613804565b5f5f5f8481526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1611156112c3575f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506116df565b5f820361140e575f5f5f81526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506116df565b5f5f90505f600160035461142291906147e1565b90505b80821161166d575f6002828461143b9190614814565b6114459190614874565b90505f5f5f60015f8581526020019081526020015f205481526020019081526020015f20905085815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16111580156114bd5750805f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168611155b1561162b575f5f825f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509450505050506116df565b805f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1686101561165f57600182039250611666565b6001820193505b5050611425565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250925050505b919050565b5f5f1b81565b60605f6116f56129ed565b9050611711815f015f8581526020019081526020015f20612a14565b915050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f5f61175d8661113f565b90505f816060015167ffffffffffffffff16036117b157856040517f96fab5950000000000000000000000000000000000000000000000000000000081526004016117a89190613cf8565b60405180910390fd5b6117c18585836080015186612a33565b915050949350505050565b5f5f1b6117d881611cd7565b5f8260ff1603611814576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61181d6123f9565b905082815f015f6101000a81548160ff021916908360ff1602179055507f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff836040516118699190613e0b565b60405180910390a1505050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6118a2613804565b5f82036119ed575f5f5f81526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050611bbf565b5f60015f8481526020019081526020015f205490505f8103611a7e576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250915050611bbf565b5f5f8281526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509150505b919050565b611bcd826109f0565b611bd681611cd7565b611be083836122fa565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b5f611c13612a4b565b905090565b611c21826109f0565b611c2a81611cd7565b5f5f90505b8251811015611c6857611c5c84848381518110611c4f57611c4e61471f565b5b6020026020010151612269565b50806001019050611c2f565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b611ce881611ce3611d2c565b612a59565b50565b611cf3610dcc565b15611d2a576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f33905090565b5f611d3c6123f9565b90505f84519050835181148015611d535750825181145b8185518551909192611d9d576040517f9089eb84000000000000000000000000000000000000000000000000000000008152600401611d94939291906148a4565b60405180910390fd5b505050815f015f9054906101000a900460ff1660ff168110158190611df8576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611def9190613cf8565b60405180910390fd5b505f5f90505f5f90505f5f90505b83811015611efa575f611e7c898381518110611e2557611e2461471f565b5b6020026020010151898481518110611e4057611e3f61471f565b5b6020026020010151898581518110611e5b57611e5a61471f565b5b6020026020010151611e6c8e612aaa565b612ac3909392919063ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16108015611edf5750611ede7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926826110ce565b5b15611eee578360010193508092505b50806001019050611e06565b50835f015f9054906101000a900460ff1660ff168210158290611f53576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611f4a9190613cf8565b60405180910390fd5b505050505050505050565b8282111580611f6e57505f5f1b81145b15611fa5576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016841180611fc6575067ffffffffffffffff801683115b80611fda575067ffffffffffffffff801682115b15612011576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8581526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff16111561208657826040517fb7aa1d9700000000000000000000000000000000000000000000000000000000815260040161207d9190613cf8565b60405180910390fd5b5f6040518060c001604052808667ffffffffffffffff1681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681526020018381526020018773ffffffffffffffffffffffffffffffffffffffff168152509050805f5f8681526020019081526020015f205f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f0160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151815f0160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001015560a0820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050508360015f8781526020019081526020015f208190555060035f81546001019190508190555083600281905550505050505050565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b5f6122748383612af1565b905080156122f4575f6122856129ed565b90506122ab83825f015f8781526020019081526020015f20612be990919063ffffffff16565b838590916122f0576040517fd180cb310000000000000000000000000000000000000000000000000000000081526004016122e79291906148d9565b60405180910390fd5b5050505b92915050565b5f6123058383612c16565b90508015612385575f6123166129ed565b905061233c83825f015f8781526020019081526020015f20612d0e90919063ffffffff16565b83859091612381576040517f054af8d80000000000000000000000000000000000000000000000000000000081526004016123789291906148d9565b60405180910390fd5b5050505b92915050565b612393612d3b565b5f61239c6127f4565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6123e1611d2c565b6040516123ee9190614900565b60405180910390a150565b5f7f303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b61244f612d7b565b565b612459612d7b565b612461612dbb565b565b61246b612d7b565b61247481612deb565b50565b61247f612d7b565b5f8160ff16036124bb576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61252f6040518060400160405280600981526020017f56616c696461746f7200000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e30000000000000000000000000000000000000000000000000000000815250612e0b565b5f6125386123f9565b905081815f015f6101000a81548160ff021916908360ff1602179055505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061260657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125ed612e21565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561263d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f1b61264b81611cd7565b5050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156126b757506040513d601f19601f820116820180604052508101906126b4919061492d565b60015b6126f857816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016126ef9190614900565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b811461275e57806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016127559190613d4b565b60405180910390fd5b6127688383612e74565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16146127f2576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b612823611ceb565b5f61282c6127f4565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612872611d2c565b60405161287f9190614900565b60405180910390a150565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100905090565b60605f6128bc61288a565b90508060020180546128cd90614985565b80601f01602080910402602001604051908101604052809291908181526020018280546128f990614985565b80156129445780601f1061291b57610100808354040283529160200191612944565b820191905f5260205f20905b81548152906001019060200180831161292757829003601f168201915b505050505091505090565b60605f61295a61288a565b905080600301805461296b90614985565b80601f016020809104026020016040519081016040528092919081815260200182805461299790614985565b80156129e25780601f106129b9576101008083540402835291602001916129e2565b820191905f5260205f20905b8154815290600101906020018083116129c557829003601f168201915b505050505091505090565b5f7f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00905090565b60605f612a22835f01612ee6565b905060608190508092505050919050565b5f82612a40868685612f3f565b149050949350505050565b5f612a54612f92565b905090565b612a6382826110ce565b612aa65780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612a9d9291906148d9565b60405180910390fd5b5050565b5f612abc612ab6612a4b565b83612ff5565b9050919050565b5f5f5f5f612ad388888888613035565b925092509250612ae3828261311c565b829350505050949350505050565b5f5f612afb612242565b9050612b0784846110ce565b612bde576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612b7a611d2c565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612be3565b5f9150505b92915050565b5f612c0e835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61327e565b905092915050565b5f5f612c20612242565b9050612c2c84846110ce565b15612d03575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612c9f611d2c565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050612d08565b5f9150505b92915050565b5f612d33835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6132e5565b905092915050565b612d43610dcc565b612d79576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612d836133e1565b612db9576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612dc3612d7b565b5f612dcc6127f4565b90505f815f015f6101000a81548160ff02191690831515021790555050565b612df3612d7b565b612dfb6133ff565b612e075f5f1b82612269565b5050565b612e13612d7b565b612e1d8282613409565b5050565b5f612e4d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61345a565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b612e7d82613463565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115612ed957612ed3828261352c565b50612ee2565b612ee16135ac565b5b5050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612f3357602002820191905f5260205f20905b815481526020019060010190808311612f1f575b50505050509050919050565b5f5f8290505f5f90505b85859050811015612f8657612f7782878784818110612f6b57612f6a61471f565b5b905060200201356135e8565b91508080600101915050612f49565b50809150509392505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612fbc613612565b612fc4613688565b4630604051602001612fda9594939291906149b5565b60405160208183030381529060405280519060200120905090565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c1115613071575f600385925092509250613112565b5f6001888888886040515f81526020016040526040516130949493929190614a06565b6020604051602081039080840390855afa1580156130b4573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613105575f60015f5f1b93509350935050613112565b805f5f5f1b935093509350505b9450945094915050565b5f600381111561312f5761312e614a49565b5b82600381111561314257613141614a49565b5b031561327a576001600381111561315c5761315b614a49565b5b82600381111561316f5761316e614a49565b5b036131a6576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260038111156131ba576131b9614a49565b5b8260038111156131cd576131cc614a49565b5b0361321157805f1c6040517ffce698f70000000000000000000000000000000000000000000000000000000081526004016132089190613cf8565b60405180910390fd5b60038081111561322457613223614a49565b5b82600381111561323757613236614a49565b5b0361327957806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016132709190613d4b565b60405180910390fd5b5b5050565b5f61328983836136ff565b6132db57825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f2081905550600190506132df565b5f90505b92915050565b5f5f836001015f8481526020019081526020015f205490505f81146133d6575f60018261331291906147e1565b90505f6001865f018054905061332891906147e1565b905080821461338e575f865f0182815481106133475761334661471f565b5b905f5260205f200154905080875f0184815481106133685761336761471f565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f018054806133a1576133a0614a76565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506133db565b5f9150505b92915050565b5f6133ea612420565b5f0160089054906101000a900460ff16905090565b613407612d7b565b565b613411612d7b565b5f61341a61288a565b90508281600201908161342d9190614c3a565b508181600301908161343f9190614c3a565b505f5f1b815f01819055505f5f1b8160010181905550505050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b036134be57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016134b59190614900565b60405180910390fd5b806134ea7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61345a565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516135559190614d4d565b5f60405180830381855af49150503d805f811461358d576040519150601f19603f3d011682016040523d82523d5f602084013e613592565b606091505b50915091506135a285838361371f565b9250505092915050565b5f3411156135e6576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8183106135ff576135fa82846137ac565b61360a565b61360983836137ac565b5b905092915050565b5f5f61361c61288a565b90505f6136276128b1565b90505f8151111561364357808051906020012092505050613685565b5f825f015490505f5f1b811461365e57809350505050613685565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f61369261288a565b90505f61369d61294f565b90505f815111156136b9578080519060200120925050506136fc565b5f826001015490505f5f1b81146136d5578093505050506136fc565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f836001015f8481526020019081526020015f20541415905092915050565b6060826137345761372f826137c0565b6137a4565b5f825114801561375a57505f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561379c57836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016137939190614900565b60405180910390fd5b8190506137a5565b5b9392505050565b5f825f528160205260405f20905092915050565b5f815111156137d25780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6138b781613883565b81146138c1575f5ffd5b50565b5f813590506138d2816138ae565b92915050565b5f602082840312156138ed576138ec61387b565b5b5f6138fa848285016138c4565b91505092915050565b5f8115159050919050565b61391781613903565b82525050565b5f6020820190506139305f83018461390e565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261395757613956613936565b5b8235905067ffffffffffffffff8111156139745761397361393a565b5b6020830191508360018202830111156139905761398f61393e565b5b9250929050565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6139dd82613997565b810181811067ffffffffffffffff821117156139fc576139fb6139a7565b5b80604052505050565b5f613a0e613872565b9050613a1a82826139d4565b919050565b5f67ffffffffffffffff821115613a3957613a386139a7565b5b602082029050602081019050919050565b5f60ff82169050919050565b613a5f81613a4a565b8114613a69575f5ffd5b50565b5f81359050613a7a81613a56565b92915050565b5f613a92613a8d84613a1f565b613a05565b90508083825260208201905060208402830185811115613ab557613ab461393e565b5b835b81811015613ade5780613aca8882613a6c565b845260208401935050602081019050613ab7565b5050509392505050565b5f82601f830112613afc57613afb613936565b5b8135613b0c848260208601613a80565b91505092915050565b5f67ffffffffffffffff821115613b2f57613b2e6139a7565b5b602082029050602081019050919050565b5f819050919050565b613b5281613b40565b8114613b5c575f5ffd5b50565b5f81359050613b6d81613b49565b92915050565b5f613b85613b8084613b15565b613a05565b90508083825260208201905060208402830185811115613ba857613ba761393e565b5b835b81811015613bd15780613bbd8882613b5f565b845260208401935050602081019050613baa565b5050509392505050565b5f82601f830112613bef57613bee613936565b5b8135613bff848260208601613b73565b91505092915050565b5f5f5f5f5f60808688031215613c2157613c2061387b565b5b5f86013567ffffffffffffffff811115613c3e57613c3d61387f565b5b613c4a88828901613942565b9550955050602086013567ffffffffffffffff811115613c6d57613c6c61387f565b5b613c7988828901613ae8565b935050604086013567ffffffffffffffff811115613c9a57613c9961387f565b5b613ca688828901613bdb565b925050606086013567ffffffffffffffff811115613cc757613cc661387f565b5b613cd388828901613bdb565b9150509295509295909350565b5f819050919050565b613cf281613ce0565b82525050565b5f602082019050613d0b5f830184613ce9565b92915050565b5f60208284031215613d2657613d2561387b565b5b5f613d3384828501613b5f565b91505092915050565b613d4581613b40565b82525050565b5f602082019050613d5e5f830184613d3c565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613d8d82613d64565b9050919050565b613d9d81613d83565b8114613da7575f5ffd5b50565b5f81359050613db881613d94565b92915050565b5f5f60408385031215613dd457613dd361387b565b5b5f613de185828601613b5f565b9250506020613df285828601613daa565b9150509250929050565b613e0581613a4a565b82525050565b5f602082019050613e1e5f830184613dfc565b92915050565b5f60208284031215613e3957613e3861387b565b5b5f613e4684828501613a6c565b91505092915050565b5f604082019050613e625f830185613ce9565b613e6f6020830184613ce9565b9392505050565b5f5ffd5b5f67ffffffffffffffff821115613e9457613e936139a7565b5b613e9d82613997565b9050602081019050919050565b828183375f83830152505050565b5f613eca613ec584613e7a565b613a05565b905082815260208101848484011115613ee657613ee5613e76565b5b613ef1848285613eaa565b509392505050565b5f82601f830112613f0d57613f0c613936565b5b8135613f1d848260208601613eb8565b91505092915050565b5f5f60408385031215613f3c57613f3b61387b565b5b5f613f4985828601613daa565b925050602083013567ffffffffffffffff811115613f6a57613f6961387f565b5b613f7685828601613ef9565b9150509250929050565b5f67ffffffffffffffff821115613f9a57613f996139a7565b5b602082029050602081019050919050565b5f613fbd613fb884613f80565b613a05565b90508083825260208201905060208402830185811115613fe057613fdf61393e565b5b835b818110156140095780613ff58882613daa565b845260208401935050602081019050613fe2565b5050509392505050565b5f82601f83011261402757614026613936565b5b8135614037848260208601613fab565b91505092915050565b5f5f604083850312156140565761405561387b565b5b5f61406385828601613b5f565b925050602083013567ffffffffffffffff8111156140845761408361387f565b5b61409085828601614013565b9150509250929050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6140ce8161409a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614106826140d4565b61411081856140de565b93506141208185602086016140ee565b61412981613997565b840191505092915050565b61413d81613d83565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61417581613ce0565b82525050565b5f614186838361416c565b60208301905092915050565b5f602082019050919050565b5f6141a882614143565b6141b2818561414d565b93506141bd8361415d565b805f5b838110156141ed5781516141d4888261417b565b97506141df83614192565b9250506001810190506141c0565b5085935050505092915050565b5f60e08201905061420d5f83018a6140c5565b818103602083015261421f81896140fc565b9050818103604083015261423381886140fc565b90506142426060830187613ce9565b61424f6080830186614134565b61425c60a0830185613d3c565b81810360c083015261426e818461419e565b905098975050505050505050565b61428581613ce0565b811461428f575f5ffd5b50565b5f813590506142a08161427c565b92915050565b5f602082840312156142bb576142ba61387b565b5b5f6142c884828501614292565b91505092915050565b5f67ffffffffffffffff82169050919050565b6142ed816142d1565b82525050565b6142fc81613b40565b82525050565b61430b81613d83565b82525050565b60c082015f8201516143255f8501826142e4565b50602082015161433860208501826142e4565b50604082015161434b60408501826142e4565b50606082015161435e60608501826142e4565b50608082015161437160808501826142f3565b5060a082015161438460a0850182614302565b50505050565b5f60c08201905061439d5f830184614311565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6143d78383614302565b60208301905092915050565b5f602082019050919050565b5f6143f9826143a3565b61440381856143ad565b935061440e836143bd565b805f5b8381101561443e57815161442588826143cc565b9750614430836143e3565b925050600181019050614411565b5085935050505092915050565b5f6020820190508181035f83015261446381846143ef565b905092915050565b5f6020820190508181035f83015261448381846140fc565b905092915050565b5f5f83601f8401126144a05761449f613936565b5b8235905067ffffffffffffffff8111156144bd576144bc61393a565b5b6020830191508360208202830111156144d9576144d861393e565b5b9250929050565b5f5f5f5f606085870312156144f8576144f761387b565b5b5f61450587828801614292565b945050602085013567ffffffffffffffff8111156145265761452561387f565b5b6145328782880161448b565b9350935050604061454587828801613b5f565b91505092959194509250565b5f5ffd5b5f60a0828403121561456a57614569614551565b5b61457460a0613a05565b90505f61458384828501614292565b5f83015250602061459684828501614292565b60208301525060406145aa84828501614292565b60408301525060606145be84828501613b5f565b60608301525060806145d284828501614292565b60808301525092915050565b5f60a082840312156145f3576145f261387b565b5b5f61460084828501614555565b91505092915050565b5f60c08201905061461c5f830189613d3c565b6146296020830188613ce9565b6146366040830187613ce9565b6146436060830186613ce9565b6146506080830185613d3c565b61465d60a0830184613ce9565b979650505050505050565b5f60408201905061467b5f830185613ce9565b6146886020830184613d3c565b9392505050565b5f6060820190506146a25f830186613ce9565b6146af6020830185613ce9565b6146bc6040830184613dfc565b949350505050565b5f819050919050565b5f819050919050565b5f6146f06146eb6146e6846146c4565b6146cd565b6142d1565b9050919050565b614700816146d6565b82525050565b5f6020820190506147195f8301846146f7565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4549503731323a20556e696e697469616c697a656400000000000000000000005f82015250565b5f6147806015836140de565b915061478b8261474c565b602082019050919050565b5f6020820190508181035f8301526147ad81614774565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6147eb82613ce0565b91506147f683613ce0565b925082820390508181111561480e5761480d6147b4565b5b92915050565b5f61481e82613ce0565b915061482983613ce0565b9250828201905080821115614841576148406147b4565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61487e82613ce0565b915061488983613ce0565b92508261489957614898614847565b5b828204905092915050565b5f6060820190506148b75f830186613ce9565b6148c46020830185613ce9565b6148d16040830184613ce9565b949350505050565b5f6040820190506148ec5f830185614134565b6148f96020830184613d3c565b9392505050565b5f6020820190506149135f830184614134565b92915050565b5f8151905061492781613b49565b92915050565b5f602082840312156149425761494161387b565b5b5f61494f84828501614919565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061499c57607f821691505b6020821081036149af576149ae614958565b5b50919050565b5f60a0820190506149c85f830188613d3c565b6149d56020830187613d3c565b6149e26040830186613d3c565b6149ef6060830185613ce9565b6149fc6080830184614134565b9695505050505050565b5f608082019050614a195f830187613d3c565b614a266020830186613dfc565b614a336040830185613d3c565b614a406060830184613d3c565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302614aff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614ac4565b614b098683614ac4565b95508019841693508086168417925050509392505050565b5f614b3b614b36614b3184613ce0565b6146cd565b613ce0565b9050919050565b5f819050919050565b614b5483614b21565b614b68614b6082614b42565b848454614ad0565b825550505050565b5f5f905090565b614b7f614b70565b614b8a818484614b4b565b505050565b5b81811015614bad57614ba25f82614b77565b600181019050614b90565b5050565b601f821115614bf257614bc381614aa3565b614bcc84614ab5565b81016020851015614bdb578190505b614bef614be785614ab5565b830182614b8f565b50505b505050565b5f82821c905092915050565b5f614c125f1984600802614bf7565b1980831691505092915050565b5f614c2a8383614c03565b9150826002028217905092915050565b614c43826140d4565b67ffffffffffffffff811115614c5c57614c5b6139a7565b5b614c668254614985565b614c71828285614bb1565b5f60209050601f831160018114614ca2575f8415614c90578287015190505b614c9a8582614c1f565b865550614d01565b601f198416614cb086614aa3565b5f5b82811015614cd757848901518255600182019150602085019450602081019050614cb2565b86831015614cf45784890151614cf0601f891682614c03565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f81905092915050565b5f614d2782614d09565b614d318185614d13565b9350614d418185602086016140ee565b80840191505092915050565b5f614d588284614d1d565b91508190509291505056fea26469706673582212202f82185f41b39990e018183f5e104cce4b3501255bac628e689d60b1fe78b0bc64736f6c634300081c0033"

// CrossCheckBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossCheckMetaData.Bin instead.
var CrossCheckBin = CrossCheckMetaData.Bin

// DeployCrossCheck deploys a new Ethereum contract, binding an instance of CrossCheck to it.
func DeployCrossCheck(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossCheck, error) {
	parsed, err := CrossCheckMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossCheckBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossCheck{CrossCheckCaller: CrossCheckCaller{contract: contract}, CrossCheckTransactor: CrossCheckTransactor{contract: contract}, CrossCheckFilterer: CrossCheckFilterer{contract: contract}}, nil
}

// CrossCheck is an auto generated Go binding around an Ethereum contract.
type CrossCheck struct {
	CrossCheckCaller     // Read-only binding to the contract
	CrossCheckTransactor // Write-only binding to the contract
	CrossCheckFilterer   // Log filterer for contract events
}

// CrossCheckCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossCheckCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossCheckTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossCheckTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossCheckFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossCheckFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossCheckSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossCheckSession struct {
	Contract     *CrossCheck       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossCheckCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossCheckCallerSession struct {
	Contract *CrossCheckCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrossCheckTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossCheckTransactorSession struct {
	Contract     *CrossCheckTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossCheckRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossCheckRaw struct {
	Contract *CrossCheck // Generic contract binding to access the raw methods on
}

// CrossCheckCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossCheckCallerRaw struct {
	Contract *CrossCheckCaller // Generic read-only contract binding to access the raw methods on
}

// CrossCheckTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossCheckTransactorRaw struct {
	Contract *CrossCheckTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossCheck creates a new instance of CrossCheck, bound to a specific deployed contract.
func NewCrossCheck(address common.Address, backend bind.ContractBackend) (*CrossCheck, error) {
	contract, err := bindCrossCheck(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossCheck{CrossCheckCaller: CrossCheckCaller{contract: contract}, CrossCheckTransactor: CrossCheckTransactor{contract: contract}, CrossCheckFilterer: CrossCheckFilterer{contract: contract}}, nil
}

// NewCrossCheckCaller creates a new read-only instance of CrossCheck, bound to a specific deployed contract.
func NewCrossCheckCaller(address common.Address, caller bind.ContractCaller) (*CrossCheckCaller, error) {
	contract, err := bindCrossCheck(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossCheckCaller{contract: contract}, nil
}

// NewCrossCheckTransactor creates a new write-only instance of CrossCheck, bound to a specific deployed contract.
func NewCrossCheckTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossCheckTransactor, error) {
	contract, err := bindCrossCheck(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossCheckTransactor{contract: contract}, nil
}

// NewCrossCheckFilterer creates a new log filterer instance of CrossCheck, bound to a specific deployed contract.
func NewCrossCheckFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossCheckFilterer, error) {
	contract, err := bindCrossCheck(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossCheckFilterer{contract: contract}, nil
}

// bindCrossCheck binds a generic wrapper to an already deployed contract.
func bindCrossCheck(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossCheckMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossCheck *CrossCheckRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossCheck.Contract.CrossCheckCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossCheck *CrossCheckRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossCheck.Contract.CrossCheckTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossCheck *CrossCheckRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossCheck.Contract.CrossCheckTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossCheck *CrossCheckCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossCheck.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossCheck *CrossCheckTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossCheck.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossCheck *CrossCheckTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossCheck.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossCheck.Contract.DEFAULTADMINROLE(&_CrossCheck.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossCheck.Contract.DEFAULTADMINROLE(&_CrossCheck.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) OPERATORROLE() ([32]byte, error) {
	return _CrossCheck.Contract.OPERATORROLE(&_CrossCheck.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) OPERATORROLE() ([32]byte, error) {
	return _CrossCheck.Contract.OPERATORROLE(&_CrossCheck.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossCheck *CrossCheckCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossCheck *CrossCheckSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossCheck.Contract.UPGRADEINTERFACEVERSION(&_CrossCheck.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossCheck *CrossCheckCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossCheck.Contract.UPGRADEINTERFACEVERSION(&_CrossCheck.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) VALIDATORROLE() ([32]byte, error) {
	return _CrossCheck.Contract.VALIDATORROLE(&_CrossCheck.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _CrossCheck.Contract.VALIDATORROLE(&_CrossCheck.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) DomainSeparator() ([32]byte, error) {
	return _CrossCheck.Contract.DomainSeparator(&_CrossCheck.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) DomainSeparator() ([32]byte, error) {
	return _CrossCheck.Contract.DomainSeparator(&_CrossCheck.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossCheck *CrossCheckCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossCheck *CrossCheckSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossCheck.Contract.Eip712Domain(&_CrossCheck.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossCheck *CrossCheckCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossCheck.Contract.Eip712Domain(&_CrossCheck.CallOpts)
}

// GetCheckBlock is a free data retrieval call binding the contract method 0x871a782e.
//
// Solidity: function getCheckBlock(uint256 startBlockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCaller) GetCheckBlock(opts *bind.CallOpts, startBlockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getCheckBlock", startBlockNumber)

	if err != nil {
		return *new(CrossCheckBlockCheckBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(CrossCheckBlockCheckBlock)).(*CrossCheckBlockCheckBlock)

	return out0, err

}

// GetCheckBlock is a free data retrieval call binding the contract method 0x871a782e.
//
// Solidity: function getCheckBlock(uint256 startBlockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckSession) GetCheckBlock(startBlockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlock(&_CrossCheck.CallOpts, startBlockNumber)
}

// GetCheckBlock is a free data retrieval call binding the contract method 0x871a782e.
//
// Solidity: function getCheckBlock(uint256 startBlockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCallerSession) GetCheckBlock(startBlockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlock(&_CrossCheck.CallOpts, startBlockNumber)
}

// GetCheckBlockByBlockNumber is a free data retrieval call binding the contract method 0x9f3e14cf.
//
// Solidity: function getCheckBlockByBlockNumber(uint256 blockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCaller) GetCheckBlockByBlockNumber(opts *bind.CallOpts, blockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getCheckBlockByBlockNumber", blockNumber)

	if err != nil {
		return *new(CrossCheckBlockCheckBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(CrossCheckBlockCheckBlock)).(*CrossCheckBlockCheckBlock)

	return out0, err

}

// GetCheckBlockByBlockNumber is a free data retrieval call binding the contract method 0x9f3e14cf.
//
// Solidity: function getCheckBlockByBlockNumber(uint256 blockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckSession) GetCheckBlockByBlockNumber(blockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlockByBlockNumber(&_CrossCheck.CallOpts, blockNumber)
}

// GetCheckBlockByBlockNumber is a free data retrieval call binding the contract method 0x9f3e14cf.
//
// Solidity: function getCheckBlockByBlockNumber(uint256 blockNumber) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCallerSession) GetCheckBlockByBlockNumber(blockNumber *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlockByBlockNumber(&_CrossCheck.CallOpts, blockNumber)
}

// GetCheckBlockByNonce is a free data retrieval call binding the contract method 0xcf7910b6.
//
// Solidity: function getCheckBlockByNonce(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCaller) GetCheckBlockByNonce(opts *bind.CallOpts, nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getCheckBlockByNonce", nonce)

	if err != nil {
		return *new(CrossCheckBlockCheckBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(CrossCheckBlockCheckBlock)).(*CrossCheckBlockCheckBlock)

	return out0, err

}

// GetCheckBlockByNonce is a free data retrieval call binding the contract method 0xcf7910b6.
//
// Solidity: function getCheckBlockByNonce(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckSession) GetCheckBlockByNonce(nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlockByNonce(&_CrossCheck.CallOpts, nonce)
}

// GetCheckBlockByNonce is a free data retrieval call binding the contract method 0xcf7910b6.
//
// Solidity: function getCheckBlockByNonce(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCallerSession) GetCheckBlockByNonce(nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlockByNonce(&_CrossCheck.CallOpts, nonce)
}

// GetNextBlockInfo is a free data retrieval call binding the contract method 0x4d581907.
//
// Solidity: function getNextBlockInfo() view returns(uint256 nextNonce, uint256 nextStart)
func (_CrossCheck *CrossCheckCaller) GetNextBlockInfo(opts *bind.CallOpts) (struct {
	NextNonce *big.Int
	NextStart *big.Int
}, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getNextBlockInfo")

	outstruct := new(struct {
		NextNonce *big.Int
		NextStart *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NextNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextStart = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNextBlockInfo is a free data retrieval call binding the contract method 0x4d581907.
//
// Solidity: function getNextBlockInfo() view returns(uint256 nextNonce, uint256 nextStart)
func (_CrossCheck *CrossCheckSession) GetNextBlockInfo() (struct {
	NextNonce *big.Int
	NextStart *big.Int
}, error) {
	return _CrossCheck.Contract.GetNextBlockInfo(&_CrossCheck.CallOpts)
}

// GetNextBlockInfo is a free data retrieval call binding the contract method 0x4d581907.
//
// Solidity: function getNextBlockInfo() view returns(uint256 nextNonce, uint256 nextStart)
func (_CrossCheck *CrossCheckCallerSession) GetNextBlockInfo() (struct {
	NextNonce *big.Int
	NextStart *big.Int
}, error) {
	return _CrossCheck.Contract.GetNextBlockInfo(&_CrossCheck.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossCheck *CrossCheckSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossCheck.Contract.GetRoleAdmin(&_CrossCheck.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossCheck.Contract.GetRoleAdmin(&_CrossCheck.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossCheck *CrossCheckCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossCheck *CrossCheckSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossCheck.Contract.GetRoleMembers(&_CrossCheck.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CrossCheck *CrossCheckCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CrossCheck.Contract.GetRoleMembers(&_CrossCheck.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossCheck *CrossCheckCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossCheck *CrossCheckSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossCheck.Contract.HasRole(&_CrossCheck.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossCheck *CrossCheckCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossCheck.Contract.HasRole(&_CrossCheck.CallOpts, role, account)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_CrossCheck *CrossCheckCaller) LatestBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "latestBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_CrossCheck *CrossCheckSession) LatestBlock() (*big.Int, error) {
	return _CrossCheck.Contract.LatestBlock(&_CrossCheck.CallOpts)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_CrossCheck *CrossCheckCallerSession) LatestBlock() (*big.Int, error) {
	return _CrossCheck.Contract.LatestBlock(&_CrossCheck.CallOpts)
}

// NumCheckBlocks is a free data retrieval call binding the contract method 0x17894ff4.
//
// Solidity: function numCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckCaller) NumCheckBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "numCheckBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumCheckBlocks is a free data retrieval call binding the contract method 0x17894ff4.
//
// Solidity: function numCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckSession) NumCheckBlocks() (*big.Int, error) {
	return _CrossCheck.Contract.NumCheckBlocks(&_CrossCheck.CallOpts)
}

// NumCheckBlocks is a free data retrieval call binding the contract method 0x17894ff4.
//
// Solidity: function numCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckCallerSession) NumCheckBlocks() (*big.Int, error) {
	return _CrossCheck.Contract.NumCheckBlocks(&_CrossCheck.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossCheck *CrossCheckCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossCheck *CrossCheckSession) Paused() (bool, error) {
	return _CrossCheck.Contract.Paused(&_CrossCheck.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossCheck *CrossCheckCallerSession) Paused() (bool, error) {
	return _CrossCheck.Contract.Paused(&_CrossCheck.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) ProxiableUUID() ([32]byte, error) {
	return _CrossCheck.Contract.ProxiableUUID(&_CrossCheck.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossCheck.Contract.ProxiableUUID(&_CrossCheck.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossCheck *CrossCheckCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossCheck *CrossCheckSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossCheck.Contract.SupportsInterface(&_CrossCheck.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossCheck *CrossCheckCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossCheck.Contract.SupportsInterface(&_CrossCheck.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossCheck *CrossCheckCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossCheck *CrossCheckSession) Threshold() (uint8, error) {
	return _CrossCheck.Contract.Threshold(&_CrossCheck.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_CrossCheck *CrossCheckCallerSession) Threshold() (uint8, error) {
	return _CrossCheck.Contract.Threshold(&_CrossCheck.CallOpts)
}

// VerifyBlock is a free data retrieval call binding the contract method 0xaea3679f.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes32[] proof, bytes32 blockHash) view returns(bool)
func (_CrossCheck *CrossCheckCaller) VerifyBlock(opts *bind.CallOpts, blockNumber *big.Int, proof [][32]byte, blockHash [32]byte) (bool, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "verifyBlock", blockNumber, proof, blockHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlock is a free data retrieval call binding the contract method 0xaea3679f.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes32[] proof, bytes32 blockHash) view returns(bool)
func (_CrossCheck *CrossCheckSession) VerifyBlock(blockNumber *big.Int, proof [][32]byte, blockHash [32]byte) (bool, error) {
	return _CrossCheck.Contract.VerifyBlock(&_CrossCheck.CallOpts, blockNumber, proof, blockHash)
}

// VerifyBlock is a free data retrieval call binding the contract method 0xaea3679f.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes32[] proof, bytes32 blockHash) view returns(bool)
func (_CrossCheck *CrossCheckCallerSession) VerifyBlock(blockNumber *big.Int, proof [][32]byte, blockHash [32]byte) (bool, error) {
	return _CrossCheck.Contract.VerifyBlock(&_CrossCheck.CallOpts, blockNumber, proof, blockHash)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossCheck *CrossCheckTransactor) ChangeThreshold(opts *bind.TransactOpts, threshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "changeThreshold", threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossCheck *CrossCheckSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.ChangeThreshold(&_CrossCheck.TransactOpts, threshold_)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0xb7f3358d.
//
// Solidity: function changeThreshold(uint8 threshold_) returns()
func (_CrossCheck *CrossCheckTransactorSession) ChangeThreshold(threshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.ChangeThreshold(&_CrossCheck.TransactOpts, threshold_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.GrantRole(&_CrossCheck.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.GrantRole(&_CrossCheck.TransactOpts, role, account)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckTransactor) GrantRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "grantRoleBatch", role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.GrantRoleBatch(&_CrossCheck.TransactOpts, role, accounts)
}

// GrantRoleBatch is a paid mutator transaction binding the contract method 0xfd47852b.
//
// Solidity: function grantRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckTransactorSession) GrantRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.GrantRoleBatch(&_CrossCheck.TransactOpts, role, accounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 validatorThreshold) returns()
func (_CrossCheck *CrossCheckTransactor) Initialize(opts *bind.TransactOpts, validatorThreshold uint8) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "initialize", validatorThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 validatorThreshold) returns()
func (_CrossCheck *CrossCheckSession) Initialize(validatorThreshold uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.Initialize(&_CrossCheck.TransactOpts, validatorThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x4351e6b6.
//
// Solidity: function initialize(uint8 validatorThreshold) returns()
func (_CrossCheck *CrossCheckTransactorSession) Initialize(validatorThreshold uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.Initialize(&_CrossCheck.TransactOpts, validatorThreshold)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossCheck *CrossCheckTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossCheck *CrossCheckSession) Pause() (*types.Transaction, error) {
	return _CrossCheck.Contract.Pause(&_CrossCheck.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossCheck *CrossCheckTransactorSession) Pause() (*types.Transaction, error) {
	return _CrossCheck.Contract.Pause(&_CrossCheck.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossCheck *CrossCheckTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossCheck *CrossCheckSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RenounceRole(&_CrossCheck.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CrossCheck *CrossCheckTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RenounceRole(&_CrossCheck.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RevokeRole(&_CrossCheck.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossCheck *CrossCheckTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RevokeRole(&_CrossCheck.TransactOpts, role, account)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckTransactor) RevokeRoleBatch(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "revokeRoleBatch", role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RevokeRoleBatch(&_CrossCheck.TransactOpts, role, accounts)
}

// RevokeRoleBatch is a paid mutator transaction binding the contract method 0x82f51fa6.
//
// Solidity: function revokeRoleBatch(bytes32 role, address[] accounts) returns()
func (_CrossCheck *CrossCheckTransactorSession) RevokeRoleBatch(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _CrossCheck.Contract.RevokeRoleBatch(&_CrossCheck.TransactOpts, role, accounts)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x07a9eef2.
//
// Solidity: function submitCheckpoint(bytes data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_CrossCheck *CrossCheckTransactor) SubmitCheckpoint(opts *bind.TransactOpts, data []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "submitCheckpoint", data, v, r, s)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x07a9eef2.
//
// Solidity: function submitCheckpoint(bytes data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_CrossCheck *CrossCheckSession) SubmitCheckpoint(data []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossCheck.Contract.SubmitCheckpoint(&_CrossCheck.TransactOpts, data, v, r, s)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x07a9eef2.
//
// Solidity: function submitCheckpoint(bytes data, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_CrossCheck *CrossCheckTransactorSession) SubmitCheckpoint(data []byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _CrossCheck.Contract.SubmitCheckpoint(&_CrossCheck.TransactOpts, data, v, r, s)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossCheck *CrossCheckTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossCheck *CrossCheckSession) Unpause() (*types.Transaction, error) {
	return _CrossCheck.Contract.Unpause(&_CrossCheck.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossCheck *CrossCheckTransactorSession) Unpause() (*types.Transaction, error) {
	return _CrossCheck.Contract.Unpause(&_CrossCheck.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossCheck *CrossCheckTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossCheck *CrossCheckSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossCheck.Contract.UpgradeToAndCall(&_CrossCheck.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossCheck *CrossCheckTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossCheck.Contract.UpgradeToAndCall(&_CrossCheck.TransactOpts, newImplementation, data)
}

// CrossCheckCrossCheckInitializedIterator is returned from FilterCrossCheckInitialized and is used to iterate over the raw logs and unpacked data for CrossCheckInitialized events raised by the CrossCheck contract.
type CrossCheckCrossCheckInitializedIterator struct {
	Event *CrossCheckCrossCheckInitialized // Event containing the contract specifics and raw log

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
func (it *CrossCheckCrossCheckInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckCrossCheckInitialized)
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
		it.Event = new(CrossCheckCrossCheckInitialized)
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
func (it *CrossCheckCrossCheckInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckCrossCheckInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckCrossCheckInitialized represents a CrossCheckInitialized event raised by the CrossCheck contract.
type CrossCheckCrossCheckInitialized struct {
	ChainID            *big.Int
	BlocksPerCheck     *big.Int
	ValidatorThreshold uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCrossCheckInitialized is a free log retrieval operation binding the contract event 0xb80103afdb269ad46586734e4733e996c5fd0ae2882ba36f5f8f2033ae8be277.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint8 validatorThreshold)
func (_CrossCheck *CrossCheckFilterer) FilterCrossCheckInitialized(opts *bind.FilterOpts) (*CrossCheckCrossCheckInitializedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CrossCheckInitialized")
	if err != nil {
		return nil, err
	}
	return &CrossCheckCrossCheckInitializedIterator{contract: _CrossCheck.contract, event: "CrossCheckInitialized", logs: logs, sub: sub}, nil
}

// WatchCrossCheckInitialized is a free log subscription operation binding the contract event 0xb80103afdb269ad46586734e4733e996c5fd0ae2882ba36f5f8f2033ae8be277.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint8 validatorThreshold)
func (_CrossCheck *CrossCheckFilterer) WatchCrossCheckInitialized(opts *bind.WatchOpts, sink chan<- *CrossCheckCrossCheckInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "CrossCheckInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckCrossCheckInitialized)
				if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckInitialized", log); err != nil {
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

// ParseCrossCheckInitialized is a log parse operation binding the contract event 0xb80103afdb269ad46586734e4733e996c5fd0ae2882ba36f5f8f2033ae8be277.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint8 validatorThreshold)
func (_CrossCheck *CrossCheckFilterer) ParseCrossCheckInitialized(log types.Log) (*CrossCheckCrossCheckInitialized, error) {
	event := new(CrossCheckCrossCheckInitialized)
	if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossCheck contract.
type CrossCheckEIP712DomainChangedIterator struct {
	Event *CrossCheckEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossCheckEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckEIP712DomainChanged)
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
		it.Event = new(CrossCheckEIP712DomainChanged)
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
func (it *CrossCheckEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckEIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossCheck contract.
type CrossCheckEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossCheck *CrossCheckFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossCheckEIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossCheckEIP712DomainChangedIterator{contract: _CrossCheck.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossCheck *CrossCheckFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossCheckEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckEIP712DomainChanged)
				if err := _CrossCheck.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseEIP712DomainChanged(log types.Log) (*CrossCheckEIP712DomainChanged, error) {
	event := new(CrossCheckEIP712DomainChanged)
	if err := _CrossCheck.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossCheck contract.
type CrossCheckInitializedIterator struct {
	Event *CrossCheckInitialized // Event containing the contract specifics and raw log

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
func (it *CrossCheckInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckInitialized)
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
		it.Event = new(CrossCheckInitialized)
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
func (it *CrossCheckInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckInitialized represents a Initialized event raised by the CrossCheck contract.
type CrossCheckInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossCheck *CrossCheckFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossCheckInitializedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossCheckInitializedIterator{contract: _CrossCheck.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossCheck *CrossCheckFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossCheckInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckInitialized)
				if err := _CrossCheck.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseInitialized(log types.Log) (*CrossCheckInitialized, error) {
	event := new(CrossCheckInitialized)
	if err := _CrossCheck.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckNewCheckBlockIterator is returned from FilterNewCheckBlock and is used to iterate over the raw logs and unpacked data for NewCheckBlock events raised by the CrossCheck contract.
type CrossCheckNewCheckBlockIterator struct {
	Event *CrossCheckNewCheckBlock // Event containing the contract specifics and raw log

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
func (it *CrossCheckNewCheckBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckNewCheckBlock)
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
		it.Event = new(CrossCheckNewCheckBlock)
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
func (it *CrossCheckNewCheckBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckNewCheckBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckNewCheckBlock represents a NewCheckBlock event raised by the CrossCheck contract.
type CrossCheckNewCheckBlock struct {
	Proposer common.Address
	Nonce    *big.Int
	Start    *big.Int
	End      *big.Int
	RootHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewCheckBlock is a free log retrieval operation binding the contract event 0xbb765a5098e4a96564f8633892153a94a950dadf5b371b7b46fa9fa170e034c2.
//
// Solidity: event NewCheckBlock(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) FilterNewCheckBlock(opts *bind.FilterOpts, proposer []common.Address, nonce []*big.Int, start []*big.Int) (*CrossCheckNewCheckBlockIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var startRule []interface{}
	for _, startItem := range start {
		startRule = append(startRule, startItem)
	}

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "NewCheckBlock", proposerRule, nonceRule, startRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckNewCheckBlockIterator{contract: _CrossCheck.contract, event: "NewCheckBlock", logs: logs, sub: sub}, nil
}

// WatchNewCheckBlock is a free log subscription operation binding the contract event 0xbb765a5098e4a96564f8633892153a94a950dadf5b371b7b46fa9fa170e034c2.
//
// Solidity: event NewCheckBlock(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) WatchNewCheckBlock(opts *bind.WatchOpts, sink chan<- *CrossCheckNewCheckBlock, proposer []common.Address, nonce []*big.Int, start []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var startRule []interface{}
	for _, startItem := range start {
		startRule = append(startRule, startItem)
	}

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "NewCheckBlock", proposerRule, nonceRule, startRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckNewCheckBlock)
				if err := _CrossCheck.contract.UnpackLog(event, "NewCheckBlock", log); err != nil {
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

// ParseNewCheckBlock is a log parse operation binding the contract event 0xbb765a5098e4a96564f8633892153a94a950dadf5b371b7b46fa9fa170e034c2.
//
// Solidity: event NewCheckBlock(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) ParseNewCheckBlock(log types.Log) (*CrossCheckNewCheckBlock, error) {
	event := new(CrossCheckNewCheckBlock)
	if err := _CrossCheck.contract.UnpackLog(event, "NewCheckBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossCheck contract.
type CrossCheckPausedIterator struct {
	Event *CrossCheckPaused // Event containing the contract specifics and raw log

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
func (it *CrossCheckPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckPaused)
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
		it.Event = new(CrossCheckPaused)
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
func (it *CrossCheckPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckPaused represents a Paused event raised by the CrossCheck contract.
type CrossCheckPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossCheck *CrossCheckFilterer) FilterPaused(opts *bind.FilterOpts) (*CrossCheckPausedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossCheckPausedIterator{contract: _CrossCheck.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossCheck *CrossCheckFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossCheckPaused) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckPaused)
				if err := _CrossCheck.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossCheck *CrossCheckFilterer) ParsePaused(log types.Log) (*CrossCheckPaused, error) {
	event := new(CrossCheckPaused)
	if err := _CrossCheck.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossCheck contract.
type CrossCheckRoleAdminChangedIterator struct {
	Event *CrossCheckRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossCheckRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckRoleAdminChanged)
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
		it.Event = new(CrossCheckRoleAdminChanged)
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
func (it *CrossCheckRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckRoleAdminChanged represents a RoleAdminChanged event raised by the CrossCheck contract.
type CrossCheckRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossCheck *CrossCheckFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossCheckRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckRoleAdminChangedIterator{contract: _CrossCheck.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossCheck *CrossCheckFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossCheckRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckRoleAdminChanged)
				if err := _CrossCheck.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseRoleAdminChanged(log types.Log) (*CrossCheckRoleAdminChanged, error) {
	event := new(CrossCheckRoleAdminChanged)
	if err := _CrossCheck.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossCheck contract.
type CrossCheckRoleGrantedIterator struct {
	Event *CrossCheckRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossCheckRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckRoleGranted)
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
		it.Event = new(CrossCheckRoleGranted)
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
func (it *CrossCheckRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckRoleGranted represents a RoleGranted event raised by the CrossCheck contract.
type CrossCheckRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossCheck *CrossCheckFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossCheckRoleGrantedIterator, error) {

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

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckRoleGrantedIterator{contract: _CrossCheck.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossCheck *CrossCheckFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossCheckRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckRoleGranted)
				if err := _CrossCheck.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseRoleGranted(log types.Log) (*CrossCheckRoleGranted, error) {
	event := new(CrossCheckRoleGranted)
	if err := _CrossCheck.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossCheck contract.
type CrossCheckRoleRevokedIterator struct {
	Event *CrossCheckRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossCheckRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckRoleRevoked)
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
		it.Event = new(CrossCheckRoleRevoked)
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
func (it *CrossCheckRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckRoleRevoked represents a RoleRevoked event raised by the CrossCheck contract.
type CrossCheckRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossCheck *CrossCheckFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossCheckRoleRevokedIterator, error) {

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

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckRoleRevokedIterator{contract: _CrossCheck.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossCheck *CrossCheckFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossCheckRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckRoleRevoked)
				if err := _CrossCheck.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseRoleRevoked(log types.Log) (*CrossCheckRoleRevoked, error) {
	event := new(CrossCheckRoleRevoked)
	if err := _CrossCheck.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossCheck contract.
type CrossCheckThresholdChangedIterator struct {
	Event *CrossCheckThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossCheckThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckThresholdChanged)
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
		it.Event = new(CrossCheckThresholdChanged)
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
func (it *CrossCheckThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckThresholdChanged represents a ThresholdChanged event raised by the CrossCheck contract.
type CrossCheckThresholdChanged struct {
	Threshold uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossCheck *CrossCheckFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossCheckThresholdChangedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossCheckThresholdChangedIterator{contract: _CrossCheck.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossCheck *CrossCheckFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossCheckThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckThresholdChanged)
				if err := _CrossCheck.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff.
//
// Solidity: event ThresholdChanged(uint8 threshold)
func (_CrossCheck *CrossCheckFilterer) ParseThresholdChanged(log types.Log) (*CrossCheckThresholdChanged, error) {
	event := new(CrossCheckThresholdChanged)
	if err := _CrossCheck.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossCheck contract.
type CrossCheckUnpausedIterator struct {
	Event *CrossCheckUnpaused // Event containing the contract specifics and raw log

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
func (it *CrossCheckUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckUnpaused)
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
		it.Event = new(CrossCheckUnpaused)
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
func (it *CrossCheckUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckUnpaused represents a Unpaused event raised by the CrossCheck contract.
type CrossCheckUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossCheck *CrossCheckFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossCheckUnpausedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossCheckUnpausedIterator{contract: _CrossCheck.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossCheck *CrossCheckFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossCheckUnpaused) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckUnpaused)
				if err := _CrossCheck.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossCheck *CrossCheckFilterer) ParseUnpaused(log types.Log) (*CrossCheckUnpaused, error) {
	event := new(CrossCheckUnpaused)
	if err := _CrossCheck.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossCheck contract.
type CrossCheckUpgradedIterator struct {
	Event *CrossCheckUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossCheckUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckUpgraded)
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
		it.Event = new(CrossCheckUpgraded)
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
func (it *CrossCheckUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckUpgraded represents a Upgraded event raised by the CrossCheck contract.
type CrossCheckUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossCheck *CrossCheckFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossCheckUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckUpgradedIterator{contract: _CrossCheck.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossCheck *CrossCheckFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossCheckUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckUpgraded)
				if err := _CrossCheck.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossCheck *CrossCheckFilterer) ParseUpgraded(log types.Log) (*CrossCheckUpgraded, error) {
	event := new(CrossCheckUpgraded)
	if err := _CrossCheck.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
