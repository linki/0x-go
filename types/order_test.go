package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/util"

	"github.com/stretchr/testify/suite"
)

type OrderSuite struct {
	suite.Suite
}

func (suite *OrderSuite) TestCalculateOrderHash() {
	for _, tt := range []struct {
		order    Order
		expected common.Hash
	}{
		// RadarRelay
		{
			Order{
				ExchangeContractAddress:    common.HexToAddress("0x4f833a24e1f95d70f028921e27040ca56e09ab0b"),
				Sender:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Maker:                      common.HexToAddress("0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c"),
				Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
				MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"), // incomplete
				TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"), // incomplete
				FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
				MakerTokenAmount:           util.StrToBig("12070000000000000"),
				TakerTokenAmount:           util.StrToBig("17000000000000000000"),
				MakerFee:                   common.Big0,
				TakerFee:                   common.Big0,
				ExpirationUnixTimestampSec: time.Unix(1570973282, 0).UTC(),
				Signature:                  "0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602",
				Salt:                       util.StrToBig("1562419682770"),
			},
			common.HexToHash("0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb"),
		},
		// // The Ocean X
		// {
		// 	Order{
		// 		ExchangeContractAddress:    common.HexToAddress("0x90fe2af704b34e0224bf2299c838e04d4dcf1364"),
		// 		Maker:                      common.HexToAddress("0x00a6d07f3530430f87e19c25d999b627f4fe32e2"),
		// 		Taker:                      common.HexToAddress("0x00ba938cc0df182c25108d7bf2ee3d37bce07513"),
		// 		MakerTokenAddress:          common.HexToAddress("0x6ff6c0ff1d68b964901f986d4c9fa3ac68346570"),
		// 		TakerTokenAddress:          common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c"),
		// 		FeeRecipient:               common.HexToAddress("0x88a64b5e882e5ad851bea5e7a3c8ba7c523fecbe"),
		// 		MakerTokenAmount:           util.StrToBig("1000000000"),
		// 		TakerTokenAmount:           util.StrToBig("1100000000"),
		// 		MakerFee:                   common.Big0,
		// 		TakerFee:                   common.Big0,
		// 		ExpirationUnixTimestampSec: time.Unix(1523097537, 0).UTC(),
		// 		Salt:                       "96779178608164233712795994683330674094398651784855349948764786357549104359274",
		// 	},
		// 	common.HexToHash("0x0bf9344a08234507fadfdc6040b4d941fdf82c89a5f8455508499e8b3d4739e7"),
		// },
	} {
		suite.Equal(tt.expected, tt.order.CalculateOrderHash())
	}
}

func (suite *OrderSuite) TestUnmarshal() {
	expectedOrders := []Order{
		{
			OrderHash:                  common.HexToHash("0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb"),
			ExchangeContractAddress:    common.HexToAddress("0x4f833a24e1f95d70f028921e27040ca56e09ab0b"),
			Sender:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Maker:                      common.HexToAddress("0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c"),
			Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MakerTokenAddress:          common.HexToAddress("0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
			TakerTokenAddress:          common.HexToAddress("0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498"),
			FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
			MakerTokenAmount:           util.StrToBig("12070000000000000"),
			TakerTokenAmount:           util.StrToBig("17000000000000000000"),
			MakerFee:                   common.Big0,
			TakerFee:                   common.Big0,
			ExpirationUnixTimestampSec: time.Unix(1570973282, 0).UTC(),
			Signature:                  "0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602",
			Salt:                       util.StrToBig("1562419682770"),
		},
	}

	for _, tt := range []struct {
		jsonStr string
	}{
		// JSON doesn't contain order hash (RadarRelay)

		{
			`[
					{
						"order": {
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
							"salt":                  "1562419682770"
					  },
						"metaData": {
					    "orderHash": "0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb"
					  }
					}
			]`,
		},
		// "ecSignature": {
		// 	"v": 28,
		// 	"r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
		// 	"s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709"
		// }

		// JSON contains order hash (The Ocean X)
		// {
		// 	`[
		// 			{
		// 				"orderHash":                  "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
		// 				"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
		// 				"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
		// 				"taker":                      "0x0000000000000000000000000000000000000000",
		// 				"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// 				"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
		// 				"feeRecipient":               "0xa258b39954cef5cb142fd567a46cddb31a670124",
		// 				"makerTokenAmount":           "18981000000000000",
		// 				"takerTokenAmount":           "19000000000000000000",
		// 				"makerFee":                   "0",
		// 				"takerFee":                   "0",
		// 				"expirationUnixTimestampSec": "1518201120",
		// 				"salt": "58600101225676680041453168589125977076540694791976419610199695339725548478315",
		// 				"ecSignature": {
		// 					"v": 28,
		// 					"r": "0x2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d9",
		// 					"s": "0x44ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde127709"
		// 				}
		// 			}
		// 	]`,
		// },
	} {
		orders := []Order{}

		err := json.Unmarshal([]byte(tt.jsonStr), &orders)
		suite.Require().NoError(err)

		suite.Equal(expectedOrders, orders)
	}
}

func (suite *OrderSuite) TestMarshal() {
	for _, tt := range []struct {
		orders   []Order
		expected []map[string]interface{}
	}{
		{
			[]Order{
				{
					OrderHash:                  common.HexToHash("0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb"),
					ExchangeContractAddress:    common.HexToAddress("0x4f833a24e1f95d70f028921e27040ca56e09ab0b"),
					Sender:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
					Maker:                      common.HexToAddress("0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c"),
					Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
					MakerTokenAddress:          common.HexToAddress("0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					TakerTokenAddress:          common.HexToAddress("0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498"),
					FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
					MakerTokenAmount:           util.StrToBig("12070000000000000"),
					TakerTokenAmount:           util.StrToBig("17000000000000000000"),
					MakerFee:                   common.Big0,
					TakerFee:                   common.Big0,
					ExpirationUnixTimestampSec: time.Unix(1570973282, 0).UTC(),
					Signature:                  "0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602",
					Salt:                       util.StrToBig("1562419682770"),
				},
			},
			[]map[string]interface{}{
				{
					"order": map[string]string{
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
		},
	} {
		orders, err := json.Marshal(tt.orders)
		suite.Require().NoError(err)

		expectedJSON, err := json.Marshal(tt.expected)
		suite.Require().NoError(err)

		suite.Equal(string(expectedJSON), string(orders))
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
