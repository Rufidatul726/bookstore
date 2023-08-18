package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Rufidatul726/bookstore/pkg/models"
	"github.com/Rufidatul726/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// GetBook is a function to get all books
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById is a function to get a book by id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		panic(error)
	}
	bookDetail, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook is a function to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	bookDetail := CreateBook.CreateBook()
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook is a function to update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		panic(error)
	}
	bookDetail, db := models.GetBookByID(ID)
	utils.ParseBody(r, &bookDetail)
	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook is a function to delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, error := strconv.ParseInt(bookId, 0, 0)
	if error != nil {
		panic(error)
	}
	bookDetail := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
