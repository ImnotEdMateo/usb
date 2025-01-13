package handlers

import (
	"net/http"
	"path/filepath"
)

var staticPath = filepath.Join("static", "index.html")

func WebPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticPath)
}
