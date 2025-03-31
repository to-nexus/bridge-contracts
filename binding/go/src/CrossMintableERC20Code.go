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

// CrossMintableERC20CodeMetaData contains all meta data concerning the CrossMintableERC20Code contract.
var CrossMintableERC20CodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createCrossMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20CodeNotBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e78cea92": "bridge()",
		"f88d3d42": "createCrossMintableERC20(uint256,address,string,uint8)",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50604051611b38380380611b38833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b611ab2806100865f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063e78cea9214610038578063f88d3d4214610066575b5f5ffd5b5f5461004a906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b61004a610074366004610291565b5f805433906001600160a01b031681146100b25760405163430ad06360e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f85856040516020016100e292919091825260601b6bffffffffffffffffffffffff1916602082015260340190565b6040516020818303038152906040528051906020012090505f6040518060200161010b9061025b565b601f1982820381018352601f9091011660408190525f546001600160a01b03169061013a908890602001610392565b6040516020818303038152906040528760405160200161015a91906103b3565b60408051601f198184030181529082905261017b93929189906020016103fd565b60408051601f19818403018152908290526101999291602001610447565b60405160208183030381529060405290506101b55f83836101c0565b979650505050505050565b5f834710156101eb5760405163cf47918160e01b8152476004820152602481018590526044016100a9565b81515f0361020c57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d15198115161561022d576040513d5f823e3d81fd5b6001600160a01b0381166102545760405163b06ebf3d60e01b815260040160405180910390fd5b9392505050565b6116198061046483390190565b634e487b7160e01b5f52604160045260245ffd5b803560ff8116811461028c575f5ffd5b919050565b5f5f5f5f608085870312156102a4575f5ffd5b8435935060208501356001600160a01b03811681146102c1575f5ffd5b9250604085013567ffffffffffffffff8111156102dc575f5ffd5b8501601f810187136102ec575f5ffd5b803567ffffffffffffffff81111561030657610306610268565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561033557610335610268565b60405281815282820160200189101561034c575f5ffd5b816020840160208301375f602083830101528094505050506103706060860161027c565b905092959194509250565b5f81518060208401855e5f93019283525090919050565b6c021b937b9b990213934b233b29609d1b81525f610254600d83018461037b565b5f6103be828461037b565b600f60fb1b81526001019392505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03851681526080602082018190525f90610420908301866103cf565b828103604084015261043281866103cf565b91505060ff8316606083015295945050505050565b5f61045b610455838661037b565b8461037b565b94935050505056fe6101a0604052348015610010575f5ffd5b5060405161161938038061161983398101604081905261002f9161024e565b6040805180820190915260018152603160f81b60208201528390819081856003610059838261036c565b506004610066828261036c565b5061007691508390506005610139565b61012052610085816006610139565b61014052815160208084019190912060e052815190820120610100524660a05261011160e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052506001600160a01b039390931661018052505060ff166101605261047e565b5f6020835110156101545761014d8361016b565b9050610165565b8161015f848261036c565b5060ff90505b92915050565b5f5f829050601f8151111561019e578260405163305a27a960e01b81526004016101959190610426565b60405180910390fd5b80516101a98261045b565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126101d4575f5ffd5b81516001600160401b038111156101ed576101ed6101b1565b604051601f8201601f19908116603f011681016001600160401b038111828210171561021b5761021b6101b1565b604052818152838201602001851015610232575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f60808587031215610261575f5ffd5b84516001600160a01b0381168114610277575f5ffd5b60208601519094506001600160401b03811115610292575f5ffd5b61029e878288016101c5565b604087015190945090506001600160401b038111156102bb575f5ffd5b6102c7878288016101c5565b925050606085015160ff811681146102dd575f5ffd5b939692955090935050565b600181811c908216806102fc57607f821691505b60208210810361031a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561036757805f5260205f20601f840160051c810160208510156103455750805b601f840160051c820191505b81811015610364575f8155600101610351565b50505b505050565b81516001600160401b03811115610385576103856101b1565b6103998161039384546102e8565b84610320565b6020601f8211600181146103cb575f83156103b45750848201515b5f19600385901b1c1916600184901b178455610364565b5f84815260208120601f198516915b828110156103fa57878501518255602094850194600190920191016103da565b508482101561041757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561031a575f1960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051610160516101805161111a6104ff5f395f81816102870152818161039f0152818161046a01528181610a2f0152610a6201525f61016c01525f6108a801525f61087b01525f6107cf01525f6107a701525f61070201525f61072c01525f610756015261111a5ff3fe608060405234801561000f575f5ffd5b50600436106100fb575f3560e01c80637ecebe0011610093578063a9059cbb11610063578063a9059cbb14610222578063d505accf14610235578063dd62ed3e1461024a578063e78cea9214610282575f5ffd5b80637ecebe00146101d957806384b0196e146101ec57806395d89b41146102075780639dc29fac1461020f575f5ffd5b8063313ce567116100ce578063313ce567146101655780633644e5151461019657806340c10f191461019e57806370a08231146101b1575f5ffd5b806306fdde03146100ff578063095ea7b31461011d57806318160ddd1461014057806323b872dd14610152575b5f5ffd5b6101076102c1565b6040516101149190610e96565b60405180910390f35b61013061012b366004610eca565b610351565b6040519015158152602001610114565b6002545b604051908152602001610114565b610130610160366004610ef2565b61036a565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610114565b61014461038d565b6101306101ac366004610eca565b61039b565b6101446101bf366004610f2c565b6001600160a01b03165f9081526020819052604090205490565b6101446101e7366004610f2c565b61040b565b6101f4610415565b6040516101149796959493929190610f45565b610107610457565b61013061021d366004610eca565b610466565b610130610230366004610eca565b6104c8565b610248610243366004610fdb565b6104d5565b005b610144610258366004611048565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6102a97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610114565b6060600380546102d090611079565b80601f01602080910402602001604051908101604052809291908181526020018280546102fc90611079565b80156103475780601f1061031e57610100808354040283529160200191610347565b820191905f5260205f20905b81548152906001019060200180831161032a57829003601f168201915b5050505050905090565b5f3361035e81858561060b565b60019150505b92915050565b5f3361037785828561061d565b610382858585610699565b506001949350505050565b5f6103966106f6565b905090565b5f337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681146103f757604051630b56edc360e21b81526001600160a01b0390911660048201526024015b60405180910390fd5b50610402838361081f565b50600192915050565b5f61036482610857565b5f6060805f5f5f6060610426610874565b61042e6108a1565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6060600480546102d090611079565b5f337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681146104bd57604051630b56edc360e21b81526001600160a01b0390911660048201526024016103ee565b5061040283836108ce565b5f3361035e818585610699565b834211156104f95760405163313c898160e11b8152600481018590526024016103ee565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105448c6001600160a01b03165f90815260076020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61059e82610902565b90505f6105ad8287878761092e565b9050896001600160a01b0316816001600160a01b0316146105f4576040516325c0072360e11b81526001600160a01b0380831660048301528b1660248201526044016103ee565b6105ff8a8a8a61060b565b50505050505050505050565b610618838383600161095a565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811015610693578181101561068557604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103ee565b61069384848484035f61095a565b50505050565b6001600160a01b0383166106c257604051634b637e8f60e11b81525f60048201526024016103ee565b6001600160a01b0382166106eb5760405163ec442f0560e01b81525f60048201526024016103ee565b610618838383610a2c565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561074e57507f000000000000000000000000000000000000000000000000000000000000000046145b1561077857507f000000000000000000000000000000000000000000000000000000000000000090565b610396604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b6001600160a01b0382166108485760405163ec442f0560e01b81525f60048201526024016103ee565b6108535f8383610a2c565b5050565b6001600160a01b0381165f90815260076020526040812054610364565b60606103967f00000000000000000000000000000000000000000000000000000000000000006005610abf565b60606103967f00000000000000000000000000000000000000000000000000000000000000006006610abf565b6001600160a01b0382166108f757604051634b637e8f60e11b81525f60048201526024016103ee565b610853825f83610a2c565b5f61036461090e6106f6565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f61093e88888888610b68565b92509250925061094e8282610c26565b50909695505050505050565b6001600160a01b0384166109835760405163e602df0560e01b81525f60048201526024016103ee565b6001600160a01b0383166109ac57604051634a1406b160e11b81525f60048201526024016103ee565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561069357826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610a1e91815260200190565b60405180910390a350505050565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614801590610a9657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316145b15610ab45760405163246a105b60e21b815260040160405180910390fd5b610618838383610cde565b606060ff8314610ad957610ad283610e04565b9050610364565b818054610ae590611079565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1190611079565b8015610b5c5780601f10610b3357610100808354040283529160200191610b5c565b820191905f5260205f20905b815481529060010190602001808311610b3f57829003601f168201915b50505050509050610364565b5f80806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03841115610b9757505f91506003905082610c1c565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610be8573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610c1357505f925060019150829050610c1c565b92505f91508190505b9450945094915050565b5f826003811115610c3957610c396110b1565b03610c42575050565b6001826003811115610c5657610c566110b1565b03610c745760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610c8857610c886110b1565b03610ca95760405163fce698f760e01b8152600481018290526024016103ee565b6003826003811115610cbd57610cbd6110b1565b03610853576040516335e2f38360e21b8152600481018290526024016103ee565b6001600160a01b038316610d08578060025f828254610cfd91906110c5565b90915550610d789050565b6001600160a01b0383165f9081526020819052604090205481811015610d5a5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103ee565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610d9457600280548290039055610db2565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610df791815260200190565b60405180910390a3505050565b60605f610e1083610e41565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561036457604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610ea86020830184610e68565b9392505050565b80356001600160a01b0381168114610ec5575f5ffd5b919050565b5f5f60408385031215610edb575f5ffd5b610ee483610eaf565b946020939093013593505050565b5f5f5f60608486031215610f04575f5ffd5b610f0d84610eaf565b9250610f1b60208501610eaf565b929592945050506040919091013590565b5f60208284031215610f3c575f5ffd5b610ea882610eaf565b60ff60f81b8816815260e060208201525f610f6360e0830189610e68565b8281036040840152610f758189610e68565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015610fca578351835260209384019390920191600101610fac565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215610ff1575f5ffd5b610ffa88610eaf565b965061100860208901610eaf565b95506040880135945060608801359350608088013560ff8116811461102b575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215611059575f5ffd5b61106283610eaf565b915061107060208401610eaf565b90509250929050565b600181811c9082168061108d57607f821691505b6020821081036110ab57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b8082018082111561036457634e487b7160e01b5f52601160045260245ffdfea26469706673582212209cf99f6454b2f540875908e736a606fa80d46b057a2f0e8d8be234ea9885c5ae64736f6c634300081c0033a26469706673582212202fb623525d8cc5fa809af80d7c6bad492738f4c37ef6b4a1310a81447ee09fa464736f6c634300081c0033",
}

// CrossMintableERC20CodeABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossMintableERC20CodeMetaData.ABI instead.
var CrossMintableERC20CodeABI = CrossMintableERC20CodeMetaData.ABI

// Deprecated: Use CrossMintableERC20CodeMetaData.Sigs instead.
// CrossMintableERC20CodeFuncSigs maps the 4-byte function signature to its string representation.
var CrossMintableERC20CodeFuncSigs = CrossMintableERC20CodeMetaData.Sigs

// CrossMintableERC20CodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossMintableERC20CodeMetaData.Bin instead.
var CrossMintableERC20CodeBin = CrossMintableERC20CodeMetaData.Bin

// DeployCrossMintableERC20Code deploys a new Ethereum contract, binding an instance of CrossMintableERC20Code to it.
func DeployCrossMintableERC20Code(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *CrossMintableERC20Code, error) {
	parsed, err := CrossMintableERC20CodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossMintableERC20CodeBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossMintableERC20Code{CrossMintableERC20CodeCaller: CrossMintableERC20CodeCaller{contract: contract}, CrossMintableERC20CodeTransactor: CrossMintableERC20CodeTransactor{contract: contract}, CrossMintableERC20CodeFilterer: CrossMintableERC20CodeFilterer{contract: contract}}, nil
}

// CrossMintableERC20Code is an auto generated Go binding around an Ethereum contract.
type CrossMintableERC20Code struct {
	CrossMintableERC20CodeCaller     // Read-only binding to the contract
	CrossMintableERC20CodeTransactor // Write-only binding to the contract
	CrossMintableERC20CodeFilterer   // Log filterer for contract events
}

// CrossMintableERC20CodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossMintableERC20CodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20CodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossMintableERC20CodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20CodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossMintableERC20CodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20CodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossMintableERC20CodeSession struct {
	Contract     *CrossMintableERC20Code // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrossMintableERC20CodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossMintableERC20CodeCallerSession struct {
	Contract *CrossMintableERC20CodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CrossMintableERC20CodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossMintableERC20CodeTransactorSession struct {
	Contract     *CrossMintableERC20CodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CrossMintableERC20CodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossMintableERC20CodeRaw struct {
	Contract *CrossMintableERC20Code // Generic contract binding to access the raw methods on
}

// CrossMintableERC20CodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossMintableERC20CodeCallerRaw struct {
	Contract *CrossMintableERC20CodeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossMintableERC20CodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossMintableERC20CodeTransactorRaw struct {
	Contract *CrossMintableERC20CodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossMintableERC20Code creates a new instance of CrossMintableERC20Code, bound to a specific deployed contract.
func NewCrossMintableERC20Code(address common.Address, backend bind.ContractBackend) (*CrossMintableERC20Code, error) {
	contract, err := bindCrossMintableERC20Code(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20Code{CrossMintableERC20CodeCaller: CrossMintableERC20CodeCaller{contract: contract}, CrossMintableERC20CodeTransactor: CrossMintableERC20CodeTransactor{contract: contract}, CrossMintableERC20CodeFilterer: CrossMintableERC20CodeFilterer{contract: contract}}, nil
}

// NewCrossMintableERC20CodeCaller creates a new read-only instance of CrossMintableERC20Code, bound to a specific deployed contract.
func NewCrossMintableERC20CodeCaller(address common.Address, caller bind.ContractCaller) (*CrossMintableERC20CodeCaller, error) {
	contract, err := bindCrossMintableERC20Code(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20CodeCaller{contract: contract}, nil
}

// NewCrossMintableERC20CodeTransactor creates a new write-only instance of CrossMintableERC20Code, bound to a specific deployed contract.
func NewCrossMintableERC20CodeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossMintableERC20CodeTransactor, error) {
	contract, err := bindCrossMintableERC20Code(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20CodeTransactor{contract: contract}, nil
}

// NewCrossMintableERC20CodeFilterer creates a new log filterer instance of CrossMintableERC20Code, bound to a specific deployed contract.
func NewCrossMintableERC20CodeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossMintableERC20CodeFilterer, error) {
	contract, err := bindCrossMintableERC20Code(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20CodeFilterer{contract: contract}, nil
}

// bindCrossMintableERC20Code binds a generic wrapper to an already deployed contract.
func bindCrossMintableERC20Code(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossMintableERC20CodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20Code *CrossMintableERC20CodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20Code.Contract.CrossMintableERC20CodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20Code *CrossMintableERC20CodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.CrossMintableERC20CodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20Code *CrossMintableERC20CodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.CrossMintableERC20CodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20Code *CrossMintableERC20CodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20Code.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20Code *CrossMintableERC20CodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20Code *CrossMintableERC20CodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Code *CrossMintableERC20CodeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20Code.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Code *CrossMintableERC20CodeSession) Bridge() (common.Address, error) {
	return _CrossMintableERC20Code.Contract.Bridge(&_CrossMintableERC20Code.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Code *CrossMintableERC20CodeCallerSession) Bridge() (common.Address, error) {
	return _CrossMintableERC20Code.Contract.Bridge(&_CrossMintableERC20Code.CallOpts)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Code *CrossMintableERC20CodeTransactor) CreateCrossMintableERC20(opts *bind.TransactOpts, remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Code.contract.Transact(opts, "createCrossMintableERC20", remoteChainID, remoteToken, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Code *CrossMintableERC20CodeSession) CreateCrossMintableERC20(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.CreateCrossMintableERC20(&_CrossMintableERC20Code.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0xf88d3d42.
//
// Solidity: function createCrossMintableERC20(uint256 remoteChainID, address remoteToken, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Code *CrossMintableERC20CodeTransactorSession) CreateCrossMintableERC20(remoteChainID *big.Int, remoteToken common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Code.Contract.CreateCrossMintableERC20(&_CrossMintableERC20Code.TransactOpts, remoteChainID, remoteToken, symbol, decimals)
}
