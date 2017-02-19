package main

import "fmt"

func main() {

	fmt.Println(half(2))
}

func half(x int) (int, bool) {

	y := x / 2

	if x%2 == 0 {
		return y, true
	}
	return y, false

}
