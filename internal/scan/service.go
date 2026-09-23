package scan

import (
	"context"
	"errors"
	"net/url"

	"github.com/leomarqueseh/kestrel/internal/asset"
	"github.com/leomarqueseh/kestrel/internal/enum"
	"github.com/leomarqueseh/kestrel/internal/recon"
	"github.com/leomarqueseh/kestrel/internal/target"
)

var ErrTargetNotAuthorized = errors.New("scan: target is not authorized for testing")

type Service struct {
	scans   Repository
	assets  asset.Repository
	targets target.Repository
}

func NewService(scans Repository, assets asset.Repository, targets target.Repository) *Service {
	return &Service{scans: scans, assets: assets, targets: targets}
}

// RunRecon executes the reconnaissance module against targetID.
func (s *Service) RunRecon(ctx context.Context, targetID string) (*Scan, error) {
	return s.run(ctx, targetID, "recon", recon.Run)
}

// RunEnumeration executes the port enumeration module against targetID.
func (s *Service) RunEnumeration(ctx context.Context, targetID string) (*Scan, error) {
	return s.run(ctx, targetID, "enumeration", enum.Run)
}

// run centralizes the shared lifecycle every module follows: check
// authorization, create the scan record, execute, persist results,
// and mark the outcome. Only the module function itself differs.
func (s *Service) run(ctx context.Context, targetID, module string, execute func(context.Context, string) ([]asset.Asset, error)) (*Scan, error) {
	t, err := s.targets.GetByID(ctx, targetID)
	if err != nil {
		return nil, err
	}

	// Scope-enforcement gate: no module runs against an unauthorized target.
	if !t.Authorized {
		return nil, ErrTargetNotAuthorized
	}

	sc, err := s.scans.Create(ctx, targetID, module)
	if err != nil {
		return nil, err
	}

	if err := s.scans.MarkRunning(ctx, sc.ID); err != nil {
		return nil, err
	}

	host := hostOnly(t.Value, t.Type)

	discovered, err := execute(ctx, host)
	if err != nil {
		_ = s.scans.MarkFailed(ctx, sc.ID, err.Error())
		return nil, err
	}

	for _, a := range discovered {
		a.ScanID = sc.ID
		if _, err := s.assets.Create(ctx, a); err != nil {
			_ = s.scans.MarkFailed(ctx, sc.ID, err.Error())
			return nil, err
		}
	}

	if err := s.scans.MarkCompleted(ctx, sc.ID); err != nil {
		return nil, err
	}

	return s.scans.GetByID(ctx, sc.ID)
}

func (s *Service) ListByTarget(ctx context.Context, targetID string) ([]Scan, error) {
	return s.scans.ListByTarget(ctx, targetID)
}

// hostOnly normalizes a target's stored value into a bare hostname, since
// scan modules dial host:port directly — a URL target ("https://foo.com/x")
// needs its scheme and path stripped before it can be scanned.
func hostOnly(value string, targetType target.Type) string {
	if targetType == target.TypeURL {
		if u, err := url.Parse(value); err == nil && u.Hostname() != "" {
			return u.Hostname()
		}
	}
	return value
}
