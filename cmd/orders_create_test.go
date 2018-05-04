package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
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
				"--relayer-url", suite.url,
				"--keystore-file", suite.keystoreFile,
				"--passphrase", "not-secure-do-not-use-me-for-anything-else",
			},
			map[string]interface{}{
				"orderHash":                  "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
				"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
				"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"taker":                      "0x0000000000000000000000000000000000000000",
				"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"feeRecipient":               "0xa258b39954cef5cb142fd567a46cddb31a670124",
				"makerTokenAmount":           "18981000000000000",
				"takerTokenAmount":           "19000000000000000000",
				"makerFee":                   "0",
				"takerFee":                   "0",
				"expirationUnixTimestampSec": "1518201120",
				"salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
				"ecSignature": map[string]interface{}{
					"v": 28,
					"r": "0x5ac1007d026a3fe62b08262fac34daa5b311e72aac81b326c028c253bfc54284",
					"s": "0x25f08e4ba49abfd4f44205e9e01e1c552ad447747beac9b0308493b8127949c6",
				},
			},
			"0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942\n",
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

func (suite *OrdersCreateSuite) TestOrdersCreateWithFeeDetection() {
	for _, tt := range []struct {
		expectedFeesBody map[string]string
		feesResponse     map[string]string
		flags            []string
		expectedBody     map[string]interface{}
		expected         string
	}{
		{
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
				"--keystore-file", suite.keystoreFile,
				"--passphrase", "not-secure-do-not-use-me-for-anything-else",
				"--autodetect-fees", "true",
			},
			map[string]interface{}{
				"orderHash":                  "0xf52f1d9d197eb6ec0add6fcdd16ac99738c264f2f4e9ebba11c641fcda94dbf5",
				"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
				"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"taker":                      "0x0000000000000000000000000000000000000000",
				"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"feeRecipient":               "0xb046140686d052fff581f63f8136cce132e857da",
				"makerTokenAmount":           "18981000000000000",
				"takerTokenAmount":           "19000000000000000000",
				"makerFee":                   "100000000000000",
				"takerFee":                   "200000000000000",
				"expirationUnixTimestampSec": "1518201120",
				"salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
				"ecSignature": map[string]interface{}{
					"v": 28,
					"r": "0xce571ecc4be0bbb004a9b7274b9adf07b8e1d0b473d409d6c285aace3cf2c57f",
					"s": "0x66b2ed7465071f9a8c3ae998ea8cc816a7050931067a372fc1a1965b6a46463f",
				},
			},
			"0xf52f1d9d197eb6ec0add6fcdd16ac99738c264f2f4e9ebba11c641fcda94dbf5\n",
		},
	} {
		gock.New(suite.url).
			Post("/fees").
			JSON(tt.expectedFeesBody).
			Reply(http.StatusOK).
			JSON(tt.feesResponse)

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
