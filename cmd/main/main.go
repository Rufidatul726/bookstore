package main

import (
	"log"
	"net/http"

	"github.com/Rufidatul726/bookstore/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookstoreRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
