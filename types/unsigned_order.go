package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type UnsignedOrder struct {
	ExchangeContractAddress    common.Address
	Maker                      common.Address
	Taker                      common.Address
	MakerTokenAddress          common.Address
	TakerTokenAddress          common.Address
	MakerTokenAmount           *big.Int
	TakerTokenAmount           *big.Int
	ExpirationUnixTimestampSec time.Time
	Salt                       *big.Int
}

func (o *UnsignedOrder) MarshalJSON() ([]byte, error) {
	order := map[string]interface{}{
		"exchangeContractAddress":    o.ExchangeContractAddress,
		"maker":                      o.Maker,
		"taker":                      o.Taker,
		"makerTokenAddress":          o.MakerTokenAddress,
		"takerTokenAddress":          o.TakerTokenAddress,
		"makerTokenAmount":           o.MakerTokenAmount.String(),
		"takerTokenAmount":           o.TakerTokenAmount.String(),
		"expirationUnixTimestampSec": fmt.Sprintf("%d", o.ExpirationUnixTimestampSec.Unix()),
		"salt": o.Salt.String(),
	}
	return json.Marshal(order)
}
