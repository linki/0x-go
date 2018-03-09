package util

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func EmptyAddress(a common.Address) bool {
	return common.EmptyHash(a.Hash())
}

func StrToBig(str string) *big.Int {
	bigInt, ok := new(big.Int).SetString(str, 10)
	if !ok {
		panic("StrToBig failed")
	}
	return bigInt
}
