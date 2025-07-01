package image_file

import (
	imagefile "go-image-api/internal/entities"
	"io"
	"log"
	"net/http"
)

func (ir *ImageHandler) UploadImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
		buf, err := io.ReadAll(file)
		if err != nil {
			log.Println("ERROR to read file", err)
			http.Error(w, "Failed to upload image", http.StatusBadRequest)
			return
		}
		img, err := imagefile.New(int(h.Size), h.Filename, h.Header.Get("Content-Type"), buf)
		if err != nil {
			http.Error(w, "Wrong File", http.StatusBadRequest)
		}

		if !img.CheckRestrictions() {
			http.Error(w, "Invalid file: too large or invalid format.", http.StatusBadRequest)
			return
		}

		err = img.Save()
		if err != nil {
			log.Println("Failed to create uploaded image:", err)
			http.Error(w, "Failed to upload image", http.StatusBadRequest)
			return
		}

		err = ir.ImageRepo.Create(img)
		if err != nil {
			log.Println("Error writing to database:", err)
			http.Error(w, "Failed to upload image", http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Image uploaded"))
		log.Println("Image uploaded:", img.Name, img.Format, img.Size, img.Hash)

	}
}
