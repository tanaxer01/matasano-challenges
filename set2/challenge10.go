package set2

import (
	"Matasano/set1"
)


func CBC_encode(key []byte, iv []byte, plain []byte) []byte {
	ret := make([]byte, len(plain))

	tmp := make([]byte, 16)
	aux := iv

	for i := 0; i < len(plain); i += 16 {
		tmp = set1.XOR_Bytes(aux, plain[i:i+16])
		tmp = set1.AES_encode(key, tmp)

		aux = tmp
		for j := 0; j < 16; j++ { ret[i+j] = tmp[j] }
	}
	
	return ret 
}

func CBC_decode(key []byte, iv []byte, cipher []byte) []byte {
	ret := make([]byte, len(cipher))

	tmp := make([]byte, 16) 
	aux := iv

	for i := 0; i < len(cipher); i += 16 {
		tmp = set1.AES_decode(key, cipher[i:i+16])
		tmp = set1.XOR_Bytes(tmp, aux)

		aux = tmp
		for j := 0; j < 16; j++ { ret[j] = tmp[j] }
	}

	return ret 
}
