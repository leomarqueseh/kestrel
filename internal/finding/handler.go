package finding

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/leomarqueseh/kestrel/internal/evidence"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Assess(w http.ResponseWriter, r *http.Request) {
	findings, err := h.service.Assess(r.Context(), chi.URLParam(r, "targetID"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "assessment failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, findings)
}

func (h *Handler) ListByTarget(w http.ResponseWriter, r *http.Request) {
	findings, err := h.service.ListByTarget(r.Context(), chi.URLParam(r, "targetID"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list findings")
		return
	}
	writeJSON(w, http.StatusOK, findings)
}

func (h *Handler) StartValidation(w http.ResponseWriter, r *http.Request) {
	f, err := h.service.StartValidation(r.Context(), chi.URLParam(r, "findingID"))
	h.respond(w, f, err)
}

type confirmRequest struct {
	Request  string `json:"request"`
	Response string `json:"response"`
	Notes    string `json:"notes"`
}

func (h *Handler) Confirm(w http.ResponseWriter, r *http.Request) {
	var req confirmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	f, err := h.service.Confirm(r.Context(), chi.URLParam(r, "findingID"), evidence.Evidence{
		Request:  req.Request,
		Response: req.Response,
		Notes:    req.Notes,
	})
	h.respond(w, f, err)
}

type rejectRequest struct {
	Reason string `json:"reason"`
}

func (h *Handler) Reject(w http.ResponseWriter, r *http.Request) {
	var req rejectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	f, err := h.service.Reject(r.Context(), chi.URLParam(r, "findingID"), req.Reason)
	h.respond(w, f, err)
}

func (h *Handler) respond(w http.ResponseWriter, f *Finding, err error) {
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeError(w, http.StatusNotFound, "finding not found")
			return
		}
		if errors.Is(err, ErrInvalidTransition) {
			writeError(w, http.StatusConflict, "invalid status transition for this finding's current state")
			return
		}
		writeError(w, http.StatusInternalServerError, "operation failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, f)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
