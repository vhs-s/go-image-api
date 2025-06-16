package imagefile

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

func HashName(name, salt string) string {
	hashedName := sha256.New()
	_, err := hashedName.Write([]byte(name + salt))
	if err != nil {
		log.Println("Error hashing file name:", err)
	}
	return fmt.Sprintf("%x", hashedName.Sum(nil))
}

func NewImage(id, size int, name, format, mimeType string, uploadedAt string) *Image {
	return &Image{ID: id, Size: size, Name: name, Format: format, MimeType: mimeType, UploadedAt: uploadedAt}
}

func CreateUploadedImage(file multipart.File, hashedname string, format string) error {
	buf, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading multipart.File:", err)
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		log.Println("Error while decoding:", err)
		return err
	}

	_, err = os.Stat("uploads/")
	if os.IsNotExist(err) {
		err = os.Mkdir("uploads", 0777)
		if err != nil {
			log.Println("Error creating uploads folder:", err)
			return err
		}
	}
	out, err := os.Create("uploads/" + hashedname + format)
	if err != nil {
		log.Println("Error creating file", hashedname+format, ":", err)
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println("Failed to encode image to JPEG:", err)
		return err
	}
	return nil
}

func CheckRestrictions(size int, format string) bool {
	if size > ValidSizeImage || (format != ".jpg" && format != ".jpeg") {
		return false
	}
	return true
}

func GetFileFormat(filename string) string {
	subs := strings.Split(filename, ".")
	return "." + subs[len(subs)-1]
}

func GetName(filename string) string {
	subs := strings.Split(filename, ".")
	return subs[0]
}
