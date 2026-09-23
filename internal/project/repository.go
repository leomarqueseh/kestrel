package project

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, name, description, ownerID string) (*Project, error)
	ListByOwner(ctx context.Context, ownerID string) ([]Project, error)
	GetByID(ctx context.Context, id string) (*Project, error)
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) Create(ctx context.Context, name, description, ownerID string) (*Project, error) {
	var p Project
	err := r.pool.QueryRow(ctx,
		`INSERT INTO projects (name, description, owner_id) VALUES ($1, $2, $3)
		 RETURNING id, name, description, owner_id, created_at`,
		name, description, ownerID,
	).Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postgresRepository) ListByOwner(ctx context.Context, ownerID string) ([]Project, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, description, owner_id, created_at FROM projects
		 WHERE owner_id = $1 ORDER BY created_at DESC`,
		ownerID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, rows.Err()
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*Project, error) {
	var p Project
	err := r.pool.QueryRow(ctx,
		`SELECT id, name, description, owner_id, created_at FROM projects WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
