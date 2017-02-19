package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(wall(22, 10))
}

func wall(width, height int) int64 {

	var total int64

	combinations := findCombinations(width)
	combiPointer := &combinations

	var tmp int

	m := make(map[string][][]int)

	var comboArray [][][]int

	for i := 0; i < 4; i++ {

		for index, a := range combinations[i] {
			if i%2 == 0 {
				buildMap(a, combiPointer[i+1], tmp+index, &comboArray, &m)
			} else {
				buildMap(a, combiPointer[i-1], tmp+index, &comboArray, &m)
			}

		}
		tmp += len(combinations[i])
	}
	// fmt.Println(combinations)
	// fmt.Println(m)
	// fmt.Println("----------------------------------------------------------------------")
	//tmp = 0

	for i := range m {

		buildWall(m[i], &m, height, 0, &total)

	}
	// fmt.Print("HERE COMES I: ")
	// fmt.Println(i)
	//tmp += len(combinations[i])
	return total
}

func buildWall(a [][]int, m *map[string][][]int, height, level int, total *int64) {
	level++

	if level >= height {
		*total++
		return
	}

	for _, v := range a {
		vs := stringIt(v)
		buildWall((*m)[vs], m, height, level, total)
	}
}

func buildMap(a []int, combi [][]int, index int, combiArray *[][][]int, m *map[string][][]int) {
	//var c1 int

	result := stringIt(a)

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
		if canCombine == true {

			// if index >= len(*combiArray) {
			// 	(*combiArray) = append((*combiArray), [][]int{})
			// }
			// (*combiArray)[index] = append((*combiArray)[index], a1)

			(*m)[result] = append((*m)[result], a1)

		}
	}
}

func stringIt(a []int) string {
	valuesText := []string{}
	for i := range a {
		number := a[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	result := strings.Join(valuesText, " ")
	return result
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
