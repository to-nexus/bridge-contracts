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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"allTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createCrossMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossMintableERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"errCrossMintableERC20FactoryNotBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"errCrossMintableERC20FactoryNotExist\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"6ff97f1d": "allTokens()",
		"e78cea92": "bridge()",
		"4804a041": "createCrossMintableERC20(bytes32,string,string,uint8)",
		"76a67a51": "pause(address)",
		"4f6ccce7": "tokenByIndex(uint256)",
		"d92fc67b": "tokensLength()",
		"57b001f9": "unpause(address)",
	},
	Bin: "0x6080604052348015600e575f5ffd5b50604051612051380380612051833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b611fcb806100865f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80636ff97f1d116100585780636ff97f1d146100d657806376a67a51146100eb578063d92fc67b146100fe578063e78cea9214610114575f5ffd5b80634804a0411461007e5780634f6ccce7146100ae57806357b001f9146100c1575b5f5ffd5b61009161008c36600461058d565b610126565b6040516001600160a01b0390911681526020015b60405180910390f35b6100916100bc366004610612565b6101ef565b6100d46100cf366004610629565b610201565b005b6100de6102c2565b6040516100a5919061064f565b6100d46100f9366004610629565b6102d3565b61010661037b565b6040519081526020016100a5565b5f54610091906001600160a01b031681565b5f805433906001600160a01b03168114610164576040516346570c9d60e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b505f60405180602001610176906104e1565b601f1982820381018352601f9091011660408190525f546101ab916001600160a01b03909116908890889088906020016106c8565b60408051601f19818403018152908290526101c99291602001610729565b60405160208183030381529060405290506101e55f8783610386565b9695505050505050565b5f6101fb600183610421565b92915050565b5f5433906001600160a01b03168114610239576040516346570c9d60e11b81526001600160a01b03909116600482015260240161015b565b5061024560018261042c565b81906102705760405163ca4c700160e01b81526001600160a01b03909116600482015260240161015b565b50806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156102a9575f5ffd5b505af11580156102bb573d5f5f3e3d5ffd5b5050505050565b60606102ce600161044d565b905090565b5f5433906001600160a01b0316811461030b576040516346570c9d60e11b81526001600160a01b03909116600482015260240161015b565b5061031760018261042c565b81906103425760405163ca4c700160e01b81526001600160a01b03909116600482015260240161015b565b50806001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156102a9575f5ffd5b5f6102ce6001610459565b5f834710156103b15760405163cf47918160e01b81524760048201526024810185905260440161015b565b81515f036103d257604051631328927760e21b815260040160405180910390fd5b8282516020840186f590503d1519811516156103f3576040513d5f823e3d81fd5b6001600160a01b03811661041a5760405163b06ebf3d60e01b815260040160405180910390fd5b9392505050565b5f61041a8383610462565b6001600160a01b0381165f908152600183016020526040812054151561041a565b60605f61041a83610488565b5f6101fb825490565b5f825f01828154811061047757610477610745565b905f5260205f200154905092915050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156104d557602002820191905f5260205f20905b8154815260200190600101908083116104c1575b50505050509050919050565b61183c8061075a83390190565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610511575f5ffd5b813567ffffffffffffffff81111561052b5761052b6104ee565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561055a5761055a6104ee565b604052818152838201602001851015610571575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f608085870312156105a0575f5ffd5b84359350602085013567ffffffffffffffff8111156105bd575f5ffd5b6105c987828801610502565b935050604085013567ffffffffffffffff8111156105e5575f5ffd5b6105f187828801610502565b925050606085013560ff81168114610607575f5ffd5b939692955090935050565b5f60208284031215610622575f5ffd5b5035919050565b5f60208284031215610639575f5ffd5b81356001600160a01b038116811461041a575f5ffd5b602080825282518282018190525f918401906040840190835b8181101561068f5783516001600160a01b0316835260209384019390920191600101610668565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03851681526080602082018190525f906106eb9083018661069a565b82810360408401526106fd818661069a565b91505060ff8316606083015295945050505050565b5f81518060208401855e5f93019283525090919050565b5f61073d6107378386610712565b84610712565b949350505050565b634e487b7160e01b5f52603260045260245ffdfe610180604052348015610010575f5ffd5b5060405161183c38038061183c83398101604081905261002f916102fb565b8280604051806040016040528060018152602001603160f81b815250858561005b61019c60201b60201c565b6001600160a01b03811661008957604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610092816101a0565b505f805460ff60a01b1916905560046100ab8382610419565b5060056100b88282610419565b506100c8915083905060066101ef565b610120526100d78160076101ef565b61014052815160208084019190912060e052815190820120610100524660a05261016360e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250600980546001600160a01b0319166001600160a01b039590951694909417909355505060ff166101605261052b565b3390565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208351101561020a5761020383610221565b905061021b565b816102158482610419565b5060ff90505b92915050565b5f5f829050601f8151111561024b578260405163305a27a960e01b815260040161008091906104d3565b805161025682610508565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610281575f5ffd5b81516001600160401b0381111561029a5761029a61025e565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102c8576102c861025e565b6040528181528382016020018510156102df575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f6080858703121561030e575f5ffd5b84516001600160a01b0381168114610324575f5ffd5b60208601519094506001600160401b0381111561033f575f5ffd5b61034b87828801610272565b604087015190945090506001600160401b03811115610368575f5ffd5b61037487828801610272565b925050606085015160ff8116811461038a575f5ffd5b939692955090935050565b600181811c908216806103a957607f821691505b6020821081036103c757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561041457805f5260205f20601f840160051c810160208510156103f25750805b601f840160051c820191505b81811015610411575f81556001016103fe565b50505b505050565b81516001600160401b038111156104325761043261025e565b610446816104408454610395565b846103cd565b6020601f821160018114610478575f83156104615750848201515b5f19600385901b1c1916600184901b178455610411565b5f84815260208120601f198516915b828110156104a75787850151825560209485019460019092019101610487565b50848210156104c457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156103c7575f1960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051610160516112b66105865f395f6101ae01525f610a6701525f610a3a01525f61087d01525f61085501525f6107b001525f6107da01525f61080401526112b65ff3fe608060405234801561000f575f5ffd5b506004361061013d575f3560e01c80637ecebe00116100b45780639dc29fac116100795780639dc29fac146102a0578063a9059cbb146102b3578063d505accf146102c6578063dd62ed3e146102d9578063e78cea9214610311578063f2fde38b14610324575f5ffd5b80637ecebe001461023e5780638456cb591461025157806384b0196e146102595780638da5cb5b1461027457806395d89b4114610298575f5ffd5b80633644e515116101055780633644e515146101d85780633f4ba83a146101e057806340c10f19146101ea5780635c975abb146101fd57806370a082311461020e578063715018a614610236575f5ffd5b806306fdde0314610141578063095ea7b31461015f57806318160ddd1461018257806323b872dd14610194578063313ce567146101a7575b5f5ffd5b610149610337565b6040516101569190611032565b60405180910390f35b61017261016d366004611066565b6103c7565b6040519015158152602001610156565b6003545b604051908152602001610156565b6101726101a236600461108e565b6103e0565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610156565b610186610403565b6101e8610411565b005b6101726101f8366004611066565b610423565b5f54600160a01b900460ff16610172565b61018661021c3660046110c8565b6001600160a01b03165f9081526001602052604090205490565b6101e8610477565b61018661024c3660046110c8565b610488565b6101e8610492565b6102616104a2565b60405161015697969594939291906110e1565b5f546001600160a01b03165b6040516001600160a01b039091168152602001610156565b6101496104e4565b6101726102ae366004611066565b6104f3565b6101726102c1366004611066565b610539565b6101e86102d4366004611177565b610546565b6101866102e73660046111e4565b6001600160a01b039182165f90815260026020908152604080832093909416825291909152205490565b600954610280906001600160a01b031681565b6101e86103323660046110c8565b61067c565b60606004805461034690611215565b80601f016020809104026020016040519081016040528092919081815260200182805461037290611215565b80156103bd5780601f10610394576101008083540402835291602001916103bd565b820191905f5260205f20905b8154815290600101906020018083116103a057829003601f168201915b5050505050905090565b5f336103d48185856106b9565b60019150505b92915050565b5f336103ed8582856106cb565b6103f8858585610747565b506001949350505050565b5f61040c6107a4565b905090565b6104196108cd565b6104216108f9565b565b6009545f9033906001600160a01b03168114610463576040516338d2446d60e11b81526001600160a01b0390911660048201526024015b60405180910390fd5b5061046e838361094d565b50600192915050565b61047f6108cd565b6104215f610985565b5f6103da826109d4565b61049a6108cd565b6104216109f1565b5f6060805f5f5f60606104b3610a33565b6104bb610a60565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606005805461034690611215565b6009545f9033906001600160a01b0316811461052e576040516338d2446d60e11b81526001600160a01b03909116600482015260240161045a565b5061046e8383610a8d565b5f336103d4818585610747565b8342111561056a5760405163313c898160e11b81526004810185905260240161045a565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886105b58c6001600160a01b03165f90815260086020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61060f82610ac1565b90505f61061e82878787610aed565b9050896001600160a01b0316816001600160a01b031614610665576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161045a565b6106708a8a8a6106b9565b50505050505050505050565b6106846108cd565b6001600160a01b0381166106ad57604051631e4fbdf760e01b81525f600482015260240161045a565b6106b681610985565b50565b6106c68383836001610b19565b505050565b6001600160a01b038381165f908152600260209081526040808320938616835292905220545f19811015610741578181101561073357604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161045a565b61074184848484035f610b19565b50505050565b6001600160a01b03831661077057604051634b637e8f60e11b81525f600482015260240161045a565b6001600160a01b0382166107995760405163ec442f0560e01b81525f600482015260240161045a565b6106c6838383610beb565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156107fc57507f000000000000000000000000000000000000000000000000000000000000000046145b1561082657507f000000000000000000000000000000000000000000000000000000000000000090565b61040c604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f546001600160a01b031633146104215760405163118cdaa760e01b815233600482015260240161045a565b610901610bfe565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b0382166109765760405163ec442f0560e01b81525f600482015260240161045a565b6109815f8383610beb565b5050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381165f908152600860205260408120546103da565b6109f9610c27565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586109303390565b606061040c7f00000000000000000000000000000000000000000000000000000000000000006006610c51565b606061040c7f00000000000000000000000000000000000000000000000000000000000000006007610c51565b6001600160a01b038216610ab657604051634b637e8f60e11b81525f600482015260240161045a565b610981825f83610beb565b5f6103da610acd6107a4565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610afd88888888610cfa565b925092509250610b0d8282610dc2565b50909695505050505050565b6001600160a01b038416610b425760405163e602df0560e01b81525f600482015260240161045a565b6001600160a01b038316610b6b57604051634a1406b160e11b81525f600482015260240161045a565b6001600160a01b038085165f908152600260209081526040808320938716835292905220829055801561074157826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bdd91815260200190565b60405180910390a350505050565b610bf3610c27565b6106c6838383610e7a565b5f54600160a01b900460ff1661042157604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff16156104215760405163d93c066560e01b815260040160405180910390fd5b606060ff8314610c6b57610c6483610fa0565b90506103da565b818054610c7790611215565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca390611215565b8015610cee5780601f10610cc557610100808354040283529160200191610cee565b820191905f5260205f20905b815481529060010190602001808311610cd157829003601f168201915b505050505090506103da565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610d3357505f91506003905082610db8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610d84573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610daf57505f925060019150829050610db8565b92505f91508190505b9450945094915050565b5f826003811115610dd557610dd561124d565b03610dde575050565b6001826003811115610df257610df261124d565b03610e105760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610e2457610e2461124d565b03610e455760405163fce698f760e01b81526004810182905260240161045a565b6003826003811115610e5957610e5961124d565b03610981576040516335e2f38360e21b81526004810182905260240161045a565b6001600160a01b038316610ea4578060035f828254610e999190611261565b90915550610f149050565b6001600160a01b0383165f9081526001602052604090205481811015610ef65760405163391434e360e21b81526001600160a01b0385166004820152602481018290526044810183905260640161045a565b6001600160a01b0384165f9081526001602052604090209082900390555b6001600160a01b038216610f3057600380548290039055610f4e565b6001600160a01b0382165f9081526001602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f9391815260200190565b60405180910390a3505050565b60605f610fac83610fdd565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f8111156103da57604051632cd44ac360e21b815260040160405180910390fd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6110446020830184611004565b9392505050565b80356001600160a01b0381168114611061575f5ffd5b919050565b5f5f60408385031215611077575f5ffd5b6110808361104b565b946020939093013593505050565b5f5f5f606084860312156110a0575f5ffd5b6110a98461104b565b92506110b76020850161104b565b929592945050506040919091013590565b5f602082840312156110d8575f5ffd5b6110448261104b565b60ff60f81b8816815260e060208201525f6110ff60e0830189611004565b82810360408401526111118189611004565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611166578351835260209384019390920191600101611148565b50909b9a5050505050505050505050565b5f5f5f5f5f5f5f60e0888a03121561118d575f5ffd5b6111968861104b565b96506111a46020890161104b565b95506040880135945060608801359350608088013560ff811681146111c7575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156111f5575f5ffd5b6111fe8361104b565b915061120c6020840161104b565b90509250929050565b600181811c9082168061122957607f821691505b60208210810361124757634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b808201808211156103da57634e487b7160e01b5f52601160045260245ffdfea26469706673582212204500fb9c1f281ed08b6af7229bf86f9d987a88b7ac8404e69da0aca9d08ba62064736f6c634300081c0033a2646970667358221220688a4119e5b178861e2786b92191f38971ab4df4ee83fc5f417922d52fec65f364736f6c634300081c0033",
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

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactor) Pause(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.contract.Transact(opts, "pause", token)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) Pause(token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.Pause(&_CrossMintableERC20Factory.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactorSession) Pause(token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.Pause(&_CrossMintableERC20Factory.TransactOpts, token)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactor) Unpause(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.contract.Transact(opts, "unpause", token)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactorySession) Unpause(token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.Unpause(&_CrossMintableERC20Factory.TransactOpts, token)
}

// Unpause is a paid mutator transaction binding the contract method 0x57b001f9.
//
// Solidity: function unpause(address token) returns()
func (_CrossMintableERC20Factory *CrossMintableERC20FactoryTransactorSession) Unpause(token common.Address) (*types.Transaction, error) {
	return _CrossMintableERC20Factory.Contract.Unpause(&_CrossMintableERC20Factory.TransactOpts, token)
}
