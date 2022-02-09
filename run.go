package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var BuildVersion string = "0.0.0"

func loadRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/health", statusHandler)
	return router
}

func main() {
	Run(loadRoutes())
}
