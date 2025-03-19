package set2

import "log"

func Main(input int) {
	log.Println(" --> Set 2")

	challengeMap := map[int]func(){
		9:  Challenge9,
		10: Challenge10,
		11: Challenge11,
		12: Challenge12,
		13: Challenge13,
		14: Challenge14,
	}

	if input >= 9 && input <= 14 {
		challengeMap[input]()
		return
	}

	for _, chall := range challengeMap {
		chall()
	}
}
