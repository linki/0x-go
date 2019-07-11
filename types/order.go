package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Order struct {
	OrderHash                  common.Hash
	ExchangeContractAddress    common.Address
	Sender                     common.Address
	Maker                      common.Address
	Taker                      common.Address
	MakerTokenAddress          common.Address
	TakerTokenAddress          common.Address
	FeeRecipient               common.Address
	MakerTokenAmount           *big.Int
	TakerTokenAmount           *big.Int
	MakerFee                   *big.Int
	TakerFee                   *big.Int
	ExpirationUnixTimestampSec time.Time
	Signature                  string
	Salt                       *big.Int
}

func (o *Order) CalculateOrderHash() common.Hash {
	sig := crypto.Keccak256([]byte("ERC20Token(address)"))[:4]

	eip191Header := []byte{25, 1}
	domainSchemaSha := sha3.NewLegacyKeccak256()
	domainSchemaSha.Write([]byte("EIP712Domain(string name,string version,address verifyingContract)"))
	domainSha := sha3.NewLegacyKeccak256()
	nameSha := sha3.NewLegacyKeccak256()
	nameSha.Write([]byte("0x Protocol"))
	versionSha := sha3.NewLegacyKeccak256()
	versionSha.Write([]byte("2"))
	domainSha.Write(domainSchemaSha.Sum(nil))
	domainSha.Write(nameSha.Sum(nil))
	domainSha.Write(versionSha.Sum(nil))
	domainSha.Write(o.ExchangeContractAddress.Hash().Bytes())

	orderSchemaSha := sha3.NewLegacyKeccak256()
	orderSchemaSha.Write([]byte("Order(address makerAddress,address takerAddress,address feeRecipientAddress,address senderAddress,uint256 makerAssetAmount,uint256 takerAssetAmount,uint256 makerFee,uint256 takerFee,uint256 expirationTimeSeconds,uint256 salt,bytes makerAssetData,bytes takerAssetData)"))
	exchangeSha := sha3.NewLegacyKeccak256()
	exchangeSha.Write(o.ExchangeContractAddress[:])
	makerAssetDataSha := sha3.NewLegacyKeccak256()
	makerAssetDataSha.Write(sig)
	makerAssetDataSha.Write(o.MakerTokenAddress.Hash().Bytes())
	takerAssetDataSha := sha3.NewLegacyKeccak256()
	takerAssetDataSha.Write(sig)
	takerAssetDataSha.Write(o.TakerTokenAddress.Hash().Bytes())
	orderSha := sha3.NewLegacyKeccak256()
	orderSha.Write(orderSchemaSha.Sum(nil))
	orderSha.Write(o.Maker.Hash().Bytes())
	orderSha.Write(o.Taker.Hash().Bytes())
	orderSha.Write(o.FeeRecipient.Hash().Bytes())
	orderSha.Write(o.Sender.Hash().Bytes())
	orderSha.Write(common.BigToHash(o.MakerTokenAmount).Bytes())
	orderSha.Write(common.BigToHash(o.TakerTokenAmount).Bytes())
	orderSha.Write(common.BigToHash(o.MakerFee).Bytes())
	orderSha.Write(common.BigToHash(o.TakerFee).Bytes())
	orderSha.Write(common.BigToHash(big.NewInt(o.ExpirationUnixTimestampSec.Unix())).Bytes())
	orderSha.Write(common.BigToHash(o.Salt).Bytes())
	orderSha.Write(makerAssetDataSha.Sum(nil))
	orderSha.Write(takerAssetDataSha.Sum(nil))

	sha := sha3.NewLegacyKeccak256()
	sha.Write(eip191Header)
	sha.Write(domainSha.Sum(nil))
	sha.Write(orderSha.Sum(nil))

	return common.BytesToHash(sha.Sum(nil))
}

func (o *Order) UnmarshalJSON(b []byte) error {
	order := map[string]interface{}{}

	err := json.Unmarshal(b, &order)
	if err != nil {
		return err
	}

	orderData := order["order"].(map[string]interface{})

	o.ExchangeContractAddress = common.HexToAddress(orderData["exchangeAddress"].(string))
	o.Sender = common.HexToAddress(orderData["senderAddress"].(string))
	o.Maker = common.HexToAddress(orderData["makerAddress"].(string))
	o.Taker = common.HexToAddress(orderData["takerAddress"].(string))
	o.MakerTokenAddress = common.HexToAddress(orderData["makerAssetData"].(string))
	o.TakerTokenAddress = common.HexToAddress(orderData["takerAssetData"].(string))
	o.FeeRecipient = common.HexToAddress(orderData["feeRecipientAddress"].(string))

	o.MakerTokenAmount = new(big.Int)
	o.TakerTokenAmount = new(big.Int)
	o.MakerFee = new(big.Int)
	o.TakerFee = new(big.Int)
	o.Salt = new(big.Int)

	o.MakerTokenAmount.UnmarshalJSON([]byte(orderData["makerAssetAmount"].(string)))
	o.TakerTokenAmount.UnmarshalJSON([]byte(orderData["takerAssetAmount"].(string)))
	o.MakerFee.UnmarshalJSON([]byte(orderData["makerFee"].(string)))
	o.TakerFee.UnmarshalJSON([]byte(orderData["takerFee"].(string)))
	o.Signature = orderData["signature"].(string)
	o.Salt.UnmarshalJSON([]byte(orderData["salt"].(string)))

	// sig := orderData["ecSignature"].(map[string]interface{})
	// o.Signature = Signature{
	// 	V: byte(sig["v"].(float64)),
	// 	R: common.HexToHash(sig["r"].(string)),
	// 	S: common.HexToHash(sig["s"].(string)),
	// }

	timestamp, err := strconv.ParseInt(orderData["expirationTimeSeconds"].(string), 10, 64)
	if err != nil {
		return err
	}
	o.ExpirationUnixTimestampSec = time.Unix(timestamp, 0).UTC()

	// When the JSON doesn't contain an order hash, we calculate it ourself.
	metaData := order["metaData"].(map[string]interface{})
	if metaData["orderHash"] != nil {
		o.OrderHash = common.HexToHash(metaData["orderHash"].(string))
	} else {
		o.OrderHash = o.CalculateOrderHash()
	}

	return nil
}

func (o *Order) MarshalJSON() ([]byte, error) {
	order := map[string]map[string]string{
		"order": {
			"exchangeAddress":       strings.ToLower(o.ExchangeContractAddress.Hex()),
			"senderAddress":         strings.ToLower(o.Sender.Hex()),
			"makerAddress":          strings.ToLower(o.Maker.Hex()),
			"takerAddress":          strings.ToLower(o.Taker.Hex()),
			"makerAssetData":        fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(o.MakerTokenAddress.Hash().String())[2:]),
			"takerAssetData":        fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(o.TakerTokenAddress.Hash().String())[2:]),
			"feeRecipientAddress":   strings.ToLower(o.FeeRecipient.Hex()),
			"makerAssetAmount":      o.MakerTokenAmount.String(),
			"takerAssetAmount":      o.TakerTokenAmount.String(),
			"makerFee":              o.MakerFee.String(),
			"takerFee":              o.TakerFee.String(),
			"expirationTimeSeconds": fmt.Sprintf("%d", o.ExpirationUnixTimestampSec.Unix()),
			"signature":             o.Signature,
			"salt":                  o.Salt.String(),
		},
		"metaData": {
			"orderHash": o.OrderHash.Hex(),
		},
	}
	return json.Marshal(order)
}

func (o *Order) MarshalJSONPlain() ([]byte, error) {
	order := map[string]string{
		"exchangeAddress":       strings.ToLower(o.ExchangeContractAddress.Hex()),
		"senderAddress":         strings.ToLower(o.Sender.Hex()),
		"makerAddress":          strings.ToLower(o.Maker.Hex()),
		"takerAddress":          strings.ToLower(o.Taker.Hex()),
		"makerAssetData":        fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(o.MakerTokenAddress.Hash().String())[2:]),
		"takerAssetData":        fmt.Sprintf("%s%s", "0xf47261b0", strings.ToLower(o.TakerTokenAddress.Hash().String())[2:]),
		"feeRecipientAddress":   strings.ToLower(o.FeeRecipient.Hex()),
		"makerAssetAmount":      o.MakerTokenAmount.String(),
		"takerAssetAmount":      o.TakerTokenAmount.String(),
		"makerFee":              o.MakerFee.String(),
		"takerFee":              o.TakerFee.String(),
		"expirationTimeSeconds": fmt.Sprintf("%d", o.ExpirationUnixTimestampSec.Unix()),
		"signature":             o.Signature,
		"salt":                  o.Salt.String(),
	}
	return json.Marshal(order)
}
