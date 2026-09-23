package finding

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/leomarqueseh/kestrel/internal/asset"
	"github.com/leomarqueseh/kestrel/internal/nvd"
)

type Service struct {
	repo   Repository
	assets asset.Repository
}

func NewService(repo Repository, assets asset.Repository) *Service {
	return &Service{repo: repo, assets: assets}
}

// Assess correlates every network-service asset discovered for targetID
// against the NVD and records the results as findings. Every finding
// starts as StatusDetected — nothing here is a confirmed vulnerability.
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
			time.Sleep(6 * time.Second) // stay within NVD's public rate limit
		}

		created, err := s.assessAsset(ctx, a, identifier)
		if identifier != "" {
			requestsMade++
		}
		if err != nil {
			continue // one failed asset shouldn't abort the whole assessment
		}
		findings = append(findings, created...)
	}
	return findings, nil
}

func (s *Service) assessAsset(ctx context.Context, a asset.Asset, identifier string) ([]Finding, error) {
	// No reliable software identifier — status codes like "HTTP 200" don't
	// count. Recording an informational finding is more honest than
	// guessing, and avoids noisy, irrelevant NVD keyword matches.
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
		// Skip matches with no CVSS v3 score: usually pre-2016 CVEs that
		// keyword search over-matched on generic terms — not reliable
		// enough to surface as an actionable finding.
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

// softwareIdentifier picks the best available fingerprint for correlation:
// a real technology banner (e.g. "Apache/2.4.7 (Ubuntu)") over a generic
// protocol status like "HTTP 200", which isn't a software identifier at all.
func softwareIdentifier(a asset.Asset) string {
	if a.Technology != "" {
		return normalize(a.Technology)
	}
	if a.Version != "" && !strings.HasPrefix(a.Version, "HTTP ") {
		return a.Service + " " + a.Version
	}
	return ""
}

// normalize turns "Apache/2.4.7 (Ubuntu)" into "Apache 2.4.7" — closer to
// how NVD's keyword index tokenizes software names, and drops the OS
// suffix that adds noise without adding precision.
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
