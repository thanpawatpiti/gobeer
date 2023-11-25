package helpers

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	// Get file name
	fileName := file.Filename

	// Get file extension
	fileExtension := filepath.Ext(fileName)

	// Create random file name
	randomFileName := uuid.New().String()

	// Create new file name
	newFileName := randomFileName + fileExtension

	// Get file
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	// Create new file
	dst, err := os.Create("./public/images/" + newFileName)
	if err != nil {
		return "", err
	}

	// Copy file
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Close file
	defer dst.Close()
	defer src.Close()

	return "/public/images/" + newFileName, nil
}
