package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	addr string = "localhost:8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info(r.RequestURI)
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
