package main

import (
	"chrome-nnwallet-server/internal/status"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	// Set-up Route
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// CORS
	//router.Use(cors.Handler(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
	//	ExposedHeaders:   []string{"Content-Type", "JWT-Token"},
	//	AllowCredentials: false,
	//	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//}))

	// Handlers
	statusHandler := status.NewHandler()
	statusHandler.SetupRoutes(router)

	// Start App
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic("ERROR: Something wrong with start app!")
	}
}
