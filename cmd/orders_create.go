package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/linki/0x-go/relayer"
	"github.com/linki/0x-go/types"
	"github.com/linki/0x-go/util"
)

var (
	ordersCreateCmd = &cobra.Command{
		Use: "create",
		Run: createOrder,
	}
)

func init() {
	ordersCmd.AddCommand(ordersCreateCmd)
}

func createOrder(cmd *cobra.Command, _ []string) {
	client := relayer.NewClient(relayerURL)

	fees := types.Fees{}

	if autodetectFees {
		quoteOrder := types.UnsignedOrder{
			ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
			Maker:                      common.HexToAddress(maker),
			Taker:                      common.HexToAddress(taker),
			MakerTokenAddress:          common.HexToAddress(makerTokenAddress),
			TakerTokenAddress:          common.HexToAddress(takerTokenAddress),
			MakerTokenAmount:           util.StrToBig(makerTokenAmount),
			TakerTokenAmount:           util.StrToBig(takerTokenAmount),
			ExpirationUnixTimestampSec: time.Unix(expirationUnixTimestampSec, 0),
			Salt: util.StrToBig(salt),
		}

		detectedFees, err := client.GetFees(context.Background(), quoteOrder)
		if err != nil {
			log.Fatal(err)
		}
		fees = detectedFees
	} else {
		fees = types.Fees{
			FeeRecipient: common.HexToAddress(feeRecipient),
			MakerFee:     util.StrToBig(makerFee),
			TakerFee:     util.StrToBig(takerFee),
		}
	}

	order := types.Order{
		ExchangeContractAddress: common.HexToAddress(exchangeContractAddress),
		Maker:                      common.HexToAddress(maker),
		Taker:                      common.HexToAddress(taker),
		MakerTokenAddress:          common.HexToAddress(makerTokenAddress),
		TakerTokenAddress:          common.HexToAddress(takerTokenAddress),
		MakerTokenAmount:           util.StrToBig(makerTokenAmount),
		TakerTokenAmount:           util.StrToBig(takerTokenAmount),
		ExpirationUnixTimestampSec: time.Unix(expirationUnixTimestampSec, 0),
		Salt: util.StrToBig(salt),
	}
	order.FeeRecipient = fees.FeeRecipient
	order.MakerFee = fees.MakerFee
	order.TakerFee = fees.TakerFee

	order.OrderHash = order.CalculateOrderHash()

	privateKey, err := loadPrivateKey()
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	signature, err := types.SignHash(order.OrderHash, privateKey.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to calculate the order's signature: %v", err)
	}
	order.Signature = signature

	if err := client.CreateOrder(context.Background(), order); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", order.OrderHash.Hex())
}
