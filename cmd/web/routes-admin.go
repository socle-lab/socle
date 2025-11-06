package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) AdminRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		// add any admin routes here
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			var payload struct {
				Content string `json:"content"`
			}
			payload.Content = "Hello, world"
			app.Core.WriteJSON(w, http.StatusOK, payload)
		})

	})

	return r
}
