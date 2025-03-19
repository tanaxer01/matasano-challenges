package set1

import "log"

func Main(input int) {
	log.Println(" --> Set 1")

	challengeMap := map[int]func(){
		1: Challenge1,
		2: Challenge2,
		3: Challenge3,
		4: Challenge4,
		5: Challenge5,
		6: Challenge6,
		7: Challenge7,
		8: Challenge8,
	}

	if input >= 1 && input <= 8 {
		challengeMap[input]()
		return
	}

	for _, chall := range challengeMap {
		chall()
	}
}
