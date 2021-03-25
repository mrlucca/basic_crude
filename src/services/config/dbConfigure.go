package config

import (
	"database/sql"
	"log"
)

func ConfigureDb() {
	db, _ := GetDbConnection()
	setSchema(db)
	setTables(db)
	setAdmin(db)
	defer db.Close()
}

func setSchema(db *sql.DB) {
	sqlStatement := "CREATE SCHEMA IF NOT EXISTS basic_crud;"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	log.Println("Set schema successful")
}

func setTables(db *sql.DB) {
	t1 := `
	CREATE TABLE basic_crud."book" (
		id SERIAL NOT NULL PRIMARY KEY,
		title VARCHAR(255),
		subtitle TEXT,
		price INTEGER,
		publication_date TIMESTAMPTZ,
		image BYTEA,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`
	t2 := `
	CREATE TABLE basic_crud."user" (
		id SERIAL NOT NULL PRIMARY KEY,
		"name" varchar NOT NULL,
		admin boolean NOT NULL, 
		nickname varchar NOT NULL,
		email text UNIQUE,
		passworld varchar NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	
	`
	tables := [2]string{t1, t2}

	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			panic(err)
		}
	}
	log.Println("Set tables successful")
}

func setAdmin(db *sql.DB) {
	sqlStatement := `
		INSERT INTO basic_crud."user" ("name", nickname, email, admin, passworld)
		VALUES('root', 'root', 'root@root.com', true, md5('masteruser'));
		`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	log.Println("Set admin successful")
}
