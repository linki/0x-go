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
	rootCmd.SetOutput(suite.console)
	suite.url = "http://127.0.0.1:8080"
}

func (suite *OrdersListSuite) TearDownTest() {
	rootCmd.SetOutput(nil)
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
					"order": map[string]interface{}{
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
						"signature":             "0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602",
						"salt":                  "1562419682770",
					},
					"metaData": map[string]string{
						"orderHash": "0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb",
					},
				},
			},
			[]string{
				"--relayer-url", suite.url,
			},
			map[string]string{},
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb 17.000000 ZRX 0.012070 WETH2 0.000710 ZRX/WETH2\n",
		},
		// {
		// 	[]map[string]interface{}{
		// 		{
		// 			"orderHash":                  "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
		// 			"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
		// 			"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 			"taker":                      "0x0000000000000000000000000000000000000000",
		// 			"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 			"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
		// 			"feeRecipient":               "0xa258b39954cef5cb142fd567a46cddb31a670124",
		// 			"makerTokenAmount":           "18981000000000000",
		// 			"takerTokenAmount":           "19000000000000000000",
		// 			"makerFee":                   "0",
		// 			"takerFee":                   "0",
		// 			"expirationUnixTimestampSec": "1518201120",
		// 			"salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
		// 			"ecSignature": map[string]interface{}{
		// 				"v": 28,
		// 				"r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
		// 				"s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709",
		// 			},
		// 		},
		// 	},
		// 	[]string{
		// 		"--relayer-url", suite.url,
		// 		"--exchange-contract-address", "0x12459c951127e0c374ff9105dda097662a027093",
		// 		"--token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 		"--maker-token-address", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 		"--taker-token-address", "0xe41d2489571d322189246dafa5ebde1f4699f498",
		// 		"--maker", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 		"--taker", "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
		// 		"--trader", "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 		"--fee-recipient", "0xa258b39954cef5cb142fd567a46cddb31a670124",
		// 	},
		// 	map[string]string{
		// 		"exchangeContractAddress": "0x12459c951127e0c374ff9105dda097662a027093",
		// 		"tokenAddress":            "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 		"makerTokenAddress":       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 		"takerTokenAddress":       "0xe41d2489571d322189246dafa5ebde1f4699f498",
		// 		"maker":                   "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 		"taker":                   "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
		// 		"trader":                  "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 		"feeRecipient":            "0xa258b39954cef5cb142fd567a46cddb31a670124",
		// 	},
		// 	"0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942 19.000000 ZRX 0.018981 WETH2 0.000999 ZRX/WETH2\n",
		// },
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
