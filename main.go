package main

import (
	"log"
	"net/http"
	"os"

	"github.com/imnotedmateo/usb/handlers"
)

func main() {
	port := os.Getenv("USB_PORT")
	if port == "" {
		log.Fatal("PORT is not defined")
	}

	log.Printf("Running server on http://0.0.0.0:%s", port)

  // serve static files
  http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
  
	http.HandleFunc("/", handlers.FileOrPageHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download/", handlers.DownloadHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
