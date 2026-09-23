package evidence

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, e Evidence) (*Evidence, error)
	ListByFinding(ctx context.Context, findingID string) ([]Evidence, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, e Evidence) (*Evidence, error) {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO evidence (finding_id, request, response, notes)
		 VALUES ($1, $2, $3, $4) RETURNING id, captured_at`,
		e.FindingID, e.Request, e.Response, e.Notes,
	).Scan(&e.ID, &e.CapturedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *postgresRepository) ListByFinding(ctx context.Context, findingID string) ([]Evidence, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, finding_id, request, response, notes, captured_at
		 FROM evidence WHERE finding_id = $1 ORDER BY captured_at`,
		findingID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Evidence
	for rows.Next() {
		var e Evidence
		if err := rows.Scan(&e.ID, &e.FindingID, &e.Request, &e.Response, &e.Notes, &e.CapturedAt); err != nil {
			return nil, err
		}
		items = append(items, e)
	}
	return items, rows.Err()
}
