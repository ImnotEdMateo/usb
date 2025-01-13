package main

import (
	"fmt"
  "log"
	"net/http"

	"github.com/imnotedmateo/ubs/handlers"
)

func main() {
	fmt.Println("Running server...")

	// Serve Static Files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.FileOrPageHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	if err := http.ListenAndServe(":1488", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
