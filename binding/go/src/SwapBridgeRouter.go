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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapRouter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crossToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossChainID_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBaseBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPath\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getPriceImpact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserveA\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserveB\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getSwapBridgeIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"getSwapBridgeInCross\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getSwapBridgeOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bridgeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getSwapBridgeOutCross\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bridgeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceImpactBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"removeTokenPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"setTokenPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactCrossTokensBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokensBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForCrossTokensBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForCrossTokensBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractIPancakeRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactCrossTokensBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETHBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNetworkFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokensBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToPath\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"SwapBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"TokenPathRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"TokenPathSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeBridgeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeBridgeValueTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeCanNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeExFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInsufficientReserve\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInvalidBridgeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeInvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeNetworkFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgeOnlySwapRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapBridgePairDoesNotExist\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"e78cea92": "bridge()",
		"2c7b395f": "crossChainID()",
		"90f1b5a4": "crossToken()",
		"85f8c259": "getAmountIn(uint256,uint256,uint256)",
		"054d50d4": "getAmountOut(uint256,uint256,uint256)",
		"1f00ca74": "getAmountsIn(uint256,address[])",
		"d06ca61f": "getAmountsOut(uint256,address[])",
		"4f973bac": "getPath(address)",
		"0a353a5e": "getPriceImpact(address[],uint256)",
		"d52bb6f4": "getReserves(address,address)",
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
		"ec7d9618": "swapETHForExactCrossTokensBridge(address,address,uint256,uint256,uint256,uint256)",
		"ad921379": "swapETHForExactTokensBridge(uint256,address,uint256,uint256,uint256,address[],uint256)",
		"38208f67": "swapExactETHForCrossTokensBridge(address,address,uint256,uint256,uint256,uint256)",
		"88ad9911": "swapExactETHForTokensBridge(uint256,address,uint256,uint256,uint256,address[],uint256)",
		"f0a85637": "swapExactTokensForCrossTokensBridge(address,address,uint256,uint256,uint256,uint256,uint256)",
		"7e5262b4": "swapExactTokensForETHBridge(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"a964ea54": "swapExactTokensForTokensBridge(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"c31c9c07": "swapRouter()",
		"5ce53779": "swapTokensForExactCrossTokensBridge(address,address,uint256,uint256,uint256,uint256,uint256)",
		"bb1024f8": "swapTokensForExactETHBridge(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"235feb6a": "swapTokensForExactTokensBridge(uint256,address,uint256,uint256,uint256,uint256,address[],uint256)",
		"3e42c51f": "tokenToPath(address,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x610100604052348015610010575f5ffd5b50604051614cb1380380614cb183398101604081905261002f916101bc565b846001600160a01b03811661005d57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61006681610152565b505f805460ff60a01b191690556001600160a01b03851661009a576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b0384166100c1576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b0383166100e8576040516332746c8560e01b815260040160405180910390fd5b6001600160a01b03821661010f576040516332746c8560e01b815260040160405180910390fd5b805f0361012f576040516395e19d4f60e01b815260040160405180910390fd5b6001600160a01b0393841660805291831660a05290911660c05260e05250610216565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146101b7575f5ffd5b919050565b5f5f5f5f5f60a086880312156101d0575f5ffd5b6101d9866101a1565b94506101e7602087016101a1565b93506101f5604087016101a1565b9250610203606087016101a1565b9150608086015190509295509295909350565b60805160a05160c05160e0516148fc6103b55f395f818161031f01528181610f190152818161118201528181611470015281816123a80152818161251f015261268901525f81816104db01528181610aac01528181610f5501526111be01525f818161020f0152818161057e01528181610735015281816109bc01528181610be901528181611260015281816114dd0152818161156d0152818161178401528181611932015281816119c201528181611ba401528181611e1b01528181611fa7015281816120dd015281816121190152818161274d015281816128d901528181612bd001528181612d0501528181612d9601528181612e6101528181612f1901528181612ff501528181613070015281816130d501528181613164015281816135e9015261388201525f818161067d01528181610c9c01528181611313015281816116200152818161183701528181611a7501528181611c5701528181611ece0152818161280001528181612976015281816133da01528181613686015281816139c101528181613a5301528181613b000152613b6e01526148fc5ff3fe6080604052600436106101ff575f3560e01c806388ad991111610113578063cdd19b7e1161009d578063e78cea921161006d578063e78cea921461066c578063ec7d96181461069f578063f0a85637146106b2578063f2fde38b146106d1578063fa768893146106f0575f5ffd5b8063cdd19b7e146105bf578063d06ca61f146105ef578063d52bb6f41461060e578063dd6255091461064d575f5ffd5b8063ad615dec116100e3578063ad615dec1461051c578063ad9213791461053b578063bb1024f81461054e578063c31c9c071461056d578063cdbed967146105a0575f5ffd5b806388ad99111461049b5780638da5cb5b146104ae57806390f1b5a4146104ca578063a964ea54146104fd575f5ffd5b80633f4ba83a11610194578063715018a611610164578063715018a6146104125780637e5262b4146104265780637f332b52146104395780638456cb591461046857806385f8c2591461047c575f5ffd5b80633f4ba83a1461038b5780634f973bac1461039f5780635c975abb146103cb5780635ce53779146103f3575f5ffd5b8063235feb6a116101cf578063235feb6a146102ef5780632c7b395f1461030e57806338208f67146103415780633e42c51f14610354575f5ffd5b8063054d50d4146102535780630a353a5e146102855780631f00ca74146102a45780631fc810c4146102d0575f5ffd5b3661024f57336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461024d576040516313e21f5560e31b815260040160405180910390fd5b005b5f5ffd5b34801561025e575f5ffd5b5061027261026d366004613f49565b61070f565b6040519081526020015b60405180910390f35b348015610290575f5ffd5b5061027261029f366004614087565b6107b1565b3480156102af575f5ffd5b506102c36102be3660046140c8565b6109a2565b60405161027c9190614145565b3480156102db575f5ffd5b5061024d6102ea366004614157565b610a34565b3480156102fa575f5ffd5b5061024d61030936600461418d565b610b90565b348015610319575f5ffd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b61024d61034f366004614212565b610ddc565b34801561035f575f5ffd5b5061037361036e366004614267565b610f8c565b6040516001600160a01b03909116815260200161027c565b348015610396575f5ffd5b5061024d610fc0565b3480156103aa575f5ffd5b506103be6103b9366004614291565b610fd2565b60405161027c91906142e5565b3480156103d6575f5ffd5b505f54600160a01b900460ff16604051901515815260200161027c565b3480156103fe575f5ffd5b5061024d61040d3660046142f7565b611045565b34801561041d575f5ffd5b5061024d6111f6565b61024d61043436600461418d565b611207565b348015610444575f5ffd5b50610458610453366004614267565b6113e9565b60405161027c9493929190614354565b348015610473575f5ffd5b5061024d6114a7565b348015610487575f5ffd5b50610272610496366004613f49565b6114b7565b61024d6104a9366004614382565b611514565b3480156104b9575f5ffd5b505f546001600160a01b0316610373565b3480156104d5575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b348015610508575f5ffd5b5061024d61051736600461418d565b61172b565b348015610527575f5ffd5b50610272610536366004613f49565b61190c565b61024d610549366004614382565b611969565b348015610559575f5ffd5b5061024d61056836600461418d565b611b4b565b348015610578575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b3480156105ab575f5ffd5b5061024d6105ba366004614291565b611d2d565b3480156105ca575f5ffd5b506105de6105d93660046143fe565b611dc4565b60405161027c959493929190614449565b3480156105fa575f5ffd5b506102c36106093660046140c8565b6120c3565b348015610619575f5ffd5b5061062d61062836600461447f565b612114565b604080516001600160701b0393841681529290911660208301520161027c565b348015610658575f5ffd5b506105de610667366004614267565b61232e565b348015610677575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b61024d6106ad366004614212565b6123e2565b3480156106bd575f5ffd5b5061024d6106cc3660046142f7565b61254c565b3480156106dc575f5ffd5b5061024d6106eb366004614291565b6126b5565b3480156106fb575f5ffd5b5061045861070a3660046143fe565b6126f7565b604051630153543560e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063054d50d4906064015b602060405180830381865afa158015610783573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a791906144b6565b90505b9392505050565b5f6002835110806107c0575081155b156107cc57505f61099c565b5f6107d783856120c3565b90505f5b600185516107e991906144e1565b811015610999575f858281518110610803576108036144f4565b602002602001015190505f8683600161081c9190614508565b8151811061082c5761082c6144f4565b602002602001015190505f5f6108428484612114565b915091505f826001600160701b031611801561086657505f816001600160701b0316115b6108835760405163a008318960e01b815260040160405180910390fd5b5f826001600160701b0316826001600160701b0316670de0b6b3a76400006108ab919061451b565b6108b59190614546565b90505f8787815181106108ca576108ca6144f4565b602002602001015190505f888860016108e39190614508565b815181106108f3576108f36144f4565b602002602001015190505f82866001600160701b03166109139190614508565b610926836001600160701b0388166144e1565b61093890670de0b6b3a764000061451b565b6109429190614546565b90505f81851115610975578461095883826144e1565b6109649061271061451b565b61096e9190614546565b9050610978565b505f5b610982818d614508565b9b5050600190980197506107db9650505050505050565b50505b92915050565b6040516307c0329d60e21b81526060906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631f00ca74906109f39086908690600401614565565b5f60405180830381865afa158015610a0d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107aa919081019061457d565b610a3c612ae7565b600281511015610a5f5760405163bb52eed160e01b815260040160405180910390fd5b816001600160a01b0316815f81518110610a7b57610a7b6144f4565b60200260200101516001600160a01b031614610aaa5760405163bb52eed160e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168160018351610ae391906144e1565b81518110610af357610af36144f4565b60200260200101516001600160a01b031614610b225760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0382165f9081526001602090815260409091208251610b4a92840190613ebb565b50816001600160a01b03167f654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec82604051610b8491906142e5565b60405180910390a25050565b610b98612b13565b8782600281511015610bbd5760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351610bcd91906144e1565b81518110610bdd57610bdd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c43573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c67919061460d565b6001600160a01b0316816001600160a01b031603610c83575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590610cd39087908690600401614628565b60e060405180830381865afa158015610cee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d129190614653565b9050816001600160a01b0316815f01516001600160a01b031614610d49576040516337088ab560e11b815260040160405180910390fd5b8589610d6e825f81518110610d6057610d606144f4565b602002602001015182612b5a565b610d808e8e8d8f8e8e8e8e6003612e16565b610da2825f81518110610d9557610d956144f4565b6020026020010151613203565b610dc48260018451610db491906144e1565b81518110610d9557610d956144f4565b505050505050610dd26132d7565b5050505050505050565b610de4612b13565b6001600160a01b0386165f90815260016020908152604080832080548251818502810185019093528083528a9434949093929190830182828015610e4f57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610e31575b50505050509050600281511015610e795760405163bb52eed160e01b815260040160405180910390fd5b610e838383612b5a565b6001600160a01b0389165f90815260016020908152604080832080548251818502810185019093528083529192909190830182828015610eea57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610ecc575b50505050509050600281511015610f145760405163bb52eed160e01b815260040160405180910390fd5b610f467f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c6001612e16565b50610f5083613203565b610f797f0000000000000000000000000000000000000000000000000000000000000000613203565b505050610f846132d7565b505050505050565b6001602052815f5260405f208181548110610fa5575f80fd5b5f918252602090912001546001600160a01b03169150829050565b610fc8612ae7565b610fd06132ee565b565b6001600160a01b0381165f9081526001602090815260409182902080548351818402810184019094528084526060939283018282801561103957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161101b575b50505050509050919050565b61104d612b13565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b94899490939291908301828280156110b857602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161109a575b505050505090506002815110156110e25760405163bb52eed160e01b815260040160405180910390fd5b6110ec8383612b5a565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561115357602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611135575b5050505050905060028151101561117d5760405163bb52eed160e01b815260040160405180910390fd5b6111af7f00000000000000000000000000000000000000000000000000000000000000008b8a8c8b8b878c6003612e16565b506111b983613203565b6111e27f0000000000000000000000000000000000000000000000000000000000000000613203565b5050506111ed6132d7565b50505050505050565b6111fe612ae7565b610fd05f613342565b61120f612b13565b87826002815110156112345760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161124491906144e1565b81518110611254576112546144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ba573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112de919061460d565b6001600160a01b0316816001600160a01b0316036112fa575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b59061134a9087908690600401614628565b60e060405180830381865afa158015611365573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113899190614653565b9050816001600160a01b0316815f01516001600160a01b0316146113c0576040516337088ab560e11b815260040160405180910390fd5b858a6113d7825f81518110610d6057610d606144f4565b610d808e8e8e8e8e8e8e8e6002612e16565b60605f5f5f5f60015f886001600160a01b03166001600160a01b031681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561146457602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611446575b505050505090506114967f000000000000000000000000000000000000000000000000000000000000000087836126f7565b929a91995097509095509350505050565b6114af612ae7565b610fd0613391565b6040516385f8c25960e01b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906385f8c25990606401610768565b61151c612b13565b86826002815110156115415760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161155191906144e1565b81518110611561576115616144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115eb919061460d565b6001600160a01b0316816001600160a01b031603611607575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906116579087908690600401614628565b60e060405180830381865afa158015611672573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116969190614653565b9050816001600160a01b0316815f01516001600160a01b0316146116cd576040516337088ab560e11b815260040160405180910390fd5b85346116e4825f81518110610d6057610d606144f4565b6116f68d8d348e8e8e8e8e6001612e16565b61170b825f81518110610d9557610d956144f4565b61171d8260018451610db491906144e1565b5050505050506111ed6132d7565b611733612b13565b87826002815110156117585760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161176891906144e1565b81518110611778576117786144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117de573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611802919061460d565b6001600160a01b0316816001600160a01b03160361181e575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b59061186e9087908690600401614628565b60e060405180830381865afa158015611889573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118ad9190614653565b9050816001600160a01b0316815f01516001600160a01b0316146118e4576040516337088ab560e11b815260040160405180910390fd5b858a6118fb825f81518110610d6057610d606144f4565b610d808e8e8e8e8e8e8e8e5f612e16565b604051632b58577b60e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ad615dec90606401610768565b611971612b13565b86826002815110156119965760405163bb52eed160e01b815260040160405180910390fd5b5f81600183516119a691906144e1565b815181106119b6576119b66144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a1c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a40919061460d565b6001600160a01b0316816001600160a01b031603611a5c575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611aac9087908690600401614628565b60e060405180830381865afa158015611ac7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611aeb9190614653565b9050816001600160a01b0316815f01516001600160a01b031614611b22576040516337088ab560e11b815260040160405180910390fd5b8534611b39825f81518110610d6057610d606144f4565b6116f68d8d348e8e8e8e8e6004612e16565b611b53612b13565b8782600281511015611b785760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611b8891906144e1565b81518110611b9857611b986144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bfe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c22919061460d565b6001600160a01b0316816001600160a01b031603611c3e575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611c8e9087908690600401614628565b60e060405180830381865afa158015611ca9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ccd9190614653565b9050816001600160a01b0316815f01516001600160a01b031614611d04576040516337088ab560e11b815260040160405180910390fd5b8589611d1b825f81518110610d6057610d606144f4565b610d808e8e8d8f8e8e8e8e6005612e16565b611d35612ae7565b6001600160a01b0381165f9081526001602052604090205460021115611d6e5760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0381165f908152600160205260408120611d8e91613f1e565b6040516001600160a01b038216907fcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193905f90a250565b60605f5f5f5f8786600281511015611def5760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611dff91906144e1565b81518110611e0f57611e0f6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e75573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e99919061460d565b6001600160a01b0316816001600160a01b031603611eb5575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611f059087908690600401614628565b60e060405180830381865afa158015611f20573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f449190614653565b9050816001600160a01b0316815f01516001600160a01b031614611f7b576040516337088ab560e11b815260040160405180910390fd5b5f8a60018c51611f8b91906144e1565b81518110611f9b57611f9b6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612001573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612025919061460d565b6001600160a01b0316816001600160a01b031603612041575060015b61204b8c8c6120c3565b99505f61207f8e838d60018f5161206291906144e1565b81518110612072576120726144f4565b60200260200101516133d3565b919c509a5098509050806120a6576040516326b56d9d60e21b815260040160405180910390fd5b6120b08c8e6107b1565b9650505050505050939792965093509350565b60405163d06ca61f60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d06ca61f906109f39086908690600401614565565b5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c45a01556040518163ffffffff1660e01b8152600401602060405180830381865afa158015612173573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612197919061460d565b60405163e6a4390560e01b81526001600160a01b0387811660048301528681166024830152919091169063e6a4390590604401602060405180830381865afa1580156121e5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612209919061460d565b90506001600160a01b03811661223257604051630240c4d360e11b815260040160405180910390fd5b5f8190505f5f826001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa158015612274573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061229891906146ec565b5091509150826001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122fd919061460d565b6001600160a01b0316886001600160a01b03161461231c57808261231f565b81815b90999098509650505050505050565b6001600160a01b0382165f908152600160209081526040808320805482518185028101850190935280835260609493849384938493849383018282801561239c57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161237e575b505050505090506123ce7f00000000000000000000000000000000000000000000000000000000000000008883611dc4565b939c929b5090995097509095509350505050565b6123ea612b13565b6001600160a01b0386165f90815260016020908152604080832080548251818502810185019093528083528a943494909392919083018282801561245557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612437575b5050505050905060028151101561247f5760405163bb52eed160e01b815260040160405180910390fd5b6124898383612b5a565b6001600160a01b0389165f908152600160209081526040808320805482518185028101850190935280835291929091908301828280156124f057602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116124d2575b5050505050905060028151101561251a5760405163bb52eed160e01b815260040160405180910390fd5b610f467f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c6004612e16565b612554612b13565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b948a9490939291908301828280156125bf57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116125a1575b505050505090506002815110156125e95760405163bb52eed160e01b815260040160405180910390fd5b6125f38383612b5a565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561265a57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161263c575b505050505090506002815110156126845760405163bb52eed160e01b815260040160405180910390fd5b6111af7f00000000000000000000000000000000000000000000000000000000000000008b8b8b8b8b878c5f612e16565b6126bd612ae7565b6001600160a01b0381166126eb57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6126f481613342565b50565b60605f5f5f86856002815110156127215760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161273191906144e1565b81518110612741576127416144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127a7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127cb919061460d565b6001600160a01b0316816001600160a01b0316036127e7575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906128379087908690600401614628565b60e060405180830381865afa158015612852573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128769190614653565b9050816001600160a01b0316815f01516001600160a01b0316146128ad576040516337088ab560e11b815260040160405180910390fd5b5f8960018b516128bd91906144e1565b815181106128cd576128cd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612933573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612957919061460d565b6001600160a01b0316816001600160a01b031603612973575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129f4919061460d565b6040516337dba1f760e21b8152600481018f90526001600160a01b038481166024830152604482018f9052919091169063df6e87dc90606401606060405180830381865afa158015612a48573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a6c9190614738565b909a5098509050808c11612a93576040516326b56d9d60e21b815260040160405180910390fd5b612ab188612aa18b8f614508565b612aab9190614508565b8c6109a2565b9950612ad68b8b5f81518110612ac957612ac96144f4565b60200260200101516107b1565b965050505050505093509350935093565b5f546001600160a01b03163314610fd05760405163118cdaa760e01b81523360048201526024016126e2565b5f5160206148a75f395f51905f525c15612b4057604051633ee5aeb560e01b815260040160405180910390fd5b610fd060015f5160206148a75f395f51905f525b906135b5565b8015612b665780612bcc565b6040516370a0823160e01b81523360048201526001600160a01b038316906370a0823190602401602060405180830381865afa158015612ba8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bcc91906144b6565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c2a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c4e919061460d565b6001600160a01b0316826001600160a01b0316148015612c6e5750803410155b612ce8576040516323b872dd60e01b8152336004820152306024820152604481018290526001600160a01b038316906323b872dd906064016020604051808303815f875af1158015612cc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ce69190614763565b505b604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919084169063dd62ed3e90604401602060405180830381865afa158015612d55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d7991906144b6565b1015612e125760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015283169063095ea7b3906044015b6020604051808303815f875af1158015612dec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e109190614763565b505b5050565b5f80606081846005811115612e2d57612e2d61477c565b03612ee75786612e3d898b614508565b612e479190614508565b6040516338ed173960e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906338ed173990612e9e908d9087908b9030908c90600401614790565b5f604051808303815f875af1158015612eb9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612ee0919081019061457d565b90506131e6565b6001846005811115612efb57612efb61477c565b03612fad5786612f0b898b614508565b612f159190614508565b92507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff36ab58b8589308a6040518663ffffffff1660e01b8152600401612f6a94939291906147cb565b5f6040518083038185885af1158015612f85573d5f5f3e3d5ffd5b50505050506040513d5f823e601f3d908101601f19168201604052612ee0919081019061457d565b6002846005811115612fc157612fc161477c565b036130325786612fd1898b614508565b612fdb9190614508565b6040516318cbafe560e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906318cbafe590612e9e908d9087908b9030908c90600401614790565b60038460058111156130465761304661477c565b036130ad576130568c8a886135bc565b604051634401edf760e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638803dbee90612e9e9085908e908b9030908c90600401614790565b60048460058111156130c1576130c161477c565b03613126576130d18c8a886135bc565b91507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fb3bdb418b8489308a6040518663ffffffff1660e01b8152600401612f6a94939291906147cb565b600584600581111561313a5761313a61477c565b036131e65761314a8c8a886135bc565b604051632512eca560e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634a25d94a906131a19085908e908b9030908c90600401614790565b5f604051808303815f875af11580156131bc573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526131e3919081019061457d565b90505b6131f58c8c8b8b8b8b87613833565b505050505050505050505050565b47156132345760405133904780156108fc02915f818181858888f19350505050158015613232573d5f5f3e3d5ffd5b505b6040516370a0823160e01b81523060048201525f906001600160a01b038316906370a0823190602401602060405180830381865afa158015613278573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061329c91906144b6565b90508015612e125760405163a9059cbb60e01b8152336004820152602481018290526001600160a01b0383169063a9059cbb90604401612dd0565b610fd05f5f5160206148a75f395f51905f52612b54565b6132f6613cd1565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b613399613cfa565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586133253390565b5f5f5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613434573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613458919061460d565b90505f5f826001600160a01b0316636332aec68b8b6040518363ffffffff1660e01b815260040161348a929190614628565b606060405180830381865afa1580156134a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134c99190614738565b90965090925090508488116134ec575f5f5f5f96509650965096505050506135ac565b5f836001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa158015613529573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061354d91906144b6565b90505f61355a878b6144e1565b90506135728261356a8582614508565b839190613d24565b975061357f888484613d24565b955083881180156135a4575085613596888a614508565b6135a09190614508565b8a10155b985050505050505b93509350935093565b80825d5050565b5f5f82600184516135cd91906144e1565b815181106135dd576135dd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015613643573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613667919061460d565b6001600160a01b0316816001600160a01b031603613683575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136e0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613704919061460d565b90505f5f5f836001600160a01b0316636332aec68a876040518363ffffffff1660e01b8152600401613737929190614628565b606060405180830381865afa158015613752573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137769190614738565b919450909250905082881161379e576040516326b56d9d60e21b815260040160405180910390fd5b5f846001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137db573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137ff91906144b6565b90505f61380d8a8584613d24565b90508261381a828c614508565b6138249190614508565b9b9a5050505050505050505050565b5f61385560405180606001604052805f81526020015f81526020015f81525090565b5f5f856001875161386691906144e1565b81518110613876576138766144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156138dc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613900919061460d565b6001600160a01b0316816001600160a01b0316149350836139215780613924565b60015b91505f84613997576040516370a0823160e01b81523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa15801561396e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061399291906144b6565b613999565b475b905084158015613a375750604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919085169063dd62ed3e90604401602060405180830381865afa158015613a11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a3591906144b6565b105b15613ace5760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015284169063095ea7b3906044016020604051808303815f875af1158015613aa8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613acc9190614763565b505b613adc8c84838d8d8d613dda565b6040516315cd127f60e31b8152600481018e90529094505f92506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915063ae6893f890602401602060405180830381865afa158015613b46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b6a91906144b6565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635fd262de85613ba6575f613ba9565b84515b8d858e885f015189602001518a604001515f6001600160401b03811115613bd257613bd2613f72565b6040519080825280601f01601f191660200182016040528015613bfc576020820181803683370190505b506040518963ffffffff1660e01b8152600401613c1f97969594939291906147ff565b60206040518083038185885af1158015613c3b573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190613c609190614763565b613c7d57604051632fba5fef60e21b815260040160405180910390fd5b336001600160a01b0316818c7f72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca8d8a8a604051613cbc93929190614871565b60405180910390a45050505050505050505050565b5f54600160a01b900460ff16610fd057604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff1615610fd05760405163d93c066560e01b815260040160405180910390fd5b5f838302815f1985870982811083820303915050805f03613d5857838281613d4e57613d4e614532565b04925050506107aa565b808411613d6f57613d6f6003851502601118613eaa565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b613dfb60405180606001604052805f81526020015f81526020015f81525090565b5f613e078888886133d3565b604086015260208501528352905080613e33576040516326b56d9d60e21b815260040160405180910390fd5b8151851115613e5557604051637c3bec0960e11b815260040160405180910390fd5b8382602001511115613e7a5760405163de7539a960e01b815260040160405180910390fd5b8282604001511115613e9f5760405163cce42f3560e01b815260040160405180910390fd5b509695505050505050565b634e487b715f52806020526024601cfd5b828054828255905f5260205f20908101928215613f0e579160200282015b82811115613f0e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613ed9565b50613f1a929150613f35565b5090565b5080545f8255905f5260205f20908101906126f491905b5b80821115613f1a575f8155600101613f36565b5f5f5f60608486031215613f5b575f5ffd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715613fa857613fa8613f72565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613fd657613fd6613f72565b604052919050565b5f6001600160401b03821115613ff657613ff6613f72565b5060051b60200190565b6001600160a01b03811681146126f4575f5ffd5b5f82601f830112614023575f5ffd5b813561403661403182613fde565b613fae565b8082825260208201915060208360051b860101925085831115614057575f5ffd5b602085015b8381101561407d57803561406f81614000565b83526020928301920161405c565b5095945050505050565b5f5f60408385031215614098575f5ffd5b82356001600160401b038111156140ad575f5ffd5b6140b985828601614014565b95602094909401359450505050565b5f5f604083850312156140d9575f5ffd5b8235915060208301356001600160401b038111156140f5575f5ffd5b61410185828601614014565b9150509250929050565b5f8151808452602084019350602083015f5b8281101561413b57815186526020958601959091019060010161411d565b5093949350505050565b602081525f6107aa602083018461410b565b5f5f60408385031215614168575f5ffd5b823561417381614000565b915060208301356001600160401b038111156140f5575f5ffd5b5f5f5f5f5f5f5f5f610100898b0312156141a5575f5ffd5b8835975060208901356141b781614000565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b038111156141ed575f5ffd5b6141f98b828c01614014565b989b979a50959894979396929550929360e00135925050565b5f5f5f5f5f5f60c08789031215614227575f5ffd5b863561423281614000565b9550602087013561424281614000565b95989597505050506040840135936060810135936080820135935060a0909101359150565b5f5f60408385031215614278575f5ffd5b823561428381614000565b946020939093013593505050565b5f602082840312156142a1575f5ffd5b81356107aa81614000565b5f8151808452602084019350602083015f5b8281101561413b5781516001600160a01b03168652602095860195909101906001016142be565b602081525f6107aa60208301846142ac565b5f5f5f5f5f5f5f60e0888a03121561430d575f5ffd5b873561431881614000565b9650602088013561432881614000565b96999698505050506040850135946060810135946080820135945060a0820135935060c0909101359150565b608081525f614366608083018761410b565b6020830195909552506040810192909252606090910152919050565b5f5f5f5f5f5f5f60e0888a031215614398575f5ffd5b8735965060208801356143aa81614000565b955060408801359450606088013593506080880135925060a08801356001600160401b038111156143d9575f5ffd5b6143e58a828b01614014565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614410575f5ffd5b833592506020840135915060408401356001600160401b03811115614433575f5ffd5b61443f86828701614014565b9150509250925092565b60a081525f61445b60a083018861410b565b90508560208301528460408301528360608301528260808301529695505050505050565b5f5f60408385031215614490575f5ffd5b823561449b81614000565b915060208301356144ab81614000565b809150509250929050565b5f602082840312156144c6575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561099c5761099c6144cd565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561099c5761099c6144cd565b808202811582820484141761099c5761099c6144cd565b634e487b7160e01b5f52601260045260245ffd5b5f8261456057634e487b7160e01b5f52601260045260245ffd5b500490565b828152604060208201525f6107a760408301846142ac565b5f6020828403121561458d575f5ffd5b81516001600160401b038111156145a2575f5ffd5b8201601f810184136145b2575f5ffd5b80516145c061403182613fde565b8082825260208201915060208360051b8501019250868311156145e1575f5ffd5b6020840193505b828410156146035783518252602093840193909101906145e8565b9695505050505050565b5f6020828403121561461d575f5ffd5b81516107aa81614000565b9182526001600160a01b0316602082015260400190565b8051801515811461464e575f5ffd5b919050565b5f60e0828403128015614664575f5ffd5b5061466d613f86565b825161467881614000565b8152602083015161468881614000565b60208201526146996040840161463f565b60408201526146aa6060840161463f565b60608201526080838101519082015260a0808401519082015260c0928301519281019290925250919050565b80516001600160701b038116811461464e575f5ffd5b5f5f5f606084860312156146fe575f5ffd5b614707846146d6565b9250614715602085016146d6565b9150604084015163ffffffff8116811461472d575f5ffd5b809150509250925092565b5f5f5f6060848603121561474a575f5ffd5b5050815160208301516040909301519094929350919050565b5f60208284031215614773575f5ffd5b6107aa8261463f565b634e487b7160e01b5f52602160045260245ffd5b85815284602082015260a060408201525f6147ae60a08301866142ac565b6001600160a01b0394909416606083015250608001529392505050565b848152608060208201525f6147e360808301866142ac565b6001600160a01b03949094166040830152506060015292915050565b87815260018060a01b038716602082015260018060a01b03861660408201528460608201528360808201528260a082015260e060c08201525f82518060e0840152806020850161010085015e5f6101008285010152610100601f19601f83011684010191505098975050505050505050565b6001600160a01b03841681526060602082018190525f90614894908301856142ac565b8281036040840152614603818561410b56fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122054f4b24286abd2281838a36e562f7e27a1b813ce734071a8c623bc8adf20769164736f6c634300081c0033",
}

// SwapBridgeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapBridgeRouterMetaData.ABI instead.
var SwapBridgeRouterABI = SwapBridgeRouterMetaData.ABI

// SwapBridgeRouterBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const SwapBridgeRouterBinRuntime = "6080604052600436106101ff575f3560e01c806388ad991111610113578063cdd19b7e1161009d578063e78cea921161006d578063e78cea921461066c578063ec7d96181461069f578063f0a85637146106b2578063f2fde38b146106d1578063fa768893146106f0575f5ffd5b8063cdd19b7e146105bf578063d06ca61f146105ef578063d52bb6f41461060e578063dd6255091461064d575f5ffd5b8063ad615dec116100e3578063ad615dec1461051c578063ad9213791461053b578063bb1024f81461054e578063c31c9c071461056d578063cdbed967146105a0575f5ffd5b806388ad99111461049b5780638da5cb5b146104ae57806390f1b5a4146104ca578063a964ea54146104fd575f5ffd5b80633f4ba83a11610194578063715018a611610164578063715018a6146104125780637e5262b4146104265780637f332b52146104395780638456cb591461046857806385f8c2591461047c575f5ffd5b80633f4ba83a1461038b5780634f973bac1461039f5780635c975abb146103cb5780635ce53779146103f3575f5ffd5b8063235feb6a116101cf578063235feb6a146102ef5780632c7b395f1461030e57806338208f67146103415780633e42c51f14610354575f5ffd5b8063054d50d4146102535780630a353a5e146102855780631f00ca74146102a45780631fc810c4146102d0575f5ffd5b3661024f57336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461024d576040516313e21f5560e31b815260040160405180910390fd5b005b5f5ffd5b34801561025e575f5ffd5b5061027261026d366004613f49565b61070f565b6040519081526020015b60405180910390f35b348015610290575f5ffd5b5061027261029f366004614087565b6107b1565b3480156102af575f5ffd5b506102c36102be3660046140c8565b6109a2565b60405161027c9190614145565b3480156102db575f5ffd5b5061024d6102ea366004614157565b610a34565b3480156102fa575f5ffd5b5061024d61030936600461418d565b610b90565b348015610319575f5ffd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b61024d61034f366004614212565b610ddc565b34801561035f575f5ffd5b5061037361036e366004614267565b610f8c565b6040516001600160a01b03909116815260200161027c565b348015610396575f5ffd5b5061024d610fc0565b3480156103aa575f5ffd5b506103be6103b9366004614291565b610fd2565b60405161027c91906142e5565b3480156103d6575f5ffd5b505f54600160a01b900460ff16604051901515815260200161027c565b3480156103fe575f5ffd5b5061024d61040d3660046142f7565b611045565b34801561041d575f5ffd5b5061024d6111f6565b61024d61043436600461418d565b611207565b348015610444575f5ffd5b50610458610453366004614267565b6113e9565b60405161027c9493929190614354565b348015610473575f5ffd5b5061024d6114a7565b348015610487575f5ffd5b50610272610496366004613f49565b6114b7565b61024d6104a9366004614382565b611514565b3480156104b9575f5ffd5b505f546001600160a01b0316610373565b3480156104d5575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b348015610508575f5ffd5b5061024d61051736600461418d565b61172b565b348015610527575f5ffd5b50610272610536366004613f49565b61190c565b61024d610549366004614382565b611969565b348015610559575f5ffd5b5061024d61056836600461418d565b611b4b565b348015610578575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b3480156105ab575f5ffd5b5061024d6105ba366004614291565b611d2d565b3480156105ca575f5ffd5b506105de6105d93660046143fe565b611dc4565b60405161027c959493929190614449565b3480156105fa575f5ffd5b506102c36106093660046140c8565b6120c3565b348015610619575f5ffd5b5061062d61062836600461447f565b612114565b604080516001600160701b0393841681529290911660208301520161027c565b348015610658575f5ffd5b506105de610667366004614267565b61232e565b348015610677575f5ffd5b506103737f000000000000000000000000000000000000000000000000000000000000000081565b61024d6106ad366004614212565b6123e2565b3480156106bd575f5ffd5b5061024d6106cc3660046142f7565b61254c565b3480156106dc575f5ffd5b5061024d6106eb366004614291565b6126b5565b3480156106fb575f5ffd5b5061045861070a3660046143fe565b6126f7565b604051630153543560e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063054d50d4906064015b602060405180830381865afa158015610783573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a791906144b6565b90505b9392505050565b5f6002835110806107c0575081155b156107cc57505f61099c565b5f6107d783856120c3565b90505f5b600185516107e991906144e1565b811015610999575f858281518110610803576108036144f4565b602002602001015190505f8683600161081c9190614508565b8151811061082c5761082c6144f4565b602002602001015190505f5f6108428484612114565b915091505f826001600160701b031611801561086657505f816001600160701b0316115b6108835760405163a008318960e01b815260040160405180910390fd5b5f826001600160701b0316826001600160701b0316670de0b6b3a76400006108ab919061451b565b6108b59190614546565b90505f8787815181106108ca576108ca6144f4565b602002602001015190505f888860016108e39190614508565b815181106108f3576108f36144f4565b602002602001015190505f82866001600160701b03166109139190614508565b610926836001600160701b0388166144e1565b61093890670de0b6b3a764000061451b565b6109429190614546565b90505f81851115610975578461095883826144e1565b6109649061271061451b565b61096e9190614546565b9050610978565b505f5b610982818d614508565b9b5050600190980197506107db9650505050505050565b50505b92915050565b6040516307c0329d60e21b81526060906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631f00ca74906109f39086908690600401614565565b5f60405180830381865afa158015610a0d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526107aa919081019061457d565b610a3c612ae7565b600281511015610a5f5760405163bb52eed160e01b815260040160405180910390fd5b816001600160a01b0316815f81518110610a7b57610a7b6144f4565b60200260200101516001600160a01b031614610aaa5760405163bb52eed160e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168160018351610ae391906144e1565b81518110610af357610af36144f4565b60200260200101516001600160a01b031614610b225760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0382165f9081526001602090815260409091208251610b4a92840190613ebb565b50816001600160a01b03167f654e7e70ddd53f06d83a071616e02b594dee891d4e42f5055198d9e72bbe3cec82604051610b8491906142e5565b60405180910390a25050565b610b98612b13565b8782600281511015610bbd5760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351610bcd91906144e1565b81518110610bdd57610bdd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c43573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c67919061460d565b6001600160a01b0316816001600160a01b031603610c83575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590610cd39087908690600401614628565b60e060405180830381865afa158015610cee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d129190614653565b9050816001600160a01b0316815f01516001600160a01b031614610d49576040516337088ab560e11b815260040160405180910390fd5b8589610d6e825f81518110610d6057610d606144f4565b602002602001015182612b5a565b610d808e8e8d8f8e8e8e8e6003612e16565b610da2825f81518110610d9557610d956144f4565b6020026020010151613203565b610dc48260018451610db491906144e1565b81518110610d9557610d956144f4565b505050505050610dd26132d7565b5050505050505050565b610de4612b13565b6001600160a01b0386165f90815260016020908152604080832080548251818502810185019093528083528a9434949093929190830182828015610e4f57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610e31575b50505050509050600281511015610e795760405163bb52eed160e01b815260040160405180910390fd5b610e838383612b5a565b6001600160a01b0389165f90815260016020908152604080832080548251818502810185019093528083529192909190830182828015610eea57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610ecc575b50505050509050600281511015610f145760405163bb52eed160e01b815260040160405180910390fd5b610f467f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c6001612e16565b50610f5083613203565b610f797f0000000000000000000000000000000000000000000000000000000000000000613203565b505050610f846132d7565b505050505050565b6001602052815f5260405f208181548110610fa5575f80fd5b5f918252602090912001546001600160a01b03169150829050565b610fc8612ae7565b610fd06132ee565b565b6001600160a01b0381165f9081526001602090815260409182902080548351818402810184019094528084526060939283018282801561103957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161101b575b50505050509050919050565b61104d612b13565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b94899490939291908301828280156110b857602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161109a575b505050505090506002815110156110e25760405163bb52eed160e01b815260040160405180910390fd5b6110ec8383612b5a565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561115357602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611135575b5050505050905060028151101561117d5760405163bb52eed160e01b815260040160405180910390fd5b6111af7f00000000000000000000000000000000000000000000000000000000000000008b8a8c8b8b878c6003612e16565b506111b983613203565b6111e27f0000000000000000000000000000000000000000000000000000000000000000613203565b5050506111ed6132d7565b50505050505050565b6111fe612ae7565b610fd05f613342565b61120f612b13565b87826002815110156112345760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161124491906144e1565b81518110611254576112546144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ba573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112de919061460d565b6001600160a01b0316816001600160a01b0316036112fa575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b59061134a9087908690600401614628565b60e060405180830381865afa158015611365573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113899190614653565b9050816001600160a01b0316815f01516001600160a01b0316146113c0576040516337088ab560e11b815260040160405180910390fd5b858a6113d7825f81518110610d6057610d606144f4565b610d808e8e8e8e8e8e8e8e6002612e16565b60605f5f5f5f60015f886001600160a01b03166001600160a01b031681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561146457602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611446575b505050505090506114967f000000000000000000000000000000000000000000000000000000000000000087836126f7565b929a91995097509095509350505050565b6114af612ae7565b610fd0613391565b6040516385f8c25960e01b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906385f8c25990606401610768565b61151c612b13565b86826002815110156115415760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161155191906144e1565b81518110611561576115616144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115eb919061460d565b6001600160a01b0316816001600160a01b031603611607575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906116579087908690600401614628565b60e060405180830381865afa158015611672573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116969190614653565b9050816001600160a01b0316815f01516001600160a01b0316146116cd576040516337088ab560e11b815260040160405180910390fd5b85346116e4825f81518110610d6057610d606144f4565b6116f68d8d348e8e8e8e8e6001612e16565b61170b825f81518110610d9557610d956144f4565b61171d8260018451610db491906144e1565b5050505050506111ed6132d7565b611733612b13565b87826002815110156117585760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161176891906144e1565b81518110611778576117786144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117de573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611802919061460d565b6001600160a01b0316816001600160a01b03160361181e575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b59061186e9087908690600401614628565b60e060405180830381865afa158015611889573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118ad9190614653565b9050816001600160a01b0316815f01516001600160a01b0316146118e4576040516337088ab560e11b815260040160405180910390fd5b858a6118fb825f81518110610d6057610d606144f4565b610d808e8e8e8e8e8e8e8e5f612e16565b604051632b58577b60e21b81526004810184905260248101839052604481018290525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ad615dec90606401610768565b611971612b13565b86826002815110156119965760405163bb52eed160e01b815260040160405180910390fd5b5f81600183516119a691906144e1565b815181106119b6576119b66144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a1c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a40919061460d565b6001600160a01b0316816001600160a01b031603611a5c575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611aac9087908690600401614628565b60e060405180830381865afa158015611ac7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611aeb9190614653565b9050816001600160a01b0316815f01516001600160a01b031614611b22576040516337088ab560e11b815260040160405180910390fd5b8534611b39825f81518110610d6057610d606144f4565b6116f68d8d348e8e8e8e8e6004612e16565b611b53612b13565b8782600281511015611b785760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611b8891906144e1565b81518110611b9857611b986144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bfe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c22919061460d565b6001600160a01b0316816001600160a01b031603611c3e575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611c8e9087908690600401614628565b60e060405180830381865afa158015611ca9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ccd9190614653565b9050816001600160a01b0316815f01516001600160a01b031614611d04576040516337088ab560e11b815260040160405180910390fd5b8589611d1b825f81518110610d6057610d606144f4565b610d808e8e8d8f8e8e8e8e6005612e16565b611d35612ae7565b6001600160a01b0381165f9081526001602052604090205460021115611d6e5760405163bb52eed160e01b815260040160405180910390fd5b6001600160a01b0381165f908152600160205260408120611d8e91613f1e565b6040516001600160a01b038216907fcd237a66ab933b37859cd3402fcd457c1c7988af691ce485b3f98552cc87a193905f90a250565b60605f5f5f5f8786600281511015611def5760405163bb52eed160e01b815260040160405180910390fd5b5f8160018351611dff91906144e1565b81518110611e0f57611e0f6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e75573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e99919061460d565b6001600160a01b0316816001600160a01b031603611eb5575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b590611f059087908690600401614628565b60e060405180830381865afa158015611f20573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f449190614653565b9050816001600160a01b0316815f01516001600160a01b031614611f7b576040516337088ab560e11b815260040160405180910390fd5b5f8a60018c51611f8b91906144e1565b81518110611f9b57611f9b6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612001573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612025919061460d565b6001600160a01b0316816001600160a01b031603612041575060015b61204b8c8c6120c3565b99505f61207f8e838d60018f5161206291906144e1565b81518110612072576120726144f4565b60200260200101516133d3565b919c509a5098509050806120a6576040516326b56d9d60e21b815260040160405180910390fd5b6120b08c8e6107b1565b9650505050505050939792965093509350565b60405163d06ca61f60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d06ca61f906109f39086908690600401614565565b5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c45a01556040518163ffffffff1660e01b8152600401602060405180830381865afa158015612173573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612197919061460d565b60405163e6a4390560e01b81526001600160a01b0387811660048301528681166024830152919091169063e6a4390590604401602060405180830381865afa1580156121e5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612209919061460d565b90506001600160a01b03811661223257604051630240c4d360e11b815260040160405180910390fd5b5f8190505f5f826001600160a01b0316630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa158015612274573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061229891906146ec565b5091509150826001600160a01b0316630dfe16816040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122fd919061460d565b6001600160a01b0316886001600160a01b03161461231c57808261231f565b81815b90999098509650505050505050565b6001600160a01b0382165f908152600160209081526040808320805482518185028101850190935280835260609493849384938493849383018282801561239c57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161237e575b505050505090506123ce7f00000000000000000000000000000000000000000000000000000000000000008883611dc4565b939c929b5090995097509095509350505050565b6123ea612b13565b6001600160a01b0386165f90815260016020908152604080832080548251818502810185019093528083528a943494909392919083018282801561245557602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612437575b5050505050905060028151101561247f5760405163bb52eed160e01b815260040160405180910390fd5b6124898383612b5a565b6001600160a01b0389165f908152600160209081526040808320805482518185028101850190935280835291929091908301828280156124f057602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116124d2575b5050505050905060028151101561251a5760405163bb52eed160e01b815260040160405180910390fd5b610f467f00000000000000000000000000000000000000000000000000000000000000008a348b8b8b878c6004612e16565b612554612b13565b6001600160a01b0387165f90815260016020908152604080832080548251818502810185019093528083528b948a9490939291908301828280156125bf57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116125a1575b505050505090506002815110156125e95760405163bb52eed160e01b815260040160405180910390fd5b6125f38383612b5a565b6001600160a01b038a165f9081526001602090815260408083208054825181850281018501909352808352919290919083018282801561265a57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161263c575b505050505090506002815110156126845760405163bb52eed160e01b815260040160405180910390fd5b6111af7f00000000000000000000000000000000000000000000000000000000000000008b8b8b8b8b878c5f612e16565b6126bd612ae7565b6001600160a01b0381166126eb57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6126f481613342565b50565b60605f5f5f86856002815110156127215760405163bb52eed160e01b815260040160405180910390fd5b5f816001835161273191906144e1565b81518110612741576127416144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127a7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127cb919061460d565b6001600160a01b0316816001600160a01b0316036127e7575060015b60405163814914b560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063814914b5906128379087908690600401614628565b60e060405180830381865afa158015612852573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128769190614653565b9050816001600160a01b0316815f01516001600160a01b0316146128ad576040516337088ab560e11b815260040160405180910390fd5b5f8960018b516128bd91906144e1565b815181106128cd576128cd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612933573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612957919061460d565b6001600160a01b0316816001600160a01b031603612973575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129f4919061460d565b6040516337dba1f760e21b8152600481018f90526001600160a01b038481166024830152604482018f9052919091169063df6e87dc90606401606060405180830381865afa158015612a48573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a6c9190614738565b909a5098509050808c11612a93576040516326b56d9d60e21b815260040160405180910390fd5b612ab188612aa18b8f614508565b612aab9190614508565b8c6109a2565b9950612ad68b8b5f81518110612ac957612ac96144f4565b60200260200101516107b1565b965050505050505093509350935093565b5f546001600160a01b03163314610fd05760405163118cdaa760e01b81523360048201526024016126e2565b5f5160206148a75f395f51905f525c15612b4057604051633ee5aeb560e01b815260040160405180910390fd5b610fd060015f5160206148a75f395f51905f525b906135b5565b8015612b665780612bcc565b6040516370a0823160e01b81523360048201526001600160a01b038316906370a0823190602401602060405180830381865afa158015612ba8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bcc91906144b6565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c2a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c4e919061460d565b6001600160a01b0316826001600160a01b0316148015612c6e5750803410155b612ce8576040516323b872dd60e01b8152336004820152306024820152604481018290526001600160a01b038316906323b872dd906064016020604051808303815f875af1158015612cc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ce69190614763565b505b604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919084169063dd62ed3e90604401602060405180830381865afa158015612d55573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d7991906144b6565b1015612e125760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015283169063095ea7b3906044015b6020604051808303815f875af1158015612dec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e109190614763565b505b5050565b5f80606081846005811115612e2d57612e2d61477c565b03612ee75786612e3d898b614508565b612e479190614508565b6040516338ed173960e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906338ed173990612e9e908d9087908b9030908c90600401614790565b5f604051808303815f875af1158015612eb9573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612ee0919081019061457d565b90506131e6565b6001846005811115612efb57612efb61477c565b03612fad5786612f0b898b614508565b612f159190614508565b92507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637ff36ab58b8589308a6040518663ffffffff1660e01b8152600401612f6a94939291906147cb565b5f6040518083038185885af1158015612f85573d5f5f3e3d5ffd5b50505050506040513d5f823e601f3d908101601f19168201604052612ee0919081019061457d565b6002846005811115612fc157612fc161477c565b036130325786612fd1898b614508565b612fdb9190614508565b6040516318cbafe560e01b81529093506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906318cbafe590612e9e908d9087908b9030908c90600401614790565b60038460058111156130465761304661477c565b036130ad576130568c8a886135bc565b604051634401edf760e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638803dbee90612e9e9085908e908b9030908c90600401614790565b60048460058111156130c1576130c161477c565b03613126576130d18c8a886135bc565b91507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fb3bdb418b8489308a6040518663ffffffff1660e01b8152600401612f6a94939291906147cb565b600584600581111561313a5761313a61477c565b036131e65761314a8c8a886135bc565b604051632512eca560e11b81529092506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690634a25d94a906131a19085908e908b9030908c90600401614790565b5f604051808303815f875af11580156131bc573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526131e3919081019061457d565b90505b6131f58c8c8b8b8b8b87613833565b505050505050505050505050565b47156132345760405133904780156108fc02915f818181858888f19350505050158015613232573d5f5f3e3d5ffd5b505b6040516370a0823160e01b81523060048201525f906001600160a01b038316906370a0823190602401602060405180830381865afa158015613278573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061329c91906144b6565b90508015612e125760405163a9059cbb60e01b8152336004820152602481018290526001600160a01b0383169063a9059cbb90604401612dd0565b610fd05f5f5160206148a75f395f51905f52612b54565b6132f6613cd1565b5f805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b613399613cfa565b5f805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586133253390565b5f5f5f5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613434573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613458919061460d565b90505f5f826001600160a01b0316636332aec68b8b6040518363ffffffff1660e01b815260040161348a929190614628565b606060405180830381865afa1580156134a5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134c99190614738565b90965090925090508488116134ec575f5f5f5f96509650965096505050506135ac565b5f836001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa158015613529573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061354d91906144b6565b90505f61355a878b6144e1565b90506135728261356a8582614508565b839190613d24565b975061357f888484613d24565b955083881180156135a4575085613596888a614508565b6135a09190614508565b8a10155b985050505050505b93509350935093565b80825d5050565b5f5f82600184516135cd91906144e1565b815181106135dd576135dd6144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa158015613643573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613667919061460d565b6001600160a01b0316816001600160a01b031603613683575060015b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f0702e8e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136e0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613704919061460d565b90505f5f5f836001600160a01b0316636332aec68a876040518363ffffffff1660e01b8152600401613737929190614628565b606060405180830381865afa158015613752573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137769190614738565b919450909250905082881161379e576040516326b56d9d60e21b815260040160405180910390fd5b5f846001600160a01b03166396ce07956040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137db573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137ff91906144b6565b90505f61380d8a8584613d24565b90508261381a828c614508565b6138249190614508565b9b9a5050505050505050505050565b5f61385560405180606001604052805f81526020015f81526020015f81525090565b5f5f856001875161386691906144e1565b81518110613876576138766144f4565b602002602001015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ad5c46486040518163ffffffff1660e01b8152600401602060405180830381865afa1580156138dc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613900919061460d565b6001600160a01b0316816001600160a01b0316149350836139215780613924565b60015b91505f84613997576040516370a0823160e01b81523060048201526001600160a01b038416906370a0823190602401602060405180830381865afa15801561396e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061399291906144b6565b613999565b475b905084158015613a375750604051636eb1769f60e11b81523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116602483015282919085169063dd62ed3e90604401602060405180830381865afa158015613a11573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a3591906144b6565b105b15613ace5760405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f19602483015284169063095ea7b3906044016020604051808303815f875af1158015613aa8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613acc9190614763565b505b613adc8c84838d8d8d613dda565b6040516315cd127f60e31b8152600481018e90529094505f92506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915063ae6893f890602401602060405180830381865afa158015613b46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b6a91906144b6565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635fd262de85613ba6575f613ba9565b84515b8d858e885f015189602001518a604001515f6001600160401b03811115613bd257613bd2613f72565b6040519080825280601f01601f191660200182016040528015613bfc576020820181803683370190505b506040518963ffffffff1660e01b8152600401613c1f97969594939291906147ff565b60206040518083038185885af1158015613c3b573d5f5f3e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190613c609190614763565b613c7d57604051632fba5fef60e21b815260040160405180910390fd5b336001600160a01b0316818c7f72bf47e60197f642b1bd182e1cbdf5a8f80f8eee36d2ba83c2f14abcea5c9eca8d8a8a604051613cbc93929190614871565b60405180910390a45050505050505050505050565b5f54600160a01b900460ff16610fd057604051638dfc202b60e01b815260040160405180910390fd5b5f54600160a01b900460ff1615610fd05760405163d93c066560e01b815260040160405180910390fd5b5f838302815f1985870982811083820303915050805f03613d5857838281613d4e57613d4e614532565b04925050506107aa565b808411613d6f57613d6f6003851502601118613eaa565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b613dfb60405180606001604052805f81526020015f81526020015f81525090565b5f613e078888886133d3565b604086015260208501528352905080613e33576040516326b56d9d60e21b815260040160405180910390fd5b8151851115613e5557604051637c3bec0960e11b815260040160405180910390fd5b8382602001511115613e7a5760405163de7539a960e01b815260040160405180910390fd5b8282604001511115613e9f5760405163cce42f3560e01b815260040160405180910390fd5b509695505050505050565b634e487b715f52806020526024601cfd5b828054828255905f5260205f20908101928215613f0e579160200282015b82811115613f0e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613ed9565b50613f1a929150613f35565b5090565b5080545f8255905f5260205f20908101906126f491905b5b80821115613f1a575f8155600101613f36565b5f5f5f60608486031215613f5b575f5ffd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715613fa857613fa8613f72565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613fd657613fd6613f72565b604052919050565b5f6001600160401b03821115613ff657613ff6613f72565b5060051b60200190565b6001600160a01b03811681146126f4575f5ffd5b5f82601f830112614023575f5ffd5b813561403661403182613fde565b613fae565b8082825260208201915060208360051b860101925085831115614057575f5ffd5b602085015b8381101561407d57803561406f81614000565b83526020928301920161405c565b5095945050505050565b5f5f60408385031215614098575f5ffd5b82356001600160401b038111156140ad575f5ffd5b6140b985828601614014565b95602094909401359450505050565b5f5f604083850312156140d9575f5ffd5b8235915060208301356001600160401b038111156140f5575f5ffd5b61410185828601614014565b9150509250929050565b5f8151808452602084019350602083015f5b8281101561413b57815186526020958601959091019060010161411d565b5093949350505050565b602081525f6107aa602083018461410b565b5f5f60408385031215614168575f5ffd5b823561417381614000565b915060208301356001600160401b038111156140f5575f5ffd5b5f5f5f5f5f5f5f5f610100898b0312156141a5575f5ffd5b8835975060208901356141b781614000565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b038111156141ed575f5ffd5b6141f98b828c01614014565b989b979a50959894979396929550929360e00135925050565b5f5f5f5f5f5f60c08789031215614227575f5ffd5b863561423281614000565b9550602087013561424281614000565b95989597505050506040840135936060810135936080820135935060a0909101359150565b5f5f60408385031215614278575f5ffd5b823561428381614000565b946020939093013593505050565b5f602082840312156142a1575f5ffd5b81356107aa81614000565b5f8151808452602084019350602083015f5b8281101561413b5781516001600160a01b03168652602095860195909101906001016142be565b602081525f6107aa60208301846142ac565b5f5f5f5f5f5f5f60e0888a03121561430d575f5ffd5b873561431881614000565b9650602088013561432881614000565b96999698505050506040850135946060810135946080820135945060a0820135935060c0909101359150565b608081525f614366608083018761410b565b6020830195909552506040810192909252606090910152919050565b5f5f5f5f5f5f5f60e0888a031215614398575f5ffd5b8735965060208801356143aa81614000565b955060408801359450606088013593506080880135925060a08801356001600160401b038111156143d9575f5ffd5b6143e58a828b01614014565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614410575f5ffd5b833592506020840135915060408401356001600160401b03811115614433575f5ffd5b61443f86828701614014565b9150509250925092565b60a081525f61445b60a083018861410b565b90508560208301528460408301528360608301528260808301529695505050505050565b5f5f60408385031215614490575f5ffd5b823561449b81614000565b915060208301356144ab81614000565b809150509250929050565b5f602082840312156144c6575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561099c5761099c6144cd565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561099c5761099c6144cd565b808202811582820484141761099c5761099c6144cd565b634e487b7160e01b5f52601260045260245ffd5b5f8261456057634e487b7160e01b5f52601260045260245ffd5b500490565b828152604060208201525f6107a760408301846142ac565b5f6020828403121561458d575f5ffd5b81516001600160401b038111156145a2575f5ffd5b8201601f810184136145b2575f5ffd5b80516145c061403182613fde565b8082825260208201915060208360051b8501019250868311156145e1575f5ffd5b6020840193505b828410156146035783518252602093840193909101906145e8565b9695505050505050565b5f6020828403121561461d575f5ffd5b81516107aa81614000565b9182526001600160a01b0316602082015260400190565b8051801515811461464e575f5ffd5b919050565b5f60e0828403128015614664575f5ffd5b5061466d613f86565b825161467881614000565b8152602083015161468881614000565b60208201526146996040840161463f565b60408201526146aa6060840161463f565b60608201526080838101519082015260a0808401519082015260c0928301519281019290925250919050565b80516001600160701b038116811461464e575f5ffd5b5f5f5f606084860312156146fe575f5ffd5b614707846146d6565b9250614715602085016146d6565b9150604084015163ffffffff8116811461472d575f5ffd5b809150509250925092565b5f5f5f6060848603121561474a575f5ffd5b5050815160208301516040909301519094929350919050565b5f60208284031215614773575f5ffd5b6107aa8261463f565b634e487b7160e01b5f52602160045260245ffd5b85815284602082015260a060408201525f6147ae60a08301866142ac565b6001600160a01b0394909416606083015250608001529392505050565b848152608060208201525f6147e360808301866142ac565b6001600160a01b03949094166040830152506060015292915050565b87815260018060a01b038716602082015260018060a01b03861660408201528460608201528360808201528260a082015260e060c08201525f82518060e0840152806020850161010085015e5f6101008285010152610100601f19601f83011684010191505098975050505050505050565b6001600160a01b03841681526060602082018190525f90614894908301856142ac565b8281036040840152614603818561410b56fe9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122054f4b24286abd2281838a36e562f7e27a1b813ce734071a8c623bc8adf20769164736f6c634300081c0033"

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

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) Bridge() (common.Address, error) {
	return _SwapBridgeRouter.Contract.Bridge(&_SwapBridgeRouter.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) Bridge() (common.Address, error) {
	return _SwapBridgeRouter.Contract.Bridge(&_SwapBridgeRouter.CallOpts)
}

// CrossChainID is a free data retrieval call binding the contract method 0x2c7b395f.
//
// Solidity: function crossChainID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) CrossChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "crossChainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossChainID is a free data retrieval call binding the contract method 0x2c7b395f.
//
// Solidity: function crossChainID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterSession) CrossChainID() (*big.Int, error) {
	return _SwapBridgeRouter.Contract.CrossChainID(&_SwapBridgeRouter.CallOpts)
}

// CrossChainID is a free data retrieval call binding the contract method 0x2c7b395f.
//
// Solidity: function crossChainID() view returns(uint256)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) CrossChainID() (*big.Int, error) {
	return _SwapBridgeRouter.Contract.CrossChainID(&_SwapBridgeRouter.CallOpts)
}

// CrossToken is a free data retrieval call binding the contract method 0x90f1b5a4.
//
// Solidity: function crossToken() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) CrossToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "crossToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossToken is a free data retrieval call binding the contract method 0x90f1b5a4.
//
// Solidity: function crossToken() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) CrossToken() (common.Address, error) {
	return _SwapBridgeRouter.Contract.CrossToken(&_SwapBridgeRouter.CallOpts)
}

// CrossToken is a free data retrieval call binding the contract method 0x90f1b5a4.
//
// Solidity: function crossToken() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) CrossToken() (common.Address, error) {
	return _SwapBridgeRouter.Contract.CrossToken(&_SwapBridgeRouter.CallOpts)
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

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapRouter() (common.Address, error) {
	return _SwapBridgeRouter.Contract.SwapRouter(&_SwapBridgeRouter.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) SwapRouter() (common.Address, error) {
	return _SwapBridgeRouter.Contract.SwapRouter(&_SwapBridgeRouter.CallOpts)
}

// TokenToPath is a free data retrieval call binding the contract method 0x3e42c51f.
//
// Solidity: function tokenToPath(address , uint256 ) view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCaller) TokenToPath(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SwapBridgeRouter.contract.Call(opts, &out, "tokenToPath", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToPath is a free data retrieval call binding the contract method 0x3e42c51f.
//
// Solidity: function tokenToPath(address , uint256 ) view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterSession) TokenToPath(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SwapBridgeRouter.Contract.TokenToPath(&_SwapBridgeRouter.CallOpts, arg0, arg1)
}

// TokenToPath is a free data retrieval call binding the contract method 0x3e42c51f.
//
// Solidity: function tokenToPath(address , uint256 ) view returns(address)
func (_SwapBridgeRouter *SwapBridgeRouterCallerSession) TokenToPath(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SwapBridgeRouter.Contract.TokenToPath(&_SwapBridgeRouter.CallOpts, arg0, arg1)
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

// SwapETHForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0xec7d9618.
//
// Solidity: function swapETHForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapETHForExactCrossTokensBridge(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapETHForExactCrossTokensBridge", sourceToken, to, amountOut, maxNetworkFee, maxExFee, deadline)
}

// SwapETHForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0xec7d9618.
//
// Solidity: function swapETHForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapETHForExactCrossTokensBridge(sourceToken common.Address, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapETHForExactCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, maxNetworkFee, maxExFee, deadline)
}

// SwapETHForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0xec7d9618.
//
// Solidity: function swapETHForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapETHForExactCrossTokensBridge(sourceToken common.Address, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapETHForExactCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, maxNetworkFee, maxExFee, deadline)
}

// SwapETHForExactTokensBridge is a paid mutator transaction binding the contract method 0xad921379.
//
// Solidity: function swapETHForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapETHForExactTokensBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapETHForExactTokensBridge", toChainID, to, amountOut, maxNetworkFee, maxExFee, path, deadline)
}

// SwapETHForExactTokensBridge is a paid mutator transaction binding the contract method 0xad921379.
//
// Solidity: function swapETHForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapETHForExactTokensBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapETHForExactTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, maxNetworkFee, maxExFee, path, deadline)
}

// SwapETHForExactTokensBridge is a paid mutator transaction binding the contract method 0xad921379.
//
// Solidity: function swapETHForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapETHForExactTokensBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapETHForExactTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactETHForCrossTokensBridge is a paid mutator transaction binding the contract method 0x38208f67.
//
// Solidity: function swapExactETHForCrossTokensBridge(address sourceToken, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapExactETHForCrossTokensBridge(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapExactETHForCrossTokensBridge", sourceToken, to, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactETHForCrossTokensBridge is a paid mutator transaction binding the contract method 0x38208f67.
//
// Solidity: function swapExactETHForCrossTokensBridge(address sourceToken, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapExactETHForCrossTokensBridge(sourceToken common.Address, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactETHForCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactETHForCrossTokensBridge is a paid mutator transaction binding the contract method 0x38208f67.
//
// Solidity: function swapExactETHForCrossTokensBridge(address sourceToken, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapExactETHForCrossTokensBridge(sourceToken common.Address, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactETHForCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactETHForTokensBridge is a paid mutator transaction binding the contract method 0x88ad9911.
//
// Solidity: function swapExactETHForTokensBridge(uint256 toChainID, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapExactETHForTokensBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapExactETHForTokensBridge", toChainID, to, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactETHForTokensBridge is a paid mutator transaction binding the contract method 0x88ad9911.
//
// Solidity: function swapExactETHForTokensBridge(uint256 toChainID, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapExactETHForTokensBridge(toChainID *big.Int, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactETHForTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactETHForTokensBridge is a paid mutator transaction binding the contract method 0x88ad9911.
//
// Solidity: function swapExactETHForTokensBridge(uint256 toChainID, address to, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapExactETHForTokensBridge(toChainID *big.Int, to common.Address, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactETHForTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForCrossTokensBridge is a paid mutator transaction binding the contract method 0xf0a85637.
//
// Solidity: function swapExactTokensForCrossTokensBridge(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapExactTokensForCrossTokensBridge(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapExactTokensForCrossTokensBridge", sourceToken, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactTokensForCrossTokensBridge is a paid mutator transaction binding the contract method 0xf0a85637.
//
// Solidity: function swapExactTokensForCrossTokensBridge(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapExactTokensForCrossTokensBridge(sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactTokensForCrossTokensBridge is a paid mutator transaction binding the contract method 0xf0a85637.
//
// Solidity: function swapExactTokensForCrossTokensBridge(address sourceToken, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapExactTokensForCrossTokensBridge(sourceToken common.Address, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, deadline)
}

// SwapExactTokensForETHBridge is a paid mutator transaction binding the contract method 0x7e5262b4.
//
// Solidity: function swapExactTokensForETHBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapExactTokensForETHBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapExactTokensForETHBridge", toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForETHBridge is a paid mutator transaction binding the contract method 0x7e5262b4.
//
// Solidity: function swapExactTokensForETHBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapExactTokensForETHBridge(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForETHBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForETHBridge is a paid mutator transaction binding the contract method 0x7e5262b4.
//
// Solidity: function swapExactTokensForETHBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) payable returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapExactTokensForETHBridge(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForETHBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForTokensBridge is a paid mutator transaction binding the contract method 0xa964ea54.
//
// Solidity: function swapExactTokensForTokensBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapExactTokensForTokensBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapExactTokensForTokensBridge", toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForTokensBridge is a paid mutator transaction binding the contract method 0xa964ea54.
//
// Solidity: function swapExactTokensForTokensBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapExactTokensForTokensBridge(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapExactTokensForTokensBridge is a paid mutator transaction binding the contract method 0xa964ea54.
//
// Solidity: function swapExactTokensForTokensBridge(uint256 toChainID, address to, uint256 amountIn, uint256 amountOutMin, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapExactTokensForTokensBridge(toChainID *big.Int, to common.Address, amountIn *big.Int, amountOutMin *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapExactTokensForTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountIn, amountOutMin, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0x5ce53779.
//
// Solidity: function swapTokensForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapTokensForExactCrossTokensBridge(opts *bind.TransactOpts, sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapTokensForExactCrossTokensBridge", sourceToken, to, amountOut, amountInMax, maxNetworkFee, maxExFee, deadline)
}

// SwapTokensForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0x5ce53779.
//
// Solidity: function swapTokensForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapTokensForExactCrossTokensBridge(sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, amountInMax, maxNetworkFee, maxExFee, deadline)
}

// SwapTokensForExactCrossTokensBridge is a paid mutator transaction binding the contract method 0x5ce53779.
//
// Solidity: function swapTokensForExactCrossTokensBridge(address sourceToken, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapTokensForExactCrossTokensBridge(sourceToken common.Address, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactCrossTokensBridge(&_SwapBridgeRouter.TransactOpts, sourceToken, to, amountOut, amountInMax, maxNetworkFee, maxExFee, deadline)
}

// SwapTokensForExactETHBridge is a paid mutator transaction binding the contract method 0xbb1024f8.
//
// Solidity: function swapTokensForExactETHBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapTokensForExactETHBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapTokensForExactETHBridge", toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactETHBridge is a paid mutator transaction binding the contract method 0xbb1024f8.
//
// Solidity: function swapTokensForExactETHBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapTokensForExactETHBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactETHBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactETHBridge is a paid mutator transaction binding the contract method 0xbb1024f8.
//
// Solidity: function swapTokensForExactETHBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapTokensForExactETHBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactETHBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactTokensBridge is a paid mutator transaction binding the contract method 0x235feb6a.
//
// Solidity: function swapTokensForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactor) SwapTokensForExactTokensBridge(opts *bind.TransactOpts, toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.contract.Transact(opts, "swapTokensForExactTokensBridge", toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactTokensBridge is a paid mutator transaction binding the contract method 0x235feb6a.
//
// Solidity: function swapTokensForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterSession) SwapTokensForExactTokensBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
}

// SwapTokensForExactTokensBridge is a paid mutator transaction binding the contract method 0x235feb6a.
//
// Solidity: function swapTokensForExactTokensBridge(uint256 toChainID, address to, uint256 amountOut, uint256 amountInMax, uint256 maxNetworkFee, uint256 maxExFee, address[] path, uint256 deadline) returns()
func (_SwapBridgeRouter *SwapBridgeRouterTransactorSession) SwapTokensForExactTokensBridge(toChainID *big.Int, to common.Address, amountOut *big.Int, amountInMax *big.Int, maxNetworkFee *big.Int, maxExFee *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _SwapBridgeRouter.Contract.SwapTokensForExactTokensBridge(&_SwapBridgeRouter.TransactOpts, toChainID, to, amountOut, amountInMax, maxNetworkFee, maxExFee, path, deadline)
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
