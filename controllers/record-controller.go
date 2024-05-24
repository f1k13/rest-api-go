package controllers

import (
	"github.com/go-chi/chi"
	"rest_api_go/handlers"
)

func RecordController() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", handlers.CreateRec)
	r.Get("/getCat", handlers.GetCat)
	r.Get("/getByAcc", handlers.GetCreatingRec)
	return r
}
