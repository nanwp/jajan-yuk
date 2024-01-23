package helper

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadImage(image multipart.File, handler *multipart.FileHeader) (string, error) {
	dir := "images"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	// generate file name
	fileName := time.Now().Format("20060102150405") + filepath.Ext(handler.Filename)

	// create file
	file, err := os.Create(dir + "/" + fileName)
	if err != nil {
		return "", err
	}

	// copy image to file
	_, err = io.Copy(file, image)
	if err != nil {
		return "", err
	}

	fullImagePath, err := filepath.Abs(dir + "/" + fileName)

	return fullImagePath, nil
}

func DeleteImage(fileName string) error {

	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}
