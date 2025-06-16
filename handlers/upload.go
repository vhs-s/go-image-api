package handlers

import (
	"go-image-api/imagefile"
	"go-image-api/metadb"
	"log"
	"net/http"
	"strconv"
	"time"
)

const TimeFormat string = "2006-01-02 15:04:05"

func (a *App) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, h, err := r.FormFile("image")
	if err != nil {
		log.Println("Failed to get info FormFile:", err)
		http.Error(w, "Failed to upload image", http.StatusBadRequest)
		return
	}
	defer file.Close()
	if !imagefile.CheckRestrictions(int(h.Size), imagefile.GetFileFormat(h.Filename)) {
		http.Error(w, "Invalid file: too large or invalid format.", http.StatusBadRequest)
		return
	}

	err = imagefile.CreateUploadedImage(file, imagefile.HashName(imagefile.GetName(h.Filename), strconv.Itoa(metadb.GetMaxId(a.DB)+1)), imagefile.GetFileFormat(h.Filename))
	if err != nil {
		log.Println("Failed to create uploaded image:", err)
		http.Error(w, "Failed to upload image", http.StatusBadRequest)
		return
	}

	img := imagefile.NewImage(metadb.GetMaxId(a.DB)+1, int(h.Size), imagefile.GetName(h.Filename), imagefile.GetFileFormat(h.Filename), h.Header.Get("Content-Type"), time.Now().Format("2006-01-02 15:04:05"))

	err = metadb.CreateMetaRow(a.DB, img)
	if err != nil {
		log.Println("Error writing to database:", err)
		http.Error(w, "Failed to upload image", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image uploaded"))
	log.Println("Image uploaded:", img)
}
