package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Order struct {
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
