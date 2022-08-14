package set1

import (
	"bufio"
	"os"
)

func DetectXor(path string) (word string){
	fd, _  := os.Open(path)
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	 word = ""
	var maxScore float32 = 0
	
	for scanner.Scan() {
		line := scanner.Text()
		
		curr, _, score := Xor_cipher(line)
		
		if score > maxScore {
			maxScore = score
			word = curr
		}
	}

	return
}
