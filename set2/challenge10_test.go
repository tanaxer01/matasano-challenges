package set2

import (
	"testing"
	"strings"
	"bufio"
	"os"
)

func TestChallenge10(t *testing.T) {
	var path string = "test_data/10.txt"


	a := []byte("YELLOW SUBMARINE")
	b := []byte("TESTINGTHISSHIET")
	c := []byte("RANDOMWORDFORSHT")

	t.Log(c)
	val := CBC_encode(a,b,c)
	t.Log(val)
	val2 := CBC_decode(a,b,val)
	t.Log(string(val2))

}
