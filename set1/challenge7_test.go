package set1

import (
	"testing"
	"strings"
	"bufio"
	"os"
)

func TestChallenge7(t *testing.T) {
	var key  string = "YELLOW SUBMARINE"
	var path string = "test_data/7.txt"

	fd, _ := os.Open(path)
	defer fd.Close()

	var input string = ""
	
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() { input += strings.TrimSuffix(scanner.Text(),"\n") }

	var input_decoded = []byte(B642Hex(input))

	decoded  := AES_decode([]byte(key), input_decoded)
	received := AES_encode([]byte(key), decoded)

	t.Log("Output1 ==> ",string(decoded))
	t.Log("Output2 ==> ",string(received) == string(input_decoded))
	//t.Log("Output2 (", received == input_decoded, ") ==> ",string(decoded))


}
