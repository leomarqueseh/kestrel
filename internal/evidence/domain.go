// Package evidence stores the proof attached to a finding when it moves
// from needs_validation to confirmed — request, response, and tester notes.
package evidence

import "time"

type Evidence struct {
	ID         string    `json:"id"`
	FindingID  string    `json:"finding_id"`
	Request    string    `json:"request,omitempty"`
	Response   string    `json:"response,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	CapturedAt time.Time `json:"captured_at"`
}
