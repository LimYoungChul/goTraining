package main

import (
	"fmt"
	"os"
)

var openList []Coordinate
var closedList []Coordinate
var destination Coordinate

func aSTAR(origin Coordinate, dest Coordinate, matrix [][]byte) (bool, []Coordinate) {
	if !validPoint(origin, matrix) && !validPoint(dest, matrix) {
		fmt.Println("origin or destination is not a valid point")
		os.Exit(-1)
	}
	destination = dest
	closedList = nil //reset this fucker
	openList = append(openList, origin)
	succes, finalPath := findPath(matrix)
	return succes, finalPath
}

func validPoint(point Coordinate, matrix [][]byte) bool {
	if matrix[point.X][point.Y] == X {
		return false
	}
	return true
}

func findPath(matrix [][]byte) (bool, []Coordinate) {
	for len(openList) != 0 {
		currentNode := getClosestNode()
		if addToClosedList(currentNode) {
			finalPath := generatePath(currentNode)
			return true, finalPath
		}
		removeFromOpenList(currentNode)
		newNodes := getValidNodes(currentNode, matrix)

		for _, n := range newNodes {
			addToOpenList(n)
		}
	}
	return false, nil
}

func generatePath(c Coordinate) []Coordinate {
	var finalPath []Coordinate
	node := c
	for node.Parent != nil {
		finalPath = append(finalPath, node)
		node = *node.Parent
	}
	return finalPath
}

func addToClosedList(c Coordinate) bool {

	if c.X == destination.X && c.Y == destination.Y {
		return true
	}
	closedList = append(closedList, c)

	return false
}

func addToOpenList(c Coordinate) {
	if checkExist(c, closedList) {
		return
	}
	if !checkExist(c, openList) {
		openList = append(openList, c)
	} else {
		if openList[findPoint(c, openList)].heuristic > c.heuristic {
			openList[findPoint(c, openList)].Parent = c.Parent
		}
	}

}

func checkExist(c Coordinate, arr []Coordinate) bool {
	for _, point := range arr {
		if c.X == point.X && c.Y == point.Y {
			return true
		}
	}
	return false
}

func findPoint(c Coordinate, arr []Coordinate) int {
	for index, point := range arr {
		if c.X == point.X && c.Y == point.Y {
			return index
		}
	}
	return -1
}

func removeFromOpenList(c Coordinate) {
	index := findPoint(c, openList)
	openList = append(openList[:index], openList[index+1:]...)
}

func getClosestNode() Coordinate {
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

func getValidNodes(c Coordinate, matrix [][]byte) []Coordinate {

	var walkable []Coordinate

	row := c.X
	column := c.Y

	left := matrix[row][column-1]
	right := matrix[row][column+1]
	up := matrix[row-1][column]
	down := matrix[row+1][column]

	if left == DOT || left == G {
		walkable = append(walkable, newCoordinate(row, column-1, &c))
	}
	if right == DOT || right == G {
		walkable = append(walkable, newCoordinate(row, column+1, &c))
	}
	if up == DOT || up == G {
		walkable = append(walkable, newCoordinate(row-1, column, &c))
	}
	if down == DOT || down == G {
		walkable = append(walkable, newCoordinate(row+1, column, &c))
	}

	return walkable
}

func distance(x, y int, destination Coordinate) int {
	distanceX := absInt(x - destination.X)
	distanceY := absInt(y - destination.Y)
	totalDistance := distanceX + distanceY
	return (totalDistance)
}

func absInt(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
