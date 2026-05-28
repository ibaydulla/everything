package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func salam(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	n, _ := strconv.Atoi(login)

	fmt.Fprintf(w, "%d-nji yonekey san: %d", n, yonekey(n))
}

func hosh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye World, Method: ", "Header", r.Header, "Host", r.Host, "Pattern", r.Pattern, "RemoteAddr", r.RemoteAddr, "RequestUrI", r.RequestURI, "Url.Host", r.URL.Host)
}

func yonekey(n int) int {
	var i int = 2
	var a []int
	for len(a) <= n {
		isPrime := true
		for _, value := range a {
			if i%value == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			a = append(a, i)
		}
		i++
	}
	return a[n]
}

func main() {
	http.HandleFunc("/hello", salam)
	http.HandleFunc("/hosh", hosh)
	http.ListenAndServe(":8080", nil)
}
