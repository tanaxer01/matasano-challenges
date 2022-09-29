package set1

import (
	"Matasano/utils"
	"strings"
)

func Score(word string) (score float32) {
	word = strings.ToLower(word)
	freq := map[string]float32 {
		"a": .08167, "b": .01492, "c": .02782, "d": .04253,
		"e": .12702, "f": .02228, "g": .02015, "h": .06094,
		"i": .06094, "j": .00153, "k": .00772, "l": .04025,
		"m": .02406, "n": .06749, "o": .07507, "p": .01929,
		"q": .00095, "r": .05987, "s": .06327, "t": .09056,
		"u": .02758, "v": .00978, "w": .02360, "x": .00150,
		"y": .01974, "z": .00074, " ": .13000,
	}

	for i := 0; i < len(word); i++ {
		value, present := freq[ string(word[i]) ]
		if present { score += value }

	}

	return
}

func Xor_cipher(input string) (output string, char int, maximo float32) {
	ints := utils.From_hex(input)
	curr := make([]byte,len(ints))

	for i := 0; i<256; i++ {
		for j:=0; j < len(curr); j++ { curr[j] = ints[j] ^ byte(i) }
		score := Score( string(curr) )

		if score > maximo { output, maximo, char = string(curr), score, i }
	}

	return
}
