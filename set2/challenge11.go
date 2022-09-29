package set2

import (
	"Matasano/utils"
	"math/rand"
	"time"
)

func EncryptionOracle(input []byte) []byte {
	rand.Seed(time.Now().Unix())

	/*
	key := utils.RandomBytes(16) 
	iv  := utils.RandomBytes(16)
	*/

	pre := utils.RandomBytes( utils.RandomInt(5,10) )
	pos := utils.RandomBytes( utils.RandomInt(5,10) )
	end := utils.PKCS7(append(append(pre, input...), pos...), 16 * (1+(len(input)/16)) )

	return end
	/*
	rand.Seed(time.Now().UnixNano())

	key := make([]byte, 16)
	rand.Read(key)

	iv := make([]byte, 16)
	rand.Read(iv)

	pre := make([]byte, rand.Intn(6)+5)
	rand.Read(pre)

	pos := make([]byte, rand.Intn(6)+5)
	rand.Read(pos)

	if rand.Intn(2) & 1 {
		enco := CBC_decode([]byte(key), iv, input_decoded)
	} else {
		enco := set1.AES_decode([]byte(key), input_decoded)
	}

	*/

}


