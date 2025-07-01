package image_file

import (
	"bytes"
	imagefile "go-image-api/internal/entities"
	"io"
	"log"
	"net/http"
	"strings"
)

func (ir *ImageHandler) ImagesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method != http.MethodGet) && (r.Method != http.MethodDelete) {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		urlSubs := strings.Split(r.URL.Path, "/")
		id := urlSubs[len(urlSubs)-1]

		imgmeta, err := ir.ImageRepo.GetById(id)
		if err != nil {
			http.Error(w, "Failed to get image", http.StatusNotFound)
			return
		}
		image, err := imagefile.Open(imgmeta.Id, imgmeta.Format, imgmeta.Size, imgmeta.Name, imgmeta.MimeType, imgmeta.UploadedAt)
		if err != nil {
			log.Println("Failed to open file named:", imgmeta.Name+"."+imgmeta.Format)
			http.Error(w, "Failed to get image", http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			originalfilename := image.Name + "." + image.Format
			w.Header().Set("Content-Disposition", "attachment; filename=\""+originalfilename+"\"")
			w.Header().Set("Content-Type", image.MimeType)
			io.Copy(w, bytes.NewReader(image.Content))
		case http.MethodDelete:
			image.Delete()
			_ = ir.ImageRepo.DeleteById(image.Hash)
			w.Write([]byte("The image has been successfully deleted."))
		}
	}
}
