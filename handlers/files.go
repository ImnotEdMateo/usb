package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/imnotedmateo/usb/config"
)

func FileOrPageHandler(w http.ResponseWriter, r *http.Request) {
	// Get the directory path from the URL (removes the leading and trailing slashes)
	path := r.URL.Path[1:]
	if len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	// Show the main page if there is no path
	if path == "" {
		WebPageHandler(w, r)
		return
	}

	// Determine the valid pattern for the path
	var validPathPattern string
	if config.RandomPath == "GUID" {
		validPathPattern = `^[a-f0-9\-]{36}$`
	} else {
		numChars, _ := strconv.Atoi(config.RandomPath)
		validPathPattern = fmt.Sprintf(`^[a-zA-Z0-9]{%d}$`, numChars)
	}

	// Validate the path against the pattern
	matched, err := regexp.MatchString(validPathPattern, path)
	if err != nil || !matched {
		http.NotFound(w, r)
		return
	}

	// Build the full path to the directory
	dirPath := filepath.Join("uploads", path)

	// Check if the directory exists
	dirInfo, err := os.Stat(dirPath)
	if os.IsNotExist(err) || !dirInfo.IsDir() {
		http.NotFound(w, r)
		return
	}

	// Look for files inside the directory
	files, err := os.ReadDir(dirPath)
	if err != nil || len(files) == 0 {
		http.NotFound(w, r)
		return
	}

	// Take the first file inside the directory
	fileName := files[0].Name()
	filePath := filepath.Join(dirPath, fileName)

	// Check if the file exists and is a regular file
	fileInfo, err := os.Stat(filePath)
	if err != nil || fileInfo.IsDir() {
		http.NotFound(w, r)
		return
	}

	// Get the MIME type of the file
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		// Detect the MIME type if not found by extension
		mimeType = "application/octet-stream"
	}

	// Set headers
	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, fileName))
	w.Header().Set("Content-Type", mimeType)

	// If the file is not compatible with the browser, force download
	if mimeType == "application/octet-stream" {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	}

	http.ServeFile(w, r, filePath)
}
