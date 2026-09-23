package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"

	"github.com/leomarqueseh/kestrel/internal/asset"
	"github.com/leomarqueseh/kestrel/internal/auth"
	"github.com/leomarqueseh/kestrel/internal/config"
	"github.com/leomarqueseh/kestrel/internal/evidence"
	"github.com/leomarqueseh/kestrel/internal/finding"
	"github.com/leomarqueseh/kestrel/internal/health"
	"github.com/leomarqueseh/kestrel/internal/platform/postgres"
	"github.com/leomarqueseh/kestrel/internal/project"
	"github.com/leomarqueseh/kestrel/internal/scan"
	"github.com/leomarqueseh/kestrel/internal/target"
	"github.com/leomarqueseh/kestrel/internal/version"
)

func main() {
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

	tokenIssuer := auth.NewTokenIssuer(cfg.JWTSecret)
	authHandler := auth.NewHandler(auth.NewService(auth.NewPostgresRepository(dbPool), tokenIssuer))

	projectHandler := project.NewHandler(project.NewService(project.NewPostgresRepository(dbPool)))

	targetRepo := target.NewPostgresRepository(dbPool)
	targetHandler := target.NewHandler(target.NewService(targetRepo))

	assetRepo := asset.NewPostgresRepository(dbPool)
	scanRepo := scan.NewPostgresRepository(dbPool)
	scanHandler := scan.NewHandler(scan.NewService(scanRepo, assetRepo, targetRepo), assetRepo)

	evidenceRepo := evidence.NewPostgresRepository(dbPool)
	findingHandler := finding.NewHandler(finding.NewService(finding.NewPostgresRepository(dbPool), assetRepo, evidenceRepo))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthHandler.Check)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/version", versionHandler.Get)

		r.Route("/auth", func(r chi.Router) {
			r.With(httprate.LimitByIP(5, time.Minute)).Post("/login", authHandler.Login)
			r.Post("/refresh", authHandler.Refresh)
			r.Post("/logout", authHandler.Logout)
		})

		r.Group(func(r chi.Router) {
			r.Use(auth.RequireAuth(tokenIssuer))

			r.Get("/me", func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"message":"if you can read this, your token is valid"}`))
			})

			r.Route("/projects", func(r chi.Router) {
				r.Post("/", projectHandler.Create)
				r.Get("/", projectHandler.List)

				r.Route("/{projectID}/targets", func(r chi.Router) {
					r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/", targetHandler.Create)
					r.Get("/", targetHandler.ListByProject)
				})
			})

			r.Route("/targets/{targetID}", func(r chi.Router) {
				r.Get("/", targetHandler.Get)
				r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Delete("/", targetHandler.Delete)
				r.With(auth.RequireRole(auth.RoleAdmin)).Post("/authorize", targetHandler.Authorize)

				r.Get("/assets", scanHandler.ListAssetsByTarget)

				r.Route("/scans", func(r chi.Router) {
					r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/recon", scanHandler.RunRecon)
					r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/enumeration", scanHandler.RunEnumeration)
					r.Get("/", scanHandler.ListByTarget)
				})

				r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/assessment", findingHandler.Assess)
				r.Get("/findings", findingHandler.ListByTarget)
			})

			r.Get("/scans/{scanID}/assets", scanHandler.ListAssets)

			r.Route("/findings/{findingID}", func(r chi.Router) {
				r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/start-validation", findingHandler.StartValidation)
				r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/confirm", findingHandler.Confirm)
				r.With(auth.RequireRole(auth.RoleAdmin, auth.RoleAnalyst)).Post("/reject", findingHandler.Reject)
			})
		})
	})

	addr := ":" + cfg.Port
	log.Printf("kestrel-api listening on %s (env=%s)", addr, cfg.Env)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
