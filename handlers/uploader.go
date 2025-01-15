package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/imnotedmateo/ubs/config"
	"github.com/imnotedmateo/ubs/utils"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(w, r, "Method not Allowed")
		return
	}

	// Gets the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.HandleError(w, r, "Error retrieving file from the form")
		return
	}
	defer file.Close()

	utils.LogUpload(r, header.Filename)

	// Temporarily saves the file in the system
	tempFile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		utils.HandleError(w, r, "Error creating temporary file")
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		utils.HandleError(w, r, "Error saving temporary file")
		return
	}

	// Validates the file
	if err := utils.ValidateFile(tempFile, header.Filename, config.MaxFileSize); err != nil {
		utils.HandleError(w, r, err.Error())
		return
	}

	// Generates a unique path for the file
	uniquePath, err := utils.GenerateRandomPath()
	if err != nil {
		utils.HandleError(w, r, "Error generating final file")
		return
	}

	// Creates the final file in the "uploads/" dir
	uploadPath := filepath.Join("uploads", uniquePath)
	dest, err := os.Create(uploadPath)
	if err != nil {
		utils.HandleError(w, r, "Error creating final file, maybe uploads/ does not exist.")
		return
	}
	defer dest.Close()

	// Copies the content of the temporary file to the final file
	if _, err := io.Copy(dest, tempFile); err != nil {
		utils.HandleError(w, r, "Error saving final file")
		return
	}

	// Schedules automatic deletion
	time.AfterFunc(config.FileExpirationTime, func() {
		os.Remove(uploadPath)
    log.Printf("%s Deleted", uploadPath)
	})

	http.Redirect(w, r, "/"+uniquePath, http.StatusSeeOther)
}
