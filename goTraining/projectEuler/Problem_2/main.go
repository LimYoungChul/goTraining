package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	f1 := 1 // Prior number
	f2 := 2 // Current number
	var sum int

	for {
		if f2%2 == 0 {
			sum += f2
		} else if f2 > 4000000 {
			break
		}
		f2 = f1 + f2 //update current number
		f1 = f2 - f1 //update prior number
	}
	fmt.Println("sum is equal to:", sum) //4613732, which is the correct answer
	elapsed := time.Since(start)
	fmt.Println("total runtime:", elapsed)

}
