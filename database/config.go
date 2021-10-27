package database

import "fmt"

var (
	dbUsername = "postgres"
	dbPassword = "postgres"
	dbHost     = "localhost"
	dbName     = "database"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbName, dbPassword)
)
