package main

import "fmt"

func main() {
	findCombinations(9)
	//fmt.Println(buildWall(9, 3))
}

func buildWall(width, height int) int {
	array := make([][]int, height)

	var counter int

	for i := range array {
		array[i] = make([]int, width-1)
	}

	return counter
}

func findCombinations(width int) {
	var i int
	var c int
	var open []int
	var tmpArray []int
	var solutions []int

	open = append([]int{3, 2}, open...) //appends entry on the left
	tmpArray = append([]int{0}, tmpArray...)
	//open = append(open[:0], open[:1]...) // deletes first entry from the left
	for {
		if len(open) > 0 {
			tmpArray = append(tmpArray[:i], open[0])
			open = append(open[:0], open[1:]...)
			fmt.Println(tmpArray)
			i++

			counter := 0
			for _, x := range tmpArray {
				counter += x
			}
			if counter == width {
				fmt.Println("Succes")
				solutions = append(solutions[:c], tmpArray[0:]...)
				c++
				i--
			} else if counter > width {
				fmt.Println("overshot ")
				i--
			} else if counter < width {
				open = append([]int{3, 2}, open...)
			}
		} else {
			fmt.Print("this is a solution: ")
			fmt.Println(solutions)
			break
		}
	}
}
