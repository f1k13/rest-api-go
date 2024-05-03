package controllers

import (
	"github.com/go-chi/chi"
	"rest_api_go/handlers"
)

func AuthController() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)
	return r
}
