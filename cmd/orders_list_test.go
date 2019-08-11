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
				"--exchange-contract-address", "0x4f833a24e1f95d70f028921e27040ca56e09ab0b",
				"--token-address", "0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--maker-token-address", "0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"--taker-token-address", "0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498",
				"--maker", "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"--taker", "0x0000000000000000000000000000000000000000",
				"--trader", "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"--fee-recipient", "0xa258b39954cef5cb142fd567a46cddb31a670124",
			},
			map[string]string{
				"exchangeContractAddress": "0x4f833a24e1f95d70f028921e27040ca56e09ab0b",
				"assetAddress":            "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"makerTokenAddress":       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerTokenAddress":       "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"maker":                   "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"taker":                   "0x0000000000000000000000000000000000000000",
				"trader":                  "0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
				"feeRecipient":            "0xa258b39954cef5cb142fd567a46cddb31a670124",
			},
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb 17.000000 ZRX 0.012070 WETH2 0.000710 ZRX/WETH2\n",
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
