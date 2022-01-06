package main

import (
	"chrome-nnwallet-server/config"
	"chrome-nnwallet-server/internal/health"
	"chrome-nnwallet-server/pkg/crypto"
	"chrome-nnwallet-server/pkg/helpers"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

type PasswordRequest struct {
	Password string `json:"password"`
}

func main() {
	// Init config
	cfg, err := config.Get(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(cfg)

	// Set-up Route
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//router.Use(middleware.BasicAuth("authentication", map[string]string{cfg.User: cfg.Password}))

	// Handlers
	healthHandler := health.NewHandler()
	healthHandler.SetupRoutes(router)

	router.Post("/password", func(writer http.ResponseWriter, request *http.Request) {
		data := &PasswordRequest{}
		if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
			helpers.Respond(writer, http.StatusInternalServerError, err)
			return
		}

		decrypted, err := crypto.Decrypt(data.Password)
		if err != nil {
			helpers.Respond(writer, http.StatusInternalServerError, err)
			return
		}

		helpers.Respond(writer, http.StatusAccepted, decrypted)
	})

	// Start App
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		panic("ERROR: Something wrong with start app!")
	}
}
