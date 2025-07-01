package imagefile

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

const ValidSizeImage = 1048576

type Image struct {
	Size       int
	Name       string
	Format     string
	Hash       string
	Content    []byte
	MimeType   string
	UploadedAt string
}

const TimeFormat string = "2006-01-02 15:04:05"

func New(size int, name, mimeType string, content []byte) (*Image, error) {
	hash := uuid.New().String()
	subs := strings.Split(name, ".")
	if len(subs) != 2 {
		return nil, errors.New("invalid name")
	}
	uploadedAt := time.Now().Format(TimeFormat)
	return &Image{Size: size, Name: subs[0], Format: subs[1], MimeType: mimeType, UploadedAt: uploadedAt, Hash: hash, Content: content}, nil
}

func (i *Image) CheckRestrictions() bool {
	if i.Size > ValidSizeImage || (i.Format != "jpg" && i.Format != "jpeg") {
		return false
	}
	return true
}

func Open(hash string, format string, size int, originalname string, mimetype string, uploadedAt string) (*Image, error) {
	filename := hash + "." + format

	file, err := os.Open("uploads/" + filename)
	if err != nil {
		log.Println("Failed to open file named:", filename)
		return nil, errors.New("ERROR open file :" + filename)
	}
	defer file.Close()
	buf, _ := io.ReadAll(file)
	return &Image{
		Size:       size,
		Name:       originalname,
		Format:     format,
		Hash:       hash,
		Content:    buf,
		MimeType:   mimetype,
		UploadedAt: uploadedAt,
	}, nil
}

func (i *Image) Delete() error {
	filename := i.Hash + "." + i.Format
	err := os.Remove("uploads/" + filename)
	if err != nil {
		return errors.New("ERROR while deleting")
	}
	return nil
}

func (i *Image) Save() error {
	buf := i.Content //io.ReadAll(i.Content)

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
	out, err := os.Create("uploads/" + i.Hash + "." + i.Format)
	if err != nil {
		log.Println("Error creating file", i.Hash+"."+i.Format, ":", err)
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
