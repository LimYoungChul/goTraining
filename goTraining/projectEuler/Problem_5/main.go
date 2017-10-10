package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(test())
	elapsed := time.Since(start)
	//takes about 2.3 seconds to run. Wildly inefficient, but nonetheless below 1 minute so ill keep it that way :D
	fmt.Println("runtime: ", elapsed) //as a side note, it takes twice the time to run if you don't modulo the highest numbers first.
}

func test() int {
	n := 20
	var found bool
	for {
		n++
		found = true
		for i := 20; i > 0; i-- { //since any number can do % of 1 and half the numbers can do % of 2 we need to start from the opposite direction to save resources
			if n%i != 0 {
				found = false
				break
			}
		}
		if found == true {
			return n
		}
	}
}
