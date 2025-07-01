package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewCon(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("Error opening database at path:", path, "ERROR:", err)
	}
	if err := database.Ping(); err != nil {
		panic("Cannot establish connection to db")
	}
	return database
}
