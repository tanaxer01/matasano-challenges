package set1

import (
	"log"
)

func HammingDistance(a []byte, b []byte) int {
	var score int

	for _, c := range Xor(a, b) {
		for i := 7; i >= 0; i-- {
			score += (int(c) >> uint(i)) & 1
		}
	}

	return score
}

func EditDistance(input []byte, ks int) float32 {
	var totalDistance int = 0

	for i := 0; i < len(input)-(2*ks); i += ks {
		totalDistance += HammingDistance(input[i:i+ks], input[i+ks:i+(2*ks)]) / ks
	}

	return float32(totalDistance) / (float32(len(input)/ks + 1))
}

func BreakRepeatingXor(input []byte) string {
	var keySize int
	var minDistance float32 = 0

	for ks := 2; ks <= 42; ks++ {
		ed := EditDistance(input, ks)
		if ed < minDistance || minDistance == 0 {
			keySize = ks
			minDistance = ed
		}
	}

	var key string

	block := make([]byte, len(input)/keySize+1)
	for i := 0; i < keySize; i++ {
		for j := i; j < len(input); j += keySize {
			block[j/keySize] = input[j]
		}

		_, chr := BreakXorCipher(block)
		key += string(chr)
	}

	return key
}

func Challenge6() {
	input, err := ReadFile("inputs/6.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	ciphertext := B64Decode(input)
	key := BreakRepeatingXor(ciphertext)

	log.Printf("\t[ch 5] %s", key)
}
