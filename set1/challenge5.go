package set1


func RepeatingKeyXor(key string, input string) string {
	var ints []byte = []byte(input)
	var intsKey []byte 

	for i:=0;i<len(ints)/len(key);i++ { intsKey = append(intsKey,[]byte(key)...) }
	intsKey = append(intsKey,[]byte(key)[:len(ints)%len(key)]...)

	var xored string = XOR_Hex(ints,intsKey)
	
	return xored
}
