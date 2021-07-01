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
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "KIND Protein, Crunchy Peanut Butter",
				Price:    15.40,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "ALOHA Chocolate Chip Cookie",
				Price:    29.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "GoMacro Macrobar Protein Paradise, Cashew Caramel",
				Price:    25.75,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "Primal Kitchen Almond Spice",
				Price:    24.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "ThinkThin Protein & Superfruit Bar",
				Price:    16.50,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
			},
			Product{
				Name:     "Clif Bar Whey Protein Salted Caramel Cashew",
				Price:    9.60,
				InStock:  true,
				Rating:   5,
				ImageUrl: fmt.Sprintf("https://picsum.photos/100?color&random=%d", rand.Intn(99999)),
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
