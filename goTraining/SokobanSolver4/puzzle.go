package main

import "fmt"

type state struct {
	boxes  [4]coordinate
	player coordinate
	//hashCode int hopefully this is unique, used for exploredNodes map
}

type action struct {
	act               string
	newPlayerPosition coordinate
}

func newState(boxes [4]coordinate, player coordinate) state {
	newState := state{}
	newState.boxes = boxes
	newState.player = player
	// newState.hashCode = 17
	//
	// for _, v := range boxes {
	// 	newState.hashCode = 37*newState.hashCode + (v.X*1000 + v.Y)
	// }
	// newState.hashCode = 37*newState.hashCode + (player.X*1000 + player.Y)

	return newState
}

type puzzle struct {
	initialState state
	wallMap      map[coordinate]bool
	goalMap      map[coordinate]bool
	aStar        bool
}

func (p puzzle) goalTest(s state) bool {
	for _, v := range s.boxes {
		if !p.goalMap[v] {
			return false
		}
	}
	return true
}

//try implement a better deadlock test :)
func (p puzzle) deadTest(s state) bool {
	for _, v := range s.boxes {
		up := coordinate{v.X, v.Y - 1}
		down := coordinate{v.X, v.Y + 1}
		left := coordinate{v.X - 1, v.Y}
		right := coordinate{v.X + 1, v.Y}
		if !p.goalMap[coordinate{v.X, v.Y}] {
			if (p.wallMap[up] || p.wallMap[down]) && (p.wallMap[left] || p.wallMap[right]) {
				return true
			}
		}
	}
	return false
}

func (p puzzle) getActions(s state) []action {
	var actions []action
	var newPlayer coordinate
	var newBox coordinate
	player := s.player
	newBoxMap := make(map[coordinate]bool)
	for _, v := range s.boxes {
		newBoxMap[v] = true
	}
	for _, v := range s.boxes {
		newPlayer = coordinate{v.X, v.Y + 1}
		newBox = coordinate{v.X, v.Y - 1}
		if !p.wallMap[newPlayer] {
			if !(newBoxMap[newPlayer]) && (!newBoxMap[newBox] && !p.wallMap[newBox]) {
				if aSTAR(player, aStarNode{0, newPlayer, nil}, s, p) {
					actions = append(actions, action{"U", v})
				}
			}
		}
		newPlayer = coordinate{v.X, v.Y - 1}
		newBox = coordinate{v.X, v.Y + 1}
		if !p.wallMap[newPlayer] {
			if !(newBoxMap[newPlayer]) && (!newBoxMap[newBox] && !p.wallMap[newBox]) {
				if aSTAR(player, aStarNode{0, newPlayer, nil}, s, p) {
					actions = append(actions, action{"D", v})
				}
			}
		}
		newPlayer = coordinate{v.X + 1, v.Y}
		newBox = coordinate{v.X - 1, v.Y}
		if !p.wallMap[newPlayer] {
			if !(newBoxMap[newPlayer]) && (!newBoxMap[newBox] && !p.wallMap[newBox]) {
				if aSTAR(player, aStarNode{0, newPlayer, nil}, s, p) {
					actions = append(actions, action{"L", v})
				}
			}
		}
		newPlayer = coordinate{v.X - 1, v.Y}
		newBox = coordinate{v.X + 1, v.Y}
		if !p.wallMap[newPlayer] {
			if !(newBoxMap[newPlayer]) && (!newBoxMap[newBox] && !p.wallMap[newBox]) {
				if aSTAR(player, aStarNode{0, newPlayer, nil}, s, p) {
					actions = append(actions, action{"R", v})
				}
			}
		}
	}
	return actions
}

func initializePuzzle(goalMap map[coordinate]bool, boxMap map[coordinate]bool, wallMap map[coordinate]bool, player coordinate, useAstar bool) puzzle {
	newPuzzle := puzzle{}
	newPuzzle.goalMap = goalMap
	newPuzzle.wallMap = wallMap
	var boxes [4]coordinate
	var i int
	for key := range boxMap {
		boxes[i] = key
		i++
	}
	fmt.Println("player in initialize puzzle:", player)
	newPuzzle.initialState = newState(boxes, player)
	newPuzzle.aStar = useAstar

	return newPuzzle
}
