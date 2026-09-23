// Package report consolidates a project's scope, findings, and evidence
// into an exportable assessment report.
package report

import "time"

type Format string

const (
	FormatJSON Format = "json"
	FormatHTML Format = "html"
	FormatPDF  Format = "pdf"
)

// Report is the persisted audit record of a generation event —
// not the rendered content itself, which is returned directly and not stored.
type Report struct {
	ID          string    `json:"id"`
	ProjectID   string    `json:"project_id"`
	Format      Format    `json:"format"`
	GeneratedAt time.Time `json:"generated_at"`
}
