package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrdersSignHashSuite struct {
	suite.Suite
	console      io.ReadWriter
	keystoreFile string
}

func (suite *OrdersSignHashSuite) SetupTest() {
	suite.console = &bytes.Buffer{}
	rootCmd.SetOutput(suite.console)
	suite.keystoreFile = setupTestKeystoreFile(suite.Require())
}

func (suite *OrdersSignHashSuite) TearDownTest() {
	os.Remove(suite.keystoreFile)
	rootCmd.SetOutput(nil)
}

func (suite *OrdersSignHashSuite) TestOrdersSignHash() {
	for _, tt := range []struct {
		flags    []string
		expected string
	}{
		{
			[]string{
				"--order-hash", "0x731200d7056c1a4900cb9015d208d7dd1b2424afe21baa4e74084824112c253e",
				"--keystore-file", suite.keystoreFile,
				"--passphrase", "not-secure-do-not-use-me-for-anything-else",
			},
			"0x1b6efe8313a7aa5a06ee95e41029b7e53ba667abb8ebce84de68ae87c972a2d11c74fddf58bf162b39703916db9f5afbb9a70c8b6f21c8b5018be884b815997ed002\n",
		},
	} {
		args := append(
			[]string{"orders", "sign-hash"},
			tt.flags...,
		)
		rootCmd.SetArgs(args)

		suite.Require().NoError(rootCmd.Execute())

		output, err := ioutil.ReadAll(suite.console)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, string(output))
	}
}

func TestOrdersSignHashSuite(t *testing.T) {
	suite.Run(t, new(OrdersSignHashSuite))
}
