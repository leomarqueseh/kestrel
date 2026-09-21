package target

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrValueRequired = errors.New("target: value is required")
	ErrInvalidType   = errors.New("target: target_type must be one of domain, ip, url, cidr")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, projectID, value string, targetType Type, description string) (*Target, error) {
	if strings.TrimSpace(value) == "" {
		return nil, ErrValueRequired
	}
	if !targetType.Valid() {
		return nil, ErrInvalidType
	}
	return s.repo.Create(ctx, projectID, value, targetType, description)
}

func (s *Service) ListByProject(ctx context.Context, projectID string) ([]Target, error) {
	return s.repo.ListByProject(ctx, projectID)
}

func (s *Service) Get(ctx context.Context, id string) (*Target, error) {
	return s.repo.GetByID(ctx, id)
}

// Authorize is the scope-enforcement gate: a target only becomes usable by
// future scan modules after this explicit call succeeds.
func (s *Service) Authorize(ctx context.Context, id string) (*Target, error) {
	return s.repo.Authorize(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
