package main

import (
	"fmt"
	"time"
)

type node struct {
	s      state
	parent *node
	move   string
}

func newNode(s state, parent *node, move string) *node {
	if move == "" {
		return nil
	}
	newNode := node{}
	newNode.s = s
	newNode.parent = parent
	newNode.move = move
	return &newNode

}

func search(p puzzle) bool {
	start := time.Now()
	startNode := node{p.initialState, nil, ""}
	var n node
	totalNodes := 1
	if p.goalTest(startNode.s) {
		return done(start, totalNodes, startNode, false)
	}
	exploredNodes := make(map[state]bool)
	var queue []*node
	var actions []action
	queue = append(queue, &startNode)

	for len(queue) > 0 {
		n = *queue[0] //take first element in queue
		queue[0] = nil
		queue = queue[1:] //remove first element in queue
		exploredNodes[n.s] = true
		if totalNodes > 50000000 {
			fmt.Println("niggers")
			break
		}
		actions = p.getActions(n.s)
		for v := range actions {
			child := getChild(actions[v], n)
			if child != nil {
				totalNodes++
				if !exploredNodes[child.s] {
					if p.goalTest(child.s) {
						return done(start, totalNodes, *child, false)
					}
					if !p.deadTest(child.s) {
						queue = append(queue, child)
					}
				}
			}
		}
	}

	return done(start, totalNodes, startNode, true)
	// fmt.Println("???")
	// return done(start, totalNodes, n, 0, false)
}

func getChild(a action, n node) *node {
	boxes := make(map[coordinate]bool)
	var boxSlice [4]coordinate
	var i int
	for _, v := range n.s.boxes {
		boxes[v] = true
	}

	var newPlayer coordinate
	var newBox coordinate

	switch a.act {
	case "U":
		newPlayer = a.newPlayerPosition
		if boxes[newPlayer] {
			newBox = coordinate{newPlayer.X, newPlayer.Y - 1}
			delete(boxes, newPlayer)
			boxes[newBox] = true
		}
	case "D":
		newPlayer = a.newPlayerPosition
		if boxes[newPlayer] {
			newBox = coordinate{newPlayer.X, newPlayer.Y + 1}
			delete(boxes, newPlayer)
			boxes[newBox] = true
		}
	case "L":
		newPlayer = a.newPlayerPosition
		if boxes[newPlayer] {
			newBox = coordinate{newPlayer.X - 1, newPlayer.Y}
			delete(boxes, newPlayer)
			boxes[newBox] = true
		}
	case "R":
		newPlayer = a.newPlayerPosition
		if boxes[newPlayer] {
			newBox = coordinate{newPlayer.X + 1, newPlayer.Y}
			delete(boxes, newPlayer)
			boxes[newBox] = true
		}
	}
	for key, value := range boxes {
		if value {
			boxSlice[i] = key
			i++
		}
	}

	return newNode(newState(boxSlice, newPlayer), &n, a.act)
}

func done(start time.Time, totalNodes int, lastNode node, success bool) bool {
	steps := 0
	if !success {
		result := ""
		n := lastNode
		for n.parent != nil {
			result = n.move + " " + result
			n = *n.parent
			steps++
		}
		fmt.Println(n.move + " " + result)
		fmt.Println("number of steps: ", steps)
		fmt.Println("total number of nodes explored", totalNodes)
		elapsed := time.Since(start)
		fmt.Println("Total Runtime: ", elapsed)
	} else {
		fmt.Println("Failed to complete puzzle")
	}

	return success

}
