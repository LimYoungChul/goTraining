package main

import "fmt"

func main() {
	fmt.Println(wall(32, 10))
}

func wall(width, height int) int {

	total := 0

	combinations := findCombinations(width)

	fmt.Println(len(combinations))

	for _, a := range combinations {

		buildWall(a, combinations, 0, height, &total)

	}

	return total
}

func buildWall(a []int, combi [][]int, level, height int, total *int) {
	level++
	var canCombine bool

	for _, a1 := range combi {
		canCombine = true
		for _, value := range a {
			if canCombine == false {
				break
			}
			for _, value1 := range a1 {
				if value == value1 {
					canCombine = false
					break
				}

			}
		}
		if canCombine == true && level < height {
			buildWall(a1, combi, level, height, total)
		} else if level == height {
			*total++
			break
		}
	}
}

func findCombinations(width int) [][]int {
	var i int
	//var c int
	var tmp int
	open := make([][]int, 0, 100)
	solutions := make([][]int, 0, 100)
	open = append(open, []int{3, 2})
	tmpArray := make([]int, 0, 100)

	for {
		if len(open[i]) > 0 {
			tmpArray = append(tmpArray[:i], open[i][0])
			open[i] = append(open[i][:0], open[i][1:]...)
			counter := 0
			for _, x := range tmpArray {
				counter += x
			}
			if counter == width {
				solutionArray := make([]int, len(tmpArray)-1)
				counter2 := 0
				for n := 0; n < len(tmpArray)-1; n++ {
					counter2 += tmpArray[n]
					solutionArray[n] = counter2
				}
				solutions = append(solutions, solutionArray)

				for _, v := range open {
					tmp += len(v)
				}
				if tmp == 0 {
					return solutions
				}
				tmp = 0
			} else if counter > width {
				for _, v := range open {
					tmp += len(v)
				}
				if tmp == 0 {
					return solutions
				}
				tmp = 0
			} else if counter < width {
				i++
				if len(open) <= i {
					open = append(open, []int{3, 2})
				} else {
					open[i] = append(open[i], []int{3, 2}...)
				}
			}
		} else {
			i--
		}
	}
}
