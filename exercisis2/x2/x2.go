package main

import "fmt"

func main() {

	half := func(x int) (int, bool) {

		y := x / 2

		if x%2 == 0 {
			return y, true
		}
		return y, false
	}

	fmt.Println(half(27))
}

/*
func half(x int) (int, bool) {

	y := x / 2

	if x%2 == 0 {
		return y, true
	}
	return y, false

}
*/
