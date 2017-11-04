package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	sum := big.NewInt(2) //2^1000 is too big to handle with convetional vars
	var addSum int
	var sumStringed string
	two := big.NewInt(2)

	for i := 1; i < 1000; i++ { // 2^1000, math.Pow is not precise enough
		sum.Mul(sum, two)
	}
	sumStringed = sum.String() //convert it to a string so we can fetch each digit

	for i := 0; i < len(sumStringed); i++ {
		tmp, _ := strconv.Atoi(string(sumStringed[i])) //fetch each digit and
		addSum += tmp                                  //sum them badboys up
	}

	fmt.Println(addSum) //1366

	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed)

}
