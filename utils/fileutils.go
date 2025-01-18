package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/imnotedmateo/usb/config"
)

func SaveUploadedFile(file *os.File, filename string) (string, error) {
	// Generate a unique path for the directory
	uniquePath, err := GenerateRandomPath()
	if err != nil {
		return "", fmt.Errorf("error generating unique directory path")
	}

	// Create the directory based on the unique path
	uploadDir := filepath.Join("uploads", uniquePath)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creating directory for upload")
	}

	// Create the final file inside the directory
	uploadPath := filepath.Join(uploadDir, filename)
	dest, err := os.Create(uploadPath)
	if err != nil {
		return "", fmt.Errorf("error creating final file")
	}
	defer dest.Close()

	// Copy the content from the temporary file to the final file
	if _, err := io.Copy(dest, file); err != nil {
		return "", fmt.Errorf("error saving final file")
	}

	// Schedule automatic deletion
	time.AfterFunc(config.FileExpirationTime, func() {
		if err := os.RemoveAll(uploadDir); err != nil {
			log.Printf("Error deleting directory: %v", err)
		} else {
			log.Printf("%s Deleted", uploadDir)
		}
	})

	log.Printf("File successfully uploaded to directory: %s", uniquePath)
	return uniquePath, nil
}
