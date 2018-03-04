package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"
)

var (
	exchangeContractAddress    string
	makerTokenAddress          string
	takerTokenAddress          string
	maker                      string
	taker                      string
	feeRecipient               string
	makerTokenAmount           string
	takerTokenAmount           string
	makerFee                   string
	takerFee                   string
	expirationUnixTimestampSec int64
	salt                       string
)

var (
	ordersHashCmd = &cobra.Command{
		Use: "hash",
		Run: hashOrder,
	}
)

func init() {
	ordersHashCmd.Flags().StringVar(&exchangeContractAddress, "exchange-contract-address", "", "")
	ordersHashCmd.Flags().StringVar(&maker, "maker", "", "")
	ordersHashCmd.Flags().StringVar(&taker, "taker", "", "")
	ordersHashCmd.Flags().StringVar(&makerTokenAddress, "maker-token-address", "", "")
	ordersHashCmd.Flags().StringVar(&takerTokenAddress, "taker-token-address", "", "")
	ordersHashCmd.Flags().StringVar(&feeRecipient, "fee-recipient", "", "")
	ordersHashCmd.Flags().StringVar(&makerTokenAmount, "maker-token-amount", "", "")
	ordersHashCmd.Flags().StringVar(&takerTokenAmount, "taker-token-amount", "", "")
	ordersHashCmd.Flags().StringVar(&makerFee, "maker-fee", "", "")
	ordersHashCmd.Flags().StringVar(&takerFee, "taker-fee", "", "")
	ordersHashCmd.Flags().Int64Var(&expirationUnixTimestampSec, "expiration-unix-timestamp-sec", 0, "")
	ordersHashCmd.Flags().StringVar(&salt, "salt", "", "")

	ordersCmd.AddCommand(ordersHashCmd)
}

func hashOrder(cmd *cobra.Command, _ []string) {
	order := types.Order{
		ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
		Maker:                      common.HexToAddress(maker),
		Taker:                      common.HexToAddress(taker),
		MakerTokenAddress:          common.HexToAddress(makerTokenAddress),
		TakerTokenAddress:          common.HexToAddress(takerTokenAddress),
		FeeRecipient:               common.HexToAddress(feeRecipient),
		MakerTokenAmount:           util.StrToBig(makerTokenAmount),
		TakerTokenAmount:           util.StrToBig(takerTokenAmount),
		MakerFee:                   util.StrToBig(makerFee),
		TakerFee:                   util.StrToBig(takerFee),
		ExpirationUnixTimestampSec: time.Unix(expirationUnixTimestampSec, 0),
		Salt: util.StrToBig(salt),
	}

	fmt.Printf(`exchange-contract-address: %s
maker: %s
taker: %s
maker-token-address: %s
taker-token-address: %s
fee-recipient: %s
maker-token-amount: %s
taker-token-amount: %s
maker-fee: %s
taker-fee: %s
expiration-unix-timestamp-sec: %d
salt: %s
orderHash: %s
`, exchangeContractAddress,
		maker,
		taker,
		makerTokenAddress,
		takerTokenAddress,
		feeRecipient,
		makerTokenAmount,
		takerTokenAmount,
		makerFee,
		takerFee,
		expirationUnixTimestampSec,
		salt,
		order.CalculateOrderHash().Hex(),
	)

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", order.CalculateOrderHash().Hex())
}
