package set2

import (
	"Matasano/set1"
	"Matasano/utils"
	"log"
	"strings"
)

func oracle12(input []byte) []byte {
	target := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	plaintext := append(input, set1.B64Decode(target)...)

	return set1.ECBEncrypt(plaintext, KEY)
}

func calcPaddingLen() (int, int) {
	var padding string = ""

	empty_length := len(oracle12([]byte(padding)))
	for empty_length == len(oracle12([]byte(padding))) {
		padding = padding + "A"
	}

	return empty_length, len(padding)
}

func matchByte(filler []byte, plain []byte, offset int) string {
	options := make(map[string]byte)
	pt := append(filler, plain...)
	for i := 0; i < 256; i++ {
		possiblePt := append(pt, byte(i))
		ct := oracle12(possiblePt)

		options[string(ct[offset:offset+16])] = byte(i)
	}

	ct := oracle12(filler)
	if val, ok := options[string(ct[offset:offset+16])]; ok {
		return string(val)
	}

	return ""
}

func Byte2ByteDecrypt(targetLen int) string {
	var block string
	var plain string

	for i := 0; i < targetLen; i += 16 {
		filler := strings.Repeat("B", 15)

		for j := 0; j < 16; j++ {
			if len(plain)+len(block) == targetLen {
				break
			}

			letter := matchByte([]byte(filler), []byte(plain), i)
			if letter == "" {
				log.Fatalf("matchOneByte didn't find any match")
			}

			plain += letter
			if len(filler) > 0 {
				filler = filler[1:]
			}
		}
	}

	return plain
}

func Challenge12() {
	utils.RandomBytes(KEY)

	// 1. Discover the block size of the cipher. You know it, but do this step anyway.
	targetLen, paddingLen := calcPaddingLen()
	padding := []byte(strings.Repeat("A", paddingLen))

	blockLen := len(oracle12(padding)) - len(oracle12(nil))
	if blockLen != 16 {
		log.Fatalln("Did somethig wrong calculating the blockLen")
	}

	// 2. Detect that the function is using ECB. You already know, but do this step anyways.
	modeUsed := IsEcb(oracle12([]byte(strings.Repeat("A", 32))))
	if modeUsed != true {
		log.Fatalln("Did somethig wrong checking AES Mode")
	}

	// 3. Byte-at-a-time ECB decryption
	plainText := Byte2ByteDecrypt(targetLen - paddingLen)
	plainText = strings.ReplaceAll(plainText, "\n", "\\n")
	log.Printf("\t[ch 12] %s...", plainText[:80])
}
