package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devansh777/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Printf("serving port 9010'n")
	log.Fatal(http.ListenAndServe(":9010", r))
}
