package models

import (
	"time"
)

type Book struct {
	Title           string    `json:"title"`
	Subtitle        string    `json:"subtitle"`
	PublicationDate time.Time `json:"publication_date"`
	Image           []byte    `json:"image"`
}
