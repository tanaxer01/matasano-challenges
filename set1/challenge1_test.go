package set1

import (
	"encoding/hex"
	"testing"
)


func TestChallenge1(t *testing.T) {
	var input    string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	var received string = Hex2B64(input)
	var recoded  []byte = B642Hex(received)

	if expected != received {
		t.Errorf("Expected: %v, got: %v",expected,received)
	} else {
		t.Log("Output1 (", received == expected,      ") ==> ", received)
		t.Log("Output2 (", hex.EncodeToString(recoded) == input, ") ==> ", string(recoded))		
	}
}
