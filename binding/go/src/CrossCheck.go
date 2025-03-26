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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blocksPerCheck\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeThreshold\",\"inputs\":[{\"name\":\"threshold_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckBlock\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCrossCheckBlock.CheckBlock\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckBlockByBlockNumber\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCrossCheckBlock.CheckBlock\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"start\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextBlockInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"_nextNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nextStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRoleBatch\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"blocksPerCheck_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCheckBlocks_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validatorThreshold_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxCheckBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numCheckBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRoleBatch\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlocksPerCheck\",\"inputs\":[{\"name\":\"blocksPerCheck_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxCheckBlocks\",\"inputs\":[{\"name\":\"maxCheckBlocks_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPause\",\"inputs\":[{\"name\":\"set\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitCheckpoint\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckBlockAdded\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckBlockPruned\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossCheckBlocksPerCheckUpdated\",\"inputs\":[{\"name\":\"blocksPerCheck\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossCheckInitialized\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blocksPerCheck\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxCheckBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"validatorThreshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossCheckMaxCheckBlocksUpdated\",\"inputs\":[{\"name\":\"maxCheckBlocks\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CrossCheckBlockExists\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckBlockNotFound\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidBlocksPerCheck\",\"inputs\":[{\"name\":\"blocksPerCheck\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidChainID\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CrossCheckInvalidMaxCheckBlocks\",\"inputs\":[{\"name\":\"maxCheckBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckNonceNotFound\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CrossCheckNotNextBlock\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RoleManagerAlreadyHasRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RoleManagerDoesNotHaveRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorInsufficientSignature\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorInvalidSignatures\",\"inputs\":[{\"name\":\"vLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotAuthorized\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ValidatorThresholdCanNotZero\",\"inputs\":[]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525062015180600355606460045534801561004e575f5ffd5b5061005d61006260201b60201c565b6101c2565b5f61007161016060201b60201c565b9050805f0160089054906101000a900460ff16156100bb576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461015d5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161015491906101a9565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f67ffffffffffffffff82169050919050565b6101a381610187565b82525050565b5f6020820190506101bc5f83018461019a565b92915050565b608051614cf36101e85f395f818161209d015281816120f201526122b10152614cf35ff3fe6080604052600436106101ed575f3560e01c806384b0196e1161010c578063b7f3358d1161009f578063d547741f1161006e578063d547741f146106ee578063f5b541a614610716578063f698da2514610740578063fcc6aca91461076a578063fd47852b14610794576101ed565b8063b7f3358d1461064c578063bedb86fb14610674578063c00111141461069c578063c49baebe146106c4576101ed565b8063a217fddf116100db578063a217fddf14610580578063a3246ad3146105aa578063ad3cb1cc146105e6578063aea3679f14610610576101ed565b806384b0196e1461049c578063871a782e146104cc57806391d14854146105085780639f3e14cf14610544576101ed565b80634f1ef286116101845780635c975abb116101535780635c975abb146103f857806366cdc83b1461042257806375b238fc1461044a57806382f51fa614610474576101ed565b80634f1ef2861461036057806352d1902d1461037c578063544d7e1e146103a657806355a84bb4146103d0576101ed565b80632f2ff15d116101c05780632f2ff15d146102bb57806336568abe146102e357806342cde4e81461030b5780634d58190714610335576101ed565b806301ffc9a7146101f157806307a9eef21461022d57806317894ff414610255578063248a9ca31461027f575b5f5ffd5b3480156101fc575f5ffd5b506102176004803603810190610212919061377f565b6107bc565b60405161022491906137c4565b60405180910390f35b348015610238575f5ffd5b50610253600480360381019061024e9190613aaf565b610835565b005b348015610260575f5ffd5b50610269610ac6565b6040516102769190613b9f565b60405180910390f35b34801561028a575f5ffd5b506102a560048036038101906102a09190613bb8565b610adc565b6040516102b29190613bf2565b60405180910390f35b3480156102c6575f5ffd5b506102e160048036038101906102dc9190613c65565b610b06565b005b3480156102ee575f5ffd5b5061030960048036038101906103049190613c65565b610b28565b005b348015610316575f5ffd5b5061031f610ba3565b60405161032c9190613cb2565b60405180910390f35b348015610340575f5ffd5b50610349610bc5565b604051610357929190613ccb565b60405180910390f35b61037a60048036038101906103759190613da2565b610c4e565b005b348015610387575f5ffd5b50610390610c6d565b60405161039d9190613bf2565b60405180910390f35b3480156103b1575f5ffd5b506103ba610c9e565b6040516103c79190613b9f565b60405180910390f35b3480156103db575f5ffd5b506103f660048036038101906103f19190613e26565b610ca4565b005b348015610403575f5ffd5b5061040c610e92565b60405161041991906137c4565b60405180910390f35b34801561042d575f5ffd5b5061044860048036038101906104439190613e76565b610eb4565b005b348015610455575f5ffd5b5061045e610f22565b60405161046b9190613bf2565b60405180910390f35b34801561047f575f5ffd5b5061049a60048036038101906104959190613f61565b610f46565b005b3480156104a7575f5ffd5b506104b0610f9c565b6040516104c3979695949392919061411b565b60405180910390f35b3480156104d7575f5ffd5b506104f260048036038101906104ed9190613e76565b6110a5565b6040516104ff9190614256565b60405180910390f35b348015610513575f5ffd5b5061052e60048036038101906105299190613c65565b6111f1565b60405161053b91906137c4565b60405180910390f35b34801561054f575f5ffd5b5061056a60048036038101906105659190613e76565b611262565b6040516105779190614256565b60405180910390f35b34801561058b575f5ffd5b5061059461162e565b6040516105a19190613bf2565b60405180910390f35b3480156105b5575f5ffd5b506105d060048036038101906105cb9190613bb8565b611634565b6040516105dd9190614317565b60405180910390f35b3480156105f1575f5ffd5b506105fa611663565b6040516106079190614337565b60405180910390f35b34801561061b575f5ffd5b50610636600480360381019061063191906143ac565b61169c565b60405161064391906137c4565b60405180910390f35b348015610657575f5ffd5b50610672600480360381019061066d919061441d565b611716565b005b34801561067f575f5ffd5b5061069a60048036038101906106959190614472565b6117c0565b005b3480156106a7575f5ffd5b506106c260048036038101906106bd9190613e76565b61180a565b005b3480156106cf575f5ffd5b506106d8611878565b6040516106e59190613bf2565b60405180910390f35b3480156106f9575f5ffd5b50610714600480360381019061070f9190613c65565b61189c565b005b348015610721575f5ffd5b5061072a6118be565b6040516107379190613bf2565b60405180910390f35b34801561074b575f5ffd5b506107546118e2565b6040516107619190613bf2565b60405180910390f35b348015610775575f5ffd5b5061077e6118f0565b60405161078b9190613b9f565b60405180910390f35b34801561079f575f5ffd5b506107ba60048036038101906107b59190613f61565b6118f6565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061082e575061082d8261194c565b5b9050919050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892661085f816119b5565b6108676119c9565b5f610870611a0a565b90505f8787810190610882919061452a565b9050620956cc8160800151146108d35780608001516040517fa5acee440000000000000000000000000000000000000000000000000000000081526004016108ca9190613b9f565b60405180910390fd5b806020015181604001511115806108ef57505f5f1b8160600151145b15610926576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61092f610bc5565b9150505f5f905060016003548301039050825f0151600254141580610958575082602001518214155b80610967575082604001518114155b156109b257825f015183602001516040517f99a5875b0000000000000000000000000000000000000000000000000000000081526004016109a9929190613ccb565b60405180910390fd5b5f7ffe27049a99724ab477714cb3e8105a7b7d274fe7e5bee5cf905d63758b108490845f01518560200151866040015187606001518860800151604051602001610a0196959493929190614555565b604051602081830303815290604052805190602001209050610a25818a8a8a611a11565b610a4185855f0151866020015187604001518860600151611c3c565b60025f8154600101919050819055508360200151845f01518673ffffffffffffffffffffffffffffffffffffffff167fbadf6f34b60f61fb703b4f3ffa385fcf9f98bf0a53cfbadd15fe381c30b145bf87604001518860600151604051610aa99291906145b4565b60405180910390a4610ab9611e99565b5050505050505050505050565b5f600154600254610ad79190614608565b905090565b5f5f610ae6611f2b565b9050805f015f8481526020019081526020015f2060010154915050919050565b610b0f82610adc565b610b18816119b5565b610b228383611f52565b50505050565b610b30611a0a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b94576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b9e8282611fe3565b505050565b5f5f610bad612074565b9050805f015f9054906101000a900460ff1691505090565b5f5f60025491505f6002541115610c4a575f5f5f6001600254610be89190614608565b81526020019081526020015f2090505f815f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115610c48576001815f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff160191505b505b9091565b610c5661209b565b610c5f82612181565b610c698282612191565b5050565b5f610c766122af565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b60045481565b5f610cad612336565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610cf55750825b90505f60018367ffffffffffffffff16148015610d2857505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610d36575080155b15610d6d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610dba576001855f0160086101000a81548160ff0219169083151502179055505b610dc261235d565b610dca612367565b610dd333612379565b610ddc8661238d565b610de58861246f565b610dee876124bf565b7f38778a692f0977d8b900e132518cdf1350918c14b717a3f68825d1e067476bc1620956cc898989604051610e26949392919061463b565b60405180910390a18315610e88575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610e7f91906146c0565b60405180910390a15b5050505050505050565b5f5f610e9c61250f565b9050805f015f9054906101000a900460ff1691505090565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610ede816119b5565b610ee7826124bf565b7ff3114416d687e19a2ba81f92f6b8a93a48863d39989e5fd2f1adb0f5f89fd69982604051610f169190613b9f565b60405180910390a15050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b610f4f82610adc565b610f58816119b5565b5f5f90505b8251811015610f9657610f8a84848381518110610f7d57610f7c6146d9565b5b6020026020010151611fe3565b50806001019050610f5d565b50505050565b5f6060805f5f5f60605f610fae612536565b90505f5f1b815f0154148015610fc957505f5f1b8160010154145b611008576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fff90614750565b60405180910390fd5b61101061255d565b6110186125fb565b46305f5f1b5f67ffffffffffffffff8111156110375761103661384e565b5b6040519080825280602002602001820160405280156110655781602001602082028036833780820191505090505b507f0f0000000000000000000000000000000000000000000000000000000000000095949392919097509750975097509750975097505090919293949596565b6110ad6136ab565b5f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b5f5f6111fb611f2b565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b61126a6136ab565b600254600154106112e9576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152509050611629565b5f5f60015481526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16821015611397576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152509050611629565b5f60015490505f60016002546113ad9190614608565b90505b8082116115b7575f600282846113c6919061476e565b6113d091906147ce565b90505f5f5f8381526020019081526020015f20905085815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16111580156114375750805f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168611155b1561157557806040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050945050505050611629565b805f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168610156115a9576001820392506115b0565b6001820193505b50506113b0565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250925050505b919050565b5f5f1b81565b60605f61163f612699565b905061165b815f015f8581526020019081526020015f206126c0565b915050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f5f6116a786611262565b90505f816060015167ffffffffffffffff16036116fb57856040517f96fab5950000000000000000000000000000000000000000000000000000000081526004016116f29190613b9f565b60405180910390fd5b61170b85858360800151866126df565b915050949350505050565b5f5f1b611722816119b5565b5f8260ff160361175e576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611767612074565b905082815f015f6101000a81548160ff021916908360ff1602179055507f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff836040516117b39190613cb2565b60405180910390a1505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9296117ea816119b5565b81156117fd576117f86126f7565b611806565b611805612766565b5b5050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611834816119b5565b61183d8261246f565b7fc37172b02acb2a6b00612c1178f47d9f660e5fa7800fa7a3904e7077665659128260405161186c9190613b9f565b60405180910390a15050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6118a582610adc565b6118ae816119b5565b6118b88383611fe3565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b5f6118eb6127d4565b905090565b60035481565b6118ff82610adc565b611908816119b5565b5f5f90505b82518110156119465761193a8484838151811061192d5761192c6146d9565b5b6020026020010151611f52565b5080600101905061190d565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6119c6816119c1611a0a565b6127e2565b50565b6119d1610e92565b15611a08576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f33905090565b5f611a1a612074565b90505f84519050835181148015611a315750825181145b8185518551909192611a7b576040517f9089eb84000000000000000000000000000000000000000000000000000000008152600401611a72939291906147fe565b60405180910390fd5b505050815f015f9054906101000a900460ff1660ff168110158190611ad6576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611acd9190613b9f565b60405180910390fd5b505f5f90505f5f90505f5f90505b83811015611bd8575f611b5a898381518110611b0357611b026146d9565b5b6020026020010151898481518110611b1e57611b1d6146d9565b5b6020026020010151898581518110611b3957611b386146d9565b5b6020026020010151611b4a8e612833565b61284c909392919063ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16108015611bbd5750611bbc7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926826111f1565b5b15611bcc578360010193508092505b50806001019050611ae4565b50835f015f9054906101000a900460ff1660ff168210158290611c31576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611c289190613b9f565b60405180910390fd5b505050505050505050565b67ffffffffffffffff8016841180611c5d575067ffffffffffffffff801682115b15611c94576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8681526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115611d0957836040517fb7aa1d97000000000000000000000000000000000000000000000000000000008152600401611d009190613b9f565b60405180910390fd5b5f6040518060c001604052808667ffffffffffffffff1681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681526020018381526020018773ffffffffffffffffffffffffffffffffffffffff168152509050805f5f8781526020019081526020015f205f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f0160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151815f0160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001015560a0820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505050505050565b5f60015490505b600454600154600254611eb39190614608565b1115611ed857611ec460015461287a565b60015f815460010191905081905550611ea0565b806001541115611f2857807fb8a4623367e916c9ab30fca57e7f767aaa221385f9e2726b7cd24bffde7385e282600154611f129190614608565b604051611f1f9190613b9f565b60405180910390a25b50565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b5f611f5d8383612998565b90508015611fdd575f611f6e612699565b9050611f9483825f015f8781526020019081526020015f20612a9090919063ffffffff16565b83859091611fd9576040517fd180cb31000000000000000000000000000000000000000000000000000000008152600401611fd0929190614833565b60405180910390fd5b5050505b92915050565b5f611fee8383612abd565b9050801561206e575f611fff612699565b905061202583825f015f8781526020019081526020015f20612bb590919063ffffffff16565b8385909161206a576040517f054af8d8000000000000000000000000000000000000000000000000000000008152600401612061929190614833565b60405180910390fd5b5050505b92915050565b5f7f303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061214857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661212f612be2565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561217f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f1b61218d816119b5565b5050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156121f957506040513d601f19601f820116820180604052508101906121f6919061486e565b60015b61223a57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016122319190614899565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b81146122a057806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016122979190613bf2565b60405180910390fd5b6122aa8383612c35565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612334576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b612365612ca7565b565b61236f612ca7565b612377612ce7565b565b612381612ca7565b61238a81612d17565b50565b612395612ca7565b5f8160ff16036123d1576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124456040518060400160405280600981526020017f56616c696461746f7200000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e30000000000000000000000000000000000000000000000000000000815250612d37565b5f61244e612074565b905081815f015f6101000a81548160ff021916908360ff1602179055505050565b60648110156124b557806040517f94c570b90000000000000000000000000000000000000000000000000000000081526004016124ac9190613b9f565b60405180910390fd5b8060038190555050565b600a81101561250557806040517f7bdbff960000000000000000000000000000000000000000000000000000000081526004016124fc9190613b9f565b60405180910390fd5b8060048190555050565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100905090565b60605f612568612536565b9050806002018054612579906148df565b80601f01602080910402602001604051908101604052809291908181526020018280546125a5906148df565b80156125f05780601f106125c7576101008083540402835291602001916125f0565b820191905f5260205f20905b8154815290600101906020018083116125d357829003601f168201915b505050505091505090565b60605f612606612536565b9050806003018054612617906148df565b80601f0160208091040260200160405190810160405280929190818152602001828054612643906148df565b801561268e5780601f106126655761010080835404028352916020019161268e565b820191905f5260205f20905b81548152906001019060200180831161267157829003601f168201915b505050505091505090565b5f7f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00905090565b60605f6126ce835f01612d4d565b905060608190508092505050919050565b5f826126ec868685612da6565b149050949350505050565b6126ff6119c9565b5f61270861250f565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861274e611a0a565b60405161275b9190614899565b60405180910390a150565b61276e612df9565b5f61277761250f565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6127bc611a0a565b6040516127c99190614899565b60405180910390a150565b5f6127dd612e39565b905090565b6127ec82826111f1565b61282f5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612826929190614833565b60405180910390fd5b5050565b5f61284561283f6127d4565b83612e9c565b9050919050565b5f5f5f5f61285c88888888612edc565b92509250925061286c8282612fc3565b829350505050949350505050565b5f5f5f8381526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff16036128ee57806040517f21ee5f4b0000000000000000000000000000000000000000000000000000000081526004016128e59190613b9f565b60405180910390fd5b5f5f8281526020019081526020015f205f5f82015f6101000a81549067ffffffffffffffff02191690555f820160086101000a81549067ffffffffffffffff02191690555f820160106101000a81549067ffffffffffffffff02191690555f820160186101000a81549067ffffffffffffffff0219169055600182015f9055600282015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505050565b5f5f6129a2611f2b565b90506129ae84846111f1565b612a85576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612a21611a0a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612a8a565b5f9150505b92915050565b5f612ab5835f018373ffffffffffffffffffffffffffffffffffffffff165f1b613125565b905092915050565b5f5f612ac7611f2b565b9050612ad384846111f1565b15612baa575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612b46611a0a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050612baf565b5f9150505b92915050565b5f612bda835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61318c565b905092915050565b5f612c0e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613288565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b612c3e82613291565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115612c9a57612c94828261335a565b50612ca3565b612ca26133da565b5b5050565b612caf613416565b612ce5576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612cef612ca7565b5f612cf861250f565b90505f815f015f6101000a81548160ff02191690831515021790555050565b612d1f612ca7565b612d27613434565b612d335f5f1b82611f52565b5050565b612d3f612ca7565b612d49828261343e565b5050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612d9a57602002820191905f5260205f20905b815481526020019060010190808311612d86575b50505050509050919050565b5f5f8290505f5f90505b85859050811015612ded57612dde82878784818110612dd257612dd16146d9565b5b9050602002013561348f565b91508080600101915050612db0565b50809150509392505050565b612e01610e92565b612e37576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612e636134b9565b612e6b61352f565b4630604051602001612e8195949392919061490f565b60405160208183030381529060405280519060200120905090565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c1115612f18575f600385925092509250612fb9565b5f6001888888886040515f8152602001604052604051612f3b9493929190614960565b6020604051602081039080840390855afa158015612f5b573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612fac575f60015f5f1b93509350935050612fb9565b805f5f5f1b935093509350505b9450945094915050565b5f6003811115612fd657612fd56149a3565b5b826003811115612fe957612fe86149a3565b5b03156131215760016003811115613003576130026149a3565b5b826003811115613016576130156149a3565b5b0361304d576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026003811115613061576130606149a3565b5b826003811115613074576130736149a3565b5b036130b857805f1c6040517ffce698f70000000000000000000000000000000000000000000000000000000081526004016130af9190613b9f565b60405180910390fd5b6003808111156130cb576130ca6149a3565b5b8260038111156130de576130dd6149a3565b5b0361312057806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016131179190613bf2565b60405180910390fd5b5b5050565b5f61313083836135a6565b61318257825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050613186565b5f90505b92915050565b5f5f836001015f8481526020019081526020015f205490505f811461327d575f6001826131b99190614608565b90505f6001865f01805490506131cf9190614608565b9050808214613235575f865f0182815481106131ee576131ed6146d9565b5b905f5260205f200154905080875f01848154811061320f5761320e6146d9565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f01805480613248576132476149d0565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050613282565b5f9150505b92915050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b036132ec57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016132e39190614899565b60405180910390fd5b806133187f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613288565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516133839190614a41565b5f60405180830381855af49150503d805f81146133bb576040519150601f19603f3d011682016040523d82523d5f602084013e6133c0565b606091505b50915091506133d08583836135c6565b9250505092915050565b5f341115613414576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61341f612336565b5f0160089054906101000a900460ff16905090565b61343c612ca7565b565b613446612ca7565b5f61344f612536565b9050828160020190816134629190614bee565b50818160030190816134749190614bee565b505f5f1b815f01819055505f5f1b8160010181905550505050565b5f8183106134a6576134a18284613653565b6134b1565b6134b08383613653565b5b905092915050565b5f5f6134c3612536565b90505f6134ce61255d565b90505f815111156134ea5780805190602001209250505061352c565b5f825f015490505f5f1b81146135055780935050505061352c565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f613539612536565b90505f6135446125fb565b90505f81511115613560578080519060200120925050506135a3565b5f826001015490505f5f1b811461357c578093505050506135a3565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f836001015f8481526020019081526020015f20541415905092915050565b6060826135db576135d682613667565b61364b565b5f825114801561360157505f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561364357836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161363a9190614899565b60405180910390fd5b81905061364c565b5b9392505050565b5f825f528160205260405f20905092915050565b5f815111156136795780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61375e8161372a565b8114613768575f5ffd5b50565b5f8135905061377981613755565b92915050565b5f6020828403121561379457613793613722565b5b5f6137a18482850161376b565b91505092915050565b5f8115159050919050565b6137be816137aa565b82525050565b5f6020820190506137d75f8301846137b5565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126137fe576137fd6137dd565b5b8235905067ffffffffffffffff81111561381b5761381a6137e1565b5b602083019150836001820283011115613837576138366137e5565b5b9250929050565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6138848261383e565b810181811067ffffffffffffffff821117156138a3576138a261384e565b5b80604052505050565b5f6138b5613719565b90506138c1828261387b565b919050565b5f67ffffffffffffffff8211156138e0576138df61384e565b5b602082029050602081019050919050565b5f60ff82169050919050565b613906816138f1565b8114613910575f5ffd5b50565b5f81359050613921816138fd565b92915050565b5f613939613934846138c6565b6138ac565b9050808382526020820190506020840283018581111561395c5761395b6137e5565b5b835b8181101561398557806139718882613913565b84526020840193505060208101905061395e565b5050509392505050565b5f82601f8301126139a3576139a26137dd565b5b81356139b3848260208601613927565b91505092915050565b5f67ffffffffffffffff8211156139d6576139d561384e565b5b602082029050602081019050919050565b5f819050919050565b6139f9816139e7565b8114613a03575f5ffd5b50565b5f81359050613a14816139f0565b92915050565b5f613a2c613a27846139bc565b6138ac565b90508083825260208201905060208402830185811115613a4f57613a4e6137e5565b5b835b81811015613a785780613a648882613a06565b845260208401935050602081019050613a51565b5050509392505050565b5f82601f830112613a9657613a956137dd565b5b8135613aa6848260208601613a1a565b91505092915050565b5f5f5f5f5f60808688031215613ac857613ac7613722565b5b5f86013567ffffffffffffffff811115613ae557613ae4613726565b5b613af1888289016137e9565b9550955050602086013567ffffffffffffffff811115613b1457613b13613726565b5b613b208882890161398f565b935050604086013567ffffffffffffffff811115613b4157613b40613726565b5b613b4d88828901613a82565b925050606086013567ffffffffffffffff811115613b6e57613b6d613726565b5b613b7a88828901613a82565b9150509295509295909350565b5f819050919050565b613b9981613b87565b82525050565b5f602082019050613bb25f830184613b90565b92915050565b5f60208284031215613bcd57613bcc613722565b5b5f613bda84828501613a06565b91505092915050565b613bec816139e7565b82525050565b5f602082019050613c055f830184613be3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613c3482613c0b565b9050919050565b613c4481613c2a565b8114613c4e575f5ffd5b50565b5f81359050613c5f81613c3b565b92915050565b5f5f60408385031215613c7b57613c7a613722565b5b5f613c8885828601613a06565b9250506020613c9985828601613c51565b9150509250929050565b613cac816138f1565b82525050565b5f602082019050613cc55f830184613ca3565b92915050565b5f604082019050613cde5f830185613b90565b613ceb6020830184613b90565b9392505050565b5f5ffd5b5f67ffffffffffffffff821115613d1057613d0f61384e565b5b613d198261383e565b9050602081019050919050565b828183375f83830152505050565b5f613d46613d4184613cf6565b6138ac565b905082815260208101848484011115613d6257613d61613cf2565b5b613d6d848285613d26565b509392505050565b5f82601f830112613d8957613d886137dd565b5b8135613d99848260208601613d34565b91505092915050565b5f5f60408385031215613db857613db7613722565b5b5f613dc585828601613c51565b925050602083013567ffffffffffffffff811115613de657613de5613726565b5b613df285828601613d75565b9150509250929050565b613e0581613b87565b8114613e0f575f5ffd5b50565b5f81359050613e2081613dfc565b92915050565b5f5f5f60608486031215613e3d57613e3c613722565b5b5f613e4a86828701613e12565b9350506020613e5b86828701613e12565b9250506040613e6c86828701613913565b9150509250925092565b5f60208284031215613e8b57613e8a613722565b5b5f613e9884828501613e12565b91505092915050565b5f67ffffffffffffffff821115613ebb57613eba61384e565b5b602082029050602081019050919050565b5f613ede613ed984613ea1565b6138ac565b90508083825260208201905060208402830185811115613f0157613f006137e5565b5b835b81811015613f2a5780613f168882613c51565b845260208401935050602081019050613f03565b5050509392505050565b5f82601f830112613f4857613f476137dd565b5b8135613f58848260208601613ecc565b91505092915050565b5f5f60408385031215613f7757613f76613722565b5b5f613f8485828601613a06565b925050602083013567ffffffffffffffff811115613fa557613fa4613726565b5b613fb185828601613f34565b9150509250929050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b613fef81613fbb565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61402782613ff5565b6140318185613fff565b935061404181856020860161400f565b61404a8161383e565b840191505092915050565b61405e81613c2a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61409681613b87565b82525050565b5f6140a7838361408d565b60208301905092915050565b5f602082019050919050565b5f6140c982614064565b6140d3818561406e565b93506140de8361407e565b805f5b8381101561410e5781516140f5888261409c565b9750614100836140b3565b9250506001810190506140e1565b5085935050505092915050565b5f60e08201905061412e5f83018a613fe6565b8181036020830152614140818961401d565b90508181036040830152614154818861401d565b90506141636060830187613b90565b6141706080830186614055565b61417d60a0830185613be3565b81810360c083015261418f81846140bf565b905098975050505050505050565b5f67ffffffffffffffff82169050919050565b6141b98161419d565b82525050565b6141c8816139e7565b82525050565b6141d781613c2a565b82525050565b60c082015f8201516141f15f8501826141b0565b50602082015161420460208501826141b0565b50604082015161421760408501826141b0565b50606082015161422a60608501826141b0565b50608082015161423d60808501826141bf565b5060a082015161425060a08501826141ce565b50505050565b5f60c0820190506142695f8301846141dd565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6142a383836141ce565b60208301905092915050565b5f602082019050919050565b5f6142c58261426f565b6142cf8185614279565b93506142da83614289565b805f5b8381101561430a5781516142f18882614298565b97506142fc836142af565b9250506001810190506142dd565b5085935050505092915050565b5f6020820190508181035f83015261432f81846142bb565b905092915050565b5f6020820190508181035f83015261434f818461401d565b905092915050565b5f5f83601f84011261436c5761436b6137dd565b5b8235905067ffffffffffffffff811115614389576143886137e1565b5b6020830191508360208202830111156143a5576143a46137e5565b5b9250929050565b5f5f5f5f606085870312156143c4576143c3613722565b5b5f6143d187828801613e12565b945050602085013567ffffffffffffffff8111156143f2576143f1613726565b5b6143fe87828801614357565b9350935050604061441187828801613a06565b91505092959194509250565b5f6020828403121561443257614431613722565b5b5f61443f84828501613913565b91505092915050565b614451816137aa565b811461445b575f5ffd5b50565b5f8135905061446c81614448565b92915050565b5f6020828403121561448757614486613722565b5b5f6144948482850161445e565b91505092915050565b5f5ffd5b5f60a082840312156144b6576144b561449d565b5b6144c060a06138ac565b90505f6144cf84828501613e12565b5f8301525060206144e284828501613e12565b60208301525060406144f684828501613e12565b604083015250606061450a84828501613a06565b606083015250608061451e84828501613e12565b60808301525092915050565b5f60a0828403121561453f5761453e613722565b5b5f61454c848285016144a1565b91505092915050565b5f60c0820190506145685f830189613be3565b6145756020830188613b90565b6145826040830187613b90565b61458f6060830186613b90565b61459c6080830185613be3565b6145a960a0830184613b90565b979650505050505050565b5f6040820190506145c75f830185613b90565b6145d46020830184613be3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61461282613b87565b915061461d83613b87565b9250828203905081811115614635576146346145db565b5b92915050565b5f60808201905061464e5f830187613b90565b61465b6020830186613b90565b6146686040830185613b90565b6146756060830184613ca3565b95945050505050565b5f819050919050565b5f819050919050565b5f6146aa6146a56146a08461467e565b614687565b61419d565b9050919050565b6146ba81614690565b82525050565b5f6020820190506146d35f8301846146b1565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4549503731323a20556e696e697469616c697a656400000000000000000000005f82015250565b5f61473a601583613fff565b915061474582614706565b602082019050919050565b5f6020820190508181035f8301526147678161472e565b9050919050565b5f61477882613b87565b915061478383613b87565b925082820190508082111561479b5761479a6145db565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6147d882613b87565b91506147e383613b87565b9250826147f3576147f26147a1565b5b828204905092915050565b5f6060820190506148115f830186613b90565b61481e6020830185613b90565b61482b6040830184613b90565b949350505050565b5f6040820190506148465f830185614055565b6148536020830184613be3565b9392505050565b5f81519050614868816139f0565b92915050565b5f6020828403121561488357614882613722565b5b5f6148908482850161485a565b91505092915050565b5f6020820190506148ac5f830184614055565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806148f657607f821691505b602082108103614909576149086148b2565b5b50919050565b5f60a0820190506149225f830188613be3565b61492f6020830187613be3565b61493c6040830186613be3565b6149496060830185613b90565b6149566080830184614055565b9695505050505050565b5f6080820190506149735f830187613be3565b6149806020830186613ca3565b61498d6040830185613be3565b61499a6060830184613be3565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f81519050919050565b5f81905092915050565b5f614a1b826149fd565b614a258185614a07565b9350614a3581856020860161400f565b80840191505092915050565b5f614a4c8284614a11565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302614ab37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614a78565b614abd8683614a78565b95508019841693508086168417925050509392505050565b5f614aef614aea614ae584613b87565b614687565b613b87565b9050919050565b5f819050919050565b614b0883614ad5565b614b1c614b1482614af6565b848454614a84565b825550505050565b5f5f905090565b614b33614b24565b614b3e818484614aff565b505050565b5b81811015614b6157614b565f82614b2b565b600181019050614b44565b5050565b601f821115614ba657614b7781614a57565b614b8084614a69565b81016020851015614b8f578190505b614ba3614b9b85614a69565b830182614b43565b50505b505050565b5f82821c905092915050565b5f614bc65f1984600802614bab565b1980831691505092915050565b5f614bde8383614bb7565b9150826002028217905092915050565b614bf782613ff5565b67ffffffffffffffff811115614c1057614c0f61384e565b5b614c1a82546148df565b614c25828285614b65565b5f60209050601f831160018114614c56575f8415614c44578287015190505b614c4e8582614bd3565b865550614cb5565b601f198416614c6486614a57565b5f5b82811015614c8b57848901518255600182019150602085019450602081019050614c66565b86831015614ca85784890151614ca4601f891682614bb7565b8355505b6001600288020188555050505b50505050505056fea264697066735822122027178bf391b91fef9410f7e085a28ef520843db4b0f3352cd9338896898261a064736f6c634300081c0033",
}

// CrossCheckABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossCheckMetaData.ABI instead.
var CrossCheckABI = CrossCheckMetaData.ABI

// CrossCheckBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossCheckBinRuntime = "6080604052600436106101ed575f3560e01c806384b0196e1161010c578063b7f3358d1161009f578063d547741f1161006e578063d547741f146106ee578063f5b541a614610716578063f698da2514610740578063fcc6aca91461076a578063fd47852b14610794576101ed565b8063b7f3358d1461064c578063bedb86fb14610674578063c00111141461069c578063c49baebe146106c4576101ed565b8063a217fddf116100db578063a217fddf14610580578063a3246ad3146105aa578063ad3cb1cc146105e6578063aea3679f14610610576101ed565b806384b0196e1461049c578063871a782e146104cc57806391d14854146105085780639f3e14cf14610544576101ed565b80634f1ef286116101845780635c975abb116101535780635c975abb146103f857806366cdc83b1461042257806375b238fc1461044a57806382f51fa614610474576101ed565b80634f1ef2861461036057806352d1902d1461037c578063544d7e1e146103a657806355a84bb4146103d0576101ed565b80632f2ff15d116101c05780632f2ff15d146102bb57806336568abe146102e357806342cde4e81461030b5780634d58190714610335576101ed565b806301ffc9a7146101f157806307a9eef21461022d57806317894ff414610255578063248a9ca31461027f575b5f5ffd5b3480156101fc575f5ffd5b506102176004803603810190610212919061377f565b6107bc565b60405161022491906137c4565b60405180910390f35b348015610238575f5ffd5b50610253600480360381019061024e9190613aaf565b610835565b005b348015610260575f5ffd5b50610269610ac6565b6040516102769190613b9f565b60405180910390f35b34801561028a575f5ffd5b506102a560048036038101906102a09190613bb8565b610adc565b6040516102b29190613bf2565b60405180910390f35b3480156102c6575f5ffd5b506102e160048036038101906102dc9190613c65565b610b06565b005b3480156102ee575f5ffd5b5061030960048036038101906103049190613c65565b610b28565b005b348015610316575f5ffd5b5061031f610ba3565b60405161032c9190613cb2565b60405180910390f35b348015610340575f5ffd5b50610349610bc5565b604051610357929190613ccb565b60405180910390f35b61037a60048036038101906103759190613da2565b610c4e565b005b348015610387575f5ffd5b50610390610c6d565b60405161039d9190613bf2565b60405180910390f35b3480156103b1575f5ffd5b506103ba610c9e565b6040516103c79190613b9f565b60405180910390f35b3480156103db575f5ffd5b506103f660048036038101906103f19190613e26565b610ca4565b005b348015610403575f5ffd5b5061040c610e92565b60405161041991906137c4565b60405180910390f35b34801561042d575f5ffd5b5061044860048036038101906104439190613e76565b610eb4565b005b348015610455575f5ffd5b5061045e610f22565b60405161046b9190613bf2565b60405180910390f35b34801561047f575f5ffd5b5061049a60048036038101906104959190613f61565b610f46565b005b3480156104a7575f5ffd5b506104b0610f9c565b6040516104c3979695949392919061411b565b60405180910390f35b3480156104d7575f5ffd5b506104f260048036038101906104ed9190613e76565b6110a5565b6040516104ff9190614256565b60405180910390f35b348015610513575f5ffd5b5061052e60048036038101906105299190613c65565b6111f1565b60405161053b91906137c4565b60405180910390f35b34801561054f575f5ffd5b5061056a60048036038101906105659190613e76565b611262565b6040516105779190614256565b60405180910390f35b34801561058b575f5ffd5b5061059461162e565b6040516105a19190613bf2565b60405180910390f35b3480156105b5575f5ffd5b506105d060048036038101906105cb9190613bb8565b611634565b6040516105dd9190614317565b60405180910390f35b3480156105f1575f5ffd5b506105fa611663565b6040516106079190614337565b60405180910390f35b34801561061b575f5ffd5b50610636600480360381019061063191906143ac565b61169c565b60405161064391906137c4565b60405180910390f35b348015610657575f5ffd5b50610672600480360381019061066d919061441d565b611716565b005b34801561067f575f5ffd5b5061069a60048036038101906106959190614472565b6117c0565b005b3480156106a7575f5ffd5b506106c260048036038101906106bd9190613e76565b61180a565b005b3480156106cf575f5ffd5b506106d8611878565b6040516106e59190613bf2565b60405180910390f35b3480156106f9575f5ffd5b50610714600480360381019061070f9190613c65565b61189c565b005b348015610721575f5ffd5b5061072a6118be565b6040516107379190613bf2565b60405180910390f35b34801561074b575f5ffd5b506107546118e2565b6040516107619190613bf2565b60405180910390f35b348015610775575f5ffd5b5061077e6118f0565b60405161078b9190613b9f565b60405180910390f35b34801561079f575f5ffd5b506107ba60048036038101906107b59190613f61565b6118f6565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061082e575061082d8261194c565b5b9050919050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892661085f816119b5565b6108676119c9565b5f610870611a0a565b90505f8787810190610882919061452a565b9050620956cc8160800151146108d35780608001516040517fa5acee440000000000000000000000000000000000000000000000000000000081526004016108ca9190613b9f565b60405180910390fd5b806020015181604001511115806108ef57505f5f1b8160600151145b15610926576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61092f610bc5565b9150505f5f905060016003548301039050825f0151600254141580610958575082602001518214155b80610967575082604001518114155b156109b257825f015183602001516040517f99a5875b0000000000000000000000000000000000000000000000000000000081526004016109a9929190613ccb565b60405180910390fd5b5f7ffe27049a99724ab477714cb3e8105a7b7d274fe7e5bee5cf905d63758b108490845f01518560200151866040015187606001518860800151604051602001610a0196959493929190614555565b604051602081830303815290604052805190602001209050610a25818a8a8a611a11565b610a4185855f0151866020015187604001518860600151611c3c565b60025f8154600101919050819055508360200151845f01518673ffffffffffffffffffffffffffffffffffffffff167fbadf6f34b60f61fb703b4f3ffa385fcf9f98bf0a53cfbadd15fe381c30b145bf87604001518860600151604051610aa99291906145b4565b60405180910390a4610ab9611e99565b5050505050505050505050565b5f600154600254610ad79190614608565b905090565b5f5f610ae6611f2b565b9050805f015f8481526020019081526020015f2060010154915050919050565b610b0f82610adc565b610b18816119b5565b610b228383611f52565b50505050565b610b30611a0a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b94576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b9e8282611fe3565b505050565b5f5f610bad612074565b9050805f015f9054906101000a900460ff1691505090565b5f5f60025491505f6002541115610c4a575f5f5f6001600254610be89190614608565b81526020019081526020015f2090505f815f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115610c48576001815f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff160191505b505b9091565b610c5661209b565b610c5f82612181565b610c698282612191565b5050565b5f610c766122af565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b60045481565b5f610cad612336565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff16148015610cf55750825b90505f60018367ffffffffffffffff16148015610d2857505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610d36575080155b15610d6d576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610dba576001855f0160086101000a81548160ff0219169083151502179055505b610dc261235d565b610dca612367565b610dd333612379565b610ddc8661238d565b610de58861246f565b610dee876124bf565b7f38778a692f0977d8b900e132518cdf1350918c14b717a3f68825d1e067476bc1620956cc898989604051610e26949392919061463b565b60405180910390a18315610e88575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610e7f91906146c0565b60405180910390a15b5050505050505050565b5f5f610e9c61250f565b9050805f015f9054906101000a900460ff1691505090565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610ede816119b5565b610ee7826124bf565b7ff3114416d687e19a2ba81f92f6b8a93a48863d39989e5fd2f1adb0f5f89fd69982604051610f169190613b9f565b60405180910390a15050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b610f4f82610adc565b610f58816119b5565b5f5f90505b8251811015610f9657610f8a84848381518110610f7d57610f7c6146d9565b5b6020026020010151611fe3565b50806001019050610f5d565b50505050565b5f6060805f5f5f60605f610fae612536565b90505f5f1b815f0154148015610fc957505f5f1b8160010154145b611008576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fff90614750565b60405180910390fd5b61101061255d565b6110186125fb565b46305f5f1b5f67ffffffffffffffff8111156110375761103661384e565b5b6040519080825280602002602001820160405280156110655781602001602082028036833780820191505090505b507f0f0000000000000000000000000000000000000000000000000000000000000095949392919097509750975097509750975097505090919293949596565b6110ad6136ab565b5f5f8381526020019081526020015f206040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b5f5f6111fb611f2b565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b61126a6136ab565b600254600154106112e9576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152509050611629565b5f5f60015481526020019081526020015f205f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16821015611397576040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff168152509050611629565b5f60015490505f60016002546113ad9190614608565b90505b8082116115b7575f600282846113c6919061476e565b6113d091906147ce565b90505f5f5f8381526020019081526020015f20905085815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16111580156114375750805f0160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168611155b1561157557806040518060c00160405290815f82015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020015f820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050945050505050611629565b805f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168610156115a9576001820392506115b0565b6001820193505b50506113b0565b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f5f1b81526020015f73ffffffffffffffffffffffffffffffffffffffff16815250925050505b919050565b5f5f1b81565b60605f61163f612699565b905061165b815f015f8581526020019081526020015f206126c0565b915050919050565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f5f6116a786611262565b90505f816060015167ffffffffffffffff16036116fb57856040517f96fab5950000000000000000000000000000000000000000000000000000000081526004016116f29190613b9f565b60405180910390fd5b61170b85858360800151866126df565b915050949350505050565b5f5f1b611722816119b5565b5f8260ff160361175e576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611767612074565b905082815f015f6101000a81548160ff021916908360ff1602179055507f541ae612d90d592d7865e446bfd28eb8c1de1ced93c6f992a6df36c9e7bd33ff836040516117b39190613cb2565b60405180910390a1505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9296117ea816119b5565b81156117fd576117f86126f7565b611806565b611805612766565b5b5050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611834816119b5565b61183d8261246f565b7fc37172b02acb2a6b00612c1178f47d9f660e5fa7800fa7a3904e7077665659128260405161186c9190613b9f565b60405180910390a15050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6118a582610adc565b6118ae816119b5565b6118b88383611fe3565b50505050565b7f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b5f6118eb6127d4565b905090565b60035481565b6118ff82610adc565b611908816119b5565b5f5f90505b82518110156119465761193a8484838151811061192d5761192c6146d9565b5b6020026020010151611f52565b5080600101905061190d565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6119c6816119c1611a0a565b6127e2565b50565b6119d1610e92565b15611a08576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f33905090565b5f611a1a612074565b90505f84519050835181148015611a315750825181145b8185518551909192611a7b576040517f9089eb84000000000000000000000000000000000000000000000000000000008152600401611a72939291906147fe565b60405180910390fd5b505050815f015f9054906101000a900460ff1660ff168110158190611ad6576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611acd9190613b9f565b60405180910390fd5b505f5f90505f5f90505f5f90505b83811015611bd8575f611b5a898381518110611b0357611b026146d9565b5b6020026020010151898481518110611b1e57611b1d6146d9565b5b6020026020010151898581518110611b3957611b386146d9565b5b6020026020010151611b4a8e612833565b61284c909392919063ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16108015611bbd5750611bbc7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926826111f1565b5b15611bcc578360010193508092505b50806001019050611ae4565b50835f015f9054906101000a900460ff1660ff168210158290611c31576040517f35d7a2f2000000000000000000000000000000000000000000000000000000008152600401611c289190613b9f565b60405180910390fd5b505050505050505050565b67ffffffffffffffff8016841180611c5d575067ffffffffffffffff801682115b15611c94576040517f688e967f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8681526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff161115611d0957836040517fb7aa1d97000000000000000000000000000000000000000000000000000000008152600401611d009190613b9f565b60405180910390fd5b5f6040518060c001604052808667ffffffffffffffff1681526020018567ffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681526020018381526020018773ffffffffffffffffffffffffffffffffffffffff168152509050805f5f8781526020019081526020015f205f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f0160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151815f0160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001015560a0820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505050505050565b5f60015490505b600454600154600254611eb39190614608565b1115611ed857611ec460015461287a565b60015f815460010191905081905550611ea0565b806001541115611f2857807fb8a4623367e916c9ab30fca57e7f767aaa221385f9e2726b7cd24bffde7385e282600154611f129190614608565b604051611f1f9190613b9f565b60405180910390a25b50565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b5f611f5d8383612998565b90508015611fdd575f611f6e612699565b9050611f9483825f015f8781526020019081526020015f20612a9090919063ffffffff16565b83859091611fd9576040517fd180cb31000000000000000000000000000000000000000000000000000000008152600401611fd0929190614833565b60405180910390fd5b5050505b92915050565b5f611fee8383612abd565b9050801561206e575f611fff612699565b905061202583825f015f8781526020019081526020015f20612bb590919063ffffffff16565b8385909161206a576040517f054af8d8000000000000000000000000000000000000000000000000000000008152600401612061929190614833565b60405180910390fd5b5050505b92915050565b5f7f303ad04a409295f82fe653ea0c2830faa7f9686cca19dded40cb6b9dfd913200905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061214857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661212f612be2565b73ffffffffffffffffffffffffffffffffffffffff1614155b1561217f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f5f1b61218d816119b5565b5050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156121f957506040513d601f19601f820116820180604052508101906121f6919061486e565b60015b61223a57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016122319190614899565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b81146122a057806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016122979190613bf2565b60405180910390fd5b6122aa8383612c35565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614612334576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b612365612ca7565b565b61236f612ca7565b612377612ce7565b565b612381612ca7565b61238a81612d17565b50565b612395612ca7565b5f8160ff16036123d1576040517ff0f15b9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124456040518060400160405280600981526020017f56616c696461746f7200000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e30000000000000000000000000000000000000000000000000000000815250612d37565b5f61244e612074565b905081815f015f6101000a81548160ff021916908360ff1602179055505050565b60648110156124b557806040517f94c570b90000000000000000000000000000000000000000000000000000000081526004016124ac9190613b9f565b60405180910390fd5b8060038190555050565b600a81101561250557806040517f7bdbff960000000000000000000000000000000000000000000000000000000081526004016124fc9190613b9f565b60405180910390fd5b8060048190555050565b5f7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b5f7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100905090565b60605f612568612536565b9050806002018054612579906148df565b80601f01602080910402602001604051908101604052809291908181526020018280546125a5906148df565b80156125f05780601f106125c7576101008083540402835291602001916125f0565b820191905f5260205f20905b8154815290600101906020018083116125d357829003601f168201915b505050505091505090565b60605f612606612536565b9050806003018054612617906148df565b80601f0160208091040260200160405190810160405280929190818152602001828054612643906148df565b801561268e5780601f106126655761010080835404028352916020019161268e565b820191905f5260205f20905b81548152906001019060200180831161267157829003601f168201915b505050505091505090565b5f7f66b7cbf19cb8e36075b986e08abfd5bd2137fd1e8b0682ffea32f23305c42f00905090565b60605f6126ce835f01612d4d565b905060608190508092505050919050565b5f826126ec868685612da6565b149050949350505050565b6126ff6119c9565b5f61270861250f565b90506001815f015f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861274e611a0a565b60405161275b9190614899565b60405180910390a150565b61276e612df9565b5f61277761250f565b90505f815f015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6127bc611a0a565b6040516127c99190614899565b60405180910390a150565b5f6127dd612e39565b905090565b6127ec82826111f1565b61282f5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612826929190614833565b60405180910390fd5b5050565b5f61284561283f6127d4565b83612e9c565b9050919050565b5f5f5f5f61285c88888888612edc565b92509250925061286c8282612fc3565b829350505050949350505050565b5f5f5f8381526020019081526020015f205f0160189054906101000a900467ffffffffffffffff1667ffffffffffffffff16036128ee57806040517f21ee5f4b0000000000000000000000000000000000000000000000000000000081526004016128e59190613b9f565b60405180910390fd5b5f5f8281526020019081526020015f205f5f82015f6101000a81549067ffffffffffffffff02191690555f820160086101000a81549067ffffffffffffffff02191690555f820160106101000a81549067ffffffffffffffff02191690555f820160186101000a81549067ffffffffffffffff0219169055600182015f9055600282015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505050565b5f5f6129a2611f2b565b90506129ae84846111f1565b612a85576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612a21611a0a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612a8a565b5f9150505b92915050565b5f612ab5835f018373ffffffffffffffffffffffffffffffffffffffff165f1b613125565b905092915050565b5f5f612ac7611f2b565b9050612ad384846111f1565b15612baa575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550612b46611a0a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050612baf565b5f9150505b92915050565b5f612bda835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61318c565b905092915050565b5f612c0e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613288565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b612c3e82613291565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115612c9a57612c94828261335a565b50612ca3565b612ca26133da565b5b5050565b612caf613416565b612ce5576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b612cef612ca7565b5f612cf861250f565b90505f815f015f6101000a81548160ff02191690831515021790555050565b612d1f612ca7565b612d27613434565b612d335f5f1b82611f52565b5050565b612d3f612ca7565b612d49828261343e565b5050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612d9a57602002820191905f5260205f20905b815481526020019060010190808311612d86575b50505050509050919050565b5f5f8290505f5f90505b85859050811015612ded57612dde82878784818110612dd257612dd16146d9565b5b9050602002013561348f565b91508080600101915050612db0565b50809150509392505050565b612e01610e92565b612e37576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612e636134b9565b612e6b61352f565b4630604051602001612e8195949392919061490f565b60405160208183030381529060405280519060200120905090565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c1115612f18575f600385925092509250612fb9565b5f6001888888886040515f8152602001604052604051612f3b9493929190614960565b6020604051602081039080840390855afa158015612f5b573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612fac575f60015f5f1b93509350935050612fb9565b805f5f5f1b935093509350505b9450945094915050565b5f6003811115612fd657612fd56149a3565b5b826003811115612fe957612fe86149a3565b5b03156131215760016003811115613003576130026149a3565b5b826003811115613016576130156149a3565b5b0361304d576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026003811115613061576130606149a3565b5b826003811115613074576130736149a3565b5b036130b857805f1c6040517ffce698f70000000000000000000000000000000000000000000000000000000081526004016130af9190613b9f565b60405180910390fd5b6003808111156130cb576130ca6149a3565b5b8260038111156130de576130dd6149a3565b5b0361312057806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016131179190613bf2565b60405180910390fd5b5b5050565b5f61313083836135a6565b61318257825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050613186565b5f90505b92915050565b5f5f836001015f8481526020019081526020015f205490505f811461327d575f6001826131b99190614608565b90505f6001865f01805490506131cf9190614608565b9050808214613235575f865f0182815481106131ee576131ed6146d9565b5b905f5260205f200154905080875f01848154811061320f5761320e6146d9565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f01805480613248576132476149d0565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050613282565b5f9150505b92915050565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b036132ec57806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016132e39190614899565b60405180910390fd5b806133187f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b613288565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516133839190614a41565b5f60405180830381855af49150503d805f81146133bb576040519150601f19603f3d011682016040523d82523d5f602084013e6133c0565b606091505b50915091506133d08583836135c6565b9250505092915050565b5f341115613414576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61341f612336565b5f0160089054906101000a900460ff16905090565b61343c612ca7565b565b613446612ca7565b5f61344f612536565b9050828160020190816134629190614bee565b50818160030190816134749190614bee565b505f5f1b815f01819055505f5f1b8160010181905550505050565b5f8183106134a6576134a18284613653565b6134b1565b6134b08383613653565b5b905092915050565b5f5f6134c3612536565b90505f6134ce61255d565b90505f815111156134ea5780805190602001209250505061352c565b5f825f015490505f5f1b81146135055780935050505061352c565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f613539612536565b90505f6135446125fb565b90505f81511115613560578080519060200120925050506135a3565b5f826001015490505f5f1b811461357c578093505050506135a3565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47093505050505b90565b5f5f836001015f8481526020019081526020015f20541415905092915050565b6060826135db576135d682613667565b61364b565b5f825114801561360157505f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561364357836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161363a9190614899565b60405180910390fd5b81905061364c565b5b9392505050565b5f825f528160205260405f20905092915050565b5f815111156136795780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060c001604052805f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61375e8161372a565b8114613768575f5ffd5b50565b5f8135905061377981613755565b92915050565b5f6020828403121561379457613793613722565b5b5f6137a18482850161376b565b91505092915050565b5f8115159050919050565b6137be816137aa565b82525050565b5f6020820190506137d75f8301846137b5565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126137fe576137fd6137dd565b5b8235905067ffffffffffffffff81111561381b5761381a6137e1565b5b602083019150836001820283011115613837576138366137e5565b5b9250929050565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6138848261383e565b810181811067ffffffffffffffff821117156138a3576138a261384e565b5b80604052505050565b5f6138b5613719565b90506138c1828261387b565b919050565b5f67ffffffffffffffff8211156138e0576138df61384e565b5b602082029050602081019050919050565b5f60ff82169050919050565b613906816138f1565b8114613910575f5ffd5b50565b5f81359050613921816138fd565b92915050565b5f613939613934846138c6565b6138ac565b9050808382526020820190506020840283018581111561395c5761395b6137e5565b5b835b8181101561398557806139718882613913565b84526020840193505060208101905061395e565b5050509392505050565b5f82601f8301126139a3576139a26137dd565b5b81356139b3848260208601613927565b91505092915050565b5f67ffffffffffffffff8211156139d6576139d561384e565b5b602082029050602081019050919050565b5f819050919050565b6139f9816139e7565b8114613a03575f5ffd5b50565b5f81359050613a14816139f0565b92915050565b5f613a2c613a27846139bc565b6138ac565b90508083825260208201905060208402830185811115613a4f57613a4e6137e5565b5b835b81811015613a785780613a648882613a06565b845260208401935050602081019050613a51565b5050509392505050565b5f82601f830112613a9657613a956137dd565b5b8135613aa6848260208601613a1a565b91505092915050565b5f5f5f5f5f60808688031215613ac857613ac7613722565b5b5f86013567ffffffffffffffff811115613ae557613ae4613726565b5b613af1888289016137e9565b9550955050602086013567ffffffffffffffff811115613b1457613b13613726565b5b613b208882890161398f565b935050604086013567ffffffffffffffff811115613b4157613b40613726565b5b613b4d88828901613a82565b925050606086013567ffffffffffffffff811115613b6e57613b6d613726565b5b613b7a88828901613a82565b9150509295509295909350565b5f819050919050565b613b9981613b87565b82525050565b5f602082019050613bb25f830184613b90565b92915050565b5f60208284031215613bcd57613bcc613722565b5b5f613bda84828501613a06565b91505092915050565b613bec816139e7565b82525050565b5f602082019050613c055f830184613be3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613c3482613c0b565b9050919050565b613c4481613c2a565b8114613c4e575f5ffd5b50565b5f81359050613c5f81613c3b565b92915050565b5f5f60408385031215613c7b57613c7a613722565b5b5f613c8885828601613a06565b9250506020613c9985828601613c51565b9150509250929050565b613cac816138f1565b82525050565b5f602082019050613cc55f830184613ca3565b92915050565b5f604082019050613cde5f830185613b90565b613ceb6020830184613b90565b9392505050565b5f5ffd5b5f67ffffffffffffffff821115613d1057613d0f61384e565b5b613d198261383e565b9050602081019050919050565b828183375f83830152505050565b5f613d46613d4184613cf6565b6138ac565b905082815260208101848484011115613d6257613d61613cf2565b5b613d6d848285613d26565b509392505050565b5f82601f830112613d8957613d886137dd565b5b8135613d99848260208601613d34565b91505092915050565b5f5f60408385031215613db857613db7613722565b5b5f613dc585828601613c51565b925050602083013567ffffffffffffffff811115613de657613de5613726565b5b613df285828601613d75565b9150509250929050565b613e0581613b87565b8114613e0f575f5ffd5b50565b5f81359050613e2081613dfc565b92915050565b5f5f5f60608486031215613e3d57613e3c613722565b5b5f613e4a86828701613e12565b9350506020613e5b86828701613e12565b9250506040613e6c86828701613913565b9150509250925092565b5f60208284031215613e8b57613e8a613722565b5b5f613e9884828501613e12565b91505092915050565b5f67ffffffffffffffff821115613ebb57613eba61384e565b5b602082029050602081019050919050565b5f613ede613ed984613ea1565b6138ac565b90508083825260208201905060208402830185811115613f0157613f006137e5565b5b835b81811015613f2a5780613f168882613c51565b845260208401935050602081019050613f03565b5050509392505050565b5f82601f830112613f4857613f476137dd565b5b8135613f58848260208601613ecc565b91505092915050565b5f5f60408385031215613f7757613f76613722565b5b5f613f8485828601613a06565b925050602083013567ffffffffffffffff811115613fa557613fa4613726565b5b613fb185828601613f34565b9150509250929050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b613fef81613fbb565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61402782613ff5565b6140318185613fff565b935061404181856020860161400f565b61404a8161383e565b840191505092915050565b61405e81613c2a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61409681613b87565b82525050565b5f6140a7838361408d565b60208301905092915050565b5f602082019050919050565b5f6140c982614064565b6140d3818561406e565b93506140de8361407e565b805f5b8381101561410e5781516140f5888261409c565b9750614100836140b3565b9250506001810190506140e1565b5085935050505092915050565b5f60e08201905061412e5f83018a613fe6565b8181036020830152614140818961401d565b90508181036040830152614154818861401d565b90506141636060830187613b90565b6141706080830186614055565b61417d60a0830185613be3565b81810360c083015261418f81846140bf565b905098975050505050505050565b5f67ffffffffffffffff82169050919050565b6141b98161419d565b82525050565b6141c8816139e7565b82525050565b6141d781613c2a565b82525050565b60c082015f8201516141f15f8501826141b0565b50602082015161420460208501826141b0565b50604082015161421760408501826141b0565b50606082015161422a60608501826141b0565b50608082015161423d60808501826141bf565b5060a082015161425060a08501826141ce565b50505050565b5f60c0820190506142695f8301846141dd565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6142a383836141ce565b60208301905092915050565b5f602082019050919050565b5f6142c58261426f565b6142cf8185614279565b93506142da83614289565b805f5b8381101561430a5781516142f18882614298565b97506142fc836142af565b9250506001810190506142dd565b5085935050505092915050565b5f6020820190508181035f83015261432f81846142bb565b905092915050565b5f6020820190508181035f83015261434f818461401d565b905092915050565b5f5f83601f84011261436c5761436b6137dd565b5b8235905067ffffffffffffffff811115614389576143886137e1565b5b6020830191508360208202830111156143a5576143a46137e5565b5b9250929050565b5f5f5f5f606085870312156143c4576143c3613722565b5b5f6143d187828801613e12565b945050602085013567ffffffffffffffff8111156143f2576143f1613726565b5b6143fe87828801614357565b9350935050604061441187828801613a06565b91505092959194509250565b5f6020828403121561443257614431613722565b5b5f61443f84828501613913565b91505092915050565b614451816137aa565b811461445b575f5ffd5b50565b5f8135905061446c81614448565b92915050565b5f6020828403121561448757614486613722565b5b5f6144948482850161445e565b91505092915050565b5f5ffd5b5f60a082840312156144b6576144b561449d565b5b6144c060a06138ac565b90505f6144cf84828501613e12565b5f8301525060206144e284828501613e12565b60208301525060406144f684828501613e12565b604083015250606061450a84828501613a06565b606083015250608061451e84828501613e12565b60808301525092915050565b5f60a0828403121561453f5761453e613722565b5b5f61454c848285016144a1565b91505092915050565b5f60c0820190506145685f830189613be3565b6145756020830188613b90565b6145826040830187613b90565b61458f6060830186613b90565b61459c6080830185613be3565b6145a960a0830184613b90565b979650505050505050565b5f6040820190506145c75f830185613b90565b6145d46020830184613be3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61461282613b87565b915061461d83613b87565b9250828203905081811115614635576146346145db565b5b92915050565b5f60808201905061464e5f830187613b90565b61465b6020830186613b90565b6146686040830185613b90565b6146756060830184613ca3565b95945050505050565b5f819050919050565b5f819050919050565b5f6146aa6146a56146a08461467e565b614687565b61419d565b9050919050565b6146ba81614690565b82525050565b5f6020820190506146d35f8301846146b1565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4549503731323a20556e696e697469616c697a656400000000000000000000005f82015250565b5f61473a601583613fff565b915061474582614706565b602082019050919050565b5f6020820190508181035f8301526147678161472e565b9050919050565b5f61477882613b87565b915061478383613b87565b925082820190508082111561479b5761479a6145db565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6147d882613b87565b91506147e383613b87565b9250826147f3576147f26147a1565b5b828204905092915050565b5f6060820190506148115f830186613b90565b61481e6020830185613b90565b61482b6040830184613b90565b949350505050565b5f6040820190506148465f830185614055565b6148536020830184613be3565b9392505050565b5f81519050614868816139f0565b92915050565b5f6020828403121561488357614882613722565b5b5f6148908482850161485a565b91505092915050565b5f6020820190506148ac5f830184614055565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806148f657607f821691505b602082108103614909576149086148b2565b5b50919050565b5f60a0820190506149225f830188613be3565b61492f6020830187613be3565b61493c6040830186613be3565b6149496060830185613b90565b6149566080830184614055565b9695505050505050565b5f6080820190506149735f830187613be3565b6149806020830186613ca3565b61498d6040830185613be3565b61499a6060830184613be3565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f81519050919050565b5f81905092915050565b5f614a1b826149fd565b614a258185614a07565b9350614a3581856020860161400f565b80840191505092915050565b5f614a4c8284614a11565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302614ab37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614a78565b614abd8683614a78565b95508019841693508086168417925050509392505050565b5f614aef614aea614ae584613b87565b614687565b613b87565b9050919050565b5f819050919050565b614b0883614ad5565b614b1c614b1482614af6565b848454614a84565b825550505050565b5f5f905090565b614b33614b24565b614b3e818484614aff565b505050565b5b81811015614b6157614b565f82614b2b565b600181019050614b44565b5050565b601f821115614ba657614b7781614a57565b614b8084614a69565b81016020851015614b8f578190505b614ba3614b9b85614a69565b830182614b43565b50505b505050565b5f82821c905092915050565b5f614bc65f1984600802614bab565b1980831691505092915050565b5f614bde8383614bb7565b9150826002028217905092915050565b614bf782613ff5565b67ffffffffffffffff811115614c1057614c0f61384e565b5b614c1a82546148df565b614c25828285614b65565b5f60209050601f831160018114614c56575f8415614c44578287015190505b614c4e8582614bd3565b865550614cb5565b601f198416614c6486614a57565b5f5b82811015614c8b57848901518255600182019150602085019450602081019050614c66565b86831015614ca85784890151614ca4601f891682614bb7565b8355505b6001600288020188555050505b50505050505056fea264697066735822122027178bf391b91fef9410f7e085a28ef520843db4b0f3352cd9338896898261a064736f6c634300081c0033"

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

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckSession) ADMINROLE() ([32]byte, error) {
	return _CrossCheck.Contract.ADMINROLE(&_CrossCheck.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CrossCheck *CrossCheckCallerSession) ADMINROLE() ([32]byte, error) {
	return _CrossCheck.Contract.ADMINROLE(&_CrossCheck.CallOpts)
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

// BlocksPerCheck is a free data retrieval call binding the contract method 0xfcc6aca9.
//
// Solidity: function blocksPerCheck() view returns(uint256)
func (_CrossCheck *CrossCheckCaller) BlocksPerCheck(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "blocksPerCheck")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerCheck is a free data retrieval call binding the contract method 0xfcc6aca9.
//
// Solidity: function blocksPerCheck() view returns(uint256)
func (_CrossCheck *CrossCheckSession) BlocksPerCheck() (*big.Int, error) {
	return _CrossCheck.Contract.BlocksPerCheck(&_CrossCheck.CallOpts)
}

// BlocksPerCheck is a free data retrieval call binding the contract method 0xfcc6aca9.
//
// Solidity: function blocksPerCheck() view returns(uint256)
func (_CrossCheck *CrossCheckCallerSession) BlocksPerCheck() (*big.Int, error) {
	return _CrossCheck.Contract.BlocksPerCheck(&_CrossCheck.CallOpts)
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
// Solidity: function getCheckBlock(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCaller) GetCheckBlock(opts *bind.CallOpts, nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "getCheckBlock", nonce)

	if err != nil {
		return *new(CrossCheckBlockCheckBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(CrossCheckBlockCheckBlock)).(*CrossCheckBlockCheckBlock)

	return out0, err

}

// GetCheckBlock is a free data retrieval call binding the contract method 0x871a782e.
//
// Solidity: function getCheckBlock(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckSession) GetCheckBlock(nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlock(&_CrossCheck.CallOpts, nonce)
}

// GetCheckBlock is a free data retrieval call binding the contract method 0x871a782e.
//
// Solidity: function getCheckBlock(uint256 nonce) view returns((uint64,uint64,uint64,uint64,bytes32,address))
func (_CrossCheck *CrossCheckCallerSession) GetCheckBlock(nonce *big.Int) (CrossCheckBlockCheckBlock, error) {
	return _CrossCheck.Contract.GetCheckBlock(&_CrossCheck.CallOpts, nonce)
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

// GetNextBlockInfo is a free data retrieval call binding the contract method 0x4d581907.
//
// Solidity: function getNextBlockInfo() view returns(uint256 _nextNonce, uint256 _nextStart)
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
// Solidity: function getNextBlockInfo() view returns(uint256 _nextNonce, uint256 _nextStart)
func (_CrossCheck *CrossCheckSession) GetNextBlockInfo() (struct {
	NextNonce *big.Int
	NextStart *big.Int
}, error) {
	return _CrossCheck.Contract.GetNextBlockInfo(&_CrossCheck.CallOpts)
}

// GetNextBlockInfo is a free data retrieval call binding the contract method 0x4d581907.
//
// Solidity: function getNextBlockInfo() view returns(uint256 _nextNonce, uint256 _nextStart)
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

// MaxCheckBlocks is a free data retrieval call binding the contract method 0x544d7e1e.
//
// Solidity: function maxCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckCaller) MaxCheckBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossCheck.contract.Call(opts, &out, "maxCheckBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCheckBlocks is a free data retrieval call binding the contract method 0x544d7e1e.
//
// Solidity: function maxCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckSession) MaxCheckBlocks() (*big.Int, error) {
	return _CrossCheck.Contract.MaxCheckBlocks(&_CrossCheck.CallOpts)
}

// MaxCheckBlocks is a free data retrieval call binding the contract method 0x544d7e1e.
//
// Solidity: function maxCheckBlocks() view returns(uint256)
func (_CrossCheck *CrossCheckCallerSession) MaxCheckBlocks() (*big.Int, error) {
	return _CrossCheck.Contract.MaxCheckBlocks(&_CrossCheck.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x55a84bb4.
//
// Solidity: function initialize(uint256 blocksPerCheck_, uint256 maxCheckBlocks_, uint8 validatorThreshold_) returns()
func (_CrossCheck *CrossCheckTransactor) Initialize(opts *bind.TransactOpts, blocksPerCheck_ *big.Int, maxCheckBlocks_ *big.Int, validatorThreshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "initialize", blocksPerCheck_, maxCheckBlocks_, validatorThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x55a84bb4.
//
// Solidity: function initialize(uint256 blocksPerCheck_, uint256 maxCheckBlocks_, uint8 validatorThreshold_) returns()
func (_CrossCheck *CrossCheckSession) Initialize(blocksPerCheck_ *big.Int, maxCheckBlocks_ *big.Int, validatorThreshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.Initialize(&_CrossCheck.TransactOpts, blocksPerCheck_, maxCheckBlocks_, validatorThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0x55a84bb4.
//
// Solidity: function initialize(uint256 blocksPerCheck_, uint256 maxCheckBlocks_, uint8 validatorThreshold_) returns()
func (_CrossCheck *CrossCheckTransactorSession) Initialize(blocksPerCheck_ *big.Int, maxCheckBlocks_ *big.Int, validatorThreshold_ uint8) (*types.Transaction, error) {
	return _CrossCheck.Contract.Initialize(&_CrossCheck.TransactOpts, blocksPerCheck_, maxCheckBlocks_, validatorThreshold_)
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

// SetBlocksPerCheck is a paid mutator transaction binding the contract method 0xc0011114.
//
// Solidity: function setBlocksPerCheck(uint256 blocksPerCheck_) returns()
func (_CrossCheck *CrossCheckTransactor) SetBlocksPerCheck(opts *bind.TransactOpts, blocksPerCheck_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "setBlocksPerCheck", blocksPerCheck_)
}

// SetBlocksPerCheck is a paid mutator transaction binding the contract method 0xc0011114.
//
// Solidity: function setBlocksPerCheck(uint256 blocksPerCheck_) returns()
func (_CrossCheck *CrossCheckSession) SetBlocksPerCheck(blocksPerCheck_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetBlocksPerCheck(&_CrossCheck.TransactOpts, blocksPerCheck_)
}

// SetBlocksPerCheck is a paid mutator transaction binding the contract method 0xc0011114.
//
// Solidity: function setBlocksPerCheck(uint256 blocksPerCheck_) returns()
func (_CrossCheck *CrossCheckTransactorSession) SetBlocksPerCheck(blocksPerCheck_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetBlocksPerCheck(&_CrossCheck.TransactOpts, blocksPerCheck_)
}

// SetMaxCheckBlocks is a paid mutator transaction binding the contract method 0x66cdc83b.
//
// Solidity: function setMaxCheckBlocks(uint256 maxCheckBlocks_) returns()
func (_CrossCheck *CrossCheckTransactor) SetMaxCheckBlocks(opts *bind.TransactOpts, maxCheckBlocks_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "setMaxCheckBlocks", maxCheckBlocks_)
}

// SetMaxCheckBlocks is a paid mutator transaction binding the contract method 0x66cdc83b.
//
// Solidity: function setMaxCheckBlocks(uint256 maxCheckBlocks_) returns()
func (_CrossCheck *CrossCheckSession) SetMaxCheckBlocks(maxCheckBlocks_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetMaxCheckBlocks(&_CrossCheck.TransactOpts, maxCheckBlocks_)
}

// SetMaxCheckBlocks is a paid mutator transaction binding the contract method 0x66cdc83b.
//
// Solidity: function setMaxCheckBlocks(uint256 maxCheckBlocks_) returns()
func (_CrossCheck *CrossCheckTransactorSession) SetMaxCheckBlocks(maxCheckBlocks_ *big.Int) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetMaxCheckBlocks(&_CrossCheck.TransactOpts, maxCheckBlocks_)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossCheck *CrossCheckTransactor) SetPause(opts *bind.TransactOpts, set bool) (*types.Transaction, error) {
	return _CrossCheck.contract.Transact(opts, "setPause", set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossCheck *CrossCheckSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetPause(&_CrossCheck.TransactOpts, set)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool set) returns()
func (_CrossCheck *CrossCheckTransactorSession) SetPause(set bool) (*types.Transaction, error) {
	return _CrossCheck.Contract.SetPause(&_CrossCheck.TransactOpts, set)
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

// CrossCheckCheckBlockAddedIterator is returned from FilterCheckBlockAdded and is used to iterate over the raw logs and unpacked data for CheckBlockAdded events raised by the CrossCheck contract.
type CrossCheckCheckBlockAddedIterator struct {
	Event *CrossCheckCheckBlockAdded // Event containing the contract specifics and raw log

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
func (it *CrossCheckCheckBlockAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckCheckBlockAdded)
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
		it.Event = new(CrossCheckCheckBlockAdded)
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
func (it *CrossCheckCheckBlockAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckCheckBlockAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckCheckBlockAdded represents a CheckBlockAdded event raised by the CrossCheck contract.
type CrossCheckCheckBlockAdded struct {
	Proposer common.Address
	Nonce    *big.Int
	Start    *big.Int
	End      *big.Int
	RootHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCheckBlockAdded is a free log retrieval operation binding the contract event 0xbadf6f34b60f61fb703b4f3ffa385fcf9f98bf0a53cfbadd15fe381c30b145bf.
//
// Solidity: event CheckBlockAdded(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) FilterCheckBlockAdded(opts *bind.FilterOpts, proposer []common.Address, nonce []*big.Int, start []*big.Int) (*CrossCheckCheckBlockAddedIterator, error) {

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

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CheckBlockAdded", proposerRule, nonceRule, startRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckCheckBlockAddedIterator{contract: _CrossCheck.contract, event: "CheckBlockAdded", logs: logs, sub: sub}, nil
}

// WatchCheckBlockAdded is a free log subscription operation binding the contract event 0xbadf6f34b60f61fb703b4f3ffa385fcf9f98bf0a53cfbadd15fe381c30b145bf.
//
// Solidity: event CheckBlockAdded(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) WatchCheckBlockAdded(opts *bind.WatchOpts, sink chan<- *CrossCheckCheckBlockAdded, proposer []common.Address, nonce []*big.Int, start []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "CheckBlockAdded", proposerRule, nonceRule, startRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckCheckBlockAdded)
				if err := _CrossCheck.contract.UnpackLog(event, "CheckBlockAdded", log); err != nil {
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

// ParseCheckBlockAdded is a log parse operation binding the contract event 0xbadf6f34b60f61fb703b4f3ffa385fcf9f98bf0a53cfbadd15fe381c30b145bf.
//
// Solidity: event CheckBlockAdded(address indexed proposer, uint256 indexed nonce, uint256 indexed start, uint256 end, bytes32 rootHash)
func (_CrossCheck *CrossCheckFilterer) ParseCheckBlockAdded(log types.Log) (*CrossCheckCheckBlockAdded, error) {
	event := new(CrossCheckCheckBlockAdded)
	if err := _CrossCheck.contract.UnpackLog(event, "CheckBlockAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckCheckBlockPrunedIterator is returned from FilterCheckBlockPruned and is used to iterate over the raw logs and unpacked data for CheckBlockPruned events raised by the CrossCheck contract.
type CrossCheckCheckBlockPrunedIterator struct {
	Event *CrossCheckCheckBlockPruned // Event containing the contract specifics and raw log

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
func (it *CrossCheckCheckBlockPrunedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckCheckBlockPruned)
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
		it.Event = new(CrossCheckCheckBlockPruned)
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
func (it *CrossCheckCheckBlockPrunedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckCheckBlockPrunedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckCheckBlockPruned represents a CheckBlockPruned event raised by the CrossCheck contract.
type CrossCheckCheckBlockPruned struct {
	Nonce *big.Int
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCheckBlockPruned is a free log retrieval operation binding the contract event 0xb8a4623367e916c9ab30fca57e7f767aaa221385f9e2726b7cd24bffde7385e2.
//
// Solidity: event CheckBlockPruned(uint256 indexed nonce, uint256 count)
func (_CrossCheck *CrossCheckFilterer) FilterCheckBlockPruned(opts *bind.FilterOpts, nonce []*big.Int) (*CrossCheckCheckBlockPrunedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CheckBlockPruned", nonceRule)
	if err != nil {
		return nil, err
	}
	return &CrossCheckCheckBlockPrunedIterator{contract: _CrossCheck.contract, event: "CheckBlockPruned", logs: logs, sub: sub}, nil
}

// WatchCheckBlockPruned is a free log subscription operation binding the contract event 0xb8a4623367e916c9ab30fca57e7f767aaa221385f9e2726b7cd24bffde7385e2.
//
// Solidity: event CheckBlockPruned(uint256 indexed nonce, uint256 count)
func (_CrossCheck *CrossCheckFilterer) WatchCheckBlockPruned(opts *bind.WatchOpts, sink chan<- *CrossCheckCheckBlockPruned, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "CheckBlockPruned", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckCheckBlockPruned)
				if err := _CrossCheck.contract.UnpackLog(event, "CheckBlockPruned", log); err != nil {
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

// ParseCheckBlockPruned is a log parse operation binding the contract event 0xb8a4623367e916c9ab30fca57e7f767aaa221385f9e2726b7cd24bffde7385e2.
//
// Solidity: event CheckBlockPruned(uint256 indexed nonce, uint256 count)
func (_CrossCheck *CrossCheckFilterer) ParseCheckBlockPruned(log types.Log) (*CrossCheckCheckBlockPruned, error) {
	event := new(CrossCheckCheckBlockPruned)
	if err := _CrossCheck.contract.UnpackLog(event, "CheckBlockPruned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckCrossCheckBlocksPerCheckUpdatedIterator is returned from FilterCrossCheckBlocksPerCheckUpdated and is used to iterate over the raw logs and unpacked data for CrossCheckBlocksPerCheckUpdated events raised by the CrossCheck contract.
type CrossCheckCrossCheckBlocksPerCheckUpdatedIterator struct {
	Event *CrossCheckCrossCheckBlocksPerCheckUpdated // Event containing the contract specifics and raw log

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
func (it *CrossCheckCrossCheckBlocksPerCheckUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckCrossCheckBlocksPerCheckUpdated)
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
		it.Event = new(CrossCheckCrossCheckBlocksPerCheckUpdated)
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
func (it *CrossCheckCrossCheckBlocksPerCheckUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckCrossCheckBlocksPerCheckUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckCrossCheckBlocksPerCheckUpdated represents a CrossCheckBlocksPerCheckUpdated event raised by the CrossCheck contract.
type CrossCheckCrossCheckBlocksPerCheckUpdated struct {
	BlocksPerCheck *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrossCheckBlocksPerCheckUpdated is a free log retrieval operation binding the contract event 0xc37172b02acb2a6b00612c1178f47d9f660e5fa7800fa7a3904e707766565912.
//
// Solidity: event CrossCheckBlocksPerCheckUpdated(uint256 blocksPerCheck)
func (_CrossCheck *CrossCheckFilterer) FilterCrossCheckBlocksPerCheckUpdated(opts *bind.FilterOpts) (*CrossCheckCrossCheckBlocksPerCheckUpdatedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CrossCheckBlocksPerCheckUpdated")
	if err != nil {
		return nil, err
	}
	return &CrossCheckCrossCheckBlocksPerCheckUpdatedIterator{contract: _CrossCheck.contract, event: "CrossCheckBlocksPerCheckUpdated", logs: logs, sub: sub}, nil
}

// WatchCrossCheckBlocksPerCheckUpdated is a free log subscription operation binding the contract event 0xc37172b02acb2a6b00612c1178f47d9f660e5fa7800fa7a3904e707766565912.
//
// Solidity: event CrossCheckBlocksPerCheckUpdated(uint256 blocksPerCheck)
func (_CrossCheck *CrossCheckFilterer) WatchCrossCheckBlocksPerCheckUpdated(opts *bind.WatchOpts, sink chan<- *CrossCheckCrossCheckBlocksPerCheckUpdated) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "CrossCheckBlocksPerCheckUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckCrossCheckBlocksPerCheckUpdated)
				if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckBlocksPerCheckUpdated", log); err != nil {
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

// ParseCrossCheckBlocksPerCheckUpdated is a log parse operation binding the contract event 0xc37172b02acb2a6b00612c1178f47d9f660e5fa7800fa7a3904e707766565912.
//
// Solidity: event CrossCheckBlocksPerCheckUpdated(uint256 blocksPerCheck)
func (_CrossCheck *CrossCheckFilterer) ParseCrossCheckBlocksPerCheckUpdated(log types.Log) (*CrossCheckCrossCheckBlocksPerCheckUpdated, error) {
	event := new(CrossCheckCrossCheckBlocksPerCheckUpdated)
	if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckBlocksPerCheckUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	MaxCheckBlocks     *big.Int
	ValidatorThreshold uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCrossCheckInitialized is a free log retrieval operation binding the contract event 0x38778a692f0977d8b900e132518cdf1350918c14b717a3f68825d1e067476bc1.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint256 maxCheckBlocks, uint8 validatorThreshold)
func (_CrossCheck *CrossCheckFilterer) FilterCrossCheckInitialized(opts *bind.FilterOpts) (*CrossCheckCrossCheckInitializedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CrossCheckInitialized")
	if err != nil {
		return nil, err
	}
	return &CrossCheckCrossCheckInitializedIterator{contract: _CrossCheck.contract, event: "CrossCheckInitialized", logs: logs, sub: sub}, nil
}

// WatchCrossCheckInitialized is a free log subscription operation binding the contract event 0x38778a692f0977d8b900e132518cdf1350918c14b717a3f68825d1e067476bc1.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint256 maxCheckBlocks, uint8 validatorThreshold)
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

// ParseCrossCheckInitialized is a log parse operation binding the contract event 0x38778a692f0977d8b900e132518cdf1350918c14b717a3f68825d1e067476bc1.
//
// Solidity: event CrossCheckInitialized(uint256 chainID, uint256 blocksPerCheck, uint256 maxCheckBlocks, uint8 validatorThreshold)
func (_CrossCheck *CrossCheckFilterer) ParseCrossCheckInitialized(log types.Log) (*CrossCheckCrossCheckInitialized, error) {
	event := new(CrossCheckCrossCheckInitialized)
	if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator is returned from FilterCrossCheckMaxCheckBlocksUpdated and is used to iterate over the raw logs and unpacked data for CrossCheckMaxCheckBlocksUpdated events raised by the CrossCheck contract.
type CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator struct {
	Event *CrossCheckCrossCheckMaxCheckBlocksUpdated // Event containing the contract specifics and raw log

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
func (it *CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCheckCrossCheckMaxCheckBlocksUpdated)
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
		it.Event = new(CrossCheckCrossCheckMaxCheckBlocksUpdated)
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
func (it *CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCheckCrossCheckMaxCheckBlocksUpdated represents a CrossCheckMaxCheckBlocksUpdated event raised by the CrossCheck contract.
type CrossCheckCrossCheckMaxCheckBlocksUpdated struct {
	MaxCheckBlocks *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrossCheckMaxCheckBlocksUpdated is a free log retrieval operation binding the contract event 0xf3114416d687e19a2ba81f92f6b8a93a48863d39989e5fd2f1adb0f5f89fd699.
//
// Solidity: event CrossCheckMaxCheckBlocksUpdated(uint256 maxCheckBlocks)
func (_CrossCheck *CrossCheckFilterer) FilterCrossCheckMaxCheckBlocksUpdated(opts *bind.FilterOpts) (*CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator, error) {

	logs, sub, err := _CrossCheck.contract.FilterLogs(opts, "CrossCheckMaxCheckBlocksUpdated")
	if err != nil {
		return nil, err
	}
	return &CrossCheckCrossCheckMaxCheckBlocksUpdatedIterator{contract: _CrossCheck.contract, event: "CrossCheckMaxCheckBlocksUpdated", logs: logs, sub: sub}, nil
}

// WatchCrossCheckMaxCheckBlocksUpdated is a free log subscription operation binding the contract event 0xf3114416d687e19a2ba81f92f6b8a93a48863d39989e5fd2f1adb0f5f89fd699.
//
// Solidity: event CrossCheckMaxCheckBlocksUpdated(uint256 maxCheckBlocks)
func (_CrossCheck *CrossCheckFilterer) WatchCrossCheckMaxCheckBlocksUpdated(opts *bind.WatchOpts, sink chan<- *CrossCheckCrossCheckMaxCheckBlocksUpdated) (event.Subscription, error) {

	logs, sub, err := _CrossCheck.contract.WatchLogs(opts, "CrossCheckMaxCheckBlocksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCheckCrossCheckMaxCheckBlocksUpdated)
				if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckMaxCheckBlocksUpdated", log); err != nil {
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

// ParseCrossCheckMaxCheckBlocksUpdated is a log parse operation binding the contract event 0xf3114416d687e19a2ba81f92f6b8a93a48863d39989e5fd2f1adb0f5f89fd699.
//
// Solidity: event CrossCheckMaxCheckBlocksUpdated(uint256 maxCheckBlocks)
func (_CrossCheck *CrossCheckFilterer) ParseCrossCheckMaxCheckBlocksUpdated(log types.Log) (*CrossCheckCrossCheckMaxCheckBlocksUpdated, error) {
	event := new(CrossCheckCrossCheckMaxCheckBlocksUpdated)
	if err := _CrossCheck.contract.UnpackLog(event, "CrossCheckMaxCheckBlocksUpdated", log); err != nil {
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
