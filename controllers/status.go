package controllers

import (
	"net/http"

	"github.com/jabro86/go-webserver"
)

const version = "0.2.0"

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	webserver.Render(w, true).JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello!",
		"version": version,
	})
}
