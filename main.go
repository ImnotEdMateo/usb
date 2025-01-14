package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/imnotedmateo/ubs/handlers"
)

func main() {
	port := os.Getenv("UBS_PORT")
	if port == "" {
		log.Fatalf("PORT is not defined")
	}

	fmt.Println("Running server...")

	// Serve Static Files
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.FileOrPageHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
