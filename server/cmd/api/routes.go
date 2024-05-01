package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	// mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Route("/api/student", func(r chi.Router) {
		r.Post("/", app.CreateStudent)
		r.Get("/", app.GetStudentByName)
		r.Get("/{sid}", app.GetStudentById)
	})

	// // CREATE NEW STUDENT
	// mux.Post("/api/student", app.CreateStudent)
	// // SEARCH STUDENT
	// mux.Get("/api/student/{sid}", app.GetStudentById)
	// mux.Get("/api/student", app.GetStudentByName)
	return mux
}
