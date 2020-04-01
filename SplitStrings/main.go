package main

import (
	"fmt"
	"math"
)

func Solution(str string) (pairsArray []string) {
	sub := ""

	if math.Mod(float64(len(str)), 2) == 1 {
		fmt.Printf("The length is odd\n")
		str += "_"
	} else {
		fmt.Printf("The length is even\n")
	}

	for i, r := range str {
		sub = sub + string(r)
		if (i+1)%2 == 0 {
			pairsArray = append(pairsArray, sub)
			sub = ""
		} else if (i + 1) == len(str) {
			pairsArray = append(pairsArray, sub)
		}
		fmt.Printf("i%d r%c\n", i, r)
	}

	return pairsArray
}

func main() {
	fmt.Printf("%s\n", Solution("abc"))
	fmt.Printf("%s\n", Solution("abcdef"))
}
