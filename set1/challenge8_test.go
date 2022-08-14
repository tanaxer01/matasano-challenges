package set1

import "testing"

func TestChallenge8(t *testing.T) {
	var path = "test_data/8.txt"

	text, score := DetectECB(path)

	t.Log("Output1 ==> ", score," - ", text)
}
