package utils

import (
	"github.com/imnotedmateo/usb/config"
	"log"
	"net/http"
)

func LogUpload(r *http.Request, filename string) {
  clientIP := GetClientIP(r)

	// Apply color to the output
	ipColor := "\033[32m"   // Green color for IP
  fileColor := "\033[35m" // Purple color for Filename
	resetColor := "\033[0m" // Reset color

	if config.Doxxing {
		log.Printf("Attempt to upload File from IP: %s%s%s | Filename: %s%s%s\n",
			ipColor, clientIP, resetColor,
			fileColor, filename, resetColor)
	} else {
		log.Printf("Attempt to upload File: %s%s%s\n",
			fileColor, filename, resetColor)
  }
}
