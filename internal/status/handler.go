package status

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
	router.Get("/status", h.statusOk)
}

func (h *Handler) statusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
