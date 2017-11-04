package main

import (
	"fmt"
	"time"
)

// 1 10 09 04
// stringArray[0] = 2 "     XXXXX"
// 3 "XXXXXX...X"
// 4 "X..J.....X"
// 5 "X..JXXX.XX"
// 6 "XXGJG.G.GX"
// 7 " X.JX....X"
// 8 " X.MXXXXXX"
// 9 " X..X     "
// 10" XXXX     "

//X dwadwa
var X = byte('X')

//G dwadwa
var G = byte('G')

//J dwadwa
var J = byte('J')

//DOT dwadwa
var DOT = byte('.')

func main() {
	start := time.Now()
	var stringArray [9]string
	stringArray[0] = "     XXXXX"
	stringArray[1] = "XXXXXX...X"
	stringArray[2] = "X..J.....X"
	stringArray[3] = "X..JXXX.XX"
	stringArray[4] = "XXGJG.G.GX"
	stringArray[5] = " X.JX....X"
	stringArray[6] = " X..XXXXXX"
	stringArray[7] = " X..X     "
	stringArray[8] = " XXXX     "

	matrix := make([][]byte, 0)
	matrixCopy := matrix
	matrixPointer := &matrix

	for i := range stringArray {
		tmp := []byte(stringArray[i])
		matrix = append(matrix[:], tmp)
	}
	//i just make x = y and y = x, so we can write x,y instead of y,x.
	transpose(matrix)
	fmt.Println(matrix[3][6])
	fmt.Println(matrix[4][6])
	fmt.Println(matrix[2][6])
	fmt.Println(matrix[3][7])
	fmt.Println(matrix[3][5])
	fmt.Println(possibleTurns(3, 6, matrix))
	elapsed := time.Since(start)
	fmt.Println("Runtime: ", elapsed)
}

func transpose(a [][]byte) {
	n := len(a)
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			b[i][j] = a[j][i]
		}
	}
	copy(a, b)
}

func step(x, y int, matrix [][]byte, state []string) {
	moves := possibleTurns(x, y, matrix)
	if isCompleted(matrix) == true {
		return
	}

}

func isCompleted(stepCounter, maxSteps int, matrix [][]byte) bool {
	if matrix[2][4] == J && matrix[4][4] == J && matrix[6][4] == J && matrix[8][4] == J {
		return true
	}
	if stepCounter == maxSteps {
		return true
	}

	return false
}

func possibleTurns(x, y int, matrix [][]byte) []string {
	var posTurns = make([]string, 0, 4)
	leftMove := matrix[x-1][y]
	rightMove := matrix[x+1][y]
	upMove := matrix[x][y-1]
	downMove := matrix[x][y+1]
	doubleLeftMove := matrix[x-2][y]
	doubleRightMove := matrix[x+2][y]
	doubleUpMove := matrix[x][y-2]
	doubleDownMove := matrix[x][y+2]

	if leftMove != X {
		if leftMove == DOT || leftMove == G {
			posTurns = append(posTurns, "L")
		} else if leftMove == J {
			if doubleLeftMove != J && doubleLeftMove != X {
				posTurns = append(posTurns, "LJ")
			}
		}
	}
	if rightMove != X {
		if rightMove == DOT || rightMove == G {
			posTurns = append(posTurns, "R")
		} else if rightMove == J {
			if doubleRightMove != J && doubleRightMove != X {
				posTurns = append(posTurns, "RJ")
			}
		}
	}
	if upMove != X {
		if upMove == DOT || upMove == G {
			posTurns = append(posTurns, "U")
		} else if upMove == J {
			if doubleUpMove != J && doubleUpMove != X {
				posTurns = append(posTurns, "UJ")
			}
		}
	}
	if downMove != X {
		if downMove == DOT || downMove == G {
			posTurns = append(posTurns, "D")
		} else if downMove == J {
			if doubleDownMove != J && doubleDownMove != X {
				posTurns = append(posTurns, "DJ")
			}
		}
	}

	return posTurns
}
