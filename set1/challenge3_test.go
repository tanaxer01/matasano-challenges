package set1

import "testing"

func TestChallenge3(t *testing.T) {
	var input    string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var expected string = "Cooking MC's like a pound of bacon"
	received, _, maximum := Xor_cipher(input)
	
	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,received)
	} else {
		t.Log("Output1 (", received == expected, ") ==>", received, maximum)
	}
}
