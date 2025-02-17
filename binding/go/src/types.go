package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// IBridgeFeeManagerFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFeeManagerFeeInfo struct {
	Token      common.Address
	GasFee     *big.Int
	ServiceFee *big.Int
}

// IBridgeStandardFinalizeArguments is an auto generated low-level Go binding around an user-defined struct.
type IBridgeStandardFinalizeArguments struct {
	Index     *big.Int
	Token     common.Address
	To        common.Address
	Value     *big.Int
	ExtraData [][]byte
}

// IBridgeStandardPermitArguments is an auto generated low-level Go binding around an user-defined struct.
type IBridgeStandardPermitArguments struct {
	Token    common.Address
	Account  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}
