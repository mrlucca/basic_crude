package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mrlucca/basic_crud/env"
)

func GetDbConnection() (*sql.DB, error) {
	dbConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
	)
	return sql.Open("postgres", dbConn)
}

func TestDbConnection() {
	db, err := GetDbConnection()
	if err != nil {
		errMensage := fmt.Sprintf("Database not found! err: %v", err)
		panic(errMensage)
	}
	if !databaseConnected(db) {
		panic("Db not connected")
	}
	defer db.Close()
	log.Println("Successfully connected!")
}

func databaseConnected(db *sql.DB) bool {
	err := db.Ping()
	if err != nil {
		log.Fatalln("Not response of database")
		return false
	}
	return true
}
