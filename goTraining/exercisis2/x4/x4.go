package main

import "fmt"

func main() {
	if (true && false) || (false && true) || !(false && true) {
		fmt.Println("dwadwa")
	}
}
