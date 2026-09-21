// Package auth handles user authentication (login, tokens) and
// authorization (role-based access control).
package auth

import "time"

// Role represents what a user is allowed to do in Kestrel.
type Role string

const (
	RoleAdmin   Role = "admin"
	RoleAnalyst Role = "analyst"
	RoleViewer  Role = "viewer"
)

// User mirrors the users table. PasswordHash is never exposed in JSON
// (the json:"-" tag excludes it from any response).
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
