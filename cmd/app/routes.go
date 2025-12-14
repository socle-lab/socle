package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	socle "github.com/socle-lab/core"
	"github.com/socle-lab/core/pkg/env"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	// middleware must come before any routes
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "*")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", func(r chi.Router) {
		// add routes here
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(`{"data":"API 📺 Up and Running"}`))
		// })
		r.Get("/", app.Handler.HomeHandler)

	})
	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from socle
	r.Mount("/socle", socle.Routes())

	// routes from web inner admin
	r.Mount("/admin-console", app.AdminRoutes())

	// routes from web inner api
	//r.Mount("/api", app.ApiRoutes())
	return r
}
