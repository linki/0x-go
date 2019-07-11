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

type AssetsPairsSuite struct {
	suite.Suite
	console io.ReadWriter
	url     string
}

func (suite *AssetsPairsSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	rootCmd.SetOutput(suite.console)
	suite.url = "http://127.0.0.1:8080"
}

func (suite *AssetsPairsSuite) TearDownTest() {
	rootCmd.SetOutput(nil)
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *AssetsPairsSuite) TestAssetsPairs() {
	for _, tt := range []struct {
		response       map[string]interface{}
		flags          []string
		expectedParams map[string]string
		expected       string
	}{
		{
			map[string]interface{}{
				"records": []map[string]interface{}{
					{
						"assetDataA": map[string]interface{}{
							"assetData": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
							"maxAmount": "100",
						},
						"assetDataB": map[string]interface{}{
							"assetData": "0xef7fff64389b814a946f3e92105513705ca6b990",
							"maxAmount": "42",
						},
					},
				},
			},
			[]string{
				"--relayer-url", suite.url,
				"--asset-data-a", "0x323b5d4c32345ced77393b3530b1eed0f346429d",
				"--asset-data-b", "0xef7fff64389b814a946f3e92105513705ca6b990",
			},
			map[string]string{
				"assetDataA": "0x323b5d4c32345ced77393b3530b1eed0f346429d",
				"assetDataB": "0xef7fff64389b814a946f3e92105513705ca6b990",
			},
			"0x323b5d4c32345ced77393b3530b1eed0f346429d 100 0xef7fff64389b814a946f3e92105513705ca6b990 42\n",
		},
	} {
		gock.New(suite.url).
			Get("/asset_pairs").
			MatchParams(tt.expectedParams).
			Reply(http.StatusOK).
			JSON(tt.response)

		args := append(
			[]string{"assets", "pairs"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestAssetsPairsSuite(t *testing.T) {
	suite.Run(t, new(AssetsPairsSuite))
}
