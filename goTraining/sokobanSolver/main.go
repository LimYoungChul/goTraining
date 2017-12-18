package main

import (
	"fmt"
	"strconv"
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

var goalArray []Coordinate
var emptyMatrix [][]byte

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

func setGoals(matrix [][]byte) {

	for i := range matrix {
		for i2 := range matrix[i] {
			if matrix[i][i2] == G {
				goalArray = append(goalArray, newCoordinate(i, i2, nil))
			}
		}
	}
}

//Box represents a box in the matrix
type Box struct {
	coordinate     Coordinate
	possiblePushes []string
	inGoal         bool
	stuck          bool
}

var closedNodeList []matrixHolder

func (b Box) push(s string) Box {
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
	return b
}

func identifier(node []Box) string {
	var s string
	for i := range node {
		s += strconv.Itoa(node[i].coordinate.X)
		s += strconv.Itoa(node[i].coordinate.Y)
	}
	return s
}

func createBox(x, y int) Box {
	box := Box{}
	box.coordinate = newCoordinate(x, y, nil)

	return box
}

type matrixHolder struct {
	boxes          []Box
	parentMatrix   *matrixHolder
	robotX, robotY int
	id             string
}

func newMatrixHolder(x, y int, b []Box, parent *matrixHolder) matrixHolder {
	m := matrixHolder{}
	m.parentMatrix = parent
	m.robotX = x
	m.robotY = y
	m.boxes = m.updateBoxPushes(b)
	m.id = identifier(m.boxes)

	return m
}

//var finalPath []Coordinate

func main() {
	start := time.Now()
	stringArray := []string{
		"     XXXXX",
		"XXXXXX...X",
		"X..J.....X",
		"X..JXXX.XX",
		"XXGJG.G.GX",
		" X.JX....X",
		" X..XXXXXX",
		" X..X     ",
		" XXXX     "}
	stingArray := []string{
		"XXXXXXXXXX",
		"XXXXXXXXXX",
		"XXG....GXX",
		"XX..J...XX",
		"XX..JJ..XX",
		"XX..J...XX",
		"XXG....GXX",
		"XXXXXXXXXX",
		"XXXXXXXXXX"}

	emptyArray := []string{
		"XXXXXXXXXX",
		"XXXXXXXXXX",
		"XXG....GXX",
		"XX......XX",
		"XX......XX",
		"XX......XX",
		"XXG....GXX",
		"XXXXXXXXXX",
		"XXXXXXXXXX"}

	for i := range emptyArray {
		fmt.Println(emptyArray[i])
		tmp := []byte(emptyArray[i])
		emptyMatrix = append(emptyMatrix[:], tmp)
	}

	originalMatrix := make([][]byte, 0)
	// matrixCopy := matrix
	// matrixPointer := &matrix
	fmt.Println(stringArray)
	for i := range stingArray {
		fmt.Println(stingArray[i])
		tmp := []byte(stingArray[i])
		originalMatrix = append(originalMatrix[:], tmp)
	}
	setGoals(originalMatrix)
	fmt.Println(goalArray)
	b := createBoxes(originalMatrix)
	fmt.Println(b)
	m := newMatrixHolder(6, 3, b, nil)
	for i := range m.boxes {
		fmt.Println(m.boxes[i].possiblePushes)
	}
	fmt.Println("---- POSSIBLEPUSHES UPDATED ----")
	for i := range m.boxes {
		fmt.Println(m.boxes[i].possiblePushes)
	}
	i := 0
	var stateArray [][]matrixHolder
	var finalMatrix matrixHolder
	var goalCounter int
	stateArray = append(stateArray, []matrixHolder{m})

Loop:
	for {
		stateArray = append(stateArray, []matrixHolder{})
		if len(stateArray[i]) == 0 {
			fmt.Println("ERROR OR COULD NOT SOLVE")
			break
		}
	Loop2:
		for n := range stateArray[i] {

			for x := range closedNodeList {
				if stateArray[i][n].id == closedNodeList[x].id {
					continue Loop2
				}
			}
			// for _, v := range stateArray[i][n].matrix {
			// 	s := string(v)
			// 	fmt.Println(s)
			// }
			goalCounter = 0
			for n2 := range stateArray[i][n].boxes {
				if stateArray[i][n].boxes[n2].inGoal {
					goalCounter++
				}
				if goalCounter == 4 {
					finalMatrix = stateArray[i][n]
					break Loop
				}
			}
			closedNodeList = append(closedNodeList, stateArray[i][n])
			// for n2, v := range stateArray[i][n].boxes {
			// 	fmt.Println("State:", i, "Box number:", n2, " pushes:", v.possiblePushes)
			// }
			stateArray[i+1] = append(stateArray[i+1], pushesToMatrix(stateArray[i][n])...)

		}
		i++
		fmt.Println("-------", i, "-------")

	}

	fmt.Println(finalMatrix)

	elapsed := time.Since(start)
	fmt.Println("Runtime: ", elapsed)
}

func pushesToMatrix(m matrixHolder) []matrixHolder {
	var matrixHolderArray []matrixHolder
	tempBoxes := make([]Box, len(m.boxes))
	var x, y int
	for i, v := range m.boxes {
		for _, v2 := range v.possiblePushes {
			copy(tempBoxes, m.boxes)
			tempBoxes[i] = tempBoxes[i].push(v2)
			x = v.coordinate.X
			y = v.coordinate.Y
			switch v2 {
			case "left":
				x++
			case "right":
				x--
			case "up":
				y++
			case "down":
				y--
			}
			matrixHolderArray = append(matrixHolderArray, newMatrixHolder(x, y, tempBoxes, &m))
		}
	}
	return matrixHolderArray
}

func updateMatrix(boxes []Box) [][]byte {
	matrix := make([][]byte, len(emptyMatrix))

	for i := range emptyMatrix {
		matrix[i] = make([]byte, len(emptyMatrix[i]))
		copy(matrix[i], emptyMatrix[i])
	}

	for i := range matrix {
		for i2 := range matrix[i] {
			for i3 := range boxes {
				if i == boxes[i3].coordinate.X && i2 == boxes[i3].coordinate.Y {
					matrix[i][i2] = J
				}
			}
		}
	}
	return matrix
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
func (m matrixHolder) updateBoxPushes(boxes []Box) []Box {
	newBoxes := make([]Box, len(boxes))
	copy(newBoxes, boxes)
	matrix := updateMatrix(newBoxes)
	x := m.robotX
	y := m.robotY
	origin := newCoordinate(x, y, nil)

	left := X
	right := X
	up := X
	down := X
	for i := range newBoxes {
		if newBoxes[i].coordinate.Y > 0 {
			left = matrix[newBoxes[i].coordinate.X][newBoxes[i].coordinate.Y-1]
		}
		if newBoxes[i].coordinate.Y < len(matrix[newBoxes[i].coordinate.X]) {
			right = matrix[newBoxes[i].coordinate.X][newBoxes[i].coordinate.Y+1]
		}
		if newBoxes[i].coordinate.X > 0 {
			up = matrix[newBoxes[i].coordinate.X-1][newBoxes[i].coordinate.Y]
		}
		if newBoxes[i].coordinate.Y < len(matrix[newBoxes[i].coordinate.X]) {
			down = matrix[newBoxes[i].coordinate.X+1][newBoxes[i].coordinate.Y]
		}
		newBoxes[i].possiblePushes = []string{}

		for i2 := range goalArray {
			if newBoxes[i].coordinate.X == goalArray[i2].X && newBoxes[i].coordinate.Y == goalArray[i2].Y {
				newBoxes[i].inGoal = true
			} else {
				newBoxes[i].inGoal = false
			}
		}
		if (left == X || right == X) && (up == X || down == X) {
			newBoxes[i].stuck = true
			if newBoxes[i].inGoal {
				continue
			} else {
				for n := range newBoxes {
					newBoxes[n].possiblePushes = []string{}
				}
				break
			}
		}

		if left == DOT || left == G { // If i can stand to the left of the box.
			if right == DOT || right == G { // And i can push it to the right of the box.
				test, _ := aSTAR(origin, newCoordinate(newBoxes[i].coordinate.X, newBoxes[i].coordinate.Y-1, nil), matrix) // and i can move to the left of the box from where i stand.
				if test {
					//newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, newCoordinate(newBoxes[i].coordinate.X, newBoxes[i].coordinate.Y-1, nil)) //then add that option to the box
					newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, "right")
				}
			}
		}
		if right == DOT || right == G {
			if left == DOT || left == G {
				test, _ := aSTAR(origin, newCoordinate(newBoxes[i].coordinate.X, newBoxes[i].coordinate.Y+1, nil), matrix)
				if test {
					//newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, newCoordinate(newBoxes[i].coordinate.X, newBoxes[i].coordinate.Y+1, nil))
					newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, "left")
				}
			}
		}
		if up == DOT || up == G {
			if down == DOT || down == G {
				test, _ := aSTAR(origin, newCoordinate(newBoxes[i].coordinate.X-1, newBoxes[i].coordinate.Y, nil), matrix)
				if test {
					//newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, newCoordinate(newBoxes[i].coordinate.X-1, newBoxes[i].coordinate.Y, nil))
					newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, "down")
				}
			}
		}
		if down == DOT || down == G {
			if up == DOT || up == G {
				test, _ := aSTAR(origin, newCoordinate(newBoxes[i].coordinate.X+1, newBoxes[i].coordinate.Y, nil), matrix)
				if test {
					//newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, newCoordinate(newBoxes[i].coordinate.X+1, newBoxes[i].coordinate.Y, nil))
					newBoxes[i].possiblePushes = append(newBoxes[i].possiblePushes, "up")
				}
			}
		}
	}
	return newBoxes
}
