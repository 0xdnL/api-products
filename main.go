package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

var (
	addr string = "localhost:8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info(r.RequestURI)
	fmt.Fprintf(w, "hello world")
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Info("[getAllProducts] ", r.RequestURI)
	json.NewEncoder(w).Encode(Products)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	log.Info("[getProductById] ", r.RequestURI)

	key := r.URL.Path[len("/product/"):]

	id, err := strconv.Atoi(key)
	if err != nil {
		panic(err)
	}

	for _, product := range Products {
		if product.Id == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleReq() {
	http.HandleFunc("/product/", getProductById)
	http.HandleFunc("/products", getAllProducts)
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}

func main() {

	Products = []Product{
		{Id: 1, Name: "chair", Quantity: 100, Price: 100.00},
		{Id: 2, Name: "table", Quantity: 150, Price: 185.00},
		{Id: 3, Name: "lamp", Quantity: 70, Price: 89.90},
	}

	handleReq()
}
