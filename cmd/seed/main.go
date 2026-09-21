// Command seed creates the first admin user. Run once after migrations:
//   go run ./cmd/seed -email admin@kestrel.local -password change-me
package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"

	"github.com/leomarqueseh/kestrel/internal/auth"
	"github.com/leomarqueseh/kestrel/internal/config"
	"github.com/leomarqueseh/kestrel/internal/platform/postgres"
)

func main() {
	email := flag.String("email", "", "admin email")
	password := flag.String("password", "", "admin password")
	flag.Parse()

	if *email == "" || *password == "" {
		log.Fatal("usage: go run ./cmd/seed -email <email> -password <password>")
	}

	_ = godotenv.Load()
	cfg := config.Load()
	ctx := context.Background()

	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()

	hash, err := auth.HashPassword(*password)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	user, err := auth.NewPostgresRepository(pool).Create(ctx, *email, hash, auth.RoleAdmin)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	log.Printf("admin user created: %s (role: %s)", user.Email, user.Role)
}
