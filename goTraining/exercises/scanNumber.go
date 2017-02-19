package scanNumber

import "fmt"

func main() {

	x := 0
	y := 0

	fmt.Println("Enter two numbers, one larger than the other:")
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x > y {
		fmt.Println(x % y)
	} else {
		fmt.Println(y % x)

	}
}
