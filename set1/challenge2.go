package set1

import "encoding/hex"


func Xor_Bytes(A []byte, B []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(A)))
	
	for i:=0;i<len(A);i++ { A[i] ^= B[i] }
	_ =  hex.Encode(dst,A)

	return dst
}
