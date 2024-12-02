package main

import (
	"net/http"
	"project2/config/logger"
	Handlers "project2/handlers"

	"github.com/go-chi/chi/v5"
)

// @title Project 2 on GoLangApi
// @version 1.0
// @description This is a sample project to demonstate CRUD ousing GoLand Rest API

// @contact.name Mohamed Farhaan
// @contact.email mfarhaan@presidio.com

// @host localhost:8080
// @BasePath /tasks
func main() {
	logger.ConfigureLogger()
	var chiMux *chi.Mux = chi.NewRouter()
	Handlers.ApiHandler(chiMux)
	Handlers.CronHandler()
	logger.Log(`Server starting at http://localhost:8080`)
	err := http.ListenAndServe("localhost:8080", chiMux)
	if err != nil {
		logger.Error(err)
	}
}
