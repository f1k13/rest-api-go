package controllers

import (
	"github.com/go-chi/chi"
	"rest_api_go/handlers"
)

func UserController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/getUser", handlers.GetUser)
	return r
}
