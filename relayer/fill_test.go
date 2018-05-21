package relayer

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/linki/0x-go/contracts/protocol"
	"github.com/linki/0x-go/contracts/tokens"
	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/suite"
)

type ClientFillSuite struct {
	suite.Suite

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
	suite.masterAuth, _ = suite.newAccount()
	suite.makerAuth, suite.makerKey = suite.newAccount()
	suite.takerAuth, _ = suite.newAccount()

	// initialize the system, give everyone 200 ETH.
	suite.backend = backends.NewSimulatedBackend(
		map[common.Address]core.GenesisAccount{
			suite.masterAuth.From: {Balance: util.EthToWei(200)},
			suite.makerAuth.From:  {Balance: util.EthToWei(200)},
			suite.takerAuth.From:  {Balance: util.EthToWei(200)},
		},
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
}

func (suite *ClientFillSuite) TestFillOrder() {
	// create an order where the maker sells ZRX for WETH.
	order := types.Order{
		Maker:                   suite.makerAuth.From,
		Taker:                   common.HexToAddress("0x0000000000000000000000000000000000000000"),
		FeeRecipient:            suite.masterAuth.From,
		MakerTokenAddress:       suite.zrxAddress,
		TakerTokenAddress:       suite.wethAddress,
		ExchangeContractAddress: suite.exchangeAddress,
		Salt:                       big.NewInt(42),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		MakerTokenAmount:           util.EthToWei(20), // 20 ZRX
		TakerTokenAmount:           util.EthToWei(10), // 10 WETH
		ExpirationUnixTimestampSec: time.Date(2018, 2, 8, 18, 0, 0, 0, time.UTC),
	}

	// calculate the order hash.
	order.OrderHash = order.CalculateOrderHash()

	// sign the order hash.
	signature, err := types.SignHash(order.OrderHash, suite.makerKey)
	suite.Require().NoError(err)

	// attach the signature to the order.
	order.Signature = signature

	// call the `FillOrder` function on the exchange contract.
	_, err = suite.exchangeContract.FillOrder(suite.takerAuth,
		[5]common.Address{
			order.Maker,
			order.Taker,
			order.MakerTokenAddress,
			order.TakerTokenAddress,
			order.FeeRecipient,
		},
		[6]*big.Int{
			order.MakerTokenAmount,
			order.TakerTokenAmount,
			order.MakerFee,
			order.TakerFee,
			big.NewInt(order.ExpirationUnixTimestampSec.Unix()),
			order.Salt,
		},
		util.EthToWei(1), // taker fills 1 WETH for 2 ZRX
		true,
		order.Signature.V,
		order.Signature.R,
		order.Signature.S,
	)
	suite.Require().NoError(err)

	// commit all of the above.
	suite.backend.Commit()

	// ensure that maker has 1 WETH more and 2 ZRX less.
	suite.Equal(util.EthToWei(51), suite.getTokenBalance(suite.makerAuth, suite.wethToken))
	suite.Equal(util.EthToWei(48), suite.getTokenBalance(suite.makerAuth, suite.zrxToken))

	// ensure that taker has 1 WETH less and 2 ZRX more.
	suite.Equal(util.EthToWei(49), suite.getTokenBalance(suite.takerAuth, suite.wethToken))
	suite.Equal(util.EthToWei(52), suite.getTokenBalance(suite.takerAuth, suite.zrxToken))
}

func (suite *ClientFillSuite) TearDownTest() {
	suite.True(gock.IsDone())
	gock.Off()
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

func TestClientFillSuite(t *testing.T) {
	suite.Run(t, new(ClientFillSuite))
}
