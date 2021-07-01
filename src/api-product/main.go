package main

import (
	"encoding/json"
	"fmt"
	"log"
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
				Name:     "Rise Bar Lemon Cashew",
				Price:    25.90,
				InStock:  true,
				Rating:   4,
				ImageUrl: "https://i1.wp.com/www.eatthis.com/wp-content/uploads/media/images/ext/254043726/protein-bars-rise.jpg",
			},
			Product{
				Name:     "RXBar Chocolate Sea Salt",
				Price:    38.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i1.wp.com/www.eatthis.com/wp-content/uploads/2017/03/rxbar-dark-chocolate-sea-salt-protein-bar.jpg",
			},
			Product{
				Name:     "KIND Protein, Crunchy Peanut Butter",
				Price:    15.40,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i0.wp.com/www.eatthis.com/wp-content/uploads/2018/05/kind-protein-crunchy-peanut-butter.jpg",
			},
			Product{
				Name:     "ALOHA Chocolate Chip Cookie",
				Price:    29.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i0.wp.com/www.eatthis.com/wp-content/uploads/2017/03/aloha-chocolate-chip-cookie-dough-protein-bar.jpg",
			},
			Product{
				Name:     "GoMacro Macrobar Protein Paradise, Cashew Caramel",
				Price:    25.75,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i2.wp.com/www.eatthis.com/wp-content/uploads/media/images/ext/700050229/protein-bars-gomacro.jpg",
			},
			Product{
				Name:     "Primal Kitchen Almond Spice",
				Price:    24.90,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i1.wp.com/www.eatthis.com/wp-content/uploads/2018/09/primal-kitchen-almond-spice-protein-bar.jpg",
			},
			Product{
				Name:     "ThinkThin Protein & Superfruit Bar",
				Price:    16.50,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i0.wp.com/www.eatthis.com/wp-content/uploads/2019/03/thinkthin-coconut-almond-chia-bar.jpg",
			},
			Product{
				Name:     "Clif Bar Whey Protein Salted Caramel Cashew",
				Price:    9.60,
				InStock:  true,
				Rating:   5,
				ImageUrl: "https://i1.wp.com/www.eatthis.com/wp-content/uploads/media/images/ext/721025516/protein-bars-clif-whey.jpg",
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
