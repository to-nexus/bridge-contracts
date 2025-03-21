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
	Bin: "0x6101a0604052348015610010575f5ffd5b5060405161298f38038061298f83398181016040528101906100329190610478565b82806040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152508585816003908161007b9190610724565b50806004908161008b9190610724565b5050506100a260058361018660201b90919060201c565b61012081815250506100be60068261018660201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a081815250506100fb6101d360201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250505050508373ffffffffffffffffffffffffffffffffffffffff166101808173ffffffffffffffffffffffffffffffffffffffff16815250508060ff166101608160ff168152505050505050610975565b5f6020835110156101a7576101a08361022d60201b60201c565b90506101cd565b826101b78361029260201b60201c565b5f0190816101c59190610724565b5060ff5f1b90505b92915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e051610100514630604051602001610212959493929190610829565b60405160208183030381529060405280519060200120905090565b5f5f829050601f8151111561027957826040517f305a27a900000000000000000000000000000000000000000000000000000000815260040161027091906108c2565b60405180910390fd5b8051816102859061090f565b5f1c175f1b915050919050565b5f819050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102d5826102ac565b9050919050565b6102e5816102cb565b81146102ef575f5ffd5b50565b5f81519050610300816102dc565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103548261030e565b810181811067ffffffffffffffff821117156103735761037261031e565b5b80604052505050565b5f61038561029b565b9050610391828261034b565b919050565b5f67ffffffffffffffff8211156103b0576103af61031e565b5b6103b98261030e565b9050602081019050919050565b8281835e5f83830152505050565b5f6103e66103e184610396565b61037c565b9050828152602081018484840111156104025761040161030a565b5b61040d8482856103c6565b509392505050565b5f82601f83011261042957610428610306565b5b81516104398482602086016103d4565b91505092915050565b5f60ff82169050919050565b61045781610442565b8114610461575f5ffd5b50565b5f815190506104728161044e565b92915050565b5f5f5f5f608085870312156104905761048f6102a4565b5b5f61049d878288016102f2565b945050602085015167ffffffffffffffff8111156104be576104bd6102a8565b5b6104ca87828801610415565b935050604085015167ffffffffffffffff8111156104eb576104ea6102a8565b5b6104f787828801610415565b925050606061050887828801610464565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061056257607f821691505b6020821081036105755761057461051e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026105d77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261059c565b6105e1868361059c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61062561062061061b846105f9565b610602565b6105f9565b9050919050565b5f819050919050565b61063e8361060b565b61065261064a8261062c565b8484546105a8565b825550505050565b5f5f905090565b61066961065a565b610674818484610635565b505050565b5b818110156106975761068c5f82610661565b60018101905061067a565b5050565b601f8211156106dc576106ad8161057b565b6106b68461058d565b810160208510156106c5578190505b6106d96106d18561058d565b830182610679565b50505b505050565b5f82821c905092915050565b5f6106fc5f19846008026106e1565b1980831691505092915050565b5f61071483836106ed565b9150826002028217905092915050565b61072d82610514565b67ffffffffffffffff8111156107465761074561031e565b5b610750825461054b565b61075b82828561069b565b5f60209050601f83116001811461078c575f841561077a578287015190505b6107848582610709565b8655506107eb565b601f19841661079a8661057b565b5f5b828110156107c15784890151825560018201915060208501945060208101905061079c565b868310156107de57848901516107da601f8916826106ed565b8355505b6001600288020188555050505b505050505050565b5f819050919050565b610805816107f3565b82525050565b610814816105f9565b82525050565b610823816102cb565b82525050565b5f60a08201905061083c5f8301886107fc565b61084960208301876107fc565b61085660408301866107fc565b610863606083018561080b565b610870608083018461081a565b9695505050505050565b5f82825260208201905092915050565b5f61089482610514565b61089e818561087a565b93506108ae8185602086016103c6565b6108b78161030e565b840191505092915050565b5f6020820190508181035f8301526108da818461088a565b905092915050565b5f81519050919050565b5f819050602082019050919050565b5f61090682516107f3565b80915050919050565b5f610919826108e2565b82610923846108ec565b905061092e816108fb565b9250602082101561096e576109697fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080261059c565b831692505b5050919050565b60805160a05160c05160e0516101005161012051610140516101605161018051611f996109f65f395f81816104ae015281816106e00152818161095801528181610fbc015261101a01525f61046201525f610ca001525f610c6501525f6110f401525f6110d301525f610b1901525f610b6f01525f610b980152611f995ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c80637ecebe0011610095578063a9059cbb11610064578063a9059cbb146102dc578063d505accf1461030c578063dd62ed3e14610328578063e78cea9214610358576100fe565b80637ecebe001461023a57806384b0196e1461026a57806395d89b411461028e5780639dc29fac146102ac576100fe565b8063313ce567116100d1578063313ce5671461019e5780633644e515146101bc57806340c10f19146101da57806370a082311461020a576100fe565b806306fdde0314610102578063095ea7b31461012057806318160ddd1461015057806323b872dd1461016e575b5f5ffd5b61010a610376565b6040516101179190611809565b60405180910390f35b61013a600480360381019061013591906118ba565b610406565b6040516101479190611912565b60405180910390f35b610158610428565b604051610165919061193a565b60405180910390f35b61018860048036038101906101839190611953565b610431565b6040516101959190611912565b60405180910390f35b6101a661045f565b6040516101b391906119be565b60405180910390f35b6101c4610486565b6040516101d191906119ef565b60405180910390f35b6101f460048036038101906101ef91906118ba565b610494565b6040516102019190611912565b60405180910390f35b610224600480360381019061021f9190611a08565b61053b565b604051610231919061193a565b60405180910390f35b610254600480360381019061024f9190611a08565b610580565b604051610261919061193a565b60405180910390f35b610272610591565b6040516102859796959493929190611b33565b60405180910390f35b610296610636565b6040516102a39190611809565b60405180910390f35b6102c660048036038101906102c191906118ba565b6106c6565b6040516102d39190611912565b60405180910390f35b6102f660048036038101906102f191906118ba565b61076d565b6040516103039190611912565b60405180910390f35b61032660048036038101906103219190611c09565b61078f565b005b610342600480360381019061033d9190611ca6565b6108d4565b60405161034f919061193a565b60405180910390f35b610360610956565b60405161036d9190611ce4565b60405180910390f35b60606003805461038590611d2a565b80601f01602080910402602001604051908101604052809291908181526020018280546103b190611d2a565b80156103fc5780601f106103d3576101008083540402835291602001916103fc565b820191905f5260205f20905b8154815290600101906020018083116103df57829003601f168201915b5050505050905090565b5f5f61041061097a565b905061041d818585610981565b600191505092915050565b5f600254905090565b5f5f61043b61097a565b9050610448858285610993565b610453858585610a26565b60019150509392505050565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f61048f610b16565b905090565b5f3373ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16143390610526576040517f2d5bb70c00000000000000000000000000000000000000000000000000000000815260040161051d9190611ce4565b60405180910390fd5b506105318383610bcc565b6001905092915050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f61058a82610c4b565b9050919050565b5f6060805f5f5f60606105a2610c5c565b6105aa610c97565b46305f5f1b5f67ffffffffffffffff8111156105c9576105c8611d5a565b5b6040519080825280602002602001820160405280156105f75781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b60606004805461064590611d2a565b80601f016020809104026020016040519081016040528092919081815260200182805461067190611d2a565b80156106bc5780601f10610693576101008083540402835291602001916106bc565b820191905f5260205f20905b81548152906001019060200180831161069f57829003601f168201915b5050505050905090565b5f3373ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16143390610758576040517f2d5bb70c00000000000000000000000000000000000000000000000000000000815260040161074f9190611ce4565b60405180910390fd5b506107638383610cd2565b6001905092915050565b5f5f61077761097a565b9050610784818585610a26565b600191505092915050565b834211156107d457836040517f627913020000000000000000000000000000000000000000000000000000000081526004016107cb919061193a565b60405180910390fd5b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886108028c610d51565b8960405160200161081896959493929190611d87565b6040516020818303038152906040528051906020012090505f61083a82610da4565b90505f61084982878787610dbd565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146108bd57808a6040517f4b800e460000000000000000000000000000000000000000000000000000000081526004016108b4929190611de6565b60405180910390fd5b6108c88a8a8a610981565b50505050505050505050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f33905090565b61098e8383836001610deb565b505050565b5f61099e84846108d4565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610a205781811015610a11578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610a0893929190611e0d565b60405180910390fd5b610a1f84848484035f610deb565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a96575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610a8d9190611ce4565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b06575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610afd9190611ce4565b60405180910390fd5b610b11838383610fba565b505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148015610b9157507f000000000000000000000000000000000000000000000000000000000000000046145b15610bbe577f00000000000000000000000000000000000000000000000000000000000000009050610bc9565b610bc66110af565b90505b90565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c3c575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610c339190611ce4565b60405180910390fd5b610c475f8383610fba565b5050565b5f610c5582611144565b9050919050565b6060610c9260057f000000000000000000000000000000000000000000000000000000000000000061118a90919063ffffffff16565b905090565b6060610ccd60067f000000000000000000000000000000000000000000000000000000000000000061118a90919063ffffffff16565b905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d42575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610d399190611ce4565b60405180910390fd5b610d4d825f83610fba565b5050565b5f60075f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050559050919050565b5f610db6610db0610b16565b83611237565b9050919050565b5f5f5f5f610dcd88888888611277565b925092509250610ddd828261135e565b829350505050949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610e5b575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610e529190611ce4565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ecb575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610ec29190611ce4565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610fb4578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610fab919061193a565b60405180910390a35b50505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610ff961097a565b73ffffffffffffffffffffffffffffffffffffffff161415801561106857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b1561109f576040517f91a8416c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110aa8383836114c0565b505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000004630604051602001611129959493929190611e42565b60405160208183030381529060405280519060200120905090565b5f60075f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b606060ff5f1b83146111a65761119f836116d9565b9050611231565b8180546111b290611d2a565b80601f01602080910402602001604051908101604052809291908181526020018280546111de90611d2a565b80156112295780601f1061120057610100808354040283529160200191611229565b820191905f5260205f20905b81548152906001019060200180831161120c57829003601f168201915b505050505090505b92915050565b5f6040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b5f5f5f7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0845f1c11156112b3575f600385925092509250611354565b5f6001888888886040515f81526020016040526040516112d69493929190611e93565b6020604051602081039080840390855afa1580156112f6573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611347575f60015f5f1b93509350935050611354565b805f5f5f1b935093509350505b9450945094915050565b5f600381111561137157611370611ed6565b5b82600381111561138457611383611ed6565b5b03156114bc576001600381111561139e5761139d611ed6565b5b8260038111156113b1576113b0611ed6565b5b036113e8576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260038111156113fc576113fb611ed6565b5b82600381111561140f5761140e611ed6565b5b0361145357805f1c6040517ffce698f700000000000000000000000000000000000000000000000000000000815260040161144a919061193a565b60405180910390fd5b60038081111561146657611465611ed6565b5b82600381111561147957611478611ed6565b5b036114bb57806040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004016114b291906119ef565b60405180910390fd5b5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611510578060025f8282546115049190611f30565b925050819055506115de565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015611599578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161159093929190611e0d565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611625578060025f828254039250508190555061166f565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516116cc919061193a565b60405180910390a3505050565b60605f6116e58361174b565b90505f602067ffffffffffffffff81111561170357611702611d5a565b5b6040519080825280601f01601f1916602001820160405280156117355781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b5f5f60ff835f1c169050601f811115611790576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6117db82611799565b6117e581856117a3565b93506117f58185602086016117b3565b6117fe816117c1565b840191505092915050565b5f6020820190508181035f83015261182181846117d1565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6118568261182d565b9050919050565b6118668161184c565b8114611870575f5ffd5b50565b5f813590506118818161185d565b92915050565b5f819050919050565b61189981611887565b81146118a3575f5ffd5b50565b5f813590506118b481611890565b92915050565b5f5f604083850312156118d0576118cf611829565b5b5f6118dd85828601611873565b92505060206118ee858286016118a6565b9150509250929050565b5f8115159050919050565b61190c816118f8565b82525050565b5f6020820190506119255f830184611903565b92915050565b61193481611887565b82525050565b5f60208201905061194d5f83018461192b565b92915050565b5f5f5f6060848603121561196a57611969611829565b5b5f61197786828701611873565b935050602061198886828701611873565b9250506040611999868287016118a6565b9150509250925092565b5f60ff82169050919050565b6119b8816119a3565b82525050565b5f6020820190506119d15f8301846119af565b92915050565b5f819050919050565b6119e9816119d7565b82525050565b5f602082019050611a025f8301846119e0565b92915050565b5f60208284031215611a1d57611a1c611829565b5b5f611a2a84828501611873565b91505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b611a6781611a33565b82525050565b611a768161184c565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611aae81611887565b82525050565b5f611abf8383611aa5565b60208301905092915050565b5f602082019050919050565b5f611ae182611a7c565b611aeb8185611a86565b9350611af683611a96565b805f5b83811015611b26578151611b0d8882611ab4565b9750611b1883611acb565b925050600181019050611af9565b5085935050505092915050565b5f60e082019050611b465f83018a611a5e565b8181036020830152611b5881896117d1565b90508181036040830152611b6c81886117d1565b9050611b7b606083018761192b565b611b886080830186611a6d565b611b9560a08301856119e0565b81810360c0830152611ba78184611ad7565b905098975050505050505050565b611bbe816119a3565b8114611bc8575f5ffd5b50565b5f81359050611bd981611bb5565b92915050565b611be8816119d7565b8114611bf2575f5ffd5b50565b5f81359050611c0381611bdf565b92915050565b5f5f5f5f5f5f5f60e0888a031215611c2457611c23611829565b5b5f611c318a828b01611873565b9750506020611c428a828b01611873565b9650506040611c538a828b016118a6565b9550506060611c648a828b016118a6565b9450506080611c758a828b01611bcb565b93505060a0611c868a828b01611bf5565b92505060c0611c978a828b01611bf5565b91505092959891949750929550565b5f5f60408385031215611cbc57611cbb611829565b5b5f611cc985828601611873565b9250506020611cda85828601611873565b9150509250929050565b5f602082019050611cf75f830184611a6d565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d4157607f821691505b602082108103611d5457611d53611cfd565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60c082019050611d9a5f8301896119e0565b611da76020830188611a6d565b611db46040830187611a6d565b611dc1606083018661192b565b611dce608083018561192b565b611ddb60a083018461192b565b979650505050505050565b5f604082019050611df95f830185611a6d565b611e066020830184611a6d565b9392505050565b5f606082019050611e205f830186611a6d565b611e2d602083018561192b565b611e3a604083018461192b565b949350505050565b5f60a082019050611e555f8301886119e0565b611e6260208301876119e0565b611e6f60408301866119e0565b611e7c606083018561192b565b611e896080830184611a6d565b9695505050505050565b5f608082019050611ea65f8301876119e0565b611eb360208301866119af565b611ec060408301856119e0565b611ecd60608301846119e0565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f3a82611887565b9150611f4583611887565b9250828201905080821115611f5d57611f5c611f03565b5b9291505056fea2646970667358221220d33afb95830ada4403f4b69d2c3e807bc60dc6f878d012b2624468b6ec249e3e64736f6c634300081c0033",
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
