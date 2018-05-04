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
			`v: 28
r: 0x38c4e69b77f5e85e577337b2738fb27ddffa640a9b90971c260a4618c41bc4b2
s: 0x12a308633603f6575bd73495d4cffd37d61ddc788b96ebe4a03b8c499a559d32
`,
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
