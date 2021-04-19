package main

import (
	"log"
	"net/http"

	"github.com/Ajmalazeem/flurn/api"
)

func main() {

	loan := api.NewLoanService()

	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(loan))
}
