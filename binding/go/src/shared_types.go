// Code generated - DO NOT EDIT.
// This file contains shared type definitions.

package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IBridgeRegistryPendingData struct {
	Args            IBridgeRegistryFinalizeArguments
	Status          uint8
	DelayExpiration *big.Int
}


type IBridgeRegistryTokenPair struct {
	LocalToken    common.Address
	RemoteToken   common.Address
	IsOrigin      bool
	Paused        bool
	PendingAmount *big.Int
	Deposited     *big.Int
	Minted        *big.Int
}


type IBaseBridgePermitArguments struct {
	Token    common.Address
	Account  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}


type IBridgeRegistryFinalizeArguments struct {
	FromChainID *big.Int
	Index       *big.Int
	ToToken     common.Address
	To          common.Address
	Value       *big.Int
	ExtraData   []byte
}


type IBaseBridgeBridgeTokenArguments struct {
	ToChainID  *big.Int
	FromToken  common.Address
	From       common.Address
	To         common.Address
	Value      *big.Int
	NetworkFee *big.Int
	ExFee      *big.Int
	ExtraData  []byte
}


