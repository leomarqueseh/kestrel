package report

import (
	"context"
	"time"

	"github.com/leomarqueseh/kestrel/internal/evidence"
	"github.com/leomarqueseh/kestrel/internal/finding"
	"github.com/leomarqueseh/kestrel/internal/project"
	"github.com/leomarqueseh/kestrel/internal/target"
)

// Data holds everything a rendered report needs, gathered once and reused
// across whichever output format was requested.
type Data struct {
	Project        project.Project
	GeneratedAt    time.Time
	Targets        []target.Target
	Findings       []finding.Finding
	EvidenceByID   map[string][]evidence.Evidence
	SeverityCounts map[finding.Severity]int
}

type Builder struct {
	projects project.Repository
	targets  target.Repository
	findings finding.Repository
	evidence evidence.Repository
}

func NewBuilder(projects project.Repository, targets target.Repository, findings finding.Repository, evidenceRepo evidence.Repository) *Builder {
	return &Builder{projects: projects, targets: targets, findings: findings, evidence: evidenceRepo}
}

func (b *Builder) Build(ctx context.Context, projectID string) (*Data, error) {
	proj, err := b.projects.GetByID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	targets, err := b.targets.ListByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	findings, err := b.findings.ListByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	evidenceByID := make(map[string][]evidence.Evidence)
	severityCounts := make(map[finding.Severity]int)

	for _, f := range findings {
		severityCounts[f.Severity]++

		// Evidence is only meaningful once a finding has been reviewed —
		// fetching it for every "detected" finding would be wasted queries.
		if f.Status == finding.StatusConfirmed || f.Status == finding.StatusFalsePositive {
			if ev, err := b.evidence.ListByFinding(ctx, f.ID); err == nil {
				evidenceByID[f.ID] = ev
			}
		}
	}

	return &Data{
		Project:        *proj,
		GeneratedAt:    time.Now().UTC(),
		Targets:        targets,
		Findings:       findings,
		EvidenceByID:   evidenceByID,
		SeverityCounts: severityCounts,
	}, nil
}
