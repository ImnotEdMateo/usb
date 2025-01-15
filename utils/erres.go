package utils

import (
	"errors"
	"fmt"
	"net/http"
)

func ExtremelySeriousErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	http.ServeFile(w, r, "static/assets/owtffd.webp")
}

func HandleError(w http.ResponseWriter, r *http.Request, errMsg string) {
	fmt.Println("Error:", errMsg)
	ExtremelySeriousErrorResponse(w, r, errors.New(errMsg))
}
