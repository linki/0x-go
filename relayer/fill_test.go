package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/linki/0x-go/contracts/protocol"
	"github.com/linki/0x-go/contracts/tokens"
	"github.com/linki/0x-go/util"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type ClientFillSuite struct {
	suite.Suite
	client *Client
	url    string

	backend    *backends.SimulatedBackend
	masterAuth *bind.TransactOpts
	makerAuth  *bind.TransactOpts
	takerAuth  *bind.TransactOpts
	makerKey   *ecdsa.PrivateKey

	zrxAddress       common.Address
	zrxToken         *tokens.ZRXToken
	wethAddress      common.Address
	wethToken        *tokens.WETH9
	exchangeAddress  common.Address
	exchangeContract *protocol.Exchange
}

func (suite *ClientFillSuite) SetupTest() {
	var err error

	// setup a master, maker and taker account.
	suite.masterAuth, _ = suite.newAccountFromHex("f5f7b32d9fa2d52b22ec3cd8aed0298ac6ebccbd54b340f870c734161fc7efbc")
	suite.makerAuth, suite.makerKey = suite.newAccountFromHex("c5010f1aaa410683399ea6bc85044360545917e7d61279681fd506cedc81bc23")
	suite.takerAuth, _ = suite.newAccountFromHex("acfb7e627bead064eac670aa6d9ee28ea3b876980808c79d3aac6881447b66ae")

	// initialize the system, give everyone 200 ETH.
	suite.backend = backends.NewSimulatedBackend(
		map[common.Address]core.GenesisAccount{
			suite.masterAuth.From: {Balance: util.EthToWei(200)},
			suite.makerAuth.From:  {Balance: util.EthToWei(200)},
			suite.takerAuth.From:  {Balance: util.EthToWei(200)},
		}, 10000000,
	)

	// deploy the ZRX token contract.
	suite.zrxAddress, _, suite.zrxToken, err = tokens.DeployZRXToken(suite.masterAuth, suite.backend)
	suite.Require().NoError(err)

	// deploy the WETH token contract.
	suite.wethAddress, _, suite.wethToken, err = tokens.DeployWETH9(suite.masterAuth, suite.backend)
	suite.Require().NoError(err)

	// deploy 0x's token transfer proxy contract.
	proxyAddress, _, proxyContract, err := protocol.DeployTokenTransferProxy(suite.masterAuth, suite.backend)
	suite.Require().NoError(err)

	// deploy 0x's exchange contract.
	suite.exchangeAddress, _, suite.exchangeContract, err = protocol.DeployExchange(suite.masterAuth, suite.backend, suite.zrxAddress, proxyAddress)
	suite.Require().NoError(err)

	// authorize the exchange contract to call the token transfer proxy contract.
	_, err = proxyContract.AddAuthorizedAddress(suite.masterAuth, suite.exchangeAddress)
	suite.Require().NoError(err)

	// give the master account 100 WETH in order to send them to maker and taker.
	suite.depositWeth(suite.masterAuth, suite.wethToken, util.EthToWei(100))

	// send 50 WETH to both maker and taker.
	suite.transferToken(suite.masterAuth, suite.wethToken, suite.makerAuth.From, util.EthToWei(50))
	suite.transferToken(suite.masterAuth, suite.wethToken, suite.takerAuth.From, util.EthToWei(50))

	// send 50 ZRX to both maker and taker.
	suite.transferToken(suite.masterAuth, suite.zrxToken, suite.makerAuth.From, util.EthToWei(50))
	suite.transferToken(suite.masterAuth, suite.zrxToken, suite.takerAuth.From, util.EthToWei(50))

	// maker sets allowance to 50 ZRX for token transfer proxy contract.
	_, err = suite.zrxToken.Approve(suite.makerAuth, proxyAddress, util.EthToWei(50))
	suite.Require().NoError(err)

	// taker sets allowance to 50 WETH for token transfer proxy contract.
	_, err = suite.wethToken.Approve(suite.takerAuth, proxyAddress, util.EthToWei(50))
	suite.Require().NoError(err)

	// commit all of the above.
	suite.backend.Commit()

	// ensure that maker has 50 ZRX and 50 WETH.
	suite.Require().Equal(util.EthToWei(50), suite.getTokenBalance(suite.makerAuth, suite.wethToken))
	suite.Require().Equal(util.EthToWei(50), suite.getTokenBalance(suite.makerAuth, suite.zrxToken))

	// ensure that taker has 50 ZRX and 50 WETH.
	suite.Require().Equal(util.EthToWei(50), suite.getTokenBalance(suite.takerAuth, suite.wethToken))
	suite.Require().Equal(util.EthToWei(50), suite.getTokenBalance(suite.takerAuth, suite.zrxToken))

	suite.url = "http://127.0.0.1:8080"
	suite.client = NewClient(suite.url)
	suite.client.exchangeContract = suite.exchangeContract
}

func (suite *ClientFillSuite) TearDownTest() {
	suite.True(gock.IsDone())
	gock.Off()
}

func (suite *ClientFillSuite) TestFillOrder() {
	gock.New(suite.url).
		Get("/order/0x9e4ecffa4c7da98177505139b5411d87bc62637a4a9c64a8c87d924b45669f09").
		Reply(http.StatusOK).
		JSON(map[string]interface{}{
			"orderHash":                  "0x9e4ecffa4c7da98177505139b5411d87bc62637a4a9c64a8c87d924b45669f09",
			"exchangeContractAddress":    suite.exchangeAddress,
			"maker":                      "0xeF5cd3E1F525f899f1d383920209cC890fF7CF95",
			"taker":                      "0x0000000000000000000000000000000000000000",
			"makerTokenAddress":          suite.zrxAddress,
			"takerTokenAddress":          suite.wethAddress,
			"feeRecipient":               "0xD714dDDCaf8fCB719E3310E39c0A1824daA2BED0",
			"makerTokenAmount":           "20000000000000000000",
			"takerTokenAmount":           "10000000000000000000",
			"makerFee":                   "0",
			"takerFee":                   "0",
			"expirationUnixTimestampSec": "1518112800",
			"salt": "42",
			"ecSignature": map[string]interface{}{
				"v": 27,
				"r": "0x57dd2e211aefdee2ffe8dbae676fe3804b61bc2418678226262963fe8e61f954",
				"s": "0x7165cfa75eff1eaf8427e68c1d0be91142f03aacad736ad093baa061396bc3ef",
			},
		})

	order, err := suite.client.GetOrder(context.Background(), common.HexToHash("0x9e4ecffa4c7da98177505139b5411d87bc62637a4a9c64a8c87d924b45669f09"))
	suite.NoError(err)

	err = suite.client.FillOrder(context.Background(), suite.takerAuth, order, util.EthToWei(1)) // taker fills 1 WETH for 2 ZRX
	suite.NoError(err)

	// commit all of the above.
	suite.backend.Commit()

	// ensure that maker has 1 WETH more and 2 ZRX less.
	suite.Equal(util.EthToWei(51), suite.getTokenBalance(suite.makerAuth, suite.wethToken))
	suite.Equal(util.EthToWei(48), suite.getTokenBalance(suite.makerAuth, suite.zrxToken))

	// ensure that taker has 1 WETH less and 2 ZRX more.
	suite.Equal(util.EthToWei(49), suite.getTokenBalance(suite.takerAuth, suite.wethToken))
	suite.Equal(util.EthToWei(52), suite.getTokenBalance(suite.takerAuth, suite.zrxToken))
}

// Token is a common interface for all ERC20 tokens.
type Token interface {
	BalanceOf(opts *bind.CallOpts, address common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*ethtypes.Transaction, error)
}

// depositWeth converts ETH to WETH.
func (suite *ClientFillSuite) depositWeth(auth *bind.TransactOpts, token *tokens.WETH9, value *big.Int) {
	_, err := token.Deposit(&bind.TransactOpts{
		From:     auth.From,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    value,
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
	})
	suite.Require().NoError(err)
}

// transferToken transfers an amount of the given tokens to the recipient.
func (suite *ClientFillSuite) transferToken(auth *bind.TransactOpts, token Token, recipient common.Address, value *big.Int) {
	_, err := token.Transfer(auth, recipient, value)
	suite.Require().NoError(err)
}

// getTokenBalance returns the given token's balance of the sender.
func (suite *ClientFillSuite) getTokenBalance(auth *bind.TransactOpts, token Token) *big.Int {
	balance, err := token.BalanceOf(nil, auth.From)
	suite.Require().NoError(err)
	return balance
}

// newAccount creates a new random ethereum account.
func (suite *ClientFillSuite) newAccount() (*bind.TransactOpts, *ecdsa.PrivateKey) {
	key, err := crypto.GenerateKey()
	suite.Require().NoError(err)
	return bind.NewKeyedTransactor(key), key
}

func (suite *ClientFillSuite) newAccountFromHex(hex string) (*bind.TransactOpts, *ecdsa.PrivateKey) {
	key, err := crypto.HexToECDSA(hex)
	suite.Require().NoError(err)
	return bind.NewKeyedTransactor(key), key
}

func TestClientFillSuite(t *testing.T) {
	suite.Run(t, new(ClientFillSuite))
}
