package helper

import (
	"fmt"
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

	fullPath := fmt.Sprintf("%s/%s", dir, fileName)

	return fullPath, nil
}

func DeleteImage(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func GetImage(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
