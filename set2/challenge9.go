package set2

func PKCS7(input []byte, target int) []byte {
	var padd  int = target - (len(input) % target)
	var res = make([]byte, len(input)+padd)

	for i := 0; i < len(input)+padd; i++ {
		if i < len(input) {
			res[i] = input[i]
		} else {
			res[i] = byte(padd)
		}
	}

	return res 
}
