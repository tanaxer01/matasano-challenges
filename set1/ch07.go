package set1

import (
	"crypto/aes"
	"log"
	"strings"
)

func Pkcs7(input []byte, targetLength int) []byte {
	output := make([]byte, targetLength)

	for i := 0; i < targetLength; i++ {
		if i < len(input) {
			output[i] = input[i]
		} else {
			output[i] = byte(targetLength - len(input))
		}
	}

	return output
}

func ECBEncrypt(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	input = Pkcs7(input, len(input)+16-(len(input)%16))

	output := make([]byte, len(input))
	for i := 0; i < len(input); i += 16 {
		block.Encrypt(output[i:i+16], input[i:i+16])
	}

	return output
}

func ECBDecrypt(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	output := make([]byte, len(input))
	for i := 0; i < len(input); i += 16 {
		block.Decrypt(output[i:i+16], input[i:i+16])
	}

	return output
}

func Challenge7() {
	key := "YELLOW SUBMARINE"
	input, err := ReadFile("inputs/7.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	ciphertext := B64Decode(input)
	plaintext := ECBDecrypt(ciphertext, []byte(key))
	text, _, _ := strings.Cut(string(plaintext), "\n")

	log.Printf("\t[ch 7] %s...", text)
}
