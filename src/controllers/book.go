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
		geted, err := book.Get()

		if err != nil {
			log.Fatalln(err)
		}
		if geted {
			fmt.Fprintf(w, "Book geted!")
		} else {
			http.Error(w, "Book not geted!", http.StatusNotFound)

		}
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
			http.Error(w, "Book not created!", http.StatusNotFound)

		}

	case "PUT":
		updated, err := book.Update()
		if err != nil {
			log.Fatalln(err)
		}
		if updated {
			fmt.Fprintf(w, "Book updated")
		} else {
			http.Error(w, "Book not updated", http.StatusNotFound)
		}

	case "DELETE":
		deleted, err := book.Delete()
		if err != nil {
			log.Fatalln(err)
		}
		if deleted {
			fmt.Fprintf(w, "Book deleted!")
		} else {
			http.Error(w, "Book not deleted!", http.StatusNotFound)
		}

	default:
		msg := fmt.Sprintf("Method %v not implemented", r.Method)
		http.Error(w, msg, http.StatusNotFound)

	}
}
