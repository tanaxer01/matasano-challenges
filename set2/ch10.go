package set2

import (
	"Matasano/set1"
	"crypto/aes"
	"log"
	"strings"
)

func CBCEncrypt(input []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	input = set1.Pkcs7(input, len(input)+16-(len(input)%16))

	output := make([]byte, len(input))
	temp := make([]byte, 16)
	for i := 0; i < len(input); i += 16 {
		temp = set1.Xor(input[i:i+16], iv)
		block.Encrypt(output[i:i+16], temp)
		iv = output[i : i+16]
	}

	return output
}

func CBCDecrypt(input []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creating AES cipher: %s", err)
	}

	output := make([]byte, len(input))
	temp := make([]byte, 16)
	for i := 0; i < len(input); i += 16 {
		block.Decrypt(temp, input[i:i+16])
		temp = set1.Xor(temp, iv)

		for j := 0; j < 16; j++ {
			output[i+j] = temp[j]
		}

		iv = input[i : i+16]
	}

	return output
}

func Challenge10() {
	key := "YELLOW SUBMARINE"
	iv := make([]byte, 16)

	input, err := set1.ReadFile("inputs/10.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	ciphertext := set1.B64Decode(input)
	plaintext := string(CBCDecrypt(ciphertext, []byte(key), iv))
	plaintext = strings.ReplaceAll(plaintext, "\n", "\\n")

	log.Printf("\t[ch 10] %s...", plaintext[:80])
}
