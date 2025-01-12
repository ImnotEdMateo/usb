package handlers

import (
  "net/http"
)

func WebPageHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w , r, "static/index.html")
}
