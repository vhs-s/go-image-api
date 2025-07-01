package image_file

type ImageHandler struct {
	ImageRepo imageRepository
}

func New(ImageRepo imageRepository) *ImageHandler {
	return &ImageHandler{ImageRepo: ImageRepo}
}
