package main

import "log"

func reverse(s string) string {
	if len(s) == 0 {
		return ""
	}
	k := ""
	k += reverse(s[1:])
	k += string(s[0])

	return k
}

func main() {

	s1 := reverse("Hello, World!")
	log.Println(s1)
}
