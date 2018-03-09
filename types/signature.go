package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type Signature struct {
	V byte
	R common.Hash
	S common.Hash
}
