package main

import (
	"fmt"
	"time"
)

type node struct {
	s      state
	parent *node
	branch int
	move   string
}

func newNode(s state, parent *node, branch int, move string) *node {
	if move == "" {
		return nil
	}
	newNode := node{}
	newNode.s = s
	newNode.parent = parent
	newNode.branch = branch
	newNode.move = move
	return &newNode

}

func search(p puzzle, searchMethod string) bool {
	start := time.Now()
	startNode := node{p.initialState, nil, 0, ""}
	var n node
	totalNodes := 1
	if p.goalTest(startNode.s) {
		return done(start, totalNodes, startNode, false)
	}
	exploredNodes := make(map[state]bool)

	var actions []string

	if searchMethod == "bfs" {
		//var currentBranch int
		var queue []*node
		queue = append(queue, &startNode)
		for len(queue) > 0 {
			// n = *queue[0] //take first element in queue
			// queue[0] = nil
			// queue = queue[1:] //remove first element in queue
			n = *queue[len(queue)-1]
			//queue[len(queue)-1] = nil
			copy(queue[len(queue)-1:], queue[len(queue):])
			queue[len(queue)-1] = nil
			queue = queue[:len(queue)-1]
			//fmt.Println(n.branch)
			exploredNodes[n.s] = true
			// if n.branch > currentBranch {
			// 	currentBranch = n.branch
			// 	fmt.Println(currentBranch)
			// }
			actions = p.getActions(n.s)
			for _, v := range actions {
				child := getChild(v, n)
				totalNodes++
				if child.branch > 200 {
					break
				}
				if !exploredNodes[child.s] {
					if p.goalTest(child.s) {
						return done(start, totalNodes, *child, true)
					}
					if !p.deadTest(child.s) {
						queue = append(queue, child)
					}
				}
			}
		}
		fmt.Println("explored:", len(exploredNodes))
		fmt.Println(queue)
		fmt.Println(totalNodes)
		return done(start, totalNodes, n, false)

	} else if searchMethod == "dfs" {
		index := 0
		stack := make([][]*node, 0, 100)
		x := 0
		for x < 500 {
			stack = append(stack, []*node{})
			x++
		}
		var addToIndex bool
		stack[0] = append(stack[0], &startNode)
		fmt.Println(stack)
		fmt.Println(len(stack[index]) > 0)
		for {
			if len(stack[index]) > 0 {
				n = *stack[index][len(stack[index])-1]
				stack[index][len(stack[index])-1] = nil
				stack[index] = append(stack[index][:len(stack[index])-1], stack[index][len(stack[index]):]...)
				exploredNodes[n.s] = true
				actions = p.getActions(n.s)
				addToIndex = false
				for _, v := range actions {
					child := getChild(v, n)
					totalNodes++
					if child.branch < 499 {
						addToIndex = true
						if !exploredNodes[child.s] {
							if p.goalTest(child.s) {
								return done(start, totalNodes, *child, true)
							}
							if !p.deadTest(child.s) {
								stack[child.branch] = append(stack[child.branch], child)
							}
						}
					}
				}
				if addToIndex {
					index++
				}
			} else {
				if index != 0 {
					index--
				} else {
					return done(start, totalNodes, n, false)
				}
			}
		}
	}
	return false
}

func getChild(action string, n node) *node {
	newBoxes := n.s.boxes

	x := n.s.player.X
	y := n.s.player.Y

	var newPlayer coordinate
	var newBox coordinate
	var newBranch = n.branch + 1

	switch action {
	case "u":
		newPlayer = coordinate{x, y - 1}
		newBox = coordinate{x, y - 2}
		for index := range newBoxes {
			if newPlayer == newBoxes[index] {
				newBoxes[index] = newBox
				action = "U"
				break
			}
		}
	case "d":
		newPlayer = coordinate{x, y + 1}
		newBox = coordinate{x, y + 2}
		for index := range newBoxes {
			if newPlayer == newBoxes[index] {
				newBoxes[index] = newBox
				action = "D"
				break
			}
		}
	case "l":
		newPlayer = coordinate{x - 1, y}
		newBox = coordinate{x - 2, y}
		for index := range newBoxes {
			if newPlayer == newBoxes[index] {
				newBoxes[index] = newBox
				action = "L"
				break
			}
		}
	case "r":
		newPlayer = coordinate{x + 1, y}
		newBox = coordinate{x + 2, y}
		for index := range newBoxes {
			if newPlayer == newBoxes[index] {
				newBoxes[index] = newBox
				action = "R"
				break
			}
		}
	}
	return newNode(newState(newBoxes, newPlayer), &n, newBranch, action)
}

func done(start time.Time, totalNodes int, lastNode node, success bool) bool {
	steps := 0
	if success {
		result := ""
		n := lastNode
		for n.parent != nil {
			result = n.move + " " + result
			n = *n.parent
			steps++
		}
		fmt.Println(n.move + " " + result)
		fmt.Println("number of steps: ", steps)
	} else {
		fmt.Println("Failed to complete puzzle")
	}
	fmt.Println("last branch: ", lastNode.branch)
	fmt.Println("total number of nodes generated", totalNodes)
	elapsed := time.Since(start)
	fmt.Println("Total Runtime: ", elapsed)

	return success

}
