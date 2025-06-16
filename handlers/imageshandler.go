package handlers

import (
	"go-image-api/imagefile"
	"go-image-api/metadb"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (a App) ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodGet) && (r.Method != http.MethodDelete) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlSubs := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlSubs[len(urlSubs)-1])
	if err != nil {
		log.Println("Error converting string:", err)
		http.Error(w, "Failed to get image", http.StatusBadRequest)
		return
	}

	img, err := metadb.GetImageWithId(a.DB, id)
	if err != nil {
		http.Error(w, "Failed to get image", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		filename := imagefile.HashName(img.Name, strconv.Itoa(id)) + img.Format
		file, err := os.Open("uploads/" + filename)
		if err != nil {
			log.Println("Failed to open file named:", filename)
			http.Error(w, "Failed to get image", http.StatusNotFound)
			return
		}
		defer file.Close()
		originalfilename := img.Name + img.Format
		w.Header().Set("Content-Disposition", "attachment; filename=\""+originalfilename+"\"")
		w.Header().Set("Content-Type", img.MimeType)
		io.Copy(w, file)
	case http.MethodDelete:
		filename := imagefile.HashName(img.Name, strconv.Itoa(id)) + img.Format
		os.Remove("uploads/" + filename)
		_ = metadb.DeleteImageWithId(a.DB, id)
		w.Write([]byte("The image has been successfully deleted."))
	}
}
