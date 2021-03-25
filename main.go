package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mrlucca/basic_crud/env"
	"github.com/mrlucca/basic_crud/src/controllers"
	"github.com/mrlucca/basic_crud/src/services/config"
)

func main() {
	config.TestDbConnection()
	config.ConfigureDb()
	go http.HandleFunc("/", controllers.Home)
	go http.HandleFunc("/login", controllers.SignIn)
	go http.HandleFunc("/login/register", controllers.SignUp)
	go http.HandleFunc("/book", controllers.Book)

	port := fmt.Sprintf(":%v", env.PORT)
	fmt.Printf("Server started in %v \n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
