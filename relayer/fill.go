package relayer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
)

func (c *Client) FillOrder(_ context.Context, auth *bind.TransactOpts, order types.Order, amount *big.Int) error {
	// call the `FillOrder` function on the exchange contract.
	_, err := c.exchangeContract.FillOrder(auth,
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
		amount,
		true,
		order.Signature.V,
		order.Signature.R,
		order.Signature.S,
	)
	if err != nil {
		return err
	}

	return nil
}
