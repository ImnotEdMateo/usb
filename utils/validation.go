package utils

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"path/filepath"
)

func ValidateFile(file *os.File, filename string, maxSize int64) error {
	// Checks the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() > maxSize {
		return fmt.Errorf("file too large")
	}

	// Checks the file extension
	ext := filepath.Ext(filename)
	if ext == ".exe" || ext == ".sh" || ext == ".bat" || ext == ".apk" {
		return fmt.Errorf("file type not allowed")
	}

	// Checks the actual MIME type by reading the content
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

	return nil
}
