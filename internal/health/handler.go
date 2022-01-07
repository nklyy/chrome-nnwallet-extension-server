package health

import (
	"chrome-nnwallet-server/pkg/crypto"
	"chrome-nnwallet-server/pkg/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct {
}

type EncDecRequest struct {
	Value string `json:"value"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", h.statusOk)
		r.Post("/test/enc/dec", h.testEncryptionDecryption)
	})
}

func (h *Handler) statusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (h *Handler) testEncryptionDecryption(writer http.ResponseWriter, request *http.Request) {
	data := &EncDecRequest{}
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		helpers.Respond(writer, http.StatusInternalServerError, err)
		return
	}

	if data.Value == "" {
		helpers.Respond(writer, http.StatusInternalServerError, errors.New("Value does not valid!"))
		return
	}

	decrypted, err := crypto.Decrypt(data.Value)
	if err != nil {
		helpers.Respond(writer, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Decrypted: ", decrypted)

	encrypted, err := crypto.Encrypt(decrypted)
	if err != nil {
		helpers.Respond(writer, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Encrypted: ", encrypted)

	helpers.Respond(writer, http.StatusAccepted, encrypted)
}
