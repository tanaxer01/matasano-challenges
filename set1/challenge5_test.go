package set1

import "testing"

func TestChallenge5(t *testing.T) {
	var key string = "ICE"
	var input string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	
	var expected string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	var received string = RepeatingKeyXor(key,input)

	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,received)
	} else {
		t.Log("Output1 (", received == expected, ") ==>", received)
	}
	
}
