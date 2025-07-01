package imagemeta

import (
	"database/sql"
	imagefile "go-image-api/internal/entities"
	"log"
)

type Imagemeta struct {
	Id         string
	Size       int
	Name       string
	Format     string
	MimeType   string
	UploadedAt string
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(imagefile *imagefile.Image) error {
	query := QueryInsertImageMeta
	_, err := r.db.Exec(query, imagefile.Hash, imagefile.Name, imagefile.Format, imagefile.MimeType, imagefile.Size, imagefile.UploadedAt)
	if err != nil {
		log.Println("Error creating record in database:", err)
		return err
	}
	return nil
}

func (r *repository) GetById(id string) (*Imagemeta, error) {
	query := QueryGetRowWithId
	row := r.db.QueryRow(query, id)
	imagemeta := Imagemeta{}
	err := row.Scan(&imagemeta.Id, &imagemeta.Name, &imagemeta.Format, &imagemeta.MimeType, &imagemeta.Size, &imagemeta.UploadedAt)
	if err != nil {
		log.Println("Failed to Scan QueryRow:", err)
		return nil, err
	}
	return &imagemeta, nil
}

func (r *repository) DeleteById(id string) error {
	query := QueryDeleteRowWithId
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Failed to DELETE image with id:", id, "ERROR:", err)
		return err
	}
	return nil

}
