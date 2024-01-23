package services

import (
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

type AssetService struct {
	logger *slog.Logger
}

func (as *AssetService) SaveFile(file multipart.File, handler *multipart.FileHeader, path string) (string, error) {
	// Save the file
	fileUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	fileExt := strings.ToLower(strings.TrimPrefix(handler.Filename, "."))
	fileName := fileUUID.String() + "." + fileExt

	savePath := path + fileName

	newFile, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	// Copy the file to the destination
	_, err = io.Copy(newFile, file)
	if err != nil {
		return "", err
	}

	return savePath, nil
}
