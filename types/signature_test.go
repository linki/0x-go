package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/suite"
)

type SignatureSuite struct {
	suite.Suite
}

func (suite *SignatureSuite) TestSignHash() {
	for _, tt := range []struct {
		orderHash  string
		privateKey string
		expected   Signature
	}{
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"c325e0261d889e2cfd581be2eef17405e4a872ef7a69ada4f09e7375a08c556b",
			Signature{
				V: 28,
				R: common.HexToHash("0x38ae2328af4da36d8f63b1b034e6de91f8488c4c0226754b0107b9a88532b148"),
				S: common.HexToHash("0x7932e9f29fecd5d483bae7b07df57a04b7ca61ce6dae029274f0cc0358cc6b7c"),
			},
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			Signature{
				V: 28,
				R: common.HexToHash("0xa7be996ed9d2d754ffe235074db015647227b623cd505657924d996dbb6ba1d2"),
				S: common.HexToHash("0x2e817c35cdb3e20bd0ffa29e8a0b6b0ede18b8b9c83b2275ce86a42841844563"),
			},
		},
		{
			"0xbbde8ab3d1e178726af9f8802380e4de79e483fff8daf29ee9b86d87e3aebf12",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			Signature{
				V: 27,
				R: common.HexToHash("0x917b943b6a2751eec7408376a07eca9b530bce2564e360377c73402bb16d797d"),
				S: common.HexToHash("0x24fbd2428bfde0cda5614abee424525a76c52e9e46da55441e283ac09e1d535f"),
			},
		},
	} {
		orderHash := common.HexToHash(tt.orderHash)

		privateKey, err := crypto.HexToECDSA(tt.privateKey)
		suite.Require().NoError(err)

		signature, err := SignHash(orderHash, privateKey)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, signature)
	}
}

func TestSignatureSuite(t *testing.T) {
	suite.Run(t, new(SignatureSuite))
}
