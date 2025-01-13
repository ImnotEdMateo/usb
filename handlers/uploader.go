package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"github.com/imnotedmateo/ubs/utils"
)

const maxFileSize = 1 << 20

func handleError(w http.ResponseWriter, r *http.Request, errMsg string) {
	fmt.Println("Error:", errMsg) 
	utils.ExtremelySeriousErrorResponse(w, r, fmt.Errorf(errMsg))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, r, "Method Not Allowed")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
    handleError(w, r, "Error getting the file from the form")
		return
	}
	defer file.Close()

	if header.Size > maxFileSize {
    handleError(w, r, "The file exceeds the maximum allowed size")
		return
	}

	uniqueFilename := filepath.Join("uploads", header.Filename)
	dest, err := os.Create(uniqueFilename)
	if err != nil {
    handleError(w, r, "Error creating the destination file")
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
    handleError(w, r, "Error copying the file to the destination")
		return
	}

	w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "File successfully uploaded: %s\n", header.Filename)
}
