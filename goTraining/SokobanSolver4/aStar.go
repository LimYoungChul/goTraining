package main

import (
	"fmt"
	"os"
)

type aStarNode struct {
	heuristic int
	c         coordinate
	parent    *aStarNode
}

func newaStarNode(c coordinate, parent *aStarNode, destination aStarNode) aStarNode {
	newNode := aStarNode{}
	newNode.c = c
	newNode.parent = parent
	newNode.heuristic = (absInt(newNode.c.X-destination.c.X) + absInt(newNode.c.Y-destination.c.Y))

	return newNode
}

var openList []aStarNode
var closedList []aStarNode
var destination aStarNode

func aSTAR(originC coordinate, dest aStarNode, s state, p puzzle) bool {
	origin := newaStarNode(originC, nil, destination)
	if !validPoint(origin, p) && !validPoint(dest, p) {
		fmt.Println("origin or destination is not a valid point")
		os.Exit(-1)
	}
	destination = dest
	closedList = nil //reset this fucker
	openList = nil   //reset this fucker too
	openList = append(openList, origin)
	success := findPath(p, s, destination)
	return success
}

func validPoint(n aStarNode, p puzzle) bool {
	if p.wallMap[n.c] {
		return false
	}
	return true
}

func findPath(p puzzle, s state, destination aStarNode) bool {
	for len(openList) != 0 {
		currentNode := getClosestNode()
		if addToClosedList(currentNode, destination) {
			//finalPath := generatePath(currentNode)
			return true
		}
		removeFromOpenList(currentNode)
		newNodes := getValidNodes(currentNode, destination, s, p)

		for _, n := range newNodes {
			addToOpenList(n)
		}
	}
	return false
}

func generatePath(node aStarNode) []aStarNode {
	var finalPath []aStarNode
	n := node
	for n.parent != nil {
		finalPath = append([]aStarNode{node}, finalPath...)
		n = *node.parent
	}
	finalPath = append([]aStarNode{node}, finalPath...)
	return finalPath
}

func addToClosedList(node, destination aStarNode) bool {
	closedList = append(closedList, node)
	if node.c.X == destination.c.X && node.c.Y == destination.c.Y {
		return true
	}

	return false
}

func addToOpenList(c aStarNode) {
	if checkExist(c, closedList) {
		return
	}
	if !checkExist(c, openList) {
		openList = append(openList, c)
	} else {
		if openList[findPoint(c, openList)].heuristic > c.heuristic {
			openList[findPoint(c, openList)].parent = c.parent
		}
	}

}

func checkExist(node aStarNode, arr []aStarNode) bool {
	for _, point := range arr {
		if node.c.X == point.c.X && node.c.Y == point.c.Y {
			return true
		}
	}
	return false
}

func findPoint(node aStarNode, arr []aStarNode) int {
	for index, point := range arr {
		if node.c.X == point.c.X && node.c.Y == point.c.Y {
			return index
		}
	}
	return -1
}

func removeFromOpenList(node aStarNode) {
	index := findPoint(node, openList)
	openList = append(openList[:index], openList[index+1:]...)
}

func getClosestNode() aStarNode {
	if len(openList) == 0 {
		fmt.Println("no viable nodes left")
		os.Exit(-1)
	}
	index := 0
	for i, c := range openList {
		if (i > 0) && (c.heuristic <= openList[index].heuristic) {
			index = i
		}
	}
	return openList[index]
}

func getValidNodes(node, destination aStarNode, s state, p puzzle) []aStarNode {

	var walkable []aStarNode
	boxMap := make(map[coordinate]bool)
	for _, value := range s.boxes {
		boxMap[value] = true
	}
	var newC coordinate
	newC = coordinate{node.c.X - 1, node.c.Y}
	if !p.wallMap[newC] && !boxMap[newC] {
		walkable = append(walkable, newaStarNode(newC, &node, destination))
	}
	newC = coordinate{node.c.X + 1, node.c.Y}
	if !p.wallMap[newC] && !boxMap[newC] {
		walkable = append(walkable, newaStarNode(newC, &node, destination))
	}
	newC = coordinate{node.c.X, node.c.Y - 1}
	if !p.wallMap[newC] && !boxMap[newC] {
		walkable = append(walkable, newaStarNode(newC, &node, destination))
	}
	newC = coordinate{node.c.X - 1, node.c.Y + 1}
	if !p.wallMap[newC] && !boxMap[newC] {
		walkable = append(walkable, newaStarNode(newC, &node, destination))
	}

	return walkable
}

// func distance(x, y int, destination Coordinate) int {
// 	distanceX := absInt(x - destination.X)
// 	distanceY := absInt(y - destination.Y)
// 	totalDistance := distanceX + distanceY
// 	return (totalDistance)
// }

func absInt(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
