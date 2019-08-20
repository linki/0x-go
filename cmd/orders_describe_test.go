package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
)

type OrdersDescribeSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *OrdersDescribeSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	rootCmd.SetOutput(suite.console)
	suite.url = "http://127.0.0.1:8080"
}

func (suite *OrdersDescribeSuite) TearDownTest() {
	rootCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *OrdersDescribeSuite) TestOrdersDescribe() {
	for _, tt := range []struct {
		response map[string]interface{}
		flags    []string
		expected string
	}{
		{
			map[string]interface{}{
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
			[]string{
				"--relayer-url", suite.url,
				"--order-hash", "0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb",
			},
			`orderHash: 0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb
exchangeAddress: 0x4f833a24e1f95d70f028921e27040ca56e09ab0b
senderAddress: 0x0000000000000000000000000000000000000000
makerAddress: 0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c
takerAddress: 0x0000000000000000000000000000000000000000
makerAssetData: 0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
takerAssetData: 0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498
feeRecipientAddress: 0xa258b39954cef5cb142fd567a46cddb31a670124
makerAssetAmount: 12070000000000000
takerAssetAmount: 17000000000000000000
makerFee: 0
takerFee: 0
expirationTimeSeconds: 1570973282
signature: 0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602
salt: 1562419682770
`,
		},
	} {
		gock.New(suite.url).
			Get("/order/0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb").
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"orders", "describe"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersDescribeSuite(t *testing.T) {
	suite.Run(t, new(OrdersDescribeSuite))
}
