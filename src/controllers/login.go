package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("SignIn accessed")
	fmt.Fprintf(w, "This is sign in!")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("SignUp accessed")
	fmt.Fprintf(w, "This is sign up!")
}
