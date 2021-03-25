package book

import (
	"log"

	"github.com/mrlucca/basic_crud/src/models"
)

func Create(b models.Book) (bool, error) {
	log.Println("Create book")
	log.Println(b)
	return true, nil
}
