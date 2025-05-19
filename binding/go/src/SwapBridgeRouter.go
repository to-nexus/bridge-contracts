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

// SwapBridgeRouterMetaData contains all meta data concerning the SwapBridgeRouter contract.
var SwapBridgeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapRouter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crossToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"contractIBaseBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWAP_ROUTER\",\"outputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPath\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getPriceImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserveA\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserveB\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSourceTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getSwapBridgeIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"getSwapBridgeInCross\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getSwapBridgeOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bridgeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getSwapBridgeOutCross\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bridgeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"removeTokenPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"setTokenPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeETHForExactCrossTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeETHForExactTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeExactETHForCrossTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeExactETHForTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeExactTokensForCrossTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeExactTokensForETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeExactTokensForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeTokensForExactCrossTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeTokensForExactETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFeeMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBridgeTokensForExactTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"SwapBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"TokenPathRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"TokenPathSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeBridgeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeBridgeValueTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeExFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInsufficientReserve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInvalidBridgeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInvalidSourceToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeNetworkFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeOnlySwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgePairDoesNotExist\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"ee9a31a2": "BRIDGE()",
		"3b269e4e": "CROSS_CHAIN_ID()",
		"1e10ebdc": "CROSS_TOKEN()",
		"c6005893": "SWAP_ROUTER()",
		"ad5c4648": "WETH()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"4f973bac": "getPath(address)",
		"0a353a5e": "getPriceImpact(address[],uint256)",
		"d52bb6f4": "getReserves(address,address)",
		"92583a1a": "getSourceTokens()",
		"fa768893": "getSwapBridgeIn(uint256,uint256,address[])",
		"7f332b52": "getSwapBridgeInCross(address,uint256)",
		"cdd19b7e": "getSwapBridgeOut(uint256,uint256,address[])",
		"dd625509": "getSwapBridgeOutCross(address,uint256)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"ad615dec": "quote(uint256,uint256,uint256)",
		"cdbed967": "removeTokenPath(address)",
		"715018a6": "renounceOwnership()",
		"1fc810c4": "setTokenPath(address,address[])",
		"fd24b30b": "swapBridgeETHForExactCrossTokens(address,uint256,uint256,uint256,uint256)",
		"60f71c20": "swapBridgeETHForExactTokens(uint256,address,uint256,uint256,uint256,address[],uint256)",
		"7d173b7a": "swapBridgeExactETHForCrossTokens(address,uint256,uint256,uint256,uint256)",
		"c78cb094": "swapBridgeExactETHForTokens(uint256,address,uint256,uint256,uint256,address[],uint256)",
		"b379262d": "swapBridgeExactTokensForCrossTokens(address,address,uint256,uint256,uint256,uint256,uint256)",
		"af69d7fe": "swapBridgeExactTokensForETH(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"be835388": "swapBridgeExactTokensForTokens(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"02a827f2": "swapBridgeTokensForExactCrossTokens(address,address,uint256,uint256,uint256,uint256,uint256)",
		"a88c3f7c": "swapBridgeTokensForExactETH(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"bf557c55": "swapBridgeTokensForExactTokens(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x610120604052348015610010575f5ffd5b50604051614b3c380380614b3c83398101604081905261002f9161022f565b846001600160a01b03811661005d57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b610066816101c5565b505f805460ff60a01b191690556001600160a01b03851661009a576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b0384166100c1576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b0383166100e8576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b03821661010f576040516332746c8560e01b815260040160405180910390fd5b805f0361012f576040516395e19d4f60e01b815260040160405180910390fd5b6001600160a01b0380851660805283811660a081905290831660c05260e0829052604080516315ab88c960e31b8152905163ad5c4648916004808201926020929091908290030181865afa158015610189573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101ad9190610289565b6001600160a01b031661010052506102a99350505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b038116811461022a575f5ffd5b919050565b5f5f5f5f5f60a08688031215610243575f5ffd5b61024c86610214565b945061025a60208701610214565b935061026860408701610214565b925061027660608701610214565b9150608086015190509295509295909350565b5f60208284031215610299575f5ffd5b6102a282610214565b9392505050565b60805160a05160c05160e051610100516146ce61046e5f395f81816104fb01528181610e4401528181610fe9015281816110ab01528181611354015281816115730152818161183c015281816119be01528181611b4101528181611d5401528181611e810152818161231c0152818161244901528181612602015281816126c4015281816128370152818161325e015261349901525f81816103750152818161085c0152818161113401528181611246015281816117b70152818161224a015261274d01525f81816102df0152818161089801528181610c66015261117001525f818161021a015281816105bd015281816108f601528181610b76015281816112b3015281816114e301528181611f5801528181611f94015281816128ac0152818161293d01528181612a0701528181612abf01528181612b9b01528181612c1601528181612c7b0152612d0a01525f81816106cf01528181610e98015281816113a8015281816115c70152818161189001528181611a1201528181611b9501528181611da8015281816123700152818161248701528181612fe20152818161329c0152818161357001528181613602015281816136af015261371d01526146ce5ff3fe60806040526004361061020a575f3560e01c8063a88c3f7c11610113578063cdbed9671161009d578063dd6255091161006d578063dd6255091461069f578063ee9a31a2146106be578063f2fde38b146106f1578063fa76889314610710578063fd24b30b1461072f575f5ffd5b8063cdbed967146105f2578063cdd19b7e14610611578063d06ca61f14610641578063d52bb6f414610660575f5ffd5b8063b379262d116100e3578063b379262d1461054f578063be8353881461056e578063bf557c551461058d578063c6005893146105ac578063c78cb094146105df575f5ffd5b8063a88c3f7c146104cb578063ad5c4648146104ea578063ad615dec1461051d578063af69d7fe1461053c575f5ffd5b80635c975abb116101945780637f332b52116101645780637f332b52146104395780638456cb591461046857806385f8c2591461047c5780638da5cb5b1461049b57806392583a1a146104b7575f5ffd5b80635c975abb146103d757806360f71c20146103ff578063715018a6146104125780637d173b7a14610426575f5ffd5b80631f00ca74116101da5780631f00ca74146103195780631fc810c4146103455780633b269e4e146103645780633f4ba83a146103975780634f973bac146103ab575f5ffd5b806302a827f21461025e578063054d50d41461027d5780630a353a5e146102af5780631e10ebdc146102ce575f5ffd5b3661025a57336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610258576040516313e21f5560e31b815260040160405180910390fd5b005b5f5ffd5b348015610269575f5ffd5b50610258610278366004613d30565b610742565b348015610288575f5ffd5b5061029c610297366004613d8d565b6108d0565b6040519081526020015b60405180910390f35b3480156102ba575f5ffd5b5061029c6102c9366004613eb7565b610972565b3480156102d9575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102a6565b348015610324575f5ffd5b50610338610333366004613ef8565b610b5c565b6040516102a69190613f75565b348015610350575f5ffd5b5061025861035f366004613f87565b610bee565b34801561036f575f5ffd5b5061029c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103a2575f5ffd5b50610258610d66565b3480156103b6575f5ffd5b506103ca6103c5366004613fbd565b610d78565b6040516102a69190614011565b3480156103e2575f5ffd5b505f54600160a01b900460ff1660405190151581526020016102a6565b61025861040d366004614023565b610deb565b34801561041d575f5ffd5b50610258610fce565b61025861043436600461409f565b610fdf565b348015610444575f5ffd5b506104586104533660046140df565b6111a6565b6040516102a69493929190614109565b348015610473575f5ffd5b5061025861127d565b348015610487575f5ffd5b5061029c610496366004613d8d565b61128d565b3480156104a6575f5ffd5b505f546001600160a01b0316610301565b3480156104c2575f5ffd5b506103ca6112ea565b3480156104d6575f5ffd5b506102586104e5366004614137565b6112fb565b3480156104f5575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b348015610528575f5ffd5b5061029c610537366004613d8d565b6114bd565b61025861054a366004614137565b61151a565b34801561055a575f5ffd5b50610258610569366004613d30565b61169d565b348015610579575f5ffd5b50610258610588366004614137565b6117e3565b348015610598575f5ffd5b506102586105a7366004614137565b611965565b3480156105b7575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b6102586105ed366004614023565b611ae8565b3480156105fd575f5ffd5b5061025861060c366004613fbd565b611c6b565b34801561061c575f5ffd5b5061063061062b3660046141bc565b611cfd565b6040516102a6959493929190614207565b34801561064c575f5ffd5b5061033861065b366004613ef8565b611f3e565b34801561066b575f5ffd5b5061067f61067a36600461423d565b611f8f565b604080516001600160701b039384168152929091166020830152016102a6565b3480156106aa575f5ffd5b506106306106b93660046140df565b6121a9565b3480156106c9575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b3480156106fc575f5ffd5b5061025861070b366004613fbd565b612284565b34801561071b575f5ffd5b5061045861072a3660046141bc565b6122c6565b61025861073d36600461409f565b6125f8565b61074a61277a565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b94899490939291908301828280156107b557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610797575b505050505090506002815110156107df5760405163bb52eed160e01b815260040160405180910390fd5b6107e983836127c1565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561085057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610832575b505050505090506108897f00000000000000000000000000000000000000000000000000000000000000008b8a8c8b8b878c60036129bc565b5061089383612da9565b6108bc7f0000000000000000000000000000000000000000000000000000000000000000612da9565b5050506108c7612e5e565b50505050505050565b604051630153543560e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063054d50d4906064015b602060405180830381865afa158015610944573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109689190614274565b90505b9392505050565b5f600283511080610981575081155b1561098d57505f610b56565b5f6109988385611f3e565b90505f5b600185516109aa919061429f565b811015610b53575f8582815181106109c4576109c46142b2565b602002602001015190505f868360016109dd91906142c6565b815181106109ed576109ed6142b2565b602002602001015190505f5f610a038484611f8f565b915091505f826001600160701b0316118015610a2757505f816001600160701b0316115b610a445760405163a008318960e01b815260040160405180910390fd5b5f826001600160701b0316826001600160701b0316670de0b6b3a7640000610a6c91906142d9565b610a769190614304565b90505f878781518110610a8b57610a8b6142b2565b602002602001015190505f88886001610aa491906142c6565b81518110610ab457610ab46142b2565b602002602001015190505f82866001600160701b0316610ad491906142c6565b610ae7836001600160701b03881661429f565b610af990670de0b6b3a76400006142d9565b610b039190614304565b90505f81851115610b325784610b19838261429f565b610b25906127106142d9565b610b2f9190614304565b90505b610b3c818d6142c6565b9b50506001909801975061099c9650505050505050565b50505b92915050565b6040516307c0329d60e21b81526060906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631f00ca7490610bad9086908690600401614323565b5f60405180830381865afa158015610bc7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261096b919081019061433b565b610bf6612e75565b600281511015610c195760405163bb52eed160e01b815260040160405180910390fd5b816001600160a01b0316815f81518110610c3557610c356142b2565b60200260200101516001600160a01b031614610c645760405163bb52eed160e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168160018351610c9d919061429f565b81518110610cad57610cad6142b2565b60200260200101516001600160a01b031614610cdc5760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0382165f9081526001602090815260409091208251610d0492840190613c8e565b50610d10600283612ea1565b610d2157610d1f600283612ec2565b505b816001600160a01b03167f654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec82604051610d5a9190614011565b60405180910390a25050565b610d6e612e75565b610d76612ed6565b565b6001600160a01b0381165f90815260016020908152604091829020805483518184028101840190945280845260609392830182828015610ddf57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610dc1575b50505050509050919050565b610df361277a565b8682600281511015610e185760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351610e28919061429f565b81518110610e3857610e386142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603610e7f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590610ecf90879086906004016143cb565b60e060405180830381865afa158015610eea573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f0e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614610f45576040516337088ab560e11b815260040160405180910390fd5b8534610f6a825f81518110610f5c57610f5c6142b2565b6020026020010151826127c1565b610f7c8d8d348e8e8e8e8e60046129bc565b610f9e825f81518110610f9157610f916142b2565b6020026020010151612da9565b610fc08260018451610fb0919061429f565b81518110610f9157610f916142b2565b5050505050506108c7612e5e565b610fd6612e75565b610d765f612f2a565b610fe761277a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0381165f908152600160209081526040808320805482518185028101850190935280835234949383018282801561106d57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161104f575b505050505090506002815110156110975760405163bb52eed160e01b815260040160405180910390fd5b6110a183836127c1565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561112857602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161110a575b505050505090506111617f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c60016129bc565b5061116b83612da9565b6111947f0000000000000000000000000000000000000000000000000000000000000000612da9565b50505061119f612e5e565b5050505050565b60605f80806111b6600287612ea1565b6111d3576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0386165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561123a57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161121c575b5050505050905061126c7f000000000000000000000000000000000000000000000000000000000000000087836122c6565b929a91995097509095509350505050565b611285612e75565b610d76612f79565b6040516385f8c25960e01b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906385f8c25990606401610929565b60606112f66002612fbb565b905090565b61130361277a565b87826002815110156113285760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611338919061429f565b81518110611348576113486142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03160361138f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906113df90879086906004016143cb565b60e060405180830381865afa1580156113fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611455576040516337088ab560e11b815260040160405180910390fd5b858961146c825f81518110610f5c57610f5c6142b2565b61147e8e8e8d8f8e8e8e8e60056129bc565b611493825f81518110610f9157610f916142b2565b6114a58260018451610fb0919061429f565b5050505050506114b3612e5e565b5050505050505050565b604051632b58577b60e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ad615dec90606401610929565b61152261277a565b87826002815110156115475760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611557919061429f565b81518110611567576115676142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036115ae575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906115fe90879086906004016143cb565b60e060405180830381865afa158015611619573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061163d91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611674576040516337088ab560e11b815260040160405180910390fd5b858a61168b825f81518110610f5c57610f5c6142b2565b61147e8e8e8e8e8e8e8e8e60026129bc565b6116a561277a565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b948a94909392919083018282801561171057602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116116f2575b5050505050905060028151101561173a5760405163bb52eed160e01b815260040160405180910390fd5b61174483836127c1565b6001600160a01b038a165f908152600160209081526040808320805482518185028101850190935280835291929091908301828280156117ab57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161178d575b505050505090506108897f00000000000000000000000000000000000000000000000000000000000000008b8b8b8b8b878c5f6129bc565b6117eb61277a565b87826002815110156118105760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611820919061429f565b81518110611830576118306142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611877575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906118c790879086906004016143cb565b60e060405180830381865afa1580156118e2573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061190691906143f6565b9050816001600160a01b0316815f01516001600160a01b03161461193d576040516337088ab560e11b815260040160405180910390fd5b858a611954825f81518110610f5c57610f5c6142b2565b61147e8e8e8e8e8e8e8e8e5f6129bc565b61196d61277a565b87826002815110156119925760405163bb52eed160e01b815260040160405180910390fd5b5f81600183516119a2919061429f565b815181106119b2576119b26142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036119f9575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611a4990879086906004016143cb565b60e060405180830381865afa158015611a64573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a8891906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611abf576040516337088ab560e11b815260040160405180910390fd5b8589611ad6825f81518110610f5c57610f5c6142b2565b61147e8e8e8d8f8e8e8e8e60036129bc565b611af061277a565b8682600281511015611b155760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611b25919061429f565b81518110611b3557611b356142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611b7c575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611bcc90879086906004016143cb565b60e060405180830381865afa158015611be7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c0b91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611c42576040516337088ab560e11b815260040160405180910390fd5b8534611c59825f81518110610f5c57610f5c6142b2565b610f7c8d8d348e8e8e8e8e60016129bc565b611c73612e75565b611c7e600282612ea1565b611c9b576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0381165f908152600160205260408120611cbb91613cf1565b611cc6600282612fc7565b506040516001600160a01b038216907fcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193905f90a250565b60605f5f5f5f8786600281511015611d285760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611d38919061429f565b81518110611d4857611d486142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611d8f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611ddf90879086906004016143cb565b60e060405180830381865afa158015611dfa573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e1e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611e55576040516337088ab560e11b815260040160405180910390fd5b5f8a60018c51611e65919061429f565b81518110611e7557611e756142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611ebc575060015b611ec68c8c611f3e565b99505f611efa8e838d60018f51611edd919061429f565b81518110611eed57611eed6142b2565b6020026020010151612fdb565b919c509a509850905080611f21576040516326b56d9d60e21b815260040160405180910390fd5b611f2b8c8e610972565b9650505050505050939792965093509350565b60405163d06ca61f60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d06ca61f90610bad9086908690600401614323565b5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c45a01556040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fee573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120129190614479565b60405163e6a4390560e01b81526001600160a01b0387811660048301528681166024830152919091169063e6a4390590604401602060405180830381865afa158015612060573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120849190614479565b90506001600160a01b0381166120ad57604051630240c4d360e11b815260040160405180910390fd5b5f8190505f5f826001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa1580156120ef573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061211391906144aa565b5091509150826001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015612154573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121789190614479565b6001600160a01b0316886001600160a01b03161461219757808261219a565b81815b90999098509650505050505050565b60605f8080806121ba600288612ea1565b6121d7576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0387165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561223e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612220575b505050505090506122707f00000000000000000000000000000000000000000000000000000000000000008883611cfd565b939c929b5090995097509095509350505050565b61228c612e75565b6001600160a01b0381166122ba57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6122c381612f2a565b50565b60605f5f5f86856002815110156122f05760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351612300919061429f565b81518110612310576123106142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612357575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906123a790879086906004016143cb565b60e060405180830381865afa1580156123c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123e691906143f6565b9050816001600160a01b0316815f01516001600160a01b03161461241d576040516337088ab560e11b815260040160405180910390fd5b5f8960018b5161242d919061429f565b8151811061243d5761243d6142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612484575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125059190614479565b6040516337dba1f760e21b8152600481018f90526001600160a01b038481166024830152604482018f9052919091169063df6e87dc90606401606060405180830381865afa158015612559573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061257d91906144f6565b909a5098509050808c116125a4576040516326b56d9d60e21b815260040160405180910390fd5b6125c2886125b28b8f6142c6565b6125bc91906142c6565b8c610b5c565b99506125e78b8b5f815181106125da576125da6142b2565b6020026020010151610972565b965050505050505093509350935093565b61260061277a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0381165f908152600160209081526040808320805482518185028101850190935280835234949383018282801561268657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612668575b505050505090506002815110156126b05760405163bb52eed160e01b815260040160405180910390fd5b6126ba83836127c1565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561274157602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612723575b505050505090506111617f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c60046129bc565b5f5160206146795f395f51905f525c156127a757604051633ee5aeb560e01b815260040160405180910390fd5b610d7660015f5160206146795f395f51905f525b906131bd565b80156127cd5780612833565b6040516370a0823160e01b81523360048201526001600160a01b038316906370a0823190602401602060405180830381865afa15801561280f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128339190614274565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480156128765750803410155b61288f5761288f6001600160a01b0383163330846131c4565b604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919084169063dd62ed3e90604401602060405180830381865afa1580156128fc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129209190614274565b10156129b85760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015283169063095ea7b3906044016020604051808303815f875af1158015612992573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129b69190614521565b505b5050565b5f806060818460058111156129d3576129d361453a565b03612a8d57866129e3898b6142c6565b6129ed91906142c6565b6040516338ed173960e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906338ed173990612a44908d9087908b9030908c9060040161454e565b5f604051808303815f875af1158015612a5f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612a86919081019061433b565b9050612d8c565b6001846005811115612aa157612aa161453a565b03612b535786612ab1898b6142c6565b612abb91906142c6565b92507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff36ab58b8589308a6040518663ffffffff1660e01b8152600401612b109493929190614589565b5f6040518083038185885af1158015612b2b573d5f5f3e3d5ffd5b50505050506040513d5f823e601f3d908101601f19168201604052612a86919081019061433b565b6002846005811115612b6757612b6761453a565b03612bd85786612b77898b6142c6565b612b8191906142c6565b6040516318cbafe560e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906318cbafe590612a44908d9087908b9030908c9060040161454e565b6003846005811115612bec57612bec61453a565b03612c5357612bfc8c8a88613231565b604051634401edf760e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638803dbee90612a449085908e908b9030908c9060040161454e565b6004846005811115612c6757612c6761453a565b03612ccc57612c778c8a88613231565b91507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fb3bdb418b8489308a6040518663ffffffff1660e01b8152600401612b109493929190614589565b6005846005811115612ce057612ce061453a565b03612d8c57612cf08c8a88613231565b604051632512eca560e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634a25d94a90612d479085908e908b9030908c9060040161454e565b5f604051808303815f875af1158015612d62573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612d89919081019061433b565b90505b612d9b8c8c8b8b8b8b8761344a565b505050505050505050505050565b4715612dda5760405133904780156108fc02915f818181858888f19350505050158015612dd8573d5f5f3e3d5ffd5b505b6040516370a0823160e01b81523060048201525f906001600160a01b038316906370a0823190602401602060405180830381865afa158015612e1e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e429190614274565b905080156129b8576129b86001600160a01b0383163383613880565b610d765f5f5160206146795f395f51905f526127bb565b5f546001600160a01b03163314610d765760405163118cdaa760e01b81523360048201526024016122b1565b6001600160a01b0381165f908152600183016020526040812054151561096b565b5f61096b836001600160a01b0384166138b1565b612ede6138fd565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b612f81613926565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612f0d3390565b60605f61096b83613950565b5f61096b836001600160a01b0384166139a8565b5f5f5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561303c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130609190614479565b90505f5f826001600160a01b0316636332aec68b8b6040518363ffffffff1660e01b81526004016130929291906143cb565b606060405180830381865afa1580156130ad573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130d191906144f6565b90965090925090508488116130f4575f5f5f5f96509650965096505050506131b4565b5f836001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa158015613131573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131559190614274565b90505f613162878b61429f565b905061317a8261317285826142c6565b839190613a8b565b9750613187888484613a8b565b955083881180156131ac57508561319e888a6142c6565b6131a891906142c6565b8a10155b985050505050505b93509350935093565b80825d5050565b6040516001600160a01b03848116602483015283811660448301526064820183905261322b9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613b41565b50505050565b5f5f8260018451613242919061429f565b81518110613252576132526142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603613299575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061331a9190614479565b90505f5f5f836001600160a01b0316636332aec68a876040518363ffffffff1660e01b815260040161334d9291906143cb565b606060405180830381865afa158015613368573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061338c91906144f6565b9194509092509050828810156133b5576040516326b56d9d60e21b815260040160405180910390fd5b5f846001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133f2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134169190614274565b90505f6134248a8584613a8b565b905082613431828c6142c6565b61343b91906142c6565b9b9a5050505050505050505050565b5f61346c60405180606001604052805f81526020015f81526020015f81525090565b5f5f856001875161347d919061429f565b8151811061348d5761348d6142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03161493505f84613541576040516370a0823160e01b815230600482015282906001600160a01b038216906370a0823190602401602060405180830381865afa158015613518573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061353c9190614274565b613545565b6001475b9093509050841580156135e65750604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919085169063dd62ed3e90604401602060405180830381865afa1580156135c0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135e49190614274565b105b1561367d5760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015284169063095ea7b3906044016020604051808303815f875af1158015613657573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061367b9190614521565b505b61368b8c84838d8d8d613bad565b6040516315cd127f60e31b8152600481018e90529094505f92506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915063ae6893f890602401602060405180830381865afa1580156136f5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137199190614274565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635fd262de85613755575f613758565b84515b8d858e885f015189602001518a604001515f6001600160401b0381111561378157613781613db6565b6040519080825280601f01601f1916602001820160405280156137ab576020820181803683370190505b506040518963ffffffff1660e01b81526004016137ce97969594939291906145bd565b60206040518083038185885af11580156137ea573d5f5f3e3d5ffd5b50505050506040513d601f19601f8201168201806040525081019061380f9190614521565b61382c57604051632fba5fef60e21b815260040160405180910390fd5b336001600160a01b0316818c7f72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca8d8a8a60405161386b9392919061462f565b60405180910390a45050505050505050505050565b6040516001600160a01b038381166024830152604482018390526129b691859182169063a9059cbb906064016131f9565b5f8181526001830160205260408120546138f657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b56565b505f610b56565b5f54600160a01b900460ff16610d7657604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff1615610d765760405163d93c066560e01b815260040160405180910390fd5b6060815f01805480602002602001604051908101604052809291908181526020018280548015610ddf57602002820191905f5260205f20905b8154815260200190600101908083116139895750505050509050919050565b5f8181526001830160205260408120548015613a82575f6139ca60018361429f565b85549091505f906139dd9060019061429f565b9050808214613a3c575f865f0182815481106139fb576139fb6142b2565b905f5260205f200154905080875f018481548110613a1b57613a1b6142b2565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080613a4d57613a4d614664565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610b56565b5f915050610b56565b5f838302815f1985870982811083820303915050805f03613abf57838281613ab557613ab56142f0565b049250505061096b565b808411613ad657613ad66003851502601118613c7d565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f5f60205f8451602086015f885af180613b60576040513d5f823e3d81fd5b50505f513d91508115613b77578060011415613b84565b6001600160a01b0384163b155b1561322b57604051635274afe760e01b81526001600160a01b03851660048201526024016122b1565b613bce60405180606001604052805f81526020015f81526020015f81525090565b5f613bda888888612fdb565b604086015260208501528352905080613c06576040516326b56d9d60e21b815260040160405180910390fd5b8151851115613c2857604051637c3bec0960e11b815260040160405180910390fd5b8382602001511115613c4d5760405163de7539a960e01b815260040160405180910390fd5b8282604001511115613c725760405163cce42f3560e01b815260040160405180910390fd5b509695505050505050565b634e487b715f52806020526024601cfd5b828054828255905f5260205f20908101928215613ce1579160200282015b82811115613ce157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613cac565b50613ced929150613d08565b5090565b5080545f8255905f5260205f20908101906122c391905b5b80821115613ced575f8155600101613d09565b6001600160a01b03811681146122c3575f5ffd5b5f5f5f5f5f5f5f60e0888a031215613d46575f5ffd5b8735613d5181613d1c565b96506020880135613d6181613d1c565b96999698505050506040850135946060810135946080820135945060a0820135935060c0909101359150565b5f5f5f60608486031215613d9f575f5ffd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715613dec57613dec613db6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613e1a57613e1a613db6565b604052919050565b5f6001600160401b03821115613e3a57613e3a613db6565b5060051b60200190565b5f82601f830112613e53575f5ffd5b8135613e66613e6182613e22565b613df2565b8082825260208201915060208360051b860101925085831115613e87575f5ffd5b602085015b83811015613ead578035613e9f81613d1c565b835260209283019201613e8c565b5095945050505050565b5f5f60408385031215613ec8575f5ffd5b82356001600160401b03811115613edd575f5ffd5b613ee985828601613e44565b95602094909401359450505050565b5f5f60408385031215613f09575f5ffd5b8235915060208301356001600160401b03811115613f25575f5ffd5b613f3185828601613e44565b9150509250929050565b5f8151808452602084019350602083015f5b82811015613f6b578151865260209586019590910190600101613f4d565b5093949350505050565b602081525f61096b6020830184613f3b565b5f5f60408385031215613f98575f5ffd5b8235613fa381613d1c565b915060208301356001600160401b03811115613f25575f5ffd5b5f60208284031215613fcd575f5ffd5b813561096b81613d1c565b5f8151808452602084019350602083015f5b82811015613f6b5781516001600160a01b0316865260209586019590910190600101613fea565b602081525f61096b6020830184613fd8565b5f5f5f5f5f5f5f60e0888a031215614039575f5ffd5b87359650602088013561404b81613d1c565b955060408801359450606088013593506080880135925060a08801356001600160401b0381111561407a575f5ffd5b6140868a828b01613e44565b979a969950949793969295929450505060c09091013590565b5f5f5f5f5f60a086880312156140b3575f5ffd5b85356140be81613d1c565b97602087013597506040870135966060810135965060800135945092505050565b5f5f604083850312156140f0575f5ffd5b82356140fb81613d1c565b946020939093013593505050565b608081525f61411b6080830187613f3b565b6020830195909552506040810192909252606090910152919050565b5f5f5f5f5f5f5f5f610100898b03121561414f575f5ffd5b88359750602089013561416181613d1c565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b03811115614197575f5ffd5b6141a38b828c01613e44565b989b979a50959894979396929550929360e00135925050565b5f5f5f606084860312156141ce575f5ffd5b833592506020840135915060408401356001600160401b038111156141f1575f5ffd5b6141fd86828701613e44565b9150509250925092565b60a081525f61421960a0830188613f3b565b90508560208301528460408301528360608301528260808301529695505050505050565b5f5f6040838503121561424e575f5ffd5b823561425981613d1c565b9150602083013561426981613d1c565b809150509250929050565b5f60208284031215614284575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610b5657610b5661428b565b634e487b7160e01b5f52603260045260245ffd5b80820180821115610b5657610b5661428b565b8082028115828204841417610b5657610b5661428b565b634e487b7160e01b5f52601260045260245ffd5b5f8261431e57634e487b7160e01b5f52601260045260245ffd5b500490565b828152604060208201525f6109686040830184613fd8565b5f6020828403121561434b575f5ffd5b81516001600160401b03811115614360575f5ffd5b8201601f81018413614370575f5ffd5b805161437e613e6182613e22565b8082825260208201915060208360051b85010192508683111561439f575f5ffd5b6020840193505b828410156143c15783518252602093840193909101906143a6565b9695505050505050565b9182526001600160a01b0316602082015260400190565b805180151581146143f1575f5ffd5b919050565b5f60e0828403128015614407575f5ffd5b50614410613dca565b825161441b81613d1c565b8152602083015161442b81613d1c565b602082015261443c604084016143e2565b604082015261444d606084016143e2565b60608201526080838101519082015260a0808401519082015260c0928301519281019290925250919050565b5f60208284031215614489575f5ffd5b815161096b81613d1c565b80516001600160701b03811681146143f1575f5ffd5b5f5f5f606084860312156144bc575f5ffd5b6144c584614494565b92506144d360208501614494565b9150604084015163ffffffff811681146144eb575f5ffd5b809150509250925092565b5f5f5f60608486031215614508575f5ffd5b5050815160208301516040909301519094929350919050565b5f60208284031215614531575f5ffd5b61096b826143e2565b634e487b7160e01b5f52602160045260245ffd5b85815284602082015260a060408201525f61456c60a0830186613fd8565b6001600160a01b0394909416606083015250608001529392505050565b848152608060208201525f6145a16080830186613fd8565b6001600160a01b03949094166040830152506060015292915050565b87815260018060a01b038716602082015260018060a01b03861660408201528460608201528360808201528260a082015260e060c08201525f82518060e0840152806020850161010085015e5f6101008285010152610100601f19601f83011684010191505098975050505050505050565b6001600160a01b03841681526060602082018190525f9061465290830185613fd8565b82810360408401526143c18185613f3b565b634e487b7160e01b5f52603160045260245ffdfe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122001b235830d738541fc14a42aeb2ee1f027559c74608910dc8f054dee77914f4864736f6c634300081c0033",
}

// SwapBridgeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapBridgeRouterMetaData.ABI instead.
var SwapBridgeRouterABI = SwapBridgeRouterMetaData.ABI

// SwapBridgeRouterBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const SwapBridgeRouterBinRuntime = "60806040526004361061020a575f3560e01c8063a88c3f7c11610113578063cdbed9671161009d578063dd6255091161006d578063dd6255091461069f578063ee9a31a2146106be578063f2fde38b146106f1578063fa76889314610710578063fd24b30b1461072f575f5ffd5b8063cdbed967146105f2578063cdd19b7e14610611578063d06ca61f14610641578063d52bb6f414610660575f5ffd5b8063b379262d116100e3578063b379262d1461054f578063be8353881461056e578063bf557c551461058d578063c6005893146105ac578063c78cb094146105df575f5ffd5b8063a88c3f7c146104cb578063ad5c4648146104ea578063ad615dec1461051d578063af69d7fe1461053c575f5ffd5b80635c975abb116101945780637f332b52116101645780637f332b52146104395780638456cb591461046857806385f8c2591461047c5780638da5cb5b1461049b57806392583a1a146104b7575f5ffd5b80635c975abb146103d757806360f71c20146103ff578063715018a6146104125780637d173b7a14610426575f5ffd5b80631f00ca74116101da5780631f00ca74146103195780631fc810c4146103455780633b269e4e146103645780633f4ba83a146103975780634f973bac146103ab575f5ffd5b806302a827f21461025e578063054d50d41461027d5780630a353a5e146102af5780631e10ebdc146102ce575f5ffd5b3661025a57336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610258576040516313e21f5560e31b815260040160405180910390fd5b005b5f5ffd5b348015610269575f5ffd5b50610258610278366004613d30565b610742565b348015610288575f5ffd5b5061029c610297366004613d8d565b6108d0565b6040519081526020015b60405180910390f35b3480156102ba575f5ffd5b5061029c6102c9366004613eb7565b610972565b3480156102d9575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102a6565b348015610324575f5ffd5b50610338610333366004613ef8565b610b5c565b6040516102a69190613f75565b348015610350575f5ffd5b5061025861035f366004613f87565b610bee565b34801561036f575f5ffd5b5061029c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103a2575f5ffd5b50610258610d66565b3480156103b6575f5ffd5b506103ca6103c5366004613fbd565b610d78565b6040516102a69190614011565b3480156103e2575f5ffd5b505f54600160a01b900460ff1660405190151581526020016102a6565b61025861040d366004614023565b610deb565b34801561041d575f5ffd5b50610258610fce565b61025861043436600461409f565b610fdf565b348015610444575f5ffd5b506104586104533660046140df565b6111a6565b6040516102a69493929190614109565b348015610473575f5ffd5b5061025861127d565b348015610487575f5ffd5b5061029c610496366004613d8d565b61128d565b3480156104a6575f5ffd5b505f546001600160a01b0316610301565b3480156104c2575f5ffd5b506103ca6112ea565b3480156104d6575f5ffd5b506102586104e5366004614137565b6112fb565b3480156104f5575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b348015610528575f5ffd5b5061029c610537366004613d8d565b6114bd565b61025861054a366004614137565b61151a565b34801561055a575f5ffd5b50610258610569366004613d30565b61169d565b348015610579575f5ffd5b50610258610588366004614137565b6117e3565b348015610598575f5ffd5b506102586105a7366004614137565b611965565b3480156105b7575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b6102586105ed366004614023565b611ae8565b3480156105fd575f5ffd5b5061025861060c366004613fbd565b611c6b565b34801561061c575f5ffd5b5061063061062b3660046141bc565b611cfd565b6040516102a6959493929190614207565b34801561064c575f5ffd5b5061033861065b366004613ef8565b611f3e565b34801561066b575f5ffd5b5061067f61067a36600461423d565b611f8f565b604080516001600160701b039384168152929091166020830152016102a6565b3480156106aa575f5ffd5b506106306106b93660046140df565b6121a9565b3480156106c9575f5ffd5b506103017f000000000000000000000000000000000000000000000000000000000000000081565b3480156106fc575f5ffd5b5061025861070b366004613fbd565b612284565b34801561071b575f5ffd5b5061045861072a3660046141bc565b6122c6565b61025861073d36600461409f565b6125f8565b61074a61277a565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b94899490939291908301828280156107b557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610797575b505050505090506002815110156107df5760405163bb52eed160e01b815260040160405180910390fd5b6107e983836127c1565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561085057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610832575b505050505090506108897f00000000000000000000000000000000000000000000000000000000000000008b8a8c8b8b878c60036129bc565b5061089383612da9565b6108bc7f0000000000000000000000000000000000000000000000000000000000000000612da9565b5050506108c7612e5e565b50505050505050565b604051630153543560e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063054d50d4906064015b602060405180830381865afa158015610944573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109689190614274565b90505b9392505050565b5f600283511080610981575081155b1561098d57505f610b56565b5f6109988385611f3e565b90505f5b600185516109aa919061429f565b811015610b53575f8582815181106109c4576109c46142b2565b602002602001015190505f868360016109dd91906142c6565b815181106109ed576109ed6142b2565b602002602001015190505f5f610a038484611f8f565b915091505f826001600160701b0316118015610a2757505f816001600160701b0316115b610a445760405163a008318960e01b815260040160405180910390fd5b5f826001600160701b0316826001600160701b0316670de0b6b3a7640000610a6c91906142d9565b610a769190614304565b90505f878781518110610a8b57610a8b6142b2565b602002602001015190505f88886001610aa491906142c6565b81518110610ab457610ab46142b2565b602002602001015190505f82866001600160701b0316610ad491906142c6565b610ae7836001600160701b03881661429f565b610af990670de0b6b3a76400006142d9565b610b039190614304565b90505f81851115610b325784610b19838261429f565b610b25906127106142d9565b610b2f9190614304565b90505b610b3c818d6142c6565b9b50506001909801975061099c9650505050505050565b50505b92915050565b6040516307c0329d60e21b81526060906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631f00ca7490610bad9086908690600401614323565b5f60405180830381865afa158015610bc7573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261096b919081019061433b565b610bf6612e75565b600281511015610c195760405163bb52eed160e01b815260040160405180910390fd5b816001600160a01b0316815f81518110610c3557610c356142b2565b60200260200101516001600160a01b031614610c645760405163bb52eed160e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168160018351610c9d919061429f565b81518110610cad57610cad6142b2565b60200260200101516001600160a01b031614610cdc5760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0382165f9081526001602090815260409091208251610d0492840190613c8e565b50610d10600283612ea1565b610d2157610d1f600283612ec2565b505b816001600160a01b03167f654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec82604051610d5a9190614011565b60405180910390a25050565b610d6e612e75565b610d76612ed6565b565b6001600160a01b0381165f90815260016020908152604091829020805483518184028101840190945280845260609392830182828015610ddf57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610dc1575b50505050509050919050565b610df361277a565b8682600281511015610e185760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351610e28919061429f565b81518110610e3857610e386142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603610e7f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590610ecf90879086906004016143cb565b60e060405180830381865afa158015610eea573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f0e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614610f45576040516337088ab560e11b815260040160405180910390fd5b8534610f6a825f81518110610f5c57610f5c6142b2565b6020026020010151826127c1565b610f7c8d8d348e8e8e8e8e60046129bc565b610f9e825f81518110610f9157610f916142b2565b6020026020010151612da9565b610fc08260018451610fb0919061429f565b81518110610f9157610f916142b2565b5050505050506108c7612e5e565b610fd6612e75565b610d765f612f2a565b610fe761277a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0381165f908152600160209081526040808320805482518185028101850190935280835234949383018282801561106d57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161104f575b505050505090506002815110156110975760405163bb52eed160e01b815260040160405180910390fd5b6110a183836127c1565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561112857602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161110a575b505050505090506111617f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c60016129bc565b5061116b83612da9565b6111947f0000000000000000000000000000000000000000000000000000000000000000612da9565b50505061119f612e5e565b5050505050565b60605f80806111b6600287612ea1565b6111d3576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0386165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561123a57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161121c575b5050505050905061126c7f000000000000000000000000000000000000000000000000000000000000000087836122c6565b929a91995097509095509350505050565b611285612e75565b610d76612f79565b6040516385f8c25960e01b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906385f8c25990606401610929565b60606112f66002612fbb565b905090565b61130361277a565b87826002815110156113285760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611338919061429f565b81518110611348576113486142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03160361138f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906113df90879086906004016143cb565b60e060405180830381865afa1580156113fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611455576040516337088ab560e11b815260040160405180910390fd5b858961146c825f81518110610f5c57610f5c6142b2565b61147e8e8e8d8f8e8e8e8e60056129bc565b611493825f81518110610f9157610f916142b2565b6114a58260018451610fb0919061429f565b5050505050506114b3612e5e565b5050505050505050565b604051632b58577b60e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ad615dec90606401610929565b61152261277a565b87826002815110156115475760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611557919061429f565b81518110611567576115676142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036115ae575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906115fe90879086906004016143cb565b60e060405180830381865afa158015611619573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061163d91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611674576040516337088ab560e11b815260040160405180910390fd5b858a61168b825f81518110610f5c57610f5c6142b2565b61147e8e8e8e8e8e8e8e8e60026129bc565b6116a561277a565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b948a94909392919083018282801561171057602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116116f2575b5050505050905060028151101561173a5760405163bb52eed160e01b815260040160405180910390fd5b61174483836127c1565b6001600160a01b038a165f908152600160209081526040808320805482518185028101850190935280835291929091908301828280156117ab57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161178d575b505050505090506108897f00000000000000000000000000000000000000000000000000000000000000008b8b8b8b8b878c5f6129bc565b6117eb61277a565b87826002815110156118105760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611820919061429f565b81518110611830576118306142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611877575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906118c790879086906004016143cb565b60e060405180830381865afa1580156118e2573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061190691906143f6565b9050816001600160a01b0316815f01516001600160a01b03161461193d576040516337088ab560e11b815260040160405180910390fd5b858a611954825f81518110610f5c57610f5c6142b2565b61147e8e8e8e8e8e8e8e8e5f6129bc565b61196d61277a565b87826002815110156119925760405163bb52eed160e01b815260040160405180910390fd5b5f81600183516119a2919061429f565b815181106119b2576119b26142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036119f9575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611a4990879086906004016143cb565b60e060405180830381865afa158015611a64573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a8891906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611abf576040516337088ab560e11b815260040160405180910390fd5b8589611ad6825f81518110610f5c57610f5c6142b2565b61147e8e8e8d8f8e8e8e8e60036129bc565b611af061277a565b8682600281511015611b155760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611b25919061429f565b81518110611b3557611b356142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611b7c575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611bcc90879086906004016143cb565b60e060405180830381865afa158015611be7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c0b91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611c42576040516337088ab560e11b815260040160405180910390fd5b8534611c59825f81518110610f5c57610f5c6142b2565b610f7c8d8d348e8e8e8e8e60016129bc565b611c73612e75565b611c7e600282612ea1565b611c9b576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0381165f908152600160205260408120611cbb91613cf1565b611cc6600282612fc7565b506040516001600160a01b038216907fcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193905f90a250565b60605f5f5f5f8786600281511015611d285760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611d38919061429f565b81518110611d4857611d486142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611d8f575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611ddf90879086906004016143cb565b60e060405180830381865afa158015611dfa573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e1e91906143f6565b9050816001600160a01b0316815f01516001600160a01b031614611e55576040516337088ab560e11b815260040160405180910390fd5b5f8a60018c51611e65919061429f565b81518110611e7557611e756142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603611ebc575060015b611ec68c8c611f3e565b99505f611efa8e838d60018f51611edd919061429f565b81518110611eed57611eed6142b2565b6020026020010151612fdb565b919c509a509850905080611f21576040516326b56d9d60e21b815260040160405180910390fd5b611f2b8c8e610972565b9650505050505050939792965093509350565b60405163d06ca61f60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d06ca61f90610bad9086908690600401614323565b5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c45a01556040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fee573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120129190614479565b60405163e6a4390560e01b81526001600160a01b0387811660048301528681166024830152919091169063e6a4390590604401602060405180830381865afa158015612060573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120849190614479565b90506001600160a01b0381166120ad57604051630240c4d360e11b815260040160405180910390fd5b5f8190505f5f826001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa1580156120ef573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061211391906144aa565b5091509150826001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa158015612154573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121789190614479565b6001600160a01b0316886001600160a01b03161461219757808261219a565b81815b90999098509650505050505050565b60605f8080806121ba600288612ea1565b6121d7576040516360c1ec5f60e01b815260040160405180910390fd5b6001600160a01b0387165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561223e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612220575b505050505090506122707f00000000000000000000000000000000000000000000000000000000000000008883611cfd565b939c929b5090995097509095509350505050565b61228c612e75565b6001600160a01b0381166122ba57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6122c381612f2a565b50565b60605f5f5f86856002815110156122f05760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351612300919061429f565b81518110612310576123106142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612357575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906123a790879086906004016143cb565b60e060405180830381865afa1580156123c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123e691906143f6565b9050816001600160a01b0316815f01516001600160a01b03161461241d576040516337088ab560e11b815260040160405180910390fd5b5f8960018b5161242d919061429f565b8151811061243d5761243d6142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603612484575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125059190614479565b6040516337dba1f760e21b8152600481018f90526001600160a01b038481166024830152604482018f9052919091169063df6e87dc90606401606060405180830381865afa158015612559573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061257d91906144f6565b909a5098509050808c116125a4576040516326b56d9d60e21b815260040160405180910390fd5b6125c2886125b28b8f6142c6565b6125bc91906142c6565b8c610b5c565b99506125e78b8b5f815181106125da576125da6142b2565b6020026020010151610972565b965050505050505093509350935093565b61260061277a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0381165f908152600160209081526040808320805482518185028101850190935280835234949383018282801561268657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612668575b505050505090506002815110156126b05760405163bb52eed160e01b815260040160405180910390fd5b6126ba83836127c1565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561274157602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612723575b505050505090506111617f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c60046129bc565b5f5160206146795f395f51905f525c156127a757604051633ee5aeb560e01b815260040160405180910390fd5b610d7660015f5160206146795f395f51905f525b906131bd565b80156127cd5780612833565b6040516370a0823160e01b81523360048201526001600160a01b038316906370a0823190602401602060405180830381865afa15801561280f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128339190614274565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480156128765750803410155b61288f5761288f6001600160a01b0383163330846131c4565b604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919084169063dd62ed3e90604401602060405180830381865afa1580156128fc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129209190614274565b10156129b85760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015283169063095ea7b3906044016020604051808303815f875af1158015612992573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129b69190614521565b505b5050565b5f806060818460058111156129d3576129d361453a565b03612a8d57866129e3898b6142c6565b6129ed91906142c6565b6040516338ed173960e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906338ed173990612a44908d9087908b9030908c9060040161454e565b5f604051808303815f875af1158015612a5f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612a86919081019061433b565b9050612d8c565b6001846005811115612aa157612aa161453a565b03612b535786612ab1898b6142c6565b612abb91906142c6565b92507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff36ab58b8589308a6040518663ffffffff1660e01b8152600401612b109493929190614589565b5f6040518083038185885af1158015612b2b573d5f5f3e3d5ffd5b50505050506040513d5f823e601f3d908101601f19168201604052612a86919081019061433b565b6002846005811115612b6757612b6761453a565b03612bd85786612b77898b6142c6565b612b8191906142c6565b6040516318cbafe560e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906318cbafe590612a44908d9087908b9030908c9060040161454e565b6003846005811115612bec57612bec61453a565b03612c5357612bfc8c8a88613231565b604051634401edf760e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638803dbee90612a449085908e908b9030908c9060040161454e565b6004846005811115612c6757612c6761453a565b03612ccc57612c778c8a88613231565b91507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fb3bdb418b8489308a6040518663ffffffff1660e01b8152600401612b109493929190614589565b6005846005811115612ce057612ce061453a565b03612d8c57612cf08c8a88613231565b604051632512eca560e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634a25d94a90612d479085908e908b9030908c9060040161454e565b5f604051808303815f875af1158015612d62573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612d89919081019061433b565b90505b612d9b8c8c8b8b8b8b8761344a565b505050505050505050505050565b4715612dda5760405133904780156108fc02915f818181858888f19350505050158015612dd8573d5f5f3e3d5ffd5b505b6040516370a0823160e01b81523060048201525f906001600160a01b038316906370a0823190602401602060405180830381865afa158015612e1e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e429190614274565b905080156129b8576129b86001600160a01b0383163383613880565b610d765f5f5160206146795f395f51905f526127bb565b5f546001600160a01b03163314610d765760405163118cdaa760e01b81523360048201526024016122b1565b6001600160a01b0381165f908152600183016020526040812054151561096b565b5f61096b836001600160a01b0384166138b1565b612ede6138fd565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b612f81613926565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612f0d3390565b60605f61096b83613950565b5f61096b836001600160a01b0384166139a8565b5f5f5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561303c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130609190614479565b90505f5f826001600160a01b0316636332aec68b8b6040518363ffffffff1660e01b81526004016130929291906143cb565b606060405180830381865afa1580156130ad573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130d191906144f6565b90965090925090508488116130f4575f5f5f5f96509650965096505050506131b4565b5f836001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa158015613131573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131559190614274565b90505f613162878b61429f565b905061317a8261317285826142c6565b839190613a8b565b9750613187888484613a8b565b955083881180156131ac57508561319e888a6142c6565b6131a891906142c6565b8a10155b985050505050505b93509350935093565b80825d5050565b6040516001600160a01b03848116602483015283811660448301526064820183905261322b9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613b41565b50505050565b5f5f8260018451613242919061429f565b81518110613252576132526142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b031603613299575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061331a9190614479565b90505f5f5f836001600160a01b0316636332aec68a876040518363ffffffff1660e01b815260040161334d9291906143cb565b606060405180830381865afa158015613368573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061338c91906144f6565b9194509092509050828810156133b5576040516326b56d9d60e21b815260040160405180910390fd5b5f846001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133f2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134169190614274565b90505f6134248a8584613a8b565b905082613431828c6142c6565b61343b91906142c6565b9b9a5050505050505050505050565b5f61346c60405180606001604052805f81526020015f81526020015f81525090565b5f5f856001875161347d919061429f565b8151811061348d5761348d6142b2565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03161493505f84613541576040516370a0823160e01b815230600482015282906001600160a01b038216906370a0823190602401602060405180830381865afa158015613518573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061353c9190614274565b613545565b6001475b9093509050841580156135e65750604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919085169063dd62ed3e90604401602060405180830381865afa1580156135c0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135e49190614274565b105b1561367d5760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015284169063095ea7b3906044016020604051808303815f875af1158015613657573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061367b9190614521565b505b61368b8c84838d8d8d613bad565b6040516315cd127f60e31b8152600481018e90529094505f92506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915063ae6893f890602401602060405180830381865afa1580156136f5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137199190614274565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635fd262de85613755575f613758565b84515b8d858e885f015189602001518a604001515f6001600160401b0381111561378157613781613db6565b6040519080825280601f01601f1916602001820160405280156137ab576020820181803683370190505b506040518963ffffffff1660e01b81526004016137ce97969594939291906145bd565b60206040518083038185885af11580156137ea573d5f5f3e3d5ffd5b50505050506040513d601f19601f8201168201806040525081019061380f9190614521565b61382c57604051632fba5fef60e21b815260040160405180910390fd5b336001600160a01b0316818c7f72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca8d8a8a60405161386b9392919061462f565b60405180910390a45050505050505050505050565b6040516001600160a01b038381166024830152604482018390526129b691859182169063a9059cbb906064016131f9565b5f8181526001830160205260408120546138f657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b56565b505f610b56565b5f54600160a01b900460ff16610d7657604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff1615610d765760405163d93c066560e01b815260040160405180910390fd5b6060815f01805480602002602001604051908101604052809291908181526020018280548015610ddf57602002820191905f5260205f20905b8154815260200190600101908083116139895750505050509050919050565b5f8181526001830160205260408120548015613a82575f6139ca60018361429f565b85549091505f906139dd9060019061429f565b9050808214613a3c575f865f0182815481106139fb576139fb6142b2565b905f5260205f200154905080875f018481548110613a1b57613a1b6142b2565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080613a4d57613a4d614664565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610b56565b5f915050610b56565b5f838302815f1985870982811083820303915050805f03613abf57838281613ab557613ab56142f0565b049250505061096b565b808411613ad657613ad66003851502601118613c7d565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b5f5f60205f8451602086015f885af180613b60576040513d5f823e3d81fd5b50505f513d91508115613b77578060011415613b84565b6001600160a01b0384163b155b1561322b57604051635274afe760e01b81526001600160a01b03851660048201526024016122b1565b613bce60405180606001604052805f81526020015f81526020015f81525090565b5f613bda888888612fdb565b604086015260208501528352905080613c06576040516326b56d9d60e21b815260040160405180910390fd5b8151851115613c2857604051637c3bec0960e11b815260040160405180910390fd5b8382602001511115613c4d5760405163de7539a960e01b815260040160405180910390fd5b8282604001511115613c725760405163cce42f3560e01b815260040160405180910390fd5b509695505050505050565b634e487b715f52806020526024601cfd5b828054828255905f5260205f20908101928215613ce1579160200282015b82811115613ce157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613cac565b50613ced929150613d08565b5090565b5080545f8255905f5260205f20908101906122c391905b5b80821115613ced575f8155600101613d09565b6001600160a01b03811681146122c3575f5ffd5b5f5f5f5f5f5f5f60e0888a031215613d46575f5ffd5b8735613d5181613d1c565b96506020880135613d6181613d1c565b96999698505050506040850135946060810135946080820135945060a0820135935060c0909101359150565b5f5f5f60608486031215613d9f575f5ffd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715613dec57613dec613db6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613e1a57613e1a613db6565b604052919050565b5f6001600160401b03821115613e3a57613e3a613db6565b5060051b60200190565b5f82601f830112613e53575f5ffd5b8135613e66613e6182613e22565b613df2565b8082825260208201915060208360051b860101925085831115613e87575f5ffd5b602085015b83811015613ead578035613e9f81613d1c565b835260209283019201613e8c565b5095945050505050565b5f5f60408385031215613ec8575f5ffd5b82356001600160401b03811115613edd575f5ffd5b613ee985828601613e44565b95602094909401359450505050565b5f5f60408385031215613f09575f5ffd5b8235915060208301356001600160401b03811115613f25575f5ffd5b613f3185828601613e44565b9150509250929050565b5f8151808452602084019350602083015f5b82811015613f6b578151865260209586019590910190600101613f4d565b5093949350505050565b602081525f61096b6020830184613f3b565b5f5f60408385031215613f98575f5ffd5b8235613fa381613d1c565b915060208301356001600160401b03811115613f25575f5ffd5b5f60208284031215613fcd575f5ffd5b813561096b81613d1c565b5f8151808452602084019350602083015f5b82811015613f6b5781516001600160a01b0316865260209586019590910190600101613fea565b602081525f61096b6020830184613fd8565b5f5f5f5f5f5f5f60e0888a031215614039575f5ffd5b87359650602088013561404b81613d1c565b955060408801359450606088013593506080880135925060a08801356001600160401b0381111561407a575f5ffd5b6140868a828b01613e44565b979a969950949793969295929450505060c09091013590565b5f5f5f5f5f60a086880312156140b3575f5ffd5b85356140be81613d1c565b97602087013597506040870135966060810135965060800135945092505050565b5f5f604083850312156140f0575f5ffd5b82356140fb81613d1c565b946020939093013593505050565b608081525f61411b6080830187613f3b565b6020830195909552506040810192909252606090910152919050565b5f5f5f5f5f5f5f5f610100898b03121561414f575f5ffd5b88359750602089013561416181613d1c565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b03811115614197575f5ffd5b6141a38b828c01613e44565b989b979a50959894979396929550929360e00135925050565b5f5f5f606084860312156141ce575f5ffd5b833592506020840135915060408401356001600160401b038111156141f1575f5ffd5b6141fd86828701613e44565b9150509250925092565b60a081525f61421960a0830188613f3b565b90508560208301528460408301528360608301528260808301529695505050505050565b5f5f6040838503121561424e575f5ffd5b823561425981613d1c565b9150602083013561426981613d1c565b809150509250929050565b5f60208284031215614284575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610b5657610b5661428b565b634e487b7160e01b5f52603260045260245ffd5b80820180821115610b5657610b5661428b565b8082028115828204841417610b5657610b5661428b565b634e487b7160e01b5f52601260045260245ffd5b5f8261431e57634e487b7160e01b5f52601260045260245ffd5b500490565b828152604060208201525f6109686040830184613fd8565b5f6020828403121561434b575f5ffd5b81516001600160401b03811115614360575f5ffd5b8201601f81018413614370575f5ffd5b805161437e613e6182613e22565b8082825260208201915060208360051b85010192508683111561439f575f5ffd5b6020840193505b828410156143c15783518252602093840193909101906143a6565b9695505050505050565b9182526001600160a01b0316602082015260400190565b805180151581146143f1575f5ffd5b919050565b5f60e0828403128015614407575f5ffd5b50614410613dca565b825161441b81613d1c565b8152602083015161442b81613d1c565b602082015261443c604084016143e2565b604082015261444d606084016143e2565b60608201526080838101519082015260a0808401519082015260c0928301519281019290925250919050565b5f60208284031215614489575f5ffd5b815161096b81613d1c565b80516001600160701b03811681146143f1575f5ffd5b5f5f5f606084860312156144bc575f5ffd5b6144c584614494565b92506144d360208501614494565b9150604084015163ffffffff811681146144eb575f5ffd5b809150509250925092565b5f5f5f60608486031215614508575f5ffd5b5050815160208301516040909301519094929350919050565b5f60208284031215614531575f5ffd5b61096b826143e2565b634e487b7160e01b5f52602160045260245ffd5b85815284602082015260a060408201525f61456c60a0830186613fd8565b6001600160a01b0394909416606083015250608001529392505050565b848152608060208201525f6145a16080830186613fd8565b6001600160a01b03949094166040830152506060015292915050565b87815260018060a01b038716602082015260018060a01b03861660408201528460608201528360808201528260a082015260e060c08201525f82518060e0840152806020850161010085015e5f6101008285010152610100601f19601f83011684010191505098975050505050505050565b6001600160a01b03841681526060602082018190525f9061465290830185613fd8565b82810360408401526143c18185613f3b565b634e487b7160e01b5f52603160045260245ffdfe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122001b235830d738541fc14a42aeb2ee1f027559c74608910dc8f054dee77914f4864736f6c634300081c0033"

// Deprecated: Use SwapBridgeRouterMetaData.Sigs instead.
// SwapBridgeRouterFuncSigs maps the 4-byte function signature to its string representation.
var SwapBridgeRouterFuncSigs = SwapBridgeRouterMetaData.Sigs

// SwapBridgeRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapBridgeRouterMetaData.Bin instead.
var SwapBridgeRouterBin = SwapBridgeRouterMetaData.Bin

// DeploySwapBridgeRouter deploys a new Ethereum contract, binding an instance of SwapBridgeRouter to it.
func DeploySwapBridgeRouter(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, bridge_ common.Address, swapRouter_ common.Address, crossToken_ common.Address, crossChainID_ *big.Int) (common.Address, *types.Transaction, *SwapBridgeRouter, error) {
	parsed, err := SwapBridgeRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapBridgeRouterBin), backend, owner_, bridge_, swapRouter_, crossToken_, crossChainID_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapBridgeRouter{SwapBridgeRouterCaller: SwapBridgeRouterCaller{contract: contract}, SwapBridgeRouterTransactor: SwapBridgeRouterTransactor{contract: contract}, SwapBridgeRouterFilterer: SwapBridgeRouterFilterer{contract: contract}}, nil
}

// SwapBridgeRouter is an auto generated Go binding around an Ethereum contract.
type SwapBridgeRouter struct {
	SwapBridgeRouterCaller     // Read-only binding to the contract
	SwapBridgeRouterTransactor // Write-only binding to the contract
	SwapBridgeRouterFilterer   // Log filterer for contract events
}

// SwapBridgeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapBridgeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapBridgeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapBridgeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapBridgeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapBridgeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapBridgeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapBridgeRouterSession struct {
	Contract     *SwapBridgeRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapBridgeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapBridgeRouterCallerSession struct {
	Contract *SwapBridgeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SwapBridgeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapBridgeRouterTransactorSession struct {
	Contract     *SwapBridgeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SwapBridgeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapBridgeRouterRaw struct {
	Contract *SwapBridgeRouter // Generic contract binding to access the raw methods on
}

// SwapBridgeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapBridgeRouterCallerRaw struct {
	Contract *SwapBridgeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SwapBridgeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapBridgeRouterTransactorRaw struct {
	Contract *SwapBridgeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapBridgeRouter creates a new instance of SwapBridgeRouter, bound to a specific deployed contract.
func NewSwapBridgeRouter(address common.Address, backend bind.ContractBackend) (*SwapBridgeRouter, error) {
	contract, err := bindSwapBridgeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouter{SwapBridgeRouterCaller: SwapBridgeRouterCaller{contract: contract}, SwapBridgeRouterTransactor: SwapBridgeRouterTransactor{contract: contract}, SwapBridgeRouterFilterer: SwapBridgeRouterFilterer{contract: contract}}, nil
}

// NewSwapBridgeRouterCaller creates a new read-only instance of SwapBridgeRouter, bound to a specific deployed contract.
func NewSwapBridgeRouterCaller(address common.Address, caller bind.ContractCaller) (*SwapBridgeRouterCaller, error) {
	contract, err := bindSwapBridgeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterCaller{contract: contract}, nil
}

// NewSwapBridgeRouterTransactor creates a new write-only instance of SwapBridgeRouter, bound to a specific deployed contract.
func NewSwapBridgeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapBridgeRouterTransactor, error) {
	contract, err := bindSwapBridgeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterTransactor{contract: contract}, nil
}

// NewSwapBridgeRouterFilterer creates a new log filterer instance of SwapBridgeRouter, bound to a specific deployed contract.
func NewSwapBridgeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapBridgeRouterFilterer, error) {
	contract, err := bindSwapBridgeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterFilterer{contract: contract}, nil
}

// bindSwapBridgeRouter binds a generic wrapper to an already deployed contract.
func bindSwapBridgeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapBridgeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapBridgeRouter *SwapBridgeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapBridgeRouter.Contract.SwapBridgeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapBridgeRouter *SwapBridgeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapBridgeRouter *SwapBridgeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapBridgeRouter *SwapBridgeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapBridgeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapBridgeRouter *SwapBridgeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapBridgeRouter *SwapBridgeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) BRIDGE() (common.Address, error) {
	return _SwapBridgeRouter.Contract.BRIDGE(&_SwapBridgeRouter.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) BRIDGE() (common.Address, error) {
	return _SwapBridgeRouter.Contract.BRIDGE(&_SwapBridgeRouter.CallOpts)
}

// CROSSCHAINID is a free data retrieval call binding the contract method 0x3b269e4e.
//
// Solidity: function CROSS_CHAIN_ID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) CROSSCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "CROSS_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CROSSCHAINID is a free data retrieval call binding the contract method 0x3b269e4e.
//
// Solidity: function CROSS_CHAIN_ID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterSession) CROSSCHAINID() (*big.Int, error) {
	return _SwapBridgeRouter.Contract.CROSSCHAINID(&_SwapBridgeRouter.CallOpts)
}

// CROSSCHAINID is a free data retrieval call binding the contract method 0x3b269e4e.
//
// Solidity: function CROSS_CHAIN_ID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) CROSSCHAINID() (*big.Int, error) {
	return _SwapBridgeRouter.Contract.CROSSCHAINID(&_SwapBridgeRouter.CallOpts)
}

// CROSSTOKEN is a free data retrieval call binding the contract method 0x1e10ebdc.
//
// Solidity: function CROSS_TOKEN() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) CROSSTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "CROSS_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSTOKEN is a free data retrieval call binding the contract method 0x1e10ebdc.
//
// Solidity: function CROSS_TOKEN() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) CROSSTOKEN() (common.Address, error) {
	return _SwapBridgeRouter.Contract.CROSSTOKEN(&_SwapBridgeRouter.CallOpts)
}

// CROSSTOKEN is a free data retrieval call binding the contract method 0x1e10ebdc.
//
// Solidity: function CROSS_TOKEN() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) CROSSTOKEN() (common.Address, error) {
	return _SwapBridgeRouter.Contract.CROSSTOKEN(&_SwapBridgeRouter.CallOpts)
}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) SWAPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "SWAP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) SWAPROUTER() (common.Address, error) {
	return _SwapBridgeRouter.Contract.SWAPROUTER(&_SwapBridgeRouter.CallOpts)
}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) SWAPROUTER() (common.Address, error) {
	return _SwapBridgeRouter.Contract.SWAPROUTER(&_SwapBridgeRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) WETH() (common.Address, error) {
	return _SwapBridgeRouter.Contract.WETH(&_SwapBridgeRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) WETH() (common.Address, error) {
	return _SwapBridgeRouter.Contract.WETH(&_SwapBridgeRouter.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountIn(&_SwapBridgeRouter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountIn(&_SwapBridgeRouter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountOut(&_SwapBridgeRouter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountOut(&_SwapBridgeRouter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountsIn(&_SwapBridgeRouter.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountsIn(&_SwapBridgeRouter.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountsOut(&_SwapBridgeRouter.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetAmountsOut(&_SwapBridgeRouter.CallOpts, amountIn, path)
}

// GetPath is a free data retrieval call binding the contract method 0x4f973bac.
//
// Solidity: function getPath(address token) view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetPath(opts *bind.CallOpts, token common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getPath", token)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPath is a free data retrieval call binding the contract method 0x4f973bac.
//
// Solidity: function getPath(address token) view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetPath(token common.Address) ([]common.Address, error) {
	return _SwapBridgeRouter.Contract.GetPath(&_SwapBridgeRouter.CallOpts, token)
}

// GetPath is a free data retrieval call binding the contract method 0x4f973bac.
//
// Solidity: function getPath(address token) view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetPath(token common.Address) ([]common.Address, error) {
	return _SwapBridgeRouter.Contract.GetPath(&_SwapBridgeRouter.CallOpts, token)
}

// GetPriceImpact is a free data retrieval call binding the contract method 0x0a353a5e.
//
// Solidity: function getPriceImpact(address[] path, uint256 amountIn) view returns(uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetPriceImpact(opts *bind.CallOpts, path []common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getPriceImpact", path, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceImpact is a free data retrieval call binding the contract method 0x0a353a5e.
//
// Solidity: function getPriceImpact(address[] path, uint256 amountIn) view returns(uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetPriceImpact(path []common.Address, amountIn *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetPriceImpact(&_SwapBridgeRouter.CallOpts, path, amountIn)
}

// GetPriceImpact is a free data retrieval call binding the contract method 0x0a353a5e.
//
// Solidity: function getPriceImpact(address[] path, uint256 amountIn) view returns(uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetPriceImpact(path []common.Address, amountIn *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.GetPriceImpact(&_SwapBridgeRouter.CallOpts, path, amountIn)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint112 reserveA, uint112 reserveB)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getReserves", tokenA, tokenB)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint112 reserveA, uint112 reserveB)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetReserves(tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetReserves(&_SwapBridgeRouter.CallOpts, tokenA, tokenB)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint112 reserveA, uint112 reserveB)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetReserves(tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetReserves(&_SwapBridgeRouter.CallOpts, tokenA, tokenB)
}

// GetSourceTokens is a free data retrieval call binding the contract method 0x92583a1a.
//
// Solidity: function getSourceTokens() view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetSourceTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getSourceTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSourceTokens is a free data retrieval call binding the contract method 0x92583a1a.
//
// Solidity: function getSourceTokens() view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetSourceTokens() ([]common.Address, error) {
	return _SwapBridgeRouter.Contract.GetSourceTokens(&_SwapBridgeRouter.CallOpts)
}

// GetSourceTokens is a free data retrieval call binding the contract method 0x92583a1a.
//
// Solidity: function getSourceTokens() view returns(address[])
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetSourceTokens() ([]common.Address, error) {
	return _SwapBridgeRouter.Contract.GetSourceTokens(&_SwapBridgeRouter.CallOpts)
}

// GetSwapBridgeIn is a free data retrieval call binding the contract method 0xfa768893.
//
// Solidity: function getSwapBridgeIn(uint256 toChainID, uint256 amountOut, address[] path) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetSwapBridgeIn(opts *bind.CallOpts, toChainID *big.Int, amountOut *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getSwapBridgeIn", toChainID, amountOut, path)

	outstruct := new(struct {
		Amounts        []*big.Int
		NetworkFee     *big.Int
		ExFee          *big.Int
		PriceImpactBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PriceImpactBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapBridgeIn is a free data retrieval call binding the contract method 0xfa768893.
//
// Solidity: function getSwapBridgeIn(uint256 toChainID, uint256 amountOut, address[] path) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetSwapBridgeIn(toChainID *big.Int, amountOut *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeIn(&_SwapBridgeRouter.CallOpts, toChainID, amountOut, path)
}

// GetSwapBridgeIn is a free data retrieval call binding the contract method 0xfa768893.
//
// Solidity: function getSwapBridgeIn(uint256 toChainID, uint256 amountOut, address[] path) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetSwapBridgeIn(toChainID *big.Int, amountOut *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeIn(&_SwapBridgeRouter.CallOpts, toChainID, amountOut, path)
}

// GetSwapBridgeInCross is a free data retrieval call binding the contract method 0x7f332b52.
//
// Solidity: function getSwapBridgeInCross(address token, uint256 amountOut) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetSwapBridgeInCross(opts *bind.CallOpts, token common.Address, amountOut *big.Int) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getSwapBridgeInCross", token, amountOut)

	outstruct := new(struct {
		Amounts        []*big.Int
		NetworkFee     *big.Int
		ExFee          *big.Int
		PriceImpactBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PriceImpactBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapBridgeInCross is a free data retrieval call binding the contract method 0x7f332b52.
//
// Solidity: function getSwapBridgeInCross(address token, uint256 amountOut) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetSwapBridgeInCross(token common.Address, amountOut *big.Int) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeInCross(&_SwapBridgeRouter.CallOpts, token, amountOut)
}

// GetSwapBridgeInCross is a free data retrieval call binding the contract method 0x7f332b52.
//
// Solidity: function getSwapBridgeInCross(address token, uint256 amountOut) view returns(uint256[] amounts, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetSwapBridgeInCross(token common.Address, amountOut *big.Int) (struct {
	Amounts        []*big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeInCross(&_SwapBridgeRouter.CallOpts, token, amountOut)
}

// GetSwapBridgeOut is a free data retrieval call binding the contract method 0xcdd19b7e.
//
// Solidity: function getSwapBridgeOut(uint256 toChainID, uint256 amountIn, address[] path) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetSwapBridgeOut(opts *bind.CallOpts, toChainID *big.Int, amountIn *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getSwapBridgeOut", toChainID, amountIn, path)

	outstruct := new(struct {
		Amounts        []*big.Int
		BridgeValue    *big.Int
		NetworkFee     *big.Int
		ExFee          *big.Int
		PriceImpactBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.BridgeValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PriceImpactBps = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapBridgeOut is a free data retrieval call binding the contract method 0xcdd19b7e.
//
// Solidity: function getSwapBridgeOut(uint256 toChainID, uint256 amountIn, address[] path) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetSwapBridgeOut(toChainID *big.Int, amountIn *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeOut(&_SwapBridgeRouter.CallOpts, toChainID, amountIn, path)
}

// GetSwapBridgeOut is a free data retrieval call binding the contract method 0xcdd19b7e.
//
// Solidity: function getSwapBridgeOut(uint256 toChainID, uint256 amountIn, address[] path) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetSwapBridgeOut(toChainID *big.Int, amountIn *big.Int, path []common.Address) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeOut(&_SwapBridgeRouter.CallOpts, toChainID, amountIn, path)
}

// GetSwapBridgeOutCross is a free data retrieval call binding the contract method 0xdd625509.
//
// Solidity: function getSwapBridgeOutCross(address token, uint256 amountIn) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) GetSwapBridgeOutCross(opts *bind.CallOpts, token common.Address, amountIn *big.Int) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "getSwapBridgeOutCross", token, amountIn)

	outstruct := new(struct {
		Amounts        []*big.Int
		BridgeValue    *big.Int
		NetworkFee     *big.Int
		ExFee          *big.Int
		PriceImpactBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.BridgeValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NetworkFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PriceImpactBps = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapBridgeOutCross is a free data retrieval call binding the contract method 0xdd625509.
//
// Solidity: function getSwapBridgeOutCross(address token, uint256 amountIn) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterSession) GetSwapBridgeOutCross(token common.Address, amountIn *big.Int) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeOutCross(&_SwapBridgeRouter.CallOpts, token, amountIn)
}

// GetSwapBridgeOutCross is a free data retrieval call binding the contract method 0xdd625509.
//
// Solidity: function getSwapBridgeOutCross(address token, uint256 amountIn) view returns(uint256[] amounts, uint256 bridgeValue, uint256 networkFee, uint256 exFee, uint256 priceImpactBps)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) GetSwapBridgeOutCross(token common.Address, amountIn *big.Int) (struct {
	Amounts        []*big.Int
	BridgeValue    *big.Int
	NetworkFee     *big.Int
	ExFee          *big.Int
	PriceImpactBps *big.Int
}, error) {
	return _SwapBridgeRouter.Contract.GetSwapBridgeOutCross(&_SwapBridgeRouter.CallOpts, token, amountIn)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) Owner() (common.Address, error) {
	return _SwapBridgeRouter.Contract.Owner(&_SwapBridgeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) Owner() (common.Address, error) {
	return _SwapBridgeRouter.Contract.Owner(&_SwapBridgeRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapBridgeRouter *SwapBridgeRouterSession) Paused() (bool, error) {
	return _SwapBridgeRouter.Contract.Paused(&_SwapBridgeRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) Paused() (bool, error) {
	return _SwapBridgeRouter.Contract.Paused(&_SwapBridgeRouter.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_SwapBridgeRouter *SwapBridgeRouterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.Quote(&_SwapBridgeRouter.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _SwapBridgeRouter.Contract.Quote(&_SwapBridgeRouter.CallOpts, amountA, reserveA, reserveB)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) Pause() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Pause(&_SwapBridgeRouter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) Pause() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Pause(&_SwapBridgeRouter.TransactOpts)
}

// RemoveTokenPath is a paid mutator transaction binding the contract method 0xcdbed967.
//
// Solidity: function removeTokenPath(address sourceToken) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) RemoveTokenPath(opts *bind.TransactOpts, sourceToken common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "removeTokenPath", sourceToken)
}

// RemoveTokenPath is a paid mutator transaction binding the contract method 0xcdbed967.
//
// Solidity: function removeTokenPath(address sourceToken) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) RemoveTokenPath(sourceToken common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.RemoveTokenPath(&_SwapBridgeRouter.TransactOpts, sourceToken)
}

// RemoveTokenPath is a paid mutator transaction binding the contract method 0xcdbed967.
//
// Solidity: function removeTokenPath(address sourceToken) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) RemoveTokenPath(sourceToken common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.RemoveTokenPath(&_SwapBridgeRouter.TransactOpts, sourceToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.RenounceOwnership(&_SwapBridgeRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.RenounceOwnership(&_SwapBridgeRouter.TransactOpts)
}

// SetTokenPath is a paid mutator transaction binding the contract method 0x1fc810c4.
//
// Solidity: function setTokenPath(address sourceToken, address[] path) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SetTokenPath(opts *bind.TransactOpts, sourceToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "setTokenPath", sourceToken, path)
}

// SetTokenPath is a paid mutator transaction binding the contract method 0x1fc810c4.
//
// Solidity: function setTokenPath(address sourceToken, address[] path) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SetTokenPath(sourceToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SetTokenPath(&_SwapBridgeRouter.TransactOpts, sourceToken, path)
}

// SetTokenPath is a paid mutator transaction binding the contract method 0x1fc810c4.
//
// Solidity: function setTokenPath(address sourceToken, address[] path) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SetTokenPath(sourceToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SetTokenPath(&_SwapBridgeRouter.TransactOpts, sourceToken, path)
}

// SwapBridgeETHForExactCrossTokens is a paid mutator transaction binding the contract method 0xfd24b30b.
//
// Solidity: function swapBridgeETHForExactCrossTokens(address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeETHForExactCrossTokens(opts *bind.TransactOpts, to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeETHForExactCrossTokens", to, amountOut, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeETHForExactCrossTokens is a paid mutator transaction binding the contract method 0xfd24b30b.
//
// Solidity: function swapBridgeETHForExactCrossTokens(address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeETHForExactCrossTokens(to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeETHForExactCrossTokens(&_SwapBridgeRouter.TransactOpts, to, amountOut, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeETHForExactCrossTokens is a paid mutator transaction binding the contract method 0xfd24b30b.
//
// Solidity: function swapBridgeETHForExactCrossTokens(address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeETHForExactCrossTokens(to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeETHForExactCrossTokens(&_SwapBridgeRouter.TransactOpts, to, amountOut, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeETHForExactTokens is a paid mutator transaction binding the contract method 0x60f71c20.
//
// Solidity: function swapBridgeETHForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeETHForExactTokens(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeETHForExactTokens", toChainID, to, amountOut, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeETHForExactTokens is a paid mutator transaction binding the contract method 0x60f71c20.
//
// Solidity: function swapBridgeETHForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeETHForExactTokens(toChainID *big.Int, to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeETHForExactTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeETHForExactTokens is a paid mutator transaction binding the contract method 0x60f71c20.
//
// Solidity: function swapBridgeETHForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeETHForExactTokens(toChainID *big.Int, to common.Address, amountOut *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeETHForExactTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactETHForCrossTokens is a paid mutator transaction binding the contract method 0x7d173b7a.
//
// Solidity: function swapBridgeExactETHForCrossTokens(address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeExactETHForCrossTokens(opts *bind.TransactOpts, to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeExactETHForCrossTokens", to, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactETHForCrossTokens is a paid mutator transaction binding the contract method 0x7d173b7a.
//
// Solidity: function swapBridgeExactETHForCrossTokens(address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeExactETHForCrossTokens(to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactETHForCrossTokens(&_SwapBridgeRouter.TransactOpts, to, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactETHForCrossTokens is a paid mutator transaction binding the contract method 0x7d173b7a.
//
// Solidity: function swapBridgeExactETHForCrossTokens(address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeExactETHForCrossTokens(to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactETHForCrossTokens(&_SwapBridgeRouter.TransactOpts, to, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactETHForTokens is a paid mutator transaction binding the contract method 0xc78cb094.
//
// Solidity: function swapBridgeExactETHForTokens(uint256 toChainID, address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeExactETHForTokens(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeExactETHForTokens", toChainID, to, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactETHForTokens is a paid mutator transaction binding the contract method 0xc78cb094.
//
// Solidity: function swapBridgeExactETHForTokens(uint256 toChainID, address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeExactETHForTokens(toChainID *big.Int, to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactETHForTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactETHForTokens is a paid mutator transaction binding the contract method 0xc78cb094.
//
// Solidity: function swapBridgeExactETHForTokens(uint256 toChainID, address to, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeExactETHForTokens(toChainID *big.Int, to common.Address, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactETHForTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForCrossTokens is a paid mutator transaction binding the contract method 0xb379262d.
//
// Solidity: function swapBridgeExactTokensForCrossTokens(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeExactTokensForCrossTokens(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeExactTokensForCrossTokens", sourceToken, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactTokensForCrossTokens is a paid mutator transaction binding the contract method 0xb379262d.
//
// Solidity: function swapBridgeExactTokensForCrossTokens(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeExactTokensForCrossTokens(sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForCrossTokens(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactTokensForCrossTokens is a paid mutator transaction binding the contract method 0xb379262d.
//
// Solidity: function swapBridgeExactTokensForCrossTokens(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeExactTokensForCrossTokens(sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForCrossTokens(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeExactTokensForETH is a paid mutator transaction binding the contract method 0xaf69d7fe.
//
// Solidity: function swapBridgeExactTokensForETH(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeExactTokensForETH(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeExactTokensForETH", toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForETH is a paid mutator transaction binding the contract method 0xaf69d7fe.
//
// Solidity: function swapBridgeExactTokensForETH(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeExactTokensForETH(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForETH(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForETH is a paid mutator transaction binding the contract method 0xaf69d7fe.
//
// Solidity: function swapBridgeExactTokensForETH(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeExactTokensForETH(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForETH(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForTokens is a paid mutator transaction binding the contract method 0xbe835388.
//
// Solidity: function swapBridgeExactTokensForTokens(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeExactTokensForTokens(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeExactTokensForTokens", toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForTokens is a paid mutator transaction binding the contract method 0xbe835388.
//
// Solidity: function swapBridgeExactTokensForTokens(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeExactTokensForTokens(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeExactTokensForTokens is a paid mutator transaction binding the contract method 0xbe835388.
//
// Solidity: function swapBridgeExactTokensForTokens(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeExactTokensForTokens(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeExactTokensForTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactCrossTokens is a paid mutator transaction binding the contract method 0x02a827f2.
//
// Solidity: function swapBridgeTokensForExactCrossTokens(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeTokensForExactCrossTokens(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeTokensForExactCrossTokens", sourceToken, to, amountOut, amountInMax, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeTokensForExactCrossTokens is a paid mutator transaction binding the contract method 0x02a827f2.
//
// Solidity: function swapBridgeTokensForExactCrossTokens(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeTokensForExactCrossTokens(sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactCrossTokens(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, amountInMax, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeTokensForExactCrossTokens is a paid mutator transaction binding the contract method 0x02a827f2.
//
// Solidity: function swapBridgeTokensForExactCrossTokens(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeTokensForExactCrossTokens(sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactCrossTokens(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, amountInMax, networkFeeMax, exFeeMax, deadline)
}

// SwapBridgeTokensForExactETH is a paid mutator transaction binding the contract method 0xa88c3f7c.
//
// Solidity: function swapBridgeTokensForExactETH(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeTokensForExactETH(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeTokensForExactETH", toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactETH is a paid mutator transaction binding the contract method 0xa88c3f7c.
//
// Solidity: function swapBridgeTokensForExactETH(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeTokensForExactETH(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactETH(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactETH is a paid mutator transaction binding the contract method 0xa88c3f7c.
//
// Solidity: function swapBridgeTokensForExactETH(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeTokensForExactETH(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactETH(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactTokens is a paid mutator transaction binding the contract method 0xbf557c55.
//
// Solidity: function swapBridgeTokensForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapBridgeTokensForExactTokens(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapBridgeTokensForExactTokens", toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactTokens is a paid mutator transaction binding the contract method 0xbf557c55.
//
// Solidity: function swapBridgeTokensForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapBridgeTokensForExactTokens(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// SwapBridgeTokensForExactTokens is a paid mutator transaction binding the contract method 0xbf557c55.
//
// Solidity: function swapBridgeTokensForExactTokens(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 networkFeeMax, uint256 exFeeMax, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapBridgeTokensForExactTokens(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, networkFeeMax *big.Int, exFeeMax *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapBridgeTokensForExactTokens(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, networkFeeMax, exFeeMax, path, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.TransferOwnership(&_SwapBridgeRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.TransferOwnership(&_SwapBridgeRouter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) Unpause() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Unpause(&_SwapBridgeRouter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) Unpause() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Unpause(&_SwapBridgeRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) Receive() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Receive(&_SwapBridgeRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.Receive(&_SwapBridgeRouter.TransactOpts)
}

// SwapBridgeRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterOwnershipTransferredIterator struct {
	Event *SwapBridgeRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterOwnershipTransferred)
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
		it.Event = new(SwapBridgeRouterOwnershipTransferred)
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
func (it *SwapBridgeRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterOwnershipTransferred represents a OwnershipTransferred event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapBridgeRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterOwnershipTransferredIterator{contract: _SwapBridgeRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterOwnershipTransferred)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParseOwnershipTransferred(log types.Log) (*SwapBridgeRouterOwnershipTransferred, error) {
	event := new(SwapBridgeRouterOwnershipTransferred)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBridgeRouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterPausedIterator struct {
	Event *SwapBridgeRouterPaused // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterPaused)
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
		it.Event = new(SwapBridgeRouterPaused)
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
func (it *SwapBridgeRouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterPaused represents a Paused event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterPaused(opts *bind.FilterOpts) (*SwapBridgeRouterPausedIterator, error) {

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterPausedIterator{contract: _SwapBridgeRouter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterPaused) (event.Subscription, error) {

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterPaused)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParsePaused(log types.Log) (*SwapBridgeRouterPaused, error) {
	event := new(SwapBridgeRouterPaused)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBridgeRouterSwapBridgeInitiatedIterator is returned from FilterSwapBridgeInitiated and is used to iterate over the raw logs and unpacked data for SwapBridgeInitiated events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterSwapBridgeInitiatedIterator struct {
	Event *SwapBridgeRouterSwapBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterSwapBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterSwapBridgeInitiated)
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
		it.Event = new(SwapBridgeRouterSwapBridgeInitiated)
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
func (it *SwapBridgeRouterSwapBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterSwapBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterSwapBridgeInitiated represents a SwapBridgeInitiated event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterSwapBridgeInitiated struct {
	ToChainID *big.Int
	Index     *big.Int
	From      common.Address
	To        common.Address
	Path      []common.Address
	Amounts   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapBridgeInitiated is a free log retrieval operation binding the contract event 0x72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca.
//
// Solidity: event SwapBridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address indexed from, address to, address[] path, uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterSwapBridgeInitiated(opts *bind.FilterOpts, toChainID []*big.Int, index []*big.Int, from []common.Address) (*SwapBridgeRouterSwapBridgeInitiatedIterator, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "SwapBridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterSwapBridgeInitiatedIterator{contract: _SwapBridgeRouter.contract, event: "SwapBridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchSwapBridgeInitiated is a free log subscription operation binding the contract event 0x72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca.
//
// Solidity: event SwapBridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address indexed from, address to, address[] path, uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchSwapBridgeInitiated(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterSwapBridgeInitiated, toChainID []*big.Int, index []*big.Int, from []common.Address) (event.Subscription, error) {

	var toChainIDRule []interface{}
	for _, toChainIDItem := range toChainID {
		toChainIDRule = append(toChainIDRule, toChainIDItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "SwapBridgeInitiated", toChainIDRule, indexRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterSwapBridgeInitiated)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "SwapBridgeInitiated", log); err != nil {
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

// ParseSwapBridgeInitiated is a log parse operation binding the contract event 0x72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca.
//
// Solidity: event SwapBridgeInitiated(uint256 indexed toChainID, uint256 indexed index, address indexed from, address to, address[] path, uint256[] amounts)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParseSwapBridgeInitiated(log types.Log) (*SwapBridgeRouterSwapBridgeInitiated, error) {
	event := new(SwapBridgeRouterSwapBridgeInitiated)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "SwapBridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBridgeRouterTokenPathRemovedIterator is returned from FilterTokenPathRemoved and is used to iterate over the raw logs and unpacked data for TokenPathRemoved events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterTokenPathRemovedIterator struct {
	Event *SwapBridgeRouterTokenPathRemoved // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterTokenPathRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterTokenPathRemoved)
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
		it.Event = new(SwapBridgeRouterTokenPathRemoved)
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
func (it *SwapBridgeRouterTokenPathRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterTokenPathRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterTokenPathRemoved represents a TokenPathRemoved event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterTokenPathRemoved struct {
	SourceToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPathRemoved is a free log retrieval operation binding the contract event 0xcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193.
//
// Solidity: event TokenPathRemoved(address indexed sourceToken)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterTokenPathRemoved(opts *bind.FilterOpts, sourceToken []common.Address) (*SwapBridgeRouterTokenPathRemovedIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "TokenPathRemoved", sourceTokenRule)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterTokenPathRemovedIterator{contract: _SwapBridgeRouter.contract, event: "TokenPathRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenPathRemoved is a free log subscription operation binding the contract event 0xcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193.
//
// Solidity: event TokenPathRemoved(address indexed sourceToken)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchTokenPathRemoved(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterTokenPathRemoved, sourceToken []common.Address) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "TokenPathRemoved", sourceTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterTokenPathRemoved)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "TokenPathRemoved", log); err != nil {
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

// ParseTokenPathRemoved is a log parse operation binding the contract event 0xcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193.
//
// Solidity: event TokenPathRemoved(address indexed sourceToken)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParseTokenPathRemoved(log types.Log) (*SwapBridgeRouterTokenPathRemoved, error) {
	event := new(SwapBridgeRouterTokenPathRemoved)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "TokenPathRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBridgeRouterTokenPathSetIterator is returned from FilterTokenPathSet and is used to iterate over the raw logs and unpacked data for TokenPathSet events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterTokenPathSetIterator struct {
	Event *SwapBridgeRouterTokenPathSet // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterTokenPathSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterTokenPathSet)
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
		it.Event = new(SwapBridgeRouterTokenPathSet)
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
func (it *SwapBridgeRouterTokenPathSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterTokenPathSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterTokenPathSet represents a TokenPathSet event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterTokenPathSet struct {
	SourceToken common.Address
	Path        []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPathSet is a free log retrieval operation binding the contract event 0x654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec.
//
// Solidity: event TokenPathSet(address indexed sourceToken, address[] path)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterTokenPathSet(opts *bind.FilterOpts, sourceToken []common.Address) (*SwapBridgeRouterTokenPathSetIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "TokenPathSet", sourceTokenRule)
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterTokenPathSetIterator{contract: _SwapBridgeRouter.contract, event: "TokenPathSet", logs: logs, sub: sub}, nil
}

// WatchTokenPathSet is a free log subscription operation binding the contract event 0x654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec.
//
// Solidity: event TokenPathSet(address indexed sourceToken, address[] path)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchTokenPathSet(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterTokenPathSet, sourceToken []common.Address) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "TokenPathSet", sourceTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterTokenPathSet)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "TokenPathSet", log); err != nil {
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

// ParseTokenPathSet is a log parse operation binding the contract event 0x654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec.
//
// Solidity: event TokenPathSet(address indexed sourceToken, address[] path)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParseTokenPathSet(log types.Log) (*SwapBridgeRouterTokenPathSet, error) {
	event := new(SwapBridgeRouterTokenPathSet)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "TokenPathSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBridgeRouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SwapBridgeRouter contract.
type SwapBridgeRouterUnpausedIterator struct {
	Event *SwapBridgeRouterUnpaused // Event containing the contract specifics and raw log

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
func (it *SwapBridgeRouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBridgeRouterUnpaused)
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
		it.Event = new(SwapBridgeRouterUnpaused)
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
func (it *SwapBridgeRouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBridgeRouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBridgeRouterUnpaused represents a Unpaused event raised by the SwapBridgeRouter contract.
type SwapBridgeRouterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SwapBridgeRouterUnpausedIterator, error) {

	logs, sub, err := _SwapBridgeRouter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SwapBridgeRouterUnpausedIterator{contract: _SwapBridgeRouter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SwapBridgeRouterUnpaused) (event.Subscription, error) {

	logs, sub, err := _SwapBridgeRouter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBridgeRouterUnpaused)
				if err := _SwapBridgeRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SwapBridgeRouter *SwapBridgeRouterFilterer) ParseUnpaused(log types.Log) (*SwapBridgeRouterUnpaused, error) {
	event := new(SwapBridgeRouterUnpaused)
	if err := _SwapBridgeRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
