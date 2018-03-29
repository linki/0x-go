package types

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/util"

	"github.com/stretchr/testify/suite"
)

type FeesSuite struct {
	suite.Suite
}

func (suite *FeesSuite) TestUnmarshal() {
	for _, tt := range []struct {
		fees  map[string]string
		expected Fees
	}{
		{
			map[string]string{
				"feeRecipient": "0xb046140686d052fff581f63f8136cce132e857da",
				"makerFee":     "100000000000000",
				"takerFee":     "200000000000000",
			},
			Fees{
				FeeRecipient: common.HexToAddress("0xb046140686d052fff581f63f8136cce132e857da"),
				MakerFee:     util.StrToBig("100000000000000"),
				TakerFee:     util.StrToBig("200000000000000"),
			},
		},
	} {
		feesJSON, err := json.Marshal(tt.fees)
		suite.Require().NoError(err)

		fees := Fees{}

		err = json.Unmarshal(feesJSON, &fees)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, fees)
	}
}

func TestFeesSuite(t *testing.T) {
	suite.Run(t, new(FeesSuite))
}
