package report

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, projectID string, format Format) (*Report, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, projectID string, format Format) (*Report, error) {
	var rep Report
	err := r.pool.QueryRow(ctx,
		`INSERT INTO reports (project_id, format) VALUES ($1, $2)
		 RETURNING id, project_id, format, generated_at`,
		projectID, format,
	).Scan(&rep.ID, &rep.ProjectID, &rep.Format, &rep.GeneratedAt)
	if err != nil {
		return nil, err
	}
	return &rep, nil
}
