package types

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	typeEIP712  = 2
	typeEthSign = 3
)

type Signature struct {
	V byte
	R common.Hash
	S common.Hash
}

func SignHash(orderHash common.Hash, privateKey *ecdsa.PrivateKey) (string, error) {
	return SignHashEIP712(orderHash, privateKey)
}

func SignHashEIP712(orderHash common.Hash, privateKey *ecdsa.PrivateKey) (string, error) {
	sigBytes, err := crypto.Sign(orderHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	signature := make([]byte, 66)
	signature[0] = sigBytes[64] + 27
	copy(signature[1:33], sigBytes[0:32])
	copy(signature[33:65], sigBytes[32:64])
	signature[65] = typeEIP712

	return hexutil.Encode(signature), nil
}

func SignHashEthSign(orderHash common.Hash, privateKey *ecdsa.PrivateKey) (string, error) {
	magicMessage := crypto.Keccak256(
		[]byte("\x19Ethereum Signed Message:\n32"),
		orderHash.Bytes(),
	)

	sigBytes, err := crypto.Sign(magicMessage, privateKey)
	if err != nil {
		return "", err
	}

	signature := make([]byte, 66)
	signature[0] = sigBytes[64] + 27
	copy(signature[1:33], sigBytes[0:32])
	copy(signature[33:65], sigBytes[32:64])
	signature[65] = typeEthSign

	return hexutil.Encode(signature), nil
}

func VerifySignature(signature []byte, address common.Address, hash common.Hash) bool {
	if len(signature) != 66 {
		if len(signature) != 65 {
			log.Printf("Invalid length: %v", len(signature))
			return false
		}
		return VerifySignatureEthSign(append(signature, typeEthSign), address, hash)
	}
	if signature[65] == typeEIP712 {
		return VerifySignatureEIP712(signature, address, hash)
	}
	if signature[65] == typeEthSign {
		return VerifySignatureEthSign(signature, address, hash)
	}
	return false
}

func VerifySignatureEIP712(signature []byte, address common.Address, hash common.Hash) bool {
	if len(signature) != 66 {
		log.Printf("Invalid length: %v", len(signature))
		return false
	}
	sigBytes := make([]byte, len(signature))
	copy(sigBytes, signature)
	v := sigBytes[0]
	r := sigBytes[1:33]
	s := sigBytes[33:65]
	if v < 27 {
		return false
	}
	pub, err := crypto.Ecrecover(hash.Bytes(), append(append(r, s...), v-27))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	recoverAddress := common.BytesToAddress(crypto.Keccak256(pub[1:])[12:])
	return bytes.Equal(address[:], recoverAddress[:])
}

func VerifySignatureEthSign(signature []byte, address common.Address, hash common.Hash) bool {
	if len(signature) != 66 {
		log.Printf("Invalid length: %v", len(signature))
		return false
	}
	sigBytes := make([]byte, len(signature))
	copy(sigBytes, signature)
	v := sigBytes[0]
	r := sigBytes[1:33]
	s := sigBytes[33:65]
	if v < 27 {
		return false
	}
	hashedBytes := append([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash[:]))), hash[:]...)
	signedBytes := crypto.Keccak256(hashedBytes)
	pub, err := crypto.Ecrecover(signedBytes, append(append(r, s...), v-27))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	recoverAddress := common.BytesToAddress(crypto.Keccak256(pub[1:])[12:])
	return bytes.Equal(address[:], recoverAddress[:])
}
