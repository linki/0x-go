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

type OrdersListSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *OrdersListSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	ordersListCmd.SetOutput(suite.console)

	suite.url = "http://127.0.0.1:8080"
}

func (suite *OrdersListSuite) TearDownTest() {
	ordersListCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *OrdersListSuite) TestOrdersList() {
	for _, tt := range []struct {
		response       []map[string]interface{}
		flags          []string
		expectedParams map[string]string
		expected       string
	}{
		{
			[]map[string]interface{}{
				{
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
						"r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
						"s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709",
					},
				},
			},
			[]string{
				"--relayer-url", suite.url,
			},
			map[string]string{},
			"0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942\n",
		},
		{
			[]map[string]interface{}{
				{
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
						"r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
						"s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709",
					},
				},
			},
			[]string{
				"--relayer-url", suite.url,
				"--exchange-contract-address", "0x12459c951127e0c374ff9105dda097662a027093",
				"--token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--maker-token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--taker-token-address", "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"--maker", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"--taker", "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
				"--trader", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"--fee-recipient", "0xa258b39954cef5cb142fd567a46cddb31a670124",
			},
			map[string]string{
				"exchangeContractAddress": "0x12459c951127e0c374ff9105dda097662a027093",
				"tokenAddress":            "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"makerTokenAddress":       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerTokenAddress":       "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"maker":                   "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"taker":                   "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
				"trader":                  "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"feeRecipient":            "0xa258b39954cef5cb142fd567a46cddb31a670124",
			},
			"0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942\n",
		},
	} {
		gock.New(suite.url).
			Get("/orders").
			MatchParams(tt.expectedParams).
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"orders", "list"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersListSuite(t *testing.T) {
	suite.Run(t, new(OrdersListSuite))
}
