package cmd

import (
	"github.com/spf13/cobra"
)

var (
	tokensCmd = &cobra.Command{
		Use: "tokens",
	}
)

func init() {
	rootCmd.AddCommand(tokensCmd)
}
