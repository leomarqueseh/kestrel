package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/leomarqueseh/kestrel/internal/config"
	"github.com/leomarqueseh/kestrel/internal/health"
	"github.com/leomarqueseh/kestrel/internal/version"
)

func main() {
	cfg := config.Load()

	healthHandler := health.NewHandler(health.NewService())
	versionHandler := version.NewHandler(version.NewService())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthHandler.Check)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/version", versionHandler.Get)
	})

	addr := ":" + cfg.Port
	log.Printf("kestrel-api listening on %s (env=%s)", addr, cfg.Env)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
