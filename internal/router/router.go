package router

import (
	"warehouse-app/internal/category"

	"github.com/go-chi/chi/v5"
)

func Setup(handler *category.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.GetByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})

	return r
}
