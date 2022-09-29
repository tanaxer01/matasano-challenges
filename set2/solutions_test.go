package set2

import (
	"Matasano/utils"
	"testing"
)

func TestChallenge9(t *testing.T) {
	val := utils.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	t.Log("[9] PKCS7 => ", string(val))
}

func TestChallenge10(t *testing.T) {
	var key  string = "YELLOW SUBMARINE"

	var input = utils.ReadFile("test_data/10.txt")
	var input_decoded = utils.From_base64(input)

	iv := make([]byte, 16)
	for i := 0; i < 16; i++ { iv[i] = 0x00 }

	deco := CBC_decode([]byte(key), iv, input_decoded)
	enco := CBC_encode([]byte(key), iv, deco)
	t.Logf("[10] CBC Mode => %v", string(enco) == string(input_decoded) )
}

func TestChallenge11(t *testing.T) {
	key := EncryptionOracle([]byte{ 1, 2, 3 })

	t.Logf("%v", key)
}
