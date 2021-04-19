package main

import (
	"log"
	"net/http"

	"github.com/Ajmalazeem/flurn/api"
	"github.com/Ajmalazeem/flurn/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=example dbname=flurn"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	loanStore := store.NewLoanStore(db)

	loan := api.NewLoanService(loanStore)

	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(loan))
}
