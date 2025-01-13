package utils

import (
  "fmt"
	"mime"
	"os"
	"path/filepath"
)

func ValidateFile(file *os.File, filename string, maxSize int64) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() > maxSize {
		return fmt.Errorf("archivo demasiado grande")
	}

	ext := filepath.Ext(filename)
	if ext == ".exe" || ext == ".sh" || ext == ".bat" {
		return fmt.Errorf("tipo de archivo no permitido")
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "application/x-msdownload" || mimeType == "application/x-sh" {
		return fmt.Errorf("tipo MIME no permitido")
	}

	return nil
}
