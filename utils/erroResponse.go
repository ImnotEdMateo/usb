package utils

import (
	"fmt"
	"net/http"
)

func SeriousErrorResponse(w http.ResponseWriter, r *http.Request, errMsg string) {
	fmt.Println("Error:", errMsg)
  http.ServeFile(w, r, "static/assets/owtffd.webp")
}
