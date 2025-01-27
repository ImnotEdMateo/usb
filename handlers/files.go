package handlers

import (
	"fmt"
	"html/template"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/imnotedmateo/usb/config"
)

const downloadTemplateStr = `
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.FileName}} - USB by edmateo.site</title>
    <link href="/static/index.css" rel="stylesheet">
    <link href="/static/themes/{{.Theme}}" rel="stylesheet">
    <link rel="icon" type="image/x-icon" href="/static/assets/favicon.ico">
</head>
<body>
    <div class="wrapper">
        <h1><span class="initial">U</span>pload <span class="initial">S</span>erver for <span class="initial">B</span>ullshit</h1>
        <h2>File: {{.FileName}}</h2>
        <p class="disclaimer">This file cannot be displayed in the browser. You can download here:</p>
        <div class="form-shit">
          <form action="/download/{{.Path}}" method="get" class="download-form">
            <input type="submit" value="DOWNLOAD">
          </form>
        </div>
        <p>
            <b>
                <a href="https://github.com/ImnotEdMateo/ubs">SOURCE CODE</a>
            </b>
        </p>
    </div>
</body>
</html>
`

func FileOrPageHandler(w http.ResponseWriter, r *http.Request) {
	path := getSanitizedPath(r)
	if path == "" {
		WebPageHandler(w, r)
		return
	}

	if err := validatePath(path); err != nil {
		http.NotFound(w, r)
		return
	}

	dirPath := filepath.Join("uploads", path)
	if err := validateDirectory(dirPath); err != nil {
		http.NotFound(w, r)
		return
	}

	fileName, err := getFirstFileName(dirPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	mimeType := getMimeType(fileName)
	if !isBrowserSupported(mimeType) {
		renderDownloadTemplate(w, fileName, path)
		return
	}

	serveFileInline(w, r, filepath.Join(dirPath, fileName), fileName, mimeType)
}

func getSanitizedPath(r *http.Request) string {
	path := r.URL.Path[1:]
	if len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}

func validatePath(path string) error {
	pattern := getPathPattern()
	matched, err := regexp.MatchString(pattern, path)
	if err != nil || !matched {
		return fmt.Errorf("invalid path")
	}
	return nil
}

func getPathPattern() string {
	if config.RandomPath == "GUID" {
		return `^[a-f0-9\-]{36}$`
	}
	numChars, _ := strconv.Atoi(config.RandomPath)
	return fmt.Sprintf(`^[a-zA-Z0-9]{%d}$`, numChars)
}

func validateDirectory(dirPath string) error {
	dirInfo, err := os.Stat(dirPath)
	if err != nil || !dirInfo.IsDir() {
		return fmt.Errorf("invalid directory")
	}
	return nil
}

func getFirstFileName(dirPath string) (string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil || len(files) == 0 {
		return "", fmt.Errorf("no files found")
	}
	return files[0].Name(), nil
}

func getMimeType(fileName string) string {
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}

func renderDownloadTemplate(w http.ResponseWriter, fileName, path string) {
    tmpl, err := template.New("download").Parse(downloadTemplateStr)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }

    data := struct {
        FileName string
        Path     string
        Theme    string
    }{
        FileName: fileName,
        Path:     path,
        Theme:    config.Theme,
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    tmpl.Execute(w, data)
}

func serveFileInline(w http.ResponseWriter, r *http.Request, filePath, fileName, mimeType string) {
	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, fileName))
	w.Header().Set("Content-Type", mimeType)
	http.ServeFile(w, r, filePath)
}

func isBrowserSupported(mimeType string) bool {
	supportedTypes := map[string]bool{
		"text/html":              true,
		"text/plain":             true,
		"text/css":               true,
		"application/javascript": true,
		"image/jpeg":             true,
		"image/png":              true,
		"image/gif":              true,
		"image/svg+xml":          true,
		"application/pdf":        true,
		"video/mp4":              true,
		"audio/mpeg":             true,
		"audio/wav":              true,
	}
	return supportedTypes[mimeType]
}
