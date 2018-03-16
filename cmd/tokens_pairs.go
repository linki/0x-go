package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/relayer"
)

var (
	tokenA string
	tokenB string
)

var (
	tokensPairsCmd = &cobra.Command{
		Use: "pairs",
		Run: listTokenPairs,
	}
)

func init() {
	tokensPairsCmd.Flags().StringVar(&tokenA, "token-a", "", "")
	tokensPairsCmd.Flags().StringVar(&tokenB, "token-b", "", "")

	tokensCmd.AddCommand(tokensPairsCmd)
}

func listTokenPairs(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	tokenPairs, err := client.GetTokenPairs(context.Background(), relayer.GetTokenPairsOpts{
		TokenA: common.HexToAddress(tokenA),
		TokenB: common.HexToAddress(tokenB),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, tp := range tokenPairs {
		fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", tp.TokenA.Address, tp.TokenB.Address)
	}
}
