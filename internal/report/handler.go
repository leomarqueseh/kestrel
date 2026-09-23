package report

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Generate(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")

	format := Format(r.URL.Query().Get("format"))
	if format == "" {
		format = FormatJSON
	}

	content, contentType, err := h.service.Generate(r.Context(), projectID, format)
	if err != nil {
		status := http.StatusInternalServerError
		message := "report generation failed"
		if err == ErrUnsupportedFormat {
			status = http.StatusNotImplemented
			message = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(`{"error":"` + message + `"}`))
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
