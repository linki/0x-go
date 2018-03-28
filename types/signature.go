package types

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signature struct {
	V byte
	R common.Hash
	S common.Hash
}

func SignHash(orderHash common.Hash, privateKey *ecdsa.PrivateKey) (Signature, error) {
	magicMessage := crypto.Keccak256(
		[]byte("\x19Ethereum Signed Message:\n32"),
		orderHash.Bytes(),
	)

	sigBytes, err := crypto.Sign(magicMessage, privateKey)
	if err != nil {
		return Signature{}, err
	}

	signature := Signature{
		R: common.BytesToHash(sigBytes[0:32]),
		S: common.BytesToHash(sigBytes[32:64]),
		V: sigBytes[64] + 27,
	}

	return signature, nil
}
