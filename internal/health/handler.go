package health

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", h.statusOk)
	})
}

func (h *Handler) statusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
