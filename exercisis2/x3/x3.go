package main

import "fmt"

func main() {
	fmt.Println(greatestNumber(100000, 2, 3332, 58, 3554, 43, 321, 322, 4763, 327000))
}

func greatestNumber(is ...int) int {
	var x int

	for _, num := range is {

		if num > x {
			x = num
		}
	}
	return x
}
