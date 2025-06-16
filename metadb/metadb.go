package metadb

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = nil
		_, err = os.Create(path)
		if err != nil {
			log.Println("Failed to create file at path:", path, "ERROR:", err)
		}
	} else {
		log.Println("File", path, "already exists")
	}
}

func CreateTable(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Println("An error occurred while creating the table:", err)
	}
}
