package storage

import (
	"io"
	"mime/multipart"
	"os"
)

const (
	Prefix string = "public/"
)

func LocalPut(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(Prefix + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}  