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

// CrossMintableERC20MetaData contains all meta data concerning the CrossMintableERC20 contract.
var CrossMintableERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CrossMintableERC20NotMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"e78cea92": "bridge()",
		"9dc29fac": "burn(address,uint256)",
		"313ce567": "decimals()",
		"84b0196e": "eip712Domain()",
		"726600ce": "isBridge(address)",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"7ecebe00": "nonces(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"715018a6": "renounceOwnership()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x610180604052348015610010575f5ffd5b5060405161196638038061196683398101604081905261002f916103c1565b8280604051806040016040528060018152602001603160f81b815250858561005b6101f760201b60201c565b6001600160a01b03811661008957604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610092816101fb565b505f805460ff60a01b1916905560046100ab83826104c2565b5060056100b882826104c2565b506100c89150839050600661024a565b610120526100d781600761024a565b61014052815160208084019190912060e052815190820120610100524660a05261016360e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600980546001600160a01b0319163390811790915560408051638da5cb5b60e01b815290516101e89291638da5cb5b9160048083019260209291908290030181865afa1580156101bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e3919061057c565b61027c565b60ff1661016052506106019050565b3390565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020835110156102655761025e836102b9565b9050610276565b8161027084826104c2565b5060ff90505b92915050565b6102846102f6565b6001600160a01b0381166102ad57604051631e4fbdf760e01b81525f6004820152602401610080565b6102b6816101fb565b50565b5f5f829050601f815111156102e3578260405163305a27a960e01b815260040161008091906105a9565b80516102ee826105de565b179392505050565b5f546001600160a01b031633146103225760405163118cdaa760e01b8152336004820152602401610080565b565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610347575f5ffd5b81516001600160401b0381111561036057610360610324565b604051601f8201601f19908116603f011681016001600160401b038111828210171561038e5761038e610324565b6040528181528382016020018510156103a5575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f606084860312156103d3575f5ffd5b83516001600160401b038111156103e8575f5ffd5b6103f486828701610338565b602086015190945090506001600160401b03811115610411575f5ffd5b61041d86828701610338565b925050604084015160ff81168114610433575f5ffd5b809150509250925092565b600181811c9082168061045257607f821691505b60208210810361047057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156104bd57805f5260205f20601f840160051c8101602085101561049b5750805b601f840160051c820191505b818110156104ba575f81556001016104a7565b50505b505050565b81516001600160401b038111156104db576104db610324565b6104ef816104e9845461043e565b84610476565b6020601f821160018114610521575f831561050a5750848201515b5f19600385901b1c1916600184901b1784556104ba565b5f84815260208120601f198516915b828110156105505787850151825560209485019460019092019101610530565b508482101561056d57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6020828403121561058c575f5ffd5b81516001600160a01b03811681146105a2575f5ffd5b9392505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610470575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516101605161130a61065c5f395f6101b901525f610a4401525f610a1701525f61085a01525f61083201525f61078d01525f6107b701525f6107e1015261130a5ff3fe608060405234801561000f575f5ffd5b5060043610610148575f3560e01c8063726600ce116100bf5780639dc29fac116100795780639dc29fac146102cd578063a9059cbb146102e0578063d505accf146102f3578063dd62ed3e14610306578063e78cea921461033e578063f2fde38b14610351575f5ffd5b8063726600ce146102495780637ecebe001461026b5780638456cb591461027e57806384b0196e146102865780638da5cb5b146102a157806395d89b41146102c5575f5ffd5b80633644e515116101105780633644e515146101e35780633f4ba83a146101eb57806340c10f19146101f55780635c975abb1461020857806370a0823114610219578063715018a614610241575f5ffd5b806306fdde031461014c578063095ea7b31461016a57806318160ddd1461018d57806323b872dd1461019f578063313ce567146101b2575b5f5ffd5b610154610364565b6040516101619190611086565b60405180910390f35b61017d6101783660046110ba565b6103f4565b6040519015158152602001610161565b6003545b604051908152602001610161565b61017d6101ad3660046110e2565b61040d565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610161565b610191610445565b6101f3610453565b005b61017d6102033660046110ba565b610465565b5f54600160a01b900460ff1661017d565b61019161022736600461111c565b6001600160a01b03165f9081526001602052604090205490565b6101f36104b3565b61017d61025736600461111c565b6009546001600160a01b0390811691161490565b61019161027936600461111c565b6104c4565b6101f36104ce565b61028e6104de565b6040516101619796959493929190611135565b5f546001600160a01b03165b6040516001600160a01b039091168152602001610161565b610154610520565b61017d6102db3660046110ba565b61052f565b61017d6102ee3660046110ba565b61056f565b6101f36103013660046111cb565b61057c565b610191610314366004611238565b6001600160a01b039182165f90815260026020908152604080832093909416825291909152205490565b6009546102ad906001600160a01b031681565b6101f361035f36600461111c565b6106b2565b60606004805461037390611269565b80601f016020809104026020016040519081016040528092919081815260200182805461039f90611269565b80156103ea5780601f106103c1576101008083540402835291602001916103ea565b820191905f5260205f20905b8154815290600101906020018083116103cd57829003601f168201915b5050505050905090565b5f336104018185856106ef565b60019150505b92915050565b5f5f1961041a8533610314565b146104305761042a848484610701565b5061043b565b61043b848484610724565b5060019392505050565b5f61044e610781565b905090565b61045b6108aa565b6104636108d6565b565b5f61046f33610257565b339061049f576040516313a3eb9960e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b506104aa838361092a565b50600192915050565b6104bb6108aa565b6104635f610962565b5f610407826109b1565b6104d66108aa565b6104636109ce565b5f6060805f5f5f60606104ef610a10565b6104f7610a3d565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606005805461037390611269565b5f61053933610257565b3390610564576040516313a3eb9960e01b81526001600160a01b039091166004820152602401610496565b506104aa8383610a6a565b5f33610401818585610724565b834211156105a05760405163313c898160e11b815260048101859052602401610496565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105eb8c6001600160a01b03165f90815260086020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61064582610a9e565b90505f61065482878787610aca565b9050896001600160a01b0316816001600160a01b03161461069b576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610496565b6106a68a8a8a6106ef565b50505050505050505050565b6106ba6108aa565b6001600160a01b0381166106e357604051631e4fbdf760e01b81525f6004820152602401610496565b6106ec81610962565b50565b6106fc8383836001610af6565b505050565b5f3361070e858285610bc9565b610719858585610724565b506001949350505050565b6001600160a01b03831661074d57604051634b637e8f60e11b81525f6004820152602401610496565b6001600160a01b0382166107765760405163ec442f0560e01b81525f6004820152602401610496565b6106fc838383610c3f565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156107d957507f000000000000000000000000000000000000000000000000000000000000000046145b1561080357507f000000000000000000000000000000000000000000000000000000000000000090565b61044e604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f546001600160a01b031633146104635760405163118cdaa760e01b8152336004820152602401610496565b6108de610c52565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b0382166109535760405163ec442f0560e01b81525f6004820152602401610496565b61095e5f8383610c3f565b5050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381165f90815260086020526040812054610407565b6109d6610c7b565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861090d3390565b606061044e7f00000000000000000000000000000000000000000000000000000000000000006006610ca5565b606061044e7f00000000000000000000000000000000000000000000000000000000000000006007610ca5565b6001600160a01b038216610a9357604051634b637e8f60e11b81525f6004820152602401610496565b61095e825f83610c3f565b5f610407610aaa610781565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610ada88888888610d4e565b925092509250610aea8282610e16565b50909695505050505050565b6001600160a01b038416610b1f5760405163e602df0560e01b81525f6004820152602401610496565b6001600160a01b038316610b4857604051634a1406b160e11b81525f6004820152602401610496565b6001600160a01b038085165f9081526002602090815260408083209387168352929052208290558015610bc357826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bba91815260200190565b60405180910390a35b50505050565b6001600160a01b038381165f908152600260209081526040808320938616835292905220545f19811015610bc35781811015610c3157604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610496565b610bc384848484035f610af6565b610c47610c7b565b6106fc838383610ece565b5f54600160a01b900460ff1661046357604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff16156104635760405163d93c066560e01b815260040160405180910390fd5b606060ff8314610cbf57610cb883610ff4565b9050610407565b818054610ccb90611269565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf790611269565b8015610d425780601f10610d1957610100808354040283529160200191610d42565b820191905f5260205f20905b815481529060010190602001808311610d2557829003601f168201915b50505050509050610407565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610d8757505f91506003905082610e0c565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610dd8573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610e0357505f925060019150829050610e0c565b92505f91508190505b9450945094915050565b5f826003811115610e2957610e296112a1565b03610e32575050565b6001826003811115610e4657610e466112a1565b03610e645760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610e7857610e786112a1565b03610e995760405163fce698f760e01b815260048101829052602401610496565b6003826003811115610ead57610ead6112a1565b0361095e576040516335e2f38360e21b815260048101829052602401610496565b6001600160a01b038316610ef8578060035f828254610eed91906112b5565b90915550610f689050565b6001600160a01b0383165f9081526001602052604090205481811015610f4a5760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610496565b6001600160a01b0384165f9081526001602052604090209082900390555b6001600160a01b038216610f8457600380548290039055610fa2565b6001600160a01b0382165f9081526001602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fe791815260200190565b60405180910390a3505050565b60605f61100083611031565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561040757604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6110986020830184611058565b9392505050565b80356001600160a01b03811681146110b5575f5ffd5b919050565b5f5f604083850312156110cb575f5ffd5b6110d48361109f565b946020939093013593505050565b5f5f5f606084860312156110f4575f5ffd5b6110fd8461109f565b925061110b6020850161109f565b929592945050506040919091013590565b5f6020828403121561112c575f5ffd5b6110988261109f565b60ff60f81b8816815260e060208201525f61115360e0830189611058565b82810360408401526111658189611058565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b818110156111ba57835183526020938401939092019160010161119c565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a0312156111e1575f5ffd5b6111ea8861109f565b96506111f86020890161109f565b95506040880135945060608801359350608088013560ff8116811461121b575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215611249575f5ffd5b6112528361109f565b91506112606020840161109f565b90509250929050565b600181811c9082168061127d57607f821691505b60208210810361129b57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b8082018082111561040757634e487b7160e01b5f52601160045260245ffdfea264697066735822122076e6997885caeb35884c196bee40080226912da8b7dac8c0f37c79ba5aa19ee964736f6c634300081c0033",
}

// CrossMintableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossMintableERC20MetaData.ABI instead.
var CrossMintableERC20ABI = CrossMintableERC20MetaData.ABI

// Deprecated: Use CrossMintableERC20MetaData.Sigs instead.
// CrossMintableERC20FuncSigs maps the 4-byte function signature to its string representation.
var CrossMintableERC20FuncSigs = CrossMintableERC20MetaData.Sigs

// CrossMintableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossMintableERC20MetaData.Bin instead.
var CrossMintableERC20Bin = CrossMintableERC20MetaData.Bin

// DeployCrossMintableERC20 deploys a new Ethereum contract, binding an instance of CrossMintableERC20 to it.
func DeployCrossMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *CrossMintableERC20, error) {
	parsed, err := CrossMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossMintableERC20Bin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossMintableERC20{CrossMintableERC20Caller: CrossMintableERC20Caller{contract: contract}, CrossMintableERC20Transactor: CrossMintableERC20Transactor{contract: contract}, CrossMintableERC20Filterer: CrossMintableERC20Filterer{contract: contract}}, nil
}

// CrossMintableERC20 is an auto generated Go binding around an Ethereum contract.
type CrossMintableERC20 struct {
	CrossMintableERC20Caller     // Read-only binding to the contract
	CrossMintableERC20Transactor // Write-only binding to the contract
	CrossMintableERC20Filterer   // Log filterer for contract events
}

// CrossMintableERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossMintableERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossMintableERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossMintableERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossMintableERC20Session struct {
	Contract     *CrossMintableERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CrossMintableERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossMintableERC20CallerSession struct {
	Contract *CrossMintableERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CrossMintableERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossMintableERC20TransactorSession struct {
	Contract     *CrossMintableERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CrossMintableERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossMintableERC20Raw struct {
	Contract *CrossMintableERC20 // Generic contract binding to access the raw methods on
}

// CrossMintableERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossMintableERC20CallerRaw struct {
	Contract *CrossMintableERC20Caller // Generic read-only contract binding to access the raw methods on
}

// CrossMintableERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossMintableERC20TransactorRaw struct {
	Contract *CrossMintableERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossMintableERC20 creates a new instance of CrossMintableERC20, bound to a specific deployed contract.
func NewCrossMintableERC20(address common.Address, backend bind.ContractBackend) (*CrossMintableERC20, error) {
	contract, err := bindCrossMintableERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20{CrossMintableERC20Caller: CrossMintableERC20Caller{contract: contract}, CrossMintableERC20Transactor: CrossMintableERC20Transactor{contract: contract}, CrossMintableERC20Filterer: CrossMintableERC20Filterer{contract: contract}}, nil
}

// NewCrossMintableERC20Caller creates a new read-only instance of CrossMintableERC20, bound to a specific deployed contract.
func NewCrossMintableERC20Caller(address common.Address, caller bind.ContractCaller) (*CrossMintableERC20Caller, error) {
	contract, err := bindCrossMintableERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20Caller{contract: contract}, nil
}

// NewCrossMintableERC20Transactor creates a new write-only instance of CrossMintableERC20, bound to a specific deployed contract.
func NewCrossMintableERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossMintableERC20Transactor, error) {
	contract, err := bindCrossMintableERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20Transactor{contract: contract}, nil
}

// NewCrossMintableERC20Filterer creates a new log filterer instance of CrossMintableERC20, bound to a specific deployed contract.
func NewCrossMintableERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossMintableERC20Filterer, error) {
	contract, err := bindCrossMintableERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20Filterer{contract: contract}, nil
}

// bindCrossMintableERC20 binds a generic wrapper to an already deployed contract.
func bindCrossMintableERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossMintableERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20 *CrossMintableERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20.Contract.CrossMintableERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20 *CrossMintableERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.CrossMintableERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20 *CrossMintableERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.CrossMintableERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20 *CrossMintableERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20 *CrossMintableERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20 *CrossMintableERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20 *CrossMintableERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20 *CrossMintableERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _CrossMintableERC20.Contract.DOMAINSEPARATOR(&_CrossMintableERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CrossMintableERC20.Contract.DOMAINSEPARATOR(&_CrossMintableERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.Allowance(&_CrossMintableERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.Allowance(&_CrossMintableERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.BalanceOf(&_CrossMintableERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.BalanceOf(&_CrossMintableERC20.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20Session) Bridge() (common.Address, error) {
	return _CrossMintableERC20.Contract.Bridge(&_CrossMintableERC20.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Bridge() (common.Address, error) {
	return _CrossMintableERC20.Contract.Bridge(&_CrossMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20 *CrossMintableERC20Session) Decimals() (uint8, error) {
	return _CrossMintableERC20.Contract.Decimals(&_CrossMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Decimals() (uint8, error) {
	return _CrossMintableERC20.Contract.Decimals(&_CrossMintableERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "eip712Domain")

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
func (_CrossMintableERC20 *CrossMintableERC20Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossMintableERC20.Contract.Eip712Domain(&_CrossMintableERC20.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CrossMintableERC20.Contract.Eip712Domain(&_CrossMintableERC20.CallOpts)
}

// IsBridge is a free data retrieval call binding the contract method 0x726600ce.
//
// Solidity: function isBridge(address _account) view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Caller) IsBridge(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "isBridge", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridge is a free data retrieval call binding the contract method 0x726600ce.
//
// Solidity: function isBridge(address _account) view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) IsBridge(_account common.Address) (bool, error) {
	return _CrossMintableERC20.Contract.IsBridge(&_CrossMintableERC20.CallOpts, _account)
}

// IsBridge is a free data retrieval call binding the contract method 0x726600ce.
//
// Solidity: function isBridge(address _account) view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) IsBridge(_account common.Address) (bool, error) {
	return _CrossMintableERC20.Contract.IsBridge(&_CrossMintableERC20.CallOpts, _account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20Session) Name() (string, error) {
	return _CrossMintableERC20.Contract.Name(&_CrossMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Name() (string, error) {
	return _CrossMintableERC20.Contract.Name(&_CrossMintableERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.Nonces(&_CrossMintableERC20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _CrossMintableERC20.Contract.Nonces(&_CrossMintableERC20.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20Session) Owner() (common.Address, error) {
	return _CrossMintableERC20.Contract.Owner(&_CrossMintableERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Owner() (common.Address, error) {
	return _CrossMintableERC20.Contract.Owner(&_CrossMintableERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) Paused() (bool, error) {
	return _CrossMintableERC20.Contract.Paused(&_CrossMintableERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Paused() (bool, error) {
	return _CrossMintableERC20.Contract.Paused(&_CrossMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20Session) Symbol() (string, error) {
	return _CrossMintableERC20.Contract.Symbol(&_CrossMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) Symbol() (string, error) {
	return _CrossMintableERC20.Contract.Symbol(&_CrossMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20Session) TotalSupply() (*big.Int, error) {
	return _CrossMintableERC20.Contract.TotalSupply(&_CrossMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrossMintableERC20 *CrossMintableERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _CrossMintableERC20.Contract.TotalSupply(&_CrossMintableERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Approve(&_CrossMintableERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Approve(&_CrossMintableERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Burn(&_CrossMintableERC20.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Burn(&_CrossMintableERC20.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Mint(&_CrossMintableERC20.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Mint(&_CrossMintableERC20.TransactOpts, _account, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossMintableERC20 *CrossMintableERC20Session) Pause() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Pause(&_CrossMintableERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Pause(&_CrossMintableERC20.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20 *CrossMintableERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Permit(&_CrossMintableERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Permit(&_CrossMintableERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossMintableERC20 *CrossMintableERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossMintableERC20 *CrossMintableERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.RenounceOwnership(&_CrossMintableERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.RenounceOwnership(&_CrossMintableERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Transfer(&_CrossMintableERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Transfer(&_CrossMintableERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferFrom(&_CrossMintableERC20.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferFrom(&_CrossMintableERC20.TransactOpts, _sender, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossMintableERC20 *CrossMintableERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossMintableERC20 *CrossMintableERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferOwnership(&_CrossMintableERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferOwnership(&_CrossMintableERC20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossMintableERC20 *CrossMintableERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossMintableERC20 *CrossMintableERC20Session) Unpause() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Unpause(&_CrossMintableERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.Unpause(&_CrossMintableERC20.TransactOpts)
}

// CrossMintableERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CrossMintableERC20 contract.
type CrossMintableERC20ApprovalIterator struct {
	Event *CrossMintableERC20Approval // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20Approval)
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
		it.Event = new(CrossMintableERC20Approval)
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
func (it *CrossMintableERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20Approval represents a Approval event raised by the CrossMintableERC20 contract.
type CrossMintableERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossMintableERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20ApprovalIterator{contract: _CrossMintableERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20Approval)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParseApproval(log types.Log) (*CrossMintableERC20Approval, error) {
	event := new(CrossMintableERC20Approval)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the CrossMintableERC20 contract.
type CrossMintableERC20EIP712DomainChangedIterator struct {
	Event *CrossMintableERC20EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20EIP712DomainChanged)
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
		it.Event = new(CrossMintableERC20EIP712DomainChanged)
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
func (it *CrossMintableERC20EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20EIP712DomainChanged represents a EIP712DomainChanged event raised by the CrossMintableERC20 contract.
type CrossMintableERC20EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*CrossMintableERC20EIP712DomainChangedIterator, error) {

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20EIP712DomainChangedIterator{contract: _CrossMintableERC20.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20EIP712DomainChanged)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParseEIP712DomainChanged(log types.Log) (*CrossMintableERC20EIP712DomainChanged, error) {
	event := new(CrossMintableERC20EIP712DomainChanged)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossMintableERC20 contract.
type CrossMintableERC20OwnershipTransferredIterator struct {
	Event *CrossMintableERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20OwnershipTransferred)
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
		it.Event = new(CrossMintableERC20OwnershipTransferred)
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
func (it *CrossMintableERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20OwnershipTransferred represents a OwnershipTransferred event raised by the CrossMintableERC20 contract.
type CrossMintableERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossMintableERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20OwnershipTransferredIterator{contract: _CrossMintableERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20OwnershipTransferred)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParseOwnershipTransferred(log types.Log) (*CrossMintableERC20OwnershipTransferred, error) {
	event := new(CrossMintableERC20OwnershipTransferred)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossMintableERC20 contract.
type CrossMintableERC20PausedIterator struct {
	Event *CrossMintableERC20Paused // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20Paused)
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
		it.Event = new(CrossMintableERC20Paused)
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
func (it *CrossMintableERC20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20Paused represents a Paused event raised by the CrossMintableERC20 contract.
type CrossMintableERC20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterPaused(opts *bind.FilterOpts) (*CrossMintableERC20PausedIterator, error) {

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20PausedIterator{contract: _CrossMintableERC20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20Paused) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20Paused)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParsePaused(log types.Log) (*CrossMintableERC20Paused, error) {
	event := new(CrossMintableERC20Paused)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CrossMintableERC20 contract.
type CrossMintableERC20TransferIterator struct {
	Event *CrossMintableERC20Transfer // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20Transfer)
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
		it.Event = new(CrossMintableERC20Transfer)
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
func (it *CrossMintableERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20Transfer represents a Transfer event raised by the CrossMintableERC20 contract.
type CrossMintableERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossMintableERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20TransferIterator{contract: _CrossMintableERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20Transfer)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParseTransfer(log types.Log) (*CrossMintableERC20Transfer, error) {
	event := new(CrossMintableERC20Transfer)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossMintableERC20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossMintableERC20 contract.
type CrossMintableERC20UnpausedIterator struct {
	Event *CrossMintableERC20Unpaused // Event containing the contract specifics and raw log

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
func (it *CrossMintableERC20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossMintableERC20Unpaused)
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
		it.Event = new(CrossMintableERC20Unpaused)
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
func (it *CrossMintableERC20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossMintableERC20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossMintableERC20Unpaused represents a Unpaused event raised by the CrossMintableERC20 contract.
type CrossMintableERC20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossMintableERC20UnpausedIterator, error) {

	logs, sub, err := _CrossMintableERC20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20UnpausedIterator{contract: _CrossMintableERC20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossMintableERC20 *CrossMintableERC20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossMintableERC20Unpaused) (event.Subscription, error) {

	logs, sub, err := _CrossMintableERC20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossMintableERC20Unpaused)
				if err := _CrossMintableERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CrossMintableERC20 *CrossMintableERC20Filterer) ParseUnpaused(log types.Log) (*CrossMintableERC20Unpaused, error) {
	event := new(CrossMintableERC20Unpaused)
	if err := _CrossMintableERC20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
