package main

import (
	"Matasano/set1"
	"Matasano/set2"
	"os"
	"strconv"
)

func main() {
	var args [2]int

	setMap := map[int]func(int){
		1: set1.Main,
		2: set2.Main,
	}

	for i, arg := range os.Args[1:] {
		if i == 2 {
			break
		}

		num, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}

		args[i] = num
	}

	if args[0] > 0 {
		setMap[args[0]](args[1])
		return
	}

	for i := 1; i < 3; i++ {
		setMap[i](args[1])
	}
}
