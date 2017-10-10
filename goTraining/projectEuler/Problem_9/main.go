package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a, b, c := finder()
	fmt.Print("a:", a)
	fmt.Print(" b:", b)
	fmt.Println(" c:", c)
	fmt.Println("product:", a*b*c)
	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)

}

func finder() (int, int, int) {

	for c := 0; c <= 1000; c++ {
		for b := 0; b <= 1000; b++ {
			if b >= c {
				break
			}
			for a := 0; a <= 1000; a++ {
				if a >= c { //if a or b is larger than c, then c just needs to move on to the next number
					break
				}
				if a+b+c == 1000 && (a*a)+(b*b) == (c*c) {
					return a, b, c
				}
			}
		}
	}
	return 0, 0, 0 //no such number
}
