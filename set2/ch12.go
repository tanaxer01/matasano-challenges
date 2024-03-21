package set2

import (
	"Matasano/set1"
	"log"
	"strings"
)

func Oracle12(input []byte) []byte {
	target := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	plaintext := append(input, set1.B64Decode(target)...)

	return set1.ECBEncrypt(plaintext, KEY)
}

func calcPaddingLength() (int, int) {
	var padding string = ""

	empty_length := len(Oracle12([]byte(padding)))
	for empty_length == len(Oracle12([]byte(padding))) {
		padding = padding + "A"
	}

	return len(padding), empty_length
}

func Challenge12() {
	RandomBytes(KEY)

	paddingLength, Length := calcPaddingLength()
	neededPad := strings.Repeat("A", paddingLength)

	blockSize := len(Oracle12([]byte(neededPad))) - len(Oracle12(nil))
	modeUsed := IsEcb(Oracle12([]byte(strings.Repeat("A", 32))))

	log.Println(blockSize, modeUsed)

	/*
			var plaintext []byte
			padding := strings.Repeat("A", 15)
		    for i :=0; i < Length; i++ {

		    }
	*/
}
