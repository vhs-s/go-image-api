package main

import (
	"go-image-api/internal/database"
	"go-image-api/internal/handlers/image_file"
	imagemeta "go-image-api/internal/repositories/image_meta"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	databasePath  string = "database.db"
	sqlDriver     string = "sqlite3"
	ServerAddress string = ":8080"
)

func main() {
	db := database.NewCon(databasePath)
	imagerepo := imagemeta.New(db)
	app := image_file.New(imagerepo)

	// defer database.Close()

	http.HandleFunc("/api/upload", app.UploadImageHandler())
	http.HandleFunc("/api/images/", app.ImagesHandler())
	err := http.ListenAndServe(ServerAddress, nil)
	if err != nil {
		log.Fatal("Error when starting the server:", err)
	}

}
