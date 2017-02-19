package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	fmt.Println(wall(10, 3))

	elapsed := time.Since(start)

	fmt.Print("Total runtime: ")
	fmt.Println(elapsed)
}

func wall(width, height int) int64 {

	var total int64

	combinations := findCombinations(width)

	//combiPointer := &combinations

	var matrix [][]int

	var vector []int
	var resultVector []int

	matrix = make([][]int, len(combinations), len(combinations))

	vector = make([]int, len(combinations))
	resultVector = make([]int, len(combinations))

	for t := range matrix {
		matrix[t] = make([]int, len(combinations), len(combinations))
		vector[t] = 1

	}
	for n, a := range combinations {
		var canCombine bool
		for n1, a1 := range combinations {
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
				matrix[n][n1] = 1
			}
		}
	}

	tmp := powerMatrix(matrix, height, len(combinations))

	for i := range tmp {
		for j := range tmp[i] {
			resultVector[i] += tmp[i][j] * vector[j]
		}
	}

	for _, value := range resultVector {
		total += int64(value)
	}

	return total
}

func multiplyMatrices(a, b [][]int, length int) [][]int {
	var resultMatrix [][]int
	resultMatrix = make([][]int, length, length)
	var sum int
	for i := range b {
		resultMatrix[i] = make([]int, length, length)
		for j := range b {
			sum = 0
			for k := range b {
				sum += a[i][k] * b[k][j]
			}
			resultMatrix[i][j] = sum
		}
	}
	return resultMatrix
}

func powerMatrix(a [][]int, power int, length int) [][]int {
	result := a
	for n := 2; n < power; n++ {
		result = multiplyMatrices(a, result, length)
	}
	return result
}

func findCombinations(width int) [][]int {
	var i int
	var tmp int
	open := make([][]int, 0, 100)
	var solutionsHolder [][]int
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

				solutionsHolder = append(solutionsHolder, solutionArray)

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
