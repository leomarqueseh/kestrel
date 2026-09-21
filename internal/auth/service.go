package auth

import (
	"context"
	"errors"
)

var ErrInvalidCredentials = errors.New("auth: invalid email or password")

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Service struct {
	repo   Repository
	tokens *TokenIssuer
}

func NewService(repo Repository, tokens *TokenIssuer) *Service {
	return &Service{repo: repo, tokens: tokens}
}

func (s *Service) Login(ctx context.Context, email, password string) (*TokenPair, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		// Same error for "user not found" and "wrong password" — never reveal
		// which one it was, so an attacker can't enumerate valid emails.
		return nil, ErrInvalidCredentials
	}

	if !CheckPassword(user.PasswordHash, password) {
		return nil, ErrInvalidCredentials
	}

	return s.issueTokenPair(user)
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (*TokenPair, error) {
	claims, err := s.tokens.Parse(refreshToken)
	if err != nil || claims.TokenType != "refresh" {
		return nil, ErrInvalidToken
	}

	access, err := s.tokens.GenerateAccessToken(claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}

	// Simplification: reuses the same refresh token instead of rotating it.
	// Full rotation + revocation needs persistent storage — a Phase 19
	// hardening item, not required for this MVP.
	return &TokenPair{AccessToken: access, RefreshToken: refreshToken}, nil
}

func (s *Service) issueTokenPair(user *User) (*TokenPair, error) {
	access, err := s.tokens.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}
	refresh, err := s.tokens.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}
	return &TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}
