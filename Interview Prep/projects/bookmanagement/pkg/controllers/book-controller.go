package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/suraj8108/bookmanage/pkg/models"
	"github.com/suraj8108/bookmanage/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allBooks := models.GetAllBooks()

	res, _ := json.Marshal(allBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 10, 64)

	if err != nil {
		log.Fatal("Error while parsing")
	}

	bookDetails, db := models.GetBookById(bookId)

	if db.RowsAffected < 1 {
		resp, _ := json.Marshal(map[string]string{"message": "Data Not Present"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
	} else {
		resp, _ := json.Marshal(&bookDetails)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyDetails := r.Body
	booksInByte, err := ioutil.ReadAll(bodyDetails)
	if err != nil {
		fmt.Println("Error while reading the request body")
	}

	err = json.Unmarshal(booksInByte, &NewBook)
	if err != nil {
		fmt.Println("Error while unmashalling the body")
	}

	createdBook := NewBook.CreateBook()
	fmt.Println("Book created")

	w.WriteHeader(http.StatusCreated)

	response, _ := json.Marshal(*createdBook)
	w.Write(response)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	deletedBook := models.DeleteBook(bookId)

	resp, err := json.Marshal(deletedBook)

	if err != nil {
		fmt.Println("Error while marshaling the data")
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newBook models.Book
	// Parse the body
	utils.ParseBody(r, &newBook)

	// Book Details
	bookId := newBook.ID
	fmt.Println(newBook)
	_, db := models.GetBookById(int64(bookId))

	if db.RecordNotFound() {
		fmt.Println("Creating a new record")
	}
	db.Save(newBook)

	resp, _ := json.Marshal(&newBook)
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)

}
