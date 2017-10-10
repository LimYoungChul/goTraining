package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(findAllThePrimes(2000000))
	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)
	//Duration notes
	//Without a primeCap it takes 1min 47 seconds, which fails our requirements
	//With a primeCap it takes ~385ms. Quite a nice upgrade :D
}

func findAllThePrimes(n int) int {

	var primeArray []int
	var primeCap int
	var isPrime bool

	primeArray = append(primeArray, 2, 3, 5)
	sum := 2 + 3 + 5

	for i := 6; i < n; i++ {
		primeCap = int(math.Sqrt(float64(i))) //squareroot of number we are checking.
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
			sum += i
		}
	}

	return sum

}
