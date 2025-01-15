package utils

import (
  "net"
  "net/http"
)

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

