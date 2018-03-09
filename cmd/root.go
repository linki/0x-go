package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	relayerURL string
)

var (
	rootCmd = &cobra.Command{}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&relayerURL, "relayer-url", "", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
