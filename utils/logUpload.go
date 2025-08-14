package utils

import (
	"log"
	"net/http"
)

func LogUpload(r *http.Request, filename string, doxxing bool) {
	clientIP := GetClientIP(r)

	// Apply color to the output
	ipColor := "\033[32m"   // Green
	fileColor := "\033[35m" // Purple
	resetColor := "\033[0m" // Reset

	if doxxing {
		log.Printf("Attempt to upload File from IP: %s%s%s | Filename: %s%s%s\n",
			ipColor, clientIP, resetColor,
			fileColor, filename, resetColor)
	} else {
		log.Printf("Attempt to upload File: %s%s%s\n",
			fileColor, filename, resetColor)
	}
}
