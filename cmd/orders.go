package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var (
	orderHash                  string
	exchangeContractAddress    string
	makerTokenAddress          string
	takerTokenAddress          string
	sender                     string
	maker                      string
	taker                      string
	feeRecipient               string
	makerTokenAmount           string
	takerTokenAmount           string
	makerFee                   string
	takerFee                   string
	expirationUnixTimestampSec int64
	salt                       string

	keystoreFile string
	passphrase   string
)

var (
	ordersCmd = &cobra.Command{
		Use: "orders",
	}
)

func init() {
	ordersCmd.PersistentFlags().StringVar(&orderHash, "order-hash", "", "")
	ordersCmd.PersistentFlags().StringVar(&exchangeContractAddress, "exchange-contract-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&sender, "sender", "", "")
	ordersCmd.PersistentFlags().StringVar(&maker, "maker", "", "")
	ordersCmd.PersistentFlags().StringVar(&taker, "taker", "", "")
	ordersCmd.PersistentFlags().StringVar(&makerTokenAddress, "maker-token-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&takerTokenAddress, "taker-token-address", "", "")
	ordersCmd.PersistentFlags().StringVar(&feeRecipient, "fee-recipient", "", "")
	ordersCmd.PersistentFlags().StringVar(&makerTokenAmount, "maker-token-amount", "", "")
	ordersCmd.PersistentFlags().StringVar(&takerTokenAmount, "taker-token-amount", "", "")
	ordersCmd.PersistentFlags().StringVar(&makerFee, "maker-fee", "", "")
	ordersCmd.PersistentFlags().StringVar(&takerFee, "taker-fee", "", "")
	ordersCmd.PersistentFlags().Int64Var(&expirationUnixTimestampSec, "expiration-unix-timestamp-sec", 0, "")
	ordersCmd.PersistentFlags().StringVar(&salt, "salt", "", "")

	ordersCmd.PersistentFlags().StringVar(&keystoreFile, "keystore-file", "", "")
	ordersCmd.PersistentFlags().StringVar(&passphrase, "passphrase", "", "")

	rootCmd.AddCommand(ordersCmd)
}

func loadPrivateKey() (*keystore.Key, error) {
	// Open the provided Keystore file.
	keyStoreFile, err := os.OpenFile(keystoreFile, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("Failed to open Keystore file: %v", err)
	}
	defer keyStoreFile.Close()

	// Read the Keystore file's JSON content.
	keyStoreJSON, err := ioutil.ReadAll(keyStoreFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Keystore file: %v", err)
	}

	// Decrypt the encrypted private key using the provided passphrase.
	privateKey, err := keystore.DecryptKey(keyStoreJSON, passphrase)
	if err != nil {
		return nil, fmt.Errorf("Failed to decrypt private key: %v", err)
	}

	return privateKey, nil
}
