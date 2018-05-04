package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type OrdersFeesSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *OrdersFeesSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	rootCmd.SetOutput(suite.console)
	suite.url = "http://127.0.0.1:8080"
}

func (suite *OrdersFeesSuite) TearDownTest() {
	rootCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *OrdersFeesSuite) TestOrdersFees() {
	for _, tt := range []struct {
		response     map[string]string
		flags        []string
		expectedBody map[string]string
		expected     string
	}{
		{
			map[string]string{
				"feeRecipient": "0xb046140686d052fff581f63f8136cce132e857da",
				"makerFee":     "100000000000000",
				"takerFee":     "200000000000000",
			},
			[]string{
				"--exchange-contract-address", "0x12459c951127e0c374ff9105dda097662a027093",
				"--maker", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"--taker", "0x0000000000000000000000000000000000000000",
				"--maker-token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--taker-token-address", "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"--maker-token-amount", "18981000000000000",
				"--taker-token-amount", "19000000000000000000",
				"--expiration-unix-timestamp-sec", "1518201120",
				"--salt", "58600101225676680041453168589125977076540694791976419610199695339725548478315",
				"--relayer-url", suite.url,
			},
			map[string]string{
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
			`feeRecipient: 0xb046140686d052fff581f63f8136cce132e857da
makerFee: 100000000000000
takerFee: 200000000000000
`,
		},
	} {
		gock.New(suite.url).
			Post("/fees").
			JSON(tt.expectedBody).
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"orders", "fees"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersFeesSuite(t *testing.T) {
	suite.Run(t, new(OrdersFeesSuite))
}
