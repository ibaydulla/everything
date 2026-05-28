package main

import "fmt"

func add(n int) int {
	if n == 1 {
		return 1
	}
	n += add(n - 1)
	return n
}

func main() {
	k := 0
	fmt.Scan(&k)
	s := add(k)
	fmt.Printf("1-den %d cenli sanlaryn jemi: %d ", k, s)
}
