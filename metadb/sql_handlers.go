package metadb

import (
	"database/sql"
	"go-image-api/imagefile"
	"log"
)

func CreateMetaRow(db *sql.DB, imagefile *imagefile.Image) error {
	query := QueryInsertImageMeta
	_, err := db.Exec(query, imagefile.ID, imagefile.Name, imagefile.Format, imagefile.MimeType, imagefile.Size, imagefile.UploadedAt)
	if err != nil {
		log.Println("Error creating record in database:", err)
		return err
	}
	return nil
}

func GetMaxId(db *sql.DB) int {
	query := QueryGetMaxId
	row := db.QueryRow(query)
	var id sql.NullInt64
	err := row.Scan(&id)
	if err != nil {
		log.Println("Failed to Scan QueryRow:", err)
	}
	if !id.Valid {
		id.Int64 = 0
	}
	return int(id.Int64)
}

func GetImageWithId(db *sql.DB, id int) (*imagefile.Image, error) {
	query := QueryGetRowWithId
	row := db.QueryRow(query, id)
	img := imagefile.Image{}
	err := row.Scan(&img.ID, &img.Name, &img.Format, &img.MimeType, &img.Size, &img.UploadedAt)
	if err != nil {
		log.Println("Failed to Scan QueryRow:", err)
		return nil, err
	}
	return &img, nil
}

func DeleteImageWithId(db *sql.DB, id int) error {
	query := QueryDeleteRowWithId
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Failed to DELETE image with id:", id, "ERROR:", err)
		return err
	}
	return nil

}
