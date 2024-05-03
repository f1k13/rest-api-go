package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"rest_api_go/controllers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	auth := controllers.AuthController()
	users := controllers.UserController()
	r.Mount("/auth", auth)
	r.Mount("/users", users)
	return r
}
