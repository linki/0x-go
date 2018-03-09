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
		// RadarRelay
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
		// The Ocean X
		{
			[]string{
				"--exchange-contract-address", "0x90fe2af704b34e0224bf2299c838e04d4dcf1364",
				"--maker", "0x00a6d07f3530430f87e19c25d999b627f4fe32e2",
				"--taker", "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
				"--maker-token-address", "0x6ff6c0ff1d68b964901f986d4c9fa3ac68346570",
				"--taker-token-address", "0xd0a1e359811322d97991e03f863a0c30c2cf029c",
				"--fee-recipient", "0x88a64b5e882e5ad851bea5e7a3c8ba7c523fecbe",
				"--maker-token-amount", "1000000000",
				"--taker-token-amount", "1100000000",
				"--maker-fee", "0",
				"--taker-fee", "0",
				"--expiration-unix-timestamp-sec", "1523097537",
				"--salt", "96779178608164233712795994683330674094398651784855349948764786357549104359274",
			},
			"0x0bf9344a08234507fadfdc6040b4d941fdf82c89a5f8455508499e8b3d4739e7\n",
		},
	} {
		args := append(
			[]string{"orders", "hash"},
			tt.flags...,
		)

		rootCmd.SetArgs(args)
		rootCmd.Execute()

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersHashSuite(t *testing.T) {
	suite.Run(t, new(OrdersHashSuite))
}
