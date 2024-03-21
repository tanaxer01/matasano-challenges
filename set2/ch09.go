package set2

import (
	"Matasano/set1"
	"log"
)

func Challenge9() {
	input := "YELLOW SUBMARINE"
	output := set1.Pkcs7([]byte(input), 20)

	if len(output) != 20 {
		log.Fatalf("padding result doesn't match expected length")
	}

	log.Printf("\t[ch 9] %s with target length = 20", output)
}
