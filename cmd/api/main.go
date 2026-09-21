package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/leomarqueseh/kestrel/internal/config"
	"github.com/leomarqueseh/kestrel/internal/health"
	"github.com/leomarqueseh/kestrel/internal/platform/postgres"
	"github.com/leomarqueseh/kestrel/internal/version"
)

func main() {
	// .env is optional in production (real env vars take over), but convenient
	// for local development — the error is intentionally ignored here.
	_ = godotenv.Load()

	cfg := config.Load()
	ctx := context.Background()

	dbPool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer dbPool.Close()

	healthHandler := health.NewHandler(health.NewService(dbPool))
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
