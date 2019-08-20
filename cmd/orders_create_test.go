package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
)

type OrdersCreateSuite struct {
	suite.Suite
	console      io.ReadWriter
	url          string
	keystoreFile string
}

func (suite *OrdersCreateSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	rootCmd.SetOutput(suite.console)
	suite.url = "http://127.0.0.1:8080"
	suite.keystoreFile = setupTestKeystoreFile(suite.Require())
}

func (suite *OrdersCreateSuite) TearDownTest() {
	os.Remove(suite.keystoreFile)
	rootCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *OrdersCreateSuite) TestOrdersCreate() {
	for _, tt := range []struct {
		flags        []string
		expectedBody map[string]interface{}
		expected     string
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
				"--relayer-url", suite.url,
				"--keystore-file", suite.keystoreFile,
				"--passphrase", "not-secure-do-not-use-me-for-anything-else",
			},
			map[string]interface{}{
				"exchangeAddress":       "0x4f833a24e1f95d70f028921e27040ca56e09ab0b",
				"senderAddress":         "0x0000000000000000000000000000000000000000",
				"makerAddress":          "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"takerAddress":          "0x0000000000000000000000000000000000000000",
				"makerAssetData":        "0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerAssetData":        "0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498",
				"feeRecipientAddress":   "0xa258b39954cef5cb142fd567a46cddb31a670124",
				"makerAssetAmount":      "12070000000000000",
				"takerAssetAmount":      "17000000000000000000",
				"makerFee":              "0",
				"takerFee":              "0",
				"expirationTimeSeconds": "1570973282",
				"signature":             "0x1b8b35942cff4f0a4008170e31e851be197a98febaa54dc0b8249d718eacadd774321946510517717df90f718b36af5f46423a4fc0105e5dbf175a99f018bd933202",
				"salt":                  "1562419682770",
			},
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb\n",
		},
	} {
		gock.New(suite.url).
			Post("/order").
			JSON(tt.expectedBody).
			Reply(http.StatusCreated)

		args := append(
			[]string{"orders", "create"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersCreateSuite(t *testing.T) {
	suite.Run(t, new(OrdersCreateSuite))
}
