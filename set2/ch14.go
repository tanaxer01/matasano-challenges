package set2

import (
	"Matasano/set1"
	"Matasano/utils"
	"bytes"
	"errors"
	"log"
	"strings"
)

var (
	prefix_len int
	PREFIX     []byte
)

func oracle14(input []byte) []byte {
	target := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	plaintext := append(input, set1.B64Decode(target)...)
	plaintext = append(PREFIX, plaintext...)

	return set1.ECBEncrypt(plaintext, KEY)
}

func countRepeatedBlocks(a []byte, b []byte) int {
	for i := 0; i < len(a); i += 16 {
		if !bytes.Equal(a[i:i+16], b[i:i+16]) {
			return i / 16
		}
	}

	return len(a) / 16
}

func calcFillerLen() int {
	var prevCipher, currCipher []byte

	currCipher = oracle14(nil)
	for i := 1; i <= 32; i++ {
		prevCipher = currCipher
		currCipher = oracle14(bytes.Repeat([]byte("a"), i))

		if countRepeatedBlocks(prevCipher, currCipher) == 2 {
			return i
		}
	}

	return 32
}

func calcPadding(filler int) int {
	input := bytes.Repeat([]byte("A"), filler)
	base_len := len(oracle14(input))

	for len(oracle14(input)) == base_len {
		input = append(input, 'A')
	}

	return len(input) - filler
}

func matchOneByte(padding []byte, plain []byte, offset int) (rune, error) {
	ct := oracle14(padding)
	pt := append(padding, plain...)

	for i := 0; i < 256; i++ {
		option := append(pt, byte(i))
		option_ct := oracle14(option)

		if bytes.Equal(ct[offset:offset+16], option_ct[offset:offset+16]) {
			return rune(i), nil
		}
	}

	return 0, errors.New("no match found: " + string(pt[offset:offset+16]))
}

func Challenge14() {
	prefix_len, err := utils.RandomRange(1, 16)
	if err != nil {
		log.Fatalf("Error generating rand number: %s", err)
	}
	PREFIX = make([]byte, prefix_len)

	utils.RandomBytes(KEY)
	utils.RandomBytes(PREFIX)

	// We need to add 2 blocks of filler_len including the prefix
	filler_len := calcFillerLen()
	// We still might have padding at the end
	padding_len := calcPadding(filler_len)
	// Having the filler len and the padding len, we can calculate the length of the payload
	payload_len := len(oracle14(bytes.Repeat([]byte("a"), filler_len))) - 32 - padding_len

	// Having calculated those values, we can do the byte by byte decryption
	filler := bytes.Repeat([]byte("B"), filler_len-2)
	var plain []byte
	for j := 16; j < payload_len+16; j += 16 {
		for i := 0; i < 16; i++ {
			if len(plain) == payload_len {
				break
			}

			letter, err := matchOneByte(filler[i:], plain, j)
			if err != nil {
				log.Println("ERR", i, string(plain))
				panic(err)
			}

			plain = append(plain, byte(letter))
		}
	}

	log.Printf("\t[ch 14] %s...", strings.ReplaceAll(string(plain), "\n", "\\n"))
}
