package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/linki/0x-go/relayer"
)

var (
	tokensPairsCmd = &cobra.Command{
		Use: "pairs",
		Run: listTokenPairs,
	}
)

func init() {
	tokensCmd.AddCommand(tokensPairsCmd)
}

func listTokenPairs(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	tokenPairs, err := client.GetTokenPairs(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, tp := range tokenPairs {
		fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", tp.TokenA.Address, tp.TokenB.Address)
	}
}
