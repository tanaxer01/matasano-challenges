package set1

import "log"

func Xor(a []byte, b []byte) []byte {
	var lengthC int
	if len(a) > len(b) {
		lengthC = len(a)
	} else {
		lengthC = len(b)
	}

	c := make([]byte, lengthC)
	for i := 0; i < lengthC; i++ {
		c[i] = a[i] ^ b[i]
	}

	return c
}

func Challenge2() {
	a := HexDecode("1c0111001f010100061a024b53535009181c")
	b := HexDecode("686974207468652062756c6c277320657965")
	c := "746865206b696420646f6e277420706c6179"

	xored := Xor(a, b)
	if HexEncode(xored) != c {
		log.Fatalf("Error while xoring bytes. Expected: %s, Got: %s", c, HexEncode(xored))
	}

	log.Printf("\t[ch 2] %s", HexEncode(xored))
}
