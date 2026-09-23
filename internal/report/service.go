package report

import (
	"context"
	"encoding/json"
	"errors"
)

// PDF isn't wired yet — it needs a rendering engine (headless Chrome or
// wkhtmltopdf) not present in this environment. Deliberate MVP scope, not
// an oversight: HTML already prints cleanly to PDF from any browser.
var ErrUnsupportedFormat = errors.New("pdf export is not implemented yet — use format=html and print to PDF from your browser")

type Service struct {
	builder *Builder
	repo    Repository
}

func NewService(builder *Builder, repo Repository) *Service {
	return &Service{builder: builder, repo: repo}
}

func (s *Service) Generate(ctx context.Context, projectID string, format Format) ([]byte, string, error) {
	data, err := s.builder.Build(ctx, projectID)
	if err != nil {
		return nil, "", err
	}

	var content []byte
	var contentType string

	switch format {
	case FormatJSON:
		content, err = json.MarshalIndent(data, "", "  ")
		contentType = "application/json"
	case FormatHTML:
		content, err = RenderHTML(data)
		contentType = "text/html"
	default:
		return nil, "", ErrUnsupportedFormat
	}
	if err != nil {
		return nil, "", err
	}

	if _, err := s.repo.Create(ctx, projectID, format); err != nil {
		return nil, "", err
	}

	return content, contentType, nil
}
