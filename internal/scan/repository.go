package scan

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("scan: not found")

type Repository interface {
	Create(ctx context.Context, targetID, module string) (*Scan, error)
	MarkRunning(ctx context.Context, id string) error
	MarkCompleted(ctx context.Context, id string) error
	MarkFailed(ctx context.Context, id string, errMsg string) error
	ListByTarget(ctx context.Context, targetID string) ([]Scan, error)
	GetByID(ctx context.Context, id string) (*Scan, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

const selectColumns = `id, target_id, module, status, started_at, finished_at, error, created_at`

func scanRow(row pgx.Row) (*Scan, error) {
	var s Scan
	err := row.Scan(&s.ID, &s.TargetID, &s.Module, &s.Status, &s.StartedAt, &s.FinishedAt, &s.Error, &s.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *postgresRepository) Create(ctx context.Context, targetID, module string) (*Scan, error) {
	row := r.pool.QueryRow(ctx,
		`INSERT INTO scans (target_id, module, status) VALUES ($1, $2, 'pending')
		 RETURNING `+selectColumns,
		targetID, module,
	)
	return scanRow(row)
}

func (r *postgresRepository) MarkRunning(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE scans SET status = 'running', started_at = $2 WHERE id = $1`,
		id, time.Now().UTC(),
	)
	return err
}

func (r *postgresRepository) MarkCompleted(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE scans SET status = 'completed', finished_at = $2 WHERE id = $1`,
		id, time.Now().UTC(),
	)
	return err
}

func (r *postgresRepository) MarkFailed(ctx context.Context, id string, errMsg string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE scans SET status = 'failed', finished_at = $2, error = $3 WHERE id = $1`,
		id, time.Now().UTC(), errMsg,
	)
	return err
}

func (r *postgresRepository) ListByTarget(ctx context.Context, targetID string) ([]Scan, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT `+selectColumns+` FROM scans WHERE target_id = $1 ORDER BY created_at DESC`,
		targetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scans []Scan
	for rows.Next() {
		s, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		scans = append(scans, *s)
	}
	return scans, rows.Err()
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*Scan, error) {
	row := r.pool.QueryRow(ctx, `SELECT `+selectColumns+` FROM scans WHERE id = $1`, id)
	return scanRow(row)
}
