// Package finding represents potential vulnerabilities correlated from
// discovered assets. Every finding starts as Detected — promotion to
// Confirmed happens only through the Phase 09 validation workflow.
package finding

import "time"

type Severity string

const (
	SeverityCritical      Severity = "critical"
	SeverityHigh          Severity = "high"
	SeverityMedium        Severity = "medium"
	SeverityLow           Severity = "low"
	SeverityInformational Severity = "informational"
)

type Status string

const (
	StatusDetected        Status = "detected"
	StatusNeedsValidation Status = "needs_validation"
	StatusConfirmed       Status = "confirmed"
	StatusFalsePositive   Status = "false_positive"
)

type Finding struct {
	ID             string    `json:"id"`
	AssetID        string    `json:"asset_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description,omitempty"`
	Severity       Severity  `json:"severity"`
	CVSS           *float64  `json:"cvss,omitempty"`
	CWE            string    `json:"cwe,omitempty"`
	Status         Status    `json:"status"`
	Recommendation string    `json:"recommendation,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
