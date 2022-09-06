package set2

import (
	"crypto/rand"
	"math/big"
)

func EncryptionOracle(input []byte) int64 {
	len, _ := rand.Int(rand.Reader, big.NewInt(6)) + 5

	key := make([]byte, 16)
	rand.Read(key)

	return 
}


