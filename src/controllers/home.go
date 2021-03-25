package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Metadata struct {
	Method      string
	Body        interface{}
	QueryParams map[string][]string
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home acessed!")
	var m Metadata
	m = Metadata{
		Method:      r.Method,
		QueryParams: r.URL.Query(),
	}

	err := json.NewDecoder(r.Body).Decode(&m.Body)

	metadata, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(metadata))
}
