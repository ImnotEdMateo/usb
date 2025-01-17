package handlers

import (
  "fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
  "strconv"
	
  "github.com/imnotedmateo/usb/config"
)

func FileOrPageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if path == "" {
		WebPageHandler(w, r)
		return
	}

	var validPathPattern string
	if config.RandomPath == "GUID" {
		validPathPattern = `^[a-f0-9\-]{36}$`
	} else {
		numChars, _ := strconv.Atoi(config.RandomPath)
		validPathPattern = fmt.Sprintf(`^[a-zA-Z0-9]{%d}$`, numChars)
	}

	matched, err := regexp.MatchString(validPathPattern, path)
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
