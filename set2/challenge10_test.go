package set2

import "testing"

func TestChallenge10(t *testing.T) {
	a := []byte("YELLOW SUBMARINE")
	b := []byte("TESTINGTHISSHIET")
	c := []byte("RANDOMWORDFORSHT")

	val := CBC_encode(a,b,c)
	t.Log(val)
	val2 := CBC_decode(a,b,val)
	t.Log(string(val2))

}
