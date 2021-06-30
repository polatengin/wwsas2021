package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Name    string
	Price   float32
	InStock bool
	Rating  int
}
type ProductList []Product

func main() {
	http.HandleFunc("/get-list", func(w http.ResponseWriter, r *http.Request) {
		productList := ProductList{
			Product{
				Name:    "Berry Mix Organic Bar",
				Price:   14.80,
				InStock: true,
				Rating:  4,
			},
			Product{
				Name:    "Coffee Maker",
				Price:   57.90,
				InStock: true,
				Rating:  5},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productList)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
