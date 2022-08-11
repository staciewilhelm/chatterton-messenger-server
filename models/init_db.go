package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

func InitializeDB() error {
	var (
		dbPort   = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_DATABASE")
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbPort, username, password, dbName)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Println("Error pinging database", err)
		return err
	}

	fmt.Println("Database successfully connected!")

	DB = conn
	return nil
}
