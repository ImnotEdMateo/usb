package main

import (
	"net/http"
  "github.com/imnotedmateo/ubs/handlers"
)

func main() {
	// Sirve archivos estáticos
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.FileOrPageHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	http.ListenAndServe(":1488", nil)
}
