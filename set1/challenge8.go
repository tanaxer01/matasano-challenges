package set1

import (
	"encoding/hex"
	"strings"
	"bufio"
	"os"
)

func DetectECB(path string) (ecb_block string, score int) {
	fd, _ := os.Open(path)
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		curr, _ := hex.DecodeString(strings.TrimSuffix(scanner.Text(),"\n"))

		blocks := make(map[string]int, 0)
		for i := 0; i < len(curr); i += 16 {
			curr := string(AES_decode([]byte("AAAAAAAAAAAAAAAA"), curr[i:i+16]))
			if _, ok := blocks[curr]; !ok { blocks[curr] = 1 } 
		}

		var total int = 0
		for _, value := range blocks { total += value }

		if total < score || score == 0 {
			ecb_block = hex.EncodeToString(curr)
			score = total
		}
	}

	return 
}
