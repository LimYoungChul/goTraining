package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(squareSum(100) - sumSquare(100))

	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)
}

func sumSquare(n int) int {

	var sum int

	for i := 1; i < n+1; i++ {
		sum += i * i
	}
	return sum
}

func squareSum(n int) int {

	var sum int
	var sumSquared int
	for i := 1; i < n+1; i++ {
		sum += i
	}
	sumSquared = sum * sum

	return sumSquared

}
