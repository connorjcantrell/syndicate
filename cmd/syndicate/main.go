package main

import (
	"log"
	"net/http"

	"github.com/connorjcantrell/syndicate/postgres"
	"github.com/connorjcantrell/syndicate/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@5432:5432/syndicate?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
