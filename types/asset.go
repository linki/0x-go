package types

import (
	"math"
	"math/big"
)

type Asset struct {
	AssetData string
	MinAmount string
	MaxAmount string
	Precision int

	Symbol string
	Digits int
}

var (
	UnknownAsset = Asset{Symbol: "UNKNOWN"}
)

func Price(base, quote *big.Float) *big.Float {
	return new(big.Float).Quo(quote, base)
}

func (t Asset) NormalizedValue(amount *big.Int) *big.Float {
	divisor := big.NewFloat(math.Pow10(t.Digits))
	return new(big.Float).Quo(new(big.Float).SetInt(amount), divisor)
}
