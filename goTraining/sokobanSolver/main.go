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
	coordinate     Coordinate
	possiblePushes []Coordinate
	inGoal         bool
	stuck          bool
	path           []Coordinate
}

func (b Box) push(ss []string, matrix [][]byte) {
	for _, s := range ss {
		switch s {
		case "left":
			b.coordinate.Y--
		case "right":
			b.coordinate.Y++
		case "up":
			b.coordinate.X--
		case "down":
			b.coordinate.X++
		}
	}
}

func createBox(x, y int) Box {
	box := Box{}
	box.coordinate = newCoordinate(x, y, nil)

	return box
}

//var finalPath []Coordinate

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

	var stingArray [9]string
	stingArray[0] = "XXXXXXXXXX"
	stingArray[1] = "XXXXXXXXXX"
	stingArray[2] = "XXG.J..GXX"
	stingArray[3] = "XX......XX"
	stingArray[4] = "XX..JJ..XX"
	stingArray[5] = "XX......XX"
	stingArray[6] = "XXG.J..GXX"
	stingArray[7] = "XXXXXXXXXX"
	stingArray[8] = "XXXXXXXXXX"

	originalMatrix := make([][]byte, 0)
	// matrixCopy := matrix
	// matrixPointer := &matrix

	for i := range stingArray {
		tmp := []byte(stingArray[i])
		originalMatrix = append(originalMatrix[:], tmp)
	}
	//fmt.Println(step(6, 3, 0, 15, originalMatrix))

	//fmt.Sprintf("matrix value %s", string(matrix[6][3])))
	// fmt.Println(fmt.Sprintf("matrix up %s", string(matrix[5][3])))
	// fmt.Println(fmt.Sprintf("matrix down %s", string(matrix[7][3])))
	// fmt.Println(fmt.Sprintf("matrix left %s", string(matrix[6][2])))
	// fmt.Println(fmt.Sprintf("matrix right %s", string(matrix[6][4])))

	elapsed := time.Since(start)
	fmt.Println("Runtime: ", elapsed)
}

func solveSokoban(x, y int, matrix [][]byte) {
	maxCounter := 5
	counter := 0
	startBoxes := createBoxes(matrix)
	var stateArray [][]Coordinate
	solved := false

	for _, v := range startBoxes {

		stateArray[0] = append(stateArray[0], v.possiblePushes...)

	}

	for !solved {

		for counter < maxCounter {

		}
	}

}

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
		left := matrix[boxes[i].coordinate.X][boxes[i].coordinate.Y-1]
		right := matrix[boxes[i].coordinate.X][boxes[i].coordinate.Y+1]
		up := matrix[boxes[i].coordinate.X-1][boxes[i].coordinate.Y]
		down := matrix[boxes[i].coordinate.X+1][boxes[i].coordinate.Y]

		if (up == X || down == X) && (left == X || right == X) {
			boxes[i].stuck = true
		}
		// if (boxes[i].coordinate.X == 4) && (boxes[i].coordinate.Y == 2 || boxes[i].coordinate.Y == 4 || boxes[i].coordinate.Y == 6 || boxes[i].coordinate.Y == 8) {
		// 	boxes[i].inGoal = true
		// } else {
		// 	boxes[i].inGoal = false
		// }
		if (boxes[i].coordinate.X == 2 && boxes[i].coordinate.Y == 2) || (boxes[i].coordinate.X == 2 && boxes[i].coordinate.Y == 7) || (boxes[i].coordinate.X == 6 && boxes[i].coordinate.Y == 2) || (boxes[i].coordinate.X == 6 && boxes[i].coordinate.Y == 7) {
			boxes[i].inGoal = true
		} else {
			boxes[i].inGoal = false
		}
		if boxes[i].stuck == false {
			if left == DOT || left == G { // If i can stand to the left of the box.
				if right == DOT || right == G { // And i can push it to the right of the box.
					test, path := aSTAR(origin, newCoordinate(boxes[i].coordinate.X, boxes[i].coordinate.Y-1, nil), matrix) // and i can move to the left of the box from where i stand.
					if test {
						boxes[i].possiblePushes = append(boxes[i].possiblePushes, newCoordinate(boxes[i].coordinate.X, boxes[i].coordinate.Y-1, nil)) //then add that option to the box
						boxes[i].path = path
					}
				}
			}
			if right == DOT || right == G {
				if left == DOT || left == G {
					test, path := aSTAR(origin, newCoordinate(boxes[i].coordinate.X, boxes[i].coordinate.Y+1, nil), matrix)
					if test {
						boxes[i].possiblePushes = append(boxes[i].possiblePushes, newCoordinate(boxes[i].coordinate.X, boxes[i].coordinate.Y+1, nil))
						boxes[i].path = path
					}
				}
			}
			if up == DOT || up == G {
				if down == DOT || down == G {
					test, path := aSTAR(origin, newCoordinate(boxes[i].coordinate.X-1, boxes[i].coordinate.Y, nil), matrix)
					if test {
						boxes[i].possiblePushes = append(boxes[i].possiblePushes, newCoordinate(boxes[i].coordinate.X-1, boxes[i].coordinate.Y, nil))
						boxes[i].path = path
					}
				}
			}
			if down == DOT || down == G {
				if up == DOT || up == G {
					test, path := aSTAR(origin, newCoordinate(boxes[i].coordinate.X+1, boxes[i].coordinate.Y, nil), matrix)
					if test {
						boxes[i].possiblePushes = append(boxes[i].possiblePushes, newCoordinate(boxes[i].coordinate.X+1, boxes[i].coordinate.Y, nil))
						boxes[i].path = path
					}
				}
			}
		}
	}
}
