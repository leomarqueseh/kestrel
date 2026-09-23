package finding

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/leomarqueseh/kestrel/internal/asset"
	"github.com/leomarqueseh/kestrel/internal/evidence"
	"github.com/leomarqueseh/kestrel/internal/nvd"
)

var ErrInvalidTransition = errors.New("finding: invalid status transition")

type Service struct {
	repo     Repository
	assets   asset.Repository
	evidence evidence.Repository
}

func NewService(repo Repository, assets asset.Repository, evidenceRepo evidence.Repository) *Service {
	return &Service{repo: repo, assets: assets, evidence: evidenceRepo}
}

// --- Phase 08: assessment (unchanged logic, just carried over) ---

func (s *Service) Assess(ctx context.Context, targetID string) ([]Finding, error) {
	assets, err := s.assets.ListByTarget(ctx, targetID)
	if err != nil {
		return nil, err
	}

	var findings []Finding
	requestsMade := 0

	for _, a := range assets {
		if a.Protocol == "dns" {
			continue
		}

		identifier := softwareIdentifier(a)
		if identifier != "" && requestsMade > 0 {
			time.Sleep(6 * time.Second)
		}

		created, err := s.assessAsset(ctx, a, identifier)
		if identifier != "" {
			requestsMade++
		}
		if err != nil {
			continue
		}
		findings = append(findings, created...)
	}
	return findings, nil
}

func (s *Service) assessAsset(ctx context.Context, a asset.Asset, identifier string) ([]Finding, error) {
	if identifier == "" {
		f, err := s.repo.Create(ctx, Finding{
			AssetID:        a.ID,
			Title:          fmt.Sprintf("Unidentified service exposed: %s on %s", a.Service, a.Host),
			Description:    "The service responded but no reliable version/software banner was captured. Manual investigation is recommended.",
			Severity:       SeverityInformational,
			Status:         StatusDetected,
			Recommendation: "Confirm whether this service should be exposed; identify its version manually if so.",
		})
		if err != nil {
			return nil, err
		}
		return []Finding{*f}, nil
	}

	cves, err := nvd.SearchByKeyword(ctx, identifier, 5)
	if err != nil {
		return nil, err
	}

	var created []Finding
	for _, cve := range cves {
		if cve.CVSS == 0 {
			continue
		}
		cvss := cve.CVSS
		f, err := s.repo.Create(ctx, Finding{
			AssetID:        a.ID,
			Title:          fmt.Sprintf("%s — %s", cve.ID, identifier),
			Description:    cve.Description,
			Severity:       mapSeverity(cve.Severity),
			CVSS:           &cvss,
			Status:         StatusDetected,
			Recommendation: "Verify the affected version range and apply the vendor's patch or documented mitigation.",
		})
		if err != nil {
			continue
		}
		created = append(created, *f)
	}
	return created, nil
}

func (s *Service) ListByTarget(ctx context.Context, targetID string) ([]Finding, error) {
	return s.repo.ListByTarget(ctx, targetID)
}

// --- Phase 09: validation workflow ---

// StartValidation moves a finding from detected to needs_validation,
// signaling that an analyst has picked it up for manual review.
func (s *Service) StartValidation(ctx context.Context, id string) (*Finding, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if f.Status != StatusDetected {
		return nil, ErrInvalidTransition
	}
	if err := s.repo.UpdateStatus(ctx, id, StatusNeedsValidation); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

// Confirm promotes a finding to confirmed, but only from needs_validation,
// and only together with evidence — a confirmed finding without evidence
// is never allowed to exist.
func (s *Service) Confirm(ctx context.Context, id string, ev evidence.Evidence) (*Finding, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if f.Status != StatusNeedsValidation {
		return nil, ErrInvalidTransition
	}

	ev.FindingID = id
	if _, err := s.evidence.Create(ctx, ev); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateStatus(ctx, id, StatusConfirmed); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

// Reject marks a finding as false_positive, recording the reason as
// evidence for audit purposes even though nothing was confirmed.
func (s *Service) Reject(ctx context.Context, id, reason string) (*Finding, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if f.Status != StatusNeedsValidation {
		return nil, ErrInvalidTransition
	}

	if reason != "" {
		if _, err := s.evidence.Create(ctx, evidence.Evidence{FindingID: id, Notes: "Rejected: " + reason}); err != nil {
			return nil, err
		}
	}
	if err := s.repo.UpdateStatus(ctx, id, StatusFalsePositive); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func softwareIdentifier(a asset.Asset) string {
	if a.Technology != "" {
		return normalize(a.Technology)
	}
	if a.Version != "" && !strings.HasPrefix(a.Version, "HTTP ") {
		return a.Service + " " + a.Version
	}
	return ""
}

func normalize(tech string) string {
	if idx := strings.Index(tech, "("); idx != -1 {
		tech = tech[:idx]
	}
	return strings.TrimSpace(strings.ReplaceAll(tech, "/", " "))
}

func mapSeverity(nvdSeverity string) Severity {
	switch strings.ToUpper(nvdSeverity) {
	case "CRITICAL":
		return SeverityCritical
	case "HIGH":
		return SeverityHigh
	case "MEDIUM":
		return SeverityMedium
	case "LOW":
		return SeverityLow
	default:
		return SeverityInformational
	}
}
