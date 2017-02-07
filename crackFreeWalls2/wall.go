package main

import "fmt"

func main() {
	fmt.Println(wall(20, 10))
}

func wall(width, height int) int64 {

	var total int64

	combinations := findCombinations(width)

	combiPointer := &combinations

	for i := 0; i < 4; i++ {
		for _, a := range combinations[i] {
			if i%2 == 0 {
				buildWall(a, combiPointer[i+1], 0, height, i+1, &total, combiPointer)
			} else {
				buildWall(a, combiPointer[i-1], 0, height, i-1, &total, combiPointer)
			}
		}
	}
	return total
}

func buildWall(a []int, combi [][]int, level, height, index int, total *int64, pointer *[4][][]int) {
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

			if index%2 == 0 {
				buildWall(a1, pointer[index+1], level, height, index+1, total, pointer)
			} else {
				buildWall(a1, pointer[index-1], level, height, index-1, total, pointer)
			}
		} else if level == height {
			*total++
			break
		}
	}
}

func findCombinations(width int) [4][][]int {
	var i int
	var tmp int
	var tmpInt1 int
	var tmpInt2 int
	open := make([][]int, 0, 100)
	var solutionsHolder [4][][]int
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
					if n == 0 {
						tmpInt1 = tmpArray[n] % 2
					}
					counter2 += tmpArray[n]
					solutionArray[n] = counter2
				}
				tmpInt2 = counter2 % 2
				if tmpInt1 == 0 && tmpInt2 == 0 {
					solutionsHolder[0] = append(solutionsHolder[0], solutionArray)
				} else if tmpInt1 == 1 && tmpInt2 == 1 {
					solutionsHolder[1] = append(solutionsHolder[1], solutionArray)
				} else if tmpInt1 == 1 && tmpInt2 == 0 {
					solutionsHolder[2] = append(solutionsHolder[2], solutionArray)
				} else {
					solutionsHolder[3] = append(solutionsHolder[3], solutionArray)
				}

				for _, v := range open {
					tmp += len(v)
				}
				if tmp == 0 {
					return solutionsHolder
				}
				tmp = 0
			} else if counter > width {
				for _, v := range open {
					tmp += len(v)
				}
				if tmp == 0 {
					return solutionsHolder
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
