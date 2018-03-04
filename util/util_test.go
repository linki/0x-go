package util

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilSuite struct {
	suite.Suite
}

func (suite *UtilSuite) TestStrToBig() {
	suite.Equal(big.NewInt(100), StrToBig("100"))
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}
