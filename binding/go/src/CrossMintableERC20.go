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
	Bin: "0x6101a080604052346104675761166f803803809161001d828561046b565b833981016080828203126104675781516001600160a01b03811681036104675760208301516001600160401b038111610467578261005c91850161048e565b60408401519092906001600160401b0381116104675760609161008091860161048e565b9301519160ff831683036104675760409384519161009e868461046b565b60018352603160f81b6020840190815281519092906001600160401b03811161037757600354600181811c9116801561045d575b602082101461035957601f81116103fa575b50806020601f8211600114610396575f9161038b575b508160011b915f199060031b1c1916176003555b8051906001600160401b0382116103775760045490600182811c9216801561036d575b60208310146103595781601f8493116102eb575b50602090601f8311600114610285575f9261027a575b50508160011b915f199060031b1c1916176004555b610179816104e3565b610120526101868361066a565b6101405260208151910120918260e05251902080610100524660a05284519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f84528683015260608201524660808201523060a082015260a081526101f160c08261046b565b5190206080523060c052610180526101605251610ecc90816107a3823960805181610b9f015260a05181610c5c015260c05181610b69015260e05181610bee01526101005181610c14015261012051816104d9015261014051816105020152610160518161074e01526101805181818160e1015281816102f9015281816106430152610ac80152f35b015190505f8061015b565b60045f9081528281209350601f198516905b8181106102d357509084600195949392106102bb575b505050811b01600455610170565b01515f1960f88460031b161c191690555f80806102ad565b92936020600181928786015181550195019301610297565b60045f529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c8101916020851061034f575b90601f859493920160051c01905b8181106103415750610145565b5f8155849350600101610334565b9091508190610326565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610131565b634e487b7160e01b5f52604160045260245ffd5b90508301515f6100fa565b60035f9081528181209250601f198416905b8181106103e2575090836001949392106103ca575b5050811b0160035561010e565b8501515f1960f88460031b161c191690555f806103bd565b9192602060018192868a0151815501940192016103a8565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c81019160208410610453575b601f0160051c01905b81811061044857506100e4565b5f815560010161043b565b9091508190610432565b90607f16906100d2565b5f80fd5b601f909101601f19168101906001600160401b0382119082101761037757604052565b81601f82011215610467578051906001600160401b03821161037757604051926104c2601f8401601f19166020018561046b565b8284526020838301011161046757815f9260208093018386015e8301015290565b908151602081105f1461055d575090601f81511161051d57602081519101516020821061050e571790565b5f198260200360031b1b161790565b604460209160405192839163305a27a960e01b83528160048401528051918291826024860152018484015e5f828201840152601f01601f19168101030190fd5b6001600160401b03811161037757600554600181811c91168015610660575b602082101461035957601f811161062d575b50602092601f82116001146105cc57928192935f926105c1575b50508160011b915f199060031b1c19161760055560ff90565b015190505f806105a8565b601f1982169360055f52805f20915f5b86811061061557508360019596106105fd575b505050811b0160055560ff90565b01515f1960f88460031b161c191690555f80806105ef565b919260206001819286850151815501940192016105dc565b60055f52601f60205f20910160051c810190601f830160051c015b818110610655575061058e565b5f8155600101610648565b90607f169061057c565b908151602081105f14610695575090601f81511161051d57602081519101516020821061050e571790565b6001600160401b03811161037757600654600181811c91168015610798575b602082101461035957601f8111610765575b50602092601f821160011461070457928192935f926106f9575b50508160011b915f199060031b1c19161760065560ff90565b015190505f806106e0565b601f1982169360065f52805f20915f5b86811061074d5750836001959610610735575b505050811b0160065560ff90565b01515f1960f88460031b161c191690555f8080610727565b91926020600181928685015181550194019201610714565b60065f52601f60205f20910160051c810190601f830160051c015b81811061078d57506106c6565b5f8155600101610780565b90607f16906106b456fe6080806040526004361015610012575f80fd5b5f3560e01c90816306fdde031461089457508063095ea7b31461086e57806318160ddd1461085157806323b872dd14610772578063313ce567146107355780633644e5151461071357806340c10f191461062857806370a08231146105f15780637ecebe00146105b957806384b0196e146104c157806395d89b41146103df5780639dc29fac146102da578063a9059cbb146102a9578063d505accf14610164578063dd62ed3e146101145763e78cea92146100cc575f80fd5b34610110575f366003190112610110576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5f80fd5b346101105760403660031901126101105761012d61095a565b610135610970565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b346101105760e03660031901126101105761017d61095a565b610185610970565b604435906064359260843560ff8116810361011057844211610296576102596102629160018060a01b03841696875f52600760205260405f20908154916001830190556040519060208201927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c984528a604084015260018060a01b038916606084015289608084015260a083015260c082015260c0815261022760e082610a3f565b519020610232610b66565b906040519161190160f01b83526002830152602282015260c43591604260a4359220610d7f565b90929192610e02565b6001600160a01b031684810361027f575061027d9350610c82565b005b84906325c0072360e11b5f5260045260245260445ffd5b8463313c898160e11b5f5260045260245ffd5b34610110576040366003190112610110576102cf6102c561095a565b6024359033610a9e565b602060405160018152f35b34610110576040366003190112610110576102f361095a565b602435907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169061032e33808414610a76565b6001600160a01b03169081156103cc578033141590816103c3575b506103b457805f525f60205260405f205482811061039b576020835f945f516020610e775f395f51905f52938587528684520360408620558060025403600255604051908152a3602060405160018152f35b9063391434e360e21b5f5260045260245260445260645ffd5b63246a105b60e21b5f5260045ffd5b90501583610349565b634b637e8f60e11b5f525f60045260245ffd5b34610110575f366003190112610110576040515f6004546103ff81610986565b808452906001811690811561049d575060011461043f575b61043b8361042781850382610a3f565b604051918291602083526020830190610936565b0390f35b60045f9081527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b80821061048357509091508101602001610427610417565b91926001816020925483858801015201910190929161046b565b60ff191660208086019190915291151560051b840190910191506104279050610417565b34610110575f3660031901126101105761055d6104fd7f0000000000000000000000000000000000000000000000000000000000000000610ce5565b6105267f0000000000000000000000000000000000000000000000000000000000000000610d48565b602061056b604051926105398385610a3f565b5f84525f368137604051958695600f60f81b875260e08588015260e0870190610936565b908582036040870152610936565b4660608501523060808501525f60a085015283810360c08501528180845192838152019301915f5b8281106105a257505050500390f35b835185528695509381019392810192600101610593565b34610110576020366003190112610110576001600160a01b036105da61095a565b165f526007602052602060405f2054604051908152f35b34610110576020366003190112610110576001600160a01b0361061261095a565b165f525f602052602060405f2054604051908152f35b346101105760403660031901126101105761064161095a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906024359061067c33808514610a76565b6001600160a01b0316918215610700578033141590816106f6575b506103b457600254908082018092116106e25760205f516020610e775f395f51905f52915f9360025584845283825260408420818154019055604051908152a3602060405160018152f35b634e487b7160e01b5f52601160045260245ffd5b9050821483610697565b63ec442f0560e01b5f525f60045260245ffd5b34610110575f36600319011261011057602061072d610b66565b604051908152f35b34610110575f36600319011261011057602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101105760603660031901126101105761078b61095a565b610793610970565b6001600160a01b0382165f818152600160209081526040808320338452909152902054909260443592915f1981106107d1575b506102cf9350610a9e565b838110610836578415610823573315610810576102cf945f52600160205260405f2060018060a01b0333165f526020528360405f2091039055846107c6565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b34610110575f366003190112610110576020600254604051908152f35b34610110576040366003190112610110576102cf61088a61095a565b6024359033610c82565b34610110575f366003190112610110575f6003546108b181610986565b808452906001811690811561049d57506001146108d85761043b8361042781850382610a3f565b60035f9081527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b80821061091c57509091508101602001610427610417565b919260018160209254838588010152019101909291610904565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361011057565b602435906001600160a01b038216820361011057565b90600182811c921680156109b4575b60208310146109a057565b634e487b7160e01b5f52602260045260245ffd5b91607f1691610995565b5f92918154916109cd83610986565b8083529260018116908115610a2257506001146109e957505050565b5f9081526020812093945091925b838310610a08575060209250010190565b6001816020929493945483858701015201910191906109f7565b915050602093945060ff929192191683830152151560051b010190565b601f909101601f19168101906001600160401b03821190821017610a6257604052565b634e487b7160e01b5f52604160045260245ffd5b15610a7e5750565b630b56edc360e21b5f9081526001600160a01b0391909116600452602490fd5b6001600160a01b03169081156103cc576001600160a01b0316918215610700576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016338114159081610b5c575b506103b457815f525f60205260405f2054818110610b4357815f516020610e775f395f51905f5292602092855f525f84520360405f2055845f525f825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b905083145f610af3565b307f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161480610c59575b15610bc1577f000000000000000000000000000000000000000000000000000000000000000090565b60405160208101907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82527f000000000000000000000000000000000000000000000000000000000000000060408201527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260a08152610c5360c082610a3f565b51902090565b507f00000000000000000000000000000000000000000000000000000000000000004614610b98565b6001600160a01b0316908115610823576001600160a01b03169182156108105760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60ff8114610d2b5760ff811690601f8211610d1c5760405191610d09604084610a3f565b6020808452838101919036833783525290565b632cd44ac360e21b5f5260045ffd5b50604051610d4581610d3e8160056109be565b0382610a3f565b90565b60ff8114610d6c5760ff811690601f8211610d1c5760405191610d09604084610a3f565b50604051610d4581610d3e8160066109be565b91906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038411610df7579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa15610dec575f516001600160a01b03811615610de257905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f9160039190565b6004811015610e625780610e14575050565b60018103610e2b5763f645eedf60e01b5f5260045ffd5b60028103610e46575063fce698f760e01b5f5260045260245ffd5b600314610e505750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212202b471c796cbd6ee8a35b194727b619e4aeb8dc56c6617fd35fbc1c3dadcf41c764736f6c634300081c0033",
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
