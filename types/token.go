package types

import (
	"math"
	"math/big"
)

type Token struct {
	Address   string
	MinAmount string
	MaxAmount string
	Precision int
	Symbol    string
	Digits    int
}

var (
	UnknownToken = Token{Symbol: "UNKNOWN"}
)

func Price(base, quote *big.Float) *big.Float {
	return new(big.Float).Quo(quote, base)
}

func (t Token) NormalizedValue(amount *big.Int) *big.Float {
	divisor := big.NewFloat(math.Pow10(t.Digits))
	return new(big.Float).Quo(new(big.Float).SetInt(amount), divisor)
}
