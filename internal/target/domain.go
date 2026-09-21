// Package target manages authorized assets in scope for a project —
// the boundary that later scan phases must never cross.
package target

import "time"

type Type string

const (
	TypeDomain Type = "domain"
	TypeIP     Type = "ip"
	TypeURL    Type = "url"
	TypeCIDR   Type = "cidr"
)

func (t Type) Valid() bool {
	switch t {
	case TypeDomain, TypeIP, TypeURL, TypeCIDR:
		return true
	}
	return false
}

type Target struct {
	ID          string    `json:"id"`
	ProjectID   string    `json:"project_id"`
	Value       string    `json:"value"`
	Type        Type      `json:"target_type"`
	Authorized  bool      `json:"authorized"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
