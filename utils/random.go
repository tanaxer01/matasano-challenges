package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomRange(min int, max int) (int, error) {
	randInt, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return min + int(randInt.Int64()), nil
}

func RandomBytes(input []byte) {
	rand.Read(input)
}
