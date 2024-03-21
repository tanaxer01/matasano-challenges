package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func B64Decode(msg string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		log.Fatalf("Error while decoding base64 string: %s", err)
	}

	return decoded
}

func HexDecode(msg string) []byte {
	decoded, err := hex.DecodeString(msg)
	if err != nil {
		log.Fatalf("Error while decoding hex string: %s", err)
	}

	return decoded
}

func B64Encode(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func HexEncode(msg []byte) string {
	return hex.EncodeToString(msg)
}

func Challenge1() {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	hexBytes := HexDecode(input)
	b64Bytes := B64Decode(output)

	if output != B64Encode(hexBytes) {
		log.Printf("Solution doesn't match. Expected: %s, Got: %s", output, B64Encode(hexBytes))
	}

	if input != HexEncode(b64Bytes) {
		log.Printf("Solution doesn't match. Expected: %s, Got: %s", input, HexEncode(b64Bytes))
	}

	log.Printf("\t[ch 1] %s", output)
}
