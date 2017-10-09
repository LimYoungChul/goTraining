package main

import (
	"fmt"
	"strconv"
)

func main() {

	var largest int //holder for the largest palindrome
	var x int       //current number
	var s string    //stringified version of x
	var s2 string   //reversed version of s
	//basically if a reversed version of a number is equal to itself then it is a palindrome, and that is how this problem is computed.

	for n := 1; n < 1000; n++ {
		for i := 1; i < 1000; i++ {
			x = n * i
			s = strconv.Itoa(x)
			s2 = reverseString(s)
			if s == s2 {
				if largest < x {
					largest = x
				}
			}
		}
	}
	fmt.Println(largest)
}

func reverseString(s string) string {

	n := len(s)
	secondHalf := make([]rune, n)

	for _, rune := range s {
		n--
		secondHalf[n] = rune
	}

	return string(secondHalf)

}
