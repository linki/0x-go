package types

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Order struct {
	OrderHash                  common.Hash
	ExchangeContractAddress    common.Address
	Maker                      common.Address
	Taker                      common.Address
	MakerTokenAddress          common.Address
	TakerTokenAddress          common.Address
	FeeRecipient               common.Address
	MakerTokenAmount           *big.Int
	TakerTokenAmount           *big.Int
	MakerFee                   *big.Int
	TakerFee                   *big.Int
	ExpirationUnixTimestampSec time.Time
	Salt                       *big.Int
	Signature                  Signature
}

func (o *Order) CalculateOrderHash() common.Hash {
	sha := sha3.NewKeccak256()

	sha.Write(o.ExchangeContractAddress.Bytes())
	sha.Write(o.Maker.Bytes())
	sha.Write(o.Taker.Bytes())
	sha.Write(o.MakerTokenAddress.Bytes())
	sha.Write(o.TakerTokenAddress.Bytes())
	sha.Write(o.FeeRecipient.Bytes())
	sha.Write(common.BigToHash(o.MakerTokenAmount).Bytes())
	sha.Write(common.BigToHash(o.TakerTokenAmount).Bytes())
	sha.Write(common.BigToHash(o.MakerFee).Bytes())
	sha.Write(common.BigToHash(o.TakerFee).Bytes())
	sha.Write(common.BigToHash(big.NewInt(o.ExpirationUnixTimestampSec.Unix())).Bytes())
	sha.Write(common.BigToHash(o.Salt).Bytes())

	return common.BytesToHash(sha.Sum(nil))
}

func (o *Order) UnmarshalJSON(b []byte) error {
	order := map[string]interface{}{}

	err := json.Unmarshal(b, &order)
	if err != nil {
		return err
	}

	o.OrderHash = common.HexToHash(order["orderHash"].(string))
	o.ExchangeContractAddress = common.HexToAddress(order["exchangeContractAddress"].(string))
	o.Maker = common.HexToAddress(order["maker"].(string))
	o.Taker = common.HexToAddress(order["taker"].(string))
	o.MakerTokenAddress = common.HexToAddress(order["makerTokenAddress"].(string))
	o.TakerTokenAddress = common.HexToAddress(order["takerTokenAddress"].(string))
	o.FeeRecipient = common.HexToAddress(order["feeRecipient"].(string))

	o.MakerTokenAmount = new(big.Int)
	o.TakerTokenAmount = new(big.Int)
	o.MakerFee = new(big.Int)
	o.TakerFee = new(big.Int)
	o.Salt = new(big.Int)

	o.MakerTokenAmount.UnmarshalJSON([]byte(order["makerTokenAmount"].(string)))
	o.TakerTokenAmount.UnmarshalJSON([]byte(order["takerTokenAmount"].(string)))
	o.MakerFee.UnmarshalJSON([]byte(order["makerFee"].(string)))
	o.TakerFee.UnmarshalJSON([]byte(order["takerFee"].(string)))
	o.Salt.UnmarshalJSON([]byte(order["salt"].(string)))

	sig := order["ecSignature"].(map[string]interface{})
	o.Signature = Signature{
		V: byte(sig["v"].(float64)),
		R: common.HexToHash(sig["r"].(string)),
		S: common.HexToHash(sig["s"].(string)),
	}

	timestamp, err := strconv.ParseInt(order["expirationUnixTimestampSec"].(string), 10, 64)
	if err != nil {
		return err
	}
	o.ExpirationUnixTimestampSec = time.Unix(timestamp, 0)

	return nil
}
