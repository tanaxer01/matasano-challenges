package set1

import (
	"testing"
	"encoding/hex"
)

func TestChallenge2(t *testing.T) {
	intsA, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	intsB, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	
	var expected string = "746865206b696420646f6e277420706c6179"
	var received string = XOR_Hex(intsA,intsB)

	if expected != received {
		t.Errorf("Expected: %v, got: %v",expected,received)
	} else {
		t.Log("Output1 (", expected == received,") ==> ", received)
		t.Log( []byte(received) ) 
	}
}
