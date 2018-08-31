package types

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/suite"
)

type TokenSuite struct {
	suite.Suite
}

func (suite *TokenSuite) TestPrice() {
	for _, tt := range []struct {
		baseTokenAmount   *big.Int
		baseTokenDigits   int
		quoteTokenAmount  *big.Int
		quoteTokendDigits int
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
				Token{Digits: tt.baseTokenDigits}.NormalizedValue(tt.baseTokenAmount),
				Token{Digits: tt.quoteTokendDigits}.NormalizedValue(tt.quoteTokenAmount),
			)),
		)
	}
}

func (suite *TokenSuite) TestTokenNormalizedValue() {
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
		token := Token{Digits: tt.digits}
		suite.Equal(
			fmt.Sprintf("%.18f", tt.expected),
			fmt.Sprintf("%.18f", token.NormalizedValue(tt.amount)),
		)
	}
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}
