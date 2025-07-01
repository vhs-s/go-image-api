package image_file

import (
	imagefile "go-image-api/internal/entities"
	imagemeta "go-image-api/internal/repositories/image_meta"
)

type imageRepository interface {
	Create(*imagefile.Image) error
	GetById(string) (*imagemeta.Imagemeta, error)
	DeleteById(string) error
}
