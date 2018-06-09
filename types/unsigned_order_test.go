package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/util"

	"github.com/stretchr/testify/suite"
)

type UnsignedOrderSuite struct {
	suite.Suite
}

func (suite *UnsignedOrderSuite) TestMarshal() {
	for _, tt := range []struct {
		orders   []UnsignedOrder
		expected []map[string]string
	}{
		{
			[]UnsignedOrder{
				{
					Maker:                   common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
					Taker:                   common.HexToAddress("0x0000000000000000000000000000000000000000"),
					MakerTokenAddress:       common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					TakerTokenAddress:       common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
					ExchangeContractAddress: common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
					Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
					MakerTokenAmount:           util.StrToBig("18981000000000000"),
					TakerTokenAmount:           util.StrToBig("19000000000000000000"),
					ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
				},
			},
			[]map[string]string{
				{
					"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
					"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
					"taker":                      "0x0000000000000000000000000000000000000000",
					"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
					"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
					"makerTokenAmount":           "18981000000000000",
					"takerTokenAmount":           "19000000000000000000",
					"expirationUnixTimestampSec": "1518201120",
					"salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
				},
			},
		},
	} {
		orders, err := json.Marshal(tt.orders)
		suite.Require().NoError(err)

		expectedJSON, err := json.Marshal(tt.expected)
		suite.Require().NoError(err)

		suite.Equal(string(expectedJSON), string(orders))
	}
}

func TestUnsignedOrderSuite(t *testing.T) {
	suite.Run(t, new(UnsignedOrderSuite))
}
