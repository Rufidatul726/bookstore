package models

import (
	"github.com/jinzhu/gorm"
)

// Book is a struct to mapping the book table in the database
type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

// TableName is a function to change the table name
func (b *Book) TableName() string {
	return "book"
}

// CreateBook is a function to create a new book
func CreateBook(book *Book) (*Book, error) {
	if err := config.DB.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// GetBook is a function to get all books
func GetBook() (*[]Book, error) {
	var books []Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}
