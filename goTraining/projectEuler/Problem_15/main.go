package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := 20
	fmt.Println(pathFinder(2*n, n)) // Takes around 1ms to compute result. Hardly changes when increasing n
	//fmt.Println(routeFinder(0, 0, 0)) // Takes around an hour to compute result for n=20. n++ quadruples the time it takes to compute.
	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)
}

//i have no idea how this works to be honest.
//Need to read up on https://en.wikipedia.org/wiki/Binomial_coefficient but not right now. Im lazy :D
func pathFinder(n, k int) int {

	if k > n {
		panic(fmt.Sprintf("%v, %v", n, k))
	}
	if k == 0 {
		return 1
	}
	if k > n/2 {
		return pathFinder(n, n-k)
	}
	return n * pathFinder(n-1, k-1) / k
}

//works, but too slow. Bruteforce way to find all solutions by counting them.
func routeFinder(x, y, solutions int) int {
	movedRight := false
	movedDown := false
	size := 17

	if movedRight == false && x < size {
		movedRight = true

		solutions = routeFinder(x+1, y, solutions)
	}

	if movedDown == false && y < size {
		movedDown = true

		solutions = routeFinder(x, y+1, solutions)
	}

	if x == size && y == size {
		solutions++
		return solutions
	}

	if x == size {
		movedRight = true
	}

	if y == size {
		movedDown = true
	}

	if movedDown == true && movedRight == true {
		return solutions
	}
	//should NOT reach this
	return 0
}
