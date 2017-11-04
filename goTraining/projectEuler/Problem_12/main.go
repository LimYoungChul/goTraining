package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	i := 0
	var x int
	var n int

	for x < 500 {
		i++
		n = i * (i + 1) / 2
		x = factorCounter(n)
	}

	fmt.Println(n)
	fmt.Println(x)
	elapsed := time.Since(start)
	fmt.Println("Runtime", elapsed)
}

func factorCounter(number int) int {
	counter := 0
	d := number

	for i := 1; i < number; i++ {
		if i >= d { //this should ALWAYS break the loop. i guess unless the number is a prime, but we never pass the function any primes i think
			break
		}
		if number%i == 0 {
			counter += 2 //we add 2, cause everytime we find a number we also find one in the other end. eg: 20 gives 1 and 20, 2 and 10, 4 and 5. "d will be = 5 and "i" will be 5 next time so it breaks the loop
			d = number / i
		}
	}

	return counter
}
