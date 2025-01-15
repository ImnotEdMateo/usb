package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func ValidateFile(file *os.File, filename string, maxSize int64, uploadDir string) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
  if fileInfo.Size() == 0 {
		return fmt.Errorf("File has nothing")
	}
	if fileInfo.Size() > maxSize {
		return fmt.Errorf("file too large")
	}

	// Checks the file extension
	ext := filepath.Ext(filename)
	if ext == ".exe" || ext == ".sh" || ext == ".bat" || ext == ".apk" {
		return fmt.Errorf("file type not allowed")
	}

	// Checks the MIME type by reading the content
	buffer := make([]byte, 512) // Sufficient size to detect MIME
	if _, err := file.Read(buffer); err != nil && err != io.EOF {
		return fmt.Errorf("error reading the file: %v", err)
	}

	// Resets the file position to allow subsequent reads
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("error resetting the file: %v", err)
	}

	// Detects the MIME type of the content
	mimeType := http.DetectContentType(buffer)
	if mimeType == "application/x-msdownload" || mimeType == "application/x-sh" {
		return fmt.Errorf("MIME type not allowed")
	}

	// Calculate and check the hash
	fileHash, err := CalculateFileHash(file)
	if err != nil {
		return fmt.Errorf("error calculating file hash: %v", err)
	}
	if CheckHashExists(fileHash, uploadDir) {
		return fmt.Errorf("duplicate file detected")
	}

	return nil
}
