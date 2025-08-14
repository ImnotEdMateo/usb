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
	err := filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}

 		// If it's a file, calculate its hash
		if !info.IsDir() {
			existingFile, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return nil
			}
			defer existingFile.Close()

			existingHash, err := CalculateFileHash(existingFile)
			if err != nil {
				fmt.Printf("Error calculating hash for file %q: %v\n", path, err)
				return nil 
			}

			// Compare hashes
			if hash == existingHash {
				return fmt.Errorf("hash found")
			}
		}
		return nil
	})

	return err != nil && err.Error() == "hash found"
}
