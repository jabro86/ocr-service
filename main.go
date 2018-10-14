package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jabro86/go-webserver"
	"github.com/jabro86/ocr-service/controllers"
	"github.com/jabro86/ocr-service/filters"
)

var logger *log.Logger

func main() {

	webserver.LoadViews("./app/views")

	r := webserver.NewRouter()
	// API
	r.GET("/status", controllers.Status)
	r.POST("/base64", controllers.Base64)
	r.POST("/file", controllers.FileUpload)
	// Sample Page
	r.GET("/", controllers.Index)
	r.Static("/assets", "./app/assets")

	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", "ocrserver"), 0)
	r.Apply(&filters.LogFilter{Logger: logger})

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatalln("Required env `PORT` is not specified.")
	}
	logger.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Println(err)
	}
}
