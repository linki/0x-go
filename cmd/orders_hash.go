package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"
)

var (
	ordersHashCmd = &cobra.Command{
		Use: "hash",
		Run: hashOrder,
	}
)

func init() {
	ordersCmd.AddCommand(ordersHashCmd)
}

func hashOrder(cmd *cobra.Command, _ []string) {
	order := types.Order{
		ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
		Maker:                      common.HexToAddress(maker),
		Taker:                      common.HexToAddress(taker),
		MakerTokenAddress:          common.HexToAddress(makerTokenAddress),
		TakerTokenAddress:          common.HexToAddress(takerTokenAddress),
		FeeRecipient:               common.HexToAddress(feeRecipient),
		MakerTokenAmount:           util.StrToBig(makerTokenAmount),
		TakerTokenAmount:           util.StrToBig(takerTokenAmount),
		MakerFee:                   util.StrToBig(makerFee),
		TakerFee:                   util.StrToBig(takerFee),
		ExpirationUnixTimestampSec: time.Unix(expirationUnixTimestampSec, 0).UTC(),
		Salt: util.StrToBig(salt),
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", order.CalculateOrderHash().Hex())
}
