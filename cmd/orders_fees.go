package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/relayer"
	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"
)

var (
	ordersFeesCmd = &cobra.Command{
		Use: "fees",
		Run: getFees,
	}
)

func init() {
	ordersCmd.AddCommand(ordersFeesCmd)
}

func getFees(cmd *cobra.Command, _ []string) {
	order := types.UnsignedOrder{
		ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
		Maker:                      common.HexToAddress(maker),
		Taker:                      common.HexToAddress(taker),
		MakerTokenAddress:          common.HexToAddress(makerTokenAddress),
		TakerTokenAddress:          common.HexToAddress(takerTokenAddress),
		MakerTokenAmount:           util.StrToBig(makerTokenAmount),
		TakerTokenAmount:           util.StrToBig(takerTokenAmount),
		ExpirationUnixTimestampSec: time.Unix(expirationUnixTimestampSec, 0),
		Salt: util.StrToBig(salt),
	}

	client := relayer.NewClient(relayerURL)

	fees, err := client.GetFees(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}

	feesFmt := `feeRecipient: %s
makerFee: %s
takerFee: %s
`

	fmt.Fprintf(cmd.OutOrStdout(), feesFmt,
		strings.ToLower(fees.FeeRecipient.Hex()),
		fees.MakerFee,
		fees.TakerFee,
	)
}
