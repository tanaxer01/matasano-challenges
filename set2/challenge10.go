package set2

import "Matasano/set1"


func CBC_encode(key []byte, iv []byte, plain []byte) []byte {
	ret := make([]byte, len(plain))

	var temp []byte
	for i := 0; i < len(plain); i += 16 {
		if i == 0 {
			temp = set1.AES_encode(key,iv)
		} else {
			temp = set1.AES_encode(key, ret[i-16:i])
		}

		for j := 0; j < 16; j++ { ret[j] = temp[j] }
	}

	return ret
}

func CBC_decode(key []byte, iv []byte, cipher []byte) []byte {
	ret := make([]byte, len(cipher))

	var temp []byte
	for i := len(cipher); i > 0 ; i -= 16 {
		if i == 0 {
			temp = set1.AES_decode(key,iv)
		} else {
			temp = set1.AES_encode(key, ret[i-16:i])
		}

		for j := 0; j < 16; j++ { ret[i-j] = temp[i-j] }
	}

	return ret
}
