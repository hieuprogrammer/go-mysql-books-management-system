package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hieuprogrammer/go-mysql-books-management-system/pkg/route"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	route.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Printf("Starting Go HTTP server on port: 9010.\n")
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
