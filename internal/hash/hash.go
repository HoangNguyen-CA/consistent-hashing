package hash

import (
	"crypto/sha256"
	"math/big"
)

func Hash(data []byte) *big.Int {
	h := sha256.Sum256(data)
	return new(big.Int).SetBytes(h[:])
}
