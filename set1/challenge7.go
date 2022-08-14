package set1

import (
	"crypto/aes"
)

func AES_encode(key []byte, plain []byte) []byte {
	cipher, _ := aes.NewCipher(key)

	out := make([]byte, len(plain))
	for i := 0; i < len(plain); i += 16 {
		cipher.Encrypt(out[i:i+16], plain[i:i+16])
	}

	return out
}

func AES_decode(key []byte, encoded []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	
	out := make([]byte, len(encoded))
	for i := 0; i < len(encoded); i += 16 {
		cipher.Decrypt(out[i:i+16], encoded[i:i+16])
	}

	return out
}
