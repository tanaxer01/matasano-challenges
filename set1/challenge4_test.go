package set1

import "testing"

func TestChallenge4(t *testing.T) {
	var path string = "test_data/4.txt"
	var expected string = "Now that the party is jumping\n"
	var received string = DetectXor(path)

	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,received)
	} else {
		t.Log("Output1 (", received == expected,") ==>", received)
	}
}
