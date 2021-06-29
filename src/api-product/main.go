package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Name            string
	Price           float32
	InStock         bool
	CampaignStartOn time.Time
}
type ProductList []Product

func main() {
	http.HandleFunc("/get-product-list", func(w http.ResponseWriter, r *http.Request) {
		productList := ProductList{
			Product{Name: "LED TV"},
			Product{Name: "Coffee Maker"},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productList)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
