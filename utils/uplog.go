package utils

import (
	"github.com/imnotedmateo/ubs/config"
	"log"
	"net"
	"net/http"
)

// GetClientIP extracts the client IP address from the HTTP request.
func GetClientIP(r *http.Request) string {
	// Check for the X-Forwarded-For header (useful if behind a reverse proxy)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}
	// Fall back to the remote address
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// LogUpload logs details about the file upload attempt with colors in the desired format.
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
