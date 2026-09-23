package scan

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/leomarqueseh/kestrel/internal/asset"
)

type Handler struct {
	service *Service
	assets  asset.Repository
}

func NewHandler(service *Service, assets asset.Repository) *Handler {
	return &Handler{service: service, assets: assets}
}

func (h *Handler) RunRecon(w http.ResponseWriter, r *http.Request) {
	sc, err := h.service.RunRecon(r.Context(), chi.URLParam(r, "targetID"))
	h.respondScan(w, sc, err)
}

func (h *Handler) RunEnumeration(w http.ResponseWriter, r *http.Request) {
	sc, err := h.service.RunEnumeration(r.Context(), chi.URLParam(r, "targetID"))
	h.respondScan(w, sc, err)
}

func (h *Handler) respondScan(w http.ResponseWriter, sc *Scan, err error) {
	if err != nil {
		if errors.Is(err, ErrTargetNotAuthorized) {
			writeError(w, http.StatusForbidden, "target is not authorized — an admin must approve it first")
			return
		}
		writeError(w, http.StatusInternalServerError, "scan failed: "+err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, sc)
}

func (h *Handler) ListByTarget(w http.ResponseWriter, r *http.Request) {
	scans, err := h.service.ListByTarget(r.Context(), chi.URLParam(r, "targetID"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list scans")
		return
	}
	writeJSON(w, http.StatusOK, scans)
}

func (h *Handler) ListAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := h.assets.ListByScan(r.Context(), chi.URLParam(r, "scanID"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list assets")
		return
	}
	writeJSON(w, http.StatusOK, assets)
}

// ListAssetsByTarget returns the target's full Attack Surface Inventory —
// every asset discovered across every scan module ever run against it.
func (h *Handler) ListAssetsByTarget(w http.ResponseWriter, r *http.Request) {
	assets, err := h.assets.ListByTarget(r.Context(), chi.URLParam(r, "targetID"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list assets")
		return
	}
	writeJSON(w, http.StatusOK, assets)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
