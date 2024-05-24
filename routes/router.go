package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"rest_api_go/controllers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	auth := controllers.AuthController()
	users := controllers.UserController()
	account := controllers.AccountController()
	income := controllers.RecordController()
	r.Mount("/auth", auth)
	r.Mount("/users", users)
	r.Mount("/account", account)
	r.Mount("/record", income)
	return r
}
