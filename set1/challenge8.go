package set1

import "Matasano/utils"

func DetectECB(path string) (ecb_block string, score int) {
	for _, line := range utils.ReadLines(path) {
		curr := utils.From_hex(line)

		blocks := make(map[string]int, 0)
		for i := 0; i < len(curr); i += 16 {
			curr := string(AES_decode([]byte("AAAAAAAAAAAAAAAA"), curr[i:i+16]))
			if _, ok := blocks[curr]; !ok { blocks[curr] = 1 } 
		}

		var total int = 0
		for _, value := range blocks { total += value }

		if total < score || score == 0 {
			ecb_block = string(utils.To_hex(curr))
			score = total
		}
	}

	return 
}
