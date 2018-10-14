package controllers

import (
	"net/http"

	"github.com/jabro86/go-webserver"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	webserver.Render(w).HTML("index", map[string]interface{}{
		"AppName": "ocrserver",
	})
}
