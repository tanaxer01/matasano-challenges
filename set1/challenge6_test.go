package set1

import "testing"

func TestChallenge6(t *testing.T) {
	var a []byte = []byte("this is a test")
	var b []byte = []byte("wokka wokka!!!")

	t.Log("Output1 (", CalcHamming(a,b) == 37, ") ==> ", CalcHamming(a,b))
	
	var path string = "test_data/6.txt"
	key, text := BreakRepetingKeyXor(path)
	t.Log("Output2 ==>", key )
	t.Log("Output3 ==>", text)
}
