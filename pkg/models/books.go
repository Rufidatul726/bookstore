package models

import (
	"github.com/jinzhu/gorm"

	"github.com/Rufidatul726/bookstore/pkg/config"
)

var db *gorm.DB

// Book is a struct to mapping the book table in the database
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func init() {
	config.ConnectDB()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook is a function to create a new book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks is a function to get all books
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookByID is a function to get a book by id
func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

// UpdateBook is a function to update a book
func (b *Book) UpdateBook(id int64) *Book {
	db.Save(&b)
	return b
}

// DeleteBook is a function to delete a book
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}
