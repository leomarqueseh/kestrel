package project

import (
	"context"
	"errors"
	"strings"
)

var ErrNameRequired = errors.New("project: name is required")

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, description, ownerID string) (*Project, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrNameRequired
	}
	return s.repo.Create(ctx, name, description, ownerID)
}

func (s *Service) ListByOwner(ctx context.Context, ownerID string) ([]Project, error) {
	return s.repo.ListByOwner(ctx, ownerID)
}
