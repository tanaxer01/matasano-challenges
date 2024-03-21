package set1

import "log"

func DetectECB(input [][]byte) int {
	var best int
	var bestLine int

	for i, line := range input {
		pt := ECBDecrypt(line, []byte("AAAAAAAAAAAAAAAA"))

		// Create a "set" with all the blocks
		var count int
		ocurrences := make(map[string]int)
		for i := 0; i < len(pt); i += 16 {
			curr := string(pt[i : i+16])
			if _, ok := ocurrences[curr]; !ok {
				ocurrences[curr] = 1
				count++
			}
		}

		// Get the line with most repeated blocks
		if best > count || best == 0 {
			best = count
			bestLine = i
		}
	}

	return bestLine
}

func Challenge8() {
	hexLines, err := ReadLines("inputs/8.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	lines := make([][]byte, len(hexLines))
	for i, line := range hexLines {
		lines[i] = HexDecode(line)
	}

	ecbLine := DetectECB(lines)
	log.Printf("\t[ch 5] line %d - %s...", ecbLine, hexLines[ecbLine][:50])
}
