package cmd

import (
	"github.com/spf13/cobra"
)

var (
	assetsCmd = &cobra.Command{
		Use: "assets",
	}
)

func init() {
	rootCmd.AddCommand(assetsCmd)
}
