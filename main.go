package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/imnotedmateo/ubs/handlers"
)

func main() {
	fmt.Println("Iniciando la aplicación...")

	// Sirve archivos estáticos
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.FileOrPageHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	fmt.Println("Aplicación ejecutada correctamente")

	// Escuchar y servir el servidor con manejo de errores
	if err := http.ListenAndServe(":1488", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
