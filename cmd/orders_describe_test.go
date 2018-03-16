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

type OrdersDescribeSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *OrdersDescribeSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	ordersDescribeCmd.SetOutput(suite.console)

	suite.url = "http://127.0.0.1:8080"
}

func (suite *OrdersDescribeSuite) TearDownTest() {
	ordersDescribeCmd.SetOutput(nil)
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
			[]string{
				"--relayer-url", suite.url,
				"--order-hash", "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
			},
			`orderHash: 0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942
exchange-contract-address: 0x12459c951127e0c374ff9105dda097662a027093
maker: 0xc9b32e9563fe99612ce3a2695ac2a6404c111dde
taker: 0x0000000000000000000000000000000000000000
maker-token-address: 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
taker-token-address: 0xe41d2489571d322189246dafa5ebde1f4699f498
fee-recipient: 0xa258b39954cef5cb142fd567a46cddb31a670124
maker-token-amount: 18981000000000000
taker-token-amount: 19000000000000000000
maker-fee: 0
taker-fee: 0
expiration-unix-timestamp-sec: 1518201120
salt: 58600101225676680041453168589125977076540694791976419610199695339725548478315
v: 28
r: 0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9
s: 0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709
`,
		},
	} {
		gock.New(suite.url).
			Get("/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942").
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"orders", "describe"},
			tt.flags...,
		)

		rootCmd.SetArgs(args)
		rootCmd.Execute()

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersDescribeSuite(t *testing.T) {
	suite.Run(t, new(OrdersDescribeSuite))
}
