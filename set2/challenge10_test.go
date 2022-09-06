package set2

import (
	"Matasano/set1"
	"testing"
	"strings"
	"bufio"
	"os"
)

func TestChallenge10(t *testing.T) {
	var key  string = "YELLOW SUBMARINE"
	var path string = "test_data/10.txt"

	iv := make([]byte, 16)
	for i := 0; i < 16; i++ { iv[i] = 0x00 }

	fd, _ := os.Open(path)
	defer fd.Close()

	var input string = ""
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() { input += strings.TrimSuffix(scanner.Text(),"\n") }

	var input_decoded = []byte(set1.B642Hex(input))

	deco := CBC_decode([]byte(key), iv, input_decoded)
	t.Log("Output 1 => ", string(deco))

	enco := CBC_encode([]byte(key), iv, deco)
	t.Log("Output 2 => ", string(enco) == string(input_decoded))

}
