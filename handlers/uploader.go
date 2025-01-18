package handlers

import (
  "io"
  "os"
  "log"
	"net/http"

	"github.com/imnotedmateo/usb/utils"
	"github.com/imnotedmateo/usb/config"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(w, r, "Method not Allowed")
		return
	}

	// Get the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.HandleError(w, r, "Error retrieving file from the form")
		return
	}
	defer file.Close()

	utils.LogUpload(r, header.Filename)

	// Temporarily save the file in the system
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

	// Validate the file
	err = utils.ValidateFile(tempFile, header.Filename, config.MaxFileSize, "uploads")
	if err != nil {
		utils.HandleError(w, r, err.Error())
		return
	}

	// Save the file using the modularized function
	dirPath, err := utils.SaveUploadedFile(tempFile, header.Filename)
	if err != nil {
		utils.HandleError(w, r, err.Error())
		return
	}

	http.Redirect(w, r, "/"+dirPath+"/", http.StatusSeeOther)
}
