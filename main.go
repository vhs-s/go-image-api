package main

import (
	"database/sql"
	"go-image-api/handlers"
	"go-image-api/metadb"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	databasePath  string = "metadb/database.db"
	sqlDriver     string = "sqlite3"
	ServerAddress string = ":8080"
)

func main() {
	metadb.CreateDB(databasePath)
	database, err := sql.Open(sqlDriver, databasePath)
	if err != nil {
		log.Fatal("Error opening database at path:", databasePath, "ERROR:", err)
	}
	metadb.CreateTable(database, metadb.Metatable)
	defer database.Close()
	app := &handlers.App{DB: database}

	http.HandleFunc("/api/upload", app.UploadImageHandler)
	http.HandleFunc("/api/images/", app.ImagesHandler)
	err = http.ListenAndServe(ServerAddress, nil)
	if err != nil {
		log.Fatal("Error when starting the server:", err)
	}

}
