package models

import (
	"time"
)

type Book struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Subtitle        string    `json:"subtitle"`
	Price           int       `json:"price"`
	PublicationDate time.Time `json:"publication_date"`
	Image           []byte    `json:"image"`
}
