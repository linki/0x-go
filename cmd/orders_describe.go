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
	ordersDescribeCmd = &cobra.Command{
		Use: "describe",
		Run: describeOrders,
	}
)

func init() {
	ordersCmd.AddCommand(ordersDescribeCmd)
}

func describeOrders(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	order, err := client.GetOrder(context.Background(), common.HexToHash(orderHash))
	if err != nil {
		log.Fatal(err)
	}

	orderFmt := `orderHash: %s
exchangeAddress: %s
senderAddress: %s
makerAddress: %s
takerAddress: %s
makerAssetData: %s
takerAssetData: %s
feeRecipientAddress: %s
makerAssetAmount: %s
takerAssetAmount: %s
makerFee: %s
takerFee: %s
expirationTimeSeconds: %d
signature: %s
salt: %s
`
	fmt.Fprintf(cmd.OutOrStdout(), orderFmt,
		order.OrderHash.Hex(),
		strings.ToLower(order.ExchangeContractAddress.Hex()),
		strings.ToLower(order.Sender.Hex()),
		strings.ToLower(order.Maker.Hex()),
		strings.ToLower(order.Taker.Hex()),
		fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(order.MakerTokenAddress.Hash().String())[2:]),
		fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(order.TakerTokenAddress.Hash().String())[2:]),
		strings.ToLower(order.FeeRecipient.Hex()),
		order.MakerTokenAmount,
		order.TakerTokenAmount,
		order.MakerFee,
		order.TakerFee,
		order.ExpirationUnixTimestampSec.Unix(),
		order.Signature,
		order.Salt,
	)
}
