package set2

import "Matasano/set1"


func CBC_encode(key []byte, iv []byte, plain []byte) []byte {
	ret := make([]byte, len(plain))
	tmp := make([]byte, 16)

	for i := 0; i < len(plain); i += 16 {
		/*
		if i == 0 {
			tmp = set1.Xor_Bytes(plain[i:i+16],iv) 
		} else {
			tmp = set1.Xor_Bytes(plain[i:i+16],ret[i-16:i]) 
		}
		*/

		tmp  = set1.AES_encode(key, tmp)
		for j := 0; j < 16; j++ { ret[i+j] = tmp[j] }
	}
	
	return ret 
}

func CBC_decode(key []byte, iv []byte, cipher []byte) []byte {
	ret := make([]byte, len(cipher))
	tmp := make([]byte, 16) 

	for i := 0; i < len(cipher); i += 16 {
		/*
		if i == 0 {
			tmp = set1.Xor_Bytes(tmp, iv)
		} else {
			tmp = set1.Xor_Bytes(tmp, cipher[i-16:i])
		}
		*/
		tmp = set1.AES_decode(key, cipher[:16])
		for j := 0; j < 16; j++ { ret[j] = tmp[j] }
	}

	return ret 
}
