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
	rootCmd.SetOutput(suite.console)
}

func (suite *OrdersHashSuite) TearDownTest() {
	rootCmd.SetOutput(nil)
}

func (suite *OrdersHashSuite) TestOrdersHash() {
	for _, tt := range []struct {
		flags    []string
		expected string
	}{
		{
			[]string{
				"--exchange-contract-address", "0x4f833a24e1f95d70f028921e27040ca56e09ab0b",
				"--sender", "0x0000000000000000000000000000000000000000",
				"--maker", "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"--taker", "0x0000000000000000000000000000000000000000",
				"--maker-token-address", "0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--taker-token-address", "0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498",
				"--fee-recipient", "0xa258b39954cef5cb142fd567a46cddb31a670124",
				"--maker-token-amount", "12070000000000000",
				"--taker-token-amount", "17000000000000000000",
				"--maker-fee", "0",
				"--taker-fee", "0",
				"--expiration-unix-timestamp-sec", "1570973282",
				"--salt", "1562419682770",
			},
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb\n",
		},
	} {
		args := append(
			[]string{"orders", "hash"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersHashSuite(t *testing.T) {
	suite.Run(t, new(OrdersHashSuite))
}
