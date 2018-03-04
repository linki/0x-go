package types

import (
	"math/big"
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
	order := Order{
		ExchangeContractAddress: common.HexToAddress("0xCcEF5Ce0a59D4CA4a942302002Eb0bf64Ab00D10"),
		Maker:                      common.HexToAddress("0xadD132AD2b945CDBDC4Dd51C32744bfA4D5fC28d"),
		Taker:                      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		MakerTokenAddress:          common.HexToAddress("0x6ba494043E22755c8D626739b2Ff7305D76ef7ec"),
		TakerTokenAddress:          common.HexToAddress("0x672E9578A90d79f1d817e0197937635c00c76352"),
		FeeRecipient:               common.HexToAddress("0xBdF0bF6Bba29B55199Ba7f0895a797a2037080DC"),
		MakerTokenAmount:           util.StrToBig("20000000000000000000"),
		TakerTokenAmount:           util.StrToBig("10000000000000000000"),
		MakerFee:                   common.Big0,
		TakerFee:                   common.Big0,
		ExpirationUnixTimestampSec: time.Date(2018, 2, 8, 18, 0, 0, 0, time.UTC),
		Salt: big.NewInt(42),
	}
	orderHash := common.HexToHash("0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7")

	suite.Equal(orderHash, order.CalculateOrderHash())
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
