package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		countS := 0
		countT := 0

		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				countS++
			}
			if s[i] == t[j] {
				countT++
			}
		}

		if countS != countT {
			return false
		}
	}

	return true
}
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(isAnagram(s, t))
}
