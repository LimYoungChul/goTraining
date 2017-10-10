package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	number := 10001

	lotsOfPrimes := findAlotOfPrimes(number)

	fmt.Println(lotsOfPrimes[number-1]) //10000 in the array is number 10001 prime.

	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)
}

func findAlotOfPrimes(n int) []int {

	var primeArray []int
	var primeCap int
	var isPrime bool

	primeArray = append(primeArray, 2, 3, 5)

	for i := 6; len(primeArray) < n; i++ {
		primeCap = int(math.Sqrt(float64(i))) //squareroot of number we are checking
		isPrime = true
		for x := 0; x < len(primeArray); x++ {
			if primeCap < primeArray[x] { //If the position x in our primearray exceeds our primecap, it means that this is a prime!
				break
			}
			if i%primeArray[x] == 0 { //If i is divisible by a prime in our array then it is not a prime itself.
				isPrime = false
				break
			}
		}
		if isPrime == true {
			primeArray = append(primeArray, i) //append when true
		}
	}

	return primeArray

}
