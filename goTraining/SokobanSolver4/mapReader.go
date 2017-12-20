package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func readMap(fileName string) (map[coordinate]bool, map[coordinate]bool, map[coordinate]bool, coordinate, int, int, int) {
	problemFolder := "Maps/"
	b, err := ioutil.ReadFile(problemFolder + fileName)
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	fmt.Println(str)
	wallMap := make(map[coordinate]bool)
	boxMap := make(map[coordinate]bool)
	goalMap := make(map[coordinate]bool)
	var player coordinate
	var mapLengthString string
	var mapHeightString string
	var mapCansString string
	var mapLength int
	var mapHeight int
	var mapCans int
	var spaceCounter int
	var lengthCounter int
	var heightCounter int
	for _, v := range b {
		if spaceCounter < 3 {
			if v == byte(' ') {
				spaceCounter++
			} else {
				switch spaceCounter {
				case 0:
					mapLengthString += string(v)
				case 1:
					mapHeightString += string(v)
				case 2:
					mapCansString += string(v)
				}
			}
		} else {
			switch v {
			case '\n':
				lengthCounter = 0
				heightCounter++
			case byte('X'):
				wallMap[coordinate{lengthCounter, heightCounter}] = true
			case byte('J'):
				boxMap[coordinate{lengthCounter, heightCounter}] = true
			case byte('G'):
				goalMap[coordinate{lengthCounter, heightCounter}] = true
			case byte('M'):
				player = coordinate{lengthCounter, heightCounter}
			}
			lengthCounter++
		}
	}
	mapLength, _ = strconv.Atoi(mapLengthString)
	mapHeight, _ = strconv.Atoi(mapHeightString)
	mapCans, _ = strconv.Atoi(mapCansString)
	return wallMap, boxMap, goalMap, player, mapLength, mapHeight, mapCans

}
