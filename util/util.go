package util

import (
	"math/big"
)

func StrToBig(str string) *big.Int {
	bigInt, ok := new(big.Int).SetString(str, 10)
	if !ok {
		panic("StrToBig failed")
	}
	return bigInt
}
