package set1

import "encoding/hex"


func XOR_Bytes(A []byte, B []byte) []byte {
	dst := make([]byte, len(A))
	for i:=0;i<len(A);i++ { dst[i] =  A[i] ^ B[i] }
	
	return dst
}
	
func XOR_Hex(A []byte, B []byte) string {
	dst := make([]byte, hex.EncodedLen(len(A)))
	_ =  hex.Encode(dst, XOR_Bytes(A, B))

	return string(dst)
}
