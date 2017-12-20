package main

import "fmt"

type state struct {
	boxes  [4]coordinate
	player coordinate
	//hashCode int hopefully this is unique, used for exploredNodes map
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
	boxMap := make(map[coordinate]bool)
	for _, v := range s.boxes {
		boxMap[v] = true
	}
	for _, v := range s.boxes {
		up := coordinate{v.X, v.Y - 1}
		down := coordinate{v.X, v.Y + 1}
		left := coordinate{v.X - 1, v.Y}
		right := coordinate{v.X + 1, v.Y}
		upnright := coordinate{v.X + 1, v.Y - 1}
		upnleft := coordinate{v.X - 1, v.Y - 1}
		downnright := coordinate{v.X + 1, v.Y + 1}
		downnleft := coordinate{v.X - 1, v.Y + 1}
		doubleup := coordinate{v.X, v.Y - 2}
		doubledown := coordinate{v.X, v.Y + 2}
		doubleleft := coordinate{v.X - 2, v.Y}
		doubleright := coordinate{v.X + 2, v.Y}

		if !p.goalMap[coordinate{v.X, v.Y}] {
			if (p.wallMap[up] || p.wallMap[down]) && (p.wallMap[left] || p.wallMap[right]) {
				return true
			}
		} else if p.wallMap[up] && p.wallMap[upnright] && boxMap[right] {
			return true
		} else if p.wallMap[up] && p.wallMap[upnleft] && boxMap[left] {
			return true
		} else if p.wallMap[down] && p.wallMap[downnright] && boxMap[right] {
			return true
		} else if p.wallMap[down] && p.wallMap[downnleft] && boxMap[left] {
			return true
		} else if p.wallMap[left] && p.wallMap[upnleft] && boxMap[up] {
			return true
		} else if p.wallMap[left] && p.wallMap[downnleft] && boxMap[down] {
			return true
		} else if p.wallMap[right] && p.wallMap[downnright] && boxMap[down] {
			return true
		} else if p.wallMap[right] && p.wallMap[upnright] && boxMap[up] {
			return true
		} else if p.wallMap[right] && p.wallMap[upnright] && p.wallMap[downnright] && p.wallMap[doubledown] && p.wallMap[doubleup] && !p.goalMap[up] && !p.goalMap[down] {
			return true
		} else if p.wallMap[left] && p.wallMap[upnleft] && p.wallMap[downnleft] && p.wallMap[doubledown] && p.wallMap[doubleup] && !p.goalMap[up] && !p.goalMap[down] {
			return true
		} else if p.wallMap[up] && p.wallMap[upnright] && p.wallMap[upnleft] && p.wallMap[doubleleft] && p.wallMap[doubleright] && !p.goalMap[left] && !p.goalMap[right] {
			return true
		} else if p.wallMap[down] && p.wallMap[downnright] && p.wallMap[downnleft] && p.wallMap[doubleleft] && p.wallMap[doubleright] && !p.goalMap[left] && !p.goalMap[right] {
			return true
		}
	}
	return false
}

func (p puzzle) getActions(s state) []string {
	var actions []string
	var newPlayer coordinate
	var newBox coordinate
	newBoxMap := make(map[coordinate]bool)
	for _, v := range s.boxes {
		newBoxMap[v] = true
	}
	// copy(boxes, n.s.boxes)
	if !p.aStar {
		newPlayer = coordinate{s.player.X, s.player.Y - 1}
		newBox = coordinate{s.player.X, s.player.Y - 2}
		if !p.wallMap[newPlayer] {
			if (newBoxMap[newPlayer]) && (newBoxMap[newBox] || p.wallMap[newBox]) {

			} else {
				actions = append(actions, "u")
			}
		}

		newPlayer = coordinate{s.player.X, s.player.Y + 1}
		newBox = coordinate{s.player.X, s.player.Y + 2}
		if !p.wallMap[newPlayer] {
			if (newBoxMap[newPlayer]) && (newBoxMap[newBox] || p.wallMap[newBox]) {

			} else {
				actions = append(actions, "d")
			}
		}
		newPlayer = coordinate{s.player.X - 1, s.player.Y}
		newBox = coordinate{s.player.X - 2, s.player.Y}
		if !p.wallMap[newPlayer] {
			if (newBoxMap[newPlayer]) && (newBoxMap[newBox] || p.wallMap[newBox]) {

			} else {
				actions = append(actions, "l")
			}
		}
		newPlayer = coordinate{s.player.X + 1, s.player.Y}
		newBox = coordinate{s.player.X + 2, s.player.Y}
		if !p.wallMap[newPlayer] {
			if (newBoxMap[newPlayer]) && (newBoxMap[newBox] || p.wallMap[newBox]) {

			} else {
				actions = append(actions, "r")
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
