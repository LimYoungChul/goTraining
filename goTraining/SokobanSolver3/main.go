package main

import "fmt"

type coordinate struct {
	X, Y int
}

func main() {

	var wallMap map[coordinate]bool
	var boxMap map[coordinate]bool
	var goalMap map[coordinate]bool
	var player coordinate
	var mapLength int
	var mapHeight int
	// var mapCans int

	wallMap, boxMap, goalMap, player, mapLength, mapHeight, _ = readMap("test_map.txt")

	i := 1
	var i2 int
	var dontPrintDot bool

	for i < mapHeight+1 {
		i2 = 1
		for i2 < mapLength+1 {
			dontPrintDot = false
			for keys := range wallMap {
				if i == keys.Y && i2 == keys.X {
					fmt.Print("X")
					dontPrintDot = true
				}
			}
			for keys := range goalMap {
				if i == keys.Y && i2 == keys.X {
					fmt.Print("G")
					dontPrintDot = true
				}
			}
			for keys := range boxMap {
				if i == keys.Y && i2 == keys.X {
					fmt.Print("J")
					dontPrintDot = true
				}
			}
			if !dontPrintDot {
				fmt.Print(".")
			}
			i2++
		}
		fmt.Print("\n")
		i++
	}

	puzz := initializePuzzle(goalMap, boxMap, wallMap, player, false)
	success := search(puzz, "bfs")
	fmt.Println("puzzle was completed: ", success)
}
