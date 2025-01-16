package handlers

import (
	"fmt"
	"net/http"
  
  "github.com/imnotedmateo/usb/utils"
  "github.com/imnotedmateo/usb/config"
)

func WebPageHandler(w http.ResponseWriter, r *http.Request) {
  theme := config.Theme
  maxFileSizeReadable := utils.BytesToHumanReadable(config.MaxFileSize)

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="es">
	  <head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1">
	    <title>USB by edmateo.site</title>
	    <link href="/static/index.css" rel="stylesheet">
	    <link href="/static/themes/%s" rel="stylesheet">
	    <link rel="icon" type="image/x-icon" href="/static/assets/favicon.ico">
	  </head>
	  <body>
	    <div class="wrapper">
	      <h1><span class="initial">U</span>pload <span class="initial">S</span>erver for <span class="initial">B</span>ullshit</h1> 
        <p>Temporary uploads less than %s. Made by <a href="http://edmateo.site">edmateo.site</a></p>
	      <div class="form-shit">
	        <form action="/upload" method="post" enctype="multipart/form-data">
	          <input type="file" name="file" id="file" required>
	          <input type="submit" name="send" value="UPLOAD">
	        </form>
	      </div>
	      <p>
	        <b>
	          <a href="https://github.com/ImnotEdMateo/ubs">SOURCE CODE</a>
	        </b>
	      </p>
	    </div>
	  </body>
	</html>`, theme, maxFileSizeReadable)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
