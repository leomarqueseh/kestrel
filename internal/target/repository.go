package target

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("target: not found")

type Repository interface {
	Create(ctx context.Context, projectID, value string, targetType Type, description string) (*Target, error)
	ListByProject(ctx context.Context, projectID string) ([]Target, error)
	GetByID(ctx context.Context, id string) (*Target, error)
	Authorize(ctx context.Context, id string) (*Target, error)
	Delete(ctx context.Context, id string) error
}

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return &postgresRepository{pool: pool}
}

const selectColumns = `id, project_id, value, target_type, authorized, description, created_at`

func scanTarget(row pgx.Row) (*Target, error) {
	var t Target
	err := row.Scan(&t.ID, &t.ProjectID, &t.Value, &t.Type, &t.Authorized, &t.Description, &t.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *postgresRepository) Create(ctx context.Context, projectID, value string, targetType Type, description string) (*Target, error) {
	// authorized is never taken from client input — it always starts false
	// and can only become true through the explicit Authorize call below.
	row := r.pool.QueryRow(ctx,
		`INSERT INTO targets (project_id, value, target_type, authorized, description)
		 VALUES ($1, $2, $3, false, $4) RETURNING `+selectColumns,
		projectID, value, targetType, description,
	)
	return scanTarget(row)
}

func (r *postgresRepository) ListByProject(ctx context.Context, projectID string) ([]Target, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT `+selectColumns+` FROM targets WHERE project_id = $1 ORDER BY created_at DESC`,
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var targets []Target
	for rows.Next() {
		t, err := scanTarget(rows)
		if err != nil {
			return nil, err
		}
		targets = append(targets, *t)
	}
	return targets, rows.Err()
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*Target, error) {
	row := r.pool.QueryRow(ctx, `SELECT `+selectColumns+` FROM targets WHERE id = $1`, id)
	return scanTarget(row)
}

func (r *postgresRepository) Authorize(ctx context.Context, id string) (*Target, error) {
	row := r.pool.QueryRow(ctx,
		`UPDATE targets SET authorized = true WHERE id = $1 RETURNING `+selectColumns,
		id,
	)
	return scanTarget(row)
}

func (r *postgresRepository) Delete(ctx context.Context, id string) error {
	tag, err := r.pool.Exec(ctx, `DELETE FROM targets WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
