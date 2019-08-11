package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/linki/0x-go/relayer"
)

var (
	assetDataA string
	assetDataB string
)

var (
	assetsPairsCmd = &cobra.Command{
		Use: "pairs",
		Run: listAssetPairs,
	}
)

func init() {
	assetsPairsCmd.Flags().StringVar(&assetDataA, "asset-data-a", "", "")
	assetsPairsCmd.Flags().StringVar(&assetDataB, "asset-data-b", "", "")

	assetsCmd.AddCommand(assetsPairsCmd)
}

func listAssetPairs(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	assetPairs, err := client.GetAssetPairs(context.Background(), relayer.GetAssetPairsOpts{
		AssetDataA: assetDataA,
		AssetDataB: assetDataB,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, tp := range assetPairs.Records {
		fmt.Fprintf(cmd.OutOrStdout(), "%s %s %s %s\n", tp.AssetDataA.AssetData, tp.AssetDataA.MaxAmount, tp.AssetDataB.AssetData, tp.AssetDataB.MaxAmount)
	}
}
