package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	number := 600851475143
	largestFactor := 0

	for i := 2; i <= number; i++ {
		if number%i == 0 {

			number /= i
			if largestFactor < i {
				largestFactor = i
			}
			i = 2
		}

	}
	fmt.Println("The largest factor of our number is:", largestFactor)
	elapsed := time.Since(start)
	fmt.Println("Time elapsed:", elapsed)
}
