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
	// var mapLength int
	// var mapHeight int
	// var mapCans int

	wallMap, boxMap, goalMap, player, _, _, _ = readMap("test_map2.txt")

	puzz := initializePuzzle(goalMap, boxMap, wallMap, player, false)
	//fmt.Println(puzz.getActions(puzz.initialState))
	fmt.Println(search(puzz))
	fmt.Println(aSTAR(puzz.initialState.player, aStarNode{0, coordinate{5, 3}, nil}, puzz.initialState, puzz))
	//fmt.Println("puzzle failed:", success)
}
