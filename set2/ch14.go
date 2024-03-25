package set2

import (
	"Matasano/set1"
	"log"
)

var (
	prefix_len int
	PREFIX     []byte
)

func oracle14(input []byte) []byte {
	target := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	plaintext := append(input, set1.B64Decode(target)...)
	plaintext = append(PREFIX, plaintext...)

	log.Println(string(plaintext))

	return set1.ECBEncrypt(plaintext, KEY)
}

func Challenge14() {
	prefix_len, err := RandomRange(1, 16)
	if err != nil {
		log.Fatalf("Error generating rand number: %s", err)
	}
	PREFIX = make([]byte, prefix_len)

	RandomBytes(KEY)
	RandomBytes(PREFIX)

	oracle14(nil)
}
