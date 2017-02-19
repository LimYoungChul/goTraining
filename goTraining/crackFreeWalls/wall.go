/*
Made by Henning Vestergaard - Github: LimYoungChul
Copyright : Feel free to use this for any personal use

This is a solution to Project euler Problem 215. Read the problem description for clarity.

This part explains how you get from having an array with all possible combinations for a wall row with width X to combining them into a wall with Y rows.
How to get all possible combinations is explained further below down at the function 'findCombinations'. *Not added yet

The explanation might be a little over the top, but i personally had a hard time understanding how matrix multiplication can be used to find possible solutions.
You can also just read the code and understand what's going on if you're a fookin legend ;)

Once we have all possible combinations as an array of values where each value represents the position of a crack (except the start and the end of the wall) we can use it to see which combinations can be put on top of otherwise

Say we are creating a wall of 9 width and 3 height so wall(9,3).

Our combination array would look like this, given a width of 9

combinations = [[3,6][3,5,7][2,5,7][2,4,7][2,4,6]] - Notice no 9's or 0's in our array, these values are irrelevant so we don't add them.

We can then conclude that if any entry in one of our arrays matches a value in another that it won't be able to be put on top of another.
Why? Because the values represents the positions of cracks, and we don't want any cracks on top of eachother.

We use this to create a matrix of NxN size, where N is the number of combinations. This matrix shows which combinations can go with eachother.
0 will represent the combinations that are not possbile and 1 will represent the ones that are possible

[3,6]		-> [0,0,1,1,0] - cause [3,6] can combine with [2,5,7] and [2,4,7]
[3,5,7] -> [0,0,0,0,1] - cause [3,5,7] can combine with [2,4,6]
[2,5,7] -> [1,0,0,0,0] - cause [2,5,7] can combine with [3,6]
[2,4,7] -> [1,0,0,0,0] - cause [2,4,7] can combine with [3,6]
[2,4,6] -> [0,1,0,0,0] - cause [2,4,6] cam combine with [3,5,7]

We then have a sparse matrix that represents which combinations of bricks can go on top of other combinations of bricks.

We then create a vector of 1's that represents the possible number of combinations for a wall of height 1, which in this case is [1,1,1,1,1]

Vector = [1,1,1,1,1]  - Because for a wall of height 1 we can put in our 5 initial combinations.

We then multiply this vector with the matrix, and this will represent the number of combinations of a wall of height 2 in a new vector

				 [0,0,1,1,0]		[1]		[0*1 + 0*1 + 1*1 + 1*1 + 0*1]		[2]	- Because there are 2 walls we can put on top, thus expanding our possibilties by 1
				 [0,0,0,0,1]	  [1]		[0*1 + 0*1 + 0*1 + 0*1 + 1*1]		[1] - Because there is only 1 wall we can put on top
Vector = [1,0,0,0,0] *  [1] =	[1*1 + 0*1 + 0*1 + 0*1 + 0*1] = [1]
				 [1,0,0,0,0]		[1]		[1*1 + 0*1 + 0*1 + 0*1 + 0*1]		[1]
				 [0,1,0,0,0]		[1]		[0*1 + 1*1 + 0*1 + 0*1 + 0*1]		[1]

If we take the sum of the resulting vector we find the possible solutions for a wall of height 2
Vector = [2,1,1,1,1] = 2+1+1+1+1 = 6 - Cause there was only 1 additional combination.

If we want to find the possible combinations of a wall with the height 3 we simply take the resulting vector and multiply it with our matrix again

				 [0,0,1,1,0]		[2]		[0*2 + 0*1 + 1*1 + 1*1 + 0*1]		[2]
				 [0,0,0,0,1]	  [1]		[0*2 + 0*1 + 0*1 + 0*1 + 1*1]		[1]
Vector = [1,0,0,0,0] *  [1] =	[1*2 + 0*1 + 0*1 + 0*1 + 0*1] = [2] - Now we get more combinations cause we have 2 possible solutions in our first matrix row, and this multiplies the number of entries that can go on top of this entry
				 [1,0,0,0,0]		[1]		[1*2 + 0*1 + 0*1 + 0*1 + 0*1]		[2]
				 [0,1,0,0,0]		[1]		[0*2 + 1*1 + 0*1 + 0*1 + 0*1]		[1]

Summing our vector up we see that wall(9,3) = 8
[2,1,2,2,1] = 2+1+2+2+1 = 8

Let's try to work the logic out differently just for clarity.

So with a wall of height 2 we get the following wall combinations

Level:  1               2
		  [3,6]   + [2,5,7] || [2,4,7] 			 = [3,6][2,5,7] && [3,6][2,4,7] = 2 combinations	- so by adding 1 height to the wall we got 1 more combination
		  [3,5,7] + [2,4,6] 								 = [3,5,7][2,4,6] 						  = 1 combination - we didnt get any more combinations than when our height was 1
		  [2,5,7] + [3,6] 									 = [2,5,7][3,6]     						= 1 combination
		  [2,4,7] + [3,6] 									 = [2,4,7][3,6]     						= 1 combination
		  [2,4,6] + [3,5,7] 								 = [2,4,6][3,5,7] 							= 1 combination

Now we add another level to our wall and we see what happens.

Level:  1     2     				            3
		  [3,6][2,5,7] && [3,6][2,4,7] + [3,6] 							= [3,6][2,5,7][3,6] && [3,6][2,4,7][3,6] 		 = 2 combinations - We could only add 1 thing so we didnt get any more combinations
		  [3,5,7][2,4,6] 							 + [3,5,7]						= [3,5,7][2,4,6][3,5,7] 								  	 = 1 combination - Here we are stuck in a loop of adding the same walls on top eachother over and over so we never increase in combinations
		  [2,5,7][3,6] 								 + [2,5,7] || [2,4,7] = [2,5,7][3,6][2,5,7] && [2,5,7][3,6][2,4,7] = 2 combinations - Since the combination we put on top of our original combination has 2 combinations, the total combinations is now also 2
		  [2,4,7][3,6] 								 + [2,4,7] || [2,5,7]	= [2,4,7][3,6][2,4,7] && [2,4,7][3,6][2,5,7] = 2 combinations - Same as above
		  [2,4,6][3,5,7] 							 + [2,4,6]						= [2,4,6][3,5,7][2,4,6]											 = 1 combination - Again, we are stuck in a loop here so combination never increases.


			We add up the number of combinations: 2+1+2+2+1 = 8
			And we see that there are 8 combinations for a wall with a width of 9 and a height of 3.

The matrix-vector multiplication represents the same as above, yet much more elegantly so it can be processed relatively quickly.
The vector keeps track of every point where multiple possiblities are possible and sum them up in one value.
Since an extra possibility anywhere in the height of the wall affects the number of possibilities when adding a new layer it is simply multiplied with the initial matrix.

We could even go on and find wall(9,4)

					[0,0,1,1,0]		[2]		[0*2 + 0*1 + 1*2 + 1*2 + 0*1]		[4]
					[0,0,0,0,1]	  [1]		[0*2 + 0*1 + 0*2 + 0*2 + 1*1]		[1]
Vector = 	[1,0,0,0,0] * [2] =	[1*2 + 0*1 + 0*2 + 0*2 + 0*1] = [2]
					[1,0,0,0,0]		[2]		[1*2 + 0*1 + 0*2 + 0*2 + 0*1]		[2]
					[0,1,0,0,0]		[1]		[0*2 + 1*1 + 0*2 + 0*2 + 0*1]		[1]

Summing our vector we see that wall(9,4) = 10
[4,1,2,2,1] = 4+1+2+2+1 = 10

wall(9,5) = [4,1,4,4,1] = 4+1+4+4+1 = 14
wall(9,6) = [8,1,4,4,1] = 8+1+4+4+1 = 18
etc...

This generalized solution can be computed with any width or height. Be vary though it gets exponentially harder for the computer to calculate it
wall(40,40) takes about ~5 minutes to compute, and going much higher than that will take a long time and eventually return a result too big for an uint64 to hold so you would need to modify the code a bit
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
	var i int
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
