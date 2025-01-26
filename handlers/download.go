package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/download/"):]
	if path == "" {
		http.NotFound(w, r)
		return
	}

	dirPath := filepath.Join("uploads", path)
	
	dirInfo, err := os.Stat(dirPath)
	if os.IsNotExist(err) || !dirInfo.IsDir() {
		http.NotFound(w, r)
		return
	}

	files, err := os.ReadDir(dirPath)
	if err != nil || len(files) == 0 {
		http.NotFound(w, r)
		return
	}

	fileInfo, err := files[0].Info()
	if err != nil || !fileInfo.Mode().IsRegular() {
		http.NotFound(w, r)
		return
	}

	filePath := filepath.Join(dirPath, fileInfo.Name())
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileInfo.Name()))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}
