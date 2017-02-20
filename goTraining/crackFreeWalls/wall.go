/*
Made by Henning Vestergaard - Github: LimYoungChul
Copyright : Feel free to use this for any personal use

This is a solution to Project euler Problem 215. Read the problem description for clarity.

*/

package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	fmt.Println(wall(32, 10)) // give result to project euler problem 215. wall(32,10) = 806844323190414 combinations

	elapsed := time.Since(start)

	fmt.Print("Total runtime: ")
	fmt.Println(elapsed)
}

func wall(width, height int) uint64 {

	var total uint64
	var sum uint64
	combinations := findCombinations(width)
	var matrix [][]int
	//we need 2 vectors, one to hold the initial 1 array which will be updated and one to hold the result of a multiplication

	var vector []uint64       // our initial array which we fill with 1's
	var resultVector []uint64 // the result of multiplying  vector with matrix, need 2 vectors as the entire mul

	matrix = make([][]int, len(combinations), len(combinations))
	vector = make([]uint64, len(combinations))
	resultVector = make([]uint64, len(combinations))

	//populate our vectors and arrays
	for t := range matrix {
		matrix[t] = make([]int, len(combinations), len(combinations))
		vector[t] = 1
	}
	//create a spare matrix[i][j] where if j = 1 it can combine with i, otherwise j = 0
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
			if canCombine == true { // canCombine is true if no values are the same in two combinations
				matrix[n][n1] = 1
			}
		}
	}

	//Multiplying the matrix with the vector height-1 times. For each iteration the vector
	for n := 0; n < height-1; n++ {
		for i := range matrix {
			sum = 0
			for j := range matrix[i] {
				sum += uint64(matrix[i][j]) * vector[j]
			}
			resultVector[i] = sum
		}
		vector = append(vector[:0], resultVector...)
	}
	for _, v := range resultVector {
		total += v
	}
	return total
}

func findCombinations(width int) [][]int {
	var i int //represents the number of bricks currently placed
	var tmp int
	open := make([][]int, 0, 100)
	var solutionsHolder [][]int
	open = append(open, []int{3, 2}) //append the starting options.
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
				solutionArray := make([]int, len(tmpArray)-1) // we dont want the last value as it will be the width of the wall
				counter2 := 0
				for n := 0; n < len(tmpArray)-1; n++ {

					counter2 += tmpArray[n]
					solutionArray[n] = counter2
				}
				solutionsHolder = append(solutionsHolder, solutionArray)

				for _, v := range open { //this checks if there are no more values in open, if so we end our search
					tmp += len(v)
				}
				if tmp == 0 {
					return solutionsHolder
				}
				tmp = 0
			} else if counter > width {
				for _, v := range open { //we also check if open is empty here. I should probably put this in a function ;)
					tmp += len(v)
				}
				if tmp == 0 {
					return solutionsHolder
				}
				tmp = 0
			} else if counter < width {
				i++
				if len(open) <= i { //we don't want panic errors
					open = append(open, []int{3, 2})
				} else {
					open[i] = append(open[i], []int{3, 2}...) //we can't just append to the end if we are not at the end of a branch
				}
			}
		} else {
			i-- //traverse backwards
		}
	}
}
