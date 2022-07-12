package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hieuprogrammer/go-mysql-books-management-system/pkg/model"
	"github.com/hieuprogrammer/go-mysql-books-management-system/pkg/util"
	"net/http"
	"strconv"
)

var book model.Book

func CreateBook(response http.ResponseWriter, request *http.Request) {
	book := &model.Book{}
	util.ParseBody(request, book)
	newBook := book.CreateBook()
	res, _ := json.Marshal(newBook)
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBooks(response http.ResponseWriter, request *http.Request) {
	books := model.GetBooks()
	res, _ := json.Marshal(books)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Printf("Could not find book with ID: %s.\n", bookId)
	}
	book, _ := model.GetBookById(ID)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBookById(response http.ResponseWriter, request *http.Request) {
	var book = &model.Book{}
	util.ParseBody(request, book)
	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Printf("Could not find book with ID: %s.\n", bookId)
	}
	bookDetail, db := model.GetBookById(ID)
	if book.Name != "" {
		bookDetail.Name = book.Name
	}
	if book.Author != "" {
		bookDetail.Author = book.Author
	}
	if book.Publication != "" {
		bookDetail.Author = book.Publication
	}
	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func DeleteBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		fmt.Printf("Could not find book with ID: %s.\n", bookId)
	}
	book := model.DeleteBookById(ID)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
