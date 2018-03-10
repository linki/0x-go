package cmd

import (
	"github.com/spf13/cobra"
)

var (
	exchangeContractAddress string
	makerTokenAddress       string
	takerTokenAddress       string
	maker                   string
	taker                   string
	feeRecipient            string
)

var (
	ordersCmd = &cobra.Command{
		Use: "orders",
	}
)

func init() {
	ordersCmd.PersistentFlags().StringVar(&exchangeContractAddress, "exchange-contract-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&maker, "maker", "", "")
	ordersCmd.PersistentFlags().StringVar(&taker, "taker", "", "")
	ordersCmd.PersistentFlags().StringVar(&makerTokenAddress, "maker-token-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&takerTokenAddress, "taker-token-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&feeRecipient, "fee-recipient", "", "")

	rootCmd.AddCommand(ordersCmd)
}
