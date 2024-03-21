package set1

import "log"

func ScoreWord(word string) float32 {
	charFreq := map[rune]float32{
		'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835, 'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610,
		'h': 0.0492888, 'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490, 'm': 0.0202124, 'n': 0.0564513,
		'o': 0.0596302, 'p': 0.0137645, 'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357, 'u': 0.0225134,
		'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692, 'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182,
	}

	var score float32 = 0
	for _, c := range word {
		score += charFreq[c]
	}

	return score
}

func BreakXorCipher(input []byte) (float32, rune) {
	var bestScore float32 = 0
	var bestRune rune

	key := make([]byte, len(input))
	for i := 0; i < 256; i++ {
		for j := range key {
			key[j] = byte(i)
		}

		xored := Xor(input, key)
		score := ScoreWord(string(xored))

		if score > bestScore {
			bestScore = score
			bestRune = rune(i)
		}
	}

	return bestScore, bestRune
}

func SingleByteXor(input []byte, key byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key
	}

	return output
}

func Challenge3() {
	encoded := HexDecode("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	score, key := BreakXorCipher(encoded)
	decoded := SingleByteXor(encoded, byte(key))

	log.Printf("\t[ch 3] %s with score %f", string(decoded), score)
}
