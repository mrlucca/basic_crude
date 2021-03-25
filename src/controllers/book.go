package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mrlucca/basic_crud/src/models"
	"github.com/mrlucca/basic_crud/src/services/book"
)

func Book(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get all books")
	case "POST":
		var newBook models.Book
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			log.Fatalln(err)
		}
		created, err := book.Create(newBook)
		if err != nil {
			log.Fatalln(err)
		}
		if created {
			fmt.Fprintf(w, "Book created!")
		} else {
			fmt.Fprintf(w, "Book not created!")
		}
	case "PUT":
		fmt.Fprintf(w, "Change Book")
	case "DELETE":
		fmt.Fprintf(w, "Delete Book")
	default:
		msg := fmt.Sprintf("Method %v not implemented", r.Method)
		http.Error(w, msg, http.StatusNotFound)

	}
}
