package handlers

import (
  "os"
  "regexp"
	"net/http"
	"path/filepath"
)

func FileOrPageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if path == "" {
		WebPageHandler(w, r)
		return
	}

	matched, err := regexp.MatchString(`^[a-zA-Z0-9]{3}$`, path)
	if err != nil || !matched {
		http.NotFound(w, r)
		return
	}

	filePath := filepath.Join("uploads", path)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
