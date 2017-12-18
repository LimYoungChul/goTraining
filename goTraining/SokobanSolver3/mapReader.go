package main

import (
	"fmt"
	"io/ioutil"
)

func readMap(fileName string) {
	problemFolder := "Maps/"
	b, err := ioutil.ReadFile(problemFolder + fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}
