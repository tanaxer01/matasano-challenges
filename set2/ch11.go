package set2

import (
	"Matasano/set1"
	"crypto/rand"
	"log"
	"math/big"
)

var (
	KEY = make([]byte, 16)
	IV  = make([]byte, 16)
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

func oracle11(input []byte) ([]byte, bool) {
	var output []byte

	prefixLen, err := RandomRange(5, 11)
	if err != nil {
		log.Fatalf("Error generating random number: %s", err)
	}
	sufixLen, err := RandomRange(5, 11)
	if err != nil {
		log.Fatalf("Error generating random number: %s", err)
	}

	prefix := make([]byte, prefixLen)
	sufix := make([]byte, sufixLen)

	rand.Read(prefix)
	rand.Read(sufix)

	modedInput := append(append(prefix, input...), sufix...)
	choice, err := RandomRange(0, 2)
	if err != nil {
		log.Fatalf("Error generating random number: %s", err)
	}

	if choice == 1 {
		output = set1.ECBEncrypt(modedInput, KEY)
	} else {
		output = CBCEncrypt(modedInput, KEY, IV)
	}

	return output, choice == 1
}

func IsEcb(input []byte) bool {
	var count int
	ocurrences := make(map[string]int)
	for i := 0; i < len(input); i += 16 {
		curr := string(input[i : i+16])
		if _, ok := ocurrences[curr]; !ok {
			ocurrences[curr] = 1
			count++
		}
	}

	return count != len(input)/16
}

func Challenge11() {
	RandomBytes(KEY)
	RandomBytes(IV)

	input := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	output, choice := oracle11([]byte(input))
	res := IsEcb(output)

	if res != choice {
		log.Fatalf("Mode guess didn't match the expected result. Expected: %b, Got: %b", choice, res)
	}

	log.Printf("\t[ch 11] ECB Mode used: %b", res)
}
