package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CalculateFileHash(file *os.File) (string, error) {
	hasher := sha256.New()
	if _, err := file.Seek(0, io.SeekStart); err != nil { // Reset the file pointer
		return "", fmt.Errorf("error resetting the file: %v", err)
	}
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("error calculating file hash: %v", err)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil { // Reset again for subsequent reads
		return "", fmt.Errorf("error resetting the file after hashing: %v", err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func CheckHashExists(hash string, uploadDir string) bool {
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return false
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(uploadDir, file.Name())
			existingFile, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				continue
			}
			defer existingFile.Close()

			existingHash, err := CalculateFileHash(existingFile)
			if err != nil {
				fmt.Printf("Error calculating hash for existing file: %v\n", err)
				continue
			}

			if hash == existingHash {
				return true
			}
		}
	}
	return false
}
