package imagefile

const ValidSizeImage = 1048576

type Image struct {
	ID         int
	Size       int
	Name       string
	Format     string
	MimeType   string
	UploadedAt string
}
