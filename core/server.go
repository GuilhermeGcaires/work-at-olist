package main

import "github.com/guilhermegcaires/olist/core/db"

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
