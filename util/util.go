package util

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func EmptyAddress(a common.Address) bool {
	return a.Hash() == common.Hash{}
}

func StrToBig(str string) *big.Int {
	bigInt, ok := new(big.Int).SetString(str, 10)
	if !ok {
		panic("StrToBig failed")
	}
	return bigInt
}

func EthToWei(eth int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(eth), big.NewInt(1000000000000000000))
}
