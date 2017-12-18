package main

import (
	"fmt"
	"time"
)

// 1 10 09 04
// 2 "     XXXXX"
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

// var boxes []Box
var goalArray []Coordinate

//Coordinate storage
type Coordinate struct {
	X, Y      int
	heuristic int
	Parent    *Coordinate
}

func newCoordinate(x, y int, parent *Coordinate) Coordinate {
	coordinate := Coordinate{}
	coordinate.X = x
	coordinate.Y = y
	coordinate.Parent = parent
	coordinate.heuristic = distance(x, y, destination)

	return coordinate
}

//Box represents a box in the matrix
type Box struct {
	X, Y                                  int
	possiblePushes                        []string
	inGoal                                bool
	stuck                                 bool
	leftPath, rightPath, downPath, upPath []Coordinate
}

func (b Box) push(ss []string, matrix [][]byte) {
	for _, s := range ss {
		switch s {
		case "left":
			b.Y--
		case "right":
			b.Y++
		case "up":
			b.X--
		case "down":
			b.X++
		}
	}
}

func createBox(x, y int) Box {
	box := Box{}
	box.X = x
	box.Y = y

	return box
}

var finalPath [][]Coordinate

func main() {
	start := time.Now()

	// stringArray := []string{
	// 	"XXXXXXXXXX",
	// 	"XXXXXX...X",
	// 	"X..J.....X",
	// 	"X..JXXX.XX",
	// 	"XXGJG.G.GX",
	// 	"XX.JX....X",
	// 	"XX..XXXXXX",
	// 	"XX..XXXXXX",
	// 	"XXXXXXXXXX"}
	stingArray := []string{
		"XXXXXXXXXX",
		"XXXXXXXXXX",
		"XXG.J..GXX",
		"XX......XX",
		"XX..JJ..XX",
		"XX......XX",
		"XXG.J..GXX",
		"XXXXXXXXXX",
		"XXXXXXXXXX"}

	originalMatrix := make([][]byte, 0)
	// matrixCopy := matrix
	// matrixPointer := &matrix

	for i := range stingArray {
		fmt.Println(stingArray[i])
		tmp := []byte(stingArray[i])
		originalMatrix = append(originalMatrix[:], tmp)
	}

	setGoals(originalMatrix)
	fmt.Println(len(originalMatrix))
	fmt.Println(len(originalMatrix[0]))
	startingPoint := newCoordinate(3, 2, nil)
	endPoint := newCoordinate(3, 5, nil)
	fmt.Println(aSTAR(startingPoint, endPoint, originalMatrix))
	fmt.Println(step(6, 3, 0, 15, originalMatrix))
	//
	// for _, v := range finalPath {
	// 	for _, v2 := range v {
	// 		lard := originalMatrix[v2.X][v2.Y]
	// 		originalMatrix[v2.X][v2.Y] = byte('M')
	// 		for i3 := range originalMatrix {
	// 			s := string(originalMatrix[i3])
	// 			fmt.Println(s)
	// 		}
	// 		originalMatrix[v2.X][v2.Y] = lard
	// 		time.Sleep(1000)
	// 		fmt.Println("         ")
	// 	}
	//
	// }

	//fmt.Sprintf("matrix value %s", string(matrix[6][3])))
	// fmt.Println(fmt.Sprintf("matrix up %s", string(matrix[5][3])))
	// fmt.Println(fmt.Sprintf("matrix down %s", string(matrix[7][3])))
	// fmt.Println(fmt.Sprintf("matrix left %s", string(matrix[6][2])))
	// fmt.Println(fmt.Sprintf("matrix right %s", string(matrix[6][4])))

	elapsed := time.Since(start)
	fmt.Println("Runtime: ", elapsed)
}

func setGoals(matrix [][]byte) {

	for i := range matrix {
		for i2 := range matrix[i] {
			if matrix[i][i2] == G {
				goalArray = append(goalArray, newCoordinate(i, i2, nil))
			}
		}
	}
}

func step(x, y, counter, maxCounter int, matrix [][]byte) bool {
	boxes := createBoxes(matrix)
	updateBoxPushes(x, y, boxes, matrix)
	tmpMatrix := make([][]byte, len(matrix))
	var goalCounter int
	var success bool
	for _, b2 := range boxes {
		if b2.stuck && !b2.inGoal {
			return false
		}
		if b2.inGoal {
			goalCounter++
			if goalCounter >= 4 {
				fmt.Println(matrix)
				finalPath = make([][]Coordinate, counter)
				return true
			}
		}
	}
	if counter >= maxCounter {
		return false
	}

	for i := range boxes {
		for n := range boxes[i].possiblePushes {
			for c := range matrix {
				tmpMatrix[c] = make([]byte, len(matrix[c]))
				copy(tmpMatrix[c], matrix[c])
			}
			switch boxes[i].possiblePushes[n] {
			case "left":
				tmpMatrix[boxes[i].X][boxes[i].Y] = DOT
				tmpMatrix[boxes[i].X][boxes[i].Y-1] = J
				success = step(boxes[i].X, boxes[i].Y+1, counter+1, maxCounter, tmpMatrix)
			case "right":
				tmpMatrix[boxes[i].X][boxes[i].Y] = DOT
				tmpMatrix[boxes[i].X][boxes[i].Y+1] = J
				success = step(boxes[i].X, boxes[i].Y-1, counter+1, maxCounter, tmpMatrix)
			case "down":
				tmpMatrix[boxes[i].X][boxes[i].Y] = DOT
				tmpMatrix[boxes[i].X+1][boxes[i].Y] = J
				success = step(boxes[i].X-1, boxes[i].Y, counter+1, maxCounter, tmpMatrix)
			case "up":
				tmpMatrix[boxes[i].X][boxes[i].Y] = DOT
				tmpMatrix[boxes[i].X-1][boxes[i].Y] = J
				success = step(boxes[i].X+1, boxes[i].Y, counter+1, maxCounter, tmpMatrix)
			}
			if success {
				switch boxes[i].possiblePushes[n] {
				case "left":
					finalPath[counter] = append(finalPath[counter], boxes[i].leftPath...)
				case "right":
					finalPath[counter] = append(finalPath[counter], boxes[i].rightPath...)
				case "down":
					finalPath[counter] = append(finalPath[counter], boxes[i].downPath...)
				case "up":
					finalPath[counter] = append(finalPath[counter], boxes[i].upPath...)
				}
				return true
			}
		}
	}
	return false
}

// func solveSokoban(x, y int, matrix [][]byte) {
// 	maxCounter := 5
// 	counter := 0
// 	startBoxes := createBoxes(matrix)
// 	var stateArray [][]Box
// 	solved := false
//
// 	for _, v := range startBoxes {
//
// 		stateArray[0] = append(stateArray[0], v.possiblePushes)
//
// 	}
//
// 	for !solved {
//
// 		for counter < maxCounter {
//
// 		}
// 	}
//
// }

func createBoxes(matrix [][]byte) []Box {
	var boxes []Box
	for i, v := range matrix {
		for i2, v2 := range v {
			if v2 == J {
				boxes = append(boxes, createBox(i, i2))
			}
		}
	}
	return boxes
}
func updateBoxPushes(x, y int, boxes []Box, matrix [][]byte) {

	origin := newCoordinate(x, y, nil)

	for i := range boxes {
		left := matrix[boxes[i].X][boxes[i].Y-1]
		right := matrix[boxes[i].X][boxes[i].Y+1]
		up := matrix[boxes[i].X-1][boxes[i].Y]
		down := matrix[boxes[i].X+1][boxes[i].Y]

		for i2 := range goalArray {
			if boxes[i].X == goalArray[i2].X && boxes[i].Y == goalArray[i2].Y {
				boxes[i].inGoal = true
			} else {
				boxes[i].inGoal = false
			}
		}
		if (up == X || down == X) && (left == X || right == X) {
			boxes[i].stuck = true
			continue
		}
		if left == DOT || left == G { // If i can stand to the left of the box.
			if right == DOT || right == G { // And i can push it to the right of the box.
				test, _ := aSTAR(origin, newCoordinate(boxes[i].X, boxes[i].Y-1, nil), matrix) // and i can move to the left of the box from where i stand.
				if test {
					boxes[i].possiblePushes = append(boxes[i].possiblePushes, "right") //then add that option to the box
				}
			}
		}
		if right == DOT || right == G {
			if left == DOT || left == G {
				test, _ := aSTAR(origin, newCoordinate(boxes[i].X, boxes[i].Y+1, nil), matrix)
				if test {
					boxes[i].possiblePushes = append(boxes[i].possiblePushes, "left")
				}
			}
		}
		if up == DOT || up == G {
			if down == DOT || down == G {
				test, _ := aSTAR(origin, newCoordinate(boxes[i].X-1, boxes[i].Y, nil), matrix)
				if test {
					boxes[i].possiblePushes = append(boxes[i].possiblePushes, "down")
				}
			}
		}
		if down == DOT || down == G {
			if up == DOT || up == G {
				test, _ := aSTAR(origin, newCoordinate(boxes[i].X+1, boxes[i].Y, nil), matrix)
				if test {
					boxes[i].possiblePushes = append(boxes[i].possiblePushes, "up")
				}
			}
		}
	}
}
