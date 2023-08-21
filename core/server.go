package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilhermegcaires/olist/core/db"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Setup()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the homepage!")
	})

	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the about page.")
	})

	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Contact us at: contact@example.com")
	})

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
