package set1

import "Matasano/utils"

func RepeatingKeyXor(key string, input string) (xored string) {
	ints := []byte(input)
	var intsKey []byte 

	for i := 0; i < len(ints)/len(key); i++ { intsKey = append(intsKey,[]byte(key)...) }
	intsKey = append(intsKey,[]byte(key)[:len(ints)%len(key)]...) 

	xored = string(utils.To_hex(utils.XOR(ints,intsKey)))
	return 
}
