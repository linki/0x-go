package relayer

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client      *Client
	url         string
	assetA      types.Asset
	assetB      types.Asset
	assetPairAB types.AssetPair
}

func (suite *ClientSuite) SetupTest() {
	suite.url = "http://127.0.0.1:8080"
	suite.client = NewClient(suite.url)
	suite.assetA = types.Asset{
		AssetData: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
		MinAmount: "0",
		MaxAmount: "10000000000000000000",
		Precision: 5,
	}
	suite.assetB = types.Asset{
		AssetData: "0xef7fff64389b814a946f3e92105513705ca6b990",
		MinAmount: "0",
		MaxAmount: "50000000000000000000",
		Precision: 5,
	}
	suite.assetPairAB = types.AssetPair{
		AssetDataA: suite.assetA,
		AssetDataB: suite.assetB,
	}
}

func (suite *ClientSuite) TearDownTest() {
	suite.True(gock.IsDone())
	gock.Off()
}

// GET /asset_pairs

func (suite *ClientSuite) TestGetAssetPairs() {
	gock.New(suite.url).
		Get("/asset_pairs").
		Reply(http.StatusOK).
		JSON(map[string]interface{}{
			"records": []map[string]interface{}{
				{
					"assetDataA": map[string]interface{}{
						"assetData": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
						"minAmount": "0",
						"maxAmount": "10000000000000000000",
						"precision": 5,
					},
					"assetDataB": map[string]interface{}{
						"assetData": "0xef7fff64389b814a946f3e92105513705ca6b990",
						"minAmount": "0",
						"maxAmount": "50000000000000000000",
						"precision": 5,
					},
				},
			},
		})

	assetPairs, err := suite.client.GetAssetPairs(context.Background(), GetAssetPairsOpts{})
	suite.NoError(err)

	suite.Require().Len(assetPairs.Records, 1)
	suite.Equal(suite.assetPairAB, assetPairs.Records[0])
}

func (suite *ClientSuite) TestGetAssetPairsWithFilters() {
	gock.New(suite.url).
		Get("/asset_pairs").
		MatchParams(map[string]string{
			"assetDataA": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			"assetDataB": "0xef7fff64389b814a946f3e92105513705ca6b990",
		}).
		Reply(http.StatusOK).
		JSON(map[string][]map[string]interface{}{})

	assetPairs, err := suite.client.GetAssetPairs(context.Background(), GetAssetPairsOpts{
		AssetDataA: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
		AssetDataB: "0xef7fff64389b814a946f3e92105513705ca6b990",
	})
	suite.Require().NoError(err)

	suite.Len(assetPairs.Records, 0)
}

func (suite *ClientSuite) TestGetAssetPairsWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Get("/asset_pairs").
		Reply(http.StatusServiceUnavailable)

	_, err := suite.client.GetAssetPairs(context.Background(), GetAssetPairsOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestGetAssetPairsWithMalformedJSONResponse() {
	gock.New(suite.url).
		Get("/asset_pairs").
		Reply(http.StatusOK).
		BodyString("//\\")

	_, err := suite.client.GetAssetPairs(context.Background(), GetAssetPairsOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "error parsing json response")
}

func (suite *ClientSuite) TestGetAssetPairsWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := suite.client.GetAssetPairs(ctx, GetAssetPairsOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

// GET /orders

func (suite *ClientSuite) TestGetOrders() {
	gock.New(suite.url).
		Get("/orders").
		Reply(http.StatusOK).
		JSON([]map[string]interface{}{
			{
				"order": map[string]interface{}{
					"exchangeAddress":       "0x12459c951127e0c374ff9105dda097662a027093",
					"senderAddress":         "0x0000000000000000000000000000000000000000",
					"makerAddress":          "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
					"takerAddress":          "0x0000000000000000000000000000000000000000",
					"makerAssetData":        "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
					"takerAssetData":        "0xe41d2489571d322189246dafa5ebde1f4699f498",
					"feeRecipientAddress":   "0xa258b39954cef5cb142fd567a46cddb31a670124",
					"makerAssetAmount":      "18981000000000000",
					"takerAssetAmount":      "19000000000000000000",
					"makerFee":              "0",
					"takerFee":              "0",
					"expirationTimeSeconds": "1518201120",
					"signature":             "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
					"salt":                  "58600101225676680041453168589125977076540694791976419610199695339725548478315",
				},
				"metaData": map[string]string{
					"orderHash": "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
				},
			},
		})

	order := types.Order{
		OrderHash:                  common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"),
		Maker:                      common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
		MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		ExchangeContractAddress:    common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		MakerTokenAmount:           util.StrToBig("18981000000000000"),
		TakerTokenAmount:           util.StrToBig("19000000000000000000"),
		ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
		Signature:                  "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
	}

	orders, err := suite.client.GetOrders(context.Background(), GetOrdersOpts{})
	suite.NoError(err)

	suite.Require().Len(orders, 1)
	suite.Equal(order, orders[0])
}

func (suite *ClientSuite) TestGetOrdersWithFilters() {
	gock.New(suite.url).
		Get("/orders").
		MatchParams(map[string]string{
			"exchangeContractAddress": "0x12459c951127e0c374ff9105dda097662a027093",
			"assetAddress":            "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"makerTokenAddress":       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"takerTokenAddress":       "0xe41d2489571d322189246dafa5ebde1f4699f498",
			"maker":                   "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
			"taker":                   "0x00ba938cc0df182c25108d7bf2ee3d37bce07513",
			"trader":                  "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
			"feeRecipient":            "0xa258b39954cef5cb142fd567a46cddb31a670124",
		}).
		Reply(http.StatusOK).
		JSON([]map[string]interface{}{})

	orders, err := suite.client.GetOrders(context.Background(), GetOrdersOpts{
		ExchangeContractAddress: common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		TokenAddress:            common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		MakerTokenAddress:       common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:       common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		Maker:                   common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                   common.HexToAddress("0x00ba938cc0df182c25108d7bf2ee3d37bce07513"),
		Trader:                  common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		FeeRecipient:            common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
	})
	suite.NoError(err)

	suite.Len(orders, 0)
}

func (suite *ClientSuite) TestGetOrdersWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Get("/orders").
		Reply(http.StatusServiceUnavailable)

	_, err := suite.client.GetOrders(context.Background(), GetOrdersOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestGetOrdersWithMalformedJSONResponse() {
	gock.New(suite.url).
		Get("/orders").
		Reply(http.StatusOK).
		BodyString("//\\")

	_, err := suite.client.GetOrders(context.Background(), GetOrdersOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "error parsing json response")
}

func (suite *ClientSuite) TestGetOrdersWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := suite.client.GetOrders(ctx, GetOrdersOpts{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

// GET /orders/orderHash

func (suite *ClientSuite) TestGetOrder() {
	gock.New(suite.url).
		Get("/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942").
		Reply(http.StatusOK).
		JSON(map[string]interface{}{
			"order": map[string]interface{}{
				"exchangeAddress":       "0x12459c951127e0c374ff9105dda097662a027093",
				"senderAddress":         "0x0000000000000000000000000000000000000000",
				"makerAddress":          "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
				"takerAddress":          "0x0000000000000000000000000000000000000000",
				"makerAssetData":        "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				"takerAssetData":        "0xe41d2489571d322189246dafa5ebde1f4699f498",
				"feeRecipientAddress":   "0xa258b39954cef5cb142fd567a46cddb31a670124",
				"makerAssetAmount":      "18981000000000000",
				"takerAssetAmount":      "19000000000000000000",
				"makerFee":              "0",
				"takerFee":              "0",
				"expirationTimeSeconds": "1518201120",
				"signature":             "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
				"salt":                  "58600101225676680041453168589125977076540694791976419610199695339725548478315",
			},
			"metaData": map[string]string{
				"orderHash": "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942",
			},
		})

	expectedOrder := types.Order{
		OrderHash:                  common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"),
		Sender:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Maker:                      common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
		MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		ExchangeContractAddress:    common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		MakerTokenAmount:           util.StrToBig("18981000000000000"),
		TakerTokenAmount:           util.StrToBig("19000000000000000000"),
		ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
		Signature:                  "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
	}

	order, err := suite.client.GetOrder(context.Background(), common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"))
	suite.NoError(err)

	suite.Equal(expectedOrder, order)
}

func (suite *ClientSuite) TestGetOrderOrderNotFound() {
	gock.New(suite.url).
		Get("/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942").
		Reply(http.StatusNotFound).
		JSON(map[string]interface{}{})

	_, err := suite.client.GetOrder(context.Background(), common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"))
	suite.Require().Error(err)
	suite.Contains(err.Error(), "order not found")
	suite.Contains(err.Error(), "0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942")
}

func (suite *ClientSuite) TestGetOrderWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Get("/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942").
		Reply(http.StatusServiceUnavailable)

	_, err := suite.client.GetOrder(context.Background(), common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"))
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestGetOrderWithMalformedJSONResponse() {
	gock.New(suite.url).
		Get("/order/0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942").
		Reply(http.StatusOK).
		BodyString("//\\")

	_, err := suite.client.GetOrder(context.Background(), common.HexToHash("0x10d750751d98bc8a9c29542118fbcf2fdb5b4977a3e5abf7cf38d03a6c149942"))
	suite.Require().Error(err)
	suite.Contains(err.Error(), "error parsing json response")
}

func (suite *ClientSuite) TestGetOrderWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := suite.client.GetOrder(ctx, common.Hash{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

// POST /order

func (suite *ClientSuite) TestCreateOrder() {
	gock.New(suite.url).
		Post("/order").
		JSON(map[string]interface{}{
			"exchangeAddress":       "0x12459c951127e0c374ff9105dda097662a027093",
			"senderAddress":         "0x0000000000000000000000000000000000000000",
			"makerAddress":          "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
			"takerAddress":          "0x0000000000000000000000000000000000000000",
			"makerAssetData":        "0xf47261b0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"takerAssetData":        "0xf47261b0000000000000000000000000e41d2489571d322189246dafa5ebde1f4699f498",
			"feeRecipientAddress":   "0xa258b39954cef5cb142fd567a46cddb31a670124",
			"makerAssetAmount":      "18981000000000000",
			"takerAssetAmount":      "19000000000000000000",
			"makerFee":              "0",
			"takerFee":              "0",
			"expirationTimeSeconds": "1518201120",
			"signature":             "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
			"salt":                  "58600101225676680041453168589125977076540694791976419610199695339725548478315",
		}).
		Reply(http.StatusCreated)

	order := types.Order{
		Sender:                     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		Maker:                      common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
		MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		ExchangeContractAddress:    common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		MakerTokenAmount:           util.StrToBig("18981000000000000"),
		TakerTokenAmount:           util.StrToBig("19000000000000000000"),
		ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
		Signature:                  "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
	}

	err := suite.client.CreateOrder(context.Background(), order)
	suite.NoError(err)
}

func (suite *ClientSuite) TestCreateOrderWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Post("/order").
		Reply(http.StatusServiceUnavailable)

	order := types.Order{
		Maker:                      common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		FeeRecipient:               common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
		MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		ExchangeContractAddress:    common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		MakerTokenAmount:           util.StrToBig("18981000000000000"),
		TakerTokenAmount:           util.StrToBig("19000000000000000000"),
		ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
		Signature:                  "0x1c2ffe986adb2ba48a800fe153ec0ec2af8b65856a34a67648e65a4bd6639c54d944ea4220aec0676a41ae7d0bc2433407f2ce892217be30e39d4e44dcde12770903",
	}

	err := suite.client.CreateOrder(context.Background(), order)
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestCreateOrderWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := suite.client.CreateOrder(ctx, types.Order{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

// POST /fees

func (suite *ClientSuite) TestGetFees() {
	gock.New(suite.url).
		Post("/fees").
		JSON(map[string]interface{}{
			"exchangeContractAddress":    "0x12459c951127e0c374ff9105dda097662a027093",
			"maker":                      "0xc9b32e9563fe99612ce3a2695ac2a6404c111dde",
			"taker":                      "0x0000000000000000000000000000000000000000",
			"makerTokenAddress":          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			"takerTokenAddress":          "0xe41d2489571d322189246dafa5ebde1f4699f498",
			"makerTokenAmount":           "18981000000000000",
			"takerTokenAmount":           "19000000000000000000",
			"expirationUnixTimestampSec": "1518201120",
			"salt":                       "58600101225676680041453168589125977076540694791976419610199695339725548478315",
		}).
		Reply(http.StatusOK).
		JSON(map[string]string{
			"feeRecipient": "0xb046140686d052fff581f63f8136cce132e857da",
			"makerFee":     "100000000000000",
			"takerFee":     "200000000000000",
		})

	order := types.UnsignedOrder{
		Maker:                      common.HexToAddress("0xc9b32e9563fe99612ce3a2695ac2a6404c111dde"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		MakerTokenAddress:          common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		TakerTokenAddress:          common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		ExchangeContractAddress:    common.HexToAddress("0x12459c951127e0c374ff9105dda097662a027093"),
		Salt:                       util.StrToBig("58600101225676680041453168589125977076540694791976419610199695339725548478315"),
		MakerTokenAmount:           util.StrToBig("18981000000000000"),
		TakerTokenAmount:           util.StrToBig("19000000000000000000"),
		ExpirationUnixTimestampSec: time.Unix(1518201120, 0).UTC(),
	}

	fees, err := suite.client.GetFees(context.Background(), order)
	suite.Require().NoError(err)

	expectedFees := types.Fees{
		FeeRecipient: common.HexToAddress("0xb046140686d052fff581f63f8136cce132e857da"),
		MakerFee:     util.StrToBig("100000000000000"),
		TakerFee:     util.StrToBig("200000000000000"),
	}

	suite.Equal(expectedFees, fees)
}

func (suite *ClientSuite) TestGetFeesWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Post("/fees").
		Reply(http.StatusServiceUnavailable)

	_, err := suite.client.GetFees(context.Background(), types.UnsignedOrder{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestGetFeesWithMalformedJSONResponse() {
	gock.New(suite.url).
		Post("/fees").
		Reply(http.StatusOK).
		BodyString("//\\")

	_, err := suite.client.GetFees(context.Background(), types.UnsignedOrder{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "error parsing json response")
}

func (suite *ClientSuite) TestGetFeesWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := suite.client.GetFees(ctx, types.UnsignedOrder{})
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
