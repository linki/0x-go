package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
)

var (
	ordersSignHashCmd = &cobra.Command{
		Use: "sign-hash",
		Run: signOrderHash,
	}
)

func init() {
	ordersCmd.AddCommand(ordersSignHashCmd)
}

func signOrderHash(cmd *cobra.Command, _ []string) {
	privateKey, err := loadPrivateKey()
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	signature, err := types.SignHash(common.HexToHash(orderHash), privateKey.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to calculate the order's signature: %v", err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "v: %d\nr: %s\ns: %s\n", signature.V, signature.R.Hex(), signature.S.Hex())
}
