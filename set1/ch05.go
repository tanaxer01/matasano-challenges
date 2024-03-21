package set1

import (
	"log"
	"strings"
)

func RepeatingKeyXor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = key[i%3]
	}

	return Xor(output, input)
}

func Challenge5() {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	output := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	key := "ICE"

	encoded := RepeatingKeyXor([]byte(input), []byte(key))

	if HexEncode(encoded) != output {
		log.Fatalf("Solution doesnt match expected result. Expected: %s, Got: %s", output, HexEncode(encoded))
	}

	input = strings.ReplaceAll(input, "\n", "\\n")
	log.Printf("\t[ch 5] %s", input)
}
