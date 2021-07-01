package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Product struct {
	Name     string
	Price    float32
	InStock  bool
	Rating   int
	ImageUrl string
}
type ProductList []Product

func main() {
	http.HandleFunc("/get-list", func(w http.ResponseWriter, r *http.Request) {
		productList := ProductList{
			Product{
				Name:     "Berry Mix Organic Bar",
				Price:    14.80,
				InStock:  true,
				Rating:   4,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "RXBar Chocolate Sea Salt",
				Price:    38.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/50?color&random=%d", rand.Intn(10000)),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productList)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
