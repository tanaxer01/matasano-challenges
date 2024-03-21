package set1

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr := scanner.Text()

		lines = append(lines, curr)
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

func ReadFile(path string) (string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return "", err
	}

	return strings.Join(lines, ""), nil
}

func Challenge4() {
	lines, err := ReadLines("inputs/4.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	var bestKey rune
	var bestLine string
	var bestScore float32
	for _, l := range lines {
		lineBytes := HexDecode(l)
		score, key := BreakXorCipher(lineBytes)

		if score > bestScore {
			bestKey = rune(key)
			bestLine = l
			bestScore = score
		}
	}

	decoded := string(SingleByteXor(HexDecode(bestLine), byte(bestKey)))
	decoded = strings.ReplaceAll(decoded, "\n", "\\n")
	log.Printf("\t[ch 4] %s with score %f", decoded, bestScore)
}
