package main

import (
	"fmt"
)

func mySqrt(x int) int {
	i := 1

	for i <= x/i {
		i++
	}

	return i - 1
}
func main() {
	var x int
	fmt.Print("san giriz:")
	fmt.Scan(&x)
	fmt.Println(mySqrt(x))
}
