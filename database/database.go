package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {
	pg, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to Database!")
	log.Println("Creating schemas!")
	createSchemas(pg)
	return pg
}

func createSchemas(db *sql.DB) {
	_, err := db.Exec(createProducts)
	if err != nil {
		panic(err)
	}
}
