package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var chainCounter int
	var highestChain int
	var hv int

	for i := 14; i < 1000000; i++ {
		chainCounter = 0

		for x := i; x != 1; chainCounter++ {
			if x%2 == 0 {
				x /= 2
			} else {
				x = 3*x + 1
			}
		}
		if chainCounter > highestChain {
			hv = i
			highestChain = chainCounter
		}
	}

	fmt.Println(hv)              //837799, correct answer
	fmt.Println(highestChain)    //524
	elapsed := time.Since(start) // 379ms

	fmt.Println("Runtime:", elapsed)
}
