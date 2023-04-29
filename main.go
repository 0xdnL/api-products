package main

import (
	"fmt"
	"net/http"
)

var (
	addr string = "localhost:8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
