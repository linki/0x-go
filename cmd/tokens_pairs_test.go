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

type TokensPairsSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *TokensPairsSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	tokensPairsCmd.SetOutput(suite.console)

	suite.url = "http://127.0.0.1:8080"
}

func (suite *TokensPairsSuite) TearDownTest() {
	tokensPairsCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *TokensPairsSuite) TestTokensPairs() {
	for _, tt := range []struct {
		response       []map[string]interface{}
		flags          []string
		expectedParams map[string]string
		expected       string
	}{
		{
			[]map[string]interface{}{
				{
					"tokenA": map[string]interface{}{
						"address": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
					},
					"tokenB": map[string]interface{}{
						"address": "0xef7fff64389b814a946f3e92105513705ca6b990",
					},
				},
			},
			[]string{
				"--relayer-url", suite.url,
				"--token-a", "0x323b5d4c32345ced77393b3530b1eed0f346429d",
				"--token-b", "0xef7fff64389b814a946f3e92105513705ca6b990",
			},
			map[string]string{
				"tokenA": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
				"tokenB": "0xef7fff64389b814a946f3e92105513705ca6b990",
			},
			"0x323b5d4c32345ced77393b3530b1eed0f346429d 0xef7fff64389b814a946f3e92105513705ca6b990\n",
		},
	} {
		gock.New(suite.url).
			Get("/token_pairs").
			MatchParams(tt.expectedParams).
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"tokens", "pairs"},
			tt.flags...,
		)

		rootCmd.SetArgs(args)
		rootCmd.Execute()

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestTokensPairsSuite(t *testing.T) {
	suite.Run(t, new(TokensPairsSuite))
}
