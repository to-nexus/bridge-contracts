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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20CanNotTransferToBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20NotBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"}]",
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
		"d505accf": "permit(address,address,uint256,uint256,uint8,bytes32,bytes32)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x6101a0604052348015610010575f5ffd5b5060405161162338038061162383398101604081905261002f9161024e565b6040805180820190915260018152603160f81b60208201528390819081856003610059838261036c565b506004610066828261036c565b5061007691508390506005610139565b61012052610085816006610139565b61014052815160208084019190912060e052815190820120610100524660a05261011160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052506001600160a01b039390931661018052505060ff166101605261047e565b5f6020835110156101545761014d8361016b565b9050610165565b8161015f848261036c565b5060ff90505b92915050565b5f5f829050601f8151111561019e578260405163305a27a960e01b81526004016101959190610426565b60405180910390fd5b80516101a98261045b565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126101d4575f5ffd5b81516001600160401b038111156101ed576101ed6101b1565b604051601f8201601f19908116603f011681016001600160401b038111828210171561021b5761021b6101b1565b604052818152838201602001851015610232575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f60808587031215610261575f5ffd5b84516001600160a01b0381168114610277575f5ffd5b60208601519094506001600160401b03811115610292575f5ffd5b61029e878288016101c5565b604087015190945090506001600160401b038111156102bb575f5ffd5b6102c7878288016101c5565b925050606085015160ff811681146102dd575f5ffd5b939692955090935050565b600181811c908216806102fc57607f821691505b60208210810361031a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561036757805f5260205f20601f840160051c810160208510156103455750805b601f840160051c820191505b81811015610364575f8155600101610351565b50505b505050565b81516001600160401b03811115610385576103856101b1565b6103998161039384546102e8565b84610320565b6020601f8211600181146103cb575f83156103b45750848201515b5f19600385901b1c1916600184901b178455610364565b5f84815260208120601f198516915b828110156103fa57878501518255602094850194600190920191016103da565b508482101561041757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561031a575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516111246104ff5f395f81816102870152818161039f0152818161046a01528181610a2f0152610a6201525f61016c01525f6108a801525f61087b01525f6107cf01525f6107a701525f61070201525f61072c01525f61075601526111245ff3fe608060405234801561000f575f5ffd5b50600436106100fb575f3560e01c80637ecebe0011610093578063a9059cbb11610063578063a9059cbb14610222578063d505accf14610235578063dd62ed3e1461024a578063e78cea9214610282575f5ffd5b80637ecebe00146101d957806384b0196e146101ec57806395d89b41146102075780639dc29fac1461020f575f5ffd5b8063313ce567116100ce578063313ce567146101655780633644e5151461019657806340c10f191461019e57806370a08231146101b1575f5ffd5b806306fdde03146100ff578063095ea7b31461011d57806318160ddd1461014057806323b872dd14610152575b5f5ffd5b6101076102c1565b6040516101149190610ea0565b60405180910390f35b61013061012b366004610ed4565b610351565b6040519015158152602001610114565b6002545b604051908152602001610114565b610130610160366004610efc565b61036a565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610114565b61014461038d565b6101306101ac366004610ed4565b61039b565b6101446101bf366004610f36565b6001600160a01b03165f9081526020819052604090205490565b6101446101e7366004610f36565b61040b565b6101f4610415565b6040516101149796959493929190610f4f565b610107610457565b61013061021d366004610ed4565b610466565b610130610230366004610ed4565b6104c8565b610248610243366004610fe5565b6104d5565b005b610144610258366004611052565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610114565b6060600380546102d090611083565b80601f01602080910402602001604051908101604052809291908181526020018280546102fc90611083565b80156103475780601f1061031e57610100808354040283529160200191610347565b820191905f5260205f20905b81548152906001019060200180831161032a57829003601f168201915b5050505050905090565b5f3361035e81858561060b565b60019150505b92915050565b5f3361037785828561061d565b610382858585610699565b506001949350505050565b5f6103966106f6565b905090565b5f337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681146103f757604051630b56edc360e21b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610402838361081f565b50600192915050565b5f61036482610857565b5f6060805f5f5f6060610426610874565b61042e6108a1565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6060600480546102d090611083565b5f337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681146104bd57604051630b56edc360e21b81526001600160a01b0390911660048201526024016103ee565b5061040283836108ce565b5f3361035e818585610699565b834211156104f95760405163313c898160e11b8152600481018590526024016103ee565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105448c6001600160a01b03165f90815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61059e82610902565b90505f6105ad8287878761092e565b9050896001600160a01b0316816001600160a01b0316146105f4576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016103ee565b6105ff8a8a8a61060b565b50505050505050505050565b610618838383600161095a565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610693578181101561068557604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103ee565b61069384848484035f61095a565b50505050565b6001600160a01b0383166106c257604051634b637e8f60e11b81525f60048201526024016103ee565b6001600160a01b0382166106eb5760405163ec442f0560e01b81525f60048201526024016103ee565b610618838383610a2c565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561074e57507f000000000000000000000000000000000000000000000000000000000000000046145b1561077857507f000000000000000000000000000000000000000000000000000000000000000090565b610396604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b6001600160a01b0382166108485760405163ec442f0560e01b81525f60048201526024016103ee565b6108535f8383610a2c565b5050565b6001600160a01b0381165f90815260076020526040812054610364565b60606103967f00000000000000000000000000000000000000000000000000000000000000006005610abf565b60606103967f00000000000000000000000000000000000000000000000000000000000000006006610abf565b6001600160a01b0382166108f757604051634b637e8f60e11b81525f60048201526024016103ee565b610853825f83610a2c565b5f61036461090e6106f6565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f61093e88888888610b68565b92509250925061094e8282610c30565b50909695505050505050565b6001600160a01b0384166109835760405163e602df0560e01b81525f60048201526024016103ee565b6001600160a01b0383166109ac57604051634a1406b160e11b81525f60048201526024016103ee565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561069357826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610a1e91815260200190565b60405180910390a350505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614801590610a9657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316145b15610ab45760405163246a105b60e21b815260040160405180910390fd5b610618838383610ce8565b606060ff8314610ad957610ad283610e0e565b9050610364565b818054610ae590611083565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1190611083565b8015610b5c5780601f10610b3357610100808354040283529160200191610b5c565b820191905f5260205f20905b815481529060010190602001808311610b3f57829003601f168201915b50505050509050610364565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610ba157505f91506003905082610c26565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610bf2573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610c1d57505f925060019150829050610c26565b92505f91508190505b9450945094915050565b5f826003811115610c4357610c436110bb565b03610c4c575050565b6001826003811115610c6057610c606110bb565b03610c7e5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610c9257610c926110bb565b03610cb35760405163fce698f760e01b8152600481018290526024016103ee565b6003826003811115610cc757610cc76110bb565b03610853576040516335e2f38360e21b8152600481018290526024016103ee565b6001600160a01b038316610d12578060025f828254610d0791906110cf565b90915550610d829050565b6001600160a01b0383165f9081526020819052604090205481811015610d645760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103ee565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610d9e57600280548290039055610dbc565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e0191815260200190565b60405180910390a3505050565b60605f610e1a83610e4b565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561036457604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610eb26020830184610e72565b9392505050565b80356001600160a01b0381168114610ecf575f5ffd5b919050565b5f5f60408385031215610ee5575f5ffd5b610eee83610eb9565b946020939093013593505050565b5f5f5f60608486031215610f0e575f5ffd5b610f1784610eb9565b9250610f2560208501610eb9565b929592945050506040919091013590565b5f60208284031215610f46575f5ffd5b610eb282610eb9565b60ff60f81b8816815260e060208201525f610f6d60e0830189610e72565b8281036040840152610f7f8189610e72565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015610fd4578351835260209384019390920191600101610fb6565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215610ffb575f5ffd5b61100488610eb9565b965061101260208901610eb9565b95506040880135945060608801359350608088013560ff81168114611035575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215611063575f5ffd5b61106c83610eb9565b915061107a60208401610eb9565b90509250929050565b600181811c9082168061109757607f821691505b6020821081036110b557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b8082018082111561036457634e487b7160e01b5f52601160045260245ffdfea26469706673582212202837ffe4daaf9d2f9434330b2daaa1d7c897f853cb9511460859c269946ef2e564736f6c634300081c0033",
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
