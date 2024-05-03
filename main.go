package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"rest_api_go/internal/config"
	"rest_api_go/internal/storage"
	"rest_api_go/routes"
)

func main() {
	config.InitViperEnv()
	storage.ConnectDb()
	r := routes.SetupRouter()
	start(r)

}

func start(r *chi.Mux) {
	port := viper.GetString("PORT")
	srv := &http.Server{Addr: ":" + port, Handler: r}

	fmt.Println("Server started on port", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
