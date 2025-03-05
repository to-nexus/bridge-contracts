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

// CrossMintableERC20FactoryMetaData contains all meta data concerning the CrossMintableERC20Factory contract.
var CrossMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createCrossMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ERC20FactoryNotBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ERC20FactoryNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"6ff97f1d": "allTokens()",
		"e78cea92": "bridge()",
		"4804a041": "createCrossMintableERC20(bytes32,string,string,uint8)",
		"4f6ccce7": "tokenByIndex(uint256)",
		"d92fc67b": "tokensLength()",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50604051611d44380380611d44833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b611cbe806100865f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80634804a041146100595780634f6ccce7146100895780636ff97f1d1461009c578063d92fc67b146100b1578063e78cea92146100c7575b5f5ffd5b61006c6100673660046103b6565b6100d9565b6040516001600160a01b0390911681526020015b60405180910390f35b61006c61009736600461043b565b6101a2565b6100a46101b4565b6040516100809190610452565b6100b96101c5565b604051908152602001610080565b5f5461006c906001600160a01b031681565b5f805433906001600160a01b0316811461011757604051633a33fb4d60e01b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f604051806020016101299061030a565b601f1982820381018352601f9091011660408190525f5461015e916001600160a01b03909116908890889088906020016104cb565b60408051601f198184030181529082905261017c929160200161052c565b60405160208183030381529060405290506101985f87836101d0565b9695505050505050565b5f6101ae60018361026b565b92915050565b60606101c06001610276565b905090565b5f6101c06001610282565b5f834710156101fb5760405163cf47918160e01b81524760048201526024810185905260440161010e565b81515f0361021c57604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d15198115161561023d576040513d5f823e3d81fd5b6001600160a01b0381166102645760405163b06ebf3d60e01b815260040160405180910390fd5b9392505050565b5f610264838361028b565b60605f610264836102b1565b5f6101ae825490565b5f825f0182815481106102a0576102a0610548565b905f5260205f200154905092915050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156102fe57602002820191905f5260205f20905b8154815260200190600101908083116102ea575b50505050509050919050565b61172c8061055d83390190565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261033a575f5ffd5b813567ffffffffffffffff81111561035457610354610317565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561038357610383610317565b60405281815283820160200185101561039a575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f608085870312156103c9575f5ffd5b84359350602085013567ffffffffffffffff8111156103e6575f5ffd5b6103f28782880161032b565b935050604085013567ffffffffffffffff81111561040e575f5ffd5b61041a8782880161032b565b925050606085013560ff81168114610430575f5ffd5b939692955090935050565b5f6020828403121561044b575f5ffd5b5035919050565b602080825282518282018190525f918401906040840190835b818110156104925783516001600160a01b031683526020938401939092019160010161046b565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03851681526080602082018190525f906104ee9083018661049d565b8281036040840152610500818661049d565b91505060ff8316606083015295945050505050565b5f81518060208401855e5f93019283525090919050565b5f61054061053a8386610515565b84610515565b949350505050565b634e487b7160e01b5f52603260045260245ffdfe610180604052348015610010575f5ffd5b5060405161172c38038061172c83398101604081905261002f916102ef565b8280604051806040016040528060018152602001603160f81b815250858561005b61019060201b60201c565b6001600160a01b03811661008957604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61009281610194565b50600461009f838261040d565b5060056100ac828261040d565b506100bc915083905060066101e3565b610120526100cb8160076101e3565b61014052815160208084019190912060e052815190820120610100524660a05261015760e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600980546001600160a01b0319166001600160a01b039590951694909417909355505060ff166101605261051f565b3390565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020835110156101fe576101f783610215565b905061020f565b81610209848261040d565b5060ff90505b92915050565b5f5f829050601f8151111561023f578260405163305a27a960e01b815260040161008091906104c7565b805161024a826104fc565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610275575f5ffd5b81516001600160401b0381111561028e5761028e610252565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102bc576102bc610252565b6040528181528382016020018510156102d3575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f60808587031215610302575f5ffd5b84516001600160a01b0381168114610318575f5ffd5b60208601519094506001600160401b03811115610333575f5ffd5b61033f87828801610266565b604087015190945090506001600160401b0381111561035c575f5ffd5b61036887828801610266565b925050606085015160ff8116811461037e575f5ffd5b939692955090935050565b600181811c9082168061039d57607f821691505b6020821081036103bb57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561040857805f5260205f20601f840160051c810160208510156103e65750805b601f840160051c820191505b81811015610405575f81556001016103f2565b50505b505050565b81516001600160401b0381111561042657610426610252565b61043a816104348454610389565b846103c1565b6020601f82116001811461046c575f83156104555750848201515b5f19600385901b1c1916600184901b178455610405565b5f84815260208120601f198516915b8281101561049b578785015182556020948501946001909201910161047b565b50848210156104b857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156103bb575f1960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051610160516111b261057a5f395f61018d01525f61096f01525f61094201525f61081b01525f6107f301525f61074e01525f61077801525f6107a201526111b25ff3fe608060405234801561000f575f5ffd5b506004361061011c575f3560e01c80637ecebe00116100a9578063a9059cbb1161006e578063a9059cbb14610271578063d505accf14610284578063dd62ed3e14610297578063e78cea92146102cf578063f2fde38b146102e2575f5ffd5b80637ecebe001461020457806384b0196e146102175780638da5cb5b1461023257806395d89b41146102565780639dc29fac1461025e575f5ffd5b8063313ce567116100ef578063313ce567146101865780633644e515146101b757806340c10f19146101bf57806370a08231146101d2578063715018a6146101fa575f5ffd5b806306fdde0314610120578063095ea7b31461013e57806318160ddd1461016157806323b872dd14610173575b5f5ffd5b6101286102f5565b6040516101359190610f2e565b60405180910390f35b61015161014c366004610f62565b610385565b6040519015158152602001610135565b6003545b604051908152602001610135565b610151610181366004610f8a565b61039e565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610135565b6101656103c1565b6101516101cd366004610f62565b6103cf565b6101656101e0366004610fc4565b6001600160a01b03165f9081526001602052604090205490565b610202610423565b005b610165610212366004610fc4565b610436565b61021f610440565b6040516101359796959493929190610fdd565b5f546001600160a01b03165b6040516001600160a01b039091168152602001610135565b610128610482565b61015161026c366004610f62565b610491565b61015161027f366004610f62565b6104d7565b610202610292366004611073565b6104e4565b6101656102a53660046110e0565b6001600160a01b039182165f90815260026020908152604080832093909416825291909152205490565b60095461023e906001600160a01b031681565b6102026102f0366004610fc4565b61061a565b60606004805461030490611111565b80601f016020809104026020016040519081016040528092919081815260200182805461033090611111565b801561037b5780601f106103525761010080835404028352916020019161037b565b820191905f5260205f20905b81548152906001019060200180831161035e57829003601f168201915b5050505050905090565b5f33610392818585610657565b60019150505b92915050565b5f336103ab858285610669565b6103b68585856106e5565b506001949350505050565b5f6103ca610742565b905090565b6009545f9033906001600160a01b0316811461040f57604051630b56edc360e21b81526001600160a01b0390911660048201526024015b60405180910390fd5b5061041a838361086b565b50600192915050565b61042b6108a3565b6104345f6108cf565b565b5f6103988261091e565b5f6060805f5f5f606061045161093b565b610459610968565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606005805461030490611111565b6009545f9033906001600160a01b031681146104cc57604051630b56edc360e21b81526001600160a01b039091166004820152602401610406565b5061041a8383610995565b5f336103928185856106e5565b834211156105085760405163313c898160e11b815260048101859052602401610406565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105538c6001600160a01b03165f90815260086020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6105ad826109c9565b90505f6105bc828787876109f5565b9050896001600160a01b0316816001600160a01b031614610603576040516325c0072360e11b81526001600160a01b0380831660048301528b166024820152604401610406565b61060e8a8a8a610657565b50505050505050505050565b6106226108a3565b6001600160a01b03811661064b57604051631e4fbdf760e01b81525f6004820152602401610406565b610654816108cf565b50565b6106648383836001610a21565b505050565b6001600160a01b038381165f908152600260209081526040808320938616835292905220545f198110156106df57818110156106d157604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610406565b6106df84848484035f610a21565b50505050565b6001600160a01b03831661070e57604051634b637e8f60e11b81525f6004820152602401610406565b6001600160a01b0382166107375760405163ec442f0560e01b81525f6004820152602401610406565b610664838383610af3565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561079a57507f000000000000000000000000000000000000000000000000000000000000000046145b156107c457507f000000000000000000000000000000000000000000000000000000000000000090565b6103ca604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b6001600160a01b0382166108945760405163ec442f0560e01b81525f6004820152602401610406565b61089f5f8383610af3565b5050565b5f546001600160a01b031633146104345760405163118cdaa760e01b8152336004820152602401610406565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381165f90815260086020526040812054610398565b60606103ca7f00000000000000000000000000000000000000000000000000000000000000006006610b4d565b60606103ca7f00000000000000000000000000000000000000000000000000000000000000006007610b4d565b6001600160a01b0382166109be57604051634b637e8f60e11b81525f6004820152602401610406565b61089f825f83610af3565b5f6103986109d5610742565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610a0588888888610bf6565b925092509250610a158282610cbe565b50909695505050505050565b6001600160a01b038416610a4a5760405163e602df0560e01b81525f6004820152602401610406565b6001600160a01b038316610a7357604051634a1406b160e11b81525f6004820152602401610406565b6001600160a01b038085165f90815260026020908152604080832093871683529290522082905580156106df57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610ae591815260200190565b60405180910390a350505050565b6009546001600160a01b0316336001600160a01b031614158015610b2457506009546001600160a01b038381169116145b15610b425760405163246a105b60e21b815260040160405180910390fd5b610664838383610d76565b606060ff8314610b6757610b6083610e9c565b9050610398565b818054610b7390611111565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9f90611111565b8015610bea5780601f10610bc157610100808354040283529160200191610bea565b820191905f5260205f20905b815481529060010190602001808311610bcd57829003601f168201915b50505050509050610398565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610c2f57505f91506003905082610cb4565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610c80573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610cab57505f925060019150829050610cb4565b92505f91508190505b9450945094915050565b5f826003811115610cd157610cd1611149565b03610cda575050565b6001826003811115610cee57610cee611149565b03610d0c5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610d2057610d20611149565b03610d415760405163fce698f760e01b815260048101829052602401610406565b6003826003811115610d5557610d55611149565b0361089f576040516335e2f38360e21b815260048101829052602401610406565b6001600160a01b038316610da0578060035f828254610d95919061115d565b90915550610e109050565b6001600160a01b0383165f9081526001602052604090205481811015610df25760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610406565b6001600160a01b0384165f9081526001602052604090209082900390555b6001600160a01b038216610e2c57600380548290039055610e4a565b6001600160a01b0382165f9081526001602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e8f91815260200190565b60405180910390a3505050565b60605f610ea883610ed9565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f81111561039857604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610f406020830184610f00565b9392505050565b80356001600160a01b0381168114610f5d575f5ffd5b919050565b5f5f60408385031215610f73575f5ffd5b610f7c83610f47565b946020939093013593505050565b5f5f5f60608486031215610f9c575f5ffd5b610fa584610f47565b9250610fb360208501610f47565b929592945050506040919091013590565b5f60208284031215610fd4575f5ffd5b610f4082610f47565b60ff60f81b8816815260e060208201525f610ffb60e0830189610f00565b828103604084015261100d8189610f00565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611062578351835260209384019390920191600101611044565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a031215611089575f5ffd5b61109288610f47565b96506110a060208901610f47565b95506040880135945060608801359350608088013560ff811681146110c3575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156110f1575f5ffd5b6110fa83610f47565b915061110860208401610f47565b90509250929050565b600181811c9082168061112557607f821691505b60208210810361114357634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b8082018082111561039857634e487b7160e01b5f52601160045260245ffdfea2646970667358221220e4cb566be26bffbad53c11af2837a9447d3f4b9c82491a0f6a551883c867b97e64736f6c634300081c0033a26469706673582212204f6a481b735ac690465b0b4bc2215f0ff1d34c4de3d1b4101cf82ac4249984d964736f6c634300081c0033",
}

// CrossMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossMintableERC20FactoryMetaData.ABI instead.
var CrossMintableERC20FactoryABI = CrossMintableERC20FactoryMetaData.ABI

// Deprecated: Use CrossMintableERC20FactoryMetaData.Sigs instead.
// CrossMintableERC20FactoryFuncSigs maps the 4-byte function signature to its string representation.
var CrossMintableERC20FactoryFuncSigs = CrossMintableERC20FactoryMetaData.Sigs

// CrossMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossMintableERC20FactoryMetaData.Bin instead.
var CrossMintableERC20FactoryBin = CrossMintableERC20FactoryMetaData.Bin

// DeployCrossMintableERC20Factory deploys a new Ethereum contract, binding an instance of CrossMintableERC20Factory to it.
func DeployCrossMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *CrossMintableERC20Factory, error) {
	parsed, err := CrossMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossMintableERC20Factory{CrossMintableERC20FactoryCaller: CrossMintableERC20FactoryCaller{contract: contract}, CrossMintableERC20FactoryTransactor: CrossMintableERC20FactoryTransactor{contract: contract}, CrossMintableERC20FactoryFilterer: CrossMintableERC20FactoryFilterer{contract: contract}}, nil
}

// CrossMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type CrossMintableERC20Factory struct {
	CrossMintableERC20FactoryCaller     // Read-only binding to the contract
	CrossMintableERC20FactoryTransactor // Write-only binding to the contract
	CrossMintableERC20FactoryFilterer   // Log filterer for contract events
}

// CrossMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossMintableERC20FactorySession struct {
	Contract     *CrossMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CrossMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossMintableERC20FactoryCallerSession struct {
	Contract *CrossMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CrossMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossMintableERC20FactoryTransactorSession struct {
	Contract     *CrossMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CrossMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossMintableERC20FactoryRaw struct {
	Contract *CrossMintableERC20Factory // Generic contract binding to access the raw methods on
}

// CrossMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossMintableERC20FactoryCallerRaw struct {
	Contract *CrossMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CrossMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossMintableERC20FactoryTransactorRaw struct {
	Contract *CrossMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossMintableERC20Factory creates a new instance of CrossMintableERC20Factory, bound to a specific deployed contract.
func NewCrossMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*CrossMintableERC20Factory, error) {
	contract, err := bindCrossMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20Factory{CrossMintableERC20FactoryCaller: CrossMintableERC20FactoryCaller{contract: contract}, CrossMintableERC20FactoryTransactor: CrossMintableERC20FactoryTransactor{contract: contract}, CrossMintableERC20FactoryFilterer: CrossMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewCrossMintableERC20FactoryCaller creates a new read-only instance of CrossMintableERC20Factory, bound to a specific deployed contract.
func NewCrossMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*CrossMintableERC20FactoryCaller, error) {
	contract, err := bindCrossMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20FactoryCaller{contract: contract}, nil
}

// NewCrossMintableERC20FactoryTransactor creates a new write-only instance of CrossMintableERC20Factory, bound to a specific deployed contract.
func NewCrossMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossMintableERC20FactoryTransactor, error) {
	contract, err := bindCrossMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewCrossMintableERC20FactoryFilterer creates a new log filterer instance of CrossMintableERC20Factory, bound to a specific deployed contract.
func NewCrossMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossMintableERC20FactoryFilterer, error) {
	contract, err := bindCrossMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindCrossMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindCrossMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20Factory.Contract.CrossMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.CrossMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.CrossMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[])
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCaller) AllTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20Factory.contract.Call(opts, &out, "allTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[])
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) AllTokens() ([]common.Address, error) {
	return _CrossMintableERC20Factory.Contract.AllTokens(&_CrossMintableERC20Factory.CallOpts)
}

// AllTokens is a free data retrieval call binding the contract method 0x6ff97f1d.
//
// Solidity: function allTokens() view returns(address[])
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCallerSession) AllTokens() ([]common.Address, error) {
	return _CrossMintableERC20Factory.Contract.AllTokens(&_CrossMintableERC20Factory.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20Factory.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) Bridge() (common.Address, error) {
	return _CrossMintableERC20Factory.Contract.Bridge(&_CrossMintableERC20Factory.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCallerSession) Bridge() (common.Address, error) {
	return _CrossMintableERC20Factory.Contract.Bridge(&_CrossMintableERC20Factory.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossMintableERC20Factory.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) TokenByIndex(index *big.Int) (common.Address, error) {
	return _CrossMintableERC20Factory.Contract.TokenByIndex(&_CrossMintableERC20Factory.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(address)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCallerSession) TokenByIndex(index *big.Int) (common.Address, error) {
	return _CrossMintableERC20Factory.Contract.TokenByIndex(&_CrossMintableERC20Factory.CallOpts, index)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossMintableERC20Factory.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) TokensLength() (*big.Int, error) {
	return _CrossMintableERC20Factory.Contract.TokensLength(&_CrossMintableERC20Factory.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryCallerSession) TokensLength() (*big.Int, error) {
	return _CrossMintableERC20Factory.Contract.TokensLength(&_CrossMintableERC20Factory.CallOpts)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0x4804a041.
//
// Solidity: function createCrossMintableERC20(bytes32 salt, string name, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactor) CreateCrossMintableERC20(opts *bind.TransactOpts, salt [32]byte, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.contract.Transact(opts, "createCrossMintableERC20", salt, name, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0x4804a041.
//
// Solidity: function createCrossMintableERC20(bytes32 salt, string name, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) CreateCrossMintableERC20(salt [32]byte, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.CreateCrossMintableERC20(&_CrossMintableERC20Factory.TransactOpts, salt, name, symbol, decimals)
}

// CreateCrossMintableERC20 is a paid mutator transaction binding the contract method 0x4804a041.
//
// Solidity: function createCrossMintableERC20(bytes32 salt, string name, string symbol, uint8 decimals) returns(address tokenAddress)
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactorSession) CreateCrossMintableERC20(salt [32]byte, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.CreateCrossMintableERC20(&_CrossMintableERC20Factory.TransactOpts, salt, name, symbol, decimals)
}
