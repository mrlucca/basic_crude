package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrlucca/basic_crud/src/controllers"
)

const PORT string = ":7000"

func main() {
	go http.HandleFunc("/", controllers.Home)
	go http.HandleFunc("/login", controllers.SignIn)
	go http.HandleFunc("/login/register", controllers.SignUp)
	go http.HandleFunc("/book", controllers.Book)

	fmt.Printf("Server started in %v \n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}

}
