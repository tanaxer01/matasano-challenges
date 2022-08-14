package set1

import (
	"encoding/base64"
	"encoding/hex"
)


func Hex2B64(input string) string {
	ints, _ := hex.DecodeString(input)
	encoded := base64.StdEncoding.EncodeToString(ints)

	return encoded
}

func B642Hex(input string) []byte {
	ints, _ := base64.StdEncoding.DecodeString(input)
	//decoded := hex.EncodeToString(ints)
	
	return ints
}
