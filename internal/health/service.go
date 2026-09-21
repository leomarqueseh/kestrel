package health

import (
	"context"
	"time"
)

// Pinger is satisfied by anything that can confirm it's reachable —
// such as a *pgxpool.Pool. It lets health checks verify dependencies
// without this package knowing which database driver is behind them.
type Pinger interface {
	Ping(ctx context.Context) error
}

type Status struct {
	Status   string `json:"status"`
	Service  string `json:"service"`
	Time     string `json:"time"`
	Database string `json:"database"`
}

type Service struct {
	db Pinger
}

func NewService(db Pinger) *Service {
	return &Service{db: db}
}

func (s *Service) Check(ctx context.Context) Status {
	dbStatus := "unavailable"
	if s.db != nil {
		if err := s.db.Ping(ctx); err == nil {
			dbStatus = "ok"
		}
	}

	return Status{
		Status:   "ok",
		Service:  "kestrel-api",
		Time:     time.Now().UTC().Format(time.RFC3339),
		Database: dbStatus,
	}
}
