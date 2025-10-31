package api

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/download", handleDownload)
	r.Get("/history", handleHistory)

	return r
}
