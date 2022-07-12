package model

import (
	"github.com/hieuprogrammer/go-mysql-books-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)

	return book
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)

	return &book, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)

	return book
}
