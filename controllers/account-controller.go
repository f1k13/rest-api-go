package controllers

import (
	"github.com/go-chi/chi"
	"rest_api_go/handlers"
)

func AccountController() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", handlers.CreateAccount)
	r.Get("/getOfUser", handlers.GetAccOfUserId)
	return r
}
