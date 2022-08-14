package main

import (
	"Matasano/set1"
	"fmt"
)

func main(){
	var a []byte = []byte("this is a test")
	var b []byte = []byte("wokka wokka!!!")

	fmt.Println("CalcHamming ->",set1.CalcHamming(a,b))
	
	var path string = "set1/test_data/6.txt"
	aa := set1.BreakRepetingKeyXor(path)
	fmt.Println("AAA -> ", aa)
}
