package set2

import "testing"

func TestChallenge11(t *testing.T) {
	a := EncryptionOracle([]byte{10, 10})
	t.Log(a)
}
