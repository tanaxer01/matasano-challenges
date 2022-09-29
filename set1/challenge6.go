package set1

import "Matasano/utils"

func CountOnes(word []byte) int {
	var count int = 0
	
	for i:=0;i < len(word);i++{
		curr := int(word[i])

		for ;curr!=0;curr>>=1 { count += curr&1 }
	}
	
	return count
}

func CalcHamming(a []byte, b []byte) int {
	dst := make([]byte,len(a))
	for i := 0;i < len(a);i++ { dst[i] = a[i]^b[i] }
	
	return CountOnes(dst)
}

func CalcScore(num int, encoded []byte) float32 {
	var largo int = len(encoded)/(2*num) - 1
	var suma float32 = 0
	
	for i:=0;i<largo;i++ {
		score := float32( CalcHamming(encoded[i*num:(i+1)*num],encoded[(i+1)*num:(i+2)*num]))	
		suma += score/float32(num)
	}

	return suma/float32(largo)
}

func BreakRepetingKeyXor(path string) (KEY string,TEXT string) {
	//1. Fetch input
	var input = utils.ReadFile(path)
	var input_decoded = utils.From_base64(input)

	// 2. Calc KEYSIZE
	var KEYSIZE int
	var promedio float32 = 0
	
	for i:=2;i < 42;i++ {
		curr := CalcScore(i,input_decoded)
		if curr < promedio || promedio == 0 { promedio, KEYSIZE = curr, i }
	}

	//2. -> CalcHamming
	byte_key := make([]byte, KEYSIZE)

	for i := 0; i < KEYSIZE; i++ {
		temp := make([]byte, len(input_decoded)/KEYSIZE + 1)
		for j := i; j < len(input_decoded); j += KEYSIZE { temp[j/KEYSIZE] = input_decoded[j] }

		_, char, _ := Xor_cipher( string(utils.To_hex(temp)) )
		byte_key[i] = byte(char)  
	}

	KEY  = string(byte_key)
	TEXT = string( utils.From_hex( RepeatingKeyXor(KEY, string(input_decoded)) ) )

	return 
}
