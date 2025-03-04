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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errCrossMintableERC20NotBridge\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"3644e515": "DOMAIN_SEPARATOR()",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"e78cea92": "bridge()",
		"9dc29fac": "burn(address,uint256)",
		"313ce567": "decimals()",
		"84b0196e": "eip712Domain()",
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
	Bin: "0x610180604052348015610010575f5ffd5b5060405161183c38038061183c83398101604081905261002f916102fb565b8280604051806040016040528060018152602001603160f81b815250858561005b61019c60201b60201c565b6001600160a01b03811661008957604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610092816101a0565b505f805460ff60a01b1916905560046100ab8382610419565b5060056100b88282610419565b506100c8915083905060066101ef565b610120526100d78160076101ef565b61014052815160208084019190912060e052815190820120610100524660a05261016360e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600980546001600160a01b0319166001600160a01b039590951694909417909355505060ff166101605261052b565b3390565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208351101561020a5761020383610221565b905061021b565b816102158482610419565b5060ff90505b92915050565b5f5f829050601f8151111561024b578260405163305a27a960e01b815260040161008091906104d3565b805161025682610508565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610281575f5ffd5b81516001600160401b0381111561029a5761029a61025e565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102c8576102c861025e565b6040528181528382016020018510156102df575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f6080858703121561030e575f5ffd5b84516001600160a01b0381168114610324575f5ffd5b60208601519094506001600160401b0381111561033f575f5ffd5b61034b87828801610272565b604087015190945090506001600160401b03811115610368575f5ffd5b61037487828801610272565b925050606085015160ff8116811461038a575f5ffd5b939692955090935050565b600181811c908216806103a957607f821691505b6020821081036103c757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561041457805f5260205f20601f840160051c810160208510156103f25750805b601f840160051c820191505b81811015610411575f81556001016103fe565b50505b505050565b81516001600160401b038111156104325761043261025e565b610446816104408454610395565b846103cd565b6020601f821160018114610478575f83156104615750848201515b5f19600385901b1c1916600184901b178455610411565b5f84815260208120601f198516915b828110156104a75787850151825560209485019460019092019101610487565b50848210156104c457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156103c7575f1960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051610160516112b66105865f395f6101ae01525f610a6701525f610a3a01525f61087d01525f61085501525f6107b001525f6107da01525f61080401526112b65ff3fe608060405234801561000f575f5ffd5b506004361061013d575f3560e01c80637ecebe00116100b45780639dc29fac116100795780639dc29fac146102a0578063a9059cbb146102b3578063d505accf146102c6578063dd62ed3e146102d9578063e78cea9214610311578063f2fde38b14610324575f5ffd5b80637ecebe001461023e5780638456cb591461025157806384b0196e146102595780638da5cb5b1461027457806395d89b4114610298575f5ffd5b80633644e515116101055780633644e515146101d85780633f4ba83a146101e057806340c10f19146101ea5780635c975abb146101fd57806370a082311461020e578063715018a614610236575f5ffd5b806306fdde0314610141578063095ea7b31461015f57806318160ddd1461018257806323b872dd14610194578063313ce567146101a7575b5f5ffd5b610149610337565b6040516101569190611032565b60405180910390f35b61017261016d366004611066565b6103c7565b6040519015158152602001610156565b6003545b604051908152602001610156565b6101726101a236600461108e565b6103e0565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610156565b610186610403565b6101e8610411565b005b6101726101f8366004611066565b610423565b5f54600160a01b900460ff16610172565b61018661021c3660046110c8565b6001600160a01b03165f9081526001602052604090205490565b6101e8610477565b61018661024c3660046110c8565b610488565b6101e8610492565b6102616104a2565b60405161015697969594939291906110e1565b5f546001600160a01b03165b6040516001600160a01b039091168152602001610156565b6101496104e4565b6101726102ae366004611066565b6104f3565b6101726102c1366004611066565b610539565b6101e86102d4366004611177565b610546565b6101866102e73660046111e4565b6001600160a01b039182165f90815260026020908152604080832093909416825291909152205490565b600954610280906001600160a01b031681565b6101e86103323660046110c8565b61067c565b60606004805461034690611215565b80601f016020809104026020016040519081016040528092919081815260200182805461037290611215565b80156103bd5780601f10610394576101008083540402835291602001916103bd565b820191905f5260205f20905b8154815290600101906020018083116103a057829003601f168201915b5050505050905090565b5f336103d48185856106b9565b60019150505b92915050565b5f336103ed8582856106cb565b6103f8858585610747565b506001949350505050565b5f61040c6107a4565b905090565b6104196108cd565b6104216108f9565b565b6009545f9033906001600160a01b03168114610463576040516338d2446d60e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b5061046e838361094d565b50600192915050565b61047f6108cd565b6104215f610985565b5f6103da826109d4565b61049a6108cd565b6104216109f1565b5f6060805f5f5f60606104b3610a33565b6104bb610a60565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606005805461034690611215565b6009545f9033906001600160a01b0316811461052e576040516338d2446d60e11b81526001600160a01b03909116600482015260240161045a565b5061046e8383610a8d565b5f336103d4818585610747565b8342111561056a5760405163313c898160e11b81526004810185905260240161045a565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105b58c6001600160a01b03165f90815260086020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61060f82610ac1565b90505f61061e82878787610aed565b9050896001600160a01b0316816001600160a01b031614610665576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161045a565b6106708a8a8a6106b9565b50505050505050505050565b6106846108cd565b6001600160a01b0381166106ad57604051631e4fbdf760e01b81525f600482015260240161045a565b6106b681610985565b50565b6106c68383836001610b19565b505050565b6001600160a01b038381165f908152600260209081526040808320938616835292905220545f19811015610741578181101561073357604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161045a565b61074184848484035f610b19565b50505050565b6001600160a01b03831661077057604051634b637e8f60e11b81525f600482015260240161045a565b6001600160a01b0382166107995760405163ec442f0560e01b81525f600482015260240161045a565b6106c6838383610beb565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156107fc57507f000000000000000000000000000000000000000000000000000000000000000046145b1561082657507f000000000000000000000000000000000000000000000000000000000000000090565b61040c604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f546001600160a01b031633146104215760405163118cdaa760e01b815233600482015260240161045a565b610901610bfe565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b0382166109765760405163ec442f0560e01b81525f600482015260240161045a565b6109815f8383610beb565b5050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381165f908152600860205260408120546103da565b6109f9610c27565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586109303390565b606061040c7f00000000000000000000000000000000000000000000000000000000000000006006610c51565b606061040c7f00000000000000000000000000000000000000000000000000000000000000006007610c51565b6001600160a01b038216610ab657604051634b637e8f60e11b81525f600482015260240161045a565b610981825f83610beb565b5f6103da610acd6107a4565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610afd88888888610cfa565b925092509250610b0d8282610dc2565b50909695505050505050565b6001600160a01b038416610b425760405163e602df0560e01b81525f600482015260240161045a565b6001600160a01b038316610b6b57604051634a1406b160e11b81525f600482015260240161045a565b6001600160a01b038085165f908152600260209081526040808320938716835292905220829055801561074157826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bdd91815260200190565b60405180910390a350505050565b610bf3610c27565b6106c6838383610e7a565b5f54600160a01b900460ff1661042157604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff16156104215760405163d93c066560e01b815260040160405180910390fd5b606060ff8314610c6b57610c6483610fa0565b90506103da565b818054610c7790611215565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca390611215565b8015610cee5780601f10610cc557610100808354040283529160200191610cee565b820191905f5260205f20905b815481529060010190602001808311610cd157829003601f168201915b505050505090506103da565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610d3357505f91506003905082610db8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610d84573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610daf57505f925060019150829050610db8565b92505f91508190505b9450945094915050565b5f826003811115610dd557610dd561124d565b03610dde575050565b6001826003811115610df257610df261124d565b03610e105760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610e2457610e2461124d565b03610e455760405163fce698f760e01b81526004810182905260240161045a565b6003826003811115610e5957610e5961124d565b03610981576040516335e2f38360e21b81526004810182905260240161045a565b6001600160a01b038316610ea4578060035f828254610e999190611261565b90915550610f149050565b6001600160a01b0383165f9081526001602052604090205481811015610ef65760405163391434e360e21b81526001600160a01b0385166004820152602481018290526044810183905260640161045a565b6001600160a01b0384165f9081526001602052604090209082900390555b6001600160a01b038216610f3057600380548290039055610f4e565b6001600160a01b0382165f9081526001602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f9391815260200190565b60405180910390a3505050565b60605f610fac83610fdd565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f8111156103da57604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6110446020830184611004565b9392505050565b80356001600160a01b0381168114611061575f5ffd5b919050565b5f5f60408385031215611077575f5ffd5b6110808361104b565b946020939093013593505050565b5f5f5f606084860312156110a0575f5ffd5b6110a98461104b565b92506110b76020850161104b565b929592945050506040919091013590565b5f602082840312156110d8575f5ffd5b6110448261104b565b60ff60f81b8816815260e060208201525f6110ff60e0830189611004565b82810360408401526111118189611004565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611166578351835260209384019390920191600101611148565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a03121561118d575f5ffd5b6111968861104b565b96506111a46020890161104b565b95506040880135945060608801359350608088013560ff811681146111c7575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156111f5575f5ffd5b6111fe8361104b565b915061120c6020840161104b565b90509250929050565b600181811c9082168061122957607f821691505b60208210810361124757634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b808201808211156103da57634e487b7160e01b5f52601160045260245ffdfea26469706673582212204500fb9c1f281ed08b6af7229bf86f9d987a88b7ac8404e69da0aca9d08ba62064736f6c634300081c0033",
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
func DeployCrossMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *CrossMintableERC20, error) {
	parsed, err := CrossMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossMintableERC20Bin), backend, _bridge, name_, symbol_, decimals_)
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
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferFrom(&_CrossMintableERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CrossMintableERC20 *CrossMintableERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossMintableERC20.Contract.TransferFrom(&_CrossMintableERC20.TransactOpts, from, to, value)
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
