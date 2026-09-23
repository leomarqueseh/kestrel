// Package scan tracks executions of a reconnaissance/enumeration module
// against a target, through a pending → running → completed/failed lifecycle.
package scan

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Scan struct {
	ID         string     `json:"id"`
	TargetID   string     `json:"target_id"`
	Module     string     `json:"module"`
	Status     Status     `json:"status"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// Error is a pointer because the "error" column is NULL when the scan
	// succeeds — a plain string can't represent "no value", only "".
	Error     *string   `json:"error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
