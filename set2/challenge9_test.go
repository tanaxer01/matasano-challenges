package set2

import "testing"

func TestChallenge9(t *testing.T) {
	val := PKCS7([]byte("YELLOW SUBMARINE"), 20)
	t.Log("Output1", string(val))
}
