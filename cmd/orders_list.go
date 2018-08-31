package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/linki/0x-go/registry"
	"github.com/linki/0x-go/relayer"
	"github.com/linki/0x-go/types"
)

var (
	tokenAddress string
	trader       string
)

var (
	ordersListCmd = &cobra.Command{
		Use: "list",
		Run: listOrders,
	}
)

func init() {
	ordersListCmd.Flags().StringVar(&tokenAddress, "token-address", "", "")
	ordersListCmd.Flags().StringVar(&trader, "trader", "", "")

	ordersCmd.AddCommand(ordersListCmd)
}

func listOrders(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	orders, err := client.GetOrders(context.Background(), relayer.GetOrdersOpts{
		ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
		TokenAddress:            common.HexToAddress(tokenAddress),
		MakerTokenAddress:       common.HexToAddress(makerTokenAddress),
		TakerTokenAddress:       common.HexToAddress(takerTokenAddress),
		Maker:                   common.HexToAddress(maker),
		Taker:                   common.HexToAddress(taker),
		Trader:                  common.HexToAddress(trader),
		FeeRecipient:            common.HexToAddress(feeRecipient),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range orders {
		makerToken := registry.Lookup(o.MakerTokenAddress.Hex())
		takerToken := registry.Lookup(o.TakerTokenAddress.Hex())

		fmt.Fprintf(cmd.OutOrStdout(), "%s %.6f %s %.6f %s %.6f %s/%s\n",
			o.OrderHash.Hex(),
			takerToken.NormalizedValue(o.TakerTokenAmount),
			takerToken.Symbol,
			makerToken.NormalizedValue(o.MakerTokenAmount),
			makerToken.Symbol,
			types.Price(takerToken.NormalizedValue(o.TakerTokenAmount), makerToken.NormalizedValue(o.MakerTokenAmount)),
			takerToken.Symbol,
			makerToken.Symbol,
		)
	}
}
