package main

import "fmt"

func main() {
	fmt.Println(buildWall(9, 3))
}

func buildWall(width, height int) int {
	array := make([][]int, height)
	var counter int

	for i := range array {
		array[i] = make([]int, width)
	}

	return counter
}
