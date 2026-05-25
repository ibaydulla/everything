package main

import "fmt"

func isPowerOfFour(n int) bool {
	pow := 1
	for pow <= n {
		if n == pow {
			return true
		}

		pow = pow * 4
	}

	return false
}

func main() {
	var n int
	fmt.Println("san giriz:")
	fmt.Scan(&n)
	fmt.Println(isPowerOfFour(n))
}
  