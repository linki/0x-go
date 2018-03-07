package relayer

import (
	"context"
	"net/http"
	"testing"

	"github.com/linki/0x-go/types"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client      *Client
	url         string
	tokenA      types.Token
	tokenB      types.Token
	tokenPairAB types.TokenPair
}

func (suite *ClientSuite) SetupTest() {
	suite.url = "http://127.0.0.1:8080"
	suite.client = NewClient(suite.url)
	suite.tokenA = types.Token{
		Address:   "0x323b5d4c32345ced77393b3530b1eed0f346429d",
		MinAmount: "0",
		MaxAmount: "10000000000000000000",
		Precision: 5,
	}
	suite.tokenB = types.Token{
		Address:   "0xef7fff64389b814a946f3e92105513705ca6b990",
		MinAmount: "0",
		MaxAmount: "50000000000000000000",
		Precision: 5,
	}
	suite.tokenPairAB = types.TokenPair{
		TokenA: suite.tokenA,
		TokenB: suite.tokenB,
	}
}

func (suite *ClientSuite) TearDownTest() {
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *ClientSuite) TestGetTokenPairs() {
	gock.New(suite.url).
		Get("/token_pairs").
		Reply(http.StatusOK).
		JSON([]map[string]interface{}{
			{
				"tokenA": map[string]interface{}{
					"address":   "0x323b5d4c32345ced77393b3530b1eed0f346429d",
					"minAmount": "0",
					"maxAmount": "10000000000000000000",
					"precision": 5,
				},
				"tokenB": map[string]interface{}{
					"address":   "0xef7fff64389b814a946f3e92105513705ca6b990",
					"minAmount": "0",
					"maxAmount": "50000000000000000000",
					"precision": 5,
				},
			},
		})

	tokenPairs, err := suite.client.GetTokenPairs(context.Background())
	suite.NoError(err)

	suite.Require().Len(tokenPairs, 1)
	suite.Equal(suite.tokenPairAB, tokenPairs[0])
}

func (suite *ClientSuite) TestGetTokenPairsWithUnsuccessfulResponse() {
	gock.New(suite.url).
		Get("/token_pairs").
		Reply(http.StatusServiceUnavailable)

	_, err := suite.client.GetTokenPairs(context.Background())
	suite.Require().Error(err)
	suite.Contains(err.Error(), "erroneous status code")
	suite.Contains(err.Error(), "503 Service Unavailable")
}

func (suite *ClientSuite) TestGetTokenPairsWithMalformedJSON() {
	gock.New(suite.url).
		Get("/token_pairs").
		Reply(http.StatusOK).
		BodyString("//\\")

	_, err := suite.client.GetTokenPairs(context.Background())
	suite.Require().Error(err)
	suite.Contains(err.Error(), "error parsing json response")
}

func (suite *ClientSuite) TestGetTokenPairsWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := suite.client.GetTokenPairs(ctx)
	suite.Require().Error(err)
	suite.Contains(err.Error(), "context canceled")
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
