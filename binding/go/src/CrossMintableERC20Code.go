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
	Bin: "0x608034606f57601f611a3c38819003918201601f19168301916001600160401b03831184841017607357808492602094604052833981010312606f57516001600160a01b03811690819003606f575f80546001600160a01b0319169190911790556040516119b490816100888239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c908163e78cea92146102a6575063f88d3d4214610032575f80fd5b3461028e57608036600319011261028e576024356001600160a01b038116810361028e576044356001600160401b03811161028e573660238201121561028e576004810135916001600160401b038311610292576040519261009e601f8201601f1916602001856102c8565b8084526020840192366024838301011161028e57815f9260246020930186378501015260643560ff811680910361028e575f546001600160a01b0316903382900361027b576020610223926101f26101cc9560405184810191600435835260018060601b03199060601b1660408201526034815261011d6054826102c8565b5190209761166f946101de604051996101388789018c6102c8565b878b52868b019761031089396101af600188604051966c021b937b9b990213934b233b29609d1b82890152610187602d89835180898484015e81015f838201520301601f1981018a52896102c8565b6040519485915180918484015e8101600f60fb1b838201520301601e198101845201826102c8565b604051998a948886019788526080604087015260a08601906102eb565b848103601f19016060860152906102eb565b90608083015203601f1981018752866102c8565b60405194859383850197518091895e840190838201905f8252519283915e01015f815203601f1981018352826102c8565b80511561026c5751905ff53d1519811516610261576001600160a01b0316801561025257602090604051908152f35b63b06ebf3d60e01b5f5260045ffd5b6040513d5f823e3d90fd5b631328927760e21b5f5260045ffd5b63430ad06360e11b5f523360045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b3461028e575f36600319011261028e575f546001600160a01b03168152602090f35b601f909101601f19168101906001600160401b0382119082101761029257604052565b805180835260209291819084018484015e5f828201840152601f01601f191601019056fe6101a080604052346104675761166f803803809161001d828561046b565b833981016080828203126104675781516001600160a01b03811681036104675760208301516001600160401b038111610467578261005c91850161048e565b60408401519092906001600160401b0381116104675760609161008091860161048e565b9301519160ff831683036104675760409384519161009e868461046b565b60018352603160f81b6020840190815281519092906001600160401b03811161037757600354600181811c9116801561045d575b602082101461035957601f81116103fa575b50806020601f8211600114610396575f9161038b575b508160011b915f199060031b1c1916176003555b8051906001600160401b0382116103775760045490600182811c9216801561036d575b60208310146103595781601f8493116102eb575b50602090601f8311600114610285575f9261027a575b50508160011b915f199060031b1c1916176004555b610179816104e3565b610120526101868361066a565b6101405260208151910120918260e05251902080610100524660a05284519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f84528683015260608201524660808201523060a082015260a081526101f160c08261046b565b5190206080523060c052610180526101605251610ecc90816107a3823960805181610b9f015260a05181610c5c015260c05181610b69015260e05181610bee01526101005181610c14015261012051816104d9015261014051816105020152610160518161074e01526101805181818160e1015281816102f9015281816106430152610ac80152f35b015190505f8061015b565b60045f9081528281209350601f198516905b8181106102d357509084600195949392106102bb575b505050811b01600455610170565b01515f1960f88460031b161c191690555f80806102ad565b92936020600181928786015181550195019301610297565b60045f529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c8101916020851061034f575b90601f859493920160051c01905b8181106103415750610145565b5f8155849350600101610334565b9091508190610326565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610131565b634e487b7160e01b5f52604160045260245ffd5b90508301515f6100fa565b60035f9081528181209250601f198416905b8181106103e2575090836001949392106103ca575b5050811b0160035561010e565b8501515f1960f88460031b161c191690555f806103bd565b9192602060018192868a0151815501940192016103a8565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c81019160208410610453575b601f0160051c01905b81811061044857506100e4565b5f815560010161043b565b9091508190610432565b90607f16906100d2565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761037757604052565b81601f82011215610467578051906001600160401b03821161037757604051926104c2601f8401601f19166020018561046b565b8284526020838301011161046757815f9260208093018386015e8301015290565b908151602081105f1461055d575090601f81511161051d57602081519101516020821061050e571790565b5f198260200360031b1b161790565b604460209160405192839163305a27a960e01b83528160048401528051918291826024860152018484015e5f828201840152601f01601f19168101030190fd5b6001600160401b03811161037757600554600181811c91168015610660575b602082101461035957601f811161062d575b50602092601f82116001146105cc57928192935f926105c1575b50508160011b915f199060031b1c19161760055560ff90565b015190505f806105a8565b601f1982169360055f52805f20915f5b86811061061557508360019596106105fd575b505050811b0160055560ff90565b01515f1960f88460031b161c191690555f80806105ef565b919260206001819286850151815501940192016105dc565b60055f52601f60205f20910160051c810190601f830160051c015b818110610655575061058e565b5f8155600101610648565b90607f169061057c565b908151602081105f14610695575090601f81511161051d57602081519101516020821061050e571790565b6001600160401b03811161037757600654600181811c91168015610798575b602082101461035957601f8111610765575b50602092601f821160011461070457928192935f926106f9575b50508160011b915f199060031b1c19161760065560ff90565b015190505f806106e0565b601f1982169360065f52805f20915f5b86811061074d5750836001959610610735575b505050811b0160065560ff90565b01515f1960f88460031b161c191690555f8080610727565b91926020600181928685015181550194019201610714565b60065f52601f60205f20910160051c810190601f830160051c015b81811061078d57506106c6565b5f8155600101610780565b90607f16906106b456fe6080806040526004361015610012575f80fd5b5f3560e01c90816306fdde031461089457508063095ea7b31461086e57806318160ddd1461085157806323b872dd14610772578063313ce567146107355780633644e5151461071357806340c10f191461062857806370a08231146105f15780637ecebe00146105b957806384b0196e146104c157806395d89b41146103df5780639dc29fac146102da578063a9059cbb146102a9578063d505accf14610164578063dd62ed3e146101145763e78cea92146100cc575f80fd5b34610110575f366003190112610110576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5f80fd5b346101105760403660031901126101105761012d61095a565b610135610970565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b346101105760e03660031901126101105761017d61095a565b610185610970565b604435906064359260843560ff8116810361011057844211610296576102596102629160018060a01b03841696875f52600760205260405f20908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c0815261022760e082610a3f565b519020610232610b66565b906040519161190160f01b83526002830152602282015260c43591604260a4359220610d7f565b90929192610e02565b6001600160a01b031684810361027f575061027d9350610c82565b005b84906325c0072360e11b5f5260045260245260445ffd5b8463313c898160e11b5f5260045260245ffd5b34610110576040366003190112610110576102cf6102c561095a565b6024359033610a9e565b602060405160018152f35b34610110576040366003190112610110576102f361095a565b602435907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169061032e33808414610a76565b6001600160a01b03169081156103cc578033141590816103c3575b506103b457805f525f60205260405f205482811061039b576020835f945f516020610e775f395f51905f52938587528684520360408620558060025403600255604051908152a3602060405160018152f35b9063391434e360e21b5f5260045260245260445260645ffd5b63246a105b60e21b5f5260045ffd5b90501583610349565b634b637e8f60e11b5f525f60045260245ffd5b34610110575f366003190112610110576040515f6004546103ff81610986565b808452906001811690811561049d575060011461043f575b61043b8361042781850382610a3f565b604051918291602083526020830190610936565b0390f35b60045f9081527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b80821061048357509091508101602001610427610417565b91926001816020925483858801015201910190929161046b565b60ff191660208086019190915291151560051b840190910191506104279050610417565b34610110575f3660031901126101105761055d6104fd7f0000000000000000000000000000000000000000000000000000000000000000610ce5565b6105267f0000000000000000000000000000000000000000000000000000000000000000610d48565b602061056b604051926105398385610a3f565b5f84525f368137604051958695600f60f81b875260e08588015260e0870190610936565b908582036040870152610936565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b8281106105a257505050500390f35b835185528695509381019392810192600101610593565b34610110576020366003190112610110576001600160a01b036105da61095a565b165f526007602052602060405f2054604051908152f35b34610110576020366003190112610110576001600160a01b0361061261095a565b165f525f602052602060405f2054604051908152f35b346101105760403660031901126101105761064161095a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906024359061067c33808514610a76565b6001600160a01b0316918215610700578033141590816106f6575b506103b457600254908082018092116106e25760205f516020610e775f395f51905f52915f9360025584845283825260408420818154019055604051908152a3602060405160018152f35b634e487b7160e01b5f52601160045260245ffd5b9050821483610697565b63ec442f0560e01b5f525f60045260245ffd5b34610110575f36600319011261011057602061072d610b66565b604051908152f35b34610110575f36600319011261011057602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101105760603660031901126101105761078b61095a565b610793610970565b6001600160a01b0382165f818152600160209081526040808320338452909152902054909260443592915f1981106107d1575b506102cf9350610a9e565b838110610836578415610823573315610810576102cf945f52600160205260405f2060018060a01b0333165f526020528360405f2091039055846107c6565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b34610110575f366003190112610110576020600254604051908152f35b34610110576040366003190112610110576102cf61088a61095a565b6024359033610c82565b34610110575f366003190112610110575f6003546108b181610986565b808452906001811690811561049d57506001146108d85761043b8361042781850382610a3f565b60035f9081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b80821061091c57509091508101602001610427610417565b919260018160209254838588010152019101909291610904565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361011057565b602435906001600160a01b038216820361011057565b90600182811c921680156109b4575b60208310146109a057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610995565b5f92918154916109cd83610986565b8083529260018116908115610a2257506001146109e957505050565b5f9081526020812093945091925b838310610a08575060209250010190565b6001816020929493945483858701015201910191906109f7565b915050602093945060ff929192191683830152151560051b010190565b601f909101601f19168101906001600160401b03821190821017610a6257604052565b634e487b7160e01b5f52604160045260245ffd5b15610a7e5750565b630b56edc360e21b5f9081526001600160a01b0391909116600452602490fd5b6001600160a01b03169081156103cc576001600160a01b0316918215610700576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016338114159081610b5c575b506103b457815f525f60205260405f2054818110610b4357815f516020610e775f395f51905f5292602092855f525f84520360405f2055845f525f825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b905083145f610af3565b307f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161480610c59575b15610bc1577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a08152610c5360c082610a3f565b51902090565b507f00000000000000000000000000000000000000000000000000000000000000004614610b98565b6001600160a01b0316908115610823576001600160a01b03169182156108105760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60ff8114610d2b5760ff811690601f8211610d1c5760405191610d09604084610a3f565b6020808452838101919036833783525290565b632cd44ac360e21b5f5260045ffd5b50604051610d4581610d3e8160056109be565b0382610a3f565b90565b60ff8114610d6c5760ff811690601f8211610d1c5760405191610d09604084610a3f565b50604051610d4581610d3e8160066109be565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411610df7579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610dec575f516001600160a01b03811615610de257905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f9160039190565b6004811015610e625780610e14575050565b60018103610e2b5763f645eedf60e01b5f5260045ffd5b60028103610e46575063fce698f760e01b5f5260045260245ffd5b600314610e505750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212202b471c796cbd6ee8a35b194727b619e4aeb8dc56c6617fd35fbc1c3dadcf41c764736f6c634300081c0033a26469706673582212204a71968660e4af1f7eea53bd93c45c54757a4bab6c5e4b249d6c3d0215ebb5c364736f6c634300081c0033",
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
