package finding

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Assess(w http.ResponseWriter, r *http.Request) {
	targetID := chi.URLParam(r, "targetID")

	findings, err := h.service.Assess(r.Context(), targetID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "assessment failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, findings)
}

func (h *Handler) ListByTarget(w http.ResponseWriter, r *http.Request) {
	targetID := chi.URLParam(r, "targetID")

	findings, err := h.service.ListByTarget(r.Context(), targetID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list findings")
		return
	}
	writeJSON(w, http.StatusOK, findings)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
