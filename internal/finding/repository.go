package finding

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, f Finding) (*Finding, error)
	ListByTarget(ctx context.Context, targetID string) ([]Finding, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, f Finding) (*Finding, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO findings (asset_id, title, description, severity, cvss, cwe, status, recommendation)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, created_at`,
		f.AssetID, f.Title, f.Description, f.Severity, f.CVSS, f.CWE, f.Status, f.Recommendation,
	).Scan(&f.ID, &f.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// ListByTarget joins findings → assets → scans to return every finding
// ever recorded for a target, regardless of which scan produced the asset.
func (r *postgresRepository) ListByTarget(ctx context.Context, targetID string) ([]Finding, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT f.id, f.asset_id, f.title, f.description, f.severity, f.cvss, f.cwe, f.status, f.recommendation, f.created_at
		 FROM findings f
		 JOIN assets a ON a.id = f.asset_id
		 JOIN scans s ON s.id = a.scan_id
		 WHERE s.target_id = $1
		 ORDER BY f.severity, f.created_at DESC`,
		targetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var findings []Finding
	for rows.Next() {
		var f Finding
		if err := rows.Scan(&f.ID, &f.AssetID, &f.Title, &f.Description, &f.Severity, &f.CVSS, &f.CWE, &f.Status, &f.Recommendation, &f.CreatedAt); err != nil {
			return nil, err
		}
		findings = append(findings, f)
	}
	return findings, rows.Err()
}
