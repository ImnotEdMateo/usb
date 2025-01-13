package handlers

import (
	"net/http"
	"path/filepath"
)

func WebPageHandler(w http.ResponseWriter, r *http.Request) {
	absPath, _ := filepath.Abs(filepath.Join("static", "index.html"))
	http.ServeFile(w, r, absPath)
}
