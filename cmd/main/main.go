package main

import (
	"go_book_store/pkg/config"
	"go_book_store/pkg/models"
	"go_book_store/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	models.InitBookModel()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
