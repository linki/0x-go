package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/relayer"
)

var (
	orderHash string
)

var (
	ordersDescribeCmd = &cobra.Command{
		Use: "describe",
		Run: describeOrders,
	}
)

func init() {
	ordersDescribeCmd.Flags().StringVar(&orderHash, "order-hash", "", "")

	ordersCmd.AddCommand(ordersDescribeCmd)
}

func describeOrders(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	order, err := client.GetOrder(context.Background(), common.HexToHash(orderHash))
	if err != nil {
		log.Fatal(err)
	}

	orderFmt := `orderHash: %s
exchange-contract-address: %s
maker: %s
taker: %s
maker-token-address: %s
taker-token-address: %s
fee-recipient: %s
maker-token-amount: %s
taker-token-amount: %s
maker-fee: %s
taker-fee: %s
expiration-unix-timestamp-sec: %d
salt: %s
v: %d
r: %s
s: %s
`

	fmt.Fprintf(cmd.OutOrStdout(), orderFmt,
		order.OrderHash.Hex(),
		strings.ToLower(order.ExchangeContractAddress.Hex()),
		strings.ToLower(order.Maker.Hex()),
		strings.ToLower(order.Taker.Hex()),
		strings.ToLower(order.MakerTokenAddress.Hex()),
		strings.ToLower(order.TakerTokenAddress.Hex()),
		strings.ToLower(order.FeeRecipient.Hex()),
		order.MakerTokenAmount,
		order.TakerTokenAmount,
		order.MakerFee,
		order.TakerFee,
		order.ExpirationUnixTimestampSec.Unix(),
		order.Salt,
		order.Signature.V,
		order.Signature.R.Hex(),
		order.Signature.S.Hex(),
	)
}
