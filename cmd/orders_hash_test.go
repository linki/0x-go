package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrdersHashSuite struct {
	suite.Suite
	console io.ReadWriter
}

func (suite *OrdersHashSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	ordersHashCmd.SetOutput(suite.console)
}

func (suite *OrdersHashSuite) TearDownTest() {
	ordersHashCmd.SetOutput(nil)
}

func (suite *OrdersHashSuite) TestOrdersHash() {
	for _, tt := range []struct {
		flags    []string
		expected string
	}{
		{
			[]string{
				"--exchange-contract-address", "exchange-contract-address",
				"--maker", "maker",
				"--taker", "taker",
				"--maker-token-address", "maker-token-address",
				"--taker-token-address", "taker-token-address",
				"--fee-recipient", "fee-recipient",
				"--maker-token-amount", "1",
				"--taker-token-amount", "2",
				"--maker-fee", "3",
				"--taker-fee", "4",
				"--expiration-unix-timestamp-sec", "5",
				"--salt", "6",
			},
			"0xb9522c02a6e351361a01403f7c140768baf9a7e5446beb1733bb0578f6612f94\n",
		},
		{
			[]string{
				"--exchange-contract-address", "0x12459c951127e0c374ff9105dda097662a027093",
				"--maker", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"--taker", "0x0000000000000000000000000000000000000000",
				"--maker-token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--taker-token-address", "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"--fee-recipient", "0xa258b39954cef5cb142fd567a46cddb31a670124",
				"--maker-token-amount", "18981000000000000",
				"--taker-token-amount", "19000000000000000000",
				"--maker-fee", "0",
				"--taker-fee", "0",
				"--expiration-unix-timestamp-sec", "1518201120",
				"--salt", "58600101225676680041453168589125977076540694791976419610199695339725548478315",
			},
			"0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942\n",
		},
	} {
		err := ordersHashCmd.ParseFlags(tt.flags)
		suite.Require().NoError(err)

		ordersHashCmd.Run(ordersHashCmd, nil)

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersHashSuite(t *testing.T) {
	suite.Run(t, new(OrdersHashSuite))
}
