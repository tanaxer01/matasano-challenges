package set1

import "Matasano/utils"

func DetectXor(path string) (word string){
	var maxScore float32 = 0
		
	for _, line := range utils.ReadLines(path) {
		curr, _, score := Xor_cipher(line)
		if score > maxScore { maxScore, word = score, curr }
	}

	return
}
