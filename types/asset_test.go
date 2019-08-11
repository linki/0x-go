package types

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/suite"
)

type AssetSuite struct {
	suite.Suite
}

func (suite *AssetSuite) TestPrice() {
	for _, tt := range []struct {
		baseAssetAmount   *big.Int
		baseAssetDigits   int
		quoteAssetAmount  *big.Int
		quoteAssetdDigits int
		expected          *big.Float
	}{
		{
			big.NewInt(100),
			2,
			big.NewInt(100),
			2,
			big.NewFloat(1),
		},
		{
			big.NewInt(100),
			2,
			big.NewInt(1),
			2,
			big.NewFloat(0.01),
		},
		{
			big.NewInt(1),
			2,
			big.NewInt(100),
			2,
			big.NewFloat(100),
		},
		{
			big.NewInt(1),
			0,
			big.NewInt(100),
			2,
			big.NewFloat(1),
		},
		{
			big.NewInt(100),
			2,
			big.NewInt(1),
			0,
			big.NewFloat(1),
		},
		{
			big.NewInt(1),
			0,
			big.NewInt(1),
			2,
			big.NewFloat(0.01),
		},
		{
			big.NewInt(1),
			2,
			big.NewInt(1),
			0,
			big.NewFloat(100),
		},
	} {
		suite.Equal(
			fmt.Sprintf("%.18f", tt.expected),
			fmt.Sprintf("%.18f", Price(
				Asset{Digits: tt.baseAssetDigits}.NormalizedValue(tt.baseAssetAmount),
				Asset{Digits: tt.quoteAssetdDigits}.NormalizedValue(tt.quoteAssetAmount),
			)),
		)
	}
}

func (suite *AssetSuite) TestAssetNormalizedValue() {
	for _, tt := range []struct {
		amount   *big.Int
		digits   int
		expected *big.Float
	}{
		{
			amount:   common.Big0,
			digits:   0,
			expected: big.NewFloat(0),
		},
		{
			amount:   common.Big0,
			digits:   2,
			expected: big.NewFloat(0),
		},
		{
			amount:   big.NewInt(100),
			digits:   0,
			expected: big.NewFloat(100),
		},
		{
			amount:   big.NewInt(100),
			digits:   2,
			expected: big.NewFloat(1),
		},
		{
			amount:   big.NewInt(100),
			digits:   4,
			expected: big.NewFloat(0.01),
		},
	} {
		asset := Asset{Digits: tt.digits}
		suite.Equal(
			fmt.Sprintf("%.18f", tt.expected),
			fmt.Sprintf("%.18f", asset.NormalizedValue(tt.amount)),
		)
	}
}

func TestAssetSuite(t *testing.T) {
	suite.Run(t, new(AssetSuite))
}
