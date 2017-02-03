package scan

import "fmt"

func main() {

	var name string

	fmt.Println("please enter your name here:")
	fmt.Scan(&name)
	fmt.Println("Hello " + name)
}
