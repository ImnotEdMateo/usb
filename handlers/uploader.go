package handlers

import (
	"fmt"
  "os"
	"io"
  "time"
	"net/http"
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
  
  // Get the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
    handleError(w, r, "Error getting the file from the form")
		return
	}
	defer file.Close()

  // Checks if the file is to large
	if header.Size > maxFileSize {
    handleError(w, r, "The file exceeds the maximum allowed size")
		return
	}

  // Create unique file Route
  uniquePath, err := utils.GenerateRandomPath()
	if err != nil {
		http.Error(w, "Error ", http.StatusInternalServerError)
		return
	}

	// Create temporal file Route
	uploadPath := filepath.Join("uploads", uniquePath)
	dest, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "Error creating Route", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

  // Copy to /uploads/ dir
	_, err = io.Copy(dest, file)
	if err != nil {
    handleError(w, r, "Error saving the file")
		return
	}

  // Automatic deletion after one hour
	time.AfterFunc(1*time.Hour, func() {
		os.Remove(uploadPath)
	})

  http.Redirect(w, r, "/"+uniquePath, http.StatusSeeOther)
}
