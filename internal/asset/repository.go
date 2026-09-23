package asset

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, a Asset) (*Asset, error)
	ListByScan(ctx context.Context, scanID string) ([]Asset, error)
	ListByTarget(ctx context.Context, targetID string) ([]Asset, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, a Asset) (*Asset, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO assets (scan_id, host, port, protocol, service, version, technology)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id, created_at`,
		a.ScanID, a.Host, a.Port, a.Protocol, a.Service, a.Version, a.Technology,
	).Scan(&a.ID, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *postgresRepository) ListByScan(ctx context.Context, scanID string) ([]Asset, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, scan_id, host, port, protocol, service, version, technology, created_at
		 FROM assets WHERE scan_id = $1 ORDER BY host`,
		scanID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var a Asset
		if err := rows.Scan(&a.ID, &a.ScanID, &a.Host, &a.Port, &a.Protocol, &a.Service, &a.Version, &a.Technology, &a.CreatedAt); err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}
	return assets, rows.Err()
}

// ListByTarget aggregates assets across every scan ever run against
// targetID — this is the Attack Surface Inventory: everything discovered
// about a target, regardless of which module or execution found it.
func (r *postgresRepository) ListByTarget(ctx context.Context, targetID string) ([]Asset, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT a.id, a.scan_id, a.host, a.port, a.protocol, a.service, a.version, a.technology, a.created_at
		 FROM assets a
		 JOIN scans s ON s.id = a.scan_id
		 WHERE s.target_id = $1
		 ORDER BY a.host, a.port`,
		targetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var a Asset
		if err := rows.Scan(&a.ID, &a.ScanID, &a.Host, &a.Port, &a.Protocol, &a.Service, &a.Version, &a.Technology, &a.CreatedAt); err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}
	return assets, rows.Err()
}
