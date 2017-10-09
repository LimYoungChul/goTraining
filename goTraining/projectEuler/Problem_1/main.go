package main

import "fmt"

func main() {

	var sum int
	// we simply add all the multies of either 3 or 5 together and print out the sum, answers question 1.
	for n := 0; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}
	fmt.Println(sum) //233168 should be the returned value.
}
