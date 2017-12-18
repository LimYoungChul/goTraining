package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Println("Runtime: ", elapsed)
}
