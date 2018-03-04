package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ordersCmd = &cobra.Command{
		Use: "orders",
	}
)

func init() {
	rootCmd.AddCommand(ordersCmd)
}
