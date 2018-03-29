package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Fees struct {
	FeeRecipient common.Address
	MakerFee     *big.Int
	TakerFee     *big.Int
}

func (f *Fees) UnmarshalJSON(b []byte) error {
	fees := map[string]string{}

	err := json.Unmarshal(b, &fees)
	if err != nil {
		return err
	}

	f.FeeRecipient = common.HexToAddress(fees["feeRecipient"])

	f.MakerFee = new(big.Int)
	f.TakerFee = new(big.Int)

	f.MakerFee.UnmarshalJSON([]byte(fees["makerFee"]))
	f.TakerFee.UnmarshalJSON([]byte(fees["takerFee"]))

	return nil
}
