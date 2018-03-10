package util

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/suite"
)

type UtilSuite struct {
	suite.Suite
}

func (suite *UtilSuite) TestEmptyAddress() {
	suite.True(EmptyAddress(common.HexToAddress("0x0000000000000000000000000000000000000000")))
	suite.False(EmptyAddress(common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")))
}

func (suite *UtilSuite) TestStrToBig() {
	suite.Equal(big.NewInt(100), StrToBig("100"))
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}
