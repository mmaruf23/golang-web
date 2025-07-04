package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "All Products")
	})
	mux.HandleFunc("/products/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Product with id = 1")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
