package handlers

import (
	"project2/api/task"
	"project2/config/logger"

	chiMid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ApiHandler(chiMux *chi.Mux) {
	chiMux.Use(chiMid.StripSlashes)
	chiMux.Route("/tasks", func(r chi.Router) {

		r.Get("/", task.GetAlltask)
		r.Post("/", task.CreateTask)
		r.Get("/{taskId}", task.GetTask)
		r.Put("/{taskId}", task.UpdateTask)
		r.Delete("/{taskId}", task.DeleteTask)

		logger.Log(`
	API available:
	* GET		GetAllTasks		http://localhost:8080/tasks
	* GET 		GetTask			http://localhost:8080/tasks/{taskId}
	* POST 		CreateTask		http://localhost:8080/tasks
	* PUT 		UpdateTask		http://localhost:8080/tasks/{taskId}
	* DELETE	DeleteTask	 	http://localhost:8080/tasks/{taskId}



	Documentation http://localhost:8080/swagger/index.html
		`)
	})

	chiMux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))
}
