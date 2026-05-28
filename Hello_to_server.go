package main

import (
	"fmt"
	"log/slog"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("HelloWorld\n"))
	if err != nil {
		slog.Error("error writing response",err)
		return
	}
	fmt.Printf("%d bytes written",wc)
}
