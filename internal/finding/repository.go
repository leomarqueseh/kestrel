package finding

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("finding: not found")

type Repository interface {
	Create(ctx context.Context, f Finding) (*Finding, error)
	GetByID(ctx context.Context, id string) (*Finding, error)
	ListByTarget(ctx context.Context, targetID string) ([]Finding, error)
	UpdateStatus(ctx context.Context, id string, status Status) error
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

const selectColumns = `id, asset_id, title, description, severity, cvss, cwe, status, recommendation, created_at`

func scanFinding(row pgx.Row) (*Finding, error) {
	var f Finding
	err := row.Scan(&f.ID, &f.AssetID, &f.Title, &f.Description, &f.Severity, &f.CVSS, &f.CWE, &f.Status, &f.Recommendation, &f.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *postgresRepository) Create(ctx context.Context, f Finding) (*Finding, error) {
	row := r.pool.QueryRow(ctx,
		`INSERT INTO findings (asset_id, title, description, severity, cvss, cwe, status, recommendation)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING `+selectColumns,
		f.AssetID, f.Title, f.Description, f.Severity, f.CVSS, f.CWE, f.Status, f.Recommendation,
	)
	return scanFinding(row)
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*Finding, error) {
	row := r.pool.QueryRow(ctx, `SELECT `+selectColumns+` FROM findings WHERE id = $1`, id)
	return scanFinding(row)
}

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
		f, err := scanFinding(rows)
		if err != nil {
			return nil, err
		}
		findings = append(findings, *f)
	}
	return findings, rows.Err()
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id string, status Status) error {
	tag, err := r.pool.Exec(ctx, `UPDATE findings SET status = $2 WHERE id = $1`, id, status)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
