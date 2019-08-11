package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		expected   string
	}{
		// EIP712
		{
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb",
			"5f57c21cece9fad96863b031a1e68192d03671b4158c3337bb44a77b7411b509",
			"0x1c344ec69afd5ebba01527cb6726527cbd25f97d0a3efc9b05b41386f9cf05227e656c960ed46bd63a3ba320ce0494f4ec76968d49b53f51ea12883c8d87a9fb3902",
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"c325e0261d889e2cfd581be2eef17405e4a872ef7a69ada4f09e7375a08c556b",
			"0x1b776d88743e0fa952352fbc08fb91d5509972e063e81bc9b85e442becf29f09bd44631c2fc86995e88aef373888945775bd47ded46abbe85bf9bdae20789ff66a02",
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			"0x1bac170f678d3694e5a2d3943df793310d29f94ea8aa4ef5f153e2377ea056542b700f9f097a41685307f158e024ad91d817eb8db89626af69f09424c4be2f6b5802",
		},
		{
			"0xbbde8ab3d1e178726af9f8802380e4de79e483fff8daf29ee9b86d87e3aebf12",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			"0x1c087cc811a91872c76c24ddb9b170653befd2643510cdd04e5226cfe2f68b7dec5697533c41447634d3f30b40f57ee7c5fd6e9822a9fd53de49b91a1f3fd897c702",
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

func (suite *SignatureSuite) TestSignHashEthSign() {
	for _, tt := range []struct {
		orderHash  string
		privateKey string
		expected   string
	}{
		// EthSign
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"c325e0261d889e2cfd581be2eef17405e4a872ef7a69ada4f09e7375a08c556b",
			"0x1c38ae2328af4da36d8f63b1b034e6de91f8488c4c0226754b0107b9a88532b1487932e9f29fecd5d483bae7b07df57a04b7ca61ce6dae029274f0cc0358cc6b7c03",
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			"0x1ca7be996ed9d2d754ffe235074db015647227b623cd505657924d996dbb6ba1d22e817c35cdb3e20bd0ffa29e8a0b6b0ede18b8b9c83b2275ce86a4284184456303",
		},
		{
			"0xbbde8ab3d1e178726af9f8802380e4de79e483fff8daf29ee9b86d87e3aebf12",
			"9df70a8bde54b8fd0121b6f28e885e85457ad81cffeea2c99ebd8c2fae863c87",
			"0x1b917b943b6a2751eec7408376a07eca9b530bce2564e360377c73402bb16d797d24fbd2428bfde0cda5614abee424525a76c52e9e46da55441e283ac09e1d535f03",
		},
	} {
		orderHash := common.HexToHash(tt.orderHash)

		privateKey, err := crypto.HexToECDSA(tt.privateKey)
		suite.Require().NoError(err)

		signature, err := SignHashEthSign(orderHash, privateKey)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, signature)
	}
}

func (suite *SignatureSuite) TestSignHashEIP712() {
	for _, tt := range []struct {
		orderHash  string
		privateKey string
		expected   string
	}{
		// EIP712
		{
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb",
			"5f57c21cece9fad96863b031a1e68192d03671b4158c3337bb44a77b7411b509",
			"0x1c344ec69afd5ebba01527cb6726527cbd25f97d0a3efc9b05b41386f9cf05227e656c960ed46bd63a3ba320ce0494f4ec76968d49b53f51ea12883c8d87a9fb3902",
		},
	} {
		orderHash := common.HexToHash(tt.orderHash)

		privateKey, err := crypto.HexToECDSA(tt.privateKey)
		suite.Require().NoError(err)

		signature, err := SignHashEIP712(orderHash, privateKey)
		suite.Require().NoError(err)

		suite.Equal(tt.expected, signature)
	}
}

func (suite *SignatureSuite) TestVerifySignature() {
	for _, tt := range []struct {
		orderHash string
		address   string
		expected  string
	}{
		// EIP712
		{
			"0x39a90feeca9c0568526df9e3ceecdb7113bfe9f7b981ae19c4570d702e8302bb",
			"0x001eeaf1ec3c4aceaed088d1e7e151dd6dd0098c",
			"0x1c6b2caaf983908bb83e2e0db4e6a782405c8e827129ede38855ddd68420c3f3530cecc18e48ca4f0478e456eb4213ae21a54c5b01fcfb24cdefadd16de72b9e4602",
		},
		// EthSign
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"0xadD132AD2b945CDBDC4Dd51C32744bfA4D5fC28d",
			"0x1c38ae2328af4da36d8f63b1b034e6de91f8488c4c0226754b0107b9a88532b1487932e9f29fecd5d483bae7b07df57a04b7ca61ce6dae029274f0cc0358cc6b7c03",
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"0x0CFa1FaF595594a7609c824530d5623c84d6df50",
			"0x1ca7be996ed9d2d754ffe235074db015647227b623cd505657924d996dbb6ba1d22e817c35cdb3e20bd0ffa29e8a0b6b0ede18b8b9c83b2275ce86a4284184456303",
		},
		{
			"0xbbde8ab3d1e178726af9f8802380e4de79e483fff8daf29ee9b86d87e3aebf12",
			"0x0CFa1FaF595594a7609c824530d5623c84d6df50",
			"0x1b917b943b6a2751eec7408376a07eca9b530bce2564e360377c73402bb16d797d24fbd2428bfde0cda5614abee424525a76c52e9e46da55441e283ac09e1d535f03",
		},
		// Legacy
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"0xadD132AD2b945CDBDC4Dd51C32744bfA4D5fC28d",
			"0x1c38ae2328af4da36d8f63b1b034e6de91f8488c4c0226754b0107b9a88532b1487932e9f29fecd5d483bae7b07df57a04b7ca61ce6dae029274f0cc0358cc6b7c",
		},
		{
			"0x8aa8e6cbe04f63443a71fd43d511883087df205e7b47f479bc616713d13ce0c7",
			"0x0CFa1FaF595594a7609c824530d5623c84d6df50",
			"0x1ca7be996ed9d2d754ffe235074db015647227b623cd505657924d996dbb6ba1d22e817c35cdb3e20bd0ffa29e8a0b6b0ede18b8b9c83b2275ce86a42841844563",
		},
		{
			"0xbbde8ab3d1e178726af9f8802380e4de79e483fff8daf29ee9b86d87e3aebf12",
			"0x0CFa1FaF595594a7609c824530d5623c84d6df50",
			"0x1b917b943b6a2751eec7408376a07eca9b530bce2564e360377c73402bb16d797d24fbd2428bfde0cda5614abee424525a76c52e9e46da55441e283ac09e1d535f",
		},
	} {
		suite.True(VerifySignature(
			hexutil.MustDecode(tt.expected),
			common.HexToAddress(tt.address),
			common.HexToHash(tt.orderHash)),
		)
	}
}

func TestSignatureSuite(t *testing.T) {
	suite.Run(t, new(SignatureSuite))
}
