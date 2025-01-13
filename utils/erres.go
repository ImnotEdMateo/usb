package utils

import (
	"net/http"
)

func ExtremelySeriousErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	http.ServeFile(w, r, "static/assets/owtffd.webp")
}
