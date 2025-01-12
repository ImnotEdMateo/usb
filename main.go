package main

import (
  "net/http"
  "github.com/imnotedmateo/ubs/handlers"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", handlers.WebPageHandler)

  http.ListenAndServe(":8000", nil)
}
