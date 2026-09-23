// Package dashboard aggregates counts across a project — targets, scans,
// and findings by severity/status — for an at-a-glance security posture view.
package dashboard

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Summary struct {
	TotalTargets      int `json:"total_targets"`
	TotalScans        int `json:"total_scans"`
	Critical          int `json:"critical"`
	High              int `json:"high"`
	Medium            int `json:"medium"`
	Low               int `json:"low"`
	Informational     int `json:"informational"`
	ConfirmedFindings int `json:"confirmed_findings"`
	OpenFindings      int `json:"open_findings"`     // detected + needs_validation
	ResolvedFindings  int `json:"resolved_findings"` // confirmed + false_positive
}

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

func (s *Service) ForProject(ctx context.Context, projectID string) (*Summary, error) {
	var sum Summary

	if err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM targets WHERE project_id = $1`, projectID,
	).Scan(&sum.TotalTargets); err != nil {
		return nil, err
	}

	if err := s.pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM scans s JOIN targets t ON t.id = s.target_id WHERE t.project_id = $1`, projectID,
	).Scan(&sum.TotalScans); err != nil {
		return nil, err
	}

	rows, err := s.pool.Query(ctx,
		`SELECT f.severity, f.status, COUNT(*)
		 FROM findings f
		 JOIN assets a ON a.id = f.asset_id
		 JOIN scans sc ON sc.id = a.scan_id
		 JOIN targets t ON t.id = sc.target_id
		 WHERE t.project_id = $1
		 GROUP BY f.severity, f.status`,
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var severity, status string
		var count int
		if err := rows.Scan(&severity, &status, &count); err != nil {
			return nil, err
		}

		switch severity {
		case "critical":
			sum.Critical += count
		case "high":
			sum.High += count
		case "medium":
			sum.Medium += count
		case "low":
			sum.Low += count
		case "informational":
			sum.Informational += count
		}

		switch status {
		case "confirmed":
			sum.ConfirmedFindings += count
			sum.ResolvedFindings += count
		case "false_positive":
			sum.ResolvedFindings += count
		case "detected", "needs_validation":
			sum.OpenFindings += count
		}
	}

	return &sum, rows.Err()
}
