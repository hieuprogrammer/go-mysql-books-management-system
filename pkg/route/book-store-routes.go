package route

import (
	"github.com/gorilla/mux"
	"github.com/hieuprogrammer/go-mysql-books-management-system/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controller.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBookById).Methods("DELETE")
}
